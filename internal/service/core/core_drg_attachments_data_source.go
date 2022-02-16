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

func CoreDrgAttachmentsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreDrgAttachments,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"attachment_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"drg_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"drg_route_table_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_id": {
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
			"drg_attachments": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(CoreDrgAttachmentResource()),
			},
		},
	}
}

func readCoreDrgAttachments(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDrgAttachmentsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
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

	if attachmentType, ok := s.D.GetOkExists("attachment_type"); ok {
		request.AttachmentType = oci_core.ListDrgAttachmentsAttachmentTypeEnum(attachmentType.(string))
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if drgId, ok := s.D.GetOkExists("drg_id"); ok {
		tmp := drgId.(string)
		request.DrgId = &tmp
	}

	if drgRouteTableId, ok := s.D.GetOkExists("drg_route_table_id"); ok {
		tmp := drgRouteTableId.(string)
		request.DrgRouteTableId = &tmp
	}

	if networkId, ok := s.D.GetOkExists("network_id"); ok {
		tmp := networkId.(string)
		request.NetworkId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_core.DrgAttachmentLifecycleStateEnum(state.(string))
	}

	if vcnId, ok := s.D.GetOkExists("vcn_id"); ok {
		tmp := vcnId.(string)
		request.VcnId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

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

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreDrgAttachmentsDataSource-", CoreDrgAttachmentsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		drgAttachment := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			drgAttachment["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			drgAttachment["display_name"] = *r.DisplayName
		}

		if r.DrgId != nil {
			drgAttachment["drg_id"] = *r.DrgId
		}

		if r.DrgRouteTableId != nil {
			drgAttachment["drg_route_table_id"] = *r.DrgRouteTableId
		}

		if r.ExportDrgRouteDistributionId != nil {
			drgAttachment["export_drg_route_distribution_id"] = *r.ExportDrgRouteDistributionId
		}

		drgAttachment["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			drgAttachment["id"] = *r.Id
		}

		if r.IsCrossTenancy != nil {
			drgAttachment["is_cross_tenancy"] = *r.IsCrossTenancy
		}

		if r.NetworkDetails != nil {
			networkDetailsArray := []interface{}{}
			if networkDetailsMap := DrgAttachmentNetworkDetailsToMap(&r.NetworkDetails); networkDetailsMap != nil {
				networkDetailsArray = append(networkDetailsArray, networkDetailsMap)
			}
			drgAttachment["network_details"] = networkDetailsArray
		} else {
			drgAttachment["network_details"] = nil
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
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, CoreDrgAttachmentsDataSource().Schema["drg_attachments"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("drg_attachments", resources); err != nil {
		return err
	}

	return nil
}
