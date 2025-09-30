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
	displayNameForCreate = `tersi-dvh-display-name-create`
	displayNameForUpdate = `tersi-dvh-display-name-update`

	// Configurations for E4 shape
	dvhShape                    = `DVH.Standard.E4.128`
	capacityConfig              = `standard_e4_flex`
	capacityBinsSizeForDvhShape = `2`
	isMemoryEncryptionEnabled   = `true`

	CoreDedicatedVmHostRequiredOnlyResource = CoreDedicatedVmHostResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_dedicated_vm_host", "test_dedicated_vm_host",
			acctest.Required, acctest.Create, CoreDedicatedVmHostRepresentation)

	CoreDedicatedVmHostResourceConfig = CoreDedicatedVmHostResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_dedicated_vm_host", "test_dedicated_vm_host",
			acctest.Optional, acctest.Update, CoreDedicatedVmHostRepresentation)

	// This is the path parameter for the Get DVH API call
	CoreCoreDedicatedVmHostSingularDataSourceRepresentation = map[string]interface{}{
		"dedicated_vm_host_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_dedicated_vm_host.test_dedicated_vm_host.id}`},
	}

	// These are the query parameters for the List DVH API call
	CoreDedicatedVmHostDataSourceRepresentation = map[string]interface{}{
		"compartment_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"availability_domain":          acctest.Representation{RepType: acctest.Required, Create: `${var.availability_domain}`},
		"display_name":                 acctest.Representation{RepType: acctest.Optional, Create: displayNameForCreate, Update: displayNameForUpdate},
		"is_memory_encryption_enabled": acctest.Representation{RepType: acctest.Optional, Create: isMemoryEncryptionEnabled},
		"state":                        acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":                       acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreDedicatedVmHostDataSourceFilterRepresentation},

		// When using instanceShapeName query parameter, we also need to provide the remaining CPU query
		// parameter if using FLEX shapes otherwise List DVH API call (Step 6) fails
		// "instance_shape_name":                              acctest.Representation{RepType: acctest.Optional, Create: vmShape},
		// "remaining_memory_in_gbs_greater_than_or_equal_to": acctest.Representation{RepType: acctest.Optional, Create: `1.0`},
		// "remaining_ocpus_greater_than_or_equal_to":         acctest.Representation{RepType: acctest.Optional, Create: `1.0`},
	}

	CoreDedicatedVmHostDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_dedicated_vm_host.test_dedicated_vm_host.id}`}},
	}

	CoreDedicatedVmHostRepresentation = map[string]interface{}{
		"availability_domain":          acctest.Representation{RepType: acctest.Required, Create: `${var.availability_domain}`},
		"compartment_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"dedicated_vm_host_shape":      acctest.Representation{RepType: acctest.Required, Create: dvhShape},
		"display_name":                 acctest.Representation{RepType: acctest.Required, Create: displayNameForCreate, Update: displayNameForUpdate},
		"capacity_config":              acctest.Representation{RepType: acctest.Optional, Create: capacityConfig},
		"is_memory_encryption_enabled": acctest.Representation{RepType: acctest.Optional, Create: isMemoryEncryptionEnabled},
		"freeform_tags":                acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"FreeFormState": "Created"}, Update: map[string]string{"FreeFormState": "Updated"}},
	}

	CoreDedicatedVmHostResourceDependencies = AvailabilityDomainConfig
)

// issue-routing-tag: core/default
func TestCoreDedicatedVmHostResource_basic(t *testing.T) {
	if err := httpreplay.SetScenario("TestCoreDedicatedVmHostResource_basic"); err != nil {
		fmt.Printf("error occurred in httpreplay.SetScenario, %s", err)
		return
	}
	defer func() {
		if err := httpreplay.SaveScenario(); err != nil {
			fmt.Printf("error occurred in httpreplay.SaveScenario, %s", err)
			return
		}
	}()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	availabilityDomain := utils.GetEnvSettingWithBlankDefault("availability_domain")
	availabilityDomainVariableStr := fmt.Sprintf("variable \"availability_domain\" { default = \"%s\" }\n", availabilityDomain)

	resourceName := "oci_core_dedicated_vm_host.test_dedicated_vm_host"
	datasourceName := "data.oci_core_dedicated_vm_hosts.test_dedicated_vm_hosts"
	singularDatasourceName := "data.oci_core_dedicated_vm_host.test_dedicated_vm_host"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties.
	// This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+availabilityDomainVariableStr+CoreDedicatedVmHostResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_core_dedicated_vm_host", "test_dedicated_vm_host",
			acctest.Optional, acctest.Create, CoreDedicatedVmHostRepresentation), "core", "dedicatedVmHost", t)

	acctest.ResourceTest(t, testAccCheckCoreDedicatedVmHostDestroy, []resource.TestStep{
		// Step 1 - verify Create
		{
			Config: config + compartmentIdVariableStr + availabilityDomainVariableStr + CoreDedicatedVmHostResourceDependencies +
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

				// Verify default capacity config is selected when no value is passed, for the shape in this
				// example, there is only one capacity config so ensure that is picked even when not passed
				resource.TestCheckResourceAttr(resourceName, "capacity_config", capacityConfig),
				// This property will be true if you're creating a Confidential DVH, false otherwise
				resource.TestCheckResourceAttr(resourceName, "is_memory_encryption_enabled", "false"),

				func(s *terraform.State) (err error) {
					if resId, err = acctest.FromInstanceState(s, resourceName, "id"); err != nil {
						return fmt.Errorf("[TERSI-LOG] Step 1: resource with id=%s could not be created. Error: %s\n", resId, err)
					}
					fmt.Printf("[TERSI-LOG] Step 1: Resource created with id: %s\n", resId)
					return err
				},
			),
		},

		// Step 2 - delete before next Create
		{
			Config: config + compartmentIdVariableStr + availabilityDomainVariableStr + CoreDedicatedVmHostResourceDependencies,
		},

		// Step 3 - verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + availabilityDomainVariableStr + CoreDedicatedVmHostResourceDependencies +
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
				resource.TestCheckResourceAttr(resourceName, "placement_constraint_details.#", "0"),

				resource.TestCheckResourceAttr(resourceName, "capacity_config", capacityConfig),
				// This property will be true if you're creating a Confidential DVH, false otherwise
				resource.TestCheckResourceAttr(resourceName, "is_memory_encryption_enabled", isMemoryEncryptionEnabled),

				func(s *terraform.State) (err error) {
					if resId, err = acctest.FromInstanceState(s, resourceName, "id"); err != nil {
						return fmt.Errorf("[TERSI-LOG] Step 3: resource with id=%s could not be created. Error: %s\n", resId, err)
					}
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					fmt.Printf("[TERSI-LOG] Step 3: Resource created with id: %s\n", resId)
					return err
				},
			),
		},

		// Step 4 - verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + availabilityDomainVariableStr + compartmentIdUVariableStr + CoreDedicatedVmHostResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_dedicated_vm_host", "test_dedicated_vm_host",
					acctest.Optional, acctest.Create,
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
				resource.TestCheckResourceAttr(resourceName, "placement_constraint_details.#", "0"),

				resource.TestCheckResourceAttr(resourceName, "capacity_config", capacityConfig),
				// This property will be true if you're creating a Confidential DVH, false otherwise
				resource.TestCheckResourceAttr(resourceName, "is_memory_encryption_enabled", isMemoryEncryptionEnabled),

				func(s *terraform.State) (err error) {
					if resId2, err = acctest.FromInstanceState(s, resourceName, "id"); err != nil {
						return fmt.Errorf("[TERSI-LOG] Step 4: resource with id=%s could not be created. Error: %s\n", resId2, err)
					}
					if resId != resId2 {
						return fmt.Errorf("[TERSI-LOG] Step 4: resource recreated when it was supposed to be updated, resource values - "+
							"resId: %s, resId2: %s", resId, resId2)
					}
					fmt.Printf("[TERSI-LOG] Step 4: Resource updated with id: %s\n", resId2)
					return err
				},
			),
		},

		// Step 5 - verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + availabilityDomainVariableStr + CoreDedicatedVmHostResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_dedicated_vm_host", "test_dedicated_vm_host",
					acctest.Optional, acctest.Update, CoreDedicatedVmHostRepresentation),
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

				resource.TestCheckResourceAttr(resourceName, "capacity_config", capacityConfig),
				// This property will be true if you're creating a Confidential DVH, false otherwise
				resource.TestCheckResourceAttr(resourceName, "is_memory_encryption_enabled", isMemoryEncryptionEnabled),

				func(s *terraform.State) (err error) {
					if resId2, err = acctest.FromInstanceState(s, resourceName, "id"); err != nil {
						return fmt.Errorf("[TERSI-LOG] Step 5: resource with id=%s could not be created. Error: %s\n", resId2, err)
					}
					if resId != resId2 {
						return fmt.Errorf("[TERSI-LOG] Step 5: resource recreated when it was supposed to be updated, resource values - "+
							"resId: %s, resId2: %s", resId, resId2)
					}
					fmt.Printf("[TERSI-LOG] Step 5: Resource updated with id: %s\n", resId2)
					return err
				},
			),
		},

		// Step 6 - verify datasource (List DVH call)
		{
			Config: config + acctest.GenerateDataSourceFromRepresentationMap("oci_core_dedicated_vm_hosts", "test_dedicated_vm_hosts",
				acctest.Optional, acctest.Update, CoreDedicatedVmHostDataSourceRepresentation) + compartmentIdVariableStr + availabilityDomainVariableStr + CoreDedicatedVmHostResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_dedicated_vm_host", "test_dedicated_vm_host", acctest.Optional, acctest.Update,
					CoreDedicatedVmHostRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "dedicated_vm_hosts.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_hosts.0.availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "dedicated_vm_hosts.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "dedicated_vm_hosts.0.dedicated_vm_host_shape", dvhShape),
				resource.TestCheckResourceAttr(datasourceName, "dedicated_vm_hosts.0.display_name", displayNameForUpdate),
				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_hosts.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "dedicated_vm_hosts.0.is_memory_encryption_enabled", isMemoryEncryptionEnabled),
				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_hosts.0.remaining_memory_in_gbs"),
				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_hosts.0.remaining_ocpus"),
				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_hosts.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_hosts.0.total_memory_in_gbs"),
				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_hosts.0.total_ocpus"),
			),
		},

		// Step 7 - verify singular datasource (Get DVH call)
		{
			Config: config + acctest.GenerateDataSourceFromRepresentationMap("oci_core_dedicated_vm_host", "test_dedicated_vm_host", acctest.Required, acctest.Create,
				CoreCoreDedicatedVmHostSingularDataSourceRepresentation) + compartmentIdVariableStr + availabilityDomainVariableStr + CoreDedicatedVmHostResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "dedicated_vm_host_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "capacity_config", capacityConfig),
				resource.TestCheckResourceAttr(singularDatasourceName, "capacity_bins.#", capacityBinsSizeForDvhShape),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "dedicated_vm_host_shape", dvhShape),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", displayNameForUpdate),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_memory_encryption_enabled", isMemoryEncryptionEnabled),
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
				resource.TestCheckResourceAttr(resourceName, "placement_constraint_details.#", "0"),

				resource.TestCheckResourceAttr(resourceName, "capacity_config", capacityConfig),
				// This property will be true if you're creating a Confidential DVH, false otherwise
				resource.TestCheckResourceAttr(resourceName, "is_memory_encryption_enabled", isMemoryEncryptionEnabled),
			),
		},

		// Step 8 - verify resource import
		{
			Config:                  config + CoreDedicatedVmHostRequiredOnlyResource + compartmentIdVariableStr + availabilityDomainVariableStr,
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
					// resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
				// resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
			}

			// Verify that exception is for '404 not found'.
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
			_, err2 := computeClient.DeleteDedicatedVmHost(context.Background(), deleteDedicatedVmHostRequest)
			if err2 != nil {
				fmt.Printf("Error deleting DedicatedVmHost %s %s, It is possible that the resource is already deleted. Please verify manually \n", dedicatedVmHostId, err2)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &dedicatedVmHostId, CoreDedicatedVmHostSweepWaitCondition, 3*time.Minute,
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
