// Code generated by ent, DO NOT EDIT.

package connection

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the connection type in the database.
	Label = "connection"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldLocalID holds the string denoting the local_id field in the database.
	FieldLocalID = "local_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldHost holds the string denoting the host field in the database.
	FieldHost = "host"
	// FieldPort holds the string denoting the port field in the database.
	FieldPort = "port"
	// FieldUser holds the string denoting the user field in the database.
	FieldUser = "user"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldArgs holds the string denoting the args field in the database.
	FieldArgs = "args"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// Table holds the table name of the connection in the database.
	Table = "connections"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "connections"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "user_connections"
)

// Columns holds all SQL columns for connection fields.
var Columns = []string{
	FieldID,
	FieldLocalID,
	FieldName,
	FieldHost,
	FieldPort,
	FieldUser,
	FieldPassword,
	FieldArgs,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "connections"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_connections",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// HostValidator is a validator for the "host" field. It is called by the builders before save.
	HostValidator func(string) error
	// UserValidator is a validator for the "user" field. It is called by the builders before save.
	UserValidator func(string) error
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator func(string) error
)

// OrderOption defines the ordering options for the Connection queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByLocalID orders the results by the local_id field.
func ByLocalID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLocalID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByHost orders the results by the host field.
func ByHost(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHost, opts...).ToFunc()
}

// ByPort orders the results by the port field.
func ByPort(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPort, opts...).ToFunc()
}

// ByUser orders the results by the user field.
func ByUser(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUser, opts...).ToFunc()
}

// ByPassword orders the results by the password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// ByArgs orders the results by the args field.
func ByArgs(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldArgs, opts...).ToFunc()
}

// ByOwnerField orders the results by owner field.
func ByOwnerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOwnerStep(), sql.OrderByField(field, opts...))
	}
}
func newOwnerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OwnerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
	)
}
