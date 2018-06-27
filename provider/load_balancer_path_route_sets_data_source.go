// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_load_balancer "github.com/oracle/oci-go-sdk/loadbalancer"

	"github.com/oracle/terraform-provider-oci/crud"
)

func PathRouteSetsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readPathRouteSets,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"load_balancer_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"path_route_sets": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     PathRouteSetResource(),
			},
		},
	}
}

func readPathRouteSets(d *schema.ResourceData, m interface{}) error {
	sync := &PathRouteSetsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient

	return crud.ReadResource(sync)
}

type PathRouteSetsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_load_balancer.LoadBalancerClient
	Res    *oci_load_balancer.ListPathRouteSetsResponse
}

func (s *PathRouteSetsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *PathRouteSetsDataSourceCrud) Get() error {
	request := oci_load_balancer.ListPathRouteSetsRequest{}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "load_balancer")

	response, err := s.Client.ListPathRouteSets(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *PathRouteSetsDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())
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
		resources = ApplyFilters(f.(*schema.Set), resources, PathRouteSetsDataSource().Schema["path_route_sets"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("path_route_sets", resources); err != nil {
		panic(err)
	}

	return
}
