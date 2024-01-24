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
	BlockchainPeerRequiredOnlyResource = BlockchainPeerResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_blockchain_peer", "test_peer", acctest.Required, acctest.Create, peerRepresentation)

	BlockchainPeerResourceConfig = BlockchainPeerResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_blockchain_peer", "test_peer", acctest.Optional, acctest.Update, peerRepresentation)

	BlockchainpeerSingularDataSourceRepresentation = map[string]interface{}{
		"blockchain_platform_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_blockchain_blockchain_platform.test_blockchain_platform.id}`},
		"peer_id":                acctest.Representation{RepType: acctest.Required, Create: `${oci_blockchain_peer.test_peer.id}`},
	}

	BlockchainpeerDataSourceRepresentation = map[string]interface{}{
		"blockchain_platform_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_blockchain_blockchain_platform.test_blockchain_platform.id}`},
		"display_name":           acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"filter":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: peerDataSourceFilterRepresentation}}
	peerDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `peer_key`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_blockchain_peer.test_peer.id}`}},
	}

	peerRepresentation = map[string]interface{}{
		"ad":                     acctest.Representation{RepType: acctest.Required, Create: `AD1`},
		"blockchain_platform_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_blockchain_blockchain_platform.test_blockchain_platform.id}`},
		"ocpu_allocation_param":  acctest.RepresentationGroup{RepType: acctest.Required, Group: peerOcpuAllocationParamRepresentation},
		"role":                   acctest.Representation{RepType: acctest.Required, Create: `MEMBER`},
		"alias":                  acctest.Representation{RepType: acctest.Optional, Create: `alias`},
	}
	peerOcpuAllocationParamRepresentation = map[string]interface{}{
		"ocpu_allocation_number": acctest.Representation{RepType: acctest.Required, Create: `0.5`, Update: `0.6`},
	}

	BlockchainPeerResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_blockchain_blockchain_platform", "test_blockchain_platform", acctest.Required, acctest.Create, blockchainPlatformRepresentation)
)

// issue-routing-tag: blockchain/default
func TestBlockchainPeerResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestBlockchainPeerResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	idcsAccessToken := utils.GetEnvSettingWithBlankDefault("idcs_access_token")
	idcsAccessTokenVariableStr := fmt.Sprintf("variable \"idcs_access_token\" { default = \"%s\" }\n", idcsAccessToken)

	resourceName := "oci_blockchain_peer.test_peer"
	datasourceName := "data.oci_blockchain_peers.test_peers"
	singularDatasourceName := "data.oci_blockchain_peer.test_peer"

	var resId, resId2, compositeId string

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+BlockchainPeerResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_blockchain_peer", "test_peer", acctest.Optional, acctest.Create, peerRepresentation), "blockchain", "peer", t)

	acctest.ResourceTest(t, testAccCheckBlockchainPeerDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + BlockchainPeerResourceDependencies + idcsAccessTokenVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_blockchain_peer", "test_peer", acctest.Required, acctest.Create, peerRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "ad", "AD1"),
				resource.TestCheckResourceAttrSet(resourceName, "blockchain_platform_id"),
				resource.TestCheckResourceAttr(resourceName, "ocpu_allocation_param.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ocpu_allocation_param.0.ocpu_allocation_number", "0.5"),
				resource.TestCheckResourceAttr(resourceName, "role", "MEMBER"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + BlockchainPeerResourceDependencies + idcsAccessTokenVariableStr,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + BlockchainPeerResourceDependencies + idcsAccessTokenVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_blockchain_peer", "test_peer", acctest.Optional, acctest.Create, peerRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "ad", "AD1"),
				resource.TestCheckResourceAttr(resourceName, "alias", "alias"),
				resource.TestCheckResourceAttrSet(resourceName, "blockchain_platform_id"),
				resource.TestCheckResourceAttrSet(resourceName, "host"),
				resource.TestCheckResourceAttr(resourceName, "ocpu_allocation_param.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ocpu_allocation_param.0.ocpu_allocation_number", "0.5"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_key"),
				resource.TestCheckResourceAttr(resourceName, "role", "MEMBER"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					blockchainPlatformId, _ := acctest.FromInstanceState(s, resourceName, "blockchain_platform_id")
					compositeId = "blockchainPlatforms/" + blockchainPlatformId + "/peers/" + resId
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
			Config: config + compartmentIdVariableStr + BlockchainPeerResourceDependencies + idcsAccessTokenVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_blockchain_peer", "test_peer", acctest.Optional, acctest.Update, peerRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "ad", "AD1"),
				resource.TestCheckResourceAttr(resourceName, "alias", "alias"),
				resource.TestCheckResourceAttrSet(resourceName, "blockchain_platform_id"),
				resource.TestCheckResourceAttrSet(resourceName, "host"),
				resource.TestCheckResourceAttr(resourceName, "ocpu_allocation_param.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ocpu_allocation_param.0.ocpu_allocation_number", "0.6"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_key"),
				resource.TestCheckResourceAttr(resourceName, "role", "MEMBER"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_blockchain_peers", "test_peers", acctest.Optional, acctest.Update, BlockchainpeerDataSourceRepresentation) +
				compartmentIdVariableStr + BlockchainPeerResourceDependencies + idcsAccessTokenVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_blockchain_peer", "test_peer", acctest.Optional, acctest.Update, peerRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "blockchain_platform_id"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_blockchain_peer", "test_peer", acctest.Required, acctest.Create, BlockchainpeerSingularDataSourceRepresentation) +
				compartmentIdVariableStr + idcsAccessTokenVariableStr + BlockchainPeerResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "blockchain_platform_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "peer_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "ad", "AD1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "alias", "alias"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "host"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ocpu_allocation_param.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ocpu_allocation_param.0.ocpu_allocation_number", "0.6"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "peer_key"),
				resource.TestCheckResourceAttr(singularDatasourceName, "role", "MEMBER"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
			),
		},
		// verify resource import
		{
			Config:                  config + BlockchainPeerRequiredOnlyResource,
			ImportState:             true,
			ImportStateIdFunc:       getBlockchainPeerCompositeId(resourceName),
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func getBlockchainPeerCompositeId(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("not found: %s", resourceName)
		}

		return fmt.Sprintf("blockchainPlatforms/%s/peers/%s", rs.Primary.Attributes["blockchain_platform_id"], rs.Primary.Attributes["id"]), nil
	}
}

func testAccCheckBlockchainPeerDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).BlockchainPlatformClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_blockchain_peer" {
			noResourceFound = false
			request := oci_blockchain.GetPeerRequest{}

			if value, ok := rs.Primary.Attributes["blockchain_platform_id"]; ok {
				request.BlockchainPlatformId = &value
			}

			tmp := rs.Primary.ID
			request.PeerId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "blockchain")

			_, err := client.GetPeer(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("BlockchainPeer") {
		resource.AddTestSweepers("BlockchainPeer", &resource.Sweeper{
			Name:         "BlockchainPeer",
			Dependencies: acctest.DependencyGraph["peer"],
			F:            sweepBlockchainPeerResource,
		})
	}
}

func sweepBlockchainPeerResource(compartment string) error {
	blockchainPlatformClient := acctest.GetTestClients(&schema.ResourceData{}).BlockchainPlatformClient()
	peerIds, err := getBlockchainPeerIds(compartment)
	if err != nil {
		return err
	}
	for _, peerId := range peerIds {
		if ok := acctest.SweeperDefaultResourceId[peerId]; !ok {
			deletePeerRequest := oci_blockchain.DeletePeerRequest{}

			deletePeerRequest.PeerId = &peerId

			deletePeerRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "blockchain")
			_, error := blockchainPlatformClient.DeletePeer(context.Background(), deletePeerRequest)
			if error != nil {
				fmt.Printf("Error deleting Peer %s %s, It is possible that the resource is already deleted. Please verify manually \n", peerId, error)
				continue
			}
		}
	}
	return nil
}

func getBlockchainPeerIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "PeerId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	blockchainPlatformClient := acctest.GetTestClients(&schema.ResourceData{}).BlockchainPlatformClient()

	listPeersRequest := oci_blockchain.ListPeersRequest{}

	blockchainPlatformIds, error := getBlockchainPlatformIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting blockchainPlatformId required for Peer resource requests \n")
	}
	for _, blockchainPlatformId := range blockchainPlatformIds {
		listPeersRequest.BlockchainPlatformId = &blockchainPlatformId

		listPeersResponse, err := blockchainPlatformClient.ListPeers(context.Background(), listPeersRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting Peer list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, peer := range listPeersResponse.Items {
			id := *peer.PeerKey
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "PeerId", id)
		}

	}
	return resourceIds, nil
}
