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
	oci_compute_cloud_at_customer "github.com/oracle/oci-go-sdk/v65/computecloudatcustomer"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	ComputeCloudAtCustomerCccInfrastructureRequiredOnlyResource = ComputeCloudAtCustomerCccInfrastructureResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_compute_cloud_at_customer_ccc_infrastructure", "test_ccc_infrastructure", acctest.Required, acctest.Create, ComputeCloudAtCustomerCccInfrastructureRepresentation)

	ComputeCloudAtCustomerCccInfrastructureResourceConfig = ComputeCloudAtCustomerCccInfrastructureResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_compute_cloud_at_customer_ccc_infrastructure", "test_ccc_infrastructure", acctest.Optional, acctest.Update, ComputeCloudAtCustomerCccInfrastructureRepresentation)

	ComputeCloudAtCustomerCccInfrastructureSingularDataSourceRepresentation = map[string]interface{}{
		"ccc_infrastructure_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_compute_cloud_at_customer_ccc_infrastructure.test_ccc_infrastructure.id}`},
	}

	ComputeCloudAtCustomerCccInfrastructureDataSourceRepresentation = map[string]interface{}{
		"access_level":              acctest.Representation{RepType: acctest.Optional, Create: `RESTRICTED`},
		"ccc_infrastructure_id":     acctest.Representation{RepType: acctest.Optional, Create: `${oci_compute_cloud_at_customer_ccc_infrastructure.test_ccc_infrastructure.id}`},
		"compartment_id":            acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"display_name":              acctest.Representation{RepType: acctest.Optional, Create: `example_cccInfrastructure`, Update: `displayName2`},
		"display_name_contains":     acctest.Representation{RepType: acctest.Optional, Create: `displayNameContains`},
		"state":                     acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: ComputeCloudAtCustomerCccInfrastructureDataSourceFilterRepresentation}}
	ComputeCloudAtCustomerCccInfrastructureDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_compute_cloud_at_customer_ccc_infrastructure.test_ccc_infrastructure.id}`}},
	}

	ComputeCloudAtCustomerCccInfrastructureRepresentation = map[string]interface{}{
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":            acctest.Representation{RepType: acctest.Required, Create: `example_cccInfrastructure`, Update: `displayName2`},
		"subnet_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"ccc_upgrade_schedule_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_compute_cloud_at_customer_ccc_upgrade_schedule.test_ccc_upgrade_schedule.id}`},
		"connection_details":      acctest.Representation{RepType: acctest.Optional, Create: `connectionDetails`, Update: `connectionDetails2`},
		"connection_state":        acctest.Representation{RepType: acctest.Optional, Create: `REJECT`, Update: `REJECT`},
		"defined_tags":            acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":             acctest.Representation{RepType: acctest.Optional, Create: `Datacenter 231`, Update: `description2`},
		"freeform_tags":           acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
	}

	ComputeCloudAtCustomerCccInfrastructureResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_compute_cloud_at_customer_ccc_infrastructure", "test_ccc_infrastructure_def", acctest.Required, acctest.Create, ComputeCloudAtCustomerCccInfrastructureRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_compute_cloud_at_customer_ccc_upgrade_schedule", "test_ccc_upgrade_schedule", acctest.Required, acctest.Create, ComputeCloudAtCustomerCccUpgradeScheduleRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: compute_cloud_at_customer/default
func TestComputeCloudAtCustomerCccInfrastructureResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestComputeCloudAtCustomerCccInfrastructureResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_compute_cloud_at_customer_ccc_infrastructure.test_ccc_infrastructure"
	// resourceNameDef := "oci_compute_cloud_at_customer_ccc_infrastructure.test_ccc_infrastructure_default"
	datasourceName := "data.oci_compute_cloud_at_customer_ccc_infrastructures.test_ccc_infrastructures"
	singularDatasourceName := "data.oci_compute_cloud_at_customer_ccc_infrastructure.test_ccc_infrastructure"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ComputeCloudAtCustomerCccInfrastructureResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_compute_cloud_at_customer_ccc_infrastructure", "test_ccc_infrastructure", acctest.Optional, acctest.Create, ComputeCloudAtCustomerCccInfrastructureRepresentation), "computecloudatcustomer", "cccInfrastructure", t)

	acctest.ResourceTest(t, testAccCheckComputeCloudAtCustomerCccInfrastructureDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ComputeCloudAtCustomerCccInfrastructureResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_compute_cloud_at_customer_ccc_infrastructure", "test_ccc_infrastructure", acctest.Required, acctest.Create, ComputeCloudAtCustomerCccInfrastructureRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_cccInfrastructure"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + ComputeCloudAtCustomerCccInfrastructureResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ComputeCloudAtCustomerCccInfrastructureResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_compute_cloud_at_customer_ccc_infrastructure", "test_ccc_infrastructure", acctest.Optional, acctest.Create, ComputeCloudAtCustomerCccInfrastructureRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "ccc_upgrade_schedule_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "connection_details", "connectionDetails"),
				resource.TestCheckResourceAttr(resourceName, "connection_state", "REJECT"),
				resource.TestCheckResourceAttr(resourceName, "description", "Datacenter 231"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_cccInfrastructure"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + ComputeCloudAtCustomerCccInfrastructureResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_compute_cloud_at_customer_ccc_infrastructure", "test_ccc_infrastructure", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(ComputeCloudAtCustomerCccInfrastructureRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "ccc_upgrade_schedule_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "connection_details", "connectionDetails"),
				resource.TestCheckResourceAttr(resourceName, "connection_state", "REJECT"),
				resource.TestCheckResourceAttr(resourceName, "description", "Datacenter 231"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_cccInfrastructure"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
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
			Config: config + compartmentIdVariableStr + ComputeCloudAtCustomerCccInfrastructureResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_compute_cloud_at_customer_ccc_infrastructure", "test_ccc_infrastructure", acctest.Optional, acctest.Update, ComputeCloudAtCustomerCccInfrastructureRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "ccc_upgrade_schedule_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "connection_details", "connectionDetails2"),
				resource.TestCheckResourceAttr(resourceName, "connection_state", "REJECT"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_compute_cloud_at_customer_ccc_infrastructures", "test_ccc_infrastructures", acctest.Optional, acctest.Update, ComputeCloudAtCustomerCccInfrastructureDataSourceRepresentation) +
				compartmentIdVariableStr + ComputeCloudAtCustomerCccInfrastructureResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_compute_cloud_at_customer_ccc_infrastructure", "test_ccc_infrastructure", acctest.Optional, acctest.Update, ComputeCloudAtCustomerCccInfrastructureRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "access_level", "RESTRICTED"),
				resource.TestCheckResourceAttrSet(datasourceName, "ccc_infrastructure_id"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "false"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "display_name_contains", "displayNameContains"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "ccc_infrastructure_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "ccc_infrastructure_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_compute_cloud_at_customer_ccc_infrastructure", "test_ccc_infrastructure", acctest.Required, acctest.Create, ComputeCloudAtCustomerCccInfrastructureSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ComputeCloudAtCustomerCccInfrastructureResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ccc_infrastructure_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "connection_details", "connectionDetails2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "connection_state", "REJECT"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "infrastructure_inventory.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "infrastructure_network_configuration.#", "1"),
				// resource.TestCheckResourceAttrSet(singularDatasourceName, "provisioning_fingerprint"),
				// resource.TestCheckResourceAttrSet(singularDatasourceName, "provisioning_pin"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "short_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "upgrade_information.#", "1"),
			),
		},
		// verify resource import
		{
			Config:                  config + ComputeCloudAtCustomerCccInfrastructureRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckComputeCloudAtCustomerCccInfrastructureDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ComputeCloudAtCustomerClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_compute_cloud_at_customer_ccc_infrastructure" {
			noResourceFound = false
			request := oci_compute_cloud_at_customer.GetCccInfrastructureRequest{}

			tmp := rs.Primary.ID
			request.CccInfrastructureId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "compute_cloud_at_customer")

			response, err := client.GetCccInfrastructure(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_compute_cloud_at_customer.CccInfrastructureLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("ComputeCloudAtCustomerCccInfrastructure") {
		resource.AddTestSweepers("ComputeCloudAtCustomerCccInfrastructure", &resource.Sweeper{
			Name:         "ComputeCloudAtCustomerCccInfrastructure",
			Dependencies: acctest.DependencyGraph["cccInfrastructure"],
			F:            sweepComputeCloudAtCustomerCccInfrastructureResource,
		})
	}
}

func sweepComputeCloudAtCustomerCccInfrastructureResource(compartment string) error {
	computeCloudAtCustomerClient := acctest.GetTestClients(&schema.ResourceData{}).ComputeCloudAtCustomerClient()
	cccInfrastructureIds, err := getComputeCloudAtCustomerCccInfrastructureIds(compartment)
	if err != nil {
		return err
	}
	for _, cccInfrastructureId := range cccInfrastructureIds {
		if ok := acctest.SweeperDefaultResourceId[cccInfrastructureId]; !ok {
			deleteCccInfrastructureRequest := oci_compute_cloud_at_customer.DeleteCccInfrastructureRequest{}

			deleteCccInfrastructureRequest.CccInfrastructureId = &cccInfrastructureId

			deleteCccInfrastructureRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "compute_cloud_at_customer")
			_, error := computeCloudAtCustomerClient.DeleteCccInfrastructure(context.Background(), deleteCccInfrastructureRequest)
			if error != nil {
				fmt.Printf("Error deleting CccInfrastructure %s %s, It is possible that the resource is already deleted. Please verify manually \n", cccInfrastructureId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &cccInfrastructureId, ComputeCloudAtCustomerCccInfrastructureSweepWaitCondition, time.Duration(3*time.Minute),
				ComputeCloudAtCustomerCccInfrastructureSweepResponseFetchOperation, "compute_cloud_at_customer", true)
		}
	}
	return nil
}

func getComputeCloudAtCustomerCccInfrastructureIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "CccInfrastructureId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	computeCloudAtCustomerClient := acctest.GetTestClients(&schema.ResourceData{}).ComputeCloudAtCustomerClient()

	listCccInfrastructuresRequest := oci_compute_cloud_at_customer.ListCccInfrastructuresRequest{}
	listCccInfrastructuresRequest.CompartmentId = &compartmentId
	listCccInfrastructuresRequest.LifecycleState = oci_compute_cloud_at_customer.CccInfrastructureLifecycleStateActive
	listCccInfrastructuresResponse, err := computeCloudAtCustomerClient.ListCccInfrastructures(context.Background(), listCccInfrastructuresRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting CccInfrastructure list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, cccInfrastructure := range listCccInfrastructuresResponse.Items {
		id := *cccInfrastructure.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "CccInfrastructureId", id)
	}
	return resourceIds, nil
}

func ComputeCloudAtCustomerCccInfrastructureSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if cccInfrastructureResponse, ok := response.Response.(oci_compute_cloud_at_customer.GetCccInfrastructureResponse); ok {
		return cccInfrastructureResponse.LifecycleState != oci_compute_cloud_at_customer.CccInfrastructureLifecycleStateDeleted
	}
	return false
}

func ComputeCloudAtCustomerCccInfrastructureSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ComputeCloudAtCustomerClient().GetCccInfrastructure(context.Background(), oci_compute_cloud_at_customer.GetCccInfrastructureRequest{
		CccInfrastructureId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
