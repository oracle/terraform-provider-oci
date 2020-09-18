// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	oci_blockchain "github.com/oracle/oci-go-sdk/v25/blockchain"
	"github.com/oracle/oci-go-sdk/v25/common"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	OsnRequiredOnlyResource = OsnResourceDependencies +
		generateResourceFromRepresentationMap("oci_blockchain_osn", "test_osn", Required, Create, osnRepresentation)

	OsnResourceConfig = OsnResourceDependencies +
		generateResourceFromRepresentationMap("oci_blockchain_osn", "test_osn", Optional, Update, osnRepresentation)

	osnSingularDataSourceRepresentation = map[string]interface{}{
		"blockchain_platform_id": Representation{repType: Required, create: `${oci_blockchain_blockchain_platform.test_blockchain_platform.id}`},
		"osn_id":                 Representation{repType: Required, create: `${oci_blockchain_osn.test_osn.id}`},
	}

	osnDataSourceRepresentation = map[string]interface{}{
		"blockchain_platform_id": Representation{repType: Required, create: `${oci_blockchain_blockchain_platform.test_blockchain_platform.id}`},
		"display_name":           Representation{repType: Optional, create: `displayName`},
		"filter":                 RepresentationGroup{Required, osnDataSourceFilterRepresentation}}
	osnDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `osn_key`},
		"values": Representation{repType: Required, create: []string{`${oci_blockchain_osn.test_osn.id}`}},
	}

	osnRepresentation = map[string]interface{}{
		"ad":                     Representation{repType: Required, create: `AD1`},
		"blockchain_platform_id": Representation{repType: Required, create: `${oci_blockchain_blockchain_platform.test_blockchain_platform.id}`},
		"ocpu_allocation_param":  RepresentationGroup{Optional, osnOcpuAllocationParamRepresentation},
	}
	osnOcpuAllocationParamRepresentation = map[string]interface{}{
		"ocpu_allocation_number": Representation{repType: Required, create: `1.0`, update: `1.1`},
	}

	OsnResourceDependencies = generateResourceFromRepresentationMap("oci_blockchain_blockchain_platform", "test_blockchain_platform", Required, Create, blockchainPlatformRepresentation)
)

func TestBlockchainOsnResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestBlockchainOsnResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	idcsAccessToken := getEnvSettingWithBlankDefault("idcs_access_token")
	idcsAccessTokenVariableStr := fmt.Sprintf("variable \"idcs_access_token\" { default = \"%s\" }\n", idcsAccessToken)

	resourceName := "oci_blockchain_osn.test_osn"
	datasourceName := "data.oci_blockchain_osns.test_osns"
	singularDatasourceName := "data.oci_blockchain_osn.test_osn"

	var resId, resId2, compositeId string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckBlockchainOsnDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + OsnResourceDependencies + idcsAccessTokenVariableStr +
					generateResourceFromRepresentationMap("oci_blockchain_osn", "test_osn", Required, Create, osnRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "ad", "AD1"),
					resource.TestCheckResourceAttrSet(resourceName, "blockchain_platform_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + OsnResourceDependencies + idcsAccessTokenVariableStr,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + OsnResourceDependencies + idcsAccessTokenVariableStr +
					generateResourceFromRepresentationMap("oci_blockchain_osn", "test_osn", Optional, Create, osnRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "ad", "AD1"),
					resource.TestCheckResourceAttrSet(resourceName, "blockchain_platform_id"),
					//resource.TestCheckResourceAttr(resourceName, "ocpu_allocation_param.#", "1"),
					//resource.TestCheckResourceAttr(resourceName, "ocpu_allocation_param.0.ocpu_allocation_number", "1.0"),
					resource.TestCheckResourceAttrSet(resourceName, "osn_key"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						blockchainPlatformId, _ := fromInstanceState(s, resourceName, "blockchain_platform_id")
						compositeId = "blockchainPlatforms/" + blockchainPlatformId + "/osns/" + resId
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&compositeId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + OsnResourceDependencies + idcsAccessTokenVariableStr +
					generateResourceFromRepresentationMap("oci_blockchain_osn", "test_osn", Optional, Update, osnRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "ad", "AD1"),
					resource.TestCheckResourceAttrSet(resourceName, "blockchain_platform_id"),
					//resource.TestCheckResourceAttr(resourceName, "ocpu_allocation_param.#", "1"),
					//resource.TestCheckResourceAttr(resourceName, "ocpu_allocation_param.0.ocpu_allocation_number", "1.1"),
					resource.TestCheckResourceAttrSet(resourceName, "osn_key"),

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
					generateDataSourceFromRepresentationMap("oci_blockchain_osns", "test_osns", Optional, Update, osnDataSourceRepresentation) +
					compartmentIdVariableStr + OsnResourceDependencies + idcsAccessTokenVariableStr +
					generateResourceFromRepresentationMap("oci_blockchain_osn", "test_osn", Optional, Update, osnRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "blockchain_platform_id"),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName"),

					resource.TestCheckResourceAttr(datasourceName, "osn_collection.#", "1"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_blockchain_osn", "test_osn", Required, Create, osnSingularDataSourceRepresentation) +
					compartmentIdVariableStr + idcsAccessTokenVariableStr + OsnResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "blockchain_platform_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "osn_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "ad", "AD1"),
					//resource.TestCheckResourceAttr(singularDatasourceName, "ocpu_allocation_param.#", "1"),
					//resource.TestCheckResourceAttr(singularDatasourceName, "ocpu_allocation_param.0.ocpu_allocation_number", "1.1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "osn_key"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr + OsnResourceConfig,
			},
			// verify resource import
			{
				Config:                  config + generateResourceImportConfig("oci_blockchain_osn", "test_osn"),
				ImportState:             true,
				ImportStateIdFunc:       getBlockchainOsnCompositeId(resourceName),
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
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
	client := testAccProvider.Meta().(*OracleClients).blockchainPlatformClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_blockchain_osn" {
			noResourceFound = false
			request := oci_blockchain.GetOsnRequest{}

			if value, ok := rs.Primary.Attributes["blockchain_platform_id"]; ok {
				request.BlockchainPlatformId = &value
			}

			tmp := rs.Primary.ID
			request.OsnId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "blockchain")

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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("BlockchainOsn") {
		resource.AddTestSweepers("BlockchainOsn", &resource.Sweeper{
			Name:         "BlockchainOsn",
			Dependencies: DependencyGraph["osn"],
			F:            sweepBlockchainOsnResource,
		})
	}
}

func sweepBlockchainOsnResource(compartment string) error {
	blockchainPlatformClient := GetTestClients(&schema.ResourceData{}).blockchainPlatformClient()
	osnIds, err := getOsnIds(compartment)
	if err != nil {
		return err
	}
	for _, osnId := range osnIds {
		if ok := SweeperDefaultResourceId[osnId]; !ok {
			deleteOsnRequest := oci_blockchain.DeleteOsnRequest{}

			deleteOsnRequest.OsnId = &osnId

			deleteOsnRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "blockchain")
			_, error := blockchainPlatformClient.DeleteOsn(context.Background(), deleteOsnRequest)
			if error != nil {
				fmt.Printf("Error deleting Osn %s %s, It is possible that the resource is already deleted. Please verify manually \n", osnId, error)
				continue
			}
		}
	}
	return nil
}

func getOsnIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "OsnId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	blockchainPlatformClient := GetTestClients(&schema.ResourceData{}).blockchainPlatformClient()

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
			addResourceIdToSweeperResourceIdMap(compartmentId, "OsnId", id)
		}

	}
	return resourceIds, nil
}
