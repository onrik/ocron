package cron

import (
	"bytes"
	"time"

	"github.com/robfig/cron/v3"
)

type Cron struct {
	c *cron.Cron
}

func New(tz string) (*Cron, error) {
	options := []cron.Option{}
	if tz != "" {
		location, err := time.LoadLocation(tz)
		if err != nil {
			return nil, err
		}
		options = append(options, cron.WithLocation(location))
	}
	return &Cron{
		c: cron.New(options...),
	}, nil
}

func (c *Cron) AddTask(name, spec string, script, onSuccess, onError, finally []string, env map[string]string) (Task, error) {
	task := Task{
		name:      name,
		env:       env,
		script:    script,
		onSuccess: onSuccess,
		onError:   onError,
		finally:   finally,
		logger:    bytes.Buffer{},
	}

	_, err := c.c.AddJob(spec, &task)
	return task, err
}

func (c *Cron) Start() {
	c.c.Start()

	<-make(chan struct{})
}
