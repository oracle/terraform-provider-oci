// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_load_balancer "github.com/oracle/oci-go-sdk/loadbalancer"

	"github.com/oracle/terraform-provider-oci/crud"
)

func LoadBalancerShapesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readLoadBalancerShapes,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"shapes": {
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

func readLoadBalancerShapes(d *schema.ResourceData, m interface{}) error {
	sync := &LoadBalancerShapesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient

	return crud.ReadResource(sync)
}

type LoadBalancerShapesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_load_balancer.LoadBalancerClient
	Res    *oci_load_balancer.ListShapesResponse
}

func (s *LoadBalancerShapesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LoadBalancerShapesDataSourceCrud) Get() error {
	request := oci_load_balancer.ListShapesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "load_balancer")

	response, err := s.Client.ListShapes(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *LoadBalancerShapesDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		loadBalancerShape := map[string]interface{}{}

		if r.Name != nil {
			loadBalancerShape["name"] = *r.Name
		}

		resources = append(resources, loadBalancerShape)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, LoadBalancerShapesDataSource().Schema["shapes"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("shapes", resources); err != nil {
		panic(err)
	}

	return
}
