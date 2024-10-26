package main

type Process struct {
	Port        string
	Application string
	PID         string
}

func GetProcesses() []Process {
	var p []Process

	p = append(p, Process{
		Port:        "3000",
		PID:         "93940",
		Application: "Nodejs runtime",
	},
	)

	return p
}
