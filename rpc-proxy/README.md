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
 -port int
      Server port (default 8000)
```

## RPC Format

Json post requst to %kodi-ip%:8000/jsonrpc

{"ur": "youtube.com/kjasdlkajsld"}

