// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package entity

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"dm/db"
	"dm/fieldtype"
	. "dm/query"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// Location is an object representing the database table.
// Implement dm.model.ContentTyper interface
type Location struct {
	ID          int    `boil:"id" json:"id" toml:"id" yaml:"id"`
	ParentID    int    `boil:"parent_id" json:"parent_id" toml:"parent_id" yaml:"parent_id"`
	MainID      int    `boil:"main_id" json:"main_id" toml:"main_id" yaml:"main_id"`
	Hierarchy   string `boil:"hierarchy" json:"hierarchy" toml:"hierarchy" yaml:"hierarchy"`
	ContentType string `boil:"content_type" json:"content_type" toml:"content_type" yaml:"content_type"`
	ContentID   int    `boil:"content_id" json:"content_id" toml:"content_id" yaml:"content_id"`
	Language    string `boil:"language" json:"language" toml:"language" yaml:"language"`
	Name        string `boil:"name" json:"name" toml:"name" yaml:"name"`
	IsHidden    bool   `boil:"is_hidden" json:"is_hidden" toml:"is_hidden" yaml:"is_hidden"`
	IsInvisible bool   `boil:"is_invisible" json:"is_invisible" toml:"is_invisible" yaml:"is_invisible"`
	Priority    int    `boil:"priority" json:"priority" toml:"priority" yaml:"priority"`
	UID         string `boil:"uid" json:"uid" toml:"uid" yaml:"uid"`
	Section     string `boil:"section" json:"section" toml:"section" yaml:"section"`
	P           string `boil:"p" json:"p" toml:"p" yaml:"p"`
}

func (c *Location) Fields() map[string]fieldtype.Fielder {
	return nil
}

func (c *Location) Values() map[string]interface{} {
	result := make(map[string]interface{})
	result["id"] = c.ID
	result["parent_id"] = c.ParentID
	result["main_id"] = c.MainID
	result["hierarchy"] = c.Hierarchy
	result["content_type"] = c.ContentType
	result["content_id"] = c.ContentID
	result["language"] = c.Language
	result["name"] = c.Name
	result["is_hidden"] = c.IsHidden
	result["is_invisible"] = c.IsInvisible
	result["priority"] = c.Priority
	result["uid"] = c.UID
	result["section"] = c.Section
	result["p"] = c.P
	return result
}

func (c *Location) TableName() string {
	return "dm_location"
}

func (c *Location) Field(name string) interface{} {
	var result interface{}
	switch name {
	case "id", "ID":
		result = c.ID
	case "parent_id", "ParentID":
		result = c.ParentID
	case "main_id", "MainID":
		result = c.MainID
	case "hierarchy", "Hierarchy":
		result = c.Hierarchy
	case "content_type", "ContentType":
		result = c.ContentType
	case "content_id", "ContentID":
		result = c.ContentID
	case "language", "Language":
		result = c.Language
	case "name", "Name":
		result = c.Name
	case "is_hidden", "IsHidden":
		result = c.IsHidden
	case "is_invisible", "IsInvisible":
		result = c.IsInvisible
	case "priority", "Priority":
		result = c.Priority
	case "uid", "UID":
		result = c.UID
	case "section", "Section":
		result = c.Section
	case "p", "P":
		result = c.P
	default:
	}
	return result
}

func (c Location) Store() error {
	handler := db.DBHanlder()
	if c.ID == 0 {
		id, err := handler.Insert(c.TableName(), c.Values())
		c.ID = id
		if err != nil {
			return err
		}
	} else {
		err := handler.Update(c.TableName(), c.Values(), Cond("id", c.ID))
		return err
	}
	return nil
}

var (
	locationColumns               = []string{"id", "parent_id", "main_id", "hierarchy", "content_type", "content_id", "language", "name", "is_hidden", "is_invisible", "priority", "uid", "section", "p"}
	locationColumnsWithoutDefault = []string{"main_id", "hierarchy", "content_type", "content_id", "language", "name", "uid", "section"}
	locationColumnsWithDefault    = []string{"id", "parent_id", "is_hidden", "is_invisible", "priority", "p"}
	locationPrimaryKeyColumns     = []string{"id"}
)

type (
	// LocationSlice is an alias for a slice of pointers to Location.
	// This should generally be used opposed to []Location.
	LocationSlice []*Location
	// LocationHook is the signature for custom Location hook methods
	LocationHook func(context.Context, boil.ContextExecutor, *Location) error

	locationQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	locationType                 = reflect.TypeOf(&Location{})
	locationMapping              = queries.MakeStructMapping(locationType)
	locationPrimaryKeyMapping, _ = queries.BindMapping(locationType, locationMapping, locationPrimaryKeyColumns)
	locationInsertCacheMut       sync.RWMutex
	locationInsertCache          = make(map[string]insertCache)
	locationUpdateCacheMut       sync.RWMutex
	locationUpdateCache          = make(map[string]updateCache)
	locationUpsertCacheMut       sync.RWMutex
	locationUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var locationBeforeInsertHooks []LocationHook
var locationBeforeUpdateHooks []LocationHook
var locationBeforeDeleteHooks []LocationHook
var locationBeforeUpsertHooks []LocationHook

var locationAfterInsertHooks []LocationHook
var locationAfterSelectHooks []LocationHook
var locationAfterUpdateHooks []LocationHook
var locationAfterDeleteHooks []LocationHook
var locationAfterUpsertHooks []LocationHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Location) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range locationBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Location) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range locationBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Location) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range locationBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Location) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range locationBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Location) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range locationAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Location) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range locationAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Location) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range locationAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Location) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range locationAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Location) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range locationAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddLocationHook registers your hook function for all future operations.
func AddLocationHook(hookPoint boil.HookPoint, locationHook LocationHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		locationBeforeInsertHooks = append(locationBeforeInsertHooks, locationHook)
	case boil.BeforeUpdateHook:
		locationBeforeUpdateHooks = append(locationBeforeUpdateHooks, locationHook)
	case boil.BeforeDeleteHook:
		locationBeforeDeleteHooks = append(locationBeforeDeleteHooks, locationHook)
	case boil.BeforeUpsertHook:
		locationBeforeUpsertHooks = append(locationBeforeUpsertHooks, locationHook)
	case boil.AfterInsertHook:
		locationAfterInsertHooks = append(locationAfterInsertHooks, locationHook)
	case boil.AfterSelectHook:
		locationAfterSelectHooks = append(locationAfterSelectHooks, locationHook)
	case boil.AfterUpdateHook:
		locationAfterUpdateHooks = append(locationAfterUpdateHooks, locationHook)
	case boil.AfterDeleteHook:
		locationAfterDeleteHooks = append(locationAfterDeleteHooks, locationHook)
	case boil.AfterUpsertHook:
		locationAfterUpsertHooks = append(locationAfterUpsertHooks, locationHook)
	}
}

// One returns a single location record from the query.
func (q locationQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Location, error) {
	o := &Location{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "entity: failed to execute a one query for dm_location")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Location records from the query.
func (q locationQuery) All(ctx context.Context, exec boil.ContextExecutor) (LocationSlice, error) {
	var o []*Location

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "entity: failed to assign all query results to Location slice")
	}

	if len(locationAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Location records in the query.
func (q locationQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to count dm_location rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q locationQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "entity: failed to check if dm_location exists")
	}

	return count > 0, nil
}

var mySQLLocationUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Location) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("entity: no dm_location provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(locationColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLLocationUniqueColumns, o)

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

	locationUpsertCacheMut.RLock()
	cache, cached := locationUpsertCache[key]
	locationUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			locationColumns,
			locationColumnsWithDefault,
			locationColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			locationColumns,
			locationPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("entity: unable to upsert dm_location, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "dm_location", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `dm_location` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(locationType, locationMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(locationType, locationMapping, ret)
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
		return errors.Wrap(err, "entity: unable to upsert for dm_location")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == locationMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(locationType, locationMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "entity: unable to retrieve unique values for dm_location")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "entity: unable to populate default values for dm_location")
	}

CacheNoHooks:
	if !cached {
		locationUpsertCacheMut.Lock()
		locationUpsertCache[key] = cache
		locationUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}
