package htask

import (
	"errors"
	"testing"
	"time"
)

func TestRunTask(t *testing.T) {
	tests := []struct {
		name      string
		call      func(errCh chan error)
		wantError bool
	}{
		{
			name: "ErrorCase",
			call: func(errCh chan error) {
				errCh <- errors.New("error")
			},
			wantError: true,
		},
		{
			name: "SucceedCase",
			call: func(errCh chan error) {
				return
			},
			wantError: false,
		},
		{
			name: "ErrorAfterDelay",
			call: func(errCh chan error) {
				for i := 0; i < 10; i++ {
					t.Log(i)
					errCh <- nil
				}
			},
			wantError: false,
		},
		{
			name: "ErrorAfterDelay2",
			call: func(errCh chan error) {
				for i := 0; i < 10; i++ {
					t.Log(i)
					if i > 0 {
						errCh <- errors.New("err")
					}
					errCh <- nil
				}
			},
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			task := CreateTask(tt.call)
			err := RunTask(task)
			if (err != nil) != tt.wantError {
				t.Errorf("RunTask() error = %v, wantErr %v", err, tt.wantError)
				return
			}
			time.Sleep(time.Second)
		})
	}
}
