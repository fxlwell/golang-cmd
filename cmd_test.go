package cmd

import (
	"testing"
	"time"
)

func TestRunCmd(t *testing.T) {
	cmds := map[string]bool{
		"ls":                  true,
		"lss":                 false,
		"ls | grep cmd":       true,
		"ping www.baidu.comx": false,
	}

	for cmd, result := range cmds {
		out1, err1 := Run(cmd)
		out2, err2 := RunWithTimeOut(cmd, time.Second)
		ok1 := (err1 == nil)
		ok2 := (err2 == nil)

		if ok1 != ok2 {
			t.Fatalf("cmd '%s' run error, ok1:%T, ok2:%T, out1:%v, out2:%v", cmd, ok1, ok2, out1, out2)
		}

		if ok1 != result {
			t.Fatalf("cmd '%s' run error, ok1:%T, ok2:%T, out1:%v, out2:%v", cmd, ok1, ok2, out1, out2)
		}

		if cmd == "ls" {
			if out1[0] != "LICENSE" || out2[0] != "LICENSE" {
				t.Fatalf("cmd '%s' run error", cmd)
			}
		}

		if cmd == "ls | grep cmd" {
			if out1[0] != "cmd.go" || out2[0] != "cmd.go" {
				t.Fatalf("cmd '%s' run error", cmd)
			}
		}
	}
}

func TestRunWithTimeOut(t *testing.T) {
	cmd := "ping -c 2 -i 2 192.168.1.1"
	out, err := RunWithTimeOut(cmd, time.Second)
	if err == nil {
		t.Fatal("timeout cmd run error", out)
	}

	cmd = "ping -c 2 -i 2 192.168.1.1"
	out, err = RunWithTimeOut(cmd, time.Second*10)
	if err != nil {
		t.Fatal("timeout cmd run error", err, out)
	}
}
