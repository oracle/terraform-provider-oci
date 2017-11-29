// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/oracle/bmcs-go-sdk"

	"github.com/oracle/terraform-provider-oci/options"

	"github.com/oracle/terraform-provider-oci/crud"
)

func SecurityListDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readSecurityLists,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"limit": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"page": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcn_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"security_lists": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     SecurityListResource(),
			},
		},
	}
}

func readSecurityLists(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*OracleClients)
	sync := &SecurityListDatasourceCrud{}
	sync.D = d
	sync.Client = client.client
	return crud.ReadResource(sync)
}

type SecurityListDatasourceCrud struct {
	crud.BaseCrud
	Res *baremetal.ListSecurityLists
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
	if s.Res == nil {
		return
	}

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
				&egressRule.IsStateless,
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
				nil,
			)
			confIngressRules = append(confIngressRules, confIngressRule)
		}
		res["ingress_security_rules"] = confIngressRules

		resources = append(resources, res)
	}

	if f, fOk := s.D.GetOk("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources)
	}

	if err := s.D.Set("security_lists", resources); err != nil {
		panic(err)
	}

	return
}
