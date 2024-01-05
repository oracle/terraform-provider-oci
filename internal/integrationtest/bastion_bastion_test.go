// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	oci_bastion "github.com/oracle/oci-go-sdk/v65/bastion"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	BastionBastionRequiredOnlyResource = BastionBastionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_bastion_bastion", "test_bastion", acctest.Required, acctest.Create, BastionbastionRepresentation)

	BastionBastionResourceConfig = BastionBastionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_bastion_bastion", "test_bastion", acctest.Optional, acctest.Update, BastionbastionRepresentation)

	BastionBastionbastionSingularDataSourceRepresentation = map[string]interface{}{
		"bastion_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_bastion_bastion.test_bastion.id}`},
	}

	bastionName = utils.RandomString(15, utils.CharsetWithoutDigits)

	BastionBastionbastionDataSourceRepresentation = map[string]interface{}{
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"bastion_id":              acctest.Representation{RepType: acctest.Optional, Create: `${oci_bastion_bastion.test_bastion.id}`},
		"bastion_lifecycle_state": acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"name":                    acctest.Representation{RepType: acctest.Optional, Create: bastionName},
		"filter":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: BastionbastionDataSourceFilterRepresentation}}
	BastionbastionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_bastion_bastion.test_bastion.id}`}},
	}

	BastionbastionRepresentation = map[string]interface{}{
		"bastion_type":                 acctest.Representation{RepType: acctest.Required, Create: `STANDARD`},
		"compartment_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"target_subnet_id":             acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"client_cidr_block_allow_list": acctest.Representation{RepType: acctest.Required, Create: []string{`0.0.0.0/0`}, Update: []string{`0.0.0.0/0`}},
		"defined_tags":                 acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"dns_proxy_status":             acctest.Representation{RepType: acctest.Optional, Create: `DISABLED`},
		"freeform_tags":                acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"max_session_ttl_in_seconds":   acctest.Representation{RepType: acctest.Optional, Create: `1800`, Update: `3600`},
		"name":                         acctest.Representation{RepType: acctest.Required, Create: bastionName},
	}

	BastionBastionResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: bastion/default
func TestBastionBastionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestBastionBastionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_bastion_bastion.test_bastion"
	datasourceName := "data.oci_bastion_bastions.test_bastions"
	singularDatasourceName := "data.oci_bastion_bastion.test_bastion"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+BastionBastionResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_bastion_bastion", "test_bastion", acctest.Optional, acctest.Create, BastionbastionRepresentation), "bastion", "bastion", t)

	acctest.ResourceTest(t, testAccCheckBastionBastionDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + BastionBastionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bastion_bastion", "test_bastion", acctest.Required, acctest.Create, BastionbastionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "bastion_type", "STANDARD"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "target_subnet_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + BastionBastionResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + BastionBastionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bastion_bastion", "test_bastion", acctest.Optional, acctest.Create, BastionbastionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "bastion_type", "STANDARD"),
				resource.TestCheckResourceAttr(resourceName, "client_cidr_block_allow_list.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "dns_proxy_status", "DISABLED"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "max_session_ttl_in_seconds", "1800"),
				resource.TestCheckResourceAttr(resourceName, "name", bastionName),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "target_subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "target_vcn_id"),
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

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + BastionBastionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bastion_bastion", "test_bastion", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(BastionbastionRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "bastion_type", "STANDARD"),
				resource.TestCheckResourceAttr(resourceName, "client_cidr_block_allow_list.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "dns_proxy_status", "DISABLED"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "max_session_ttl_in_seconds", "1800"),
				resource.TestCheckResourceAttr(resourceName, "name", bastionName),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "target_subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "target_vcn_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
			Config: config + compartmentIdVariableStr + BastionBastionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bastion_bastion", "test_bastion", acctest.Optional, acctest.Update, BastionbastionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "bastion_type", "STANDARD"),
				resource.TestCheckResourceAttr(resourceName, "client_cidr_block_allow_list.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "dns_proxy_status", "DISABLED"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "max_session_ttl_in_seconds", "3600"),
				resource.TestCheckResourceAttr(resourceName, "name", bastionName),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "target_subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "target_vcn_id"),
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
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_bastion_bastions", "test_bastions", acctest.Optional, acctest.Update, BastionBastionbastionDataSourceRepresentation) +
				compartmentIdVariableStr + BastionBastionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bastion_bastion", "test_bastion", acctest.Optional, acctest.Update, BastionbastionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "bastion_id"),
				resource.TestCheckResourceAttr(datasourceName, "bastion_lifecycle_state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "name", bastionName),

				resource.TestCheckResourceAttr(datasourceName, "bastions.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "bastions.0.bastion_type", "STANDARD"),
				resource.TestCheckResourceAttr(datasourceName, "bastions.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "bastions.0.dns_proxy_status", "DISABLED"),
				resource.TestCheckResourceAttr(datasourceName, "bastions.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "bastions.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "bastions.0.name", bastionName),
				resource.TestCheckResourceAttrSet(datasourceName, "bastions.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "bastions.0.target_subnet_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "bastions.0.target_vcn_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "bastions.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "bastions.0.time_updated"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_bastion_bastion", "test_bastion", acctest.Required, acctest.Create, BastionBastionbastionSingularDataSourceRepresentation) +
				compartmentIdVariableStr + BastionBastionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "bastion_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "bastion_type", "STANDARD"),
				resource.TestCheckResourceAttr(singularDatasourceName, "client_cidr_block_allow_list.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "dns_proxy_status", "DISABLED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "max_session_ttl_in_seconds", "3600"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "max_sessions_allowed"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", bastionName),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "private_endpoint_ip_address"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_vcn_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + BastionBastionRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckBastionBastionDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).BastionClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_bastion_bastion" {
			noResourceFound = false
			request := oci_bastion.GetBastionRequest{}

			tmp := rs.Primary.ID
			request.BastionId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "bastion")

			response, err := client.GetBastion(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_bastion.BastionLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("BastionBastion") {
		resource.AddTestSweepers("BastionBastion", &resource.Sweeper{
			Name:         "BastionBastion",
			Dependencies: acctest.DependencyGraph["bastion"],
			F:            sweepBastionBastionResource,
		})
	}
}

func sweepBastionBastionResource(compartment string) error {
	bastionClient := acctest.GetTestClients(&schema.ResourceData{}).BastionClient()
	bastionIds, err := getBastionBastionIds(compartment)
	if err != nil {
		return err
	}
	for _, bastionId := range bastionIds {
		if ok := acctest.SweeperDefaultResourceId[bastionId]; !ok {
			deleteBastionRequest := oci_bastion.DeleteBastionRequest{}

			deleteBastionRequest.BastionId = &bastionId

			deleteBastionRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "bastion")
			_, error := bastionClient.DeleteBastion(context.Background(), deleteBastionRequest)
			if error != nil {
				fmt.Printf("Error deleting Bastion %s %s, It is possible that the resource is already deleted. Please verify manually \n", bastionId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &bastionId, BastionbastionsSweepWaitCondition, time.Duration(3*time.Minute),
				BastionbastionsSweepResponseFetchOperation, "bastion", true)
		}
	}
	return nil
}

func getBastionBastionIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "BastionId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	bastionClient := acctest.GetTestClients(&schema.ResourceData{}).BastionClient()

	listBastionsRequest := oci_bastion.ListBastionsRequest{}
	listBastionsRequest.CompartmentId = &compartmentId
	listBastionsResponse, err := bastionClient.ListBastions(context.Background(), listBastionsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Bastion list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, bastion := range listBastionsResponse.Items {
		id := *bastion.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "BastionId", id)
	}
	return resourceIds, nil
}

func BastionbastionsSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if bastionResponse, ok := response.Response.(oci_bastion.GetBastionResponse); ok {
		return bastionResponse.LifecycleState != oci_bastion.BastionLifecycleStateDeleted
	}
	return false
}

func BastionbastionsSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.BastionClient().GetBastion(context.Background(), oci_bastion.GetBastionRequest{
		BastionId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
