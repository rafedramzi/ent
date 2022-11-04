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

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/edgeschema/ent/group"
	"entgo.io/ent/entc/integration/edgeschema/ent/user"
	"entgo.io/ent/entc/integration/edgeschema/ent/usergroup"
	"entgo.io/ent/schema/field"
)

// UserGroupCreate is the builder for creating a UserGroup entity.
type UserGroupCreate struct {
	config
	mutation *UserGroupMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetJoinedAt sets the "joined_at" field.
func (ugc *UserGroupCreate) SetJoinedAt(t time.Time) *UserGroupCreate {
	ugc.mutation.SetJoinedAt(t)
	return ugc
}

// SetNillableJoinedAt sets the "joined_at" field if the given value is not nil.
func (ugc *UserGroupCreate) SetNillableJoinedAt(t *time.Time) *UserGroupCreate {
	if t != nil {
		ugc.SetJoinedAt(*t)
	}
	return ugc
}

// SetUserID sets the "user_id" field.
func (ugc *UserGroupCreate) SetUserID(i int) *UserGroupCreate {
	ugc.mutation.SetUserID(i)
	return ugc
}

// SetGroupID sets the "group_id" field.
func (ugc *UserGroupCreate) SetGroupID(i int) *UserGroupCreate {
	ugc.mutation.SetGroupID(i)
	return ugc
}

// SetUser sets the "user" edge to the User entity.
func (ugc *UserGroupCreate) SetUser(u *User) *UserGroupCreate {
	return ugc.SetUserID(u.ID)
}

// SetGroup sets the "group" edge to the Group entity.
func (ugc *UserGroupCreate) SetGroup(g *Group) *UserGroupCreate {
	return ugc.SetGroupID(g.ID)
}

// Mutation returns the UserGroupMutation object of the builder.
func (ugc *UserGroupCreate) Mutation() *UserGroupMutation {
	return ugc.mutation
}

// Save creates the UserGroup in the database.
func (ugc *UserGroupCreate) Save(ctx context.Context) (*UserGroup, error) {
	var (
		err  error
		node *UserGroup
	)
	ugc.defaults()
	if len(ugc.hooks) == 0 {
		if err = ugc.check(); err != nil {
			return nil, err
		}
		node, err = ugc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserGroupMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ugc.check(); err != nil {
				return nil, err
			}
			ugc.mutation = mutation
			if node, err = ugc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ugc.hooks) - 1; i >= 0; i-- {
			if ugc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ugc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ugc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*UserGroup)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from UserGroupMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ugc *UserGroupCreate) SaveX(ctx context.Context) *UserGroup {
	v, err := ugc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ugc *UserGroupCreate) Exec(ctx context.Context) error {
	_, err := ugc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ugc *UserGroupCreate) ExecX(ctx context.Context) {
	if err := ugc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ugc *UserGroupCreate) defaults() {
	if _, ok := ugc.mutation.JoinedAt(); !ok {
		v := usergroup.DefaultJoinedAt()
		ugc.mutation.SetJoinedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ugc *UserGroupCreate) check() error {
	if _, ok := ugc.mutation.JoinedAt(); !ok {
		return &ValidationError{Name: "joined_at", err: errors.New(`ent: missing required field "UserGroup.joined_at"`)}
	}
	if _, ok := ugc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "UserGroup.user_id"`)}
	}
	if _, ok := ugc.mutation.GroupID(); !ok {
		return &ValidationError{Name: "group_id", err: errors.New(`ent: missing required field "UserGroup.group_id"`)}
	}
	if _, ok := ugc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "UserGroup.user"`)}
	}
	if _, ok := ugc.mutation.GroupID(); !ok {
		return &ValidationError{Name: "group", err: errors.New(`ent: missing required edge "UserGroup.group"`)}
	}
	return nil
}

func (ugc *UserGroupCreate) sqlSave(ctx context.Context) (*UserGroup, error) {
	_node, _spec := ugc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ugc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ugc *UserGroupCreate) createSpec() (*UserGroup, *sqlgraph.CreateSpec) {
	var (
		_node = &UserGroup{config: ugc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: usergroup.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: usergroup.FieldID,
			},
		}
	)
	_spec.OnConflict = ugc.conflict
	if value, ok := ugc.mutation.JoinedAt(); ok {
		_spec.SetField(usergroup.FieldJoinedAt, field.TypeTime, value)
		_node.JoinedAt = value
	}
	if nodes := ugc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   usergroup.UserTable,
			Columns: []string{usergroup.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ugc.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   usergroup.GroupTable,
			Columns: []string{usergroup.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: group.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.GroupID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.UserGroup.Create().
//		SetJoinedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserGroupUpsert) {
//			SetJoinedAt(v+v).
//		}).
//		Exec(ctx)
func (ugc *UserGroupCreate) OnConflict(opts ...sql.ConflictOption) *UserGroupUpsertOne {
	ugc.conflict = opts
	return &UserGroupUpsertOne{
		create: ugc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.UserGroup.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ugc *UserGroupCreate) OnConflictColumns(columns ...string) *UserGroupUpsertOne {
	ugc.conflict = append(ugc.conflict, sql.ConflictColumns(columns...))
	return &UserGroupUpsertOne{
		create: ugc,
	}
}

type (
	// UserGroupUpsertOne is the builder for "upsert"-ing
	//  one UserGroup node.
	UserGroupUpsertOne struct {
		create *UserGroupCreate
	}

	// UserGroupUpsert is the "OnConflict" setter.
	UserGroupUpsert struct {
		*sql.UpdateSet
	}
)

// SetJoinedAt sets the "joined_at" field.
func (u *UserGroupUpsert) SetJoinedAt(v time.Time) *UserGroupUpsert {
	u.Set(usergroup.FieldJoinedAt, v)
	return u
}

// UpdateJoinedAt sets the "joined_at" field to the value that was provided on create.
func (u *UserGroupUpsert) UpdateJoinedAt() *UserGroupUpsert {
	u.SetExcluded(usergroup.FieldJoinedAt)
	return u
}

// SetUserID sets the "user_id" field.
func (u *UserGroupUpsert) SetUserID(v int) *UserGroupUpsert {
	u.Set(usergroup.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *UserGroupUpsert) UpdateUserID() *UserGroupUpsert {
	u.SetExcluded(usergroup.FieldUserID)
	return u
}

// SetGroupID sets the "group_id" field.
func (u *UserGroupUpsert) SetGroupID(v int) *UserGroupUpsert {
	u.Set(usergroup.FieldGroupID, v)
	return u
}

// UpdateGroupID sets the "group_id" field to the value that was provided on create.
func (u *UserGroupUpsert) UpdateGroupID() *UserGroupUpsert {
	u.SetExcluded(usergroup.FieldGroupID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.UserGroup.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *UserGroupUpsertOne) UpdateNewValues() *UserGroupUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.UserGroup.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *UserGroupUpsertOne) Ignore() *UserGroupUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserGroupUpsertOne) DoNothing() *UserGroupUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserGroupCreate.OnConflict
// documentation for more info.
func (u *UserGroupUpsertOne) Update(set func(*UserGroupUpsert)) *UserGroupUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserGroupUpsert{UpdateSet: update})
	}))
	return u
}

// SetJoinedAt sets the "joined_at" field.
func (u *UserGroupUpsertOne) SetJoinedAt(v time.Time) *UserGroupUpsertOne {
	return u.Update(func(s *UserGroupUpsert) {
		s.SetJoinedAt(v)
	})
}

// UpdateJoinedAt sets the "joined_at" field to the value that was provided on create.
func (u *UserGroupUpsertOne) UpdateJoinedAt() *UserGroupUpsertOne {
	return u.Update(func(s *UserGroupUpsert) {
		s.UpdateJoinedAt()
	})
}

// SetUserID sets the "user_id" field.
func (u *UserGroupUpsertOne) SetUserID(v int) *UserGroupUpsertOne {
	return u.Update(func(s *UserGroupUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *UserGroupUpsertOne) UpdateUserID() *UserGroupUpsertOne {
	return u.Update(func(s *UserGroupUpsert) {
		s.UpdateUserID()
	})
}

// SetGroupID sets the "group_id" field.
func (u *UserGroupUpsertOne) SetGroupID(v int) *UserGroupUpsertOne {
	return u.Update(func(s *UserGroupUpsert) {
		s.SetGroupID(v)
	})
}

// UpdateGroupID sets the "group_id" field to the value that was provided on create.
func (u *UserGroupUpsertOne) UpdateGroupID() *UserGroupUpsertOne {
	return u.Update(func(s *UserGroupUpsert) {
		s.UpdateGroupID()
	})
}

// Exec executes the query.
func (u *UserGroupUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UserGroupCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserGroupUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID. Will return error on MYSQL dialect.
func (u *UserGroupUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *UserGroupUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// Save upsert the UserGroup in the database and returns the last inserted record. Will return error on MYSQL dialect.
func (u *UserGroupUpsertOne) Save(ctx context.Context) (*UserGroup, error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back the record
		// fields from the database since MySQL does not support the RETURNING clause.
		return nil, errors.New("ent: UserGroupUpsertOne.Save is not supported by MySQL driver. Use UserGroupUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return node, err
	}
	return node, nil
}

// SaveX calls Save and panics if Save returns an error.
func (u *UserGroupUpsertOne) SaveX(ctx context.Context) *UserGroup {
	node, err := u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// UserGroupCreateBulk is the builder for creating many UserGroup entities in bulk.
type UserGroupCreateBulk struct {
	config
	builders []*UserGroupCreate
	conflict []sql.ConflictOption
}

// Save creates the UserGroup entities in the database.
func (ugcb *UserGroupCreateBulk) Save(ctx context.Context) ([]*UserGroup, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ugcb.builders))
	nodes := make([]*UserGroup, len(ugcb.builders))
	mutators := make([]Mutator, len(ugcb.builders))
	for i := range ugcb.builders {
		func(i int, root context.Context) {
			builder := ugcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserGroupMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ugcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ugcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ugcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
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
		if _, err := mutators[0].Mutate(ctx, ugcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ugcb *UserGroupCreateBulk) SaveX(ctx context.Context) []*UserGroup {
	v, err := ugcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ugcb *UserGroupCreateBulk) Exec(ctx context.Context) error {
	_, err := ugcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ugcb *UserGroupCreateBulk) ExecX(ctx context.Context) {
	if err := ugcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.UserGroup.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserGroupUpsert) {
//			SetJoinedAt(v+v).
//		}).
//		Exec(ctx)
func (ugcb *UserGroupCreateBulk) OnConflict(opts ...sql.ConflictOption) *UserGroupUpsertBulk {
	ugcb.conflict = opts
	return &UserGroupUpsertBulk{
		create: ugcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.UserGroup.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ugcb *UserGroupCreateBulk) OnConflictColumns(columns ...string) *UserGroupUpsertBulk {
	ugcb.conflict = append(ugcb.conflict, sql.ConflictColumns(columns...))
	return &UserGroupUpsertBulk{
		create: ugcb,
	}
}

// UserGroupUpsertBulk is the builder for "upsert"-ing
// a bulk of UserGroup nodes.
type UserGroupUpsertBulk struct {
	create *UserGroupCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.UserGroup.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *UserGroupUpsertBulk) UpdateNewValues() *UserGroupUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.UserGroup.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *UserGroupUpsertBulk) Ignore() *UserGroupUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserGroupUpsertBulk) DoNothing() *UserGroupUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserGroupCreateBulk.OnConflict
// documentation for more info.
func (u *UserGroupUpsertBulk) Update(set func(*UserGroupUpsert)) *UserGroupUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserGroupUpsert{UpdateSet: update})
	}))
	return u
}

// SetJoinedAt sets the "joined_at" field.
func (u *UserGroupUpsertBulk) SetJoinedAt(v time.Time) *UserGroupUpsertBulk {
	return u.Update(func(s *UserGroupUpsert) {
		s.SetJoinedAt(v)
	})
}

// UpdateJoinedAt sets the "joined_at" field to the value that was provided on create.
func (u *UserGroupUpsertBulk) UpdateJoinedAt() *UserGroupUpsertBulk {
	return u.Update(func(s *UserGroupUpsert) {
		s.UpdateJoinedAt()
	})
}

// SetUserID sets the "user_id" field.
func (u *UserGroupUpsertBulk) SetUserID(v int) *UserGroupUpsertBulk {
	return u.Update(func(s *UserGroupUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *UserGroupUpsertBulk) UpdateUserID() *UserGroupUpsertBulk {
	return u.Update(func(s *UserGroupUpsert) {
		s.UpdateUserID()
	})
}

// SetGroupID sets the "group_id" field.
func (u *UserGroupUpsertBulk) SetGroupID(v int) *UserGroupUpsertBulk {
	return u.Update(func(s *UserGroupUpsert) {
		s.SetGroupID(v)
	})
}

// UpdateGroupID sets the "group_id" field to the value that was provided on create.
func (u *UserGroupUpsertBulk) UpdateGroupID() *UserGroupUpsertBulk {
	return u.Update(func(s *UserGroupUpsert) {
		s.UpdateGroupID()
	})
}

// Exec executes the query.
func (u *UserGroupUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the UserGroupCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UserGroupCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserGroupUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
