package internal

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestResponse_Write(t *testing.T) {
	type args struct {
		status int
		data   any
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Writing success",
			args: args{
				status: http.StatusOK,
				data:   map[string]string{"message": "success"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Response{}
			w := httptest.NewRecorder()
			r.Write(w, tt.args.status, tt.args.data)

			if w.Code != tt.args.status {
				t.Errorf("Expected status %v, got %v", tt.args.status, w.Code)
			}
		})
	}
}

func TestResponse_WriteError(t *testing.T) {
	type args struct {
		status  int
		message string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Writing error",
			args: args{
				status:  http.StatusBadRequest,
				message: "error message",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Response{}
			w := httptest.NewRecorder()
			r.WriteError(w, tt.args.status, tt.args.message)

			if w.Code != tt.args.status {
				t.Errorf("Expected status %v, got %v", tt.args.status, w.Code)
			}
		})
	}
}
