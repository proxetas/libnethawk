package api

import "time"

type API struct {
	BlockIP      func(ip string)
	UnblockIP    func(ip string)
	QuarantineIP func(ip string, duration time.Duration)
	ReleaseIP    func(ip string)
	SendMessage  func(extension string, message string)
	GetPeerAlias func(ip string) string
	GetPeerMeta  func(ip string) string
	SetKnownPeer func(ip string, alias string, meta string)
	Lockdown     func()
	EndLockdown  func()
}
