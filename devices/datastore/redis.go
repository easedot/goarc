package datastore
import(
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
	//"github.com/heroku/x/hredis/redigo"

	"github.com/easedot/goarc/config"
)

func NewRedis() (*redis.Pool, error){
	//redis://[:password@]host[:port][/db-number][?option=value]
	url := fmt.Sprintf(
		"redis://%s@%s:%s/%d",
		config.C.Database.Password,
		config.C.Redis.Host, config.C.Redis.Port,
		config.C.Redis.Db,
	)

	dialFunc:=func() (redis.Conn, error){
		r,c:=redis.DialURL(url)
		return r,c
	}

	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial:        dialFunc,
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}, nil
}
