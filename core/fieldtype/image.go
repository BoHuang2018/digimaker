package fieldtype

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/digimakergo/digimaker/core/definition"
	"github.com/digimakergo/digimaker/core/log"
	"github.com/digimakergo/digimaker/core/util"
)

type ImageHandler struct {
	definition.FieldDef
}

func (handler ImageHandler) LoadInput(input interface{}, mode string) (interface{}, error) {
	//todo: check image format
	str := fmt.Sprint(input)
	return str, nil
}

//Image can be loaded from rest, or local api
func (handler ImageHandler) BeforeStore(value interface{}, existing interface{}, mode string) (interface{}, error) {
	imagepath := value.(string)

	if imagepath != "" && (existing == nil || existing != nil && imagepath != existing.(string)) { //means there is a valid image change
		//todo: support other image services or remote image
		oldAbsPath := util.VarFolder() + "/" + imagepath

		//image path should be under temp
		temp := util.GetConfig("general", "upload_tempfolder")

		if _, err := os.Stat(oldAbsPath); err != nil {
			return nil, errors.New("Can't find file on " + oldAbsPath)
		}

		arr := strings.Split(imagepath, "/")
		filename := arr[len(arr)-1]

		//create 2 level folder
		rand := util.RandomStr(3)
		secondLevel := string(rand)
		firstLevel := string(rand[0])

		newFolder := "images/" + firstLevel + "/" + secondLevel
		newFolderAbs := util.VarFolder() + "/" + newFolder
		_, err := os.Stat(newFolderAbs)
		if os.IsNotExist(err) {
			err = os.MkdirAll(newFolderAbs, 0775)
			if err != nil {
				return nil, err
			}
		}

		newPath := newFolder + "/" + filename
		newPathAbs := util.VarFolder() + "/" + newPath

		underTemp := strings.HasPrefix(imagepath, temp)
		if underTemp {
			err = os.Rename(oldAbsPath, newPathAbs)
		} else {
			err = os.Link(oldAbsPath, newPathAbs) //todo: use better copy
		}
		if err != nil {
			errorMessage := "Can not copy/move image to target " + imagepath + ". error: " + err.Error()
			return nil, errors.New(errorMessage)
		}

		err = GenerateThumbnail(newPath)
		if err != nil {
			return nil, err
		}
		return newPath, nil
	} else {
		return existing, nil
	}
}

//After delete content, delete image&thumbnail, skip error if there is any wrong(eg. image not existing).
func (handler ImageHandler) AfterDelete(value interface{}) error {
	path := util.VarFolder() + "/" + value.(string)
	err := os.Remove(path)
	if err != nil {
		message := fmt.Sprintf("Deleting image(path: %v) of %v error: %v. Deleting continued.", path, handler.FieldDef.Identifier, err.Error())
		log.Warning(message, "system")
	}

	thumbnail := ThumbnailFolder() + "/" + path
	err := os.Remove(thumbnail)
	if err != nil{
		message := fmt.Sprintf("Deleting image thumnail(path: %v) of %v error: %v. Deleting continued.", path, handler.FieldDef.Identifier, err.Error()), "system" )
		log.Warning(message, "system")
	}
	return nil
}

func (handler ImageHandler) DBField() string {
	return "VARCHAR (500) NOT NULL DEFAULT ''"
}

func GenerateThumbnail(imagePath string) error {
	varFolder := util.VarFolder()
	size := util.GetConfig("general", "image_thumbnail_size")
	thumbCacheFolder := ThumbnailFolder()
	return util.ResizeImage(varFolder+"/"+imagePath, thumbCacheFolder+"/"+imagePath, size)
}

func ThumbnailFolder() string {
	thumbFolder := util.VarFolder() + "/" + util.GetConfig("general", "image_thumbnail_folder")
	return thumbFolder
}

func init() {
	Register(
		Definition{
			Name:     "image",
			DataType: "string",
			NewHandler: func(def definition.FieldDef) Handler {
				return ImageHandler{FieldDef: def}
			}})
}
