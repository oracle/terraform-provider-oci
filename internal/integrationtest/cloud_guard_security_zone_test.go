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
	oci_cloud_guard "github.com/oracle/oci-go-sdk/v65/cloudguard"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	CloudGuardSecurityZoneRequiredOnlyResource = CloudGuardSecurityZoneResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_security_zone", "test_security_zone", acctest.Required, acctest.Create, CloudGuardSecurityZoneRepresentation)

	CloudGuardSecurityZoneResourceConfig = CloudGuardSecurityZoneResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_security_zone", "test_security_zone", acctest.Optional, acctest.Update, CloudGuardSecurityZoneRepresentation)

	CloudGuardCloudGuardSecurityZoneSingularDataSourceRepresentation = map[string]interface{}{
		"security_zone_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_cloud_guard_security_zone.test_security_zone.id}`},
	}

	CloudGuardCloudGuardSecurityZoneDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":                          acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":                                    acctest.Representation{RepType: acctest.Optional, Create: `${oci_cloud_guard_security_zone.test_security_zone.id}`},
		"is_required_security_zones_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"security_recipe_id":                    acctest.Representation{RepType: acctest.Optional, Create: `${oci_cloud_guard_security_recipe.test_security_recipe.id}`},
		"state":                                 acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":                                acctest.RepresentationGroup{RepType: acctest.Required, Group: CloudGuardSecurityZoneDataSourceFilterRepresentation}}
	CloudGuardSecurityZoneDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_cloud_guard_security_zone.test_security_zone.id}`}},
	}

	CloudGuardSecurityZoneRepresentation = map[string]interface{}{
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":            acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"security_zone_recipe_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_cloud_guard_security_recipe.test_security_recipe.id}`},
		"defined_tags":            acctest.Representation{RepType: acctest.Optional, Create: `${tomap({"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "value"})}`, Update: `${tomap({"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "updatedValue"})}`},
		"description":             acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":           acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
	}

	CloudGuardSecurityZoneResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_security_recipe", "test_security_recipe", acctest.Required, acctest.Create, CloudGuardSecurityRecipeRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_guard_security_policies", "oracle_security_policy", acctest.Required, acctest.Create, CloudGuardSecurityPolicyDataSourceRepresentationPluralDataSource) +
		DefinedTagsDependencies
)

// issue-routing-tag: cloud_guard/default
func TestCloudGuardSecurityZoneResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCloudGuardSecurityZoneResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_id_for_update")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_cloud_guard_security_zone.test_security_zone"
	datasourceName := "data.oci_cloud_guard_security_zones.test_security_zones"
	singularDatasourceName := "data.oci_cloud_guard_security_zone.test_security_zone"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CloudGuardSecurityZoneResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_security_zone", "test_security_zone", acctest.Optional, acctest.Create, CloudGuardSecurityZoneRepresentation), "cloudguard", "securityZone", t)

	acctest.ResourceTest(t, testAccCheckCloudGuardSecurityZoneDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + CloudGuardSecurityZoneResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_security_zone", "test_security_zone", acctest.Required, acctest.Create, CloudGuardSecurityZoneRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "security_zone_recipe_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + CloudGuardSecurityZoneResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + CloudGuardSecurityZoneResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_security_zone", "test_security_zone", acctest.Optional, acctest.Create, CloudGuardSecurityZoneRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "security_zone_recipe_id"),

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
			Config: config + compartmentIdVariableStr + CloudGuardSecurityZoneResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_security_zone", "test_security_zone", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(CloudGuardSecurityZoneRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "security_zone_recipe_id"),

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
			Config: config + compartmentIdVariableStr + CloudGuardSecurityZoneResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_security_zone", "test_security_zone", acctest.Optional, acctest.Update, CloudGuardSecurityZoneRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "security_zone_recipe_id"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_guard_security_zones", "test_security_zones", acctest.Optional, acctest.Update, CloudGuardCloudGuardSecurityZoneDataSourceRepresentation) +
				compartmentIdVariableStr + CloudGuardSecurityZoneResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_security_zone", "test_security_zone", acctest.Optional, acctest.Update, CloudGuardSecurityZoneRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				//resource.TestCheckResourceAttr(datasourceName, "id", "id"),
				resource.TestCheckResourceAttr(datasourceName, "is_required_security_zones_in_subtree", "false"),
				resource.TestCheckResourceAttrSet(datasourceName, "security_recipe_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "security_zone_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "security_zone_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_guard_security_zone", "test_security_zone", acctest.Required, acctest.Create, CloudGuardCloudGuardSecurityZoneSingularDataSourceRepresentation) +
				compartmentIdVariableStr + CloudGuardSecurityZoneResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "security_zone_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "inherited_by_compartments.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "security_zone_target_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + CloudGuardSecurityZoneRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckCloudGuardSecurityZoneDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).CloudGuardClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_cloud_guard_security_zone" {
			noResourceFound = false
			request := oci_cloud_guard.GetSecurityZoneRequest{}

			tmp := rs.Primary.ID
			request.SecurityZoneId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "cloud_guard")

			response, err := client.GetSecurityZone(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_cloud_guard.LifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("CloudGuardSecurityZone") {
		resource.AddTestSweepers("CloudGuardSecurityZone", &resource.Sweeper{
			Name:         "CloudGuardSecurityZone",
			Dependencies: acctest.DependencyGraph["securityZone"],
			F:            sweepCloudGuardSecurityZoneResource,
		})
	}
}

func sweepCloudGuardSecurityZoneResource(compartment string) error {
	cloudGuardClient := acctest.GetTestClients(&schema.ResourceData{}).CloudGuardClient()
	securityZoneIds, err := getCloudGuardSecurityZoneIds(compartment)
	if err != nil {
		return err
	}
	for _, securityZoneId := range securityZoneIds {
		if ok := acctest.SweeperDefaultResourceId[securityZoneId]; !ok {
			deleteSecurityZoneRequest := oci_cloud_guard.DeleteSecurityZoneRequest{}

			deleteSecurityZoneRequest.SecurityZoneId = &securityZoneId

			deleteSecurityZoneRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "cloud_guard")
			_, error := cloudGuardClient.DeleteSecurityZone(context.Background(), deleteSecurityZoneRequest)
			if error != nil {
				fmt.Printf("Error deleting SecurityZone %s %s, It is possible that the resource is already deleted. Please verify manually \n", securityZoneId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &securityZoneId, CloudGuardSecurityZoneSweepWaitCondition, time.Duration(3*time.Minute),
				CloudGuardSecurityZoneSweepResponseFetchOperation, "cloud_guard", true)
		}
	}
	return nil
}

func getCloudGuardSecurityZoneIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "SecurityZoneId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	cloudGuardClient := acctest.GetTestClients(&schema.ResourceData{}).CloudGuardClient()

	listSecurityZonesRequest := oci_cloud_guard.ListSecurityZonesRequest{}
	listSecurityZonesRequest.CompartmentId = &compartmentId
	listSecurityZonesRequest.LifecycleState = oci_cloud_guard.ListSecurityZonesLifecycleStateActive
	listSecurityZonesResponse, err := cloudGuardClient.ListSecurityZones(context.Background(), listSecurityZonesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting SecurityZone list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, securityZone := range listSecurityZonesResponse.Items {
		id := *securityZone.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "SecurityZoneId", id)
	}
	return resourceIds, nil
}

func CloudGuardSecurityZoneSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if securityZoneResponse, ok := response.Response.(oci_cloud_guard.GetSecurityZoneResponse); ok {
		return securityZoneResponse.LifecycleState != oci_cloud_guard.LifecycleStateDeleted
	}
	return false
}

func CloudGuardSecurityZoneSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.CloudGuardClient().GetSecurityZone(context.Background(), oci_cloud_guard.GetSecurityZoneRequest{
		SecurityZoneId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
