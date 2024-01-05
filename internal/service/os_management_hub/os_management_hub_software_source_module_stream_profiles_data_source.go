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

func OsManagementHubSoftwareSourceModuleStreamProfilesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOsManagementHubSoftwareSourceModuleStreamProfiles,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"module_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
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
			"module_stream_profile_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_default": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"module_name": {
										Type:     schema.TypeString,
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
									"stream_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func readOsManagementHubSoftwareSourceModuleStreamProfiles(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubSoftwareSourceModuleStreamProfilesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SoftwareSourceClient()

	return tfresource.ReadResource(sync)
}

type OsManagementHubSoftwareSourceModuleStreamProfilesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_os_management_hub.SoftwareSourceClient
	Res    *oci_os_management_hub.ListModuleStreamProfilesResponse
}

func (s *OsManagementHubSoftwareSourceModuleStreamProfilesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OsManagementHubSoftwareSourceModuleStreamProfilesDataSourceCrud) Get() error {
	request := oci_os_management_hub.ListModuleStreamProfilesRequest{}

	if moduleName, ok := s.D.GetOkExists("module_name"); ok {
		tmp := moduleName.(string)
		request.ModuleName = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
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

func (s *OsManagementHubSoftwareSourceModuleStreamProfilesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OsManagementHubSoftwareSourceModuleStreamProfilesDataSource-", OsManagementHubSoftwareSourceModuleStreamProfilesDataSource(), s.D))
	resources := []map[string]interface{}{}
	softwareSourceModuleStreamProfile := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ModuleStreamProfileSummaryToMap(item))
	}
	softwareSourceModuleStreamProfile["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, OsManagementHubSoftwareSourceModuleStreamProfilesDataSource().Schema["module_stream_profile_collection"].Elem.(*schema.Resource).Schema)
		softwareSourceModuleStreamProfile["items"] = items
	}

	resources = append(resources, softwareSourceModuleStreamProfile)
	if err := s.D.Set("module_stream_profile_collection", resources); err != nil {
		return err
	}

	return nil
}

func ModuleStreamProfileSummaryToMap(obj oci_os_management_hub.ModuleStreamProfileSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsDefault != nil {
		result["is_default"] = bool(*obj.IsDefault)
	}

	if obj.ModuleName != nil {
		result["module_name"] = string(*obj.ModuleName)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.StreamName != nil {
		result["stream_name"] = string(*obj.StreamName)
	}

	return result
}
