package cmd

import (
	"bytes"
	"context"
	"os/exec"
	"strings"
	"time"
)

func Run(command string) ([]string, error) {
	var (
		stdout bytes.Buffer
		stderr bytes.Buffer
	)

	c := exec.Command("bash", "-c", command)
	c.Stdout = &stdout
	c.Stderr = &stderr

	if err := c.Run(); err != nil {
		return parseCmdOutput(stderr.String()), err
	}

	return parseCmdOutput(stdout.String()), nil
}

func parseCmdOutput(output string) []string {
	out := []string{}
	arr := strings.Split(output, "\n")
	for _, v := range arr {
		out = append(out, strings.TrimSpace(v))
	}
	return out
}

func RunWithTimeOut(command string, timeout time.Duration) ([]string, error) {
	var (
		stdout bytes.Buffer
		stderr bytes.Buffer
	)

	ctx, _ := context.WithTimeout(context.Background(), timeout)
	c := exec.CommandContext(ctx, "bash", "-c", command)
	c.Stdout = &stdout
	c.Stderr = &stderr

	if err := c.Run(); err != nil {
		return parseCmdOutput(stderr.String()), err
	}

	return parseCmdOutput(stdout.String()), nil
}
