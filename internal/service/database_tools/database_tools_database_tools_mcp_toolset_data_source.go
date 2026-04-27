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

func DatabaseToolsDatabaseToolsMcpToolsetDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["database_tools_mcp_toolset_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchemaWithContext(DatabaseToolsDatabaseToolsMcpToolsetResource(), fieldMap, readSingularDatabaseToolsDatabaseToolsMcpToolsetWithContext)
}

func readSingularDatabaseToolsDatabaseToolsMcpToolsetWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseToolsDatabaseToolsMcpToolsetDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DatabaseToolsDatabaseToolsMcpToolsetDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_tools.DatabaseToolsClient
	Res    *oci_database_tools.GetDatabaseToolsMcpToolsetResponse
}

func (s *DatabaseToolsDatabaseToolsMcpToolsetDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseToolsDatabaseToolsMcpToolsetDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_database_tools.GetDatabaseToolsMcpToolsetRequest{}

	if databaseToolsMcpToolsetId, ok := s.D.GetOkExists("database_tools_mcp_toolset_id"); ok {
		tmp := databaseToolsMcpToolsetId.(string)
		request.DatabaseToolsMcpToolsetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_tools")

	response, err := s.Client.GetDatabaseToolsMcpToolset(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseToolsDatabaseToolsMcpToolsetDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.GetId())
	switch v := (s.Res.DatabaseToolsMcpToolset).(type) {
	case oci_database_tools.DatabaseToolsMcpToolsetBuiltInSqlTools:
		s.D.Set("type", "BUILT_IN_SQL_TOOLS")

		s.D.Set("default_execution_type", v.DefaultExecutionType)

		tools := []interface{}{}
		for _, item := range v.Tools {
			tools = append(tools, DatabaseToolsMcpToolsetToolDetailsToMap(item))
		}
		s.D.Set("tools", tools)

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DatabaseToolsMcpServerId != nil {
			s.D.Set("database_tools_mcp_server_id", *v.DatabaseToolsMcpServerId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, DbtoolsMcpToolsetResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.Version != nil {
			s.D.Set("version", *v.Version)
		}
	case oci_database_tools.DatabaseToolsMcpToolsetCustomizableReportingTools:
		s.D.Set("type", "CUSTOMIZABLE_REPORTING_TOOLS")

		s.D.Set("default_execution_type", v.DefaultExecutionType)

		reports := []interface{}{}
		for _, item := range v.Reports {
			reports = append(reports, DatabaseToolsMcpToolsetCustomizableReportingToolsReportToMap(item))
		}
		s.D.Set("reports", reports)

		tools := []interface{}{}
		for _, item := range v.Tools {
			tools = append(tools, DatabaseToolsMcpToolsetToolDetailsToMap(item))
		}
		s.D.Set("tools", tools)

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DatabaseToolsMcpServerId != nil {
			s.D.Set("database_tools_mcp_server_id", *v.DatabaseToolsMcpServerId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, DbtoolsMcpToolsetResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.Version != nil {
			s.D.Set("version", *v.Version)
		}
	case oci_database_tools.DatabaseToolsMcpToolsetCustomSqlTool:
		s.D.Set("type", "CUSTOM_SQL_TOOL")

		s.D.Set("allowed_roles", v.AllowedRoles)

		s.D.Set("default_execution_type", v.DefaultExecutionType)

		if v.Source != nil {
			s.D.Set("source", []interface{}{DatabaseToolsCustomSqlToolToolsetSourceToMap(v.Source)})
		} else {
			s.D.Set("source", nil)
		}

		if v.ToolDescription != nil {
			s.D.Set("tool_description", *v.ToolDescription)
		}

		if v.ToolName != nil {
			s.D.Set("tool_name", *v.ToolName)
		}

		variables := []interface{}{}
		for _, item := range v.Variables {
			variables = append(variables, DatabaseToolsMcpToolsetCustomSqlToolVariableToMap(item))
		}
		s.D.Set("variables", variables)

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DatabaseToolsMcpServerId != nil {
			s.D.Set("database_tools_mcp_server_id", *v.DatabaseToolsMcpServerId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, DbtoolsMcpToolsetResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.Version != nil {
			s.D.Set("version", *v.Version)
		}
	case oci_database_tools.DatabaseToolsMcpToolsetGenAiSqlAssistant:
		s.D.Set("type", "GENAI_SQL_ASSISTANT")

		s.D.Set("default_execution_type", v.DefaultExecutionType)

		if v.GenerativeAiSemanticStoreId != nil {
			s.D.Set("generative_ai_semantic_store_id", *v.GenerativeAiSemanticStoreId)
		}

		tools := []interface{}{}
		for _, item := range v.Tools {
			tools = append(tools, DatabaseToolsMcpToolsetToolDetailsToMap(item))
		}
		s.D.Set("tools", tools)

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DatabaseToolsMcpServerId != nil {
			s.D.Set("database_tools_mcp_server_id", *v.DatabaseToolsMcpServerId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, DbtoolsMcpToolsetResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.Version != nil {
			s.D.Set("version", *v.Version)
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", s.Res.DatabaseToolsMcpToolset)
		return nil
	}

	return nil
}
