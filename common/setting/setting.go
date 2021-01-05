package setting

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Setting struct {
	viper *viper.Viper
}
/**
 * @summary 通过判断CLI是否传递config路径来判断采用配置信息
 * @param string 配置文件路径
 * @return *Setting error
 */
func NewSetting(configs ....string) (*Setting,error) {
	viper := viper.New()
	viper.SetConfigName("config")
	for _,config := range configs {
		if config != "" {
			viper.AddConfigPath(config)
		}
	}
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	s := &Setting{viper}
	s.WatchConfig()
	return s,err
}
/**
 * @summary 监控配置信息
 * @return void
 */
func (s *Setting) WatchConfig(){
	go func() {
		s.viper.WatchConfig()
		s.viper.OnConfigChange(func(in fsnotify.Event) {
			fmt.Println("配置信息改变了")
			_ = s.ReadAllSection()
		})
	}()
}
