// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/entc/integration/ent/fieldtype"
	"github.com/facebookincubator/ent/schema/field"
)

// FieldTypeCreate is the builder for creating a FieldType entity.
type FieldTypeCreate struct {
	config
	int                     *int
	int8                    *int8
	int16                   *int16
	int32                   *int32
	int64                   *int64
	optional_int            *int
	optional_int8           *int8
	optional_int16          *int16
	optional_int32          *int32
	optional_int64          *int64
	nillable_int            *int
	nillable_int8           *int8
	nillable_int16          *int16
	nillable_int32          *int32
	nillable_int64          *int64
	validate_optional_int32 *int32
	optional_uint           *uint
	optional_uint8          *uint8
	optional_uint16         *uint16
	optional_uint32         *uint32
	optional_uint64         *uint64
	state                   *fieldtype.State
	optional_float          *float64
	optional_float32        *float32
}

// SetInt sets the int field.
func (ftc *FieldTypeCreate) SetInt(i int) *FieldTypeCreate {
	ftc.int = &i
	return ftc
}

// SetInt8 sets the int8 field.
func (ftc *FieldTypeCreate) SetInt8(i int8) *FieldTypeCreate {
	ftc.int8 = &i
	return ftc
}

// SetInt16 sets the int16 field.
func (ftc *FieldTypeCreate) SetInt16(i int16) *FieldTypeCreate {
	ftc.int16 = &i
	return ftc
}

// SetInt32 sets the int32 field.
func (ftc *FieldTypeCreate) SetInt32(i int32) *FieldTypeCreate {
	ftc.int32 = &i
	return ftc
}

// SetInt64 sets the int64 field.
func (ftc *FieldTypeCreate) SetInt64(i int64) *FieldTypeCreate {
	ftc.int64 = &i
	return ftc
}

// SetOptionalInt sets the optional_int field.
func (ftc *FieldTypeCreate) SetOptionalInt(i int) *FieldTypeCreate {
	ftc.optional_int = &i
	return ftc
}

// SetNillableOptionalInt sets the optional_int field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalInt(i *int) *FieldTypeCreate {
	if i != nil {
		ftc.SetOptionalInt(*i)
	}
	return ftc
}

// SetOptionalInt8 sets the optional_int8 field.
func (ftc *FieldTypeCreate) SetOptionalInt8(i int8) *FieldTypeCreate {
	ftc.optional_int8 = &i
	return ftc
}

// SetNillableOptionalInt8 sets the optional_int8 field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalInt8(i *int8) *FieldTypeCreate {
	if i != nil {
		ftc.SetOptionalInt8(*i)
	}
	return ftc
}

// SetOptionalInt16 sets the optional_int16 field.
func (ftc *FieldTypeCreate) SetOptionalInt16(i int16) *FieldTypeCreate {
	ftc.optional_int16 = &i
	return ftc
}

// SetNillableOptionalInt16 sets the optional_int16 field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalInt16(i *int16) *FieldTypeCreate {
	if i != nil {
		ftc.SetOptionalInt16(*i)
	}
	return ftc
}

// SetOptionalInt32 sets the optional_int32 field.
func (ftc *FieldTypeCreate) SetOptionalInt32(i int32) *FieldTypeCreate {
	ftc.optional_int32 = &i
	return ftc
}

// SetNillableOptionalInt32 sets the optional_int32 field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalInt32(i *int32) *FieldTypeCreate {
	if i != nil {
		ftc.SetOptionalInt32(*i)
	}
	return ftc
}

// SetOptionalInt64 sets the optional_int64 field.
func (ftc *FieldTypeCreate) SetOptionalInt64(i int64) *FieldTypeCreate {
	ftc.optional_int64 = &i
	return ftc
}

// SetNillableOptionalInt64 sets the optional_int64 field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalInt64(i *int64) *FieldTypeCreate {
	if i != nil {
		ftc.SetOptionalInt64(*i)
	}
	return ftc
}

// SetNillableInt sets the nillable_int field.
func (ftc *FieldTypeCreate) SetNillableInt(i int) *FieldTypeCreate {
	ftc.nillable_int = &i
	return ftc
}

// SetNillableNillableInt sets the nillable_int field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableNillableInt(i *int) *FieldTypeCreate {
	if i != nil {
		ftc.SetNillableInt(*i)
	}
	return ftc
}

// SetNillableInt8 sets the nillable_int8 field.
func (ftc *FieldTypeCreate) SetNillableInt8(i int8) *FieldTypeCreate {
	ftc.nillable_int8 = &i
	return ftc
}

// SetNillableNillableInt8 sets the nillable_int8 field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableNillableInt8(i *int8) *FieldTypeCreate {
	if i != nil {
		ftc.SetNillableInt8(*i)
	}
	return ftc
}

// SetNillableInt16 sets the nillable_int16 field.
func (ftc *FieldTypeCreate) SetNillableInt16(i int16) *FieldTypeCreate {
	ftc.nillable_int16 = &i
	return ftc
}

// SetNillableNillableInt16 sets the nillable_int16 field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableNillableInt16(i *int16) *FieldTypeCreate {
	if i != nil {
		ftc.SetNillableInt16(*i)
	}
	return ftc
}

// SetNillableInt32 sets the nillable_int32 field.
func (ftc *FieldTypeCreate) SetNillableInt32(i int32) *FieldTypeCreate {
	ftc.nillable_int32 = &i
	return ftc
}

// SetNillableNillableInt32 sets the nillable_int32 field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableNillableInt32(i *int32) *FieldTypeCreate {
	if i != nil {
		ftc.SetNillableInt32(*i)
	}
	return ftc
}

// SetNillableInt64 sets the nillable_int64 field.
func (ftc *FieldTypeCreate) SetNillableInt64(i int64) *FieldTypeCreate {
	ftc.nillable_int64 = &i
	return ftc
}

// SetNillableNillableInt64 sets the nillable_int64 field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableNillableInt64(i *int64) *FieldTypeCreate {
	if i != nil {
		ftc.SetNillableInt64(*i)
	}
	return ftc
}

// SetValidateOptionalInt32 sets the validate_optional_int32 field.
func (ftc *FieldTypeCreate) SetValidateOptionalInt32(i int32) *FieldTypeCreate {
	ftc.validate_optional_int32 = &i
	return ftc
}

// SetNillableValidateOptionalInt32 sets the validate_optional_int32 field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableValidateOptionalInt32(i *int32) *FieldTypeCreate {
	if i != nil {
		ftc.SetValidateOptionalInt32(*i)
	}
	return ftc
}

// SetOptionalUint sets the optional_uint field.
func (ftc *FieldTypeCreate) SetOptionalUint(u uint) *FieldTypeCreate {
	ftc.optional_uint = &u
	return ftc
}

// SetNillableOptionalUint sets the optional_uint field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalUint(u *uint) *FieldTypeCreate {
	if u != nil {
		ftc.SetOptionalUint(*u)
	}
	return ftc
}

// SetOptionalUint8 sets the optional_uint8 field.
func (ftc *FieldTypeCreate) SetOptionalUint8(u uint8) *FieldTypeCreate {
	ftc.optional_uint8 = &u
	return ftc
}

// SetNillableOptionalUint8 sets the optional_uint8 field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalUint8(u *uint8) *FieldTypeCreate {
	if u != nil {
		ftc.SetOptionalUint8(*u)
	}
	return ftc
}

// SetOptionalUint16 sets the optional_uint16 field.
func (ftc *FieldTypeCreate) SetOptionalUint16(u uint16) *FieldTypeCreate {
	ftc.optional_uint16 = &u
	return ftc
}

// SetNillableOptionalUint16 sets the optional_uint16 field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalUint16(u *uint16) *FieldTypeCreate {
	if u != nil {
		ftc.SetOptionalUint16(*u)
	}
	return ftc
}

// SetOptionalUint32 sets the optional_uint32 field.
func (ftc *FieldTypeCreate) SetOptionalUint32(u uint32) *FieldTypeCreate {
	ftc.optional_uint32 = &u
	return ftc
}

// SetNillableOptionalUint32 sets the optional_uint32 field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalUint32(u *uint32) *FieldTypeCreate {
	if u != nil {
		ftc.SetOptionalUint32(*u)
	}
	return ftc
}

// SetOptionalUint64 sets the optional_uint64 field.
func (ftc *FieldTypeCreate) SetOptionalUint64(u uint64) *FieldTypeCreate {
	ftc.optional_uint64 = &u
	return ftc
}

// SetNillableOptionalUint64 sets the optional_uint64 field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalUint64(u *uint64) *FieldTypeCreate {
	if u != nil {
		ftc.SetOptionalUint64(*u)
	}
	return ftc
}

// SetState sets the state field.
func (ftc *FieldTypeCreate) SetState(f fieldtype.State) *FieldTypeCreate {
	ftc.state = &f
	return ftc
}

// SetNillableState sets the state field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableState(f *fieldtype.State) *FieldTypeCreate {
	if f != nil {
		ftc.SetState(*f)
	}
	return ftc
}

// SetOptionalFloat sets the optional_float field.
func (ftc *FieldTypeCreate) SetOptionalFloat(f float64) *FieldTypeCreate {
	ftc.optional_float = &f
	return ftc
}

// SetNillableOptionalFloat sets the optional_float field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalFloat(f *float64) *FieldTypeCreate {
	if f != nil {
		ftc.SetOptionalFloat(*f)
	}
	return ftc
}

// SetOptionalFloat32 sets the optional_float32 field.
func (ftc *FieldTypeCreate) SetOptionalFloat32(f float32) *FieldTypeCreate {
	ftc.optional_float32 = &f
	return ftc
}

// SetNillableOptionalFloat32 sets the optional_float32 field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalFloat32(f *float32) *FieldTypeCreate {
	if f != nil {
		ftc.SetOptionalFloat32(*f)
	}
	return ftc
}

// Save creates the FieldType in the database.
func (ftc *FieldTypeCreate) Save(ctx context.Context) (*FieldType, error) {
	if ftc.int == nil {
		return nil, errors.New("ent: missing required field \"int\"")
	}
	if ftc.int8 == nil {
		return nil, errors.New("ent: missing required field \"int8\"")
	}
	if ftc.int16 == nil {
		return nil, errors.New("ent: missing required field \"int16\"")
	}
	if ftc.int32 == nil {
		return nil, errors.New("ent: missing required field \"int32\"")
	}
	if ftc.int64 == nil {
		return nil, errors.New("ent: missing required field \"int64\"")
	}
	if ftc.validate_optional_int32 != nil {
		if err := fieldtype.ValidateOptionalInt32Validator(*ftc.validate_optional_int32); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"validate_optional_int32\": %v", err)
		}
	}
	if ftc.state != nil {
		if err := fieldtype.StateValidator(*ftc.state); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"state\": %v", err)
		}
	}
	return ftc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (ftc *FieldTypeCreate) SaveX(ctx context.Context) *FieldType {
	v, err := ftc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ftc *FieldTypeCreate) sqlSave(ctx context.Context) (*FieldType, error) {
	var (
		ft    = &FieldType{config: ftc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: fieldtype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: fieldtype.FieldID,
			},
		}
	)
	if value := ftc.int; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: fieldtype.FieldInt,
		})
		ft.Int = *value
	}
	if value := ftc.int8; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  *value,
			Column: fieldtype.FieldInt8,
		})
		ft.Int8 = *value
	}
	if value := ftc.int16; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt16,
			Value:  *value,
			Column: fieldtype.FieldInt16,
		})
		ft.Int16 = *value
	}
	if value := ftc.int32; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  *value,
			Column: fieldtype.FieldInt32,
		})
		ft.Int32 = *value
	}
	if value := ftc.int64; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  *value,
			Column: fieldtype.FieldInt64,
		})
		ft.Int64 = *value
	}
	if value := ftc.optional_int; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: fieldtype.FieldOptionalInt,
		})
		ft.OptionalInt = *value
	}
	if value := ftc.optional_int8; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  *value,
			Column: fieldtype.FieldOptionalInt8,
		})
		ft.OptionalInt8 = *value
	}
	if value := ftc.optional_int16; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt16,
			Value:  *value,
			Column: fieldtype.FieldOptionalInt16,
		})
		ft.OptionalInt16 = *value
	}
	if value := ftc.optional_int32; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  *value,
			Column: fieldtype.FieldOptionalInt32,
		})
		ft.OptionalInt32 = *value
	}
	if value := ftc.optional_int64; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  *value,
			Column: fieldtype.FieldOptionalInt64,
		})
		ft.OptionalInt64 = *value
	}
	if value := ftc.nillable_int; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: fieldtype.FieldNillableInt,
		})
		ft.NillableInt = value
	}
	if value := ftc.nillable_int8; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  *value,
			Column: fieldtype.FieldNillableInt8,
		})
		ft.NillableInt8 = value
	}
	if value := ftc.nillable_int16; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt16,
			Value:  *value,
			Column: fieldtype.FieldNillableInt16,
		})
		ft.NillableInt16 = value
	}
	if value := ftc.nillable_int32; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  *value,
			Column: fieldtype.FieldNillableInt32,
		})
		ft.NillableInt32 = value
	}
	if value := ftc.nillable_int64; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  *value,
			Column: fieldtype.FieldNillableInt64,
		})
		ft.NillableInt64 = value
	}
	if value := ftc.validate_optional_int32; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  *value,
			Column: fieldtype.FieldValidateOptionalInt32,
		})
		ft.ValidateOptionalInt32 = *value
	}
	if value := ftc.optional_uint; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  *value,
			Column: fieldtype.FieldOptionalUint,
		})
		ft.OptionalUint = *value
	}
	if value := ftc.optional_uint8; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  *value,
			Column: fieldtype.FieldOptionalUint8,
		})
		ft.OptionalUint8 = *value
	}
	if value := ftc.optional_uint16; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint16,
			Value:  *value,
			Column: fieldtype.FieldOptionalUint16,
		})
		ft.OptionalUint16 = *value
	}
	if value := ftc.optional_uint32; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  *value,
			Column: fieldtype.FieldOptionalUint32,
		})
		ft.OptionalUint32 = *value
	}
	if value := ftc.optional_uint64; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  *value,
			Column: fieldtype.FieldOptionalUint64,
		})
		ft.OptionalUint64 = *value
	}
	if value := ftc.state; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  *value,
			Column: fieldtype.FieldState,
		})
		ft.State = *value
	}
	if value := ftc.optional_float; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  *value,
			Column: fieldtype.FieldOptionalFloat,
		})
		ft.OptionalFloat = *value
	}
	if value := ftc.optional_float32; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  *value,
			Column: fieldtype.FieldOptionalFloat32,
		})
		ft.OptionalFloat32 = *value
	}
	if err := sqlgraph.CreateNode(ctx, ftc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	ft.ID = strconv.FormatInt(id, 10)
	return ft, nil
}
