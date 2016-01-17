package vcron

import (
	"github.com/yangmls/vcron/cronexpr"
	"time"
)

type Cron struct {
	Expression *cronexpr.Expression
}

func (cron *Cron) SetExpression(value string) {
	cron.Expression = cronexpr.MustParse(value)
}

func (cron *Cron) GetNextTimer() *time.Timer {
	now := time.Now()
	next := cron.Expression.Next(now)
	return time.NewTimer(next.Sub(now))
}

func NewCron(value string) *Cron {
	cron := &Cron{}

	cron.SetExpression(value)

	return cron
}
