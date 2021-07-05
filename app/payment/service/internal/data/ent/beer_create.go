// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/go-kratos/beer-shop/app/payment/service/internal/data/ent/beer"
)

// BeerCreate is the builder for creating a Beer entity.
type BeerCreate struct {
	config
	mutation *BeerMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (bc *BeerCreate) SetName(s string) *BeerCreate {
	bc.mutation.SetName(s)
	return bc
}

// SetDescription sets the "description" field.
func (bc *BeerCreate) SetDescription(s string) *BeerCreate {
	bc.mutation.SetDescription(s)
	return bc
}

// SetCount sets the "count" field.
func (bc *BeerCreate) SetCount(i int64) *BeerCreate {
	bc.mutation.SetCount(i)
	return bc
}

// SetPrice sets the "price" field.
func (bc *BeerCreate) SetPrice(i int64) *BeerCreate {
	bc.mutation.SetPrice(i)
	return bc
}

// SetCreatedAt sets the "created_at" field.
func (bc *BeerCreate) SetCreatedAt(t time.Time) *BeerCreate {
	bc.mutation.SetCreatedAt(t)
	return bc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (bc *BeerCreate) SetNillableCreatedAt(t *time.Time) *BeerCreate {
	if t != nil {
		bc.SetCreatedAt(*t)
	}
	return bc
}

// SetUpdatedAt sets the "updated_at" field.
func (bc *BeerCreate) SetUpdatedAt(t time.Time) *BeerCreate {
	bc.mutation.SetUpdatedAt(t)
	return bc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (bc *BeerCreate) SetNillableUpdatedAt(t *time.Time) *BeerCreate {
	if t != nil {
		bc.SetUpdatedAt(*t)
	}
	return bc
}

// SetID sets the "id" field.
func (bc *BeerCreate) SetID(i int64) *BeerCreate {
	bc.mutation.SetID(i)
	return bc
}

// Mutation returns the BeerMutation object of the builder.
func (bc *BeerCreate) Mutation() *BeerMutation {
	return bc.mutation
}

// Save creates the Beer in the database.
func (bc *BeerCreate) Save(ctx context.Context) (*Beer, error) {
	var (
		err  error
		node *Beer
	)
	bc.defaults()
	if len(bc.hooks) == 0 {
		if err = bc.check(); err != nil {
			return nil, err
		}
		node, err = bc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BeerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bc.check(); err != nil {
				return nil, err
			}
			bc.mutation = mutation
			node, err = bc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(bc.hooks) - 1; i >= 0; i-- {
			mut = bc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BeerCreate) SaveX(ctx context.Context) *Beer {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (bc *BeerCreate) defaults() {
	if _, ok := bc.mutation.CreatedAt(); !ok {
		v := beer.DefaultCreatedAt()
		bc.mutation.SetCreatedAt(v)
	}
	if _, ok := bc.mutation.UpdatedAt(); !ok {
		v := beer.DefaultUpdatedAt()
		bc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bc *BeerCreate) check() error {
	if _, ok := bc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if _, ok := bc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New("ent: missing required field \"description\"")}
	}
	if _, ok := bc.mutation.Count(); !ok {
		return &ValidationError{Name: "count", err: errors.New("ent: missing required field \"count\"")}
	}
	if _, ok := bc.mutation.Price(); !ok {
		return &ValidationError{Name: "price", err: errors.New("ent: missing required field \"price\"")}
	}
	if _, ok := bc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("ent: missing required field \"created_at\"")}
	}
	if _, ok := bc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New("ent: missing required field \"updated_at\"")}
	}
	return nil
}

func (bc *BeerCreate) sqlSave(ctx context.Context) (*Beer, error) {
	_node, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	if _node.ID == 0 {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	return _node, nil
}

func (bc *BeerCreate) createSpec() (*Beer, *sqlgraph.CreateSpec) {
	var (
		_node = &Beer{config: bc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: beer.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: beer.FieldID,
			},
		}
	)
	if id, ok := bc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := bc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: beer.FieldName,
		})
		_node.Name = value
	}
	if value, ok := bc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: beer.FieldDescription,
		})
		_node.Description = value
	}
	if value, ok := bc.mutation.Count(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: beer.FieldCount,
		})
		_node.Count = value
	}
	if value, ok := bc.mutation.Price(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: beer.FieldPrice,
		})
		_node.Price = value
	}
	if value, ok := bc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: beer.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := bc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: beer.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// BeerCreateBulk is the builder for creating many Beer entities in bulk.
type BeerCreateBulk struct {
	config
	builders []*BeerCreate
}

// Save creates the Beer entities in the database.
func (bcb *BeerCreateBulk) Save(ctx context.Context) ([]*Beer, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bcb.builders))
	nodes := make([]*Beer, len(bcb.builders))
	mutators := make([]Mutator, len(bcb.builders))
	for i := range bcb.builders {
		func(i int, root context.Context) {
			builder := bcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BeerMutation)
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
					_, err = mutators[i+1].Mutate(root, bcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				if nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, bcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bcb *BeerCreateBulk) SaveX(ctx context.Context) []*Beer {
	v, err := bcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
