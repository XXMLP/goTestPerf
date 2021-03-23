package client

import (
	"github.com/miekg/dns"
)

var (
	// DNSClient global dns client object
	DNSClient *dns.Client = &dns.Client{}
)
