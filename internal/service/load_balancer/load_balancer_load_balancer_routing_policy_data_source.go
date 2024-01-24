// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package load_balancer

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_load_balancer "github.com/oracle/oci-go-sdk/v65/loadbalancer"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func LoadBalancerLoadBalancerRoutingPolicyDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["load_balancer_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["routing_policy_name"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(LoadBalancerLoadBalancerRoutingPolicyResource(), fieldMap, readSingularLoadBalancerLoadBalancerRoutingPolicy)
}

func readSingularLoadBalancerLoadBalancerRoutingPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &LoadBalancerLoadBalancerRoutingPolicyDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LoadBalancerClient()

	return tfresource.ReadResource(sync)
}

type LoadBalancerLoadBalancerRoutingPolicyDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_load_balancer.LoadBalancerClient
	Res    *oci_load_balancer.GetRoutingPolicyResponse
}

func (s *LoadBalancerLoadBalancerRoutingPolicyDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LoadBalancerLoadBalancerRoutingPolicyDataSourceCrud) Get() error {
	request := oci_load_balancer.GetRoutingPolicyRequest{}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	if routingPolicyName, ok := s.D.GetOkExists("routing_policy_name"); ok {
		tmp := routingPolicyName.(string)
		request.RoutingPolicyName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "load_balancer")

	response, err := s.Client.GetRoutingPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *LoadBalancerLoadBalancerRoutingPolicyDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("LoadBalancerLoadBalancerRoutingPolicyDataSource-", LoadBalancerLoadBalancerRoutingPolicyDataSource(), s.D))

	s.D.Set("condition_language_version", s.Res.ConditionLanguageVersion)

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	rules := []interface{}{}
	for _, item := range s.Res.Rules {
		rules = append(rules, RoutingRuleToMap(item))
	}
	s.D.Set("rules", rules)

	return nil
}
