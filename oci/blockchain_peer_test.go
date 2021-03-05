// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
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
	oci_blockchain "github.com/oracle/oci-go-sdk/v36/blockchain"
	"github.com/oracle/oci-go-sdk/v36/common"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	PeerRequiredOnlyResource = PeerResourceDependencies +
		generateResourceFromRepresentationMap("oci_blockchain_peer", "test_peer", Required, Create, peerRepresentation)

	PeerResourceConfig = PeerResourceDependencies +
		generateResourceFromRepresentationMap("oci_blockchain_peer", "test_peer", Optional, Update, peerRepresentation)

	peerSingularDataSourceRepresentation = map[string]interface{}{
		"blockchain_platform_id": Representation{repType: Required, create: `${oci_blockchain_blockchain_platform.test_blockchain_platform.id}`},
		"peer_id":                Representation{repType: Required, create: `${oci_blockchain_peer.test_peer.id}`},
	}

	peerDataSourceRepresentation = map[string]interface{}{
		"blockchain_platform_id": Representation{repType: Required, create: `${oci_blockchain_blockchain_platform.test_blockchain_platform.id}`},
		"display_name":           Representation{repType: Optional, create: `displayName`},
		"filter":                 RepresentationGroup{Required, peerDataSourceFilterRepresentation}}
	peerDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `peer_key`},
		"values": Representation{repType: Required, create: []string{`${oci_blockchain_peer.test_peer.id}`}},
	}

	peerRepresentation = map[string]interface{}{
		"ad":                     Representation{repType: Required, create: `AD1`},
		"blockchain_platform_id": Representation{repType: Required, create: `${oci_blockchain_blockchain_platform.test_blockchain_platform.id}`},
		"ocpu_allocation_param":  RepresentationGroup{Required, peerOcpuAllocationParamRepresentation},
		"role":                   Representation{repType: Required, create: `MEMBER`},
		"alias":                  Representation{repType: Optional, create: `alias`},
	}
	peerOcpuAllocationParamRepresentation = map[string]interface{}{
		"ocpu_allocation_number": Representation{repType: Required, create: `0.5`, update: `0.6`},
	}

	PeerResourceDependencies = generateResourceFromRepresentationMap("oci_blockchain_blockchain_platform", "test_blockchain_platform", Required, Create, blockchainPlatformRepresentation)
)

func TestBlockchainPeerResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestBlockchainPeerResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	idcsAccessToken := getEnvSettingWithBlankDefault("idcs_access_token")
	idcsAccessTokenVariableStr := fmt.Sprintf("variable \"idcs_access_token\" { default = \"%s\" }\n", idcsAccessToken)

	resourceName := "oci_blockchain_peer.test_peer"
	datasourceName := "data.oci_blockchain_peers.test_peers"
	singularDatasourceName := "data.oci_blockchain_peer.test_peer"

	var resId, resId2, compositeId string

	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+PeerResourceDependencies+
		generateResourceFromRepresentationMap("oci_blockchain_peer", "test_peer", Optional, Create, peerRepresentation), "blockchain", "peer", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckBlockchainPeerDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + PeerResourceDependencies + idcsAccessTokenVariableStr +
					generateResourceFromRepresentationMap("oci_blockchain_peer", "test_peer", Required, Create, peerRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "ad", "AD1"),
					resource.TestCheckResourceAttrSet(resourceName, "blockchain_platform_id"),
					resource.TestCheckResourceAttr(resourceName, "ocpu_allocation_param.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ocpu_allocation_param.0.ocpu_allocation_number", "0.5"),
					resource.TestCheckResourceAttr(resourceName, "role", "MEMBER"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + PeerResourceDependencies + idcsAccessTokenVariableStr,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + PeerResourceDependencies + idcsAccessTokenVariableStr +
					generateResourceFromRepresentationMap("oci_blockchain_peer", "test_peer", Optional, Create, peerRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "ad", "AD1"),
					resource.TestCheckResourceAttr(resourceName, "alias", "alias"),
					resource.TestCheckResourceAttrSet(resourceName, "blockchain_platform_id"),
					resource.TestCheckResourceAttrSet(resourceName, "host"),
					resource.TestCheckResourceAttr(resourceName, "ocpu_allocation_param.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ocpu_allocation_param.0.ocpu_allocation_number", "0.5"),
					resource.TestCheckResourceAttrSet(resourceName, "peer_key"),
					resource.TestCheckResourceAttr(resourceName, "role", "MEMBER"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						blockchainPlatformId, _ := fromInstanceState(s, resourceName, "blockchain_platform_id")
						compositeId = "blockchainPlatforms/" + blockchainPlatformId + "/peers/" + resId
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
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
				Config: config + compartmentIdVariableStr + PeerResourceDependencies + idcsAccessTokenVariableStr +
					generateResourceFromRepresentationMap("oci_blockchain_peer", "test_peer", Optional, Update, peerRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "ad", "AD1"),
					resource.TestCheckResourceAttr(resourceName, "alias", "alias"),
					resource.TestCheckResourceAttrSet(resourceName, "blockchain_platform_id"),
					resource.TestCheckResourceAttrSet(resourceName, "host"),
					resource.TestCheckResourceAttr(resourceName, "ocpu_allocation_param.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ocpu_allocation_param.0.ocpu_allocation_number", "0.6"),
					resource.TestCheckResourceAttrSet(resourceName, "peer_key"),
					resource.TestCheckResourceAttr(resourceName, "role", "MEMBER"),

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
					generateDataSourceFromRepresentationMap("oci_blockchain_peers", "test_peers", Optional, Update, peerDataSourceRepresentation) +
					compartmentIdVariableStr + PeerResourceDependencies + idcsAccessTokenVariableStr +
					generateResourceFromRepresentationMap("oci_blockchain_peer", "test_peer", Optional, Update, peerRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "blockchain_platform_id"),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_blockchain_peer", "test_peer", Required, Create, peerSingularDataSourceRepresentation) +
					compartmentIdVariableStr + idcsAccessTokenVariableStr + PeerResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
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
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + PeerResourceConfig + idcsAccessTokenVariableStr,
			},
			// verify resource import
			{
				Config:                  config,
				ImportState:             true,
				ImportStateIdFunc:       getBlockchainPeerCompositeId(resourceName),
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
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
	client := testAccProvider.Meta().(*OracleClients).blockchainPlatformClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_blockchain_peer" {
			noResourceFound = false
			request := oci_blockchain.GetPeerRequest{}

			if value, ok := rs.Primary.Attributes["blockchain_platform_id"]; ok {
				request.BlockchainPlatformId = &value
			}

			tmp := rs.Primary.ID
			request.PeerId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "blockchain")

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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("BlockchainPeer") {
		resource.AddTestSweepers("BlockchainPeer", &resource.Sweeper{
			Name:         "BlockchainPeer",
			Dependencies: DependencyGraph["peer"],
			F:            sweepBlockchainPeerResource,
		})
	}
}

func sweepBlockchainPeerResource(compartment string) error {
	blockchainPlatformClient := GetTestClients(&schema.ResourceData{}).blockchainPlatformClient()
	peerIds, err := getPeerIds(compartment)
	if err != nil {
		return err
	}
	for _, peerId := range peerIds {
		if ok := SweeperDefaultResourceId[peerId]; !ok {
			deletePeerRequest := oci_blockchain.DeletePeerRequest{}

			deletePeerRequest.PeerId = &peerId

			deletePeerRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "blockchain")
			_, error := blockchainPlatformClient.DeletePeer(context.Background(), deletePeerRequest)
			if error != nil {
				fmt.Printf("Error deleting Peer %s %s, It is possible that the resource is already deleted. Please verify manually \n", peerId, error)
				continue
			}
		}
	}
	return nil
}

func getPeerIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "PeerId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	blockchainPlatformClient := GetTestClients(&schema.ResourceData{}).blockchainPlatformClient()

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
			addResourceIdToSweeperResourceIdMap(compartmentId, "PeerId", id)
		}

	}
	return resourceIds, nil
}
