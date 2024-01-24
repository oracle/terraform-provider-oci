// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package osmanagement

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_osmanagement "github.com/oracle/oci-go-sdk/v65/osmanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OsmanagementSoftwareSourceModuleStreamProfileDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularOsmanagementSoftwareSourceModuleStreamProfile,
		Schema: map[string]*schema.Schema{
			"module_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"profile_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"software_source_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"stream_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_default": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"packages": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func readSingularOsmanagementSoftwareSourceModuleStreamProfile(d *schema.ResourceData, m interface{}) error {
	sync := &OsmanagementSoftwareSourceModuleStreamProfileDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OsManagementClient()

	return tfresource.ReadResource(sync)
}

type OsmanagementSoftwareSourceModuleStreamProfileDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_osmanagement.OsManagementClient
	Res    *oci_osmanagement.GetModuleStreamProfileResponse
}

func (s *OsmanagementSoftwareSourceModuleStreamProfileDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OsmanagementSoftwareSourceModuleStreamProfileDataSourceCrud) Get() error {
	request := oci_osmanagement.GetModuleStreamProfileRequest{}

	if moduleName, ok := s.D.GetOkExists("module_name"); ok {
		tmp := moduleName.(string)
		request.ModuleName = &tmp
	}

	if profileName, ok := s.D.GetOkExists("profile_name"); ok {
		tmp := profileName.(string)
		request.ProfileName = &tmp
	}

	if softwareSourceId, ok := s.D.GetOkExists("software_source_id"); ok {
		tmp := softwareSourceId.(string)
		request.SoftwareSourceId = &tmp
	}

	if streamName, ok := s.D.GetOkExists("stream_name"); ok {
		tmp := streamName.(string)
		request.StreamName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "osmanagement")

	response, err := s.Client.GetModuleStreamProfile(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OsmanagementSoftwareSourceModuleStreamProfileDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OsmanagementSoftwareSourceModuleStreamProfileDataSource-", OsmanagementSoftwareSourceModuleStreamProfileDataSource(), s.D))

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.IsDefault != nil {
		s.D.Set("is_default", *s.Res.IsDefault)
	}

	s.D.Set("packages", s.Res.Packages)

	return nil
}
