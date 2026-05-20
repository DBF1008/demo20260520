package consts

const (

	CachePrefix = "APP:"

	CacheModelMem   = "memory"
	CacheModelRedis = "redis"
	CacheModelDist  = "dist"


	CacheSysDict = CachePrefix + "sysDict"


	CacheSysDictTag = CachePrefix + "sysDictTag"

	CacheSysConfigTag = CachePrefix + "sysConfigTag"
)
