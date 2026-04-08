// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package cloud_bridge

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_cloud_bridge "github.com/oracle/oci-go-sdk/v65/cloudbridge"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CloudBridgeAssetDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["asset_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(CloudBridgeAssetResource(), fieldMap, readSingularCloudBridgeAsset)
}

func readSingularCloudBridgeAsset(d *schema.ResourceData, m interface{}) error {
	sync := &CloudBridgeAssetDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).InventoryClient()

	return tfresource.ReadResource(sync)
}

type CloudBridgeAssetDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_cloud_bridge.InventoryClient
	Res    *oci_cloud_bridge.GetAssetResponse
}

func (s *CloudBridgeAssetDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CloudBridgeAssetDataSourceCrud) Get() error {
	request := oci_cloud_bridge.GetAssetRequest{}

	if assetId, ok := s.D.GetOkExists("asset_id"); ok {
		tmp := assetId.(string)
		request.AssetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "cloud_bridge")

	response, err := s.Client.GetAsset(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CloudBridgeAssetDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.GetId())
	switch v := (s.Res.Asset).(type) {
	case oci_cloud_bridge.AwsEbsAsset:
		s.D.Set("asset_type", "AWS_EBS")

		if v.AwsEbs != nil {
			s.D.Set("aws_ebs", []interface{}{AwsEbsPropertiesToMap(v.AwsEbs)})
		} else {
			s.D.Set("aws_ebs", nil)
		}

		s.D.Set("asset_source_ids", v.AssetSourceIds)

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("environment_type", v.EnvironmentType)

		if v.ExternalAssetKey != nil {
			s.D.Set("external_asset_key", *v.ExternalAssetKey)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.InventoryId != nil {
			s.D.Set("inventory_id", *v.InventoryId)
		}

		if v.SourceKey != nil {
			s.D.Set("source_key", *v.SourceKey)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_cloud_bridge.AwsEc2Asset:
		s.D.Set("asset_type", "AWS_EC2")

		if v.AttachedEbsVolumesCost != nil {
			s.D.Set("attached_ebs_volumes_cost", []interface{}{MonthlyCostSummaryToMap(v.AttachedEbsVolumesCost)})
		} else {
			s.D.Set("attached_ebs_volumes_cost", nil)
		}

		if v.AwsEc2 != nil {
			s.D.Set("aws_ec2", []interface{}{AwsEc2PropertiesToMap(v.AwsEc2)})
		} else {
			s.D.Set("aws_ec2", nil)
		}

		if v.AwsEc2Cost != nil {
			s.D.Set("aws_ec2cost", []interface{}{MonthlyCostSummaryToMap(v.AwsEc2Cost)})
		} else {
			s.D.Set("aws_ec2cost", nil)
		}

		if v.Compute != nil {
			s.D.Set("compute", []interface{}{ComputePropertiesToMap(v.Compute)})
		} else {
			s.D.Set("compute", nil)
		}

		if v.Vm != nil {
			s.D.Set("vm", []interface{}{VmPropertiesToMap(v.Vm)})
		} else {
			s.D.Set("vm", nil)
		}

		s.D.Set("asset_source_ids", v.AssetSourceIds)

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("environment_type", v.EnvironmentType)

		if v.ExternalAssetKey != nil {
			s.D.Set("external_asset_key", *v.ExternalAssetKey)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.InventoryId != nil {
			s.D.Set("inventory_id", *v.InventoryId)
		}

		if v.SourceKey != nil {
			s.D.Set("source_key", *v.SourceKey)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_cloud_bridge.InventoryAsset:
		s.D.Set("asset_type", "INVENTORY_ASSET")

		if v.AssetClassName != nil {
			s.D.Set("asset_class_name", *v.AssetClassName)
		}

		if v.AssetClassVersion != nil {
			s.D.Set("asset_class_version", *v.AssetClassVersion)
		}

		if v.AssetDetails != nil {
			tmp, err := tfresource.ConvertObjectToJsonString(v.AssetDetails)
			if err != nil {
				return err
			}
			s.D.Set("asset_details", tmp)
		} else {
			s.D.Set("asset_details", nil)
		}

		s.D.Set("asset_source_ids", v.AssetSourceIds)

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("environment_type", v.EnvironmentType)

		if v.ExternalAssetKey != nil {
			s.D.Set("external_asset_key", *v.ExternalAssetKey)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.InventoryId != nil {
			s.D.Set("inventory_id", *v.InventoryId)
		}

		if v.SourceKey != nil {
			s.D.Set("source_key", *v.SourceKey)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_cloud_bridge.VmAsset:
		s.D.Set("asset_type", "VM")

		if v.Compute != nil {
			s.D.Set("compute", []interface{}{ComputePropertiesToMap(v.Compute)})
		} else {
			s.D.Set("compute", nil)
		}

		if v.Vm != nil {
			s.D.Set("vm", []interface{}{VmPropertiesToMap(v.Vm)})
		} else {
			s.D.Set("vm", nil)
		}

		s.D.Set("asset_source_ids", v.AssetSourceIds)

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("environment_type", v.EnvironmentType)

		if v.ExternalAssetKey != nil {
			s.D.Set("external_asset_key", *v.ExternalAssetKey)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.InventoryId != nil {
			s.D.Set("inventory_id", *v.InventoryId)
		}

		if v.SourceKey != nil {
			s.D.Set("source_key", *v.SourceKey)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_cloud_bridge.VmwareVmAsset:
		s.D.Set("asset_type", "VMWARE_VM")

		if v.Compute != nil {
			s.D.Set("compute", []interface{}{ComputePropertiesToMap(v.Compute)})
		} else {
			s.D.Set("compute", nil)
		}

		if v.Vm != nil {
			s.D.Set("vm", []interface{}{VmPropertiesToMap(v.Vm)})
		} else {
			s.D.Set("vm", nil)
		}

		if v.VmwareVCenter != nil {
			s.D.Set("vmware_vcenter", []interface{}{VmwareVCenterPropertiesToMap(v.VmwareVCenter)})
		} else {
			s.D.Set("vmware_vcenter", nil)
		}

		if v.VmwareVm != nil {
			s.D.Set("vmware_vm", []interface{}{VmwareVmPropertiesToMap(v.VmwareVm)})
		} else {
			s.D.Set("vmware_vm", nil)
		}

		s.D.Set("asset_source_ids", v.AssetSourceIds)

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("environment_type", v.EnvironmentType)

		if v.ExternalAssetKey != nil {
			s.D.Set("external_asset_key", *v.ExternalAssetKey)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.InventoryId != nil {
			s.D.Set("inventory_id", *v.InventoryId)
		}

		if v.SourceKey != nil {
			s.D.Set("source_key", *v.SourceKey)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	default:
		log.Printf("[WARN] Received 'asset_type' of unknown type %v", s.Res.Asset)
		return nil
	}

	return nil
}
