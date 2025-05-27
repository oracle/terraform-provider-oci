// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package wlms

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_wlms "github.com/oracle/oci-go-sdk/v65/wlms"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func WlmsWlsDomainsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readWlmsWlsDomains,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"middleware_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"patch_readiness_status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"weblogic_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"wls_domain_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"compartment_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"state": {
										Type:             schema.TypeString,
										Optional:         true,
										Computed:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											string(oci_wlms.WlsDomainLifecycleStateActive),
											string(oci_wlms.WlsDomainLifecycleStateNeedsAttention),
											string(oci_wlms.WlsDomainLifecycleStateDeleting),
											string(oci_wlms.WlsDomainLifecycleStateDeleted),
											string(oci_wlms.WlsDomainLifecycleStateCreating),
											string(oci_wlms.WlsDomainLifecycleStateFailed),
											string(oci_wlms.WlsDomainLifecycleStateUpdating),
										}, true),
									},

									// Computed
									"configuration": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"admin_server_control_mode": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"admin_server_start_script_path": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"admin_server_stop_script_path": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"is_patch_enabled": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"is_rollback_on_failure": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"managed_server_control_mode": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"managed_server_start_script_path": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"managed_server_stop_script_path": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"servers_shutdown_timeout": {
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},
									"defined_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"freeform_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_accepted_terms_and_conditions": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"lifecycle_details": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"middleware_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"patch_readiness_status": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"system_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"time_created": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_updated": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"weblogic_version": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func readWlmsWlsDomains(d *schema.ResourceData, m interface{}) error {
	sync := &WlmsWlsDomainsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).WeblogicManagementServiceClient()

	return tfresource.ReadResource(sync)
}

type WlmsWlsDomainsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_wlms.WeblogicManagementServiceClient
	Res    *oci_wlms.ListWlsDomainsResponse
}

func (s *WlmsWlsDomainsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *WlmsWlsDomainsDataSourceCrud) Get() error {
	request := oci_wlms.ListWlsDomainsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if middlewareType, ok := s.D.GetOkExists("middleware_type"); ok {
		request.MiddlewareType = oci_wlms.ListWlsDomainsMiddlewareTypeEnum(middlewareType.(string))
	}

	if patchReadinessStatus, ok := s.D.GetOkExists("patch_readiness_status"); ok {
		request.PatchReadinessStatus = oci_wlms.ListWlsDomainsPatchReadinessStatusEnum(patchReadinessStatus.(string))
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_wlms.WlsDomainLifecycleStateEnum(state.(string))
	}

	if weblogicVersion, ok := s.D.GetOkExists("weblogic_version"); ok {
		request.WeblogicVersion = oci_wlms.ListWlsDomainsWeblogicVersionEnum(weblogicVersion.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "wlms")

	response, err := s.Client.ListWlsDomains(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListWlsDomains(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *WlmsWlsDomainsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("WlmsWlsDomainsDataSource-", WlmsWlsDomainsDataSource(), s.D))
	resources := []map[string]interface{}{}
	wlsDomain := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, WlsDomainSummaryToMap(item))
	}
	wlsDomain["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, WlmsWlsDomainsDataSource().Schema["wls_domain_collection"].Elem.(*schema.Resource).Schema)
		wlsDomain["items"] = items
	}

	resources = append(resources, wlsDomain)
	if err := s.D.Set("wls_domain_collection", resources); err != nil {
		return err
	}

	return nil
}

func WlsDomainConfigurationToMap(obj *oci_wlms.WlsDomainConfiguration) map[string]interface{} {
	result := map[string]interface{}{}

	result["admin_server_control_mode"] = string(obj.AdminServerControlMode)

	if obj.AdminServerStartScriptPath != nil {
		result["admin_server_start_script_path"] = string(*obj.AdminServerStartScriptPath)
	}

	if obj.AdminServerStopScriptPath != nil {
		result["admin_server_stop_script_path"] = string(*obj.AdminServerStopScriptPath)
	}

	if obj.IsPatchEnabled != nil {
		result["is_patch_enabled"] = bool(*obj.IsPatchEnabled)
	}

	if obj.IsRollbackOnFailure != nil {
		result["is_rollback_on_failure"] = bool(*obj.IsRollbackOnFailure)
	}

	result["managed_server_control_mode"] = string(obj.ManagedServerControlMode)

	if obj.ManagedServerStartScriptPath != nil {
		result["managed_server_start_script_path"] = string(*obj.ManagedServerStartScriptPath)
	}

	if obj.ManagedServerStopScriptPath != nil {
		result["managed_server_stop_script_path"] = string(*obj.ManagedServerStopScriptPath)
	}

	if obj.ServersShutdownTimeout != nil {
		result["servers_shutdown_timeout"] = int(*obj.ServersShutdownTimeout)
	}

	return result
}

func WlsDomainSummaryToMap(obj oci_wlms.WlsDomainSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.MiddlewareType != nil {
		result["middleware_type"] = string(*obj.MiddlewareType)
	}

	result["patch_readiness_status"] = string(obj.PatchReadinessStatus)

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	if obj.WeblogicVersion != nil {
		result["weblogic_version"] = string(*obj.WeblogicVersion)
	}

	return result
}
