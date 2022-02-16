// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package load_balancer

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_load_balancer "github.com/oracle/oci-go-sdk/v58/loadbalancer"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func LoadBalancerRuleSetsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readLoadBalancerRuleSets,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"load_balancer_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"rule_sets": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(LoadBalancerRuleSetResource()),
			},
		},
	}
}

func readLoadBalancerRuleSets(d *schema.ResourceData, m interface{}) error {
	sync := &LoadBalancerRuleSetsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LoadBalancerClient()

	return tfresource.ReadResource(sync)
}

type LoadBalancerRuleSetsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_load_balancer.LoadBalancerClient
	Res    *oci_load_balancer.ListRuleSetsResponse
}

func (s *LoadBalancerRuleSetsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LoadBalancerRuleSetsDataSourceCrud) Get() error {
	request := oci_load_balancer.ListRuleSetsRequest{}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "load_balancer")

	response, err := s.Client.ListRuleSets(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *LoadBalancerRuleSetsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("LoadBalancerRuleSetsDataSource-", LoadBalancerRuleSetsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		ruleSet := map[string]interface{}{}

		items := []interface{}{}
		for _, item := range r.Items {
			items = append(items, RuleToMap(item, true))
		}
		ruleSet["items"] = items

		if r.Name != nil {
			ruleSet["name"] = *r.Name
		}

		resources = append(resources, ruleSet)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, LoadBalancerRuleSetsDataSource().Schema["rule_sets"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("rule_sets", resources); err != nil {
		return err
	}

	return nil
}
