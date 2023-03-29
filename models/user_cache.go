package models

import (
	"Gin_WebSocket_IM/utils"
	"context"
	"time"
)

func SetUserOnlineInfo(key string, val []byte, timeTTL time.Duration) {
	ctx := context.Background()
	utils.Rdb.Set(ctx, key, val, timeTTL)
}
