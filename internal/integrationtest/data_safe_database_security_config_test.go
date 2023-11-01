// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DataSafeDatabaseSecurityConfigRequiredOnlyResource = DataSafeDatabaseSecurityConfigResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_database_security_config", "test_database_security_config", acctest.Required, acctest.Create, DataSafeDatabaseSecurityConfigUpdateRepresentation)

	DataSafeDatabaseSecurityConfigResourceConfig = DataSafeDatabaseSecurityConfigResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_database_security_config", "test_database_security_config", acctest.Optional, acctest.Update, DataSafeDatabaseSecurityConfigUpdateRepresentation)

	DataSafeDatabaseSecurityConfigSingularDataSourceRepresentation = map[string]interface{}{
		"database_security_config_id": acctest.Representation{RepType: acctest.Required, Create: `${var.database_security_config_id}`},
	}

	DataSafeDatabaseSecurityConfigDataSourceRepresentation = map[string]interface{}{
		"compartment_id":              acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"access_level":                acctest.Representation{RepType: acctest.Optional, Create: `RESTRICTED`},
		"compartment_id_in_subtree":   acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"database_security_config_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.database_security_config_id}`},
	}

	DataSafeDatabaseSecurityConfigUpdateRepresentation = map[string]interface{}{
		"database_security_config_id": acctest.Representation{RepType: acctest.Required, Create: `${var.database_security_config_id}`},
		"defined_tags":                acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":                 acctest.Representation{RepType: acctest.Optional, Create: `Updated database security configuration description`, Update: `description2`},
		"display_name":                acctest.Representation{RepType: acctest.Optional, Create: `DatabaseSecurityConfig_updated`, Update: `displayName2`},
		"freeform_tags":               acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"sql_firewall_config":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataSafeDatabaseSecurityConfigSqlFirewallConfigRepresentation},
		"refresh_trigger":             acctest.Representation{RepType: acctest.Optional, Create: `1`, Update: `1`},
		"lifecycle":                   acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDatabaseSecurityConfigTagsChangesRep},
	}
	DataSafeDatabaseSecurityConfigSqlFirewallConfigRepresentation = map[string]interface{}{
		"exclude_job":              acctest.Representation{RepType: acctest.Optional, Create: `EXCLUDED`, Update: `INCLUDED`},
		"status":                   acctest.Representation{RepType: acctest.Optional, Create: `DISABLED`, Update: `ENABLED`},
		"violation_log_auto_purge": acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`, Update: `DISABLED`},
	}
	ignoreDatabaseSecurityConfigTagsChangesRep = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`system_tags`, `defined_tags`, `freeform_tags`}},
	}

	DataSafeDatabaseSecurityConfigRefreshRepresentation = map[string]interface{}{
		"display_name":                acctest.Representation{RepType: acctest.Optional, Create: `DatabaseSecurityConfig_updated`, Update: `displayName2`},
		"database_security_config_id": acctest.Representation{RepType: acctest.Required, Create: `${var.database_security_config_id}`},
		"refresh_trigger":             acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
		"lifecycle":                   acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDatabaseSecurityConfigTagsChangesRep},
	}

	DataSafeDatabaseSecurityConfigResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: data_safe/default
func TestDataSafeDatabaseSecurityConfigResource_basic(t *testing.T) {
	t.Skip("Skipping this test as the database security config ocid is hardcoded and may not exist when the test runs")
	httpreplay.SetScenario("TestDataSafeDatabaseSecurityConfigResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	databaseSecurityConfigId := utils.GetEnvSettingWithBlankDefault("database_security_config_ocid")
	databaseSecurityConfigIdVariableStr := fmt.Sprintf("variable \"database_security_config_id\" { default = \"%s\" }\n", databaseSecurityConfigId)

	resourceName := "oci_data_safe_database_security_config.test_database_security_config"
	datasourceName := "data.oci_data_safe_database_security_configs.test_database_security_configs"
	singularDatasourceName := "data.oci_data_safe_database_security_config.test_database_security_config"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+compartmentIdUVariableStr+databaseSecurityConfigIdVariableStr+DataSafeDatabaseSecurityConfigResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_database_security_config", "test_database_security_config", acctest.Optional, acctest.Create,
			acctest.RepresentationCopyWithNewProperties(DataSafeDatabaseSecurityConfigUpdateRepresentation, map[string]interface{}{
				"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
			})), "datasafe", "databaseSecurityConfig", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify refresh and change compartment
		{
			Config: config + compartmentIdVariableStr + databaseSecurityConfigIdVariableStr + compartmentIdUVariableStr + DataSafeDatabaseSecurityConfigResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_database_security_config", "test_database_security_config", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(DataSafeDatabaseSecurityConfigRefreshRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "database_security_config_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + databaseSecurityConfigIdVariableStr + DataSafeDatabaseSecurityConfigResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_database_security_config", "test_database_security_config", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(DataSafeDatabaseSecurityConfigUpdateRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "database_security_config_id"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "sql_firewall_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "sql_firewall_config.0.exclude_job", "INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "sql_firewall_config.0.status", "ENABLED"),
				resource.TestCheckResourceAttr(resourceName, "sql_firewall_config.0.violation_log_auto_purge", "DISABLED"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "target_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_database_security_configs", "test_database_security_configs", acctest.Optional, acctest.Update, DataSafeDatabaseSecurityConfigDataSourceRepresentation) +
				compartmentIdVariableStr + databaseSecurityConfigIdVariableStr + DataSafeDatabaseSecurityConfigResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_database_security_config", "test_database_security_config", acctest.Optional, acctest.Update, DataSafeDatabaseSecurityConfigUpdateRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "database_security_config_id"),

				resource.TestCheckResourceAttr(datasourceName, "database_security_config_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "database_security_config_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_database_security_config", "test_database_security_config", acctest.Required, acctest.Create, DataSafeDatabaseSecurityConfigSingularDataSourceRepresentation) +
				compartmentIdVariableStr + databaseSecurityConfigIdVariableStr + DataSafeDatabaseSecurityConfigResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "database_security_config_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "sql_firewall_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "sql_firewall_config.0.exclude_job", "INCLUDED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "sql_firewall_config.0.status", "ENABLED"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "sql_firewall_config.0.time_status_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "sql_firewall_config.0.violation_log_auto_purge", "DISABLED"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_last_refreshed"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + DataSafeDatabaseSecurityConfigResourceConfig,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{`database_security_config_id`, `refresh_trigger`},
			ResourceName:            resourceName,
		},
	})
}
