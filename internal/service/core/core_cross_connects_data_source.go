// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v58/core"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func CoreCrossConnectsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreCrossConnects,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cross_connect_group_id": {
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
			"cross_connects": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(CoreCrossConnectResource()),
			},
		},
	}
}

func readCoreCrossConnects(d *schema.ResourceData, m interface{}) error {
	sync := &CoreCrossConnectsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

type CoreCrossConnectsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.ListCrossConnectsResponse
}

func (s *CoreCrossConnectsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreCrossConnectsDataSourceCrud) Get() error {
	request := oci_core.ListCrossConnectsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if crossConnectGroupId, ok := s.D.GetOkExists("cross_connect_group_id"); ok {
		tmp := crossConnectGroupId.(string)
		request.CrossConnectGroupId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_core.CrossConnectLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.ListCrossConnects(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListCrossConnects(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreCrossConnectsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreCrossConnectsDataSource-", CoreCrossConnectsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		crossConnect := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.CrossConnectGroupId != nil {
			crossConnect["cross_connect_group_id"] = *r.CrossConnectGroupId
		}

		if r.CustomerReferenceName != nil {
			crossConnect["customer_reference_name"] = *r.CustomerReferenceName
		}

		if r.DefinedTags != nil {
			crossConnect["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			crossConnect["display_name"] = *r.DisplayName
		}

		crossConnect["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			crossConnect["id"] = *r.Id
		}

		if r.LocationName != nil {
			crossConnect["location_name"] = *r.LocationName
		}

		if r.MacsecProperties != nil {
			crossConnect["macsec_properties"] = []interface{}{MacsecPropertiesToMap(r.MacsecProperties)}
		} else {
			crossConnect["macsec_properties"] = nil
		}

		if r.PortName != nil {
			crossConnect["port_name"] = *r.PortName
		}

		if r.PortSpeedShapeName != nil {
			crossConnect["port_speed_shape_name"] = *r.PortSpeedShapeName
		}

		crossConnect["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			crossConnect["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, crossConnect)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, CoreCrossConnectsDataSource().Schema["cross_connects"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("cross_connects", resources); err != nil {
		return err
	}

	return nil
}
