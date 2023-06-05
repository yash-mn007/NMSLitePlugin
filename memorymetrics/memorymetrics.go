package memorymetrics

import (
	"errors"
	"fmt"
	"main/constant"
	"main/utility"
	"strconv"
	"strings"
)

func Collect(credentials map[string]interface{}) {

	var result utility.ResultStorage

	result.MetricGroup = constant.MemoryMetric

	defer func() {

		if criticalError := recover(); criticalError != nil {

			result.Err = fmt.Errorf("%v", criticalError)

			utility.Formatresult(result, credentials)
		}

	}()

	memoryOutput, execErr := utility.ExecuteCommand(credentials, constant.MEMORY)

	if execErr != nil {

		result.Output = nil

		result.Err = execErr

		utility.Formatresult(result, credentials)

		return
	}

	memoryOutput = strings.Trim(memoryOutput, "\r\n")

	data := strings.Split(memoryOutput, constant.SemiColon)

	memoryStatistics := make(map[string]interface{}, len(data))

	if len(data) >= 6 {

		memoryStatistics["system.memory.installed"] = data[0]

		memoryStatistics["system.memory.free.bytes"] = data[1]

		memoryStatistics["system.memory.used.bytes"] = data[2]

		memoryStatistics["system.memory.free.percentage"], _ = strconv.ParseFloat(data[3], 64)

		memoryStatistics["system.memory.used.percentage"], _ = strconv.ParseFloat(data[4], 64)

		memoryStatistics["system.memory.swap"] = data[5]

	} else {

		result.Err = errors.New("data split error in memory ")

		utility.Formatresult(result, credentials)

	}

	result.Output = memoryStatistics

	utility.Formatresult(result, credentials)

	return

}
