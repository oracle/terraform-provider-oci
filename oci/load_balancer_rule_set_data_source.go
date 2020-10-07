// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_load_balancer "github.com/oracle/oci-go-sdk/v27/loadbalancer"
)

func init() {
	RegisterDatasource("oci_load_balancer_rule_set", LoadBalancerRuleSetDataSource())
}

func LoadBalancerRuleSetDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["load_balancer_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["name"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(LoadBalancerRuleSetResource(), fieldMap, readSingularLoadBalancerRuleSet)
}

func readSingularLoadBalancerRuleSet(d *schema.ResourceData, m interface{}) error {
	sync := &LoadBalancerRuleSetDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient()

	return ReadResource(sync)
}

type LoadBalancerRuleSetDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_load_balancer.LoadBalancerClient
	Res    *oci_load_balancer.GetRuleSetResponse
}

func (s *LoadBalancerRuleSetDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LoadBalancerRuleSetDataSourceCrud) Get() error {
	request := oci_load_balancer.GetRuleSetRequest{}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.RuleSetName = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "load_balancer")

	response, err := s.Client.GetRuleSet(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *LoadBalancerRuleSetDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("LoadBalancerRuleSetDataSource-", LoadBalancerRuleSetDataSource(), s.D))

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, RuleToMap(item, true))
	}
	s.D.Set("items", items)

	return nil
}
