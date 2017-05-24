package control

import (
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/ClimbingMoon/camm/dto"
	m "github.com/ClimbingMoon/camm/model"
	"github.com/labstack/echo"
)

// GetUserProfile 获取用户信息
func GetUserProfile(c echo.Context) error {
	cc := c.(*dto.CustomContext)

	userProfile := &dto.UserProfile{
		ID:          cc.User.ID.Hex(),
		Username:    cc.User.Username,
		Email:       cc.User.Email,
		IsAdmin:     cc.User.IsAdmin,
		IsAbleLogin: cc.User.IsAbleLogin,
		CreateTime:  cc.User.CreateTime,
	}

	return cc.JSON(http.StatusOK, userProfile)
}

// GetUserProfileByID 根据id获取用户信息
func GetUserProfileByID(c echo.Context) error {
	idStr := c.Param("id")

	if !bson.IsObjectIdHex(idStr) {
		// TODO 404
		return c.String(http.StatusOK, "Invalid user id")
	}

	query := m.User{
		ID: bson.ObjectIdHex(idStr),
	}

	var targetUser m.User
	var err error
	if targetUser, err = query.GetOneByID(); err != nil {
		c.String(http.StatusOK, "user "+err.Error())
	}

	userProfile := &dto.UserProfile{
		ID:          targetUser.ID.Hex(),
		Username:    targetUser.Username,
		Email:       targetUser.Email,
		IsAdmin:     targetUser.IsAdmin,
		IsAbleLogin: targetUser.IsAbleLogin,
		CreateTime:  targetUser.CreateTime,
	}

	return c.JSON(http.StatusOK, userProfile)
}
