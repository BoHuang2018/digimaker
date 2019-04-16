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
	"dm/model"
	. "dm/query"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// Article is an object representing the database table.
// Implement dm.model.ContentTyper interface
type Article struct {
	CID       int                     `boil:"id" json:"id" toml:"id" yaml:"id"`
	Status    int                     `boil:"status" json:"status" toml:"status" yaml:"status"`
	Author    int                     `boil:"author" json:"author" toml:"author" yaml:"author"`
	Title     fieldtype.TextField     `boil:"title" json:"title" toml:"title" yaml:"title"`
	Body      fieldtype.RichTextField `boil:"body" json:"body" toml:"body" yaml:"body"`
	Published int                     `boil:"published" json:"published,omitempty" toml:"published" yaml:"published,omitempty"`
	Modified  int                     `boil:"modified" json:"modified,omitempty" toml:"modified" yaml:"modified,omitempty"`
	RemoteID  string                  `boil:"remote_id" json:"remote_id" toml:"remote_id" yaml:"remote_id"`

	R        *articleR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L        articleL  `boil:"-" json:"-" toml:"-" yaml:"-"`
	Location `boil:"dm_location,bind"`
}

func (c Article) Fields() map[string]model.Fielder {
	return nil
}

func (c *Article) Values() map[string]interface{} {
	result := make(map[string]interface{})
	result["id"] = c.CID
	result["status"] = c.Status
	result["author"] = c.Author
	result["title"] = c.Title
	result["body"] = c.Body
	result["published"] = c.Published
	result["modified"] = c.Modified
	result["remote_id"] = c.RemoteID
	return result
}

func (c *Article) Value(identifier string, value interface{}) {

}

func (c *Article) TableName() string {
	return "dm_article"
}

func (c Article) Field(name string) interface{} {
	var result interface{}
	switch name {
	case "id", "CID":
		result = c.CID
	case "status", "Status":
		result = c.Status
	case "author", "Author":
		result = c.Author
	case "title", "Title":
		result = c.Title
	case "body", "Body":
		result = c.Body
	case "published", "Published":
		result = c.Published
	case "modified", "Modified":
		result = c.Modified
	case "remote_id", "RemoteID":
		result = c.RemoteID
	default:
	}
	return result
}

func (c Article) Store() error {
	handler := db.DBHanlder()
	if c.CID == 0 {
		id, err := handler.Insert(c.TableName(), c.Values())
		c.CID = id
		if err != nil {
			return err
		}
	} else {
		err := handler.Update(c.TableName(), c.Values(), Cond("id", c.CID))
		return err
	}
	return nil
}

var ArticleColumns = struct {
	CID       string
	Status    string
	Author    string
	Title     string
	Body      string
	Published string
	Modified  string
	RemoteID  string
}{
	CID:       "id",
	Status:    "status",
	Author:    "author",
	Title:     "title",
	Body:      "body",
	Published: "published",
	Modified:  "modified",
	RemoteID:  "remote_id",
}

// ArticleRels is where relationship names are stored.
var ArticleRels = struct {
}{}

// articleR is where relationships are stored.
type articleR struct {
}

// NewStruct creates a new relationship struct
func (*articleR) NewStruct() *articleR {
	return &articleR{}
}

// articleL is where Load methods for each relationship are stored.
type articleL struct{}

var (
	articleColumns               = []string{"id", "status", "author", "title", "body", "published", "modified", "remote_id"}
	articleColumnsWithoutDefault = []string{"title", "body", "published", "modified", "remote_id"}
	articleColumnsWithDefault    = []string{"id", "status", "author"}
	articlePrimaryKeyColumns     = []string{"id"}
)

type (
	// ArticleSlice is an alias for a slice of pointers to Article.
	// This should generally be used opposed to []Article.
	ArticleSlice []*Article
	// ArticleHook is the signature for custom Article hook methods
	ArticleHook func(context.Context, boil.ContextExecutor, *Article) error

	articleQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	articleType                 = reflect.TypeOf(&Article{})
	articleMapping              = queries.MakeStructMapping(articleType)
	articlePrimaryKeyMapping, _ = queries.BindMapping(articleType, articleMapping, articlePrimaryKeyColumns)
	articleInsertCacheMut       sync.RWMutex
	articleInsertCache          = make(map[string]insertCache)
	articleUpdateCacheMut       sync.RWMutex
	articleUpdateCache          = make(map[string]updateCache)
	articleUpsertCacheMut       sync.RWMutex
	articleUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var articleBeforeInsertHooks []ArticleHook
var articleBeforeUpdateHooks []ArticleHook
var articleBeforeDeleteHooks []ArticleHook
var articleBeforeUpsertHooks []ArticleHook

var articleAfterInsertHooks []ArticleHook
var articleAfterSelectHooks []ArticleHook
var articleAfterUpdateHooks []ArticleHook
var articleAfterDeleteHooks []ArticleHook
var articleAfterUpsertHooks []ArticleHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Article) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range articleBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Article) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range articleBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Article) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range articleBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Article) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range articleBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Article) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range articleAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Article) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range articleAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Article) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range articleAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Article) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range articleAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Article) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range articleAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddArticleHook registers your hook function for all future operations.
func AddArticleHook(hookPoint boil.HookPoint, articleHook ArticleHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		articleBeforeInsertHooks = append(articleBeforeInsertHooks, articleHook)
	case boil.BeforeUpdateHook:
		articleBeforeUpdateHooks = append(articleBeforeUpdateHooks, articleHook)
	case boil.BeforeDeleteHook:
		articleBeforeDeleteHooks = append(articleBeforeDeleteHooks, articleHook)
	case boil.BeforeUpsertHook:
		articleBeforeUpsertHooks = append(articleBeforeUpsertHooks, articleHook)
	case boil.AfterInsertHook:
		articleAfterInsertHooks = append(articleAfterInsertHooks, articleHook)
	case boil.AfterSelectHook:
		articleAfterSelectHooks = append(articleAfterSelectHooks, articleHook)
	case boil.AfterUpdateHook:
		articleAfterUpdateHooks = append(articleAfterUpdateHooks, articleHook)
	case boil.AfterDeleteHook:
		articleAfterDeleteHooks = append(articleAfterDeleteHooks, articleHook)
	case boil.AfterUpsertHook:
		articleAfterUpsertHooks = append(articleAfterUpsertHooks, articleHook)
	}
}

// One returns a single article record from the query.
func (q articleQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Article, error) {
	o := &Article{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "entity: failed to execute a one query for dm_article")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Article records from the query.
func (q articleQuery) All(ctx context.Context, exec boil.ContextExecutor) (ArticleSlice, error) {
	var o []*Article

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "entity: failed to assign all query results to Article slice")
	}

	if len(articleAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Article records in the query.
func (q articleQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to count dm_article rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q articleQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "entity: failed to check if dm_article exists")
	}

	return count > 0, nil
}

var mySQLArticleUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Article) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("entity: no dm_article provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(articleColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLArticleUniqueColumns, o)

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

	articleUpsertCacheMut.RLock()
	cache, cached := articleUpsertCache[key]
	articleUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			articleColumns,
			articleColumnsWithDefault,
			articleColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			articleColumns,
			articlePrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("entity: unable to upsert dm_article, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "dm_article", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `dm_article` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(articleType, articleMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(articleType, articleMapping, ret)
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
		return errors.Wrap(err, "entity: unable to upsert for dm_article")
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

	o.CID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == articleMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(articleType, articleMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "entity: unable to retrieve unique values for dm_article")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "entity: unable to populate default values for dm_article")
	}

CacheNoHooks:
	if !cached {
		articleUpsertCacheMut.Lock()
		articleUpsertCache[key] = cache
		articleUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}
