// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalRequiredOnlyResource = DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_tools_runtime_database_tools_database_api_gateway_config_global", "test_database_tools_database_api_gateway_config_global", acctest.Required, acctest.Create, DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalRepresentation)

	DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalResourceConfig = DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_tools_runtime_database_tools_database_api_gateway_config_global", "test_database_tools_database_api_gateway_config_global", acctest.Optional, acctest.Update, DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalUpdateRepresentation)

	DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalSingularDataSourceRepresentation = map[string]interface{}{
		"database_tools_database_api_gateway_config_id": acctest.Representation{RepType: acctest.Required, Create: `${var.database_tools_database_api_gateway_config_id}`},
		"global_key": acctest.Representation{RepType: acctest.Required, Create: `SETTINGS`},
	}

	DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalRepresentation = map[string]interface{}{
		"database_tools_database_api_gateway_config_id": acctest.Representation{RepType: acctest.Required, Create: `${var.database_tools_database_api_gateway_config_id}`},
		"global_key":          acctest.Representation{RepType: acctest.Required, Create: `SETTINGS`},
		"type":                acctest.Representation{RepType: acctest.Required, Create: `DEFAULT`},
		"certificate_bundle":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalCertificateBundleRepresentation},
		"database_api_status": acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`, Update: `ENABLED`},
		"document_root":       acctest.Representation{RepType: acctest.Optional, Create: `/var/www/webroots`, Update: `/var/www/webroots`},
		"http_port":           acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `0`},
		"https_port":          acctest.Representation{RepType: acctest.Optional, Create: `8443`, Update: `8443`},
		"pool_route":          acctest.Representation{RepType: acctest.Optional, Create: `HEADER`, Update: `HEADER`},
		"pool_routing_header": acctest.Representation{RepType: acctest.Optional, Create: `Host1`, Update: `Host1`},
	}
	DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalUpdateRepresentation = acctest.GetRepresentationCopyWithMultipleRemovedProperties(
		[]string{
			"certificate_bundle.certificate_private_key",
			"certificate_bundle.certificate_public",
		},
		DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalRepresentation,
	)
	DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalCertificateBundleRepresentation = map[string]interface{}{
		"type":                    acctest.Representation{RepType: acctest.Required, Create: `FILENAME`, Update: `SELF_SIGNED`},
		"certificate_private_key": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalCertificateBundleCertificatePrivateKeyRepresentation},
		"certificate_public":      acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalCertificateBundleCertificatePublicRepresentation},
	}
	DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalCertificateBundleCertificatePrivateKeyRepresentation = map[string]interface{}{
		"format": acctest.Representation{RepType: acctest.Optional, Create: `DER`, Update: `PEM`},
		"path":   acctest.Representation{RepType: acctest.Optional, Create: `path`, Update: `path2`},
	}
	DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalCertificateBundleCertificatePublicRepresentation = map[string]interface{}{
		"format": acctest.Representation{RepType: acctest.Optional, Create: `PEM`},
		"path":   acctest.Representation{RepType: acctest.Optional, Create: `path`, Update: `path2`},
	}

	DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalResourceDependencies = ""
)

// issue-routing-tag: database_tools_runtime/default
func TestDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	databaseToolsDatabaseApiGatewayConfigId := utils.GetEnvSettingWithBlankDefault("database_tools_database_api_gateway_config_id")
	if databaseToolsDatabaseApiGatewayConfigId == "" {
		t.Skip("set database_tools_database_api_gateway_config_id to run this test")
	}

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	databaseToolsDatabaseApiGatewayConfigIdVariableStr := terraformStringVariable("database_tools_database_api_gateway_config_id", databaseToolsDatabaseApiGatewayConfigId)

	resourceName := "oci_database_tools_runtime_database_tools_database_api_gateway_config_global.test_database_tools_database_api_gateway_config_global"

	singularDatasourceName := "data.oci_database_tools_runtime_database_tools_database_api_gateway_config_global.test_database_tools_database_api_gateway_config_global"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+databaseToolsDatabaseApiGatewayConfigIdVariableStr+DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_tools_runtime_database_tools_database_api_gateway_config_global", "test_database_tools_database_api_gateway_config_global", acctest.Optional, acctest.Create, DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalRepresentation), "databasetoolsruntime", "databaseToolsDatabaseApiGatewayConfigGlobal", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + databaseToolsDatabaseApiGatewayConfigIdVariableStr + DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_runtime_database_tools_database_api_gateway_config_global", "test_database_tools_database_api_gateway_config_global", acctest.Required, acctest.Create, DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "database_tools_database_api_gateway_config_id"),
				resource.TestCheckResourceAttr(resourceName, "global_key", "SETTINGS"),
				resource.TestCheckResourceAttr(resourceName, "type", "DEFAULT"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + databaseToolsDatabaseApiGatewayConfigIdVariableStr + DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + databaseToolsDatabaseApiGatewayConfigIdVariableStr + DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_runtime_database_tools_database_api_gateway_config_global", "test_database_tools_database_api_gateway_config_global", acctest.Optional, acctest.Create, DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "certificate_bundle.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "certificate_bundle.0.certificate_private_key.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "certificate_bundle.0.certificate_private_key.0.format", "DER"),
				resource.TestCheckResourceAttr(resourceName, "certificate_bundle.0.certificate_private_key.0.path", "path"),
				resource.TestCheckResourceAttr(resourceName, "certificate_bundle.0.certificate_public.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "certificate_bundle.0.certificate_public.0.format", "PEM"),
				resource.TestCheckResourceAttr(resourceName, "certificate_bundle.0.certificate_public.0.path", "path"),
				resource.TestCheckResourceAttr(resourceName, "certificate_bundle.0.type", "FILENAME"),
				resource.TestCheckResourceAttr(resourceName, "database_api_status", "ENABLED"),
				resource.TestCheckResourceAttrSet(resourceName, "database_tools_database_api_gateway_config_id"),
				resource.TestCheckResourceAttr(resourceName, "document_root", "/var/www/webroots"),
				resource.TestCheckResourceAttr(resourceName, "global_key", "SETTINGS"),
				resource.TestCheckResourceAttr(resourceName, "http_port", "0"),
				resource.TestCheckResourceAttr(resourceName, "https_port", "8443"),
				resource.TestCheckResourceAttrSet(resourceName, "key"),
				resource.TestCheckResourceAttr(resourceName, "pool_route", "HEADER"),
				resource.TestCheckResourceAttr(resourceName, "pool_routing_header", "Host1"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "type", "DEFAULT"),

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

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + databaseToolsDatabaseApiGatewayConfigIdVariableStr + DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_runtime_database_tools_database_api_gateway_config_global", "test_database_tools_database_api_gateway_config_global", acctest.Optional, acctest.Update, DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalUpdateRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "certificate_bundle.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "certificate_bundle.0.certificate_private_key.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "certificate_bundle.0.certificate_public.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "certificate_bundle.0.type", "SELF_SIGNED"),
				resource.TestCheckResourceAttr(resourceName, "database_api_status", "ENABLED"),
				resource.TestCheckResourceAttrSet(resourceName, "database_tools_database_api_gateway_config_id"),
				resource.TestCheckResourceAttr(resourceName, "document_root", "/var/www/webroots"),
				resource.TestCheckResourceAttr(resourceName, "global_key", "SETTINGS"),
				resource.TestCheckResourceAttr(resourceName, "http_port", "0"),
				resource.TestCheckResourceAttr(resourceName, "https_port", "8443"),
				resource.TestCheckResourceAttrSet(resourceName, "key"),
				resource.TestCheckResourceAttr(resourceName, "pool_route", "HEADER"),
				resource.TestCheckResourceAttr(resourceName, "pool_routing_header", "Host1"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "type", "DEFAULT"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_runtime_database_tools_database_api_gateway_config_global", "test_database_tools_database_api_gateway_config_global", acctest.Required, acctest.Create, DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalSingularDataSourceRepresentation) +
				compartmentIdVariableStr + databaseToolsDatabaseApiGatewayConfigIdVariableStr + DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "database_tools_database_api_gateway_config_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "global_key", "SETTINGS"),

				resource.TestCheckResourceAttr(singularDatasourceName, "certificate_bundle.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "certificate_bundle.0.certificate_private_key.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "certificate_bundle.0.certificate_public.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "certificate_bundle.0.type", "SELF_SIGNED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "database_api_status", "ENABLED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "document_root", "/var/www/webroots"),
				resource.TestCheckResourceAttr(singularDatasourceName, "http_port", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "https_port", "8443"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "key"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "metadata_source"),
				resource.TestCheckResourceAttr(singularDatasourceName, "pool_route", "HEADER"),
				resource.TestCheckResourceAttr(singularDatasourceName, "pool_routing_header", "Host1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "type", "DEFAULT"),
			),
		},
		// verify resource import
		{
			Config:            config + compartmentIdVariableStr + databaseToolsDatabaseApiGatewayConfigIdVariableStr + DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateCheck: checkImportedDatabaseToolsRuntimeCompositeID(
				resourceName,
				parseDatabaseToolsRuntimeDatabaseApiGatewayConfigGlobalCompositeIDToAttributes,
			),
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}
