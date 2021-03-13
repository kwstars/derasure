package services

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/kwstars/derasure/internal/app/model"
	"net/http"
	"strconv"
)

var Set = wire.NewSet(wire.Struct(new(Services), "*"))

const (
	MINIGAME_ELIMINATE = iota + 1
	MINIGAME_BANQUET
	MINIGAME_FISHING
	MINIGAME_KITE
)

type Services struct {
	Repostiory *model.Repostiory
}

func (f *Services) Del(c *gin.Context) {
	var deleteType string
	uid, ok1 := c.GetPostForm("uid")
	tmpType, ok2 := c.GetPostForm("type")
	if !ok1 || !ok2 {
		fmt.Println(ok1, ok2)
		c.HTML(http.StatusBadRequest, "index.tmpl", gin.H{
			"msg":  "请求参数不正确",
			"uid":  uid,
			"type": tmpType,
		})
		return
	}
	tp, _ := strconv.Atoi(tmpType)

	if err := f.Repostiory.CheckAccountExist(context.Background(), uid); err != nil {
		c.HTML(http.StatusBadRequest, "index.tmpl", gin.H{
			"msg":  err,
			"uid":  uid,
			"type": tp,
		})
		return
	}

	switch tp {
	case MINIGAME_ELIMINATE:
		deleteType = "eliminate:" + uid
	case MINIGAME_BANQUET:
		deleteType = "banquet:" + uid
	case MINIGAME_FISHING:
		deleteType = "fishing:" + uid
	case MINIGAME_KITE:
		deleteType = "kite:" + uid
	default:
		c.HTML(http.StatusInternalServerError, "index.tmpl", gin.H{"msg": "无效选项", "uid": uid, "type": tp})
		return
	}

	if err := f.Repostiory.DelKey(context.Background(), deleteType); err != nil {
		c.HTML(http.StatusInternalServerError, "index.tmpl", gin.H{
			"msg":  err,
			"uid":  uid,
			"type": tp,
		})
		return
	}

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"msg":  "执行成功",
		"uid":  uid,
		"type": tp,
	})
}
