// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_containerengine "github.com/oracle/oci-go-sdk/containerengine"
)

func NodePoolsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readNodePools,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"cluster_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"node_pools": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(NodePoolDataSource()),
			},
		},
	}
}

func readNodePools(d *schema.ResourceData, m interface{}) error {
	sync := &NodePoolsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).containerEngineClient

	return ReadResource(sync)
}

type NodePoolsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_containerengine.ContainerEngineClient
	Res    *oci_containerengine.ListNodePoolsResponse
}

func (s *NodePoolsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *NodePoolsDataSourceCrud) Get() error {
	request := oci_containerengine.ListNodePoolsRequest{}

	if clusterId, ok := s.D.GetOkExists("cluster_id"); ok {
		tmp := clusterId.(string)
		request.ClusterId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "containerengine")

	response, err := s.Client.ListNodePools(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListNodePools(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *NodePoolsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		nodePool := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.ClusterId != nil {
			nodePool["cluster_id"] = *r.ClusterId
		}

		if r.Id != nil {
			nodePool["id"] = *r.Id
		}

		initialNodeLabels := []interface{}{}
		for _, item := range r.InitialNodeLabels {
			initialNodeLabels = append(initialNodeLabels, KeyValueToMap(item))
		}
		nodePool["initial_node_labels"] = initialNodeLabels

		if r.KubernetesVersion != nil {
			nodePool["kubernetes_version"] = *r.KubernetesVersion
		}

		if r.Name != nil {
			nodePool["name"] = *r.Name
		}

		if r.NodeImageId != nil {
			nodePool["node_image_id"] = *r.NodeImageId
		}

		if r.NodeImageName != nil {
			nodePool["node_image_name"] = *r.NodeImageName
		}

		if r.NodeShape != nil {
			nodePool["node_shape"] = *r.NodeShape
		}

		if r.QuantityPerSubnet != nil {
			nodePool["quantity_per_subnet"] = *r.QuantityPerSubnet
		}

		if r.SshPublicKey != nil {
			nodePool["ssh_public_key"] = *r.SshPublicKey
		}

		nodePool["subnet_ids"] = r.SubnetIds

		resources = append(resources, nodePool)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, NodePoolsDataSource().Schema["node_pools"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("node_pools", resources); err != nil {
		return err
	}

	return nil
}
