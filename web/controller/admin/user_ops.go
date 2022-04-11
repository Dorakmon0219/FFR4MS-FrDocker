package admin

import (
	"context"
	"net/http"

	"gitee.com/zengtao321/frdocker/db"
	"gitee.com/zengtao321/frdocker/web/entity"
	"gitee.com/zengtao321/frdocker/web/entity/R"
	"gitee.com/zengtao321/frdocker/web/service/token"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

var userMgo = db.GetUserMgo()

func AddUser(c *gin.Context) {
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

func GetUserList(c *gin.Context) {
	var users []entity.UserEntity
	userMgo.FindAll().All(context.TODO(), &users)
	c.JSON(http.StatusOK, R.OK(users))
}

func DeleteUser(c *gin.Context) {
	var user entity.UserEntity
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, R.Error(http.StatusBadRequest, "", nil))
		return
	}
	tokenStr := c.Request.Header["Authorization"][0]
	claims, _ := token.ParseToken(tokenStr)
	currentUserId := claims.UserId
	if currentUserId == user.Id {
		c.JSON(http.StatusBadRequest, R.Error(http.StatusBadRequest, "Cannot delete yourself!", nil))
	}
	var filter = bson.D{{Key: "id", Value: user.Id}}
	deletecCount := userMgo.Delete(filter)
	if deletecCount == 0 {
		c.JSON(http.StatusBadRequest, R.Error(http.StatusBadRequest, "No such user!", nil))
		return
	}
	c.JSON(http.StatusOK, R.OK(nil))
}

func UpdateUser(c *gin.Context) {
	var user entity.UserEntity
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, R.Error(http.StatusBadRequest, "", nil))
		return
	}
	var filter = bson.D{{Key: "id", Value: user.Id}}
	var tempUser *entity.UserEntity
	userMgo.FindOne(filter).Decode(&tempUser)
	if tempUser == nil {
		c.JSON(http.StatusBadRequest, R.Error(http.StatusBadRequest, "No such user!", nil))
		return
	}
	if tempUser.Password != user.Password {
		cryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		user.Password = string(cryptedPassword)
	}
	userMgo.ReplaceOne(filter, user)
	c.JSON(http.StatusOK, R.OK(user))
}
