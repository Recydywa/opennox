package discover

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"path/filepath"
)

const (
	StaticFile = "game_ip.txt"
)

func staticIPs(path string) ([]Server, error) {
	name := filepath.Base(path)
	f, err := os.Open(path)
	if os.IsNotExist(err) {
		Log.Printf("no %s file", name)
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot read %s: %w", name, err)
	}
	defer f.Close()

	var (
		out  []Server
		last error
	)
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
		ip := net.ParseIP(line).To4()
		if ip == nil {
			last = fmt.Errorf("cannot parse server IP in %s: %q", name, line)
			Log.Println(last)
			continue
		}
		Log.Printf("%s: %v", name, ip)
		out = append(out, Server{Addr: ip})
	}
	if err := sc.Err(); err != nil {
		last = fmt.Errorf("error reading %s: %w", name, err)
		Log.Println(last)
	}
	return out, last
}
