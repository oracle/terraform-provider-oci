// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_kms "github.com/oracle/oci-go-sdk/v65/keymanagement"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	KmsEkmsPrivateEndpointRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_kms_ekms_private_endpoint", "test_ekms_private_endpoint", acctest.Required, acctest.Create, KmsEkmsPrivateEndpointRepresentation)

	KmsEkmsPrivateEndpointResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_kms_ekms_private_endpoint", "test_ekms_private_endpoint", acctest.Optional, acctest.Update, KmsEkmsPrivateEndpointRepresentation)

	KmsEkmsPrivateEndpointSingularDataSourceRepresentation = map[string]interface{}{
		"ekms_private_endpoint_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_kms_ekms_private_endpoint.test_ekms_private_endpoint.id}`},
	}

	KmsEkmsPrivateEndpointDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: KmsEkmsPrivateEndpointDataSourceFilterRepresentation}}
	KmsEkmsPrivateEndpointDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_kms_ekms_private_endpoint.test_ekms_private_endpoint.id}`}},
	}

	KmsEkmsPrivateEndpointRepresentation = map[string]interface{}{
		"ca_bundle":               acctest.Representation{RepType: acctest.Required, Create: `${var.ca_bundle}`},
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":            acctest.Representation{RepType: acctest.Required, Create: `EKMS Private Endpoint 1`, Update: `displayName2`},
		"external_key_manager_ip": acctest.Representation{RepType: acctest.Required, Create: `10.0.0.31`},
		"subnet_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.subnet_id}`},
		"defined_tags":            acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":           acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		//"port":          acctest.Representation{RepType: acctest.Optional, Create: `443`},
		"lifecycle": acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesEkmsPeRepresentation},
	}
	KmsEkmsPrivateEndpointResourceDependencies = DefinedTagsDependencies
	ignoreChangesEkmsPeRepresentation          = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`ca_bundle`, `defined_tags`, `freeform_tags`}},
	}
)

// issue-routing-tag: kms/default
func TestKmsEkmsPrivateEndpointResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestKmsEkmsPrivateEndpointResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	subnet_id := utils.GetEnvSettingWithBlankDefault("ekms_subnet_id")
	subnetIdVariableStr := fmt.Sprintf("variable \"subnet_id\" { default = \"%s\" }\n", subnet_id)

	ca_bundle := utils.GetEnvSettingWithBlankDefault("ekms_ca_bundle")
	caBundleVariableStr := fmt.Sprintf("variable \"ca_bundle\" { default = \"%s\" }\n", ca_bundle)

	resourceName := "oci_kms_ekms_private_endpoint.test_ekms_private_endpoint"
	datasourceName := "data.oci_kms_ekms_private_endpoints.test_ekms_private_endpoints"
	singularDatasourceName := "data.oci_kms_ekms_private_endpoint.test_ekms_private_endpoint"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+subnetIdVariableStr+caBundleVariableStr+KmsEkmsPrivateEndpointResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_kms_ekms_private_endpoint", "test_ekms_private_endpoint", acctest.Optional, acctest.Create, KmsEkmsPrivateEndpointRepresentation), "keymanagement", "ekmsPrivateEndpoint", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + subnetIdVariableStr + caBundleVariableStr + KmsEkmsPrivateEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_kms_ekms_private_endpoint", "test_ekms_private_endpoint", acctest.Required, acctest.Create, KmsEkmsPrivateEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "subnet_id", subnet_id),
				resource.TestCheckResourceAttr(resourceName, "ca_bundle", strings.Replace(ca_bundle, "\\n", "\n", -1)),
				resource.TestCheckResourceAttr(resourceName, "display_name", "EKMS Private Endpoint 1"),
				resource.TestCheckResourceAttr(resourceName, "external_key_manager_ip", "10.0.0.31"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr,
		},
		//verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + subnetIdVariableStr + caBundleVariableStr + KmsEkmsPrivateEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_kms_ekms_private_endpoint", "test_ekms_private_endpoint", acctest.Optional, acctest.Create, KmsEkmsPrivateEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "EKMS Private Endpoint 1"),
				resource.TestCheckResourceAttr(resourceName, "external_key_manager_ip", "10.0.0.31"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "subnet_id", subnet_id),
				resource.TestCheckResourceAttr(resourceName, "ca_bundle", strings.Replace(ca_bundle, "\\n", "\n", -1)),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "port", "443"),
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

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + subnetIdVariableStr + caBundleVariableStr + KmsEkmsPrivateEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_kms_ekms_private_endpoint", "test_ekms_private_endpoint", acctest.Optional, acctest.Update, KmsEkmsPrivateEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "external_key_manager_ip", "10.0.0.31"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "subnet_id", subnet_id),
				resource.TestCheckResourceAttr(resourceName, "ca_bundle", strings.Replace(ca_bundle, "\\n", "\n", -1)),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "port", "443"),
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
			Config: config + subnetIdVariableStr + caBundleVariableStr + KmsEkmsPrivateEndpointResourceDependencies +
				acctest.GenerateDataSourceFromRepresentationMap("oci_kms_ekms_private_endpoints", "test_ekms_private_endpoints", acctest.Optional, acctest.Update, KmsEkmsPrivateEndpointDataSourceRepresentation) +
				compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_kms_ekms_private_endpoint", "test_ekms_private_endpoint", acctest.Optional, acctest.Update, KmsEkmsPrivateEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttr(datasourceName, "ekms_private_endpoints.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "ekms_private_endpoints.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "ekms_private_endpoints.0.display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "ekms_private_endpoints.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "ekms_private_endpoints.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "ekms_private_endpoints.0.subnet_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "ekms_private_endpoints.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "ekms_private_endpoints.0.time_updated"),
			),
		},
		// verify singular datasource
		{
			Config: config + subnetIdVariableStr + caBundleVariableStr + KmsEkmsPrivateEndpointResourceDependencies +
				acctest.GenerateDataSourceFromRepresentationMap("oci_kms_ekms_private_endpoint", "test_ekms_private_endpoint", acctest.Required, acctest.Create, KmsEkmsPrivateEndpointSingularDataSourceRepresentation) +
				compartmentIdVariableStr + KmsEkmsPrivateEndpointResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ekms_private_endpoint_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "external_key_manager_ip", "10.0.0.31"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "port", "443"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "private_endpoint_ip"),
				resource.TestCheckResourceAttr(resourceName, "ca_bundle", strings.Replace(ca_bundle, "\\n", "\n", -1)),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + KmsEkmsPrivateEndpointRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckKmsEkmsPrivateEndpointDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).EkmClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_kms_ekms_private_endpoint" {
			noResourceFound = false
			request := oci_kms.GetEkmsPrivateEndpointRequest{}

			tmp := rs.Primary.ID
			request.EkmsPrivateEndpointId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "kms")

			response, err := client.GetEkmsPrivateEndpoint(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_kms.EkmsPrivateEndpointLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("KmsEkmsPrivateEndpoint") {
		resource.AddTestSweepers("KmsEkmsPrivateEndpoint", &resource.Sweeper{
			Name:         "KmsEkmsPrivateEndpoint",
			Dependencies: acctest.DependencyGraph["ekmsPrivateEndpoint"],
			F:            sweepKmsEkmsPrivateEndpointResource,
		})
	}
}

func sweepKmsEkmsPrivateEndpointResource(compartment string) error {
	ekmClient := acctest.GetTestClients(&schema.ResourceData{}).EkmClient()
	ekmsPrivateEndpointIds, err := getKmsEkmsPrivateEndpointIds(compartment)
	if err != nil {
		return err
	}
	for _, ekmsPrivateEndpointId := range ekmsPrivateEndpointIds {
		if ok := acctest.SweeperDefaultResourceId[ekmsPrivateEndpointId]; !ok {
			deleteEkmsPrivateEndpointRequest := oci_kms.DeleteEkmsPrivateEndpointRequest{}

			deleteEkmsPrivateEndpointRequest.EkmsPrivateEndpointId = &ekmsPrivateEndpointId

			deleteEkmsPrivateEndpointRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "kms")
			_, error := ekmClient.DeleteEkmsPrivateEndpoint(context.Background(), deleteEkmsPrivateEndpointRequest)
			if error != nil {
				fmt.Printf("Error deleting EkmsPrivateEndpoint %s %s, It is possible that the resource is already deleted. Please verify manually \n", ekmsPrivateEndpointId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &ekmsPrivateEndpointId, KmsEkmsPrivateEndpointSweepWaitCondition, time.Duration(3*time.Minute),
				KmsEkmsPrivateEndpointSweepResponseFetchOperation, "kms", true)
		}
	}
	return nil
}

func getKmsEkmsPrivateEndpointIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "EkmsPrivateEndpointId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	ekmClient := acctest.GetTestClients(&schema.ResourceData{}).EkmClient()

	listEkmsPrivateEndpointsRequest := oci_kms.ListEkmsPrivateEndpointsRequest{}
	listEkmsPrivateEndpointsRequest.CompartmentId = &compartmentId
	listEkmsPrivateEndpointsResponse, err := ekmClient.ListEkmsPrivateEndpoints(context.Background(), listEkmsPrivateEndpointsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting EkmsPrivateEndpoint list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, ekmsPrivateEndpoint := range listEkmsPrivateEndpointsResponse.Items {
		id := *ekmsPrivateEndpoint.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "EkmsPrivateEndpointId", id)
	}
	return resourceIds, nil
}

func KmsEkmsPrivateEndpointSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if ekmsPrivateEndpointResponse, ok := response.Response.(oci_kms.GetEkmsPrivateEndpointResponse); ok {
		return ekmsPrivateEndpointResponse.LifecycleState != oci_kms.EkmsPrivateEndpointLifecycleStateDeleted
	}
	return false
}

func KmsEkmsPrivateEndpointSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.EkmClient().GetEkmsPrivateEndpoint(context.Background(), oci_kms.GetEkmsPrivateEndpointRequest{
		EkmsPrivateEndpointId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
