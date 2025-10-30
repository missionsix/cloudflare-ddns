#!/usr/bin/bash

export WORKDIR=/data/dev/ddns

# Load environment variables from .env file if it exists
if [ -f $WORKDIR/.env ]; then
    export $(grep -v '^#' $WORKDIR/.env | xargs)
fi

WAN_IP=$(curl -s 'https://api.ipify.org?format=text')

echo "Current WAN IP: $WAN_IP"

DDNS_BIN=$WORKDIR/ddns

if ! [ -f $DDNS_BIN ]; then
    echo "ddns binary not found."
    exit 1
fi

$DDNS_BIN -zone-id $CLOUDFLARE_ZONE_ID -record-name $CLOUDFLARE_DNS_RECORD_NAME -ip $WAN_IP

exit $?
