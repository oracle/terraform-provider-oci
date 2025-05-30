// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package apiaccesscontrol

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_apiaccesscontrol "github.com/oracle/oci-go-sdk/v65/apiaccesscontrol"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ApiaccesscontrolApiMetadataByEntityTypesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readApiaccesscontrolApiMetadataByEntityTypes,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"api_metadata_by_entity_type_collection": {
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
									"api_metadatas": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"api_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"attributes": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
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
												"entity_type": {
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
												"service_name": {
													Type:     schema.TypeString,
													Computed: true,
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
											},
										},
									},
									"defined_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"entity_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"freeform_tags": {
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

func readApiaccesscontrolApiMetadataByEntityTypes(d *schema.ResourceData, m interface{}) error {
	sync := &ApiaccesscontrolApiMetadataByEntityTypesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApiMetadataClient()

	return tfresource.ReadResource(sync)
}

type ApiaccesscontrolApiMetadataByEntityTypesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_apiaccesscontrol.ApiMetadataClient
	Res    *oci_apiaccesscontrol.ListApiMetadataByEntityTypesResponse
}

func (s *ApiaccesscontrolApiMetadataByEntityTypesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ApiaccesscontrolApiMetadataByEntityTypesDataSourceCrud) Get() error {
	request := oci_apiaccesscontrol.ListApiMetadataByEntityTypesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if resourceType, ok := s.D.GetOkExists("resource_type"); ok {
		tmp := resourceType.(string)
		request.ResourceType = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_apiaccesscontrol.ApiMetadataLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "apiaccesscontrol")

	response, err := s.Client.ListApiMetadataByEntityTypes(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListApiMetadataByEntityTypes(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ApiaccesscontrolApiMetadataByEntityTypesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ApiaccesscontrolApiMetadataByEntityTypesDataSource-", ApiaccesscontrolApiMetadataByEntityTypesDataSource(), s.D))
	resources := []map[string]interface{}{}
	apiMetadataByEntityType := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ApiMetadataByEntityTypeSummaryToMap(item))
	}
	apiMetadataByEntityType["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, ApiaccesscontrolApiMetadataByEntityTypesDataSource().Schema["api_metadata_by_entity_type_collection"].Elem.(*schema.Resource).Schema)
		apiMetadataByEntityType["items"] = items
	}

	resources = append(resources, apiMetadataByEntityType)
	if err := s.D.Set("api_metadata_by_entity_type_collection", resources); err != nil {
		return err
	}

	return nil
}

func ApiMetadataByEntityTypeSummaryToMap(obj oci_apiaccesscontrol.ApiMetadataByEntityTypeSummary) map[string]interface{} {
	result := map[string]interface{}{}

	apiMetadatas := []interface{}{}
	for _, item := range obj.ApiMetadatas {
		apiMetadatas = append(apiMetadatas, ApiMetadataSummaryToMap(item))
	}
	result["api_metadatas"] = apiMetadatas

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.EntityType != nil {
		result["entity_type"] = string(*obj.EntityType)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	return result
}
