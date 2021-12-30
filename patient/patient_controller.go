package patient

import (
	"encoding/json"
	"io"
	"main/common"
	"main/database"
	"main/patient/model"
	"main/token"
	"net/http"
	"os"
	"reflect"
	"strings"

	"github.com/labstack/echo/v4"
	pusher "github.com/pusher/pusher-http-go"
	"gopkg.in/go-playground/validator.v9"
)

type (
	PatientController struct{}
)

type (
	Auth_PatientController struct{}
)

func (controller PatientController) Routes() []common.Route {

	return []common.Route{
		{
			Method:  echo.POST,
			Path:    "/patient",
			Handler: controller.AddPatient,
		},
		{
			Method:     echo.GET,
			Path:       "/patients",
			Handler:    controller.GetPatients,
			Middleware: []echo.MiddlewareFunc{common.JwtMiddleWare()},
		},
		{
			Method:  echo.POST,
			Path:    "/upload-photo",
			Handler: controller.UploadPhoto,
		},
		{
			Method:  echo.POST,
			Path:    "/login",
			Handler: controller.Login,
		},
	}
}

func (controller PatientController) AddPatient(ctx echo.Context) error {

	patient := &model.Patient{}
	err := json.NewDecoder(ctx.Request().Body).Decode(&patient)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	if ok, errors := validate_input(*patient); !ok {
		return ctx.JSON(http.StatusBadRequest, errors)
	}

	db := database.GetInstance()
	db.Save(patient)

	return ctx.JSON(http.StatusCreated, patient)
}

func validate_input(data interface{}) (bool, string) {

	var validate *validator.Validate
	validate = validator.New()
	err := validate.Struct(data)
	if err != nil {
		if err, ok := err.(*validator.InvalidValidationError); ok {
			panic(err)
		}

		var errors string
		reflected := reflect.ValueOf(data)
		for _, err := range err.(validator.ValidationErrors) {
			field, _ := reflected.Type().FieldByName(err.StructField())
			var name string
			if name = field.Tag.Get("json"); name == "" {
				name = strings.ToLower(err.StructField())
			}

			switch err.Tag() {
			case "required":
				errors = "The " + name + " is required"
			}
		}
		return false, errors
	}

	return true, ""
}

func (controller PatientController) GetPatients(ctx echo.Context) error {

	var patient model.Patient
	db := database.GetInstance()
	data := db.Find(&patient)

	return ctx.JSON(http.StatusOK, data.Value)
}

func (controller PatientController) UploadPhoto(ctx echo.Context) error {

	var client = pusher.Client{
		AppID:   "1323924",
		Key:     "99b69fd4b4bbc2c8265d",
		Secret:  "ba832086f62551866ac3",
		Cluster: "ap1",
		Secure:  true,
	}
	file, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	src, err := file.Open()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	defer src.Close()
	filePath := "./patient/" + file.Filename
	fileSrc := " http://localhost:1323/api/upload-photo/" + file.Filename
	dst, err := os.Create(filePath)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	defer dst.Close()
	if _, err = io.Copy(dst, src); err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	client.Trigger("photo-stream", "new-photo", fileSrc)
	return ctx.JSON(http.StatusCreated, fileSrc)

}

func (controller PatientController) Login(ctx echo.Context) error {

	user := model.User{}
	err := json.NewDecoder(ctx.Request().Body).Decode(&user)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	if user.Username != "admin" || user.Password != "1234" {
		return ctx.JSON(http.StatusUnauthorized, echo.Map{"error": "Please, try enter passowrd or username again."})
	}
	accessToken, err := token.CreateToken(user.Username, token.AccessToken)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	user.AccessToken = accessToken
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, echo.Map{"token": user.AccessToken})
}
