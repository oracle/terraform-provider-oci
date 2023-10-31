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

func IdentityDomainsGroupDataSource() *schema.Resource {
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
	fieldMap["group_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["idcs_endpoint"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["resource_type_schema_version"] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(IdentityDomainsGroupResource(), fieldMap, readSingularIdentityDomainsGroup)
}

func readSingularIdentityDomainsGroup(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsGroupDataSourceCrud{}
	sync.D = d
	idcsEndpoint, err := getIdcsEndpointForRead(d, "groups")
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

type IdentityDomainsGroupDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity_domains.IdentityDomainsClient
	Res    *oci_identity_domains.GetGroupResponse
}

func (s *IdentityDomainsGroupDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityDomainsGroupDataSourceCrud) Get() error {
	request := oci_identity_domains.GetGroupRequest{}

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

	if groupId, ok := s.D.GetOkExists("group_id"); ok {
		tmp := groupId.(string)
		request.GroupId = &tmp
	}

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "identity_domains")

	response, err := s.Client.GetGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *IdentityDomainsGroupDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentOcid != nil {
		s.D.Set("compartment_ocid", *s.Res.CompartmentOcid)
	}

	if s.Res.DeleteInProgress != nil {
		s.D.Set("delete_in_progress", *s.Res.DeleteInProgress)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
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

	members := []interface{}{}
	for _, item := range s.Res.Members {
		members = append(members, GroupMembersToMap(item))
	}
	s.D.Set("members", members)

	if s.Res.Meta != nil {
		s.D.Set("meta", []interface{}{metaToMap(s.Res.Meta)})
	} else {
		s.D.Set("meta", nil)
	}

	if s.Res.NonUniqueDisplayName != nil {
		s.D.Set("non_unique_display_name", *s.Res.NonUniqueDisplayName)
	}

	if s.Res.Ocid != nil {
		s.D.Set("ocid", *s.Res.Ocid)
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

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionOciTags != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextension_oci_tags", []interface{}{ExtensionOCITagsToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionOciTags)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextension_oci_tags", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionDbcsGroup != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensiondbcs_group", []interface{}{ExtensionDbcsGroupToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionDbcsGroup)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensiondbcs_group", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionDynamicGroup != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensiondynamic_group", []interface{}{ExtensionDynamicGroupToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionDynamicGroup)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensiondynamic_group", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionGroupGroup != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensiongroup_group", []interface{}{ExtensionGroupGroupToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionGroupGroup)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensiongroup_group", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionPosixGroup != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionposix_group", []interface{}{ExtensionPosixGroupToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionPosixGroup)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionposix_group", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionRequestableGroup != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionrequestable_group", []interface{}{ExtensionRequestableGroupToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionRequestableGroup)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionrequestable_group", nil)
	}

	return nil
}
