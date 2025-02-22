package pb

import (
	"os"
	"testing"

	"github.com/bebrochkas/rural_potatoes/core/config"
)

func TestMain(m *testing.M) {
	config.Initialize("../../../.env")
	Initialize()
	os.Exit(m.Run())
}

func Test_getTags(t *testing.T) {
	type args struct {
		description string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Valid description with tags",
			args: args{
				description: "This is a test #golang #testing",
			},
			wantErr: false,
		},
		{
			name: "Valid description without tags",
			args: args{
				description: "This is a test without tags",
			},
			wantErr: false,
		},
		{
			name: "Empty description",
			args: args{
				description: "",
			},
			wantErr: false,
		},
		{
			name: "Description with special characters",
			args: args{
				description: "Special characters test #!@#%^&*()_+",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetTags(tt.args.description)
			if (err != nil) != tt.wantErr {
				t.Errorf("getTags() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}
