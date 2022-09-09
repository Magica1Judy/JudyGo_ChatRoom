package main
import(
	"github.com/gomodule/redigo/redis"
	"time"
	
)

var pool *redis.Pool

func initPool(address string, maxidle int, mexactive int, idletimeout time.Duration)(){
	pool = &redis.Pool{
		MaxIdle: maxidle, // 最大空闲链接数
		MaxActive: mexactive, // 和数据库最大的链接数 0是无限制
		IdleTimeout: idletimeout, // 最大空闲时间 100s没用这个链接后 放入连接池
		Dial: func()(redis.Conn,error){ // 初始化链接的代码，链接哪个ip的redis
			return redis.Dial("tcp",address)
		},
	}
}