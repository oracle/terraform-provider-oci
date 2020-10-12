// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v27/core"
)

func init() {
	RegisterDatasource("oci_core_ipsec_connections", CoreIpSecConnectionsDataSource())
}

func CoreIpSecConnectionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreIpSecConnections,
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
			"connections": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(CoreIpSecConnectionResource()),
			},
		},
	}
}

func readCoreIpSecConnections(d *schema.ResourceData, m interface{}) error {
	sync := &CoreIpSecConnectionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient()

	return ReadResource(sync)
}

type CoreIpSecConnectionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.ListIPSecConnectionsResponse
}

func (s *CoreIpSecConnectionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreIpSecConnectionsDataSourceCrud) Get() error {
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

func (s *CoreIpSecConnectionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("CoreIpSecConnectionsDataSource-", CoreIpSecConnectionsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		ipSecConnection := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.CpeId != nil {
			ipSecConnection["cpe_id"] = *r.CpeId
		}

		if r.CpeLocalIdentifier != nil {
			ipSecConnection["cpe_local_identifier"] = *r.CpeLocalIdentifier
		}

		ipSecConnection["cpe_local_identifier_type"] = r.CpeLocalIdentifierType

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
		resources = ApplyFilters(f.(*schema.Set), resources, CoreIpSecConnectionsDataSource().Schema["connections"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("connections", resources); err != nil {
		return err
	}

	return nil
}
