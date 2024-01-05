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

func OsmanagementManagedInstanceModuleStreamsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOsmanagementManagedInstanceModuleStreams,
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
			"stream_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"stream_status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"module_stream_on_managed_instances": {
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
						"profiles": {
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
						"software_source_id": {
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

func readOsmanagementManagedInstanceModuleStreams(d *schema.ResourceData, m interface{}) error {
	sync := &OsmanagementManagedInstanceModuleStreamsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OsManagementClient()

	return tfresource.ReadResource(sync)
}

type OsmanagementManagedInstanceModuleStreamsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_osmanagement.OsManagementClient
	Res    *oci_osmanagement.ListModuleStreamsOnManagedInstanceResponse
}

func (s *OsmanagementManagedInstanceModuleStreamsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OsmanagementManagedInstanceModuleStreamsDataSourceCrud) Get() error {
	request := oci_osmanagement.ListModuleStreamsOnManagedInstanceRequest{}

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

	if streamName, ok := s.D.GetOkExists("stream_name"); ok {
		tmp := streamName.(string)
		request.StreamName = &tmp
	}

	if streamStatus, ok := s.D.GetOkExists("stream_status"); ok {
		request.StreamStatus = oci_osmanagement.ListModuleStreamsOnManagedInstanceStreamStatusEnum(streamStatus.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "osmanagement")

	response, err := s.Client.ListModuleStreamsOnManagedInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListModuleStreamsOnManagedInstance(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OsmanagementManagedInstanceModuleStreamsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OsmanagementManagedInstanceModuleStreamsDataSource-", OsmanagementManagedInstanceModuleStreamsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		managedInstanceModuleStream := map[string]interface{}{}

		if r.ModuleName != nil {
			managedInstanceModuleStream["module_name"] = *r.ModuleName
		}

		profiles := []interface{}{}
		for _, item := range r.Profiles {
			profiles = append(profiles, ModuleStreamProfileOnManagedInstanceSummaryToMap(item))
		}
		managedInstanceModuleStream["profiles"] = profiles

		if r.SoftwareSourceId != nil {
			managedInstanceModuleStream["software_source_id"] = *r.SoftwareSourceId
		}

		managedInstanceModuleStream["status"] = r.Status

		if r.StreamName != nil {
			managedInstanceModuleStream["stream_name"] = *r.StreamName
		}

		if r.TimeModified != nil {
			managedInstanceModuleStream["time_modified"] = r.TimeModified.String()
		}

		resources = append(resources, managedInstanceModuleStream)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, OsmanagementManagedInstanceModuleStreamsDataSource().Schema["module_stream_on_managed_instances"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("module_stream_on_managed_instances", resources); err != nil {
		return err
	}

	return nil
}

func ModuleStreamProfileOnManagedInstanceSummaryToMap(obj oci_osmanagement.ModuleStreamProfileOnManagedInstanceSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ModuleName != nil {
		result["module_name"] = string(*obj.ModuleName)
	}

	if obj.ProfileName != nil {
		result["profile_name"] = string(*obj.ProfileName)
	}

	result["status"] = string(obj.Status)

	if obj.StreamName != nil {
		result["stream_name"] = string(*obj.StreamName)
	}

	if obj.TimeModified != nil {
		result["time_modified"] = obj.TimeModified.String()
	}

	return result
}
