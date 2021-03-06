// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"repo.nefrosovet.ru/go-lms/api-video/ent/subscriber"
	"repo.nefrosovet.ru/go-lms/api-video/ent/user"
	"repo.nefrosovet.ru/go-lms/api-video/ent/useraccount"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	created_at  *time.Time
	updated_at  *time.Time
	meta_data   *string
	subscriber  map[int]struct{}
	useraccount map[int]struct{}
}

// SetCreatedAt sets the created_at field.
func (uc *UserCreate) SetCreatedAt(t time.Time) *UserCreate {
	uc.created_at = &t
	return uc
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (uc *UserCreate) SetNillableCreatedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetCreatedAt(*t)
	}
	return uc
}

// SetUpdatedAt sets the updated_at field.
func (uc *UserCreate) SetUpdatedAt(t time.Time) *UserCreate {
	uc.updated_at = &t
	return uc
}

// SetNillableUpdatedAt sets the updated_at field if the given value is not nil.
func (uc *UserCreate) SetNillableUpdatedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetUpdatedAt(*t)
	}
	return uc
}

// SetMetaData sets the meta_data field.
func (uc *UserCreate) SetMetaData(s string) *UserCreate {
	uc.meta_data = &s
	return uc
}

// SetSubscriberID sets the subscriber edge to Subscriber by id.
func (uc *UserCreate) SetSubscriberID(id int) *UserCreate {
	if uc.subscriber == nil {
		uc.subscriber = make(map[int]struct{})
	}
	uc.subscriber[id] = struct{}{}
	return uc
}

// SetNillableSubscriberID sets the subscriber edge to Subscriber by id if the given value is not nil.
func (uc *UserCreate) SetNillableSubscriberID(id *int) *UserCreate {
	if id != nil {
		uc = uc.SetSubscriberID(*id)
	}
	return uc
}

// SetSubscriber sets the subscriber edge to Subscriber.
func (uc *UserCreate) SetSubscriber(s *Subscriber) *UserCreate {
	return uc.SetSubscriberID(s.ID)
}

// SetUseraccountID sets the useraccount edge to UserAccount by id.
func (uc *UserCreate) SetUseraccountID(id int) *UserCreate {
	if uc.useraccount == nil {
		uc.useraccount = make(map[int]struct{})
	}
	uc.useraccount[id] = struct{}{}
	return uc
}

// SetNillableUseraccountID sets the useraccount edge to UserAccount by id if the given value is not nil.
func (uc *UserCreate) SetNillableUseraccountID(id *int) *UserCreate {
	if id != nil {
		uc = uc.SetUseraccountID(*id)
	}
	return uc
}

// SetUseraccount sets the useraccount edge to UserAccount.
func (uc *UserCreate) SetUseraccount(u *UserAccount) *UserCreate {
	return uc.SetUseraccountID(u.ID)
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	if uc.created_at == nil {
		v := user.DefaultCreatedAt()
		uc.created_at = &v
	}
	if uc.meta_data == nil {
		return nil, errors.New("ent: missing required field \"meta_data\"")
	}
	if len(uc.subscriber) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"subscriber\"")
	}
	if len(uc.useraccount) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"useraccount\"")
	}
	return uc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	var (
		u     = &User{config: uc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: user.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		}
	)
	if value := uc.created_at; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: user.FieldCreatedAt,
		})
		u.CreatedAt = value
	}
	if value := uc.updated_at; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: user.FieldUpdatedAt,
		})
		u.UpdatedAt = value
	}
	if value := uc.meta_data; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldMetaData,
		})
		u.MetaData = *value
	}
	if nodes := uc.subscriber; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.SubscriberTable,
			Columns: []string{user.SubscriberColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: subscriber.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.useraccount; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.UseraccountTable,
			Columns: []string{user.UseraccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: useraccount.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	u.ID = int(id)
	return u, nil
}
