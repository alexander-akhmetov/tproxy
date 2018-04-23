package main

import (
	"context"
	"fmt"
	"log"
	"net"

	socks5 "github.com/armon/go-socks5"
)

func filterTelegram() socks5.RuleSet {
	ipList := []string{
		"91.108.4.0/22",
		"91.108.8.0/22",
		"91.108.12.0/22",
		"91.108.16.0/22",
		"91.108.56.0/22",
		"91.108.56.0/23",
		"91.108.56.0/24",
		"149.154.160.0/20",
		"149.154.160.0/22",
		"149.154.164.0/22",
		"149.154.168.0/22",
		"149.154.168.0/23",
		"149.154.170.0/23",
		"2001:67c:4e8::/48",
		"2001:b28:f23d::/48",
		"2001:b28:f23e::/48",
		"2001:b28:f23f::/48",
	}
	var ipNets []*net.IPNet
	for i := 0; i < len(ipList); i++ {
		ip, ipnet, _ := net.ParseCIDR(ipList[i])
		fmt.Println(ipList[i], "-> ip:", ip, " net:", ipnet)
		ipNets = append(ipNets, ipnet)
	}
	return &permitOnlyTelegram{
		ipNets: ipNets,
	}
}

type permitOnlyTelegram struct {
	ipNets []*net.IPNet
}

func (p *permitOnlyTelegram) Allow(ctx context.Context, req *socks5.Request) (context.Context, bool) {
	for _, ipNet := range p.ipNets {
		if ipNet.Contains(req.DestAddr.IP) {
			log.Printf("[INFO] Allow: IP address %s is in the Telegram Networks", req.DestAddr.IP.String())
			return ctx, true
		}
	}
	log.Printf("[WARNING] Deny: IP address %s is not in the Telegram Networks!", req.DestAddr.IP.String())
	return ctx, false
}
