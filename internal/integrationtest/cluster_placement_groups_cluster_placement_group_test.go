// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	oci_cluster_placement_groups "github.com/oracle/oci-go-sdk/v65/clusterplacementgroups"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	ClusterPlacementGroupsClusterPlacementGroupRequiredOnlyResource = ClusterPlacementGroupsClusterPlacementGroupResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_cluster_placement_groups_cluster_placement_group", "test_cluster_placement_group", acctest.Required, acctest.Create, ClusterPlacementGroupsClusterPlacementGroupRepresentation)

	ClusterPlacementGroupsClusterPlacementGroupResourceConfig = ClusterPlacementGroupsClusterPlacementGroupResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_cluster_placement_groups_cluster_placement_group", "test_cluster_placement_group", acctest.Optional, acctest.Update, ClusterPlacementGroupsClusterPlacementGroupRepresentationWithCaps)

	ClusterPlacementGroupsClusterPlacementGroupSingularDataSourceRepresentation = map[string]interface{}{
		"cluster_placement_group_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_cluster_placement_groups_cluster_placement_group.test_cluster_placement_group.id}`},
	}

	nameWithRequired                  = utils.RandomString(8, utils.CharsetLowerCaseWithoutDigits)
	nameWithCapabilities              = utils.RandomString(8, utils.CharsetLowerCaseWithoutDigits)
	nameWithCapabilitiesVariableStr   = fmt.Sprintf("variable \"name\" { default = \"%s\" }\n", nameWithCapabilities)
	nameWithPlacement                 = utils.RandomString(8, utils.CharsetLowerCaseWithoutDigits)
	nameUpdateCompartments            = utils.RandomString(8, utils.CharsetLowerCaseWithoutDigits)
	nameUpdateCompartmentsVariableStr = fmt.Sprintf("variable \"name\" { default = \"%s\" }\n", nameUpdateCompartments)
	nameUpdateUpdates                 = utils.RandomString(8, utils.CharsetLowerCaseWithoutDigits)
	nameUpdateUpdatesVariableStr      = fmt.Sprintf("variable \"name\" { default = \"%s\" }\n", nameUpdateUpdates)

	nameDs1    = utils.RandomString(8, utils.CharsetLowerCaseWithoutDigits)
	nameDs1Str = fmt.Sprintf("variable \"name\" { default = \"%s\" }\n", nameDs1)

	ClusterPlacementGroupsClusterPlacementGroupDataSourceRepresentation = map[string]interface{}{
		"ad":                        acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":            acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"id":                        acctest.Representation{RepType: acctest.Optional, Create: `${oci_cluster_placement_groups_cluster_placement_group.test_cluster_placement_group.id}`},
		"name":                      acctest.Representation{RepType: acctest.Optional, Create: nameWithCapabilities},
		"state":                     acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: ClusterPlacementGroupsClusterPlacementGroupDataSourceFilterRepresentation}}
	ClusterPlacementGroupsClusterPlacementGroupDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_cluster_placement_groups_cluster_placement_group.test_cluster_placement_group.id}`}},
	}

	ClusterPlacementGroupsClusterPlacementGroupRepresentation = map[string]interface{}{
		"availability_domain":          acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"cluster_placement_group_type": acctest.Representation{RepType: acctest.Required, Create: `STANDARD`},
		"compartment_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"description":                  acctest.Representation{RepType: acctest.Required, Create: `description`, Update: `description2`},
		"name":                         acctest.Representation{RepType: acctest.Required, Create: nameWithRequired},
		"capabilities":                 acctest.RepresentationGroup{RepType: acctest.Optional, Group: ClusterPlacementGroupsClusterPlacementGroupCapabilitiesRepresentation},
		"defined_tags":                 acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":                acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"opc_dry_run":                  acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"placement_instruction":        acctest.RepresentationGroup{RepType: acctest.Optional, Group: ClusterPlacementGroupsClusterPlacementGroupPlacementInstructionRepresentation},
		"state":                        acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`, Update: `ACTIVE`},
		"lifecycle":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreCPGChangesRepresentation},
	}

	ClusterPlacementGroupsClusterPlacementGroupRepresentationWithCaps = map[string]interface{}{
		"availability_domain":          acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"cluster_placement_group_type": acctest.Representation{RepType: acctest.Required, Create: `STANDARD`},
		"compartment_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"description":                  acctest.Representation{RepType: acctest.Required, Create: `description`, Update: `description2`},
		"name":                         acctest.Representation{RepType: acctest.Required, Create: `${var.name}`},
		"capabilities":                 acctest.RepresentationGroup{RepType: acctest.Optional, Group: ClusterPlacementGroupsClusterPlacementGroupCapabilitiesRepresentation},
		"defined_tags":                 acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":                acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"opc_dry_run":                  acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"state":                        acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`, Update: `ACTIVE`},
		"lifecycle":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreCPGChangesRepresentation},
	}

	ClusterPlacementGroupsClusterPlacementGroupRepresentationWithPlacement = map[string]interface{}{
		"availability_domain":          acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"cluster_placement_group_type": acctest.Representation{RepType: acctest.Required, Create: `STANDARD`},
		"compartment_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"description":                  acctest.Representation{RepType: acctest.Required, Create: `description`, Update: `description2`},
		"name":                         acctest.Representation{RepType: acctest.Required, Create: nameWithPlacement},
		"defined_tags":                 acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":                acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"opc_dry_run":                  acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"placement_instruction":        acctest.RepresentationGroup{RepType: acctest.Optional, Group: ClusterPlacementGroupsClusterPlacementGroupPlacementInstructionRepresentation},
		"state":                        acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`, Update: `ACTIVE`},
		"lifecycle":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreCPGChangesRepresentation},
	}
	ClusterPlacementGroupsClusterPlacementGroupCapabilitiesRepresentation = map[string]interface{}{
		"items": acctest.RepresentationGroup{RepType: acctest.Required, Group: ClusterPlacementGroupsClusterPlacementGroupCapabilitiesItemsRepresentation},
	}
	ClusterPlacementGroupsClusterPlacementGroupPlacementInstructionRepresentation = map[string]interface{}{
		"type":  acctest.Representation{RepType: acctest.Required, Create: `TOKEN`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `canary-placementtoken`},
	}
	ClusterPlacementGroupsClusterPlacementGroupCapabilitiesItemsRepresentation = map[string]interface{}{
		"name":    acctest.Representation{RepType: acctest.Required, Create: `volume`},
		"service": acctest.Representation{RepType: acctest.Required, Create: `block-storage`},
	}

	ClusterPlacementGroupsClusterPlacementGroupResourceDependencies = AvailabilityDomainConfig +
		DefinedTagsDependencies

	ignoreCPGChangesRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`, `freeform_tags`}, Update: []string{`defined_tags`, `freeform_tags`}},
	}
)

// issue-routing-tag: cluster_placement_groups/default
func TestClusterPlacementGroupsClusterPlacementGroupResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestClusterPlacementGroupsClusterPlacementGroupResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_cluster_placement_groups_cluster_placement_group.test_cluster_placement_group"
	datasourceName := "data.oci_cluster_placement_groups_cluster_placement_groups.test_cluster_placement_groups"
	singularDatasourceName := "data.oci_cluster_placement_groups_cluster_placement_group.test_cluster_placement_group"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ClusterPlacementGroupsClusterPlacementGroupResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_cluster_placement_groups_cluster_placement_group", "test_cluster_placement_group", acctest.Optional, acctest.Create, ClusterPlacementGroupsClusterPlacementGroupRepresentation), "clusterplacementgroups", "clusterPlacementGroup", t)

	acctest.ResourceTest(t, testAccCheckClusterPlacementGroupsClusterPlacementGroupDestroy, []resource.TestStep{
		// verify Create 0
		{
			Config: config + compartmentIdVariableStr + ClusterPlacementGroupsClusterPlacementGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cluster_placement_groups_cluster_placement_group", "test_cluster_placement_group", acctest.Required, acctest.Create, ClusterPlacementGroupsClusterPlacementGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "cluster_placement_group_type", "STANDARD"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "name", nameWithRequired),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create 1
		{
			Config: config + compartmentIdVariableStr + ClusterPlacementGroupsClusterPlacementGroupResourceDependencies,
		},

		// verify Create with optionals 2
		{
			Config: config + compartmentIdVariableStr + ClusterPlacementGroupsClusterPlacementGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cluster_placement_groups_cluster_placement_group", "test_cluster_placement_group", acctest.Optional, acctest.Create, ClusterPlacementGroupsClusterPlacementGroupRepresentationWithPlacement),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "cluster_placement_group_type", "STANDARD"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", nameWithPlacement),
				resource.TestCheckResourceAttr(resourceName, "opc_dry_run", "false"),
				resource.TestCheckResourceAttr(resourceName, "placement_instruction.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "placement_instruction.0.type", "TOKEN"),
				resource.TestCheckResourceAttr(resourceName, "placement_instruction.0.value", "canary-placementtoken"),
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

		// verify Create with optionals 3
		{
			Config: config + compartmentIdVariableStr + nameWithCapabilitiesVariableStr + ClusterPlacementGroupsClusterPlacementGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cluster_placement_groups_cluster_placement_group", "test_cluster_placement_group", acctest.Optional, acctest.Create, ClusterPlacementGroupsClusterPlacementGroupRepresentationWithCaps),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "capabilities.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "capabilities.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "capabilities.0.items.0.name", "volume"),
				resource.TestCheckResourceAttr(resourceName, "capabilities.0.items.0.service", "block-storage"),
				resource.TestCheckResourceAttr(resourceName, "cluster_placement_group_type", "STANDARD"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", nameWithCapabilities),
				resource.TestCheckResourceAttr(resourceName, "opc_dry_run", "false"),
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

		// verify Update to the compartment (the compartment will be switched back in the next step) 4
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + nameWithCapabilitiesVariableStr + ClusterPlacementGroupsClusterPlacementGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cluster_placement_groups_cluster_placement_group", "test_cluster_placement_group", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(ClusterPlacementGroupsClusterPlacementGroupRepresentationWithCaps, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "capabilities.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "capabilities.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "capabilities.0.items.0.name", "volume"),
				resource.TestCheckResourceAttr(resourceName, "capabilities.0.items.0.service", "block-storage"),
				resource.TestCheckResourceAttr(resourceName, "cluster_placement_group_type", "STANDARD"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", nameWithCapabilities),
				resource.TestCheckResourceAttr(resourceName, "opc_dry_run", "false"),
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

		// verify updates to updatable parameters 5
		{
			Config: config + compartmentIdVariableStr + nameWithCapabilitiesVariableStr + ClusterPlacementGroupsClusterPlacementGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cluster_placement_groups_cluster_placement_group", "test_cluster_placement_group", acctest.Optional, acctest.Update, ClusterPlacementGroupsClusterPlacementGroupRepresentationWithCaps),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "capabilities.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "capabilities.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "capabilities.0.items.0.name", "volume"),
				resource.TestCheckResourceAttr(resourceName, "capabilities.0.items.0.service", "block-storage"),
				resource.TestCheckResourceAttr(resourceName, "cluster_placement_group_type", "STANDARD"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", nameWithCapabilities),
				resource.TestCheckResourceAttr(resourceName, "opc_dry_run", "false"),
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
		// verify datasource 6
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_cluster_placement_groups_cluster_placement_groups", "test_cluster_placement_groups", acctest.Optional, acctest.Update, ClusterPlacementGroupsClusterPlacementGroupDataSourceRepresentation) +
				compartmentIdVariableStr + nameWithCapabilitiesVariableStr + ClusterPlacementGroupsClusterPlacementGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cluster_placement_groups_cluster_placement_group", "test_cluster_placement_group", acctest.Optional, acctest.Update, ClusterPlacementGroupsClusterPlacementGroupRepresentationWithCaps),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "ad"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "false"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "name", nameWithCapabilities),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "cluster_placement_group_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "cluster_placement_group_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource 7
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_cluster_placement_groups_cluster_placement_group", "test_cluster_placement_group", acctest.Required, acctest.Create, ClusterPlacementGroupsClusterPlacementGroupSingularDataSourceRepresentation) +
				compartmentIdVariableStr + nameWithCapabilitiesVariableStr + ClusterPlacementGroupsClusterPlacementGroupResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cluster_placement_group_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "capabilities.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "capabilities.0.items.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "capabilities.0.items.0.name", "volume"),
				resource.TestCheckResourceAttr(singularDatasourceName, "capabilities.0.items.0.service", "block-storage"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cluster_placement_group_type", "STANDARD"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", nameWithCapabilities),
				resource.TestCheckResourceAttr(singularDatasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import 8
		{
			Config:            config + ClusterPlacementGroupsClusterPlacementGroupRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"opc_dry_run",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckClusterPlacementGroupsClusterPlacementGroupDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ClusterPlacementGroupsCPClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_cluster_placement_groups_cluster_placement_group" {
			noResourceFound = false
			request := oci_cluster_placement_groups.GetClusterPlacementGroupRequest{}

			tmp := rs.Primary.ID
			request.ClusterPlacementGroupId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "cluster_placement_groups")

			response, err := client.GetClusterPlacementGroup(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_cluster_placement_groups.ClusterPlacementGroupLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("ClusterPlacementGroupsClusterPlacementGroup") {
		resource.AddTestSweepers("ClusterPlacementGroupsClusterPlacementGroup", &resource.Sweeper{
			Name:         "ClusterPlacementGroupsClusterPlacementGroup",
			Dependencies: acctest.DependencyGraph["clusterPlacementGroup"],
			F:            sweepClusterPlacementGroupsClusterPlacementGroupResource,
		})
	}
}

func sweepClusterPlacementGroupsClusterPlacementGroupResource(compartment string) error {
	clusterPlacementGroupsCPClient := acctest.GetTestClients(&schema.ResourceData{}).ClusterPlacementGroupsCPClient()
	clusterPlacementGroupIds, err := getClusterPlacementGroupsClusterPlacementGroupIds(compartment)
	if err != nil {
		return err
	}
	for _, clusterPlacementGroupId := range clusterPlacementGroupIds {
		if ok := acctest.SweeperDefaultResourceId[clusterPlacementGroupId]; !ok {
			deleteClusterPlacementGroupRequest := oci_cluster_placement_groups.DeleteClusterPlacementGroupRequest{}

			deleteClusterPlacementGroupRequest.ClusterPlacementGroupId = &clusterPlacementGroupId

			deleteClusterPlacementGroupRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "cluster_placement_groups")
			_, error := clusterPlacementGroupsCPClient.DeleteClusterPlacementGroup(context.Background(), deleteClusterPlacementGroupRequest)
			if error != nil {
				fmt.Printf("Error deleting ClusterPlacementGroup %s %s, It is possible that the resource is already deleted. Please verify manually \n", clusterPlacementGroupId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &clusterPlacementGroupId, ClusterPlacementGroupsClusterPlacementGroupSweepWaitCondition, time.Duration(3*time.Minute),
				ClusterPlacementGroupsClusterPlacementGroupSweepResponseFetchOperation, "cluster_placement_groups", true)
		}
	}
	return nil
}

func getClusterPlacementGroupsClusterPlacementGroupIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ClusterPlacementGroupId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	clusterPlacementGroupsCPClient := acctest.GetTestClients(&schema.ResourceData{}).ClusterPlacementGroupsCPClient()

	listClusterPlacementGroupsRequest := oci_cluster_placement_groups.ListClusterPlacementGroupsRequest{}
	listClusterPlacementGroupsRequest.CompartmentId = &compartmentId
	listClusterPlacementGroupsRequest.LifecycleState = oci_cluster_placement_groups.ClusterPlacementGroupLifecycleStateActive
	listClusterPlacementGroupsResponse, err := clusterPlacementGroupsCPClient.ListClusterPlacementGroups(context.Background(), listClusterPlacementGroupsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ClusterPlacementGroup list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, clusterPlacementGroup := range listClusterPlacementGroupsResponse.Items {
		id := *clusterPlacementGroup.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ClusterPlacementGroupId", id)
	}
	return resourceIds, nil
}

func ClusterPlacementGroupsClusterPlacementGroupSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if clusterPlacementGroupResponse, ok := response.Response.(oci_cluster_placement_groups.GetClusterPlacementGroupResponse); ok {
		return clusterPlacementGroupResponse.LifecycleState != oci_cluster_placement_groups.ClusterPlacementGroupLifecycleStateDeleted
	}
	return false
}

func ClusterPlacementGroupsClusterPlacementGroupSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ClusterPlacementGroupsCPClient().GetClusterPlacementGroup(context.Background(), oci_cluster_placement_groups.GetClusterPlacementGroupRequest{
		ClusterPlacementGroupId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
