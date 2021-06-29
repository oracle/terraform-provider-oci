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
	"github.com/oracle/oci-go-sdk/v43/common"
	oci_golden_gate "github.com/oracle/oci-go-sdk/v43/goldengate"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	DeploymentBackupRequiredOnlyResource = DeploymentBackupResourceDependencies +
		generateResourceFromRepresentationMap("oci_golden_gate_deployment_backup", "test_deployment_backup", Required, Create, deploymentBackupRepresentation)

	DeploymentBackupResourceConfig = DeploymentBackupResourceDependencies +
		generateResourceFromRepresentationMap("oci_golden_gate_deployment_backup", "test_deployment_backup", Optional, Update, deploymentBackupRepresentation)

	deploymentBackupSingularDataSourceRepresentation = map[string]interface{}{
		"deployment_backup_id": Representation{repType: Required, create: `${oci_golden_gate_deployment_backup.test_deployment_backup.id}`},
	}

	deploymentBackupDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"deployment_id":  Representation{repType: Optional, create: `${oci_golden_gate_deployment.test_backup_deployment.id}`},
		"display_name":   Representation{repType: Optional, create: `demoDeploymentBackup`},
		"state":          Representation{repType: Optional, create: `ACTIVE`},
		"filter":         RepresentationGroup{Required, deploymentBackupDataSourceFilterRepresentation}}
	deploymentBackupDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_golden_gate_deployment_backup.test_deployment_backup.id}`}},
	}

	deploymentBackupRepresentation = map[string]interface{}{
		"bucket":         Representation{repType: Required, create: `${oci_objectstorage_bucket.test_bucket.name}`},
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"deployment_id":  Representation{repType: Required, create: `${oci_golden_gate_deployment.test_backup_deployment.id}`},
		"display_name":   Representation{repType: Required, create: `demoDeploymentBackup`},
		"namespace":      Representation{repType: Required, create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"object":         Representation{repType: Required, create: `object`},
		"defined_tags":   Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":  Representation{repType: Optional, create: map[string]string{"bar-key": "value"}, update: map[string]string{"Department": "Accounting"}},
	}

	DeploymentBackupResourceDependencies = generateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, subnetRepresentation) +
		generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, vcnRepresentation) +
		generateResourceFromRepresentationMap("oci_golden_gate_deployment", "test_backup_deployment", Required, Create, goldenGateDeploymentRepresentation) +
		DefinedTagsDependencies +
		generateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", Required, Create, bucketRepresentation) +
		generateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", Required, Create, namespaceSingularDataSourceRepresentation)
)

func TestGoldenGateDeploymentBackupResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestGoldenGateDeploymentBackupResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_golden_gate_deployment_backup.test_deployment_backup"
	datasourceName := "data.oci_golden_gate_deployment_backups.test_deployment_backups"
	singularDatasourceName := "data.oci_golden_gate_deployment_backup.test_deployment_backup"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+DeploymentBackupResourceDependencies+
		generateResourceFromRepresentationMap("oci_golden_gate_deployment_backup", "test_deployment_backup", Optional, Create, deploymentBackupRepresentation), "goldengate", "deploymentBackup", t)

	fmt.Printf("Terraform generated %s", config+compartmentIdVariableStr+DeploymentBackupResourceDependencies+
		generateResourceFromRepresentationMap("oci_golden_gate_deployment_backup", "test_deployment_backup", Optional, Create, deploymentBackupRepresentation))

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckGoldenGateDeploymentBackupDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + DeploymentBackupResourceDependencies +
					generateResourceFromRepresentationMap("oci_golden_gate_deployment_backup", "test_deployment_backup", Required, Create, deploymentBackupRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "bucket"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "deployment_id"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "demoDeploymentBackup"),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),
					resource.TestCheckResourceAttr(resourceName, "object", "object"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + DeploymentBackupResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + DeploymentBackupResourceDependencies +
					generateResourceFromRepresentationMap("oci_golden_gate_deployment_backup", "test_deployment_backup", Optional, Create, deploymentBackupRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "bucket"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "deployment_id"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "demoDeploymentBackup"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),
					resource.TestCheckResourceAttr(resourceName, "object", "object"),
					resource.TestCheckResourceAttrSet(resourceName, "ogg_version"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),

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
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DeploymentBackupResourceDependencies +
					generateResourceFromRepresentationMap("oci_golden_gate_deployment_backup", "test_deployment_backup", Optional, Create,
						representationCopyWithNewProperties(deploymentBackupRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "bucket"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "deployment_id"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "demoDeploymentBackup"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),
					resource.TestCheckResourceAttr(resourceName, "object", "object"),
					resource.TestCheckResourceAttrSet(resourceName, "ogg_version"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),

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
				Config: config + compartmentIdVariableStr + DeploymentBackupResourceDependencies +
					generateResourceFromRepresentationMap("oci_golden_gate_deployment_backup", "test_deployment_backup", Optional, Update, deploymentBackupRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "bucket"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "deployment_id"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "demoDeploymentBackup"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),
					resource.TestCheckResourceAttr(resourceName, "object", "object"),
					resource.TestCheckResourceAttrSet(resourceName, "ogg_version"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),

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
					generateDataSourceFromRepresentationMap("oci_golden_gate_deployment_backups", "test_deployment_backups", Optional, Update, deploymentBackupDataSourceRepresentation) +
					compartmentIdVariableStr + DeploymentBackupResourceDependencies +
					generateResourceFromRepresentationMap("oci_golden_gate_deployment_backup", "test_deployment_backup", Optional, Update, deploymentBackupRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
					generateDataSourceFromRepresentationMap("oci_golden_gate_deployment_backup", "test_deployment_backup", Required, Create, deploymentBackupSingularDataSourceRepresentation) +
					compartmentIdVariableStr + DeploymentBackupResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "deployment_backup_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "backup_type"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "bucket"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "demoDeploymentBackup"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "is_automatic"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "namespace"),
					resource.TestCheckResourceAttr(singularDatasourceName, "object", "object"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "ogg_version"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
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
		},
	})
}

func testAccCheckGoldenGateDeploymentBackupDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).goldenGateClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_golden_gate_deployment_backup" {
			noResourceFound = false
			request := oci_golden_gate.GetDeploymentBackupRequest{}

			tmp := rs.Primary.ID
			request.DeploymentBackupId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "golden_gate")

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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("GoldenGateDeploymentBackup") {
		resource.AddTestSweepers("GoldenGateDeploymentBackup", &resource.Sweeper{
			Name:         "GoldenGateDeploymentBackup",
			Dependencies: DependencyGraph["deploymentBackup"],
			F:            sweepGoldenGateDeploymentBackupResource,
		})
	}
}

func sweepGoldenGateDeploymentBackupResource(compartment string) error {
	goldenGateClient := GetTestClients(&schema.ResourceData{}).goldenGateClient()
	deploymentBackupIds, err := getDeploymentBackupIds(compartment)
	if err != nil {
		return err
	}
	for _, deploymentBackupId := range deploymentBackupIds {
		if ok := SweeperDefaultResourceId[deploymentBackupId]; !ok {
			deleteDeploymentBackupRequest := oci_golden_gate.DeleteDeploymentBackupRequest{}

			deleteDeploymentBackupRequest.DeploymentBackupId = &deploymentBackupId

			deleteDeploymentBackupRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "golden_gate")
			_, error := goldenGateClient.DeleteDeploymentBackup(context.Background(), deleteDeploymentBackupRequest)
			if error != nil {
				fmt.Printf("Error deleting DeploymentBackup %s %s, It is possible that the resource is already deleted. Please verify manually \n", deploymentBackupId, error)
				continue
			}
			waitTillCondition(testAccProvider, &deploymentBackupId, deploymentBackupSweepWaitCondition, time.Duration(3*time.Minute),
				deploymentBackupSweepResponseFetchOperation, "golden_gate", true)
		}
	}
	return nil
}

func getDeploymentBackupIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "DeploymentBackupId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	goldenGateClient := GetTestClients(&schema.ResourceData{}).goldenGateClient()

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
		addResourceIdToSweeperResourceIdMap(compartmentId, "DeploymentBackupId", id)
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

func deploymentBackupSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.goldenGateClient().GetDeploymentBackup(context.Background(), oci_golden_gate.GetDeploymentBackupRequest{
		DeploymentBackupId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
