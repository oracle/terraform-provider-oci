// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_load_balancer "github.com/oracle/oci-go-sdk/loadbalancer"

	"github.com/oracle/terraform-provider-oci/crud"
)

func LoadBalancerProtocolsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readLoadBalancerProtocols,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"protocols": {
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

func readLoadBalancerProtocols(d *schema.ResourceData, m interface{}) error {
	sync := &LoadBalancerProtocolsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient

	return crud.ReadResource(sync)
}

type LoadBalancerProtocolsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_load_balancer.LoadBalancerClient
	Res    *oci_load_balancer.ListProtocolsResponse
}

func (s *LoadBalancerProtocolsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LoadBalancerProtocolsDataSourceCrud) Get() error {
	request := oci_load_balancer.ListProtocolsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "load_balancer")

	response, err := s.Client.ListProtocols(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *LoadBalancerProtocolsDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		loadBalancerProtocol := map[string]interface{}{}

		if r.Name != nil {
			loadBalancerProtocol["name"] = *r.Name
		}

		resources = append(resources, loadBalancerProtocol)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, LoadBalancerProtocolsDataSource().Schema["protocols"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("protocols", resources); err != nil {
		panic(err)
	}

	return
}
