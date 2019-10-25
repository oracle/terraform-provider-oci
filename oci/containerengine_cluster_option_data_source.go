// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_containerengine "github.com/oracle/oci-go-sdk/containerengine"
)

func ContainerengineClusterOptionDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularContainerengineClusterOption,
		Schema: map[string]*schema.Schema{
			"cluster_option_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"kubernetes_versions": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func readSingularContainerengineClusterOption(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerengineClusterOptionDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).containerEngineClient

	return ReadResource(sync)
}

type ContainerengineClusterOptionDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_containerengine.ContainerEngineClient
	Res    *oci_containerengine.GetClusterOptionsResponse
}

func (s *ContainerengineClusterOptionDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ContainerengineClusterOptionDataSourceCrud) Get() error {
	request := oci_containerengine.GetClusterOptionsRequest{}

	if clusterOptionId, ok := s.D.GetOkExists("cluster_option_id"); ok {
		tmp := clusterOptionId.(string)
		request.ClusterOptionId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "containerengine")

	response, err := s.Client.GetClusterOptions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ContainerengineClusterOptionDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())

	s.D.Set("kubernetes_versions", s.Res.KubernetesVersions)

	return nil
}
