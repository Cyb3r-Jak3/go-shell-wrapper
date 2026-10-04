package goshellwrapper

import "testing"

func TestRunCommand(t *testing.T) {
	status := RunCommand(t.Context(), "echo Hello World!")
	if status.Error != nil {
		t.Errorf("Expected no error, got %v", status.Error)
	}
	if len(status.Stderr) != 0 {
		t.Error("Expected empty stderr, got slice with elements")
	}
	if len(status.Stdout) == 0 {
		t.Error("Expected stdout, got empty slice")
	}
	if status.Stdout[0] != "Hello" {
		t.Errorf("Expected 'Hello', got %s", status.Stdout[0])
	}
	if status.Stdout[1] != "World!" {
		t.Errorf("Expected 'World!', got %s", status.Stdout[1])
	}
}
