// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-baremetal/client"
	"github.com/oracle/terraform-provider-baremetal/crud"
	"github.com/oracle/terraform-provider-baremetal/options"
)

func InstanceDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readInstances,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"availability_domain": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"page": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"instances": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     resourceCoreInstance(),
			},
		},
	}
}

func resourceCoreInstance() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"availability_domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"image": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ipxe_script": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"metadata": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"region": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"shape": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readInstances(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	reader := &InstanceDatasourceCrud{}
	reader.D = d
	reader.Client = client

	return crud.ReadResource(reader)
}

type InstanceDatasourceCrud struct {
	crud.BaseCrud
	Res *baremetal.ListInstances
}

func (s *InstanceDatasourceCrud) Get() (e error) {
	compartmentID := s.D.Get("compartment_id").(string)

	opts := &baremetal.ListInstancesOptions{}
	options.SetListOptions(s.D, &opts.ListOptions)
	if val, ok := s.D.GetOk("availability_domain"); ok {
		opts.AvailabilityDomain = val.(string)
	}

	s.Res = &baremetal.ListInstances{Instances: []baremetal.Instance{}}

	for {
		var list *baremetal.ListInstances
		if list, e = s.Client.ListInstances(compartmentID, opts); e != nil {
			break
		}

		s.Res.Instances = append(s.Res.Instances, list.Instances...)

		if hasNextPage := options.SetNextPageOption(list.NextPage, &opts.ListOptions.PageListOptions); !hasNextPage {
			break
		}
	}

	return
}

func (s *InstanceDatasourceCrud) SetData() {
	if s.Res != nil {
		// Important, if you don't have an ID, make one up for your datasource
		// or things will end in tears
		s.D.SetId(time.Now().UTC().String())
		resources := []map[string]interface{}{}
		for _, v := range s.Res.Instances {
			res := map[string]interface{}{
				"availability_domain": v.AvailabilityDomain,
				"compartment_id":      v.CompartmentID,
				"display_name":        v.DisplayName,
				"id":                  v.ID,
				"image":               v.ImageID,
				"ipxe_script":         v.IpxeScript,
				"metadata":            v.Metadata,
				"region":              v.Region,
				"shape":               v.Shape,
				"state":               v.State,
				"time_created":        v.TimeCreated.String(),
			}
			resources = append(resources, res)
		}
		s.D.Set("instances", resources)
	}
	return
}
