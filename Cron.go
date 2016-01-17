package vcron

import (
	"github.com/gorhill/cronexpr"
	"time"
)

type Cron struct {
	Expression *cronexpr.Expression
}

func (cron *Cron) SetExpression(value string) {
	cron.Expression = cronexpr.MustParse(value)
}

func (cron *Cron) GetNextTimer() *time.Timer {
	next := cron.Expression.Next(time.Now())
	return time.NewTimer(next.Sub(time.Now()))
}

func NewCron(value string) *Cron {
	cron := &Cron{}

	cron.SetExpression(value)

	return cron
}
