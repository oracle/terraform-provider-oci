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

func NetworkLoadBalancerBackendSetDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["backend_set_name"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["network_load_balancer_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(NetworkLoadBalancerBackendSetResource(), fieldMap, readSingularNetworkLoadBalancerBackendSet)
}

func readSingularNetworkLoadBalancerBackendSet(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkLoadBalancerBackendSetDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkLoadBalancerClient()

	return tfresource.ReadResource(sync)
}

type NetworkLoadBalancerBackendSetDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_network_load_balancer.NetworkLoadBalancerClient
	Res    *oci_network_load_balancer.GetBackendSetResponse
}

func (s *NetworkLoadBalancerBackendSetDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *NetworkLoadBalancerBackendSetDataSourceCrud) Get() error {
	request := oci_network_load_balancer.GetBackendSetRequest{}

	if backendSetName, ok := s.D.GetOkExists("backend_set_name"); ok {
		tmp := backendSetName.(string)
		request.BackendSetName = &tmp
	}

	if networkLoadBalancerId, ok := s.D.GetOkExists("network_load_balancer_id"); ok {
		tmp := networkLoadBalancerId.(string)
		request.NetworkLoadBalancerId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "network_load_balancer")

	response, err := s.Client.GetBackendSet(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *NetworkLoadBalancerBackendSetDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("NetworkLoadBalancerBackendSetDataSource-", NetworkLoadBalancerBackendSetDataSource(), s.D))

	backends := []interface{}{}
	for _, item := range s.Res.Backends {
		backends = append(backends, NlbBackendToMap(item))
	}
	s.D.Set("backends", backends)

	if s.Res.HealthChecker != nil {
		s.D.Set("health_checker", []interface{}{NlbHealthCheckerToMap(s.Res.HealthChecker)})
	} else {
		s.D.Set("health_checker", nil)
	}
	s.D.Set("ip_version", s.Res.IpVersion)

	if s.Res.IsFailOpen != nil {
		s.D.Set("is_fail_open", *s.Res.IsFailOpen)
	}

	if s.Res.IsPreserveSource != nil {
		s.D.Set("is_preserve_source", *s.Res.IsPreserveSource)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("policy", s.Res.Policy)

	return nil
}
