// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package containerengine

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_containerengine "github.com/oracle/oci-go-sdk/v56/containerengine"
)

func ContainerengineNodePoolDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["node_pool_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(ContainerengineNodePoolResource(), fieldMap, readSingularContainerengineNodePool)
}

func readSingularContainerengineNodePool(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerengineNodePoolDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ContainerEngineClient()

	return tfresource.ReadResource(sync)
}

type ContainerengineNodePoolDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_containerengine.ContainerEngineClient
	Res    *oci_containerengine.GetNodePoolResponse
}

func (s *ContainerengineNodePoolDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ContainerengineNodePoolDataSourceCrud) Get() error {
	request := oci_containerengine.GetNodePoolRequest{}

	if nodePoolId, ok := s.D.GetOkExists("node_pool_id"); ok {
		tmp := nodePoolId.(string)
		request.NodePoolId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "containerengine")

	response, err := s.Client.GetNodePool(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ContainerengineNodePoolDataSourceCrud) SetData() error {
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

	initialNodeLabels := []interface{}{}
	for _, item := range s.Res.InitialNodeLabels {
		initialNodeLabels = append(initialNodeLabels, KeyValueToMap(item))
	}
	s.D.Set("initial_node_labels", initialNodeLabels)

	if s.Res.KubernetesVersion != nil {
		s.D.Set("kubernetes_version", *s.Res.KubernetesVersion)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.NodeConfigDetails != nil {
		s.D.Set("node_config_details", []interface{}{NodePoolNodeConfigDetailsToMap(s.Res.NodeConfigDetails, true)})
	} else {
		s.D.Set("node_config_details", nil)
	}

	if s.Res.NodeImageId != nil {
		s.D.Set("node_image_id", *s.Res.NodeImageId)
	}

	if s.Res.NodeImageName != nil {
		s.D.Set("node_image_name", *s.Res.NodeImageName)
	}

	s.D.Set("node_metadata", s.Res.NodeMetadata)

	if s.Res.NodeShape != nil {
		s.D.Set("node_shape", *s.Res.NodeShape)
	}

	if s.Res.NodeShapeConfig != nil {
		s.D.Set("node_shape_config", []interface{}{NodeShapeConfigToMap(s.Res.NodeShapeConfig)})
	} else {
		s.D.Set("node_shape_config", nil)
	}

	if s.Res.NodeSource != nil {
		nodeSourceArray := []interface{}{}
		if nodeSourceMap := NodeSourceOptionToMap(&s.Res.NodeSource); nodeSourceMap != nil {
			nodeSourceArray = append(nodeSourceArray, nodeSourceMap)
		}
		s.D.Set("node_source", nodeSourceArray)
	} else {
		s.D.Set("node_source", nil)
	}

	if s.Res.NodeSourceDetails != nil {
		nodeSourceDetailsArray := []interface{}{}
		if nodeSourceDetailsMap := NodeSourceDetailsToMap(&s.Res.NodeSourceDetails); nodeSourceDetailsMap != nil {
			nodeSourceDetailsArray = append(nodeSourceDetailsArray, nodeSourceDetailsMap)
		}
		s.D.Set("node_source_details", nodeSourceDetailsArray)
	} else {
		s.D.Set("node_source_details", nil)
	}

	nodes := []interface{}{}
	for _, item := range s.Res.Nodes {
		nodes = append(nodes, NodeToMap(item))
	}
	s.D.Set("nodes", nodes)

	if s.Res.QuantityPerSubnet != nil {
		s.D.Set("quantity_per_subnet", *s.Res.QuantityPerSubnet)
	}

	if s.Res.SshPublicKey != nil {
		s.D.Set("ssh_public_key", *s.Res.SshPublicKey)
	}

	s.D.Set("subnet_ids", s.Res.SubnetIds)

	return nil
}
