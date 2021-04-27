package main

import (
	"reflect"
	"testing"
)

func TestStrip(t *testing.T) {
	type args struct {
		code []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name:    "TestStrip:#1",
			args:    args{
				code: []byte(`
<?php

echo "hello";

`),
			},
			want:    []byte("<?php echo\"hello\";"),
			wantErr: false,
		},

		{
			name:    "TestStrip:#2",
			args:    args{
				code: []byte(`
<?php

use App\Models\User;
use Illuminate\Support\Facades\DB;

// Customer Support
// when a user does not receive a password reset email

$user = User::where('email', 'pmoore@example.net')->first();

$user->password = bcrypt('your-new-secure-password');

$user->save();

$user;

`),
			},
			want:    []byte("<?php use App\\Models\\User;use Illuminate\\Support\\Facades\\DB;$user=User::where('email','pmoore@example.net')->first();$user->password=bcrypt('your-new-secure-password');$user->save();$user;"),
			wantErr: false,
		},
		{
			name:    "TestStrip:#3",
			args:    args{
				code: []byte(`
<?php
$users = User::where('name', 'LIKE', '%B01%')
->get();



$users;

`),
			},
			want:    []byte("<?php $users=User::where('name','LIKE','%B01%')->get();$users;"),
			wantErr: false,
		},

		{
			name:    "TestStrip:#4",
			args:    args{
				code: []byte(`
// User::query()
// ->all()

DB::select('
SELECT COUNT(*)
FROM users
');

// hello

`),
			},
			want:    []byte("<?php DB::select('\nSELECT COUNT(*)\nFROM users\n');"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Strip(tt.args.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("Strip() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Strip() got = %v, want %v", got, tt.want)
			}
		})
	}
}
