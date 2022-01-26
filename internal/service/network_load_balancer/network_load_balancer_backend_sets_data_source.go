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

func NetworkLoadBalancerBackendSetsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readNetworkLoadBalancerBackendSets,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"network_load_balancer_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"backend_set_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(NetworkLoadBalancerBackendSetResource()),
						},
					},
				},
			},
		},
	}
}

func readNetworkLoadBalancerBackendSets(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkLoadBalancerBackendSetsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkLoadBalancerClient()

	return tfresource.ReadResource(sync)
}

type NetworkLoadBalancerBackendSetsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_network_load_balancer.NetworkLoadBalancerClient
	Res    *oci_network_load_balancer.ListBackendSetsResponse
}

func (s *NetworkLoadBalancerBackendSetsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *NetworkLoadBalancerBackendSetsDataSourceCrud) Get() error {
	request := oci_network_load_balancer.ListBackendSetsRequest{}

	if networkLoadBalancerId, ok := s.D.GetOkExists("network_load_balancer_id"); ok {
		tmp := networkLoadBalancerId.(string)
		request.NetworkLoadBalancerId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "network_load_balancer")

	response, err := s.Client.ListBackendSets(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListBackendSets(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *NetworkLoadBalancerBackendSetsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("NetworkLoadBalancerBackendSetsDataSource-", NetworkLoadBalancerBackendSetsDataSource(), s.D))
	resources := []map[string]interface{}{}
	backendSet := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, NlbBackendSetSummaryToMap(item))
	}
	backendSet["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, NetworkLoadBalancerBackendSetsDataSource().Schema["backend_set_collection"].Elem.(*schema.Resource).Schema)
		backendSet["items"] = items
	}

	resources = append(resources, backendSet)
	if err := s.D.Set("backend_set_collection", resources); err != nil {
		return err
	}

	return nil
}
