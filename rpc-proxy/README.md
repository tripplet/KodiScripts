# RPC-Proxy

Small rpc-proxy for easier sending of urls to kodi.
(For example with iOS Workflow)

Supports:

- youtube.com & youtu.be links
	- Timestamp supported 
- Generic urls

## Usage
Systemd service file

```
Usage of ./rpc-proxy:
 -http int
 	http port of kodi
 -rpc
 	rpc port of kodi
 -server
      Server port (default 8000)
```

## RPC Format

Json post requst to %kodi-ip%:8000/jsonrpc

{"url": "youtube.com/kjasdlkajsld"}

