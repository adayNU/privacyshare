package main

import (
	"net/url"
	"os"
	"strings"

	"github.com/atotto/clipboard"
)

func main() {
	var s, err = clipboard.ReadAll()
	if err != nil {
		os.Exit(1)
		return
	}

	for {
		var temp string
		temp, err = clipboard.ReadAll()

		switch {
		case err != nil:
			os.Exit(1)
			return
		case temp != s:
			var new string
			new, err = stripQuery(temp)
			if err != nil || new == "" {
				// Silently fail on non-URL.
				break
			}

			err = clipboard.WriteAll(new)
			if err != nil {
				os.Exit(1)
				return
			}
			s = temp
		}
	}
}

// stripQuery removes any known tracking parameters from |s|.
func stripQuery(s string) (string, error) {
	var u, err = url.Parse(s)
	if err != nil {
		return "", err
	}

	// Seems virtually impossible that if the scheme is not http(s) that
	// we need to strip any query parameters.
	if !strings.HasPrefix(u.Scheme, "http") {
		return s, nil
	}

	var q = u.Query()

	switch u.Host {
	case "open.spotify.com":
		q.Del("si")
	case "www.amazon.com":
		if u.Path == "/s" {
			q = url.Values{
				"k": []string{q.Get("k")},
			}
			break
		}
		q = url.Values{}
	case "twitter.com", "x.com":
		q.Del("t")
	case "www.threads.net":
		q.Del("xmt")
	}

	for k, _ := range q {
		if strings.HasPrefix(k, "utm_") {
			q.Del(k)
		}
	}

	u.RawQuery = q.Encode()

	return u.String(), nil
}
