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

func NetworkLoadBalancerNetworkLoadBalancersProtocolsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readNetworkLoadBalancerNetworkLoadBalancersProtocols,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"network_load_balancers_protocol_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func readNetworkLoadBalancerNetworkLoadBalancersProtocols(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkLoadBalancerNetworkLoadBalancersProtocolsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkLoadBalancerClient()

	return tfresource.ReadResource(sync)
}

type NetworkLoadBalancerNetworkLoadBalancersProtocolsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_network_load_balancer.NetworkLoadBalancerClient
	Res    *oci_network_load_balancer.ListNetworkLoadBalancersProtocolsResponse
}

func (s *NetworkLoadBalancerNetworkLoadBalancersProtocolsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *NetworkLoadBalancerNetworkLoadBalancersProtocolsDataSourceCrud) Get() error {
	request := oci_network_load_balancer.ListNetworkLoadBalancersProtocolsRequest{}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "network_load_balancer")

	response, err := s.Client.ListNetworkLoadBalancersProtocols(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response

	return nil
}

func (s *NetworkLoadBalancerNetworkLoadBalancersProtocolsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("NetworkLoadBalancerNetworkLoadBalancersProtocolsDataSource-", NetworkLoadBalancerNetworkLoadBalancersProtocolsDataSource(), s.D))
	resources := []map[string]interface{}{}
	networkLoadBalancersProtocol := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, item)
	}
	networkLoadBalancersProtocol["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, NetworkLoadBalancerNetworkLoadBalancersProtocolsDataSource().Schema["network_load_balancers_protocol_collection"].Elem.(*schema.Resource).Schema)
		networkLoadBalancersProtocol["items"] = items
	}

	resources = append(resources, networkLoadBalancersProtocol)
	if err := s.D.Set("network_load_balancers_protocol_collection", resources); err != nil {
		return err
	}

	return nil
}
