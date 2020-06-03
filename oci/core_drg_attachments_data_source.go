// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/core"
)

func init() {
	RegisterDatasource("oci_core_drg_attachments", CoreDrgAttachmentsDataSource())
}

func CoreDrgAttachmentsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreDrgAttachments,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"drg_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcn_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"drg_attachments": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(CoreDrgAttachmentResource()),
			},
		},
	}
}

func readCoreDrgAttachments(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDrgAttachmentsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient()

	return ReadResource(sync)
}

type CoreDrgAttachmentsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.ListDrgAttachmentsResponse
}

func (s *CoreDrgAttachmentsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreDrgAttachmentsDataSourceCrud) Get() error {
	request := oci_core.ListDrgAttachmentsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if drgId, ok := s.D.GetOkExists("drg_id"); ok {
		tmp := drgId.(string)
		request.DrgId = &tmp
	}

	if vcnId, ok := s.D.GetOkExists("vcn_id"); ok {
		tmp := vcnId.(string)
		request.VcnId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

	response, err := s.Client.ListDrgAttachments(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDrgAttachments(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreDrgAttachmentsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		drgAttachment := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DisplayName != nil {
			drgAttachment["display_name"] = *r.DisplayName
		}

		if r.DrgId != nil {
			drgAttachment["drg_id"] = *r.DrgId
		}

		if r.Id != nil {
			drgAttachment["id"] = *r.Id
		}

		if r.RouteTableId != nil {
			drgAttachment["route_table_id"] = *r.RouteTableId
		}

		drgAttachment["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			drgAttachment["time_created"] = r.TimeCreated.String()
		}

		if r.VcnId != nil {
			drgAttachment["vcn_id"] = *r.VcnId
		}

		resources = append(resources, drgAttachment)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, CoreDrgAttachmentsDataSource().Schema["drg_attachments"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("drg_attachments", resources); err != nil {
		return err
	}

	return nil
}
