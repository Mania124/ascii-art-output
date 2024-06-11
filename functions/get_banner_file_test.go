package functions

import (
	"os"
	"testing"
)

func TestFileName(t *testing.T) {
	testCases := []struct {
		name     string
		args     []string
		expected string
	}{
		{"NoArgs", []string{"progName"}, ""}, // No CLI args
		{"TwoArgs", []string{"progName", "arg1"}, "standard.txt"}, // Two CLI args
		{"ThreeArgs", []string{"progName", "arg1", "arg2"}, "standard.txt"}, // Three CLI args
		{"FourArgsStandard", []string{"progName", "arg1", "arg2", "standard"}, "standard.txt"}, // Four args, style "standard"
		{"FourArgsThinkertoy", []string{"progName", "arg1", "arg2", "thinkertoy"}, "thinkertoy.txt"}, // Four args, style "thinkertoy"
		{"FourArgsShadow", []string{"progName", "arg1", "arg2", "shadow"}, "shadow.txt"}, // Four args, style "shadow"
		{"FourArgsInvalid", []string{"progName", "arg1", "arg2", "invalid"}, "standard.txt"}, // Four args, invalid style
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			originalArgs := os.Args            // Save the original os.Args
			defer func() { os.Args = originalArgs }() // Restore os.Args after the test
			os.Args = tc.args                  // Set os.Args to the test case arguments

			if got := FileName(); got != tc.expected { 
				t.Errorf("FileName() = %v, want %v", got, tc.expected) 
			}
		})
	}
}
