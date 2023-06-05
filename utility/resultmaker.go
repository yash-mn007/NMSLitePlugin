package utility

import (
	"main/constant"
)

func Formatresult(result ResultStorage, credentials map[string]interface{}) {

	if result.Err == nil && result.Output != nil {

		credentials["status"] = constant.SuccessStatus

		credentials["status.code"] = constant.OK

		credentials["result"] = result.Output

	} else {

		credentials["result"] = make([]interface{}, 0)

		credentials["status"] = constant.ErrorStatus

		credentials["status.code"] = constant.InternalServerError

		credentials["message"] = result.Err.Error()
	}
}
