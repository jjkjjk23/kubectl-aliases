package validation

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

func TestValidCases(t *testing.T) {
	validCases := [][]string{
		{"kubectl", "get", "pods", "--namespace=default", "--namespace=kube-system"},
		{"kubectl", "rollout", "history", "statefulset", "-o=yaml"},
		{"kubectl", "rollout", "undo"},
	}
	for _, command := range validCases {
		if !validArgs(command) {
			t.Fatalf("Valid case rejected: %q", command)
		}
	}
}

func TestInvalidCases(t *testing.T) {
	invalidCases := [][]string{
		{"kubectl", "rollout", "status", "statefulset", "-o=yaml"},
		{"kubectl", "rollout", "status", "statefulset", "--all-namespaces"},
		{"kubectl", "get", "get"},
		{"kubectl", "get", "--namepace=default", "pods"},
	}
	for _, command := range invalidCases {
		if validArgs(command) {
			t.Fatalf("Invalid case accepted: %q", command)
		}
	}
}

func TestGeneratedAliases(t *testing.T) {
	file, err := os.Open("../.kubectl_aliases")
	if err != nil {
		t.Fatalf("Fail: failed to read file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		splitEquation := strings.Split(string(scanner.Bytes()), "=")
		if len(splitEquation) == 1 {
			continue
		}
		command := splitEquation[1]
		if command[0] == '\'' {
			command = command[1:]
		}
		if command[len(command)-1] == '\'' {
			command = command[:len(command)-1]
		}
		args := strings.Split(command, " ")
		if !validArgs(args) {
			t.Fatalf("One of the generated cases is considered invalid: %q", args)
		}
	}
}
