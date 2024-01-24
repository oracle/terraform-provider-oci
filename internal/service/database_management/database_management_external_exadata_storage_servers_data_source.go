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

func DatabaseManagementExternalExadataStorageServersDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseManagementExternalExadataStorageServers,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_exadata_infrastructure_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"external_exadata_storage_server_collection": {
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
									"additional_details": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"connector_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"cpu_count": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"exadata_infrastructure_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"internal_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"ip_address": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"lifecycle_details": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"make_model": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"max_flash_disk_iops": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"max_flash_disk_throughput": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"max_hard_disk_iops": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"max_hard_disk_throughput": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"memory_gb": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"resource_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"status": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"storage_grid_id": {
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
									"version": {
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

func readDatabaseManagementExternalExadataStorageServers(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalExadataStorageServersDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementExternalExadataStorageServersDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.ListExternalExadataStorageServersResponse
}

func (s *DatabaseManagementExternalExadataStorageServersDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementExternalExadataStorageServersDataSourceCrud) Get() error {
	request := oci_database_management.ListExternalExadataStorageServersRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if externalExadataInfrastructureId, ok := s.D.GetOkExists("external_exadata_infrastructure_id"); ok {
		tmp := externalExadataInfrastructureId.(string)
		request.ExternalExadataInfrastructureId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.ListExternalExadataStorageServers(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListExternalExadataStorageServers(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseManagementExternalExadataStorageServersDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementExternalExadataStorageServersDataSource-", DatabaseManagementExternalExadataStorageServersDataSource(), s.D))
	resources := []map[string]interface{}{}
	externalExadataStorageServer := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ExternalExadataStorageServerSummaryToMap(item))
	}
	externalExadataStorageServer["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseManagementExternalExadataStorageServersDataSource().Schema["external_exadata_storage_server_collection"].Elem.(*schema.Resource).Schema)
		externalExadataStorageServer["items"] = items
	}

	resources = append(resources, externalExadataStorageServer)
	if err := s.D.Set("external_exadata_storage_server_collection", resources); err != nil {
		return err
	}

	return nil
}

func ExternalExadataStorageConnectorSummaryToMap1(obj *oci_database_management.ExternalExadataStorageConnectorSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["additional_details"] = obj.AdditionalDetails

	if obj.AgentId != nil {
		result["agent_id"] = string(*obj.AgentId)
	}

	if obj.ConnectionUri != nil {
		result["connection_uri"] = string(*obj.ConnectionUri)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.InternalId != nil {
		result["internal_id"] = string(*obj.InternalId)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.Status != nil {
		result["status"] = string(*obj.Status)
	}

	if obj.StorageServerId != nil {
		result["storage_server_id"] = string(*obj.StorageServerId)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	if obj.Version != nil {
		result["version"] = string(*obj.Version)
	}

	return result
}
