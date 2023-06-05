package utility

import (
	"context"
	"fmt"
	"github.com/masterzen/winrm"
	"strconv"
	"strings"
)

func ExecuteCommand(credentials map[string]interface{}, command string) (commandOutput string, err error) {

	ip, _ := credentials["ip"].(string)

	username, _ := credentials["username"].(string)

	password, _ := credentials["password"].(string)

	port, _ := strconv.Atoi(credentials["port"].(string))

	endpoint := winrm.NewEndpoint(ip, port, false, false, nil, nil, nil, 0)

	client, err := winrm.NewClient(endpoint, username, password)

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	commandOutput, _, _, err = client.RunPSWithContextWithString(ctx, command, "")

	err = fmt.Errorf("%v", err)

	if strings.EqualFold(err.Error(), "<nil>") {

		err = nil
	}

	return
}
