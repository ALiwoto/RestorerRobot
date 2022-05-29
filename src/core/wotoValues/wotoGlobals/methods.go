package wotoGlobals

import (
	"strings"

	"github.com/AnimeKaizoku/ssg/ssg"
)

func (i *PeerInfo) GetRealID() int64 {
	if i.PeerType == PeerTypeChannel {
		idStr := ssg.ToBase10(i.PeerId)
		if strings.HasPrefix(idStr, "-100") {
			idStr = strings.TrimPrefix(idStr, "-100")
			return ssg.ToInt64(idStr)
		}
	}

	return i.PeerId
}
