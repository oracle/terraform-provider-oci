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

func IdentityDomainsIdentitySettingDataSource() *schema.Resource {
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
	fieldMap["identity_setting_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["resource_type_schema_version"] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(IdentityDomainsIdentitySettingResource(), fieldMap, readSingularIdentityDomainsIdentitySetting)
}

func readSingularIdentityDomainsIdentitySetting(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsIdentitySettingDataSourceCrud{}
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

type IdentityDomainsIdentitySettingDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity_domains.IdentityDomainsClient
	Res    *oci_identity_domains.GetIdentitySettingResponse
}

func (s *IdentityDomainsIdentitySettingDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityDomainsIdentitySettingDataSourceCrud) Get() error {
	request := oci_identity_domains.GetIdentitySettingRequest{}

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

	if identitySettingId, ok := s.D.GetOkExists("identity_setting_id"); ok {
		tmp := identitySettingId.(string)
		request.IdentitySettingId = &tmp
	}

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "identity_domains")

	response, err := s.Client.GetIdentitySetting(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *IdentityDomainsIdentitySettingDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.POSIXGid != nil {
		s.D.Set("posix_gid", []interface{}{IdentitySettingsPOSIXGidToMap(s.Res.POSIXGid)})
	} else {
		s.D.Set("posix_gid", nil)
	}

	if s.Res.POSIXUid != nil {
		s.D.Set("posix_uid", []interface{}{IdentitySettingsPOSIXUidToMap(s.Res.POSIXUid)})
	} else {
		s.D.Set("posix_uid", nil)
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

	if s.Res.EmitLockedMessageWhenUserIsLocked != nil {
		s.D.Set("emit_locked_message_when_user_is_locked", *s.Res.EmitLockedMessageWhenUserIsLocked)
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

	if s.Res.Meta != nil {
		s.D.Set("meta", []interface{}{metaToMap(s.Res.Meta)})
	} else {
		s.D.Set("meta", nil)
	}

	if s.Res.MyProfile != nil {
		s.D.Set("my_profile", []interface{}{IdentitySettingsMyProfileToMap(s.Res.MyProfile)})
	} else {
		s.D.Set("my_profile", nil)
	}

	if s.Res.Ocid != nil {
		s.D.Set("ocid", *s.Res.Ocid)
	}

	if s.Res.PrimaryEmailRequired != nil {
		s.D.Set("primary_email_required", *s.Res.PrimaryEmailRequired)
	}

	if s.Res.RemoveInvalidEmails != nil {
		s.D.Set("remove_invalid_emails", *s.Res.RemoveInvalidEmails)
	}

	if s.Res.ReturnInactiveOverLockedMessage != nil {
		s.D.Set("return_inactive_over_locked_message", *s.Res.ReturnInactiveOverLockedMessage)
	}

	s.D.Set("schemas", s.Res.Schemas)

	tags := []interface{}{}
	for _, item := range s.Res.Tags {
		tags = append(tags, tagsToMap(item))
	}
	s.D.Set("tags", tags)

	if s.Res.TenancyOcid != nil {
		s.D.Set("tenancy_ocid", *s.Res.TenancyOcid)
	}

	tokens := []interface{}{}
	for _, item := range s.Res.Tokens {
		tokens = append(tokens, IdentitySettingsTokensToMap(item))
	}
	s.D.Set("tokens", tokens)

	if s.Res.UserAllowedToSetRecoveryEmail != nil {
		s.D.Set("user_allowed_to_set_recovery_email", *s.Res.UserAllowedToSetRecoveryEmail)
	}

	return nil
}
