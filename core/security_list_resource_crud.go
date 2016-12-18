package core

import (
	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/hashicorp/terraform/helper/schema"
)

type SecurityListResourceCrud struct {
	D      *schema.ResourceData
	Client client.BareMetalClient
	Res    *baremetal.SecurityList
}

func (s *SecurityListResourceCrud) ID() string {
	return s.Res.ID
}

func (s *SecurityListResourceCrud) CreatedPending() []string {
	return []string{baremetal.ResourceProvisioning}
}

func (s *SecurityListResourceCrud) CreatedTarget() []string {
	return []string{baremetal.ResourceAvailable}
}

func (s *SecurityListResourceCrud) DeletedPending() []string {
	return []string{baremetal.ResourceTerminating}
}

func (s *SecurityListResourceCrud) DeletedTarget() []string {
	return []string{baremetal.ResourceTerminated}
}

func (s *SecurityListResourceCrud) State() string {
	return s.Res.State
}

func (s *SecurityListResourceCrud) Create() (e error) {
	compartmentID := s.D.Get("compartment_id").(string)
	egress := s.buildEgressRules()
	ingress := s.buildIngressRules()
	vcnID := s.D.Get("vcn_id").(string)

	opts := &baremetal.CreateOptions{}
	opts.DisplayName = s.D.Get("display_name").(string)

	s.Res, e = s.Client.CreateSecurityList(compartmentID, vcnID, egress, ingress, opts)

	return
}

func (s *SecurityListResourceCrud) Get() (e error) {
	s.Res, e = s.Client.GetSecurityList(s.D.Id())
	return
}

func (s *SecurityListResourceCrud) Update() (e error) {
	opts := &baremetal.UpdateSecurityListOptions{}

	if egress := s.buildEgressRules(); egress != nil {
		opts.EgressRules = egress
	}
	if ingress := s.buildIngressRules(); ingress != nil {
		opts.IngressRules = ingress
	}

	s.Res, e = s.Client.UpdateSecurityList(s.D.Id(), opts)
	return
}

func (s *SecurityListResourceCrud) SetData() {
	s.D.Set("compartment_id", s.Res.CompartmentID)
	s.D.Set("display_name", s.Res.DisplayName)

	confEgressRules := []map[string]interface{}{}
	for _, egressRule := range s.Res.EgressSecurityRules {
		confEgressRule := map[string]interface{}{
			"destination": egressRule.Destination,
			"protocol":    egressRule.Protocol,
		}
		if icmpOpts := egressRule.ICMPOptions; icmpOpts != nil {
			confEgressRule["icmp_options"] = s.buildConfICMPOptions(icmpOpts)
		}
		if tcpOpts := egressRule.TCPOptions; tcpOpts != nil {
			confEgressRule["tcp_options"] = s.buildConfTCPOptions(tcpOpts)
		}
		if udpOpts := egressRule.UDPOptions; udpOpts != nil {
			confEgressRule["udp_options"] = s.buildConfUDPOptions(udpOpts)
		}
		confEgressRules = append(confEgressRules, confEgressRule)
	}
	s.D.Set("egress_security_rules", confEgressRules)

	confIngressRules := []map[string]interface{}{}
	for _, ingressRule := range s.Res.IngressSecurityRules {
		confIngressRule := map[string]interface{}{
			"protocol": ingressRule.Protocol,
			"source":   ingressRule.Source,
		}
		if icmpOpts := ingressRule.ICMPOptions; icmpOpts != nil {
			confIngressRule["icmp_options"] = s.buildConfICMPOptions(icmpOpts)
		}
		if tcpOpts := ingressRule.TCPOptions; tcpOpts != nil {
			confIngressRule["tcp_options"] = s.buildConfTCPOptions(tcpOpts)
		}
		if udpOpts := ingressRule.UDPOptions; udpOpts != nil {
			confIngressRule["udp_options"] = s.buildConfUDPOptions(udpOpts)
		}
		confIngressRules = append(confIngressRules, confIngressRule)
	}
	s.D.Set("ingress_security_rules", confIngressRules)

	s.D.Set("state", s.Res.State)
	s.D.Set("time_created", s.Res.TimeCreated.String())
	s.D.Set("vcn_id", s.Res.VcnID)
}

func (s *SecurityListResourceCrud) Delete() (e error) {
	return s.Client.DeleteSecurityList(s.D.Id(), nil)
}

func (s *SecurityListResourceCrud) buildConfICMPOptions(opts *baremetal.ICMPOptions) (list []interface{}) {
	confOpts := map[string]interface{}{
		"code": int(opts.Code),
		"type": int(opts.Type),
	}
	return []interface{}{confOpts}
}

func (s *SecurityListResourceCrud) buildConfTCPOptions(opts *baremetal.TCPOptions) (list []interface{}) {
	portRange := opts.DestinationPortRange
	confOpts := map[string]interface{}{
		"max": int(portRange.Max),
		"min": int(portRange.Min),
	}
	return []interface{}{confOpts}
}

func (s *SecurityListResourceCrud) buildConfUDPOptions(opts *baremetal.UDPOptions) (list []interface{}) {
	portRange := opts.DestinationPortRange
	confOpts := map[string]interface{}{
		"max": int(portRange.Max),
		"min": int(portRange.Min),
	}
	return []interface{}{confOpts}
}

func (s *SecurityListResourceCrud) buildEgressRules() (sdkRules []baremetal.EgressSecurityRule) {
	sdkRules = []baremetal.EgressSecurityRule{}
	for _, val := range s.D.Get("egress_security_rules").([]interface{}) {
		confRule := val.(map[string]interface{})

		sdkRule := baremetal.EgressSecurityRule{
			Destination: confRule["destination"].(string),
			ICMPOptions: s.buildICMPOptions(confRule),
			Protocol:    confRule["protocol"].(string),
			TCPOptions:  s.buildTCPOptions(confRule),
			UDPOptions:  s.buildUDPOptions(confRule),
		}

		sdkRules = append(sdkRules, sdkRule)
	}
	return
}

func (s *SecurityListResourceCrud) buildIngressRules() (sdkRules []baremetal.IngressSecurityRule) {
	sdkRules = []baremetal.IngressSecurityRule{}
	for _, val := range s.D.Get("ingress_security_rules").([]interface{}) {
		confRule := val.(map[string]interface{})

		sdkRule := baremetal.IngressSecurityRule{
			ICMPOptions: s.buildICMPOptions(confRule),
			Protocol:    confRule["protocol"].(string),
			Source:      confRule["source"].(string),
			TCPOptions:  s.buildTCPOptions(confRule),
			UDPOptions:  s.buildUDPOptions(confRule),
		}

		sdkRules = append(sdkRules, sdkRule)
	}
	return
}

func (s *SecurityListResourceCrud) buildICMPOptions(conf map[string]interface{}) (opts *baremetal.ICMPOptions) {
	l := conf["icmp_options"].([]interface{})
	if len(l) > 0 {
		confOpts := l[0].(map[string]interface{})
		opts = &baremetal.ICMPOptions{
			Code: uint64(confOpts["code"].(int)),
			Type: uint64(confOpts["type"].(int)),
		}
	}
	return
}

func (s *SecurityListResourceCrud) buildTCPOptions(conf map[string]interface{}) (opts *baremetal.TCPOptions) {
	l := conf["tcp_options"].([]interface{})
	if len(l) > 0 {
		confOpts := l[0].(map[string]interface{})
		opts = &baremetal.TCPOptions{
			baremetal.PortRange{
				Max: uint64(confOpts["max"].(int)),
				Min: uint64(confOpts["min"].(int)),
			},
		}
	}
	return
}

func (s *SecurityListResourceCrud) buildUDPOptions(conf map[string]interface{}) (opts *baremetal.UDPOptions) {
	l := conf["udp_options"].([]interface{})
	if len(l) > 0 {
		confOpts := l[0].(map[string]interface{})
		opts = &baremetal.UDPOptions{
			baremetal.PortRange{
				Max: uint64(confOpts["max"].(int)),
				Min: uint64(confOpts["min"].(int)),
			},
		}
	}
	return
}
