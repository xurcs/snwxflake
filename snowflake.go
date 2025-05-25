package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

const SnowflakeEpoch = 1288834974657

func SnowFlakeToTime(snowflakeID uint64) time.Time {
	timestampMs := (snowflakeID >> 22) + SnowflakeEpoch
	return time.UnixMilli(int64(timestampMs)).UTC()
}

func ExtractSnowflakeID(input string) (uint64, error) {
	if strings.Contains(input, "x.com") {
		parts := strings.Split(input, "/")
		if len(parts) < 6 {
			return 0, fmt.Errorf("invalid URL format")
		}
		snowflakeStr := parts[len(parts)-1]
		return strconv.ParseUint(snowflakeStr, 10, 64)
	}
	return strconv.ParseUint(input, 10, 64)
}

func main() {
	var input string
	fmt.Println("Tweet url must be in the format: https://x.com/elonmusk/status/<snowflake_id>")
	fmt.Println("Enter a full tweet URL or snowflake ID:")
	fmt.Scanln(&input)

	snowflakeID, err := ExtractSnowflakeID(input)
	if err != nil {
		fmt.Println("Invalid input. Enter a valid tweet URL or snowflake ID.")
		return
	}

	t := SnowFlakeToTime(snowflakeID)
	currentTime := time.Now()
	epochTime := time.UnixMilli(SnowflakeEpoch)
	if t.Before(epochTime) || t.After(currentTime) {
		fmt.Println("The snowflake ID is invalid or does not correspond to a valid time.")
		return
	}

	tokyo, _ := time.LoadLocation("Asia/Tokyo")
	const iso8601WithMillis = "2006-01-02T15:04:05.000Z"
	const iso8601WithOffset = "2006-01-02T15:04:05.000Z07:00"

	cyan := color.New(color.FgCyan).SprintFunc()

	fmt.Printf("The tweet was created at: %s (UTC) / %s (Asia/Tokyo)\n",
		cyan(t.Format(iso8601WithMillis)),
		cyan(t.In(tokyo).Format(iso8601WithOffset)))

	fmt.Println("\nUNIX milliseconds:")
	fmt.Printf("%s\n", cyan(t.UnixMilli()))

	fmt.Println("\nISO 8601:")
	fmt.Printf("%s\n", cyan(t.Format(iso8601WithMillis)))
}
