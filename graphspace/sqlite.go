package main

import (
	"database/sql"
	"fmt"
	"io"

	_ "github.com/mattn/go-sqlite3"
)

const schema = `
create table if not exists graphs(id integer not null primary key, graph_string text);
`

type sqlGraphviz struct {
	dbpath      string
	db          *sql.DB
	stmt_insert *sql.Stmt
	stmt_select *sql.Stmt
	stmt_delete *sql.Stmt
}

// prepare the database for usage
func PrepareSqliteDb(dbpath string) error {

	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		return err
	}

	_, err = db.Exec(schema)
	if err != nil {
		return err
	}

	return nil
}

func NewsqlGraphviz(dbpath string) (*sqlGraphviz, error) {
	err := PrepareSqliteDb(dbpath)
	if err != nil {
		return nil, fmt.Errorf("prepare %s: %s", dbpath, err)
	}

	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		return nil, err
	}

	stmt_insert, err := db.Prepare("insert into graphs (graph_string) values(?)")
	if err != nil {
		return nil, err
	}

	stmt_select, err := db.Prepare("select graph_string from graphs where id = ?")
	if err != nil {
		return nil, err
	}

	stmt_delete, err := db.Prepare("delete from graphs where id = ?")
	if err != nil {
		return nil, err
	}

	q := &sqlGraphviz{
		db:          db,
		dbpath:      dbpath,
		stmt_insert: stmt_insert,
		stmt_select: stmt_select,
		stmt_delete: stmt_delete,
	}
	return q, nil
}

func (q *sqlGraphviz) Create(graph string) error {
	log.Tracef("create %s", graph)
	_, err := q.stmt_insert.Exec(graph)
	if err != nil {
		return err
	}
	return nil
}

func (q *sqlGraphviz) Get(id int) (string, error) {
	row := q.stmt_select.QueryRow(id)
	if row == nil {
		return "", io.EOF
	}

	var graph string
	err := row.Scan( &graph)
	if err == sql.ErrNoRows {
		return "", io.EOF
	} else if err != nil {
		return "", err
	}

	return graph, nil
}

func (q *sqlGraphviz) Delete(id int) error {
	_, err := q.stmt_delete.Exec(id)
	if err != nil {
		return err
	}
	return err
}

func (q *sqlGraphviz) Close() error {
	return q.db.Close()
}
