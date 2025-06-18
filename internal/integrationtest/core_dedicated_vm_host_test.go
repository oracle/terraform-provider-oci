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
	displayNameForCreate        = `tf-test-display-name-create`
	displayNameForUpdate        = `tf-test-display-name-update`
	dvhShape                    = "DVH.DenseIO.E4.128"
	capacityBinsSizeForDvhShape = "4"
	vmShape                     = `VM.DenseIO.E4.Flex`

	CoreDedicatedVmHostRequiredOnlyResource = CoreDedicatedVmHostResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_dedicated_vm_host", "test_dedicated_vm_host", acctest.Required, acctest.Create, CoreDedicatedVmHostRepresentation)

	CoreDedicatedVmHostResourceConfig = CoreDedicatedVmHostResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_dedicated_vm_host", "test_dedicated_vm_host", acctest.Optional, acctest.Update, CoreDedicatedVmHostRepresentation)

	CoreCoreDedicatedVmHostSingularDataSourceRepresentation = map[string]interface{}{
		"dedicated_vm_host_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_dedicated_vm_host.test_dedicated_vm_host.id}`},
	}

	CoreDedicatedVmHostDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"availability_domain": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: displayNameForCreate, Update: displayNameForUpdate},
		"instance_shape_name": acctest.Representation{RepType: acctest.Optional, Create: vmShape},
		"remaining_memory_in_gbs_greater_than_or_equal_to": acctest.Representation{RepType: acctest.Optional, Create: `512.0`},
		"remaining_ocpus_greater_than_or_equal_to":         acctest.Representation{RepType: acctest.Optional, Create: `32.0`},
		"state":  acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter": acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreDedicatedVmHostDataSourceFilterRepresentation}}
	CoreDedicatedVmHostDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_dedicated_vm_host.test_dedicated_vm_host.id}`}},
	}

	CoreDedicatedVmHostRepresentation = map[string]interface{}{
		"availability_domain":          acctest.Representation{RepType: acctest.Required, Create: `UgLr:MX-QUERETARO-1-AD-1`},
		"compartment_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"dedicated_vm_host_shape":      acctest.Representation{RepType: acctest.Required, Create: dvhShape},
		"display_name":                 acctest.Representation{RepType: acctest.Optional, Create: displayNameForCreate, Update: displayNameForUpdate},
		"freeform_tags":                acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"FreeFormState": "Created"}, Update: map[string]string{"FreeFormState": "Updated"}},
		"placement_constraint_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreDedicatedVmHostPlacementConstraintDetailsRepresentation},
	}
	CoreDedicatedVmHostPlacementConstraintDetailsRepresentation = map[string]interface{}{
		"type": acctest.Representation{RepType: acctest.Required, Create: `COMPUTE_BARE_METAL_HOST`},
		// For this variable to work, make sure to assign a valid compute_bare_metal_host_id value from the region you
		// are testing in, to the TF_var_compute_bare_metal_host_id in your ~/.zshrc (or the likes). For example:
		// export TF_VAR_compute_bare_metal_host_id=valid_compute_bare_metal_host_id
		// Refer
		"compute_bare_metal_host_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compute_bare_metal_host_id}`},
	}

	CoreDedicatedVmHostResourceDependencies = AvailabilityDomainConfig
)

// issue-routing-tag: core/default
func TestCoreDedicatedVmHostResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreDedicatedVmHostResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_dedicated_vm_host.test_dedicated_vm_host"
	datasourceName := "data.oci_core_dedicated_vm_hosts.test_dedicated_vm_hosts"
	singularDatasourceName := "data.oci_core_dedicated_vm_host.test_dedicated_vm_host"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CoreDedicatedVmHostResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_core_dedicated_vm_host", "test_dedicated_vm_host", acctest.Optional, acctest.Create,
			CoreDedicatedVmHostRepresentation), "core", "dedicatedVmHost", t)

	acctest.ResourceTest(t, testAccCheckCoreDedicatedVmHostDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + CoreDedicatedVmHostResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_dedicated_vm_host", "test_dedicated_vm_host",
					acctest.Required, acctest.Create, CoreDedicatedVmHostRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "dedicated_vm_host_shape", dvhShape),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttr(resourceName, "capacity_bins.#", capacityBinsSizeForDvhShape),

				// This property will only be set if the testing tenancy mimics a dedicated capacity customer
				// Meaning that this tenancy has a dedicated pool for its DVH needs. Update the value of
				// this property accordingly based on your test requirements & inputs.
				resource.TestCheckResourceAttrSet(resourceName, "compute_bare_metal_host_id"),

				// This property will only have a non-zero value during create if the user passes it in the input
				resource.TestCheckResourceAttr(resourceName, "placement_constraint_details.#", "0"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + CoreDedicatedVmHostResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + CoreDedicatedVmHostResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_dedicated_vm_host", "test_dedicated_vm_host",
					acctest.Optional, acctest.Create, CoreDedicatedVmHostRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "dedicated_vm_host_shape", dvhShape),
				resource.TestCheckResourceAttr(resourceName, "display_name", displayNameForCreate),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "remaining_ocpus"),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "total_ocpus"),
				resource.TestCheckResourceAttr(resourceName, "capacity_bins.#", capacityBinsSizeForDvhShape),

				// This property will only be set if the testing tenancy mimics a dedicated capacity customer
				// Meaning that this tenancy has a dedicated pool for its DVH needs. Update the value of
				// this property accordingly based on your test requirements & inputs.
				resource.TestCheckResourceAttrSet(resourceName, "compute_bare_metal_host_id"),

				// This property will only have a non-zero value during create if the user passes it in the input
				// Since we're passing this for Optional tests, this field should be populated
				resource.TestCheckResourceAttr(resourceName, "placement_constraint_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "placement_constraint_details.0.compute_bare_metal_host_id"),
				resource.TestCheckResourceAttr(resourceName, "placement_constraint_details.0.type", "COMPUTE_BARE_METAL_HOST"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + CoreDedicatedVmHostResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_dedicated_vm_host", "test_dedicated_vm_host", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(CoreDedicatedVmHostRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "dedicated_vm_host_shape", dvhShape),
				resource.TestCheckResourceAttr(resourceName, "display_name", displayNameForCreate),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "remaining_ocpus"),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "total_ocpus"),
				resource.TestCheckResourceAttr(resourceName, "capacity_bins.#", capacityBinsSizeForDvhShape),

				// This property will only be set if the testing tenancy mimics a dedicated capacity customer
				// Meaning that this tenancy has a dedicated pool for its DVH needs. Update the value of
				// this property accordingly based on your test requirements & inputs.
				resource.TestCheckResourceAttrSet(resourceName, "compute_bare_metal_host_id"),

				// This property will only have a non-zero value during create if the user passes it in the input
				// Since we're passing this for Optional tests, this field should be populated
				resource.TestCheckResourceAttr(resourceName, "placement_constraint_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "placement_constraint_details.0.compute_bare_metal_host_id"),
				resource.TestCheckResourceAttr(resourceName, "placement_constraint_details.0.type", "COMPUTE_BARE_METAL_HOST"),

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
			Config: config + compartmentIdVariableStr + CoreDedicatedVmHostResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_dedicated_vm_host", "test_dedicated_vm_host", acctest.Optional, acctest.Update,
					CoreDedicatedVmHostRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "dedicated_vm_host_shape", dvhShape),
				resource.TestCheckResourceAttr(resourceName, "display_name", displayNameForUpdate),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "remaining_ocpus"),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "total_ocpus"),
				resource.TestCheckResourceAttr(resourceName, "capacity_bins.#", capacityBinsSizeForDvhShape),

				// This property will only be set if the testing tenancy mimics a dedicated capacity customer
				// Meaning that this tenancy has a dedicated pool for its DVH needs. Update the value of
				// this property accordingly based on your test requirements & inputs.
				resource.TestCheckResourceAttrSet(resourceName, "compute_bare_metal_host_id"),

				// This property will only have a non-zero value during create if the user passes it in the input
				// Since we're passing this for Optional tests, this field should be populated
				resource.TestCheckResourceAttr(resourceName, "placement_constraint_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "placement_constraint_details.0.compute_bare_metal_host_id"),
				resource.TestCheckResourceAttr(resourceName, "placement_constraint_details.0.type", "COMPUTE_BARE_METAL_HOST"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource (list DVH call)
		{
			Config: config + acctest.GenerateDataSourceFromRepresentationMap("oci_core_dedicated_vm_hosts", "test_dedicated_vm_hosts",
				acctest.Optional, acctest.Update, CoreDedicatedVmHostDataSourceRepresentation) + compartmentIdVariableStr + CoreDedicatedVmHostResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_dedicated_vm_host", "test_dedicated_vm_host", acctest.Optional, acctest.Update,
					CoreDedicatedVmHostRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "dedicated_vm_hosts.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_hosts.0.availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "dedicated_vm_hosts.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "dedicated_vm_hosts.0.dedicated_vm_host_shape", dvhShape),
				resource.TestCheckResourceAttr(datasourceName, "dedicated_vm_hosts.0.display_name", displayNameForUpdate),
				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_hosts.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_hosts.0.remaining_memory_in_gbs"),
				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_hosts.0.remaining_ocpus"),
				resource.TestCheckResourceAttr(datasourceName, "dedicated_vm_hosts.0.state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_hosts.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_hosts.0.total_memory_in_gbs"),
				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_hosts.0.total_ocpus"),
			),
		},
		// verify singular datasource (get DVH call)
		{
			Config: config + acctest.GenerateDataSourceFromRepresentationMap("oci_core_dedicated_vm_host", "test_dedicated_vm_host", acctest.Required, acctest.Create,
				CoreCoreDedicatedVmHostSingularDataSourceRepresentation) + compartmentIdVariableStr + CoreDedicatedVmHostResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "dedicated_vm_host_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "capacity_bins.#", capacityBinsSizeForDvhShape),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "dedicated_vm_host_shape", dvhShape),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", displayNameForUpdate),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "remaining_memory_in_gbs"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "remaining_ocpus"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "total_memory_in_gbs"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "total_ocpus"),

				// This property will only be set if the testing tenancy mimics a dedicated capacity customer
				// Meaning that this tenancy has a dedicated pool for its DVH needs. Update the value of
				// this property accordingly based on your test requirements & inputs.
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compute_bare_metal_host_id"),

				// This property will only have a non-zero value during create if the user passes it in the input
				// Since we're passing this for Optional tests, this field should be populated
				resource.TestCheckResourceAttr(resourceName, "placement_constraint_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "placement_constraint_details.0.compute_bare_metal_host_id"),
				resource.TestCheckResourceAttr(resourceName, "placement_constraint_details.0.type", "COMPUTE_BARE_METAL_HOST"),
			),
		},
		// verify resource import
		{
			Config:                  config + CoreDedicatedVmHostRequiredOnlyResource + compartmentIdVariableStr,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckCoreDedicatedVmHostDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ComputeClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_dedicated_vm_host" {
			noResourceFound = false
			request := oci_core.GetDedicatedVmHostRequest{}

			tmp := rs.Primary.ID
			request.DedicatedVmHostId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")

			response, err := client.GetDedicatedVmHost(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.DedicatedVmHostLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("CoreDedicatedVmHost") {
		resource.AddTestSweepers("CoreDedicatedVmHost", &resource.Sweeper{
			Name:         "CoreDedicatedVmHost",
			Dependencies: acctest.DependencyGraph["dedicatedVmHost"],
			F:            sweepCoreDedicatedVmHostResource,
		})
	}
}

func sweepCoreDedicatedVmHostResource(compartment string) error {
	computeClient := acctest.GetTestClients(&schema.ResourceData{}).ComputeClient()
	dedicatedVmHostIds, err := getCoreDedicatedVmHostIds(compartment)
	if err != nil {
		return err
	}
	for _, dedicatedVmHostId := range dedicatedVmHostIds {
		if ok := acctest.SweeperDefaultResourceId[dedicatedVmHostId]; !ok {
			deleteDedicatedVmHostRequest := oci_core.DeleteDedicatedVmHostRequest{}

			deleteDedicatedVmHostRequest.DedicatedVmHostId = &dedicatedVmHostId

			deleteDedicatedVmHostRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")
			_, error := computeClient.DeleteDedicatedVmHost(context.Background(), deleteDedicatedVmHostRequest)
			if error != nil {
				fmt.Printf("Error deleting DedicatedVmHost %s %s, It is possible that the resource is already deleted. Please verify manually \n", dedicatedVmHostId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &dedicatedVmHostId, CoreDedicatedVmHostSweepWaitCondition, time.Duration(3*time.Minute),
				CoreDedicatedVmHostSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getCoreDedicatedVmHostIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DedicatedVmHostId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	computeClient := acctest.GetTestClients(&schema.ResourceData{}).ComputeClient()

	listDedicatedVmHostsRequest := oci_core.ListDedicatedVmHostsRequest{}
	listDedicatedVmHostsRequest.CompartmentId = &compartmentId
	listDedicatedVmHostsRequest.LifecycleState = oci_core.ListDedicatedVmHostsLifecycleStateActive
	listDedicatedVmHostsResponse, err := computeClient.ListDedicatedVmHosts(context.Background(), listDedicatedVmHostsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting DedicatedVmHost list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, dedicatedVmHost := range listDedicatedVmHostsResponse.Items {
		id := *dedicatedVmHost.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DedicatedVmHostId", id)
	}
	return resourceIds, nil
}

func CoreDedicatedVmHostSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if dedicatedVmHostResponse, ok := response.Response.(oci_core.GetDedicatedVmHostResponse); ok {
		return dedicatedVmHostResponse.LifecycleState != oci_core.DedicatedVmHostLifecycleStateDeleted
	}
	return false
}

func CoreDedicatedVmHostSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ComputeClient().GetDedicatedVmHost(context.Background(), oci_core.GetDedicatedVmHostRequest{
		DedicatedVmHostId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
