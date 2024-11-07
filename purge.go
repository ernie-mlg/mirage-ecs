package mirageecs

import (
	"fmt"

	"github.com/winebarrel/cronplan"
)

type Purge struct {
	Schedule string           `json:"schedule"`
	Request  *APIPurgeRequest `json:"request"`

	purgeParams *PurgeParams
	cron        *cronplan.Expression
}

func (p *Purge) Validate() error {
	cron, err := cronplan.Parse(p.Schedule)
	if err != nil {
		return fmt.Errorf("invalid schedule expression %s: %w", p.Schedule, err)
	}
	p.cron = cron

	if p.Request == nil {
		return fmt.Errorf("purge request is required")
	}
	purgeParams, err := p.Request.Validate()
	if err != nil {
		return fmt.Errorf("invalid purge request: %w", err)
	}
	p.purgeParams = purgeParams

	return nil
}
