// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package network_load_balancer

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_network_load_balancer "github.com/oracle/oci-go-sdk/v65/networkloadbalancer"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func NetworkLoadBalancerNetworkLoadBalancerBackendSetBackendOperationalStatusDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularNetworkLoadBalancerNetworkLoadBalancerBackendSetBackendOperationalStatus,
		Schema: map[string]*schema.Schema{
			"backend_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"backend_set_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"network_load_balancer_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularNetworkLoadBalancerNetworkLoadBalancerBackendSetBackendOperationalStatus(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkLoadBalancerNetworkLoadBalancerBackendSetBackendOperationalStatusDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkLoadBalancerClient()

	return tfresource.ReadResource(sync)
}

type NetworkLoadBalancerNetworkLoadBalancerBackendSetBackendOperationalStatusDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_network_load_balancer.NetworkLoadBalancerClient
	Res    *oci_network_load_balancer.GetBackendOperationalStatusResponse
}

func (s *NetworkLoadBalancerNetworkLoadBalancerBackendSetBackendOperationalStatusDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *NetworkLoadBalancerNetworkLoadBalancerBackendSetBackendOperationalStatusDataSourceCrud) Get() error {
	request := oci_network_load_balancer.GetBackendOperationalStatusRequest{}

	if backendName, ok := s.D.GetOkExists("backend_name"); ok {
		tmp := backendName.(string)
		request.BackendName = &tmp
	}

	if backendSetName, ok := s.D.GetOkExists("backend_set_name"); ok {
		tmp := backendSetName.(string)
		request.BackendSetName = &tmp
	}

	if networkLoadBalancerId, ok := s.D.GetOkExists("network_load_balancer_id"); ok {
		tmp := networkLoadBalancerId.(string)
		request.NetworkLoadBalancerId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "network_load_balancer")

	response, err := s.Client.GetBackendOperationalStatus(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *NetworkLoadBalancerNetworkLoadBalancerBackendSetBackendOperationalStatusDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("NetworkLoadBalancerNetworkLoadBalancerBackendSetBackendOperationalStatusDataSource-", NetworkLoadBalancerNetworkLoadBalancerBackendSetBackendOperationalStatusDataSource(), s.D))

	s.D.Set("status", s.Res.Status)

	return nil
}
