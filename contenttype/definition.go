//Author xc, Created on 2019-03-28 20:00
//{COPYRIGHTS}

package contenttype

import (
	"dm/util"
)

type ContentTypeSetting struct {
	TableName  string                  `json:"table_name"`
	Versioning bool                    `json:"versioning"`
	Fields     map[string]ContentField `json:"fields"`
}

type ContentField struct {
	FieldType string `json:"type"`
	Required  bool   `json:"required"`
}

//ContentTypeDefinition Content types which defined in contenttype.json
var contentTypeDefinition map[string]ContentTypeSetting

//LoadDefinition Load all setting in file into memory.
//
// It will not load anything unless all json' format matches the struct definition.
//
func LoadDefinition() error {

	//Load contenttype.json into ContentTypeDefinition
	var contentDef map[string]ContentTypeSetting
	err := util.UnmarshalData(util.ConfigPath()+"/contenttype.json", &contentDef)
	if err != nil {
		return err
	}

	contentTypeDefinition = contentDef

	return nil
}

func GetContentDefinition(contentType string) ContentTypeSetting {
	definition := contentTypeDefinition
	result, ok := definition[contentType]
	if ok {
		return result
	} else {
		return ContentTypeSetting{}
	}

}
