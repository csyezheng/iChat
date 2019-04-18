package common

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"time"
)

type Config interface {
	Debug() bool
	Db() *gorm.DB

	MigrateDb()

	ConfigFile() string

	AppName() string
	AppVersion() string
	AppCopyright() string

	SqlServerHost() string
	SqlServerPort() uint
	SqlServerPath() string
	SqlServerPassword() string

	HttpServerHost() string
	HttpServerPort() int
	HttpServerMode() string
	HttpServerPassword() string
	HttpTemplatesPath() string
	HttpFaviconPath() string
	HttpPublicPath() string
	HttpPublicBuildPath() string

	DatabaseDriver() string
	DatabaseDsn() string

	AssetsPath() string
	ServerPath() string
	CachePath() string

	NewClientStore(lifetime time.Duration)
	ClientStore() *ClientStore
	NewHub() *Hub
	GetHub() *Hub
}
