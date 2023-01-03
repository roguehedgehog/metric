package http

import (
	database "github.com/roguehedgehog/metric/infra"
)

type Health struct {
	App bool
	Db  bool
}

func HealthCheck() *Health {
	return &Health{
		App: true,
		Db:  database.Healthy(),
	}
}

func (h *Health) isHealthy() bool {
	return h.App && h.Db
}
