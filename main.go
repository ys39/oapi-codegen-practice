/*
* エントリーポイント
 */

package main

import (
	"go-oapi-codegen/api"

	"github.com/gin-gonic/gin"
)

func main() {
	server := api.NewServer()

	r := gin.Default()

	// GinServerOptions で BaseURL を設定
	options := api.GinServerOptions{
		BaseURL: "/api/v1",
	}

	// 生成されたハンドラを Gin に登録 (BaseURL オプションを使用)
	api.RegisterHandlersWithOptions(r, server, options)

	r.Run(":8080")
}
