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

func MulticloudExternalLocationMappingMetadataDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readMulticloudExternalLocationMappingMetadata,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"subscription_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"subscription_service_name": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"limit": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"external_location_mapping_metadatum_summary_collection": {
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
												"csp_physical_az": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"csp_physical_az_display_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"csp_region": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"csp_region_display_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"service_name": {
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
									"oci_logical_ad": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"oci_physical_ad": {
										Type:     schema.TypeString,
										Computed: true,
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

func readMulticloudExternalLocationMappingMetadata(d *schema.ResourceData, m interface{}) error {
	sync := &MulticloudExternalLocationMappingMetadataDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MetadataClient()

	return tfresource.ReadResource(sync)
}

type MulticloudExternalLocationMappingMetadataDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_multicloud.MetadataClient
	Res    *oci_multicloud.ListExternalLocationMappingMetadataResponse
}

func (s *MulticloudExternalLocationMappingMetadataDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MulticloudExternalLocationMappingMetadataDataSourceCrud) Get() error {
	request := oci_multicloud.ListExternalLocationMappingMetadataRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
		tmp := subscriptionId.(string)
		request.SubscriptionId = &tmp
	}

	if subscriptionServiceName, ok := s.D.GetOkExists("subscription_service_name"); ok {
		interfaces := subscriptionServiceName.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("subscription_service_name") {
			for _, tmpStr := range tmp {
				request.SubscriptionServiceName = append(request.SubscriptionServiceName, oci_multicloud.SubscriptionTypeEnum(tmpStr))
			}
		}
	}

	if limit, ok := s.D.GetOkExists("limit"); ok {
		tmp := limit.(int)
		request.Limit = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "multicloud")

	response, err := s.Client.ListExternalLocationMappingMetadata(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListExternalLocationMappingMetadata(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *MulticloudExternalLocationMappingMetadataDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("MulticloudExternalLocationMappingMetadataDataSource-", MulticloudExternalLocationMappingMetadataDataSource(), s.D))
	resources := []map[string]interface{}{}
	externalLocationMappingMetadata := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ExternalLocationMappingMetadatumSummaryToMap(item))
	}
	externalLocationMappingMetadata["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, MulticloudExternalLocationMappingMetadataDataSource().Schema["external_location_mapping_metadatum_summary_collection"].Elem.(*schema.Resource).Schema)
		externalLocationMappingMetadata["items"] = items
	}

	resources = append(resources, externalLocationMappingMetadata)
	if err := s.D.Set("external_location_mapping_metadatum_summary_collection", resources); err != nil {
		return err
	}

	return nil
}

func ExternalLocationToMap(obj *oci_multicloud.ExternalLocation) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CspPhysicalAz != nil {
		result["csp_physical_az"] = string(*obj.CspPhysicalAz)
	}

	if obj.CspPhysicalAzDisplayName != nil {
		result["csp_physical_az_display_name"] = string(*obj.CspPhysicalAzDisplayName)
	}

	if obj.CspRegion != nil {
		result["csp_region"] = string(*obj.CspRegion)
	}

	if obj.CspRegionDisplayName != nil {
		result["csp_region_display_name"] = string(*obj.CspRegionDisplayName)
	}

	result["service_name"] = string(obj.ServiceName)

	return result
}

func ExternalLocationMappingMetadatumSummaryToMap(obj oci_multicloud.ExternalLocationMappingMetadatumSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.ExternalLocation != nil {
		result["external_location"] = []interface{}{ExternalLocationToMap(obj.ExternalLocation)}
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.OciLogicalAd != nil {
		result["oci_logical_ad"] = string(*obj.OciLogicalAd)
	}

	if obj.OciPhysicalAd != nil {
		result["oci_physical_ad"] = string(*obj.OciPhysicalAd)
	}

	if obj.OciRegion != nil {
		result["oci_region"] = string(*obj.OciRegion)
	}

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	return result
}
