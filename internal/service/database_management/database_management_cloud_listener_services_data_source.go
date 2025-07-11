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

func DatabaseManagementCloudListenerServicesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseManagementCloudListenerServices,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"cloud_listener_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"managed_database_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"opc_named_credential_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"cloud_listener_service_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"listener_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"managed_database_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": {
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

func readDatabaseManagementCloudListenerServices(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementCloudListenerServicesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementCloudListenerServicesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.ListCloudListenerServicesResponse
}

func (s *DatabaseManagementCloudListenerServicesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementCloudListenerServicesDataSourceCrud) Get() error {
	request := oci_database_management.ListCloudListenerServicesRequest{}

	if cloudListenerId, ok := s.D.GetOkExists("cloud_listener_id"); ok {
		tmp := cloudListenerId.(string)
		request.CloudListenerId = &tmp
	}

	if managedDatabaseId, ok := s.D.GetOkExists("managed_database_id"); ok {
		tmp := managedDatabaseId.(string)
		request.ManagedDatabaseId = &tmp
	}

	if opcNamedCredentialId, ok := s.D.GetOkExists("opc_named_credential_id"); ok {
		tmp := opcNamedCredentialId.(string)
		request.OpcNamedCredentialId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.ListCloudListenerServices(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListCloudListenerServices(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseManagementCloudListenerServicesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementCloudListenerServicesDataSource-", DatabaseManagementCloudListenerServicesDataSource(), s.D))
	resources := []map[string]interface{}{}
	cloudListenerService := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, CloudListenerServiceSummaryToMap(item))
	}
	cloudListenerService["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseManagementCloudListenerServicesDataSource().Schema["cloud_listener_service_collection"].Elem.(*schema.Resource).Schema)
		cloudListenerService["items"] = items
	}

	resources = append(resources, cloudListenerService)
	if err := s.D.Set("cloud_listener_service_collection", resources); err != nil {
		return err
	}

	return nil
}

func CloudListenerServiceSummaryToMap(obj oci_database_management.CloudListenerServiceSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ListenerId != nil {
		result["listener_id"] = string(*obj.ListenerId)
	}

	if obj.ManagedDatabaseId != nil {
		result["managed_database_id"] = string(*obj.ManagedDatabaseId)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}
