package main

import (
	"fmt"
	"github.com/juanjosedemiguel/loadbalancingsim/manager"
	"github.com/juanjosedemiguel/loadbalancingsim/message"
	"github.com/juanjosedemiguel/loadbalancingsim/server"
	"github.com/juanjosedemiguel/loadbalancingsim/task"
)

// Sets the environment for the simulation and executes it.
func main() {
	// Starts the Manager
	manager := manager.NewManager()
	go manager.Run()

	// 16 cores, 128 GB of RAM (one of each server type)
	types := []message.ServerType{message.HighCPU, message.HighMemory, message.Combined}
	servers := make([]*server.Server, len(types))
	for i, servertype := range types {
		port := 8081 + i
		serveraddress := fmt.Sprintf("localhost:%d", port)
		servers[i] = server.NewServer(i, 16, 64, serveraddress, "localhost:8080", servertype)
	}
	for _, server := range servers {
		go server.Run()
	}

	// starts the Task
	task := task.NewTask("generic task", 6)
	go task.Run()
}