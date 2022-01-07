package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"inherited/internal/models"
	"strconv"
)

func ToExcel(form []models.Print, ctx *gin.Context) {
	categories := map[string]string{
		"A1": "序号",
		"B1": "工号",
		"C1": "姓名",
		"D1": "方案",
		"E1": "总额",
	}

	values := map[int]map[string]string{}

	for i, v := range form {
		fmt.Println(i+1, v)
		values[i+1] = map[string]string{
			"A" + strconv.Itoa(i+2): strconv.FormatUint(uint64(v.ID), 10),
			"B" + strconv.Itoa(i+2): strconv.Itoa(v.Num),
			"C" + strconv.Itoa(i+2): v.UserName,
			"D" + strconv.Itoa(i+2): v.Scheme,
			"E" + strconv.Itoa(i+2): strconv.Itoa(v.Money),
		}

	}
	f := excelize.NewFile()
	for k, v := range categories {
		err := f.SetCellValue("Sheet1", k, v)
		if err != nil {
			return
		}
	}
	for _, v := range values {
		for k1, v1 := range v {
			err := f.SetCellValue("Sheet1", k1, v1)
			if err != nil {
				return
			}
		}

	}
	ctx.Header("Content-Type", "application/octet-stream")
	ctx.Header("Content-Disposition", "attachment; 湖南工学院教职工年终奖一览表.xlsx")
	ctx.Header("Content-Transfer-Encoding", "binary")
	_ = f.Write(ctx.Writer)
	err := f.SaveAs("testoo1.xlsx")
	if err != nil {
		panic(err)
	}

}
