package config

import (
	"errors"
	"time"
)

type App struct {
	Skin        string       `yaml:"skin"`
	Connections []Connection `yaml:connections`
}

func NewApp() *App {
	return &App{
		Skin:        "dracula",
		Connections: []Connection{},
	}
}

func (app *App) HasConnections() bool {
	return len(app.Connections) > 0
}

func (app *App) Get(uuidOrAlias string) *Connection {
	for _, connection := range app.Connections {
		if connection.Uuid == uuidOrAlias || connection.Alias == uuidOrAlias {
			return &connection
		}
	}
	return nil
}

func (app *App) AddConnection(connection *Connection) {
	app.Connections = append(app.Connections, *connection)
}

func (app *App) RemoveConnection(connection *Connection) bool {
	index := app.IndexOfConnection(connection.Uuid)
	if index == -1 {
		return false
	}

	app.Connections = append(app.Connections[:index], app.Connections[index+1:]...)

	return true
}

func (app *App) RemoveAllConnections() {
	for _, connection := range app.Connections {
		app.RemoveConnection(&connection)
	}
}

func (app *App) IndexOfConnection(uuid string) int {
	for k, v := range app.Connections {
		if uuid == v.Uuid {
			return k
		}
	}
	return -1
}

func (app *App) SetConnectionAlias(uuid string, alias string) bool {
	index := app.IndexOfConnection(uuid)
	if index == -1 {
		return false
	}

	app.Connections[index].Alias = alias
	app.Connections[index].UpdatedAt = time.Now()

	return true
}

func (app *App) CheckForduplicates(connection *Connection) error {
	for _, c := range app.Connections {
		if c.Host == connection.Host && c.Port == connection.Port && c.Password == connection.Password && c.KeyPath == connection.KeyPath && c.KeyPass == connection.KeyPass {
			return errors.New("This connection alredy exists.")
		}
	}

	return nil
}
