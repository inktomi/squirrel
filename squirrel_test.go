package main_test

import (
	main "github.com/inktomi/squirrel"
	"github.com/inktomi/squirrel/telemetry"
	"testing"
	"time"
)

func TestReportWeightIfNeeded(t *testing.T) {
	type args struct {
		lastReported   int64
		adafruitClient *telemetry.Adafruit
		weight         float64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Too Fast",
			args{lastReported: time.Now().Unix(), adafruitClient: &telemetry.Adafruit{}, weight: 10},
			true},
		{"Correct Timing",
			args{lastReported: time.Now().Unix() - 30, adafruitClient: &telemetry.Adafruit{}, weight: 10},
			true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := main.ReportWeightIfNeeded(tt.args.lastReported, tt.args.adafruitClient, tt.args.weight); (err != nil) != tt.wantErr {
				t.Errorf("ReportWeightIfNeeded() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
