// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/core"

	"github.com/oracle/terraform-provider-oci/crud"
)

func InstanceCredentialsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readInstanceCredentials,
		Schema: map[string]*schema.Schema{
			// InstanceCredentials is a single-value data source.
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

func readInstanceCredentials(d *schema.ResourceData, m interface{}) error {
	sync := &InstanceCredentialsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient

	return crud.ReadResource(sync)
}

type InstanceCredentialsDataSourceCrud struct {
	crud.BaseCrud
	Client *oci_core.ComputeClient
	Res    *oci_core.GetWindowsInstanceInitialCredentialsResponse
}

func (s *InstanceCredentialsDataSourceCrud) Get() error {
	request := oci_core.GetWindowsInstanceInitialCredentialsRequest{}

	if instanceId, ok := s.D.GetOkExists("instance_id"); ok {
		tmp := instanceId.(string)
		request.InstanceId = &tmp
	}

	response, err := s.Client.GetWindowsInstanceInitialCredentials(context.Background(), request, getRetryOptions(false, "core")...)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *InstanceCredentialsDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())

	if s.Res.Password != nil {
		s.D.Set("password", *s.Res.Password)
	}

	if s.Res.Username != nil {
		s.D.Set("username", *s.Res.Username)
	}

	return
}
