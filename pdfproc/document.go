package pdfproc

import (
	"fmt"
	"live/embeded"
	"time"

	"github.com/mechiko/maroto/v2"
	"github.com/mechiko/utility"

	"github.com/mechiko/maroto/v2/pkg/config"
	"github.com/mechiko/maroto/v2/pkg/consts/fontstyle"
	"github.com/mechiko/maroto/v2/pkg/repository"

	"github.com/mechiko/maroto/v2/pkg/props"
)

func (p *pdfProc) PdfDocument() (err error) {
	start := time.Now()
	err = p.BuildPages(true, true)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	fmt.Printf("buid pages %v\n", time.Since(start))
	start = time.Now()
	err = p.DocumentGenerate()
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	fmt.Printf("generate document %v\n", time.Since(start))
	fileName := "PDF.pdf"
	err = p.document.Save(fileName)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}

func (p *pdfProc) buildMaroto() (err error) {
	customFont := "roboto"
	customFonts, err := repository.New().
		AddUTF8FontFromBytes(customFont, fontstyle.Normal, embeded.Regular).
		AddUTF8FontFromBytes(customFont, fontstyle.Italic, embeded.Italic).
		AddUTF8FontFromBytes(customFont, fontstyle.Bold, embeded.Bold).
		AddUTF8FontFromBytes(customFont, fontstyle.BoldItalic, embeded.BoldItalic).
		Load()
	if err != nil {
		return err
	}
	// …custom font setup…
	builder := config.NewBuilder().WithCustomFonts(customFonts)
	cfg := builder.WithDefaultFont(&props.Font{Family: customFont}).Build()
	cfg.Dimensions.Height = p.templateDatamatrix.PageHeight
	cfg.Dimensions.Width = p.templateDatamatrix.PageWidth
	cfg.Margins.Bottom = 0
	cfg.Margins.Top = 1
	cfg.Margins.Left = 2
	cfg.Margins.Right = 2
	cfg.DefaultFont.Size = 4
	p.maroto = maroto.New(cfg)
	p.maroto = maroto.NewMetricsDecorator(p.maroto)
	return nil
}

func (p *pdfProc) DocumentGenerate() (err error) {
	doc, err := p.maroto.Generate()
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	p.document = doc
	return nil
}

func (p *pdfProc) BuildPages(datamatrix, bar bool) error {
	if datamatrix {
		code := `0105000213100066215aDos=X93a2MS`
		codeCis, _ := utility.ParseCisInfo(code)
		idx := fmt.Sprintf("%06d", 123)
		if err := p.addPageByTemplate(p.templateDatamatrix, codeCis, "Б1", idx); err != nil {
			return fmt.Errorf("%w", err)
		}
	}
	if bar {
		code := `0105000213100066215aDos=X93a2MS`
		codeCis, _ := utility.ParseCisInfo(code)
		if err := p.addPageByTemplate(p.templateBar, codeCis, "Б1", "000001"); err != nil {
			return fmt.Errorf("%w", err)
		}
	}
	return nil
}

func (p *pdfProc) addPageByTemplate(tmpl *MarkTemplate, cis *utility.CisInfo, party string, idx string) error {
	if tmpl == nil {
		return fmt.Errorf("add page template is nil")
	}
	page, err := p.Page(tmpl, cis, party, idx)
	if err != nil {
		return fmt.Errorf("page error %w", err)
	}
	if page == nil {
		return fmt.Errorf("page nil")
	}
	p.maroto.AddPages(page)
	return nil
}
