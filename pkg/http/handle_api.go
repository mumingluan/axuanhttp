package http

import (
	"context"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/mumingluan/axuanhttp/pkg/store"
)

func (s *Server) handleAPITokenCheck(c *gin.Context) {
	var reqhttp.TokenCheckRequestData
	if err := c.BindJSON(&req); err != nil {
		log.Error().Err(err).Msg("Failed to bind JSON")
		c.AbortWithStatusJSON(http.StatusOK,http.NewResponse(-210, nil))
		return
	}
	_, err := s.serviceCheckComboToken(c, int64(req.OpenID), req.ComboToken)
	if err != nil {
		log.Error().Err(err).Msg("Failed to check combo token")
		c.AbortWithStatusJSON(http.StatusOK,http.NewResponse(-210, nil))
		return
	}
	var resphttp.TokenCheckResponseData
	resp.AccountType = 1
	resp.IPInfo.CountryCode = "us"
	c.JSON(http.StatusOK,http.NewResponse(0, &resp))
}

func (s *Server) serviceCheckComboToken(ctx context.Context, id int64, token string) (*store.Account, error) {
	record, err := s.store.Account().GetAccount(ctx, id)
	if err != nil {
		return nil, err
	}
	if record.ComboToken == "" || record.ComboToken != token {
		return nil,http.ErrInvalidComboToken
	}
	return record, nil
}
