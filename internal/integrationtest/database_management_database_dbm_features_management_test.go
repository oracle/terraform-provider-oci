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
	DatabaseManagementDatabaseDbmFeaturesManagementRepresentation = map[string]interface{}{
		"database_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.cloud_database_id}`},
		"feature_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseManagementDatabaseDbmFeaturesManagementFeatureDetailsRepresentation},
		//Uncomment "Update : 'false'"" to run disable API or comment to run Modify API"
		"enable_database_dbm_feature": acctest.Representation{RepType: acctest.Required, Create: `true` /*, Update: `false`*/},
	}
	DatabaseManagementDatabaseDbmFeaturesManagementFeatureDetailsRepresentation = map[string]interface{}{
		"connector_details":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseManagementDatabaseDbmFeaturesManagementFeatureDetailsConnectorDetailsRepresentation},
		"database_connection_details":       acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseManagementDatabaseDbmFeaturesManagementFeatureDetailsDatabaseConnectionDetailsRepresentation},
		"feature":                           acctest.Representation{RepType: acctest.Required, Create: `DIAGNOSTICS_AND_MANAGEMENT`},
		"management_type":                   acctest.Representation{RepType: acctest.Required, Create: `ADVANCED`},
		"is_auto_enable_pluggable_database": acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}
	DatabaseManagementDatabaseDbmFeaturesManagementFeatureDetailsConnectorDetailsRepresentation = map[string]interface{}{
		"connector_type":        acctest.Representation{RepType: acctest.Required, Create: `PE`},
		"database_connector_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_management_database_connector.test_database_connector.id}`},
		"management_agent_id":   acctest.Representation{RepType: acctest.Optional, Create: `${oci_management_agent_management_agent.test_management_agent.id}`},
		"private_end_point_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.private_end_point_id}`},
	}
	DatabaseManagementDatabaseDbmFeaturesManagementFeatureDetailsDatabaseConnectionDetailsRepresentation = map[string]interface{}{
		"connection_credentials": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseManagementDatabaseDbmFeaturesManagementFeatureDetailsDatabaseConnectionDetailsConnectionCredentialsRepresentation},
		"connection_string":      acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseManagementDatabaseDbmFeaturesManagementFeatureDetailsDatabaseConnectionDetailsConnectionStringRepresentation},
	}
	DatabaseManagementDatabaseDbmFeaturesManagementFeatureDetailsDatabaseConnectionDetailsConnectionCredentialsRepresentation = map[string]interface{}{
		"credential_name":    acctest.Representation{RepType: acctest.Optional, Create: `credentialName`},
		"credential_type":    acctest.Representation{RepType: acctest.Required, Create: `DETAILS`},
		"password_secret_id": acctest.Representation{RepType: acctest.Required, Create: `${var.password_secret_id}`, Update: `${var.modified_password_secret_id}`},
		"role":               acctest.Representation{RepType: acctest.Required, Create: `SYSDBA`},
		"ssl_secret_id":      acctest.Representation{RepType: acctest.Optional, Create: `${oci_vault_secret.test_secret.id}`},
		"user_name":          acctest.Representation{RepType: acctest.Required, Create: `dbsnmp`},
	}
	DatabaseManagementDatabaseDbmFeaturesManagementFeatureDetailsDatabaseConnectionDetailsConnectionStringRepresentation = map[string]interface{}{
		"connection_type": acctest.Representation{RepType: acctest.Required, Create: `BASIC`},
		"port":            acctest.Representation{RepType: acctest.Required, Create: `1521`},
		"protocol":        acctest.Representation{RepType: acctest.Required, Create: `TCP`},
		"service":         acctest.Representation{RepType: acctest.Required, Create: `${var.service_name}`},
	}

	DatabaseDbmFeaturesManagementResourceDependencies = ""
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementDatabaseDbmFeaturesManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementDatabaseDbmFeaturesManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("cloud_compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	cloudDatabaseId := utils.GetEnvSettingWithBlankDefault("cloud_database_id")
	cloudDatabaseIdVariableStr := fmt.Sprintf("variable \"cloud_database_id\" { default = \"%s\" }\n", cloudDatabaseId)

	privateEndpointId := utils.GetEnvSettingWithBlankDefault("private_end_point_id")
	privateEndpointIdVariableStr := fmt.Sprintf("variable \"private_end_point_id\" { default = \"%s\" }\n", privateEndpointId)

	pwdSecretId := utils.GetEnvSettingWithBlankDefault("password_secret_id")
	pwdSecretIdVariableStr := fmt.Sprintf("variable \"password_secret_id\" { default = \"%s\" }\n", pwdSecretId)
	log.Printf("[INFO] pwdSecretIdVariableStr is %v", pwdSecretIdVariableStr)

	modifiedPwdSecretId := utils.GetEnvSettingWithBlankDefault("modified_password_secret_id")
	modifiedPwdSecretIdVariableStr := fmt.Sprintf("variable \"modified_password_secret_id\" { default = \"%s\" }\n", modifiedPwdSecretId)
	log.Printf("[INFO] modifiedPwdSecretIdVariableStr is %v", modifiedPwdSecretIdVariableStr)

	serviceName := utils.GetEnvSettingWithBlankDefault("service_name")
	serviceNameVariableStr := fmt.Sprintf("variable \"service_name\" { default = \"%s\" }\n", serviceName)

	variableStr := compartmentIdVariableStr + cloudDatabaseIdVariableStr + privateEndpointIdVariableStr + pwdSecretIdVariableStr + modifiedPwdSecretIdVariableStr + serviceNameVariableStr

	resourceName := "oci_database_management_database_dbm_features_management.test_database_dbm_features_management"
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+variableStr+DatabaseDbmFeaturesManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_database_dbm_features_management", "test_database_dbm_features_management", acctest.Required, acctest.Create, DatabaseManagementDatabaseDbmFeaturesManagementRepresentation), "databasemanagement", "databaseDbmFeaturesManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// create with enable
		{
			Config: config + variableStr + DatabaseDbmFeaturesManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_database_dbm_features_management", "test_database_dbm_features_management", acctest.Required, acctest.Create, DatabaseManagementDatabaseDbmFeaturesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "database_id"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.0.connector_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.0.connector_details.0.connector_type", "PE"),
				resource.TestCheckResourceAttrSet(resourceName, "feature_details.0.connector_details.0.private_end_point_id"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.0.database_connection_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.0.database_connection_details.0.connection_credentials.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "feature_details.0.database_connection_details.0.connection_credentials.0.password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.0.database_connection_details.0.connection_credentials.0.role", "SYSDBA"),
				resource.TestCheckResourceAttrSet(resourceName, "feature_details.0.database_connection_details.0.connection_credentials.0.user_name"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.0.database_connection_details.0.connection_string.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.0.database_connection_details.0.connection_string.0.port", "1521"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.0.database_connection_details.0.connection_string.0.protocol", "TCP"),
				resource.TestCheckResourceAttrSet(resourceName, "feature_details.0.database_connection_details.0.connection_string.0.service"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.0.feature", "DIAGNOSTICS_AND_MANAGEMENT"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.0.management_type", "ADVANCED"),
			),
		},
		// update to Modify
		{
			Config: config + variableStr + DatabaseDbmFeaturesManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_database_dbm_features_management", "test_database_dbm_features_management", acctest.Required, acctest.Update, DatabaseManagementDatabaseDbmFeaturesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "database_id"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.0.connector_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.0.connector_details.0.connector_type", "PE"),
				resource.TestCheckResourceAttrSet(resourceName, "feature_details.0.connector_details.0.private_end_point_id"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.0.database_connection_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.0.database_connection_details.0.connection_credentials.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "feature_details.0.database_connection_details.0.connection_credentials.0.password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.0.database_connection_details.0.connection_credentials.0.role", "SYSDBA"),
				resource.TestCheckResourceAttrSet(resourceName, "feature_details.0.database_connection_details.0.connection_credentials.0.user_name"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.0.database_connection_details.0.connection_string.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.0.database_connection_details.0.connection_string.0.port", "1521"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.0.database_connection_details.0.connection_string.0.protocol", "TCP"),
				resource.TestCheckResourceAttrSet(resourceName, "feature_details.0.database_connection_details.0.connection_string.0.service"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.0.feature", "DIAGNOSTICS_AND_MANAGEMENT"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.0.management_type", "ADVANCED"),
			),
		},
	})
}
