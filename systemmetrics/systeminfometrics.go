package systemmetrics

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

	result.MetricGroup = constant.SystemInfoMetric

	defer func() {

		if criticalError := recover(); criticalError != nil {

			result.Err = fmt.Errorf("%v", criticalError)

			utility.Formatresult(result, credentials)

		}

	}()

	systemOutput, execErr := utility.ExecuteCommand(credentials, constant.SYSTEMINFO)

	systemOutput = strings.Trim(systemOutput, "\r\n")

	if execErr != nil {

		result.Output = nil

		result.Err = execErr

		utility.Formatresult(result, credentials)

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

	} else {

		result.Err = errors.New("data split error in memory ")

		utility.Formatresult(result, credentials)

	}

	result.Output = systemStatistics

	utility.Formatresult(result, credentials)

	return

}
