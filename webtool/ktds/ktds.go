package ktds

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type V struct {
	DateTime string
	Message  string
}

func Get(ctx *gin.Context) {
	// 空调默认关闭时间 明天 04:00
	V := V{
		time.Now().AddDate(0, 0, 1).Format("2006-01-02") + "T04:00",
		"",
	}

	if ctx.Query("kttime") != "" {
		V.DateTime = fmt.Sprintf("%sT%s",
			time.Now().AddDate(0, 0, 1).Format("2006-01-02"),
			ctx.Query("kttime"),
		)
	}

	ctx.HTML(200, "ktds.html", V)
}

func Post(ctx *gin.Context) {
	V := V{
		ctx.PostForm("kttime"),
		""}

	kttime, _ := time.Parse("2006-01-02T15:04", V.DateTime)
	// cst 转 utc
	kttime = kttime.Add(-time.Hour * 8)

	sinceMinute := int(kttime.Sub(time.Now().UTC()).Minutes())

	V.Message = fmt.Sprintf("设置定时 %d 小时 %d 分钟", sinceMinute/60, sinceMinute%60)
	ctx.HTML(200, "ktds.html", V)
}
