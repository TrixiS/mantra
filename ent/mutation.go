// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/TrixiS/mantra/ent/connection"
	"github.com/TrixiS/mantra/ent/predicate"
	"github.com/TrixiS/mantra/ent/user"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeConnection = "Connection"
	TypeUser       = "User"
)

// ConnectionMutation represents an operation that mutates the Connection nodes in the graph.
type ConnectionMutation struct {
	config
	op            Op
	typ           string
	id            *int
	local_id      *int
	addlocal_id   *int
	name          *string
	host          *string
	port          *uint
	addport       *int
	user          *string
	password      *string
	args          *string
	clearedFields map[string]struct{}
	owner         *int
	clearedowner  bool
	done          bool
	oldValue      func(context.Context) (*Connection, error)
	predicates    []predicate.Connection
}

var _ ent.Mutation = (*ConnectionMutation)(nil)

// connectionOption allows management of the mutation configuration using functional options.
type connectionOption func(*ConnectionMutation)

// newConnectionMutation creates new mutation for the Connection entity.
func newConnectionMutation(c config, op Op, opts ...connectionOption) *ConnectionMutation {
	m := &ConnectionMutation{
		config:        c,
		op:            op,
		typ:           TypeConnection,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withConnectionID sets the ID field of the mutation.
func withConnectionID(id int) connectionOption {
	return func(m *ConnectionMutation) {
		var (
			err   error
			once  sync.Once
			value *Connection
		)
		m.oldValue = func(ctx context.Context) (*Connection, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Connection.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withConnection sets the old Connection of the mutation.
func withConnection(node *Connection) connectionOption {
	return func(m *ConnectionMutation) {
		m.oldValue = func(context.Context) (*Connection, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m ConnectionMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m ConnectionMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *ConnectionMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *ConnectionMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Connection.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetLocalID sets the "local_id" field.
func (m *ConnectionMutation) SetLocalID(i int) {
	m.local_id = &i
	m.addlocal_id = nil
}

// LocalID returns the value of the "local_id" field in the mutation.
func (m *ConnectionMutation) LocalID() (r int, exists bool) {
	v := m.local_id
	if v == nil {
		return
	}
	return *v, true
}

// OldLocalID returns the old "local_id" field's value of the Connection entity.
// If the Connection object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ConnectionMutation) OldLocalID(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldLocalID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldLocalID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldLocalID: %w", err)
	}
	return oldValue.LocalID, nil
}

// AddLocalID adds i to the "local_id" field.
func (m *ConnectionMutation) AddLocalID(i int) {
	if m.addlocal_id != nil {
		*m.addlocal_id += i
	} else {
		m.addlocal_id = &i
	}
}

// AddedLocalID returns the value that was added to the "local_id" field in this mutation.
func (m *ConnectionMutation) AddedLocalID() (r int, exists bool) {
	v := m.addlocal_id
	if v == nil {
		return
	}
	return *v, true
}

// ResetLocalID resets all changes to the "local_id" field.
func (m *ConnectionMutation) ResetLocalID() {
	m.local_id = nil
	m.addlocal_id = nil
}

// SetName sets the "name" field.
func (m *ConnectionMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *ConnectionMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Connection entity.
// If the Connection object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ConnectionMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *ConnectionMutation) ResetName() {
	m.name = nil
}

// SetHost sets the "host" field.
func (m *ConnectionMutation) SetHost(s string) {
	m.host = &s
}

// Host returns the value of the "host" field in the mutation.
func (m *ConnectionMutation) Host() (r string, exists bool) {
	v := m.host
	if v == nil {
		return
	}
	return *v, true
}

// OldHost returns the old "host" field's value of the Connection entity.
// If the Connection object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ConnectionMutation) OldHost(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldHost is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldHost requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldHost: %w", err)
	}
	return oldValue.Host, nil
}

// ResetHost resets all changes to the "host" field.
func (m *ConnectionMutation) ResetHost() {
	m.host = nil
}

// SetPort sets the "port" field.
func (m *ConnectionMutation) SetPort(u uint) {
	m.port = &u
	m.addport = nil
}

// Port returns the value of the "port" field in the mutation.
func (m *ConnectionMutation) Port() (r uint, exists bool) {
	v := m.port
	if v == nil {
		return
	}
	return *v, true
}

// OldPort returns the old "port" field's value of the Connection entity.
// If the Connection object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ConnectionMutation) OldPort(ctx context.Context) (v uint, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPort is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPort requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPort: %w", err)
	}
	return oldValue.Port, nil
}

// AddPort adds u to the "port" field.
func (m *ConnectionMutation) AddPort(u int) {
	if m.addport != nil {
		*m.addport += u
	} else {
		m.addport = &u
	}
}

// AddedPort returns the value that was added to the "port" field in this mutation.
func (m *ConnectionMutation) AddedPort() (r int, exists bool) {
	v := m.addport
	if v == nil {
		return
	}
	return *v, true
}

// ResetPort resets all changes to the "port" field.
func (m *ConnectionMutation) ResetPort() {
	m.port = nil
	m.addport = nil
}

// SetUser sets the "user" field.
func (m *ConnectionMutation) SetUser(s string) {
	m.user = &s
}

// User returns the value of the "user" field in the mutation.
func (m *ConnectionMutation) User() (r string, exists bool) {
	v := m.user
	if v == nil {
		return
	}
	return *v, true
}

// OldUser returns the old "user" field's value of the Connection entity.
// If the Connection object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ConnectionMutation) OldUser(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUser is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUser requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUser: %w", err)
	}
	return oldValue.User, nil
}

// ResetUser resets all changes to the "user" field.
func (m *ConnectionMutation) ResetUser() {
	m.user = nil
}

// SetPassword sets the "password" field.
func (m *ConnectionMutation) SetPassword(s string) {
	m.password = &s
}

// Password returns the value of the "password" field in the mutation.
func (m *ConnectionMutation) Password() (r string, exists bool) {
	v := m.password
	if v == nil {
		return
	}
	return *v, true
}

// OldPassword returns the old "password" field's value of the Connection entity.
// If the Connection object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ConnectionMutation) OldPassword(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPassword is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPassword requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPassword: %w", err)
	}
	return oldValue.Password, nil
}

// ResetPassword resets all changes to the "password" field.
func (m *ConnectionMutation) ResetPassword() {
	m.password = nil
}

// SetArgs sets the "args" field.
func (m *ConnectionMutation) SetArgs(s string) {
	m.args = &s
}

// Args returns the value of the "args" field in the mutation.
func (m *ConnectionMutation) Args() (r string, exists bool) {
	v := m.args
	if v == nil {
		return
	}
	return *v, true
}

// OldArgs returns the old "args" field's value of the Connection entity.
// If the Connection object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ConnectionMutation) OldArgs(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldArgs is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldArgs requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldArgs: %w", err)
	}
	return oldValue.Args, nil
}

// ResetArgs resets all changes to the "args" field.
func (m *ConnectionMutation) ResetArgs() {
	m.args = nil
}

// SetOwnerID sets the "owner" edge to the User entity by id.
func (m *ConnectionMutation) SetOwnerID(id int) {
	m.owner = &id
}

// ClearOwner clears the "owner" edge to the User entity.
func (m *ConnectionMutation) ClearOwner() {
	m.clearedowner = true
}

// OwnerCleared reports if the "owner" edge to the User entity was cleared.
func (m *ConnectionMutation) OwnerCleared() bool {
	return m.clearedowner
}

// OwnerID returns the "owner" edge ID in the mutation.
func (m *ConnectionMutation) OwnerID() (id int, exists bool) {
	if m.owner != nil {
		return *m.owner, true
	}
	return
}

// OwnerIDs returns the "owner" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// OwnerID instead. It exists only for internal usage by the builders.
func (m *ConnectionMutation) OwnerIDs() (ids []int) {
	if id := m.owner; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetOwner resets all changes to the "owner" edge.
func (m *ConnectionMutation) ResetOwner() {
	m.owner = nil
	m.clearedowner = false
}

// Where appends a list predicates to the ConnectionMutation builder.
func (m *ConnectionMutation) Where(ps ...predicate.Connection) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the ConnectionMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *ConnectionMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Connection, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *ConnectionMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *ConnectionMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Connection).
func (m *ConnectionMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *ConnectionMutation) Fields() []string {
	fields := make([]string, 0, 7)
	if m.local_id != nil {
		fields = append(fields, connection.FieldLocalID)
	}
	if m.name != nil {
		fields = append(fields, connection.FieldName)
	}
	if m.host != nil {
		fields = append(fields, connection.FieldHost)
	}
	if m.port != nil {
		fields = append(fields, connection.FieldPort)
	}
	if m.user != nil {
		fields = append(fields, connection.FieldUser)
	}
	if m.password != nil {
		fields = append(fields, connection.FieldPassword)
	}
	if m.args != nil {
		fields = append(fields, connection.FieldArgs)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *ConnectionMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case connection.FieldLocalID:
		return m.LocalID()
	case connection.FieldName:
		return m.Name()
	case connection.FieldHost:
		return m.Host()
	case connection.FieldPort:
		return m.Port()
	case connection.FieldUser:
		return m.User()
	case connection.FieldPassword:
		return m.Password()
	case connection.FieldArgs:
		return m.Args()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *ConnectionMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case connection.FieldLocalID:
		return m.OldLocalID(ctx)
	case connection.FieldName:
		return m.OldName(ctx)
	case connection.FieldHost:
		return m.OldHost(ctx)
	case connection.FieldPort:
		return m.OldPort(ctx)
	case connection.FieldUser:
		return m.OldUser(ctx)
	case connection.FieldPassword:
		return m.OldPassword(ctx)
	case connection.FieldArgs:
		return m.OldArgs(ctx)
	}
	return nil, fmt.Errorf("unknown Connection field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ConnectionMutation) SetField(name string, value ent.Value) error {
	switch name {
	case connection.FieldLocalID:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLocalID(v)
		return nil
	case connection.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case connection.FieldHost:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetHost(v)
		return nil
	case connection.FieldPort:
		v, ok := value.(uint)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPort(v)
		return nil
	case connection.FieldUser:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUser(v)
		return nil
	case connection.FieldPassword:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPassword(v)
		return nil
	case connection.FieldArgs:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetArgs(v)
		return nil
	}
	return fmt.Errorf("unknown Connection field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *ConnectionMutation) AddedFields() []string {
	var fields []string
	if m.addlocal_id != nil {
		fields = append(fields, connection.FieldLocalID)
	}
	if m.addport != nil {
		fields = append(fields, connection.FieldPort)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *ConnectionMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case connection.FieldLocalID:
		return m.AddedLocalID()
	case connection.FieldPort:
		return m.AddedPort()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ConnectionMutation) AddField(name string, value ent.Value) error {
	switch name {
	case connection.FieldLocalID:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddLocalID(v)
		return nil
	case connection.FieldPort:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddPort(v)
		return nil
	}
	return fmt.Errorf("unknown Connection numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *ConnectionMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *ConnectionMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *ConnectionMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Connection nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *ConnectionMutation) ResetField(name string) error {
	switch name {
	case connection.FieldLocalID:
		m.ResetLocalID()
		return nil
	case connection.FieldName:
		m.ResetName()
		return nil
	case connection.FieldHost:
		m.ResetHost()
		return nil
	case connection.FieldPort:
		m.ResetPort()
		return nil
	case connection.FieldUser:
		m.ResetUser()
		return nil
	case connection.FieldPassword:
		m.ResetPassword()
		return nil
	case connection.FieldArgs:
		m.ResetArgs()
		return nil
	}
	return fmt.Errorf("unknown Connection field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *ConnectionMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.owner != nil {
		edges = append(edges, connection.EdgeOwner)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *ConnectionMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case connection.EdgeOwner:
		if id := m.owner; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *ConnectionMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *ConnectionMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *ConnectionMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedowner {
		edges = append(edges, connection.EdgeOwner)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *ConnectionMutation) EdgeCleared(name string) bool {
	switch name {
	case connection.EdgeOwner:
		return m.clearedowner
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *ConnectionMutation) ClearEdge(name string) error {
	switch name {
	case connection.EdgeOwner:
		m.ClearOwner()
		return nil
	}
	return fmt.Errorf("unknown Connection unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *ConnectionMutation) ResetEdge(name string) error {
	switch name {
	case connection.EdgeOwner:
		m.ResetOwner()
		return nil
	}
	return fmt.Errorf("unknown Connection edge %s", name)
}

// UserMutation represents an operation that mutates the User nodes in the graph.
type UserMutation struct {
	config
	op                 Op
	typ                string
	id                 *int
	username           *string
	password_hash      *[]byte
	clearedFields      map[string]struct{}
	connections        map[int]struct{}
	removedconnections map[int]struct{}
	clearedconnections bool
	done               bool
	oldValue           func(context.Context) (*User, error)
	predicates         []predicate.User
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows management of the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for the User entity.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the ID field of the mutation.
func withUserID(id int) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *UserMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().User.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetUsername sets the "username" field.
func (m *UserMutation) SetUsername(s string) {
	m.username = &s
}

// Username returns the value of the "username" field in the mutation.
func (m *UserMutation) Username() (r string, exists bool) {
	v := m.username
	if v == nil {
		return
	}
	return *v, true
}

// OldUsername returns the old "username" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldUsername(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUsername is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUsername requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUsername: %w", err)
	}
	return oldValue.Username, nil
}

// ResetUsername resets all changes to the "username" field.
func (m *UserMutation) ResetUsername() {
	m.username = nil
}

// SetPasswordHash sets the "password_hash" field.
func (m *UserMutation) SetPasswordHash(b []byte) {
	m.password_hash = &b
}

// PasswordHash returns the value of the "password_hash" field in the mutation.
func (m *UserMutation) PasswordHash() (r []byte, exists bool) {
	v := m.password_hash
	if v == nil {
		return
	}
	return *v, true
}

// OldPasswordHash returns the old "password_hash" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldPasswordHash(ctx context.Context) (v []byte, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPasswordHash is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPasswordHash requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPasswordHash: %w", err)
	}
	return oldValue.PasswordHash, nil
}

// ResetPasswordHash resets all changes to the "password_hash" field.
func (m *UserMutation) ResetPasswordHash() {
	m.password_hash = nil
}

// AddConnectionIDs adds the "connections" edge to the Connection entity by ids.
func (m *UserMutation) AddConnectionIDs(ids ...int) {
	if m.connections == nil {
		m.connections = make(map[int]struct{})
	}
	for i := range ids {
		m.connections[ids[i]] = struct{}{}
	}
}

// ClearConnections clears the "connections" edge to the Connection entity.
func (m *UserMutation) ClearConnections() {
	m.clearedconnections = true
}

// ConnectionsCleared reports if the "connections" edge to the Connection entity was cleared.
func (m *UserMutation) ConnectionsCleared() bool {
	return m.clearedconnections
}

// RemoveConnectionIDs removes the "connections" edge to the Connection entity by IDs.
func (m *UserMutation) RemoveConnectionIDs(ids ...int) {
	if m.removedconnections == nil {
		m.removedconnections = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.connections, ids[i])
		m.removedconnections[ids[i]] = struct{}{}
	}
}

// RemovedConnections returns the removed IDs of the "connections" edge to the Connection entity.
func (m *UserMutation) RemovedConnectionsIDs() (ids []int) {
	for id := range m.removedconnections {
		ids = append(ids, id)
	}
	return
}

// ConnectionsIDs returns the "connections" edge IDs in the mutation.
func (m *UserMutation) ConnectionsIDs() (ids []int) {
	for id := range m.connections {
		ids = append(ids, id)
	}
	return
}

// ResetConnections resets all changes to the "connections" edge.
func (m *UserMutation) ResetConnections() {
	m.connections = nil
	m.clearedconnections = false
	m.removedconnections = nil
}

// Where appends a list predicates to the UserMutation builder.
func (m *UserMutation) Where(ps ...predicate.User) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the UserMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *UserMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.User, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *UserMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.username != nil {
		fields = append(fields, user.FieldUsername)
	}
	if m.password_hash != nil {
		fields = append(fields, user.FieldPasswordHash)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldUsername:
		return m.Username()
	case user.FieldPasswordHash:
		return m.PasswordHash()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldUsername:
		return m.OldUsername(ctx)
	case user.FieldPasswordHash:
		return m.OldPasswordHash(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldUsername:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUsername(v)
		return nil
	case user.FieldPasswordHash:
		v, ok := value.([]byte)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPasswordHash(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldUsername:
		m.ResetUsername()
		return nil
	case user.FieldPasswordHash:
		m.ResetPasswordHash()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.connections != nil {
		edges = append(edges, user.EdgeConnections)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeConnections:
		ids := make([]ent.Value, 0, len(m.connections))
		for id := range m.connections {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedconnections != nil {
		edges = append(edges, user.EdgeConnections)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeConnections:
		ids := make([]ent.Value, 0, len(m.removedconnections))
		for id := range m.removedconnections {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedconnections {
		edges = append(edges, user.EdgeConnections)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	switch name {
	case user.EdgeConnections:
		return m.clearedconnections
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	switch name {
	case user.EdgeConnections:
		m.ResetConnections()
		return nil
	}
	return fmt.Errorf("unknown User edge %s", name)
}
