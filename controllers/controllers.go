package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"../model"
)

var (
	modul_i = model.Image{}
	modul_t = model.Tag{}
)

func errorHandler(w http.ResponseWriter, httpStatus int, errMsg string) {
	result := map[string]interface{}{
		"response": errMsg,
	}
	responseWithJson(w, httpStatus, result)
}

func FindImage_all(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	var image []model.Image

	image = modul_i.GetAllImage()

	responseWithJson(w, http.StatusOK, image)
}

func FindImage_by_tag(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	params := map[string]interface{}{
		"tag": getParameter(r.URL.Query(), "tag", "_string"),
	}

	tag_id := modul_t.GetTagId(params["tag"].(string))

	var image []model.Image
	image = modul_i.GetImageByTag(tag_id)

	responseWithJson(w, http.StatusOK, image)
}

func CreateImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Content-Type", "application/json")

	defer r.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(r.Body)
	r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes)) // 讀取後恢復

	var images []model.Image
	if err := json.NewDecoder(r.Body).Decode(&images); err != nil {
		errorHandler(w, http.StatusBadRequest, err.Error())
		return
	}

	for _, image := range images {
		// image.Create_time = time.Now()
		// image.Create_time = image.Create_time
		image = modul_i.InsertImage(image)
	}
	// update tag: image_num
	tag_id := images[0].Tag
	count := modul_i.CountImageNum(tag_id)
	modul_t.UpdateTagImageNum(tag_id, count)

	response := map[string]string{
		"result": "success",
	}
	responseWithJson(w, http.StatusOK, response)

}

func DeleteImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	params := map[string]interface{}{
		"id": getParameter(r.URL.Query(), "id", "_int"),
	}
	image_id := params["id"].(int)

	// find image tag
	var image model.Image
	image = modul_i.GetImageById(image_id)
	tag_id := image.Tag

	// delete image
	is_success := modul_i.DeleteImageById(image_id)
	response := map[string]bool{
		"result": is_success,
	}
	responseWithJson(w, http.StatusOK, response)

	// update tag: image_num
	count := modul_i.CountImageNum(tag_id)
	modul_t.UpdateTagImageNum(tag_id, count)
}

func NewPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	var image []model.Image

	image = modul_i.GetNewPost()

	responseWithJson(w, http.StatusOK, image)
}

func FindImageNum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	params := map[string]interface{}{
		"tag": getParameter(r.URL.Query(), "tag", "_string"),
	}

	image_num := modul_t.GetTagNum(params["tag"].(string))
	response := map[string]int{
		"image_num": image_num,
	}

	responseWithJson(w, http.StatusOK, response)
}

func AddFavorite(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	params := map[string]interface{}{
		"id": getParameter(r.URL.Query(), "id", "_int"),
	}

	var image model.Image
	image = modul_i.UpdateFavorite(params["id"].(int))

	responseWithJson(w, http.StatusOK, image)
}

func AddLocation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	params := map[string]interface{}{
		"id":       getParameter(r.URL.Query(), "id", "_int"),
		"location": getParameter(r.URL.Query(), "location", "_string"),
	}

	var image model.Image
	image = modul_i.UpdateLocation(params["id"].(int), params["location"].(string))

	responseWithJson(w, http.StatusOK, image)
}
