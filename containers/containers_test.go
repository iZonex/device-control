package containers

import (
	"testing"
)

func TestListImages(t *testing.T) {
	type args struct {
		contextImages string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "test1",
			args:    args{contextImages: "example"},
			wantErr: false},
		{name: "test2",
			args:    args{contextImages: "nonexis?ted"},
			wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ListImages(tt.args.contextImages); (err != nil) != tt.wantErr {
				t.Errorf("ListImages() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestListContainers(t *testing.T) {
	type args struct {
		c container
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "test1",
			args:    args{c: container{context: "example"}},
			wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ListContainers(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("ListContainers() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
