# Cloudflare DDNS updater

This is a simple go tool for updating cloudflare A records using the Cloudflare API.

## Usage:

### Build:
```
$ make build
```

### Run
```
$ export CLOUDFLARE_API_TOKEN="YOUR_API_KEY_HERE"

$ ./ddns -zone-id "YOUR_ZONE_ID" -record-name 'YOUR_RECORD_NAME' -ip "IP ADDRESS"
```

### Helper script

There's a helper script: `ddsh.sh` which uses *ipify* to get your WAN IP. It sources configuration from the .env file.
```
$ touch .env

$ echo "# Cloudflare API Token \n CLOUDFLARE_API_TOKEN=" >> .env
$ echo "# Cloudflare Zone ID \n CLOUDFLARE_ZONE_ID=" >> .env
$ echo "# DNS Record Name \n CLOUDFLARE_DNS_RECORD_NAME=" >> .env

$ ./ddns.sh
```

