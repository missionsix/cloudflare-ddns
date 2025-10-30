package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"

	cloudflare "github.com/cloudflare/cloudflare-go/v6"
	dns "github.com/cloudflare/cloudflare-go/v6/dns"
	option "github.com/cloudflare/cloudflare-go/v6/option"
)

var (
	dnsRecordId string = ""
)

func main() {
	// Read environment variables
	cloudflareApiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if cloudflareApiToken == "" {
		fmt.Println("Error: CLOUDFLARE_API_TOKEN environment variable is required")
		os.Exit(1)
	}

	// Define command line flag for IPv4 address
	zoneId := flag.String("zone-id", "", "DNZ Zone ID (required)")
	dnsRecordName := flag.String("record-name", "", "DNS record name (required)")
	ipAddress := flag.String("ip", "", "IPv4 address to update DNS record with (required)")
	flag.Parse()

	if *zoneId == "" {
		fmt.Println("Error: Zone ID is required")
		fmt.Println("Usage: go run ddns.go -zone-id <ZONE_ID> -record-name <RECORD_NAME> -ip <IPv4_ADDRESS>")
		os.Exit(1)
	}

	if *dnsRecordName == "" {
		fmt.Println("Error: DNS record name is required")
		fmt.Println("Usage: go run ddns.go -zone-id <ZONE_ID> -record-name <RECORD_NAME> -ip <IPv4_ADDRESS>")
		os.Exit(1)
	}

	// Validate that IP address was provided
	if *ipAddress == "" {
		fmt.Println("Error: IPv4 address is required")
		fmt.Println("Usage: go run ddns.go -ip <IPv4_ADDRESS>")
		os.Exit(1)
	}

	// Validate that the provided string is a valid IPv4 address
	parsedIP := net.ParseIP(*ipAddress)
	if parsedIP == nil || parsedIP.To4() == nil {
		fmt.Printf("Error: '%s' is not a valid IPv4 address\n", *ipAddress)
		os.Exit(1)
	}

	client := cloudflare.NewClient(
		option.WithAPIToken(cloudflareApiToken),
	)

	ctx := context.Background()

	page, err := client.DNS.Records.List(ctx, dns.RecordListParams{
		ZoneID: cloudflare.F(*zoneId),
		Name:   cloudflare.F(dns.RecordListParamsName{Exact: cloudflare.F(*dnsRecordName)}),
	})
	if err != nil {
		panic(err.Error())
	}

	record := page.Result[0]
	dnsRecordId = record.ID

	if record.Content == *ipAddress {
		fmt.Println("DNS record is already up to date. No changes made.")
		return
	}

	fmt.Printf("DNS Record ID: %s\n", dnsRecordId)

	_, err = client.DNS.Records.Update(ctx, dnsRecordId, dns.RecordUpdateParams{
		ZoneID: cloudflare.F(*zoneId),
		Body: &dns.ARecordParam{
			Name:    cloudflare.F(*dnsRecordName),
			TTL:     cloudflare.F(record.TTL),
			Type:    cloudflare.F(dns.ARecordType(record.Type)),
			Comment: cloudflare.F("Updated via API"),
			Content: cloudflare.F(*ipAddress), // Use the provided IP address
			Proxied: cloudflare.Bool(record.Proxied),
		},
	})
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Successfully updated DNS record '%s' to IP address '%s'\n", dnsRecordName, *ipAddress)
}
