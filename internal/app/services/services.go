package services

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/kwstars/derasure/internal/app/model"
	"log"
	"net/http"
)

var Set = wire.NewSet(wire.Struct(new(Services), "*"))

type Services struct {
	Repostiory *model.Repostiory
}

func (f *Services) DelEliminate(c *gin.Context) {
	if err := f.Repostiory.Del(context.Background(), "eliminate"); err != nil {
		log.Printf("%+v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "删除失败",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "删除成功",
	})
}

func (f *Services) DelBanquet(c *gin.Context) {
	if err := f.Repostiory.Del(context.Background(), "banquet"); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "删除失败",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "删除成功",
	})
}

func (f *Services) DelFishing(c *gin.Context) {
	if err := f.Repostiory.Del(context.Background(), "fishing"); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "删除失败",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "删除成功",
	})
}

func (f *Services) DelKite(c *gin.Context) {
	if err := f.Repostiory.Del(context.Background(), "kite"); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "删除失败",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "删除成功",
	})
}
