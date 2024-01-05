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

func CoreCpesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreCpes,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cpes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(CoreCpeResource()),
			},
		},
	}
}

func readCoreCpes(d *schema.ResourceData, m interface{}) error {
	sync := &CoreCpesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

type CoreCpesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.ListCpesResponse
}

func (s *CoreCpesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreCpesDataSourceCrud) Get() error {
	request := oci_core.ListCpesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.ListCpes(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListCpes(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreCpesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreCpesDataSource-", CoreCpesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		cpe := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.CpeDeviceShapeId != nil {
			cpe["cpe_device_shape_id"] = *r.CpeDeviceShapeId
		}

		if r.DefinedTags != nil {
			cpe["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			cpe["display_name"] = *r.DisplayName
		}

		cpe["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			cpe["id"] = *r.Id
		}

		if r.IpAddress != nil {
			cpe["ip_address"] = *r.IpAddress
		}

		if r.IsPrivate != nil {
			cpe["is_private"] = *r.IsPrivate
		}

		if r.TimeCreated != nil {
			cpe["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, cpe)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, CoreCpesDataSource().Schema["cpes"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("cpes", resources); err != nil {
		return err
	}

	return nil
}
