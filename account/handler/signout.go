<<<<<<< HEAD
package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thienbui11/MindStack/account/model"
	"github.com/thienbui11/MindStack/account/model/apperrors"
)

// Signout handler
func (h *Handler) Signout(c *gin.Context) {
	user := c.MustGet("user")

	ctx := c.Request.Context()
	if err := h.TokenService.Signout(ctx, user.(*model.User).UID); err != nil {
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user signed out successfully!",
	})
}
=======
package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jacobsngoodwin/memrizr/account/model"
	"github.com/jacobsngoodwin/memrizr/account/model/apperrors"
)

// Signout handler
func (h *Handler) Signout(c *gin.Context) {
	user := c.MustGet("user")

	ctx := c.Request.Context()
	if err := h.TokenService.Signout(ctx, user.(*model.User).UID); err != nil {
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user signed out successfully!",
	})
}
>>>>>>> 3999c5b64d9e82200d4e850d267e0d3ed56f0643
