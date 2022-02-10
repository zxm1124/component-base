package viper

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Setting struct {
	viper *viper.Viper
}

func NewSetting() (*Setting, error) {
	vp := viper.New()
	// 配置文件名
	vp.SetConfigName("config")
	// 配置路径
	vp.AddConfigPath("configs")
	vp.AddConfigPath("./configs")
	vp.AddConfigPath("../configs")
	// 配置后缀
	vp.SetConfigType("yaml")
	// 读取配置文件
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &Setting{
		viper: vp,
	}, nil
}

func (s Setting) ReadSection(configName string, v interface{}) error {
	err := s.viper.UnmarshalKey(configName, v)
	if err != nil {
		return err
	}

	return nil
}

// SetupSetting 读取配置文件信息
func SetupSetting(settingName string, setObj interface{}) error {
	setting, err := NewSetting()
	if err != nil {
		logrus.Panicf("setupSetting.NewSetting code:%v", err)
		return err
	}

	// 读取APP配置文件信息
	err = setting.ReadSection(settingName, &setObj)
	if err != nil {
		logrus.Panicf("setupSetting.ReadSection code: %v", err)
		return err
	}

	return nil
}
