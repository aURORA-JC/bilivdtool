package core

import "testing"

func TestDoMergeOperations(t *testing.T) {
	type args struct {
		videoPath  string
		audioPath  string
		outputPath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test 1",
			args: args{
				videoPath: "./v.m4s",
				audioPath: "./a.m4s",
				outputPath: "./ouput.mp4",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DoMergeOperations(tt.args.videoPath, tt.args.audioPath, tt.args.outputPath); (err != nil) != tt.wantErr {
				t.Errorf("DoMergeOperations() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
