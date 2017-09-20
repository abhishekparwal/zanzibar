// Code generated by thriftrw v1.6.0. DO NOT EDIT.
// @generated

package bar

import (
	"errors"
	"fmt"
	"go.uber.org/thriftrw/wire"
	"strings"
)

type Bar_NormalRecur_Args struct {
	Request *BarRequestRecur `json:"request,required"`
}

func (v *Bar_NormalRecur_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Request == nil {
		return w, errors.New("field Request of Bar_NormalRecur_Args is required")
	}
	w, err = v.Request.ToWire()
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *Bar_NormalRecur_Args) FromWire(w wire.Value) error {
	var err error
	requestIsSet := false
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.Request, err = _BarRequestRecur_Read(field.Value)
				if err != nil {
					return err
				}
				requestIsSet = true
			}
		}
	}
	if !requestIsSet {
		return errors.New("field Request of Bar_NormalRecur_Args is required")
	}
	return nil
}

func (v *Bar_NormalRecur_Args) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [1]string
	i := 0
	fields[i] = fmt.Sprintf("Request: %v", v.Request)
	i++
	return fmt.Sprintf("Bar_NormalRecur_Args{%v}", strings.Join(fields[:i], ", "))
}

func (v *Bar_NormalRecur_Args) Equals(rhs *Bar_NormalRecur_Args) bool {
	if !v.Request.Equals(rhs.Request) {
		return false
	}
	return true
}

func (v *Bar_NormalRecur_Args) MethodName() string {
	return "normalRecur"
}

func (v *Bar_NormalRecur_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

var Bar_NormalRecur_Helper = struct {
	Args           func(request *BarRequestRecur) *Bar_NormalRecur_Args
	IsException    func(error) bool
	WrapResponse   func(*BarResponseRecur, error) (*Bar_NormalRecur_Result, error)
	UnwrapResponse func(*Bar_NormalRecur_Result) (*BarResponseRecur, error)
}{}

func init() {
	Bar_NormalRecur_Helper.Args = func(request *BarRequestRecur) *Bar_NormalRecur_Args {
		return &Bar_NormalRecur_Args{Request: request}
	}
	Bar_NormalRecur_Helper.IsException = func(err error) bool {
		switch err.(type) {
		case *BarException:
			return true
		default:
			return false
		}
	}
	Bar_NormalRecur_Helper.WrapResponse = func(success *BarResponseRecur, err error) (*Bar_NormalRecur_Result, error) {
		if err == nil {
			return &Bar_NormalRecur_Result{Success: success}, nil
		}
		switch e := err.(type) {
		case *BarException:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for Bar_NormalRecur_Result.BarException")
			}
			return &Bar_NormalRecur_Result{BarException: e}, nil
		}
		return nil, err
	}
	Bar_NormalRecur_Helper.UnwrapResponse = func(result *Bar_NormalRecur_Result) (success *BarResponseRecur, err error) {
		if result.BarException != nil {
			err = result.BarException
			return
		}
		if result.Success != nil {
			success = result.Success
			return
		}
		err = errors.New("expected a non-void result")
		return
	}
}

type Bar_NormalRecur_Result struct {
	Success      *BarResponseRecur `json:"success,omitempty"`
	BarException *BarException     `json:"barException,omitempty"`
}

func (v *Bar_NormalRecur_Result) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Success != nil {
		w, err = v.Success.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}
	if v.BarException != nil {
		w, err = v.BarException.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	if i != 1 {
		return wire.Value{}, fmt.Errorf("Bar_NormalRecur_Result should have exactly one field: got %v fields", i)
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _BarResponseRecur_Read(w wire.Value) (*BarResponseRecur, error) {
	var v BarResponseRecur
	err := v.FromWire(w)
	return &v, err
}

func (v *Bar_NormalRecur_Result) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TStruct {
				v.Success, err = _BarResponseRecur_Read(field.Value)
				if err != nil {
					return err
				}
			}
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.BarException, err = _BarException_Read(field.Value)
				if err != nil {
					return err
				}
			}
		}
	}
	count := 0
	if v.Success != nil {
		count++
	}
	if v.BarException != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("Bar_NormalRecur_Result should have exactly one field: got %v fields", count)
	}
	return nil
}

func (v *Bar_NormalRecur_Result) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [2]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", v.Success)
		i++
	}
	if v.BarException != nil {
		fields[i] = fmt.Sprintf("BarException: %v", v.BarException)
		i++
	}
	return fmt.Sprintf("Bar_NormalRecur_Result{%v}", strings.Join(fields[:i], ", "))
}

func (v *Bar_NormalRecur_Result) Equals(rhs *Bar_NormalRecur_Result) bool {
	if !((v.Success == nil && rhs.Success == nil) || (v.Success != nil && rhs.Success != nil && v.Success.Equals(rhs.Success))) {
		return false
	}
	if !((v.BarException == nil && rhs.BarException == nil) || (v.BarException != nil && rhs.BarException != nil && v.BarException.Equals(rhs.BarException))) {
		return false
	}
	return true
}

func (v *Bar_NormalRecur_Result) MethodName() string {
	return "normalRecur"
}

func (v *Bar_NormalRecur_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}