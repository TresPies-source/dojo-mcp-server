package main

import (
	"context"
	"log"
	"os"

	"github.com/TresPies-source/dojo-mcp-server/internal/dojo"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	// Create MCP server
	s := server.NewMCPServer(
		"dojo-genesis",
		"1.0.0",
		server.WithToolCapabilities(true),
		server.WithPromptCapabilities(true),
		server.WithResourceCapabilities(true),
	)

	// Initialize Dojo handler
	dojoHandler := dojo.NewHandler()

	// Register tools
	dojoHandler.RegisterTools(s)

	// Register prompts
	dojoHandler.RegisterPrompts(s)

	// Register resources
	dojoHandler.RegisterResources(s)

	// Start server with stdio transport
	if err := s.Serve(context.Background(), server.WithStdio()); err != nil {
		log.Fatalf("Server error: %v", err)
		os.Exit(1)
	}
}
