package main

import (
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const youtube = `{"jsonrpc": "2.0", "method": "Player.Open", "params": {` +
	`"item": {"file": "plugin://plugin.video.youtube/?action=` +
	`play_video&videoid=%s"}}, "id" : 1}`
const youtubeAndSeek = `{"jsonrpc": "2.0","method": "Player.Open","params": {` +
	`"item": {"file": "plugin://plugin.video.youtube/?action=` +
	`play_video&videoid=%s"}, ` +
	`"options": ` +
	`{"resume": {"hours": %d, "minutes": %d, "seconds": %d}}` +
	`},"id": 1}`
const genericURL = `{"jsonrpc": "2.0", "method": "Player.open", "params": ` +
	`{"item": {"file": "%s"}}, "id": 1}`

func divmod(value int64, modulo int64) (int64, int64) {
	mod := (value % modulo)
	return (value - mod) / modulo, mod
}

func generatePayload(video string) string {
	videoID := ""

	if !strings.HasPrefix(video, "http") {
		video = "http://" + video
	}

	// Parse the url
	videoURL, err := url.Parse(video)
	if err != nil {
		panic(err)
	}

	// Extract video id for youtube
	if regexp.MustCompile("[/\\.]youtube\\.com/").FindString(video) != "" {
		videoID = videoURL.Query().Get("v")
	} else if regexp.MustCompile("[/\\.]youtu\\.be/").FindString(video) != "" {
		videoID = videoURL.EscapedPath()[1:]
	} else {
		return fmt.Sprintf(genericURL, video)
	}

	if videoURL.Query().Get("t") != "" {
		seek := 0 * time.Second
		timestr := videoURL.Query().Get("t")

		if value, err := strconv.ParseInt(timestr, 10, 64); err == nil {
			seek = time.Duration(value) * time.Second
		} else {
			seek, _ = time.ParseDuration(videoURL.Query().Get("t"))
		}

		// go back 15s for viewing convenience
		seek -= 15 * time.Second
		if seek < 0*time.Second {
			seek = 0 * time.Second
		}

		minutes, seconds := divmod(int64(seek.Seconds()), 60)
		hours, minutes := divmod(minutes, 60)

		return fmt.Sprintf(youtubeAndSeek, videoID, hours, minutes, seconds)
	}

	return fmt.Sprintf(youtube, videoID)
}
