package sitekit

import (
	"context"
	"strconv"
	"strings"

	"github.com/digimakergo/digimaker/core/contenttype"
	"github.com/digimakergo/digimaker/core/log"
	"github.com/digimakergo/digimaker/core/util"
)

const templateViewContent = "content_view"
const overrideFile = "template_override"

var overrideFieldtypes = []string{"radio", "select", "checkbox"}

//TemplateFolder() returns folder of templates. eg. under "templates" or "web/templates"
func TemplateFolder() string {
	path := util.AbsHomePath() + "/" + util.GetConfig("general", "template_folder")
	return path
}

//Get content view template.
func GetContentTemplate(content contenttype.ContentTyper, viewmode string, settings SiteSettings, ctx context.Context) string {
	templateFolder := TemplateFolder()

	matchData := map[string]interface{}{}
	matchData["viewmode"] = viewmode
	matchData["site"] = settings.Site
	matchData["contenttype"] = content.ContentType()
	location := content.GetLocation()
	if location != nil {
		matchData["id"] = location.ID
		matchData["under"] = location.Path()
		matchData["level"] = location.Depth
		matchData["section"] = location.Section
	}
	for field, fieldDef := range content.Definition().FieldMap {
		if util.Contains(overrideFieldtypes, fieldDef.FieldType) {
			matchData["field_"+field] = content.Value(field)
		}
	}

	matchLog := []string{}
	path, matchLog := MatchTemplate(templateViewContent, matchData)

	log.Debug("Matching on "+content.GetName()+", got: "+path+"\n "+strings.Join(matchLog, "\n"), "template-match", ctx)

	result := ""
	if path != "" {
		pathWithTemplateFolder := settings.TemplateFolder + "/" + path
		if util.FileExists(templateFolder + "/" + pathWithTemplateFolder) {
			result = pathWithTemplateFolder
		} else {
			log.Warning("Matched file not found: "+path, "template", ctx)
		}
	}
	return result
}

//MatchTemplate returns overrided template based on override config(eg. template_override.yaml)
func MatchTemplate(viewSection string, matchData map[string]interface{}, fileName ...string) (string, []string) {
	overrideFileName := ""
	result := ""
	matchLog := []string{}
	if len(fileName) == 0 {
		overrideFileName = overrideFile
		//if there is include, match in included file
		includeI := util.GetConfigSectionAll("include", overrideFileName)
		if includeI != nil {
			for _, item := range includeI.([]interface{}) {
				includeRules := map[string]interface{}{}
				includedFile := ""
				for key, value := range item.(map[interface{}]interface{}) {
					keyS := key.(string)
					if keyS == "file" {
						includedFile = value.(string)
					} else {
						includeRules[keyS] = value
					}
				}
				includeMatched, _ := util.MatchCondition(includeRules, matchData)
				if includeMatched {
					matchLog = append(matchLog, "Matching on include file: "+includedFile)
					var includedMatchLog []string
					result, includedMatchLog = MatchTemplate(viewSection, matchData, includedFile)
					matchLog = append(matchLog, includedMatchLog...)
					if result != "" {
						return result, matchLog
					} else {
						matchLog = append(matchLog, "Not matched on include file : "+includedFile)
					}
				}
			}
		}
	} else {
		overrideFileName = fileName[0]
	}
	rulesI := util.GetConfigSectionAll(viewSection, overrideFileName)
	if rulesI == nil {
		return "", []string{"view section not found: " + viewSection}
	}
	rules := rulesI.([]interface{})
	for i, item := range rules {
		conditions := map[string]interface{}{}
		to := ""
		for key, value := range item.(map[interface{}]interface{}) {
			keyStr := key.(string)
			if keyStr == "to" {
				to = value.(string) //todo: have a better name instead of to
				continue
			}
			conditions[keyStr] = value
		}

		matchLog = append(matchLog, "Matching on rule"+strconv.Itoa(i)+" on file "+overrideFile)
		matched, currentMatchLog := util.MatchCondition(conditions, matchData)
		matchLog = append(matchLog, currentMatchLog...)
		if matched {
			washedVars := map[string]string{}
			for key, value := range matchData {
				switch value.(type) {
				case string:
					washedVars[key] = value.(string)
				case int:
					washedVars[key] = strconv.Itoa(value.(int))
				}
			}
			result = util.ReplaceStrVar(to, washedVars)
			break
		}
	}

	return result, matchLog
}
