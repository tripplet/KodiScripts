
### Play youtube
```json
{
  "jsonrpc": "2.0",
  "method": "Player.Open",
  "params": {
    "item": {
      "file": "plugin://plugin.video.youtube/?action=play_video&videoid={VIDEOID}"
    }
  },
  "id": 1
}
```

### Play url
```json
{
  "id": 1,
  "jsonrpc": "2.0",
  "method": "Player.open",
  "params": {
    "item": {
      "file": "{http://URL-HERE}"
    }
  }
}
```

### Multiple commands in chain (not working properly)
```json
[
  {
    "id": 1,
    "jsonrpc": "2.0",
    "method": "Player.Open",
    "params": {
      "item": {
        "file": "plugin://plugin.video.youtube/?action=play_video&videoid={VIDEOID}"
      }
    }
  },
  {
    "id": 2,
    "jsonrpc": "2.0",
    "method": "Player.PlayPause",
    "params": {
      "playerid": 1
    }
  },
  {
    "id": 3,
    "jsonrpc": "2.0",
    "method": "Player.Seek",
    "params": {
      "playerid": 1,
      "value": {
        "hours": 0,
        "minutes": 0,
        "seconds": 30
      }
    }
  },
  {
    "id": 4,
    "jsonrpc": "2.0",
    "method": "Player.PlayPause",
    "params": {
      "playerid": 1
    }
  }
]
```

### Play and start at certain time
```json
{
  "jsonrpc": "2.0",
  "method": "Player.Open",
  "params": {
    "item": {
      "file": "plugin://plugin.video.youtube/?action=play_video&videoid={VIDEOID}"
    },
    "options": {
        "resume": {
              "hours": 0,
              "minutes": 0,
              "seconds": 30
        }
    }
  },
  "id": 1
}
```

### Seek playback
```json
{
  "jsonrpc": "2.0",
  "method": "Player.Seek",
  "params": {
    "playerid": 1,
    "value": {
      "hours": 0,
      "minutes": 0,
      "seconds": 30
    }
  },
  "id": 2
}
```
