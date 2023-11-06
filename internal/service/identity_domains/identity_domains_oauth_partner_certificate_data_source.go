// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity_domains

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_identity_domains "github.com/oracle/oci-go-sdk/v65/identitydomains"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func IdentityDomainsOAuthPartnerCertificateDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["authorization"] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	}
	fieldMap["idcs_endpoint"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["o_auth_partner_certificate_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["resource_type_schema_version"] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(IdentityDomainsOAuthPartnerCertificateResource(), fieldMap, readSingularIdentityDomainsOAuthPartnerCertificate)
}

func readSingularIdentityDomainsOAuthPartnerCertificate(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsOAuthPartnerCertificateDataSourceCrud{}
	sync.D = d
	idcsEndpoint, err := getIdcsEndpoint(d)
	if err != nil {
		return err
	}
	client, err := m.(*client.OracleClients).IdentityDomainsClientWithEndpoint(idcsEndpoint)
	if err != nil {
		return err
	}
	sync.Client = client

	return tfresource.ReadResource(sync)
}

type IdentityDomainsOAuthPartnerCertificateDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity_domains.IdentityDomainsClient
	Res    *oci_identity_domains.GetOAuthPartnerCertificateResponse
}

func (s *IdentityDomainsOAuthPartnerCertificateDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityDomainsOAuthPartnerCertificateDataSourceCrud) Get() error {
	request := oci_identity_domains.GetOAuthPartnerCertificateRequest{}

	if authorization, ok := s.D.GetOkExists("authorization"); ok {
		tmp := authorization.(string)
		request.Authorization = &tmp
	}

	if oAuthPartnerCertificateId, ok := s.D.GetOkExists("o_auth_partner_certificate_id"); ok {
		tmp := oAuthPartnerCertificateId.(string)
		request.OAuthPartnerCertificateId = &tmp
	}

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "identity_domains")

	response, err := s.Client.GetOAuthPartnerCertificate(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *IdentityDomainsOAuthPartnerCertificateDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CertEndDate != nil {
		s.D.Set("cert_end_date", *s.Res.CertEndDate)
	}

	if s.Res.CertStartDate != nil {
		s.D.Set("cert_start_date", *s.Res.CertStartDate)
	}

	if s.Res.CertificateAlias != nil {
		s.D.Set("certificate_alias", *s.Res.CertificateAlias)
	}

	if s.Res.CompartmentOcid != nil {
		s.D.Set("compartment_ocid", *s.Res.CompartmentOcid)
	}

	if s.Res.DeleteInProgress != nil {
		s.D.Set("delete_in_progress", *s.Res.DeleteInProgress)
	}

	if s.Res.DomainOcid != nil {
		s.D.Set("domain_ocid", *s.Res.DomainOcid)
	}

	if s.Res.ExternalId != nil {
		s.D.Set("external_id", *s.Res.ExternalId)
	}

	if s.Res.IdcsCreatedBy != nil {
		s.D.Set("idcs_created_by", []interface{}{idcsCreatedByToMap(s.Res.IdcsCreatedBy)})
	} else {
		s.D.Set("idcs_created_by", nil)
	}

	if s.Res.IdcsLastModifiedBy != nil {
		s.D.Set("idcs_last_modified_by", []interface{}{idcsLastModifiedByToMap(s.Res.IdcsLastModifiedBy)})
	} else {
		s.D.Set("idcs_last_modified_by", nil)
	}

	if s.Res.IdcsLastUpgradedInRelease != nil {
		s.D.Set("idcs_last_upgraded_in_release", *s.Res.IdcsLastUpgradedInRelease)
	}

	s.D.Set("idcs_prevented_operations", s.Res.IdcsPreventedOperations)

	if s.Res.KeyStoreId != nil {
		s.D.Set("key_store_id", *s.Res.KeyStoreId)
	}

	if s.Res.KeyStoreName != nil {
		s.D.Set("key_store_name", *s.Res.KeyStoreName)
	}

	if s.Res.KeyStorePassword != nil {
		s.D.Set("key_store_password", *s.Res.KeyStorePassword)
	}

	if s.Res.Map != nil {
		s.D.Set("map", *s.Res.Map)
	}

	if s.Res.Meta != nil {
		s.D.Set("meta", []interface{}{metaToMap(s.Res.Meta)})
	} else {
		s.D.Set("meta", nil)
	}

	if s.Res.Ocid != nil {
		s.D.Set("ocid", *s.Res.Ocid)
	}

	s.D.Set("schemas", s.Res.Schemas)

	if s.Res.Sha1Thumbprint != nil {
		s.D.Set("sha1thumbprint", *s.Res.Sha1Thumbprint)
	}

	if s.Res.Sha256Thumbprint != nil {
		s.D.Set("sha256thumbprint", *s.Res.Sha256Thumbprint)
	}

	tags := []interface{}{}
	for _, item := range s.Res.Tags {
		tags = append(tags, tagsToMap(item))
	}
	s.D.Set("tags", tags)

	if s.Res.TenancyOcid != nil {
		s.D.Set("tenancy_ocid", *s.Res.TenancyOcid)
	}

	if s.Res.X509Base64Certificate != nil {
		s.D.Set("x509base64certificate", *s.Res.X509Base64Certificate)
	}

	return nil
}
