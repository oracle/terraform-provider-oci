// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package waas

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_waas "github.com/oracle/oci-go-sdk/v65/waas"
)

func WaasCertificateDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["certificate_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(WaasCertificateResource(), fieldMap, readSingularWaasCertificate)
}

func readSingularWaasCertificate(d *schema.ResourceData, m interface{}) error {
	sync := &WaasCertificateDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).WaasClient()

	return tfresource.ReadResource(sync)
}

type WaasCertificateDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_waas.WaasClient
	Res    *oci_waas.GetCertificateResponse
}

func (s *WaasCertificateDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *WaasCertificateDataSourceCrud) Get() error {
	request := oci_waas.GetCertificateRequest{}

	if certificateId, ok := s.D.GetOkExists("certificate_id"); ok {
		tmp := certificateId.(string)
		request.CertificateId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "waas")

	response, err := s.Client.GetCertificate(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *WaasCertificateDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CertificateData != nil {
		s.D.Set("certificate_data", *s.Res.CertificateData)
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

	extensions := []interface{}{}
	for _, item := range s.Res.Extensions {
		extensions = append(extensions, CertificateExtensionToMap(item))
	}
	s.D.Set("extensions", extensions)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsTrustVerificationDisabled != nil {
		s.D.Set("is_trust_verification_disabled", *s.Res.IsTrustVerificationDisabled)
	}

	if s.Res.IssuedBy != nil {
		s.D.Set("issued_by", *s.Res.IssuedBy)
	}

	if s.Res.IssuerName != nil {
		s.D.Set("issuer_name", []interface{}{CertificateIssuerNameToMap(s.Res.IssuerName)})
	} else {
		s.D.Set("issuer_name", nil)
	}

	if s.Res.PublicKeyInfo != nil {
		s.D.Set("public_key_info", []interface{}{CertificatePublicKeyInfoToMap(s.Res.PublicKeyInfo)})
	} else {
		s.D.Set("public_key_info", nil)
	}

	if s.Res.SerialNumber != nil {
		s.D.Set("serial_number", *s.Res.SerialNumber)
	}

	if s.Res.SignatureAlgorithm != nil {
		s.D.Set("signature_algorithm", *s.Res.SignatureAlgorithm)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SubjectName != nil {
		s.D.Set("subject_name", []interface{}{CertificateSubjectNameToMap(s.Res.SubjectName)})
	} else {
		s.D.Set("subject_name", nil)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeNotValidAfter != nil {
		s.D.Set("time_not_valid_after", s.Res.TimeNotValidAfter.String())
	}

	if s.Res.TimeNotValidBefore != nil {
		s.D.Set("time_not_valid_before", s.Res.TimeNotValidBefore.String())
	}

	if s.Res.Version != nil {
		s.D.Set("version", *s.Res.Version)
	}

	return nil
}
