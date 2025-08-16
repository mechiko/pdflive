package main

import (
	"github.com/mechiko/maroto/v2/pkg/components/col"
	"github.com/mechiko/maroto/v2/pkg/components/row"
	"github.com/mechiko/maroto/v2/pkg/components/text"
	"github.com/mechiko/maroto/v2/pkg/consts/align"
	"github.com/mechiko/maroto/v2/pkg/consts/fontstyle"
	"github.com/mechiko/maroto/v2/pkg/core"
	"github.com/mechiko/maroto/v2/pkg/props"
)

func textRowsBar(pg core.Page) {
	pg.Add(
		text.NewAutoRow("«HARP LAGER»", props.Text{
			Style: fontstyle.Bold,
			Size:  5,
			Align: align.Center,
		}),
		text.NewAutoRow("ПИВО СВЕТЛОЕ ФИЛЬТРОВАННОЕ ПАСТЕРИЗОВАННОЕ «ХАРП ЛАГЕР»", props.Text{
			Style: fontstyle.Bold,
			Size:  5,
			Align: align.Center,
		}),
		row.New(2).Add(),
		row.New(1.5).Add(
			text.NewCol(3, "ПРОИЗВЕДЕНО И РАЗЛИТО:", props.Text{
				Top:   0.1,
				Style: fontstyle.Bold,
				Size:  3.5,
				Align: align.Left,
			}),
			text.NewCol(9, "\"GUINNESS & CO\" ST. JAMES S GALE, DUBLIN 8, IRELAND, ИРЛАНДИЯ.", props.Text{
				Top:   0.1,
				Style: fontstyle.Normal,
				Size:  3.5,
				Align: align.Left,
			}),
		),
		row.New(1.5).Add(
			col.New(4).Add(
				text.New("ИМПОРТЕР В РЕСПУБЛИКУ БЕЛАРУСЬ:", props.Text{
					Top:   0.1,
					Style: fontstyle.Bold,
					Size:  3.4,
					Align: align.Justify,
				}),
			),
			col.New(8).Add(
				text.New("ООО \"ПЕКАРСКОЕ ДЕЛО\", БЕЛАРУСЬ, 220037, РЕСПУБЛИКА БЕЛАРУСЬ, Г. МИНСК", props.Text{
					Top:   0.1,
					Style: fontstyle.Normal,
					Size:  3.5,
					Align: align.Justify,
				}),
			),
		),
		text.NewAutoRow("ПЕР.ТВЕРДЫЙ, 1-ЫЙ, ДОМ 11, ЛИТ. И, 4/К., КАБ. 49. ТЕЛ.  +375 25 523 06 28, e-mail: pekarskaia-sprava@mail.ru.", props.Text{
			Style: fontstyle.Normal,
			Size:  4,
			Align: align.Left,
		}),
		row.New(1.5).Add(
			col.New(2).Add(
				text.New("ИМПОРТЕР В РФ:", props.Text{
					Top:   0.1,
					Style: fontstyle.Bold,
					Size:  3.4,
					Align: align.Justify,
				}),
			),
			col.New(10).Add(
				text.New("ООО «ЛЮКАТ-ПЕКАРНИ», РОССИЯ, САНКТ-ПЕТЕРБУРГ, 191167, НАБ. ОБВОДНОГО КАНАЛА,", props.Text{
					Top:   0.1,
					Style: fontstyle.Normal,
					Size:  3.5,
					Align: align.Left,
				}),
			),
		),
		text.NewAutoRow("Д. 21, ЛИТЕР А, ПОМ. 1-Н/1, ЭТАЖ 1, ТЕЛ.: +7 (812) 429-39-46, e-mail: lucate-import@mail.ru;", props.Text{
			Style: fontstyle.Normal,
			Size:  3.5,
			Align: align.Left,
		}),
		text.NewAutoRow("СРОК ГОДНОСТИ 305 ДНЕЙ ПРИ СОБЛЮДЕНИИ УСЛОВИЙ ХРАНЕНИЯ И ТРАНСПОРТИРОВАНИЯ.", props.Text{
			Style: fontstyle.Bold,
			Size:  4,
			Align: align.Center,
		}),
		text.NewAutoRow(txt3, props.Text{
			Style: fontstyle.Bold,
			Size:  4,
			Align: align.Left,
		}),
		text.NewAutoRow("Количество: 24 б. Х 0,44л", props.Text{
			Style: fontstyle.Bold,
			Size:  4,
			Align: align.Left,
		}),
		text.NewAutoRow("Дата розлива: 24.03.2025", props.Text{
			Style: fontstyle.Bold,
			Size:  4,
			Align: align.Left,
		}),
		text.NewAutoRow("Употребить до:23.01.2026", props.Text{
			Style: fontstyle.Bold,
			Size:  4,
			Align: align.Left,
		}),
		text.NewAutoRow("Номер партии соответствует дате розлива.", props.Text{
			Style: fontstyle.Bold,
			Size:  4,
			Align: align.Left,
		}),
	)
}

var txt3 = `Хранить в вентилируемых, не имеющих посторонних запахов помещениях, исключающих воздействие прямого солнечного света, при температуре от +2°С до +25°С и относительной влажности не более 85 %.`
