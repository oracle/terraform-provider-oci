// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabaseManagementExternalNonContainerDbmFeaturesManagementRepresentation = map[string]interface{}{
		"external_non_container_database_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.external_non_cdb_id}`},
		"feature_details":                           acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseManagementExternalNonContainerDbmFeaturesManagementFeatureDetailsRepresentation},
		"enable_external_non_container_dbm_feature": acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `false`},
	}
	DatabaseManagementExternalNonContainerDbmFeaturesManagementFeatureDetailsRepresentation = map[string]interface{}{
		"connector_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseManagementExternalNonContainerDbmFeaturesManagementFeatureDetailsConnectorDetailsRepresentation},
		"feature":           acctest.Representation{RepType: acctest.Required, Create: `DIAGNOSTICS_AND_MANAGEMENT`},
		"license_model":     acctest.Representation{RepType: acctest.Required, Create: `LICENSE_INCLUDED`},
	}
	DatabaseManagementExternalNonContainerDbmFeaturesManagementFeatureDetailsConnectorDetailsRepresentation = map[string]interface{}{
		"connector_type":        acctest.Representation{RepType: acctest.Required, Create: `EXTERNAL`},
		"database_connector_id": acctest.Representation{RepType: acctest.Required, Create: `${var.connector_id}`},
		"management_agent_id":   acctest.Representation{RepType: acctest.Optional, Create: `${oci_management_agent_management_agent.test_management_agent.id}`},
		"private_end_point_id":  acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_management_private_end_point.test_private_end_point.id}`},
	}

	ExternalnoncontainerdatabaseExternalNonContainerDbmFeaturesManagementResourceDependencies = ""
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementExternalnoncontainerdatabaseExternalNonContainerDbmFeaturesManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementExternalnoncontainerdatabaseExternalNonContainerDbmFeaturesManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	externalNonCdbId := utils.GetEnvSettingWithBlankDefault("dbmgmt_external_non_cdb_id")
	externalNonCdbIdStr := fmt.Sprintf("variable \"external_non_cdb_id\" { default = \"%s\" }\n", externalNonCdbId)
	log.Printf("[INFO] External Non CDB OCID is %v", externalNonCdbId)

	connectorId := utils.GetEnvSettingWithBlankDefault("dbmgmt_non_cdb_connector_id")
	connectorIdStr := fmt.Sprintf("variable \"connector_id\" { default = \"%s\" }\n", connectorId)
	log.Printf("[INFO] Connector OCID is %v", connectorId)

	externalVariableStr := compartmentIdVariableStr + externalNonCdbIdStr + connectorIdStr

	resourceName := "oci_database_management_externalnoncontainerdatabase_external_non_container_dbm_features_management.test_externalnoncontainerdatabase_external_non_container_dbm_features_management"
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+externalVariableStr+ExternalnoncontainerdatabaseExternalNonContainerDbmFeaturesManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_externalnoncontainerdatabase_external_non_container_dbm_features_management", "test_externalnoncontainerdatabase_external_non_container_dbm_features_management", acctest.Required, acctest.Create, DatabaseManagementExternalNonContainerDbmFeaturesManagementRepresentation), "databasemanagement", "externalnoncontainerdatabaseExternalNonContainerDbmFeaturesManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// create with enable
		{
			Config: config + externalVariableStr + ExternalnoncontainerdatabaseExternalNonContainerDbmFeaturesManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_externalnoncontainerdatabase_external_non_container_dbm_features_management", "test_externalnoncontainerdatabase_external_non_container_dbm_features_management", acctest.Required, acctest.Create, DatabaseManagementExternalNonContainerDbmFeaturesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "external_non_container_database_id"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.0.connector_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.0.connector_details.0.connector_type", "EXTERNAL"),
				resource.TestCheckResourceAttrSet(resourceName, "feature_details.0.connector_details.0.database_connector_id"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.0.feature", "DIAGNOSTICS_AND_MANAGEMENT"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.0.license_model", "LICENSE_INCLUDED"),
			),
		},
		// Update to disable
		{
			Config: config + externalVariableStr + ExternalnoncontainerdatabaseExternalNonContainerDbmFeaturesManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_externalnoncontainerdatabase_external_non_container_dbm_features_management", "test_externalnoncontainerdatabase_external_non_container_dbm_features_management", acctest.Required, acctest.Update, DatabaseManagementExternalNonContainerDbmFeaturesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "external_non_container_database_id"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.0.feature", "DIAGNOSTICS_AND_MANAGEMENT"),
			),
		},
	})
}
