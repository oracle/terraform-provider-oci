// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package load_balancer

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_load_balancer "github.com/oracle/oci-go-sdk/v56/loadbalancer"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func LoadBalancerSslCipherSuitesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readLoadBalancerSslCipherSuites,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"load_balancer_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_cipher_suites": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     LoadBalancerSslCipherSuiteResource(),
			},
		},
	}
}

func readLoadBalancerSslCipherSuites(d *schema.ResourceData, m interface{}) error {
	sync := &LoadBalancerSslCipherSuitesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LoadBalancerClient()

	return tfresource.ReadResource(sync)
}

type LoadBalancerSslCipherSuitesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_load_balancer.LoadBalancerClient
	Res    *oci_load_balancer.ListSSLCipherSuitesResponse
}

func (s *LoadBalancerSslCipherSuitesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LoadBalancerSslCipherSuitesDataSourceCrud) Get() error {
	request := oci_load_balancer.ListSSLCipherSuitesRequest{}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "load_balancer")

	response, err := s.Client.ListSSLCipherSuites(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *LoadBalancerSslCipherSuitesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("LoadBalancerSslCipherSuitesDataSource-", LoadBalancerSslCipherSuitesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		sslCipherSuite := map[string]interface{}{}

		sslCipherSuite["ciphers"] = r.Ciphers

		if r.Name != nil {
			sslCipherSuite["name"] = *r.Name
		}

		resources = append(resources, sslCipherSuite)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, LoadBalancerSslCipherSuitesDataSource().Schema["ssl_cipher_suites"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("ssl_cipher_suites", resources); err != nil {
		return err
	}

	return nil
}
