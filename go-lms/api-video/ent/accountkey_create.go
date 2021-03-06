// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"repo.nefrosovet.ru/go-lms/api-video/ent/accountkey"
	"repo.nefrosovet.ru/go-lms/api-video/ent/useraccount"
)

// AccountKeyCreate is the builder for creating a AccountKey entity.
type AccountKeyCreate struct {
	config
	created_at  *time.Time
	updated_at  *time.Time
	account_id  *int16
	key         *string
	options     *string
	meta_data   *string
	useraccount map[int]struct{}
}

// SetCreatedAt sets the created_at field.
func (akc *AccountKeyCreate) SetCreatedAt(t time.Time) *AccountKeyCreate {
	akc.created_at = &t
	return akc
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (akc *AccountKeyCreate) SetNillableCreatedAt(t *time.Time) *AccountKeyCreate {
	if t != nil {
		akc.SetCreatedAt(*t)
	}
	return akc
}

// SetUpdatedAt sets the updated_at field.
func (akc *AccountKeyCreate) SetUpdatedAt(t time.Time) *AccountKeyCreate {
	akc.updated_at = &t
	return akc
}

// SetNillableUpdatedAt sets the updated_at field if the given value is not nil.
func (akc *AccountKeyCreate) SetNillableUpdatedAt(t *time.Time) *AccountKeyCreate {
	if t != nil {
		akc.SetUpdatedAt(*t)
	}
	return akc
}

// SetAccountID sets the account_id field.
func (akc *AccountKeyCreate) SetAccountID(i int16) *AccountKeyCreate {
	akc.account_id = &i
	return akc
}

// SetKey sets the key field.
func (akc *AccountKeyCreate) SetKey(s string) *AccountKeyCreate {
	akc.key = &s
	return akc
}

// SetOptions sets the options field.
func (akc *AccountKeyCreate) SetOptions(s string) *AccountKeyCreate {
	akc.options = &s
	return akc
}

// SetNillableOptions sets the options field if the given value is not nil.
func (akc *AccountKeyCreate) SetNillableOptions(s *string) *AccountKeyCreate {
	if s != nil {
		akc.SetOptions(*s)
	}
	return akc
}

// SetMetaData sets the meta_data field.
func (akc *AccountKeyCreate) SetMetaData(s string) *AccountKeyCreate {
	akc.meta_data = &s
	return akc
}

// SetUseraccountID sets the useraccount edge to UserAccount by id.
func (akc *AccountKeyCreate) SetUseraccountID(id int) *AccountKeyCreate {
	if akc.useraccount == nil {
		akc.useraccount = make(map[int]struct{})
	}
	akc.useraccount[id] = struct{}{}
	return akc
}

// SetNillableUseraccountID sets the useraccount edge to UserAccount by id if the given value is not nil.
func (akc *AccountKeyCreate) SetNillableUseraccountID(id *int) *AccountKeyCreate {
	if id != nil {
		akc = akc.SetUseraccountID(*id)
	}
	return akc
}

// SetUseraccount sets the useraccount edge to UserAccount.
func (akc *AccountKeyCreate) SetUseraccount(u *UserAccount) *AccountKeyCreate {
	return akc.SetUseraccountID(u.ID)
}

// Save creates the AccountKey in the database.
func (akc *AccountKeyCreate) Save(ctx context.Context) (*AccountKey, error) {
	if akc.account_id == nil {
		return nil, errors.New("ent: missing required field \"account_id\"")
	}
	if akc.key == nil {
		return nil, errors.New("ent: missing required field \"key\"")
	}
	if err := accountkey.KeyValidator(*akc.key); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"key\": %v", err)
	}
	if akc.meta_data == nil {
		return nil, errors.New("ent: missing required field \"meta_data\"")
	}
	if len(akc.useraccount) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"useraccount\"")
	}
	return akc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (akc *AccountKeyCreate) SaveX(ctx context.Context) *AccountKey {
	v, err := akc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (akc *AccountKeyCreate) sqlSave(ctx context.Context) (*AccountKey, error) {
	var (
		ak    = &AccountKey{config: akc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: accountkey.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: accountkey.FieldID,
			},
		}
	)
	if value := akc.created_at; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: accountkey.FieldCreatedAt,
		})
		ak.CreatedAt = value
	}
	if value := akc.updated_at; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: accountkey.FieldUpdatedAt,
		})
		ak.UpdatedAt = value
	}
	if value := akc.account_id; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt16,
			Value:  *value,
			Column: accountkey.FieldAccountID,
		})
		ak.AccountID = *value
	}
	if value := akc.key; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: accountkey.FieldKey,
		})
		ak.Key = *value
	}
	if value := akc.options; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: accountkey.FieldOptions,
		})
		ak.Options = value
	}
	if value := akc.meta_data; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: accountkey.FieldMetaData,
		})
		ak.MetaData = *value
	}
	if nodes := akc.useraccount; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   accountkey.UseraccountTable,
			Columns: []string{accountkey.UseraccountColumn},
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
	if err := sqlgraph.CreateNode(ctx, akc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	ak.ID = int(id)
	return ak, nil
}
