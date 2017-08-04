// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-baremetal/client"
	"github.com/oracle/terraform-provider-baremetal/crud"
)

func InstanceCredentialsDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readInstanceCredentials,
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"username": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"password": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readInstanceCredentials(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &InstanceCredentialsDatasourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.ReadResource(sync)
}

type InstanceCredentialsDatasourceCrud struct {
	crud.BaseCrud
	Res *baremetal.InstanceCredentials
}

func (s *InstanceCredentialsDatasourceCrud) Get() (e error) {
	instanceId := s.D.Get("instance_id").(string)
	s.Res, e = s.Client.GetWindowsInstanceInitialCredentials(instanceId)
	return
}

func (s *InstanceCredentialsDatasourceCrud) SetData() {
	if s.Res != nil {
		s.D.SetId(time.Now().UTC().String())
		s.D.Set("username", s.Res.Username)
		s.D.Set("password", s.Res.Password)
	}
	return
}
