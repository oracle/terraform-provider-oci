// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package blockchain

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	oci_blockchain "github.com/oracle/oci-go-sdk/v56/blockchain"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func blockchainPlatformComputeShapeDiffSuppressFunction(key string, old string, new string, d *schema.ResourceData) bool {
	if old == "" || new == "" {
		return false
	}
	upperCaseNewValue := strings.ToUpper(new)
	upperCaseOldValue := strings.ToUpper(old)

	// ENTERPRISE_CUSTOM is auto changed when Update totalOcpuCapacity
	if upperCaseNewValue == "ENTERPRISE_CUSTOM" || upperCaseOldValue == "ENTERPRISE_CUSTOM" {
		return true
	}

	return upperCaseOldValue == upperCaseNewValue
}

func sendUpdateBlockchainPlatformRequest(s *BlockchainBlockchainPlatformResourceCrud, request oci_blockchain.UpdateBlockchainPlatformRequest) error {
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "blockchain")

	response, err := s.Client.UpdateBlockchainPlatform(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	err = s.getBlockchainPlatformFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "blockchain"), oci_blockchain.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
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
