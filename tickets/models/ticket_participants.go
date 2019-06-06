// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// TicketParticipant is an object representing the database table.
type TicketParticipant struct {
	TicketGuildID int64  `boil:"ticket_guild_id" json:"ticket_guild_id" toml:"ticket_guild_id" yaml:"ticket_guild_id"`
	TicketLocalID int64  `boil:"ticket_local_id" json:"ticket_local_id" toml:"ticket_local_id" yaml:"ticket_local_id"`
	UserID        int64  `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	Username      string `boil:"username" json:"username" toml:"username" yaml:"username"`
	Discrim       string `boil:"discrim" json:"discrim" toml:"discrim" yaml:"discrim"`
	IsStaff       bool   `boil:"is_staff" json:"is_staff" toml:"is_staff" yaml:"is_staff"`

	R *ticketParticipantR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L ticketParticipantL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TicketParticipantColumns = struct {
	TicketGuildID string
	TicketLocalID string
	UserID        string
	Username      string
	Discrim       string
	IsStaff       string
}{
	TicketGuildID: "ticket_guild_id",
	TicketLocalID: "ticket_local_id",
	UserID:        "user_id",
	Username:      "username",
	Discrim:       "discrim",
	IsStaff:       "is_staff",
}

// Generated where

var TicketParticipantWhere = struct {
	TicketGuildID whereHelperint64
	TicketLocalID whereHelperint64
	UserID        whereHelperint64
	Username      whereHelperstring
	Discrim       whereHelperstring
	IsStaff       whereHelperbool
}{
	TicketGuildID: whereHelperint64{field: `ticket_guild_id`},
	TicketLocalID: whereHelperint64{field: `ticket_local_id`},
	UserID:        whereHelperint64{field: `user_id`},
	Username:      whereHelperstring{field: `username`},
	Discrim:       whereHelperstring{field: `discrim`},
	IsStaff:       whereHelperbool{field: `is_staff`},
}

// TicketParticipantRels is where relationship names are stored.
var TicketParticipantRels = struct {
}{}

// ticketParticipantR is where relationships are stored.
type ticketParticipantR struct {
}

// NewStruct creates a new relationship struct
func (*ticketParticipantR) NewStruct() *ticketParticipantR {
	return &ticketParticipantR{}
}

// ticketParticipantL is where Load methods for each relationship are stored.
type ticketParticipantL struct{}

var (
	ticketParticipantColumns               = []string{"ticket_guild_id", "ticket_local_id", "user_id", "username", "discrim", "is_staff"}
	ticketParticipantColumnsWithoutDefault = []string{"ticket_guild_id", "ticket_local_id", "user_id", "username", "discrim", "is_staff"}
	ticketParticipantColumnsWithDefault    = []string{}
	ticketParticipantPrimaryKeyColumns     = []string{"ticket_guild_id", "ticket_local_id", "user_id"}
)

type (
	// TicketParticipantSlice is an alias for a slice of pointers to TicketParticipant.
	// This should generally be used opposed to []TicketParticipant.
	TicketParticipantSlice []*TicketParticipant

	ticketParticipantQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	ticketParticipantType                 = reflect.TypeOf(&TicketParticipant{})
	ticketParticipantMapping              = queries.MakeStructMapping(ticketParticipantType)
	ticketParticipantPrimaryKeyMapping, _ = queries.BindMapping(ticketParticipantType, ticketParticipantMapping, ticketParticipantPrimaryKeyColumns)
	ticketParticipantInsertCacheMut       sync.RWMutex
	ticketParticipantInsertCache          = make(map[string]insertCache)
	ticketParticipantUpdateCacheMut       sync.RWMutex
	ticketParticipantUpdateCache          = make(map[string]updateCache)
	ticketParticipantUpsertCacheMut       sync.RWMutex
	ticketParticipantUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// OneG returns a single ticketParticipant record from the query using the global executor.
func (q ticketParticipantQuery) OneG(ctx context.Context) (*TicketParticipant, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single ticketParticipant record from the query.
func (q ticketParticipantQuery) One(ctx context.Context, exec boil.ContextExecutor) (*TicketParticipant, error) {
	o := &TicketParticipant{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for ticket_participants")
	}

	return o, nil
}

// AllG returns all TicketParticipant records from the query using the global executor.
func (q ticketParticipantQuery) AllG(ctx context.Context) (TicketParticipantSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all TicketParticipant records from the query.
func (q ticketParticipantQuery) All(ctx context.Context, exec boil.ContextExecutor) (TicketParticipantSlice, error) {
	var o []*TicketParticipant

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to TicketParticipant slice")
	}

	return o, nil
}

// CountG returns the count of all TicketParticipant records in the query, and panics on error.
func (q ticketParticipantQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all TicketParticipant records in the query.
func (q ticketParticipantQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count ticket_participants rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q ticketParticipantQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q ticketParticipantQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if ticket_participants exists")
	}

	return count > 0, nil
}

// TicketParticipants retrieves all the records using an executor.
func TicketParticipants(mods ...qm.QueryMod) ticketParticipantQuery {
	mods = append(mods, qm.From("\"ticket_participants\""))
	return ticketParticipantQuery{NewQuery(mods...)}
}

// FindTicketParticipantG retrieves a single record by ID.
func FindTicketParticipantG(ctx context.Context, ticketGuildID int64, ticketLocalID int64, userID int64, selectCols ...string) (*TicketParticipant, error) {
	return FindTicketParticipant(ctx, boil.GetContextDB(), ticketGuildID, ticketLocalID, userID, selectCols...)
}

// FindTicketParticipant retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTicketParticipant(ctx context.Context, exec boil.ContextExecutor, ticketGuildID int64, ticketLocalID int64, userID int64, selectCols ...string) (*TicketParticipant, error) {
	ticketParticipantObj := &TicketParticipant{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"ticket_participants\" where \"ticket_guild_id\"=$1 AND \"ticket_local_id\"=$2 AND \"user_id\"=$3", sel,
	)

	q := queries.Raw(query, ticketGuildID, ticketLocalID, userID)

	err := q.Bind(ctx, exec, ticketParticipantObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from ticket_participants")
	}

	return ticketParticipantObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *TicketParticipant) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *TicketParticipant) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no ticket_participants provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(ticketParticipantColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	ticketParticipantInsertCacheMut.RLock()
	cache, cached := ticketParticipantInsertCache[key]
	ticketParticipantInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			ticketParticipantColumns,
			ticketParticipantColumnsWithDefault,
			ticketParticipantColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(ticketParticipantType, ticketParticipantMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(ticketParticipantType, ticketParticipantMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"ticket_participants\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"ticket_participants\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into ticket_participants")
	}

	if !cached {
		ticketParticipantInsertCacheMut.Lock()
		ticketParticipantInsertCache[key] = cache
		ticketParticipantInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single TicketParticipant record using the global executor.
// See Update for more documentation.
func (o *TicketParticipant) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the TicketParticipant.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *TicketParticipant) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	ticketParticipantUpdateCacheMut.RLock()
	cache, cached := ticketParticipantUpdateCache[key]
	ticketParticipantUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			ticketParticipantColumns,
			ticketParticipantPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update ticket_participants, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"ticket_participants\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, ticketParticipantPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(ticketParticipantType, ticketParticipantMapping, append(wl, ticketParticipantPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update ticket_participants row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for ticket_participants")
	}

	if !cached {
		ticketParticipantUpdateCacheMut.Lock()
		ticketParticipantUpdateCache[key] = cache
		ticketParticipantUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (q ticketParticipantQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q ticketParticipantQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for ticket_participants")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for ticket_participants")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o TicketParticipantSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TicketParticipantSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), ticketParticipantPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"ticket_participants\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, ticketParticipantPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in ticketParticipant slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all ticketParticipant")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *TicketParticipant) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *TicketParticipant) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no ticket_participants provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(ticketParticipantColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
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
	key := buf.String()
	strmangle.PutBuffer(buf)

	ticketParticipantUpsertCacheMut.RLock()
	cache, cached := ticketParticipantUpsertCache[key]
	ticketParticipantUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			ticketParticipantColumns,
			ticketParticipantColumnsWithDefault,
			ticketParticipantColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			ticketParticipantColumns,
			ticketParticipantPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert ticket_participants, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(ticketParticipantPrimaryKeyColumns))
			copy(conflict, ticketParticipantPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"ticket_participants\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(ticketParticipantType, ticketParticipantMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(ticketParticipantType, ticketParticipantMapping, ret)
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

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert ticket_participants")
	}

	if !cached {
		ticketParticipantUpsertCacheMut.Lock()
		ticketParticipantUpsertCache[key] = cache
		ticketParticipantUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteG deletes a single TicketParticipant record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *TicketParticipant) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single TicketParticipant record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *TicketParticipant) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no TicketParticipant provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), ticketParticipantPrimaryKeyMapping)
	sql := "DELETE FROM \"ticket_participants\" WHERE \"ticket_guild_id\"=$1 AND \"ticket_local_id\"=$2 AND \"user_id\"=$3"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from ticket_participants")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for ticket_participants")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q ticketParticipantQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no ticketParticipantQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from ticket_participants")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for ticket_participants")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o TicketParticipantSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TicketParticipantSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no TicketParticipant slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), ticketParticipantPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"ticket_participants\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, ticketParticipantPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from ticketParticipant slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for ticket_participants")
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *TicketParticipant) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no TicketParticipant provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *TicketParticipant) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTicketParticipant(ctx, exec, o.TicketGuildID, o.TicketLocalID, o.UserID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TicketParticipantSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty TicketParticipantSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TicketParticipantSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TicketParticipantSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), ticketParticipantPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"ticket_participants\".* FROM \"ticket_participants\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, ticketParticipantPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in TicketParticipantSlice")
	}

	*o = slice

	return nil
}

// TicketParticipantExistsG checks if the TicketParticipant row exists.
func TicketParticipantExistsG(ctx context.Context, ticketGuildID int64, ticketLocalID int64, userID int64) (bool, error) {
	return TicketParticipantExists(ctx, boil.GetContextDB(), ticketGuildID, ticketLocalID, userID)
}

// TicketParticipantExists checks if the TicketParticipant row exists.
func TicketParticipantExists(ctx context.Context, exec boil.ContextExecutor, ticketGuildID int64, ticketLocalID int64, userID int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"ticket_participants\" where \"ticket_guild_id\"=$1 AND \"ticket_local_id\"=$2 AND \"user_id\"=$3 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, ticketGuildID, ticketLocalID, userID)
	}

	row := exec.QueryRowContext(ctx, sql, ticketGuildID, ticketLocalID, userID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if ticket_participants exists")
	}

	return exists, nil
}
