package constants_test

import (
	"testing"

	. "github.com/halfbakedio/saas/constants"

	"github.com/stretchr/testify/assert"
)

func TestEnv_String(t *testing.T) {
	tests := []struct {
		name string
		env  Env
		want string
	}{
		{
			name: "Development",
			env:  Development,
			want: "development",
		},
		{
			name: "Test",
			env:  Test,
			want: "test",
		},
		{
			name: "Staging",
			env:  Staging,
			want: "staging",
		},
		{
			name: "Production",
			env:  Production,
			want: "production",
		},
		{
			name: "Undefined",
			env:  Undefined,
			want: "undefined",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.env.String())
		})
	}
}

func TestIsDevelopment(t *testing.T) {
	tests := []struct {
		name string
		env  string
		want bool
	}{
		{
			name: "Development",
			env:  "development",
			want: true,
		},
		{
			name: "Test",
			env:  "test",
			want: false,
		},
		{
			name: "Staging",
			env:  "staging",
			want: false,
		},
		{
			name: "Production",
			env:  "production",
			want: false,
		},
		{
			name: "Undefined",
			env:  "undefined",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, IsDevelopment(tt.env))
		})
	}
}

func TestIsTest(t *testing.T) {
	tests := []struct {
		name string
		env  string
		want bool
	}{
		{
			name: "Development",
			env:  "development",
			want: false,
		},
		{
			name: "Test",
			env:  "test",
			want: true,
		},
		{
			name: "Staging",
			env:  "staging",
			want: false,
		},
		{
			name: "Production",
			env:  "production",
			want: false,
		},
		{
			name: "Undefined",
			env:  "undefined",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, IsTest(tt.env))
		})
	}
}

func TestIsStaging(t *testing.T) {
	tests := []struct {
		name string
		env  string
		want bool
	}{
		{
			name: "Development",
			env:  "development",
			want: false,
		},
		{
			name: "Test",
			env:  "test",
			want: false,
		},
		{
			name: "Staging",
			env:  "staging",
			want: true,
		},
		{
			name: "Production",
			env:  "production",
			want: false,
		},
		{
			name: "Undefined",
			env:  "undefined",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, IsStaging(tt.env))
		})
	}
}

func TestIsProduction(t *testing.T) {
	tests := []struct {
		name string
		env  string
		want bool
	}{
		{
			name: "Development",
			env:  "development",
			want: false,
		},
		{
			name: "Test",
			env:  "test",
			want: false,
		},
		{
			name: "Staging",
			env:  "staging",
			want: false,
		},
		{
			name: "Production",
			env:  "production",
			want: true,
		},
		{
			name: "Undefined",
			env:  "undefined",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, IsProduction(tt.env))
		})
	}
}
