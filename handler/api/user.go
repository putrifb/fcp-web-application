package api

import (
	"a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/service"
	"net/http"

	// "time"

	"github.com/gin-gonic/gin"
	// "github.com/golang-jwt/jwt"
)

type UserAPI interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	GetUserTaskCategory(c *gin.Context)
}

type userAPI struct {
	userService service.UserService
}

func NewUserAPI(userService service.UserService) *userAPI {
	return &userAPI{userService}
}

func (u *userAPI) Register(c *gin.Context) {
	var user model.UserRegister

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("invalid decode json"))
		return
	}

	if user.Email == "" || user.Password == "" || user.Fullname == "" {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("register data is empty"))
		return
	}

	var recordUser = model.User{
		Fullname: user.Fullname,
		Email:    user.Email,
		Password: user.Password,
	}

	recordUser, err := u.userService.Register(&recordUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse("error internal server"))
		return
	}

	c.JSON(http.StatusCreated, model.NewSuccessResponse("register success"))
}

func (u *userAPI) Login(c *gin.Context) {
	//memanggil userService.Login dengan parameter context
	var usLogin model.UserLogin
	var user model.User

	if err := c.ShouldBind(&usLogin); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid decode json"}) //400
		return
	}

	user.Email = usLogin.Email
	user.Password = usLogin.Password

	token, err := u.userService.Login(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: "error internal server"}) //500
		return
	}
	c.SetCookie("session_token", *token, 3600, "", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{"user_id": user.ID, "message": "login success"}) //200
}

func (u *userAPI) GetUserTaskCategory(c *gin.Context) {
	//mendapatkan daftar tugas pengguna dengan kategori yang terkait.
	usList, err := u.userService.GetUserTaskCategory()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err) //500
		return
	}
	c.JSON(http.StatusOK, usList) //200
	// TODO: answer here
}
