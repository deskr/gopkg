package rdb

import (
	"time"

	"github.com/dancannon/gorethink"
)

// OpenConnection initializes the rethink db
func OpenConnection(address string, database string) (c *gorethink.Session, err error) {

	gorethink.SetTags("gorethink", "json")

	for i := 0; i < 10; i++ {
		c, err = gorethink.Connect(gorethink.ConnectOpts{
			Address:  address,
			Database: database,
			MaxIdle:  10,
			MaxOpen:  10,
		})
		if err == nil {
			break
		} else {
			time.Sleep(time.Second * 2)
		}
	}

	if err != nil {
		return
	}

	// Create database if not already exists
	exists, err := DatabaseExists(database, c)
	if err != nil {
		return
	}

	if !exists {
		err = gorethink.DBCreate(database).Exec(c)
		if err != nil {
			return
		}
	}

	c.Use(database)

	return
}

// DatabaseExists checks if a database exists
func DatabaseExists(name string, s *gorethink.Session) (exists bool, err error) {
	exists = false
	res, err := gorethink.DBList().Run(s)
	if err != nil {
		return
	}
	defer res.Close()
	var dbs []string
	if err = res.All(&dbs); err != nil {
		return
	}

	for _, t := range dbs {
		if t == name {
			exists = true
			break
		}
	}
	return
}

// TableExists checks if a table exists
func TableExists(name string, s *gorethink.Session) (exists bool, err error) {
	exists = false
	res, err := gorethink.TableList().Run(s)
	if err != nil {
		return
	}
	defer res.Close()
	var tables []string
	if err = res.All(&tables); err != nil {
		return
	}
	for _, t := range tables {
		if t == name {
			exists = true
			return
		}
	}
	return
}
