// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_network_load_balancer "github.com/oracle/oci-go-sdk/v46/networkloadbalancer"
)

func init() {
	RegisterDatasource("oci_network_load_balancer_backend_set_health", NetworkLoadBalancerBackendSetHealthDataSource())
}

func NetworkLoadBalancerBackendSetHealthDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularNetworkLoadBalancerBackendSetHealth,
		Schema: map[string]*schema.Schema{
			"backend_set_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"network_load_balancer_id": {
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

func readSingularNetworkLoadBalancerBackendSetHealth(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkLoadBalancerBackendSetHealthDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).networkLoadBalancerClient()

	return ReadResource(sync)
}

type NetworkLoadBalancerBackendSetHealthDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_network_load_balancer.NetworkLoadBalancerClient
	Res    *oci_network_load_balancer.GetBackendSetHealthResponse
}

func (s *NetworkLoadBalancerBackendSetHealthDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *NetworkLoadBalancerBackendSetHealthDataSourceCrud) Get() error {
	request := oci_network_load_balancer.GetBackendSetHealthRequest{}

	if backendSetName, ok := s.D.GetOkExists("backend_set_name"); ok {
		tmp := backendSetName.(string)
		request.BackendSetName = &tmp
	}

	if networkLoadBalancerId, ok := s.D.GetOkExists("network_load_balancer_id"); ok {
		tmp := networkLoadBalancerId.(string)
		request.NetworkLoadBalancerId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "network_load_balancer")

	response, err := s.Client.GetBackendSetHealth(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *NetworkLoadBalancerBackendSetHealthDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("NetworkLoadBalancerBackendSetHealthDataSource-", NetworkLoadBalancerBackendSetHealthDataSource(), s.D))

	s.D.Set("critical_state_backend_names", s.Res.CriticalStateBackendNames)

	s.D.Set("status", s.Res.Status)

	if s.Res.TotalBackendCount != nil {
		s.D.Set("total_backend_count", *s.Res.TotalBackendCount)
	}

	s.D.Set("unknown_state_backend_names", s.Res.UnknownStateBackendNames)

	s.D.Set("warning_state_backend_names", s.Res.WarningStateBackendNames)

	return nil
}
