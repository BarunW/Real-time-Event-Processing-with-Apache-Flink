package configs

import(
    "time"
)

type configs struct{}

func NewConfigs() configs {
	return configs{}
}

const (
    Event_NextEventTimeRange    int             = 30
    Event_NextEventTimeUnit     time.Duration   = time.Second
    Event_IntialEventTime       time.Duration   = 13
    Event_InitalEventTimeUnit   time.Duration   = time.Millisecond    
)

const (
    Faker_ProdIdRange           int = 1000
    // This duration is in second
    Faker_SessionDurationRange  int = 1800
    Faker_UserIdRange           int = 1000
)
    


