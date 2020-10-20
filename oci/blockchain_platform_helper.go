// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	oci_blockchain "github.com/oracle/oci-go-sdk/v27/blockchain"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func blockchainPlatformComputeShapeDiffSuppressFunction(key string, old string, new string, d *schema.ResourceData) bool {
	if old == "" || new == "" {
		return false
	}
	upperCaseNewValue := strings.ToUpper(new)
	upperCaseOldValue := strings.ToUpper(old)

	// ENTERPRISE_CUSTOM is auto changed when update totalOcpuCapacity
	if upperCaseNewValue == "ENTERPRISE_CUSTOM" || upperCaseOldValue == "ENTERPRISE_CUSTOM" {
		return true
	}

	return upperCaseOldValue == upperCaseNewValue
}

func sendUpdateBlockchainPlatformRequest(s *BlockchainBlockchainPlatformResourceCrud, request oci_blockchain.UpdateBlockchainPlatformRequest) error {
	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "blockchain")

	response, err := s.Client.UpdateBlockchainPlatform(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	err = s.getBlockchainPlatformFromWorkRequest(workId, getRetryPolicy(s.DisableNotFoundRetries, "blockchain"), oci_blockchain.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
	if err != nil {
		return err
	}
	return nil
}

func blockchainPlatformPeerOcpuAllocationNumberDiffSuppressFunction(key string, old string, new string, d *schema.ResourceData) bool {
	if old == "" || new == "" {
		return false
	}
	oldFloat, err := strconv.ParseFloat(old, 32)
	if err != nil {
		return false
	}
	newFloat, err := strconv.ParseFloat(new, 32)
	if err != nil {
		return false
	}
	// Round up to 1 digit
	return fmt.Sprintf("%.1f", oldFloat) == fmt.Sprintf("%.1f", newFloat)
}

func getListOsnFromBlockChainPlatform(blockchainPlatformId *string, client *oci_blockchain.BlockchainPlatformClient) ([]*string, error) {
	request := oci_blockchain.GetBlockchainPlatformRequest{}
	request.BlockchainPlatformId = blockchainPlatformId
	response, err := client.GetBlockchainPlatform(context.Background(), request)
	if err != nil {
		return nil, nil
	}
	osnList := response.ComponentDetails.Osns
	var results []*string
	for _, osn := range osnList {
		results = append(results, osn.OsnKey)
	}
	return results, nil
}

func getListPeerFromBlockChainPlatform(blockchainPlatformId *string, client *oci_blockchain.BlockchainPlatformClient) ([]*string, error) {
	request := oci_blockchain.GetBlockchainPlatformRequest{}
	request.BlockchainPlatformId = blockchainPlatformId
	response, err := client.GetBlockchainPlatform(context.Background(), request)
	if err != nil {
		return nil, nil
	}
	peerList := response.ComponentDetails.Peers
	var results []*string
	for _, peer := range peerList {
		results = append(results, peer.PeerKey)
	}
	return results, nil
}

// a is a superset of b with 1 more item
func difference(a, b []*string) (*string, error) {
	if len(a)-len(b) != 1 {
		return nil, fmt.Errorf("[ERROR] can not determine new resource")
	}
	mb := make(map[string]struct{}, len(b))
	for _, x := range b {
		mb[*x] = struct{}{}
	}
	for _, x := range a {
		if _, found := mb[*x]; !found {
			return x, nil
		}
	}
	return nil, fmt.Errorf("[ERROR] can not determine new resource")
}
