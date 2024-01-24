// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package os_management_hub

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_os_management_hub "github.com/oracle/oci-go-sdk/v65/osmanagementhub"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OsManagementHubSoftwareSourceModuleStreamProfileDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularOsManagementHubSoftwareSourceModuleStreamProfile,
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
			"name": {
				Type:     schema.TypeString,
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

func readSingularOsManagementHubSoftwareSourceModuleStreamProfile(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubSoftwareSourceModuleStreamProfileDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SoftwareSourceClient()

	return tfresource.ReadResource(sync)
}

type OsManagementHubSoftwareSourceModuleStreamProfileDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_os_management_hub.SoftwareSourceClient
	Res    *oci_os_management_hub.GetModuleStreamProfileResponse
}

func (s *OsManagementHubSoftwareSourceModuleStreamProfileDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OsManagementHubSoftwareSourceModuleStreamProfileDataSourceCrud) Get() error {
	request := oci_os_management_hub.GetModuleStreamProfileRequest{}

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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "os_management_hub")

	response, err := s.Client.GetModuleStreamProfile(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OsManagementHubSoftwareSourceModuleStreamProfileDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OsManagementHubSoftwareSourceModuleStreamProfileDataSource-", OsManagementHubSoftwareSourceModuleStreamProfileDataSource(), s.D))

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.IsDefault != nil {
		s.D.Set("is_default", *s.Res.IsDefault)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("packages", s.Res.Packages)

	return nil
}
