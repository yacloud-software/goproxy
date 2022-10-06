package db

/*
 This file was created by mkdb-client.
 The intention is not to modify thils file, but you may extend the struct DBCachedModule
 in a seperate file (so that you can regenerate this one from time to time)
*/

/*
 PRIMARY KEY: ID
*/

/*
 postgres:
 create sequence cachedmodule_seq;

Main Table:

 CREATE TABLE cachedmodule (id integer primary key default nextval('cachedmodule_seq'),path text not null  ,version text not null  ,suffix text not null  ,key text not null  ,created integer not null  ,lastused integer not null  ,tobedeleted boolean not null  );

Alter statements:
ALTER TABLE cachedmodule ADD COLUMN path text not null default '';
ALTER TABLE cachedmodule ADD COLUMN version text not null default '';
ALTER TABLE cachedmodule ADD COLUMN suffix text not null default '';
ALTER TABLE cachedmodule ADD COLUMN key text not null default '';
ALTER TABLE cachedmodule ADD COLUMN created integer not null default 0;
ALTER TABLE cachedmodule ADD COLUMN lastused integer not null default 0;
ALTER TABLE cachedmodule ADD COLUMN tobedeleted boolean not null default false;


Archive Table: (structs can be moved from main to archive using Archive() function)

 CREATE TABLE cachedmodule_archive (id integer unique not null,path text not null,version text not null,suffix text not null,key text not null,created integer not null,lastused integer not null,tobedeleted boolean not null);
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
	default_def_DBCachedModule *DBCachedModule
)

type DBCachedModule struct {
	DB                  *sql.DB
	SQLTablename        string
	SQLArchivetablename string
}

func DefaultDBCachedModule() *DBCachedModule {
	if default_def_DBCachedModule != nil {
		return default_def_DBCachedModule
	}
	psql, err := sql.Open()
	if err != nil {
		fmt.Printf("Failed to open database: %s\n", err)
		os.Exit(10)
	}
	res := NewDBCachedModule(psql)
	ctx := context.Background()
	err = res.CreateTable(ctx)
	if err != nil {
		fmt.Printf("Failed to create table: %s\n", err)
		os.Exit(10)
	}
	default_def_DBCachedModule = res
	return res
}
func NewDBCachedModule(db *sql.DB) *DBCachedModule {
	foo := DBCachedModule{DB: db}
	foo.SQLTablename = "cachedmodule"
	foo.SQLArchivetablename = "cachedmodule_archive"
	return &foo
}

// archive. It is NOT transactionally save.
func (a *DBCachedModule) Archive(ctx context.Context, id uint64) error {

	// load it
	p, err := a.ByID(ctx, id)
	if err != nil {
		return err
	}

	// now save it to archive:
	_, e := a.DB.ExecContext(ctx, "archive_DBCachedModule", "insert into "+a.SQLArchivetablename+" (id,path, version, suffix, key, created, lastused, tobedeleted) values ($1,$2, $3, $4, $5, $6, $7, $8) ", p.ID, p.Path, p.Version, p.Suffix, p.Key, p.Created, p.LastUsed, p.ToBeDeleted)
	if e != nil {
		return e
	}

	// now delete it.
	a.DeleteByID(ctx, id)
	return nil
}

// Save (and use database default ID generation)
func (a *DBCachedModule) Save(ctx context.Context, p *savepb.CachedModule) (uint64, error) {
	qn := "DBCachedModule_Save"
	rows, e := a.DB.QueryContext(ctx, qn, "insert into "+a.SQLTablename+" (path, version, suffix, key, created, lastused, tobedeleted) values ($1, $2, $3, $4, $5, $6, $7) returning id", p.Path, p.Version, p.Suffix, p.Key, p.Created, p.LastUsed, p.ToBeDeleted)
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
func (a *DBCachedModule) SaveWithID(ctx context.Context, p *savepb.CachedModule) error {
	qn := "insert_DBCachedModule"
	_, e := a.DB.ExecContext(ctx, qn, "insert into "+a.SQLTablename+" (id,path, version, suffix, key, created, lastused, tobedeleted) values ($1,$2, $3, $4, $5, $6, $7, $8) ", p.ID, p.Path, p.Version, p.Suffix, p.Key, p.Created, p.LastUsed, p.ToBeDeleted)
	return a.Error(ctx, qn, e)
}

func (a *DBCachedModule) Update(ctx context.Context, p *savepb.CachedModule) error {
	qn := "DBCachedModule_Update"
	_, e := a.DB.ExecContext(ctx, qn, "update "+a.SQLTablename+" set path=$1, version=$2, suffix=$3, key=$4, created=$5, lastused=$6, tobedeleted=$7 where id = $8", p.Path, p.Version, p.Suffix, p.Key, p.Created, p.LastUsed, p.ToBeDeleted, p.ID)

	return a.Error(ctx, qn, e)
}

// delete by id field
func (a *DBCachedModule) DeleteByID(ctx context.Context, p uint64) error {
	qn := "deleteDBCachedModule_ByID"
	_, e := a.DB.ExecContext(ctx, qn, "delete from "+a.SQLTablename+" where id = $1", p)
	return a.Error(ctx, qn, e)
}

// get it by primary id
func (a *DBCachedModule) ByID(ctx context.Context, p uint64) (*savepb.CachedModule, error) {
	qn := "DBCachedModule_ByID"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,path, version, suffix, key, created, lastused, tobedeleted from "+a.SQLTablename+" where id = $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByID: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByID: error scanning (%s)", e))
	}
	if len(l) == 0 {
		return nil, a.Error(ctx, qn, fmt.Errorf("No CachedModule with id %v", p))
	}
	if len(l) != 1 {
		return nil, a.Error(ctx, qn, fmt.Errorf("Multiple (%d) CachedModule with id %v", len(l), p))
	}
	return l[0], nil
}

// get it by primary id (nil if no such ID row, but no error either)
func (a *DBCachedModule) TryByID(ctx context.Context, p uint64) (*savepb.CachedModule, error) {
	qn := "DBCachedModule_TryByID"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,path, version, suffix, key, created, lastused, tobedeleted from "+a.SQLTablename+" where id = $1", p)
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
		return nil, a.Error(ctx, qn, fmt.Errorf("Multiple (%d) CachedModule with id %v", len(l), p))
	}
	return l[0], nil
}

// get all rows
func (a *DBCachedModule) All(ctx context.Context) ([]*savepb.CachedModule, error) {
	qn := "DBCachedModule_all"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,path, version, suffix, key, created, lastused, tobedeleted from "+a.SQLTablename+" order by id")
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

// get all "DBCachedModule" rows with matching Path
func (a *DBCachedModule) ByPath(ctx context.Context, p string) ([]*savepb.CachedModule, error) {
	qn := "DBCachedModule_ByPath"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,path, version, suffix, key, created, lastused, tobedeleted from "+a.SQLTablename+" where path = $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByPath: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByPath: error scanning (%s)", e))
	}
	return l, nil
}

// the 'like' lookup
func (a *DBCachedModule) ByLikePath(ctx context.Context, p string) ([]*savepb.CachedModule, error) {
	qn := "DBCachedModule_ByLikePath"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,path, version, suffix, key, created, lastused, tobedeleted from "+a.SQLTablename+" where path ilike $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByPath: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByPath: error scanning (%s)", e))
	}
	return l, nil
}

// get all "DBCachedModule" rows with matching Version
func (a *DBCachedModule) ByVersion(ctx context.Context, p string) ([]*savepb.CachedModule, error) {
	qn := "DBCachedModule_ByVersion"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,path, version, suffix, key, created, lastused, tobedeleted from "+a.SQLTablename+" where version = $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByVersion: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByVersion: error scanning (%s)", e))
	}
	return l, nil
}

// the 'like' lookup
func (a *DBCachedModule) ByLikeVersion(ctx context.Context, p string) ([]*savepb.CachedModule, error) {
	qn := "DBCachedModule_ByLikeVersion"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,path, version, suffix, key, created, lastused, tobedeleted from "+a.SQLTablename+" where version ilike $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByVersion: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByVersion: error scanning (%s)", e))
	}
	return l, nil
}

// get all "DBCachedModule" rows with matching Suffix
func (a *DBCachedModule) BySuffix(ctx context.Context, p string) ([]*savepb.CachedModule, error) {
	qn := "DBCachedModule_BySuffix"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,path, version, suffix, key, created, lastused, tobedeleted from "+a.SQLTablename+" where suffix = $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("BySuffix: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("BySuffix: error scanning (%s)", e))
	}
	return l, nil
}

// the 'like' lookup
func (a *DBCachedModule) ByLikeSuffix(ctx context.Context, p string) ([]*savepb.CachedModule, error) {
	qn := "DBCachedModule_ByLikeSuffix"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,path, version, suffix, key, created, lastused, tobedeleted from "+a.SQLTablename+" where suffix ilike $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("BySuffix: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("BySuffix: error scanning (%s)", e))
	}
	return l, nil
}

// get all "DBCachedModule" rows with matching Key
func (a *DBCachedModule) ByKey(ctx context.Context, p string) ([]*savepb.CachedModule, error) {
	qn := "DBCachedModule_ByKey"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,path, version, suffix, key, created, lastused, tobedeleted from "+a.SQLTablename+" where key = $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByKey: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByKey: error scanning (%s)", e))
	}
	return l, nil
}

// the 'like' lookup
func (a *DBCachedModule) ByLikeKey(ctx context.Context, p string) ([]*savepb.CachedModule, error) {
	qn := "DBCachedModule_ByLikeKey"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,path, version, suffix, key, created, lastused, tobedeleted from "+a.SQLTablename+" where key ilike $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByKey: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByKey: error scanning (%s)", e))
	}
	return l, nil
}

// get all "DBCachedModule" rows with matching Created
func (a *DBCachedModule) ByCreated(ctx context.Context, p uint32) ([]*savepb.CachedModule, error) {
	qn := "DBCachedModule_ByCreated"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,path, version, suffix, key, created, lastused, tobedeleted from "+a.SQLTablename+" where created = $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByCreated: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByCreated: error scanning (%s)", e))
	}
	return l, nil
}

// the 'like' lookup
func (a *DBCachedModule) ByLikeCreated(ctx context.Context, p uint32) ([]*savepb.CachedModule, error) {
	qn := "DBCachedModule_ByLikeCreated"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,path, version, suffix, key, created, lastused, tobedeleted from "+a.SQLTablename+" where created ilike $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByCreated: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByCreated: error scanning (%s)", e))
	}
	return l, nil
}

// get all "DBCachedModule" rows with matching LastUsed
func (a *DBCachedModule) ByLastUsed(ctx context.Context, p uint32) ([]*savepb.CachedModule, error) {
	qn := "DBCachedModule_ByLastUsed"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,path, version, suffix, key, created, lastused, tobedeleted from "+a.SQLTablename+" where lastused = $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByLastUsed: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByLastUsed: error scanning (%s)", e))
	}
	return l, nil
}

// the 'like' lookup
func (a *DBCachedModule) ByLikeLastUsed(ctx context.Context, p uint32) ([]*savepb.CachedModule, error) {
	qn := "DBCachedModule_ByLikeLastUsed"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,path, version, suffix, key, created, lastused, tobedeleted from "+a.SQLTablename+" where lastused ilike $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByLastUsed: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByLastUsed: error scanning (%s)", e))
	}
	return l, nil
}

// get all "DBCachedModule" rows with matching ToBeDeleted
func (a *DBCachedModule) ByToBeDeleted(ctx context.Context, p bool) ([]*savepb.CachedModule, error) {
	qn := "DBCachedModule_ByToBeDeleted"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,path, version, suffix, key, created, lastused, tobedeleted from "+a.SQLTablename+" where tobedeleted = $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByToBeDeleted: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByToBeDeleted: error scanning (%s)", e))
	}
	return l, nil
}

// the 'like' lookup
func (a *DBCachedModule) ByLikeToBeDeleted(ctx context.Context, p bool) ([]*savepb.CachedModule, error) {
	qn := "DBCachedModule_ByLikeToBeDeleted"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,path, version, suffix, key, created, lastused, tobedeleted from "+a.SQLTablename+" where tobedeleted ilike $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByToBeDeleted: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByToBeDeleted: error scanning (%s)", e))
	}
	return l, nil
}

/**********************************************************************
* Helper to convert from an SQL Query
**********************************************************************/

// from a query snippet (the part after WHERE)
func (a *DBCachedModule) FromQuery(ctx context.Context, query_where string, args ...interface{}) ([]*savepb.CachedModule, error) {
	rows, err := a.DB.QueryContext(ctx, "custom_query_"+a.Tablename(), "select "+a.SelectCols()+" from "+a.Tablename()+" where "+query_where, args...)
	if err != nil {
		return nil, err
	}
	return a.FromRows(ctx, rows)
}

/**********************************************************************
* Helper to convert from an SQL Row to struct
**********************************************************************/
func (a *DBCachedModule) Tablename() string {
	return a.SQLTablename
}

func (a *DBCachedModule) SelectCols() string {
	return "id,path, version, suffix, key, created, lastused, tobedeleted"
}
func (a *DBCachedModule) SelectColsQualified() string {
	return "" + a.SQLTablename + ".id," + a.SQLTablename + ".path, " + a.SQLTablename + ".version, " + a.SQLTablename + ".suffix, " + a.SQLTablename + ".key, " + a.SQLTablename + ".created, " + a.SQLTablename + ".lastused, " + a.SQLTablename + ".tobedeleted"
}

func (a *DBCachedModule) FromRows(ctx context.Context, rows *gosql.Rows) ([]*savepb.CachedModule, error) {
	var res []*savepb.CachedModule
	for rows.Next() {
		foo := savepb.CachedModule{}
		err := rows.Scan(&foo.ID, &foo.Path, &foo.Version, &foo.Suffix, &foo.Key, &foo.Created, &foo.LastUsed, &foo.ToBeDeleted)
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
func (a *DBCachedModule) CreateTable(ctx context.Context) error {
	csql := []string{
		`create sequence if not exists ` + a.SQLTablename + `_seq;`,
		`CREATE TABLE if not exists ` + a.SQLTablename + ` (id integer primary key default nextval('` + a.SQLTablename + `_seq'),path text not null  ,version text not null  ,suffix text not null  ,key text not null  ,created integer not null  ,lastused integer not null  ,tobedeleted boolean not null  );`,
		`CREATE TABLE if not exists ` + a.SQLTablename + `_archive (id integer primary key default nextval('` + a.SQLTablename + `_seq'),path text not null  ,version text not null  ,suffix text not null  ,key text not null  ,created integer not null  ,lastused integer not null  ,tobedeleted boolean not null  );`,
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
func (a *DBCachedModule) Error(ctx context.Context, q string, e error) error {
	if e == nil {
		return nil
	}
	return fmt.Errorf("[table="+a.SQLTablename+", query=%s] Error: %s", q, e)
}
