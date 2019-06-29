package channelsfunn

import (
	"fmt"
	"testing"
)

type dummyInsertStrategy struct {}

func (dIS dummyInsertStrategy) InsertData(csvRecord csvDataModel) {
	fmt.Printf("%v", csvRecord)
}

func TestCsvProcessor_Process(t *testing.T) {
	processor := csvProcessor{
		insertStrategy: dummyInsertStrategy{},
		csvFilePath:  "../electricity-consumption-by-sectors.csv",
	}
	t.Run("it doesnt fail", func(t *testing.T) {
		err := processor.Process()
		if err != nil {
			t.Fail()
		}
	})
}
