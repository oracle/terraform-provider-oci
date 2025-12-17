// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v65/core"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CoreFirmwareBundlesDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readCoreFirmwareBundlesWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_default_bundle": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"platform": {
				Type:     schema.TypeString,
				Required: true,
			},
			"lifecycle_state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"firmware_bundles_collection": {
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

									// Computed
									"allowable_transitions": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"downgrades": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"upgrades": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
									"compartment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"platforms": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"platform": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"versions": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"component_type": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"version": {
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
														},
													},
												},
											},
										},
									},
									"lifecycle_state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_created": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_updated": {
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

func readCoreFirmwareBundlesWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &CoreFirmwareBundlesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type CoreFirmwareBundlesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.ComputeClient
	Res    *oci_core.ListFirmwareBundlesResponse
}

func (s *CoreFirmwareBundlesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreFirmwareBundlesDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_core.ListFirmwareBundlesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if isDefaultBundle, ok := s.D.GetOkExists("is_default_bundle"); ok {
		tmp := isDefaultBundle.(bool)
		request.IsDefaultBundle = &tmp
	}

	if platform, ok := s.D.GetOkExists("platform"); ok {
		tmp := platform.(string)
		request.Platform = &tmp
	}

	if lifecycleState, ok := s.D.GetOkExists("lifecycle_state"); ok {
		tmp := lifecycleState.(string)
		request.LifecycleState = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.ListFirmwareBundles(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListFirmwareBundles(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreFirmwareBundlesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreFirmwareBundlesDataSource-", CoreFirmwareBundlesDataSource(), s.D))
	resources := []map[string]interface{}{}
	firmwareBundle := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, FirmwareBundleSummaryToMap(item))
	}
	firmwareBundle["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, CoreFirmwareBundlesDataSource().Schema["firmware_bundles_collection"].Elem.(*schema.Resource).Schema)
		firmwareBundle["items"] = items
	}

	resources = append(resources, firmwareBundle)
	if err := s.D.Set("firmware_bundles_collection", resources); err != nil {
		return err
	}

	return nil
}

func ComponentVersionToMap(obj oci_core.ComponentVersion) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ComponentType != nil {
		result["component_type"] = string(*obj.ComponentType)
	}

	result["version"] = obj.Version

	return result
}

func FirmwareBundleSummaryToMap(obj oci_core.FirmwareBundleSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AllowableTransitions != nil {
		result["allowable_transitions"] = []interface{}{FirmwareBundleTransitionsToMap(obj.AllowableTransitions)}
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	platforms := []interface{}{}
	for _, item := range obj.Platforms {
		platforms = append(platforms, PlatformVersionsToMap(item))
	}
	result["platforms"] = platforms

	result["lifecycle_state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func FirmwareBundleTransitionsToMap(obj *oci_core.FirmwareBundleTransitions) map[string]interface{} {
	result := map[string]interface{}{}

	result["downgrades"] = obj.Downgrades

	result["upgrades"] = obj.Upgrades

	return result
}

func PlatformVersionsToMap(obj oci_core.PlatformVersions) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Platform != nil {
		result["platform"] = string(*obj.Platform)
	}

	versions := []interface{}{}
	for _, item := range obj.Versions {
		versions = append(versions, ComponentVersionToMap(item))
	}
	result["versions"] = versions

	return result
}
