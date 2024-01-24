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

func ContainerengineVirtualNodePoolsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readContainerengineVirtualNodePools,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"cluster_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"virtual_node_pools": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(ContainerengineVirtualNodePoolDataSource()),
			},
		},
	}
}

func readContainerengineVirtualNodePools(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerengineVirtualNodePoolsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ContainerEngineClient()

	return tfresource.ReadResource(sync)
}

type ContainerengineVirtualNodePoolsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_containerengine.ContainerEngineClient
	Res    *oci_containerengine.ListVirtualNodePoolsResponse
}

func (s *ContainerengineVirtualNodePoolsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ContainerengineVirtualNodePoolsDataSourceCrud) Get() error {
	request := oci_containerengine.ListVirtualNodePoolsRequest{}

	if clusterId, ok := s.D.GetOkExists("cluster_id"); ok {
		tmp := clusterId.(string)
		request.ClusterId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.Name = &tmp
	}

	if states, ok := s.D.GetOkExists("state"); ok {
		var enumStates []oci_containerengine.VirtualNodePoolLifecycleStateEnum
		for _, r := range states.([]interface{}) {
			enumStates = append(enumStates, oci_containerengine.VirtualNodePoolLifecycleStateEnum(r.(string)))
		}
		request.LifecycleState = enumStates
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "containerengine")

	response, err := s.Client.ListVirtualNodePools(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListVirtualNodePools(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ContainerengineVirtualNodePoolsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ContainerengineVirtualNodePoolsDataSource-", ContainerengineVirtualNodePoolsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		virtualNodePool := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.ClusterId != nil {
			virtualNodePool["cluster_id"] = *r.ClusterId
		}

		if r.DefinedTags != nil {
			virtualNodePool["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			virtualNodePool["display_name"] = *r.DisplayName
		}

		virtualNodePool["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			virtualNodePool["id"] = *r.Id
		}

		initialVirtualNodeLabels := []interface{}{}
		for _, item := range r.InitialVirtualNodeLabels {
			initialVirtualNodeLabels = append(initialVirtualNodeLabels, InitialVirtualNodeLabelToMap(item))
		}
		virtualNodePool["initial_virtual_node_labels"] = initialVirtualNodeLabels

		if r.KubernetesVersion != nil {
			virtualNodePool["kubernetes_version"] = *r.KubernetesVersion
		}

		if r.LifecycleDetails != nil {
			virtualNodePool["lifecycle_details"] = *r.LifecycleDetails
		}

		virtualNodePool["nsg_ids"] = r.NsgIds

		placementConfigurations := []interface{}{}
		for _, item := range r.PlacementConfigurations {
			placementConfigurations = append(placementConfigurations, PlacementConfigurationToMap(item))
		}
		virtualNodePool["placement_configurations"] = placementConfigurations

		if r.PodConfiguration != nil {
			virtualNodePool["pod_configuration"] = []interface{}{PodConfigurationToMap(r.PodConfiguration, true)}
		} else {
			virtualNodePool["pod_configuration"] = nil
		}

		if r.Size != nil {
			virtualNodePool["size"] = *r.Size
		}

		virtualNodePool["state"] = r.LifecycleState

		if r.SystemTags != nil {
			virtualNodePool["system_tags"] = tfresource.SystemTagsToMap(r.SystemTags)
		}

		taints := []interface{}{}
		for _, item := range r.Taints {
			taints = append(taints, TaintToMap(item))
		}
		virtualNodePool["taints"] = taints

		if r.TimeCreated != nil {
			virtualNodePool["time_created"] = r.TimeCreated.String()
		}

		if r.TimeUpdated != nil {
			virtualNodePool["time_updated"] = r.TimeUpdated.String()
		}

		resources = append(resources, virtualNodePool)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, ContainerengineVirtualNodePoolsDataSource().Schema["virtual_node_pools"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("virtual_node_pools", resources); err != nil {
		return err
	}

	return nil
}
