// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/customid/ent/link"
	"entgo.io/ent/entc/integration/customid/ent/schema"
	uuidc "entgo.io/ent/entc/integration/customid/uuidcompatible"
	"entgo.io/ent/schema/field"
)

// LinkCreate is the builder for creating a Link entity.
type LinkCreate struct {
	config
	mutation *LinkMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetLinkInformation sets the "link_information" field.
func (lc *LinkCreate) SetLinkInformation(mi map[string]schema.LinkInformation) *LinkCreate {
	lc.mutation.SetLinkInformation(mi)
	return lc
}

// SetID sets the "id" field.
func (lc *LinkCreate) SetID(u uuidc.UUIDC) *LinkCreate {
	lc.mutation.SetID(u)
	return lc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (lc *LinkCreate) SetNillableID(u *uuidc.UUIDC) *LinkCreate {
	if u != nil {
		lc.SetID(*u)
	}
	return lc
}

// Mutation returns the LinkMutation object of the builder.
func (lc *LinkCreate) Mutation() *LinkMutation {
	return lc.mutation
}

// Save creates the Link in the database.
func (lc *LinkCreate) Save(ctx context.Context) (*Link, error) {
	lc.defaults()
	return withHooks[*Link, LinkMutation](ctx, lc.sqlSave, lc.mutation, lc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (lc *LinkCreate) SaveX(ctx context.Context) *Link {
	v, err := lc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lc *LinkCreate) Exec(ctx context.Context) error {
	_, err := lc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lc *LinkCreate) ExecX(ctx context.Context) {
	if err := lc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lc *LinkCreate) defaults() {
	if _, ok := lc.mutation.LinkInformation(); !ok {
		v := link.DefaultLinkInformation
		lc.mutation.SetLinkInformation(v)
	}
	if _, ok := lc.mutation.ID(); !ok {
		v := link.DefaultID()
		lc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lc *LinkCreate) check() error {
	if _, ok := lc.mutation.LinkInformation(); !ok {
		return &ValidationError{Name: "link_information", err: errors.New(`ent: missing required field "Link.link_information"`)}
	}
	return nil
}

func (lc *LinkCreate) sqlSave(ctx context.Context) (*Link, error) {
	if err := lc.check(); err != nil {
		return nil, err
	}
	_node, _spec := lc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuidc.UUIDC); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	lc.mutation.id = &_node.ID
	lc.mutation.done = true
	return _node, nil
}

func (lc *LinkCreate) createSpec() (*Link, *sqlgraph.CreateSpec) {
	var (
		_node = &Link{config: lc.config}
		_spec = sqlgraph.NewCreateSpec(link.Table, sqlgraph.NewFieldSpec(link.FieldID, field.TypeUUID))
	)
	_spec.OnConflict = lc.conflict
	if id, ok := lc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := lc.mutation.LinkInformation(); ok {
		_spec.SetField(link.FieldLinkInformation, field.TypeJSON, value)
		_node.LinkInformation = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Link.Create().
//		SetLinkInformation(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.LinkUpsert) {
//			SetLinkInformation(v+v).
//		}).
//		Exec(ctx)
func (lc *LinkCreate) OnConflict(opts ...sql.ConflictOption) *LinkUpsertOne {
	lc.conflict = opts
	return &LinkUpsertOne{
		create: lc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Link.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (lc *LinkCreate) OnConflictColumns(columns ...string) *LinkUpsertOne {
	lc.conflict = append(lc.conflict, sql.ConflictColumns(columns...))
	return &LinkUpsertOne{
		create: lc,
	}
}

type (
	// LinkUpsertOne is the builder for "upsert"-ing
	//  one Link node.
	LinkUpsertOne struct {
		create *LinkCreate
	}

	// LinkUpsert is the "OnConflict" setter.
	LinkUpsert struct {
		*sql.UpdateSet
	}
)

// SetLinkInformation sets the "link_information" field.
func (u *LinkUpsert) SetLinkInformation(v map[string]schema.LinkInformation) *LinkUpsert {
	u.Set(link.FieldLinkInformation, v)
	return u
}

// UpdateLinkInformation sets the "link_information" field to the value that was provided on create.
func (u *LinkUpsert) UpdateLinkInformation() *LinkUpsert {
	u.SetExcluded(link.FieldLinkInformation)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Link.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(link.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *LinkUpsertOne) UpdateNewValues() *LinkUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(link.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Link.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *LinkUpsertOne) Ignore() *LinkUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *LinkUpsertOne) DoNothing() *LinkUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the LinkCreate.OnConflict
// documentation for more info.
func (u *LinkUpsertOne) Update(set func(*LinkUpsert)) *LinkUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&LinkUpsert{UpdateSet: update})
	}))
	return u
}

// SetLinkInformation sets the "link_information" field.
func (u *LinkUpsertOne) SetLinkInformation(v map[string]schema.LinkInformation) *LinkUpsertOne {
	return u.Update(func(s *LinkUpsert) {
		s.SetLinkInformation(v)
	})
}

// UpdateLinkInformation sets the "link_information" field to the value that was provided on create.
func (u *LinkUpsertOne) UpdateLinkInformation() *LinkUpsertOne {
	return u.Update(func(s *LinkUpsert) {
		s.UpdateLinkInformation()
	})
}

// Exec executes the query.
func (u *LinkUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for LinkCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *LinkUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *LinkUpsertOne) ID(ctx context.Context) (id uuidc.UUIDC, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: LinkUpsertOne.ID is not supported by MySQL driver. Use LinkUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *LinkUpsertOne) IDX(ctx context.Context) uuidc.UUIDC {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// LinkCreateBulk is the builder for creating many Link entities in bulk.
type LinkCreateBulk struct {
	config
	builders []*LinkCreate
	conflict []sql.ConflictOption
}

// Save creates the Link entities in the database.
func (lcb *LinkCreateBulk) Save(ctx context.Context) ([]*Link, error) {
	specs := make([]*sqlgraph.CreateSpec, len(lcb.builders))
	nodes := make([]*Link, len(lcb.builders))
	mutators := make([]Mutator, len(lcb.builders))
	for i := range lcb.builders {
		func(i int, root context.Context) {
			builder := lcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LinkMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, lcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = lcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, lcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lcb *LinkCreateBulk) SaveX(ctx context.Context) []*Link {
	v, err := lcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lcb *LinkCreateBulk) Exec(ctx context.Context) error {
	_, err := lcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lcb *LinkCreateBulk) ExecX(ctx context.Context) {
	if err := lcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Link.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.LinkUpsert) {
//			SetLinkInformation(v+v).
//		}).
//		Exec(ctx)
func (lcb *LinkCreateBulk) OnConflict(opts ...sql.ConflictOption) *LinkUpsertBulk {
	lcb.conflict = opts
	return &LinkUpsertBulk{
		create: lcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Link.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (lcb *LinkCreateBulk) OnConflictColumns(columns ...string) *LinkUpsertBulk {
	lcb.conflict = append(lcb.conflict, sql.ConflictColumns(columns...))
	return &LinkUpsertBulk{
		create: lcb,
	}
}

// LinkUpsertBulk is the builder for "upsert"-ing
// a bulk of Link nodes.
type LinkUpsertBulk struct {
	create *LinkCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Link.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(link.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *LinkUpsertBulk) UpdateNewValues() *LinkUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(link.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Link.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *LinkUpsertBulk) Ignore() *LinkUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *LinkUpsertBulk) DoNothing() *LinkUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the LinkCreateBulk.OnConflict
// documentation for more info.
func (u *LinkUpsertBulk) Update(set func(*LinkUpsert)) *LinkUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&LinkUpsert{UpdateSet: update})
	}))
	return u
}

// SetLinkInformation sets the "link_information" field.
func (u *LinkUpsertBulk) SetLinkInformation(v map[string]schema.LinkInformation) *LinkUpsertBulk {
	return u.Update(func(s *LinkUpsert) {
		s.SetLinkInformation(v)
	})
}

// UpdateLinkInformation sets the "link_information" field to the value that was provided on create.
func (u *LinkUpsertBulk) UpdateLinkInformation() *LinkUpsertBulk {
	return u.Update(func(s *LinkUpsert) {
		s.UpdateLinkInformation()
	})
}

// Exec executes the query.
func (u *LinkUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the LinkCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for LinkCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *LinkUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
