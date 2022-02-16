// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v58/common"
	oci_waas "github.com/oracle/oci-go-sdk/v58/waas"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	AddressListRequiredOnlyResource = AddressListResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_waas_address_list", "test_address_list", acctest.Required, acctest.Create, addressListRepresentation)

	AddressListResourceConfig = AddressListResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_waas_address_list", "test_address_list", acctest.Optional, acctest.Update, addressListRepresentation)

	addressListSingularDataSourceRepresentation = map[string]interface{}{
		"address_list_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_waas_address_list.test_address_list.id}`},
	}

	addressListDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"ids":                                   acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_waas_address_list.test_address_list.id}`}},
		"names":                                 acctest.Representation{RepType: acctest.Optional, Create: []string{`displayName2`}},
		"states":                                acctest.Representation{RepType: acctest.Optional, Create: []string{`ACTIVE`}},
		"time_created_greater_than_or_equal_to": acctest.Representation{RepType: acctest.Optional, Create: `2018-01-01T00:00:00.000Z`},
		"time_created_less_than":                acctest.Representation{RepType: acctest.Optional, Create: `2038-01-01T00:00:00.000Z`},
		"filter":                                acctest.RepresentationGroup{RepType: acctest.Required, Group: addressListDataSourceFilterRepresentation}}
	addressListDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_waas_address_list.test_address_list.id}`}},
	}

	addressListRepresentation = map[string]interface{}{
		"addresses":      acctest.Representation{RepType: acctest.Required, Create: []string{`0.0.0.0/16`}, Update: []string{`0.0.0.0/20`}},
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	AddressListResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: waas/default
func TestWaasAddressListResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestWaasAddressListResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_waas_address_list.test_address_list"
	datasourceName := "data.oci_waas_address_lists.test_address_lists"
	singularDatasourceName := "data.oci_waas_address_list.test_address_list"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+AddressListResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_waas_address_list", "test_address_list", acctest.Optional, acctest.Create, addressListRepresentation), "waas", "addressList", t)

	acctest.ResourceTest(t, testAccCheckWaasAddressListDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + AddressListResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_waas_address_list", "test_address_list", acctest.Required, acctest.Create, addressListRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "addresses.#", "1"),
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
			Config: config + compartmentIdVariableStr + AddressListResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + AddressListResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_waas_address_list", "test_address_list", acctest.Optional, acctest.Create, addressListRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "addresses.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + AddressListResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_waas_address_list", "test_address_list", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(addressListRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "addresses.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),

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
			Config: config + compartmentIdVariableStr + AddressListResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_waas_address_list", "test_address_list", acctest.Optional, acctest.Update, addressListRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "addresses.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_waas_address_lists", "test_address_lists", acctest.Optional, acctest.Update, addressListDataSourceRepresentation) +
				compartmentIdVariableStr + AddressListResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_waas_address_list", "test_address_list", acctest.Optional, acctest.Update, addressListRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "ids.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "names.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "states.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_created_greater_than_or_equal_to"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_created_less_than"),

				resource.TestCheckResourceAttr(datasourceName, "address_lists.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "address_lists.0.address_count"),
				resource.TestCheckResourceAttr(datasourceName, "address_lists.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "address_lists.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "address_lists.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "address_lists.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "address_lists.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "address_lists.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_waas_address_list", "test_address_list", acctest.Required, acctest.Create, addressListSingularDataSourceRepresentation) +
				compartmentIdVariableStr + AddressListResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "address_list_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "address_count"),
				resource.TestCheckResourceAttr(singularDatasourceName, "addresses.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + AddressListResourceConfig,
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckWaasAddressListDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).WaasClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_waas_address_list" {
			noResourceFound = false
			request := oci_waas.GetAddressListRequest{}

			tmp := rs.Primary.ID
			request.AddressListId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "waas")

			response, err := client.GetAddressList(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_waas.LifecycleStatesDeleted): true,
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
	if !acctest.InSweeperExcludeList("WaasAddressList") {
		resource.AddTestSweepers("WaasAddressList", &resource.Sweeper{
			Name:         "WaasAddressList",
			Dependencies: acctest.DependencyGraph["addressList"],
			F:            sweepWaasAddressListResource,
		})
	}
}

func sweepWaasAddressListResource(compartment string) error {
	waasClient := acctest.GetTestClients(&schema.ResourceData{}).WaasClient()
	addressListIds, err := getAddressListIds(compartment)
	if err != nil {
		return err
	}
	for _, addressListId := range addressListIds {
		if ok := acctest.SweeperDefaultResourceId[addressListId]; !ok {
			deleteAddressListRequest := oci_waas.DeleteAddressListRequest{}

			deleteAddressListRequest.AddressListId = &addressListId

			deleteAddressListRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "waas")
			_, error := waasClient.DeleteAddressList(context.Background(), deleteAddressListRequest)
			if error != nil {
				fmt.Printf("Error deleting AddressList %s %s, It is possible that the resource is already deleted. Please verify manually \n", addressListId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &addressListId, addressListSweepWaitCondition, time.Duration(3*time.Minute),
				addressListSweepResponseFetchOperation, "waas", true)
		}
	}
	return nil
}

func getAddressListIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "AddressListId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	waasClient := acctest.GetTestClients(&schema.ResourceData{}).WaasClient()

	listAddressListsRequest := oci_waas.ListAddressListsRequest{}
	listAddressListsRequest.CompartmentId = &compartmentId
	listAddressListsRequest.LifecycleState = []oci_waas.LifecycleStatesEnum{oci_waas.LifecycleStatesActive}
	listAddressListsResponse, err := waasClient.ListAddressLists(context.Background(), listAddressListsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting AddressList list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, addressList := range listAddressListsResponse.Items {
		id := *addressList.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "AddressListId", id)
	}
	return resourceIds, nil
}

func addressListSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if addressListResponse, ok := response.Response.(oci_waas.GetAddressListResponse); ok {
		return addressListResponse.LifecycleState != oci_waas.LifecycleStatesDeleted
	}
	return false
}

func addressListSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.WaasClient().GetAddressList(context.Background(), oci_waas.GetAddressListRequest{
		AddressListId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
