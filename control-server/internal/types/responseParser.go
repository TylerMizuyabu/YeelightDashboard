package types

import (
	"bufio"
	"errors"
	"net/http"
	"strings"
)

var crlf = "\r\n"

// ErrorNoMatchingHeader used when no matching header is found in the response
var ErrorNoMatchingHeader = errors.New("No Matching Header Found")

// ResponseParser contains methods for parsing parts of the response
type ResponseParser struct {
	msg string
}

// NewParser function returns a ResponseParser
func NewParser(msg string) *ResponseParser {
	if strings.HasSuffix(msg, crlf) {
		msg = msg + crlf
	}

	return &ResponseParser{
		msg: msg,
	}
}

// ParseAddr function parses the ip address part of the string value under the LOCATION header and stores it in provided string
func (rp *ResponseParser) ParseAddr(a *string) error {
	var addr string
	err := rp.ParseHeader("LOCATION", false, &addr)
	if err != nil {
		return err
	}
	*a = strings.TrimPrefix(addr, "yeelight://")
	return nil
}

// ParseHeader function reads the response and sets the value of the provided string pointer to the header value that corresponds with the provided header key
func (rp *ResponseParser) ParseHeader(header string, emptyAllowed bool, s *string) error {
	resp, err := http.ReadResponse(bufio.NewReader(strings.NewReader(rp.msg)), nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	*s = resp.Header.Get(header)
	if *s == "" && emptyAllowed {
		return ErrorNoMatchingHeader
	}
	return nil
}
