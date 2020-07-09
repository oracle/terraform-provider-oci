// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_load_balancer "github.com/oracle/oci-go-sdk/v27/loadbalancer"
)

func init() {
	RegisterDatasource("oci_load_balancer_hostnames", LoadBalancerHostnamesDataSource())
}

func LoadBalancerHostnamesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readLoadBalancerHostnames,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"load_balancer_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"hostnames": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     LoadBalancerHostnameResource(),
			},
		},
	}
}

func readLoadBalancerHostnames(d *schema.ResourceData, m interface{}) error {
	sync := &LoadBalancerHostnamesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient()

	return ReadResource(sync)
}

type LoadBalancerHostnamesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_load_balancer.LoadBalancerClient
	Res    *oci_load_balancer.ListHostnamesResponse
}

func (s *LoadBalancerHostnamesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LoadBalancerHostnamesDataSourceCrud) Get() error {
	request := oci_load_balancer.ListHostnamesRequest{}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "load_balancer")

	response, err := s.Client.ListHostnames(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *LoadBalancerHostnamesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("LoadBalancerHostnamesDataSource-", LoadBalancerHostnamesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		hostname := map[string]interface{}{}

		if r.Hostname != nil {
			hostname["hostname"] = *r.Hostname
		}

		if r.Name != nil {
			hostname["name"] = *r.Name
		}

		resources = append(resources, hostname)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, LoadBalancerHostnamesDataSource().Schema["hostnames"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("hostnames", resources); err != nil {
		return err
	}

	return nil
}
