package tui

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
	tea "github.com/charmbracelet/bubbletea"
)

type DownloadPage struct {
	client *s3.Client
}

func NewDownloadPage(c *s3.Client) DownloadPage {
	return DownloadPage{client: c}
}

func (m DownloadPage) Init() tea.Cmd {
	return nil
}

func (m DownloadPage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "q":
			return NewLandingPage(m.client), nil
		}
	}
	return m, nil
}

func (m DownloadPage) View() string {
	return "Select Bucket Page\n\nPress 'q' to go back."
}
