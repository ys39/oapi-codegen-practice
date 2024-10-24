package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct{}

func NewServer() Server {
	return Server{}
}

func (Server) GetDetailId(ctx *gin.Context, id int) {
	// PostsからIDが一致する投稿を取得
	for _, post := range Posts {
		if post.ID == id {
			ctx.JSON(http.StatusOK, post)
			return
		}
	}

	// IDが一致する投稿が見つからなかった場合
	ctx.JSON(http.StatusNotFound, gin.H{
		"message": "Not Found",
	})
}

func (Server) GetList(ctx *gin.Context, params GetListParams) {
	// Postsを返す
	ctx.JSON(http.StatusOK, Posts)
}

func (Server) PostCreate(ctx *gin.Context) {
	// リクエストボディを取得
	var req PostCreateJSONRequestBody
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}

	// 新しい投稿を作成
	var newPost Post
	newPost.ID = len(Posts) + 1
	newPost.Title = req.Title
	newPost.Content = req.Content
	Posts = append(Posts, newPost)

	ctx.JSON(200, newPost)
}

func (Server) DeleteDeleteId(ctx *gin.Context, id int) {
	// IDが一致する投稿を削除
	for i, post := range Posts {
		if post.ID == id {
			Posts = append(Posts[:i], Posts[i+1:]...)
			ctx.JSON(200, gin.H{
				"message": "Post deleted",
			})
			return
		}
	}

	// IDが一致する投稿が見つからなかった場合
	ctx.JSON(http.StatusNotFound, gin.H{
		"message": "Not Found",
	})
}

func (Server) PutUpdateId(ctx *gin.Context, id int) {
	// リクエストボディを取得
	var req PutUpdateIdJSONRequestBody
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}

	// IDが一致する投稿を更新
	for i, post := range Posts {
		if post.ID == id {
			Posts[i].Title = req.Title
			Posts[i].Content = req.Content
			ctx.JSON(200, Posts[i])
			return
		}
	}

	// IDが一致する投稿が見つからなかった場合
	ctx.JSON(http.StatusNotFound, gin.H{
		"message": "Not Found",
	})
}
