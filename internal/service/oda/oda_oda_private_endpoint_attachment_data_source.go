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

func OdaOdaPrivateEndpointAttachmentDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["oda_private_endpoint_attachment_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(OdaOdaPrivateEndpointAttachmentResource(), fieldMap, readSingularOdaOdaPrivateEndpointAttachment)
}

func readSingularOdaOdaPrivateEndpointAttachment(d *schema.ResourceData, m interface{}) error {
	sync := &OdaOdaPrivateEndpointAttachmentDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagementClient()

	return tfresource.ReadResource(sync)
}

type OdaOdaPrivateEndpointAttachmentDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_oda.ManagementClient
	Res    *oci_oda.GetOdaPrivateEndpointAttachmentResponse
}

func (s *OdaOdaPrivateEndpointAttachmentDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OdaOdaPrivateEndpointAttachmentDataSourceCrud) Get() error {
	request := oci_oda.GetOdaPrivateEndpointAttachmentRequest{}

	if odaPrivateEndpointAttachmentId, ok := s.D.GetOkExists("oda_private_endpoint_attachment_id"); ok {
		tmp := odaPrivateEndpointAttachmentId.(string)
		request.OdaPrivateEndpointAttachmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "oda")

	response, err := s.Client.GetOdaPrivateEndpointAttachment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OdaOdaPrivateEndpointAttachmentDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.OdaInstanceId != nil {
		s.D.Set("oda_instance_id", *s.Res.OdaInstanceId)
	}

	if s.Res.OdaPrivateEndpointId != nil {
		s.D.Set("oda_private_endpoint_id", *s.Res.OdaPrivateEndpointId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
