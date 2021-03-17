package erasure

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/kwstars/derasure/pkg/retno"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type IDelData interface {
	Execution(ctx context.Context, uid string) error
}

var DelServiceSet = wire.NewSet(wire.Struct(new(ErasureServices), "*"))

const (
	MINIGAME_ELIMINATE = iota + 1
	MINIGAME_BANQUET
	MINIGAME_FISHING
	MINIGAME_KITE
	LIMITED_GIFT
)

type ErasureServices struct {
	Logger      *zap.Logger
	Eliminate   *Eliminate
	Banquet     *Banquet
	Fishing     *Fishing
	Kite        *Kite
	LimitedGift *LimitedGift
}

func (s *ErasureServices) Del(c *gin.Context) {
	uid, ok1 := c.GetPostForm("uid")
	tmpType, ok2 := c.GetPostForm("type")
	if !ok1 || !ok2 {
		s.Logger.Error("Incorrect request parameters", zap.String("uid", uid), zap.String("type", tmpType))
		c.HTML(http.StatusBadRequest, "index.tmpl", gin.H{
			"msg":  retno.INVALID_DATA.String(),
			"uid":  uid,
			"type": tmpType,
		})
		return
	}
	tp, _ := strconv.Atoi(tmpType)
	var err error
	switch tp {
	case MINIGAME_ELIMINATE:
		err = s.Eliminate.Execution(context.Background(), uid)
	case MINIGAME_BANQUET:
		err = s.Banquet.Execution(context.Background(), uid)
	case MINIGAME_FISHING:
		err = s.Fishing.Execution(context.Background(), uid)
	case MINIGAME_KITE:
		err = s.Kite.Execution(context.Background(), uid)
	case LIMITED_GIFT:
		err = s.LimitedGift.Execution(context.Background(), uid)
	default:
		s.Logger.Error("Incorrect request parameters", zap.String("uid", uid), zap.String("type", tmpType))
		c.HTML(http.StatusInternalServerError, "index.tmpl", gin.H{"msg": retno.INVALID_DATA.String(), "uid": uid, "type": tp})
		return
	}

	if err != nil {
		s.Logger.Error("Failed to delete", zap.String("uid", uid), zap.String("type", tmpType))
		c.HTML(http.StatusInternalServerError, "index.tmpl", gin.H{
			"msg":  err.Error(),
			"uid":  uid,
			"type": tp,
		})
	} else {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"msg":  retno.OK,
			"uid":  uid,
			"type": tp,
		})
	}

	return
}
