// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	oci_core "github.com/oracle/oci-go-sdk/v58/core"
)

func CoreDrgAttachmentsListResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCoreDrgAttachmentsList,
		Read:     readCoreDrgAttachmentsList,
		Delete:   deleteCoreDrgAttachmentsList,
		Schema: map[string]*schema.Schema{
			// Required
			"drg_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"attachment_type": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"is_cross_tenancy": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},

			// Computed
			"drg_all_attachments": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func createCoreDrgAttachmentsList(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDrgAttachmentsListResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.CreateResource(d, sync)
}

func readCoreDrgAttachmentsList(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteCoreDrgAttachmentsList(d *schema.ResourceData, m interface{}) error {
	return nil
}

type CoreDrgAttachmentsListResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_core.VirtualNetworkClient
	Res                    []oci_core.DrgAttachmentInfo
	DisableNotFoundRetries bool
}

func (s *CoreDrgAttachmentsListResourceCrud) ID() string {
	return tfresource.GenerateDataSourceHashID("CoreDrgAttachmentsListResource-", CoreDrgAttachmentsListResource(), s.D)
}

func (s *CoreDrgAttachmentsListResourceCrud) Create() error {
	request := oci_core.GetAllDrgAttachmentsRequest{}

	if attachmentType, ok := s.D.GetOkExists("attachment_type"); ok {
		request.AttachmentType = oci_core.GetAllDrgAttachmentsAttachmentTypeEnum(attachmentType.(string))
	}

	if drgId, ok := s.D.GetOkExists("drg_id"); ok {
		tmp := drgId.(string)
		request.DrgId = &tmp
	}

	if isCrossTenancy, ok := s.D.GetOkExists("is_cross_tenancy"); ok {
		tmp := isCrossTenancy.(bool)
		request.IsCrossTenancy = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetAllDrgAttachments(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = response.Items

	request.Page = response.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.GetAllDrgAttachments(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res = append(s.Res, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreDrgAttachmentsListResourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	resources := []map[string]interface{}{}

	for _, r := range s.Res {
		resources = append(resources, map[string]interface{}{"id": r.Id})
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, CoreDrgRouteDistributionStatementsDataSource().Schema["drg_all_attachments"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("drg_all_attachments", resources); err != nil {
		return err
	}

	return nil
}
