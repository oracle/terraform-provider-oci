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
	DatabaseManagementExternalContainerDbmFeaturesManagementRepresentation = map[string]interface{}{
		"external_container_database_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.external_cdb_id}`},
		"feature_details":                       acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseManagementExternalContainerDbmFeaturesManagementFeatureDetailsRepresentation},
		"enable_external_container_dbm_feature": acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `false`},
	}
	DatabaseManagementExternalContainerDbmFeaturesManagementFeatureDetailsRepresentation = map[string]interface{}{
		"connector_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseManagementExternalContainerDbmFeaturesManagementFeatureDetailsConnectorDetailsRepresentation},
		"feature":           acctest.Representation{RepType: acctest.Required, Create: `DIAGNOSTICS_AND_MANAGEMENT`},
		"license_model":     acctest.Representation{RepType: acctest.Required, Create: `LICENSE_INCLUDED`},
	}
	DatabaseManagementExternalContainerDbmFeaturesManagementFeatureDetailsConnectorDetailsRepresentation = map[string]interface{}{
		"connector_type":        acctest.Representation{RepType: acctest.Required, Create: `EXTERNAL`},
		"database_connector_id": acctest.Representation{RepType: acctest.Required, Create: `${var.connector_id}`},
		"management_agent_id":   acctest.Representation{RepType: acctest.Optional, Create: `${oci_management_agent_management_agent.test_management_agent.id}`},
		"private_end_point_id":  acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_management_private_end_point.test_private_end_point.id}`},
	}

	DatabaseManagementExternalContainerDbmDBLMFeaturesManagementRepresentation = map[string]interface{}{
		"external_container_database_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.external_cdb_id}`},
		"feature_details":                       acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseManagementExternalContainerDbmDBLMFeaturesManagementFeatureDetailsRepresentation},
		"enable_external_container_dbm_feature": acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `false`},
	}
	DatabaseManagementExternalContainerDbmDBLMFeaturesManagementFeatureDetailsRepresentation = map[string]interface{}{
		"connector_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseManagementExternalContainerDbmFeaturesManagementFeatureDetailsConnectorDetailsRepresentation},
		"feature":           acctest.Representation{RepType: acctest.Required, Create: `DB_LIFECYCLE_MANAGEMENT`},
		"license_model":     acctest.Representation{RepType: acctest.Required, Create: `LICENSE_INCLUDED`},
	}

	DatabaseManagementExternalContainerDbmSQLWatchFeaturesManagementRepresentation = map[string]interface{}{
		"external_container_database_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.external_cdb_id}`},
		"feature_details":                       acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseManagementExternalContainerDbmSQLWatchFeaturesManagementFeatureDetailsRepresentation},
		"enable_external_container_dbm_feature": acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `false`},
	}
	DatabaseManagementExternalContainerDbmSQLWatchFeaturesManagementFeatureDetailsRepresentation = map[string]interface{}{
		"connector_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseManagementExternalContainerDbmFeaturesManagementFeatureDetailsConnectorDetailsRepresentation},
		"feature":           acctest.Representation{RepType: acctest.Required, Create: `SQLWATCH`},
		"license_model":     acctest.Representation{RepType: acctest.Required, Create: `LICENSE_INCLUDED`},
	}

	ExternalcontainerdatabaseExternalContainerDbmFeaturesManagementResourceDependencies = ""
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementExternalcontainerdatabaseExternalContainerDbmFeaturesManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementExternalcontainerdatabaseExternalContainerDbmFeaturesManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	//log.Printf("[INFO] External CDB OCID is %v", compartmentId)

	externalCdbId := utils.GetEnvSettingWithBlankDefault("dbmgmt_external_cdb_id")
	externalCdbIdStr := fmt.Sprintf("variable \"external_cdb_id\" { default = \"%s\" }\n", externalCdbId)
	//log.Printf("[INFO] External CDB OCID is %v", externalCdbId)

	connectorId := utils.GetEnvSettingWithBlankDefault("dbmgmt_cdb_connector_id")
	connectorIdStr := fmt.Sprintf("variable \"connector_id\" { default = \"%s\" }\n", connectorId)
	//log.Printf("[INFO] Connector OCID is %v", connectorId)

	externalVariableStr := compartmentIdVariableStr + externalCdbIdStr + connectorIdStr

	resourceName := "oci_database_management_externalcontainerdatabase_external_container_dbm_features_management.test_externalcontainerdatabase_external_container_dbm_features_management"
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+externalVariableStr+ExternalcontainerdatabaseExternalContainerDbmFeaturesManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_externalcontainerdatabase_external_container_dbm_features_management", "test_externalcontainerdatabase_external_container_dbm_features_management", acctest.Required, acctest.Create, DatabaseManagementExternalContainerDbmFeaturesManagementRepresentation), "databasemanagement", "externalcontainerdatabaseExternalContainerDbmFeaturesManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// create with enable DIAGNOSTICS_AND_MANAGEMENT
		{
			Config: config + externalVariableStr + ExternalcontainerdatabaseExternalContainerDbmFeaturesManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_externalcontainerdatabase_external_container_dbm_features_management", "test_externalcontainerdatabase_external_container_dbm_features_management", acctest.Required, acctest.Create, DatabaseManagementExternalContainerDbmFeaturesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "external_container_database_id"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.0.connector_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.0.connector_details.0.connector_type", "EXTERNAL"),
				resource.TestCheckResourceAttrSet(resourceName, "feature_details.0.connector_details.0.database_connector_id"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.0.feature", "DIAGNOSTICS_AND_MANAGEMENT"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.0.license_model", "LICENSE_INCLUDED"),
			),
		},
		// update to disable DIAGNOSTICS_AND_MANAGEMENT
		{
			Config: config + externalVariableStr + ExternalcontainerdatabaseExternalContainerDbmFeaturesManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_externalcontainerdatabase_external_container_dbm_features_management", "test_externalcontainerdatabase_external_container_dbm_features_management", acctest.Required, acctest.Update, DatabaseManagementExternalContainerDbmFeaturesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "external_container_database_id"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.0.feature", "DIAGNOSTICS_AND_MANAGEMENT"),
			),
		},
		/* Commenting as we do not have a release date for DBLM
		// create with enable DBLM
		{
			Config: config + externalVariableStr + ExternalcontainerdatabaseExternalContainerDbmFeaturesManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_externalcontainerdatabase_external_container_dbm_features_management", "test_externalcontainerdatabase_external_container_dbm_features_management", acctest.Required, acctest.Create, DatabaseManagementExternalContainerDbmDBLMFeaturesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "external_container_database_id"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.0.connector_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.0.connector_details.0.connector_type", "EXTERNAL"),
				resource.TestCheckResourceAttrSet(resourceName, "feature_details.0.connector_details.0.database_connector_id"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.0.feature", "DB_LIFECYCLE_MANAGEMENT"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.0.license_model", "LICENSE_INCLUDED"),
			),
		},
		// update to disable DBLM
		{
			Config: config + externalVariableStr + ExternalcontainerdatabaseExternalContainerDbmFeaturesManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_externalcontainerdatabase_external_container_dbm_features_management", "test_externalcontainerdatabase_external_container_dbm_features_management", acctest.Required, acctest.Update, DatabaseManagementExternalContainerDbmDBLMFeaturesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "external_container_database_id"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.0.feature", "DB_LIFECYCLE_MANAGEMENT"),
			),
		},
		*/
		// create with enable SQLWATCH
		{
			Config: config + externalVariableStr + ExternalcontainerdatabaseExternalContainerDbmFeaturesManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_externalcontainerdatabase_external_container_dbm_features_management", "test_externalcontainerdatabase_external_container_dbm_features_management", acctest.Required, acctest.Create, DatabaseManagementExternalContainerDbmSQLWatchFeaturesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "external_container_database_id"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.0.connector_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.0.connector_details.0.connector_type", "EXTERNAL"),
				resource.TestCheckResourceAttrSet(resourceName, "feature_details.0.connector_details.0.database_connector_id"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.0.feature", "SQLWATCH"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.0.license_model", "LICENSE_INCLUDED"),
			),
		},
		// update to disable SQLWATCH
		{
			Config: config + externalVariableStr + ExternalcontainerdatabaseExternalContainerDbmFeaturesManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_externalcontainerdatabase_external_container_dbm_features_management", "test_externalcontainerdatabase_external_container_dbm_features_management", acctest.Required, acctest.Update, DatabaseManagementExternalContainerDbmSQLWatchFeaturesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "external_container_database_id"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.0.feature", "SQLWATCH"),
			),
		},
	})
}
