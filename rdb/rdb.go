package rdb

import (
	"fmt"
	"time"

	"gopkg.in/dancannon/gorethink.v2"
)

var sessions = make(map[string]*gorethink.Session, 0)

// OpenConnection initializes the rethink db
func OpenConnection(address string, database string) (c *gorethink.Session, err error) {
	skey := fmt.Sprintf("%s#%s", address, database)

	if s, ok := sessions[skey]; ok && s.IsConnected() {
		c = s
		return
	}

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

	sessions[skey] = c

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

// EnsureTable creates a table if it does not already exist
func EnsureTable(name string, s *gorethink.Session) (err error) {
	var exists bool
	if exists, err = TableExists(name, s); err != nil || exists {
		return
	}

	err = gorethink.TableCreate(name).Exec(s)
	return
}

// IndexExists checks if a table index exists
func IndexExists(table gorethink.Term, index string, s *gorethink.Session) (exists bool, err error) {
	exists = false
	res, err := table.IndexList().Run(s)
	if err != nil {
		return
	}
	defer res.Close()
	var indices []string
	if err = res.All(&indices); err != nil {
		return
	}
	for _, t := range indices {
		if t == index {
			exists = true
			return
		}
	}
	return
}

// EnsureIndex creates an index if not already exist
func EnsureIndex(table gorethink.Term, index string, s *gorethink.Session) (err error) {
	var exists bool
	if exists, err = IndexExists(table, index, s); err != nil || exists {
		return
	}

	err = table.IndexCreate(index).Exec(s)
	return
}

// EnsureIndexFunc creates an index if not already exist
func EnsureIndexFunc(table gorethink.Term, index string, indexFunction interface{}, s *gorethink.Session) (err error) {
	var exists bool
	if exists, err = IndexExists(table, index, s); err != nil || exists {
		return
	}

	err = table.IndexCreateFunc(index, indexFunction).Exec(s)
	return
}
