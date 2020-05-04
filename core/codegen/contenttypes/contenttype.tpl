//This file is generated automatically, DO NOT EDIT IT.
//Generated time:

package entity

import (
    "database/sql"
    "github.com/xc/digimaker/core/db"
    "github.com/xc/digimaker/core/contenttype"
	"github.com/xc/digimaker/core/fieldtype"
    {{if .settings.HasLocation}}
    "github.com/xc/digimaker/core/util"
    {{end}}
	. "github.com/xc/digimaker/core/db"
)

{{$struct_name :=.name|UpperName}}

type {{$struct_name}} struct{
     contenttype.ContentCommon `boil:",bind"`

     {{range $identifier, $fieldtype := .data_fields}}
          {{$identifier|UpperName}}  {{$fieldtype}} `boil:"{{$identifier}}" json:"{{$identifier}}" toml:"{{$identifier}}" yaml:"{{$identifier}}"`
     {{end}}
    {{range $identifier, $fieldtype := .fields}}
         {{$type_settings := index $.def_fieldtype $fieldtype.FieldType}}
         {{if not $type_settings.IsRelation }}
         {{if not $fieldtype.IsOutput}}
            {{$identifier|UpperName}}  fieldtype.{{$type_settings.Value}} `boil:"{{$identifier}}" json:"{{if eq $fieldtype.FieldType "password"}}-{{else}}{{$identifier}}{{end}}" toml:"{{$identifier}}" yaml:"{{$identifier}}"`
         {{end}}
        {{end}}
    {{end}}
    {{if .settings.HasLocation}}
     contenttype.Location `boil:"location,bind"`
    {{end}}
}

func ( *{{$struct_name}} ) TableName() string{
	 return "{{.settings.TableName}}"
}

func ( *{{$struct_name}} ) ContentType() string{
	 return "{{.name}}"
}

func (c *{{$struct_name}} ) GetName() string{
	 location := c.GetLocation()
     if location != nil{
         return location.Name
     }else{
         return ""
     }
}

func (c *{{$struct_name}}) GetLocation() *contenttype.Location{
    {{if .settings.HasLocation}}
    return &c.Location
    {{else}}
    return nil
    {{end}}
}

//Get map of the all fields(including data_fields)
//todo: cache this? (then you need a reload?)
func (c *{{$struct_name}}) ToDBValues() map[string]interface{} {
	result := make(map[string]interface{})
    {{range $identifier, $fieldtype := .data_fields}}
         result["{{$identifier}}"]=c.{{$identifier|UpperName}}
    {{end}}

    {{range $identifier, $fieldtype := .fields}}
        {{if not (index $.def_fieldtype $fieldtype.FieldType).IsRelation}}
        {{if not $fieldtype.IsOutput}}
            result["{{$identifier}}"]=c.{{$identifier|UpperName}}
        {{end}}
        {{end}}
    {{end}}
	for key, value := range c.ContentCommon.ToDBValues() {
		result[key] = value
	}
	return result
}

//Get identifier list of fields(NOT including data_fields )
func (c *{{$struct_name}}) IdentifierList() []string {
	return append(c.ContentCommon.IdentifierList(),[]string{ {{range $identifier, $fieldtype := .fields}}{{if not $fieldtype.IsOutput}}"{{$identifier}}",{{end}}{{end}}}...)
}

func (c *{{$struct_name}}) Definition(language ...string) contenttype.ContentType {
	def, _ := contenttype.GetDefinition( c.ContentType(), language... )
    return def
}

//Get field value
func (c *{{$struct_name}}) Value(identifier string) interface{} {
    {{if .settings.HasLocation}}
    if util.Contains( c.Location.IdentifierList(), identifier ) {
        return c.Location.Field( identifier )
    }
    {{end}}
    var result interface{}
	switch identifier {
    {{range $identifier, $fieldtype := .data_fields}}
      case "{{$identifier}}":
         result = c.{{$identifier|UpperName}}
    {{end}}
    {{range $identifier, $fieldtype := .fields}}
    {{if not $fieldtype.IsOutput}}
    case "{{$identifier}}":
        {{if not (index $.def_fieldtype $fieldtype.FieldType).IsRelation}}
            result = c.{{$identifier|UpperName}}
        {{else}}
            result = c.Relations.Map["{{$identifier}}"]
        {{end}}
    {{end}}
    {{end}}
	case "cid":
		result = c.ContentCommon.CID
    default:
    	result = c.ContentCommon.Value( identifier )
    }
	return result
}

//Set value to a field
func (c *{{$struct_name}}) SetValue(identifier string, value interface{}) error {
	switch identifier {
        {{range $identifier, $fieldtype := .data_fields}}
          case "{{$identifier}}":
             c.{{$identifier|UpperName}} = value.({{$fieldtype}})
        {{end}}
        {{range $identifier, $fieldtype := .fields}}
            {{$type_settings := index $.def_fieldtype $fieldtype.FieldType}}
            {{if not $type_settings.IsRelation}}
            {{if not $fieldtype.IsOutput}}
            case "{{$identifier}}":
            c.{{$identifier|UpperName}} = value.(fieldtype.{{$type_settings.Value}})
            {{end}}
            {{end}}
        {{end}}
	default:
		err := c.ContentCommon.SetValue(identifier, value)
        if err != nil{
            return err
        }
	}
	//todo: check if identifier exist
	return nil
}

//Store content.
//Note: it will set id to CID after success
func (c *{{$struct_name}}) Store(transaction ...*sql.Tx) error {
	handler := db.DBHanlder()
	if c.CID == 0 {
		id, err := handler.Insert(c.TableName(), c.ToDBValues(), transaction...)
		c.CID = id
		if err != nil {
			return err
		}
	} else {
		err := handler.Update(c.TableName(), c.ToDBValues(), Cond("id", c.CID), transaction...)
		return err
	}
	return nil
}

func (c *{{$struct_name}})StoreWithLocation(){

}

//Delete content only
func (c *{{$struct_name}}) Delete(transaction ...*sql.Tx) error {
	handler := db.DBHanlder()
	contentError := handler.Delete(c.TableName(), Cond("id", c.CID), transaction...)
	return contentError
}

func init() {
	new := func() contenttype.ContentTyper {
		return &{{$struct_name}}{}
	}

	newList := func() interface{} {
		return &[]{{$struct_name}}{}
	}

    toList := func(obj interface{}) []contenttype.ContentTyper {
        contentList := *obj.(*[]{{$struct_name}})
        list := make([]contenttype.ContentTyper, len(contentList))
        for i, _ := range contentList {
            list[i] = &contentList[i]
        }
        return list
    }

	contenttype.Register("{{.name}}",
		contenttype.ContentTypeRegister{
			New:            new,
			NewList:        newList,
            ToList:         toList})
}
