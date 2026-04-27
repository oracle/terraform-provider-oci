// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_database_tools "github.com/oracle/oci-go-sdk/v65/databasetools"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

// ============================================================
// Shared sub-representations (used across all types)
// ============================================================

// Datasource filter: narrows list results to the single resource under test.
var DatabaseToolsDatabaseToolsMcpToolsetDataSourceFilterRepresentation = map[string]interface{}{
	"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
	"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_tools_database_tools_mcp_toolset.test_database_tools_mcp_toolset.id}`}},
}

// Tools: used by BUILT_IN_SQL_TOOLS, CUSTOMIZABLE_REPORTING_TOOLS, GENAI_SQL_ASSISTANT
var DatabaseToolsDatabaseToolsMcpToolsetToolsRepresentation = map[string]interface{}{
	"name":   acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
	"status": acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`, Update: `DISABLED`},
}

// No Terraform-managed resource dependencies needed for these tests
var DatabaseToolsDatabaseToolsMcpToolsetResourceDependencies = ""

// ============================================================
// Type 1: CUSTOM_SQL_TOOL
// ============================================================

var (
	DatabaseToolsDatabaseToolsMcpToolsetCustomSqlToolRequiredOnlyResource = DatabaseToolsDatabaseToolsMcpToolsetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_mcp_toolset", "test_database_tools_mcp_toolset", acctest.Required, acctest.Create, DatabaseToolsDatabaseToolsMcpToolsetCustomSqlToolRepresentation)

	DatabaseToolsDatabaseToolsMcpToolsetCustomSqlToolResourceConfig = DatabaseToolsDatabaseToolsMcpToolsetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_mcp_toolset", "test_database_tools_mcp_toolset", acctest.Optional, acctest.Update, DatabaseToolsDatabaseToolsMcpToolsetCustomSqlToolRepresentation)

	DatabaseToolsDatabaseToolsMcpToolsetCustomSqlToolSingularDataSourceRepresentation = map[string]interface{}{
		"database_tools_mcp_toolset_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_tools_database_tools_mcp_toolset.test_database_tools_mcp_toolset.id}`},
	}

	DatabaseToolsDatabaseToolsMcpToolsetCustomSqlToolDataSourceRepresentation = map[string]interface{}{
		"compartment_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"database_tools_mcp_server_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.database_tools_mcp_server_id}`},
		"display_name":                 acctest.Representation{RepType: acctest.Optional, Create: `displayName2`},
		"state":                        acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"type":                         acctest.Representation{RepType: acctest.Optional, Create: []string{`CUSTOM_SQL_TOOL`}},
		"filter":                       acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseToolsDatabaseToolsMcpToolsetDataSourceFilterRepresentation},
	}

	// source.type is always INLINE — the only enum value in the spec.
	DatabaseToolsDatabaseToolsMcpToolsetCustomSqlToolSourceRepresentation = map[string]interface{}{
		"type":  acctest.Representation{RepType: acctest.Required, Create: `INLINE`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `select 1 from dual`, Update: `select 2 from dual`},
	}

	// Variable names must match ^[A-Za-z][A-Za-z0-9_$#]{0,29}$
	DatabaseToolsDatabaseToolsMcpToolsetCustomSqlToolVariableRepresentation = map[string]interface{}{
		"name":        acctest.Representation{RepType: acctest.Required, Create: `Varname`, Update: `Varname2`},
		"type":        acctest.Representation{RepType: acctest.Required, Create: `VARCHAR2`, Update: `NUMBER`},
		"description": acctest.Representation{RepType: acctest.Required, Create: `variable description`, Update: `variable description2`},
	}

	// toolName must match ^[A-Za-z0-9_.-]+$
	DatabaseToolsDatabaseToolsMcpToolsetCustomSqlToolRepresentation = map[string]interface{}{
		// Required (base)
		"compartment_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"database_tools_mcp_server_id": acctest.Representation{RepType: acctest.Required, Create: `${var.database_tools_mcp_server_id}`},
		"display_name":                 acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"type":                         acctest.Representation{RepType: acctest.Required, Create: `CUSTOM_SQL_TOOL`},
		"version":                      acctest.Representation{RepType: acctest.Required, Create: `1`},
		// Required (type-specific)
		"tool_name": acctest.Representation{RepType: acctest.Required, Create: `tf_acc_custom_sql_tool_001`},
		"source":    acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseToolsDatabaseToolsMcpToolsetCustomSqlToolSourceRepresentation},
		// Optional
		"default_execution_type": acctest.Representation{RepType: acctest.Optional, Create: `SYNCHRONOUS`, Update: `ASYNCHRONOUS`},
		"description":            acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":          acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"tool_description":       acctest.Representation{RepType: acctest.Optional, Create: `toolDescription`, Update: `toolDescription2`},
	}
)

// issue-routing-tag: database_tools/default
func TestDatabaseToolsDatabaseToolsMcpToolsetCustomSqlToolResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseToolsDatabaseToolsMcpToolsetCustomSqlToolResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	relatedResourceId := utils.GetEnvSettingWithBlankDefault("TF_VAR_related_resource_id")
	relatedResourceIdVariableStr := fmt.Sprintf("variable \"related_resource_id\" { default = \"%s\" }\n", relatedResourceId)

	mcpServerId := utils.GetEnvSettingWithBlankDefault("database_tools_mcp_server_ocid")
	mcpServerIdVariableStr := fmt.Sprintf("variable \"database_tools_mcp_server_id\" { default = \"%s\" }\n", mcpServerId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_database_tools_database_tools_mcp_toolset.test_database_tools_mcp_toolset"
	datasourceName := "data.oci_database_tools_database_tools_mcp_toolsets.test_database_tools_mcp_toolsets"
	singularDatasourceName := "data.oci_database_tools_database_tools_mcp_toolset.test_database_tools_mcp_toolset"

	commonVars := config + compartmentIdVariableStr + relatedResourceIdVariableStr + mcpServerIdVariableStr
	commonVarsU := commonVars + compartmentIdUVariableStr

	var resId, resId2 string
	acctest.SaveConfigContent(commonVars+DatabaseToolsDatabaseToolsMcpToolsetResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_mcp_toolset", "test_database_tools_mcp_toolset", acctest.Optional, acctest.Create, DatabaseToolsDatabaseToolsMcpToolsetCustomSqlToolRepresentation),
		"databasetools", "databaseToolsMcpToolsetCustomSqlTool", t)

	acctest.ResourceTest(t, testAccCheckDatabaseToolsDatabaseToolsMcpToolsetDestroy, []resource.TestStep{
		// Step 1: verify Create with Required only
		{
			Config: commonVars + DatabaseToolsDatabaseToolsMcpToolsetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_mcp_toolset", "test_database_tools_mcp_toolset", acctest.Required, acctest.Create, DatabaseToolsDatabaseToolsMcpToolsetCustomSqlToolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "database_tools_mcp_server_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "type", "CUSTOM_SQL_TOOL"),
				resource.TestCheckResourceAttr(resourceName, "version", "1"),
				resource.TestCheckResourceAttr(resourceName, "tool_name", "tf_acc_custom_sql_tool_001"),
				resource.TestCheckResourceAttr(resourceName, "source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source.0.type", "INLINE"),
				resource.TestCheckResourceAttr(resourceName, "source.0.value", "select 1 from dual"),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// Step 2: delete before next Create
		{
			Config: commonVars + DatabaseToolsDatabaseToolsMcpToolsetResourceDependencies,
		},
		// Step 3: verify Create with optionals
		{
			Config: commonVars + DatabaseToolsDatabaseToolsMcpToolsetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_mcp_toolset", "test_database_tools_mcp_toolset", acctest.Optional, acctest.Create, DatabaseToolsDatabaseToolsMcpToolsetCustomSqlToolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "database_tools_mcp_server_id"),
				resource.TestCheckResourceAttr(resourceName, "default_execution_type", "SYNCHRONOUS"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source.0.type", "INLINE"),
				resource.TestCheckResourceAttr(resourceName, "source.0.value", "select 1 from dual"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "tool_description", "toolDescription"),
				resource.TestCheckResourceAttr(resourceName, "tool_name", "tf_acc_custom_sql_tool_001"),
				resource.TestCheckResourceAttr(resourceName, "type", "CUSTOM_SQL_TOOL"),
				resource.TestCheckResourceAttr(resourceName, "version", "1"),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},
		// Step 4: verify Update to compartment (switched back in next step)
		{
			Config: commonVarsU + DatabaseToolsDatabaseToolsMcpToolsetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_mcp_toolset", "test_database_tools_mcp_toolset", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DatabaseToolsDatabaseToolsMcpToolsetCustomSqlToolRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "type", "CUSTOM_SQL_TOOL"),
				resource.TestCheckResourceAttr(resourceName, "source.0.type", "INLINE"),
				resource.TestCheckResourceAttr(resourceName, "source.0.value", "select 1 from dual"),
				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
		// Step 5: verify updates to updatable parameters
		{
			Config: commonVars + DatabaseToolsDatabaseToolsMcpToolsetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_mcp_toolset", "test_database_tools_mcp_toolset", acctest.Optional, acctest.Update, DatabaseToolsDatabaseToolsMcpToolsetCustomSqlToolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "database_tools_mcp_server_id"),
				resource.TestCheckResourceAttr(resourceName, "default_execution_type", "ASYNCHRONOUS"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source.0.type", "INLINE"),
				resource.TestCheckResourceAttr(resourceName, "source.0.value", "select 2 from dual"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "tool_description", "toolDescription2"),
				resource.TestCheckResourceAttr(resourceName, "tool_name", "tf_acc_custom_sql_tool_001"),
				resource.TestCheckResourceAttr(resourceName, "type", "CUSTOM_SQL_TOOL"),
				resource.TestCheckResourceAttr(resourceName, "version", "1"),
				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// Step 6: verify datasource (list)
		{
			Config: commonVars + DatabaseToolsDatabaseToolsMcpToolsetResourceDependencies +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_mcp_toolsets", "test_database_tools_mcp_toolsets", acctest.Optional, acctest.Update, DatabaseToolsDatabaseToolsMcpToolsetCustomSqlToolDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_mcp_toolset", "test_database_tools_mcp_toolset", acctest.Optional, acctest.Update, DatabaseToolsDatabaseToolsMcpToolsetCustomSqlToolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "database_tools_mcp_server_id"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "type.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "database_tools_mcp_toolset_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "database_tools_mcp_toolset_collection.0.items.#", "1"),
			),
		},
		// Step 7: verify singular datasource
		{
			Config: commonVars + DatabaseToolsDatabaseToolsMcpToolsetCustomSqlToolResourceConfig +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_mcp_toolset", "test_database_tools_mcp_toolset", acctest.Required, acctest.Create, DatabaseToolsDatabaseToolsMcpToolsetCustomSqlToolSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "database_tools_mcp_toolset_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "default_execution_type", "ASYNCHRONOUS"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "source.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "source.0.type", "INLINE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "source.0.value", "select 2 from dual"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_description", "toolDescription2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "type", "CUSTOM_SQL_TOOL"),
				resource.TestCheckResourceAttr(singularDatasourceName, "version", "1"),
			),
		},
		// Step 8: verify resource import
		{
			Config:                  commonVars + DatabaseToolsDatabaseToolsMcpToolsetCustomSqlToolRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

// ============================================================
// Type 2: BUILT_IN_SQL_TOOLS
// ============================================================

var (
	DatabaseToolsDatabaseToolsMcpToolsetBuiltInSqlToolsRequiredOnlyResource = DatabaseToolsDatabaseToolsMcpToolsetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_mcp_toolset", "test_database_tools_mcp_toolset", acctest.Required, acctest.Create, DatabaseToolsDatabaseToolsMcpToolsetBuiltInSqlToolsRepresentation)

	DatabaseToolsDatabaseToolsMcpToolsetBuiltInSqlToolsResourceConfig = DatabaseToolsDatabaseToolsMcpToolsetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_mcp_toolset", "test_database_tools_mcp_toolset", acctest.Optional, acctest.Update, DatabaseToolsDatabaseToolsMcpToolsetBuiltInSqlToolsRepresentation)

	DatabaseToolsDatabaseToolsMcpToolsetBuiltInSqlToolsSingularDataSourceRepresentation = map[string]interface{}{
		"database_tools_mcp_toolset_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_tools_database_tools_mcp_toolset.test_database_tools_mcp_toolset.id}`},
	}

	DatabaseToolsDatabaseToolsMcpToolsetBuiltInSqlToolsDataSourceRepresentation = map[string]interface{}{
		"compartment_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"database_tools_mcp_server_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.database_tools_mcp_server_id}`},
		"display_name":                 acctest.Representation{RepType: acctest.Optional, Create: `displayName2`},
		"state":                        acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"type":                         acctest.Representation{RepType: acctest.Optional, Create: []string{`BUILT_IN_SQL_TOOLS`}},
		"filter":                       acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseToolsDatabaseToolsMcpToolsetDataSourceFilterRepresentation},
	}

	// BUILT_IN_SQL_TOOLS has no type-specific required fields beyond the base.
	DatabaseToolsDatabaseToolsMcpToolsetBuiltInSqlToolsRepresentation = map[string]interface{}{
		// Required (base)
		"compartment_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"database_tools_mcp_server_id": acctest.Representation{RepType: acctest.Required, Create: `${var.database_tools_mcp_server_id}`},
		"display_name":                 acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"type":                         acctest.Representation{RepType: acctest.Required, Create: `BUILT_IN_SQL_TOOLS`},
		"version":                      acctest.Representation{RepType: acctest.Required, Create: `1`},
		// Optional
		"default_execution_type": acctest.Representation{RepType: acctest.Optional, Create: `SYNCHRONOUS`, Update: `ASYNCHRONOUS`},
		"description":            acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":          acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"tools": acctest.RepresentationGroup{RepType: acctest.Optional, Group: acctest.RepresentationCopyWithNewProperties(DatabaseToolsDatabaseToolsMcpToolsetToolsRepresentation, map[string]interface{}{
			"name": acctest.Representation{RepType: acctest.Required, Create: `dbtools_execute_sql`},
		})},
	}
)

// issue-routing-tag: database_tools/default
func TestDatabaseToolsDatabaseToolsMcpToolsetBuiltInSqlToolsResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseToolsDatabaseToolsMcpToolsetBuiltInSqlToolsResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	relatedResourceId := utils.GetEnvSettingWithBlankDefault("TF_VAR_related_resource_id")
	relatedResourceIdVariableStr := fmt.Sprintf("variable \"related_resource_id\" { default = \"%s\" }\n", relatedResourceId)

	mcpServerId := utils.GetEnvSettingWithBlankDefault("database_tools_mcp_server_ocid")
	mcpServerIdVariableStr := fmt.Sprintf("variable \"database_tools_mcp_server_id\" { default = \"%s\" }\n", mcpServerId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_database_tools_database_tools_mcp_toolset.test_database_tools_mcp_toolset"
	datasourceName := "data.oci_database_tools_database_tools_mcp_toolsets.test_database_tools_mcp_toolsets"
	singularDatasourceName := "data.oci_database_tools_database_tools_mcp_toolset.test_database_tools_mcp_toolset"

	commonVars := config + compartmentIdVariableStr + relatedResourceIdVariableStr + mcpServerIdVariableStr
	commonVarsU := commonVars + compartmentIdUVariableStr

	var resId, resId2 string
	acctest.SaveConfigContent(commonVars+DatabaseToolsDatabaseToolsMcpToolsetResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_mcp_toolset", "test_database_tools_mcp_toolset", acctest.Optional, acctest.Create, DatabaseToolsDatabaseToolsMcpToolsetBuiltInSqlToolsRepresentation),
		"databasetools", "databaseToolsMcpToolsetBuiltInSqlTools", t)

	acctest.ResourceTest(t, testAccCheckDatabaseToolsDatabaseToolsMcpToolsetDestroy, []resource.TestStep{
		// Step 1: verify Create with Required only
		{
			Config: commonVars + DatabaseToolsDatabaseToolsMcpToolsetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_mcp_toolset", "test_database_tools_mcp_toolset", acctest.Required, acctest.Create, DatabaseToolsDatabaseToolsMcpToolsetBuiltInSqlToolsRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "database_tools_mcp_server_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "type", "BUILT_IN_SQL_TOOLS"),
				resource.TestCheckResourceAttr(resourceName, "version", "1"),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// Step 2: delete before next Create
		{
			Config: commonVars + DatabaseToolsDatabaseToolsMcpToolsetResourceDependencies,
		},
		// Step 3: verify Create with optionals
		{
			Config: commonVars + DatabaseToolsDatabaseToolsMcpToolsetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_mcp_toolset", "test_database_tools_mcp_toolset", acctest.Optional, acctest.Create, DatabaseToolsDatabaseToolsMcpToolsetBuiltInSqlToolsRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "database_tools_mcp_server_id"),
				resource.TestCheckResourceAttr(resourceName, "default_execution_type", "SYNCHRONOUS"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "tools.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tools.0.name", "dbtools_execute_sql"),
				resource.TestCheckResourceAttr(resourceName, "tools.0.status", "ENABLED"),
				resource.TestCheckResourceAttr(resourceName, "type", "BUILT_IN_SQL_TOOLS"),
				resource.TestCheckResourceAttr(resourceName, "version", "1"),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},
		// Step 4: verify Update to compartment (switched back in next step)
		{
			Config: commonVarsU + DatabaseToolsDatabaseToolsMcpToolsetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_mcp_toolset", "test_database_tools_mcp_toolset", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DatabaseToolsDatabaseToolsMcpToolsetBuiltInSqlToolsRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "type", "BUILT_IN_SQL_TOOLS"),
				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
		// Step 5: verify updates to updatable parameters
		{
			Config: commonVars + DatabaseToolsDatabaseToolsMcpToolsetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_mcp_toolset", "test_database_tools_mcp_toolset", acctest.Optional, acctest.Update, DatabaseToolsDatabaseToolsMcpToolsetBuiltInSqlToolsRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "database_tools_mcp_server_id"),
				resource.TestCheckResourceAttr(resourceName, "default_execution_type", "ASYNCHRONOUS"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "tools.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tools.0.name", "dbtools_execute_sql"),
				resource.TestCheckResourceAttr(resourceName, "tools.0.status", "DISABLED"),
				resource.TestCheckResourceAttr(resourceName, "type", "BUILT_IN_SQL_TOOLS"),
				resource.TestCheckResourceAttr(resourceName, "version", "1"),
				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// Step 6: verify datasource (list)
		{
			Config: commonVars + DatabaseToolsDatabaseToolsMcpToolsetResourceDependencies +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_mcp_toolsets", "test_database_tools_mcp_toolsets", acctest.Optional, acctest.Update, DatabaseToolsDatabaseToolsMcpToolsetBuiltInSqlToolsDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_mcp_toolset", "test_database_tools_mcp_toolset", acctest.Optional, acctest.Update, DatabaseToolsDatabaseToolsMcpToolsetBuiltInSqlToolsRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "database_tools_mcp_server_id"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "type.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "database_tools_mcp_toolset_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "database_tools_mcp_toolset_collection.0.items.#", "1"),
			),
		},
		// Step 7: verify singular datasource
		{
			Config: commonVars + DatabaseToolsDatabaseToolsMcpToolsetBuiltInSqlToolsResourceConfig +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_mcp_toolset", "test_database_tools_mcp_toolset", acctest.Required, acctest.Create, DatabaseToolsDatabaseToolsMcpToolsetBuiltInSqlToolsSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "database_tools_mcp_toolset_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "default_execution_type", "ASYNCHRONOUS"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tools.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tools.0.name", "dbtools_execute_sql"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "tools.0.display_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tools.0.status", "DISABLED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "type", "BUILT_IN_SQL_TOOLS"),
				resource.TestCheckResourceAttr(singularDatasourceName, "version", "1"),
			),
		},
		// Step 8: verify resource import
		{
			Config:                  commonVars + DatabaseToolsDatabaseToolsMcpToolsetBuiltInSqlToolsRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

// ============================================================
// Type 3: CUSTOMIZABLE_REPORTING_TOOLS
// ============================================================

var (
	DatabaseToolsDatabaseToolsMcpToolsetCustomizableReportingToolsRequiredOnlyResource = DatabaseToolsDatabaseToolsMcpToolsetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_mcp_toolset", "test_database_tools_mcp_toolset", acctest.Required, acctest.Create, DatabaseToolsDatabaseToolsMcpToolsetCustomizableReportingToolsRepresentation)

	DatabaseToolsDatabaseToolsMcpToolsetCustomizableReportingToolsResourceConfig = DatabaseToolsDatabaseToolsMcpToolsetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_mcp_toolset", "test_database_tools_mcp_toolset", acctest.Optional, acctest.Update, DatabaseToolsDatabaseToolsMcpToolsetCustomizableReportingToolsRepresentation)

	DatabaseToolsDatabaseToolsMcpToolsetCustomizableReportingToolsSingularDataSourceRepresentation = map[string]interface{}{
		"database_tools_mcp_toolset_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_tools_database_tools_mcp_toolset.test_database_tools_mcp_toolset.id}`},
	}

	DatabaseToolsDatabaseToolsMcpToolsetCustomizableReportingToolsDataSourceRepresentation = map[string]interface{}{
		"compartment_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"database_tools_mcp_server_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.database_tools_mcp_server_id}`},
		"display_name":                 acctest.Representation{RepType: acctest.Optional, Create: `displayName2`},
		"state":                        acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"type":                         acctest.Representation{RepType: acctest.Optional, Create: []string{`CUSTOMIZABLE_REPORTING_TOOLS`}},
		"filter":                       acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseToolsDatabaseToolsMcpToolsetDataSourceFilterRepresentation},
	}

	// reports.databaseToolsSqlReportId is required by the spec for create.
	DatabaseToolsDatabaseToolsMcpToolsetReportsRepresentation = map[string]interface{}{
		"database_tools_sql_report_id": acctest.Representation{RepType: acctest.Required, Create: `${var.database_tools_sql_report_id}`},
	}

	DatabaseToolsDatabaseToolsMcpToolsetCustomizableReportingToolsRepresentation = map[string]interface{}{
		// Required (base)
		"compartment_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"database_tools_mcp_server_id": acctest.Representation{RepType: acctest.Required, Create: `${var.database_tools_mcp_server_id}`},
		"display_name":                 acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"type":                         acctest.Representation{RepType: acctest.Required, Create: `CUSTOMIZABLE_REPORTING_TOOLS`},
		"version":                      acctest.Representation{RepType: acctest.Required, Create: `1`},
		// Required (type-specific): at least one report is required
		"reports": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseToolsDatabaseToolsMcpToolsetReportsRepresentation},
		// Optional
		"default_execution_type": acctest.Representation{RepType: acctest.Optional, Create: `SYNCHRONOUS`, Update: `ASYNCHRONOUS`},
		"description":            acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":          acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"tools": acctest.Representation{RepType: acctest.Optional, Create: []acctest.RepresentationGroup{
			{
				RepType: acctest.Optional,
				Group: acctest.RepresentationCopyWithNewProperties(DatabaseToolsDatabaseToolsMcpToolsetToolsRepresentation, map[string]interface{}{
					"name": acctest.Representation{RepType: acctest.Required, Create: `dbtools_list_reports`},
				}),
			},
			{
				RepType: acctest.Optional,
				Group: acctest.RepresentationCopyWithNewProperties(DatabaseToolsDatabaseToolsMcpToolsetToolsRepresentation, map[string]interface{}{
					"name": acctest.Representation{RepType: acctest.Required, Create: `dbtools_execute_report`},
				}),
			},
		}},
	}
)

// issue-routing-tag: database_tools/default
func TestDatabaseToolsDatabaseToolsMcpToolsetCustomizableReportingToolsResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseToolsDatabaseToolsMcpToolsetCustomizableReportingToolsResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	relatedResourceId := utils.GetEnvSettingWithBlankDefault("TF_VAR_related_resource_id")
	relatedResourceIdVariableStr := fmt.Sprintf("variable \"related_resource_id\" { default = \"%s\" }\n", relatedResourceId)

	mcpServerId := utils.GetEnvSettingWithBlankDefault("database_tools_mcp_server_ocid")
	mcpServerIdVariableStr := fmt.Sprintf("variable \"database_tools_mcp_server_id\" { default = \"%s\" }\n", mcpServerId)

	sqlReportId := utils.GetEnvSettingWithBlankDefault("database_tools_sql_report_ocid")
	sqlReportIdVariableStr := fmt.Sprintf("variable \"database_tools_sql_report_id\" { default = \"%s\" }\n", sqlReportId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_database_tools_database_tools_mcp_toolset.test_database_tools_mcp_toolset"
	datasourceName := "data.oci_database_tools_database_tools_mcp_toolsets.test_database_tools_mcp_toolsets"
	singularDatasourceName := "data.oci_database_tools_database_tools_mcp_toolset.test_database_tools_mcp_toolset"

	commonVars := config + compartmentIdVariableStr + relatedResourceIdVariableStr + mcpServerIdVariableStr + sqlReportIdVariableStr
	commonVarsU := commonVars + compartmentIdUVariableStr

	var resId, resId2 string
	acctest.SaveConfigContent(commonVars+DatabaseToolsDatabaseToolsMcpToolsetResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_mcp_toolset", "test_database_tools_mcp_toolset", acctest.Optional, acctest.Create, DatabaseToolsDatabaseToolsMcpToolsetCustomizableReportingToolsRepresentation),
		"databasetools", "databaseToolsMcpToolsetCustomizableReportingTools", t)

	acctest.ResourceTest(t, testAccCheckDatabaseToolsDatabaseToolsMcpToolsetDestroy, []resource.TestStep{
		// Step 1: verify Create with Required only
		{
			Config: commonVars + DatabaseToolsDatabaseToolsMcpToolsetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_mcp_toolset", "test_database_tools_mcp_toolset", acctest.Required, acctest.Create, DatabaseToolsDatabaseToolsMcpToolsetCustomizableReportingToolsRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "database_tools_mcp_server_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "type", "CUSTOMIZABLE_REPORTING_TOOLS"),
				resource.TestCheckResourceAttr(resourceName, "version", "1"),
				resource.TestCheckResourceAttr(resourceName, "reports.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "reports.0.database_tools_sql_report_id"),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// Step 2: delete before next Create
		{
			Config: commonVars + DatabaseToolsDatabaseToolsMcpToolsetResourceDependencies,
		},
		// Step 3: verify Create with optionals
		{
			Config: commonVars + DatabaseToolsDatabaseToolsMcpToolsetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_mcp_toolset", "test_database_tools_mcp_toolset", acctest.Optional, acctest.Create, DatabaseToolsDatabaseToolsMcpToolsetCustomizableReportingToolsRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "database_tools_mcp_server_id"),
				resource.TestCheckResourceAttr(resourceName, "default_execution_type", "SYNCHRONOUS"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "reports.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "reports.0.database_tools_sql_report_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "tools.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "tools.0.name", "dbtools_list_reports"),
				resource.TestCheckResourceAttr(resourceName, "tools.0.status", "ENABLED"),
				resource.TestCheckResourceAttr(resourceName, "tools.1.name", "dbtools_execute_report"),
				resource.TestCheckResourceAttr(resourceName, "tools.1.status", "ENABLED"),
				resource.TestCheckResourceAttr(resourceName, "type", "CUSTOMIZABLE_REPORTING_TOOLS"),
				resource.TestCheckResourceAttr(resourceName, "version", "1"),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},
		// Step 4: verify Update to compartment (switched back in next step)
		{
			Config: commonVarsU + DatabaseToolsDatabaseToolsMcpToolsetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_mcp_toolset", "test_database_tools_mcp_toolset", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DatabaseToolsDatabaseToolsMcpToolsetCustomizableReportingToolsRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "type", "CUSTOMIZABLE_REPORTING_TOOLS"),
				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
		// Step 5: verify updates to updatable parameters
		{
			Config: commonVars + DatabaseToolsDatabaseToolsMcpToolsetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_mcp_toolset", "test_database_tools_mcp_toolset", acctest.Optional, acctest.Update, DatabaseToolsDatabaseToolsMcpToolsetCustomizableReportingToolsRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "database_tools_mcp_server_id"),
				resource.TestCheckResourceAttr(resourceName, "default_execution_type", "ASYNCHRONOUS"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "reports.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "reports.0.database_tools_sql_report_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "tools.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "tools.0.name", "dbtools_list_reports"),
				resource.TestCheckResourceAttr(resourceName, "tools.0.status", "ENABLED"),
				resource.TestCheckResourceAttr(resourceName, "tools.1.name", "dbtools_execute_report"),
				resource.TestCheckResourceAttr(resourceName, "tools.1.status", "ENABLED"),
				resource.TestCheckResourceAttr(resourceName, "type", "CUSTOMIZABLE_REPORTING_TOOLS"),
				resource.TestCheckResourceAttr(resourceName, "version", "1"),
				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// Step 6: verify datasource (list)
		{
			Config: commonVars + DatabaseToolsDatabaseToolsMcpToolsetResourceDependencies +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_mcp_toolsets", "test_database_tools_mcp_toolsets", acctest.Optional, acctest.Update, DatabaseToolsDatabaseToolsMcpToolsetCustomizableReportingToolsDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_mcp_toolset", "test_database_tools_mcp_toolset", acctest.Optional, acctest.Update, DatabaseToolsDatabaseToolsMcpToolsetCustomizableReportingToolsRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "database_tools_mcp_server_id"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "type.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "database_tools_mcp_toolset_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "database_tools_mcp_toolset_collection.0.items.#", "1"),
			),
		},
		// Step 7: verify singular datasource
		{
			Config: commonVars + DatabaseToolsDatabaseToolsMcpToolsetCustomizableReportingToolsResourceConfig +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_mcp_toolset", "test_database_tools_mcp_toolset", acctest.Required, acctest.Create, DatabaseToolsDatabaseToolsMcpToolsetCustomizableReportingToolsSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "database_tools_mcp_toolset_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "default_execution_type", "ASYNCHRONOUS"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "reports.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "reports.0.database_tools_sql_report_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tools.#", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tools.0.name", "dbtools_list_reports"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "tools.0.display_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tools.0.status", "ENABLED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tools.1.name", "dbtools_execute_report"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "tools.1.display_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tools.1.status", "ENABLED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "type", "CUSTOMIZABLE_REPORTING_TOOLS"),
				resource.TestCheckResourceAttr(singularDatasourceName, "version", "1"),
			),
		},
		// Step 8: verify resource import
		{
			Config:                  commonVars + DatabaseToolsDatabaseToolsMcpToolsetCustomizableReportingToolsRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

// ============================================================
// Type 4: GENAI_SQL_ASSISTANT
// ============================================================

var (
	DatabaseToolsDatabaseToolsMcpToolsetGenAiSqlAssistantRequiredOnlyResource = DatabaseToolsDatabaseToolsMcpToolsetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_mcp_toolset", "test_database_tools_mcp_toolset", acctest.Required, acctest.Create, DatabaseToolsDatabaseToolsMcpToolsetGenAiSqlAssistantRepresentation)

	DatabaseToolsDatabaseToolsMcpToolsetGenAiSqlAssistantResourceConfig = DatabaseToolsDatabaseToolsMcpToolsetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_mcp_toolset", "test_database_tools_mcp_toolset", acctest.Optional, acctest.Update, DatabaseToolsDatabaseToolsMcpToolsetGenAiSqlAssistantRepresentation)

	DatabaseToolsDatabaseToolsMcpToolsetGenAiSqlAssistantSingularDataSourceRepresentation = map[string]interface{}{
		"database_tools_mcp_toolset_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_tools_database_tools_mcp_toolset.test_database_tools_mcp_toolset.id}`},
	}

	DatabaseToolsDatabaseToolsMcpToolsetGenAiSqlAssistantDataSourceRepresentation = map[string]interface{}{
		"compartment_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"database_tools_mcp_server_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.database_tools_mcp_server_id}`},
		"display_name":                 acctest.Representation{RepType: acctest.Optional, Create: `displayName2`},
		"state":                        acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"type":                         acctest.Representation{RepType: acctest.Optional, Create: []string{`GENAI_SQL_ASSISTANT`}},
		"filter":                       acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseToolsDatabaseToolsMcpToolsetDataSourceFilterRepresentation},
	}

	// generativeAiSemanticStoreId is required by the spec for GENAI_SQL_ASSISTANT.
	DatabaseToolsDatabaseToolsMcpToolsetGenAiSqlAssistantRepresentation = map[string]interface{}{
		// Required (base)
		"compartment_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"database_tools_mcp_server_id": acctest.Representation{RepType: acctest.Required, Create: `${var.database_tools_mcp_server_id}`},
		"display_name":                 acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"type":                         acctest.Representation{RepType: acctest.Required, Create: `GENAI_SQL_ASSISTANT`},
		"version":                      acctest.Representation{RepType: acctest.Required, Create: `1`},
		// Required (type-specific)
		"generative_ai_semantic_store_id": acctest.Representation{RepType: acctest.Required, Create: `${var.generative_ai_semantic_store_id}`},
		// Optional
		"default_execution_type": acctest.Representation{RepType: acctest.Optional, Create: `SYNCHRONOUS`, Update: `ASYNCHRONOUS`},
		"description":            acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":          acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"tools": acctest.RepresentationGroup{RepType: acctest.Optional, Group: acctest.RepresentationCopyWithNewProperties(DatabaseToolsDatabaseToolsMcpToolsetToolsRepresentation, map[string]interface{}{
			"name": acctest.Representation{RepType: acctest.Required, Create: `dbtools_translate_natural_language_query_to_sql`},
		})},
	}
)

// issue-routing-tag: database_tools/default
func TestDatabaseToolsDatabaseToolsMcpToolsetGenAiSqlAssistantResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseToolsDatabaseToolsMcpToolsetGenAiSqlAssistantResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	relatedResourceId := utils.GetEnvSettingWithBlankDefault("TF_VAR_related_resource_id")
	relatedResourceIdVariableStr := fmt.Sprintf("variable \"related_resource_id\" { default = \"%s\" }\n", relatedResourceId)

	mcpServerId := utils.GetEnvSettingWithBlankDefault("database_tools_mcp_server_ocid")
	mcpServerIdVariableStr := fmt.Sprintf("variable \"database_tools_mcp_server_id\" { default = \"%s\" }\n", mcpServerId)

	semanticStoreId := utils.GetEnvSettingWithBlankDefault("generative_ai_semantic_store_ocid")
	semanticStoreIdVariableStr := fmt.Sprintf("variable \"generative_ai_semantic_store_id\" { default = \"%s\" }\n", semanticStoreId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_database_tools_database_tools_mcp_toolset.test_database_tools_mcp_toolset"
	datasourceName := "data.oci_database_tools_database_tools_mcp_toolsets.test_database_tools_mcp_toolsets"
	singularDatasourceName := "data.oci_database_tools_database_tools_mcp_toolset.test_database_tools_mcp_toolset"

	commonVars := config + compartmentIdVariableStr + relatedResourceIdVariableStr + mcpServerIdVariableStr + semanticStoreIdVariableStr
	commonVarsU := commonVars + compartmentIdUVariableStr

	var resId, resId2 string
	acctest.SaveConfigContent(commonVars+DatabaseToolsDatabaseToolsMcpToolsetResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_mcp_toolset", "test_database_tools_mcp_toolset", acctest.Optional, acctest.Create, DatabaseToolsDatabaseToolsMcpToolsetGenAiSqlAssistantRepresentation),
		"databasetools", "databaseToolsMcpToolsetGenAiSqlAssistant", t)

	acctest.ResourceTest(t, testAccCheckDatabaseToolsDatabaseToolsMcpToolsetDestroy, []resource.TestStep{
		// Step 1: verify Create with Required only
		{
			Config: commonVars + DatabaseToolsDatabaseToolsMcpToolsetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_mcp_toolset", "test_database_tools_mcp_toolset", acctest.Required, acctest.Create, DatabaseToolsDatabaseToolsMcpToolsetGenAiSqlAssistantRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "database_tools_mcp_server_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "type", "GENAI_SQL_ASSISTANT"),
				resource.TestCheckResourceAttr(resourceName, "version", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "generative_ai_semantic_store_id"),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// Step 2: delete before next Create
		{
			Config: commonVars + DatabaseToolsDatabaseToolsMcpToolsetResourceDependencies,
		},
		// Step 3: verify Create with optionals
		{
			Config: commonVars + DatabaseToolsDatabaseToolsMcpToolsetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_mcp_toolset", "test_database_tools_mcp_toolset", acctest.Optional, acctest.Create, DatabaseToolsDatabaseToolsMcpToolsetGenAiSqlAssistantRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "database_tools_mcp_server_id"),
				resource.TestCheckResourceAttr(resourceName, "default_execution_type", "SYNCHRONOUS"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "generative_ai_semantic_store_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "tools.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tools.0.name", "dbtools_translate_natural_language_query_to_sql"),
				resource.TestCheckResourceAttr(resourceName, "tools.0.status", "ENABLED"),
				resource.TestCheckResourceAttr(resourceName, "type", "GENAI_SQL_ASSISTANT"),
				resource.TestCheckResourceAttr(resourceName, "version", "1"),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},
		// Step 4: verify Update to compartment (switched back in next step)
		{
			Config: commonVarsU + DatabaseToolsDatabaseToolsMcpToolsetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_mcp_toolset", "test_database_tools_mcp_toolset", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DatabaseToolsDatabaseToolsMcpToolsetGenAiSqlAssistantRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "type", "GENAI_SQL_ASSISTANT"),
				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
		// Step 5: verify updates to updatable parameters
		{
			Config: commonVars + DatabaseToolsDatabaseToolsMcpToolsetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_mcp_toolset", "test_database_tools_mcp_toolset", acctest.Optional, acctest.Update, DatabaseToolsDatabaseToolsMcpToolsetGenAiSqlAssistantRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "database_tools_mcp_server_id"),
				resource.TestCheckResourceAttr(resourceName, "default_execution_type", "ASYNCHRONOUS"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "generative_ai_semantic_store_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "tools.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tools.0.name", "dbtools_translate_natural_language_query_to_sql"),
				resource.TestCheckResourceAttr(resourceName, "tools.0.status", "DISABLED"),
				resource.TestCheckResourceAttr(resourceName, "type", "GENAI_SQL_ASSISTANT"),
				resource.TestCheckResourceAttr(resourceName, "version", "1"),
				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// Step 6: verify datasource (list)
		{
			Config: commonVars + DatabaseToolsDatabaseToolsMcpToolsetResourceDependencies +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_mcp_toolsets", "test_database_tools_mcp_toolsets", acctest.Optional, acctest.Update, DatabaseToolsDatabaseToolsMcpToolsetGenAiSqlAssistantDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_mcp_toolset", "test_database_tools_mcp_toolset", acctest.Optional, acctest.Update, DatabaseToolsDatabaseToolsMcpToolsetGenAiSqlAssistantRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "database_tools_mcp_server_id"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "type.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "database_tools_mcp_toolset_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "database_tools_mcp_toolset_collection.0.items.#", "1"),
			),
		},
		// Step 7: verify singular datasource
		{
			Config: commonVars + DatabaseToolsDatabaseToolsMcpToolsetGenAiSqlAssistantResourceConfig +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_mcp_toolset", "test_database_tools_mcp_toolset", acctest.Required, acctest.Create, DatabaseToolsDatabaseToolsMcpToolsetGenAiSqlAssistantSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "database_tools_mcp_toolset_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "default_execution_type", "ASYNCHRONOUS"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "generative_ai_semantic_store_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tools.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tools.0.name", "dbtools_translate_natural_language_query_to_sql"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "tools.0.display_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tools.0.status", "DISABLED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "type", "GENAI_SQL_ASSISTANT"),
				resource.TestCheckResourceAttr(singularDatasourceName, "version", "1"),
			),
		},
		// Step 8: verify resource import
		{
			Config:                  commonVars + DatabaseToolsDatabaseToolsMcpToolsetGenAiSqlAssistantRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

// ============================================================
// Destroy check and sweeper (shared across all 4 tests)
// ============================================================

func testAccCheckDatabaseToolsDatabaseToolsMcpToolsetDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DatabaseToolsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_tools_database_tools_mcp_toolset" {
			noResourceFound = false
			request := oci_database_tools.GetDatabaseToolsMcpToolsetRequest{}

			tmp := rs.Primary.ID
			request.DatabaseToolsMcpToolsetId = &tmp
			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_tools")

			response, err := client.GetDatabaseToolsMcpToolset(context.Background(), request)
			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database_tools.DatabaseToolsMcpToolsetLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.GetLifecycleState())]; !ok {
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.GetLifecycleState())
				}
				continue
			}
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}
	return nil
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("DatabaseToolsDatabaseToolsMcpToolset") {
		resource.AddTestSweepers("DatabaseToolsDatabaseToolsMcpToolset", &resource.Sweeper{
			Name:         "DatabaseToolsDatabaseToolsMcpToolset",
			Dependencies: acctest.DependencyGraph["databaseToolsMcpToolset"],
			F:            sweepDatabaseToolsDatabaseToolsMcpToolsetResource,
		})
	}
}

func sweepDatabaseToolsDatabaseToolsMcpToolsetResource(compartment string) error {
	databaseToolsClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseToolsClient()
	databaseToolsMcpToolsetIds, err := getDatabaseToolsDatabaseToolsMcpToolsetIds(compartment)
	if err != nil {
		return err
	}
	for _, databaseToolsMcpToolsetId := range databaseToolsMcpToolsetIds {
		if ok := acctest.SweeperDefaultResourceId[databaseToolsMcpToolsetId]; !ok {
			deleteDatabaseToolsMcpToolsetRequest := oci_database_tools.DeleteDatabaseToolsMcpToolsetRequest{}
			deleteDatabaseToolsMcpToolsetRequest.DatabaseToolsMcpToolsetId = &databaseToolsMcpToolsetId
			deleteDatabaseToolsMcpToolsetRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_tools")
			_, error := databaseToolsClient.DeleteDatabaseToolsMcpToolset(context.Background(), deleteDatabaseToolsMcpToolsetRequest)
			if error != nil {
				fmt.Printf("Error deleting DatabaseToolsMcpToolset %s %s, It is possible that the resource is already deleted. Please verify manually \n", databaseToolsMcpToolsetId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &databaseToolsMcpToolsetId, DatabaseToolsDatabaseToolsMcpToolsetSweepWaitCondition, time.Duration(3*time.Minute),
				DatabaseToolsDatabaseToolsMcpToolsetSweepResponseFetchOperation, "database_tools", true)
		}
	}
	return nil
}

func getDatabaseToolsDatabaseToolsMcpToolsetIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DatabaseToolsMcpToolsetId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseToolsClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseToolsClient()

	listDatabaseToolsMcpToolsetsRequest := oci_database_tools.ListDatabaseToolsMcpToolsetsRequest{}
	listDatabaseToolsMcpToolsetsRequest.CompartmentId = &compartmentId
	listDatabaseToolsMcpToolsetsRequest.LifecycleState = oci_database_tools.ListDatabaseToolsMcpToolsetsLifecycleStateActive
	listDatabaseToolsMcpToolsetsResponse, err := databaseToolsClient.ListDatabaseToolsMcpToolsets(context.Background(), listDatabaseToolsMcpToolsetsRequest)
	if err != nil {
		return resourceIds, fmt.Errorf("Error getting DatabaseToolsMcpToolset list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, databaseToolsMcpToolset := range listDatabaseToolsMcpToolsetsResponse.Items {
		id := *databaseToolsMcpToolset.GetId()
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DatabaseToolsMcpToolsetId", id)
	}
	return resourceIds, nil
}

func DatabaseToolsDatabaseToolsMcpToolsetSweepWaitCondition(response common.OCIOperationResponse) bool {
	if databaseToolsMcpToolsetResponse, ok := response.Response.(oci_database_tools.GetDatabaseToolsMcpToolsetResponse); ok {
		return databaseToolsMcpToolsetResponse.GetLifecycleState() != oci_database_tools.DatabaseToolsMcpToolsetLifecycleStateDeleted
	}
	return false
}

func DatabaseToolsDatabaseToolsMcpToolsetSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DatabaseToolsClient().GetDatabaseToolsMcpToolset(context.Background(), oci_database_tools.GetDatabaseToolsMcpToolsetRequest{
		DatabaseToolsMcpToolsetId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
