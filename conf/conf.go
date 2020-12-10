package conf

import (
	"os"
	"qa_go/cache"
	"qa_go/model"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// 全局参数
var (
	JwtSecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))
)

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	_ = godotenv.Load()

	// 设置运行模式
	gin.SetMode(os.Getenv("GIN_MODE"))

	// 启动各种连接单例
	model.Database(os.Getenv("MYSQL_DSN"))
	cache.Redis()
}
