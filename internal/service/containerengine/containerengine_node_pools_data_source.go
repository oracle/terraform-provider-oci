// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package containerengine

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_containerengine "github.com/oracle/oci-go-sdk/v58/containerengine"
)

func ContainerengineNodePoolsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readContainerengineNodePools,
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
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"node_pools": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(ContainerengineNodePoolDataSource()),
			},
		},
	}
}

func readContainerengineNodePools(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerengineNodePoolsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ContainerEngineClient()

	return tfresource.ReadResource(sync)
}

type ContainerengineNodePoolsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_containerengine.ContainerEngineClient
	Res    *oci_containerengine.ListNodePoolsResponse
}

func (s *ContainerengineNodePoolsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ContainerengineNodePoolsDataSourceCrud) Get() error {
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "containerengine")

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

func (s *ContainerengineNodePoolsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ContainerengineNodePoolsDataSource-", ContainerengineNodePoolsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		nodePool := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.ClusterId != nil {
			nodePool["cluster_id"] = *r.ClusterId
		}

		if r.DefinedTags != nil {
			nodePool["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		nodePool["freeform_tags"] = r.FreeformTags

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

		if r.NodeConfigDetails != nil {
			nodePool["node_config_details"] = []interface{}{NodePoolNodeConfigDetailsToMap(r.NodeConfigDetails, true)}
		} else {
			nodePool["node_config_details"] = nil
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

		if r.NodeShapeConfig != nil {
			nodePool["node_shape_config"] = []interface{}{NodeShapeConfigToMap(r.NodeShapeConfig)}
		} else {
			nodePool["node_shape_config"] = nil
		}

		if r.NodeSource != nil {
			nodeSourceArray := []interface{}{}
			if nodeSourceMap := NodeSourceOptionToMap(&r.NodeSource); nodeSourceMap != nil {
				nodeSourceArray = append(nodeSourceArray, nodeSourceMap)
			}
			nodePool["node_source"] = nodeSourceArray
		} else {
			nodePool["node_source"] = nil
		}

		if r.NodeSourceDetails != nil {
			nodeSourceDetailsArray := []interface{}{}
			if nodeSourceDetailsMap := NodeSourceDetailsToMap(&r.NodeSourceDetails); nodeSourceDetailsMap != nil {
				nodeSourceDetailsArray = append(nodeSourceDetailsArray, nodeSourceDetailsMap)
			}
			nodePool["node_source_details"] = nodeSourceDetailsArray
		} else {
			nodePool["node_source_details"] = nil
		}

		if r.QuantityPerSubnet != nil {
			nodePool["quantity_per_subnet"] = *r.QuantityPerSubnet
		}

		if r.SshPublicKey != nil {
			nodePool["ssh_public_key"] = *r.SshPublicKey
		}

		nodePool["subnet_ids"] = r.SubnetIds

		if r.SystemTags != nil {
			nodePool["system_tags"] = tfresource.SystemTagsToMap(r.SystemTags)
		}

		resources = append(resources, nodePool)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, ContainerengineNodePoolsDataSource().Schema["node_pools"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("node_pools", resources); err != nil {
		return err
	}

	return nil
}
