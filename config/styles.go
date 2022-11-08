package config

import (
	"os"

	"github.com/gdamore/tcell/v2"
	"gopkg.in/yaml.v2"
)

type (
	// Color represents a color.
	Color string

	// Colors tracks multiple colors.
	Colors []Color

	// Styles tracks K9s styling options.
	Styles struct {
		Main Style `yaml:"main"`
	}

	// // Style tracks K9s styles.
	Style struct {
		Body  Body  `yaml:"body"`
		Frame Frame `yaml:"frame"`
		Views Views `yaml:"views"`
	}

	// // Body tracks body styles.
	Body struct {
		FgColor   Color `yaml:"fgColor"`
		BgColor   Color `yaml:"bgColor"`
		LogoColor Color `yaml:"logoColor"`
	}

	// Frame tracks frame styles.
	Frame struct {
		Border Border `yaml:"border"`
	}

	// Views tracks individual view styles.
	Views struct {
		Table Table `yaml:"table"`
	}

	// Border tracks border styles.
	Border struct {
		FgColor    Color `yaml:"fgColor"`
		FocusColor Color `yaml:"focusColor"`
	}

	// Table tracks table styles.
	Table struct {
		FgColor       Color       `yaml:"fgColor"`
		BgColor       Color       `yaml:"bgColor"`
		CursorFgColor Color       `yaml:"cursorFgColor"`
		CursorBgColor Color       `yaml:"cursorBgColor"`
		MarkColor     Color       `yaml:"markColor"`
		Header        TableHeader `yaml:"header"`
	}

	// TableHeader tracks table header styles.
	TableHeader struct {
		FgColor     Color `yaml:"fgColor"`
		BgColor     Color `yaml:"bgColor"`
		SorterColor Color `yaml:"sorterColor"`
	}
)

const (
	// DefaultColor represents  a default color.
	DefaultColor Color = "default"

	// TransparentColor represents the terminal bg color.
	TransparentColor Color = "-"
)

// Color returns a view color.
func (c Color) Color() tcell.Color {
	if c == DefaultColor {
		return tcell.ColorDefault
	}

	return tcell.GetColor(string(c)).TrueColor()
}

func newStyle() Style {
	return Style{
		Body:  newBody(),
		Frame: newFrame(),
		Views: newViews(),
	}
}

func newViews() Views {
	return Views{
		Table: newTable(),
	}
}

func newFrame() Frame {
	return Frame{
		Border: newBorder(),
	}
}

func newBody() Body {
	return Body{
		FgColor:   "cadetblue",
		BgColor:   "black",
		LogoColor: "orange",
	}
}

func newTable() Table {
	return Table{
		FgColor:       "aqua",
		BgColor:       "black",
		CursorFgColor: "black",
		CursorBgColor: "aqua",
		MarkColor:     "palegreen",
		Header:        newTableHeader(),
	}
}

func newTableHeader() TableHeader {
	return TableHeader{
		FgColor:     "white",
		BgColor:     "black",
		SorterColor: "aqua",
	}
}

func newBorder() Border {
	return Border{
		FgColor:    "dodgerblue",
		FocusColor: "lightskyblue",
	}
}

func NewStyles() *Styles {
	return &Styles{
		Main: newStyle(),
	}
}

// FgColor returns the foreground color.
func (s *Styles) FgColor() tcell.Color {
	return s.Body().FgColor.Color()
}

// BgColor returns the background color.
func (s *Styles) BgColor() tcell.Color {
	return s.Body().BgColor.Color()
}

// Body returns body styles.
func (s *Styles) Body() Body {
	return s.Main.Body
}

// Frame returns frame styles.
func (s *Styles) Frame() Frame {
	return s.Main.Frame
}

// Table returns table styles.
func (s *Styles) Table() Table {
	return s.Main.Views.Table
}

// Load skin configuration from file.
func (s *Styles) LoadSkin(skin string) error {
	path := "./skins/" + skin + ".yaml"
	f, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(f, s); err != nil {
		return err
	}

	return nil
}
