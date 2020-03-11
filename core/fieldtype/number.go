//Author xc, Created on 2019-03-25 20:00
//{COPYRIGHTS}

//Package fieldtype implements build-in field types(value and fieldtype handler).
package fieldtype

import (
        "fmt"
	"strconv"
	"strings"
)

//TextField is a field for normal text line. It implements Datatyper
type NumberField struct {
	FieldValue
}

func (t *NumberField) Scan(src interface{}) error {
	err := t.SetData(src, "number")
	return err
}

//convert data to view data.
func (t NumberField) ViewValue() string {
	return t.Raw
}

//implement FieldtypeHandler
type NumberHandler struct{}

func (t NumberHandler) Validate(input interface{}) (bool, string) {
	//todo: support int
	s := fmt.Sprint(input)
	if s != "" {
		_, err := strconv.Atoi(s)
		if err != nil {
			return false, s + " is not a number."
		}
	}
	return true, ""
}

func (t NumberHandler) NewValueFromInput(input interface{}) interface{} {
	r := NumberField{}
	r.Scan(fmt.Sprint(input))
	return r
}

func (t NumberHandler) IsEmpty(input interface{}) bool {
	if strings.TrimSpace(fmt.Sprint(input)) == "" {
		return true
	}
	return false
}

func init() {
	RegisterHandler("number", NumberHandler{})
}
