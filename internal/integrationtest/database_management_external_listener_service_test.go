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
	DatabaseManagementDatabaseManagementExternalListenerServiceDataSourceRepresentation = map[string]interface{}{
		"external_listener_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.external_listener_id}`},
		"managed_database_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.managed_database_id}`},
		"opc_named_credential_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.opc_named_credential_id}`},
	}

	DatabaseManagementExternalListenerServiceResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_external_listeners", "test_external_listeners", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementExternalListenerDataSourceRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_external_listener", "test_external_listener", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementExternalListenerSingularDataSourceRepresentation)
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementExternalListenerServiceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementExternalListenerServiceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("dbmgmt_compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	dbSystemId := utils.GetEnvSettingWithBlankDefault("dbmgmt_external_dbsystem_id")
	dbSystemIdVariableStr := fmt.Sprintf("variable \"external_dbsystem_id\" { default = \"%s\" }\n", dbSystemId)

	externalListenerId := utils.GetEnvSettingWithBlankDefault("dbmgmt_external_listener_id")
	externalListenerIdVariableStr := fmt.Sprintf("variable \"external_listener_id\" { default = \"%s\" }\n", externalListenerId)

	managedDatabaseId := utils.GetEnvSettingWithBlankDefault("dbmgmt_managed_database_id")
	managedDatabaseIdVariableStr := fmt.Sprintf("variable \"managed_database_id\" { default = \"%s\" }\n", managedDatabaseId)

	opcNamedCredentialId := utils.GetEnvSettingWithBlankDefault("dbmgmt_named_credential_id")
	opcNamedCredentialIdStr := fmt.Sprintf("variable \"opc_named_credential_id\" { default = \"%s\" }\n", opcNamedCredentialId)

	datasourceName := "data.oci_database_management_external_listener_services.test_external_listener_services"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_external_listener_services", "test_external_listener_services", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementExternalListenerServiceDataSourceRepresentation) +
				compartmentIdVariableStr + dbSystemIdVariableStr + externalListenerIdVariableStr + managedDatabaseIdVariableStr + DatabaseManagementExternalListenerServiceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "external_listener_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "managed_database_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "external_listener_service_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "external_listener_service_collection.0.items.#"),
			),
		},
		// verify datasource with named credential
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_external_listener_services", "test_external_listener_services", acctest.Optional, acctest.Create, DatabaseManagementDatabaseManagementExternalListenerServiceDataSourceRepresentation) +
				compartmentIdVariableStr + dbSystemIdVariableStr + externalListenerIdVariableStr + managedDatabaseIdVariableStr + opcNamedCredentialIdStr + DatabaseManagementExternalListenerServiceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "external_listener_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "managed_database_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "opc_named_credential_id"),
			),
		},
	})
}
