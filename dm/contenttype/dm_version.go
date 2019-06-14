// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package contenttype

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"dm/dm/db"
	. "dm/dm/db"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// Version is an object representing the database table.
// Implement dm.contenttype.ContentTyper interface
type Version struct {
	ID          int    `boil:"id" json:"id" toml:"id" yaml:"id"`
	ContentType string `boil:"content_type" json:"content_type" toml:"content_type" yaml:"content_type"`
	ContentID   int    `boil:"content_id" json:"content_id" toml:"content_id" yaml:"content_id"`
	Version     int    `boil:"version" json:"version" toml:"version" yaml:"version"`
	Status      int8   `boil:"status" json:"status" toml:"status" yaml:"status"`
	Author      int    `boil:"author" json:"author" toml:"author" yaml:"author"`
	Data        string `boil:"data" json:"data" toml:"data" yaml:"data"`
	LocationID  int    `boil:"location_id" json:"location_id" toml:"location_id" yaml:"location_id"`
	Created     int    `boil:"created" json:"created" toml:"created" yaml:"created"`

	R        *versionR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L        versionL  `boil:"-" json:"-" toml:"-" yaml:"-"`
	Location `boil:"dm_location,bind"`
}

func (c *Version) ToMap() map[string]interface{} {
	result := make(map[string]interface{})
	result["id"] = c.ID
	result["content_type"] = c.ContentType
	result["content_id"] = c.ContentID
	result["version"] = c.Version
	result["status"] = c.Status
	result["author"] = c.Author
	result["data"] = c.Data
	result["location_id"] = c.LocationID
	result["created"] = c.Created
	return result
}

func (c *Version) TableName() string {
	return "dm_version"
}

func (c *Version) Field(name string) interface{} {
	var result interface{}
	switch name {
	case "id", "ID":
		result = c.ID
	case "content_type", "ContentType":
		result = c.ContentType
	case "content_id", "ContentID":
		result = c.ContentID
	case "version", "Version":
		result = c.Version
	case "status", "Status":
		result = c.Status
	case "author", "Author":
		result = c.Author
	case "data", "Data":
		result = c.Data
	case "location_id", "LocationID":
		result = c.LocationID
	case "created", "Created":
		result = c.Created
	default:
	}
	return result
}

func (c Version) Store(transaction ...*sql.Tx) error {
	handler := db.DBHanlder()
	if c.ID == 0 {
		id, err := handler.Insert(c.TableName(), c.ToMap(), transaction...)
		c.ID = id
		if err != nil {
			return err
		}
	} else {
		err := handler.Update(c.TableName(), c.ToMap(), Cond("id", c.ID), transaction...)
		return err
	}
	return nil
}

var VersionColumns = struct {
	ID          string
	ContentType string
	ContentID   string
	Version     string
	Status      string
	Author      string
	Data        string
	LocationID  string
	Created     string
}{
	ID:          "id",
	ContentType: "content_type",
	ContentID:   "content_id",
	Version:     "version",
	Status:      "status",
	Author:      "author",
	Data:        "data",
	LocationID:  "location_id",
	Created:     "created",
}

// VersionRels is where relationship names are stored.
var VersionRels = struct {
}{}

// versionR is where relationships are stored.
type versionR struct {
}

// NewStruct creates a new relationship struct
func (*versionR) NewStruct() *versionR {
	return &versionR{}
}

// versionL is where Load methods for each relationship are stored.
type versionL struct{}

var (
	versionColumns               = []string{"id", "content_type", "content_id", "version", "status", "author", "data", "location_id", "created"}
	versionColumnsWithoutDefault = []string{"content_type", "content_id", "version", "data", "created"}
	versionColumnsWithDefault    = []string{"id", "status", "author", "location_id"}
	versionPrimaryKeyColumns     = []string{"id"}
)

type (
	// VersionSlice is an alias for a slice of pointers to Version.
	// This should generally be used opposed to []Version.
	VersionSlice []*Version
	// VersionHook is the signature for custom Version hook methods
	VersionHook func(context.Context, boil.ContextExecutor, *Version) error

	versionQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	versionType                 = reflect.TypeOf(&Version{})
	versionMapping              = queries.MakeStructMapping(versionType)
	versionPrimaryKeyMapping, _ = queries.BindMapping(versionType, versionMapping, versionPrimaryKeyColumns)
	versionInsertCacheMut       sync.RWMutex
	versionInsertCache          = make(map[string]insertCache)
	versionUpdateCacheMut       sync.RWMutex
	versionUpdateCache          = make(map[string]updateCache)
	versionUpsertCacheMut       sync.RWMutex
	versionUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var versionBeforeInsertHooks []VersionHook
var versionBeforeUpdateHooks []VersionHook
var versionBeforeDeleteHooks []VersionHook
var versionBeforeUpsertHooks []VersionHook

var versionAfterInsertHooks []VersionHook
var versionAfterSelectHooks []VersionHook
var versionAfterUpdateHooks []VersionHook
var versionAfterDeleteHooks []VersionHook
var versionAfterUpsertHooks []VersionHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Version) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range versionBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Version) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range versionBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Version) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range versionBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Version) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range versionBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Version) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range versionAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Version) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range versionAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Version) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range versionAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Version) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range versionAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Version) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range versionAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddVersionHook registers your hook function for all future operations.
func AddVersionHook(hookPoint boil.HookPoint, versionHook VersionHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		versionBeforeInsertHooks = append(versionBeforeInsertHooks, versionHook)
	case boil.BeforeUpdateHook:
		versionBeforeUpdateHooks = append(versionBeforeUpdateHooks, versionHook)
	case boil.BeforeDeleteHook:
		versionBeforeDeleteHooks = append(versionBeforeDeleteHooks, versionHook)
	case boil.BeforeUpsertHook:
		versionBeforeUpsertHooks = append(versionBeforeUpsertHooks, versionHook)
	case boil.AfterInsertHook:
		versionAfterInsertHooks = append(versionAfterInsertHooks, versionHook)
	case boil.AfterSelectHook:
		versionAfterSelectHooks = append(versionAfterSelectHooks, versionHook)
	case boil.AfterUpdateHook:
		versionAfterUpdateHooks = append(versionAfterUpdateHooks, versionHook)
	case boil.AfterDeleteHook:
		versionAfterDeleteHooks = append(versionAfterDeleteHooks, versionHook)
	case boil.AfterUpsertHook:
		versionAfterUpsertHooks = append(versionAfterUpsertHooks, versionHook)
	}
}

// One returns a single version record from the query.
func (q versionQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Version, error) {
	o := &Version{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "contenttype: failed to execute a one query for dm_version")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Version records from the query.
func (q versionQuery) All(ctx context.Context, exec boil.ContextExecutor) (VersionSlice, error) {
	var o []*Version

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "contenttype: failed to assign all query results to Version slice")
	}

	if len(versionAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Version records in the query.
func (q versionQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "contenttype: failed to count dm_version rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q versionQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "contenttype: failed to check if dm_version exists")
	}

	return count > 0, nil
}

var mySQLVersionUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Version) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("contenttype: no dm_version provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(versionColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLVersionUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	versionUpsertCacheMut.RLock()
	cache, cached := versionUpsertCache[key]
	versionUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			versionColumns,
			versionColumnsWithDefault,
			versionColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			versionColumns,
			versionPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("contenttype: unable to upsert dm_version, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "dm_version", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `dm_version` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(versionType, versionMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(versionType, versionMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "contenttype: unable to upsert for dm_version")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == versionMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(versionType, versionMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "contenttype: unable to retrieve unique values for dm_version")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "contenttype: unable to populate default values for dm_version")
	}

CacheNoHooks:
	if !cached {
		versionUpsertCacheMut.Lock()
		versionUpsertCache[key] = cache
		versionUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}
