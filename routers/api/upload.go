package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tangxin-demo/pkg/e"
	"tangxin-demo/pkg/logging"
	"tangxin-demo/pkg/setting"
	"tangxin-demo/pkg/upload"
	"tangxin-demo/pkg/util"
	"time"
)

func UploadImage(c *gin.Context) {
	code := e.SUCCESS
	data := make(map[string]string)

	file, image, err := c.Request.FormFile("image")
	if err != nil {
		logging.Warn(err)
		code = e.ERROR
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": data,
		})
		return
	}

	if image == nil {
		code = e.INVALID_PARAMS
	} else {
		imageName := util.EncodeMD5(image.Filename + strconv.FormatInt(time.Now().Unix(), 10))
		imagePath := setting.AppSetting.ImageSavePath + imageName
		if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
			code = e.ERROR_UPLOAD_CHECK_IMAGE_FORMAT
		} else {
			err := upload.CheckImage(imagePath)
			if err != nil {
				logging.Warn(err)
				code = e.ERROR_UPLOAD_CHECK_IMAGE_FAIL
			} else {
				err = c.SaveUploadedFile(image, imagePath)
				if err != nil {
					logging.Warn(err)
					code = e.ERROR_UPLOAD_SAVE_IMAGE_FAIL
				} else {
					data["image_url"] = setting.AppSetting.ImagePrefixUrl + "/" + imageName
				}
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
