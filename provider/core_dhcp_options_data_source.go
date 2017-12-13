// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/oracle/bmcs-go-sdk"

	"github.com/oracle/terraform-provider-oci/crud"
	"github.com/oracle/terraform-provider-oci/options"
)

func DHCPOptionsDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readDHCPOptionsList,
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
			"options": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     DHCPOptionsResource(),
			},
			"vcn_id": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func readDHCPOptionsList(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*OracleClients)
	reader := &DHCPOptionsDatasourceCrud{}
	reader.D = d
	reader.Client = client.client

	return crud.ReadResource(reader)
}

type DHCPOptionsDatasourceCrud struct {
	crud.BaseCrud
	Res *baremetal.ListDHCPOptions
}

func (s *DHCPOptionsDatasourceCrud) Get() (e error) {
	compartmentID := s.D.Get("compartment_id").(string)
	vcnID := s.D.Get("vcn_id").(string)

	opts := &baremetal.ListOptions{}
	options.SetListOptions(s.D, opts)

	s.Res = &baremetal.ListDHCPOptions{DHCPOptions: []baremetal.DHCPOptions{}}

	for {
		var list *baremetal.ListDHCPOptions
		if list, e = s.Client.ListDHCPOptions(compartmentID, vcnID, opts); e != nil {
			break
		}

		s.Res.DHCPOptions = append(s.Res.DHCPOptions, list.DHCPOptions...)

		if hasNextPage := options.SetNextPageOption(list.NextPage, &opts.PageListOptions); !hasNextPage {
			break
		}
	}

	return
}

func (s *DHCPOptionsDatasourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(time.Now().UTC().String())

	resources := []map[string]interface{}{}
	for _, res := range s.Res.DHCPOptions {

		nestedStateObjs := []map[string]interface{}{}

		for _, nestedRes := range res.Options {
			nestedStateObj := map[string]interface{}{
				"type":               nestedRes.Type,
				"custom_dns_servers": nestedRes.CustomDNSServers,
				"server_type":        nestedRes.ServerType,
			}
			nestedStateObjs = append(nestedStateObjs, nestedStateObj)
		}

		stateObj := map[string]interface{}{
			"compartment_id": res.CompartmentID,
			"display_name":   res.DisplayName,
			"id":             res.ID,
			"options":        nestedStateObjs,
			"state":          res.State,
			"time_created":   res.TimeCreated.String(),
			"vcn_id":         res.VcnID,
		}
		resources = append(resources, stateObj)
	}

	if f, fOk := s.D.GetOk("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources)
	}

	if err := s.D.Set("options", resources); err != nil {
		panic(err)
	}

	return
}
