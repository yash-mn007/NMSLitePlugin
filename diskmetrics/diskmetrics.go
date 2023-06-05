package diskmetrics

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

	result.MetricGroup = constant.DiskMetric

	defer func() {

		if criticalError := recover(); criticalError != nil {

			result.Err = fmt.Errorf("%v", criticalError)

			utility.Formatresult(result, credentials)
		}

	}()

	diskOutput, execErr := utility.ExecuteCommand(credentials, constant.DISK)

	if execErr != nil {

		result.Output = nil

		result.Err = execErr

		utility.Formatresult(result, credentials)

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

			utility.Formatresult(result, credentials)

		}
	}

	result.Output = diskStatistics

	utility.Formatresult(result, credentials)

	return

}
