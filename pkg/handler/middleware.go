package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
	emptyAuthHeader     = "empty auth header"
	invalidAuthHeader   = "invalid auth header"
	userNotFound        = "user is not found"
	invalidUserId       = "user id is of invalid type"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, emptyAuthHeader)
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, invalidAuthHeader)
		return
	}

	userId, err := h.services.Validate(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
	}

	c.Set(userCtx, userId)
}

func getUserId(c *gin.Context) (string, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, userNotFound)
	}

	idStr, ok := id.(string)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, invalidUserId)
	}

	return idStr, nil
}
