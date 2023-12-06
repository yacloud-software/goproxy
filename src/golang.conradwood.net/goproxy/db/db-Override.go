package db

/*
 This file was created by mkdb-client.
 The intention is not to modify thils file, but you may extend the struct DBOverride
 in a seperate file (so that you can regenerate this one from time to time)
*/

/*
 PRIMARY KEY: ID
*/

/*
 postgres:
 create sequence override_seq;

Main Table:

 CREATE TABLE override (id integer primary key default nextval('override_seq'),package text not null  ,list bytea not null  );

Alter statements:
ALTER TABLE override ADD COLUMN IF NOT EXISTS package text not null default '';
ALTER TABLE override ADD COLUMN IF NOT EXISTS list bytea not null default 0;


Archive Table: (structs can be moved from main to archive using Archive() function)

 CREATE TABLE override_archive (id integer unique not null,package text not null,list bytea not null);
*/

import (
	"context"
	gosql "database/sql"
	"fmt"
	savepb "golang.conradwood.net/apis/goproxy"
	"golang.conradwood.net/go-easyops/sql"
	"os"
)

var (
	default_def_DBOverride *DBOverride
)

type DBOverride struct {
	DB                  *sql.DB
	SQLTablename        string
	SQLArchivetablename string
}

func DefaultDBOverride() *DBOverride {
	if default_def_DBOverride != nil {
		return default_def_DBOverride
	}
	psql, err := sql.Open()
	if err != nil {
		fmt.Printf("Failed to open database: %s\n", err)
		os.Exit(10)
	}
	res := NewDBOverride(psql)
	ctx := context.Background()
	err = res.CreateTable(ctx)
	if err != nil {
		fmt.Printf("Failed to create table: %s\n", err)
		os.Exit(10)
	}
	default_def_DBOverride = res
	return res
}
func NewDBOverride(db *sql.DB) *DBOverride {
	foo := DBOverride{DB: db}
	foo.SQLTablename = "override"
	foo.SQLArchivetablename = "override_archive"
	return &foo
}

// archive. It is NOT transactionally save.
func (a *DBOverride) Archive(ctx context.Context, id uint64) error {

	// load it
	p, err := a.ByID(ctx, id)
	if err != nil {
		return err
	}

	// now save it to archive:
	_, e := a.DB.ExecContext(ctx, "archive_DBOverride", "insert into "+a.SQLArchivetablename+" (id,package, list) values ($1,$2, $3) ", p.ID, p.Package, p.List)
	if e != nil {
		return e
	}

	// now delete it.
	a.DeleteByID(ctx, id)
	return nil
}

// Save (and use database default ID generation)
func (a *DBOverride) Save(ctx context.Context, p *savepb.Override) (uint64, error) {
	qn := "DBOverride_Save"
	rows, e := a.DB.QueryContext(ctx, qn, "insert into "+a.SQLTablename+" (package, list) values ($1, $2) returning id", p.Package, p.List)
	if e != nil {
		return 0, a.Error(ctx, qn, e)
	}
	defer rows.Close()
	if !rows.Next() {
		return 0, a.Error(ctx, qn, fmt.Errorf("No rows after insert"))
	}
	var id uint64
	e = rows.Scan(&id)
	if e != nil {
		return 0, a.Error(ctx, qn, fmt.Errorf("failed to scan id after insert: %s", e))
	}
	p.ID = id
	return id, nil
}

// Save using the ID specified
func (a *DBOverride) SaveWithID(ctx context.Context, p *savepb.Override) error {
	qn := "insert_DBOverride"
	_, e := a.DB.ExecContext(ctx, qn, "insert into "+a.SQLTablename+" (id,package, list) values ($1,$2, $3) ", p.ID, p.Package, p.List)
	return a.Error(ctx, qn, e)
}

func (a *DBOverride) Update(ctx context.Context, p *savepb.Override) error {
	qn := "DBOverride_Update"
	_, e := a.DB.ExecContext(ctx, qn, "update "+a.SQLTablename+" set package=$1, list=$2 where id = $3", p.Package, p.List, p.ID)

	return a.Error(ctx, qn, e)
}

// delete by id field
func (a *DBOverride) DeleteByID(ctx context.Context, p uint64) error {
	qn := "deleteDBOverride_ByID"
	_, e := a.DB.ExecContext(ctx, qn, "delete from "+a.SQLTablename+" where id = $1", p)
	return a.Error(ctx, qn, e)
}

// get it by primary id
func (a *DBOverride) ByID(ctx context.Context, p uint64) (*savepb.Override, error) {
	qn := "DBOverride_ByID"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,package, list from "+a.SQLTablename+" where id = $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByID: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByID: error scanning (%s)", e))
	}
	if len(l) == 0 {
		return nil, a.Error(ctx, qn, fmt.Errorf("No Override with id %v", p))
	}
	if len(l) != 1 {
		return nil, a.Error(ctx, qn, fmt.Errorf("Multiple (%d) Override with id %v", len(l), p))
	}
	return l[0], nil
}

// get it by primary id (nil if no such ID row, but no error either)
func (a *DBOverride) TryByID(ctx context.Context, p uint64) (*savepb.Override, error) {
	qn := "DBOverride_TryByID"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,package, list from "+a.SQLTablename+" where id = $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("TryByID: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("TryByID: error scanning (%s)", e))
	}
	if len(l) == 0 {
		return nil, nil
	}
	if len(l) != 1 {
		return nil, a.Error(ctx, qn, fmt.Errorf("Multiple (%d) Override with id %v", len(l), p))
	}
	return l[0], nil
}

// get all rows
func (a *DBOverride) All(ctx context.Context) ([]*savepb.Override, error) {
	qn := "DBOverride_all"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,package, list from "+a.SQLTablename+" order by id")
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("All: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, fmt.Errorf("All: error scanning (%s)", e)
	}
	return l, nil
}

/**********************************************************************
* GetBy[FIELD] functions
**********************************************************************/

// get all "DBOverride" rows with matching Package
func (a *DBOverride) ByPackage(ctx context.Context, p string) ([]*savepb.Override, error) {
	qn := "DBOverride_ByPackage"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,package, list from "+a.SQLTablename+" where package = $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByPackage: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByPackage: error scanning (%s)", e))
	}
	return l, nil
}

// the 'like' lookup
func (a *DBOverride) ByLikePackage(ctx context.Context, p string) ([]*savepb.Override, error) {
	qn := "DBOverride_ByLikePackage"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,package, list from "+a.SQLTablename+" where package ilike $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByPackage: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByPackage: error scanning (%s)", e))
	}
	return l, nil
}

// get all "DBOverride" rows with matching List
func (a *DBOverride) ByList(ctx context.Context, p []byte) ([]*savepb.Override, error) {
	qn := "DBOverride_ByList"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,package, list from "+a.SQLTablename+" where list = $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByList: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByList: error scanning (%s)", e))
	}
	return l, nil
}

// the 'like' lookup
func (a *DBOverride) ByLikeList(ctx context.Context, p []byte) ([]*savepb.Override, error) {
	qn := "DBOverride_ByLikeList"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,package, list from "+a.SQLTablename+" where list ilike $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByList: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByList: error scanning (%s)", e))
	}
	return l, nil
}

/**********************************************************************
* Helper to convert from an SQL Query
**********************************************************************/

// from a query snippet (the part after WHERE)
func (a *DBOverride) FromQuery(ctx context.Context, query_where string, args ...interface{}) ([]*savepb.Override, error) {
	rows, err := a.DB.QueryContext(ctx, "custom_query_"+a.Tablename(), "select "+a.SelectCols()+" from "+a.Tablename()+" where "+query_where, args...)
	if err != nil {
		return nil, err
	}
	return a.FromRows(ctx, rows)
}

/**********************************************************************
* Helper to convert from an SQL Row to struct
**********************************************************************/
func (a *DBOverride) Tablename() string {
	return a.SQLTablename
}

func (a *DBOverride) SelectCols() string {
	return "id,package, list"
}
func (a *DBOverride) SelectColsQualified() string {
	return "" + a.SQLTablename + ".id," + a.SQLTablename + ".package, " + a.SQLTablename + ".list"
}

func (a *DBOverride) FromRows(ctx context.Context, rows *gosql.Rows) ([]*savepb.Override, error) {
	var res []*savepb.Override
	for rows.Next() {
		foo := savepb.Override{}
		err := rows.Scan(&foo.ID, &foo.Package, &foo.List)
		if err != nil {
			return nil, a.Error(ctx, "fromrow-scan", err)
		}
		res = append(res, &foo)
	}
	return res, nil
}

/**********************************************************************
* Helper to create table and columns
**********************************************************************/
func (a *DBOverride) CreateTable(ctx context.Context) error {
	csql := []string{
		`create sequence if not exists ` + a.SQLTablename + `_seq;`,
		`CREATE TABLE if not exists ` + a.SQLTablename + ` (id integer primary key default nextval('` + a.SQLTablename + `_seq'),package text not null  ,list bytea not null  );`,
		`CREATE TABLE if not exists ` + a.SQLTablename + `_archive (id integer primary key default nextval('` + a.SQLTablename + `_seq'),package text not null  ,list bytea not null  );`,
		`ALTER TABLE override ADD COLUMN IF NOT EXISTS package text not null default '';`,
		`ALTER TABLE override ADD COLUMN IF NOT EXISTS list bytea not null default 0;`,
	}
	for i, c := range csql {
		_, e := a.DB.ExecContext(ctx, fmt.Sprintf("create_"+a.SQLTablename+"_%d", i), c)
		if e != nil {
			return e
		}
	}
	return nil
}

/**********************************************************************
* Helper to meaningful errors
**********************************************************************/
func (a *DBOverride) Error(ctx context.Context, q string, e error) error {
	if e == nil {
		return nil
	}
	return fmt.Errorf("[table="+a.SQLTablename+", query=%s] Error: %s", q, e)
}

