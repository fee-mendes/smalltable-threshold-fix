// Copyright (C) 2017 ScyllaDB

package sched

import (
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx"
	"github.com/scylladb/mermaid/sched/runner"
	"github.com/scylladb/mermaid/uuid"
)

// TaskType specifies the type of a Task.
type TaskType string

// TaskType enumeration
const (
	UnknownTask            TaskType = "unknown"
	BackupTask             TaskType = "backup"
	RepairAutoScheduleTask TaskType = "repair_auto_schedule"
	RepairTask             TaskType = "repair"

	mockTask TaskType = "mock"
)

func (t TaskType) String() string {
	return string(t)
}

// MarshalText implements encoding.TextMarshaler.
func (t TaskType) MarshalText() (text []byte, err error) {
	return []byte(t.String()), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (t *TaskType) UnmarshalText(text []byte) error {
	switch TaskType(text) {
	case UnknownTask:
		*t = UnknownTask
	case BackupTask:
		*t = BackupTask
	case RepairAutoScheduleTask:
		*t = RepairAutoScheduleTask
	case RepairTask:
		*t = RepairTask
	case mockTask:
		*t = mockTask
	default:
		return fmt.Errorf("unrecognized TaskType %q", text)
	}
	return nil
}

// Schedule defines a periodic schedule.
type Schedule struct {
	Repeat       bool      `json:"repeat"`
	StartDate    time.Time `json:"start_date"`
	IntervalDays int       `json:"interval_days"`
	NumRetries   int       `json:"num_retries"`
}

// MarshalUDT implements UDTMarshaler.
func (s Schedule) MarshalUDT(name string, info gocql.TypeInfo) ([]byte, error) {
	f := gocqlx.DefaultMapper.FieldByName(reflect.ValueOf(s), name)
	return gocql.Marshal(info, f.Interface())
}

// UnmarshalUDT implements UDTUnmarshaler.
func (s *Schedule) UnmarshalUDT(name string, info gocql.TypeInfo, data []byte) error {
	f := gocqlx.DefaultMapper.FieldByName(reflect.ValueOf(s), name)
	return gocql.Unmarshal(info, data, f.Addr().Interface())
}

func (s *Schedule) nextActivation(now time.Time, runs []*Run) time.Time {
	n := len(runs)
	if n == 0 && s.StartDate.After(now.Add(taskStartNowSlack)) {
		return s.StartDate
	}
	lastStart := s.StartDate
	if n > 0 {
		lastStart = runs[n-1].StartTime
	}

	if s.NumRetries > 0 {
		// check no more than NumRetries Runs were attempted
		retries := 0
		for i := n - 1; i >= 0; i-- {
			retries++
			if runs[i].Status == runner.StatusStopped {
				break
			}
		}
		if retries < s.NumRetries {
			t := lastStart.Add(retryTaskWait)
			if t.Before(now) {
				// previous activation was is in the past, and didn't occur. Try again now.
				return now.Add(taskStartNowSlack)
			}
			return t
		}
	}

	// advance start date with idealized periods
	if s.IntervalDays > 0 {
		for lastStart.Before(now) {
			lastStart = lastStart.AddDate(0, 0, s.IntervalDays)
		}
	}
	return lastStart
}

// Task is a schedulable entity.
type Task struct {
	ClusterID uuid.UUID `json:"cluster_id"`
	Type      TaskType  `json:"type"`
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Tags      []string  `json:"tags"`
	Metadata  string    `json:"metadata"`
	Enabled   bool      `json:"enabled"`
	Sched     Schedule  `json:"schedule"`

	Properties map[string]string `json:"properties"`
}

// Validate checks if all the required fields are properly set.
func (t *Task) Validate() error {
	if t == nil {
		return errors.New("nil task")
	}
	if t.ID == uuid.Nil {
		return errors.New("missing ID")
	}
	if t.ClusterID == uuid.Nil {
		return errors.New("missing ClusterID")
	}

	switch t.Type {
	case "", UnknownTask:
		return errors.New("no TaskType specified")

	default:
		var tmpType TaskType
		if err := tmpType.UnmarshalText([]byte(t.Type)); err != nil {
			return err
		}
	}

	return nil
}

// Run describes a running instance of a Task.
type Run struct {
	ID        uuid.UUID     `json:"id"`
	Type      TaskType      `json:"type"`
	ClusterID uuid.UUID     `json:"cluster_id"`
	TaskID    uuid.UUID     `json:"task_id"`
	Status    runner.Status `json:"status"`
	Cause     string        `json:"cause"`
	Owner     string        `json:"owner"`
	StartTime time.Time     `json:"start_time"`
	EndTime   time.Time     `json:"end_time"`
}
