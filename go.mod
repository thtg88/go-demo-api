module goDemoApi

// +heroku goVersion go1.15
go 1.15

require (
	github.com/gin-gonic/gin v1.9.1
	github.com/go-redis/redis/v8 v8.0.0
	github.com/joho/godotenv v1.3.0
	github.com/jordan-wright/email v4.0.1-0.20200917010138-e1c00e156980+incompatible
	github.com/vmihailenco/taskq/v3 v3.0.0
	gorm.io/driver/postgres v1.0.0
	gorm.io/gorm v1.20.1
)
