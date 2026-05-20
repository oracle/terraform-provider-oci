// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package resource_search

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_resource_search "github.com/oracle/oci-go-sdk/v65/resourcesearch"
	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ResourceSearchDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readSingularResourceSearchWithContext,
		Schema: map[string]*schema.Schema{
			"query": {
				Type:     schema.TypeString,
				Required: true,
			},
			"results": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"resource_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"identifier": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"compartment_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"availability_domain": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"state": {
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
						"additional_details": {
							Type:     schema.TypeMap,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readSingularResourceSearchWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &ResourceSearchDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ResourceSearchClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type ResourceSearchDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_resource_search.ResourceSearchClient
	Res    *oci_resource_search.ResourceSummaryCollection
}

func (s *ResourceSearchDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ResourceSearchDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_resource_search.SearchResourcesRequest{}
	if qRaw, ok := s.D.GetOk("query"); ok {
		queryStr := qRaw.(string)
		request.SearchDetails = oci_resource_search.StructuredSearchDetails{Query: &queryStr}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "resource_search")

	response, err := s.Client.SearchResources(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.ResourceSummaryCollection
	return nil
}

func (s *ResourceSearchDataSourceCrud) SetData() error {

	var items []interface{}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ResourceSearchResourceTypesDataSource-", ResourceSearchDataSource(), s.D))

	for _, item := range s.Res.Items {
		items = append(items, ResourceSummaryToMap(item))
	}
	// schema defines this block under the "results" attribute
	if err := s.D.Set("results", items); err != nil {
		return err
	}

	return nil
}

func ResourceSummaryToMap(obj oci_resource_search.ResourceSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AvailabilityDomain != nil {
		result["availability_domain"] = *obj.AvailabilityDomain
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = *obj.CompartmentId
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = *obj.DisplayName
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Identifier != nil {
		result["identifier"] = *obj.Identifier
	}

	if obj.ResourceType != nil {
		result["resource_type"] = *obj.ResourceType
	}

	if obj.LifecycleState != nil {
		// schema uses the key "state" for lifecycle state in the results block
		result["state"] = *obj.LifecycleState
	}

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.AdditionalDetails != nil {
		result["additional_details"] = additionalDetailsToStringMap(obj.AdditionalDetails)
	}
	return result
}

func additionalDetailsToStringMap(details map[string]interface{}) map[string]string {
	result := make(map[string]string, len(details))

	for key, value := range details {
		result[key] = additionalDetailValueToString(value)
	}

	return result
}

func additionalDetailValueToString(value interface{}) string {
	if stringValue, ok := value.(string); ok {
		return stringValue
	}

	marshaledValue, err := json.Marshal(value)
	if err != nil {
		return fmt.Sprintf("%v", value)
	}

	return string(marshaledValue)
}
