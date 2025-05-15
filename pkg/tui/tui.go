package tui

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
	tea "github.com/charmbracelet/bubbletea"
)

func Run(client *s3.Client) *tea.Program {
	return tea.NewProgram(NewLandingPage(client))
}
