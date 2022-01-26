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
	goldenGateDeploymentRequiredOnlyResource = GoldenGateDeploymentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_deployment", "depl_test_ggs_deployment", acctest.Required, acctest.Create, goldenGateDeploymentRepresentation)

	goldenGateDeploymentResourceConfig = GoldenGateDeploymentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_deployment", "depl_test_ggs_deployment", acctest.Optional, acctest.Update, goldenGateDeploymentRepresentation)

	goldenGateDeploymentSingularDataSourceRepresentation = map[string]interface{}{
		"deployment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_golden_gate_deployment.depl_test_ggs_deployment.id}`},
	}

	goldenGateDeploymentDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"fqdn":           acctest.Representation{RepType: acctest.Optional, Create: ``},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: goldenGateDeploymentDataSourceFilterRepresentation}}
	goldenGateDeploymentDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_golden_gate_deployment.depl_test_ggs_deployment.id}`}},
	}

	goldenGateDeploymentRepresentation = map[string]interface{}{
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"cpu_core_count":          acctest.Representation{RepType: acctest.Required, Create: `1`},
		"deployment_type":         acctest.Representation{RepType: acctest.Required, Create: `OGG`},
		"display_name":            acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"is_auto_scaling_enabled": acctest.Representation{RepType: acctest.Required, Create: `false`},
		"subnet_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"license_model":           acctest.Representation{RepType: acctest.Required, Create: `LICENSE_INCLUDED`},
		"description":             acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"fqdn":                    acctest.Representation{RepType: acctest.Optional, Create: ``},
		"freeform_tags":           acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"is_public":               acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"ogg_data":                acctest.RepresentationGroup{RepType: acctest.Required, Group: goldenGateDeploymentOggDataRepresentation},
		"lifecycle":               acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreGGSDefinedTagsChangesRepresentation1},
	}

	ignoreGGSDefinedTagsChangesRepresentation1 = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}

	goldenGateDeploymentOggDataRepresentation = map[string]interface{}{
		"admin_password":  acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`, Update: `BEstrO0ng_#12`},
		"admin_username":  acctest.Representation{RepType: acctest.Required, Create: `adminUsername`, Update: `adminUsername2`},
		"deployment_name": acctest.Representation{RepType: acctest.Required, Create: `depl_test_ggs_deployment_name`},
		"certificate":     acctest.Representation{RepType: acctest.Optional, Create: ``, Update: `-----BEGIN CERTIFICATE-----\nMIICljCCAX4CCQCEpaMjTCJ8WzANBgkqhkiG9w0BAQsFADANMQswCQYDVQQGEwJV\nUzAeFw0yMTAxMTkyMTI2MjRaFw0yNDAxMTkyMTI2MjRaMA0xCzAJBgNVBAYTAlVT\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAo83kaUQXpCcSoEuRVFX3\njztWDNKtWpjNG240f0RpERI1NnZtHH0qnZqfaWAQQa8kx3+W1LOeFbkkRnkJz19g\neIXR6TeavT+W5iRh4goK+N7gubYkSMa2shVf+XsoHKERSbhdhrtX+GqvKzAvplCt\nCgd4MDlsvLv/YHCLvJL4JgRxKyevjlnE1rqrICJMCLbbZMrIKTzwb/K13hGrm6Bc\n+Je9EC3MWWxd5jBwXu3vgIYRuGR4DPg/yfMKPZr2xFDLpBsv5jaqULS9t6GwoEBJ\nKN0NXp5obaQToYqMsvAZyHoEyfCBDka16Bm5hGF60FwqgUT3p/+qlBn61cAJe9t5\n8QIDAQABMA0GCSqGSIb3DQEBCwUAA4IBAQAX1rxV2hai02Pb4Cf8U44zj+1aY6wV\nLvOMWiL3zl53up4/X7PDcmWcPM9UMVCGTISZD6A6IPvNlkvbtvYCzgjhtGxDmrj7\nwTRV5gO9j3bAhxBO7XgTmwmD/9hpykM58nbhLFnkGf+Taja8qsy0U8H74Tr9w1M8\n8E5kghgGzBElNquM8AUuDakC1JL4aLO/VDMxe/1BLtmBHLZy3XTzVycjP9ZFPh6h\nT+cWJcVOjQSYY2U75sDnKD2Sg1cmK54HauA6SPh4kAkpmxyLyDZZjPBQe2sLFmmS\naZSE+g16yMR9TVHo3pTpRkxJwDEH0LePwYXA4vUIK3HHS6zgLe0ody8g\n-----END CERTIFICATE-----`},
		"key":             acctest.Representation{RepType: acctest.Optional, Create: ``, Update: `${var.golden_gate_deployment_ogg_key}`},
	}

	GoldenGateDeploymentResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, networkSecurityGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, subnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, vcnRepresentation)
)

// issue-routing-tag: golden_gate/default
func TestGoldenGateDeploymentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestGoldenGateDeploymentResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	goldenGateDeploymentOggKey := utils.GetEnvSettingWithBlankDefault("golden_gate_deployment_ogg_key")
	goldenGateDeploymentOggKeyVariableStr := fmt.Sprintf("variable \"golden_gate_deployment_ogg_key\" { default = \"%s\" }\n", goldenGateDeploymentOggKey)

	resourceName := "oci_golden_gate_deployment.depl_test_ggs_deployment"
	datasourceName := "data.oci_golden_gate_deployments.depl_test_ggs_deployments"
	singularDatasourceName := "data.oci_golden_gate_deployment.depl_test_ggs_deployment"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DeploymentResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_deployment", "depl_test_ggs_deployment", acctest.Optional, acctest.Create, goldenGateDeploymentRepresentation), "goldengate", "deployment", t)

	acctest.ResourceTest(t, testAccCheckGoldenGateDeploymentDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + GoldenGateDeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_deployment", "depl_test_ggs_deployment", acctest.Required, acctest.Create, goldenGateDeploymentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "deployment_type", "OGG"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "ogg_data.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ogg_data.0.admin_username", "adminUsername"),
				resource.TestCheckResourceAttrSet(resourceName, "ogg_data.0.deployment_name"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + GoldenGateDeploymentResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + goldenGateDeploymentOggKeyVariableStr + GoldenGateDeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_deployment", "depl_test_ggs_deployment", acctest.Optional, acctest.Create, goldenGateDeploymentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "deployment_type", "OGG"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "fqdn"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_public", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "ogg_data.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ogg_data.0.admin_username", "adminUsername"),
				resource.TestCheckResourceAttrSet(resourceName, "ogg_data.0.deployment_name"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + goldenGateDeploymentOggKeyVariableStr + GoldenGateDeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_deployment", "depl_test_ggs_deployment", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(goldenGateDeploymentRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "deployment_type", "OGG"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "fqdn"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_public", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "ogg_data.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ogg_data.0.admin_username", "adminUsername"),
				resource.TestCheckResourceAttrSet(resourceName, "ogg_data.0.deployment_name"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

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
			Config: config + compartmentIdVariableStr + goldenGateDeploymentOggKeyVariableStr + GoldenGateDeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_deployment", "depl_test_ggs_deployment", acctest.Optional, acctest.Update, goldenGateDeploymentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "deployment_type", "OGG"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "fqdn"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_public", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "ogg_data.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ogg_data.0.admin_username", "adminUsername2"),
				resource.TestCheckResourceAttrSet(resourceName, "ogg_data.0.certificate"),
				resource.TestCheckResourceAttrSet(resourceName, "ogg_data.0.deployment_name"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

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
			Config: config + compartmentIdVariableStr + goldenGateDeploymentOggKeyVariableStr + GoldenGateDeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_deployment", "depl_test_ggs_deployment", acctest.Optional, acctest.Update, goldenGateDeploymentRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_golden_gate_deployments", "depl_test_ggs_deployments", acctest.Required, acctest.Update, goldenGateDeploymentDataSourceRepresentation),

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "deployment_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "deployment_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config + compartmentIdVariableStr + goldenGateDeploymentOggKeyVariableStr + GoldenGateDeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_deployment", "depl_test_ggs_deployment", acctest.Optional, acctest.Update, goldenGateDeploymentRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_golden_gate_deployment", "depl_test_ggs_deployment", acctest.Required, acctest.Create, goldenGateDeploymentSingularDataSourceRepresentation),

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "deployment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deployment_type", "OGG"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "deployment_url"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "fqdn"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_healthy"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_latest_version"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_public", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ogg_data.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ogg_data.0.admin_username", "adminUsername2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ogg_data.0.certificate"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "private_ip_address"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + goldenGateDeploymentOggKeyVariableStr + goldenGateDeploymentResourceConfig,
		},
		// verify resource import
		{
			Config:            config,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"ogg_data.0.admin_password",
				"ogg_data.0.key",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckGoldenGateDeploymentDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).GoldenGateClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_golden_gate_deployment" {
			noResourceFound = false
			request := oci_golden_gate.GetDeploymentRequest{}

			tmp := rs.Primary.ID
			request.DeploymentId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "golden_gate")

			response, err := client.GetDeployment(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("GoldenGateDeployment") {
		resource.AddTestSweepers("GoldenGateDeployment", &resource.Sweeper{
			Name:         "GoldenGateDeployment",
			Dependencies: acctest.DependencyGraph["deployment"],
			F:            sweepGoldenGateDeploymentResource,
		})
	}
}

func sweepGoldenGateDeploymentResource(compartment string) error {
	goldenGateClient := acctest.GetTestClients(&schema.ResourceData{}).GoldenGateClient()
	deploymentIds, err := getGoldenGateDeploymentIds(compartment)
	if err != nil {
		return err
	}
	for _, deploymentId := range deploymentIds {
		if ok := acctest.SweeperDefaultResourceId[deploymentId]; !ok {
			deleteDeploymentRequest := oci_golden_gate.DeleteDeploymentRequest{}

			deleteDeploymentRequest.DeploymentId = &deploymentId

			deleteDeploymentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "golden_gate")
			_, error := goldenGateClient.DeleteDeployment(context.Background(), deleteDeploymentRequest)
			if error != nil {
				fmt.Printf("Error deleting Deployment %s %s, It is possible that the resource is already deleted. Please verify manually \n", deploymentId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &deploymentId, goldenGateDeploymentSweepWaitCondition, time.Duration(3*time.Minute),
				goldenGateDeploymentSweepResponseFetchOperation, "golden_gate", true)
		}
	}
	return nil
}

func getGoldenGateDeploymentIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DeploymentId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	goldenGateClient := acctest.GetTestClients(&schema.ResourceData{}).GoldenGateClient()

	listDeploymentsRequest := oci_golden_gate.ListDeploymentsRequest{}
	listDeploymentsRequest.CompartmentId = &compartmentId
	listDeploymentsRequest.LifecycleState = oci_golden_gate.ListDeploymentsLifecycleStateActive
	listDeploymentsResponse, err := goldenGateClient.ListDeployments(context.Background(), listDeploymentsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Deployment list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, deployment := range listDeploymentsResponse.Items {
		id := *deployment.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DeploymentId", id)
	}
	return resourceIds, nil
}

func goldenGateDeploymentSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if deploymentResponse, ok := response.Response.(oci_golden_gate.GetDeploymentResponse); ok {
		return deploymentResponse.LifecycleState != oci_golden_gate.LifecycleStateDeleted
	}
	return false
}

func goldenGateDeploymentSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.GoldenGateClient().GetDeployment(context.Background(), oci_golden_gate.GetDeploymentRequest{
		DeploymentId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
