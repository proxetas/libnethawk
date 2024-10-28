package api

import "time"

type API struct {
	BlockIP      func(ip string)
	UnblockIP    func(ip string)
	QuarantineIP func(ip string, duration time.Duration)
	ReleaseIP    func(ip string)
	SendMessage  func(extension string, message string)
}
