package main

import (
	"encoding/json"
	"fmt"
	"main/constant"
	"main/service"
	"os"
)

func main() {

	programArguments := os.Args[1]

	credentials := make(map[string]interface{})

	err := json.Unmarshal([]byte(programArguments), &credentials)

	credentials["result"] = make([]map[string]interface{}, 0)

	credentials["error"] = make([]map[string]interface{}, 0)

	if err != nil {

		credentials["status"] = "failed"

		credentials["status.code"] = 400

		credentials["message"] = err.Error()

		return
	}

	switch credentials["service"] {

	case "discover":

		service.Discover(credentials)

	case "collect":

		service.Collect(credentials)
	}

	result, err := json.Marshal(credentials)

	if err != nil {

		credentials["status"] = constant.ErrorMessage

		credentials["message"] = "json marshal error"

		credentials["status.code"] = constant.InternalServerError
	}

	fmt.Println(string(result))

}
