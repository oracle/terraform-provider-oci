// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/oracle/bmcs-go-sdk"

	"github.com/oracle/terraform-provider-oci/options"

	"github.com/oracle/terraform-provider-oci/crud"
)

func PrivateIPDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readPrivateIPs,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"ip_address": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vnic_id": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"private_ips": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     PrivateIPResource(),
			},
		},
	}
}

func readPrivateIPs(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*OracleClients)
	sync := &PrivateIPDatasourceCrud{}
	sync.D = d
	sync.Client = client.client
	return crud.ReadResource(sync)
}

type PrivateIPDatasourceCrud struct {
	crud.BaseCrud
	Res *baremetal.ListPrivateIPs
}

func (s *PrivateIPDatasourceCrud) Get() (e error) {

	opts := &baremetal.ListPrivateIPsOptions{}
	options.SetListOptions(s.D, &opts.ListOptions)
	if val, ok := s.D.GetOk("ip_address"); ok {
		opts.IPAddress = val.(string)
	}
	if val, ok := s.D.GetOk("subnet_id"); ok {
		opts.SubnetID = val.(string)
	}
	if val, ok := s.D.GetOk("vnic_id"); ok {
		opts.VnicID = val.(string)
	}

	s.Res = &baremetal.ListPrivateIPs{PrivateIPs: []baremetal.PrivateIP{}}

	for {
		var list *baremetal.ListPrivateIPs
		if list, e = s.Client.ListPrivateIPs(opts); e != nil {
			break
		}

		s.Res.PrivateIPs = append(s.Res.PrivateIPs, list.PrivateIPs...)

		if hasNextPage := options.SetNextPageOption(list.NextPage, &opts.PageListOptions); !hasNextPage {
			break
		}
	}
	return
}

func (s *PrivateIPDatasourceCrud) SetData() {
	if s.Res == nil {
		return
	}
	s.D.SetId(time.Now().UTC().String())
	resources := []map[string]interface{}{}
	for _, r := range s.Res.PrivateIPs {
		res := map[string]interface{}{
			"availability_domain": r.AvailabilityDomain,
			"compartment_id":      r.CompartmentID,
			"display_name":        r.DisplayName,
			"hostname_label":      r.HostnameLabel,
			"id":                  r.ID,
			"ip_address":          r.IPAddress,
			"is_primary":          r.IsPrimary,
			"subnet_id":           r.SubnetID,
			"time_created":        r.TimeCreated.String(),
			"vnic_id":             r.VnicID,
		}
		resources = append(resources, res)
	}
	if f, fOk := s.D.GetOk("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources)
	}

	if err := s.D.Set("private_ips", resources); err != nil {
		panic(err)
	}

	return
}
