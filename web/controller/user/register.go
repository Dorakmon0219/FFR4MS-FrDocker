package user

import (
	"net/http"

	"gitee.com/zengtao321/frdocker/db"
	"gitee.com/zengtao321/frdocker/web/entity"
	"gitee.com/zengtao321/frdocker/web/entity/R"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func RegisterController(c *gin.Context) {
	userMgo := db.GetUserMgo()
	var user entity.UserEntity
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, R.Error(http.StatusBadRequest, "", nil))
		return
	}
	var filter = bson.D{{Key: "username", Value: user.Username}}
	var tempUser *entity.UserEntity
	userMgo.FindOne(filter).Decode(&tempUser)
	if tempUser != nil {
		c.JSON(http.StatusBadRequest, R.Error(http.StatusBadRequest, "Username already exists!", nil))
		return
	}
	user.Role = "USER"
	cryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(cryptedPassword)
	user.Id = uuid.New().String()
	userMgo.InsertOne(user)
	c.JSON(http.StatusOK, R.OK(nil))
}
