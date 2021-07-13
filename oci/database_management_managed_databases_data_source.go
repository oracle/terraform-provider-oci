// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_database_management "github.com/oracle/oci-go-sdk/v44/databasemanagement"
)

func init() {
	RegisterDatasource("oci_database_management_managed_databases", DatabaseManagementManagedDatabasesDataSource())
}

func DatabaseManagementManagedDatabasesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseManagementManagedDatabases,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"managed_database_collection": {
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
									"compartment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"database_status": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"database_sub_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"database_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_cluster": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"managed_database_groups": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"compartment_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"id": {
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
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"parent_container_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_created": {
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

func readDatabaseManagementManagedDatabases(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementManagedDatabasesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dbManagementClient()

	return ReadResource(sync)
}

type DatabaseManagementManagedDatabasesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.ListManagedDatabasesResponse
}

func (s *DatabaseManagementManagedDatabasesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementManagedDatabasesDataSourceCrud) Get() error {
	request := oci_database_management.ListManagedDatabasesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "database_management")

	response, err := s.Client.ListManagedDatabases(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListManagedDatabases(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseManagementManagedDatabasesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("DatabaseManagementManagedDatabasesDataSource-", DatabaseManagementManagedDatabasesDataSource(), s.D))
	resources := []map[string]interface{}{}
	managedDatabase := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ManagedDatabaseSummaryToMap(item))
	}
	managedDatabase["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseManagementManagedDatabasesDataSource().Schema["managed_database_collection"].Elem.(*schema.Resource).Schema)
		managedDatabase["items"] = items
	}

	resources = append(resources, managedDatabase)
	if err := s.D.Set("managed_database_collection", resources); err != nil {
		return err
	}

	return nil
}

func ManagedDatabaseSummaryToMap(obj oci_database_management.ManagedDatabaseSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	result["database_sub_type"] = string(obj.DatabaseSubType)

	result["database_type"] = string(obj.DatabaseType)

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsCluster != nil {
		result["is_cluster"] = bool(*obj.IsCluster)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.ParentContainerId != nil {
		result["parent_container_id"] = string(*obj.ParentContainerId)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	return result
}

func ParentGroupToMap(obj oci_database_management.ParentGroup) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}
