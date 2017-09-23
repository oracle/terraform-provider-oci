// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"time"

	"github.com/oracle/bmcs-go-sdk"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"
	"github.com/oracle/terraform-provider-oci/options"
)

func UserGroupMembershipDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readUserGroupMemberships,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"user_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"group_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"memberships": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     UserGroupMembershipResource(),
			},
		},
	}
}

func readUserGroupMemberships(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*baremetal.Client)
	sync := &UserGroupMembershipDatasourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.ReadResource(sync)
}

type UserGroupMembershipDatasourceCrud struct {
	crud.BaseCrud
	Res *baremetal.ListUserGroupMemberships
}

func (s *UserGroupMembershipDatasourceCrud) Get() (e error) {
	opts := &baremetal.ListMembershipsOptions{}
	if id, ok := s.D.GetOk("group_id"); ok {
		opts.GroupID = id.(string)
	}
	if id, ok := s.D.GetOk("user_id"); ok {
		opts.UserID = id.(string)
	}

	s.Res = &baremetal.ListUserGroupMemberships{Memberships: []baremetal.UserGroupMembership{}}
	for {
		var list *baremetal.ListUserGroupMemberships
		if list, e = s.Client.ListUserGroupMemberships(opts); e != nil {
			break
		}

		s.Res.Memberships = append(s.Res.Memberships, list.Memberships...)

		if hasNexPage := options.SetNextPageOption(list.NextPage, &opts.PageListOptions); !hasNexPage {
			break
		}
	}
	return
}

func (s *UserGroupMembershipDatasourceCrud) SetData() {
	if s.Res != nil {
		s.D.SetId(time.Now().UTC().String())
		resources := []map[string]interface{}{}
		for _, v := range s.Res.Memberships {
			res := map[string]interface{}{
				"compartment_id": v.CompartmentID,
				"id":             v.ID,
				"user_id":        v.UserID,
				"group_id":       v.GroupID,
				"inactive_state": v.InactiveStatus,
				"state":          v.State,
				"time_created":   v.TimeCreated.String(),
			}
			resources = append(resources, res)
		}
		if err := s.D.Set("memberships", resources); err != nil {
			panic(err)
		}
	}
	return
}
