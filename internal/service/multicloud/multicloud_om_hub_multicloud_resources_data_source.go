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

func MulticloudOmHubMulticloudResourcesDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readMulticloudOmHubMulticloudResourcesWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"subscription_service_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"subscription_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"resource_anchor_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"external_location": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"multicloud_resource_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Computed
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Computed
									"resource_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"resource_display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"resource_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"compartment_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"compartment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"vcn_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"vcn_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"network_anchor_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"network_anchor_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"csp_resource_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_created": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"csp_additional_properties": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"time_updated": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"lifecycle_state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"freeform_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"defined_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
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

func readMulticloudOmHubMulticloudResourcesWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &MulticloudOmHubMulticloudResourcesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MulticloudResourcesClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type MulticloudOmHubMulticloudResourcesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_multicloud.MulticloudResourcesClient
	Res    *oci_multicloud.ListMulticloudResourcesResponse
}

func (s *MulticloudOmHubMulticloudResourcesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MulticloudOmHubMulticloudResourcesDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_multicloud.ListMulticloudResourcesRequest{}

	if subscriptionServiceName, ok := s.D.GetOkExists("subscription_service_name"); ok {
		request.SubscriptionServiceName = oci_multicloud.ListMulticloudResourcesSubscriptionServiceNameEnum(subscriptionServiceName.(string))
	}

	if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
		tmp := subscriptionId.(string)
		request.SubscriptionId = &tmp
	}

	if resourceAnchorId, ok := s.D.GetOkExists("resource_anchor_id"); ok {
		tmp := resourceAnchorId.(string)
		request.ResourceAnchorId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if limit, ok := s.D.GetOkExists("limit"); ok {
		tmp := limit.(int)
		request.Limit = &tmp
	}

	if externalLocation, ok := s.D.GetOkExists("external_location"); ok {
		tmp := externalLocation.(string)
		request.ExternalLocation = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "multicloud")

	response, err := s.Client.ListMulticloudResources(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListMulticloudResources(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *MulticloudOmHubMulticloudResourcesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("MulticloudOmHubMulticloudResourcesDataSource-", MulticloudOmHubMulticloudResourcesDataSource(), s.D))
	resources := []map[string]interface{}{}
	omHubMulticloudResource := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, MulticloudResourceSummaryToMap(item))
	}
	omHubMulticloudResource["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, MulticloudOmHubMulticloudResourcesDataSource().Schema["multicloud_resource_collection"].Elem.(*schema.Resource).Schema)
		omHubMulticloudResource["items"] = items
	}

	resources = append(resources, omHubMulticloudResource)
	if err := s.D.Set("multicloud_resource_collection", resources); err != nil {
		return err
	}

	return nil
}

func MulticloudResourceSummaryToMap(obj oci_multicloud.MulticloudResourceSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ResourceId != nil {
		result["resource_id"] = string(*obj.ResourceId)
	}

	if obj.ResourceDisplayName != nil {
		result["resource_display_name"] = string(*obj.ResourceDisplayName)
	}

	if obj.ResourceType != nil {
		result["resource_type"] = string(*obj.ResourceType)
	}

	if obj.CompartmentName != nil {
		result["compartment_name"] = string(*obj.CompartmentName)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.VcnName != nil {
		result["vcn_name"] = string(*obj.VcnName)
	}

	if obj.VcnId != nil {
		result["vcn_id"] = string(*obj.VcnId)
	}

	if obj.NetworkAnchorName != nil {
		result["network_anchor_name"] = string(*obj.NetworkAnchorName)
	}

	if obj.NetworkAnchorId != nil {
		result["network_anchor_id"] = string(*obj.NetworkAnchorId)
	}

	if obj.CspResourceId != nil {
		result["csp_resource_id"] = string(*obj.CspResourceId)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	result["csp_additional_properties"] = obj.CspAdditionalProperties

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	result["lifecycle_state"] = string(obj.LifecycleState)

	result["freeform_tags"] = obj.FreeformTags

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	return result
}
