package gutils

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"strings"
)

// from https://stackoverflow.com/a/37897238/3337885
func RealAddr(r *http.Request) string {
	remoteIP := ""
	// the default is the originating ip. but we try to find better options because this is almost
	// never the right IP
	if parts := strings.Split(r.RemoteAddr, ":"); len(parts) == 2 {
		remoteIP = parts[0]
	}
	// If we have a forwarded-for header, take the address from there
	if xff := strings.Trim(r.Header.Get("X-Forwarded-For"), ","); len(xff) > 0 {
		addrs := strings.Split(xff, ",")
		lastFwd := addrs[len(addrs)-1]
		if ip := net.ParseIP(lastFwd); ip != nil {
			remoteIP = ip.String()
		}
		// parse X-Real-Ip header
	} else if xri := r.Header.Get("X-Real-Ip"); len(xri) > 0 {
		if ip := net.ParseIP(xri); ip != nil {
			remoteIP = ip.String()
		}
	}

	return remoteIP
}

func ReverseDNS(ip string) string {
	host, err := (&net.Resolver{}).LookupAddr(context.Background(), ip)
	if err != nil || len(host) <= 0 {
		return "N/A"
	}
	return host[0]
}

// ExtractPort takes a "IP:Port" formatted string and returns the port as integer.
// If extraction fails the function returns 0.
func ExtractPort(ipv4Addr string) int {
	_, port, err := net.SplitHostPort(ipv4Addr)
	if err != nil {
		return 0
	}
	p, err := strconv.Atoi(port)
	if err != nil {
		p = 0
	}
	return p
}

// ExtractHost takes a "IP:Port" formatted string and returns the ip portion.
// If extraction fails the function returns an empty string.
func ExtractHost(ipv4Addr string) string {
	ip, _, err := net.SplitHostPort(ipv4Addr)
	if err != nil {
		return ""
	}
	return ip
}

// SplitHostPort takes a "IP:Port" formatted string and returns the port as integer and the ip portion as string.
// If splitting fails the function returns an empty string as IP and 0 as port.
func SplitHostPort(ipv4Addr string) (string, int) {
	ip, port, err := net.SplitHostPort(ipv4Addr)
	if err != nil {
		return "", 0
	}
	p, err := strconv.Atoi(port)
	if err != nil {
		p = 0
	}
	return ip, p
}

func ExtractPortFromAddr(addr net.Addr) int {
	return ExtractPort(addr.String())
}

func ExtractHostFromAddr(addr net.Addr) string {
	return ExtractHost(addr.String())
}

func SplitHostPortFromAddr(addr net.Addr) (string, int) {
	return SplitHostPort(addr.String())
}

func JoinHostPort(host string, port int) string {
	return fmt.Sprintf("%s:%d", host, port)
}
