package wotoConfig

type WotoConfiguration struct {
	Phone       string `json:"phone"`
	AppId       string `json:"app_id"`
	AppHash     string `json:"app_hash"`
	SessionFile string `json:"session_file"`
	// Owners is the IDs of the owners.
	Owners    []int64 `json:"owners"`
	SudoUsers []int64 `json:"sudo_users"`
	// MaxPanic is the maximum amount of allowed panics to be caught.
	// -1 for unlimited.
	MaxPanic int64 `json:"max_panic"`
	// Debug is the debug mode.
	Debug bool `json:"debug"`
	// MaxCacheTime is the max cache time for database.
	MaxCacheTime int64 `json:"max_cache_time"`
	// DbUrl is the of your postgresql database.
	// if `use_sqlite` is set to `true`, this variable will be ignored.
	DbUrl string `json:"db_url"`
	// DbName is the database name. if `use_sqlite` is true, this value is required;
	// in that case, if it's empty, it will be set to `sibyldb` by default.
	DbName string `json:"db_name"`
	// UseSqlite is a bool. set this to `true` if you want to use sqlite database.
	// this is not recommended for production version of Sibyl System.
	UseSqlite bool `json:"use_sqlite"`
	// StatsCacheTime is the amount of stats to be cached in memory in minutes.
	// set it to 0 or comment it out if you don't want it to be cached.
	StatsCacheTime int64 `json:"stats_cache_time"`
	// BotToken is the helper bot's token.
	// it can be commented out (or set to empty) if you don't want
	// the application to interact with telegram directly.
	BotToken string `json:"bot_token"`

	BotAPIUrl string `json:"api_url"`

	DropUpdates      bool     `json:"drop_updates"`
	OrdinaryPrefixes []string `json:"cmd_prefixes"`
	// CmdPrefixes is the command prefixes of the bot.
	CmdPrefixes []rune `json:"-"`
	// BaseChats is the base group's IDs. separate them using " " or ",".
	// values in base chats can be anything: a user's pm, a channel or a group.
	BaseChats []int64 `json:"base_chats"`
}
