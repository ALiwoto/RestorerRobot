package wotoConfig

type MainConfigSection struct {
	AppId             string  `key:"app_id"`
	AppHash           string  `key:"app_hash"`
	BotToken          string  `key:"bot_token"`
	BackupsBaseDir    string  `key:"backups_base_directory_path"`
	PgDumpCommand     string  `key:"pg_dump_command"`
	GlobalLogChannels []int64 `key:"global_log_channels"`
	Owners            []int64 `key:"owners"`
	CmdPrefixes       []rune  `key:"cmd_prefixes" type:"[]rune"`
	SessionFile       string  `key:"session_file"`
	Debug             bool    `key:"debug"`
}

type ValueSection struct {
	DbUrl       string  `key:"db_url"`
	LogChannels []int64 `key:"log_channels"`
	BackupType  string  `key:"backup_type" default:"sql"`
	sectionName string
}
