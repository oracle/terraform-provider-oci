// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oda

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_oda "github.com/oracle/oci-go-sdk/v65/oda"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OdaOdaPrivateEndpointAttachmentsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOdaOdaPrivateEndpointAttachments,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"oda_private_endpoint_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"oda_private_endpoint_attachment_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(OdaOdaPrivateEndpointAttachmentResource()),
						},
					},
				},
			},
		},
	}
}

func readOdaOdaPrivateEndpointAttachments(d *schema.ResourceData, m interface{}) error {
	sync := &OdaOdaPrivateEndpointAttachmentsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagementClient()

	return tfresource.ReadResource(sync)
}

type OdaOdaPrivateEndpointAttachmentsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_oda.ManagementClient
	Res    *oci_oda.ListOdaPrivateEndpointAttachmentsResponse
}

func (s *OdaOdaPrivateEndpointAttachmentsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OdaOdaPrivateEndpointAttachmentsDataSourceCrud) Get() error {
	request := oci_oda.ListOdaPrivateEndpointAttachmentsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if odaPrivateEndpointId, ok := s.D.GetOkExists("oda_private_endpoint_id"); ok {
		tmp := odaPrivateEndpointId.(string)
		request.OdaPrivateEndpointId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_oda.OdaPrivateEndpointAttachmentLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "oda")

	response, err := s.Client.ListOdaPrivateEndpointAttachments(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListOdaPrivateEndpointAttachments(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OdaOdaPrivateEndpointAttachmentsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OdaOdaPrivateEndpointAttachmentsDataSource-", OdaOdaPrivateEndpointAttachmentsDataSource(), s.D))
	resources := []map[string]interface{}{}
	odaPrivateEndpointAttachment := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, OdaPrivateEndpointAttachmentSummaryToMap(item))
	}
	odaPrivateEndpointAttachment["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, OdaOdaPrivateEndpointAttachmentsDataSource().Schema["oda_private_endpoint_attachment_collection"].Elem.(*schema.Resource).Schema)
		odaPrivateEndpointAttachment["items"] = items
	}

	resources = append(resources, odaPrivateEndpointAttachment)
	if err := s.D.Set("oda_private_endpoint_attachment_collection", resources); err != nil {
		return err
	}

	return nil
}
