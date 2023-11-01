// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database_management "github.com/oracle/oci-go-sdk/v65/databasemanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseManagementManagedMySqlDatabaseConfigurationDataDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseManagementManagedMySqlDatabaseConfigurationData,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"managed_my_sql_database_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"my_sql_configuration_data_collection": {
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
									"default_value": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"host_set": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_configurable": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"is_dynamic": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"is_init": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"max_value": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"min_value": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"path": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"possible_values": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"source": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"supported_versions": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_set": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"user_set": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"value": {
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

func readDatabaseManagementManagedMySqlDatabaseConfigurationData(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementManagedMySqlDatabaseConfigurationDataDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagedMySqlDatabasesClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementManagedMySqlDatabaseConfigurationDataDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.ManagedMySqlDatabasesClient
	Res    *oci_database_management.ListManagedMySqlDatabaseConfigurationDataResponse
}

func (s *DatabaseManagementManagedMySqlDatabaseConfigurationDataDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementManagedMySqlDatabaseConfigurationDataDataSourceCrud) Get() error {
	request := oci_database_management.ListManagedMySqlDatabaseConfigurationDataRequest{}

	if managedMySqlDatabaseId, ok := s.D.GetOkExists("managed_my_sql_database_id"); ok {
		tmp := managedMySqlDatabaseId.(string)
		request.ManagedMySqlDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.ListManagedMySqlDatabaseConfigurationData(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListManagedMySqlDatabaseConfigurationData(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseManagementManagedMySqlDatabaseConfigurationDataDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementManagedMySqlDatabaseConfigurationDataDataSource-", DatabaseManagementManagedMySqlDatabaseConfigurationDataDataSource(), s.D))
	resources := []map[string]interface{}{}
	managedMySqlDatabaseConfigurationData := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, MySqlConfigurationDataSummaryToMap(item))
	}
	managedMySqlDatabaseConfigurationData["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseManagementManagedMySqlDatabaseConfigurationDataDataSource().Schema["my_sql_configuration_data_collection"].Elem.(*schema.Resource).Schema)
		managedMySqlDatabaseConfigurationData["items"] = items
	}

	resources = append(resources, managedMySqlDatabaseConfigurationData)
	if err := s.D.Set("my_sql_configuration_data_collection", resources); err != nil {
		return err
	}

	return nil
}

func MySqlConfigurationDataSummaryToMap(obj oci_database_management.MySqlConfigurationDataSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DefaultValue != nil {
		result["default_value"] = string(*obj.DefaultValue)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.HostSet != nil {
		result["host_set"] = string(*obj.HostSet)
	}

	if obj.IsConfigurable != nil {
		result["is_configurable"] = bool(*obj.IsConfigurable)
	}

	if obj.IsDynamic != nil {
		result["is_dynamic"] = bool(*obj.IsDynamic)
	}

	if obj.IsInit != nil {
		result["is_init"] = bool(*obj.IsInit)
	}

	if obj.MaxValue != nil {
		result["max_value"] = float32(*obj.MaxValue)
	}

	if obj.MinValue != nil {
		result["min_value"] = float32(*obj.MinValue)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Path != nil {
		result["path"] = string(*obj.Path)
	}

	if obj.PossibleValues != nil {
		result["possible_values"] = string(*obj.PossibleValues)
	}

	result["source"] = string(obj.Source)

	if obj.SupportedVersions != nil {
		result["supported_versions"] = string(*obj.SupportedVersions)
	}

	if obj.TimeSet != nil {
		result["time_set"] = obj.TimeSet.String()
	}

	if obj.Type != nil {
		result["type"] = string(*obj.Type)
	}

	if obj.UserSet != nil {
		result["user_set"] = string(*obj.UserSet)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}
