package utils

import (
	"log"
	"strings"
	"time"
)

func ParseDateStrToTime(date_str string) (*time.Time, error) {
	strings.ReplaceAll(date_str, "/", "-")
	date_str, _, _ = strings.Cut(date_str, "T")
	datetime, err := time.Parse(time.RFC3339, date_str+"T00:00:00-03:00")
	if err != nil {
		log.Printf("fail on converting date_str to datetime %v", err)
		return nil, err
	}
	return &datetime, nil
}
