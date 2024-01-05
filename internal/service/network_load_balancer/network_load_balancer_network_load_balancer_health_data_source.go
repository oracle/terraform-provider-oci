// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package network_load_balancer

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_network_load_balancer "github.com/oracle/oci-go-sdk/v65/networkloadbalancer"
)

func NetworkLoadBalancerNetworkLoadBalancerHealthDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularNetworkLoadBalancerNetworkLoadBalancerHealth,
		Schema: map[string]*schema.Schema{
			"network_load_balancer_id": {
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

func readSingularNetworkLoadBalancerNetworkLoadBalancerHealth(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkLoadBalancerNetworkLoadBalancerHealthDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkLoadBalancerClient()

	return tfresource.ReadResource(sync)
}

type NetworkLoadBalancerNetworkLoadBalancerHealthDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_network_load_balancer.NetworkLoadBalancerClient
	Res    *oci_network_load_balancer.GetNetworkLoadBalancerHealthResponse
}

func (s *NetworkLoadBalancerNetworkLoadBalancerHealthDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *NetworkLoadBalancerNetworkLoadBalancerHealthDataSourceCrud) Get() error {
	request := oci_network_load_balancer.GetNetworkLoadBalancerHealthRequest{}

	if networkLoadBalancerId, ok := s.D.GetOkExists("network_load_balancer_id"); ok {
		tmp := networkLoadBalancerId.(string)
		request.NetworkLoadBalancerId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "network_load_balancer")

	response, err := s.Client.GetNetworkLoadBalancerHealth(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *NetworkLoadBalancerNetworkLoadBalancerHealthDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("NetworkLoadBalancerNetworkLoadBalancerHealthDataSource-", NetworkLoadBalancerNetworkLoadBalancerHealthDataSource(), s.D))

	s.D.Set("critical_state_backend_set_names", s.Res.CriticalStateBackendSetNames)

	s.D.Set("status", s.Res.Status)

	if s.Res.TotalBackendSetCount != nil {
		s.D.Set("total_backend_set_count", *s.Res.TotalBackendSetCount)
	}

	s.D.Set("unknown_state_backend_set_names", s.Res.UnknownStateBackendSetNames)

	s.D.Set("warning_state_backend_set_names", s.Res.WarningStateBackendSetNames)

	return nil
}
