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

func IdentityDomainsMyAppsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readIdentityDomainsMyApps,
		Schema: map[string]*schema.Schema{
			"idcs_endpoint": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"my_app_count": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"my_app_filter": {
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
			"sort_order": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sort_by": {
				Type:     schema.TypeString,
				Optional: true,
			},
			// Computed
			"my_apps": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

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
									"is_login_target": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"is_opc_service": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"login_mechanism": {
										Type:     schema.TypeString,
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
						"delete_in_progress": {
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
						"id": {
							Type:     schema.TypeString,
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
						"launch_url": {
							Type:     schema.TypeString,
							Computed: true,
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
						"ocid": {
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
				},
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
			"total_results": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func readIdentityDomainsMyApps(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsMyAppsDataSourceCrud{}
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

type IdentityDomainsMyAppsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity_domains.IdentityDomainsClient
	Res    *oci_identity_domains.ListMyAppsResponse
}

func (s *IdentityDomainsMyAppsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityDomainsMyAppsDataSourceCrud) Get() error {
	request := oci_identity_domains.ListMyAppsRequest{}

	if myAppCount, ok := s.D.GetOkExists("my_app_count"); ok {
		tmp := myAppCount.(int)
		request.Count = &tmp
	}

	if myAppFilter, ok := s.D.GetOkExists("my_app_filter"); ok {
		tmp := myAppFilter.(string)
		request.Filter = &tmp
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
		tmp := oci_identity_domains.ListMyAppsSortOrderEnum(sortOrder.(string))
		request.SortOrder = tmp
	}

	if sortBy, ok := s.D.GetOkExists("sort_by"); ok {
		tmp := sortBy.(string)
		request.SortBy = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "identity_domains")

	response, err := s.Client.ListMyApps(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	// IDCS pagination
	startIndex := *response.StartIndex
	for startIndex+*response.ItemsPerPage <= *response.TotalResults {
		startIndex += *response.ItemsPerPage
		request.StartIndex = &startIndex
		listResponse, err := s.Client.ListMyApps(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Resources = append(s.Res.Resources, listResponse.Resources...)
	}

	return nil
}

func (s *IdentityDomainsMyAppsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("IdentityDomainsMyAppsDataSource-", IdentityDomainsMyAppsDataSource(), s.D))

	resources := []interface{}{}
	for _, item := range s.Res.Resources {
		resources = append(resources, MyAppToMap(item))
	}
	s.D.Set("my_apps", resources)

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

func MyAppToMap(obj oci_identity_domains.MyApp) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AccountType != nil {
		result["account_type"] = string(*obj.AccountType)
	}

	if obj.Active != nil {
		result["active"] = bool(*obj.Active)
	}

	if obj.App != nil {
		result["app"] = []interface{}{MyAppAppToMap(obj.App)}
	}

	if obj.CompartmentOcid != nil {
		result["compartment_ocid"] = string(*obj.CompartmentOcid)
	}

	if obj.DeleteInProgress != nil {
		result["delete_in_progress"] = bool(*obj.DeleteInProgress)
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

	if obj.LaunchUrl != nil {
		result["launch_url"] = string(*obj.LaunchUrl)
	}

	if obj.Meta != nil {
		result["meta"] = []interface{}{metaToMap(obj.Meta)}
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Ocid != nil {
		result["ocid"] = string(*obj.Ocid)
	}

	if obj.Owner != nil {
		result["owner"] = []interface{}{MyAppOwnerToMap(obj.Owner)}
	}

	result["schemas"] = obj.Schemas

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
		result["user_wallet_artifact"] = []interface{}{MyAppUserWalletArtifactToMap(obj.UserWalletArtifact)}
	}

	return result
}

func MyAppAppToMap(obj *oci_identity_domains.MyAppApp) map[string]interface{} {
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

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.Display != nil {
		result["display"] = string(*obj.Display)
	}

	if obj.IsAliasApp != nil {
		result["is_alias_app"] = bool(*obj.IsAliasApp)
	}

	if obj.IsLoginTarget != nil {
		result["is_login_target"] = bool(*obj.IsLoginTarget)
	}

	if obj.IsOPCService != nil {
		result["is_opc_service"] = bool(*obj.IsOPCService)
	}

	if obj.LoginMechanism != nil {
		result["login_mechanism"] = string(*obj.LoginMechanism)
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

func MyAppOwnerToMap(obj *oci_identity_domains.MyAppOwner) map[string]interface{} {
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

func MyAppUserWalletArtifactToMap(obj *oci_identity_domains.MyAppUserWalletArtifact) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}
