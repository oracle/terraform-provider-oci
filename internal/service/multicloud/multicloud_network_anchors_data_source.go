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

func MulticloudNetworkAnchorsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readMulticloudNetworkAnchors,
		Schema: map[string]*schema.Schema{
			// Required
			"subscription_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"subscription_service_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"external_location": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Optional
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_anchor_lifecycle_state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_anchor_oci_subnet_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_anchor_oci_vcn_id": {
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
			// Computed
			"network_anchor_collection": {
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
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"compartment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"resource_anchor_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"vcn_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"cluster_placement_group_id": {
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
									"network_anchor_lifecycle_state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"lifecycle_details": {
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

func readMulticloudNetworkAnchors(d *schema.ResourceData, m interface{}) error {
	sync := &MulticloudNetworkAnchorsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OmhubNetworkAnchorClient()

	return tfresource.ReadResource(sync)
}

type MulticloudNetworkAnchorsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_multicloud.OmhubNetworkAnchorClient
	Res    *oci_multicloud.ListNetworkAnchorsResponse
}

func (s *MulticloudNetworkAnchorsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MulticloudNetworkAnchorsDataSourceCrud) Get() error {
	request := oci_multicloud.ListNetworkAnchorsRequest{}

	// Required

	if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
		tmp := subscriptionId.(string)
		request.SubscriptionId = &tmp
	}

	if subscriptionServiceName, ok := s.D.GetOkExists("subscription_service_name"); ok {
		request.SubscriptionServiceName = oci_multicloud.ListNetworkAnchorsSubscriptionServiceNameEnum(subscriptionServiceName.(string))
	}

	if externalLocation, ok := s.D.GetOkExists("external_location"); ok {
		tmp := externalLocation.(string)
		request.ExternalLocation = &tmp
	}

	// Optional

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if lifecycleState, ok := s.D.GetOkExists("network_anchor_lifecycle_state"); ok {
		request.LifecycleState = oci_multicloud.NetworkAnchorLifecycleStateEnum(lifecycleState.(string))
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if networkAnchorOciSubnetId, ok := s.D.GetOkExists("network_anchor_oci_subnet_id"); ok {
		tmp := networkAnchorOciSubnetId.(string)
		request.NetworkAnchorOciSubnetId = &tmp
	}

	if networkAnchorOciVcnId, ok := s.D.GetOkExists("network_anchor_oci_vcn_id"); ok {
		tmp := networkAnchorOciVcnId.(string)
		request.NetworkAnchorOciVcnId = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if limit, ok := s.D.GetOkExists("limit"); ok {
		tmp := limit.(int)
		request.Limit = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "multicloud")

	response, err := s.Client.ListNetworkAnchors(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListNetworkAnchors(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *MulticloudNetworkAnchorsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("MulticloudNetworkAnchorsDataSource-", MulticloudNetworkAnchorsDataSource(), s.D))
	resources := []map[string]interface{}{}
	networkAnchor := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, NetworkAnchorSummaryToMap(item))
	}
	networkAnchor["items"] = items
	resources = append(resources, networkAnchor)

	if err := s.D.Set("network_anchor_collection", resources); err != nil {
		return err
	}

	return nil
}

func NetworkAnchorSummaryToMap(obj oci_multicloud.NetworkAnchorSummary) map[string]interface{} {
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

	if obj.ResourceAnchorId != nil {
		result["resource_anchor_id"] = string(*obj.ResourceAnchorId)
	}

	if obj.VcnId != nil {
		result["vcn_id"] = string(*obj.VcnId)
	}

	if obj.ClusterPlacementGroupId != nil {
		result["cluster_placement_group_id"] = string(*obj.ClusterPlacementGroupId)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	result["network_anchor_lifecycle_state"] = string(obj.LifecycleState)

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	return result
}
