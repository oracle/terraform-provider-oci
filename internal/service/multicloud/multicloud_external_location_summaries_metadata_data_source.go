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

func MulticloudExternalLocationSummariesMetadataDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readMulticloudExternalLocationSummariesMetadata,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"entity_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"subscription_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"subscription_service_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"limit": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"external_location_summaries_metadatum_summary_collection": {
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
									"defined_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"external_location": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"csp_region": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"csp_region_display_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"freeform_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"oci_region": {
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

func readMulticloudExternalLocationSummariesMetadata(d *schema.ResourceData, m interface{}) error {
	sync := &MulticloudExternalLocationSummariesMetadataDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MetadataClient()

	return tfresource.ReadResource(sync)
}

type MulticloudExternalLocationSummariesMetadataDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_multicloud.MetadataClient
	Res    *oci_multicloud.ListExternalLocationSummariesMetadataResponse
}

func (s *MulticloudExternalLocationSummariesMetadataDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MulticloudExternalLocationSummariesMetadataDataSourceCrud) Get() error {
	request := oci_multicloud.ListExternalLocationSummariesMetadataRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if entityType, ok := s.D.GetOkExists("entity_type"); ok {
		request.EntityType = oci_multicloud.ListExternalLocationSummariesMetadataEntityTypeEnum(entityType.(string))
	}

	if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
		tmp := subscriptionId.(string)
		request.SubscriptionId = &tmp
	}

	if subscriptionServiceName, ok := s.D.GetOkExists("subscription_service_name"); ok {
		request.SubscriptionServiceName = oci_multicloud.ListExternalLocationSummariesMetadataSubscriptionServiceNameEnum(subscriptionServiceName.(string))
	}

	if limit, ok := s.D.GetOkExists("limit"); ok {
		tmp := limit.(int)
		request.Limit = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "multicloud")

	response, err := s.Client.ListExternalLocationSummariesMetadata(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListExternalLocationSummariesMetadata(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *MulticloudExternalLocationSummariesMetadataDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("MulticloudExternalLocationSummariesMetadataDataSource-", MulticloudExternalLocationSummariesMetadataDataSource(), s.D))
	resources := []map[string]interface{}{}
	externalLocationSummariesMetadata := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ExternalLocationSummariesMetadatumSummaryToMap(item))
	}
	externalLocationSummariesMetadata["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, MulticloudExternalLocationSummariesMetadataDataSource().Schema["external_location_summaries_metadatum_summary_collection"].Elem.(*schema.Resource).Schema)
		externalLocationSummariesMetadata["items"] = items
	}

	resources = append(resources, externalLocationSummariesMetadata)
	if err := s.D.Set("external_location_summaries_metadatum_summary_collection", resources); err != nil {
		return err
	}

	return nil
}

func ExternalLocationSummariesMetadatumSummaryToMap(obj oci_multicloud.ExternalLocationSummariesMetadatumSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.ExternalLocation != nil {
		result["external_location"] = []interface{}{ExternalLocationSummaryToMap(obj.ExternalLocation)}
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.OciRegion != nil {
		result["oci_region"] = string(*obj.OciRegion)
	}

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	return result
}

func ExternalLocationSummaryToMap(obj *oci_multicloud.ExternalLocationSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CspRegion != nil {
		result["csp_region"] = string(*obj.CspRegion)
	}

	if obj.CspRegionDisplayName != nil {
		result["csp_region_display_name"] = string(*obj.CspRegionDisplayName)
	}

	return result
}
