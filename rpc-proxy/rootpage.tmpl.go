package main

import (
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

func rootpage(w http.ResponseWriter, req *http.Request) {
  // The "/" pattern matches everything, so we need to check
  // that we're at the root here.
  if req.URL.Path != "/" {
	  http.NotFound(w, req)
	  return
  }

	// get rid of :port
	host := req.Host
	if i := strings.Index(host, ":"); i != -1 {
		host = host[:i]
	}

	templ.Execute(w, struct {
		KodiURL string
		KodiRPC string
	}{
		KodiURL: "http://" + host + ":" + strconv.Itoa(*kodiHttpPort),
		KodiRPC: host + ":" + strconv.Itoa(*kodiRpcPort),
	})
}

var templ = template.Must(template.New("").Parse(content))

// %%%%content=rootpage.htm%%%%
// START OF AUTOGENERATED CONTENT - DO NOT CHANGE ANYTHING BELOW
// GENERATED: 2018-03-20 18:28:50.3567943 +0100 CET m=+0.016000201
var content = `<!doctype html><meta charset=utf-8><meta name=viewport content="width=device-width,initial-scale=1"><title>Kodi</title><style>html,body{height:100%;background-color:#d2d2d2;margin:0}@media(max-width:750px){.InputAddOn{display:block!important;text-align:center;margin:0 auto}.InputAddOn-field,.InputAddOn-item{padding:1em!important}input[name=url]{width:84%!important}}a{color:#008cba}.container{height:100%;font-family:helvetica neue,Helvetica,Arial,sans-serif;display:-webkit-flex;display:-ms-flexbox;display:flex;align-items:center;-webkit-align-items:center;justify-content:center;-webkit-justify-content:center}.InputAddOn{display:-webkit-box;display:-webkit-flex;display:-ms-flexbox;display:flex;margin-bottom:1.5em}#sec-link{text-align:center;margin:0 auto}.InputAddOn-field{-webkit-box-flex:1;-webkit-flex:1;-ms-flex:1;flex:1}.InputAddOn-field:not(:first-child){border-left:0}.InputAddOn-field:not(:last-child){border-right:0}.InputAddOn-item{background-color:#008cba;color:#fff;box-shadow:rgba(0,0,0,.2) 0 -2px 0 0 inset;-webkit-box-shadow:rgba(0,0,0,.2) 0 -2px 0 0 inset;font-weight:400;cursor:pointer;font-weight:700;margin-left:-1px}.InputAddOn-item:disabled{background-color:#969696;cursor:default}.InputAddOn-field,.InputAddOn-item{border:1px solid rgba(147,128,108,.25);padding:1em 2em;font-size:12pt;-webkit-appearance:none}.InputAddOn-field:first-child,.InputAddOn-item:first-child{border-radius:2px 0 0 2px}.InputAddOn-field:last-child,.InputAddOn-item:last-child{border-radius:0 2px 2px 0}.sk-wave{margin:40px auto;width:50px;height:40px;text-align:center;font-size:10px}.sk-wave .sk-rect{background-color:#008cba;height:100%;width:6px;display:inline-block;-webkit-animation:sk-waveStretchDelay 1.2s infinite ease-in-out;animation:sk-waveStretchDelay 1.2s infinite ease-in-out}.sk-wave .sk-rect1{-webkit-animation-delay:-1.2s;animation-delay:-1.2s}.sk-wave .sk-rect2{-webkit-animation-delay:-1.1s;animation-delay:-1.1s}.sk-wave .sk-rect3{-webkit-animation-delay:-1s;animation-delay:-1s}.sk-wave .sk-rect4{-webkit-animation-delay:-.9s;animation-delay:-.9s}.sk-wave .sk-rect5{-webkit-animation-delay:-.8s;animation-delay:-.8s}@-webkit-keyframes sk-waveStretchDelay{0%,40%,100%{-webkit-transform:scaleY(.4);transform:scaleY(.4)}20%{-webkit-transform:scaleY(1);transform:scaleY(1)}}@keyframes sk-waveStretchDelay{0%,40%,100%{-webkit-transform:scaleY(.4);transform:scaleY(.4)}20%{-webkit-transform:scaleY(1);transform:scaleY(1)}}.slidecontainer{width:100%}.slider{-webkit-appearance:none;width:100%;height:10px;border-radius:3px;background:#fff;outline:0;opacity:.7;-webkit-transition:.2s;transition:opacity .2s}.slider:hover{opacity:1}.slider::-webkit-slider-thumb{-webkit-appearance:none;appearance:none;width:20px;height:20px;border-radius:50%;background:#008cba;cursor:pointer}.slider::-moz-range-thumb{width:20px;height:20px;border-radius:50%;background:#008cba;cursor:pointer}</style><script>var active_playerid=-1;var refresh_timer=null;var activeplayer_timer=null;var connection=new WebSocket('ws://{{ .KodiRPC }}');function resetForm(in_progress){var link=document.getElementById("link");var progress=document.getElementById("progress");var button=document.getElementById("send-link");if(in_progress){progress_state="block";}
else{progress_state="none";}
progress.style.display=progress_state;button.disabled=in_progress;link.disabled=in_progress;}
function sendUrl(){resetForm(true);fetch("/jsonrpc",{headers:{"Content-Type":"application/json"},method:"post",body:JSON.stringify({"url":link.value})}).then(function(response){resetForm(false);link.value="";}).catch(function(err){resetForm(false);alert(err);});}
function getPercentage(){connection.send(JSON.stringify({"id":1,"jsonrpc":"2.0","method":"Player.GetProperties","params":{"playerid":playerid,"properties":["percentage","time","totaltime"]}}));clearTimeout(refresh_timer);clearTimeout(activeplayer_timer);refresh_timer=window.setTimeout(getPercentage,1000);}
function getActivePlayers(){connection.send(JSON.stringify({"id":1,"jsonrpc":"2.0","method":"Player.GetActivePlayers"}));clearTimeout(activeplayer_timer);if(active_playerid!=-1){activeplayer_timer=window.setTimeout(getActivePlayers,1000);}}
function seek(value){connection.send(JSON.stringify({"id":1,"jsonrpc":"2.0","method":"Player.Seek","params":{"playerid":playerid,"value":parseFloat(value)}}));}
connection.onerror=function(e){console.log('Error: '+e.data);}
connection.onmessage=function(e){var result=JSON.parse(e.data)["result"];if(result===undefined){return;}
if(result.length==1&&result[0]===Object(result[0])&&"playerid"in result[0])
{active_playerid=result[0]["playerid"]
getPercentage();}
else if(result===Object(result)&&"percentage"in result){document.getElementById("progress").value=result["percentage"];document.getElementById("progress-text").innerHTML=result["time"]["hours"]+"h"+result["time"]["minutes"]+"m"+result["time"]["seconds"]+"s";document.getElementById("slider").style["visibility"]="visible";}}
connection.onopen=function(){getActivePlayers();};</script><div class=container><form><div class=InputAddOn><input name=url id=link class=InputAddOn-field placeholder=http://... autocomplete=off type=url size=50 autofocus required>
<input class=InputAddOn-item id=send-link type=submit onclick=sendUrl()></div><div class=sk-wave id=progress style=display:none><div class="sk-rect sk-rect1"></div><div class="sk-rect sk-rect2"></div><div class="sk-rect sk-rect3"></div><div class="sk-rect sk-rect4"></div><div class="sk-rect sk-rect5"></div></div><div id=slider class=slidecontainer style=visibility:collapse><input type=range min=1 max=100 class=slider id=progress onchange=seek(this.value)><div style=text-align:center;margin:10px;color:#00475f id=progress-text></div></div><div style=text-align:center><a href="{{ .KodiURL }}">Kodi Web Interface</a></div></form></div>`