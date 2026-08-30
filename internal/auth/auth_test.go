package auth_test

import(
	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	type test struct {
		name    string 
		headers http.Header
		want    string
		wantErr bool
	}
	// TODO: Add test cases.
	tests := []test{
		{
			name: "ApiKey Valid",
			headers: http.Header{
				"Authorization": []string{"ApiKey asjdiqwodiajdskal"},
			},
			want: "asjdiqwodiajdska",
			wantErr: false,
		},
		{
			name: "ApiKey Invalid",
			headers: http.Header{
				"Authorization": []string{"ApiKey"},
			},
			want: "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := auth.GetAPIKey(tt.headers)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("GetAPIKey() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("GetAPIKey() succeeded unexpectedly")
			}
			if got != tt.want {
				t.Errorf("GetAPIKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

