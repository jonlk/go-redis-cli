package cache

func CreateClient() *redisClient {
	return newRedisClient()
}
