// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/core"

	"github.com/oracle/terraform-provider-oci/crud"
)

func IpSecConnectionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readIpSecConnections,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cpe_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"drg_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": {
				Type:       schema.TypeInt,
				Optional:   true,
				Deprecated: crud.FieldDeprecated("limit"),
			},
			"page": {
				Type:       schema.TypeString,
				Optional:   true,
				Deprecated: crud.FieldDeprecated("page"),
			},
			"connections": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     IpSecConnectionResource(),
			},
		},
	}
}

func readIpSecConnections(d *schema.ResourceData, m interface{}) error {
	sync := &IpSecConnectionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return crud.ReadResource(sync)
}

type IpSecConnectionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.ListIPSecConnectionsResponse
}

func (s *IpSecConnectionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IpSecConnectionsDataSourceCrud) Get() error {
	request := oci_core.ListIPSecConnectionsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if cpeId, ok := s.D.GetOkExists("cpe_id"); ok {
		tmp := cpeId.(string)
		request.CpeId = &tmp
	}

	if drgId, ok := s.D.GetOkExists("drg_id"); ok {
		tmp := drgId.(string)
		request.DrgId = &tmp
	}

	if limit, ok := s.D.GetOkExists("limit"); ok {
		tmp := limit.(int)
		request.Limit = &tmp
	}

	if page, ok := s.D.GetOkExists("page"); ok {
		tmp := page.(string)
		request.Page = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

	response, err := s.Client.ListIPSecConnections(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListIPSecConnections(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *IpSecConnectionsDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		ipSecConnection := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.CpeId != nil {
			ipSecConnection["cpe_id"] = *r.CpeId
		}

		if r.DefinedTags != nil {
			ipSecConnection["defined_tags"] = definedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			ipSecConnection["display_name"] = *r.DisplayName
		}

		if r.DrgId != nil {
			ipSecConnection["drg_id"] = *r.DrgId
		}

		ipSecConnection["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			ipSecConnection["id"] = *r.Id
		}

		ipSecConnection["state"] = r.LifecycleState

		ipSecConnection["static_routes"] = r.StaticRoutes

		if r.TimeCreated != nil {
			ipSecConnection["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, ipSecConnection)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, IpSecConnectionsDataSource().Schema["connections"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("connections", resources); err != nil {
		panic(err)
	}

	return
}
