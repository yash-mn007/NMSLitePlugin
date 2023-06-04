package metrics

import (
	"errors"
	"fmt"
	"main/constant"
	"main/utility"
	"main/winrmclient"
	"strconv"
	"strings"
)

func Memory(credentials map[string]interface{}, channel chan utility.ResultStorage) {

	var result utility.ResultStorage

	result.MetricGroup = "memory"

	defer func() {

		if criticalError := recover(); criticalError != nil {

			result.Err = fmt.Errorf("%v", criticalError)

			channel <- result
		}

	}()

	memoryOutput, execErr := winrmclient.ExecuteCommand(credentials, constant.MEMORY)

	if execErr != nil {

		result.Output = nil

		result.Err = execErr

		channel <- result

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

		result.Output = memoryStatistics

		channel <- result

	} else {

		result.Err = errors.New("data split error in memory ")

		channel <- result
	}

	return

}
func SystemInfo(credentials map[string]interface{}, channel chan utility.ResultStorage) {

	var result utility.ResultStorage

	result.MetricGroup = "systeminfo"

	defer func() {

		if criticalError := recover(); criticalError != nil {

			result.Err = fmt.Errorf("%v", criticalError)

			channel <- result
		}

	}()

	systemOutput, execErr := winrmclient.ExecuteCommand(credentials, constant.SYSTEMINFO)

	systemOutput = strings.Trim(systemOutput, "\r\n")

	if execErr != nil {

		result.Output = nil

		result.Err = execErr

		channel <- result

		return

	}

	data := strings.Split(systemOutput, constant.SemiColon)

	systemStatistics := make(map[string]interface{}, len(data))

	if len(data) >= 7 {

		systemStatistics["system.name"] = data[0]

		systemStatistics["system.os.name"] = data[1]

		systemStatistics["system.os.version"] = data[2]

		systemStatistics["system.running.processes"], _ = strconv.Atoi(data[3])

		systemStatistics["system.threads"], _ = strconv.Atoi(data[4])

		systemStatistics["system.context.switches"], _ = strconv.Atoi(data[5])

		systemStatistics["system.uptime"] = data[6]

		result.Output = systemStatistics

		channel <- result

	} else {

		result.Err = errors.New("data split error in memory ")

		channel <- result
	}

	return

}
