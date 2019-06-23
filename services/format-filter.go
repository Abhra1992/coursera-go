package services

import (
	"net/url"
	"regexp"
	"strings"
)

const (
	// TrustedFormats file formats that are trusted and should be downloaded
	TrustedFormats = "^mp4$|^srt$|^docx?$|^pptx?$|^xlsx?$|^pdf$|^zip$"

	// ComplexFormats file formats which are non-standard and should be skipped
	ComplexFormats = ".*[^a-zA-Z0-9_-]"
)

func shouldSkipFormatURL(format string, link string) bool {
	if format == "" {
		return true
	}
	if match, err := regexp.MatchString(ComplexFormats, format); match || err != nil {
		return true
	}
	// Skip emails
	if strings.Contains(link, "mailto") || strings.Contains(link, "@") {
		return true
	}
	parsed, err := url.Parse(link)
	// Possible malicious link - skip
	if err != nil {
		return true
	}
	// Skip Localhost
	if parsed.Host == "localhost" {
		return true
	}
	// Skip site root
	if parsed.Path == "" || parsed.Path == "/" {
		return true
	}
	if match, err := regexp.MatchString(TrustedFormats, format); match && err == nil {
		return false
	}
	return false
}
