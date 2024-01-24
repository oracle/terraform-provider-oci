// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database_management "github.com/oracle/oci-go-sdk/v65/databasemanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseManagementExternalAsmInstancesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseManagementExternalAsmInstances,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_asm_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_asm_instance_collection": {
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
									"adr_home_directory": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"compartment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"component_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"external_asm_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"external_db_node_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"external_db_system_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"host_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"lifecycle_details": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_created": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_updated": {
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

func readDatabaseManagementExternalAsmInstances(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalAsmInstancesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementExternalAsmInstancesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.ListExternalAsmInstancesResponse
}

func (s *DatabaseManagementExternalAsmInstancesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementExternalAsmInstancesDataSourceCrud) Get() error {
	request := oci_database_management.ListExternalAsmInstancesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if externalAsmId, ok := s.D.GetOkExists("external_asm_id"); ok {
		tmp := externalAsmId.(string)
		request.ExternalAsmId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.ListExternalAsmInstances(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListExternalAsmInstances(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseManagementExternalAsmInstancesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementExternalAsmInstancesDataSource-", DatabaseManagementExternalAsmInstancesDataSource(), s.D))
	resources := []map[string]interface{}{}
	externalAsmInstance := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ExternalAsmInstanceSummaryToMap(item))
	}
	externalAsmInstance["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseManagementExternalAsmInstancesDataSource().Schema["external_asm_instance_collection"].Elem.(*schema.Resource).Schema)
		externalAsmInstance["items"] = items
	}

	resources = append(resources, externalAsmInstance)
	if err := s.D.Set("external_asm_instance_collection", resources); err != nil {
		return err
	}

	return nil
}

func ExternalAsmInstanceSummaryToMap(obj oci_database_management.ExternalAsmInstanceSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AdrHomeDirectory != nil {
		result["adr_home_directory"] = string(*obj.AdrHomeDirectory)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.ComponentName != nil {
		result["component_name"] = string(*obj.ComponentName)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.ExternalAsmId != nil {
		result["external_asm_id"] = string(*obj.ExternalAsmId)
	}

	if obj.ExternalDbNodeId != nil {
		result["external_db_node_id"] = string(*obj.ExternalDbNodeId)
	}

	if obj.ExternalDbSystemId != nil {
		result["external_db_system_id"] = string(*obj.ExternalDbSystemId)
	}

	if obj.HostName != nil {
		result["host_name"] = string(*obj.HostName)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}
