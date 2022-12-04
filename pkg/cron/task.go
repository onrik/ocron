package cron

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

type Task struct {
	name      string
	env       map[string]string
	script    []string
	onSuccess []string
	onError   []string
	finally   []string
	logger    bytes.Buffer
}

func (t *Task) log(b []byte) {
	if len(b) > 0 {
		t.logger.Write(b)
		t.logger.WriteString("\n")
	}
}

func (t *Task) Run() {
	t.logger.Reset()
	t.log([]byte(fmt.Sprintf("=== %s %s ===", time.Now().Format(time.RFC3339), t.name)))
	dir, err := os.MkdirTemp("", "")
	if err != nil {
		log.Printf("Create temp dir error: %s", err)
		return
	}
	defer os.RemoveAll(dir)
	defer t.execFinally(dir)

	for _, s := range t.script {
		err := t.execOne(s, dir)
		if err != nil {
			t.execOnError(dir)
			return
		}
	}

	t.execOnSuccess(dir)
}

func (t *Task) execOnSuccess(dir string) {
	t.exec(t.onSuccess, dir)
}

func (t *Task) execOnError(dir string) {
	t.exec(t.onError, dir)
}

func (t *Task) execFinally(dir string) {
	t.exec(t.finally, dir)

	fmt.Println(strings.TrimSpace(t.logger.String()))
}

func (t *Task) execOne(command, dir string) error {
	command, env := fillEnv(command, t.env)
	t.log([]byte(fmt.Sprintf("$ %s", command)))
	buff := &bytes.Buffer{}
	cmd := exec.Command("sh", "-c", command)
	cmd.Dir = dir
	cmd.Env = env
	cmd.Stdout = buff
	cmd.Stderr = buff
	err := cmd.Run()
	t.log(
		bytes.ReplaceAll(buff.Bytes(), []byte("bash: line 1: "), []byte{}),
	)
	return err
}

func (t *Task) exec(ss []string, dir string) {
	for _, s := range ss {
		t.execOne(s, dir)
	}
}

func fillEnv(s string, env map[string]string) (string, []string) {
	envList := []string{}
	for k, v := range env {
		s = strings.ReplaceAll(s, fmt.Sprintf("${%s}", k), v)
		envList = append(envList, fmt.Sprintf("%s=%s", k, v))
	}

	return s, envList
}
