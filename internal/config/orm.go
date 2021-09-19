package config

import "time"

// Map to table `configs`
type config struct {
	ID        int64
	Name      string
	Value     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func Orm() *config {
	return &config{}
}
