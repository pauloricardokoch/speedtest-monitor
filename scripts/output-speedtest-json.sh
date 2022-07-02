#!/bin/sh

RESULT=$(speedtest -u Mbps -f json)

DBYTES=$(echo $RESULT | jq .download.bytes)
DELAPSED=$(echo $RESULT | jq .download.elapsed)
echo Download: $(($DBYTES/$DELAPSED*8/1000))

UBYTES=$(echo $RESULT | jq .upload.bytes)
UELAPSED=$(echo $RESULT | jq .upload.elapsed)
echo Upload: $(($UBYTES/$UELAPSED*8/1000))

echo Latency: $(echo $RESULT | jq .ping.latency)

DATA=$(echo $RESULT | jq '{
    isp: .isp,
    host: .server.host, 
    ip: .server.ip, 
    location: .server.location, 
    country: .server.country, 
    "int-ip": .interface.internalIp, 
    "int-name": .interface.name, 
    "int-mac-addr": .interface.macAddr, 
    "int-is-vpn": .interface.isVpn,
    "d-bytes": .download.bytes,
    "d-elapsed": .download.elapsed,
    "u-bytes": .upload.bytes,
    "u-elapsed": .upload.elapsed,
    "latency": .ping.latency
}')

echo -e "\n\nPosting metrics..."
curl -XPOST http://localhost:3001/register --data "$(echo $DATA)" -i
