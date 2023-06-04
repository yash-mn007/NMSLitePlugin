package service

import (
	"main/constant"
	"main/metrics"
	"main/utility"
)

func Discover(credentials map[string]interface{}) {

	channel := make(chan utility.ResultStorage)

	go metrics.Memory(credentials, channel)

	commandResult := <-channel

	if commandResult.Err == nil || commandResult.Output != nil {

		credentials["status"] = constant.SuccessMessage

		credentials["status.code"] = constant.OK

	} else {

		credentials["status"] = constant.ErrorMessage

		credentials["status.code"] = constant.InternalServerError

		credentials["message"] = commandResult.Err.Error()

	}

}
