package models

import (
	"math"
	"testing"
)

func TestFM_Predict(t *testing.T) {
	type args struct {
		input []float64
	}
	tests := []struct {
		name    string
		args    args
		result float64
		wantErr bool
	}{
		{
			name:    "0",
			args:    args{
				input: []float64{0,0,0,0,0,0,0,0,0,0,0},
			},
			result: 0.00461527772391919,
			wantErr: false,
		},
		{
			name:    "1",
			args:    args{
				input: []float64{1,0,0,0,0,0,0,0,0,0,0},
			},
			result: 0.0007966143831173357,
			wantErr: false,
		},
		{
			name:    "2",
			args:    args{
				input: []float64{1,1,0,0,0,0,0,0,0,0,0},
			},
			result: 0.001830374021359199,
			wantErr: false,
		},
		{
			name:    "3",
			args:    args{
				input: []float64{1,1,1,0,0,0,0,0,0,0,0},
			},
			result: 0.011064764768756912,
			wantErr: false,
		},
		{
			name:    "4",
			args:    args{
				input: []float64{1,1,1,1,0,0,0,0,0,0,0},
			},
			result: 0.023801983502045213,
			wantErr: false,
		},
		{
			name:    "5",
			args:    args{
				input: []float64{1,1,1,1,1,0,0,0,0,0,0},
			},
			result: 0.11977732921261826,
			wantErr: false,
		},
		{
			name:    "6",
			args:    args{
				input: []float64{1,1,1,1,1,1,0,0,0,0,0},
			},
			result: 0.6491256046291641,
			wantErr: false,
		},
		{
			name:    "7",
			args:    args{
				input: []float64{1,1,1,1,1,1,1,0,0,0,0},
			},
			result: 0.98268512450241,
			wantErr: false,
		},
		{
			name:    "8",
			args:    args{
				input: []float64{1,1,1,1,1,1,1,1,0,0,0},
			},
			result: 0.21189946170472698,
			wantErr: false,
		},
		{
			name:    "9",
			args:    args{
				input: []float64{1,1,1,1,1,1,1,1,1,0,0},
			},
			result: 0.9451365715371894,
			wantErr: false,
		},
		{
			name:    "10",
			args:    args{
				input: []float64{1,1,1,1,1,1,1,1,1,1,0},
			},
			result: 0.9989646671977132,
			wantErr: false,
		},
		{
			name:    "11",
			args:    args{
				input: []float64{1,1,1,1,1,1,1,1,1,1,1},
			},
			result: 0.9999725382971929,
			wantErr: false,
		},
	}
	m := &FM{}
	err := m.LoadModelFromJsonFile("./fm_model.json")
	if err != nil {
		t.Error(err)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRet, err := m.Predict(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Predict() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if math.Abs(gotRet - tt.result) > 0.001 {
				t.Errorf("Predict() gotRet = %v, want %v", gotRet, tt.result)
			}
		})
	}
}
