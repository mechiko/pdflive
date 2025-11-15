package pdfproc

import (
	"fmt"
	"live/assets"

	"github.com/mechiko/maroto/v2/pkg/core"
)

type pdfProc struct {
	maroto             core.Maroto
	assets             *assets.Assets
	templateDatamatrix *MarkTemplate
	templateBar        *MarkTemplate
	document           core.Document
	debug              bool
}

func New(tmplDatamatrix, tmplBar *MarkTemplate, assets *assets.Assets) (*pdfProc, error) {
	p := &pdfProc{
		templateDatamatrix: tmplDatamatrix,
		templateBar:        tmplBar,
		assets:             assets,
	}
	if err := p.buildMaroto(); err != nil {
		return nil, fmt.Errorf("build maroto error %w", err)
	}
	return p, nil
}
