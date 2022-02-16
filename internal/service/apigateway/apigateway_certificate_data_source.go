// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package apigateway

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_apigateway "github.com/oracle/oci-go-sdk/v58/apigateway"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func ApigatewayCertificateDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["certificate_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(ApigatewayCertificateResource(), fieldMap, readSingularApigatewayCertificate)
}

func readSingularApigatewayCertificate(d *schema.ResourceData, m interface{}) error {
	sync := &ApigatewayCertificateDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApiGatewayClient()

	return tfresource.ReadResource(sync)
}

type ApigatewayCertificateDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_apigateway.ApiGatewayClient
	Res    *oci_apigateway.GetCertificateResponse
}

func (s *ApigatewayCertificateDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ApigatewayCertificateDataSourceCrud) Get() error {
	request := oci_apigateway.GetCertificateRequest{}

	if certificateId, ok := s.D.GetOkExists("certificate_id"); ok {
		tmp := certificateId.(string)
		request.CertificateId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "apigateway")

	response, err := s.Client.GetCertificate(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ApigatewayCertificateDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.Certificate.Certificate != nil {
		s.D.Set("certificate", s.Res.Certificate)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IntermediateCertificates != nil {
		s.D.Set("intermediate_certificates", *s.Res.IntermediateCertificates)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("subject_names", s.Res.SubjectNames)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeNotValidAfter != nil {
		s.D.Set("time_not_valid_after", s.Res.TimeNotValidAfter.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
