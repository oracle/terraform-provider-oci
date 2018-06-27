// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_load_balancer "github.com/oracle/oci-go-sdk/loadbalancer"

	"github.com/oracle/terraform-provider-oci/crud"
)

func LoadBalancerPoliciesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readLoadBalancerPolicies,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"policies": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readLoadBalancerPolicies(d *schema.ResourceData, m interface{}) error {
	sync := &LoadBalancerPoliciesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient

	return crud.ReadResource(sync)
}

type LoadBalancerPoliciesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_load_balancer.LoadBalancerClient
	Res    *oci_load_balancer.ListPoliciesResponse
}

func (s *LoadBalancerPoliciesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LoadBalancerPoliciesDataSourceCrud) Get() error {
	request := oci_load_balancer.ListPoliciesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "load_balancer")

	response, err := s.Client.ListPolicies(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *LoadBalancerPoliciesDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		loadBalancerPolicy := map[string]interface{}{}

		if r.Name != nil {
			loadBalancerPolicy["name"] = *r.Name
		}

		resources = append(resources, loadBalancerPolicy)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, LoadBalancerPoliciesDataSource().Schema["policies"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("policies", resources); err != nil {
		panic(err)
	}

	return
}
