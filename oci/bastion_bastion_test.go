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
	oci_bastion "github.com/oracle/oci-go-sdk/v52/bastion"
	"github.com/oracle/oci-go-sdk/v52/common"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	BastionRequiredOnlyResource = BastionResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_bastion_bastion", "test_bastion", Required, Create, bastionRepresentation)

	BastionResourceConfig = BastionResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_bastion_bastion", "test_bastion", Optional, Update, bastionRepresentation)

	bastionSingularDataSourceRepresentation = map[string]interface{}{
		"bastion_id": Representation{RepType: Required, Create: `${oci_bastion_bastion.test_bastion.id}`},
	}

	bastionDataSourceRepresentation = map[string]interface{}{
		"compartment_id":          Representation{RepType: Required, Create: `${var.compartment_id}`},
		"bastion_id":              Representation{RepType: Optional, Create: `${oci_bastion_bastion.test_bastion.id}`},
		"bastion_lifecycle_state": Representation{RepType: Optional, Create: `ACTIVE`},
		"name":                    Representation{RepType: Optional, Create: `bastionterraformtest`},
		"filter":                  RepresentationGroup{Required, bastionDataSourceFilterRepresentation}}
	bastionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_bastion_bastion.test_bastion.id}`}},
	}

	bastionRepresentation = map[string]interface{}{
		"bastion_type":                 Representation{RepType: Required, Create: `STANDARD`},
		"compartment_id":               Representation{RepType: Required, Create: `${var.compartment_id}`},
		"target_subnet_id":             Representation{RepType: Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"client_cidr_block_allow_list": Representation{RepType: Required, Create: []string{`0.0.0.0/0`}, Update: []string{`0.0.0.0/0`}},
		"defined_tags":                 Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":                Representation{RepType: Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"max_session_ttl_in_seconds":   Representation{RepType: Optional, Create: `1800`, Update: `3600`},
		"name":                         Representation{RepType: Required, Create: `bastionterraformtest`, Update: "BastionTerraformtest"},
	}

	BastionResourceDependencies = GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, subnetRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, vcnRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: bastion/default
func TestBastionBastionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestBastionBastionResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_bastion_bastion.test_bastion"
	datasourceName := "data.oci_bastion_bastions.test_bastions"
	singularDatasourceName := "data.oci_bastion_bastion.test_bastion"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+BastionResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_bastion_bastion", "test_bastion", Optional, Create, bastionRepresentation), "bastion", "bastion", t)

	ResourceTest(t, testAccCheckBastionBastionDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + BastionResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_bastion_bastion", "test_bastion", Required, Create, bastionRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "bastion_type", "STANDARD"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "target_subnet_id"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + BastionResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + BastionResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_bastion_bastion", "test_bastion", Optional, Create, bastionRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "bastion_type", "STANDARD"),
				resource.TestCheckResourceAttr(resourceName, "client_cidr_block_allow_list.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "max_session_ttl_in_seconds", "1800"),
				resource.TestCheckResourceAttr(resourceName, "name", "bastionterraformtest"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "target_subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "target_vcn_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + BastionResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_bastion_bastion", "test_bastion", Optional, Create,
					RepresentationCopyWithNewProperties(bastionRepresentation, map[string]interface{}{
						"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "bastion_type", "STANDARD"),
				resource.TestCheckResourceAttr(resourceName, "client_cidr_block_allow_list.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "max_session_ttl_in_seconds", "1800"),
				resource.TestCheckResourceAttr(resourceName, "name", "bastionterraformtest"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "target_subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "target_vcn_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + BastionResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_bastion_bastion", "test_bastion", Optional, Update, bastionRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "bastion_type", "STANDARD"),
				resource.TestCheckResourceAttr(resourceName, "client_cidr_block_allow_list.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "max_session_ttl_in_seconds", "3600"),
				resource.TestCheckResourceAttr(resourceName, "name", "bastionterraformtest"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "target_subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "target_vcn_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
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
				GenerateDataSourceFromRepresentationMap("oci_bastion_bastions", "test_bastions", Optional, Update, bastionDataSourceRepresentation) +
				compartmentIdVariableStr + BastionResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_bastion_bastion", "test_bastion", Optional, Update, bastionRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "bastion_id"),
				resource.TestCheckResourceAttr(datasourceName, "bastion_lifecycle_state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "name", "bastionterraformtest"),

				resource.TestCheckResourceAttr(datasourceName, "bastions.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "bastions.0.bastion_type", "STANDARD"),
				resource.TestCheckResourceAttr(datasourceName, "bastions.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "bastions.0.defined_tags.%", "1"),
				resource.TestCheckResourceAttr(datasourceName, "bastions.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "bastions.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "bastions.0.name", "bastionterraformtest"),
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
				GenerateDataSourceFromRepresentationMap("oci_bastion_bastion", "test_bastion", Required, Create, bastionSingularDataSourceRepresentation) +
				compartmentIdVariableStr + BastionResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "bastion_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "bastion_type", "STANDARD"),
				resource.TestCheckResourceAttr(singularDatasourceName, "client_cidr_block_allow_list.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "max_session_ttl_in_seconds", "3600"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "max_sessions_allowed"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "bastionterraformtest"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "private_endpoint_ip_address"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_vcn_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + BastionResourceConfig,
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

func testAccCheckBastionBastionDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).bastionClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_bastion_bastion" {
			noResourceFound = false
			request := oci_bastion.GetBastionRequest{}

			tmp := rs.Primary.ID
			request.BastionId = &tmp

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "bastion")

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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !InSweeperExcludeList("BastionBastion") {
		resource.AddTestSweepers("BastionBastion", &resource.Sweeper{
			Name:         "BastionBastion",
			Dependencies: DependencyGraph["bastion"],
			F:            sweepBastionBastionResource,
		})
	}
}

func sweepBastionBastionResource(compartment string) error {
	bastionClient := GetTestClients(&schema.ResourceData{}).bastionClient()
	bastionIds, err := getBastionIds(compartment)
	if err != nil {
		return err
	}
	for _, bastionId := range bastionIds {
		if ok := SweeperDefaultResourceId[bastionId]; !ok {
			deleteBastionRequest := oci_bastion.DeleteBastionRequest{}

			deleteBastionRequest.BastionId = &bastionId

			deleteBastionRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "bastion")
			_, error := bastionClient.DeleteBastion(context.Background(), deleteBastionRequest)
			if error != nil {
				fmt.Printf("Error deleting Bastion %s %s, It is possible that the resource is already deleted. Please verify manually \n", bastionId, error)
				continue
			}
			WaitTillCondition(testAccProvider, &bastionId, bastionSweepWaitCondition, time.Duration(3*time.Minute),
				bastionSweepResponseFetchOperation, "bastion", true)
		}
	}
	return nil
}

func getBastionIds(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "BastionId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	bastionClient := GetTestClients(&schema.ResourceData{}).bastionClient()

	listBastionsRequest := oci_bastion.ListBastionsRequest{}
	listBastionsRequest.CompartmentId = &compartmentId
	listBastionsResponse, err := bastionClient.ListBastions(context.Background(), listBastionsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Bastion list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, bastion := range listBastionsResponse.Items {
		id := *bastion.Id
		resourceIds = append(resourceIds, id)
		AddResourceIdToSweeperResourceIdMap(compartmentId, "BastionId", id)
	}
	return resourceIds, nil
}

func bastionSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if bastionResponse, ok := response.Response.(oci_bastion.GetBastionResponse); ok {
		return bastionResponse.LifecycleState != oci_bastion.BastionLifecycleStateDeleted
	}
	return false
}

func bastionSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.bastionClient().GetBastion(context.Background(), oci_bastion.GetBastionRequest{
		BastionId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
