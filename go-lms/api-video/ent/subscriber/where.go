// Code generated by entc, DO NOT EDIT.

package subscriber

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"repo.nefrosovet.ru/go-lms/api-video/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Subscriber {
	return predicate.Subscriber(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldID), id))
		},
	)
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Subscriber {
	return predicate.Subscriber(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	},
	)
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Subscriber {
	return predicate.Subscriber(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	},
	)
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Subscriber {
	return predicate.Subscriber(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	},
	)
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Subscriber {
	return predicate.Subscriber(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	},
	)
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Subscriber {
	return predicate.Subscriber(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	},
	)
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Subscriber {
	return predicate.Subscriber(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	},
	)
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Subscriber {
	return predicate.Subscriber(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	},
	)
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Subscriber {
	return predicate.Subscriber(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	},
	)
}

// Username applies equality check predicate on the "username" field. It's identical to UsernameEQ.
func Username(v string) predicate.Subscriber {
	return predicate.Subscriber(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUsername), v))
	},
	)
}

// Domain applies equality check predicate on the "domain" field. It's identical to DomainEQ.
func Domain(v string) predicate.Subscriber {
	return predicate.Subscriber(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDomain), v))
	},
	)
}

// Ha1 applies equality check predicate on the "ha1" field. It's identical to Ha1EQ.
func Ha1(v string) predicate.Subscriber {
	return predicate.Subscriber(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHa1), v))
	},
	)
}

// Ha1b applies equality check predicate on the "ha1b" field. It's identical to Ha1bEQ.
func Ha1b(v string) predicate.Subscriber {
	return predicate.Subscriber(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHa1b), v))
	},
	)
}

// UsernameEQ applies the EQ predicate on the "username" field.
func UsernameEQ(v string) predicate.Subscriber {
	return predicate.Subscriber(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUsername), v))
	},
	)
}

// UsernameNEQ applies the NEQ predicate on the "username" field.
func UsernameNEQ(v string) predicate.Subscriber {
	return predicate.Subscriber(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUsername), v))
	},
	)
}

// UsernameIn applies the In predicate on the "username" field.
func UsernameIn(vs ...string) predicate.Subscriber {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Subscriber(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUsername), v...))
	},
	)
}

// UsernameNotIn applies the NotIn predicate on the "username" field.
func UsernameNotIn(vs ...string) predicate.Subscriber {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Subscriber(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUsername), v...))
	},
	)
}

// UsernameGT applies the GT predicate on the "username" field.
func UsernameGT(v string) predicate.Subscriber {
	return predicate.Subscriber(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUsername), v))
	},
	)
}

// UsernameGTE applies the GTE predicate on the "username" field.
func UsernameGTE(v string) predicate.Subscriber {
	return predicate.Subscriber(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUsername), v))
	},
	)
}

// UsernameLT applies the LT predicate on the "username" field.
func UsernameLT(v string) predicate.Subscriber {
	return predicate.Subscriber(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUsername), v))
	},
	)
}

// UsernameLTE applies the LTE predicate on the "username" field.
func UsernameLTE(v string) predicate.Subscriber {
	return predicate.Subscriber(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUsername), v))
	},
	)
}

// UsernameContains applies the Contains predicate on the "username" field.
func UsernameContains(v string) predicate.Subscriber {
	return predicate.Subscriber(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUsername), v))
	},
	)
}

// UsernameHasPrefix applies the HasPrefix predicate on the "username" field.
func UsernameHasPrefix(v string) predicate.Subscriber {
	return predicate.Subscriber(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUsername), v))
	},
	)
}

// UsernameHasSuffix applies the HasSuffix predicate on the "username" field.
func UsernameHasSuffix(v string) predicate.Subscriber {
	return predicate.Subscriber(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUsername), v))
	},
	)
}

// UsernameEqualFold applies the EqualFold predicate on the "username" field.
func UsernameEqualFold(v string) predicate.Subscriber {
	return predicate.Subscriber(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUsername), v))
	},
	)
}

// UsernameContainsFold applies the ContainsFold predicate on the "username" field.
func UsernameContainsFold(v string) predicate.Subscriber {
	return predicate.Subscriber(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUsername), v))
	},
	)
}

// DomainEQ applies the EQ predicate on the "domain" field.
func DomainEQ(v string) predicate.Subscriber {
	return predicate.Subscriber(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDomain), v))
	},
	)
}

// DomainNEQ applies the NEQ predicate on the "domain" field.
func DomainNEQ(v string) predicate.Subscriber {
	return predicate.Subscriber(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDomain), v))
	},
	)
}

// DomainIn applies the In predicate on the "domain" field.
func DomainIn(vs ...string) predicate.Subscriber {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Subscriber(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDomain), v...))
	},
	)
}

// DomainNotIn applies the NotIn predicate on the "domain" field.
func DomainNotIn(vs ...string) predicate.Subscriber {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Subscriber(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDomain), v...))
	},
	)
}

// DomainGT applies the GT predicate on the "domain" field.
func DomainGT(v string) predicate.Subscriber {
	return predicate.Subscriber(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDomain), v))
	},
	)
}

// DomainGTE applies the GTE predicate on the "domain" field.
func DomainGTE(v string) predicate.Subscriber {
	return predicate.Subscriber(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDomain), v))
	},
	)
}

// DomainLT applies the LT predicate on the "domain" field.
func DomainLT(v string) predicate.Subscriber {
	return predicate.Subscriber(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDomain), v))
	},
	)
}

// DomainLTE applies the LTE predicate on the "domain" field.
func DomainLTE(v string) predicate.Subscriber {
	return predicate.Subscriber(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDomain), v))
	},
	)
}

// DomainContains applies the Contains predicate on the "domain" field.
func DomainContains(v string) predicate.Subscriber {
	return predicate.Subscriber(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDomain), v))
	},
	)
}

// DomainHasPrefix applies the HasPrefix predicate on the "domain" field.
func DomainHasPrefix(v string) predicate.Subscriber {
	return predicate.Subscriber(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDomain), v))
	},
	)
}

// DomainHasSuffix applies the HasSuffix predicate on the "domain" field.
func DomainHasSuffix(v string) predicate.Subscriber {
	return predicate.Subscriber(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDomain), v))
	},
	)
}

// DomainEqualFold applies the EqualFold predicate on the "domain" field.
func DomainEqualFold(v string) predicate.Subscriber {
	return predicate.Subscriber(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDomain), v))
	},
	)
}

// DomainContainsFold applies the ContainsFold predicate on the "domain" field.
func DomainContainsFold(v string) predicate.Subscriber {
	return predicate.Subscriber(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDomain), v))
	},
	)
}

// Ha1EQ applies the EQ predicate on the "ha1" field.
func Ha1EQ(v string) predicate.Subscriber {
	return predicate.Subscriber(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHa1), v))
	},
	)
}

// Ha1NEQ applies the NEQ predicate on the "ha1" field.
func Ha1NEQ(v string) predicate.Subscriber {
	return predicate.Subscriber(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHa1), v))
	},
	)
}

// Ha1In applies the In predicate on the "ha1" field.
func Ha1In(vs ...string) predicate.Subscriber {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Subscriber(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldHa1), v...))
	},
	)
}

// Ha1NotIn applies the NotIn predicate on the "ha1" field.
func Ha1NotIn(vs ...string) predicate.Subscriber {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Subscriber(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldHa1), v...))
	},
	)
}

// Ha1GT applies the GT predicate on the "ha1" field.
func Ha1GT(v string) predicate.Subscriber {
	return predicate.Subscriber(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldHa1), v))
	},
	)
}

// Ha1GTE applies the GTE predicate on the "ha1" field.
func Ha1GTE(v string) predicate.Subscriber {
	return predicate.Subscriber(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldHa1), v))
	},
	)
}

// Ha1LT applies the LT predicate on the "ha1" field.
func Ha1LT(v string) predicate.Subscriber {
	return predicate.Subscriber(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldHa1), v))
	},
	)
}

// Ha1LTE applies the LTE predicate on the "ha1" field.
func Ha1LTE(v string) predicate.Subscriber {
	return predicate.Subscriber(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldHa1), v))
	},
	)
}

// Ha1Contains applies the Contains predicate on the "ha1" field.
func Ha1Contains(v string) predicate.Subscriber {
	return predicate.Subscriber(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldHa1), v))
	},
	)
}

// Ha1HasPrefix applies the HasPrefix predicate on the "ha1" field.
func Ha1HasPrefix(v string) predicate.Subscriber {
	return predicate.Subscriber(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldHa1), v))
	},
	)
}

// Ha1HasSuffix applies the HasSuffix predicate on the "ha1" field.
func Ha1HasSuffix(v string) predicate.Subscriber {
	return predicate.Subscriber(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldHa1), v))
	},
	)
}

// Ha1EqualFold applies the EqualFold predicate on the "ha1" field.
func Ha1EqualFold(v string) predicate.Subscriber {
	return predicate.Subscriber(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldHa1), v))
	},
	)
}

// Ha1ContainsFold applies the ContainsFold predicate on the "ha1" field.
func Ha1ContainsFold(v string) predicate.Subscriber {
	return predicate.Subscriber(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldHa1), v))
	},
	)
}

// Ha1bEQ applies the EQ predicate on the "ha1b" field.
func Ha1bEQ(v string) predicate.Subscriber {
	return predicate.Subscriber(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHa1b), v))
	},
	)
}

// Ha1bNEQ applies the NEQ predicate on the "ha1b" field.
func Ha1bNEQ(v string) predicate.Subscriber {
	return predicate.Subscriber(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHa1b), v))
	},
	)
}

// Ha1bIn applies the In predicate on the "ha1b" field.
func Ha1bIn(vs ...string) predicate.Subscriber {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Subscriber(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldHa1b), v...))
	},
	)
}

// Ha1bNotIn applies the NotIn predicate on the "ha1b" field.
func Ha1bNotIn(vs ...string) predicate.Subscriber {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Subscriber(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldHa1b), v...))
	},
	)
}

// Ha1bGT applies the GT predicate on the "ha1b" field.
func Ha1bGT(v string) predicate.Subscriber {
	return predicate.Subscriber(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldHa1b), v))
	},
	)
}

// Ha1bGTE applies the GTE predicate on the "ha1b" field.
func Ha1bGTE(v string) predicate.Subscriber {
	return predicate.Subscriber(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldHa1b), v))
	},
	)
}

// Ha1bLT applies the LT predicate on the "ha1b" field.
func Ha1bLT(v string) predicate.Subscriber {
	return predicate.Subscriber(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldHa1b), v))
	},
	)
}

// Ha1bLTE applies the LTE predicate on the "ha1b" field.
func Ha1bLTE(v string) predicate.Subscriber {
	return predicate.Subscriber(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldHa1b), v))
	},
	)
}

// Ha1bContains applies the Contains predicate on the "ha1b" field.
func Ha1bContains(v string) predicate.Subscriber {
	return predicate.Subscriber(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldHa1b), v))
	},
	)
}

// Ha1bHasPrefix applies the HasPrefix predicate on the "ha1b" field.
func Ha1bHasPrefix(v string) predicate.Subscriber {
	return predicate.Subscriber(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldHa1b), v))
	},
	)
}

// Ha1bHasSuffix applies the HasSuffix predicate on the "ha1b" field.
func Ha1bHasSuffix(v string) predicate.Subscriber {
	return predicate.Subscriber(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldHa1b), v))
	},
	)
}

// Ha1bEqualFold applies the EqualFold predicate on the "ha1b" field.
func Ha1bEqualFold(v string) predicate.Subscriber {
	return predicate.Subscriber(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldHa1b), v))
	},
	)
}

// Ha1bContainsFold applies the ContainsFold predicate on the "ha1b" field.
func Ha1bContainsFold(v string) predicate.Subscriber {
	return predicate.Subscriber(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldHa1b), v))
	},
	)
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Subscriber {
	return predicate.Subscriber(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	},
	)
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Subscriber {
	return predicate.Subscriber(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	},
	)
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Subscriber) predicate.Subscriber {
	return predicate.Subscriber(
		func(s *sql.Selector) {
			s1 := s.Clone().SetP(nil)
			for _, p := range predicates {
				p(s1)
			}
			s.Where(s1.P())
		},
	)
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Subscriber) predicate.Subscriber {
	return predicate.Subscriber(
		func(s *sql.Selector) {
			s1 := s.Clone().SetP(nil)
			for i, p := range predicates {
				if i > 0 {
					s1.Or()
				}
				p(s1)
			}
			s.Where(s1.P())
		},
	)
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Subscriber) predicate.Subscriber {
	return predicate.Subscriber(
		func(s *sql.Selector) {
			p(s.Not())
		},
	)
}