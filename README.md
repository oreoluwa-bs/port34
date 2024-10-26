# Port34

I often can't remember what commands to list and terminate processes on my computer.Meet Port34, a terminal-based user interface (TUI) designed to help you visualize and manage processes directly from your terminal.

## Prerequisites

- Go (version 1.23 or higher)
- Make (for executing the Makefile commands)

## Running the Application

The Makefile provides several commands to build, run, and manage the application.

### 1. Build the Application

To compile the Go code into a binary:

```bash
make build
```

This will create an executable named port-monitor in the project directory.

### 2. Run the Application

To build and then run the application directly:

```bash
# production
make run

# development
make run-dev
```
