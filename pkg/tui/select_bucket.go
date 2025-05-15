package tui

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
	tea "github.com/charmbracelet/bubbletea"
)

type SelectBucketPage struct {
	client *s3.Client
}

func NewSelectBucketPage(c *s3.Client) SelectBucketPage {
	return SelectBucketPage{client: c}
}

func (m SelectBucketPage) Init() tea.Cmd {
	return nil
}

func (m SelectBucketPage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "q":
			return NewLandingPage(m.client), nil
		}
	}
	return m, nil
}

func (m SelectBucketPage) View() string {
	return "Select Bucket Page\n\nPress 'q' to go back."
}
