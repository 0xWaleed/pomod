package pomo

import (
	"fmt"
	"time"
)

type SessionType byte

const (
	SessionTypeWork SessionType = iota
	SessionTypeShortBreak
	SessionTypeLongBreak
)

func (s SessionType) IsWork() bool {
	return s == SessionTypeWork
}

func (s SessionType) IsShortBreak() bool {
	return s == SessionTypeShortBreak
}

func (s SessionType) IsLongBreak() bool {
	return s == SessionTypeLongBreak
}

func (s SessionType) String() string {
	if s.IsWork() {
		return "Work"
	}

	if s.IsShortBreak() {
		return "Short Break"
	}

	if s.IsLongBreak() {
		return "Long Break"
	}

	return ""
}

type Session struct {
	Elapsed time.Duration
	Task    *Task
	Type    SessionType
}

func (s *Session) String() string {
	return fmt.Sprintf("(%s:%s:%s)[%s]", s.Type, s.Elapsed, s.Remaining(), s.Task)
}

func (s *Session) Length() time.Duration {
	switch s.Type {
	case SessionTypeWork:
		return s.Task.Options.WorkLength
	case SessionTypeShortBreak:
		return s.Task.Options.ShortBreakLength
	case SessionTypeLongBreak:
		return s.Task.Options.LongBreakLength
	default:
		panic("invalid type")
	}
}

func (s *Session) Remaining() time.Duration {
	return s.Length() - s.Elapsed
}

func (s *Session) Reset() {
	s.Elapsed = 0
}

func (s *Session) Advance() {
	s.Elapsed += time.Second
}

func (s *Session) Done() bool {
	return s.Elapsed >= s.Length()
}

func NewSession(task *Task, sessionType SessionType) *Session {
	s := Session{
		Task: task,
		Type: sessionType,
	}

	return &s
}
