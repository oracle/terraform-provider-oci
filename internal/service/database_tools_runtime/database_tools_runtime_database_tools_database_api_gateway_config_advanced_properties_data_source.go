// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_tools_runtime

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database_tools_runtime "github.com/oracle/oci-go-sdk/v65/databasetoolsruntime"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigAdvancedPropertiesDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigAdvancedPropertiesWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"database_tools_database_api_gateway_config_advanced_property_summary_collection": {
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
									"category_display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"category_key": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"config_types": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"data_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"database_tools_connection_types": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"default_value": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"documentation_url": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"hint_text": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"key": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"list_of_values": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"max_value": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"min_value": {
										Type:     schema.TypeInt,
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

func readDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigAdvancedPropertiesWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigAdvancedPropertiesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsRuntimeClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigAdvancedPropertiesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_tools_runtime.DatabaseToolsRuntimeClient
	Res    *oci_database_tools_runtime.ListDatabaseToolsDatabaseApiGatewayConfigAdvancedPropertiesResponse
}

func (s *DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigAdvancedPropertiesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigAdvancedPropertiesDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_database_tools_runtime.ListDatabaseToolsDatabaseApiGatewayConfigAdvancedPropertiesRequest{}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_tools_runtime")

	response, err := s.Client.ListDatabaseToolsDatabaseApiGatewayConfigAdvancedProperties(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigAdvancedPropertiesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigAdvancedPropertiesDataSource-", DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigAdvancedPropertiesDataSource(), s.D))
	resources := []map[string]interface{}{}
	databaseToolsDatabaseApiGatewayConfigAdvancedProperty := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryToMap(item))
	}
	databaseToolsDatabaseApiGatewayConfigAdvancedProperty["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigAdvancedPropertiesDataSource().Schema["database_tools_database_api_gateway_config_advanced_property_summary_collection"].Elem.(*schema.Resource).Schema)
		databaseToolsDatabaseApiGatewayConfigAdvancedProperty["items"] = items
	}

	resources = append(resources, databaseToolsDatabaseApiGatewayConfigAdvancedProperty)
	if err := s.D.Set("database_tools_database_api_gateway_config_advanced_property_summary_collection", resources); err != nil {
		return err
	}

	return nil
}

func DatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummaryToMap(obj oci_database_tools_runtime.DatabaseToolsDatabaseApiGatewayConfigAdvancedPropertySummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CategoryDisplayName != nil {
		result["category_display_name"] = string(*obj.CategoryDisplayName)
	}

	if obj.CategoryKey != nil {
		result["category_key"] = string(*obj.CategoryKey)
	}

	result["config_types"] = obj.ConfigTypes

	result["data_type"] = string(obj.DataType)

	result["database_tools_connection_types"] = obj.DatabaseToolsConnectionTypes

	if obj.DefaultValue != nil {
		result["default_value"] = string(*obj.DefaultValue)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.DocumentationUrl != nil {
		result["documentation_url"] = string(*obj.DocumentationUrl)
	}

	if obj.HintText != nil {
		result["hint_text"] = string(*obj.HintText)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	result["list_of_values"] = obj.ListOfValues

	if obj.MaxValue != nil {
		result["max_value"] = int(*obj.MaxValue)
	}

	if obj.MinValue != nil {
		result["min_value"] = int(*obj.MinValue)
	}

	return result
}
