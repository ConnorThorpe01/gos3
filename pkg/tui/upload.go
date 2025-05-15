package tui

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
	tea "github.com/charmbracelet/bubbletea"
)

type UploadPage struct {
	client *s3.Client
}

func NewUploadPage(c *s3.Client) UploadPage {
	return UploadPage{client: c}
}

func (m UploadPage) Init() tea.Cmd {
	return nil
}

func (m UploadPage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "q":
			return NewLandingPage(m.client), nil
		}
	}
	return m, nil
}

func (m UploadPage) View() string {
	return "Select Bucket Page\n\nPress 'q' to go back."
}
