package processmetrics

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

	result.MetricGroup = constant.ProcessMetric

	defer func() {

		if criticalError := recover(); criticalError != nil {

			result.Err = fmt.Errorf("%v", criticalError)

			utility.Formatresult(result, credentials)
		}

	}()

	processOutput, execErr := utility.ExecuteCommand(credentials, constant.PROCESS)

	if execErr != nil {

		result.Output = nil

		result.Err = execErr

		utility.Formatresult(result, credentials)

		return
	}

	processOutput = strings.Trim(processOutput, constant.NewLine)

	entries := strings.Split(processOutput, constant.NewLine)

	processStatistics := make([]map[string]interface{}, len(entries))

	for index, entry := range entries {

		splitData := strings.Split(entry, constant.SemiColon)

		data := make(map[string]interface{})

		if len(splitData) >= 5 {

			data["system.process.id"], _ = strconv.Atoi(splitData[0])

			data["system.process.memory"], _ = strconv.Atoi(splitData[1])

			data["system.process.cpu.seconds"], _ = strconv.ParseFloat(splitData[2], 64)

			data["system.process.command"] = splitData[3]

			data["system.process.name"] = splitData[4]

			processStatistics[index] = data

		} else {

			result.Err = errors.New("data split error in process ")

			utility.Formatresult(result, credentials)

		}
	}

	result.Output = processStatistics

	utility.Formatresult(result, credentials)

	return
}
