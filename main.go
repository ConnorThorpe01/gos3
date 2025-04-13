package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ConnorThorpe01/gos3/pkg/config"
	"github.com/ConnorThorpe01/gos3/pkg/s3util"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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

	p := tea.NewProgram(initialModel(s3Client))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

}

type screen int

const (
	mainMenu screen = iota
	listBuckets
	createBucket
	uploadFile
	downloadFile
	// Add more as needed
)

type model struct {
	screen    screen
	client    *s3.Client
	choices   []string
	cursor    int
	selected  map[int]struct{}
	bucketMsg string // Optional: store output like bucket list or status messages
}

var (
	cursorStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("12")) // Blue
	highlightStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("15")). // White text
			Background(lipgloss.Color("12")). // Blue background
			Padding(0, 1)
)

func initialModel(client *s3.Client) model {
	return model{
		screen:   mainMenu,
		client:   client,
		choices:  []string{"List Buckets", "Create Bucket", "Download", "Upload"},
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "ctrl+q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter":
			switch m.screen {
			case mainMenu:
				switch m.cursor {
				case 0:
					m.screen = listBuckets
					names, err := s3util.ListBuckets(m.client)
					if err != nil {
						m.choices = []string{"Unable to get list of buckets"}
					}
					m.choices = names
					m.cursor = 0
				case 1:
					m.screen = createBucket
					m.choices = []string{"[Input name placeholder]", "Back"}
					m.cursor = 0
				case 2:
					m.screen = downloadFile
					m.choices = []string{"File1.txt", "File2.txt", "Back"}
					m.cursor = 0
				case 3:
					m.screen = uploadFile
					m.choices = []string{"Pick a file", "Back"}
					m.cursor = 0
				}
			default:
				// Handle the "Back" option for any screen
				if m.choices[m.cursor] == "Back" {
					m.screen = mainMenu
					m.choices = []string{"List Buckets", "Create Bucket", "Download", "Upload"}
					m.cursor = 0
				}
			}

		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m model) View() string {
	// The header
	var s string
	switch m.screen {
	case mainMenu:
		s = "Welcome to Gos3!\n\n"
	case listBuckets:
		s = "Buckets:\n\n"
	case createBucket:
		s = "Create a new bucket:\n\n"
	case uploadFile:
		s = "Upload a file:\n\n"
	case downloadFile:
		s = "Download a file:\n\n"
	}

	// Iterate over our choices
	for i, choice := range m.choices {
		line := fmt.Sprintf("  %s", choice)
		if m.cursor == i {
			line = highlightStyle.Render("> " + choice)
		}
		s += line + "\n"
	}

	// The footer
	s += "\nPress ctrl+q to quit.\n"

	// Send the UI for rendering
	return s
}
