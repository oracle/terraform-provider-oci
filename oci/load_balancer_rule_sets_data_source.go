// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_load_balancer "github.com/oracle/oci-go-sdk/loadbalancer"
)

func RuleSetsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readRuleSets,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"load_balancer_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"rule_sets": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(RuleSetResource()),
			},
		},
	}
}

func readRuleSets(d *schema.ResourceData, m interface{}) error {
	sync := &RuleSetsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient

	return ReadResource(sync)
}

type RuleSetsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_load_balancer.LoadBalancerClient
	Res    *oci_load_balancer.ListRuleSetsResponse
}

func (s *RuleSetsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *RuleSetsDataSourceCrud) Get() error {
	request := oci_load_balancer.ListRuleSetsRequest{}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "load_balancer")

	response, err := s.Client.ListRuleSets(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *RuleSetsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		ruleSet := map[string]interface{}{}

		items := []interface{}{}
		for _, item := range r.Items {
			items = append(items, RuleToMap(item))
		}
		ruleSet["items"] = items

		if r.Name != nil {
			ruleSet["name"] = *r.Name
		}

		resources = append(resources, ruleSet)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, RuleSetsDataSource().Schema["rule_sets"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("rule_sets", resources); err != nil {
		return err
	}

	return nil
}
