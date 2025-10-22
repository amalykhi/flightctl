//go:build linux

package executer

import (
	"bytes"
	"context"
	"errors"
	"os/exec"
	"syscall"
)

func (e *CommonExecuter) CommandContext(ctx context.Context, command string, args ...string) *exec.Cmd {
	cmd := exec.CommandContext(ctx, command, args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{Pdeathsig: syscall.SIGTERM}
	return cmd
}

func (e *CommonExecuter) execute(ctx context.Context, cmd *exec.Cmd) (stdout string, stderr string, exitCode int) {
	var stdoutBytes, stderrBytes bytes.Buffer
	cmd.Stdout = &stdoutBytes
	cmd.Stderr = &stderrBytes
	// Setpgid ensures that all child processes are in the same process group.
	// Pdeathsig ensures cleanup if the parent Go process dies unexpectedly.
	// ref. https://pkg.go.dev/syscall#SysProcAttr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid:   true,
		Pdeathsig: syscall.SIGTERM,
	}
	// if context is canceled, kill the entire process group
	cmd.Cancel = func() error {
		if cmd.Process != nil {
			// negative PID kills the process group
			// ref. https://man7.org/linux/man-pages/man2/kill.2.html
			return syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
		}
		return nil
	}

	if err := cmd.Run(); err != nil {
		// handle timeout error
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			return stdoutBytes.String(), context.DeadlineExceeded.Error(), 124
		}
		return stdoutBytes.String(), getErrorStr(err, &stderrBytes), getExitCode(err)
	}

	return stdoutBytes.String(), stderrBytes.String(), 0
}
