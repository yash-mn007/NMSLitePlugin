package main

import (
	"encoding/json"
	"fmt"
	"main/constant"
	"main/cpumetrics"
	"main/diskmetrics"
	"main/memorymetrics"
	"main/processmetrics"
	"main/systemmetrics"
	"main/winrm"
	"os"
)

func main() {

	programArguments := os.Args[1]

	credentials := make(map[string]interface{})

	err := json.Unmarshal([]byte(programArguments), &credentials)

	credentials[constant.Result] = make([]map[string]interface{}, 0)

	//credentials[constant.Error] = make([]map[string]interface{}, 0)

	if err != nil {

		credentials[constant.Status] = constant.Error

		credentials[constant.StatusCode] = constant.InternalServerError

		credentials[constant.StatusMessage] = err.Error()

		return
	}

	switch credentials[constant.Service] {

	case constant.Discover:

		winrm.Discover(credentials)

	case constant.Collect:

		switch credentials[constant.MetricGroup] {

		case constant.MemoryMetric:

			memorymetrics.Collect(credentials)

		case constant.SystemInfoMetric:

			systemmetrics.Collect(credentials)

		case constant.CpuMetric:

			cpumetrics.Collect(credentials)

		case constant.DiskMetric:

			diskmetrics.Collect(credentials)

		case constant.ProcessMetric:

			processmetrics.Collect(credentials)
		}
	}

	result, err := json.Marshal(credentials)

	if err != nil {

		credentials[constant.Status] = constant.ErrorStatus

		credentials[constant.StatusMessage] = "json marshal error"

		credentials[constant.StatusCode] = constant.InternalServerError
	}

	fmt.Println(string(result))

}
