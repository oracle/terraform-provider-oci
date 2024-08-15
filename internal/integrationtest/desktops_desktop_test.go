// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/oracle/oci-go-sdk/v65/common"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	oci_desktops "github.com/oracle/oci-go-sdk/v65/desktops"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DesktopsDesktopSingularDataSourceRepresentation = map[string]interface{}{
		"desktop_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_desktops_desktop_pool_desktops.test_desktop_pool_desktops.desktop_pool_desktop_collection.0.items.0.desktop_id}`},
	}

	DesktopsDesktopDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"availability_domain": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"desktop_pool_id":     acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_desktops_desktop_pools.test_desktop_pools.desktop_pool_collection[0].items[0].id}`},
		//"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		//	"id":                  acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_desktops_desktop_pool_desktops.test_desktop_pool_desktops.desktop_pool_desktop_collection.0.items.0.desktop_id}`},
		"state": acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
	}

	DesktopsDesktopFromDesktopPoolDesktopsRepresentation = map[string]interface{}{
		"desktop_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_desktops_desktop_pool_desktops.test_desktop_pool_desktops.desktop_pool_desktop_collection.0.items.0.desktop_id}`,
			Update: `${data.oci_desktops_desktop_pool_desktops.test_desktop_pool_desktops.desktop_pool_desktop_collection.0.items.0.desktop_id}`},
		"display_name":  acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	DesktopsDesktopPoolFromCompartmentDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"availability_domain": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `testPool1`, Update: `testPool2`},
		"state":               acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
	}

	DesktopsDesktopPoolDesktopFromPoolDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"desktop_pool_id":     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_desktops_desktop_pools.test_desktop_pools.desktop_pool_collection[0].items[0].id}`},
		"availability_domain": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"state":               acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
	}

	DesktopsDesktopResourceDependencies = DesktopsDesktopPoolResourceDependencies
)

// issue-routing-tag: desktops/default
func TestDesktopsDesktopResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDesktopsDesktopResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_desktops_desktops.test_desktops"
	singularDatasourceName := "data.oci_desktops_desktop.test_desktop"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{

		// Create desktop pool which creates standby desktops and wait
		{
			Config: config + compartmentIdVariableStr + DesktopsDesktopPoolResourceDependencies + acctest.GenerateResourceFromRepresentationMap("oci_desktops_desktop_pool", "test_desktop_pool", acctest.Required, acctest.Create, DesktopsDesktopPoolRepresentation),
			Check:  delaysomeAndReturnNil(),
		},

		// verify datasource
		{
			Config: config + compartmentIdVariableStr + DesktopsDesktopResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_desktops_desktop_pool", "test_desktop_pool", acctest.Required, acctest.Create, DesktopsDesktopPoolRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_desktops_desktop_pools", "test_desktop_pools", acctest.Required, acctest.Create, DesktopsDesktopPoolFromCompartmentDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_desktops_desktops", "test_desktops", acctest.Optional, acctest.Create, DesktopsDesktopDataSourceRepresentation),

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				//resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "desktop_pool_id"),
				//resource.TestCheckResourceAttrSet(datasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttrSet(datasourceName, "desktop_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config + compartmentIdVariableStr + DesktopsDesktopResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_desktops_desktop_pool", "test_desktop_pool", acctest.Required, acctest.Create, DesktopsDesktopPoolRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_desktops_desktop_pools", "test_desktop_pools", acctest.Optional, acctest.Create, DesktopsDesktopPoolFromCompartmentDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_desktops_desktop_pool_desktops", "test_desktop_pool_desktops", acctest.Optional, acctest.Create, DesktopsDesktopPoolDesktopFromPoolDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_desktops_desktop", "test_desktop", acctest.Required, acctest.Create, DesktopsDesktopSingularDataSourceRepresentation),

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "desktop_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "device_policy.#", "1"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "hosting_options.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "pool_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_name", ""),
			),
		},
	})
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("DesktopsDesktop") {
		resource.AddTestSweepers("DesktopsDesktop", &resource.Sweeper{
			Name:         "DesktopsDesktop",
			Dependencies: acctest.DependencyGraph["desktop"],
			F:            sweepDesktopsDesktopResource,
		})
	}
}

func sweepDesktopsDesktopResource(compartment string) error {
	desktopServiceClient := acctest.GetTestClients(&schema.ResourceData{}).DesktopServiceClient()
	desktopIds, err := getDesktopsDesktopIds(compartment)
	if err != nil {
		return err
	}
	for _, desktopId := range desktopIds {
		if ok := acctest.SweeperDefaultResourceId[desktopId]; !ok {
			deleteDesktopRequest := oci_desktops.DeleteDesktopRequest{}

			deleteDesktopRequest.DesktopId = &desktopId

			deleteDesktopRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "desktops")
			_, error := desktopServiceClient.DeleteDesktop(context.Background(), deleteDesktopRequest)
			if error != nil {
				fmt.Printf("Error deleting Desktop %s %s, It is possible that the resource is already deleted. Please verify manually \n", desktopId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &desktopId, DesktopsDesktopSweepWaitCondition, time.Duration(3*time.Minute),
				DesktopsDesktopSweepResponseFetchOperation, "desktops", true)
		}
	}
	return nil
}

func getDesktopsDesktopIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DesktopId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	desktopServiceClient := acctest.GetTestClients(&schema.ResourceData{}).DesktopServiceClient()

	listDesktopsRequest := oci_desktops.ListDesktopsRequest{}
	listDesktopsRequest.CompartmentId = &compartmentId
	listDesktopsRequest.LifecycleState = GetLifecycleStateEnumStringValue(oci_desktops.LifecycleStateActive)
	listDesktopsResponse, err := desktopServiceClient.ListDesktops(context.Background(), listDesktopsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Desktop list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, desktop := range listDesktopsResponse.Items {
		id := *desktop.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DesktopId", id)
	}
	return resourceIds, nil
}

func DesktopsDesktopSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if desktopResponse, ok := response.Response.(oci_desktops.GetDesktopResponse); ok {
		return desktopResponse.LifecycleState != oci_desktops.LifecycleStateDeleted
	}
	return false
}

func DesktopsDesktopSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DesktopServiceClient().GetDesktop(context.Background(), oci_desktops.GetDesktopRequest{
		DesktopId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}

func delaysomeAndReturnNil() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		log.Println("Beginning wait ...")
		time.Sleep(60 * time.Second)
		log.Println("Ending wait ...")
		return nil
	}
}
