package config

import (
	"dify_feishu_mcp/logger"
	"dify_feishu_mcp/model"

	"github.com/spf13/viper"
)

type ConfigManager struct {
	Config *model.Config
}

func NewConfigManager() *ConfigManager {
	return &ConfigManager{}
}
func (cm *ConfigManager) LoadConfig() error {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	err := v.ReadInConfig()
	if err != nil {
		logger.Logger.Fatalf("读取配置文件失败: %v", err)
	}

	cm.Config = &model.Config{
		FeiShuAppID:     v.GetString("feishu_app_id"),
		FeiShuAppSecret: v.GetString("feishu_app_secret"),
		DifyAPIKey:      v.GetString("dify_api_key"),
		DifyBaseURL:     v.GetString("dify_base_url"),
		RedisAddr:       v.GetString("redis_addr"),
		RedisPass:       v.GetString("redis_pass"),
		RedisDB:         v.GetInt("redis_db"),
	}

	return nil
}
func (cm *ConfigManager) Start() {
	if cm.Config.FeiShuAppID == "" || cm.Config.FeiShuAppSecret == "" {
		logger.Logger.Fatalf("请设置APP_ID和APP_SECRET环境变量")
	}
	if cm.Config.DifyAPIKey == "" || cm.Config.DifyBaseURL == "" {
		logger.Logger.Fatalf("请设置DIFY_API_KEY和DIFY_BASE_URL环境变量")
	}
	if cm.Config.RedisAddr == "" || cm.Config.RedisPass == "" || cm.Config.RedisDB == 0 {
		logger.Logger.Fatalf("请设置Redis_Addr，Redis_Pass和Redis_DB环境变量")
	}
}
