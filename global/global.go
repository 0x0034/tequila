package global

import (
	"github.com/0x0034/tequila/config"
	"github.com/0x0034/tequila/utils/timer"
	"github.com/redis/go-redis/v9"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"sync"
)

var (
	GlobalDB     *gorm.DB
	GlobalConfig config.Server
	GlobalViper  *viper.Viper
	GlobalLog    *zap.Logger
	GlobalTimer  = timer.NewTimerTask()
	GlobalRedis  *redis.Client
	BlackCache   local_cache.Cache
	//GlobalDockerClient *docker.Client
	//GlobalK8sClient    *kubernetes.Clientset
	//GlobalK8sConfig    *rest.Config
	lock sync.RWMutex
)
