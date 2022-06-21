package filter

import (
	"github.com/miekg/dns"
)

func createReply(r *dns.Msg, ttl uint32) *dns.Msg {
	return newNXDomainResponse(r, ttl)
}

func newNXDomainResponse(r *dns.Msg, ttl uint32) *dns.Msg {
	msg := new(dns.Msg)
	msg.SetReply(r)
	msg.SetRcode(r, dns.RcodeNameError)

	msg.Ns = []dns.RR{&dns.SOA{
		Refresh: 1800,
		Retry:   900,
		Expire:  604800,
		Minttl:  86400,
		Ns:      "fake-for-negative-caching.dns.paesa.es.",
		Serial:  100500,

		Hdr: dns.RR_Header{
			Name:   r.Question[0].Name,
			Rrtype: dns.TypeSOA,
			Ttl:    ttl,
			Class:  dns.ClassINET,
		},
		Mbox: "hostmaster.", // zone will be appended later if it's not empty or "."
	}}
	return msg
}
