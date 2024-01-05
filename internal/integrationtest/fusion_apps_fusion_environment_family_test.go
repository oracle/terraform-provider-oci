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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_fusion_apps "github.com/oracle/oci-go-sdk/v65/fusionapps"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	FusionAppsFusionEnvironmentFamilyRequiredOnlyResource = FusionAppsFusionEnvironmentFamilyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_fusion_apps_fusion_environment_family", "test_fusion_environment_family", acctest.Required, acctest.Create, FusionAppsFusionEnvironmentFamilyRepresentation)

	FusionAppsFusionEnvironmentFamilyResourceConfig = FusionAppsFusionEnvironmentFamilyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_fusion_apps_fusion_environment_family", "test_fusion_environment_family", acctest.Optional, acctest.Update, FusionAppsFusionEnvironmentFamilyRepresentation)

	FusionAppsFusionAppsFusionEnvironmentFamilySingularDataSourceRepresentation = map[string]interface{}{
		"fusion_environment_family_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_fusion_apps_fusion_environment_family.test_fusion_environment_family.id}`},
	}

	FusionAppsFusionAppsFusionEnvironmentFamilyDataSourceRepresentation = map[string]interface{}{
		"compartment_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":                 acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"fusion_environment_family_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_fusion_apps_fusion_environment_family.test_fusion_environment_family.id}`},
		"state":                        acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":                       acctest.RepresentationGroup{RepType: acctest.Required, Group: FusionAppsFusionEnvironmentFamilyDataSourceFilterRepresentation}}
	FusionAppsFusionEnvironmentFamilyDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_fusion_apps_fusion_environment_family.test_fusion_environment_family.id}`}},
	}

	FusionAppsFusionEnvironmentFamilyRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":              acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"subscription_ids":          acctest.Representation{RepType: acctest.Required, Create: []string{`ocid1.subscription.region1..BGBG0000m255pmogxppvxocajwtnscil7o36ohz3ep6flaut37dp22dxdpnq`}, Update: []string{`ocid1.subscription.region1..BGBG0000m255pmogxppvxocajwtnscil7o36ohz3ep6flaut37dp22dxdpnq`}},
		"defined_tags":              acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace-terraform.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace-terraform.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"family_maintenance_policy": acctest.RepresentationGroup{RepType: acctest.Optional, Group: FusionAppsFusionEnvironmentFamilyFamilyMaintenancePolicyRepresentation},
		"freeform_tags":             acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":                 acctest.RepresentationGroup{RepType: acctest.Optional, Group: definedTagsIgnoreRepresentation_fusionapps},
	}
	FusionAppsFusionEnvironmentFamilyFamilyMaintenancePolicyRepresentation = map[string]interface{}{
		"concurrent_maintenance":        acctest.Representation{RepType: acctest.Optional, Create: `PROD`, Update: `NON_PROD`},
		"is_monthly_patching_enabled":   acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"quarterly_upgrade_begin_times": acctest.Representation{RepType: acctest.Optional, Create: `RRULE:FREQ=YEARLY;BYMONTH=2,5,8,11`},
	}

	FusionAppsFusionEnvironmentFamilyResourceDependencies = DefinedTagsDependencies_fusionapps
)

// issue-routing-tag: fusion_apps/default
func TestFusionAppsFusionEnvironmentFamilyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFusionAppsFusionEnvironmentFamilyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_fusion_apps_fusion_environment_family.test_fusion_environment_family"
	datasourceName := "data.oci_fusion_apps_fusion_environment_families.test_fusion_environment_families"
	singularDatasourceName := "data.oci_fusion_apps_fusion_environment_family.test_fusion_environment_family"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+FusionAppsFusionEnvironmentFamilyResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_fusion_apps_fusion_environment_family", "test_fusion_environment_family", acctest.Optional, acctest.Create, FusionAppsFusionEnvironmentFamilyRepresentation), "fusionapps", "fusionEnvironmentFamily", t)

	acctest.ResourceTest(t, testAccCheckFusionAppsFusionEnvironmentFamilyDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + FusionAppsFusionEnvironmentFamilyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fusion_apps_fusion_environment_family", "test_fusion_environment_family", acctest.Required, acctest.Create, FusionAppsFusionEnvironmentFamilyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "subscription_ids.#", "1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + FusionAppsFusionEnvironmentFamilyResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + FusionAppsFusionEnvironmentFamilyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fusion_apps_fusion_environment_family", "test_fusion_environment_family", acctest.Optional, acctest.Create, FusionAppsFusionEnvironmentFamilyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "family_maintenance_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "family_maintenance_policy.0.concurrent_maintenance", "PROD"),
				resource.TestCheckResourceAttr(resourceName, "family_maintenance_policy.0.is_monthly_patching_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "family_maintenance_policy.0.quarterly_upgrade_begin_times", "RRULE:FREQ=YEARLY;BYMONTH=2,5,8,11"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "subscription_ids.#", "1"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + FusionAppsFusionEnvironmentFamilyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fusion_apps_fusion_environment_family", "test_fusion_environment_family", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(FusionAppsFusionEnvironmentFamilyRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "family_maintenance_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "family_maintenance_policy.0.concurrent_maintenance", "PROD"),
				resource.TestCheckResourceAttr(resourceName, "family_maintenance_policy.0.is_monthly_patching_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "family_maintenance_policy.0.quarterly_upgrade_begin_times", "RRULE:FREQ=YEARLY;BYMONTH=2,5,8,11"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "subscription_ids.#", "1"),

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
			Config: config + compartmentIdVariableStr + FusionAppsFusionEnvironmentFamilyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fusion_apps_fusion_environment_family", "test_fusion_environment_family", acctest.Optional, acctest.Update, FusionAppsFusionEnvironmentFamilyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "family_maintenance_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "family_maintenance_policy.0.concurrent_maintenance", "NON_PROD"),
				resource.TestCheckResourceAttr(resourceName, "family_maintenance_policy.0.is_monthly_patching_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "family_maintenance_policy.0.quarterly_upgrade_begin_times", "RRULE:FREQ=YEARLY;BYMONTH=2,5,8,11"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "subscription_ids.#", "1"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_fusion_apps_fusion_environment_families", "test_fusion_environment_families", acctest.Optional, acctest.Update, FusionAppsFusionAppsFusionEnvironmentFamilyDataSourceRepresentation) +
				compartmentIdVariableStr + FusionAppsFusionEnvironmentFamilyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fusion_apps_fusion_environment_family", "test_fusion_environment_family", acctest.Optional, acctest.Update, FusionAppsFusionEnvironmentFamilyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "fusion_environment_family_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "fusion_environment_family_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "fusion_environment_family_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_fusion_apps_fusion_environment_family", "test_fusion_environment_family", acctest.Required, acctest.Create, FusionAppsFusionAppsFusionEnvironmentFamilySingularDataSourceRepresentation) +
				compartmentIdVariableStr + FusionAppsFusionEnvironmentFamilyResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "fusion_environment_family_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "family_maintenance_policy.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "family_maintenance_policy.0.concurrent_maintenance", "NON_PROD"),
				resource.TestCheckResourceAttr(singularDatasourceName, "family_maintenance_policy.0.is_monthly_patching_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "family_maintenance_policy.0.quarterly_upgrade_begin_times", "RRULE:FREQ=YEARLY;BYMONTH=2,5,8,11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_subscription_update_needed"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "subscription_ids.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "system_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// verify resource import
		{
			Config:                  config + FusionAppsFusionEnvironmentFamilyRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckFusionAppsFusionEnvironmentFamilyDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).FusionApplicationsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_fusion_apps_fusion_environment_family" {
			noResourceFound = false
			request := oci_fusion_apps.GetFusionEnvironmentFamilyRequest{}

			tmp := rs.Primary.ID
			request.FusionEnvironmentFamilyId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "fusion_apps")

			response, err := client.GetFusionEnvironmentFamily(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_fusion_apps.FusionEnvironmentFamilyLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("FusionAppsFusionEnvironmentFamily") {
		resource.AddTestSweepers("FusionAppsFusionEnvironmentFamily", &resource.Sweeper{
			Name:         "FusionAppsFusionEnvironmentFamily",
			Dependencies: acctest.DependencyGraph["fusionEnvironmentFamily"],
			F:            sweepFusionAppsFusionEnvironmentFamilyResource,
		})
	}
}

func sweepFusionAppsFusionEnvironmentFamilyResource(compartment string) error {
	fusionApplicationsClient := acctest.GetTestClients(&schema.ResourceData{}).FusionApplicationsClient()
	fusionEnvironmentFamilyIds, err := getFusionAppsFusionEnvironmentFamilyIds(compartment)
	if err != nil {
		return err
	}
	for _, fusionEnvironmentFamilyId := range fusionEnvironmentFamilyIds {
		if ok := acctest.SweeperDefaultResourceId[fusionEnvironmentFamilyId]; !ok {
			deleteFusionEnvironmentFamilyRequest := oci_fusion_apps.DeleteFusionEnvironmentFamilyRequest{}

			deleteFusionEnvironmentFamilyRequest.FusionEnvironmentFamilyId = &fusionEnvironmentFamilyId

			deleteFusionEnvironmentFamilyRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "fusion_apps")
			_, error := fusionApplicationsClient.DeleteFusionEnvironmentFamily(context.Background(), deleteFusionEnvironmentFamilyRequest)
			if error != nil {
				fmt.Printf("Error deleting FusionEnvironmentFamily %s %s, It is possible that the resource is already deleted. Please verify manually \n", fusionEnvironmentFamilyId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &fusionEnvironmentFamilyId, FusionAppsFusionEnvironmentFamilySweepWaitCondition, time.Duration(3*time.Minute),
				FusionAppsFusionEnvironmentFamilySweepResponseFetchOperation, "fusion_apps", true)
		}
	}
	return nil
}

func getFusionAppsFusionEnvironmentFamilyIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "FusionEnvironmentFamilyId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	fusionApplicationsClient := acctest.GetTestClients(&schema.ResourceData{}).FusionApplicationsClient()

	listFusionEnvironmentFamiliesRequest := oci_fusion_apps.ListFusionEnvironmentFamiliesRequest{}
	listFusionEnvironmentFamiliesRequest.CompartmentId = &compartmentId
	listFusionEnvironmentFamiliesRequest.LifecycleState = oci_fusion_apps.FusionEnvironmentFamilyLifecycleStateActive
	listFusionEnvironmentFamiliesResponse, err := fusionApplicationsClient.ListFusionEnvironmentFamilies(context.Background(), listFusionEnvironmentFamiliesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting FusionEnvironmentFamily list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, fusionEnvironmentFamily := range listFusionEnvironmentFamiliesResponse.Items {
		id := *fusionEnvironmentFamily.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "FusionEnvironmentFamilyId", id)
	}
	return resourceIds, nil
}

func FusionAppsFusionEnvironmentFamilySweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if fusionEnvironmentFamilyResponse, ok := response.Response.(oci_fusion_apps.GetFusionEnvironmentFamilyResponse); ok {
		return fusionEnvironmentFamilyResponse.LifecycleState != oci_fusion_apps.FusionEnvironmentFamilyLifecycleStateDeleted
	}
	return false
}

func FusionAppsFusionEnvironmentFamilySweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.FusionApplicationsClient().GetFusionEnvironmentFamily(context.Background(), oci_fusion_apps.GetFusionEnvironmentFamilyRequest{
		FusionEnvironmentFamilyId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
