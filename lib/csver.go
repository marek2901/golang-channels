package channelsfunn

import (
	"bufio"
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

var CsvErDebug bool

type csvDataModel struct {
	Year            string
	Region          string
	ConsumptionType string
	Consumption     string
}

type sqliteInsertStrategy struct {
	db        *sql.DB
	statement *sql.Stmt
}

func (is sqliteInsertStrategy) InsertData(csvRecord csvDataModel) {
	_, err := is.statement.Exec(csvRecord.Year, csvRecord.Region, csvRecord.ConsumptionType, csvRecord.Consumption)
	if err != nil {
		fmt.Printf("Warning could not save %s", err)
	}
}

type InsertStrategy interface {
	InsertData(csvRecord csvDataModel)
}

type csvProcessor struct {
	insertStrategy InsertStrategy
	csvFilePath    string
	wg             *sync.WaitGroup
}

func (cPrc csvProcessor) loadCsv(queue chan<- csvDataModel) (err error) {
	csvFile, err := os.Open(cPrc.csvFilePath)
	if err != nil {
		close(queue)
		return
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))
	firstIteration := true
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		// Skip the csv headers
		if !firstIteration {
			queue <- csvDataModel{
				Year:            line[0],
				Region:          line[1],
				ConsumptionType: line[2],
				Consumption:     line[3],
			}
		}
		if CsvErDebug {
			fmt.Print("\nAdded record")
		}
		firstIteration = false
	}
	close(queue)
	return nil
}

func (cPrc csvProcessor) worker(queue <-chan csvDataModel, workerId int) {
	defer cPrc.wg.Done()
	for csvRecord := range queue {
		if CsvErDebug {
			fmt.Printf("\nWorker no %d processing", workerId)
		}
		cPrc.insertStrategy.InsertData(csvRecord)
	}
}

func (cPrc csvProcessor) Process() (err error) {
	csvQueue := make(chan csvDataModel)

	go func() {
		err = cPrc.loadCsv(csvQueue)
		if err != nil {
			return
		}
	}()

	cPrc.wg = &sync.WaitGroup{}
	for w := 1; w <= 20; w++ {
		cPrc.wg.Add(1)
		go cPrc.worker(csvQueue, w)
	}

	cPrc.wg.Wait()
	return
}

type CsvProcessor interface {
	Process() error
}

func NewSqliteStrategy(sqliteFilePath string) *sqliteInsertStrategy {
	db, err := sql.Open("sqlite3", sqliteFilePath)
	if err != nil {
		panic(err)
	}
	stmt, err := db.Prepare("INSERT INTO CSV_RECORDS(YEAR, REGION, CONSUMPTION_TYPE, CONSUMPTION) values(?,?,?,?)")
	if err != nil {
		panic(err)
	}
	return &sqliteInsertStrategy{db: db, statement: stmt}
}

func GetProcessor(csvFilePath string, sqliteFilePath string) CsvProcessor {
	CsvErDebug = os.Getenv("CSV_ER_DEBUG") == "true"
	return csvProcessor{
		insertStrategy: NewSqliteStrategy(sqliteFilePath),
		csvFilePath:    csvFilePath,
	}
}
