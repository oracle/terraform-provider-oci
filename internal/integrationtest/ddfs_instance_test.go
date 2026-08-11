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
	oci_ddfs "github.com/oracle/oci-go-sdk/v65/ddfs"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DdfsInstanceRequiredOnlyResource = DdfsInstanceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ddfs_instance", "test_instance", acctest.Required, acctest.Create, DdfsInstanceRepresentation)

	DdfsInstanceResourceConfig = DdfsInstanceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ddfs_instance", "test_instance", acctest.Optional, acctest.Update, DdfsInstanceRepresentation)

	DdfsInstanceSingularDataSourceRepresentation = map[string]interface{}{
		"instance_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_ddfs_instance.test_instance.id}`},
	}

	ignoreDdfsInstanceDefinedTagsChangesRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}

	DdfsInstanceDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_ddfs_instance.test_instance.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: DdfsInstanceDataSourceFilterRepresentation}}
	DdfsInstanceDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_ddfs_instance.test_instance.id}`}},
	}

	DdfsInstanceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"idcs_url":       acctest.Representation{RepType: acctest.Required, Create: `${var.idcs_url}`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${tomap({"${var.home_region_defined_tag_namespace_name}.${var.home_region_defined_tag_name}" = "value"})}`, Update: `${tomap({"${var.home_region_defined_tag_namespace_name}.${var.home_region_defined_tag_name}" = "updatedValue"})}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDdfsInstanceDefinedTagsChangesRepresentation},
	}

	// DDFS instances are created in a non-home region, so defined tags must be
	// pre-created in the tenancy home region and supplied to the test.
	DdfsInstanceResourceDependencies = `
variable "home_region_defined_tag_namespace_name" {}
variable "home_region_defined_tag_name" {}
`
)

// issue-routing-tag: ddfs/default
func TestDdfsInstanceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDdfsInstanceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	idcsUrl := utils.GetEnvSettingWithBlankDefault("idcs_url")
	idcsUrlVariableStr := fmt.Sprintf("variable \"idcs_url\" { default = \"%s\" }\n", idcsUrl)

	resourceName := "oci_ddfs_instance.test_instance"
	datasourceName := "data.oci_ddfs_instances.test_instances"
	singularDatasourceName := "data.oci_ddfs_instance.test_instance"

	var resId, resId2 string
	// Save TF content with optional properties for generated config/example coverage.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+idcsUrlVariableStr+DdfsInstanceResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_ddfs_instance", "test_instance", acctest.Optional, acctest.Create, DdfsInstanceRepresentation), "ddfs", "instance", t)

	acctest.ResourceTest(t, testAccCheckDdfsInstanceDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + idcsUrlVariableStr + DdfsInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ddfs_instance", "test_instance", acctest.Required, acctest.Create, DdfsInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "idcs_url", idcsUrl),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify update to optionals
		{
			Config: config + compartmentIdVariableStr + idcsUrlVariableStr + DdfsInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ddfs_instance", "test_instance", acctest.Optional, acctest.Create, DdfsInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "idcs_url", idcsUrl),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if err != nil {
						return err
					}
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return nil
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + idcsUrlVariableStr + DdfsInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ddfs_instance", "test_instance", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DdfsInstanceRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "idcs_url", idcsUrl),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if err != nil {
						return err
					}
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return nil
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + idcsUrlVariableStr + DdfsInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ddfs_instance", "test_instance", acctest.Optional, acctest.Update, DdfsInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "idcs_url", idcsUrl),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if err != nil {
						return err
					}
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return nil
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_ddfs_instances", "test_instances", acctest.Optional, acctest.Update, DdfsInstanceDataSourceRepresentation) +
				compartmentIdVariableStr + idcsUrlVariableStr + DdfsInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ddfs_instance", "test_instance", acctest.Optional, acctest.Update, DdfsInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "instance_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "instance_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_ddfs_instance", "test_instance", acctest.Required, acctest.Create, DdfsInstanceSingularDataSourceRepresentation) +
				compartmentIdVariableStr + idcsUrlVariableStr + DdfsInstanceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "instance_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "fhir_service_endpoint"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "idcs_url", idcsUrl),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "public_ip"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + compartmentIdVariableStr + idcsUrlVariableStr + DdfsInstanceRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDdfsInstanceDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).InstanceClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_ddfs_instance" {
			noResourceFound = false
			request := oci_ddfs.GetInstanceRequest{}

			tmp := rs.Primary.ID
			request.InstanceId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "ddfs")

			response, err := client.GetInstance(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_ddfs.InstanceLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DdfsInstance") {
		resource.AddTestSweepers("DdfsInstance", &resource.Sweeper{
			Name:         "DdfsInstance",
			Dependencies: acctest.DependencyGraph["instance"],
			F:            sweepDdfsInstanceResource,
		})
	}
}

func sweepDdfsInstanceResource(compartment string) error {
	instanceClient := acctest.GetTestClients(&schema.ResourceData{}).InstanceClient()
	instanceIds, err := getDdfsInstanceIds(compartment)
	if err != nil {
		return err
	}
	for _, instanceId := range instanceIds {
		if ok := acctest.SweeperDefaultResourceId[instanceId]; !ok {
			deleteInstanceRequest := oci_ddfs.DeleteInstanceRequest{}

			deleteInstanceRequest.InstanceId = &instanceId

			deleteInstanceRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "ddfs")
			_, error := instanceClient.DeleteInstance(context.Background(), deleteInstanceRequest)
			if error != nil {
				fmt.Printf("Error deleting Instance %s %s, It is possible that the resource is already deleted. Please verify manually \n", instanceId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &instanceId, DdfsInstanceSweepWaitCondition, time.Duration(3*time.Minute),
				DdfsInstanceSweepResponseFetchOperation, "ddfs", true)
		}
	}
	return nil
}

func getDdfsInstanceIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "InstanceId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	instanceClient := acctest.GetTestClients(&schema.ResourceData{}).InstanceClient()

	listInstancesRequest := oci_ddfs.ListInstancesRequest{}
	listInstancesRequest.CompartmentId = &compartmentId
	listInstancesRequest.LifecycleState = oci_ddfs.InstanceLifecycleStateActive
	listInstancesResponse, err := instanceClient.ListInstances(context.Background(), listInstancesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Instance list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, instance := range listInstancesResponse.Items {
		id := *instance.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "InstanceId", id)
	}
	return resourceIds, nil
}

func DdfsInstanceSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if instanceResponse, ok := response.Response.(oci_ddfs.GetInstanceResponse); ok {
		return instanceResponse.LifecycleState != oci_ddfs.InstanceLifecycleStateDeleted
	}
	return false
}

func DdfsInstanceSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.InstanceClient().GetInstance(context.Background(), oci_ddfs.GetInstanceRequest{
		InstanceId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
