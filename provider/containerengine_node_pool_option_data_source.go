// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_containerengine "github.com/oracle/oci-go-sdk/containerengine"

	"github.com/oracle/terraform-provider-oci/crud"
)

func NodePoolOptionDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularNodePoolOption,
		Schema: map[string]*schema.Schema{
			"node_pool_option_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"images": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"kubernetes_versions": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"shapes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func readSingularNodePoolOption(d *schema.ResourceData, m interface{}) error {
	sync := &NodePoolOptionDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).containerEngineClient

	return crud.ReadResource(sync)
}

type NodePoolOptionDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_containerengine.ContainerEngineClient
	Res    *oci_containerengine.GetNodePoolOptionsResponse
}

func (s *NodePoolOptionDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *NodePoolOptionDataSourceCrud) Get() error {
	request := oci_containerengine.GetNodePoolOptionsRequest{}

	if nodePoolOptionId, ok := s.D.GetOkExists("node_pool_option_id"); ok {
		tmp := nodePoolOptionId.(string)
		request.NodePoolOptionId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "containerengine")

	response, err := s.Client.GetNodePoolOptions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *NodePoolOptionDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())

	s.D.Set("images", s.Res.Images)

	s.D.Set("kubernetes_versions", s.Res.KubernetesVersions)

	s.D.Set("shapes", s.Res.Shapes)

	return
}
