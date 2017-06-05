package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//go:generate go run gen.go

var defaultProxyServerPort = 8000
var defaultKodiPort = 8080

var youtube = strings.Replace("{'jsonrpc': '2.0', 'method': 'Player.Open', 'params':{"+
	"'item': {'file':'plugin://plugin.video.youtube/?action="+
	"play_video&videoid=%s'}}, 'id' : 1}", "'", "\"", -1)
var youtubeTime = strings.Replace("{'jsonrpc':'2.0','method':'Player.Open','params':{"+
	"'item':{'file':'plugin://plugin.video.youtube/?action="+
	"play_video&videoid=%s'},"+
	"'options':"+
	"{'resume':{'hours':%d,'minutes':%d,'seconds':%d}}"+
	"},'id':1},", "'", "\"", -1)
var genericUrl = strings.Replace("{'jsonrpc': '2.0', 'method': 'Player.open', 'params': "+
	"{'item': {'file': '%s'}}, 'id': 1}", "'", "\"", -1)

type request struct {
	Url string
}

var kodiPort *int

func main() {
	serverPort := flag.Int("server", defaultProxyServerPort, "Kodi web port")
	kodiPort = flag.Int("kodi", defaultKodiPort, "Proxy server port")
	flag.Parse()

	fmt.Println("Starting server on port", *serverPort)
	http.HandleFunc("/jsonrpc", HandleRpc)
	http.HandleFunc("/", rootpage_htm)
	http.ListenAndServe(fmt.Sprintf(":%d", *serverPort), nil)
}

func divmod(value int64, modulo int64) (int64, int64) {
	mod := (value % modulo)
	return (value - mod) / modulo, mod
}

func generatePayload(video string) string {
	videoId := ""
	data := ""

	if !strings.HasPrefix(video, "http") {
		video = "http://" + video
	}

	// Parse the url
	videoUrl, err := url.Parse(video)
	if err != nil {
		panic(err)
	}

	// Extract video id for youtube
	if regexp.MustCompile("[/\\.]youtube\\.com/").FindString(video) != "" {
		videoId = videoUrl.Query().Get("v")
	} else if regexp.MustCompile("[/\\.]youtu\\.be/").FindString(video) != "" {
		videoId = videoUrl.EscapedPath()[1:]
	} else {
		data = fmt.Sprintf(genericUrl, video)
	}

	if data == "" && videoUrl.Query().Get("t") != "" {
		seek := 0 * time.Second
		timestr := videoUrl.Query().Get("t")

		if value, err := strconv.ParseInt(timestr, 10, 64); err == nil {
			seek = time.Duration(value) * time.Second
		} else {
			seek, _ = time.ParseDuration(videoUrl.Query().Get("t"))
		}

		// go back 15s for viewing convenience
		seek -= 15 * time.Second
		if seek < 0*time.Second {
			seek = 0 * time.Second
		}

		minutes, seconds := divmod(int64(seek.Seconds()), 60)
		hours, minutes := divmod(minutes, 60)

		data = fmt.Sprintf(youtubeTime, videoId, hours, minutes, seconds)
	}

	if data == "" {
		data = fmt.Sprintf(youtube, videoId)
	}

	return data
}

func HandleRpc(w http.ResponseWriter, req *http.Request) {
	s := request{}
	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(body, &s)
	if err != nil || s.Url == "" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	payload := bytes.NewBufferString(generatePayload(s.Url))
	rsp, err := http.Post(fmt.Sprintf("http://localhost:%d/jsonrpc", *kodiPort), "application/json", payload)
	if err != nil {
		panic(err)
	}

	response, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		panic(err)
	}

	w.Write(response)
}
