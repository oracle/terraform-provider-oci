// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
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
	DataSafeDatabaseSecurityConfigManagementResource = DataSafeDatabaseSecurityConfigManagementResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_database_security_config_management", "test_database_security_config_management", acctest.Required, acctest.Create,
			DataSafeDatabaseSecurityConfigManagementRepresentation)

	DataSafeDatabaseSecurityConfigManagementRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"target_id":           acctest.Representation{RepType: acctest.Required, Create: `${var.target_id}`},
		"defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":         acctest.Representation{RepType: acctest.Optional, Create: `Updated database security configuration description`, Update: `description2`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `DatabaseSecurityConfig_updated`, Update: `displayName2`},
		"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"sql_firewall_config": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataSafeDatabaseSecurityConfigManagementSqlFirewallConfigRepresentation},
		"refresh_trigger":     acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"lifecycle":           acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDatabaseSecurityConfigManagementTagsChangesRep},
	}

	DataSafeDatabaseSecurityConfigManagementSqlFirewallConfigRepresentation = map[string]interface{}{
		"exclude_job":              acctest.Representation{RepType: acctest.Optional, Create: `EXCLUDED`, Update: `INCLUDED`},
		"status":                   acctest.Representation{RepType: acctest.Optional, Create: `DISABLED`, Update: `ENABLED`},
		"violation_log_auto_purge": acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`, Update: `DISABLED`},
	}

	ignoreDatabaseSecurityConfigManagementTagsChangesRep = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`, `freeform_tags`}},
	}

	DataSafeDatabaseSecurityConfigManagementResourceDependencies = DefinedTagsDependencies
)

func TestDataSafeDatabaseSecurityConfigManagementResource_basic(t *testing.T) {
	t.Skip("Skipping this test as the target ocid is hardcoded and may not exist when the test runs")
	httpreplay.SetScenario("TestDataSafeDatabaseSecurityConfigManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	targetId := utils.GetEnvSettingWithBlankDefault("data_safe_target_ocid")
	targetIdVariableStr := fmt.Sprintf("variable \"target_id\" { default = \"%s\" }\n", targetId)

	resourceName := "oci_data_safe_database_security_config_management.test_database_security_config_management"

	var resId, resId2 string

	acctest.SaveConfigContent(config+compartmentIdVariableStr+compartmentIdUVariableStr+targetIdVariableStr+DataSafeDatabaseSecurityConfigManagementResource+
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_database_security_config_management", "test_database_security_config_management", acctest.Optional, acctest.Update,
			DataSafeDatabaseSecurityConfigManagementRepresentation), "datasafe", "databaseSecurityConfigManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + targetIdVariableStr + DataSafeDatabaseSecurityConfigManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_database_security_config_management", "test_database_security_config_management", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(DataSafeDatabaseSecurityConfigManagementRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify change compartment
		{
			Config: config + compartmentIdUVariableStr + targetIdVariableStr + DataSafeDatabaseSecurityConfigManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_database_security_config_management", "test_database_security_config_management", acctest.Required, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(DataSafeDatabaseSecurityConfigManagementRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// switched back compartment
		{
			Config: config + compartmentIdVariableStr + targetIdVariableStr + DataSafeDatabaseSecurityConfigManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_database_security_config_management", "test_database_security_config_management", acctest.Required, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(DataSafeDatabaseSecurityConfigManagementRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
	})
}
