package sessionDatabase

import "time"

func (u *UserInfo) SetCachedTime(t time.Time) {
	u.cachedTime = t
}

func (u *UserInfo) SetCacheTime() {
	u.cachedTime = time.Now()
}

func (u *UserInfo) IsExpired() bool {
	return time.Since(u.cachedTime) > time.Hour*4
}
