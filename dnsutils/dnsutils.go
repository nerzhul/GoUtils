package dnsutils

import (
	"github.com/miekg/dns"
	"net"
)

func Query (query string, query_type uint16) ([]string, string) {
	config, _ := dns.ClientConfigFromFile("/etc/resolv.conf")
	m := new(dns.Msg)
	m.SetQuestion(dns.Fqdn(query), query_type)
	m.RecursionDesired = true
	r, err := dns.Exchange(m, net.JoinHostPort(config.Servers[0], config.Port))
	if r == nil {
		return []string{}, "DNS error: " + err.Error()
	}

	if r.Rcode != dns.RcodeSuccess {
		return []string{}, "DNS error: didn't success to do query for '" + query + "'"
	}

	answer := make([]string, len(r.Answer))

	for i, a := range r.Answer {
		switch a.Header().Rrtype {
			case dns.TypeMX:
				answer[i] = a.(*dns.MX).Mx
				break
			case dns.TypeNS:
				answer[i] = a.(*dns.NS).Ns
				break
			case dns.TypeTXT:
				answer[i] = a.(*dns.TXT).Txt[0]
				break
			case dns.TypeSRV:
				answer[i] = a.(*dns.SRV).Target
				break
			case dns.TypeA:
				answer[i] = a.(*dns.A).A.String()
				break
			case dns.TypeAAAA:
				answer[i] = a.(*dns.AAAA).AAAA.String()
				break
			default: return []string{}, "Unhandled response type"
		}
	}

	return answer, ""
}
