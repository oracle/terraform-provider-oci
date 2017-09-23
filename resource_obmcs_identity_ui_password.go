// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"github.com/oracle/bmcs-go-sdk"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"
)

// Version is exposed to allow resetting an existing user's password.
// Incrementing the value of version will cause a new UIPassword to be created.
func UIPasswordResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createUIPassword,
		Read:     readUIPassword,
		Delete:   deleteUIPassword,
		Schema: map[string]*schema.Schema{
			"inactive_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"password": {
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
			"user_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createUIPassword(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*baremetal.Client)
	sync := &UIPasswordResourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.CreateResource(d, sync)
}

func readUIPassword(d *schema.ResourceData, m interface{}) (e error) {
	return nil
}

func deleteUIPassword(d *schema.ResourceData, m interface{}) (e error) {
	return nil
}

type UIPasswordResourceCrud struct {
	crud.BaseCrud
	Res *baremetal.UIPassword
}

func (s *UIPasswordResourceCrud) ID() string {
	return s.D.Get("user_id").(string)
}

func (s *UIPasswordResourceCrud) Create() (e error) {
	userID := s.D.Get("user_id").(string)
	s.Res, e = s.Client.CreateOrResetUIPassword(userID, nil)
	return
}

func (s *UIPasswordResourceCrud) SetData() {
	s.D.Set("inactive_status", s.Res.InactiveStatus)
	s.D.Set("state", s.Res.State)
	s.D.Set("password", s.Res.Password)
	s.D.Set("time_created", s.Res.TimeCreated.String())
	s.D.Set("user_id", s.Res.UserID)
}
