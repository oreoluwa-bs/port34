package main

import (
	"bufio"
	"context"
	"log"
	"os/exec"
	"strings"
	"time"
)

type Process struct {
	Port        string
	Application string
	PID         string
}

func GetProcesses() []Process {
	var p []Process
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "lsof", "-i", "-P", "-n")

	out, err := cmd.Output()
	if err != nil {
		log.Panic("could not run command: ", err)
	}

	scanner := bufio.NewScanner(strings.NewReader(string(out)))

	scanner.Scan()

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) < 9 {
			continue
		}

		entry := Process{
			Application: fields[0],
			PID:         fields[1],
			// User:    fields[2],
			// FD:      fields[3],
			// Type:    fields[4],
			// Device:  fields[5],
			// SizeOff: fields[6],
			// Node:    fields[7],
			Port: strings.Join(fields[8:], " "), // Handle cases where Name has spaces
		}

		p = append(p, entry)
	}

	if err := scanner.Err(); err != nil {
		log.Panic("Error reading lsof output:", err)
	}

	return p
}

func (p *Process) kill() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, "kill", "-9", p.PID)

	_, err := cmd.Output()
	if err != nil {
		log.Panicf("could not terminate process (%s %s): %v \n", p.Application, p.PID, err)
	}
}
