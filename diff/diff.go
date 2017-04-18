package diff

import (
	"reflect"
)

type Type int

const (
	TypesDiffer Type = iota
	ContentDiffer
	Identical
)

type Differ interface {
	Diff() Type
	Strings() []string
	StringIndent(key, prefix string, conf Output) string
}

type Comparaison struct {
	Type
	LHS Differ
	RHS Differ
}

func Diff(lhs, rhs interface{}) (Differ, error) {
	lhsVal := reflect.ValueOf(lhs)
	rhsVal := reflect.ValueOf(rhs)

	if lhs == nil && rhs == nil {
		return &Scalar{lhs, rhs}, nil
	}
	if lhs == nil || rhs == nil {
		return &Types{lhs, rhs}, nil
	}
	if lhsVal.Type().Comparable() && rhsVal.Type().Comparable() {
		return &Scalar{lhs, rhs}, nil
	}
	if lhsVal.Kind() != rhsVal.Kind() {
		return &Types{lhs, rhs}, nil
	}
	if lhsVal.Kind() == reflect.Slice {
		return NewSlice(lhs, rhs)
	}
	if lhsVal.Kind() == reflect.Map {
		return NewMap(lhs, rhs)
	}

	return &Types{lhs, rhs}, nil
}
