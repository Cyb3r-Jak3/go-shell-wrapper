package goshellwrapper

import (
	"context"
	"os/exec"
	"runtime"

	"github.com/go-cmd/cmd"
)

func RunCommand(ctx context.Context, command string) ([]string, []string, error) {
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
	status := <-cmd.Start() // run blocking
	stdout := status.Stdout
	stderr := status.Stderr
	return stdout, stderr, status.Error
}
