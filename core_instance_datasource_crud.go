// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"

	"github.com/oracle/terraform-provider-baremetal/options"

	"github.com/oracle/terraform-provider-baremetal/crud"
)

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
