package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ConnorThorpe01/gos3/pkg/config"
	"github.com/ConnorThorpe01/gos3/pkg/s3util"
	"github.com/ConnorThorpe01/gos3/pkg/tui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	cfg, err := config.LoadConfig(".config.yml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	fmt.Print("endpoint" + cfg.S3.Endpoint)
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		log.Fatalf("err %w", err)
	}
	defer f.Close()
	s3Client, err := s3util.NewS3Client(cfg)
	if err != nil {
		log.Fatalf("failed to create S3 client: %v", err)
	}

	p := tui.Run(s3Client)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
