// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v58/core"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func CoreInstanceCredentialDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularCoreInstanceCredential,
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"password": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"username": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularCoreInstanceCredential(d *schema.ResourceData, m interface{}) error {
	sync := &CoreInstanceCredentialDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.ReadResource(sync)
}

type CoreInstanceCredentialDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.ComputeClient
	Res    *oci_core.GetWindowsInstanceInitialCredentialsResponse
}

func (s *CoreInstanceCredentialDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreInstanceCredentialDataSourceCrud) Get() error {
	request := oci_core.GetWindowsInstanceInitialCredentialsRequest{}

	if instanceId, ok := s.D.GetOkExists("instance_id"); ok {
		tmp := instanceId.(string)
		request.InstanceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.GetWindowsInstanceInitialCredentials(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CoreInstanceCredentialDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreInstanceCredentialDataSource-", CoreInstanceCredentialDataSource(), s.D))

	if s.Res.Password != nil {
		s.D.Set("password", *s.Res.Password)
	}

	if s.Res.Username != nil {
		s.D.Set("username", *s.Res.Username)
	}

	return nil
}
