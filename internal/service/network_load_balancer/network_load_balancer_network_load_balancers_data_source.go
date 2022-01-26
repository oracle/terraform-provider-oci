// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package network_load_balancer

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_network_load_balancer "github.com/oracle/oci-go-sdk/v56/networkloadbalancer"
)

func NetworkLoadBalancerNetworkLoadBalancersDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readNetworkLoadBalancerNetworkLoadBalancers,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_load_balancer_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(NetworkLoadBalancerNetworkLoadBalancerResource()),
						},
					},
				},
			},
		},
	}
}

func readNetworkLoadBalancerNetworkLoadBalancers(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkLoadBalancerNetworkLoadBalancersDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkLoadBalancerClient()

	return tfresource.ReadResource(sync)
}

type NetworkLoadBalancerNetworkLoadBalancersDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_network_load_balancer.NetworkLoadBalancerClient
	Res    *oci_network_load_balancer.ListNetworkLoadBalancersResponse
}

func (s *NetworkLoadBalancerNetworkLoadBalancersDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *NetworkLoadBalancerNetworkLoadBalancersDataSourceCrud) Get() error {
	request := oci_network_load_balancer.ListNetworkLoadBalancersRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_network_load_balancer.ListNetworkLoadBalancersLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "network_load_balancer")

	response, err := s.Client.ListNetworkLoadBalancers(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListNetworkLoadBalancers(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *NetworkLoadBalancerNetworkLoadBalancersDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("NetworkLoadBalancerNetworkLoadBalancersDataSource-", NetworkLoadBalancerNetworkLoadBalancersDataSource(), s.D))
	resources := []map[string]interface{}{}
	networkLoadBalancer := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, NetworkLoadBalancerSummaryToMap(item, true))
	}
	networkLoadBalancer["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, NetworkLoadBalancerNetworkLoadBalancersDataSource().Schema["network_load_balancer_collection"].Elem.(*schema.Resource).Schema)
		networkLoadBalancer["items"] = items
	}

	resources = append(resources, networkLoadBalancer)
	if err := s.D.Set("network_load_balancer_collection", resources); err != nil {
		return err
	}

	return nil
}
