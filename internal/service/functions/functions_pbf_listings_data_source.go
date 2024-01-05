// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package functions

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_functions "github.com/oracle/oci-go-sdk/v65/functions"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func FunctionsPbfListingsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readFunctionsPbfListings,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name_contains": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name_starts_with": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"pbf_listing_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"trigger": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"pbf_listings_collection": {
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
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"publisher_details": {
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

func readFunctionsPbfListings(d *schema.ResourceData, m interface{}) error {
	sync := &FunctionsPbfListingsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FunctionsManagementClient()

	return tfresource.ReadResource(sync)
}

type FunctionsPbfListingsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_functions.FunctionsManagementClient
	Res    *oci_functions.ListPbfListingsResponse
}

func (s *FunctionsPbfListingsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FunctionsPbfListingsDataSourceCrud) Get() error {
	request := oci_functions.ListPbfListingsRequest{}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if nameContains, ok := s.D.GetOkExists("name_contains"); ok {
		tmp := nameContains.(string)
		request.NameContains = &tmp
	}

	if nameStartsWith, ok := s.D.GetOkExists("name_starts_with"); ok {
		tmp := nameStartsWith.(string)
		request.NameStartsWith = &tmp
	}

	if pbfListingId, ok := s.D.GetOkExists("pbf_listing_id"); ok {
		tmp := pbfListingId.(string)
		request.PbfListingId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_functions.PbfListingLifecycleStateEnum(state.(string))
	}

	if trigger, ok := s.D.GetOkExists("trigger"); ok {
		interfaces := trigger.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("trigger") {
			request.Trigger = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "functions")

	response, err := s.Client.ListPbfListings(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListPbfListings(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *FunctionsPbfListingsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("FunctionsPbfListingsDataSource-", FunctionsPbfListingsDataSource(), s.D))
	resources := []map[string]interface{}{}
	pbfListing := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, PbfListingSummaryToMap(item))
	}
	pbfListing["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, FunctionsPbfListingsDataSource().Schema["pbf_listings_collection"].Elem.(*schema.Resource).Schema)
		pbfListing["items"] = items
	}

	resources = append(resources, pbfListing)
	if err := s.D.Set("pbf_listings_collection", resources); err != nil {
		return err
	}

	return nil
}

func PbfListingSummaryToMap(obj oci_functions.PbfListingSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.PublisherDetails != nil {
		result["publisher_details"] = []interface{}{PublisherDetailsToMap(obj.PublisherDetails)}
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

func PublisherDetailsToMap(obj *oci_functions.PublisherDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}

func TriggerToMap(obj oci_functions.Trigger) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}
