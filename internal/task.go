package pomo

import (
	"fmt"
	"time"
)

const DefaultLongBreakAfter = 3

type TaskOptions struct {
	WorkLength       time.Duration
	LongBreakLength  time.Duration
	ShortBreakLength time.Duration
	AutoSwitch       bool
	Interval         time.Duration
	LongBreakAfter   int
}

type Task struct {
	CurrentSession *Session
	Options        TaskOptions
	Title          string
	PomodoroId     int
}

func (s *Task) String() string {
	return fmt.Sprintf("(%d:%s)", s.PomodoroId, s.Title)
}

func (s *Task) Tick() *Session {
	time.Sleep(s.Options.Interval)

	s.CurrentSession.Advance()

	return s.CurrentSession
}

func (s *Task) Next() {
	sessionType := s.CurrentSession.Type

	expectLongBreak := s.PomodoroId%s.Options.LongBreakAfter == 0

	if sessionType.IsWork() && expectLongBreak {
		s.CurrentSession = NewSession(s, SessionTypeLongBreak)
		return
	}

	if sessionType.IsWork() {
		s.CurrentSession = NewSession(s, SessionTypeShortBreak)
		return
	}

	if sessionType.IsShortBreak() || sessionType.IsLongBreak() {
		s.PomodoroId++
		s.CurrentSession = NewSession(s, SessionTypeWork)
	}
}

func (s *Task) RoundCompleted() bool {
	return s.CurrentSession.Done() && s.CurrentSession.Type.IsLongBreak()
}

func NewTask(title string, settings TaskOptions) Task {
	t := Task{
		Options:    settings,
		Title:      title,
		PomodoroId: 1,
	}

	if settings.LongBreakAfter == 0 {
		settings.LongBreakAfter = DefaultLongBreakAfter
	}

	s := NewSession(&t, SessionTypeWork)

	t.CurrentSession = s

	return t
}
