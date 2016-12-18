package core

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/options"
	"github.com/hashicorp/terraform/helper/schema"
)

type SecurityListDatasourceCrud struct {
	D      *schema.ResourceData
	Client client.BareMetalClient
	Res    *baremetal.ListSecurityLists
}

func (s *SecurityListDatasourceCrud) Get() (e error) {
	compartmentID := s.D.Get("compartment_id").(string)
	vcnID := s.D.Get("vcn_id").(string)

	opts := &baremetal.ListOptions{}
	options.SetListOptions(s.D, opts)

	s.Res = &baremetal.ListSecurityLists{SecurityLists: []baremetal.SecurityList{}}

	for {
		var list *baremetal.ListSecurityLists
		if list, e = s.Client.ListSecurityLists(compartmentID, vcnID, opts); e != nil {
			break
		}

		s.Res.SecurityLists = append(s.Res.SecurityLists, list.SecurityLists...)

		if hasNextPage := options.SetNextPageOption(list.NextPage, &opts.PageListOptions); !hasNextPage {
			break
		}
	}

	return
}

func (s *SecurityListDatasourceCrud) SetData() {
	if s.Res != nil {
		s.D.SetId(time.Now().UTC().String())
		resources := []map[string]interface{}{}
		for _, v := range s.Res.SecurityLists {

			res := map[string]interface{}{
				"compartment_id": v.CompartmentID,
				"display_name":   v.DisplayName,
				"id":             v.ID,
				"state":          v.State,
				"time_created":   v.TimeCreated.String(),
				"vcn_id":         v.VcnID,
			}

			confEgressRules := []map[string]interface{}{}
			for _, egressRule := range v.EgressSecurityRules {
				confEgressRule := map[string]interface{}{}
				confEgressRule["destination"] = egressRule.Destination
				confEgressRule = buildConfRule(
					confEgressRule,
					egressRule.Protocol,
					egressRule.ICMPOptions,
					egressRule.TCPOptions,
					egressRule.UDPOptions,
				)
				confEgressRules = append(confEgressRules, confEgressRule)
			}
			res["egress_security_rules"] = confEgressRules

			confIngressRules := []map[string]interface{}{}
			for _, ingressRule := range v.IngressSecurityRules {
				confIngressRule := map[string]interface{}{}
				confIngressRule["source"] = ingressRule.Source
				confIngressRule = buildConfRule(
					confIngressRule,
					ingressRule.Protocol,
					ingressRule.ICMPOptions,
					ingressRule.TCPOptions,
					ingressRule.UDPOptions,
				)
				confIngressRules = append(confIngressRules, confIngressRule)
			}
			res["ingress_security_rules"] = confIngressRules

			resources = append(resources, res)
		}
		s.D.Set("security_lists", resources)
	}
	return
}
