package handlers

import (
	_ "crypto/sha256"
	"encoding/json"
	_"fmt"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"net/http"
	"task-assign/models"
	"task-assign/utils"
)

var (
	errUnableToLoadData = "Unable to parse post data"
)

/*type register struct{
	Name string `json:"name" validate:"min=4,alpha"`
	Email string `json:"email" validate:"(required,email)`
	Password string `json:"pass" validate:"required"`
	Phone string `json:"phone"`
	Gender string `json:"gender"`
}*/

func Registration(w http.ResponseWriter, r *http.Request) {
	//read post json dat
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, errUnableToLoadData, 500)
		return
	}

	var reg models.Worker
	if err = json.Unmarshal(b, &reg); err != nil {
		http.Error(w, "Unable to load json data", 500)
		return
	}
	validate := validator.New()

	//fmt.Println(string(nerateFromPassword([]byte(reg.Password),4)

	err = validate.Struct(reg)
	if err != nil {
		http.Error(w, "Error while validating", 500)
		return
	}

	u, _ := bcrypt.GenerateFromPassword([]byte(reg.Password),4)
	reg.Password = string(u)
	//if !result{
	//	http.Error(w,"Incorrect or invalid data",500)
	//	return
	//}

	/*worker := models.Worker{
		Fullname: reg.Name,
		Email: reg.Email,
		Phone: reg.Phone,
		Password: "",
		Gender: reg.Gender,
	}
	*/
	db := utils.DB
	err = db.Create(&reg).Error
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	success := map[string]string{
		"msg": "Successfully added 😍to database",
	}
	msg, err := json.Marshal(success)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(msg)
	//
	//var worker models.Worker
	//worker.Fullname = reg.Name
	//worker.Email = reg.Email
	//worker.Phone = reg.Phone
	//worker.Password = ""
	//worker.Gender = reg.Gender

}
func Login(w http.ResponseWriter, r *http.Request) {
	//type login struct {
	//	Email    string `json:"email" validate:"required,email"`
	//	Password string `json:"pass" validate:"required"`
	//}

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	var data models.Worker
	if err = json.Unmarshal(b, &data); err != nil {
		http.Error(w, errUnableToLoadData, 500)
	}

	validate := validator.New()
	if err=validate.Struct(data);err!=nil{
		http.Error(w,err.Error(),500)
		return
	}
	var worker models.Worker
	db := utils.DB
	if err = db.First(&worker,"Email=?",data.Email).Error;err!=nil{
		http.Error(w,err.Error(),500)
		return
	}
	if err=bcrypt.CompareHashAndPassword([]byte(worker.Password),[]byte(data.Password));err!=nil{
		http.Error(w,err.Error(),500)
		return
	}
	w.Write([]byte("Successfully logged in"))
}
