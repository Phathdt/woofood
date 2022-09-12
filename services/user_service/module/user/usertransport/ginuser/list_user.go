package ginuser

import (
	"net/http"

	"user_service/common"
	"user_service/module/user/userhdl"
	"user_service/module/user/userrepo"
	"user_service/module/user/userstorage"

	goservice "github.com/phathdt/go-sdk"
	"github.com/phathdt/go-sdk/sdkcm"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListPost(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := sc.MustGet(common.KeyMainDB).(*gorm.DB)
		userStorage := userstorage.NewUserStorage(db)
		userRepo := userrepo.NewUserRepo(userStorage)
		hdl := userhdl.NewListUserHdl(userRepo)

		users, err := hdl.Response(c.Request.Context())
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, sdkcm.SimpleSuccessResponse(users))
	}
}
