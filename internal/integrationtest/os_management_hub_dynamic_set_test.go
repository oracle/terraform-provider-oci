// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_os_management_hub "github.com/oracle/oci-go-sdk/v65/osmanagementhub"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	OsManagementHubDynamicSetRequiredOnlyResource = OsManagementHubDynamicSetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_dynamic_set", "test_dynamic_set", acctest.Required, acctest.Create, OsManagementHubDynamicSetRepresentation)

	OsManagementHubDynamicSetResourceConfig = OsManagementHubDynamicSetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_dynamic_set", "test_dynamic_set", acctest.Optional, acctest.Update, OsManagementHubDynamicSetRepresentation)

	OsManagementHubDynamicSetSingularDataSourceRepresentation = map[string]interface{}{
		"dynamic_set_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_os_management_hub_dynamic_set.test_dynamic_set.id}`},
	}

	OsManagementHubDynamicSetDataSourceRepresentation = map[string]interface{}{
		"compartment_id":        acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":          acctest.Representation{RepType: acctest.Optional, Create: `My-Display-Name`, Update: `displayName2`},
		"display_name_contains": acctest.Representation{RepType: acctest.Optional, Create: `displayNameContains`},
		"dynamic_set_id":        acctest.Representation{RepType: acctest.Optional, Create: `${oci_os_management_hub_dynamic_set.test_dynamic_set.id}`},
		"filter":                acctest.RepresentationGroup{RepType: acctest.Required, Group: OsManagementHubDynamicSetDataSourceFilterRepresentation}}
	OsManagementHubDynamicSetDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_os_management_hub_dynamic_set.test_dynamic_set.id}`}},
	}

	OsManagementHubDynamicSetRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":        acctest.Representation{RepType: acctest.Required, Create: `My-Display-Name`, Update: `displayName2`},
		"match_type":          acctest.Representation{RepType: acctest.Required, Create: `ALL`, Update: `ANY`},
		"matching_rule":       acctest.RepresentationGroup{RepType: acctest.Required, Group: OsManagementHubDynamicSetMatchingRuleRepresentation},
		"target_compartments": acctest.RepresentationGroup{RepType: acctest.Required, Group: OsManagementHubDynamicSetTargetCompartmentsRepresentation},
		"defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: OsManagementHubDefinedTagsRepresentation},
		"description":         acctest.Representation{RepType: acctest.Optional, Create: `User-specified information about the dynamic set.`, Update: `description2`},
		"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	OsManagementHubDynamicSetOL8Representation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":        acctest.Representation{RepType: acctest.Required, Create: `My-Display-Name`, Update: `displayName2`},
		"match_type":          acctest.Representation{RepType: acctest.Required, Create: `ALL`},
		"matching_rule":       acctest.RepresentationGroup{RepType: acctest.Required, Group: OsManagementHubDynamicSetMatchingRuleOL8Representation},
		"target_compartments": acctest.RepresentationGroup{RepType: acctest.Required, Group: OsManagementHubDynamicSetTargetCompartmentsRepresentation},
	}

	OsManagementHubDynamicSetMatchingRuleOL8Representation = map[string]interface{}{
		"architectures": acctest.Representation{RepType: acctest.Required, Create: []string{`X86_64`}},
		"locations":     acctest.Representation{RepType: acctest.Required, Create: []string{`OCI_COMPUTE`}},
		"os_families":   acctest.Representation{RepType: acctest.Required, Create: []string{`ORACLE_LINUX_8`}},
	}

	OsManagementHubDynamicSetMatchingRuleRepresentation = map[string]interface{}{
		"architectures":              acctest.Representation{RepType: acctest.Optional, Create: []string{`X86_64`}, Update: []string{`AARCH64`}},
		"display_names":              acctest.Representation{RepType: acctest.Optional, Create: []string{`displayNames`}, Update: []string{`displayNames2`}},
		"is_reboot_required":         acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"locations":                  acctest.Representation{RepType: acctest.Optional, Create: []string{`OCI_COMPUTE`}},
		"managed_instance_group_ids": acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_os_management_hub_managed_instance_group.test_grp.id}`}},
		"managed_instance_ids":       acctest.Representation{RepType: acctest.Optional, Create: []string{utils.GetEnvSettingWithBlankDefault("osmh_managed_instance_ocid")}, Update: []string{utils.GetEnvSettingWithBlankDefault("osmh_managed_instance_ubuntu_ocid")}},
		"os_families":                acctest.Representation{RepType: acctest.Optional, Create: []string{`ORACLE_LINUX_8`}, Update: []string{`UBUNTU_24_04`}},
		"os_names":                   acctest.Representation{RepType: acctest.Optional, Create: []string{`ORACLE_LINUX`}, Update: []string{`UBUNTU`}},
		"statuses":                   acctest.Representation{RepType: acctest.Optional, Create: []string{`NORMAL`}, Update: []string{`WARNING`}},
		"tags":                       acctest.RepresentationGroup{RepType: acctest.Optional, Group: OsManagementHubDynamicSetMatchingRuleTagsRepresentation},
	}
	OsManagementHubDynamicSetTargetCompartmentsRepresentation = map[string]interface{}{
		"compartment_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"does_include_children": acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
	}
	OsManagementHubDynamicSetMatchingRuleTagsRepresentation = map[string]interface{}{
		"type":      acctest.Representation{RepType: acctest.Required, Create: `DEFINED`},
		"key":       acctest.Representation{RepType: acctest.Optional, Create: `key`, Update: `key2`},
		"namespace": acctest.Representation{RepType: acctest.Optional, Create: `namespace`, Update: `namespace2`},
		"value":     acctest.Representation{RepType: acctest.Optional, Create: `value`, Update: `value2`},
	}

	OsManagementHubDynamicSetResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance_group", "test_grp", acctest.Required, acctest.Create, OsManagementHubManagedInstanceGroupRepresentation) + OsManagementHubVendorSoftwareSourceOl8BaseosLatestX8664Config
)

// issue-routing-tag: os_management_hub/default
func TestOsManagementHubDynamicSetResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsManagementHubDynamicSetResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_os_management_hub_dynamic_set.test_dynamic_set"
	datasourceName := "data.oci_os_management_hub_dynamic_sets.test_dynamic_sets"
	singularDatasourceName := "data.oci_os_management_hub_dynamic_set.test_dynamic_set"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+OsManagementHubDynamicSetResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_dynamic_set", "test_dynamic_set", acctest.Optional, acctest.Create, OsManagementHubDynamicSetRepresentation), "osmanagementhub", "dynamicSet", t)

	acctest.ResourceTest(t, testAccCheckOsManagementHubDynamicSetDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + OsManagementHubDynamicSetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_dynamic_set", "test_dynamic_set", acctest.Required, acctest.Create, OsManagementHubDynamicSetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "My-Display-Name"),
				resource.TestCheckResourceAttr(resourceName, "match_type", "ALL"),
				resource.TestCheckResourceAttr(resourceName, "matching_rule.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target_compartments.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "target_compartments.0.compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "target_compartments.0.does_include_children", "false"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + OsManagementHubDynamicSetResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + OsManagementHubDynamicSetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_dynamic_set", "test_dynamic_set", acctest.Optional, acctest.Create, OsManagementHubDynamicSetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "User-specified information about the dynamic set."),
				resource.TestCheckResourceAttr(resourceName, "display_name", "My-Display-Name"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "match_type", "ALL"),
				resource.TestCheckResourceAttr(resourceName, "matching_rule.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "matching_rule.0.architectures.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "matching_rule.0.display_names.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "matching_rule.0.is_reboot_required", "false"),
				resource.TestCheckResourceAttr(resourceName, "matching_rule.0.locations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "matching_rule.0.managed_instance_group_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "matching_rule.0.managed_instance_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "matching_rule.0.os_families.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "matching_rule.0.os_names.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "matching_rule.0.statuses.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "matching_rule.0.tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "matching_rule.0.tags.0.key", "key"),
				resource.TestCheckResourceAttr(resourceName, "matching_rule.0.tags.0.namespace", "namespace"),
				resource.TestCheckResourceAttr(resourceName, "matching_rule.0.tags.0.type", "DEFINED"),
				resource.TestCheckResourceAttr(resourceName, "matching_rule.0.tags.0.value", "value"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "target_compartments.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "target_compartments.0.compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "target_compartments.0.does_include_children", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + OsManagementHubDynamicSetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_dynamic_set", "test_dynamic_set", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(OsManagementHubDynamicSetRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "User-specified information about the dynamic set."),
				resource.TestCheckResourceAttr(resourceName, "display_name", "My-Display-Name"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "match_type", "ALL"),
				resource.TestCheckResourceAttr(resourceName, "matching_rule.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "matching_rule.0.architectures.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "matching_rule.0.display_names.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "matching_rule.0.is_reboot_required", "false"),
				resource.TestCheckResourceAttr(resourceName, "matching_rule.0.locations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "matching_rule.0.managed_instance_group_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "matching_rule.0.managed_instance_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "matching_rule.0.os_families.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "matching_rule.0.os_names.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "matching_rule.0.statuses.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "matching_rule.0.tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "matching_rule.0.tags.0.key", "key"),
				resource.TestCheckResourceAttr(resourceName, "matching_rule.0.tags.0.namespace", "namespace"),
				resource.TestCheckResourceAttr(resourceName, "matching_rule.0.tags.0.type", "DEFINED"),
				resource.TestCheckResourceAttr(resourceName, "matching_rule.0.tags.0.value", "value"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "target_compartments.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "target_compartments.0.compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "target_compartments.0.does_include_children", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + OsManagementHubDynamicSetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_dynamic_set", "test_dynamic_set", acctest.Optional, acctest.Update, OsManagementHubDynamicSetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "match_type", "ANY"),
				resource.TestCheckResourceAttr(resourceName, "matching_rule.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "matching_rule.0.architectures.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "matching_rule.0.display_names.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "matching_rule.0.is_reboot_required", "true"),
				resource.TestCheckResourceAttr(resourceName, "matching_rule.0.locations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "matching_rule.0.managed_instance_group_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "matching_rule.0.managed_instance_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "matching_rule.0.os_families.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "matching_rule.0.os_names.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "matching_rule.0.statuses.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "matching_rule.0.tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "matching_rule.0.tags.0.key", "key2"),
				resource.TestCheckResourceAttr(resourceName, "matching_rule.0.tags.0.namespace", "namespace2"),
				resource.TestCheckResourceAttr(resourceName, "matching_rule.0.tags.0.type", "DEFINED"),
				resource.TestCheckResourceAttr(resourceName, "matching_rule.0.tags.0.value", "value2"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "target_compartments.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "target_compartments.0.compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "target_compartments.0.does_include_children", "true"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_dynamic_sets", "test_dynamic_sets", acctest.Optional, acctest.Update, OsManagementHubDynamicSetDataSourceRepresentation) +
				compartmentIdVariableStr + OsManagementHubDynamicSetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_dynamic_set", "test_dynamic_set", acctest.Optional, acctest.Update, OsManagementHubDynamicSetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "display_name_contains", "displayNameContains"),
				resource.TestCheckResourceAttrSet(datasourceName, "dynamic_set_id"),

				resource.TestCheckResourceAttr(datasourceName, "dynamic_set_collection.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "dynamic_set_collection.0.items.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_dynamic_set", "test_dynamic_set", acctest.Required, acctest.Create, OsManagementHubDynamicSetSingularDataSourceRepresentation) +
				compartmentIdVariableStr + OsManagementHubDynamicSetResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "dynamic_set_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "match_type", "ANY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "matching_rule.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "matching_rule.0.architectures.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "matching_rule.0.display_names.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "matching_rule.0.is_reboot_required", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "matching_rule.0.locations.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "matching_rule.0.managed_instance_group_ids.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "matching_rule.0.managed_instance_ids.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "matching_rule.0.os_families.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "matching_rule.0.os_names.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "matching_rule.0.statuses.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "matching_rule.0.tags.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "matching_rule.0.tags.0.key", "key2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "matching_rule.0.tags.0.namespace", "namespace2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "matching_rule.0.tags.0.type", "DEFINED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "matching_rule.0.tags.0.value", "value2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "scheduled_job_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target_compartments.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_compartments.0.compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target_compartments.0.does_include_children", "true"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + OsManagementHubDynamicSetRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckOsManagementHubDynamicSetDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DynamicSetClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_os_management_hub_dynamic_set" {
			noResourceFound = false
			request := oci_os_management_hub.GetDynamicSetRequest{}

			tmp := rs.Primary.ID
			request.DynamicSetId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "os_management_hub")

			response, err := client.GetDynamicSet(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_os_management_hub.DynamicSetLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("OsManagementHubDynamicSet") {
		resource.AddTestSweepers("OsManagementHubDynamicSet", &resource.Sweeper{
			Name:         "OsManagementHubDynamicSet",
			Dependencies: acctest.DependencyGraph["dynamicSet"],
			F:            sweepOsManagementHubDynamicSetResource,
		})
	}
}

func sweepOsManagementHubDynamicSetResource(compartment string) error {
	dynamicSetClient := acctest.GetTestClients(&schema.ResourceData{}).DynamicSetClient()
	dynamicSetIds, err := getOsManagementHubDynamicSetIds(compartment)
	if err != nil {
		return err
	}
	for _, dynamicSetId := range dynamicSetIds {
		if ok := acctest.SweeperDefaultResourceId[dynamicSetId]; !ok {
			deleteDynamicSetRequest := oci_os_management_hub.DeleteDynamicSetRequest{}

			deleteDynamicSetRequest.DynamicSetId = &dynamicSetId

			deleteDynamicSetRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "os_management_hub")
			_, error := dynamicSetClient.DeleteDynamicSet(context.Background(), deleteDynamicSetRequest)
			if error != nil {
				fmt.Printf("Error deleting DynamicSet %s %s, It is possible that the resource is already deleted. Please verify manually \n", dynamicSetId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &dynamicSetId, OsManagementHubDynamicSetSweepWaitCondition, time.Duration(3*time.Minute),
				OsManagementHubDynamicSetSweepResponseFetchOperation, "os_management_hub", true)
		}
	}
	return nil
}

func getOsManagementHubDynamicSetIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DynamicSetId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dynamicSetClient := acctest.GetTestClients(&schema.ResourceData{}).DynamicSetClient()

	listDynamicSetsRequest := oci_os_management_hub.ListDynamicSetsRequest{}
	listDynamicSetsRequest.CompartmentId = &compartmentId
	listDynamicSetsResponse, err := dynamicSetClient.ListDynamicSets(context.Background(), listDynamicSetsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting DynamicSet list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, dynamicSet := range listDynamicSetsResponse.Items {
		id := *dynamicSet.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DynamicSetId", id)
	}
	return resourceIds, nil
}

func OsManagementHubDynamicSetSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if dynamicSetResponse, ok := response.Response.(oci_os_management_hub.GetDynamicSetResponse); ok {
		return dynamicSetResponse.LifecycleState != oci_os_management_hub.DynamicSetLifecycleStateDeleted
	}
	return false
}

func OsManagementHubDynamicSetSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DynamicSetClient().GetDynamicSet(context.Background(), oci_os_management_hub.GetDynamicSetRequest{
		DynamicSetId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
