// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity_domains

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_identity_domains "github.com/oracle/oci-go-sdk/v65/identitydomains"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func IdentityDomainsIdentityPropagationTrustDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["attribute_sets"] = &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	}
	fieldMap["attributes"] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	}
	fieldMap["authorization"] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	}
	fieldMap["idcs_endpoint"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["identity_propagation_trust_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["resource_type_schema_version"] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(IdentityDomainsIdentityPropagationTrustResource(), fieldMap, readSingularIdentityDomainsIdentityPropagationTrust)
}

func readSingularIdentityDomainsIdentityPropagationTrust(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsIdentityPropagationTrustDataSourceCrud{}
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

type IdentityDomainsIdentityPropagationTrustDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity_domains.IdentityDomainsClient
	Res    *oci_identity_domains.GetIdentityPropagationTrustResponse
}

func (s *IdentityDomainsIdentityPropagationTrustDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityDomainsIdentityPropagationTrustDataSourceCrud) Get() error {
	request := oci_identity_domains.GetIdentityPropagationTrustRequest{}

	if attributeSets, ok := s.D.GetOkExists("attribute_sets"); ok {
		interfaces := attributeSets.([]interface{})
		tmp := make([]oci_identity_domains.AttributeSetsEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_identity_domains.AttributeSetsEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("attribute_sets") {
			request.AttributeSets = tmp
		}
	}

	if attributes, ok := s.D.GetOkExists("attributes"); ok {
		tmp := attributes.(string)
		request.Attributes = &tmp
	}

	if authorization, ok := s.D.GetOkExists("authorization"); ok {
		tmp := authorization.(string)
		request.Authorization = &tmp
	}

	if identityPropagationTrustId, ok := s.D.GetOkExists("identity_propagation_trust_id"); ok {
		tmp := identityPropagationTrustId.(string)
		request.IdentityPropagationTrustId = &tmp
	}

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "identity_domains")

	response, err := s.Client.GetIdentityPropagationTrust(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *IdentityDomainsIdentityPropagationTrustDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AccountId != nil {
		s.D.Set("account_id", *s.Res.AccountId)
	}

	if s.Res.Active != nil {
		s.D.Set("active", *s.Res.Active)
	}

	if s.Res.AllowImpersonation != nil {
		s.D.Set("allow_impersonation", *s.Res.AllowImpersonation)
	}

	if s.Res.ClientClaimName != nil {
		s.D.Set("client_claim_name", *s.Res.ClientClaimName)
	}

	s.D.Set("client_claim_values", s.Res.ClientClaimValues)

	if s.Res.ClockSkewSeconds != nil {
		s.D.Set("clock_skew_seconds", *s.Res.ClockSkewSeconds)
	}

	if s.Res.CompartmentOcid != nil {
		s.D.Set("compartment_ocid", *s.Res.CompartmentOcid)
	}

	if s.Res.DeleteInProgress != nil {
		s.D.Set("delete_in_progress", *s.Res.DeleteInProgress)
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DomainOcid != nil {
		s.D.Set("domain_ocid", *s.Res.DomainOcid)
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

	impersonationServiceUsers := []interface{}{}
	for _, item := range s.Res.ImpersonationServiceUsers {
		impersonationServiceUsers = append(impersonationServiceUsers, IdentityPropagationTrustImpersonationServiceUsersToMap(item))
	}
	s.D.Set("impersonation_service_users", impersonationServiceUsers)

	if s.Res.Issuer != nil {
		s.D.Set("issuer", *s.Res.Issuer)
	}

	if s.Res.Keytab != nil {
		s.D.Set("keytab", []interface{}{IdentityPropagationTrustKeytabToMap(s.Res.Keytab)})
	} else {
		s.D.Set("keytab", nil)
	}

	if s.Res.Meta != nil {
		s.D.Set("meta", []interface{}{metaToMap(s.Res.Meta)})
	} else {
		s.D.Set("meta", nil)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("oauth_clients", s.Res.OauthClients)

	if s.Res.Ocid != nil {
		s.D.Set("ocid", *s.Res.Ocid)
	}

	if s.Res.PublicCertificate != nil {
		s.D.Set("public_certificate", *s.Res.PublicCertificate)
	}

	if s.Res.PublicKeyEndpoint != nil {
		s.D.Set("public_key_endpoint", *s.Res.PublicKeyEndpoint)
	}

	s.D.Set("schemas", s.Res.Schemas)

	if s.Res.SubjectClaimName != nil {
		s.D.Set("subject_claim_name", *s.Res.SubjectClaimName)
	}

	if s.Res.SubjectMappingAttribute != nil {
		s.D.Set("subject_mapping_attribute", *s.Res.SubjectMappingAttribute)
	}

	s.D.Set("subject_type", s.Res.SubjectType)

	tags := []interface{}{}
	for _, item := range s.Res.Tags {
		tags = append(tags, tagsToMap(item))
	}
	s.D.Set("tags", tags)

	if s.Res.TenancyOcid != nil {
		s.D.Set("tenancy_ocid", *s.Res.TenancyOcid)
	}

	s.D.Set("type", s.Res.Type)

	return nil
}
