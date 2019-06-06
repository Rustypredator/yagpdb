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

func testTicketConfigs(t *testing.T) {
	t.Parallel()

	query := TicketConfigs()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testTicketConfigsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TicketConfig{}
	if err = randomize.Struct(seed, o, ticketConfigDBTypes, true, ticketConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TicketConfig struct: %s", err)
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

	count, err := TicketConfigs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTicketConfigsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TicketConfig{}
	if err = randomize.Struct(seed, o, ticketConfigDBTypes, true, ticketConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TicketConfig struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := TicketConfigs().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := TicketConfigs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTicketConfigsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TicketConfig{}
	if err = randomize.Struct(seed, o, ticketConfigDBTypes, true, ticketConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TicketConfig struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := TicketConfigSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := TicketConfigs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTicketConfigsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TicketConfig{}
	if err = randomize.Struct(seed, o, ticketConfigDBTypes, true, ticketConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TicketConfig struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := TicketConfigExists(ctx, tx, o.GuildID)
	if err != nil {
		t.Errorf("Unable to check if TicketConfig exists: %s", err)
	}
	if !e {
		t.Errorf("Expected TicketConfigExists to return true, but got false.")
	}
}

func testTicketConfigsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TicketConfig{}
	if err = randomize.Struct(seed, o, ticketConfigDBTypes, true, ticketConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TicketConfig struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	ticketConfigFound, err := FindTicketConfig(ctx, tx, o.GuildID)
	if err != nil {
		t.Error(err)
	}

	if ticketConfigFound == nil {
		t.Error("want a record, got nil")
	}
}

func testTicketConfigsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TicketConfig{}
	if err = randomize.Struct(seed, o, ticketConfigDBTypes, true, ticketConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TicketConfig struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = TicketConfigs().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testTicketConfigsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TicketConfig{}
	if err = randomize.Struct(seed, o, ticketConfigDBTypes, true, ticketConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TicketConfig struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := TicketConfigs().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testTicketConfigsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	ticketConfigOne := &TicketConfig{}
	ticketConfigTwo := &TicketConfig{}
	if err = randomize.Struct(seed, ticketConfigOne, ticketConfigDBTypes, false, ticketConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TicketConfig struct: %s", err)
	}
	if err = randomize.Struct(seed, ticketConfigTwo, ticketConfigDBTypes, false, ticketConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TicketConfig struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = ticketConfigOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = ticketConfigTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := TicketConfigs().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testTicketConfigsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	ticketConfigOne := &TicketConfig{}
	ticketConfigTwo := &TicketConfig{}
	if err = randomize.Struct(seed, ticketConfigOne, ticketConfigDBTypes, false, ticketConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TicketConfig struct: %s", err)
	}
	if err = randomize.Struct(seed, ticketConfigTwo, ticketConfigDBTypes, false, ticketConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TicketConfig struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = ticketConfigOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = ticketConfigTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TicketConfigs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testTicketConfigsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TicketConfig{}
	if err = randomize.Struct(seed, o, ticketConfigDBTypes, true, ticketConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TicketConfig struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TicketConfigs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTicketConfigsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TicketConfig{}
	if err = randomize.Struct(seed, o, ticketConfigDBTypes, true); err != nil {
		t.Errorf("Unable to randomize TicketConfig struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(ticketConfigColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := TicketConfigs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTicketConfigsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TicketConfig{}
	if err = randomize.Struct(seed, o, ticketConfigDBTypes, true, ticketConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TicketConfig struct: %s", err)
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

func testTicketConfigsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TicketConfig{}
	if err = randomize.Struct(seed, o, ticketConfigDBTypes, true, ticketConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TicketConfig struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := TicketConfigSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testTicketConfigsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TicketConfig{}
	if err = randomize.Struct(seed, o, ticketConfigDBTypes, true, ticketConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TicketConfig struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := TicketConfigs().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	ticketConfigDBTypes = map[string]string{`GuildID`: `bigint`, `Enabled`: `boolean`, `TicketOpenMSG`: `text`, `TicketsChannelCategory`: `bigint`, `StatusChannel`: `bigint`, `TicketsTranscriptsChannel`: `bigint`, `DownloadAttachments`: `boolean`, `TicketsUseTXTTranscripts`: `boolean`, `ModRoles`: `ARRAYbigint`, `AdminRoles`: `ARRAYbigint`}
	_                   = bytes.MinRead
)

func testTicketConfigsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(ticketConfigPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(ticketConfigColumns) == len(ticketConfigPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &TicketConfig{}
	if err = randomize.Struct(seed, o, ticketConfigDBTypes, true, ticketConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TicketConfig struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TicketConfigs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, ticketConfigDBTypes, true, ticketConfigPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize TicketConfig struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testTicketConfigsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(ticketConfigColumns) == len(ticketConfigPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &TicketConfig{}
	if err = randomize.Struct(seed, o, ticketConfigDBTypes, true, ticketConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TicketConfig struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TicketConfigs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, ticketConfigDBTypes, true, ticketConfigPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize TicketConfig struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(ticketConfigColumns, ticketConfigPrimaryKeyColumns) {
		fields = ticketConfigColumns
	} else {
		fields = strmangle.SetComplement(
			ticketConfigColumns,
			ticketConfigPrimaryKeyColumns,
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

	slice := TicketConfigSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testTicketConfigsUpsert(t *testing.T) {
	t.Parallel()

	if len(ticketConfigColumns) == len(ticketConfigPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := TicketConfig{}
	if err = randomize.Struct(seed, &o, ticketConfigDBTypes, true); err != nil {
		t.Errorf("Unable to randomize TicketConfig struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert TicketConfig: %s", err)
	}

	count, err := TicketConfigs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, ticketConfigDBTypes, false, ticketConfigPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize TicketConfig struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert TicketConfig: %s", err)
	}

	count, err = TicketConfigs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
