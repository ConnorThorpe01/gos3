package tui

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	tea "github.com/charmbracelet/bubbletea"
)

type LandingPage struct {
	choices []string
	cursor  int
	client  *s3.Client
}

func NewLandingPage(client *s3.Client) LandingPage {
	return LandingPage{
		choices: []string{"Create Bucket", "Select Bucket", "Download", "Upload"},
		client:  client,
	}
}

func (m LandingPage) Init() tea.Cmd {
	return nil
}

func (m LandingPage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "ctrl+q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter":
			switch m.cursor {
			case 0:
				return NewCreateBucketPage(m.client), nil
			case 1:
				return NewSelectBucketPage(m.client), nil
			case 2:
				return NewDownloadPage(m.client), nil
			case 3:
				return NewUploadPage(m.client), nil
			}
		}
	}
	return m, nil
}

func (m LandingPage) View() string {
	s := "What do you want to do?\n\n"
	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}
	s += "\nPress ctrl+q to quit.\n"
	return s
}
