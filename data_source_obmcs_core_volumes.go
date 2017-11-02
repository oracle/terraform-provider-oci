// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/oracle/bmcs-go-sdk"

	"github.com/oracle/terraform-provider-oci/options"

	"github.com/oracle/terraform-provider-oci/crud"
)

func VolumeDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readVolumes,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"availability_domain": {
				Type:     schema.TypeString,
				Optional: true,
			},
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
			"volumes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     VolumeResource(),
			},
		},
	}
}

func readVolumes(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*OracleClients)
	sync := &VolumeDatasourceCrud{}
	sync.D = d
	sync.Client = client.client
	return crud.ReadResource(sync)
}

type VolumeDatasourceCrud struct {
	crud.BaseCrud
	Res *baremetal.ListVolumes
}

func (s *VolumeDatasourceCrud) Get() (e error) {
	compartmentID := s.D.Get("compartment_id").(string)

	opts := &baremetal.ListVolumesOptions{}
	options.SetListOptions(s.D, &opts.ListOptions)

	if val, ok := s.D.GetOk("availability_domain"); ok {
		opts.AvailabilityDomain = val.(string)
	}

	s.Res = &baremetal.ListVolumes{Volumes: []baremetal.Volume{}}

	for {
		var list *baremetal.ListVolumes
		if list, e = s.Client.ListVolumes(compartmentID, opts); e != nil {
			break
		}

		s.Res.Volumes = append(s.Res.Volumes, list.Volumes...)

		if hasNextPage := options.SetNextPageOption(list.NextPage, &opts.ListOptions.PageListOptions); !hasNextPage {
			break
		}
	}

	return
}

func (s *VolumeDatasourceCrud) SetData() {
	if s.Res == nil {
		return
	}
	// Important, if you don't have an ID, make one up for your datasource
	// or things will end in tears
	s.D.SetId(time.Now().UTC().String())
	resources := []map[string]interface{}{}
	for _, v := range s.Res.Volumes {
		vol := map[string]interface{}{
			"availability_domain": v.AvailabilityDomain,
			"compartment_id":      v.CompartmentID,
			"display_name":        v.DisplayName,
			"id":                  v.ID,
			"size_in_mbs":         v.SizeInMBs,
			"size_in_gbs":         v.SizeInGBs,
			"state":               v.State,
			"time_created":        v.TimeCreated.String(),
		}

		if vsdRaw := v.VolumeSourceDetails; vsdRaw != nil {
			vsd := make(map[string]interface{})
			vsd["id"] = vsdRaw.Id
			vsd["type"] = vsdRaw.Type
			vol["source_details"] = []interface{}{vsd}
		}

		resources = append(resources, vol)
	}

	if f, fOk := s.D.GetOk("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources)
	}

	if err := s.D.Set("volumes", resources); err != nil {
		panic(err)
	}

	return
}
