package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/gbernady/git-credential-op/pkg/op"
)

type Attributes struct {
	Protocol string
	Host     string
	Path     string
	Username string
	Password string
	URL      string
}

func (a *Attributes) Parse(r io.Reader) {
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

func (a Attributes) String() string {
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

func (a Attributes) Match(item op.Item) bool {
	if v := item.Field("hostname").Value; v != "" && a.Host != "" && v != a.Host {
		return false
	}
	if v := item.Field("path").Value; v != "" && a.Path != "" && v != a.Path {
		return false
	}
	if v := item.Field("username").Value; v != "" && a.Username != "" && v != a.Username {
		return false
	}
	return true
}
