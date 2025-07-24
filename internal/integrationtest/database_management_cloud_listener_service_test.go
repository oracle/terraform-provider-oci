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
	DatabaseManagementCloudListenerServiceDataSourceRepresentation = map[string]interface{}{
		"cloud_listener_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.cloud_listener_id}`},
		"managed_database_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.managed_database_id}`},
		"opc_named_credential_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.opc_named_credential_id}`},
	}

	DatabaseManagementCloudListenerServiceResourceConfig = ""
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementCloudListenerServiceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementCloudListenerServiceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	dbSystemId := utils.GetEnvSettingWithBlankDefault("dbaas_dbsystem_id")
	dbSystemIdVariableStr := fmt.Sprintf("variable \"dbaas_dbsystem_id\" { default = \"%s\" }\n", dbSystemId)

	cloudListenerId := utils.GetEnvSettingWithBlankDefault("dbmgmt_cloud_listener_id")
	cloudListenerIdVariableStr := fmt.Sprintf("variable \"cloud_listener_id\" { default = \"%s\" }\n", cloudListenerId)

	managedDatabaseId := utils.GetEnvSettingWithBlankDefault("dbmgmt_managed_database_id")
	managedDatabaseIdVariableStr := fmt.Sprintf("variable \"managed_database_id\" { default = \"%s\" }\n", managedDatabaseId)

	opcNamedCredentialId := utils.GetEnvSettingWithBlankDefault("dbmgmt_named_credential_id")
	opcNamedCredentialIdStr := fmt.Sprintf("variable \"opc_named_credential_id\" { default = \"%s\" }\n", opcNamedCredentialId)

	datasourceName := "data.oci_database_management_cloud_listener_services.test_cloud_listener_services"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_cloud_listener_services", "test_cloud_listener_services", acctest.Required, acctest.Create, DatabaseManagementCloudListenerServiceDataSourceRepresentation) +
				compartmentIdVariableStr + dbSystemIdVariableStr + cloudListenerIdVariableStr + managedDatabaseIdVariableStr + DatabaseManagementCloudListenerServiceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_listener_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "managed_database_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_listener_service_collection.#"),
			),
		},
		//verify datasource with named credential
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_cloud_listener_services", "test_cloud_listener_services", acctest.Optional, acctest.Create, DatabaseManagementCloudListenerServiceDataSourceRepresentation) +
				compartmentIdVariableStr + dbSystemIdVariableStr + cloudListenerIdVariableStr + managedDatabaseIdVariableStr + opcNamedCredentialIdStr + DatabaseManagementCloudListenerServiceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_listener_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "managed_database_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "opc_named_credential_id"),
			),
		},
	})
}
