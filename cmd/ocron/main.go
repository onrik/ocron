package main

import (
	"flag"
	"log"

	"github.com/onrik/ocron/pkg/cron"
)

func main() {
	configFile := flag.String("config", "/etc/ocron/config.yml", "")
	run := flag.String("run", "", "task to run")
	flag.Parse()

	config, err := readConfig(*configFile)
	if err != nil {
		log.Println(err)
		return
	}
	c, err := cron.New(config.TZ)
	if err != nil {
		log.Println(err)
		return
	}

	for _, t := range config.Tasks {
		log.Printf("%s spec=%s", t.Name, t.Spec)
		task, err := c.AddTask(
			t.Name,
			t.Spec,
			t.Script,
			t.OnSuccess,
			t.OnError,
			t.Finally,
			mergeMap(t.Env, config.Env),
		)
		if err != nil {
			log.Println(err)
			return
		}
		if run != nil && *run == t.Name {
			task.Run()
			return
		}

	}

	if run != nil && *run != "" {
		log.Printf("%s - no task found", *run)
		return
	}

	c.Start()
}

func mergeMap(m1, m2 map[string]string) map[string]string {
	m := map[string]string{}
	for k, v := range m1 {
		m[k] = v
	}

	for k, v := range m2 {
		m[k] = v
	}

	return m
}
