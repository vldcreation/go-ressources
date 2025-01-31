package interop

import (
	"path/filepath"
	"testing"
)

func getAbs(path string) string {
	abs, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return abs
}

func TestRunner(t *testing.T) {
	type args struct {
		interop Interop
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test#JavascriptRunner",
			args: args{
				interop: Interop{
					language: "javascript",
					filePath: getAbs("example/example.js"),
				},
			},
			want: "Hello World from Javascript",
		},
		{
			name: "Test#PythonRunner",
			args: args{
				interop: Interop{
					language: "python",
					filePath: getAbs("example/example.py"),
				},
			},
			want: "Hello World from Python",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewInteropRunner(tt.args.interop).Run()
			t.Logf("Success Invoke from %s: %+v\n", tt.args.interop.language, got)
			if (err != nil) != tt.wantErr {
				t.Errorf("Runner.Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Runner.Run() = %v, want %v", got, tt.want)
			}
		})
	}
}
