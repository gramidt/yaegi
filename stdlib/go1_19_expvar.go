// Code generated by 'yaegi extract expvar'. DO NOT EDIT.

//go:build go1.19 && !go1.20
// +build go1.19,!go1.20

package stdlib

import (
	"expvar"
	"reflect"
)

func init() {
	Symbols["expvar/expvar"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Do":        reflect.ValueOf(expvar.Do),
		"Get":       reflect.ValueOf(expvar.Get),
		"Handler":   reflect.ValueOf(expvar.Handler),
		"NewFloat":  reflect.ValueOf(expvar.NewFloat),
		"NewInt":    reflect.ValueOf(expvar.NewInt),
		"NewMap":    reflect.ValueOf(expvar.NewMap),
		"NewString": reflect.ValueOf(expvar.NewString),
		"Publish":   reflect.ValueOf(expvar.Publish),

		// type definitions
		"Float":    reflect.ValueOf((*expvar.Float)(nil)),
		"Func":     reflect.ValueOf((*expvar.Func)(nil)),
		"Int":      reflect.ValueOf((*expvar.Int)(nil)),
		"KeyValue": reflect.ValueOf((*expvar.KeyValue)(nil)),
		"Map":      reflect.ValueOf((*expvar.Map)(nil)),
		"String":   reflect.ValueOf((*expvar.String)(nil)),
		"Var":      reflect.ValueOf((*expvar.Var)(nil)),

		// interface wrapper definitions
		"_Var": reflect.ValueOf((*_expvar_Var)(nil)),
	}
}

// _expvar_Var is an interface wrapper for Var type
type _expvar_Var struct {
	IValue  interface{}
	WString func() string
}

func (W _expvar_Var) String() string {
	if W.WString == nil {
		return ""
	}
	return W.WString()
}
