// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	oci_blockchain "github.com/oracle/oci-go-sdk/v65/blockchain"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	BlockchainOsnRequiredOnlyResource = BlockchainOsnResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_blockchain_osn", "test_osn", acctest.Required, acctest.Create, osnRepresentation)

	BlockchainOsnResourceConfig = BlockchainOsnResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_blockchain_osn", "test_osn", acctest.Optional, acctest.Update, osnRepresentation)

	BlockchainosnSingularDataSourceRepresentation = map[string]interface{}{
		"blockchain_platform_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_blockchain_blockchain_platform.test_blockchain_platform.id}`},
		"osn_id":                 acctest.Representation{RepType: acctest.Required, Create: `${oci_blockchain_osn.test_osn.id}`},
	}

	BlockchainosnDataSourceRepresentation = map[string]interface{}{
		"blockchain_platform_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_blockchain_blockchain_platform.test_blockchain_platform.id}`},
		"display_name":           acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"filter":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: osnDataSourceFilterRepresentation}}
	osnDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `osn_key`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_blockchain_osn.test_osn.id}`}},
	}

	osnRepresentation = map[string]interface{}{
		"ad":                     acctest.Representation{RepType: acctest.Required, Create: `AD1`},
		"blockchain_platform_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_blockchain_blockchain_platform.test_blockchain_platform.id}`},
		"ocpu_allocation_param":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: osnOcpuAllocationParamRepresentation},
	}
	osnOcpuAllocationParamRepresentation = map[string]interface{}{
		"ocpu_allocation_number": acctest.Representation{RepType: acctest.Required, Create: `0.0`, Update: `0.0`},
	}

	BlockchainOsnResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_blockchain_blockchain_platform", "test_blockchain_platform", acctest.Required, acctest.Create, blockchainPlatformRepresentation)
)

// issue-routing-tag: blockchain/default
func TestBlockchainOsnResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestBlockchainOsnResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	idcsAccessToken := utils.GetEnvSettingWithBlankDefault("idcs_access_token")
	idcsAccessTokenVariableStr := fmt.Sprintf("variable \"idcs_access_token\" { default = \"%s\" }\n", idcsAccessToken)

	resourceName := "oci_blockchain_osn.test_osn"
	datasourceName := "data.oci_blockchain_osns.test_osns"
	singularDatasourceName := "data.oci_blockchain_osn.test_osn"

	var resId, resId2, compositeId string

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+BlockchainOsnResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_blockchain_osn", "test_osn", acctest.Optional, acctest.Create, osnRepresentation), "blockchain", "osn", t)

	acctest.ResourceTest(t, testAccCheckBlockchainOsnDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + BlockchainOsnResourceDependencies + idcsAccessTokenVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_blockchain_osn", "test_osn", acctest.Required, acctest.Create, osnRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "ad", "AD1"),
				resource.TestCheckResourceAttrSet(resourceName, "blockchain_platform_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + BlockchainOsnResourceDependencies + idcsAccessTokenVariableStr,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + BlockchainOsnResourceDependencies + idcsAccessTokenVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_blockchain_osn", "test_osn", acctest.Optional, acctest.Create, osnRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "ad", "AD1"),
				resource.TestCheckResourceAttrSet(resourceName, "blockchain_platform_id"),
				//resource.TestCheckResourceAttr(resourceName, "ocpu_allocation_param.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "ocpu_allocation_param.0.ocpu_allocation_number", "1.0"),
				resource.TestCheckResourceAttrSet(resourceName, "osn_key"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					blockchainPlatformId, _ := acctest.FromInstanceState(s, resourceName, "blockchain_platform_id")
					compositeId = "blockchainPlatforms/" + blockchainPlatformId + "/osns/" + resId
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&compositeId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + BlockchainOsnResourceDependencies + idcsAccessTokenVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_blockchain_osn", "test_osn", acctest.Optional, acctest.Update, osnRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "ad", "AD1"),
				resource.TestCheckResourceAttrSet(resourceName, "blockchain_platform_id"),
				//resource.TestCheckResourceAttr(resourceName, "ocpu_allocation_param.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "ocpu_allocation_param.0.ocpu_allocation_number", "1.1"),
				resource.TestCheckResourceAttrSet(resourceName, "osn_key"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_blockchain_osns", "test_osns", acctest.Optional, acctest.Update, BlockchainosnDataSourceRepresentation) +
				compartmentIdVariableStr + BlockchainOsnResourceDependencies + idcsAccessTokenVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_blockchain_osn", "test_osn", acctest.Optional, acctest.Update, osnRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "blockchain_platform_id"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName"),

				resource.TestCheckResourceAttr(datasourceName, "osn_collection.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_blockchain_osn", "test_osn", acctest.Required, acctest.Create, BlockchainosnSingularDataSourceRepresentation) +
				compartmentIdVariableStr + idcsAccessTokenVariableStr + BlockchainOsnResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "blockchain_platform_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "osn_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "ad", "AD1"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "ocpu_allocation_param.#", "1"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "ocpu_allocation_param.0.ocpu_allocation_number", "1.1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "osn_key"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
			),
		},
		// verify resource import
		{
			Config:                  config + BlockchainOsnRequiredOnlyResource,
			ImportState:             true,
			ImportStateIdFunc:       getBlockchainOsnCompositeId(resourceName),
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func getBlockchainOsnCompositeId(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("not found: %s", resourceName)
		}

		return fmt.Sprintf("blockchainPlatforms/%s/osns/%s", rs.Primary.Attributes["blockchain_platform_id"], rs.Primary.Attributes["id"]), nil
	}
}

func testAccCheckBlockchainOsnDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).BlockchainPlatformClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_blockchain_osn" {
			noResourceFound = false
			request := oci_blockchain.GetOsnRequest{}

			if value, ok := rs.Primary.Attributes["blockchain_platform_id"]; ok {
				request.BlockchainPlatformId = &value
			}

			tmp := rs.Primary.ID
			request.OsnId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "blockchain")

			_, err := client.GetOsn(context.Background(), request)

			if err == nil {
				return fmt.Errorf("resource still exists")
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
	if !acctest.InSweeperExcludeList("BlockchainOsn") {
		resource.AddTestSweepers("BlockchainOsn", &resource.Sweeper{
			Name:         "BlockchainOsn",
			Dependencies: acctest.DependencyGraph["osn"],
			F:            sweepBlockchainOsnResource,
		})
	}
}

func sweepBlockchainOsnResource(compartment string) error {
	blockchainPlatformClient := acctest.GetTestClients(&schema.ResourceData{}).BlockchainPlatformClient()
	osnIds, err := getBlockchainOsnIds(compartment)
	if err != nil {
		return err
	}
	for _, osnId := range osnIds {
		if ok := acctest.SweeperDefaultResourceId[osnId]; !ok {
			deleteOsnRequest := oci_blockchain.DeleteOsnRequest{}

			deleteOsnRequest.OsnId = &osnId

			deleteOsnRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "blockchain")
			_, error := blockchainPlatformClient.DeleteOsn(context.Background(), deleteOsnRequest)
			if error != nil {
				fmt.Printf("Error deleting Osn %s %s, It is possible that the resource is already deleted. Please verify manually \n", osnId, error)
				continue
			}
		}
	}
	return nil
}

func getBlockchainOsnIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "OsnId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	blockchainPlatformClient := acctest.GetTestClients(&schema.ResourceData{}).BlockchainPlatformClient()

	listOsnsRequest := oci_blockchain.ListOsnsRequest{}

	blockchainPlatformIds, error := getBlockchainPlatformIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting blockchainPlatformId required for Osn resource requests \n")
	}
	for _, blockchainPlatformId := range blockchainPlatformIds {
		listOsnsRequest.BlockchainPlatformId = &blockchainPlatformId

		listOsnsResponse, err := blockchainPlatformClient.ListOsns(context.Background(), listOsnsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting Osn list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, osn := range listOsnsResponse.Items {
			id := *osn.OsnKey
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "OsnId", id)
		}

	}
	return resourceIds, nil
}
