// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package multicloud

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_multicloud "github.com/oracle/oci-go-sdk/v65/multicloud"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func MulticloudMulticloudpoliciesDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readMulticloudMulticloudpoliciesWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_force_refresh": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"subscription_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"multicloud_policy_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"compartment_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
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
									"groups": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"policies": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"compartment_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"compartment_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"defined_tags": {
													Type:     schema.TypeMap,
													Computed: true,
													Elem:     schema.TypeString,
												},
												"description": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"freeform_tags": {
													Type:     schema.TypeMap,
													Computed: true,
													Elem:     schema.TypeString,
												},
												"name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"lifecycle_state": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"statements": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"system_tags": {
													Type:     schema.TypeMap,
													Computed: true,
													Elem:     schema.TypeString,
												},
											},
										},
									},
									"lifecycle_state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"subscription_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"subscription_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"system_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
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

func readMulticloudMulticloudpoliciesWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &MulticloudMulticloudpoliciesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MulticloudPoliciesClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type MulticloudMulticloudpoliciesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_multicloud.MulticloudPoliciesClient
	Res    *oci_multicloud.ListMulticloudPoliciesResponse
}

func (s *MulticloudMulticloudpoliciesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MulticloudMulticloudpoliciesDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_multicloud.ListMulticloudPoliciesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if isForceRefresh, ok := s.D.GetOkExists("is_force_refresh"); ok {
		tmp := isForceRefresh.(bool)
		request.IsForceRefresh = &tmp
	}

	if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
		tmp := subscriptionId.(string)
		request.SubscriptionId = &tmp
	}

	if limit, ok := s.D.GetOkExists("limit"); ok {
		tmp := limit.(int)
		request.Limit = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "multicloud")

	response, err := s.Client.ListMulticloudPolicies(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListMulticloudPolicies(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *MulticloudMulticloudpoliciesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("MulticloudMulticloudpoliciesDataSource-", MulticloudMulticloudpoliciesDataSource(), s.D))
	resources := []map[string]interface{}{}
	multicloudpolicy := map[string]interface{}{}

	if s.Res.CompartmentId != nil {
		multicloudpolicy["compartment_id"] = string(*s.Res.CompartmentId)
	}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, MulticloudPolicySummaryToMap(item))
	}
	multicloudpolicy["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, MulticloudMulticloudpoliciesDataSource().Schema["multicloud_policy_collection"].Elem.(*schema.Resource).Schema)
		multicloudpolicy["items"] = items
	}

	resources = append(resources, multicloudpolicy)
	if err := s.D.Set("multicloud_policy_collection", resources); err != nil {
		return err
	}

	return nil
}

func MulticloudPolicyToMap(obj oci_multicloud.MulticloudPolicy) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.CompartmentName != nil {
		result["compartment_name"] = string(*obj.CompartmentName)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	result["lifecycle_state"] = string(obj.LifecycleState)

	result["statements"] = obj.Statements

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	return result
}

func MulticloudPolicySummaryToMap(obj oci_multicloud.MulticloudPolicySummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	result["freeform_tags"] = obj.FreeformTags

	result["groups"] = obj.Groups

	policies := []interface{}{}
	for _, item := range obj.Policies {
		policies = append(policies, MulticloudPolicyToMap(item))
	}
	result["policies"] = policies

	result["lifecycle_state"] = string(obj.LifecycleState)

	if obj.SubscriptionId != nil {
		result["subscription_id"] = string(*obj.SubscriptionId)
	}

	result["subscription_type"] = string(obj.SubscriptionType)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	return result
}
