// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_load_balancer "github.com/oracle/oci-go-sdk/loadbalancer"

	"github.com/oracle/terraform-provider-oci/crud"
)

func LoadBalancerHealthDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularLoadBalancerHealth,
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

func readSingularLoadBalancerHealth(d *schema.ResourceData, m interface{}) error {
	sync := &LoadBalancerHealthDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient

	return crud.ReadResource(sync)
}

type LoadBalancerHealthDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_load_balancer.LoadBalancerClient
	Res    *oci_load_balancer.GetLoadBalancerHealthResponse
}

func (s *LoadBalancerHealthDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LoadBalancerHealthDataSourceCrud) Get() error {
	request := oci_load_balancer.GetLoadBalancerHealthRequest{}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "load_balancer")

	response, err := s.Client.GetLoadBalancerHealth(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *LoadBalancerHealthDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())

	s.D.Set("critical_state_backend_set_names", s.Res.CriticalStateBackendSetNames)

	s.D.Set("status", s.Res.Status)

	if s.Res.TotalBackendSetCount != nil {
		s.D.Set("total_backend_set_count", *s.Res.TotalBackendSetCount)
	}

	s.D.Set("unknown_state_backend_set_names", s.Res.UnknownStateBackendSetNames)

	s.D.Set("warning_state_backend_set_names", s.Res.WarningStateBackendSetNames)

	return
}
