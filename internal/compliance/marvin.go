package compliance

import (
	"encoding/json"
	"os/exec"
)

type ScanResult struct {
	KubeVersion struct {
		Major        string `json:"major"`
		Minor        string `json:"minor"`
		GitVersion   string `json:"gitVersion"`
		GitCommit    string `json:"gitCommit"`
		GitTreeState string `json:"gitTreeState"`
		BuildDate    string `json:"buildDate"`
		GoVersion    string `json:"goVersion"`
		Compiler     string `json:"compiler"`
		Platform     string `json:"platform"`
	} `json:"kubeVersion"`
	Checks []struct {
		ID           string              `json:"id"`
		Message      string              `json:"message"`
		Severity     string              `json:"severity"`
		Builtin      bool                `json:"builtin"`
		Path         string              `json:"path"`
		Status       string              `json:"status"`
		Failed       map[string][]string `json:"failed"`
		Passed       map[string][]string `json:"passed"`
		Skipped      map[string][]string `json:"skipped"`
		Errors       []interface{}       `json:"errors"`
		TotalFailed  int                 `json:"totalFailed"`
		TotalPassed  int                 `json:"totalPassed"`
		TotalSkipped int                 `json:"totalSkipped"`
	} `json:"checks"`
}

func Scan(server, token string) (*ScanResult, error) {
	// Call Marvin binary and capture output
	cmd := exec.Command("marvin", "scan", "--server", server, "--token", token, "--output", "json")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	// Parse output as JSON
	var result ScanResult
	err = json.Unmarshal(output, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}
