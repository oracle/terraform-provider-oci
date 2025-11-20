// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package service_catalog

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_service_catalog "github.com/oracle/oci-go-sdk/v65/servicecatalog"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ServiceCatalogAllApplicationsDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readServiceCatalogAllApplicationsWithContext,
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
			"entity_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"entity_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_featured": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"package_type": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"pricing": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"publisher_id": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"application_collection": {
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
									"categories": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"entity_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"entity_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_featured": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"logo": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"content_url": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"display_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"mime_type": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"package_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"pricing_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"publisher": {
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
												"id": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"short_description": {
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

func readServiceCatalogAllApplicationsWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &ServiceCatalogAllApplicationsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ServiceCatalogClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type ServiceCatalogAllApplicationsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_service_catalog.ServiceCatalogClient
	Res    *oci_service_catalog.ListAllApplicationsResponse
}

func (s *ServiceCatalogAllApplicationsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ServiceCatalogAllApplicationsDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_service_catalog.ListAllApplicationsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if entityId, ok := s.D.GetOkExists("entity_id"); ok {
		tmp := entityId.(string)
		request.EntityId = &tmp
	}

	if entityType, ok := s.D.GetOkExists("entity_type"); ok {
		tmp := entityType.(string)
		request.EntityType = &tmp
	}

	if isFeatured, ok := s.D.GetOkExists("is_featured"); ok {
		tmp := isFeatured.(bool)
		request.IsFeatured = &tmp
	}

	if packageType, ok := s.D.GetOkExists("package_type"); ok {
		interfaces := packageType.([]interface{})
		tmp := make([]oci_service_catalog.PackageTypeEnumEnum, 0, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp = append(tmp, oci_service_catalog.PackageTypeEnumEnum(interfaces[i].(string)))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("package_type") {
			request.PackageType = tmp
		}
	}

	if pricing, ok := s.D.GetOkExists("pricing"); ok {
		interfaces := pricing.([]interface{})
		tmp := make([]oci_service_catalog.PricingTypeEnumEnum, 0, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp = append(tmp, oci_service_catalog.PricingTypeEnumEnum(interfaces[i].(string)))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("pricing") {
			request.Pricing = tmp
		}
	}

	if publisherId, ok := s.D.GetOkExists("publisher_id"); ok {
		interfaces := publisherId.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("publisher_id") {
			request.PublisherId = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "service_catalog")

	response, err := s.Client.ListAllApplications(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAllApplications(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ServiceCatalogAllApplicationsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ServiceCatalogAllApplicationsDataSource-", ServiceCatalogAllApplicationsDataSource(), s.D))
	resources := []map[string]interface{}{}
	allApplication := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ApplicationSummaryToMap(item))
	}
	allApplication["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, ServiceCatalogAllApplicationsDataSource().Schema["application_collection"].Elem.(*schema.Resource).Schema)
		allApplication["items"] = items
	}

	resources = append(resources, allApplication)
	if err := s.D.Set("application_collection", resources); err != nil {
		return err
	}

	return nil
}

func ApplicationSummaryToMap(obj oci_service_catalog.ApplicationSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["categories"] = obj.Categories

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.EntityId != nil {
		result["entity_id"] = string(*obj.EntityId)
	}

	if obj.EntityType != nil {
		result["entity_type"] = string(*obj.EntityType)
	}

	if obj.IsFeatured != nil {
		result["is_featured"] = bool(*obj.IsFeatured)
	}

	if obj.Logo != nil {
		result["logo"] = []interface{}{UploadDataToMap(obj.Logo)}
	}

	result["package_type"] = string(obj.PackageType)

	result["pricing_type"] = string(obj.PricingType)

	if obj.Publisher != nil {
		result["publisher"] = []interface{}{PublisherSummaryToMap(obj.Publisher)}
	}

	if obj.ShortDescription != nil {
		result["short_description"] = string(*obj.ShortDescription)
	}

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	return result
}

func PublisherSummaryToMap(obj *oci_service_catalog.PublisherSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	return result
}

func UploadDataToMap(obj *oci_service_catalog.UploadData) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ContentUrl != nil {
		result["content_url"] = string(*obj.ContentUrl)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.MimeType != nil {
		result["mime_type"] = string(*obj.MimeType)
	}

	return result
}
