package rdb

import "gopkg.in/dancannon/gorethink.v2"

// OpenSession initializes the rethink db session
func OpenSession(options gorethink.ConnectOpts) (session *gorethink.Session, err error) {
	gorethink.SetTags("gorethink", "json")

	session, err = gorethink.Connect(options)
	if err != nil {
		return
	}

	// Create database if not already exists
	exists, err := DatabaseExists(options.Database, session)
	if err != nil {
		session.Close()
		return
	}

	if !exists {
		if err = gorethink.DBCreate(options.Database).Exec(session); err != nil {
			session.Close()
			return
		}
	}

	session.Use(options.Database)
	return
}

// DatabaseExists checks if a database exists
func DatabaseExists(name string, s *gorethink.Session) (exists bool, err error) {
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
func EnsureTable(name string, s *gorethink.Session) error {
	if exists, err := TableExists(name, s); err != nil || exists {
		return err
	}

	return gorethink.TableCreate(name).Exec(s)
}

// IndexExists checks if a table index exists
func IndexExists(table gorethink.Term, index string, s *gorethink.Session) (exists bool, err error) {
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
func EnsureIndex(table gorethink.Term, index string, s *gorethink.Session) error {
	if exists, err := IndexExists(table, index, s); err != nil || exists {
		return err
	}

	return table.IndexCreate(index).Exec(s)
}

// EnsureIndexFunc creates an index if not already exist
func EnsureIndexFunc(table gorethink.Term, index string, indexFunction interface{}, s *gorethink.Session) error {
	if exists, err := IndexExists(table, index, s); err != nil || exists {
		return err
	}

	return table.IndexCreateFunc(index, indexFunction).Exec(s)
}
