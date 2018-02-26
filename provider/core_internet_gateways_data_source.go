// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/core"

	"github.com/oracle/terraform-provider-oci/crud"
)

func InternetGatewaysDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readInternetGateways,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
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
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcn_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Retaining the legacy name (internet_gateways => gateways) to avoid a breaking change.
			"gateways": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     InternetGatewayResource(),
			},
		},
	}
}

func readInternetGateways(d *schema.ResourceData, m interface{}) error {
	sync := &InternetGatewaysDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return crud.ReadResource(sync)
}

type InternetGatewaysDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.ListInternetGatewaysResponse
}

func (s *InternetGatewaysDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *InternetGatewaysDataSourceCrud) Get() error {
	request := oci_core.ListInternetGatewaysRequest{}

	if compartmentId, ok := s.D.GetOk("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOk("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if limit, ok := s.D.GetOk("limit"); ok {
		tmp := limit.(int)
		request.Limit = &tmp
	}

	if page, ok := s.D.GetOk("page"); ok {
		tmp := page.(string)
		request.Page = &tmp
	}

	if state, ok := s.D.GetOk("state"); ok {
		request.LifecycleState = oci_core.InternetGatewayLifecycleStateEnum(state.(string))
	}

	if vcnId, ok := s.D.GetOk("vcn_id"); ok {
		tmp := vcnId.(string)
		request.VcnId = &tmp
	}

	response, err := s.Client.ListInternetGateways(context.Background(), request, getRetryOptions(false, "core")...)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListInternetGateways(context.Background(), request, getRetryOptions(false, "core")...)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *InternetGatewaysDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		internetGateway := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
			"vcn_id":         *r.VcnId,
		}

		if r.DisplayName != nil {
			internetGateway["display_name"] = *r.DisplayName
		}

		if r.IsEnabled != nil {
			internetGateway["enabled"] = *r.IsEnabled
		}

		if r.Id != nil {
			internetGateway["id"] = *r.Id
		}

		internetGateway["state"] = r.LifecycleState

		internetGateway["time_created"] = r.TimeCreated.String()

		resources = append(resources, internetGateway)
	}

	if f, fOk := s.D.GetOk("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources)
	}

	// Retaining the legacy name (internet_gateways => gateways) to avoid a breaking change.
	if err := s.D.Set("gateways", resources); err != nil {
		panic(err)
	}

	return
}
