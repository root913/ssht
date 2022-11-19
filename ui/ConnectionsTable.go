package ui

import (
	"strconv"
	"strings"

	"github.com/root913/ssht/client"
	"github.com/root913/ssht/config"
	"github.com/root913/ssht/util"

	"github.com/derailed/tview"
	"github.com/gdamore/tcell/v2"
)

func NewConnectionsTable(app *tview.Application, appStyles *config.Styles, appConfig *config.Config) {
	table := tview.NewTable()

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 113 {
			app.Stop()
		}
		return event
	})

	table.SetBackgroundColor(appStyles.Table().BgColor.Color())
	table.SetBorderColor(appStyles.Frame().Border.FgColor.Color())
	table.SetBorderFocusColor(appStyles.Frame().Border.FocusColor.Color())
	table.SetSelectedStyle(tcell.StyleDefault.Foreground(appStyles.Table().CursorFgColor.Color()).Background(appStyles.Table().CursorBgColor.Color()).Attributes(tcell.AttrBold))
	tableFgColor := appStyles.Table().CursorFgColor.Color()

	table.SetBorder(true)
	table.SetBorderAttributes(tcell.AttrBold)
	table.SetBorderPadding(0, 0, 1, 1)
	table.SetSelectable(true, false)
	table.SetFixed(1, 1)
	table.SetSelectionChangedFunc(func(row, column int) {
		if row < 0 {
			return
		}
		if cell := table.GetCell(row, column); cell != nil {
			table.SetSelectedStyle(tcell.StyleDefault.Foreground(tableFgColor).Background(cell.Color).Attributes(tcell.AttrBold))
		}
	})
	table.SetSelectedFunc(func(row int, column int) {
		table.SetSelectable(false, false)
		cell := table.GetCell(row, 0)

		if len(cell.Text) == 0 {
			return
		}

		util.Logger.Debug().
			Str("cell", cell.Text).
			Msg("Selected row")

		//TODO Add suspend
		app.Stop()

		connection := appConfig.App.Get(cell.Text)
		if nil == connection {
			util.Logger.Fatal().Msg("")
		}

		client.Connect(connection)
	})

	headers := [7]string{"# Uuid", "Alias", "Host", "Port", "Username", "Key", "Type"}
	connections := appConfig.App.Connections

	headerFgColor := appStyles.Table().Header.FgColor.Color()
	headerBgColor := appStyles.Table().Header.BgColor.Color()

	cellFgColor := appStyles.Table().FgColor.Color()
	cellBgColor := appStyles.Table().BgColor.Color()

	var col int
	for _, h := range headers {
		c := tview.NewTableCell(strings.ToUpper(h)).
			SetExpansion(1).
			SetAlign(tview.AlignLeft).
			SetSelectable(false)

		table.SetCell(0, col, c)
		c.SetBackgroundColor(headerBgColor)
		c.SetTextColor(headerFgColor)
		col++
	}

	newCell := func(col int, row int, value string) {
		c := tview.NewTableCell(value).
			SetExpansion(1).
			SetAlign(tview.AlignLeft)

		table.SetCell(row, col, c)
		c.SetBackgroundColor(cellBgColor)
		c.SetTextColor(cellFgColor)
	}

	var row int = 1
	for _, connection := range connections {
		util.Logger.Debug().
			Int("row", row).
			Str("alias", connection.Alias).
			Msg("added row")
		newCell(0, row, connection.Uuid)
		newCell(1, row, connection.Alias)
		newCell(2, row, connection.Host)
		newCell(3, row, strconv.FormatInt(int64(connection.Port), 10))
		newCell(4, row, connection.Username)
		newCell(5, row, connection.KeyPath)
		newCell(6, row, connection.Type.String())

		row++
	}

	main := tview.NewFlex().SetDirection(tview.FlexRow)
	main.AddItem(table, 0, 10, true)

	main.SetBackgroundColor(tcell.NewRGBColor(40, 42, 54))

	if err := app.SetRoot(main, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
