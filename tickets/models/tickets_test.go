// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testTickets(t *testing.T) {
	t.Parallel()

	query := Tickets()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testTicketsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Ticket{}
	if err = randomize.Struct(seed, o, ticketDBTypes, true, ticketColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Ticket struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Tickets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTicketsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Ticket{}
	if err = randomize.Struct(seed, o, ticketDBTypes, true, ticketColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Ticket struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Tickets().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Tickets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTicketsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Ticket{}
	if err = randomize.Struct(seed, o, ticketDBTypes, true, ticketColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Ticket struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := TicketSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Tickets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTicketsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Ticket{}
	if err = randomize.Struct(seed, o, ticketDBTypes, true, ticketColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Ticket struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := TicketExists(ctx, tx, o.GuildID, o.LocalID)
	if err != nil {
		t.Errorf("Unable to check if Ticket exists: %s", err)
	}
	if !e {
		t.Errorf("Expected TicketExists to return true, but got false.")
	}
}

func testTicketsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Ticket{}
	if err = randomize.Struct(seed, o, ticketDBTypes, true, ticketColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Ticket struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	ticketFound, err := FindTicket(ctx, tx, o.GuildID, o.LocalID)
	if err != nil {
		t.Error(err)
	}

	if ticketFound == nil {
		t.Error("want a record, got nil")
	}
}

func testTicketsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Ticket{}
	if err = randomize.Struct(seed, o, ticketDBTypes, true, ticketColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Ticket struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Tickets().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testTicketsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Ticket{}
	if err = randomize.Struct(seed, o, ticketDBTypes, true, ticketColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Ticket struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Tickets().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testTicketsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	ticketOne := &Ticket{}
	ticketTwo := &Ticket{}
	if err = randomize.Struct(seed, ticketOne, ticketDBTypes, false, ticketColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Ticket struct: %s", err)
	}
	if err = randomize.Struct(seed, ticketTwo, ticketDBTypes, false, ticketColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Ticket struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = ticketOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = ticketTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Tickets().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testTicketsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	ticketOne := &Ticket{}
	ticketTwo := &Ticket{}
	if err = randomize.Struct(seed, ticketOne, ticketDBTypes, false, ticketColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Ticket struct: %s", err)
	}
	if err = randomize.Struct(seed, ticketTwo, ticketDBTypes, false, ticketColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Ticket struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = ticketOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = ticketTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Tickets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testTicketsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Ticket{}
	if err = randomize.Struct(seed, o, ticketDBTypes, true, ticketColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Ticket struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Tickets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTicketsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Ticket{}
	if err = randomize.Struct(seed, o, ticketDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Ticket struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(ticketColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Tickets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTicketsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Ticket{}
	if err = randomize.Struct(seed, o, ticketDBTypes, true, ticketColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Ticket struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testTicketsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Ticket{}
	if err = randomize.Struct(seed, o, ticketDBTypes, true, ticketColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Ticket struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := TicketSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testTicketsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Ticket{}
	if err = randomize.Struct(seed, o, ticketDBTypes, true, ticketColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Ticket struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Tickets().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	ticketDBTypes = map[string]string{`GuildID`: `bigint`, `LocalID`: `bigint`, `ChannelID`: `bigint`, `Title`: `text`, `CreatedAt`: `timestamp with time zone`, `ClosedAt`: `timestamp with time zone`, `LogsID`: `bigint`, `AuthorID`: `bigint`, `AuthorUsernameDiscrim`: `text`}
	_             = bytes.MinRead
)

func testTicketsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(ticketPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(ticketColumns) == len(ticketPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Ticket{}
	if err = randomize.Struct(seed, o, ticketDBTypes, true, ticketColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Ticket struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Tickets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, ticketDBTypes, true, ticketPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Ticket struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testTicketsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(ticketColumns) == len(ticketPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Ticket{}
	if err = randomize.Struct(seed, o, ticketDBTypes, true, ticketColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Ticket struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Tickets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, ticketDBTypes, true, ticketPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Ticket struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(ticketColumns, ticketPrimaryKeyColumns) {
		fields = ticketColumns
	} else {
		fields = strmangle.SetComplement(
			ticketColumns,
			ticketPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := TicketSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testTicketsUpsert(t *testing.T) {
	t.Parallel()

	if len(ticketColumns) == len(ticketPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Ticket{}
	if err = randomize.Struct(seed, &o, ticketDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Ticket struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Ticket: %s", err)
	}

	count, err := Tickets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, ticketDBTypes, false, ticketPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Ticket struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Ticket: %s", err)
	}

	count, err = Tickets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
