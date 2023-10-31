// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package functions

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_functions "github.com/oracle/oci-go-sdk/v65/functions"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func FunctionsPbfListingVersionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readFunctionsPbfListingVersions,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"is_current_version": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"pbf_listing_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"pbf_listing_version_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"pbf_listing_versions_collection": {
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
									"change_summary": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"config": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"description": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"is_optional": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"key": {
													Type:     schema.TypeString,
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
									"freeform_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"pbf_listing_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"requirements": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"min_memory_required_in_mbs": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"policies": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"description": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"policy": {
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
											},
										},
									},
									"state": {
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
									"triggers": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"name": {
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
				},
			},
		},
	}
}

func readFunctionsPbfListingVersions(d *schema.ResourceData, m interface{}) error {
	sync := &FunctionsPbfListingVersionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FunctionsManagementClient()

	return tfresource.ReadResource(sync)
}

type FunctionsPbfListingVersionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_functions.FunctionsManagementClient
	Res    *oci_functions.ListPbfListingVersionsResponse
}

func (s *FunctionsPbfListingVersionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FunctionsPbfListingVersionsDataSourceCrud) Get() error {
	request := oci_functions.ListPbfListingVersionsRequest{}

	if isCurrentVersion, ok := s.D.GetOkExists("is_current_version"); ok {
		tmp := isCurrentVersion.(bool)
		request.IsCurrentVersion = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if pbfListingId, ok := s.D.GetOkExists("pbf_listing_id"); ok {
		tmp := pbfListingId.(string)
		request.PbfListingId = &tmp
	}

	if pbfListingVersionId, ok := s.D.GetOkExists("id"); ok {
		tmp := pbfListingVersionId.(string)
		request.PbfListingVersionId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_functions.PbfListingVersionLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "functions")

	response, err := s.Client.ListPbfListingVersions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListPbfListingVersions(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *FunctionsPbfListingVersionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("FunctionsPbfListingVersionsDataSource-", FunctionsPbfListingVersionsDataSource(), s.D))
	resources := []map[string]interface{}{}
	pbfListingVersion := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, PbfListingVersionSummaryToMap(item))
	}
	pbfListingVersion["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, FunctionsPbfListingVersionsDataSource().Schema["pbf_listing_versions_collection"].Elem.(*schema.Resource).Schema)
		pbfListingVersion["items"] = items
	}

	resources = append(resources, pbfListingVersion)
	if err := s.D.Set("pbf_listing_versions_collection", resources); err != nil {
		return err
	}

	return nil
}

func ConfigDetailsToMap(obj oci_functions.ConfigDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.IsOptional != nil {
		result["is_optional"] = bool(*obj.IsOptional)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	return result
}

func PbfListingVersionSummaryToMap(obj oci_functions.PbfListingVersionSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ChangeSummary != nil {
		result["change_summary"] = string(*obj.ChangeSummary)
	}

	config := []interface{}{}
	for _, item := range obj.Config {
		config = append(config, ConfigDetailsToMap(item))
	}
	result["config"] = config

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.PbfListingId != nil {
		result["pbf_listing_id"] = string(*obj.PbfListingId)
	}

	if obj.Requirements != nil {
		result["requirements"] = []interface{}{RequirementDetailsToMap(obj.Requirements)}
	}

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

	triggers := []interface{}{}
	for _, item := range obj.Triggers {
		triggers = append(triggers, TriggerToMap(item))
	}
	result["triggers"] = triggers

	return result
}

func PolicyDetailsToMap(obj oci_functions.PolicyDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.Policy != nil {
		result["policy"] = string(*obj.Policy)
	}

	return result
}

func RequirementDetailsToMap(obj *oci_functions.RequirementDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.MinMemoryRequiredInMBs != nil {
		result["min_memory_required_in_mbs"] = strconv.FormatInt(*obj.MinMemoryRequiredInMBs, 10)
	}

	policies := []interface{}{}
	for _, item := range obj.Policies {
		policies = append(policies, PolicyDetailsToMap(item))
	}
	result["policies"] = policies

	return result
}
