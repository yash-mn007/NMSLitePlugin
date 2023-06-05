package winrm

import (
	"errors"
	"fmt"
	"main/constant"
	"main/utility"
	"strings"
)

func Discover(credentials map[string]interface{}) {

	var result utility.ResultStorage

	result.MetricGroup = constant.Discover

	hostName := make(map[string]interface{}, 1)

	defer func() {

		if criticalError := recover(); criticalError != nil {

			result.Err = fmt.Errorf("%v", criticalError)

			utility.Formatresult(result, credentials)
		}

	}()

	commandOutput, execErr := utility.ExecuteCommand(credentials, constant.HOSTNAME)

	if execErr != nil {

		result.Output = nil

		result.Err = execErr

		utility.Formatresult(result, credentials)

		return

	}

	commandOutput = strings.Trim(commandOutput, "\r\n")

	if len(commandOutput) >= 1 {

		hostName["hostname"] = commandOutput

	} else {

		result.Err = errors.New("data split error in discovery ")

		utility.Formatresult(result, credentials)

	}

	result.Output = hostName

	utility.Formatresult(result, credentials)

	return

}
