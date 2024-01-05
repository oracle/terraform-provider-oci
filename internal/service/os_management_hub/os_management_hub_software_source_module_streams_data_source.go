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

func OsManagementHubSoftwareSourceModuleStreamsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOsManagementHubSoftwareSourceModuleStreams,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"is_latest": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"module_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"module_name_contains": {
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
			"module_stream_collection": {
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
									"profiles": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"software_source_id": {
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

func readOsManagementHubSoftwareSourceModuleStreams(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubSoftwareSourceModuleStreamsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SoftwareSourceClient()

	return tfresource.ReadResource(sync)
}

type OsManagementHubSoftwareSourceModuleStreamsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_os_management_hub.SoftwareSourceClient
	Res    *oci_os_management_hub.ListModuleStreamsResponse
}

func (s *OsManagementHubSoftwareSourceModuleStreamsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OsManagementHubSoftwareSourceModuleStreamsDataSourceCrud) Get() error {
	request := oci_os_management_hub.ListModuleStreamsRequest{}

	if isLatest, ok := s.D.GetOkExists("is_latest"); ok {
		tmp := isLatest.(bool)
		request.IsLatest = &tmp
	}

	if moduleName, ok := s.D.GetOkExists("module_name"); ok {
		tmp := moduleName.(string)
		request.ModuleName = &tmp
	}

	if moduleNameContains, ok := s.D.GetOkExists("module_name_contains"); ok {
		tmp := moduleNameContains.(string)
		request.ModuleNameContains = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if softwareSourceId, ok := s.D.GetOkExists("software_source_id"); ok {
		tmp := softwareSourceId.(string)
		request.SoftwareSourceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "os_management_hub")

	response, err := s.Client.ListModuleStreams(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListModuleStreams(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OsManagementHubSoftwareSourceModuleStreamsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OsManagementHubSoftwareSourceModuleStreamsDataSource-", OsManagementHubSoftwareSourceModuleStreamsDataSource(), s.D))
	resources := []map[string]interface{}{}
	softwareSourceModuleStream := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ModuleStreamSummaryToMap(item))
	}
	softwareSourceModuleStream["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, OsManagementHubSoftwareSourceModuleStreamsDataSource().Schema["module_stream_collection"].Elem.(*schema.Resource).Schema)
		softwareSourceModuleStream["items"] = items
	}

	resources = append(resources, softwareSourceModuleStream)
	if err := s.D.Set("module_stream_collection", resources); err != nil {
		return err
	}

	return nil
}

func ModuleStreamSummaryToMap(obj oci_os_management_hub.ModuleStreamSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsLatest != nil {
		result["is_latest"] = bool(*obj.IsLatest)
	}

	if obj.ModuleName != nil {
		result["module_name"] = string(*obj.ModuleName)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	result["profiles"] = obj.Profiles

	if obj.SoftwareSourceId != nil {
		result["software_source_id"] = string(*obj.SoftwareSourceId)
	}

	return result
}
