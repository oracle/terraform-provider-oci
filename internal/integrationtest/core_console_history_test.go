// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

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
	CoreConsoleHistoryRequiredOnlyResource = CoreConsoleHistoryResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_console_history", "test_console_history", acctest.Required, acctest.Create, CoreConsoleHistoryRepresentation)

	CoreCoreConsoleHistoryDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"availability_domain": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"instance_id":         acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_instance.test_instance.id}`},
		"state":               acctest.Representation{RepType: acctest.Optional, Create: `SUCCEEDED`},
		"filter":              acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreConsoleHistoryDataSourceFilterRepresentation}}
	CoreConsoleHistoryDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_console_history.test_console_history.id}`}},
	}

	CoreConsoleHistoryRepresentation = map[string]interface{}{
		"instance_id":   acctest.Representation{RepType: acctest.Required, Create: `${oci_core_instance.test_instance.id}`},
		"defined_tags":  acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":  acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	CoreConsoleHistoryResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		utils.OciImageIdsVariable +
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, CoreInstanceRepresentation) +
		AvailabilityDomainConfig +
		DefinedTagsDependencies
)

// issue-routing-tag: core/computeSharedOwnershipVmAndBm
func TestCoreConsoleHistoryResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreConsoleHistoryResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_console_history.test_console_history"
	datasourceName := "data.oci_core_console_histories.test_console_histories"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CoreConsoleHistoryResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_core_console_history", "test_console_history", acctest.Optional, acctest.Create, CoreConsoleHistoryRepresentation), "core", "consoleHistory", t)

	acctest.ResourceTest(t, testAccCheckCoreConsoleHistoryDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + CoreConsoleHistoryResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_console_history", "test_console_history", acctest.Required, acctest.Create, CoreConsoleHistoryRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "instance_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + CoreConsoleHistoryResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + CoreConsoleHistoryResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_console_history", "test_console_history", acctest.Optional, acctest.Create, CoreConsoleHistoryRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
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

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + CoreConsoleHistoryResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_console_history", "test_console_history", acctest.Optional, acctest.Update, CoreConsoleHistoryRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_console_histories", "test_console_histories", acctest.Optional, acctest.Update, CoreCoreConsoleHistoryDataSourceRepresentation) +
				compartmentIdVariableStr + CoreConsoleHistoryResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_console_history", "test_console_history", acctest.Optional, acctest.Update, CoreConsoleHistoryRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "instance_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "SUCCEEDED"),

				resource.TestCheckResourceAttr(datasourceName, "console_histories.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "console_histories.0.availability_domain"),
				resource.TestCheckResourceAttrSet(datasourceName, "console_histories.0.compartment_id"),
				resource.TestCheckResourceAttr(datasourceName, "console_histories.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "console_histories.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "console_histories.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "console_histories.0.instance_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "console_histories.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "console_histories.0.time_created"),
			),
		},
		// verify resource import
		{
			Config:                  config + CoreConsoleHistoryRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckCoreConsoleHistoryDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ComputeClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_console_history" {
			noResourceFound = false
			request := oci_core.GetConsoleHistoryRequest{}

			tmp := rs.Primary.ID
			request.InstanceConsoleHistoryId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")

			_, err := client.GetConsoleHistory(context.Background(), request)

			if err == nil {
				return fmt.Errorf("resource still exists")
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
	if !acctest.InSweeperExcludeList("CoreConsoleHistory") {
		resource.AddTestSweepers("CoreConsoleHistory", &resource.Sweeper{
			Name:         "CoreConsoleHistory",
			Dependencies: acctest.DependencyGraph["consoleHistory"],
			F:            sweepCoreConsoleHistoryResource,
		})
	}
}

func sweepCoreConsoleHistoryResource(compartment string) error {
	computeClient := acctest.GetTestClients(&schema.ResourceData{}).ComputeClient()
	consoleHistoryIds, err := getCoreConsoleHistoryIds(compartment)
	if err != nil {
		return err
	}
	for _, consoleHistoryId := range consoleHistoryIds {
		if ok := acctest.SweeperDefaultResourceId[consoleHistoryId]; !ok {
			deleteConsoleHistoryRequest := oci_core.DeleteConsoleHistoryRequest{}

			deleteConsoleHistoryRequest.InstanceConsoleHistoryId = &consoleHistoryId

			deleteConsoleHistoryRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")
			_, error := computeClient.DeleteConsoleHistory(context.Background(), deleteConsoleHistoryRequest)
			if error != nil {
				fmt.Printf("Error deleting ConsoleHistory %s %s, It is possible that the resource is already deleted. Please verify manually \n", consoleHistoryId, error)
				continue
			}
		}
	}
	return nil
}

func getCoreConsoleHistoryIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ConsoleHistoryId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	computeClient := acctest.GetTestClients(&schema.ResourceData{}).ComputeClient()

	listConsoleHistoriesRequest := oci_core.ListConsoleHistoriesRequest{}
	listConsoleHistoriesRequest.CompartmentId = &compartmentId
	listConsoleHistoriesResponse, err := computeClient.ListConsoleHistories(context.Background(), listConsoleHistoriesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ConsoleHistory list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, consoleHistory := range listConsoleHistoriesResponse.Items {
		id := *consoleHistory.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ConsoleHistoryId", id)
	}
	return resourceIds, nil
}
