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

func IdentityDomainsAccountMgmtInfosDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readIdentityDomainsAccountMgmtInfos,
		Schema: map[string]*schema.Schema{
			"idcs_endpoint": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"account_mgmt_info_count": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"account_mgmt_info_filter": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"attribute_sets": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"attributes": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"authorization": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_type_schema_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"start_index": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			// Computed
			"account_mgmt_infos": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(IdentityDomainsAccountMgmtInfoDataSource()),
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"items_per_page": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"schemas": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"sort_order": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sort_by": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"total_results": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func readIdentityDomainsAccountMgmtInfos(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsAccountMgmtInfosDataSourceCrud{}
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

type IdentityDomainsAccountMgmtInfosDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity_domains.IdentityDomainsClient
	Res    *oci_identity_domains.ListAccountMgmtInfosResponse
}

func (s *IdentityDomainsAccountMgmtInfosDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityDomainsAccountMgmtInfosDataSourceCrud) Get() error {
	request := oci_identity_domains.ListAccountMgmtInfosRequest{}

	if accountMgmtInfoCount, ok := s.D.GetOkExists("account_mgmt_info_count"); ok {
		tmp := accountMgmtInfoCount.(int)
		request.Count = &tmp
	}

	if accountMgmtInfoFilter, ok := s.D.GetOkExists("account_mgmt_info_filter"); ok {
		tmp := accountMgmtInfoFilter.(string)
		request.Filter = &tmp
	}

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

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	if startIndex, ok := s.D.GetOkExists("start_index"); ok {
		tmp := startIndex.(int)
		request.StartIndex = &tmp
	}

	if sortOrder, ok := s.D.GetOkExists("sort_order"); ok {
		tmp := oci_identity_domains.ListAccountMgmtInfosSortOrderEnum(sortOrder.(string))
		request.SortOrder = tmp
	}

	if sortBy, ok := s.D.GetOkExists("sort_by"); ok {
		tmp := sortBy.(string)
		request.SortBy = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "identity_domains")

	response, err := s.Client.ListAccountMgmtInfos(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response

	// IDCS pagination
	startIndex := *response.StartIndex
	for startIndex+*response.ItemsPerPage <= *response.TotalResults {
		startIndex += *response.ItemsPerPage
		request.StartIndex = &startIndex
		listResponse, err := s.Client.ListAccountMgmtInfos(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Resources = append(s.Res.Resources, listResponse.Resources...)
	}

	return nil
}

func (s *IdentityDomainsAccountMgmtInfosDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("IdentityDomainsAccountMgmtInfosDataSource-", IdentityDomainsAccountMgmtInfosDataSource(), s.D))

	resources := []interface{}{}
	for _, item := range s.Res.Resources {
		resources = append(resources, AccountMgmtInfoToMap(item))
	}
	s.D.Set("account_mgmt_infos", resources)

	if s.Res.ItemsPerPage != nil {
		s.D.Set("items_per_page", *s.Res.ItemsPerPage)
	}

	s.D.Set("schemas", s.Res.Schemas)

	if s.Res.StartIndex != nil {
		s.D.Set("start_index", *s.Res.StartIndex)
	}

	if s.Res.TotalResults != nil {
		s.D.Set("total_results", *s.Res.TotalResults)
	}

	return nil
}

func AccountMgmtInfoToMap(obj oci_identity_domains.AccountMgmtInfo) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AccountType != nil {
		result["account_type"] = string(*obj.AccountType)
	}

	if obj.Active != nil {
		result["active"] = bool(*obj.Active)
	}

	if obj.App != nil {
		result["app"] = []interface{}{AccountMgmtInfoAppToMap(obj.App)}
	}

	if obj.CompartmentOcid != nil {
		result["compartment_ocid"] = string(*obj.CompartmentOcid)
	}

	if obj.CompositeKey != nil {
		result["composite_key"] = string(*obj.CompositeKey)
	}

	if obj.DeleteInProgress != nil {
		result["delete_in_progress"] = bool(*obj.DeleteInProgress)
	}

	if obj.DoNotBackFillGrants != nil {
		result["do_not_back_fill_grants"] = bool(*obj.DoNotBackFillGrants)
	}

	if obj.DoNotPerformActionOnTarget != nil {
		result["do_not_perform_action_on_target"] = bool(*obj.DoNotPerformActionOnTarget)
	}

	if obj.DomainOcid != nil {
		result["domain_ocid"] = string(*obj.DomainOcid)
	}

	if obj.Favorite != nil {
		result["favorite"] = bool(*obj.Favorite)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IdcsCreatedBy != nil {
		result["idcs_created_by"] = []interface{}{idcsCreatedByToMap(obj.IdcsCreatedBy)}
	}

	if obj.IdcsLastModifiedBy != nil {
		result["idcs_last_modified_by"] = []interface{}{idcsLastModifiedByToMap(obj.IdcsLastModifiedBy)}
	}

	if obj.IdcsLastUpgradedInRelease != nil {
		result["idcs_last_upgraded_in_release"] = string(*obj.IdcsLastUpgradedInRelease)
	}

	result["idcs_prevented_operations"] = obj.IdcsPreventedOperations

	if obj.IsAccount != nil {
		result["is_account"] = bool(*obj.IsAccount)
	}

	if obj.LastAccessed != nil {
		result["last_accessed"] = string(*obj.LastAccessed)
	}

	matchingOwners := []interface{}{}
	for _, item := range obj.MatchingOwners {
		matchingOwners = append(matchingOwners, AccountMgmtInfoMatchingOwnersToMap(item))
	}
	result["matching_owners"] = matchingOwners

	if obj.Meta != nil {
		result["meta"] = []interface{}{metaToMap(obj.Meta)}
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.ObjectClass != nil {
		result["object_class"] = []interface{}{AccountMgmtInfoObjectClassToMap(obj.ObjectClass)}
	}

	if obj.Ocid != nil {
		result["ocid"] = string(*obj.Ocid)
	}

	result["operation_context"] = string(obj.OperationContext)

	if obj.Owner != nil {
		result["owner"] = []interface{}{AccountMgmtInfoOwnerToMap(obj.Owner)}
	}

	if obj.PreviewOnly != nil {
		result["preview_only"] = bool(*obj.PreviewOnly)
	}

	if obj.ResourceType != nil {
		result["resource_type"] = []interface{}{AccountMgmtInfoResourceTypeToMap(obj.ResourceType)}
	}

	result["schemas"] = obj.Schemas

	if obj.SyncResponse != nil {
		result["sync_response"] = string(*obj.SyncResponse)
	}

	result["sync_situation"] = string(obj.SyncSituation)

	if obj.SyncTimestamp != nil {
		result["sync_timestamp"] = string(*obj.SyncTimestamp)
	}

	tags := []interface{}{}
	for _, item := range obj.Tags {
		tags = append(tags, tagsToMap(item))
	}
	result["tags"] = tags

	if obj.TenancyOcid != nil {
		result["tenancy_ocid"] = string(*obj.TenancyOcid)
	}

	if obj.Uid != nil {
		result["uid"] = string(*obj.Uid)
	}

	if obj.UserWalletArtifact != nil {
		result["user_wallet_artifact"] = []interface{}{AccountMgmtInfoUserWalletArtifactToMap(obj.UserWalletArtifact)}
	}

	return result
}

func AccountMgmtInfoAppToMap(obj *oci_identity_domains.AccountMgmtInfoApp) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Active != nil {
		result["active"] = bool(*obj.Active)
	}

	if obj.AppIcon != nil {
		result["app_icon"] = string(*obj.AppIcon)
	}

	if obj.AppThumbnail != nil {
		result["app_thumbnail"] = string(*obj.AppThumbnail)
	}

	if obj.Audience != nil {
		result["audience"] = string(*obj.Audience)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.Display != nil {
		result["display"] = string(*obj.Display)
	}

	if obj.IsAliasApp != nil {
		result["is_alias_app"] = bool(*obj.IsAliasApp)
	}

	if obj.IsAuthoritative != nil {
		result["is_authoritative"] = bool(*obj.IsAuthoritative)
	}

	if obj.IsLoginTarget != nil {
		result["is_login_target"] = bool(*obj.IsLoginTarget)
	}

	if obj.IsManagedApp != nil {
		result["is_managed_app"] = bool(*obj.IsManagedApp)
	}

	if obj.IsOAuthResource != nil {
		result["is_oauth_resource"] = bool(*obj.IsOAuthResource)
	}

	if obj.IsOPCService != nil {
		result["is_opc_service"] = bool(*obj.IsOPCService)
	}

	if obj.IsUnmanagedApp != nil {
		result["is_unmanaged_app"] = bool(*obj.IsUnmanagedApp)
	}

	if obj.LoginMechanism != nil {
		result["login_mechanism"] = string(*obj.LoginMechanism)
	}

	if obj.MeterAsOPCService != nil {
		result["meter_as_opc_service"] = bool(*obj.MeterAsOPCService)
	}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.ServiceTypeURN != nil {
		result["service_type_urn"] = string(*obj.ServiceTypeURN)
	}

	if obj.ShowInMyApps != nil {
		result["show_in_my_apps"] = bool(*obj.ShowInMyApps)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func AccountMgmtInfoMatchingOwnersToMap(obj oci_identity_domains.AccountMgmtInfoMatchingOwners) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Display != nil {
		result["display"] = string(*obj.Display)
	}

	if obj.Email != nil {
		result["email"] = string(*obj.Email)
	}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.UserName != nil {
		result["user_name"] = string(*obj.UserName)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func AccountMgmtInfoObjectClassToMap(obj *oci_identity_domains.AccountMgmtInfoObjectClass) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Display != nil {
		result["display"] = string(*obj.Display)
	}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func AccountMgmtInfoOwnerToMap(obj *oci_identity_domains.AccountMgmtInfoOwner) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Display != nil {
		result["display"] = string(*obj.Display)
	}

	if obj.Email != nil {
		result["email"] = string(*obj.Email)
	}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.UserName != nil {
		result["user_name"] = string(*obj.UserName)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func AccountMgmtInfoResourceTypeToMap(obj *oci_identity_domains.AccountMgmtInfoResourceType) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Display != nil {
		result["display"] = string(*obj.Display)
	}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func AccountMgmtInfoUserWalletArtifactToMap(obj *oci_identity_domains.AccountMgmtInfoUserWalletArtifact) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}
