// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v65/core"
)

func CoreClusterNetworkDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["cluster_network_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(CoreClusterNetworkResource(), fieldMap, readSingularCoreClusterNetwork)
}

func readSingularCoreClusterNetwork(d *schema.ResourceData, m interface{}) error {
	sync := &CoreClusterNetworkDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeManagementClient()

	return tfresource.ReadResource(sync)
}

type CoreClusterNetworkDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.ComputeManagementClient
	Res    *oci_core.GetClusterNetworkResponse
}

func (s *CoreClusterNetworkDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreClusterNetworkDataSourceCrud) Get() error {
	request := oci_core.GetClusterNetworkRequest{}

	if clusterNetworkId, ok := s.D.GetOkExists("cluster_network_id"); ok {
		tmp := clusterNetworkId.(string)
		request.ClusterNetworkId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.GetClusterNetwork(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CoreClusterNetworkDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.HpcIslandId != nil {
		s.D.Set("hpc_island_id", *s.Res.HpcIslandId)
	}

	instancePools := []interface{}{}
	for _, item := range s.Res.InstancePools {
		instancePools = append(instancePools, InstancePoolToMap(item))
	}
	s.D.Set("instance_pools", instancePools)

	s.D.Set("network_block_ids", s.Res.NetworkBlockIds)
	s.D.Set("network_block_ids", s.Res.NetworkBlockIds)

	if s.Res.PlacementConfiguration != nil {
		s.D.Set("placement_configuration", []interface{}{ClusterNetworkPlacementConfigurationDetailsToMap(s.Res.PlacementConfiguration, true)})
	} else {
		s.D.Set("placement_configuration", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
