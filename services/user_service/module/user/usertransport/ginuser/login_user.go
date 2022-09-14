package ginuser

import (
	"net/http"

	"github.com/gin-gonic/gin"
	goservice "github.com/phathdt/go-sdk"
	"github.com/phathdt/go-sdk/sdkcm"
	"gorm.io/gorm"
	"user_service/common"
	"user_service/module/user/userhdl"
	"user_service/module/user/usermodel"
	"user_service/module/user/userrepo"
	"user_service/module/user/userstorage"
)

func LoginUser(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data usermodel.UserLogin

		if err := c.ShouldBind(&data); err != nil {
			panic(sdkcm.ErrInvalidRequest(err))
		}

		db := sc.MustGet(common.KeyMainDB).(*gorm.DB)
		storage := userstorage.NewSqlStorage(db)
		repo := userrepo.NewUserRepo(storage)
		hdl := userhdl.NewLoginUserHdl(repo)

		token, err := hdl.Response(c.Request.Context(), &data)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, sdkcm.SimpleSuccessResponse(token))
	}
}
