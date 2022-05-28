package wotoConfig

import (
	"errors"
	"os"
	"strconv"
	"strings"

	ws "github.com/ALiwoto/StrongStringGo/strongStringGo"
	"github.com/bigkevmcd/go-configparser"
)

func LoadConfigFromFile(fileName string) error {
	if WotoConf != nil {
		return nil
	}

	WotoConf = &WotoConfiguration{}
	env := os.Getenv
	configContent, err := configparser.NewConfigParserFromFile(fileName)
	if err != nil {
		return err
	}

	WotoConf.Phone, err = configContent.Get("wotoub", "phone")
	if err != nil || len(WotoConf.Phone) == 0 {
		WotoConf.Phone = env("PHONE")
	}

	if len(WotoConf.Phone) < 7 || WotoConf.Phone[0] != '+' {
		return errors.New("invalid owners phone number")
	}

	WotoConf.AppId, err = configContent.Get("wotoub", "app_id")
	if err != nil || len(WotoConf.AppId) == 0 {
		WotoConf.AppId = env("APP_ID")
	}

	if len(WotoConf.AppId) < 2 {
		return errors.New("invalid app id provided")
	}

	WotoConf.AppHash, err = configContent.Get("wotoub", "app_hash")
	if err != nil || len(WotoConf.AppHash) == 0 {
		WotoConf.AppHash = env("APP_HASH")
	}

	if len(WotoConf.AppHash) < 6 {
		return errors.New("invalid app hash provided")
	}

	WotoConf.SessionFile, err = configContent.Get("wotoub", "session_file")
	if err != nil || len(WotoConf.SessionFile) < 5 {
		WotoConf.SessionFile = env("SESSION_FILE")
	}

	if len(WotoConf.SessionFile) < 5 {
		WotoConf.SessionFile = getDefaultSessionPath()
	}

	ownersStr, err := configContent.Get("wotoub", "owners")
	if err != nil || len(ownersStr) == 0 {
		ownersStr = env("OWNERS")
	}
	WotoConf.Owners = parseBaseStr(strings.TrimSpace(ownersStr))

	WotoConf.MaxPanic, err = configContent.GetInt64("wotoub", "max_panics")
	if err != nil {
		WotoConf.MaxPanic, _ = strconv.ParseInt(env("MAX_PANICS"), 10, 64)
		if WotoConf.MaxPanic == 0 {
			WotoConf.MaxPanic = -1
		}
	}

	WotoConf.Debug, err = configContent.GetBool("wotoub", "debug")
	if err != nil {
		debug := env("WOTO_DEBUG")
		WotoConf.Debug = debug == "yes" || debug == "true"
	}

	// database section variables:
	WotoConf.UseSqlite, err = configContent.GetBool("database", "use_sqlite")
	if err != nil {
		usesqlite := env("USE_SQLITE")
		WotoConf.UseSqlite = usesqlite == "yes" || usesqlite == "true"
	}

	WotoConf.MaxCacheTime, err = configContent.GetInt64("database", "max_cache_time")
	if err != nil {
		WotoConf.MaxCacheTime, _ = strconv.ParseInt(env("MAX_CACHE_TIME"), 10, 64)
	}

	WotoConf.DbUrl, err = configContent.Get("database", "url")
	if err != nil || len(WotoConf.DbUrl) == 0 {
		WotoConf.DbUrl = env("DB_URL")
		if len(WotoConf.DbUrl) == 0 && !WotoConf.UseSqlite {
			return errors.New("no database url is specified")
		}
	}

	WotoConf.DbName, err = configContent.Get("database", "db_name")
	if err != nil || len(WotoConf.DbUrl) == 0 {
		WotoConf.DbName = env("DB_NAME")
		if len(WotoConf.DbName) == 0 {
			WotoConf.DbName = "wotodb"
		}
	}

	WotoConf.StatsCacheTime, err = configContent.GetInt64("database", "max_cache_time")
	if err != nil {
		WotoConf.StatsCacheTime, _ = strconv.ParseInt(env("MAX_CACHE_TIME"), 10, 64)
	}

	// telegram section variables
	WotoConf.BotToken, err = configContent.Get("telegram", "bot_token")
	if err != nil || len(WotoConf.BotToken) == 0 {
		WotoConf.BotToken = env("BOT_TOKEN")
	}

	WotoConf.BotAPIUrl, err = configContent.Get("telegram", "api_url")
	if err != nil || len(WotoConf.BotAPIUrl) == 0 {
		WotoConf.BotAPIUrl = env("API_URL")
	}

	// database section variables:
	WotoConf.DropUpdates, err = configContent.GetBool("telegram", "drop_updates")
	if err != nil {
		dropUpdates := env("DROP_UPDATES")
		WotoConf.DropUpdates = dropUpdates == "yes" || dropUpdates == "true"
	}

	baseStr, err := configContent.Get("telegram", "base_chats")
	if err != nil || len(baseStr) == 0 {
		baseStr = env("BASE_CHATS")
	}
	WotoConf.BaseChats = parseBaseStr(strings.TrimSpace(baseStr))

	preStr, err := configContent.Get("telegram", "cmd_prefixes")
	if err != nil || len(preStr) == 0 {
		preStr = env("CMD_PREFIXES")
	}
	WotoConf.CmdPrefixes = parseCmdPrefixes(preStr)

	return nil
}

func LoadConfig() error {
	return LoadConfigFromFile("config.ini")
}

func PrepareVariables() error {
	if WotoConf == nil {
		return errors.New("woto configuration is not loaded")
	}

	err := os.Setenv("APP_ID", WotoConf.AppId)
	if err != nil {
		return err
	}

	err = os.Setenv("APP_HASH", WotoConf.AppHash)
	if err != nil {
		return err
	}

	err = os.Setenv("SESSION_FILE", WotoConf.SessionFile)
	if err != nil {
		return err
	}

	return nil
}

func getDefaultSessionPath() string {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}

	if dir[len(dir)-1] == '/' {
		dir = dir[:len(dir)-1]
	}

	return dir + "/session/session.wotoub.json"
}

func parseCmdPrefixes(value string) []rune {
	if len(value) == 0 {
		return []rune{'!', '.'}
	}

	value = strings.TrimSpace(value)
	if strings.Contains(value, " ") {
		var all []rune
		// comment for future cases: we can use `ws.Split` function as well,
		// but since we absolutely need to split the commands by white spaces,
		// it's better to use `strings.Fields` function from stdlib.
		myStrs := ws.FixSplitWhite(strings.Fields(value))
		for _, str := range myStrs {
			all = append(all, rune(str[0]))
		}
		return all
	} else {
		if len(value) > 0 {
			return []rune(value)
		}
		return nil
	}
}

func parseBaseStr(value string) []int64 {
	if !strings.Contains(value, " ") && !strings.Contains(value, ",") {
		value = strings.TrimSpace(value)
		tmp, err := strconv.ParseInt(value, 10, 64)
		if err != nil || tmp == 0 {
			return nil
		}
		return []int64{tmp}
	}

	myStrs := ws.Split(value, " ", ",")
	if len(myStrs) == 0 {
		return nil
	}

	var tmp int64
	var err error
	var all []int64
	for _, str := range myStrs {
		tmp, err = strconv.ParseInt(str, 10, 64)
		if err != nil || tmp == 0 {
			continue
		}
		all = append(all, tmp)
	}

	return all
}

func IsDebug() bool {
	if WotoConf != nil {
		return WotoConf.Debug
	}
	return true
}

func UseSqlite() bool {
	if WotoConf != nil {
		return WotoConf.UseSqlite
	}
	return true
}

func GetPrefixes() []rune {
	if WotoConf != nil {
		return WotoConf.CmdPrefixes
	}
	return []rune{'!', '>'}
}
