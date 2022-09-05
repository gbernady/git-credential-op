package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/gbernady/go-op"
)

type attributes struct {
	Protocol string
	Host     string
	Path     string
	Username string
	Password string
	URL      string
}

func (a *attributes) Parse(r io.Reader) {
	s := bufio.NewScanner(r)
	for s.Scan() {
		v := strings.Split(s.Text(), "=")
		if len(v) != 2 {
			continue
		}
		key := v[0]
		val := strings.TrimSuffix(v[1], "\n")
		switch key {
		case "protocol":
			a.Protocol = val
		case "host":
			a.Host = val
		case "path":
			a.Path = val
		case "username":
			a.Username = val
		case "password":
			a.Password = val
		case "url":
			a.URL = val
		default:
			continue
		}
	}
}

func (a attributes) String() string {
	var b strings.Builder
	if a.Protocol != "" {
		fmt.Fprintf(&b, "protocol=%s\n", a.Protocol)
	}
	if a.Host != "" {
		fmt.Fprintf(&b, "host=%s\n", a.Host)
	}
	if a.Path != "" {
		fmt.Fprintf(&b, "path=%s\n", a.Path)
	}
	if a.Username != "" {
		fmt.Fprintf(&b, "username=%s\n", a.Username)
	}
	if a.Password != "" {
		fmt.Fprintf(&b, "password=%s\n", a.Password)
	}
	if a.URL != "" {
		fmt.Fprintf(&b, "url=%s\n", a.URL)
	}
	return b.String()
}

func (a attributes) Match(item *op.Item) bool {
	// required
	if f := item.Field("hostname"); f == nil || f.Value == "" || f.Value != a.Host {
		return false
	}
	// optional
	if f := item.Field("path"); f != nil && f.Value != "" && f.Value != a.Path {
		return false
	}
	return true
}
