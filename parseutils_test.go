package main

import (
	"testing"
)

func TestParseCode_WithStmtLength(t *testing.T) {
	type args struct {
		code []byte
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "TestParseCode_WithStmtLength: #1",
			args:    args{
				code: []byte("<?php echo 'hello world';;;;;;;"),
			},
			want:    1,
			wantErr: false,
		},

		{
			name:    "TestParseCode_WithStmtLength: #2",
			args:    args{
				code: []byte("<?php echo 'hello world'"),
			},
			want:    1,
			wantErr: false,
		},

		{
			name:    "TestParseCode_WithStmtLength: #3",
			args:    args{
				code: []byte("<?php echo 'hello world';$user;;;;;;"),
			},
			want:    2,
			wantErr: false,
		},

		{
			name:    "TestParseCode_WithStmtLength: #4",
			args:    args{
				code: []byte(`<?php 

echo 'hello world'

// some thing
// some other

`),
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseCode(tt.args.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != nil && len(got.Stmts) != tt.want {
				t.Errorf("ParseCode() got = %v, want %v", len(got.Stmts), tt.want)
			}
		})
	}
}