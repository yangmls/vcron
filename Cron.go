package vcron

import (
	"time"
)

type Cron struct {
	Listeners map[string]chan string
	Exps      map[string]string
	Commands  map[string]string
}

func (cron *Cron) Run() {
	for {
		go func() {
			for key, value := range cron.Listeners {
				value <- cron.Commands[key]
			}
		}()

		time.Sleep(time.Second)
	}
}

func (cron *Cron) Add(name string, exp string, command string) {
	cron.Listeners[name] = make(chan string)
	cron.Exps[name] = exp
	cron.Commands[name] = command
}

func NewCron() *Cron {
	cron := &Cron{}
	cron.Listeners = make(map[string]chan string)
	cron.Commands = make(map[string]string)
	cron.Exps = make(map[string]string)

	return cron
}
