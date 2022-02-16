// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package certificates_management

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_certificates_management "github.com/oracle/oci-go-sdk/v58/certificatesmanagement"
)

func CertificatesManagementCertificateAuthorityDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["certificate_authority_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(CertificatesManagementCertificateAuthorityResource(), fieldMap, readSingularCertificatesManagementCertificateAuthority)
}

func readSingularCertificatesManagementCertificateAuthority(d *schema.ResourceData, m interface{}) error {
	sync := &CertificatesManagementCertificateAuthorityDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CertificatesManagementClient()

	return tfresource.ReadResource(sync)
}

type CertificatesManagementCertificateAuthorityDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_certificates_management.CertificatesManagementClient
	Res    *oci_certificates_management.GetCertificateAuthorityResponse
}

func (s *CertificatesManagementCertificateAuthorityDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CertificatesManagementCertificateAuthorityDataSourceCrud) Get() error {
	request := oci_certificates_management.GetCertificateAuthorityRequest{}

	if certificateAuthorityId, ok := s.D.GetOkExists("certificate_authority_id"); ok {
		tmp := certificateAuthorityId.(string)
		request.CertificateAuthorityId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "certificates_management")

	response, err := s.Client.GetCertificateAuthority(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CertificatesManagementCertificateAuthorityDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	certificateAuthorityRules := []interface{}{}
	for _, item := range s.Res.CertificateAuthorityRules {
		certificateAuthorityRules = append(certificateAuthorityRules, CertificateAuthorityRuleToMap(item))
	}
	s.D.Set("certificate_authority_rules", certificateAuthorityRules)

	if s.Res.CertificateRevocationListDetails != nil {
		s.D.Set("certificate_revocation_list_details", []interface{}{CertificateRevocationListDetailsToMap(s.Res.CertificateRevocationListDetails)})
	} else {
		s.D.Set("certificate_revocation_list_details", nil)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("config_type", s.Res.ConfigType)

	if s.Res.CurrentVersion != nil {
		s.D.Set("current_version", []interface{}{CertificateAuthorityVersionSummaryToMap(s.Res.CurrentVersion)})
	} else {
		s.D.Set("current_version", nil)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IssuerCertificateAuthorityId != nil {
		s.D.Set("issuer_certificate_authority_id", *s.Res.IssuerCertificateAuthorityId)
	}

	if s.Res.KmsKeyId != nil {
		s.D.Set("kms_key_id", *s.Res.KmsKeyId)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("signing_algorithm", s.Res.SigningAlgorithm)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.Subject != nil {
		s.D.Set("subject", []interface{}{CertificateSubjectToMap(s.Res.Subject)})
	} else {
		s.D.Set("subject", nil)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeOfDeletion != nil {
		s.D.Set("time_of_deletion", s.Res.TimeOfDeletion.String())
	}

	return nil
}
