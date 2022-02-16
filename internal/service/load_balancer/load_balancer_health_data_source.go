// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package load_balancer

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_load_balancer "github.com/oracle/oci-go-sdk/v58/loadbalancer"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func LoadBalancerLoadBalancerHealthDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularLoadBalancerLoadBalancerHealth,
		Schema: map[string]*schema.Schema{
			"load_balancer_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"critical_state_backend_set_names": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"total_backend_set_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"unknown_state_backend_set_names": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"warning_state_backend_set_names": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func readSingularLoadBalancerLoadBalancerHealth(d *schema.ResourceData, m interface{}) error {
	sync := &LoadBalancerLoadBalancerHealthDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LoadBalancerClient()

	return tfresource.ReadResource(sync)
}

type LoadBalancerLoadBalancerHealthDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_load_balancer.LoadBalancerClient
	Res    *oci_load_balancer.GetLoadBalancerHealthResponse
}

func (s *LoadBalancerLoadBalancerHealthDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LoadBalancerLoadBalancerHealthDataSourceCrud) Get() error {
	request := oci_load_balancer.GetLoadBalancerHealthRequest{}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "load_balancer")

	response, err := s.Client.GetLoadBalancerHealth(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *LoadBalancerLoadBalancerHealthDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("LoadBalancerLoadBalancerHealthDataSource-", LoadBalancerLoadBalancerHealthDataSource(), s.D))

	s.D.Set("critical_state_backend_set_names", s.Res.CriticalStateBackendSetNames)

	s.D.Set("status", s.Res.Status)

	if s.Res.TotalBackendSetCount != nil {
		s.D.Set("total_backend_set_count", *s.Res.TotalBackendSetCount)
	}

	s.D.Set("unknown_state_backend_set_names", s.Res.UnknownStateBackendSetNames)

	s.D.Set("warning_state_backend_set_names", s.Res.WarningStateBackendSetNames)

	return nil
}
