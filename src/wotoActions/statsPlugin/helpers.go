package statsPlugin

import (
	"context"
	"strconv"
	"time"

	"github.com/AnimeKaizoku/RestorerRobot/src/core/wotoEntry/entryManager"
	"github.com/AnimeKaizoku/RestorerRobot/src/core/wotoStyle"
	"github.com/AnimeKaizoku/ssg/ssg"
	"github.com/AnimeKaizoku/ssg/ssg/shellUtils"
)

func fetchGitStats(md wotoStyle.WStyle) {
	rawGit, ok := getRawGit(md)
	if !ok {
		return
	}

	if len(rawGit) == 0 {
		// try again; in some situations, when we recently have
		// pushed to HEAD, the git command may not be able to
		// find the HEAD commit.
		rawGit, ok = getRawGit(md)
		if len(rawGit) == 0 || !ok {
			// give up and return :(
			return
		}
	}

	allRaws := ssg.Split(rawGit, "\n")
	if len(allRaws) < 3 {
		return
	}
	shortGit := allRaws[0]
	longGit := allRaws[1]
	gitVs := ssg.Split(allRaws[2], " ", "\t")
	if len(gitVs) != 2 {
		return
	}
	upstream, err := strconv.Atoi(gitVs[0])
	if err != nil {
		return
	}
	local, err := strconv.Atoi(gitVs[1])
	if err != nil {
		return
	}
	vsInt := upstream - local
	commitUrl := gitBaseUrl + "/commit/" + longGit

	md.Normal("ℹ️ ").Link("Git ", gitBaseUrl)
	md.Bold("Status:")
	md.Normal("\n• Current commit: ").Link(shortGit, commitUrl)
	md.Normal("\n• Running behind by ").Mono(strconv.Itoa(vsInt))
	md.Normal(" commits\n\n")
}

func getRawGit(md wotoStyle.WStyle) (string, bool) {
	strChan := make(chan string)
	timeOutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go func() {
		strChan <- shellUtils.GetGitStatsString()
	}()

	var rawGit string
	select {
	case <-timeOutCtx.Done():
		md.Bold("\n • Git stats: ").Mono("Timed out")
		return "", false
	case v := <-strChan:
		rawGit = v
		return rawGit, true
	}
}

func LoadAllHandlers(manager *entryManager.EntryManager) {
	manager.AddHandlers(statusCmd, statusHandler)
}
