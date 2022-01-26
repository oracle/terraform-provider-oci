// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package containerengine

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_containerengine "github.com/oracle/oci-go-sdk/v56/containerengine"

	"io/ioutil"
)

func ContainerengineClusterKubeConfigDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularContainerengineClusterKubeConfig,
		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"endpoint": {
				Type:     schema.TypeString,
				Optional: true,
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

func readSingularContainerengineClusterKubeConfig(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerengineClusterKubeConfigDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ContainerEngineClient()

	return tfresource.ReadResource(sync)
}

type ContainerengineClusterKubeConfigDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_containerengine.ContainerEngineClient
	Res    *[]byte
}

func (s *ContainerengineClusterKubeConfigDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ContainerengineClusterKubeConfigDataSourceCrud) Get() error {
	request := oci_containerengine.CreateKubeconfigRequest{}

	if clusterId, ok := s.D.GetOkExists("cluster_id"); ok {
		tmp := clusterId.(string)
		request.ClusterId = &tmp
	}

	if endpoint, ok := s.D.GetOkExists("endpoint"); ok {
		request.Endpoint = oci_containerengine.CreateClusterKubeconfigContentDetailsEndpointEnum(endpoint.(string))
	}

	if expiration, ok := s.D.GetOkExists("expiration"); ok {
		tmp := expiration.(int)
		request.Expiration = &tmp
	}

	if tokenVersion, ok := s.D.GetOkExists("token_version"); ok {
		tmp := tokenVersion.(string)
		request.TokenVersion = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "containerengine")

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

func (s *ContainerengineClusterKubeConfigDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ContainerengineClusterKubeConfigDataSource-", ContainerengineClusterKubeConfigDataSource(), s.D))

	s.D.Set("content", string(*s.Res))

	return nil
}
