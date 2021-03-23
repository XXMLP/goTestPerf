package suite

import (
	"errors"
	"flag"
	"math/rand"
	"pef/client"
	peflag "pef/flag"
	"strings"

	"github.com/feiyuw/boomer"
	"github.com/miekg/dns"
)

func init() {
	SuiteMap.Add("dns", NewDNSSuite())
}

// DNSSuite object of dns pet tester
type DNSSuite struct {
	Protocol string
	Host     string
	Type     uint16
	Records  peflag.ListFlags

	Conn *dns.Conn
}

// NewDNSSuite is used to create a new DNSSuite object
func NewDNSSuite() *DNSSuite {
	return &DNSSuite{Records: peflag.ListFlags{}}
}

// Init is used to argument parsing and initialization
func (ds *DNSSuite) Init(flagSet *flag.FlagSet, args []string) error {
	var typeStr string

	flagSet.StringVar(&ds.Protocol, "protocol", "udp", "DNS server protocol, udp|tcp|tcp-tls")
	flagSet.StringVar(&ds.Host, "host", "127.0.0.1:53", "DNS server host, IP and port")
	flagSet.StringVar(&typeStr, "type", "A", "DNS Type, like A|AAAA|CNAME")
	flagSet.Var(&ds.Records, "record", "DNS record to test")

	flagSet.Parse(args)

	if ds.Protocol != "udp" && ds.Protocol != "tcp" && ds.Protocol != "tcp-tls" {
		return errors.New("protocol should be one of udp|tcp|tcp-tls")
	}
	if !strings.ContainsRune(ds.Host, ':') {
		ds.Host += ":53"
	}
	if typeStr != "A" && typeStr != "AAAA" && typeStr != "CNAME" {
		return errors.New("type should be one of udp|tcp|tcp-tls")
	}
	if len(ds.Records) == 0 {
		return errors.New("specify at least one record")
	}

	for idx, record := range ds.Records {
		if record[len(record)-1] != '.' {
			ds.Records[idx] = record + "."
		}
	}

	switch typeStr {
	case "A":
		ds.Type = dns.TypeA
	case "AAAA":
		ds.Type = dns.TypeAAAA
	case "CNAME":
		ds.Type = dns.TypeCNAME
	}

	client.DNSClient.Net = ds.Protocol

	return nil
}

// GetTask is used to generate a boomer Task instance
func (ds *DNSSuite) GetTask() *boomer.Task {
	var fn func()

	switch ds.Type {
	case dns.TypeA:
		fn = ds.doA
	}

	return &boomer.Task{
		Name:    "dns",
		OnStart: func() {},
		OnStop:  func() {},
		Fn:      fn,
	}
}

func (ds *DNSSuite) doA() {
	msg := new(dns.Msg)
	record := ds.randomRecord()
	msg.SetQuestion(record, ds.Type)
	start := boomer.Now()
	_, _, err := client.DNSClient.Exchange(msg, ds.Host)
	elapsed := boomer.Now() - start

	if err != nil {
		boomer.RecordFailure(record, "err", elapsed, err.Error())
		return
	}

	boomer.RecordSuccess(record, "ok", elapsed, 0)
}

func (ds *DNSSuite) randomRecord() string {
	return ds.Records[rand.Intn(len(ds.Records))]
}
