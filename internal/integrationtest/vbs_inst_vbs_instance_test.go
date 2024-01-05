// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_vbs_inst "github.com/oracle/oci-go-sdk/v65/vbsinst"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	VbsInstVbsInstanceRequiredOnlyResource = VbsInstVbsInstanceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_vbs_inst_vbs_instance", "test_vbs_instance", acctest.Required, acctest.Create, VbsInstVbsInstanceRepresentation)

	VbsInstVbsInstanceResourceConfig = VbsInstVbsInstanceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_vbs_inst_vbs_instance", "test_vbs_instance", acctest.Optional, acctest.Update, VbsInstVbsInstanceRepresentation)

	VbsInstVbsInstanceSingularDataSourceRepresentation = map[string]interface{}{
		"vbs_instance_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_vbs_inst_vbs_instance.test_vbs_instance.id}`},
	}

	VbsInstVbsInstanceDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_vbs_inst_vbs_instance.test_vbs_instance.id}`},
		"name":           acctest.Representation{RepType: acctest.Optional, Create: `name`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: VbsInstVbsInstanceDataSourceFilterRepresentation}}
	VbsInstVbsInstanceDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_vbs_inst_vbs_instance.test_vbs_instance.id}`}},
	}

	VbsInstVbsInstanceRepresentation = map[string]interface{}{
		"compartment_id":                      acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"display_name":                        acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"name":                                acctest.Representation{RepType: acctest.Required, Create: `name`},
		"defined_tags":                        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":                       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"is_resource_usage_agreement_granted": acctest.Representation{RepType: acctest.Optional, Create: `true`, Update: `false`},
		"resource_compartment_id":             acctest.Representation{RepType: acctest.Optional, Create: `${var.resource_compartment_id}`, Update: ``},
		"lifecycle":                           acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreVbsInstanceChangesRepresentation},
	}

	ignoreVbsInstanceChangesRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`name`, `defined_tags`, `freeform_tags`}, Update: []string{`name`, `defined_tags`, `freeform_tags`}},
	}

	VbsInstVbsInstanceResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: vbs_inst/default
func TestVbsInstVbsInstanceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestVbsInstVbsInstanceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	tenancyId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	tenancyIdVariableStr := fmt.Sprintf("variable \"tenancy_id\" { default = \"%s\" }\n", tenancyId)

	compartmentIdRes := utils.GetEnvSettingWithDefault("resource_compartment_id", tenancyId)
	compartmentIdResVariableStr := fmt.Sprintf("variable \"resource_compartment_id\" { default = \"%s\" }\n", compartmentIdRes)

	resourceName := "oci_vbs_inst_vbs_instance.test_vbs_instance"
	datasourceName := "data.oci_vbs_inst_vbs_instances.test_vbs_instances"
	singularDatasourceName := "data.oci_vbs_inst_vbs_instance.test_vbs_instance"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+tenancyIdVariableStr+VbsInstVbsInstanceResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_vbs_inst_vbs_instance", "test_vbs_instance", acctest.Optional, acctest.Create, VbsInstVbsInstanceRepresentation), "vbsinst", "vbsInstance", t)

	acctest.ResourceTest(t, testAccCheckVbsInstVbsInstanceDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + tenancyIdVariableStr + VbsInstVbsInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_vbs_inst_vbs_instance", "test_vbs_instance", acctest.Required, acctest.Create, VbsInstVbsInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestMatchResourceAttr(resourceName, "name", regexp.MustCompile("name*")),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + tenancyIdVariableStr + VbsInstVbsInstanceResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + tenancyIdVariableStr + compartmentIdResVariableStr + VbsInstVbsInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_vbs_inst_vbs_instance", "test_vbs_instance", acctest.Optional, acctest.Create, VbsInstVbsInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_resource_usage_agreement_granted", "true"),
				resource.TestMatchResourceAttr(resourceName, "name", regexp.MustCompile("name*")),
				resource.TestCheckResourceAttrSet(resourceName, "resource_compartment_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &tenancyId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			// It is root compartment resource. Hence using tenancy_ocid
			Config: config + tenancyIdVariableStr + compartmentIdResVariableStr + VbsInstVbsInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_vbs_inst_vbs_instance", "test_vbs_instance", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(VbsInstVbsInstanceRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_resource_usage_agreement_granted", "true"),
				resource.TestMatchResourceAttr(resourceName, "name", regexp.MustCompile("name*")),
				resource.TestCheckResourceAttrSet(resourceName, "resource_compartment_id"),

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
			Config: config + tenancyIdVariableStr + compartmentIdResVariableStr + VbsInstVbsInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_vbs_inst_vbs_instance", "test_vbs_instance", acctest.Optional, acctest.Update, VbsInstVbsInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_resource_usage_agreement_granted", "false"),
				resource.TestMatchResourceAttr(resourceName, "name", regexp.MustCompile("name*")),
				resource.TestCheckResourceAttrSet(resourceName, "resource_compartment_id"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_vbs_inst_vbs_instances", "test_vbs_instances", acctest.Optional, acctest.Update, VbsInstVbsInstanceDataSourceRepresentation) +
				tenancyIdVariableStr + VbsInstVbsInstanceResourceDependencies + compartmentIdResVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_vbs_inst_vbs_instance", "test_vbs_instance", acctest.Optional, acctest.Update, VbsInstVbsInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestMatchResourceAttr(datasourceName, "name", regexp.MustCompile("name*")),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttrSet(datasourceName, "vbs_instance_summary_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "vbs_instance_summary_collection.0.items.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_vbs_inst_vbs_instance", "test_vbs_instance", acctest.Required, acctest.Create, VbsInstVbsInstanceSingularDataSourceRepresentation) +
				tenancyIdVariableStr + compartmentIdResVariableStr + VbsInstVbsInstanceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vbs_instance_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_resource_usage_agreement_granted", "false"),
				resource.TestMatchResourceAttr(singularDatasourceName, "name", regexp.MustCompile("name*")),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vbs_access_url"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + tenancyIdVariableStr + compartmentIdResVariableStr + VbsInstVbsInstanceResourceConfig,
		},
		// verify resource import
		{
			Config:            config + VbsInstVbsInstanceRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"idcs_access_token",
				"resource_compartment_id",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckVbsInstVbsInstanceDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).VbsInstanceClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_vbs_inst_vbs_instance" {
			noResourceFound = false
			request := oci_vbs_inst.GetVbsInstanceRequest{}

			tmp := rs.Primary.ID
			request.VbsInstanceId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "vbs_inst")

			response, err := client.GetVbsInstance(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_vbs_inst.LifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("VbsInstVbsInstance") {
		resource.AddTestSweepers("VbsInstVbsInstance", &resource.Sweeper{
			Name:         "VbsInstVbsInstance",
			Dependencies: acctest.DependencyGraph["vbsInstance"],
			F:            sweepVbsInstVbsInstanceResource,
		})
	}
}

func sweepVbsInstVbsInstanceResource(compartment string) error {
	vbsInstanceClient := acctest.GetTestClients(&schema.ResourceData{}).VbsInstanceClient()
	vbsInstanceIds, err := getVbsInstVbsInstanceIds(compartment)
	if err != nil {
		return err
	}
	for _, vbsInstanceId := range vbsInstanceIds {
		if ok := acctest.SweeperDefaultResourceId[vbsInstanceId]; !ok {
			deleteVbsInstanceRequest := oci_vbs_inst.DeleteVbsInstanceRequest{}

			deleteVbsInstanceRequest.VbsInstanceId = &vbsInstanceId

			deleteVbsInstanceRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "vbs_inst")
			_, error := vbsInstanceClient.DeleteVbsInstance(context.Background(), deleteVbsInstanceRequest)
			if error != nil {
				fmt.Printf("Error deleting VbsInstance %s %s, It is possible that the resource is already deleted. Please verify manually \n", vbsInstanceId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &vbsInstanceId, VbsInstVbsInstanceSweepWaitCondition, time.Duration(3*time.Minute),
				VbsInstVbsInstanceSweepResponseFetchOperation, "vbs_inst", true)
		}
	}
	return nil
}

func getVbsInstVbsInstanceIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "VbsInstanceId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	vbsInstanceClient := acctest.GetTestClients(&schema.ResourceData{}).VbsInstanceClient()

	listVbsInstancesRequest := oci_vbs_inst.ListVbsInstancesRequest{}
	listVbsInstancesRequest.CompartmentId = &compartmentId
	listVbsInstancesRequest.LifecycleState = oci_vbs_inst.ListVbsInstancesLifecycleStateActive
	listVbsInstancesResponse, err := vbsInstanceClient.ListVbsInstances(context.Background(), listVbsInstancesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting VbsInstance list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, vbsInstance := range listVbsInstancesResponse.Items {
		id := *vbsInstance.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "VbsInstanceId", id)
	}
	return resourceIds, nil
}

func VbsInstVbsInstanceSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if vbsInstanceResponse, ok := response.Response.(oci_vbs_inst.GetVbsInstanceResponse); ok {
		return vbsInstanceResponse.LifecycleState != oci_vbs_inst.LifecycleStateDeleted
	}
	return false
}

func VbsInstVbsInstanceSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.VbsInstanceClient().GetVbsInstance(context.Background(), oci_vbs_inst.GetVbsInstanceRequest{
		VbsInstanceId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
