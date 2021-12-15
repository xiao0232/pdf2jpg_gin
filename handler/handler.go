package handler

import (
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type ResponseInfo struct {
	Count    int    `json:"count"`
	BaseName string `json:"baseName"`
}

type SuccessResponse struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Info    ResponseInfo `json:"info"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func getFileNameWithoutExt(path string) string {
	// https://qiita.com/KemoKemo/items/d135ddc93e6f87008521
	// Fixed with a nice method given by mattn-san
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}

func countFile(path string) (int, error) {
	files, err := ioutil.ReadDir(path)
	return len(files), err
}

func pdf2img(pdfBaseName string) error {
	err := exec.Command("pdftoppm", "-rx", "300", "-ry", "300", "-jpeg", "pdf/"+pdfBaseName+"/"+pdfBaseName+".pdf", "pdf/"+pdfBaseName+"/"+pdfBaseName).Run()
	return err
}

// Upload
// @Summary      recieve pdf from FrontEnd
// @Description  get pdf
// @Tags         upload
// @Accept       mpfd
// @Produce      json
// @Param        pdf  formData  file  true  "	"
// @Success      200  {object}  SuccessResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /upload/pdf [post]
func UploadHandler(c *gin.Context) {
	// get file
	file, err := c.FormFile("pdf")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	fBaseName := getFileNameWithoutExt(file.Filename)

	if f, err := os.Stat("pdf/" + fBaseName + "/"); os.IsNotExist(err) || !f.IsDir() {
		// check save directory
		err = os.MkdirAll("./pdf/"+fBaseName, os.ModePerm)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}

		// save file
		err = c.SaveUploadedFile(file, "pdf/"+fBaseName+"/"+fBaseName+".pdf")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}

		// convert pdf to jpeg
		err = pdf2img(fBaseName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
	}

	count, err := countFile("pdf/" + fBaseName + "/")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	responseInfo := ResponseInfo{
		Count:    count - 1,
		BaseName: fBaseName,
	}

	c.JSON(http.StatusOK, gin.H{"message": "success upload and convert pdf to jpg.", "info": responseInfo})
}
