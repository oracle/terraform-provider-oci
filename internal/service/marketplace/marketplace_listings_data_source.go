// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package marketplace

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_marketplace "github.com/oracle/oci-go-sdk/v65/marketplace"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func MarketplaceListingsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readMarketplaceListings,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"category": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"image_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_featured": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"listing_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"listing_types": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"name": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"operating_systems": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"package_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"pricing": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"publisher_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"listings": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"banner": {
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
									"file_extension": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"mime_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"compatible_architectures": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"default_package_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"documentation_links": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"document_category": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"url": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_featured": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"listing_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"short_description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"regions": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"code": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"countries": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"code": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"name": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"icon": {
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
									"file_extension": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"mime_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": {
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
						"pricing_types": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"supported_operating_systems": {
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
						"categories": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"publisher": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
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
	}
}

func readMarketplaceListings(d *schema.ResourceData, m interface{}) error {
	sync := &MarketplaceListingsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MarketplaceClient()

	return tfresource.ReadResource(sync)
}

type MarketplaceListingsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_marketplace.MarketplaceClient
	Res    *oci_marketplace.ListListingsResponse
}

func (s *MarketplaceListingsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MarketplaceListingsDataSourceCrud) Get() error {
	request := oci_marketplace.ListListingsRequest{}

	if category, ok := s.D.GetOkExists("category"); ok {
		interfaces := category.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("category") {
			request.Category = tmp
		}
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if imageId, ok := s.D.GetOkExists("image_id"); ok {
		tmp := imageId.(string)
		request.ImageId = &tmp
	}

	if isFeatured, ok := s.D.GetOkExists("is_featured"); ok {
		tmp := isFeatured.(bool)
		request.IsFeatured = &tmp
	}

	if listingId, ok := s.D.GetOkExists("id"); ok {
		tmp := listingId.(string)
		request.ListingId = &tmp
	}

	if listingTypes, ok := s.D.GetOkExists("listing_types"); ok {
		interfaces := listingTypes.([]interface{})
		tmp := make([]oci_marketplace.ListingTypeEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_marketplace.ListingTypeEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("listing_types") {
			request.ListingTypes = tmp
		}
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		interfaces := name.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("name") {
			request.Name = tmp
		}
	}

	if operatingSystems, ok := s.D.GetOkExists("operating_systems"); ok {
		interfaces := operatingSystems.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("operating_systems") {
			request.OperatingSystems = tmp
		}
	}

	if packageType, ok := s.D.GetOkExists("package_type"); ok {
		tmp := packageType.(string)
		request.PackageType = &tmp
	}

	if pricing, ok := s.D.GetOkExists("pricing"); ok {
		interfaces := pricing.([]interface{})
		tmp := make([]oci_marketplace.PricingTypeEnumEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_marketplace.PricingTypeEnumEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("pricing") {
			request.Pricing = tmp
		}
	}

	if publisherId, ok := s.D.GetOkExists("publisher_id"); ok {
		tmp := publisherId.(string)
		request.PublisherId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "marketplace")

	response, err := s.Client.ListListings(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListListings(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *MarketplaceListingsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("MarketplaceListingsDataSource-", MarketplaceListingsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		listing := map[string]interface{}{}

		listing["categories"] = r.Categories

		listing["compatible_architectures"] = r.CompatibleArchitectures

		if r.Icon != nil {
			listing["icon"] = []interface{}{UploadDataToMap(r.Icon)}
		} else {
			listing["icon"] = nil
		}

		if r.Id != nil {
			listing["id"] = *r.Id
		}

		if r.IsFeatured != nil {
			listing["is_featured"] = *r.IsFeatured
		}

		listing["listing_type"] = r.ListingType

		if r.Name != nil {
			listing["name"] = *r.Name
		}

		listing["package_type"] = r.PackageType

		listing["pricing_types"] = r.PricingTypes

		if r.Publisher != nil {
			listing["publisher"] = []interface{}{PublisherSummaryToMap(r.Publisher)}
		} else {
			listing["publisher"] = nil
		}

		if r.Regions != nil {
			regions := []interface{}{}
			for _, item := range r.Regions {
				regions = append(regions, MarketplaceListingsRegionToMap(item))
			}
			listing["regions"] = regions
		}

		if r.ShortDescription != nil {
			listing["short_description"] = *r.ShortDescription
		}

		supportedOperatingSystems := []interface{}{}
		for _, item := range r.SupportedOperatingSystems {
			supportedOperatingSystems = append(supportedOperatingSystems, MarketplaceListingsOperatingSystemToMap(item))
		}
		listing["supported_operating_systems"] = supportedOperatingSystems

		resources = append(resources, listing)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, MarketplaceListingsDataSource().Schema["listings"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("listings", resources); err != nil {
		return err
	}

	return nil
}

func MarketplaceListingsRegionToMap(obj oci_marketplace.Region) interface{} {
	result := map[string]interface{}{}

	if obj.Code != nil {
		result["code"] = string(*obj.Code)
	}

	countries := []interface{}{}
	for _, item := range obj.Countries {
		countries = append(countries, MarketplaceListingPackagesItemToMap(item))
	}
	result["countries"] = countries

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}

func PublisherSummaryToMap(obj *oci_marketplace.PublisherSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Id != nil {
		result["id"] = *obj.Id
	}

	if obj.Description != nil {
		result["description"] = *obj.Description
	}

	if obj.Name != nil {
		result["name"] = *obj.Name
	}
	return result
}

func MarketplaceListingsOperatingSystemToMap(obj oci_marketplace.OperatingSystem) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}
