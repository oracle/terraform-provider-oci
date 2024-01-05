// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package containerengine

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_containerengine "github.com/oracle/oci-go-sdk/v65/containerengine"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ContainerengineVirtualNodePoolDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["virtual_node_pool_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(ContainerengineVirtualNodePoolResource(), fieldMap, readSingularContainerengineVirtualNodePool)
}

func readSingularContainerengineVirtualNodePool(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerengineVirtualNodePoolDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ContainerEngineClient()

	return tfresource.ReadResource(sync)
}

type ContainerengineVirtualNodePoolDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_containerengine.ContainerEngineClient
	Res    *oci_containerengine.GetVirtualNodePoolResponse
}

func (s *ContainerengineVirtualNodePoolDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ContainerengineVirtualNodePoolDataSourceCrud) Get() error {
	request := oci_containerengine.GetVirtualNodePoolRequest{}

	if virtualNodePoolId, ok := s.D.GetOkExists("virtual_node_pool_id"); ok {
		tmp := virtualNodePoolId.(string)
		request.VirtualNodePoolId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "containerengine")

	response, err := s.Client.GetVirtualNodePool(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ContainerengineVirtualNodePoolDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.ClusterId != nil {
		s.D.Set("cluster_id", *s.Res.ClusterId)
	}

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

	initialVirtualNodeLabels := []interface{}{}
	for _, item := range s.Res.InitialVirtualNodeLabels {
		initialVirtualNodeLabels = append(initialVirtualNodeLabels, InitialVirtualNodeLabelToMap(item))
	}
	s.D.Set("initial_virtual_node_labels", initialVirtualNodeLabels)

	if s.Res.KubernetesVersion != nil {
		s.D.Set("kubernetes_version", *s.Res.KubernetesVersion)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("nsg_ids", s.Res.NsgIds)

	placementConfigurations := []interface{}{}
	for _, item := range s.Res.PlacementConfigurations {
		placementConfigurations = append(placementConfigurations, PlacementConfigurationToMap(item))
	}
	s.D.Set("placement_configurations", placementConfigurations)

	if s.Res.PodConfiguration != nil {
		s.D.Set("pod_configuration", []interface{}{PodConfigurationToMap(s.Res.PodConfiguration, true)})
	} else {
		s.D.Set("pod_configuration", nil)
	}

	if s.Res.Size != nil {
		s.D.Set("size", *s.Res.Size)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	taints := []interface{}{}
	for _, item := range s.Res.Taints {
		taints = append(taints, TaintToMap(item))
	}
	s.D.Set("taints", taints)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.VirtualNodeTags != nil {
		s.D.Set("virtual_node_tags", []interface{}{VirtualNodeTagsToMap(s.Res.VirtualNodeTags)})
	} else {
		s.D.Set("virtual_node_tags", nil)
	}

	return nil
}
