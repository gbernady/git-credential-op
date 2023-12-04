package helper

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

// Atrributes represents a set of git credential named attributes passed to the git credential helper.
// See https://git-scm.com/docs/git-credential#IOFMT for more details.
type Attributes struct {
	Protocol          string
	Host              string
	Path              string
	Username          string
	Password          string
	PasswordExpiry    time.Time
	OAuthRefreshToken string
	URL               string
	WWWAuth           []string
}

// ParseAttributes parses git credential into named attributes.
func ParseAttributes(r io.Reader) *Attributes {
	a := &Attributes{}
	s := bufio.NewScanner(r)
	for s.Scan() {
		v := strings.SplitN(s.Text(), "=", 2)
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
		case "password_expiry_utc":
			v, _ := strconv.ParseInt(val, 10, 0)
			a.PasswordExpiry = time.Unix(v, 0)
		case "oauth_refresh_token":
			a.OAuthRefreshToken = val
		case "url":
			a.URL = val
		case "wwwauth[]":
			a.WWWAuth = append(a.WWWAuth, val)
		default:
			continue
		}
	}
	return a
}

func (a *Attributes) String() string {
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
	if !a.PasswordExpiry.IsZero() {
		fmt.Fprintf(&b, "password_expiry_utc=%d\n", a.PasswordExpiry.Unix())
	}
	if a.OAuthRefreshToken != "" {
		fmt.Fprintf(&b, "oauth_refresh_token=%s\n", a.OAuthRefreshToken)
	}
	if a.URL != "" {
		fmt.Fprintf(&b, "url=%s\n", a.URL)
	}
	for _, v := range a.WWWAuth {
		fmt.Fprintf(&b, "wwwauth[]=%s\n", v)
	}
	return b.String()
}
