// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_containerengine "github.com/oracle/oci-go-sdk/containerengine"
)

func ContainerengineNodePoolOptionDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularContainerengineNodePoolOption,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
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

func readSingularContainerengineNodePoolOption(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerengineNodePoolOptionDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).containerEngineClient

	return ReadResource(sync)
}

type ContainerengineNodePoolOptionDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_containerengine.ContainerEngineClient
	Res    *oci_containerengine.GetNodePoolOptionsResponse
}

func (s *ContainerengineNodePoolOptionDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ContainerengineNodePoolOptionDataSourceCrud) Get() error {
	request := oci_containerengine.GetNodePoolOptionsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

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

func (s *ContainerengineNodePoolOptionDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())

	s.D.Set("images", s.Res.Images)

	s.D.Set("kubernetes_versions", s.Res.KubernetesVersions)

	s.D.Set("shapes", s.Res.Shapes)

	return nil
}
