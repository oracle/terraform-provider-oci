// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabaseManagementExternalDbSystemStackMonitoringsManagementRepresentation = map[string]interface{}{
		"external_db_system_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.external_dbsystem_id}`},
		"is_enabled":              acctest.Representation{RepType: acctest.Required, Create: `true`},
		"enable_stack_monitoring": acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `false`},
		"metadata":                acctest.Representation{RepType: acctest.Optional, Create: `{ }`},
	}

	ExternalDbSystemStackMonitoringsManagementResourceDependencies = ""
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementExternalDbSystemStackMonitoringsManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementExternalDbSystemStackMonitoringsManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	dbSystemId := utils.GetEnvSettingWithBlankDefault("external_dbsystem_id")
	dbSystemIdVariableStr := fmt.Sprintf("variable \"external_dbsystem_id\" { default = \"%s\" }\n", dbSystemId)

	resourceName := "oci_database_management_external_db_system_stack_monitorings_management.test_external_db_system_stack_monitorings_management"
	parentResourceName := "oci_database_management_external_db_system_stack_monitorings_management.test_external_db_system_stack_monitorings_management"
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ExternalDbSystemStackMonitoringsManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_db_system_stack_monitorings_management", "test_external_db_system_stack_monitorings_management", acctest.Optional, acctest.Create, DatabaseManagementExternalDbSystemStackMonitoringsManagementRepresentation), "databasemanagement", "externalDbSystemStackMonitoringsManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// create with enable
		{
			Config: config + compartmentIdVariableStr + dbSystemIdVariableStr + ExternalDbSystemStackMonitoringsManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_db_system_stack_monitorings_management", "test_external_db_system_stack_monitorings_management", acctest.Required, acctest.Create, DatabaseManagementExternalDbSystemStackMonitoringsManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "external_db_system_id"),
			),
		},
		// verify enable
		{
			Config: config + compartmentIdVariableStr + dbSystemIdVariableStr + ExternalDbSystemStackMonitoringsManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_db_system_stack_monitorings_management", "test_external_db_system_stack_monitorings_management", acctest.Required, acctest.Create, DatabaseManagementExternalDbSystemStackMonitoringsManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(parentResourceName, "enable_stack_monitoring", "true"),
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + dbSystemIdVariableStr + ExternalDbSystemStackMonitoringsManagementResourceDependencies,
		},
		// create with enable and optional fields
		{
			Config: config + compartmentIdVariableStr + dbSystemIdVariableStr + ExternalDbSystemStackMonitoringsManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_db_system_stack_monitorings_management", "test_external_db_system_stack_monitorings_management", acctest.Optional, acctest.Create, DatabaseManagementExternalDbSystemStackMonitoringsManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "external_db_system_id"),
			),
		},
		// update to disable
		{
			Config: config + compartmentIdVariableStr + dbSystemIdVariableStr + ExternalDbSystemStackMonitoringsManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_db_system_stack_monitorings_management", "test_external_db_system_stack_monitorings_management", acctest.Optional, acctest.Update, DatabaseManagementExternalDbSystemStackMonitoringsManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "external_db_system_id"),
			),
		},
		// verify disable
		{
			Config: config + compartmentIdVariableStr + dbSystemIdVariableStr + ExternalDbSystemStackMonitoringsManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_db_system_stack_monitorings_management", "test_external_db_system_stack_monitorings_management", acctest.Optional, acctest.Update, DatabaseManagementExternalDbSystemStackMonitoringsManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(parentResourceName, "enable_stack_monitoring", "false"),
			),
		},
	})
}
