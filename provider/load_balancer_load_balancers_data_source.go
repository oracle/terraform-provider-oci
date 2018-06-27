// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"

	oci_load_balancer "github.com/oracle/oci-go-sdk/loadbalancer"

	"github.com/oracle/terraform-provider-oci/crud"
)

func LoadBalancersDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readLoadBalancers,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"detail": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"load_balancers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     LoadBalancerResource(),
			},
		},
	}
}

func readLoadBalancers(d *schema.ResourceData, m interface{}) error {
	sync := &LoadBalancersDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient

	return crud.ReadResource(sync)
}

type LoadBalancersDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_load_balancer.LoadBalancerClient
	Res    *oci_load_balancer.ListLoadBalancersResponse
}

func (s *LoadBalancersDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LoadBalancersDataSourceCrud) Get() error {
	request := oci_load_balancer.ListLoadBalancersRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if detail, ok := s.D.GetOkExists("detail"); ok {
		tmp := detail.(string)
		request.Detail = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_load_balancer.LoadBalancerLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "load_balancer")

	response, err := s.Client.ListLoadBalancers(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListLoadBalancers(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *LoadBalancersDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		loadBalancer := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DisplayName != nil {
			loadBalancer["display_name"] = *r.DisplayName
		}

		if r.Id != nil {
			loadBalancer["id"] = *r.Id
		}

		ipAddresses := []string{}
		for _, ad := range r.IpAddresses {
			if ad.IpAddress != nil {
				ipAddresses = append(ipAddresses, *ad.IpAddress)
			}
		}
		loadBalancer["ip_addresses"] = ipAddresses

		if r.IsPrivate != nil {
			loadBalancer["is_private"] = *r.IsPrivate
		}

		if r.ShapeName != nil {
			loadBalancer["shape"] = *r.ShapeName
		}

		loadBalancer["state"] = r.LifecycleState

		loadBalancer["subnet_ids"] = r.SubnetIds

		if r.TimeCreated != nil {
			loadBalancer["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, loadBalancer)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, LoadBalancersDataSource().Schema["load_balancers"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("load_balancers", resources); err != nil {
		panic(err)
	}

	return
}
