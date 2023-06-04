package service

import (
	"main/constant"
	"main/metrics"
	"main/utility"
)

func Collect(credentials map[string]interface{}) {

	channel := make(chan utility.ResultStorage, constant.MetricGroupCount)

	go metrics.Memory(credentials, channel)

	go metrics.SystemInfo(credentials, channel)

	go metrics.Cpu(credentials, channel)

	go metrics.Disk(credentials, channel)

	go metrics.Process(credentials, channel)

	metricCounter := 0

	data := make(map[string]interface{})

	errorData := make(map[string]interface{})

	for commandResult := range channel {

		if commandResult.Err == nil || commandResult.Output != nil {

			credentials["status"] = constant.SuccessMessage

			credentials["status.code"] = constant.OK

			data[commandResult.MetricGroup] = commandResult.Output

			metricCounter++

		} else {

			data[commandResult.MetricGroup] = nil

			errorData["status"] = constant.ErrorMessage

			errorData["status.code"] = constant.InternalServerError

			errorData["message"] = commandResult.Err.Error()

			metricCounter++

		}

		if metricCounter == constant.MetricGroupCount {

			close(channel)
		}

	}

	credentials["result"] = data

	credentials["error"] = errorData

	if len(errorData) > 0 && credentials["status.code"] == constant.OK {

		credentials["status"] = constant.PartialMessage

		credentials["status.code"] = constant.PartialContent

		credentials["message"] = errorData["message"]

	} else if credentials["status.code"] == nil {

		credentials["status"] = constant.ErrorMessage

		credentials["status.code"] = constant.InternalServerError

		credentials["message"] = errorData["message"]

	}

}
