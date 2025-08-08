package ptrval

import "net/url"

func String(s string) *string {
	return &s
}

func MustURL(s string) *url.URL {
	u, err := url.Parse(s)
	if err != nil {
		panic(err) // This should not happen if the URL is valid
	}
	return u
}
