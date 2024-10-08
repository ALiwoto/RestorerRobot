package wotoConfig

import (
	"os"
	"time"

	"github.com/ALiwoto/ssg/ssg/strongParser"
)

func LoadConfigFromFile(fileName string) error {
	if WotoConf != nil {
		return nil
	}

	opts := &strongParser.ConfigParserOptions{
		ReadEnv:         true,
		MainSectionName: "main",
	}

	config, err := strongParser.ParseMainAndArrays[MainConfigSection, ValueSection](fileName, opts)
	if err != nil {
		return err
	}

	WotoConf = config
	PrepareDirectories()
	return nil
}

func LoadConfig() error {
	return LoadConfigFromFile("config.ini")
}

func PrepareDirectories() {
	makeSureDirExists(WotoConf.Main.BackupsBaseDir)
	for _, current := range WotoConf.Sections {
		makeSureDirExists(WotoConf.Main.BackupsBaseDir + "/" + current.GetSectionName())
	}
}

func makeSureDirExists(dirName string) {
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		os.MkdirAll(dirName, 0755)
	}
}

func getDefaultSessionPath() string {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}

	if dir[len(dir)-1] == '/' {
		dir = dir[:len(dir)-1]
	}

	se := string(os.PathSeparator)

	return dir + se + "session" + se + "session.restorerbot.json"
}

func IsDebug() bool {
	return WotoConf.Main.Debug
}

func GetSectionValueByName(name string) *ValueSection {
	for _, v := range WotoConf.Sections {
		if name == v.GetSectionName() {
			return v
		}
	}

	return nil
}

func GetBaseDirForBackup(value string) string {
	p := string(os.PathSeparator)
	baseDir := WotoConf.Main.BackupsBaseDir + p + value
	makeSureDirExists(baseDir)
	return baseDir + p
}

func IsOwner(id int64) bool {
	for _, v := range WotoConf.Main.Owners {
		if v == id {
			return true
		}
	}
	return false
}

func GetAppId() int {
	return WotoConf.Main.AppId
}

func GetBotToken() string {
	return WotoConf.Main.BotToken
}

func GetAppHash() string {
	return WotoConf.Main.AppHash
}

func GetProxy() string {
	return WotoConf.Main.Proxy
}

func GetDbPath() string {
	if WotoConf.Main.DatabasePath == "" {
		return "session/database.session"
	}
	return WotoConf.Main.DatabasePath
}

func GetScheduleManagerInterval() time.Duration {
	if WotoConf.Main.ScheduleManagerInterval != 0 {
		return time.Duration(WotoConf.Main.ScheduleManagerInterval) * time.Hour
	}
	return 10 * time.Hour
}

func GetDatabasesConfigs() []*ValueSection {
	return WotoConf.Sections
}

func GetGlobalLogChannels() []int64 {
	return WotoConf.Main.GlobalLogChannels
}

func GetPgDumpCommand() string {
	if WotoConf.Main.PgDumpCommand == "" {
		return "pg_dump"
	}

	return WotoConf.Main.PgDumpCommand
}

func GetSessionPath() string {
	if len(WotoConf.Main.SessionFile) > 5 {
		return WotoConf.Main.SessionFile
	}
	return getDefaultSessionPath()
}

func GetPrefixes() []rune {
	if WotoConf != nil && len(WotoConf.Main.CmdPrefixes) != 0 {
		return WotoConf.Main.CmdPrefixes
	}
	return []rune{'!', '/'}
}
