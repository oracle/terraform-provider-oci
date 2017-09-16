// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"
)

func GroupDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readGroups,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"groups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GroupResource(),
			},
		},
	}
}

func readGroups(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*baremetal.Client)
	sync := &GroupDatasourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.ReadResource(sync)
}

type GroupDatasourceCrud struct {
	crud.BaseCrud
	Res *baremetal.ListGroups
}

func (s *GroupDatasourceCrud) Get() (e error) {
	s.Res, e = s.Client.ListGroups(nil)
	return
}

func (s *GroupDatasourceCrud) SetData() {
	if s.Res != nil {
		s.D.SetId(time.Now().UTC().String())
		resources := []map[string]interface{}{}
		for _, v := range s.Res.Groups {
			res := map[string]interface{}{
				"compartment_id": v.CompartmentID,
				"description":    v.Description,
				"id":             v.ID,
				"inactive_state": v.InactiveStatus,
				"name":           v.Name,
				"state":          v.State,
				"time_created":   v.TimeCreated.String(),
			}
			resources = append(resources, res)
		}
		if err := s.D.Set("groups", resources); err != nil {
			panic(err)
		}
	}
	return
}
