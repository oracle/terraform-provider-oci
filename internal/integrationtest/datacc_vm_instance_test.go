// Copyright (c) 2017, 2026, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_datacc "github.com/oracle/oci-go-sdk/v65/datacc"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/globalvar"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

const (
	vmInstanceNumSshPublicKeys = 1
	vmInstanceNumFreeformTags  = 1
	vmInstanceNumStates        = 1
	vmInstanceNumCollections   = 1
)

var (
	DataccVmInstanceDependencyVariableConfig = GenerateTFVariableStrings("vm_cluster_network")

	DataccVmInstanceRequiredOnlyResource = DataccVmInstanceDependencyVariableConfig +
		DataccVmInstanceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_datacc_vm_instance", "test_vm_instance", acctest.Required, acctest.Create, DataccVmInstanceRepresentation)

	DataccVmInstanceResourceConfig = DataccVmInstanceDependencyVariableConfig +
		DataccVmInstanceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_datacc_vm_instance", "test_vm_instance", acctest.Optional, acctest.Update, DataccVmInstanceRepresentation)

	DataccVmInstanceSingularDataSourceRepresentation = map[string]interface{}{
		"vm_instance_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_datacc_vm_instance.test_vm_instance.id}`},
	}

	DataccVmInstanceDataSourceRepresentation = map[string]interface{}{
		"compartment_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.vm_instance_compartment_id}`},
		"display_name":      acctest.Representation{RepType: acctest.Optional, Create: `${var.vm_instance_display_name}`, Update: `${var.vm_instance_display_name_for_update}`},
		"infrastructure_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.vm_instance_infrastructure_id}`},
		"state":             acctest.Representation{RepType: acctest.Optional, Create: []string{`ACTIVE`}, Update: []string{`ACTIVE`}},
		"filter":            acctest.RepresentationGroup{RepType: acctest.Required, Group: DataccVmInstanceDataSourceFilterRepresentation}}
	DataccVmInstanceDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_datacc_vm_instance.test_vm_instance.id}`}},
	}

	DataccVmInstanceRepresentation = map[string]interface{}{
		"compartment_id":           acctest.Representation{RepType: acctest.Required, Create: `${var.vm_instance_compartment_id}`},
		"cpus_enabled":             acctest.Representation{RepType: acctest.Required, Create: `${var.vm_instance_cpus_enabled}`},
		"infrastructure_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.vm_instance_infrastructure_id}`},
		"ssh_public_keys":          acctest.Representation{RepType: acctest.Required, Create: []string{`${var.vm_instance_ssh_public_key_0}`}},
		"boot_storage_size_in_gbs": acctest.Representation{RepType: acctest.Required, Create: `${var.vm_instance_boot_storage_size_in_gbs}`},
		"data_storage_size_in_gb":  acctest.Representation{RepType: acctest.Required, Create: `${var.vm_instance_data_storage_size_in_gbs}`, Update: `${var.vm_instance_data_storage_size_in_gbs_for_update}`},
		"memory_size_in_gbs":       acctest.Representation{RepType: acctest.Required, Create: `${var.vm_instance_memory_size_in_gbs}`},
		"description":              acctest.Representation{RepType: acctest.Optional, Create: `${var.vm_instance_description}`, Update: `${var.vm_instance_description_for_update}`},
		"display_name":             acctest.Representation{RepType: acctest.Optional, Create: `${var.vm_instance_display_name}`, Update: `${var.vm_instance_display_name_for_update}`},
		"freeform_tags":            acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"vm_network_id":            acctest.Representation{RepType: acctest.Required, Create: `${oci_datacc_vm_cluster_network.test_vm_cluster_network.id}`},
	}

	DataccVmInstanceResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_datacc_vm_cluster_network", "test_vm_cluster_network", acctest.Optional, acctest.Create, DataccVmClusterNetworkRepresentation)
)

// issue-routing-tag: datacc/default
func TestDataccVmInstanceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataccVmInstanceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	// override terraform-federation-test profile with our own user profile
	if overrideProfile := os.Getenv("datacc_custom_config_file_profile_override"); overrideProfile != "" {
		t.Setenv(globalvar.TfEnvPrefix+globalvar.ConfigFileProfileAttrName, overrideProfile)
		t.Setenv(globalvar.TfEnvPrefix+globalvar.AuthAttrName, "")
		t.Setenv(globalvar.AuthAttrName, globalvar.AuthSecurityToken)
	}

	const testResourceType = "vm_instance"
	tfVariableStr := GenerateTFVariableStrings(testResourceType)
	getTFVar := func(variableName string) string {
		return os.Getenv(globalvar.TfEnvPrefix + testResourceType + "_" + variableName)
	}

	compartmentId := getTFVar("compartment_id")
	compartmentIdU := getTFVar("compartment_id_for_update")
	infrastructureId := getTFVar("infrastructure_id")
	displayName := getTFVar("display_name")
	displayNameU := getTFVar("display_name_for_update")
	description := getTFVar("description")
	descriptionU := getTFVar("description_for_update")
	cpusEnabled := getTFVar("cpus_enabled")
	memorySizeInGbs := getTFVar("memory_size_in_gbs")
	bootStorageSizeInGbs := getTFVar("boot_storage_size_in_gbs")
	dataStorageSizeInGbs := getTFVar("data_storage_size_in_gbs")
	dataStorageSizeInGbsU := getTFVar("data_storage_size_in_gbs_for_update")

	resourceName := "oci_datacc_vm_instance.test_vm_instance"
	datasourceName := "data.oci_datacc_vm_instances.test_vm_instances"
	singularDatasourceName := "data.oci_datacc_vm_instance.test_vm_instance"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+tfVariableStr+DataccVmInstanceDependencyVariableConfig+DataccVmInstanceResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_datacc_vm_instance", "test_vm_instance", acctest.Optional, acctest.Create, DataccVmInstanceRepresentation), "datacc", "vmInstance", t)

	acctest.ResourceTest(t, testAccCheckDataccVmInstanceDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + tfVariableStr + DataccVmInstanceDependencyVariableConfig + DataccVmInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datacc_vm_instance", "test_vm_instance", acctest.Required, acctest.Create, DataccVmInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpus_enabled", cpusEnabled),
				resource.TestCheckResourceAttr(resourceName, "infrastructure_id", infrastructureId),
				resource.TestCheckResourceAttr(resourceName, "ssh_public_keys.#", strconv.Itoa(vmInstanceNumSshPublicKeys)),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + tfVariableStr + DataccVmInstanceDependencyVariableConfig + DataccVmInstanceResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + tfVariableStr + DataccVmInstanceDependencyVariableConfig + DataccVmInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datacc_vm_instance", "test_vm_instance", acctest.Optional, acctest.Create, DataccVmInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "boot_storage_size_in_gbs", bootStorageSizeInGbs),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpus_enabled", cpusEnabled),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_gb", dataStorageSizeInGbs),
				resource.TestCheckResourceAttr(resourceName, "description", description),
				resource.TestCheckResourceAttr(resourceName, "display_name", displayName),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", strconv.Itoa(vmInstanceNumFreeformTags)),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "infrastructure_id", infrastructureId),
				resource.TestCheckResourceAttr(resourceName, "memory_size_in_gbs", memorySizeInGbs),
				resource.TestCheckResourceAttr(resourceName, "ssh_public_keys.#", strconv.Itoa(vmInstanceNumSshPublicKeys)),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttrPair(resourceName, "vm_network_id", "oci_datacc_vm_cluster_network.test_vm_cluster_network", "id"),

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
			Config: config + tfVariableStr + DataccVmInstanceDependencyVariableConfig + DataccVmInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datacc_vm_instance", "test_vm_instance", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DataccVmInstanceRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.vm_instance_compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "boot_storage_size_in_gbs", bootStorageSizeInGbs),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "cpus_enabled", cpusEnabled),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_gb", dataStorageSizeInGbs),
				resource.TestCheckResourceAttr(resourceName, "description", description),
				resource.TestCheckResourceAttr(resourceName, "display_name", displayName),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", strconv.Itoa(vmInstanceNumFreeformTags)),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "infrastructure_id", infrastructureId),
				resource.TestCheckResourceAttr(resourceName, "memory_size_in_gbs", memorySizeInGbs),
				resource.TestCheckResourceAttr(resourceName, "ssh_public_keys.#", strconv.Itoa(vmInstanceNumSshPublicKeys)),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttrPair(resourceName, "vm_network_id", "oci_datacc_vm_cluster_network.test_vm_cluster_network", "id"),

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
			Config: config + tfVariableStr + DataccVmInstanceDependencyVariableConfig + DataccVmInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datacc_vm_instance", "test_vm_instance", acctest.Optional, acctest.Update, DataccVmInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "boot_storage_size_in_gbs", bootStorageSizeInGbs),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpus_enabled", cpusEnabled),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_gb", dataStorageSizeInGbsU),
				resource.TestCheckResourceAttr(resourceName, "description", descriptionU),
				resource.TestCheckResourceAttr(resourceName, "display_name", displayNameU),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", strconv.Itoa(vmInstanceNumFreeformTags)),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "infrastructure_id", infrastructureId),
				resource.TestCheckResourceAttr(resourceName, "memory_size_in_gbs", memorySizeInGbs),
				resource.TestCheckResourceAttr(resourceName, "ssh_public_keys.#", strconv.Itoa(vmInstanceNumSshPublicKeys)),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttrPair(resourceName, "vm_network_id", "oci_datacc_vm_cluster_network.test_vm_cluster_network", "id"),

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
			Config: config + tfVariableStr + DataccVmInstanceDependencyVariableConfig + DataccVmInstanceResourceDependencies +
				acctest.GenerateDataSourceFromRepresentationMap("oci_datacc_vm_instances", "test_vm_instances", acctest.Optional, acctest.Update, DataccVmInstanceDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_datacc_vm_instance", "test_vm_instance", acctest.Optional, acctest.Update, DataccVmInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", displayNameU),
				resource.TestCheckResourceAttr(datasourceName, "infrastructure_id", infrastructureId),
				resource.TestCheckResourceAttr(datasourceName, "state.#", strconv.Itoa(vmInstanceNumStates)),

				resource.TestCheckResourceAttr(datasourceName, "vm_instance_collection.#", strconv.Itoa(vmInstanceNumCollections)),
			),
		},
		// verify singular datasource
		{
			Config: config + tfVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_datacc_vm_instance", "test_vm_instance", acctest.Required, acctest.Create, DataccVmInstanceSingularDataSourceRepresentation) +
				DataccVmInstanceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vm_instance_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "boot_storage_size_in_gbs", bootStorageSizeInGbs),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "cpus_enabled", cpusEnabled),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_storage_size_in_gb", dataStorageSizeInGbsU),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", descriptionU),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", displayNameU),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", strconv.Itoa(vmInstanceNumFreeformTags)),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "memory_size_in_gbs", memorySizeInGbs),
				resource.TestCheckResourceAttr(singularDatasourceName, "ssh_public_keys.#", strconv.Itoa(vmInstanceNumSshPublicKeys)),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + tfVariableStr + DataccVmInstanceRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDataccVmInstanceDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).BaseinfraClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_datacc_vm_instance" {
			noResourceFound = false
			request := oci_datacc.GetVmInstanceRequest{}

			tmp := rs.Primary.ID
			request.VmInstanceId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "datacc")

			response, err := client.GetVmInstance(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_datacc.VmInstanceLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DataccVmInstance") {
		resource.AddTestSweepers("DataccVmInstance", &resource.Sweeper{
			Name:         "DataccVmInstance",
			Dependencies: acctest.DependencyGraph["vmInstance"],
			F:            sweepDataccVmInstanceResource,
		})
	}
}

func sweepDataccVmInstanceResource(compartment string) error {
	baseinfraClient := acctest.GetTestClients(&schema.ResourceData{}).BaseinfraClient()
	vmInstanceIds, err := getDataccVmInstanceIds(compartment)
	if err != nil {
		return err
	}
	for _, vmInstanceId := range vmInstanceIds {
		if ok := acctest.SweeperDefaultResourceId[vmInstanceId]; !ok {
			deleteVmInstanceRequest := oci_datacc.DeleteVmInstanceRequest{}

			deleteVmInstanceRequest.VmInstanceId = &vmInstanceId

			deleteVmInstanceRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "datacc")
			_, error := baseinfraClient.DeleteVmInstance(context.Background(), deleteVmInstanceRequest)
			if error != nil {
				fmt.Printf("Error deleting VmInstance %s %s, It is possible that the resource is already deleted. Please verify manually \n", vmInstanceId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &vmInstanceId, DataccVmInstanceSweepWaitCondition, time.Duration(3*time.Minute),
				DataccVmInstanceSweepResponseFetchOperation, "datacc", true)
		}
	}
	return nil
}

func getDataccVmInstanceIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "VmInstanceId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	baseinfraClient := acctest.GetTestClients(&schema.ResourceData{}).BaseinfraClient()

	listVmInstancesRequest := oci_datacc.ListVmInstancesRequest{}
	listVmInstancesRequest.CompartmentId = &compartmentId
	listVmInstancesRequest.LifecycleState = []oci_datacc.VmInstanceLifecycleStateEnum{oci_datacc.VmInstanceLifecycleStateActive}
	listVmInstancesResponse, err := baseinfraClient.ListVmInstances(context.Background(), listVmInstancesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting VmInstance list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, vmInstance := range listVmInstancesResponse.Items {
		id := *vmInstance.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "VmInstanceId", id)
	}
	return resourceIds, nil
}

func DataccVmInstanceSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if vmInstanceResponse, ok := response.Response.(oci_datacc.GetVmInstanceResponse); ok {
		return vmInstanceResponse.LifecycleState != oci_datacc.VmInstanceLifecycleStateDeleted
	}
	return false
}

func DataccVmInstanceSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.BaseinfraClient().GetVmInstance(context.Background(), oci_datacc.GetVmInstanceRequest{
		VmInstanceId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
