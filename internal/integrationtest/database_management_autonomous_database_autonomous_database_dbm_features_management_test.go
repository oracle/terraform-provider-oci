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
	DatabaseManagementAutonomousDatabaseAutonomousDatabaseDbmFeaturesManagementRepresentation = map[string]interface{}{
		"autonomous_database_id":                 acctest.Representation{RepType: acctest.Required, Create: `${var.autonomous_database_id}`},
		"feature_details":                        acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseManagementAutonomousDatabaseAutonomousDatabaseDbmFeaturesManagementFeatureDetailsRepresentation},
		"enable_autonomous_database_dbm_feature": acctest.Representation{RepType: acctest.Required, Create: `true`},
	}
	DatabaseManagementAutonomousDatabaseAutonomousDatabaseDbmFeaturesManagementRepresentationDisable = map[string]interface{}{
		"autonomous_database_id":                 acctest.Representation{RepType: acctest.Required, Create: `${var.autonomous_database_id}`},
		"feature_details":                        acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseManagementAutonomousDatabaseAutonomousDatabaseDbmFeaturesManagementFeatureDetailsRepresentation},
		"enable_autonomous_database_dbm_feature": acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `false`},
	}
	DatabaseManagementAutonomousDatabaseAutonomousDatabaseDbmFeaturesManagementFeatureDetailsRepresentation = map[string]interface{}{
		"database_connection_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseManagementAutonomousDatabaseAutonomousDatabaseDbmFeaturesManagementFeatureDetailsDatabaseConnectionDetailsRepresentation},
		"feature":                     acctest.Representation{RepType: acctest.Required, Create: `DIAGNOSTICS_AND_MANAGEMENT`},
		"connector_details":           acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseManagementAutonomousDatabaseAutonomousDatabaseDbmFeaturesManagementFeatureDetailsConnectorDetailsRepresentation},
	}
	DatabaseManagementAutonomousDatabaseAutonomousDatabaseDbmFeaturesManagementFeatureDetailsDatabaseConnectionDetailsRepresentation = map[string]interface{}{
		"connection_credentials": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseManagementAutonomousDatabaseAutonomousDatabaseDbmFeaturesManagementFeatureDetailsDatabaseConnectionDetailsConnectionCredentialsRepresentation},
		"connection_string":      acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseManagementAutonomousDatabaseAutonomousDatabaseDbmFeaturesManagementFeatureDetailsDatabaseConnectionDetailsConnectionStringRepresentation},
	}
	DatabaseManagementAutonomousDatabaseAutonomousDatabaseDbmFeaturesManagementFeatureDetailsConnectorDetailsRepresentation = map[string]interface{}{
		// DIRECT - BEGIN
		/*		"connector_type":        acctest.Representation{RepType: acctest.Required, Create: `DIRECT`},
				"database_connector_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_management_database_connector.test_database_connector.id}`},
				"management_agent_id":   acctest.Representation{RepType: acctest.Optional, Create: `${oci_management_agent_management_agent.test_management_agent.id}`},
				"private_end_point_id":  acctest.Representation{RepType: acctest.Optional, Create: `${var.private_end_point_id}`},*/
		// DIRECT - END
		// PE - BEGIN
		"connector_type":        acctest.Representation{RepType: acctest.Required, Create: `PE`},
		"database_connector_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_management_database_connector.test_database_connector.id}`},
		"management_agent_id":   acctest.Representation{RepType: acctest.Optional, Create: `${oci_management_agent_management_agent.test_management_agent.id}`},
		"private_end_point_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.private_end_point_id}`},
		// PE - END
	}
	DatabaseManagementAutonomousDatabaseAutonomousDatabaseDbmFeaturesManagementFeatureDetailsDatabaseConnectionDetailsConnectionCredentialsRepresentation = map[string]interface{}{
		"credential_name":    acctest.Representation{RepType: acctest.Required, Create: `credentialName`},
		"credential_type":    acctest.Representation{RepType: acctest.Required, Create: `${var.credential_type}`},
		"password_secret_id": acctest.Representation{RepType: acctest.Required, Create: `${var.password_secret_id}`, Update: `${var.modified_password_secret_id}`},
		"role":               acctest.Representation{RepType: acctest.Required, Create: `NORMAL`},
		// NON-SSL - BEGIN
		/*"ssl_secret_id": acctest.Representation{RepType: acctest.Required, Create: `${var.ssl_secret_id}`},*/
		// NON-SSL - BEGIN
		// SSL - BEGIN
		"ssl_secret_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.ssl_secret_id}`},
		// SSL - BEGIN
		"user_name": acctest.Representation{RepType: acctest.Required, Create: `ADMIN`},
	}
	DatabaseManagementAutonomousDatabaseAutonomousDatabaseDbmFeaturesManagementFeatureDetailsDatabaseConnectionDetailsConnectionStringRepresentation = map[string]interface{}{
		"connection_type": acctest.Representation{RepType: acctest.Required, Create: `BASIC`},
		"port":            acctest.Representation{RepType: acctest.Required, Create: `${var.port}`},
		"protocol":        acctest.Representation{RepType: acctest.Required, Create: `${var.protocol}`},
		"service":         acctest.Representation{RepType: acctest.Required, Create: `${var.service_name}`},
	}

	AutonomousDatabaseAutonomousDatabaseDbmFeaturesManagementResourceDependencies = ""
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementAutonomousDatabaseAutonomousDatabaseDbmFeaturesManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementAutonomousDatabaseAutonomousDatabaseDbmFeaturesManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("dbmgmt_cloud_compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	autonomousDatabaseId := utils.GetEnvSettingWithBlankDefault("dbmgmt_autonomous_database_id")
	autonomousDatabaseIdVariableStr := fmt.Sprintf("variable \"autonomous_database_id\" { default = \"%s\" }\n", autonomousDatabaseId)

	privateEndpointId := utils.GetEnvSettingWithBlankDefault("dbmgmt_private_end_point_id")
	privateEndpointIdVariableStr := fmt.Sprintf("variable \"private_end_point_id\" { default = \"%s\" }\n", privateEndpointId)

	pwdSecretId := utils.GetEnvSettingWithBlankDefault("dbmgmt_password_secret_id")
	pwdSecretIdVariableStr := fmt.Sprintf("variable \"password_secret_id\" { default = \"%s\" }\n", pwdSecretId)

	modifiedPwdSecretId := utils.GetEnvSettingWithBlankDefault("dbmgmt_modified_password_secret_id")
	modifiedPwdSecretIdVariableStr := fmt.Sprintf("variable \"modified_password_secret_id\" { default = \"%s\" }\n", modifiedPwdSecretId)

	serviceName := utils.GetEnvSettingWithBlankDefault("dbmgmt_service_name")
	serviceNameVariableStr := fmt.Sprintf("variable \"service_name\" { default = \"%s\" }\n", serviceName)

	protocol := utils.GetEnvSettingWithBlankDefault("dbmgmt_protocol")
	protocolVariableStr := fmt.Sprintf("variable \"protocol\" { default = \"%s\" }\n", protocol)

	port := utils.GetEnvSettingWithBlankDefault("dbmgmt_port")
	portVariableStr := fmt.Sprintf("variable \"port\" { default = \"%s\" }\n", port)

	sslSecretId := utils.GetEnvSettingWithBlankDefault("dbmgmt_ssl_secret_id")
	sslSecretIdVariableStr := fmt.Sprintf("variable \"ssl_secret_id\" { default = \"%s\" }\n", sslSecretId)

	credentialType := utils.GetEnvSettingWithBlankDefault("dbmgmt_credential_type")
	credentialTypeVariableStr := fmt.Sprintf("variable \"credential_type\" { default = \"%s\" }\n", credentialType)

	log.Printf("[INFO] Data is %v", config+compartmentIdVariableStr+credentialTypeVariableStr+sslSecretIdVariableStr+portVariableStr+protocolVariableStr+pwdSecretId+modifiedPwdSecretId+
		serviceNameVariableStr+privateEndpointIdVariableStr+autonomousDatabaseIdVariableStr+AutonomousDatabaseAutonomousDatabaseDbmFeaturesManagementResourceDependencies)
	resourceName := "oci_database_management_autonomous_database_autonomous_database_dbm_features_management.test_autonomous_database_autonomous_database_dbm_features_management"
	parentResourceName := "oci_database_management_autonomous_database_autonomous_database_dbm_features_management.test_autonomous_database_autonomous_database_dbm_features_management"
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+credentialTypeVariableStr+sslSecretIdVariableStr+portVariableStr+protocolVariableStr+pwdSecretIdVariableStr+modifiedPwdSecretIdVariableStr+
		serviceNameVariableStr+privateEndpointIdVariableStr+autonomousDatabaseIdVariableStr+AutonomousDatabaseAutonomousDatabaseDbmFeaturesManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_autonomous_database_autonomous_database_dbm_features_management", "test_autonomous_database_autonomous_database_dbm_features_management", acctest.Required, acctest.Create, DatabaseManagementAutonomousDatabaseAutonomousDatabaseDbmFeaturesManagementRepresentation), "databasemanagement", "autonomousDatabaseAutonomousDatabaseDbmFeaturesManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify enable
		{
			Config: config + compartmentIdVariableStr + credentialTypeVariableStr + sslSecretIdVariableStr + portVariableStr + protocolVariableStr + pwdSecretIdVariableStr + modifiedPwdSecretIdVariableStr +
				serviceNameVariableStr + privateEndpointIdVariableStr + autonomousDatabaseIdVariableStr + AutonomousDatabaseAutonomousDatabaseDbmFeaturesManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_autonomous_database_autonomous_database_dbm_features_management", "test_autonomous_database_autonomous_database_dbm_features_management", acctest.Required, acctest.Create, DatabaseManagementAutonomousDatabaseAutonomousDatabaseDbmFeaturesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_database_id"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.0.database_connection_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.0.database_connection_details.0.connection_credentials.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.0.database_connection_details.0.connection_credentials.0.password_secret_id", pwdSecretId),
				resource.TestCheckResourceAttr(resourceName, "feature_details.0.database_connection_details.0.connection_credentials.0.role", "NORMAL"),
				// SSL - BEGIN
				//resource.TestCheckResourceAttr(resourceName, "feature_details.0.database_connection_details.0.connection_credentials.0.ssl_secret_id", sslSecretId),
				// SSL - END
				resource.TestCheckResourceAttrSet(resourceName, "feature_details.0.database_connection_details.0.connection_credentials.0.user_name"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.0.database_connection_details.0.connection_string.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.0.database_connection_details.0.connection_string.0.connection_type", "BASIC"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.0.database_connection_details.0.connection_string.0.port", port),
				resource.TestCheckResourceAttr(resourceName, "feature_details.0.database_connection_details.0.connection_string.0.protocol", protocol),
				resource.TestCheckResourceAttr(resourceName, "feature_details.0.database_connection_details.0.connection_string.0.service", serviceName),
				resource.TestCheckResourceAttr(resourceName, "feature_details.0.feature", "DIAGNOSTICS_AND_MANAGEMENT"),
			),
		},
		// verify modify
		{
			Config: config + compartmentIdVariableStr + credentialTypeVariableStr + sslSecretIdVariableStr + portVariableStr + protocolVariableStr + pwdSecretIdVariableStr + modifiedPwdSecretIdVariableStr +
				serviceNameVariableStr + privateEndpointIdVariableStr + autonomousDatabaseIdVariableStr + AutonomousDatabaseAutonomousDatabaseDbmFeaturesManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_autonomous_database_autonomous_database_dbm_features_management", "test_autonomous_database_autonomous_database_dbm_features_management", acctest.Required, acctest.Update, DatabaseManagementAutonomousDatabaseAutonomousDatabaseDbmFeaturesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_database_id"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.0.database_connection_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.0.database_connection_details.0.connection_credentials.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.0.database_connection_details.0.connection_credentials.0.password_secret_id", modifiedPwdSecretId),
				resource.TestCheckResourceAttr(resourceName, "feature_details.0.database_connection_details.0.connection_credentials.0.role", "NORMAL"),
				// SSL - BEGIN
				//resource.TestCheckResourceAttr(resourceName, "feature_details.0.database_connection_details.0.connection_credentials.0.ssl_secret_id", sslSecretId),
				// SSL - END
				resource.TestCheckResourceAttrSet(resourceName, "feature_details.0.database_connection_details.0.connection_credentials.0.user_name"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.0.database_connection_details.0.connection_string.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.0.database_connection_details.0.connection_string.0.connection_type", "BASIC"),
				resource.TestCheckResourceAttr(resourceName, "feature_details.0.database_connection_details.0.connection_string.0.port", port),
				resource.TestCheckResourceAttr(resourceName, "feature_details.0.database_connection_details.0.connection_string.0.protocol", protocol),
				resource.TestCheckResourceAttr(resourceName, "feature_details.0.database_connection_details.0.connection_string.0.service", serviceName),
				resource.TestCheckResourceAttr(resourceName, "feature_details.0.feature", "DIAGNOSTICS_AND_MANAGEMENT"),
			),
		},
		// verify disable
		{
			Config: config + compartmentIdVariableStr + credentialTypeVariableStr + sslSecretIdVariableStr + portVariableStr + protocolVariableStr + pwdSecretIdVariableStr + modifiedPwdSecretIdVariableStr +
				serviceNameVariableStr + privateEndpointIdVariableStr + autonomousDatabaseIdVariableStr + AutonomousDatabaseAutonomousDatabaseDbmFeaturesManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_autonomous_database_autonomous_database_dbm_features_management", "test_autonomous_database_autonomous_database_dbm_features_management", acctest.Required, acctest.Update, DatabaseManagementAutonomousDatabaseAutonomousDatabaseDbmFeaturesManagementRepresentationDisable),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(parentResourceName, "enable_autonomous_database_dbm_feature", "false"),
			),
		},
	})
}
