// Code generated by ent, DO NOT EDIT.

package connection

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/TrixiS/mantra/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Connection {
	return predicate.Connection(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Connection {
	return predicate.Connection(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Connection {
	return predicate.Connection(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Connection {
	return predicate.Connection(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Connection {
	return predicate.Connection(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Connection {
	return predicate.Connection(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Connection {
	return predicate.Connection(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Connection {
	return predicate.Connection(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Connection {
	return predicate.Connection(sql.FieldLTE(FieldID, id))
}

// LocalID applies equality check predicate on the "local_id" field. It's identical to LocalIDEQ.
func LocalID(v int) predicate.Connection {
	return predicate.Connection(sql.FieldEQ(FieldLocalID, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Connection {
	return predicate.Connection(sql.FieldEQ(FieldName, v))
}

// Host applies equality check predicate on the "host" field. It's identical to HostEQ.
func Host(v string) predicate.Connection {
	return predicate.Connection(sql.FieldEQ(FieldHost, v))
}

// Port applies equality check predicate on the "port" field. It's identical to PortEQ.
func Port(v uint) predicate.Connection {
	return predicate.Connection(sql.FieldEQ(FieldPort, v))
}

// User applies equality check predicate on the "user" field. It's identical to UserEQ.
func User(v string) predicate.Connection {
	return predicate.Connection(sql.FieldEQ(FieldUser, v))
}

// Password applies equality check predicate on the "password" field. It's identical to PasswordEQ.
func Password(v string) predicate.Connection {
	return predicate.Connection(sql.FieldEQ(FieldPassword, v))
}

// Args applies equality check predicate on the "args" field. It's identical to ArgsEQ.
func Args(v string) predicate.Connection {
	return predicate.Connection(sql.FieldEQ(FieldArgs, v))
}

// LocalIDEQ applies the EQ predicate on the "local_id" field.
func LocalIDEQ(v int) predicate.Connection {
	return predicate.Connection(sql.FieldEQ(FieldLocalID, v))
}

// LocalIDNEQ applies the NEQ predicate on the "local_id" field.
func LocalIDNEQ(v int) predicate.Connection {
	return predicate.Connection(sql.FieldNEQ(FieldLocalID, v))
}

// LocalIDIn applies the In predicate on the "local_id" field.
func LocalIDIn(vs ...int) predicate.Connection {
	return predicate.Connection(sql.FieldIn(FieldLocalID, vs...))
}

// LocalIDNotIn applies the NotIn predicate on the "local_id" field.
func LocalIDNotIn(vs ...int) predicate.Connection {
	return predicate.Connection(sql.FieldNotIn(FieldLocalID, vs...))
}

// LocalIDGT applies the GT predicate on the "local_id" field.
func LocalIDGT(v int) predicate.Connection {
	return predicate.Connection(sql.FieldGT(FieldLocalID, v))
}

// LocalIDGTE applies the GTE predicate on the "local_id" field.
func LocalIDGTE(v int) predicate.Connection {
	return predicate.Connection(sql.FieldGTE(FieldLocalID, v))
}

// LocalIDLT applies the LT predicate on the "local_id" field.
func LocalIDLT(v int) predicate.Connection {
	return predicate.Connection(sql.FieldLT(FieldLocalID, v))
}

// LocalIDLTE applies the LTE predicate on the "local_id" field.
func LocalIDLTE(v int) predicate.Connection {
	return predicate.Connection(sql.FieldLTE(FieldLocalID, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Connection {
	return predicate.Connection(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Connection {
	return predicate.Connection(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Connection {
	return predicate.Connection(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Connection {
	return predicate.Connection(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Connection {
	return predicate.Connection(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Connection {
	return predicate.Connection(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Connection {
	return predicate.Connection(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Connection {
	return predicate.Connection(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Connection {
	return predicate.Connection(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Connection {
	return predicate.Connection(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Connection {
	return predicate.Connection(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Connection {
	return predicate.Connection(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Connection {
	return predicate.Connection(sql.FieldContainsFold(FieldName, v))
}

// HostEQ applies the EQ predicate on the "host" field.
func HostEQ(v string) predicate.Connection {
	return predicate.Connection(sql.FieldEQ(FieldHost, v))
}

// HostNEQ applies the NEQ predicate on the "host" field.
func HostNEQ(v string) predicate.Connection {
	return predicate.Connection(sql.FieldNEQ(FieldHost, v))
}

// HostIn applies the In predicate on the "host" field.
func HostIn(vs ...string) predicate.Connection {
	return predicate.Connection(sql.FieldIn(FieldHost, vs...))
}

// HostNotIn applies the NotIn predicate on the "host" field.
func HostNotIn(vs ...string) predicate.Connection {
	return predicate.Connection(sql.FieldNotIn(FieldHost, vs...))
}

// HostGT applies the GT predicate on the "host" field.
func HostGT(v string) predicate.Connection {
	return predicate.Connection(sql.FieldGT(FieldHost, v))
}

// HostGTE applies the GTE predicate on the "host" field.
func HostGTE(v string) predicate.Connection {
	return predicate.Connection(sql.FieldGTE(FieldHost, v))
}

// HostLT applies the LT predicate on the "host" field.
func HostLT(v string) predicate.Connection {
	return predicate.Connection(sql.FieldLT(FieldHost, v))
}

// HostLTE applies the LTE predicate on the "host" field.
func HostLTE(v string) predicate.Connection {
	return predicate.Connection(sql.FieldLTE(FieldHost, v))
}

// HostContains applies the Contains predicate on the "host" field.
func HostContains(v string) predicate.Connection {
	return predicate.Connection(sql.FieldContains(FieldHost, v))
}

// HostHasPrefix applies the HasPrefix predicate on the "host" field.
func HostHasPrefix(v string) predicate.Connection {
	return predicate.Connection(sql.FieldHasPrefix(FieldHost, v))
}

// HostHasSuffix applies the HasSuffix predicate on the "host" field.
func HostHasSuffix(v string) predicate.Connection {
	return predicate.Connection(sql.FieldHasSuffix(FieldHost, v))
}

// HostEqualFold applies the EqualFold predicate on the "host" field.
func HostEqualFold(v string) predicate.Connection {
	return predicate.Connection(sql.FieldEqualFold(FieldHost, v))
}

// HostContainsFold applies the ContainsFold predicate on the "host" field.
func HostContainsFold(v string) predicate.Connection {
	return predicate.Connection(sql.FieldContainsFold(FieldHost, v))
}

// PortEQ applies the EQ predicate on the "port" field.
func PortEQ(v uint) predicate.Connection {
	return predicate.Connection(sql.FieldEQ(FieldPort, v))
}

// PortNEQ applies the NEQ predicate on the "port" field.
func PortNEQ(v uint) predicate.Connection {
	return predicate.Connection(sql.FieldNEQ(FieldPort, v))
}

// PortIn applies the In predicate on the "port" field.
func PortIn(vs ...uint) predicate.Connection {
	return predicate.Connection(sql.FieldIn(FieldPort, vs...))
}

// PortNotIn applies the NotIn predicate on the "port" field.
func PortNotIn(vs ...uint) predicate.Connection {
	return predicate.Connection(sql.FieldNotIn(FieldPort, vs...))
}

// PortGT applies the GT predicate on the "port" field.
func PortGT(v uint) predicate.Connection {
	return predicate.Connection(sql.FieldGT(FieldPort, v))
}

// PortGTE applies the GTE predicate on the "port" field.
func PortGTE(v uint) predicate.Connection {
	return predicate.Connection(sql.FieldGTE(FieldPort, v))
}

// PortLT applies the LT predicate on the "port" field.
func PortLT(v uint) predicate.Connection {
	return predicate.Connection(sql.FieldLT(FieldPort, v))
}

// PortLTE applies the LTE predicate on the "port" field.
func PortLTE(v uint) predicate.Connection {
	return predicate.Connection(sql.FieldLTE(FieldPort, v))
}

// UserEQ applies the EQ predicate on the "user" field.
func UserEQ(v string) predicate.Connection {
	return predicate.Connection(sql.FieldEQ(FieldUser, v))
}

// UserNEQ applies the NEQ predicate on the "user" field.
func UserNEQ(v string) predicate.Connection {
	return predicate.Connection(sql.FieldNEQ(FieldUser, v))
}

// UserIn applies the In predicate on the "user" field.
func UserIn(vs ...string) predicate.Connection {
	return predicate.Connection(sql.FieldIn(FieldUser, vs...))
}

// UserNotIn applies the NotIn predicate on the "user" field.
func UserNotIn(vs ...string) predicate.Connection {
	return predicate.Connection(sql.FieldNotIn(FieldUser, vs...))
}

// UserGT applies the GT predicate on the "user" field.
func UserGT(v string) predicate.Connection {
	return predicate.Connection(sql.FieldGT(FieldUser, v))
}

// UserGTE applies the GTE predicate on the "user" field.
func UserGTE(v string) predicate.Connection {
	return predicate.Connection(sql.FieldGTE(FieldUser, v))
}

// UserLT applies the LT predicate on the "user" field.
func UserLT(v string) predicate.Connection {
	return predicate.Connection(sql.FieldLT(FieldUser, v))
}

// UserLTE applies the LTE predicate on the "user" field.
func UserLTE(v string) predicate.Connection {
	return predicate.Connection(sql.FieldLTE(FieldUser, v))
}

// UserContains applies the Contains predicate on the "user" field.
func UserContains(v string) predicate.Connection {
	return predicate.Connection(sql.FieldContains(FieldUser, v))
}

// UserHasPrefix applies the HasPrefix predicate on the "user" field.
func UserHasPrefix(v string) predicate.Connection {
	return predicate.Connection(sql.FieldHasPrefix(FieldUser, v))
}

// UserHasSuffix applies the HasSuffix predicate on the "user" field.
func UserHasSuffix(v string) predicate.Connection {
	return predicate.Connection(sql.FieldHasSuffix(FieldUser, v))
}

// UserEqualFold applies the EqualFold predicate on the "user" field.
func UserEqualFold(v string) predicate.Connection {
	return predicate.Connection(sql.FieldEqualFold(FieldUser, v))
}

// UserContainsFold applies the ContainsFold predicate on the "user" field.
func UserContainsFold(v string) predicate.Connection {
	return predicate.Connection(sql.FieldContainsFold(FieldUser, v))
}

// PasswordEQ applies the EQ predicate on the "password" field.
func PasswordEQ(v string) predicate.Connection {
	return predicate.Connection(sql.FieldEQ(FieldPassword, v))
}

// PasswordNEQ applies the NEQ predicate on the "password" field.
func PasswordNEQ(v string) predicate.Connection {
	return predicate.Connection(sql.FieldNEQ(FieldPassword, v))
}

// PasswordIn applies the In predicate on the "password" field.
func PasswordIn(vs ...string) predicate.Connection {
	return predicate.Connection(sql.FieldIn(FieldPassword, vs...))
}

// PasswordNotIn applies the NotIn predicate on the "password" field.
func PasswordNotIn(vs ...string) predicate.Connection {
	return predicate.Connection(sql.FieldNotIn(FieldPassword, vs...))
}

// PasswordGT applies the GT predicate on the "password" field.
func PasswordGT(v string) predicate.Connection {
	return predicate.Connection(sql.FieldGT(FieldPassword, v))
}

// PasswordGTE applies the GTE predicate on the "password" field.
func PasswordGTE(v string) predicate.Connection {
	return predicate.Connection(sql.FieldGTE(FieldPassword, v))
}

// PasswordLT applies the LT predicate on the "password" field.
func PasswordLT(v string) predicate.Connection {
	return predicate.Connection(sql.FieldLT(FieldPassword, v))
}

// PasswordLTE applies the LTE predicate on the "password" field.
func PasswordLTE(v string) predicate.Connection {
	return predicate.Connection(sql.FieldLTE(FieldPassword, v))
}

// PasswordContains applies the Contains predicate on the "password" field.
func PasswordContains(v string) predicate.Connection {
	return predicate.Connection(sql.FieldContains(FieldPassword, v))
}

// PasswordHasPrefix applies the HasPrefix predicate on the "password" field.
func PasswordHasPrefix(v string) predicate.Connection {
	return predicate.Connection(sql.FieldHasPrefix(FieldPassword, v))
}

// PasswordHasSuffix applies the HasSuffix predicate on the "password" field.
func PasswordHasSuffix(v string) predicate.Connection {
	return predicate.Connection(sql.FieldHasSuffix(FieldPassword, v))
}

// PasswordEqualFold applies the EqualFold predicate on the "password" field.
func PasswordEqualFold(v string) predicate.Connection {
	return predicate.Connection(sql.FieldEqualFold(FieldPassword, v))
}

// PasswordContainsFold applies the ContainsFold predicate on the "password" field.
func PasswordContainsFold(v string) predicate.Connection {
	return predicate.Connection(sql.FieldContainsFold(FieldPassword, v))
}

// ArgsEQ applies the EQ predicate on the "args" field.
func ArgsEQ(v string) predicate.Connection {
	return predicate.Connection(sql.FieldEQ(FieldArgs, v))
}

// ArgsNEQ applies the NEQ predicate on the "args" field.
func ArgsNEQ(v string) predicate.Connection {
	return predicate.Connection(sql.FieldNEQ(FieldArgs, v))
}

// ArgsIn applies the In predicate on the "args" field.
func ArgsIn(vs ...string) predicate.Connection {
	return predicate.Connection(sql.FieldIn(FieldArgs, vs...))
}

// ArgsNotIn applies the NotIn predicate on the "args" field.
func ArgsNotIn(vs ...string) predicate.Connection {
	return predicate.Connection(sql.FieldNotIn(FieldArgs, vs...))
}

// ArgsGT applies the GT predicate on the "args" field.
func ArgsGT(v string) predicate.Connection {
	return predicate.Connection(sql.FieldGT(FieldArgs, v))
}

// ArgsGTE applies the GTE predicate on the "args" field.
func ArgsGTE(v string) predicate.Connection {
	return predicate.Connection(sql.FieldGTE(FieldArgs, v))
}

// ArgsLT applies the LT predicate on the "args" field.
func ArgsLT(v string) predicate.Connection {
	return predicate.Connection(sql.FieldLT(FieldArgs, v))
}

// ArgsLTE applies the LTE predicate on the "args" field.
func ArgsLTE(v string) predicate.Connection {
	return predicate.Connection(sql.FieldLTE(FieldArgs, v))
}

// ArgsContains applies the Contains predicate on the "args" field.
func ArgsContains(v string) predicate.Connection {
	return predicate.Connection(sql.FieldContains(FieldArgs, v))
}

// ArgsHasPrefix applies the HasPrefix predicate on the "args" field.
func ArgsHasPrefix(v string) predicate.Connection {
	return predicate.Connection(sql.FieldHasPrefix(FieldArgs, v))
}

// ArgsHasSuffix applies the HasSuffix predicate on the "args" field.
func ArgsHasSuffix(v string) predicate.Connection {
	return predicate.Connection(sql.FieldHasSuffix(FieldArgs, v))
}

// ArgsEqualFold applies the EqualFold predicate on the "args" field.
func ArgsEqualFold(v string) predicate.Connection {
	return predicate.Connection(sql.FieldEqualFold(FieldArgs, v))
}

// ArgsContainsFold applies the ContainsFold predicate on the "args" field.
func ArgsContainsFold(v string) predicate.Connection {
	return predicate.Connection(sql.FieldContainsFold(FieldArgs, v))
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.User) predicate.Connection {
	return predicate.Connection(func(s *sql.Selector) {
		step := newOwnerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Connection) predicate.Connection {
	return predicate.Connection(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Connection) predicate.Connection {
	return predicate.Connection(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Connection) predicate.Connection {
	return predicate.Connection(sql.NotPredicates(p))
}
