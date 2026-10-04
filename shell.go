package goshellwrapper

import (
	"context"
	"os/exec"
	"runtime"

	"github.com/go-cmd/cmd"
)

func RunCommand(ctx context.Context, command string) cmd.Status {
	var shell string
	if runtime.GOOS == "windows" {
		powershellCheck := exec.Command("powershell", "-Command", "Write-Output 'PowerShell Available'")
		if powershellErr := powershellCheck.Run(); powershellErr == nil {
			shell = "powershell"
		} else {
			shell = "cmd"
		}
	} else {
		shell = "bash"
	}
	cmd := cmd.NewCmd(shell, command)
	return <-cmd.Start()
}
