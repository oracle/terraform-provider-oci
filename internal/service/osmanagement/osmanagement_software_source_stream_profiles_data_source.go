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

func OsmanagementSoftwareSourceStreamProfilesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOsmanagementSoftwareSourceStreamProfiles,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"module_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"profile_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"software_source_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"stream_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"module_stream_profiles": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"module_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"profile_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"stream_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readOsmanagementSoftwareSourceStreamProfiles(d *schema.ResourceData, m interface{}) error {
	sync := &OsmanagementSoftwareSourceStreamProfilesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OsManagementClient()

	return tfresource.ReadResource(sync)
}

type OsmanagementSoftwareSourceStreamProfilesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_osmanagement.OsManagementClient
	Res    *oci_osmanagement.ListModuleStreamProfilesResponse
}

func (s *OsmanagementSoftwareSourceStreamProfilesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OsmanagementSoftwareSourceStreamProfilesDataSourceCrud) Get() error {
	request := oci_osmanagement.ListModuleStreamProfilesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

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

	response, err := s.Client.ListModuleStreamProfiles(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListModuleStreamProfiles(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OsmanagementSoftwareSourceStreamProfilesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OsmanagementSoftwareSourceStreamProfilesDataSource-", OsmanagementSoftwareSourceStreamProfilesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		softwareSourceStreamProfile := map[string]interface{}{}

		if r.ModuleName != nil {
			softwareSourceStreamProfile["module_name"] = *r.ModuleName
		}

		if r.ProfileName != nil {
			softwareSourceStreamProfile["profile_name"] = *r.ProfileName
		}

		if r.StreamName != nil {
			softwareSourceStreamProfile["stream_name"] = *r.StreamName
		}

		resources = append(resources, softwareSourceStreamProfile)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, OsmanagementSoftwareSourceStreamProfilesDataSource().Schema["module_stream_profiles"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("module_stream_profiles", resources); err != nil {
		return err
	}

	return nil
}
