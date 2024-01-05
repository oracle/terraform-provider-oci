// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package waf

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_waf "github.com/oracle/oci-go-sdk/v65/waf"
)

func WafProtectionCapabilitiesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readWafProtectionCapabilities,
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
			"group_tag": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"is_latest_version": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeBool,
				},
			},
			"key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"protection_capability_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"collaborative_action_threshold": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"collaborative_weights": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"display_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"key": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"weight": {
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"group_tags": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"is_latest_version": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"key": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"version": {
										Type:     schema.TypeInt,
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

func readWafProtectionCapabilities(d *schema.ResourceData, m interface{}) error {
	sync := &WafProtectionCapabilitiesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).WafClient()

	return tfresource.ReadResource(sync)
}

type WafProtectionCapabilitiesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_waf.WafClient
	Res    *oci_waf.ListProtectionCapabilitiesResponse
}

func (s *WafProtectionCapabilitiesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *WafProtectionCapabilitiesDataSourceCrud) Get() error {
	request := oci_waf.ListProtectionCapabilitiesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if groupTag, ok := s.D.GetOkExists("group_tag"); ok {
		interfaces := groupTag.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("group_tag") {
			request.GroupTag = tmp
		}
	}

	if isLatestVersion, ok := s.D.GetOkExists("is_latest_version"); ok {
		interfaces := isLatestVersion.([]interface{})
		tmp := make([]bool, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(bool)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("is_latest_version") {
			request.IsLatestVersion = tmp
		}
	}

	if key, ok := s.D.GetOkExists("key"); ok {
		tmp := key.(string)
		request.Key = &tmp
	}

	if type_, ok := s.D.GetOkExists("type"); ok {
		request.Type = oci_waf.ProtectionCapabilitySummaryTypeEnum(type_.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "waf")

	response, err := s.Client.ListProtectionCapabilities(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListProtectionCapabilities(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *WafProtectionCapabilitiesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("WafProtectionCapabilitiesDataSource-", WafProtectionCapabilitiesDataSource(), s.D))
	resources := []map[string]interface{}{}
	protectionCapability := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ProtectionCapabilitySummaryToMap(item))
	}
	protectionCapability["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, WafProtectionCapabilitiesDataSource().Schema["protection_capability_collection"].Elem.(*schema.Resource).Schema)
		protectionCapability["items"] = items
	}

	resources = append(resources, protectionCapability)
	if err := s.D.Set("protection_capability_collection", resources); err != nil {
		return err
	}

	return nil
}

func CollaborativeCapabilityWeightToMap(obj oci_waf.CollaborativeCapabilityWeight) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.Weight != nil {
		result["weight"] = int(*obj.Weight)
	}

	return result
}

func ProtectionCapabilitySummaryToMap(obj oci_waf.ProtectionCapabilitySummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CollaborativeActionThreshold != nil {
		result["collaborative_action_threshold"] = int(*obj.CollaborativeActionThreshold)
	}

	collaborativeWeights := []interface{}{}
	for _, item := range obj.CollaborativeWeights {
		collaborativeWeights = append(collaborativeWeights, CollaborativeCapabilityWeightToMap(item))
	}
	result["collaborative_weights"] = collaborativeWeights

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["group_tags"] = obj.GroupTags

	if obj.IsLatestVersion != nil {
		result["is_latest_version"] = bool(*obj.IsLatestVersion)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	result["type"] = string(obj.Type)

	if obj.Version != nil {
		result["version"] = int(*obj.Version)
	}

	return result
}
