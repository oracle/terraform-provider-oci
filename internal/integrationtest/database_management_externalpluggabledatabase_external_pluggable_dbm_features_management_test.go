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
	DatabaseManagementExternalPluggableDbmFeaturesManagementRepresentation = map[string]interface{}{
		"external_pluggable_database_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.external_pdb_id}`},
		"feature_details":                       acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseManagementExternalPluggableDbmFeaturesManagementFeatureDetailsRepresentation},
		"enable_external_pluggable_dbm_feature": acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `false`},
	}
	DatabaseManagementExternalPluggableDbmFeaturesManagementFeatureDetailsRepresentation = map[string]interface{}{
		"connector_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseManagementExternalPluggableDbmFeaturesManagementFeatureDetailsConnectorDetailsRepresentation},
		"feature":           acctest.Representation{RepType: acctest.Required, Create: `DIAGNOSTICS_AND_MANAGEMENT`},
	}
	DatabaseManagementExternalPluggableDbmFeaturesManagementFeatureDetailsConnectorDetailsRepresentation = map[string]interface{}{
		"connector_type":        acctest.Representation{RepType: acctest.Required, Create: `EXTERNAL`},
		"database_connector_id": acctest.Representation{RepType: acctest.Required, Create: `${var.connector_id}`},
		"management_agent_id":   acctest.Representation{RepType: acctest.Optional, Create: `${oci_management_agent_management_agent.test_management_agent.id}`},
		"private_end_point_id":  acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_management_private_end_point.test_private_end_point.id}`},
	}

	ExternalpluggabledatabaseExternalPluggableDbmFeaturesManagementResourceDependencies = ""
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementExternalpluggabledatabaseExternalPluggableDbmFeaturesManagementResource_basic(t *testing.T) {
	t.Skip("Skipping as PDB enablement is dependent on CDB enablement")
	httpreplay.SetScenario("TestDatabaseManagementExternalpluggabledatabaseExternalPluggableDbmFeaturesManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	externalPdbId := utils.GetEnvSettingWithBlankDefault("dbmgmt_external_pdb_id")
	externalPdbIdStr := fmt.Sprintf("variable \"external_pdb_id\" { default = \"%s\" }\n", externalPdbId)
	log.Printf("[INFO] External PDB OCID is %v", externalPdbId)

	connectorId := utils.GetEnvSettingWithBlankDefault("dbmgmt_pdb_connector_id")
	connectorIdStr := fmt.Sprintf("variable \"connector_id\" { default = \"%s\" }\n", connectorId)
	log.Printf("[INFO] External connector OCID is %v", connectorId)

	externalVariableStr := compartmentIdVariableStr + externalPdbIdStr + connectorIdStr

	resourceName := "oci_database_management_externalpluggabledatabase_external_pluggable_dbm_features_management.test_externalpluggabledatabase_external_pluggable_dbm_features_management"
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+externalVariableStr+ExternalpluggabledatabaseExternalPluggableDbmFeaturesManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_externalpluggabledatabase_external_pluggable_dbm_features_management", "test_externalpluggabledatabase_external_pluggable_dbm_features_management", acctest.Required, acctest.Create, DatabaseManagementExternalPluggableDbmFeaturesManagementRepresentation), "databasemanagement", "externalpluggabledatabaseExternalPluggableDbmFeaturesManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// create with enable
		{
			Config: config + externalVariableStr + ExternalpluggabledatabaseExternalPluggableDbmFeaturesManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_externalpluggabledatabase_external_pluggable_dbm_features_management", "test_externalpluggabledatabase_external_pluggable_dbm_features_management", acctest.Required, acctest.Create, DatabaseManagementExternalPluggableDbmFeaturesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "external_pluggable_database_id"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.0.connector_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.0.connector_details.0.connector_type", "EXTERNAL"),
				resource.TestCheckResourceAttrSet(resourceName, "feature_details.0.connector_details.0.database_connector_id"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.0.feature", "DIAGNOSTICS_AND_MANAGEMENT"),
			),
		},
		// Update to disable
		{
			Config: config + externalVariableStr + ExternalpluggabledatabaseExternalPluggableDbmFeaturesManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_externalpluggabledatabase_external_pluggable_dbm_features_management", "test_externalpluggabledatabase_external_pluggable_dbm_features_management", acctest.Required, acctest.Update, DatabaseManagementExternalPluggableDbmFeaturesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "external_pluggable_database_id"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.0.feature", "DIAGNOSTICS_AND_MANAGEMENT"),
			),
		},
	})
}
