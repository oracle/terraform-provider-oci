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

func ContainerengineClusterOptionDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularContainerengineClusterOption,
		Schema: map[string]*schema.Schema{
			"cluster_option_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
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
	sync.Client = m.(*client.OracleClients).ContainerEngineClient()

	return tfresource.ReadResource(sync)
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

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "containerengine")

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

	s.D.SetId(tfresource.GenerateDataSourceHashID("ContainerengineClusterOptionDataSource-", ContainerengineClusterOptionDataSource(), s.D))

	s.D.Set("kubernetes_versions", s.Res.KubernetesVersions)

	return nil
}
