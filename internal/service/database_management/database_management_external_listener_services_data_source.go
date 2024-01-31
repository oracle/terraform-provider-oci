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

func DatabaseManagementExternalListenerServicesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseManagementExternalListenerServices,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"external_listener_id": {
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
			"external_listener_service_collection": {
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

func readDatabaseManagementExternalListenerServices(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalListenerServicesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementExternalListenerServicesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.ListExternalListenerServicesResponse
}

func (s *DatabaseManagementExternalListenerServicesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementExternalListenerServicesDataSourceCrud) Get() error {
	request := oci_database_management.ListExternalListenerServicesRequest{}

	if externalListenerId, ok := s.D.GetOkExists("external_listener_id"); ok {
		tmp := externalListenerId.(string)
		request.ExternalListenerId = &tmp
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

	response, err := s.Client.ListExternalListenerServices(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListExternalListenerServices(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseManagementExternalListenerServicesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementExternalListenerServicesDataSource-", DatabaseManagementExternalListenerServicesDataSource(), s.D))
	resources := []map[string]interface{}{}
	externalListenerService := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ExternalListenerServiceSummaryToMap(item))
	}
	externalListenerService["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseManagementExternalListenerServicesDataSource().Schema["external_listener_service_collection"].Elem.(*schema.Resource).Schema)
		externalListenerService["items"] = items
	}

	resources = append(resources, externalListenerService)
	if err := s.D.Set("external_listener_service_collection", resources); err != nil {
		return err
	}

	return nil
}

func ExternalListenerServiceSummaryToMap(obj oci_database_management.ExternalListenerServiceSummary) map[string]interface{} {
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
