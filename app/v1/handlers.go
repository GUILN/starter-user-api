package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guiln/starter-log/logger"
	"github.com/guiln/starter-log/messages"
)

type UserHandler struct {
	lggr logger.CompanyLogger
}

// UserLogin godoc
// @BasePath /users/v1

// @Sumary User Login
// @Schemes
// @Description logs user into the application returning an JWT token.
// @Tags Example
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router /login [post]
func (uh *UserHandler) UserLogin(ctx *gin.Context) {
	uh.lggr.Info(messages.New("").Message())
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello",
	})
}

func NewUserHandler(lggr logger.CompanyLogger) *UserHandler {
	return &UserHandler{
		lggr: lggr,
	}
}
