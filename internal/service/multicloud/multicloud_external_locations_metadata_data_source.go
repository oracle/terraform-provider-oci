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

func MulticloudExternalLocationsMetadataDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readMulticloudExternalLocationsMetadata,
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
			"linked_compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"subscription_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"subscription_service_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"limit": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"external_locations_metadatum_collection": {
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
									"cpg_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
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
												"csp_logical_az": {
													Type:     schema.TypeString,
													Computed: true,
												},
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
												"csp_zone_key_reference_id": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"key_name": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"key_value": {
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
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

func readMulticloudExternalLocationsMetadata(d *schema.ResourceData, m interface{}) error {
	sync := &MulticloudExternalLocationsMetadataDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MetadataClient()

	return tfresource.ReadResource(sync)
}

type MulticloudExternalLocationsMetadataDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_multicloud.MetadataClient
	Res    *oci_multicloud.ListExternalLocationDetailsMetadataResponse
}

func (s *MulticloudExternalLocationsMetadataDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MulticloudExternalLocationsMetadataDataSourceCrud) Get() error {
	request := oci_multicloud.ListExternalLocationDetailsMetadataRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if entityType, ok := s.D.GetOkExists("entity_type"); ok {
		request.EntityType = oci_multicloud.ListExternalLocationDetailsMetadataEntityTypeEnum(entityType.(string))
	}

	if linkedCompartmentId, ok := s.D.GetOkExists("linked_compartment_id"); ok {
		tmp := linkedCompartmentId.(string)
		request.LinkedCompartmentId = &tmp
	}

	if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
		tmp := subscriptionId.(string)
		request.SubscriptionId = &tmp
	}

	if subscriptionServiceName, ok := s.D.GetOkExists("subscription_service_name"); ok {
		request.SubscriptionServiceName = oci_multicloud.ListExternalLocationDetailsMetadataSubscriptionServiceNameEnum(subscriptionServiceName.(string))
	}

	if limit, ok := s.D.GetOkExists("limit"); ok {
		tmp := limit.(int)
		request.Limit = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "multicloud")

	response, err := s.Client.ListExternalLocationDetailsMetadata(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListExternalLocationDetailsMetadata(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *MulticloudExternalLocationsMetadataDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("MulticloudExternalLocationsMetadataDataSource-", MulticloudExternalLocationsMetadataDataSource(), s.D))
	resources := []map[string]interface{}{}
	externalLocationsMetadata := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ExternalLocationsMetadatumSummaryToMap(item))
	}
	externalLocationsMetadata["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, MulticloudExternalLocationsMetadataDataSource().Schema["external_locations_metadatum_collection"].Elem.(*schema.Resource).Schema)
		externalLocationsMetadata["items"] = items
	}

	resources = append(resources, externalLocationsMetadata)
	if err := s.D.Set("external_locations_metadatum_collection", resources); err != nil {
		return err
	}

	return nil
}

func CspZoneKeyReferenceIdToMap(obj *oci_multicloud.CspZoneKeyReferenceId) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.KeyName != nil {
		result["key_name"] = string(*obj.KeyName)
	}

	if obj.KeyValue != nil {
		result["key_value"] = string(*obj.KeyValue)
	}

	return result
}

func ExternalLocationDetailToMap(obj *oci_multicloud.ExternalLocationDetail) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CspLogicalAz != nil {
		result["csp_logical_az"] = string(*obj.CspLogicalAz)
	}

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

	if obj.CspZoneKeyReferenceId != nil {
		result["csp_zone_key_reference_id"] = []interface{}{CspZoneKeyReferenceIdToMap(obj.CspZoneKeyReferenceId)}
	}

	result["service_name"] = string(obj.ServiceName)

	return result
}

func ExternalLocationsMetadatumSummaryToMap(obj oci_multicloud.ExternalLocationsMetadatumSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CpgId != nil {
		result["cpg_id"] = string(*obj.CpgId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.ExternalLocation != nil {
		result["external_location"] = []interface{}{ExternalLocationDetailToMap(obj.ExternalLocation)}
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
