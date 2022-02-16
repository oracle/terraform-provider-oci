// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package blockchain

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_blockchain "github.com/oracle/oci-go-sdk/v58/blockchain"
)

func BlockchainBlockchainPlatformDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["blockchain_platform_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(BlockchainBlockchainPlatformResource(), fieldMap, readSingularBlockchainBlockchainPlatform)
}

func readSingularBlockchainBlockchainPlatform(d *schema.ResourceData, m interface{}) error {
	sync := &BlockchainBlockchainPlatformDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BlockchainPlatformClient()

	return tfresource.ReadResource(sync)
}

type BlockchainBlockchainPlatformDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_blockchain.BlockchainPlatformClient
	Res    *oci_blockchain.GetBlockchainPlatformResponse
}

func (s *BlockchainBlockchainPlatformDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *BlockchainBlockchainPlatformDataSourceCrud) Get() error {
	request := oci_blockchain.GetBlockchainPlatformRequest{}

	if blockchainPlatformId, ok := s.D.GetOkExists("blockchain_platform_id"); ok {
		tmp := blockchainPlatformId.(string)
		request.BlockchainPlatformId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "blockchain")

	response, err := s.Client.GetBlockchainPlatform(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *BlockchainBlockchainPlatformDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ComponentDetails != nil {
		s.D.Set("component_details", []interface{}{BlockchainPlatformComponentDetailsToMap(s.Res.ComponentDetails)})
	} else {
		s.D.Set("component_details", nil)
	}

	s.D.Set("compute_shape", s.Res.ComputeShape)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	hostOcpuUtilizationInfo := []interface{}{}
	for _, item := range s.Res.HostOcpuUtilizationInfo {
		hostOcpuUtilizationInfo = append(hostOcpuUtilizationInfo, OcpuUtilizationInfoToMap(item))
	}
	s.D.Set("host_ocpu_utilization_info", hostOcpuUtilizationInfo)

	if s.Res.IsByol != nil {
		s.D.Set("is_byol", *s.Res.IsByol)
	}

	if s.Res.IsMultiAD != nil {
		s.D.Set("is_multi_ad", *s.Res.IsMultiAD)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("load_balancer_shape", s.Res.LoadBalancerShape)

	s.D.Set("platform_role", s.Res.PlatformRole)

	s.D.Set("platform_shape_type", s.Res.PlatformShapeType)

	if s.Res.PlatformVersion != nil {
		s.D.Set("platform_version", *s.Res.PlatformVersion)
	}

	if s.Res.Replicas != nil {
		s.D.Set("replicas", []interface{}{ReplicaDetailsToMap(s.Res.Replicas)})
	} else {
		s.D.Set("replicas", nil)
	}

	if s.Res.ServiceEndpoint != nil {
		s.D.Set("service_endpoint", *s.Res.ServiceEndpoint)
	}

	if s.Res.ServiceVersion != nil {
		s.D.Set("service_version", *s.Res.ServiceVersion)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.StorageSizeInTBs != nil {
		s.D.Set("storage_size_in_tbs", *s.Res.StorageSizeInTBs)
	}

	if s.Res.StorageUsedInTBs != nil {
		s.D.Set("storage_used_in_tbs", *s.Res.StorageUsedInTBs)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.TotalOcpuCapacity != nil {
		s.D.Set("total_ocpu_capacity", *s.Res.TotalOcpuCapacity)
	}

	return nil
}
