// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/TrixiS/mantra/ent/connection"
	"github.com/TrixiS/mantra/ent/user"
)

// ConnectionCreate is the builder for creating a Connection entity.
type ConnectionCreate struct {
	config
	mutation *ConnectionMutation
	hooks    []Hook
}

// SetLocalID sets the "local_id" field.
func (cc *ConnectionCreate) SetLocalID(i int) *ConnectionCreate {
	cc.mutation.SetLocalID(i)
	return cc
}

// SetName sets the "name" field.
func (cc *ConnectionCreate) SetName(s string) *ConnectionCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetHost sets the "host" field.
func (cc *ConnectionCreate) SetHost(s string) *ConnectionCreate {
	cc.mutation.SetHost(s)
	return cc
}

// SetPort sets the "port" field.
func (cc *ConnectionCreate) SetPort(u uint) *ConnectionCreate {
	cc.mutation.SetPort(u)
	return cc
}

// SetUser sets the "user" field.
func (cc *ConnectionCreate) SetUser(s string) *ConnectionCreate {
	cc.mutation.SetUser(s)
	return cc
}

// SetPassword sets the "password" field.
func (cc *ConnectionCreate) SetPassword(s string) *ConnectionCreate {
	cc.mutation.SetPassword(s)
	return cc
}

// SetArgs sets the "args" field.
func (cc *ConnectionCreate) SetArgs(s string) *ConnectionCreate {
	cc.mutation.SetArgs(s)
	return cc
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (cc *ConnectionCreate) SetOwnerID(id int) *ConnectionCreate {
	cc.mutation.SetOwnerID(id)
	return cc
}

// SetOwner sets the "owner" edge to the User entity.
func (cc *ConnectionCreate) SetOwner(u *User) *ConnectionCreate {
	return cc.SetOwnerID(u.ID)
}

// Mutation returns the ConnectionMutation object of the builder.
func (cc *ConnectionCreate) Mutation() *ConnectionMutation {
	return cc.mutation
}

// Save creates the Connection in the database.
func (cc *ConnectionCreate) Save(ctx context.Context) (*Connection, error) {
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *ConnectionCreate) SaveX(ctx context.Context) *Connection {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *ConnectionCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *ConnectionCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *ConnectionCreate) check() error {
	if _, ok := cc.mutation.LocalID(); !ok {
		return &ValidationError{Name: "local_id", err: errors.New(`ent: missing required field "Connection.local_id"`)}
	}
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Connection.name"`)}
	}
	if v, ok := cc.mutation.Name(); ok {
		if err := connection.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Connection.name": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Host(); !ok {
		return &ValidationError{Name: "host", err: errors.New(`ent: missing required field "Connection.host"`)}
	}
	if v, ok := cc.mutation.Host(); ok {
		if err := connection.HostValidator(v); err != nil {
			return &ValidationError{Name: "host", err: fmt.Errorf(`ent: validator failed for field "Connection.host": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Port(); !ok {
		return &ValidationError{Name: "port", err: errors.New(`ent: missing required field "Connection.port"`)}
	}
	if _, ok := cc.mutation.User(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required field "Connection.user"`)}
	}
	if v, ok := cc.mutation.User(); ok {
		if err := connection.UserValidator(v); err != nil {
			return &ValidationError{Name: "user", err: fmt.Errorf(`ent: validator failed for field "Connection.user": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New(`ent: missing required field "Connection.password"`)}
	}
	if v, ok := cc.mutation.Password(); ok {
		if err := connection.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "Connection.password": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Args(); !ok {
		return &ValidationError{Name: "args", err: errors.New(`ent: missing required field "Connection.args"`)}
	}
	if len(cc.mutation.OwnerIDs()) == 0 {
		return &ValidationError{Name: "owner", err: errors.New(`ent: missing required edge "Connection.owner"`)}
	}
	return nil
}

func (cc *ConnectionCreate) sqlSave(ctx context.Context) (*Connection, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *ConnectionCreate) createSpec() (*Connection, *sqlgraph.CreateSpec) {
	var (
		_node = &Connection{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(connection.Table, sqlgraph.NewFieldSpec(connection.FieldID, field.TypeInt))
	)
	if value, ok := cc.mutation.LocalID(); ok {
		_spec.SetField(connection.FieldLocalID, field.TypeInt, value)
		_node.LocalID = value
	}
	if value, ok := cc.mutation.Name(); ok {
		_spec.SetField(connection.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := cc.mutation.Host(); ok {
		_spec.SetField(connection.FieldHost, field.TypeString, value)
		_node.Host = value
	}
	if value, ok := cc.mutation.Port(); ok {
		_spec.SetField(connection.FieldPort, field.TypeUint, value)
		_node.Port = value
	}
	if value, ok := cc.mutation.User(); ok {
		_spec.SetField(connection.FieldUser, field.TypeString, value)
		_node.User = value
	}
	if value, ok := cc.mutation.Password(); ok {
		_spec.SetField(connection.FieldPassword, field.TypeString, value)
		_node.Password = value
	}
	if value, ok := cc.mutation.Args(); ok {
		_spec.SetField(connection.FieldArgs, field.TypeString, value)
		_node.Args = value
	}
	if nodes := cc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   connection.OwnerTable,
			Columns: []string{connection.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_connections = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ConnectionCreateBulk is the builder for creating many Connection entities in bulk.
type ConnectionCreateBulk struct {
	config
	err      error
	builders []*ConnectionCreate
}

// Save creates the Connection entities in the database.
func (ccb *ConnectionCreateBulk) Save(ctx context.Context) ([]*Connection, error) {
	if ccb.err != nil {
		return nil, ccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Connection, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ConnectionMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *ConnectionCreateBulk) SaveX(ctx context.Context) []*Connection {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *ConnectionCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *ConnectionCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
