// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package multicloud

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_multicloud "github.com/oracle/oci-go-sdk/v65/multicloud"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func MulticloudResourceAnchorsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readMulticloudResourceAnchors,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"linked_compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"lifecycle_state": {
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
			"limit": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"is_compartment_id_in_subtree": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"should_fetch_compartment_name": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"subscription_service_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"subscription_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_anchor_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Computed
									"compartment_id": {
										Type:     schema.TypeString,
										Computed: true,
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
									"lifecycle_details": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"compartment_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"partner_cloud_account_identifier": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"csp_resource_anchor_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"csp_resource_anchor_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"csp_additional_properties": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"lifecycle_state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"subscription_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"system_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"linked_compartment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"linked_compartment_name": {
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

func readMulticloudResourceAnchors(d *schema.ResourceData, m interface{}) error {
	sync := &MulticloudResourceAnchorsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OmhubResourceAnchorClient()

	return tfresource.ReadResource(sync)
}

type MulticloudResourceAnchorsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_multicloud.OmhubResourceAnchorClient
	Res    *oci_multicloud.ListResourceAnchorsResponse
}

func (s *MulticloudResourceAnchorsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MulticloudResourceAnchorsDataSourceCrud) Get() error {
	request := oci_multicloud.ListResourceAnchorsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if linkedCompartmentId, ok := s.D.GetOkExists("linked_compartment_id"); ok {
		tmp := linkedCompartmentId.(string)
		request.LinkedCompartmentId = &tmp
	}

	if state, ok := s.D.GetOkExists("lifecycle_state"); ok {
		request.LifecycleState = oci_multicloud.ResourceAnchorLifecycleStateEnum(state.(string))
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if limit, ok := s.D.GetOkExists("limit"); ok {
		tmp := limit.(int)
		request.Limit = &tmp
	}

	if isCompartmentIdInSubtree, ok := s.D.GetOkExists("is_compartment_id_in_subtree"); ok {
		tmp := isCompartmentIdInSubtree.(bool)
		request.IsCompartmentIdInSubtree = &tmp
	}

	if shouldFetchCompartmentName, ok := s.D.GetOkExists("should_fetch_compartment_name"); ok {
		tmp := shouldFetchCompartmentName.(bool)
		request.ShouldFetchCompartmentName = &tmp
	}

	if subscriptionServiceName, ok := s.D.GetOkExists("subscription_service_name"); ok {
		request.SubscriptionServiceName = oci_multicloud.ListResourceAnchorsSubscriptionServiceNameEnum(subscriptionServiceName.(string))
	}

	if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
		tmp := subscriptionId.(string)
		request.SubscriptionId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "multicloud")

	response, err := s.Client.ListResourceAnchors(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListResourceAnchors(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *MulticloudResourceAnchorsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("MulticloudResourceAnchorsDataSource-", MulticloudResourceAnchorsDataSource(), s.D))
	resources := []map[string]interface{}{}
	resourceAnchor := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ResourceAnchorSummaryToMap(item))
	}
	resourceAnchor["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, MulticloudResourceAnchorsDataSource().Schema["resource_anchor_collection"].Elem.(*schema.Resource).Schema)
		resourceAnchor["items"] = items
	}

	resources = append(resources, resourceAnchor)
	if err := s.D.Set("resource_anchor_collection", resources); err != nil {
		return err
	}

	return nil
}

func ResourceAnchorSummaryToMap(obj oci_multicloud.ResourceAnchorSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.CompartmentName != nil {
		result["compartment_name"] = string(*obj.CompartmentName)
	}

	if obj.PartnerCloudAccountIdentifier != nil {
		result["partner_cloud_account_identifier"] = string(*obj.PartnerCloudAccountIdentifier)
	}

	if obj.CspResourceAnchorId != nil {
		result["csp_resource_anchor_id"] = string(*obj.CspResourceAnchorId)
	}

	if obj.CspResourceAnchorName != nil {
		result["csp_resource_anchor_name"] = string(*obj.CspResourceAnchorName)
	}

	result["csp_additional_properties"] = obj.CspAdditionalProperties

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	result["lifecycle_state"] = string(obj.LifecycleState)

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.SubscriptionId != nil {
		result["subscription_id"] = string(*obj.SubscriptionId)
	}

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.LinkedCompartmentId != nil {
		result["linked_compartment_id"] = string(*obj.LinkedCompartmentId)
	}

	if obj.LinkedCompartmentName != nil {
		result["linked_compartment_name"] = string(*obj.LinkedCompartmentName)
	}

	return result
}
