// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package load_balancer

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_load_balancer "github.com/oracle/oci-go-sdk/v65/loadbalancer"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func LoadBalancerLoadBalancerProtocolsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readLoadBalancerLoadBalancerProtocols,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
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

func readLoadBalancerLoadBalancerProtocols(d *schema.ResourceData, m interface{}) error {
	sync := &LoadBalancerLoadBalancerProtocolsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LoadBalancerClient()

	return tfresource.ReadResource(sync)
}

type LoadBalancerLoadBalancerProtocolsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_load_balancer.LoadBalancerClient
	Res    *oci_load_balancer.ListProtocolsResponse
}

func (s *LoadBalancerLoadBalancerProtocolsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LoadBalancerLoadBalancerProtocolsDataSourceCrud) Get() error {
	request := oci_load_balancer.ListProtocolsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "load_balancer")

	response, err := s.Client.ListProtocols(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *LoadBalancerLoadBalancerProtocolsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("LoadBalancerLoadBalancerProtocolsDataSource-", LoadBalancerLoadBalancerProtocolsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		loadBalancerProtocol := map[string]interface{}{}

		if r.Name != nil {
			loadBalancerProtocol["name"] = *r.Name
		}

		resources = append(resources, loadBalancerProtocol)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, LoadBalancerLoadBalancerProtocolsDataSource().Schema["protocols"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("protocols", resources); err != nil {
		return err
	}

	return nil
}
