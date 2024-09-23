package core

import (
	"testing"
)

func TestCleanResponse(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{
			name:    "Valid input with splitParam",
			input:   prefix + "This is the content",
			want:    "This is the content",
			wantErr: false,
		},
		{
			name:    "Valid input with splitParam and additional content",
			input:   prefix + "This is the content" + prefix + "More content",
			want:    "This is the content" + prefix + "More content",
			wantErr: false,
		},
		{
			name:    "Invalid input without splitParam",
			input:   "This is invalid content",
			want:    "",
			wantErr: true,
		},
		{
			name:    "Empty input",
			input:   "",
			want:    "",
			wantErr: true,
		},
		{
			name:    "Input with only splitParam",
			input:   prefix,
			want:    "",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := cleanResponse(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("cleanResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("cleanResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}
