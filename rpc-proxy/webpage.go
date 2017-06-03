package main

import (
	"net/http"
)

var html = []byte(`
<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
  <title>Kodi</title>
  <style>
    html, body {
      height: 100%;
      background-color: rgb(210, 210, 210);
      margin: 0;
    }

    .container {
      height: 100%;
      font-family: 'Helvetica Neue', Helvetica, Arial, sans-serif;
      display: -webkit-flex;
      display: -ms-flexbox;
      display: flex;
      align-items: center;
      -webkit-align-items: center;
      justify-content: center;
      -webkit-justify-content: center;
    }

    .InputAddOn {
      display: -webkit-box;
      display: -webkit-flex;
      display: -ms-flexbox;
      display: flex;
      margin-bottom: 1.5em;
    }

    #sec-link {
      text-align: center;
      margin: 0 auto;
    }

    .InputAddOn-field {
      -webkit-box-flex: 1;
      -webkit-flex: 1;
      -ms-flex: 1;
      flex: 1;
    }
    .InputAddOn-field:not(:first-child) {
      border-left: 0;
    }
    .InputAddOn-field:not(:last-child) {
      border-right: 0;
    }

    .InputAddOn-item {
      background-color: rgb(0, 140, 186);
      color: rgb(255, 255, 255);
      box-shadow: rgba(0, 0, 0, 0.2) 0px -2px 0px 0px inset;
      -webkit-box-shadow: rgba(0, 0, 0, 0.2) 0px -2px 0px 0px inset;
      font-weight: normal;
      cursor: pointer;
      font-weight: bold;
      margin-left: -1px;
    }

    .InputAddOn-item:disabled {
      background-color: rgb(150, 150, 150);
      cursor: default;
    }

    .InputAddOn-field,
    .InputAddOn-item {
      border: 1px solid rgba(147, 128, 108, 0.25);
      padding: 1em 2em;
      font-size: 12pt;
    }
    .InputAddOn-field:first-child,
    .InputAddOn-item:first-child {
      border-radius: 2px 0 0 2px;
    }
    .InputAddOn-field:last-child,
    .InputAddOn-item:last-child {
      border-radius: 0 2px 2px 0;
    }

    /* Spinner from: http://tobiasahlin.com/spinkit/
     *  Usage:
     *
     *    <div class="sk-wave">
     *      <div class="sk-rect sk-rect1"></div>
     *      <div class="sk-rect sk-rect2"></div>
     *      <div class="sk-rect sk-rect3"></div>
     *      <div class="sk-rect sk-rect4"></div>
     *      <div class="sk-rect sk-rect5"></div>
     *    </div>
     *
     */
    .sk-wave {
      margin: 40px auto;
      width: 50px;
      height: 40px;
      text-align: center;
      font-size: 10px; }
      .sk-wave .sk-rect {
        background-color: rgb(0, 140, 186);
        height: 100%;
        width: 6px;
        display: inline-block;
        -webkit-animation: sk-waveStretchDelay 1.2s infinite ease-in-out;
                animation: sk-waveStretchDelay 1.2s infinite ease-in-out; }
      .sk-wave .sk-rect1 {
        -webkit-animation-delay: -1.2s;
                animation-delay: -1.2s; }
      .sk-wave .sk-rect2 {
        -webkit-animation-delay: -1.1s;
                animation-delay: -1.1s; }
      .sk-wave .sk-rect3 {
        -webkit-animation-delay: -1s;
                animation-delay: -1s; }
      .sk-wave .sk-rect4 {
        -webkit-animation-delay: -0.9s;
                animation-delay: -0.9s; }
      .sk-wave .sk-rect5 {
        -webkit-animation-delay: -0.8s;
                animation-delay: -0.8s; }
    @-webkit-keyframes sk-waveStretchDelay {
      0%, 40%, 100% {
        -webkit-transform: scaleY(0.4);
                transform: scaleY(0.4); }
      20% {
        -webkit-transform: scaleY(1);
                transform: scaleY(1); } }
    @keyframes sk-waveStretchDelay {
      0%, 40%, 100% {
        -webkit-transform: scaleY(0.4);
                transform: scaleY(0.4); }
      20% {
        -webkit-transform: scaleY(1);
                transform: scaleY(1); } }
  </style>
</head>
<body>
  <div class="container"><form>
    <div class="InputAddOn">
      <input name="url" id="link" class="InputAddOn-field" placeholder="http://..." autocomplete="off" type="url" size="50" autofocus required /><button class="InputAddOn-item" id="send-link" type="button" onclick="sendUrl()">Send</button>
    </div>
    <div class="sk-wave" id="progress" style="display: none">
      <div class="sk-rect sk-rect1"></div>
      <div class="sk-rect sk-rect2"></div>
      <div class="sk-rect sk-rect3"></div>
      <div class="sk-rect sk-rect4"></div>
      <div class="sk-rect sk-rect5"></div>
    </div>
  </form></div>
</body>
<script>
  function sendUrl() {
    var link = document.getElementById("link");
    var progress = document.getElementById("progress");
    var button = document.getElementById("send-link");

    progress.style.display = "block";
    button.disabled = true;
    link.disabled = true;

    var xhr = new XMLHttpRequest();
    xhr.onreadystatechange = function() {
      if (xhr.readyState == XMLHttpRequest.DONE) {
        progress.style.display = "none";
        button.disabled = false;
        link.disabled = false;
      }
    }

    xhr.open("POST", "/jsonrpc", true);
    xhr.setRequestHeader('Content-type', 'application/json');
    xhr.send(JSON.stringify({"url": link.value}));
  }
</script>
</html>
`)

func HandleWebpage(w http.ResponseWriter, req *http.Request) {
	w.Write(html)
}
