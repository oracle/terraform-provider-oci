// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_containerengine "github.com/oracle/oci-go-sdk/containerengine"

	"io/ioutil"
)

func ClusterKubeConfigDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularClusterKubeConfig,
		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"expiration": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"token_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			// Computed

			"content": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularClusterKubeConfig(d *schema.ResourceData, m interface{}) error {
	sync := &ClusterKubeConfigDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).containerEngineClient

	return ReadResource(sync)
}

type ClusterKubeConfigDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_containerengine.ContainerEngineClient
	Res    *[]byte
}

func (s *ClusterKubeConfigDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ClusterKubeConfigDataSourceCrud) Get() error {
	request := oci_containerengine.CreateKubeconfigRequest{}

	if clusterId, ok := s.D.GetOkExists("cluster_id"); ok {
		tmp := clusterId.(string)
		request.ClusterId = &tmp
	}

	if expiration, ok := s.D.GetOkExists("expiration"); ok {
		tmp := expiration.(int)
		request.Expiration = &tmp
	}

	if tokenVersion, ok := s.D.GetOkExists("token_version"); ok {
		tmp := tokenVersion.(string)
		request.TokenVersion = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "containerengine")

	response, err := s.Client.CreateKubeconfig(context.Background(), request)
	if err != nil {
		return err
	}

	if response.Content != nil {
		defer response.Content.Close()
		if contentBytes, err := ioutil.ReadAll(response.Content); err == nil {
			s.Res = &contentBytes
		} else {
			return err
		}
	}

	return nil
}

func (s *ClusterKubeConfigDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())

	s.D.Set("content", string(*s.Res))

	return nil
}
