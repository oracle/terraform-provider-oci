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

func NetworkLoadBalancerListenerDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["listener_name"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["network_load_balancer_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(NetworkLoadBalancerListenerResource(), fieldMap, readSingularNetworkLoadBalancerListener)
}

func readSingularNetworkLoadBalancerListener(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkLoadBalancerListenerDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkLoadBalancerClient()

	return tfresource.ReadResource(sync)
}

type NetworkLoadBalancerListenerDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_network_load_balancer.NetworkLoadBalancerClient
	Res    *oci_network_load_balancer.GetListenerResponse
}

func (s *NetworkLoadBalancerListenerDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *NetworkLoadBalancerListenerDataSourceCrud) Get() error {
	request := oci_network_load_balancer.GetListenerRequest{}

	if listenerName, ok := s.D.GetOkExists("listener_name"); ok {
		tmp := listenerName.(string)
		request.ListenerName = &tmp
	}

	if networkLoadBalancerId, ok := s.D.GetOkExists("network_load_balancer_id"); ok {
		tmp := networkLoadBalancerId.(string)
		request.NetworkLoadBalancerId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "network_load_balancer")

	response, err := s.Client.GetListener(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *NetworkLoadBalancerListenerDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("NetworkLoadBalancerListenerDataSource-", NetworkLoadBalancerListenerDataSource(), s.D))

	if s.Res.DefaultBackendSetName != nil {
		s.D.Set("default_backend_set_name", *s.Res.DefaultBackendSetName)
	}
	s.D.Set("ip_version", s.Res.IpVersion)

	if s.Res.IsPpv2Enabled != nil {
		s.D.Set("is_ppv2enabled", *s.Res.IsPpv2Enabled)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.Port != nil {
		s.D.Set("port", *s.Res.Port)
	}

	s.D.Set("protocol", s.Res.Protocol)

	return nil
}
