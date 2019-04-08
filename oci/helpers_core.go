// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"strings"
	"time"

	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/common"
	oci_core "github.com/oracle/oci-go-sdk/core"
)

const (
	subnetService = "subnet"
)

var coreServiceExpectedRetryDurationMap = map[string]serviceExpectedRetryDurationFunc{
	subnetService: getSubnetExpectedRetryDuration,
}

// This applies the differences between the regular schema and the one
// we supply for default resources, and returns the schema for a default resource
func ConvertToDefaultVcnResourceSchema(resourceSchema *schema.Resource) *schema.Resource {
	if resourceSchema == nil {
		return nil
	}

	resourceSchema.Importer = &schema.ResourceImporter{
		State: ImportDefaultVcnResource,
	}

	resourceSchema.Schema["manage_default_resource_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
		ForceNew: true,
	}

	delete(resourceSchema.Schema, "compartment_id")
	delete(resourceSchema.Schema, "vcn_id")

	return resourceSchema
}

func ImportDefaultVcnResource(d *schema.ResourceData, value interface{}) ([]*schema.ResourceData, error) {
	err := d.Set("manage_default_resource_id", d.Id())
	return []*schema.ResourceData{d}, err
}

func LaunchOptionsToMap(obj *oci_core.LaunchOptions) map[string]interface{} {
	result := map[string]interface{}{}

	result["boot_volume_type"] = string(obj.BootVolumeType)

	result["firmware"] = string(obj.Firmware)

	result["network_type"] = string(obj.NetworkType)

	result["remote_data_volume_type"] = string(obj.RemoteDataVolumeType)

	if obj.IsConsistentVolumeNamingEnabled != nil {
		result["is_consistent_volume_naming_enabled"] = bool(*obj.IsConsistentVolumeNamingEnabled)
	}

	if obj.IsPvEncryptionInTransitEnabled != nil {
		result["is_pv_encryption_in_transit_enabled"] = bool(*obj.IsPvEncryptionInTransitEnabled)
	}

	return result
}

func getBackupPolicyId(assetId *string, client *oci_core.BlockstorageClient) (*string, error) {
	request := oci_core.GetVolumeBackupPolicyAssetAssignmentRequest{}
	request.AssetId = assetId
	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

	response, err := client.GetVolumeBackupPolicyAssetAssignment(context.Background(), request)
	if err != nil {
		return nil, err
	}

	if len(response.Items) > 0 {
		policyAssignment := response.Items[0]
		return policyAssignment.PolicyId, nil
	} else {
		return nil, nil
	}
}

func (s *CoreVolumeBackupResourceCrud) createBlockStorageSourceRegionClient(region string) error {
	if s.SourceRegionClient == nil {
		sourceBlockStorageClient, err := oci_core.NewBlockstorageClientWithConfigurationProvider(*s.Client.ConfigurationProvider())
		if err != nil {
			return fmt.Errorf("cannot create client for the source region: %v", err)
		}
		err = configureClient(&sourceBlockStorageClient.BaseClient)
		if err != nil {
			return fmt.Errorf("cannot configure client for the source region: %v", err)
		}
		s.SourceRegionClient = &sourceBlockStorageClient
	}
	s.SourceRegionClient.SetRegion(region)

	return nil
}

func getCoreExpectedRetryDuration(response oci_common.OCIOperationResponse, disableNotFoundRetries bool, optionals ...string) time.Duration {
	if len(optionals) > 0 {
		if expectedRetryDurationFunc, ok := coreServiceExpectedRetryDurationMap[optionals[0]]; ok {
			return expectedRetryDurationFunc(response, disableNotFoundRetries, optionals[1:]...)
		}
	}
	return getDefaultExpectedRetryDuration(response, disableNotFoundRetries)
}

func getSubnetExpectedRetryDuration(response oci_common.OCIOperationResponse, disableNotFoundRetries bool, optionals ...string) time.Duration {
	defaultRetryTime := getDefaultExpectedRetryDuration(response, disableNotFoundRetries)
	if response.Response == nil || response.Response.HTTPResponse() == nil {
		return defaultRetryTime
	}
	if len(optionals) > 0 {
		switch optionals[0] {
		case deleteResource:
			switch statusCode := response.Response.HTTPResponse().StatusCode; statusCode {
			case 409:
				if e := response.Error; e != nil {
					if strings.Contains(e.Error(), "Conflict") {
						defaultRetryTime = longRetryTime
					}
				}
			}
		}
	}
	return defaultRetryTime
}
