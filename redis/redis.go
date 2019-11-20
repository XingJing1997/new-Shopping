package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"strings"
)

var redisPool *redis.Pool

//1.获取redis连接
func GetRedisConn() redis.Conn {
	if redisPool == nil {
		redisPool = &redis.Pool{
			MaxIdle:     10,
			MaxActive:   100,
			IdleTimeout: 10,
			Dial: func() (conn redis.Conn, e error) {
				conn, e = redis.Dial("tcp", "localhost:6379")
				_, e = conn.Do("auth", "abcdefg")
				if e != nil {
					fmt.Println("auth password error:", e)
				}
				return
			},
		}
	}
	return redisPool.Get()

}

//关闭连接
func CloseRedisConn() {
	if redisPool != nil {
		redisPool.Close()
	}
}

//通用redis命令执行方法
func DoRedisCommand(cmd string) (interface{}, error) {
	fmt.Println("DoRedisCommand ...")
	conn := GetRedisConn()

	// get->args= name lisi age 20
	strs := strings.Split(cmd, " ")
	args := make([]interface{}, 0)
	for _, arg := range strs[1:] {
		args = append(args, arg)
	}
	reply, e := conn.Do(strs[0], args...)

	return reply, e
}
