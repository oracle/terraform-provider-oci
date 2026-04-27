// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_tools

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database_tools "github.com/oracle/oci-go-sdk/v65/databasetools"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseToolsDatabaseToolsMcpToolsetVersionsDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readDatabaseToolsDatabaseToolsMcpToolsetVersionsWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"database_tools_mcp_server_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"database_tools_mcp_toolset_version_collection": {
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
									"default_version": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"versions": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"default_allowed_roles": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"default_report_allowed_roles": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"description": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"features": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"tools": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"default_allowed_roles": {
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"default_status": {
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
															"name": {
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"version": {
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
				},
			},
		},
	}
}

func readDatabaseToolsDatabaseToolsMcpToolsetVersionsWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseToolsDatabaseToolsMcpToolsetVersionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DatabaseToolsDatabaseToolsMcpToolsetVersionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_tools.DatabaseToolsClient
	Res    *oci_database_tools.ListDatabaseToolsMcpToolsetVersionsResponse
}

func (s *DatabaseToolsDatabaseToolsMcpToolsetVersionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseToolsDatabaseToolsMcpToolsetVersionsDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_database_tools.ListDatabaseToolsMcpToolsetVersionsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if databaseToolsMcpServerId, ok := s.D.GetOkExists("database_tools_mcp_server_id"); ok {
		tmp := databaseToolsMcpServerId.(string)
		request.DatabaseToolsMcpServerId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_tools")

	response, err := s.Client.ListDatabaseToolsMcpToolsetVersions(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseToolsDatabaseToolsMcpToolsetVersionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseToolsDatabaseToolsMcpToolsetVersionsDataSource-", DatabaseToolsDatabaseToolsMcpToolsetVersionsDataSource(), s.D))
	resources := []map[string]interface{}{}
	databaseToolsMcpToolsetVersion := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DatabaseToolsMcpToolsetVersionSummaryToMap(item))
	}
	databaseToolsMcpToolsetVersion["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseToolsDatabaseToolsMcpToolsetVersionsDataSource().Schema["database_tools_mcp_toolset_version_collection"].Elem.(*schema.Resource).Schema)
		databaseToolsMcpToolsetVersion["items"] = items
	}

	resources = append(resources, databaseToolsMcpToolsetVersion)
	if err := s.D.Set("database_tools_mcp_toolset_version_collection", resources); err != nil {
		return err
	}

	return nil
}

func DatabaseToolsMcpToolsetBuiltInSqlToolsVersionToMap(obj oci_database_tools.DatabaseToolsMcpToolsetBuiltInSqlToolsVersion) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	result["features"] = obj.Features

	tools := []interface{}{}
	for _, item := range obj.Tools {
		tools = append(tools, DatabaseToolsMcpToolsetVersionToolToMap(item))
	}
	result["tools"] = tools

	if obj.Version != nil {
		result["version"] = int(*obj.Version)
	}

	return result
}

func DatabaseToolsMcpToolsetCustomSqlToolVersionToMap(obj oci_database_tools.DatabaseToolsMcpToolsetCustomSqlToolVersion) map[string]interface{} {
	result := map[string]interface{}{}

	result["default_allowed_roles"] = obj.DefaultAllowedRoles

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	result["features"] = obj.Features

	if obj.Version != nil {
		result["version"] = int(*obj.Version)
	}

	return result
}

func DatabaseToolsMcpToolsetCustomizableReportingToolsVersionToMap(obj oci_database_tools.DatabaseToolsMcpToolsetCustomizableReportingToolsVersion) map[string]interface{} {
	result := map[string]interface{}{}

	result["default_report_allowed_roles"] = obj.DefaultReportAllowedRoles

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	result["features"] = obj.Features

	tools := []interface{}{}
	for _, item := range obj.Tools {
		tools = append(tools, DatabaseToolsMcpToolsetVersionToolToMap(item))
	}
	result["tools"] = tools

	if obj.Version != nil {
		result["version"] = int(*obj.Version)
	}

	return result
}

func DatabaseToolsMcpToolsetGenAiSqlAssistantVersionToMap(obj oci_database_tools.DatabaseToolsMcpToolsetGenAiSqlAssistantVersion) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	result["features"] = obj.Features

	tools := []interface{}{}
	for _, item := range obj.Tools {
		tools = append(tools, DatabaseToolsMcpToolsetVersionToolToMap(item))
	}
	result["tools"] = tools

	if obj.Version != nil {
		result["version"] = int(*obj.Version)
	}

	return result
}

func DatabaseToolsMcpToolsetVersionSummaryToMap(obj oci_database_tools.DatabaseToolsMcpToolsetVersionSummary) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_database_tools.DatabaseToolsMcpToolsetVersionBuiltInSqlToolsSummary:
		result["type"] = "BUILT_IN_SQL_TOOLS"

		versions := []interface{}{}
		for _, item := range v.Versions {
			versions = append(versions, DatabaseToolsMcpToolsetBuiltInSqlToolsVersionToMap(item))
		}
		result["versions"] = versions

		if v.DefaultVersion != nil {
			result["default_version"] = int(*v.DefaultVersion)
		}
	case oci_database_tools.DatabaseToolsMcpToolsetVersionCustomizableReportingToolsSummary:
		result["type"] = "CUSTOMIZABLE_REPORTING_TOOLS"

		versions := []interface{}{}
		for _, item := range v.Versions {
			versions = append(versions, DatabaseToolsMcpToolsetCustomizableReportingToolsVersionToMap(item))
		}
		result["versions"] = versions

		if v.DefaultVersion != nil {
			result["default_version"] = int(*v.DefaultVersion)
		}
	case oci_database_tools.DatabaseToolsMcpToolsetVersionCustomSqlToolSummary:
		result["type"] = "CUSTOM_SQL_TOOL"

		versions := []interface{}{}
		for _, item := range v.Versions {
			versions = append(versions, DatabaseToolsMcpToolsetCustomSqlToolVersionToMap(item))
		}
		result["versions"] = versions

		if v.DefaultVersion != nil {
			result["default_version"] = int(*v.DefaultVersion)
		}
	case oci_database_tools.DatabaseToolsMcpToolsetVersionGenAiSqlAssistantSummary:
		result["type"] = "GENAI_SQL_ASSISTANT"

		versions := []interface{}{}
		for _, item := range v.Versions {
			versions = append(versions, DatabaseToolsMcpToolsetGenAiSqlAssistantVersionToMap(item))
		}
		result["versions"] = versions

		if v.DefaultVersion != nil {
			result["default_version"] = int(*v.DefaultVersion)
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", obj)
		return nil
	}

	return result
}

func DatabaseToolsMcpToolsetVersionToolToMap(obj oci_database_tools.DatabaseToolsMcpToolsetVersionTool) map[string]interface{} {
	result := map[string]interface{}{}

	result["default_allowed_roles"] = obj.DefaultAllowedRoles

	result["default_status"] = string(obj.DefaultStatus)

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}
