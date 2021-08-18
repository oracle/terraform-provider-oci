// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v46/common"
	oci_core "github.com/oracle/oci-go-sdk/v46/core"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	DedicatedVmHostRequiredOnlyResource = DedicatedVmHostResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_dedicated_vm_host", "test_dedicated_vm_host", Required, Create, dedicatedVmHostRepresentation)

	DedicatedVmHostResourceConfig = DedicatedVmHostResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_dedicated_vm_host", "test_dedicated_vm_host", Optional, Update, dedicatedVmHostRepresentation)

	dedicatedVmHostSingularDataSourceRepresentation = map[string]interface{}{
		"dedicated_vm_host_id": Representation{repType: Required, create: `${oci_core_dedicated_vm_host.test_dedicated_vm_host.id}`},
	}

	dedicatedVmHostDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      Representation{repType: Required, create: `${var.compartment_id}`},
		"availability_domain": Representation{repType: Optional, create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"display_name":        Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"instance_shape_name": Representation{repType: Optional, create: `VM.Standard2.1`},
		"remaining_memory_in_gbs_greater_than_or_equal_to": Representation{repType: Optional, create: `15.0`},
		"remaining_ocpus_greater_than_or_equal_to":         Representation{repType: Optional, create: `1.0`},
		"state":  Representation{repType: Optional, create: `ACTIVE`},
		"filter": RepresentationGroup{Required, dedicatedVmHostDataSourceFilterRepresentation}}
	dedicatedVmHostDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_core_dedicated_vm_host.test_dedicated_vm_host.id}`}},
	}

	dedicatedVmHostRepresentation = map[string]interface{}{
		"availability_domain":     Representation{repType: Required, create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":          Representation{repType: Required, create: `${var.compartment_id}`},
		"dedicated_vm_host_shape": Representation{repType: Required, create: `DVH.Standard2.52`},
		"defined_tags":            Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":            Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"fault_domain":            Representation{repType: Optional, create: `FAULT-DOMAIN-3`},
		"freeform_tags":           Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
	}

	DedicatedVmHostResourceDependencies = AvailabilityDomainConfig +
		DefinedTagsDependencies
)

// issue-routing-tag: core/default
func TestCoreDedicatedVmHostResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreDedicatedVmHostResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_dedicated_vm_host.test_dedicated_vm_host"
	datasourceName := "data.oci_core_dedicated_vm_hosts.test_dedicated_vm_hosts"
	singularDatasourceName := "data.oci_core_dedicated_vm_host.test_dedicated_vm_host"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+DedicatedVmHostResourceDependencies+
		generateResourceFromRepresentationMap("oci_core_dedicated_vm_host", "test_dedicated_vm_host", Optional, Create, dedicatedVmHostRepresentation), "core", "dedicatedVmHost", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreDedicatedVmHostDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + DedicatedVmHostResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_dedicated_vm_host", "test_dedicated_vm_host", Required, Create, dedicatedVmHostRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "dedicated_vm_host_shape", "DVH.Standard2.52"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + DedicatedVmHostResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + DedicatedVmHostResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_dedicated_vm_host", "test_dedicated_vm_host", Optional, Create, dedicatedVmHostRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "dedicated_vm_host_shape", "DVH.Standard2.52"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "fault_domain", "FAULT-DOMAIN-3"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "remaining_ocpus"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "total_ocpus"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},

			// verify update to the compartment (the compartment will be switched back in the next step)
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DedicatedVmHostResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_dedicated_vm_host", "test_dedicated_vm_host", Optional, Create,
						representationCopyWithNewProperties(dedicatedVmHostRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "dedicated_vm_host_shape", "DVH.Standard2.52"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "fault_domain", "FAULT-DOMAIN-3"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "remaining_ocpus"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "total_ocpus"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + DedicatedVmHostResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_dedicated_vm_host", "test_dedicated_vm_host", Optional, Update, dedicatedVmHostRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "dedicated_vm_host_shape", "DVH.Standard2.52"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "fault_domain", "FAULT-DOMAIN-3"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "remaining_ocpus"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "total_ocpus"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateDataSourceFromRepresentationMap("oci_core_dedicated_vm_hosts", "test_dedicated_vm_hosts", Optional, Update, dedicatedVmHostDataSourceRepresentation) +
					compartmentIdVariableStr + DedicatedVmHostResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_dedicated_vm_host", "test_dedicated_vm_host", Optional, Update, dedicatedVmHostRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "instance_shape_name", "VM.Standard2.1"),
					resource.TestCheckResourceAttr(datasourceName, "remaining_memory_in_gbs_greater_than_or_equal_to", "15"),
					resource.TestCheckResourceAttr(datasourceName, "remaining_ocpus_greater_than_or_equal_to", "1"),
					resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

					resource.TestCheckResourceAttr(datasourceName, "dedicated_vm_hosts.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_hosts.0.availability_domain"),
					resource.TestCheckResourceAttr(datasourceName, "dedicated_vm_hosts.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "dedicated_vm_hosts.0.dedicated_vm_host_shape", "DVH.Standard2.52"),
					resource.TestCheckResourceAttr(datasourceName, "dedicated_vm_hosts.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "dedicated_vm_hosts.0.fault_domain", "FAULT-DOMAIN-3"),
					resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_hosts.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_hosts.0.remaining_memory_in_gbs"),
					resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_hosts.0.remaining_ocpus"),
					resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_hosts.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_hosts.0.time_created"),
					resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_hosts.0.total_memory_in_gbs"),
					resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_hosts.0.total_ocpus"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_core_dedicated_vm_host", "test_dedicated_vm_host", Required, Create, dedicatedVmHostSingularDataSourceRepresentation) +
					compartmentIdVariableStr + DedicatedVmHostResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "dedicated_vm_host_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "dedicated_vm_host_shape", "DVH.Standard2.52"),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "fault_domain", "FAULT-DOMAIN-3"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "remaining_memory_in_gbs"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "remaining_ocpus"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "total_memory_in_gbs"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "total_ocpus"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + DedicatedVmHostResourceConfig,
			},
			// verify resource import
			{
				Config:                  config,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}

func testAccCheckCoreDedicatedVmHostDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).computeClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_dedicated_vm_host" {
			noResourceFound = false
			request := oci_core.GetDedicatedVmHostRequest{}

			tmp := rs.Primary.ID
			request.DedicatedVmHostId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "core")

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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("CoreDedicatedVmHost") {
		resource.AddTestSweepers("CoreDedicatedVmHost", &resource.Sweeper{
			Name:         "CoreDedicatedVmHost",
			Dependencies: DependencyGraph["dedicatedVmHost"],
			F:            sweepCoreDedicatedVmHostResource,
		})
	}
}

func sweepCoreDedicatedVmHostResource(compartment string) error {
	computeClient := GetTestClients(&schema.ResourceData{}).computeClient()
	dedicatedVmHostIds, err := getDedicatedVmHostIds(compartment)
	if err != nil {
		return err
	}
	for _, dedicatedVmHostId := range dedicatedVmHostIds {
		if ok := SweeperDefaultResourceId[dedicatedVmHostId]; !ok {
			deleteDedicatedVmHostRequest := oci_core.DeleteDedicatedVmHostRequest{}

			deleteDedicatedVmHostRequest.DedicatedVmHostId = &dedicatedVmHostId

			deleteDedicatedVmHostRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "core")
			_, error := computeClient.DeleteDedicatedVmHost(context.Background(), deleteDedicatedVmHostRequest)
			if error != nil {
				fmt.Printf("Error deleting DedicatedVmHost %s %s, It is possible that the resource is already deleted. Please verify manually \n", dedicatedVmHostId, error)
				continue
			}
			waitTillCondition(testAccProvider, &dedicatedVmHostId, dedicatedVmHostSweepWaitCondition, time.Duration(3*time.Minute),
				dedicatedVmHostSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getDedicatedVmHostIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "DedicatedVmHostId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	computeClient := GetTestClients(&schema.ResourceData{}).computeClient()

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
		addResourceIdToSweeperResourceIdMap(compartmentId, "DedicatedVmHostId", id)
	}
	return resourceIds, nil
}

func dedicatedVmHostSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if dedicatedVmHostResponse, ok := response.Response.(oci_core.GetDedicatedVmHostResponse); ok {
		return dedicatedVmHostResponse.LifecycleState != oci_core.DedicatedVmHostLifecycleStateDeleted
	}
	return false
}

func dedicatedVmHostSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.computeClient().GetDedicatedVmHost(context.Background(), oci_core.GetDedicatedVmHostRequest{
		DedicatedVmHostId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
