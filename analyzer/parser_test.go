package analyzer

import "testing"

func TestParseLogLine(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantLevel string
		wantMsg   string
		wantErr   bool
	}{
		{
			name:      "parse INFO log",
			input:     "2026-01-03 15:30:00 INFO User logged in",
			wantLevel: "INFO",
			wantMsg:   "User logged in",
			wantErr:   false,
		},
		{
			name:      "parse ERROR log",
			input:     "2026-01-03 15:31:00 ERROR Database connection failed",
			wantLevel: "ERROR",
			wantMsg:   "Database connection failed",
			wantErr:   false,
		},
		{
			name:      "parse WARNING log",
			input:     "2026-01-03 15:32:00 WARNING High memory usage",
			wantLevel: "WARNING",
			wantMsg:   "High memory usage",
			wantErr:   false,
		},
		{
			name:    "invalid log format",
			input:   "invalid log line",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entry, err := ParseLogLine(tt.input)

			if tt.wantErr {
				if err == nil {
					t.Error("expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if entry.Level != tt.wantLevel {
				t.Errorf("level = %q, want %q", entry.Level, tt.wantLevel)
			}

			if entry.Message != tt.wantMsg {
				t.Errorf("message = %q, want %q", entry.Message, tt.wantMsg)
			}
		})
	}
}

func TestAnalyzeLogs(t *testing.T) {
	logs := []string{
		"2026-01-03 15:30:00 INFO User logged in",
		"2026-01-03 15:31:00 ERROR Database connection failed",
		"2026-01-03 15:32:00 WARNING High memory usage",
		"2026-01-03 15:33:00 INFO Request processed",
		"2026-01-03 15:34:00 ERROR Timeout",
	}

	stats := AnalyzeLogs(logs)

	if stats.TotalLines != 5 {
		t.Errorf("TotalLines = %d, want 5", stats.TotalLines)
	}

	if stats.Counts["INFO"] != 2 {
		t.Errorf("INFO count = %d, want 2", stats.Counts["INFO"])
	}

	if stats.Counts["ERROR"] != 2 {
		t.Errorf("ERROR count = %d, want 2", stats.Counts["ERROR"])
	}

	if stats.Counts["WARNING"] != 1 {
		t.Errorf("WARNING count = %d, want 1", stats.Counts["WARNING"])
	}
}
