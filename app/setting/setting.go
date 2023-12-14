package setting

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

var Config = new(config)

type config struct {
	System *SystemConfig `mapstructure:"system" json:"system"`
	Mysql  *MysqlConfig  `mapstructure:"mysql" json:"mysql"`
	Jwt    *JwtConfig    `mapstructure:"jwt" json:"jwt"`
	Ldap   *LdapConfig   `mapstructure:"ldap" json:"ldap"`
}

type SystemConfig struct {
	Mode            string `mapstructure:"mode" json:"mode"`
	UrlPathPrefix   string `mapstructure:"url-path-prefix" json:"urlPathPrefix"`
	Port            string `mapstructure:"port" json:"port"`
	InitData        bool   `mapstructure:"init-data" json:"initData"`
	RSAPublicKey    string `mapstructure:"rsa-public-key" json:"rsaPublicKey"`
	RSAPrivateKey   string `mapstructure:"rsa-private-key" json:"rsaPrivateKey"`
	RSAPublicBytes  []byte `mapstructure:"-" json:"-"`
	RSAPrivateBytes []byte `mapstructure:"-" json:"-"`
}

type MysqlConfig struct {
	Username  string `mapstructure:"username" json:"username"`
	Password  string `mapstructure:"password" json:"password"`
	Database  string `mapstructure:"database" json:"database"`
	Host      string `mapstructure:"host" json:"host"`
	Port      int    `mapstructure:"port" json:"port"`
	Query     string `mapstructure:"query" json:"query"`
	LogMode   bool   `mapstructure:"log-mode" json:"logMode"`
	Charset   string `mapstructure:"charset" json:"charset"`
	Collation string `mapstructure:"collation" json:"collation"`
}

type LdapConfig struct {
	Url                string `mapstructure:"url" json:"url"`
	MaxConn            int    `mapstructure:"max-conn" json:"maxConn"`
	BaseDN             string `mapstructure:"base-dn" json:"baseDN"`
	AdminDN            string `mapstructure:"admin-dn" json:"adminDN"`
	AdminPass          string `mapstructure:"admin-pass" json:"adminPass"`
	UserDN             string `mapstructure:"user-dn" json:"userDN"`
	UserInitPassword   string `mapstructure:"user-init-password" json:"userInitPassword"`
	GroupNameModify    bool   `mapstructure:"group-name-modify" json:"groupNameModify"`
	UserNameModify     bool   `mapstructure:"user-name-modify" json:"userNameModify"`
	DefaultEmailSuffix string `mapstructure:"default-email-suffix" json:"defaultEmailSuffix"`
}

type JwtConfig struct {
	Realm      string `mapstructure:"realm" json:"realm"`
	Key        string `mapstructure:"key" json:"key"`
	Timeout    int    `mapstructure:"timeout" json:"timeout"`
	MaxRefresh int    `mapstructure:"max-refresh" json:"maxRefresh"`
}

func InitConfig() {

	dir, err := os.Getwd()
	if err != nil {
		panic(fmt.Errorf("读取应用目录失败:%s", err))
	}
	viper.SetConfigName("env_config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(dir + "/config")
	viper.Set("log.level", "debug") // 设置日志级别为 debug
	err = viper.ReadInConfig()

	// 热更新配置
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// 将读取的配置信息保存至全局变量Conf
		if err := viper.Unmarshal(Config); err != nil {
			panic(fmt.Errorf("初始化配置文件失败:%s", err))
		}
		// 读取rsa key
		//Conf.System.RSAPublicBytes = RSAReadKeyFromFile(Conf.System.RSAPublicKey)
		//Conf.System.RSAPrivateBytes = RSAReadKeyFromFile(Conf.System.RSAPrivateKey)
	})

	if err != nil {
		panic(fmt.Errorf("读取配置文件失败:%s", err))
	}
	// 将读取的配置信息保存至全局变量Conf
	if err := viper.Unmarshal(Config); err != nil {
		panic(fmt.Errorf("初始化配置文件失败:%s", err))
	}
}
