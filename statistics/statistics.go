package statistics

import "time"

type Stat struct {
	SYN                uint32 `json:"syn"`
	ACK                uint32 `json:"ack"`
	FIN                uint32 `json:"fin"`
	RST                uint32 `json:"rst"`
	Packets            uint32 `json:"packets"`
	PacketsOutbound    uint32 `json:"packetsOutbound"`
	Fragmented         uint32 `json:"fragmented"`
	UdpPackets         uint32 `json:"udpPackets"`
	UdpPacketsOutbound uint32 `json:"udpPacketsOutbound"`
}

type RateStat struct {
	Stat
	Snapshot     Stat
	SnapshotTime time.Time
}

type Bandwidth struct {
	Upload   float64 `json:"upload"`
	Download float64 `json:"download"`
}

func SubtractStats(a *Stat, b *Stat) *Stat {
	result := &Stat{
		SYN:                a.SYN - b.SYN,
		ACK:                a.ACK - b.ACK,
		FIN:                a.FIN - b.FIN,
		RST:                a.RST - b.RST,
		Packets:            a.Packets - b.Packets,
		PacketsOutbound:    a.PacketsOutbound - b.PacketsOutbound,
		UdpPackets:         a.UdpPackets - b.UdpPackets,
		UdpPacketsOutbound: a.UdpPacketsOutbound - b.UdpPacketsOutbound,
	}

	return result
}
