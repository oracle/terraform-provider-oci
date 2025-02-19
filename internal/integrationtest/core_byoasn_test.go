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
	oci_core "github.com/oracle/oci-go-sdk/v65/core"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	CoreByoasnRequiredOnlyResource = CoreByoasnResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_byoasn", "test_byoasn", acctest.Required, acctest.Create, CoreByoasnRepresentation)

	CoreByoasnResourceConfig = CoreByoasnResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_byoasn", "test_byoasn", acctest.Optional, acctest.Update, CoreByoasnRepresentation)

	CoreByoasnSingularDataSourceRepresentation = map[string]interface{}{
		"byoasn_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_byoasn.test_byoasn.id}`},
	}

	CoreByoasnDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreByoasnDataSourceFilterRepresentation}}
	CoreByoasnDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_byoasn.test_byoasn.id}`}},
	}

	CoreByoasnRepresentation = map[string]interface{}{
		"asn":            acctest.Representation{RepType: acctest.Required, Create: `11`},
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	CoreByoasnResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: core/vcnip
func TestCoreByoasnResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreByoasnResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_byoasn.test_byoasn"
	datasourceName := "data.oci_core_byoasns.test_byoasns"
	singularDatasourceName := "data.oci_core_byoasn.test_byoasn"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CoreByoasnResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_core_byoasn", "test_byoasn", acctest.Optional, acctest.Create, CoreByoasnRepresentation), "core", "byoasn", t)

	acctest.ResourceTest(t, testAccCheckCoreByoasnDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + CoreByoasnResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_byoasn", "test_byoasn", acctest.Required, acctest.Create, CoreByoasnRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "asn", "11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + CoreByoasnResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + CoreByoasnResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_byoasn", "test_byoasn", acctest.Optional, acctest.Create, CoreByoasnRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "asn", "11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "validation_token"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + CoreByoasnResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_byoasn", "test_byoasn", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(CoreByoasnRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "asn", "11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "validation_token"),

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
			Config: config + compartmentIdVariableStr + CoreByoasnResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_byoasn", "test_byoasn", acctest.Optional, acctest.Update, CoreByoasnRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "asn", "11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "validation_token"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_byoasns", "test_byoasns", acctest.Optional, acctest.Update, CoreByoasnDataSourceRepresentation) +
				compartmentIdVariableStr + CoreByoasnResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_byoasn", "test_byoasn", acctest.Optional, acctest.Update, CoreByoasnRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttr(datasourceName, "byoasn_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "byoasn_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_byoasn", "test_byoasn", acctest.Required, acctest.Create, CoreByoasnSingularDataSourceRepresentation) +
				compartmentIdVariableStr + CoreByoasnResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "byoasn_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "asn", "11"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "byoip_ranges.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "time_validated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "validation_token"),
			),
		},
		// verify resource import
		{
			Config:                  config + CoreByoasnRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckCoreByoasnDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).VirtualNetworkClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_byoasn" {
			noResourceFound = false
			request := oci_core.GetByoasnRequest{}

			tmp := rs.Primary.ID
			request.ByoasnId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")

			response, err := client.GetByoasn(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.ByoasnLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("CoreByoasn") {
		resource.AddTestSweepers("CoreByoasn", &resource.Sweeper{
			Name:         "CoreByoasn",
			Dependencies: acctest.DependencyGraph["byoasn"],
			F:            sweepCoreByoasnResource,
		})
	}
}

func sweepCoreByoasnResource(compartment string) error {
	virtualNetworkClient := acctest.GetTestClients(&schema.ResourceData{}).VirtualNetworkClient()
	byoasnIds, err := getCoreByoasnIds(compartment)
	if err != nil {
		return err
	}
	for _, byoasnId := range byoasnIds {
		if ok := acctest.SweeperDefaultResourceId[byoasnId]; !ok {
			deleteByoasnRequest := oci_core.DeleteByoasnRequest{}

			deleteByoasnRequest.ByoasnId = &byoasnId

			deleteByoasnRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")
			_, error := virtualNetworkClient.DeleteByoasn(context.Background(), deleteByoasnRequest)
			if error != nil {
				fmt.Printf("Error deleting Byoasn %s %s, It is possible that the resource is already deleted. Please verify manually \n", byoasnId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &byoasnId, CoreByoasnSweepWaitCondition, time.Duration(3*time.Minute),
				CoreByoasnSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getCoreByoasnIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ByoasnId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	virtualNetworkClient := acctest.GetTestClients(&schema.ResourceData{}).VirtualNetworkClient()
	listByoasnsRequest := oci_core.ListByoasnsRequest{}
	listByoasnsRequest.CompartmentId = &compartmentId
	tmp := "ACTIVE"
	listByoasnsRequest.LifecycleState = &tmp
	listByoasnsResponse, err := virtualNetworkClient.ListByoasns(context.Background(), listByoasnsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Byoasn list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, byoasn := range listByoasnsResponse.Items {
		id := *byoasn.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ByoasnId", id)
	}
	return resourceIds, nil
}

func CoreByoasnSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if byoasnResponse, ok := response.Response.(oci_core.GetByoasnResponse); ok {
		return byoasnResponse.LifecycleState != oci_core.ByoasnLifecycleStateDeleted
	}
	return false
}

func CoreByoasnSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.VirtualNetworkClient().GetByoasn(context.Background(), oci_core.GetByoasnRequest{
		ByoasnId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
