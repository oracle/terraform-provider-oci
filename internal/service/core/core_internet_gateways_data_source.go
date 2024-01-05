// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v65/core"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CoreInternetGatewaysDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreInternetGateways,
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
			"vcn_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"gateways": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(CoreInternetGatewayResource()),
			},
		},
	}
}

func readCoreInternetGateways(d *schema.ResourceData, m interface{}) error {
	sync := &CoreInternetGatewaysDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

type CoreInternetGatewaysDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.ListInternetGatewaysResponse
}

func (s *CoreInternetGatewaysDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreInternetGatewaysDataSourceCrud) Get() error {
	request := oci_core.ListInternetGatewaysRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_core.InternetGatewayLifecycleStateEnum(state.(string))
	}

	if vcnId, ok := s.D.GetOkExists("vcn_id"); ok {
		tmp := vcnId.(string)
		request.VcnId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.ListInternetGateways(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListInternetGateways(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreInternetGatewaysDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreInternetGatewaysDataSource-", CoreInternetGatewaysDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		internetGateway := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			internetGateway["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			internetGateway["display_name"] = *r.DisplayName
		}

		if r.IsEnabled != nil {
			internetGateway["enabled"] = *r.IsEnabled
		}

		internetGateway["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			internetGateway["id"] = *r.Id
		}

		if r.RouteTableId != nil {
			internetGateway["route_table_id"] = *r.RouteTableId
		}

		internetGateway["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			internetGateway["time_created"] = r.TimeCreated.String()
		}

		if r.VcnId != nil {
			internetGateway["vcn_id"] = *r.VcnId
		}

		resources = append(resources, internetGateway)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, CoreInternetGatewaysDataSource().Schema["gateways"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("gateways", resources); err != nil {
		return err
	}

	return nil
}
