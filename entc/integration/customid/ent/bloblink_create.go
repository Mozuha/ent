// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/customid/ent/blob"
	"entgo.io/ent/entc/integration/customid/ent/bloblink"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// BlobLinkCreate is the builder for creating a BlobLink entity.
type BlobLinkCreate struct {
	config
	mutation *BlobLinkMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (blc *BlobLinkCreate) SetCreatedAt(t time.Time) *BlobLinkCreate {
	blc.mutation.SetCreatedAt(t)
	return blc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (blc *BlobLinkCreate) SetNillableCreatedAt(t *time.Time) *BlobLinkCreate {
	if t != nil {
		blc.SetCreatedAt(*t)
	}
	return blc
}

// SetBlobID sets the "blob_id" field.
func (blc *BlobLinkCreate) SetBlobID(u uuid.UUID) *BlobLinkCreate {
	blc.mutation.SetBlobID(u)
	return blc
}

// SetLinkID sets the "link_id" field.
func (blc *BlobLinkCreate) SetLinkID(u uuid.UUID) *BlobLinkCreate {
	blc.mutation.SetLinkID(u)
	return blc
}

// SetBlob sets the "blob" edge to the Blob entity.
func (blc *BlobLinkCreate) SetBlob(b *Blob) *BlobLinkCreate {
	return blc.SetBlobID(b.ID)
}

// SetLink sets the "link" edge to the Blob entity.
func (blc *BlobLinkCreate) SetLink(b *Blob) *BlobLinkCreate {
	return blc.SetLinkID(b.ID)
}

// Mutation returns the BlobLinkMutation object of the builder.
func (blc *BlobLinkCreate) Mutation() *BlobLinkMutation {
	return blc.mutation
}

// Save creates the BlobLink in the database.
func (blc *BlobLinkCreate) Save(ctx context.Context) (*BlobLink, error) {
	blc.defaults()
	return withHooks[*BlobLink, BlobLinkMutation](ctx, blc.sqlSave, blc.mutation, blc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (blc *BlobLinkCreate) SaveX(ctx context.Context) *BlobLink {
	v, err := blc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (blc *BlobLinkCreate) Exec(ctx context.Context) error {
	_, err := blc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (blc *BlobLinkCreate) ExecX(ctx context.Context) {
	if err := blc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (blc *BlobLinkCreate) defaults() {
	if _, ok := blc.mutation.CreatedAt(); !ok {
		v := bloblink.DefaultCreatedAt()
		blc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (blc *BlobLinkCreate) check() error {
	if _, ok := blc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "BlobLink.created_at"`)}
	}
	if _, ok := blc.mutation.BlobID(); !ok {
		return &ValidationError{Name: "blob_id", err: errors.New(`ent: missing required field "BlobLink.blob_id"`)}
	}
	if _, ok := blc.mutation.LinkID(); !ok {
		return &ValidationError{Name: "link_id", err: errors.New(`ent: missing required field "BlobLink.link_id"`)}
	}
	if _, ok := blc.mutation.BlobID(); !ok {
		return &ValidationError{Name: "blob", err: errors.New(`ent: missing required edge "BlobLink.blob"`)}
	}
	if _, ok := blc.mutation.LinkID(); !ok {
		return &ValidationError{Name: "link", err: errors.New(`ent: missing required edge "BlobLink.link"`)}
	}
	return nil
}

func (blc *BlobLinkCreate) sqlSave(ctx context.Context) (*BlobLink, error) {
	if err := blc.check(); err != nil {
		return nil, err
	}
	_node, _spec := blc.createSpec()
	if err := sqlgraph.CreateNode(ctx, blc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}

func (blc *BlobLinkCreate) createSpec() (*BlobLink, *sqlgraph.CreateSpec) {
	var (
		_node = &BlobLink{config: blc.config}
		_spec = sqlgraph.NewCreateSpec(bloblink.Table, nil)
	)
	_spec.OnConflict = blc.conflict
	if value, ok := blc.mutation.CreatedAt(); ok {
		_spec.SetField(bloblink.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := blc.mutation.BlobIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   bloblink.BlobTable,
			Columns: []string{bloblink.BlobColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(blob.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.BlobID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := blc.mutation.LinkIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   bloblink.LinkTable,
			Columns: []string{bloblink.LinkColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(blob.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.LinkID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.BlobLink.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.BlobLinkUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (blc *BlobLinkCreate) OnConflict(opts ...sql.ConflictOption) *BlobLinkUpsertOne {
	blc.conflict = opts
	return &BlobLinkUpsertOne{
		create: blc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.BlobLink.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (blc *BlobLinkCreate) OnConflictColumns(columns ...string) *BlobLinkUpsertOne {
	blc.conflict = append(blc.conflict, sql.ConflictColumns(columns...))
	return &BlobLinkUpsertOne{
		create: blc,
	}
}

type (
	// BlobLinkUpsertOne is the builder for "upsert"-ing
	//  one BlobLink node.
	BlobLinkUpsertOne struct {
		create *BlobLinkCreate
	}

	// BlobLinkUpsert is the "OnConflict" setter.
	BlobLinkUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *BlobLinkUpsert) SetCreatedAt(v time.Time) *BlobLinkUpsert {
	u.Set(bloblink.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *BlobLinkUpsert) UpdateCreatedAt() *BlobLinkUpsert {
	u.SetExcluded(bloblink.FieldCreatedAt)
	return u
}

// SetBlobID sets the "blob_id" field.
func (u *BlobLinkUpsert) SetBlobID(v uuid.UUID) *BlobLinkUpsert {
	u.Set(bloblink.FieldBlobID, v)
	return u
}

// UpdateBlobID sets the "blob_id" field to the value that was provided on create.
func (u *BlobLinkUpsert) UpdateBlobID() *BlobLinkUpsert {
	u.SetExcluded(bloblink.FieldBlobID)
	return u
}

// SetLinkID sets the "link_id" field.
func (u *BlobLinkUpsert) SetLinkID(v uuid.UUID) *BlobLinkUpsert {
	u.Set(bloblink.FieldLinkID, v)
	return u
}

// UpdateLinkID sets the "link_id" field to the value that was provided on create.
func (u *BlobLinkUpsert) UpdateLinkID() *BlobLinkUpsert {
	u.SetExcluded(bloblink.FieldLinkID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.BlobLink.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *BlobLinkUpsertOne) UpdateNewValues() *BlobLinkUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.BlobLink.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *BlobLinkUpsertOne) Ignore() *BlobLinkUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *BlobLinkUpsertOne) DoNothing() *BlobLinkUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the BlobLinkCreate.OnConflict
// documentation for more info.
func (u *BlobLinkUpsertOne) Update(set func(*BlobLinkUpsert)) *BlobLinkUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&BlobLinkUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *BlobLinkUpsertOne) SetCreatedAt(v time.Time) *BlobLinkUpsertOne {
	return u.Update(func(s *BlobLinkUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *BlobLinkUpsertOne) UpdateCreatedAt() *BlobLinkUpsertOne {
	return u.Update(func(s *BlobLinkUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetBlobID sets the "blob_id" field.
func (u *BlobLinkUpsertOne) SetBlobID(v uuid.UUID) *BlobLinkUpsertOne {
	return u.Update(func(s *BlobLinkUpsert) {
		s.SetBlobID(v)
	})
}

// UpdateBlobID sets the "blob_id" field to the value that was provided on create.
func (u *BlobLinkUpsertOne) UpdateBlobID() *BlobLinkUpsertOne {
	return u.Update(func(s *BlobLinkUpsert) {
		s.UpdateBlobID()
	})
}

// SetLinkID sets the "link_id" field.
func (u *BlobLinkUpsertOne) SetLinkID(v uuid.UUID) *BlobLinkUpsertOne {
	return u.Update(func(s *BlobLinkUpsert) {
		s.SetLinkID(v)
	})
}

// UpdateLinkID sets the "link_id" field to the value that was provided on create.
func (u *BlobLinkUpsertOne) UpdateLinkID() *BlobLinkUpsertOne {
	return u.Update(func(s *BlobLinkUpsert) {
		s.UpdateLinkID()
	})
}

// Exec executes the query.
func (u *BlobLinkUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for BlobLinkCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *BlobLinkUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// BlobLinkCreateBulk is the builder for creating many BlobLink entities in bulk.
type BlobLinkCreateBulk struct {
	config
	builders []*BlobLinkCreate
	conflict []sql.ConflictOption
}

// Save creates the BlobLink entities in the database.
func (blcb *BlobLinkCreateBulk) Save(ctx context.Context) ([]*BlobLink, error) {
	specs := make([]*sqlgraph.CreateSpec, len(blcb.builders))
	nodes := make([]*BlobLink, len(blcb.builders))
	mutators := make([]Mutator, len(blcb.builders))
	for i := range blcb.builders {
		func(i int, root context.Context) {
			builder := blcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BlobLinkMutation)
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
					_, err = mutators[i+1].Mutate(root, blcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = blcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, blcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
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
		if _, err := mutators[0].Mutate(ctx, blcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (blcb *BlobLinkCreateBulk) SaveX(ctx context.Context) []*BlobLink {
	v, err := blcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (blcb *BlobLinkCreateBulk) Exec(ctx context.Context) error {
	_, err := blcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (blcb *BlobLinkCreateBulk) ExecX(ctx context.Context) {
	if err := blcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.BlobLink.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.BlobLinkUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (blcb *BlobLinkCreateBulk) OnConflict(opts ...sql.ConflictOption) *BlobLinkUpsertBulk {
	blcb.conflict = opts
	return &BlobLinkUpsertBulk{
		create: blcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.BlobLink.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (blcb *BlobLinkCreateBulk) OnConflictColumns(columns ...string) *BlobLinkUpsertBulk {
	blcb.conflict = append(blcb.conflict, sql.ConflictColumns(columns...))
	return &BlobLinkUpsertBulk{
		create: blcb,
	}
}

// BlobLinkUpsertBulk is the builder for "upsert"-ing
// a bulk of BlobLink nodes.
type BlobLinkUpsertBulk struct {
	create *BlobLinkCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.BlobLink.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *BlobLinkUpsertBulk) UpdateNewValues() *BlobLinkUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.BlobLink.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *BlobLinkUpsertBulk) Ignore() *BlobLinkUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *BlobLinkUpsertBulk) DoNothing() *BlobLinkUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the BlobLinkCreateBulk.OnConflict
// documentation for more info.
func (u *BlobLinkUpsertBulk) Update(set func(*BlobLinkUpsert)) *BlobLinkUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&BlobLinkUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *BlobLinkUpsertBulk) SetCreatedAt(v time.Time) *BlobLinkUpsertBulk {
	return u.Update(func(s *BlobLinkUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *BlobLinkUpsertBulk) UpdateCreatedAt() *BlobLinkUpsertBulk {
	return u.Update(func(s *BlobLinkUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetBlobID sets the "blob_id" field.
func (u *BlobLinkUpsertBulk) SetBlobID(v uuid.UUID) *BlobLinkUpsertBulk {
	return u.Update(func(s *BlobLinkUpsert) {
		s.SetBlobID(v)
	})
}

// UpdateBlobID sets the "blob_id" field to the value that was provided on create.
func (u *BlobLinkUpsertBulk) UpdateBlobID() *BlobLinkUpsertBulk {
	return u.Update(func(s *BlobLinkUpsert) {
		s.UpdateBlobID()
	})
}

// SetLinkID sets the "link_id" field.
func (u *BlobLinkUpsertBulk) SetLinkID(v uuid.UUID) *BlobLinkUpsertBulk {
	return u.Update(func(s *BlobLinkUpsert) {
		s.SetLinkID(v)
	})
}

// UpdateLinkID sets the "link_id" field to the value that was provided on create.
func (u *BlobLinkUpsertBulk) UpdateLinkID() *BlobLinkUpsertBulk {
	return u.Update(func(s *BlobLinkUpsert) {
		s.UpdateLinkID()
	})
}

// Exec executes the query.
func (u *BlobLinkUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the BlobLinkCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for BlobLinkCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *BlobLinkUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
