package dfb

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type V struct {
	DateTime string
	Cooktime int
	Message  string
}

func calcSinceMinute(v *V) bool {
	mealtime, _ := time.Parse("2006-01-02T15:04", v.DateTime)
	// cst 转 utc
	mealtime = mealtime.Add(-time.Hour * 8)
	// 间隔分钟
	sinceMinute := int(mealtime.Sub(time.Now().UTC()).Minutes()) - v.Cooktime
	if sinceMinute < 0 {
		v.Message = "煮饭时间不够！"
		return false
	} else {
		v.Message = fmt.Sprintf("需要定时 %d 小时 %d 分钟", sinceMinute/60, sinceMinute%60)
		return true
	}
}

func Get(ctx *gin.Context) {
	// 默认煮饭时间 60 分钟
	V := V{
		Cooktime: 60,
	}

	if ctx.Query("cooktime") != "" {
		V.Cooktime, _ = strconv.Atoi(ctx.Query("cooktime"))
	}

	// 早晚默认吃饭时间
	am := "08:00"
	pm := "18:30"

	if ctx.Query("am") != "" {
		am = ctx.Query("am")
	}

	if ctx.Query("pm") != "" {
		pm = ctx.Query("pm")
	}

	V.DateTime = fmt.Sprintf("%sT%s",
		time.Now().Format("2006-01-02"),
		pm,
	)

	// 间隔时间小于默认晚饭时间，判断是明天早餐
	if !calcSinceMinute(&V) {
		V.DateTime = fmt.Sprintf("%sT%s",
			time.Now().AddDate(0, 0, 1).Format("2006-01-02"),
			am,
		)
	}

	calcSinceMinute(&V)

	ctx.HTML(200, "dfb.html", V)
}

func Post(ctx *gin.Context) {
	V := V{
		DateTime: ctx.PostForm("mealtime"),
	}
	V.Cooktime, _ = strconv.Atoi(ctx.PostForm("cooktime"))
	calcSinceMinute(&V)
	ctx.HTML(200, "dfb.html", V)
}
