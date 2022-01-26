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

func LoadBalancerBackendSetHealthDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularLoadBalancerBackendSetHealth,
		Schema: map[string]*schema.Schema{
			"backend_set_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"load_balancer_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"critical_state_backend_names": {
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
			"total_backend_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"unknown_state_backend_names": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"warning_state_backend_names": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func readSingularLoadBalancerBackendSetHealth(d *schema.ResourceData, m interface{}) error {
	sync := &LoadBalancerBackendSetHealthDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LoadBalancerClient()

	return tfresource.ReadResource(sync)
}

type LoadBalancerBackendSetHealthDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_load_balancer.LoadBalancerClient
	Res    *oci_load_balancer.GetBackendSetHealthResponse
}

func (s *LoadBalancerBackendSetHealthDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LoadBalancerBackendSetHealthDataSourceCrud) Get() error {
	request := oci_load_balancer.GetBackendSetHealthRequest{}

	if backendSetName, ok := s.D.GetOkExists("backend_set_name"); ok {
		tmp := backendSetName.(string)
		request.BackendSetName = &tmp
	}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "load_balancer")

	response, err := s.Client.GetBackendSetHealth(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *LoadBalancerBackendSetHealthDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("LoadBalancerBackendSetHealthDataSource-", LoadBalancerBackendSetHealthDataSource(), s.D))

	s.D.Set("critical_state_backend_names", s.Res.CriticalStateBackendNames)

	s.D.Set("status", s.Res.Status)

	if s.Res.TotalBackendCount != nil {
		s.D.Set("total_backend_count", *s.Res.TotalBackendCount)
	}

	s.D.Set("unknown_state_backend_names", s.Res.UnknownStateBackendNames)

	s.D.Set("warning_state_backend_names", s.Res.WarningStateBackendNames)

	return nil
}
