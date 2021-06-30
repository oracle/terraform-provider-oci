// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_network_load_balancer "github.com/oracle/oci-go-sdk/v43/networkloadbalancer"
)

func init() {
	RegisterDatasource("oci_network_load_balancer_network_load_balancers_policies", NetworkLoadBalancerNetworkLoadBalancersPoliciesDataSource())
}

func NetworkLoadBalancerNetworkLoadBalancersPoliciesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readNetworkLoadBalancerNetworkLoadBalancersPolicies,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"network_load_balancers_policy_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					// Required

					// Optional

					// Computed
					Schema: map[string]*schema.Schema{
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func readNetworkLoadBalancerNetworkLoadBalancersPolicies(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkLoadBalancerNetworkLoadBalancersPoliciesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).networkLoadBalancerClient()

	return ReadResource(sync)
}

type NetworkLoadBalancerNetworkLoadBalancersPoliciesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_network_load_balancer.NetworkLoadBalancerClient
	Res    *oci_network_load_balancer.ListNetworkLoadBalancersPoliciesResponse
}

func (s *NetworkLoadBalancerNetworkLoadBalancersPoliciesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *NetworkLoadBalancerNetworkLoadBalancersPoliciesDataSourceCrud) Get() error {
	request := oci_network_load_balancer.ListNetworkLoadBalancersPoliciesRequest{}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "network_load_balancer")

	response, err := s.Client.ListNetworkLoadBalancersPolicies(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response

	return nil
}

func (s *NetworkLoadBalancerNetworkLoadBalancersPoliciesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("NetworkLoadBalancerNetworkLoadBalancersPoliciesDataSource-", NetworkLoadBalancerNetworkLoadBalancersPoliciesDataSource(), s.D))
	resources := []map[string]interface{}{}
	networkLoadBalancersPolicy := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, item)
	}
	networkLoadBalancersPolicy["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = ApplyFiltersInCollection(f.(*schema.Set), items, NetworkLoadBalancerNetworkLoadBalancersPoliciesDataSource().Schema["network_load_balancers_policy_collection"].Elem.(*schema.Resource).Schema)
		networkLoadBalancersPolicy["items"] = items
	}

	resources = append(resources, networkLoadBalancersPolicy)
	if err := s.D.Set("network_load_balancers_policy_collection", resources); err != nil {
		return err
	}

	return nil
}
