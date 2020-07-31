package db

import (
	"github.com/go-redis/redis"
	"os"
	"time"
)

type RedisPool struct {
	*redis.Client
}

var (
	//RDB *redis.Client
	err   error
	RPool RedisPool
)

// 被db/init.go引用
func NewRedisPool() {
	rb := redis.NewClient(&redis.Options{
		Addr:         os.Getenv("REDIS_ADDR"),
		Password:     "",
		DB:           0,
		DialTimeout:  0,
		ReadTimeout:  time.Second,
		WriteTimeout: 0,
		PoolSize:     0,
		MinIdleConns: 10,
		MaxConnAge:   0,
	})

	_, err = rb.Ping().Result()
	if err != nil {
		panic(err)
	}
	//RDB = rb
	RPool.Client = rb
}

// cache functions...
//  string  k:v
func (r RedisPool) SetKeyValue(k string, v interface{}, exTime time.Duration) (bool, error) {
	return r.SetNX(k, v, exTime).Result()
}

func (r RedisPool) GetKeyValue(k string) (string, error) {
	return r.Get(k).Result()
}

func (r RedisPool) DelKeyValue(k string) (int64, error) {
	return r.Del(k).Result()
}

// set   sadd
/*
	return: 1->success
*/
func (r RedisPool) SetList(k string, v interface{}) (int64, error) {
	return r.SAdd(k, v).Result()
}

// 判断成员 v 是否在 k队列中
func (r RedisPool) IsKeyInList(k string, v interface{}) (bool, error) {
	return r.SIsMember(k, v).Result()
}

func (r RedisPool) DelList(k string, v interface{}) (int64, error) {
	return r.SRem(k, v).Result()
}

// zset   zadd
func (r RedisPool) SetOrderList(k string, Members ...interface{}) (int64, error) {
	Maps := make([]redis.Z,1)
	for _,value := range Members{
		m := redis.Z{
			Score:  1,
			Member: value,
		}
		Maps = append(Maps,m)
	}
	//m := redis.Z{
	//	Score:  1,
	//	Member: Member,
	//}
	return r.ZAdd(k, Maps...).Result()
}

// 获取分数值
func (r RedisPool) GetOrderListScore(k string, Member string) (float64, error) {
	return r.ZScore(k, Member).Result()
}

// 添加元素menber的score值   返回Member值当前分数值
// 通过ZINCRBY命令递增来修改之前的值，相应的他们的排序位置也会随着分数变化而改变
func (r RedisPool) AddOrderIncre(k string, score float64, Member string) (float64, error) {
	return r.ZIncrBy(k, score, Member).Result()
}

// 删除元素    返回 0 成员不存在, >0 表示删除删除元素个数
func (r RedisPool) DelOrderIncre(k string, Member ...interface{}) (int64, error) {
	return r.ZRem(k, Member...).Result()
}

// 获取元素总个数
func (r RedisPool) GetOrderCount(k string) (int64, error) {
	return r.ZCard(k).Result()
}

// 返回成员排名 排名从0开始,所以+1
func (r RedisPool) GetMemberOrder(k string, Member string) (ranking int64, err error) {
	ranking, err = r.ZRevRank(k, Member).Result()
	return ranking + 1, err
}

// 根据一定范围返回数据,从大到小
func (r RedisPool) GetSpaceOrder(k string, page, number int64) ([]string, error) {
	start := int64(page * number)
	stop := int64(start + number)
	return r.ZRevRange(k, start, stop).Result()
}
