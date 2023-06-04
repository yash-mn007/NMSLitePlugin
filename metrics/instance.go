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

func Cpu(credentials map[string]interface{}, channel chan utility.ResultStorage) {

	var result utility.ResultStorage

	result.MetricGroup = "cpu"

	defer func() {

		if criticalError := recover(); criticalError != nil {

			result.Err = fmt.Errorf("%v", criticalError)

			channel <- result
		}

	}()

	cpuOutput, execErr := winrmclient.ExecuteCommand(credentials, constant.CPU)

	if execErr != nil {

		result.Output = nil

		result.Err = execErr

		channel <- result

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

		}
	}

	result.Output = cpuStatistics

	channel <- result

	return

}

func Process(credentials map[string]interface{}, channel chan utility.ResultStorage) {

	var result utility.ResultStorage

	result.MetricGroup = "process"

	defer func() {

		if criticalError := recover(); criticalError != nil {

			result.Err = fmt.Errorf("%v", criticalError)

			channel <- result
		}

	}()

	processOutput, execErr := winrmclient.ExecuteCommand(credentials, constant.PROCESS)

	if execErr != nil {

		result.Output = nil

		result.Err = execErr

		channel <- result

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
		}
	}

	result.Output = processStatistics

	channel <- result

	return
}

func Disk(credentials map[string]interface{}, channel chan utility.ResultStorage) {

	var result utility.ResultStorage

	result.MetricGroup = "disk"

	defer func() {

		if criticalError := recover(); criticalError != nil {

			result.Err = fmt.Errorf("%v", criticalError)

			channel <- result
		}

	}()

	diskOutput, execErr := winrmclient.ExecuteCommand(credentials, constant.DISK)

	if execErr != nil {

		result.Output = nil

		result.Err = execErr

		channel <- result

		return
	}

	diskOutput = strings.Trim(diskOutput, constant.NewLine)

	entries := strings.Split(diskOutput, constant.NewLine)

	diskStatistics := make([]map[string]interface{}, len(entries))

	for index, entry := range entries {

		splitData := strings.Split(entry, constant.SemiColon)

		data := make(map[string]interface{})

		if len(splitData) >= 7 {

			data["system.disk.name"] = splitData[0]

			data["system.disk.total.bytes"] = splitData[1]

			data["system.disk.free.bytes"] = splitData[2]

			data["system.disk.used.bytes"] = splitData[3]

			data["system.disk.free.percentage"], _ = strconv.ParseFloat(splitData[4], 64)

			data["system.disk.used.percentage"], _ = strconv.ParseFloat(splitData[5], 64)

			data["system.disk.volume.name"] = splitData[6]

			diskStatistics[index] = data

		} else {

			result.Err = errors.New("data split error in disk ")
		}
	}

	result.Output = diskStatistics

	channel <- result

	return

}
