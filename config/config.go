package config

import (
	"time"
)

// PORT specifies the port
const PORT string = "8000"

// GlobalExpireDuration specifies how long the server holds the each file
const GlobalExpireDuration time.Duration = 24 * time.Hour

// GlobalCheckRate how often does the system check the global timers
const GlobalCheckRate time.Duration = 5 * time.Second
