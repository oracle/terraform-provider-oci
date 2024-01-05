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

func IdentityDomainsMyUserDbCredentialDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["authorization"] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	}
	fieldMap["idcs_endpoint"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["my_user_db_credential_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["resource_type_schema_version"] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(IdentityDomainsMyUserDbCredentialResource(), fieldMap, readSingularIdentityDomainsMyUserDbCredential)
}

func readSingularIdentityDomainsMyUserDbCredential(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsMyUserDbCredentialDataSourceCrud{}
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

type IdentityDomainsMyUserDbCredentialDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity_domains.IdentityDomainsClient
	Res    *oci_identity_domains.GetMyUserDbCredentialResponse
}

func (s *IdentityDomainsMyUserDbCredentialDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityDomainsMyUserDbCredentialDataSourceCrud) Get() error {
	request := oci_identity_domains.GetMyUserDbCredentialRequest{}

	if authorization, ok := s.D.GetOkExists("authorization"); ok {
		tmp := authorization.(string)
		request.Authorization = &tmp
	}

	if myUserDbCredentialId, ok := s.D.GetOkExists("my_user_db_credential_id"); ok {
		tmp := myUserDbCredentialId.(string)
		request.MyUserDbCredentialId = &tmp
	}

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "identity_domains")

	response, err := s.Client.GetMyUserDbCredential(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *IdentityDomainsMyUserDbCredentialDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentOcid != nil {
		s.D.Set("compartment_ocid", *s.Res.CompartmentOcid)
	}

	if s.Res.DbPassword != nil {
		s.D.Set("db_password", *s.Res.DbPassword)
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

	if s.Res.Expired != nil {
		s.D.Set("expired", *s.Res.Expired)
	}

	if s.Res.ExpiresOn != nil {
		s.D.Set("expires_on", *s.Res.ExpiresOn)
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

	if s.Res.LastSetDate != nil {
		s.D.Set("last_set_date", *s.Res.LastSetDate)
	}

	if s.Res.Meta != nil {
		s.D.Set("meta", []interface{}{metaToMap(s.Res.Meta)})
	} else {
		s.D.Set("meta", nil)
	}

	if s.Res.MixedDbPassword != nil {
		s.D.Set("mixed_db_password", *s.Res.MixedDbPassword)
	}

	if s.Res.MixedSalt != nil {
		s.D.Set("mixed_salt", *s.Res.MixedSalt)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.Ocid != nil {
		s.D.Set("ocid", *s.Res.Ocid)
	}

	if s.Res.Salt != nil {
		s.D.Set("salt", *s.Res.Salt)
	}

	s.D.Set("schemas", s.Res.Schemas)

	s.D.Set("status", s.Res.Status)

	tags := []interface{}{}
	for _, item := range s.Res.Tags {
		tags = append(tags, tagsToMap(item))
	}
	s.D.Set("tags", tags)

	if s.Res.TenancyOcid != nil {
		s.D.Set("tenancy_ocid", *s.Res.TenancyOcid)
	}

	if s.Res.User != nil {
		s.D.Set("user", []interface{}{MyUserDbCredentialsUserToMap(s.Res.User)})
	} else {
		s.D.Set("user", nil)
	}

	return nil
}
