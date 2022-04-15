// Package rove is an Ambient plugin that provides MySQL migrations.
package rove

import (
	"context"
	"fmt"

	"github.com/ambientkit/ambient"

	"github.com/josephspurrier/rove"
	"github.com/josephspurrier/rove/pkg/adapter/mysql"
)

// Plugin represents an Ambient plugin.
type Plugin struct {
	*ambient.PluginBase
	dbinfo       *DBInfo
	dbconnection *mysql.MySQL
}

// DBInfo represent the DB info.
type DBInfo struct {
	DBName       string
	Charset      string
	Collate      string
	Changeset    string
	Prefix       string
	Verbose      bool
	ChecksumMode rove.ChecksumMode
}

// New returns an Ambient plugin that provides MySQl migration.
func New(dbinfo *DBInfo) *Plugin {
	if dbinfo == nil {
		dbinfo = new(DBInfo)
	}
	if dbinfo.Charset == "" {
		dbinfo.Charset = "utf8mb4"
	}
	if dbinfo.Collate == "" {
		dbinfo.Collate = "utf8mb4_unicode_ci"
	}
	if dbinfo.DBName == "" {
		dbinfo.DBName = "main"
	}

	return &Plugin{
		PluginBase: &ambient.PluginBase{},
		dbinfo:     dbinfo,
	}
}

// PluginName returns the plugin name.
func (p *Plugin) PluginName(context.Context) string {
	return "rove"
}

// PluginVersion returns the plugin version.
func (p *Plugin) PluginVersion(context.Context) string {
	return "1.0.0"
}

// Enable accepts the toolkit.
func (p *Plugin) Enable(toolkit *ambient.Toolkit) error {
	err := p.PluginBase.Enable(toolkit)
	if err != nil {
		return err
	}

	// Create the MySQL connection information from environment variables.
	connInfo, err := mysql.NewConnection(p.dbinfo.Prefix)
	if err != nil {
		return err
	}

	// Connect without database.
	c, err := connInfo.Connect(false)
	if err != nil {
		return err
	}

	// Create the database if it doesn't exist.
	_, err = c.DB.Exec(fmt.Sprintf(
		"CREATE DATABASE IF NOT EXISTS %v DEFAULT CHARSET = %v COLLATE = %v;",
		p.dbinfo.DBName, p.dbinfo.Charset, p.dbinfo.Collate))
	if err != nil {
		return err
	}

	err = c.DB.Close()
	if err != nil {
		return err
	}

	// Create a new MySQL database object.
	p.dbconnection, err = mysql.New(connInfo)
	if err != nil {
		return err
	}

	// Migrate database.
	r := rove.NewChangesetMigration(p.dbconnection, p.dbinfo.Changeset)
	r.Verbose = p.dbinfo.Verbose
	r.Checksum = p.dbinfo.ChecksumMode
	return r.Migrate(0)
}

// Disable handles any plugin cleanup tasks.
func (p *Plugin) Disable() error {
	if p.dbconnection != nil {
		return p.dbconnection.DB.Close()
	}

	return nil
}
