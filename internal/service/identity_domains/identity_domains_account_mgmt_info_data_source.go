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

func IdentityDomainsAccountMgmtInfoDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularIdentityDomainsAccountMgmtInfo,
		Schema: map[string]*schema.Schema{
			"account_mgmt_info_id": {
				Type:     schema.TypeString,
				Required: true,
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
			"idcs_endpoint": {
				Type:     schema.TypeString,
				Required: true,
			},
			"resource_type_schema_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			// Computed
			"account_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"active": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"app": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"active": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"app_icon": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"app_thumbnail": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"audience": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"display": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_alias_app": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_authoritative": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_login_target": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_managed_app": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_oauth_resource": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_opc_service": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_unmanaged_app": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"login_mechanism": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"meter_as_opc_service": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"ref": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"service_type_urn": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"show_in_my_apps": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"compartment_ocid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"composite_key": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"delete_in_progress": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"do_not_back_fill_grants": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"do_not_perform_action_on_target": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"domain_ocid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"favorite": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"idcs_created_by": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"display": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ocid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ref": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"idcs_last_modified_by": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"display": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ocid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ref": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"idcs_last_upgraded_in_release": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"idcs_prevented_operations": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"is_account": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"last_accessed": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"matching_owners": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"display": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"email": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ref": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"user_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"meta": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"created": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"last_modified": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"location": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"resource_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"version": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"object_class": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"display": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ref": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"ocid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"operation_context": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"owner": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"display": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"email": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ref": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"user_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"preview_only": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"resource_type": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"display": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ref": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"schemas": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"sync_response": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sync_situation": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sync_timestamp": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tags": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"key": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"tenancy_ocid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"uid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"user_wallet_artifact": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"ref": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readSingularIdentityDomainsAccountMgmtInfo(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsAccountMgmtInfoDataSourceCrud{}
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

type IdentityDomainsAccountMgmtInfoDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity_domains.IdentityDomainsClient
	Res    *oci_identity_domains.GetAccountMgmtInfoResponse
}

func (s *IdentityDomainsAccountMgmtInfoDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityDomainsAccountMgmtInfoDataSourceCrud) Get() error {
	request := oci_identity_domains.GetAccountMgmtInfoRequest{}

	if accountMgmtInfoId, ok := s.D.GetOkExists("account_mgmt_info_id"); ok {
		tmp := accountMgmtInfoId.(string)
		request.AccountMgmtInfoId = &tmp
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "identity_domains")

	response, err := s.Client.GetAccountMgmtInfo(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *IdentityDomainsAccountMgmtInfoDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AccountType != nil {
		s.D.Set("account_type", *s.Res.AccountType)
	}

	if s.Res.Active != nil {
		s.D.Set("active", *s.Res.Active)
	}

	if s.Res.App != nil {
		s.D.Set("app", []interface{}{AccountMgmtInfoAppToMap(s.Res.App)})
	} else {
		s.D.Set("app", nil)
	}

	if s.Res.CompartmentOcid != nil {
		s.D.Set("compartment_ocid", *s.Res.CompartmentOcid)
	}

	if s.Res.CompositeKey != nil {
		s.D.Set("composite_key", *s.Res.CompositeKey)
	}

	if s.Res.DeleteInProgress != nil {
		s.D.Set("delete_in_progress", *s.Res.DeleteInProgress)
	}

	if s.Res.DoNotBackFillGrants != nil {
		s.D.Set("do_not_back_fill_grants", *s.Res.DoNotBackFillGrants)
	}

	if s.Res.DoNotPerformActionOnTarget != nil {
		s.D.Set("do_not_perform_action_on_target", *s.Res.DoNotPerformActionOnTarget)
	}

	if s.Res.DomainOcid != nil {
		s.D.Set("domain_ocid", *s.Res.DomainOcid)
	}

	if s.Res.Favorite != nil {
		s.D.Set("favorite", *s.Res.Favorite)
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

	if s.Res.IsAccount != nil {
		s.D.Set("is_account", *s.Res.IsAccount)
	}

	if s.Res.LastAccessed != nil {
		s.D.Set("last_accessed", *s.Res.LastAccessed)
	}

	matchingOwners := []interface{}{}
	for _, item := range s.Res.MatchingOwners {
		matchingOwners = append(matchingOwners, AccountMgmtInfoMatchingOwnersToMap(item))
	}
	s.D.Set("matching_owners", matchingOwners)

	if s.Res.Meta != nil {
		s.D.Set("meta", []interface{}{metaToMap(s.Res.Meta)})
	} else {
		s.D.Set("meta", nil)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.ObjectClass != nil {
		s.D.Set("object_class", []interface{}{AccountMgmtInfoObjectClassToMap(s.Res.ObjectClass)})
	} else {
		s.D.Set("object_class", nil)
	}

	if s.Res.Ocid != nil {
		s.D.Set("ocid", *s.Res.Ocid)
	}

	s.D.Set("operation_context", s.Res.OperationContext)

	if s.Res.Owner != nil {
		s.D.Set("owner", []interface{}{AccountMgmtInfoOwnerToMap(s.Res.Owner)})
	} else {
		s.D.Set("owner", nil)
	}

	if s.Res.PreviewOnly != nil {
		s.D.Set("preview_only", *s.Res.PreviewOnly)
	}

	if s.Res.ResourceType != nil {
		s.D.Set("resource_type", []interface{}{AccountMgmtInfoResourceTypeToMap(s.Res.ResourceType)})
	} else {
		s.D.Set("resource_type", nil)
	}

	s.D.Set("schemas", s.Res.Schemas)

	if s.Res.SyncResponse != nil {
		s.D.Set("sync_response", *s.Res.SyncResponse)
	}

	s.D.Set("sync_situation", s.Res.SyncSituation)

	if s.Res.SyncTimestamp != nil {
		s.D.Set("sync_timestamp", *s.Res.SyncTimestamp)
	}

	tags := []interface{}{}
	for _, item := range s.Res.Tags {
		tags = append(tags, tagsToMap(item))
	}
	s.D.Set("tags", tags)

	if s.Res.TenancyOcid != nil {
		s.D.Set("tenancy_ocid", *s.Res.TenancyOcid)
	}

	if s.Res.Uid != nil {
		s.D.Set("uid", *s.Res.Uid)
	}

	if s.Res.UserWalletArtifact != nil {
		s.D.Set("user_wallet_artifact", []interface{}{AccountMgmtInfoUserWalletArtifactToMap(s.Res.UserWalletArtifact)})
	} else {
		s.D.Set("user_wallet_artifact", nil)
	}

	return nil
}
