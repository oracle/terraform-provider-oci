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

func OsManagementHubSoftwareSourceModuleStreamDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularOsManagementHubSoftwareSourceModuleStream,
		Schema: map[string]*schema.Schema{
			"module_name": {
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
			"arch_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_default": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_latest": {
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
			"profiles": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func readSingularOsManagementHubSoftwareSourceModuleStream(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubSoftwareSourceModuleStreamDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SoftwareSourceClient()

	return tfresource.ReadResource(sync)
}

type OsManagementHubSoftwareSourceModuleStreamDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_os_management_hub.SoftwareSourceClient
	Res    *oci_os_management_hub.GetModuleStreamResponse
}

func (s *OsManagementHubSoftwareSourceModuleStreamDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OsManagementHubSoftwareSourceModuleStreamDataSourceCrud) Get() error {
	request := oci_os_management_hub.GetModuleStreamRequest{}

	if moduleName, ok := s.D.GetOkExists("module_name"); ok {
		tmp := moduleName.(string)
		request.ModuleName = &tmp
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

	response, err := s.Client.GetModuleStream(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OsManagementHubSoftwareSourceModuleStreamDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OsManagementHubSoftwareSourceModuleStreamDataSource-", OsManagementHubSoftwareSourceModuleStreamDataSource(), s.D))

	s.D.Set("arch_type", s.Res.ArchType)

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.IsDefault != nil {
		s.D.Set("is_default", *s.Res.IsDefault)
	}

	if s.Res.IsLatest != nil {
		s.D.Set("is_latest", *s.Res.IsLatest)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("packages", s.Res.Packages)

	s.D.Set("profiles", s.Res.Profiles)

	return nil
}
