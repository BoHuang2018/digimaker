package fieldtypes

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/digimakergo/digimaker/core/definition"
	"github.com/digimakergo/digimaker/core/fieldtype"
)

//Map defines a key - value map
type Map map[string]interface{}

//List defines an array Map
type MapList []Map

func (a Map) Value() (driver.Value, error) {
	value, err := json.Marshal(a)
	return value, err
}

func (a *Map) Scan(value interface{}) error {
	obj := Map{}
	if value != nil {
		err := json.Unmarshal(value.([]byte), &obj)
		if err != nil {
			return err
		}
		*a = obj
	} else {
		*a = nil
	}
	return nil
}

//MapHandler
type MapHandler struct {
	definition.FieldDef
}

func (handler MapHandler) LoadInput(input interface{}, mode string) (interface{}, error) {
	if _, ok := input.(Map); ok {
		return input, nil
	}
	data := fmt.Sprint(input)
	m := Map{}
	err := json.Unmarshal([]byte(data), &m)
	return m, err
}

func (handler MapHandler) DBField() string {
	return "JSON"
}

//MapListHandler
type MapListHandler struct {
	definition.FieldDef
}

func (handler MapListHandler) LoadInput(input interface{}, mode string) (interface{}, error) {
	if _, ok := input.(MapList); ok {
		return input, nil
	}
	data := fmt.Sprint(input)
	m := Map{}
	err := json.Unmarshal([]byte(data), &m)
	return m, err
}

func (handler MapListHandler) DBField() string {
	return "JSON"
}

//JSON Handler
type JSONHandler struct {
	definition.FieldDef
}

func (handler JSONHandler) LoadInput(input interface{}, mode string) (interface{}, error) {
	if input == nil {
		return "", nil
	}

	data := fmt.Sprint(input)
	isValid := json.Valid([]byte(data))
	if !isValid {
		return "", fieldtype.NewValidationError("Not a valid json")
	}

	return data, nil
}

func (handler JSONHandler) DBField() string {
	return "JSON"
}

func init() {
	fieldtype.Register(
		fieldtype.Definition{Name: "map",
			DataType: "fieldtypes.Map",
			Package:  "github.com/digimakergo/digimaker/core/fieldtype/fieldtypes",
			NewHandler: func(def definition.FieldDef) fieldtype.Handler {
				return MapHandler{FieldDef: def}
			}})
	fieldtype.Register(fieldtype.Definition{Name: "maplist",
		DataType: "fieldtypes.MapList",
		Package:  "github.com/digimakergo/digimaker/core/fieldtype/fieldtypes",
		NewHandler: func(def definition.FieldDef) fieldtype.Handler {
			return MapListHandler{FieldDef: def}
		}})
	fieldtype.Register(fieldtype.Definition{Name: "json",
		DataType: "string",
		NewHandler: func(def definition.FieldDef) fieldtype.Handler {
			return JSONHandler{FieldDef: def}
		}})
}
