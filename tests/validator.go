package validation

import (
	"os"
	"strings"

	"k8s.io/kubectl/pkg/cmd"
)

func validArgs(args []string) bool {
	//fmt.Println("Args:", args)
	res := map[string]bool{
		"pods":        true,
		"deployment":  true,
		"statefulset": true,
		"daemonset":   true,
		"replicaset":  true,
		"service":     true,
		"ingress":     true,
		"configmap":   true,
		"secret":      true,
		"nodes":       true,
		"namespaces":  true,
		"jobs":        true,
		"pv":          true,
		"pvc":         true,
	}
	os.Args = args
	command := cmd.NewDefaultKubectlCommand()
	//newCommand is a cobra command which is the complete command with all children (so if the command is 'kubectl get' then command will be kubectl and newCommand will be kubectl get
	//flags_and_unrecognized_args is a list of all the flags in the command and all the arguments that aren't kubectl child commands (so not "get", "apply", etc.)
	//We only want flags_and_unrecognized_args to contain api-resources and valid flags
	newCommand, flags_and_unrecognized_args, err := command.Find(args[1:])
	if err != nil {
		return false
	}
	for _, arg := range flags_and_unrecognized_args {
		if arg[0] == '-' {
			continue
		}
		if !res[arg] {
			return false
		}
	}
	if err := newCommand.ParseFlags(args); err != nil {
		if !strings.HasPrefix(err.Error(), "flag needs an argument") {
			return false
		}
	}
	return true
}
