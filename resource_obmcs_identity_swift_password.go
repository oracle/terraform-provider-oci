// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-baremetal/crud"
)

func SwiftPasswordResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createSwiftPassword,
		Read:     readSwiftPassword,
		Update:   updateSwiftPassword,
		Delete:   deleteSwiftPassword,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"inactive_state": {
				Type:     schema.TypeInt,
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
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"expires_on": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createSwiftPassword(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*baremetal.Client)
	sync := &SwiftPasswordResourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.CreateResource(d, sync)
}

func readSwiftPassword(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*baremetal.Client)
	sync := &SwiftPasswordResourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.ReadResource(sync)
}

func updateSwiftPassword(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*baremetal.Client)
	sync := &SwiftPasswordResourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.UpdateResource(d, sync)
}

func deleteSwiftPassword(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*baremetal.Client)
	sync := &SwiftPasswordResourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.DeleteResource(d, sync)
}

type SwiftPasswordResourceCrud struct {
	crud.BaseCrud
	Res *baremetal.SwiftPassword
}

func (s *SwiftPasswordResourceCrud) ID() string {
	return s.Res.ID
}

func (s *SwiftPasswordResourceCrud) Get() (e error) {
	// There is no get resource for swift passwords, so we list them all and match
	id := s.D.Id()
	userID := s.D.Get("user_id").(string)
	var list *baremetal.ListSwiftPasswords
	list, e = s.Client.ListSwiftPasswords(userID)
	if e != nil {
		return
	}
	for _, sp := range list.SwiftPasswords {
		if sp.ID == id {
			s.Res = &sp
		}
	}
	return
}

func (s *SwiftPasswordResourceCrud) Create() (e error) {
	userID := s.D.Get("user_id").(string)
	desc := s.D.Get("description").(string)
	s.Res, e = s.Client.CreateSwiftPassword(userID, desc, nil)
	return
}

func (s *SwiftPasswordResourceCrud) Update() (e error) {
	userID := s.D.Get("user_id").(string)
	opts := &baremetal.UpdateIdentityOptions{}
	if description, ok := s.D.GetOk("description"); ok {
		opts.Description = description.(string)
	}

	s.Res, e = s.Client.UpdateSwiftPassword(s.D.Id(), userID, opts)
	return
}

func (s *SwiftPasswordResourceCrud) Delete() (e error) {
	userID := s.D.Get("user_id").(string)
	return s.Client.DeleteSwiftPassword(s.D.Id(), userID, nil)
}

func (s *SwiftPasswordResourceCrud) SetData() {
	s.D.Set("inactive_status", s.Res.InactiveStatus)
	s.D.Set("state", s.Res.State)
	s.D.Set("password", s.Res.Password)
	s.D.Set("time_created", s.Res.TimeCreated.String())
	s.D.Set("user_id", s.Res.UserID)
	s.D.Set("description", s.Res.Description)
	s.D.Set("expires_on", s.Res.ExpiresOn.String())
}
