// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"

	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

// getSubsettingPolicyTargetAndCompartment resolves the target used by action
// requests when targetId is omitted. The service uses the target source on the
// policy, or the target associated with its sensitive data model.
func getSubsettingPolicyTargetAndCompartment(ctx context.Context, client *oci_data_safe.DataSafeClient, policyID string) (*string, *string, error) {
	policyRequest := oci_data_safe.GetSubsettingPolicyRequest{SubsettingPolicyId: &policyID}
	policyRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	policyResponse, err := client.GetSubsettingPolicy(ctx, policyRequest)
	if err != nil {
		return nil, nil, err
	}

	var targetID *string
	switch source := policyResponse.SubsettingPolicy.SchemaSource.(type) {
	case oci_data_safe.SchemaSourceFromTargetDetails:
		targetID = source.TargetId
	case oci_data_safe.SchemaSourceFromSdmDetails:
		if source.SensitiveDataModelId != nil {
			sensitiveDataModelRequest := oci_data_safe.GetSensitiveDataModelRequest{
				SensitiveDataModelId: source.SensitiveDataModelId,
			}
			sensitiveDataModelRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

			sensitiveDataModelResponse, err := client.GetSensitiveDataModel(ctx, sensitiveDataModelRequest)
			if err != nil {
				return nil, nil, err
			}
			targetID = sensitiveDataModelResponse.SensitiveDataModel.TargetId
		}
	}

	return policyResponse.SubsettingPolicy.CompartmentId, targetID, nil
}
