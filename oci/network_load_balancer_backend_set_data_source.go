// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_network_load_balancer "github.com/oracle/oci-go-sdk/v43/networkloadbalancer"
)

func init() {
	RegisterDatasource("oci_network_load_balancer_backend_set", NetworkLoadBalancerBackendSetDataSource())
}

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
	return GetSingularDataSourceItemSchema(NetworkLoadBalancerBackendSetResource(), fieldMap, readSingularNetworkLoadBalancerBackendSet)
}

func readSingularNetworkLoadBalancerBackendSet(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkLoadBalancerBackendSetDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).networkLoadBalancerClient()

	return ReadResource(sync)
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

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "network_load_balancer")

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

	s.D.SetId(GenerateDataSourceHashID("NetworkLoadBalancerBackendSetDataSource-", NetworkLoadBalancerBackendSetDataSource(), s.D))

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

	if s.Res.IsPreserveSource != nil {
		s.D.Set("is_preserve_source", *s.Res.IsPreserveSource)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("policy", s.Res.Policy)

	return nil
}
