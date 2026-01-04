package analyzer

import (
	"errors"
	"regexp"
)

type LogEntry struct {
	Timestamp string
	Level     string
	Message   string
}

type Statistics struct {
	TotalLines int
	Counts     map[string]int
}

func ParseLogLine(line string) (*LogEntry, error) {
	re := regexp.MustCompile(`^(\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}) (INFO|ERROR|WARNING|DEBUG) (.+)$`)
	matches := re.FindStringSubmatch(line)

	if len(matches) != 4 {
		return nil, errors.New("invalid log format")
	}

	return &LogEntry{
		Timestamp: matches[1],
		Level:     matches[2],
		Message:   matches[3],
	}, nil
}

func AnalyzeLogs(lines []string) Statistics {
	stats := Statistics{
		Counts: make(map[string]int),
	}

	for _, line := range lines {
		entry, err := ParseLogLine(line)
		if err != nil {
			continue
		}

		stats.TotalLines++
		stats.Counts[entry.Level]++
	}

	return stats
}
