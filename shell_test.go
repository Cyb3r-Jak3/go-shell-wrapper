package goshellwrapper

import "testing"

func TestRunCommand(t *testing.T) {
	stdout, stderr, err := RunCommand(t.Context(), "echo Hello World!")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if len(stderr) != 0 {
		t.Error("Expected empty stderr, got slice with elements")
	}
	if len(stdout) == 0 {
		t.Error("Expected stdout, got empty slice")
	}
	if stdout[0] != "Hello" {
		t.Errorf("Expected 'Hello', got %s", stdout[0])
	}
	if stdout[1] != "World!" {
		t.Errorf("Expected 'World!', got %s", stdout[1])
	}
}
