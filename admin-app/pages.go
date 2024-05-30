package admin_app

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/matheusgomes28/urchin/common"
	"github.com/matheusgomes28/urchin/database"
	"github.com/rs/zerolog/log"
)

func postPageHandler(database database.Database) func(*gin.Context) {
	return func(c *gin.Context) {
		var add_page_request AddPageRequest
		if c.Request.Body == nil {
			c.JSON(http.StatusBadRequest, common.MsgErrorRes("no body provided to create page"))
			return
		}
		decoder := json.NewDecoder(c.Request.Body)
		err := decoder.Decode(&add_page_request)

		if err != nil {
			log.Warn().Msgf("invalid create `page` request: %v", err)
			c.JSON(http.StatusBadRequest, common.ErrorRes("invalid request body", err))
			return
		}

		link := strings.ReplaceAll(strings.ToLower(add_page_request.Title), " ", "-")
		page, err := database.AddPage(add_page_request.Title, add_page_request.Content, link, add_page_request.ParentId)

		if err != nil {
			log.Warn().Msgf("could not persist page: %v", err)
			c.JSON(http.StatusBadRequest, common.ErrorRes("creating page failed", err))
			return
		}

		c.JSON(http.StatusOK, PostPageResponse{
			page.Id,
			page.Title,
			page.Link,
			page.ParentId,
		})
	}
}
