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

func LoadBalancerLoadBalancerRoutingPoliciesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readLoadBalancerLoadBalancerRoutingPolicies,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"load_balancer_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"routing_policies": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     LoadBalancerLoadBalancerRoutingPolicyResource(),
			},
		},
	}
}

func readLoadBalancerLoadBalancerRoutingPolicies(d *schema.ResourceData, m interface{}) error {
	sync := &LoadBalancerLoadBalancerRoutingPoliciesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LoadBalancerClient()

	return tfresource.ReadResource(sync)
}

type LoadBalancerLoadBalancerRoutingPoliciesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_load_balancer.LoadBalancerClient
	Res    *oci_load_balancer.ListRoutingPoliciesResponse
}

func (s *LoadBalancerLoadBalancerRoutingPoliciesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LoadBalancerLoadBalancerRoutingPoliciesDataSourceCrud) Get() error {
	request := oci_load_balancer.ListRoutingPoliciesRequest{}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "load_balancer")

	response, err := s.Client.ListRoutingPolicies(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListRoutingPolicies(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *LoadBalancerLoadBalancerRoutingPoliciesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("LoadBalancerLoadBalancerRoutingPoliciesDataSource-", LoadBalancerLoadBalancerRoutingPoliciesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		loadBalancerRoutingPolicy := map[string]interface{}{}

		loadBalancerRoutingPolicy["condition_language_version"] = r.ConditionLanguageVersion

		if r.Name != nil {
			loadBalancerRoutingPolicy["name"] = *r.Name
		}

		rules := []interface{}{}
		for _, item := range r.Rules {
			rules = append(rules, RoutingRuleToMap(item))
		}
		loadBalancerRoutingPolicy["rules"] = rules

		resources = append(resources, loadBalancerRoutingPolicy)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, LoadBalancerLoadBalancerRoutingPoliciesDataSource().Schema["routing_policies"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("routing_policies", resources); err != nil {
		return err
	}

	return nil
}
