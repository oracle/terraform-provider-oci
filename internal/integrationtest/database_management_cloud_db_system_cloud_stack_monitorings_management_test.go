// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	CloudDbSystemCloudStackMonitoringsManagementRequiredOnlyResource = CloudDbSystemCloudStackMonitoringsManagementResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_db_system_cloud_stack_monitorings_management", "test_cloud_db_system_cloud_stack_monitorings_management", acctest.Required, acctest.Create, DatabaseManagementCloudDbSystemCloudStackMonitoringsManagementRepresentation)

	DatabaseManagementCloudDbSystemCloudStackMonitoringsManagementRepresentation = map[string]interface{}{
		"cloud_db_system_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.dbaas_dbsystem_id}`},
		"is_enabled":                    acctest.Representation{RepType: acctest.Required, Create: `true`},
		"enable_cloud_stack_monitoring": acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `false`},
		"metadata":                      acctest.Representation{RepType: acctest.Optional, Create: `{ `},
	}

	CloudDbSystemCloudStackMonitoringsManagementResourceDependencies = ""
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementCloudDbSystemCloudStackMonitoringsManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementCloudDbSystemCloudStackMonitoringsManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("dbmgmt_compartment_id")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	dbSystemId := utils.GetEnvSettingWithBlankDefault("dbaas_dbsystem_id")
	dbSystemIdVariableStr := fmt.Sprintf("variable \"dbaas_dbsystem_id\" { default = \"%s\" }\n", dbSystemId)

	resourceName := "oci_database_management_cloud_db_system_cloud_stack_monitorings_management.test_cloud_db_system_cloud_stack_monitorings_management"
	//parentResourceName := "oci_database_management_cloud_db_system_cloud_stack_monitorings_management.test_cloud_db_system_cloud_stack_monitorings_management"
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CloudDbSystemCloudStackMonitoringsManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_db_system_cloud_stack_monitorings_management", "test_cloud_db_system_cloud_stack_monitorings_management", acctest.Optional, acctest.Create, DatabaseManagementCloudDbSystemCloudStackMonitoringsManagementRepresentation), "databasemanagement", "cloudDbSystemCloudStackMonitoringsManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// Sample test, All tests below the first one needs to be uncommented after stack monitoring is enabled for cloud dbsystem
		{
			Config: config + compartmentIdVariableStr + dbSystemIdVariableStr + CloudDbSystemCloudStackMonitoringsManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_db_system_cloud_stack_monitorings_management", "test_cloud_db_system_cloud_stack_monitorings_management", acctest.Required, acctest.Create, DatabaseManagementCloudDbSystemCloudStackMonitoringsManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cloud_db_system_id"),
			),
		},

		/*
			// create with enable
			{
				Config: config + compartmentIdVariableStr + dbSystemIdVariableStr + CloudDbSystemCloudStackMonitoringsManagementResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_db_system_cloud_stack_monitorings_management", "test_cloud_db_system_cloud_stack_monitorings_management", acctest.Required, acctest.Create, DatabaseManagementCloudDbSystemCloudStackMonitoringsManagementRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "cloud_db_system_id"),
				),
			},
			// verify enable
			{
				Config: config + compartmentIdVariableStr + dbSystemIdVariableStr + CloudDbSystemCloudStackMonitoringsManagementResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_db_system_cloud_stack_monitorings_management", "test_cloud_db_system_cloud_stack_monitorings_management", acctest.Required, acctest.Create, DatabaseManagementCloudDbSystemCloudStackMonitoringsManagementRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(parentResourceName, "enable_cloud_stack_monitoring", "true"),
				),
			},
			// delete before next Create
			{
				Config: config + compartmentIdVariableStr + dbSystemIdVariableStr + CloudDbSystemCloudStackMonitoringsManagementResourceDependencies,
			},
			// create with enable and optional fields
			{
				Config: config + compartmentIdVariableStr + dbSystemIdVariableStr + CloudDbSystemCloudStackMonitoringsManagementResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_db_system_cloud_stack_monitorings_management", "test_cloud_db_system_cloud_stack_monitorings_management", acctest.Optional, acctest.Create, DatabaseManagementCloudDbSystemCloudStackMonitoringsManagementRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "cloud_db_system_id"),
				),
			},
			// update to disable
			{
				Config: config + compartmentIdVariableStr + dbSystemIdVariableStr + CloudDbSystemCloudStackMonitoringsManagementResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_db_system_cloud_stack_monitorings_management", "test_cloud_db_system_cloud_stack_monitorings_management", acctest.Optional, acctest.Update, DatabaseManagementCloudDbSystemCloudStackMonitoringsManagementRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "cloud_db_system_id"),
				),
			},
			// verify disable
			{
				Config: config + compartmentIdVariableStr + dbSystemIdVariableStr + CloudDbSystemCloudStackMonitoringsManagementResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_db_system_cloud_stack_monitorings_management", "test_cloud_db_system_cloud_stack_monitorings_management", acctest.Optional, acctest.Update, DatabaseManagementCloudDbSystemCloudStackMonitoringsManagementRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(parentResourceName, "enable_cloud_stack_monitoring", "false"),
				),
			},
		*/
	})
}
