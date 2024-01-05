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

func OsmanagementManagedInstanceStreamProfilesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOsmanagementManagedInstanceStreamProfiles,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"managed_instance_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"module_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"profile_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"profile_status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"stream_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"module_stream_profile_on_managed_instances": {
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
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"stream_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_modified": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readOsmanagementManagedInstanceStreamProfiles(d *schema.ResourceData, m interface{}) error {
	sync := &OsmanagementManagedInstanceStreamProfilesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OsManagementClient()

	return tfresource.ReadResource(sync)
}

type OsmanagementManagedInstanceStreamProfilesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_osmanagement.OsManagementClient
	Res    *oci_osmanagement.ListModuleStreamProfilesOnManagedInstanceResponse
}

func (s *OsmanagementManagedInstanceStreamProfilesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OsmanagementManagedInstanceStreamProfilesDataSourceCrud) Get() error {
	request := oci_osmanagement.ListModuleStreamProfilesOnManagedInstanceRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if managedInstanceId, ok := s.D.GetOkExists("managed_instance_id"); ok {
		tmp := managedInstanceId.(string)
		request.ManagedInstanceId = &tmp
	}

	if moduleName, ok := s.D.GetOkExists("module_name"); ok {
		tmp := moduleName.(string)
		request.ModuleName = &tmp
	}

	if profileName, ok := s.D.GetOkExists("profile_name"); ok {
		tmp := profileName.(string)
		request.ProfileName = &tmp
	}

	if profileStatus, ok := s.D.GetOkExists("profile_status"); ok {
		request.ProfileStatus = oci_osmanagement.ListModuleStreamProfilesOnManagedInstanceProfileStatusEnum(profileStatus.(string))
	}

	if streamName, ok := s.D.GetOkExists("stream_name"); ok {
		tmp := streamName.(string)
		request.StreamName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "osmanagement")

	response, err := s.Client.ListModuleStreamProfilesOnManagedInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListModuleStreamProfilesOnManagedInstance(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OsmanagementManagedInstanceStreamProfilesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OsmanagementManagedInstanceStreamProfilesDataSource-", OsmanagementManagedInstanceStreamProfilesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		managedInstanceStreamProfile := map[string]interface{}{}

		if r.ModuleName != nil {
			managedInstanceStreamProfile["module_name"] = *r.ModuleName
		}

		if r.ProfileName != nil {
			managedInstanceStreamProfile["profile_name"] = *r.ProfileName
		}

		managedInstanceStreamProfile["status"] = r.Status

		if r.StreamName != nil {
			managedInstanceStreamProfile["stream_name"] = *r.StreamName
		}

		if r.TimeModified != nil {
			managedInstanceStreamProfile["time_modified"] = r.TimeModified.String()
		}

		resources = append(resources, managedInstanceStreamProfile)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, OsmanagementManagedInstanceStreamProfilesDataSource().Schema["module_stream_profile_on_managed_instances"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("module_stream_profile_on_managed_instances", resources); err != nil {
		return err
	}

	return nil
}
