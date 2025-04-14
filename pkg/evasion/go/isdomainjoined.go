package _go

import (
	"embed"

	"github.com/adamstebbing/uru/pkg/common"
	"github.com/adamstebbing/uru/pkg/models"
)

type IsDomainJoinedEvasion struct {
	Name        string
	Description string
	Debug       bool
}

func NewIsDomainJoinedEvasion() models.ObjectModel {
	return &IsDomainJoinedEvasion{
		Name:        "IsDomainJoined",
		Debug:       false,
		Description: "check if current computer is joined to a domain.",
	}
}

func (e *IsDomainJoinedEvasion) GetImports() []string {

	return []string{
		`"syscall"`,
	}
}

func (e *IsDomainJoinedEvasion) RenderInstanciationCode(data embed.FS) (string, error) {

	return common.CommonRendering(data, "templates/go/evasions/isdomainjoined/instanciation.go.tmpl", e)
}

func (e *IsDomainJoinedEvasion) RenderFunctionCode(data embed.FS) (string, error) {

	return common.CommonRendering(data, "templates/go/evasions/isdomainjoined/functions.go.tmpl", e)
}
