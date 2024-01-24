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

func CloudBridgeAssetSourceDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["asset_source_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(CloudBridgeAssetSourceResource(), fieldMap, readSingularCloudBridgeAssetSource)
}

func readSingularCloudBridgeAssetSource(d *schema.ResourceData, m interface{}) error {
	sync := &CloudBridgeAssetSourceDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DiscoveryClient()

	return tfresource.ReadResource(sync)
}

type CloudBridgeAssetSourceDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_cloud_bridge.DiscoveryClient
	Res    *oci_cloud_bridge.GetAssetSourceResponse
}

func (s *CloudBridgeAssetSourceDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CloudBridgeAssetSourceDataSourceCrud) Get() error {
	request := oci_cloud_bridge.GetAssetSourceRequest{}

	if assetSourceId, ok := s.D.GetOkExists("asset_source_id"); ok {
		tmp := assetSourceId.(string)
		request.AssetSourceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "cloud_bridge")

	response, err := s.Client.GetAssetSource(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CloudBridgeAssetSourceDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.AssetSource.GetId())
	switch v := (s.Res.AssetSource).(type) {
	case oci_cloud_bridge.VmWareAssetSource:
		s.D.Set("type", "VMWARE")

		if v.Id != nil {
			s.D.Set("id", *v.Id)
		}

		if v.AreHistoricalMetricsCollected != nil {
			s.D.Set("are_historical_metrics_collected", *v.AreHistoricalMetricsCollected)
		}

		if v.AreRealtimeMetricsCollected != nil {
			s.D.Set("are_realtime_metrics_collected", *v.AreRealtimeMetricsCollected)
		}

		if v.DiscoveryCredentials != nil {
			s.D.Set("discovery_credentials", []interface{}{AssetSourceCredentialsToMap(v.DiscoveryCredentials)})
		} else {
			s.D.Set("discovery_credentials", nil)
		}

		if v.ReplicationCredentials != nil {
			s.D.Set("replication_credentials", []interface{}{AssetSourceCredentialsToMap(v.ReplicationCredentials)})
		} else {
			s.D.Set("replication_credentials", nil)
		}

		if v.VcenterEndpoint != nil {
			s.D.Set("vcenter_endpoint", *v.VcenterEndpoint)
		}

		if v.AssetsCompartmentId != nil {
			s.D.Set("assets_compartment_id", *v.AssetsCompartmentId)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DiscoveryScheduleId != nil {
			s.D.Set("discovery_schedule_id", *v.DiscoveryScheduleId)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.EnvironmentId != nil {
			s.D.Set("environment_id", *v.EnvironmentId)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.InventoryId != nil {
			s.D.Set("inventory_id", *v.InventoryId)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
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
		log.Printf("[WARN] Received 'type' of unknown type %v", s.Res.AssetSource)
		return nil
	}

	return nil
}
