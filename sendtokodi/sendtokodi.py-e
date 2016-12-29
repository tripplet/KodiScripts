#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import platform
import re
import subprocess
import sys
import urllib.parse as urlparse
import webbrowser

import requests

try:
    # Try to load pythonista specific modules
    import console
    import appex
    import dialogs
    from dns import ParallelResolve
    _platform = 'Pythonista'
except:
    _platform = platform.system()

# Global variables
verbose = True
payload = {
    'youtube': '{"jsonrpc": "2.0", "method": "Player.Open", "params":{' +
               '"item": {"file":"plugin://plugin.video.youtube/?action=' +
               'play_video&videoid=%s"}}, "id" : 1}',
    'youtube_time': '{"jsonrpc":"2.0","method":"Player.Open","params":{' +
                    '"item":{"file":"plugin://plugin.video.youtube/?action=' +
                    'play_video&videoid=%s"},' +
                    '"options":' +
                    '{"resume":{"hours":%d,"seconds":%d,"minutes":%d}}' +
                    '},"id":1},',
    'generic_url': '{"jsonrpc": "2.0", "method": "Player.open", "params": ' +
                   '{"item": {"file": "%s"}}, "id": 1}'
}
airport_command = "/System/Library/PrivateFrameworks/Apple80211.framework/" \
                  "Versions/Current/Resources/airport"
server = {
    'kodi-host.local': ['wifi-name', 'other-wifi-name2'],
    'other-kodi-host.local': ['wifi-name', 'wifi-name2']
}


def main():
    url = None
    kodi_server = None

    # Switch to console on pythonista
    if is_pythonista():
        console.show_activity()

    if is_pythonista() and appex.is_running_extension():
        # Get url handed over to the extension
        url = appex.get_url()
        if url is None:
            url = appex.get_text()
        kodi_server = get_server(server)
    else:
        # Get url and server from command line arguments to script
        if len(sys.argv) < 2:
            info('No url given', 'error', 3)
            return
        else:
            url = sys.argv[1]
        if len(sys.argv) < 3:
            kodi_server = get_server(server)
        else:
            kodi_server = sys.argv[2]

    # Send url to kodi
    ret_msg = send_url(url=url, ip=kodi_server)

    # Handle response from kodi
    if 'error' in ret_msg:
        info(ret_msg['error']['message'], 'error', 3)
    else:
        info('Done', 'success')
        if is_pythonista():
            webbrowser.open('xbmcremote://')


def is_pythonista():
    return _platform == 'Pythonista'


def send_url(url, ip):
    data = None

    log('url: {}'.format(url))
    log('ip: {}'.format(ip))

    if not url.startswith('http'):
        url = 'http://' + url

    parsed_url = urlparse.urlparse(url)
    url_parameter = urlparse.parse_qs(parsed_url.query)

    if re.search('[/\.]youtube\.com/', url) is not None:
        video_id = url_parameter['v'][0]
    elif re.search('[/\.]youtu\.be/', url) is not None:
        video_id = parsed_url.path[1:]
    else:
        data = payload['generic_url'] % (url)

    if data is None and 't' in url_parameter:
        # Get seek time from url and go back 10s for viewing convenience
        seek_time = max(int(url_parameter['t'][0]) - 10, 0)
        minutes, seconds = divmod(seek_time, 60)
        hours, minutes = divmod(minutes, 60)
        data = payload['youtube_time'] % (video_id, hours, minutes, seconds)
        log('seek time : {}h {}m {}s'.format(hours, minutes, seconds))

    if data is None:
        data = payload['youtube'] % (video_id)

    info('Sending...', 'success')
    log('Sending: {}'.format(data))

    response = requests.request('POST', 'http://%s/jsonrpc' % (ip),
                                headers={'Content-type': 'application/json'},
                                data=data)
    ret_msg = response.json()

    log('Received: {}'.format(ret_msg))
    return ret_msg


def info(msg, type, timeout=1.8):
    if is_pythonista():
        dialogs.hud_alert(msg, type, timeout)
    else:
        print(msg)


def log(str):
    if is_pythonista() or verbose:
        print('+ ' + str)


def get_server(servers):
    if _platform == 'Darwin':
        return get_server_osx(servers)
    elif is_pythonista():
        return get_server_pythonista(servers)
    else:
        raise Exception('Platform not supported')


def get_server_osx(names_to_resolve):
    result = subprocess.check_output([airport_command, '--getinfo'])
    wifi_name = re.search('^\s*SSID:\s+(.*)$', result.decode('latin1'),
                          re.MULTILINE).group(1)

    for server, networks in iter(names_to_resolve.items()):
        if wifi_name in networks:
            return server


def get_server_pythonista(names_to_resolve):
    p = ParallelResolve(names_to_resolve.keys())
    res = p.resolve()
    if res is not None:
        print(res[0])
        return res[0]


# Script start
if __name__ == '__main__':
    main()
