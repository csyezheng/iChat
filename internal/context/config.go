package context

import (
	"github.com/csyezheng/iChat/internal/common"
	"github.com/csyezheng/iChat/internal/models"
	"github.com/csyezheng/iChat/internal/utils"
	"github.com/jinzhu/gorm"
	"github.com/urfave/cli"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kylelemons/go-gypsy/yaml"
)

var DB *gorm.DB

// Config provides a struct in which application configuration is stored.
// In order to build an abstraction, we must use it's method to get the corresponding field value.
type Config struct {
	appName            string
	appVersion         string
	appCopyright       string
	debug              bool
	configFile         string
	assetsPath         string
	cachePath          string
	originalsPath      string
	importPath         string
	exportPath         string
	sqlServerHost      string
	sqlServerPort      uint
	sqlServerPath      string
	sqlServerPassword  string
	httpServerHost     string
	httpServerPort     int
	httpServerMode     string
	httpServerPassword string
	databaseDriver     string
	databaseDsn        string
	db                 *gorm.DB
	hub                *common.Hub
	cliStore           *common.ClientStore
}

// NewConfig() creates a new configuration entity by using two methods:
// 1. SetValuesFromFile: This will initialize values from a yaml config file.
// 2. SetValuesFromCliContext: Which comes after SetValuesFromFile and overrides
//    any previous values giving an option two override file configs through the CLI.
// TODO Add SetValuesFromCliContext method
func NewConfig(ctx *cli.Context) *Config {
	c := &Config{}
	c.appName = ctx.App.Name
	c.appCopyright = ctx.App.Copyright
	c.appVersion = ctx.App.Version
	c.SetValuesFromFile(utils.ExpandedFilename(ctx.GlobalString("config-file")))

	return c
}

// SetValuesFromFile uses a yaml config file to initiate the configuration entity.
func (c *Config) SetValuesFromFile(fileName string) error {
	yamlConfig, err := yaml.ReadFile(fileName)

	if err != nil {
		return err
	}

	c.configFile = fileName
	if debug, err := yamlConfig.GetBool("debug"); err == nil {
		c.debug = debug
	}

	if sqlServerHost, err := yamlConfig.Get("sql-host"); err == nil {
		c.sqlServerHost = sqlServerHost
	}

	if sqlServerPort, err := yamlConfig.GetInt("sql-port"); err == nil {
		c.sqlServerPort = uint(sqlServerPort)
	}

	if sqlServerPassword, err := yamlConfig.Get("sql-password"); err == nil {
		c.sqlServerPassword = sqlServerPassword
	}

	if httpServerHost, err := yamlConfig.Get("http-host"); err == nil {
		c.httpServerHost = httpServerHost
	}

	if httpServerPort, err := yamlConfig.GetInt("http-port"); err == nil {
		c.httpServerPort = int(httpServerPort)
	}

	if httpServerMode, err := yamlConfig.Get("http-mode"); err == nil {
		c.httpServerMode = httpServerMode
	}

	if databaseDriver, err := yamlConfig.Get("database-driver"); err == nil {
		c.databaseDriver = databaseDriver
	}

	if databaseDsn, err := yamlConfig.Get("database-dsn"); err == nil {
		c.databaseDsn = databaseDsn
	}

	if assetsPath, err := yamlConfig.Get("assets-path"); err == nil {
		c.assetsPath = utils.ExpandedFilename(assetsPath)
	}

	return nil
}

// connectToDatabase establishes a database connection.
// It tries to do this 12 times with a 5 second sleep interval in between.
// TODO Add multiple database driver support
func (c *Config) connectToDatabase() error {

	dbDsn := c.DatabaseDsn()
	db, err := gorm.Open("mysql", dbDsn)

	if err != nil || db == nil {

		for i := 1; i <= 12; i++ {
			time.Sleep(5 * time.Second)

			db, err = gorm.Open("mysql", dbDsn)

			if db != nil && err == nil {
				break
			}
		}

		if err != nil || db == nil {
			log.Fatal(err)
		}
	}

	c.db = db

	return err
}

// AppName returns the application name.
func (c *Config) AppName() string {
	return c.appName
}

// AppVersion returns the application version.
func (c *Config) AppVersion() string {
	return c.appVersion
}

// AppCopyright returns the application copyright.
func (c *Config) AppCopyright() string {
	return c.appCopyright
}

// Debug returns true if debug mode is on.
func (c *Config) Debug() bool {
	return c.debug
}

// ConfigFile returns the config file name.
func (c *Config) ConfigFile() string {
	return c.configFile
}

// SqlServerHost returns the built-in SQL server host name or IP address (empty for all interfaces).
func (c *Config) SqlServerHost() string {
	return c.sqlServerHost
}

// SqlServerPort returns the built-in SQL server port.
func (c *Config) SqlServerPort() uint {
	return c.sqlServerPort
}

// SqlServerPath returns the database storage path for TiDB.
func (c *Config) SqlServerPath() string {
	if c.sqlServerPath != "" {
		return c.sqlServerPath
	}

	return c.ServerPath() + "/database"
}

// SqlServerPassword returns the password for the built-in database server.
func (c *Config) SqlServerPassword() string {
	return c.sqlServerPassword
}

// HttpServerHost returns the built-in HTTP server host name or IP address (empty for all interfaces).
func (c *Config) HttpServerHost() string {
	return c.httpServerHost
}

// HttpServerPort returns the built-in HTTP server port.
func (c *Config) HttpServerPort() int {
	return c.httpServerPort
}

// HttpServerMode returns the server mode.
func (c *Config) HttpServerMode() string {
	return c.httpServerMode
}

// HttpServerPassword returns the password for the user interface (optional).
func (c *Config) HttpServerPassword() string {
	return c.httpServerPassword
}

// CachePath returns the path to the cache.
func (c *Config) CachePath() string {
	return c.cachePath
}

// AssetsPath returns the path to the assets.
func (c *Config) AssetsPath() string {
	return c.assetsPath
}

// DatabaseDriver returns the database driver name.
func (c *Config) DatabaseDriver() string {
	return c.databaseDriver
}

// DatabaseDsn returns the database data source name (DSN).
func (c *Config) DatabaseDsn() string {
	return c.databaseDsn
}

// DatabasePath returns the database storage path (e.g. for SQLite or Bleve).
func (c *Config) DatabasePath() string {
	return c.AssetsPath() + "/database"
}

// ServerPath returns the server assets path (public files, favicons, templates,...).
func (c *Config) ServerPath() string {
	return c.AssetsPath() + "/server"
}

// HttpTemplatesPath returns the server templates path.
func (c *Config) HttpTemplatesPath() string {
	return c.ServerPath() + "/templates"
}

// HttpFaviconsPath returns the favicons path.
func (c *Config) HttpFaviconPath() string {
	return c.HttpPublicPath() + "/favicons"
}

// HttpPublicPath returns the public server path (//server/assets/*).
func (c *Config) HttpPublicPath() string {
	return c.ServerPath() + "/public"
}

// HttpPublicBuildPath returns the public build path (//server/assets/build/*).
func (c *Config) HttpPublicBuildPath() string {
	return c.HttpPublicPath() + "/build"
}

// Db returns the db connection.
func (c *Config) Db() *gorm.DB {
	if c.db == nil {
		c.connectToDatabase()
	}

	return c.db
}

// MigrateDb will start a migration process.
func (c *Config) MigrateDb() {
	db := c.Db()
	DB = c.Db()

	db.AutoMigrate(&models.CoreUser{})
}

// Create new client store
func (c *Config) NewClientStore(lifetime time.Duration) {
	c.cliStore = common.NewClientStore(lifetime)
}

// Returns client store
func (c *Config) ClientStore() *common.ClientStore {
	return c.cliStore
}

// Create a new hub
func (c *Config) NewHub() *common.Hub {
	c.hub = common.NewHub()
	return c.hub
}

// Get hub
func (c *Config) GetHub() *common.Hub {
	return c.hub
}
