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

func LoadBalancerPathRouteSetsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readLoadBalancerPathRouteSets,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"load_balancer_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"path_route_sets": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     LoadBalancerPathRouteSetResource(),
			},
		},
	}
}

func readLoadBalancerPathRouteSets(d *schema.ResourceData, m interface{}) error {
	sync := &LoadBalancerPathRouteSetsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LoadBalancerClient()

	return tfresource.ReadResource(sync)
}

type LoadBalancerPathRouteSetsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_load_balancer.LoadBalancerClient
	Res    *oci_load_balancer.ListPathRouteSetsResponse
}

func (s *LoadBalancerPathRouteSetsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LoadBalancerPathRouteSetsDataSourceCrud) Get() error {
	request := oci_load_balancer.ListPathRouteSetsRequest{}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "load_balancer")

	response, err := s.Client.ListPathRouteSets(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *LoadBalancerPathRouteSetsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("LoadBalancerPathRouteSetsDataSource-", LoadBalancerPathRouteSetsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		pathRouteSet := map[string]interface{}{}

		if r.Name != nil {
			pathRouteSet["name"] = *r.Name
		}

		pathRoutes := []interface{}{}
		for _, item := range r.PathRoutes {
			pathRoutes = append(pathRoutes, PathRouteToMap(item))
		}
		pathRouteSet["path_routes"] = pathRoutes

		resources = append(resources, pathRouteSet)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, LoadBalancerPathRouteSetsDataSource().Schema["path_route_sets"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("path_route_sets", resources); err != nil {
		return err
	}

	return nil
}
