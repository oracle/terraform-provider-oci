// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-baremetal/crud"
)

func UserDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readUsers,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"users": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     UserResource(),
			},
		},
	}
}

func readUsers(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*baremetal.Client)
	sync := &UserDatasourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.ReadResource(sync)
}

type UserDatasourceCrud struct {
	crud.BaseCrud
	Res *baremetal.ListUsers
}

func (s *UserDatasourceCrud) Get() (e error) {
	s.Res, e = s.Client.ListUsers(nil)
	return
}

func (s *UserDatasourceCrud) SetData() {
	if s.Res != nil {
		s.D.SetId(time.Now().UTC().String())
		resources := []map[string]interface{}{}
		for _, v := range s.Res.Users {
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
		if err := s.D.Set("users", resources); err != nil {
			panic(err)
		}
	}
	return
}
