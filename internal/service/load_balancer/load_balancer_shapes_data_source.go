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

func LoadBalancerLoadBalancerShapesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readLoadBalancerLoadBalancerShapes,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
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

func readLoadBalancerLoadBalancerShapes(d *schema.ResourceData, m interface{}) error {
	sync := &LoadBalancerLoadBalancerShapesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LoadBalancerClient()

	return tfresource.ReadResource(sync)
}

type LoadBalancerLoadBalancerShapesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_load_balancer.LoadBalancerClient
	Res    *oci_load_balancer.ListShapesResponse
}

func (s *LoadBalancerLoadBalancerShapesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LoadBalancerLoadBalancerShapesDataSourceCrud) Get() error {
	request := oci_load_balancer.ListShapesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "load_balancer")

	response, err := s.Client.ListShapes(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *LoadBalancerLoadBalancerShapesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("LoadBalancerLoadBalancerShapesDataSource-", LoadBalancerLoadBalancerShapesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		loadBalancerShape := map[string]interface{}{}

		if r.Name != nil {
			loadBalancerShape["name"] = *r.Name
		}

		resources = append(resources, loadBalancerShape)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, LoadBalancerLoadBalancerShapesDataSource().Schema["shapes"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("shapes", resources); err != nil {
		return err
	}

	return nil
}
