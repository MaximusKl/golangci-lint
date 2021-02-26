package golinters

import (
	"fmt"

	"github.com/julz/importas"
	"golang.org/x/tools/go/analysis"

	"github.com/golangci/golangci-lint/pkg/config"
	"github.com/golangci/golangci-lint/pkg/golinters/goanalysis"
)

func NewImportAs(cfg *config.ImportAsSettings) *goanalysis.Linter {
	analyzer := importas.Analyzer
	for alias, pkg := range *cfg {
		analyzer.Flags.Set("alias", fmt.Sprintf("%s:%s", pkg, alias))
	}

	return goanalysis.NewLinter(
		analyzer.Name,
		analyzer.Doc,
		[]*analysis.Analyzer{analyzer},
		nil,
	).WithLoadMode(goanalysis.LoadModeTypesInfo)
}
