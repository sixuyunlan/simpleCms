package content

import (
	"log"
	"net/http"
	"simpleCms/app/common"

	"github.com/gin-gonic/gin"
)

func Get(ctx *gin.Context) {

	id := ctx.Param("id")
	row := &Content{}
	if result := common.DB.First(row, id); result.Error != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 2,

			"message": "内容查询失败",
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "",
			"data":    row,
		})
	}
}

func Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if result := common.DB.Delete(&Content{}, id); result.Error != nil {
		log.Println(result.Error)
		ctx.JSON(http.StatusOK, gin.H{
			"code": 2,

			"message": "内容删除失败",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 1,

		"message": "",
	})

}

func Post(ctx *gin.Context) {
	row := &Content{}
	row.Status = "new"
	if result := common.DB.Create(row); result.Error != nil {
		log.Println(result.Error)

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    2,
			"message": "添加失败",
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "",
			"data":    row,
		})
	}

}

func Put(ctx *gin.Context) {
	id := ctx.Param("id")
	body := &Work{}
	if err := ctx.ShouldBind(body); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    2,
			"message": "参数错误",
		})
		return
	}
	row := &Content{}
	common.DB.First(&row, id)

	row.Work = *body
	if result := common.DB.Save(row); result.Error != nil {
		log.Println(result.Error)
		ctx.JSON(http.StatusOK, gin.H{
			"code":    2,
			"message": "内容更新失败",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 1,

		"message": "内容更新",
		"data":    row,
	})
}
