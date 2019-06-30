package channelsfunn

import (
	"testing"
)

type dummyInsertStrategy struct{
	customChecks func(csvRecord csvDataModel)
}

func (dIS dummyInsertStrategy) InsertData(csvRecord csvDataModel) {
	dIS.customChecks(csvRecord)
}

func TestCsvProcessor_Process(t *testing.T) {
	insertStrategy := dummyInsertStrategy{
		customChecks: func(csvRecord csvDataModel) {},
	}
	processor := csvProcessor{
		insertStrategy: insertStrategy,
		csvFilePath:    "../electricity-consumption-by-sectors.csv",
	}

	tearDown := func() { insertStrategy.customChecks = func(csvRecord csvDataModel) {}}

	t.Run("it doesnt fail", func(t *testing.T) {
		defer tearDown()
		err := processor.Process()
		if err != nil {
			t.Fail()
		}
	})

	t.Run("it pass all structs properly", func(t *testing.T) {
		defer tearDown()
		err := processor.Process()
		if err != nil {
			t.Fail()
		}
		insertStrategy.customChecks = func(csvRecord csvDataModel) {
			if len(csvRecord.Consumption) == 0 {
				t.Fail()
			}
			if len(csvRecord.ConsumptionType) == 0 {
				t.Fail()
			}
			if len(csvRecord.Region) == 0 {
				t.Fail()
			}
			if len(csvRecord.Year) == 0 {
				t.Fail()
			}
		}
	})
}
