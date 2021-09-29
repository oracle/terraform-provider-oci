// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_network_load_balancer "github.com/oracle/oci-go-sdk/v48/networkloadbalancer"
)

func init() {
	RegisterDatasource("oci_network_load_balancer_listener", NetworkLoadBalancerListenerDataSource())
}

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
	return GetSingularDataSourceItemSchema(NetworkLoadBalancerListenerResource(), fieldMap, readSingularNetworkLoadBalancerListener)
}

func readSingularNetworkLoadBalancerListener(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkLoadBalancerListenerDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).networkLoadBalancerClient()

	return ReadResource(sync)
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

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "network_load_balancer")

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

	s.D.SetId(GenerateDataSourceHashID("NetworkLoadBalancerListenerDataSource-", NetworkLoadBalancerListenerDataSource(), s.D))

	if s.Res.DefaultBackendSetName != nil {
		s.D.Set("default_backend_set_name", *s.Res.DefaultBackendSetName)
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
