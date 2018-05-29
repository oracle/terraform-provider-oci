// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_containerengine "github.com/oracle/oci-go-sdk/containerengine"

	"github.com/oracle/terraform-provider-oci/crud"
)

func ClusterOptionDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularClusterOption,
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

func readSingularClusterOption(d *schema.ResourceData, m interface{}) error {
	sync := &ClusterOptionDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).containerEngineClient

	return crud.ReadResource(sync)
}

type ClusterOptionDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_containerengine.ContainerEngineClient
	Res    *oci_containerengine.GetClusterOptionsResponse
}

func (s *ClusterOptionDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ClusterOptionDataSourceCrud) Get() error {
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

func (s *ClusterOptionDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())

	s.D.Set("kubernetes_versions", s.Res.KubernetesVersions)

	return
}
