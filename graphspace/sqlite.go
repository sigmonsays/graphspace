package main

import (
	"database/sql"
	"fmt"
	"io"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const schema = `
create table if not exists graphs (
	id text not null primary key, 
	created int, 
	format text, 
	text text,
	width int, 
	height int, 
	output text,
	description text
);
`

type sqlGraphviz struct {
	dbpath string
	db     *sql.DB
	stmt   map[string]*sql.Stmt
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

func NewSqlGraphviz(dbpath string) (*sqlGraphviz, error) {
	err := PrepareSqliteDb(dbpath)
	if err != nil {
		return nil, fmt.Errorf("prepare %s: %s", dbpath, err)
	}

	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		return nil, err
	}

	stmt := make(map[string]*sql.Stmt, 0)
	queries := map[string]string{
		"insert":      "insert into graphs (id, created, format, text, width, height, output, description) values(?, ?, ?, ?, ?, ?, ?, ?)",
		"select":      "select format, text, width, height, output, description from graphs where id = ?",
		"delete":      "delete from graphs where id = ?",
		"list_recent": "select id, description from graphs order by created desc limit ?",
	}
	for name, sql := range queries {
		st, err := db.Prepare(sql)
		if err != nil {
			return nil, err
		}
		stmt[name] = st
	}

	if err != nil {
		return nil, err
	}

	q := &sqlGraphviz{
		db:     db,
		dbpath: dbpath,
		stmt:   stmt,
	}
	return q, nil
}

func (q *sqlGraphviz) Create(g *Graph) (string, error) {

	id := g.GetId()
	log.Tracef("graph id=%s", id)

	t := time.Now().Unix()

	_, err := q.stmt["insert"].Exec(id, t, g.Format, g.Text, g.Width, g.Height, g.Output, g.Description)
	if err != nil {
		return "", err
	}
	log.Tracef("create id %s from graph %d bytes", id, len(g.Text))
	return id, nil
}

func (q *sqlGraphviz) Get(id string) (*Graph, error) {
	row := q.stmt["select"].QueryRow(id)
	if row == nil {
		return nil, io.EOF
	}

	g := &Graph{}

	err := row.Scan(&g.Format, &g.Text, &g.Width, &g.Height, &g.Output, &g.Description)
	if err == sql.ErrNoRows {
		return nil, io.EOF
	} else if err != nil {
		return nil, err
	}

	return g, nil
}

func (q *sqlGraphviz) ListRecent(max int) ([]listEntry, error) {
	ls := make([]listEntry, 0)
	rows, err := q.stmt["list_recent"].Query(max)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		e := listEntry{}
		err := rows.Scan(&e.Id, &e.Description)
		if err != nil {
			log.Warnf("scan %s", err)
			continue
		}
		ls = append(ls, e)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return ls, nil

}

func (q *sqlGraphviz) Delete(id string) error {
	_, err := q.stmt["delete"].Exec(id)
	if err != nil {
		return err
	}
	return err
}

func (q *sqlGraphviz) Close() error {
	return q.db.Close()
}
