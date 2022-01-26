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
	"github.com/oracle/oci-go-sdk/v56/common"
	oci_golden_gate "github.com/oracle/oci-go-sdk/v56/goldengate"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	DeploymentBackupRequiredOnlyResource = DeploymentBackupResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_deployment_backup", "test_deployment_backup", acctest.Required, acctest.Create, deploymentBackupRepresentation)

	DeploymentBackupResourceConfig = DeploymentBackupResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_deployment_backup", "test_deployment_backup", acctest.Optional, acctest.Update, deploymentBackupRepresentation)

	deploymentBackupSingularDataSourceRepresentation = map[string]interface{}{
		"deployment_backup_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_golden_gate_deployment_backup.test_deployment_backup.id}`},
	}

	deploymentBackupDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"deployment_id":  acctest.Representation{RepType: acctest.Optional, Create: `${oci_golden_gate_deployment.test_ggsdeployment.id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `demoDeploymentBackup`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: deploymentBackupDataSourceFilterRepresentation}}
	deploymentBackupDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_golden_gate_deployment_backup.test_deployment_backup.id}`}},
	}

	deploymentBackupRepresentation = map[string]interface{}{
		"bucket":         acctest.Representation{RepType: acctest.Required, Create: `${oci_objectstorage_bucket.test_bucket.name}`},
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"deployment_id":  acctest.Representation{RepType: acctest.Required, Create: `${oci_golden_gate_deployment.test_ggsdeployment.id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `demoDeploymentBackup`},
		"namespace":      acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"object":         acctest.Representation{RepType: acctest.Required, Create: `object`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreGGSDefinedTagsChangesRepresentation2},
	}

	ignoreGGSDefinedTagsChangesRepresentation2 = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}

	DeploymentBackupResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, subnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, vcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_deployment", "test_ggsdeployment", acctest.Required, acctest.Create, goldenGateDeploymentRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", acctest.Required, acctest.Create, bucketRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", acctest.Required, acctest.Create, namespaceSingularDataSourceRepresentation)
)

// issue-routing-tag: golden_gate/default
func TestGoldenGateDeploymentBackupResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestGoldenGateDeploymentBackupResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_golden_gate_deployment_backup.test_deployment_backup"
	datasourceName := "data.oci_golden_gate_deployment_backups.test_deployment_backups"
	singularDatasourceName := "data.oci_golden_gate_deployment_backup.test_deployment_backup"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DeploymentBackupResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_deployment_backup", "test_deployment_backup", acctest.Optional, acctest.Create, deploymentBackupRepresentation), "goldengate", "deploymentBackup", t)

	fmt.Printf("Terraform generated %s", config+compartmentIdVariableStr+DeploymentBackupResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_deployment_backup", "test_deployment_backup", acctest.Optional, acctest.Create, deploymentBackupRepresentation))

	acctest.ResourceTest(t, testAccCheckGoldenGateDeploymentBackupDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DeploymentBackupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_deployment_backup", "test_deployment_backup", acctest.Required, acctest.Create, deploymentBackupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "bucket"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "deployment_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "demoDeploymentBackup"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttr(resourceName, "object", "object"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DeploymentBackupResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DeploymentBackupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_deployment_backup", "test_deployment_backup", acctest.Optional, acctest.Create, deploymentBackupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "bucket"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "deployment_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "demoDeploymentBackup"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttr(resourceName, "object", "object"),
				resource.TestCheckResourceAttrSet(resourceName, "ogg_version"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DeploymentBackupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_deployment_backup", "test_deployment_backup", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(deploymentBackupRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "bucket"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttrSet(resourceName, "deployment_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "demoDeploymentBackup"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttr(resourceName, "object", "object"),
				resource.TestCheckResourceAttrSet(resourceName, "ogg_version"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

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
			Config: config + compartmentIdVariableStr + DeploymentBackupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_deployment_backup", "test_deployment_backup", acctest.Optional, acctest.Update, deploymentBackupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "bucket"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "deployment_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "demoDeploymentBackup"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttr(resourceName, "object", "object"),
				resource.TestCheckResourceAttrSet(resourceName, "ogg_version"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_golden_gate_deployment_backups", "test_deployment_backups", acctest.Optional, acctest.Update, deploymentBackupDataSourceRepresentation) +
				compartmentIdVariableStr + DeploymentBackupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_deployment_backup", "test_deployment_backup", acctest.Optional, acctest.Update, deploymentBackupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "deployment_id"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "demoDeploymentBackup"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "deployment_backup_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "deployment_backup_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_golden_gate_deployment_backup", "test_deployment_backup", acctest.Required, acctest.Create, deploymentBackupSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DeploymentBackupResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "deployment_backup_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "backup_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "bucket"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "demoDeploymentBackup"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_automatic"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "namespace"),
				resource.TestCheckResourceAttr(singularDatasourceName, "object", "object"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ogg_version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "size_in_bytes"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_backup_finished"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_of_backup"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + DeploymentBackupResourceConfig,
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

func testAccCheckGoldenGateDeploymentBackupDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).GoldenGateClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_golden_gate_deployment_backup" {
			noResourceFound = false
			request := oci_golden_gate.GetDeploymentBackupRequest{}

			tmp := rs.Primary.ID
			request.DeploymentBackupId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "golden_gate")

			response, err := client.GetDeploymentBackup(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_golden_gate.LifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("GoldenGateDeploymentBackup") {
		resource.AddTestSweepers("GoldenGateDeploymentBackup", &resource.Sweeper{
			Name:         "GoldenGateDeploymentBackup",
			Dependencies: acctest.DependencyGraph["deploymentBackup"],
			F:            sweepGoldenGateDeploymentBackupResource,
		})
	}
}

func sweepGoldenGateDeploymentBackupResource(compartment string) error {
	goldenGateClient := acctest.GetTestClients(&schema.ResourceData{}).GoldenGateClient()
	deploymentBackupIds, err := getDeploymentBackupIds(compartment)
	if err != nil {
		return err
	}
	for _, deploymentBackupId := range deploymentBackupIds {
		if ok := acctest.SweeperDefaultResourceId[deploymentBackupId]; !ok {
			deleteDeploymentBackupRequest := oci_golden_gate.DeleteDeploymentBackupRequest{}

			deleteDeploymentBackupRequest.DeploymentBackupId = &deploymentBackupId

			deleteDeploymentBackupRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "golden_gate")
			_, error := goldenGateClient.DeleteDeploymentBackup(context.Background(), deleteDeploymentBackupRequest)
			if error != nil {
				fmt.Printf("Error deleting DeploymentBackup %s %s, It is possible that the resource is already deleted. Please verify manually \n", deploymentBackupId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &deploymentBackupId, deploymentBackupSweepWaitCondition, time.Duration(3*time.Minute),
				deploymentBackupSweepResponseFetchOperation, "golden_gate", true)
		}
	}
	return nil
}

func getDeploymentBackupIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DeploymentBackupId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	goldenGateClient := acctest.GetTestClients(&schema.ResourceData{}).GoldenGateClient()

	listDeploymentBackupsRequest := oci_golden_gate.ListDeploymentBackupsRequest{}
	listDeploymentBackupsRequest.CompartmentId = &compartmentId
	listDeploymentBackupsRequest.LifecycleState = oci_golden_gate.ListDeploymentBackupsLifecycleStateActive
	listDeploymentBackupsResponse, err := goldenGateClient.ListDeploymentBackups(context.Background(), listDeploymentBackupsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting DeploymentBackup list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, deploymentBackup := range listDeploymentBackupsResponse.Items {
		id := *deploymentBackup.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DeploymentBackupId", id)
	}
	return resourceIds, nil
}

func deploymentBackupSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if deploymentBackupResponse, ok := response.Response.(oci_golden_gate.GetDeploymentBackupResponse); ok {
		return deploymentBackupResponse.LifecycleState != oci_golden_gate.LifecycleStateDeleted
	}
	return false
}

func deploymentBackupSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.GoldenGateClient().GetDeploymentBackup(context.Background(), oci_golden_gate.GetDeploymentBackupRequest{
		DeploymentBackupId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
