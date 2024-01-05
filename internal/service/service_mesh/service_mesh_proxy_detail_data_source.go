// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package service_mesh

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_service_mesh "github.com/oracle/oci-go-sdk/v65/servicemesh"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ServiceMeshProxyDetailDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularServiceMeshProxyDetail,
		Schema: map[string]*schema.Schema{
			// Computed
			"proxy_image": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularServiceMeshProxyDetail(d *schema.ResourceData, m interface{}) error {
	sync := &ServiceMeshProxyDetailDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ServiceMeshClient()

	return tfresource.ReadResource(sync)
}

type ServiceMeshProxyDetailDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_service_mesh.ServiceMeshClient
	Res    *oci_service_mesh.GetProxyDetailsResponse
}

func (s *ServiceMeshProxyDetailDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ServiceMeshProxyDetailDataSourceCrud) Get() error {
	request := oci_service_mesh.GetProxyDetailsRequest{}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "service_mesh")

	response, err := s.Client.GetProxyDetails(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ServiceMeshProxyDetailDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ServiceMeshProxyDetailDataSource-", ServiceMeshProxyDetailDataSource(), s.D))

	if s.Res.ProxyImage != nil {
		s.D.Set("proxy_image", *s.Res.ProxyImage)
	}

	return nil
}
