package main

import (
	"time"

	"github.com/google/uuid"
)

type State int

const (
	Pending State = iota
	Scheduled
	Running
	Completed
	Failed
)

type RestartPolicy string

const (
	UnlessStopped RestartPolicy = "unless-stopped"
	Always        RestartPolicy = "always"
	OnFailure     RestartPolicy = "on-failure"
)

type Task struct {
	ID            uuid.UUID
	Name          string
	State         State
	Image         string
	Cpu           int
	Memory        int
	Disk          int
	Ports         map[uint16]uint16
	RestartPolicy RestartPolicy
	StartTime     time.Time
	FinishTime    time.Time
}

type TaskEvent struct {
	ID        uuid.UUID
	State     State
	Timestamp time.Time
	Task      Task
}
