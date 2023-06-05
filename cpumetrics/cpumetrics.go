package cpumetrics

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

	result.MetricGroup = constant.CpuMetric

	defer func() {

		if criticalError := recover(); criticalError != nil {

			result.Err = fmt.Errorf("%v", criticalError)

			utility.Formatresult(result, credentials)

		}

	}()

	cpuOutput, execErr := utility.ExecuteCommand(credentials, constant.CPU)

	if execErr != nil {

		result.Output = nil

		result.Err = execErr

		utility.Formatresult(result, credentials)

		return
	}

	cpuOutput = strings.Trim(cpuOutput, constant.NewLine)

	entries := strings.Split(cpuOutput, constant.NewLine)

	cpuStatistics := make([]map[string]interface{}, len(entries))

	for index, entry := range entries {

		splitData := strings.Split(entry, constant.SemiColon)

		data := make(map[string]interface{})

		if len(splitData) >= 4 {

			data["system.cpu.core"] = splitData[0]

			data["system.cpu.percentage"], _ = strconv.Atoi(splitData[1])

			data["system.cpu.user.percentage"], _ = strconv.Atoi(splitData[2])

			data["system.cpu.idle.percentage"], _ = strconv.Atoi(splitData[3])

			cpuStatistics[index] = data

		} else {

			result.Err = errors.New("data split error in cpu ")

			utility.Formatresult(result, credentials)

		}
	}

	result.Output = cpuStatistics

	utility.Formatresult(result, credentials)

	return
}
