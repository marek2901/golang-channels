package channelsfunn

import (
	"testing"
)

type dummyInsertStrategy struct {
	customChecks func(csvRecord csvDataModel)
}

func (dIS dummyInsertStrategy) InsertData(csvRecord csvDataModel) {
	dIS.customChecks(csvRecord)
}

func TestCsvProcessor_Process(t *testing.T) {
	createTestProcessor := func(customChecks func(csvRecord csvDataModel)) csvProcessor {
		insertStrategy := dummyInsertStrategy{
			customChecks: customChecks,
		}
		return csvProcessor{
			insertStrategy: insertStrategy,
			csvFilePath:    "../electricity-consumption-by-sectors.csv",
		}
	}

	t.Run("it doesnt fail", func(t *testing.T) {
		err := createTestProcessor(func(csvRecord csvDataModel) {}).Process()
		if err != nil {
			t.Fail()
		}
	})

	t.Run("it passes all structs properly", func(t *testing.T) {
		runCount := 0
		processor := createTestProcessor(func(csvRecord csvDataModel) {
			runCount++
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
		})
		err := processor.Process()
		if err != nil {
			t.Fail()
		}
		if runCount == 0 {
			t.Fail()
		}
	})
}
