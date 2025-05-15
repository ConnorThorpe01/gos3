package tui

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
	tea "github.com/charmbracelet/bubbletea"
)

type CreateBucketPage struct {
	client *s3.Client
}

func NewCreateBucketPage(c *s3.Client) CreateBucketPage {
	return CreateBucketPage{client: c}
}

func (m CreateBucketPage) Init() tea.Cmd {
	return nil
}

func (m CreateBucketPage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "q":
			return NewLandingPage(m.client), nil
		}
	}
	return m, nil
}

func (m CreateBucketPage) View() string {
	return "Create Bucket Page\n\nPress 'q' to go back."
}
