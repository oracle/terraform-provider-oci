// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-baremetal/crud"
)

func SwiftPasswordDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readSwiftPasswords,
		Schema: map[string]*schema.Schema{
			"user_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"passwords": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     SwiftPasswordResource(),
			},
		},
	}
}

func readSwiftPasswords(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*baremetal.Client)
	sync := &SwiftPasswordDatasourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.ReadResource(sync)
}

type SwiftPasswordDatasourceCrud struct {
	crud.BaseCrud
	Res *baremetal.ListSwiftPasswords
}

func (s *SwiftPasswordDatasourceCrud) Get() (e error) {
	userID := s.D.Get("user_id").(string)

	s.Res, e = s.Client.ListSwiftPasswords(userID)
	return
}

func (s *SwiftPasswordDatasourceCrud) SetData() {
	if s.Res != nil {
		s.D.SetId(time.Now().UTC().String())
		resources := []map[string]interface{}{}
		for _, v := range s.Res.SwiftPasswords {
			res := map[string]interface{}{
				"id":             v.ID,
				"user_id":        v.UserID,
				"description":    v.Description,
				"state":          v.State,
				"inactive_state": v.InactiveStatus,
				"time_created":   v.TimeCreated.String(),
				"expires_on":     v.ExpiresOn.String(),
			}
			resources = append(resources, res)
		}
		if err := s.D.Set("passwords", resources); err != nil {
			panic(err)
		}
	}
	return
}
