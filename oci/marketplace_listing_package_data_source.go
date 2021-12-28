// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_marketplace "github.com/oracle/oci-go-sdk/v54/marketplace"
)

func init() {
	RegisterDatasource("oci_marketplace_listing_package", MarketplaceListingPackageDataSource())
}

func MarketplaceListingPackageDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularMarketplaceListingPackage,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"listing_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"package_version": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"app_catalog_listing_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"app_catalog_listing_resource_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"image_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"operating_system": {
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
			"package_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"pricing": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"currency": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"pay_go_strategy": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"rate": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
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
			"resource_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"resource_link": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"variables": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"data_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"default_value": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"hint_message": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_mandatory": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"version": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularMarketplaceListingPackage(d *schema.ResourceData, m interface{}) error {
	sync := &MarketplaceListingPackageDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).marketplaceClient()

	return ReadResource(sync)
}

type MarketplaceListingPackageDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_marketplace.MarketplaceClient
	Res    *oci_marketplace.ListingPackage
}

func (s *MarketplaceListingPackageDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MarketplaceListingPackageDataSourceCrud) Get() error {
	request := oci_marketplace.GetPackageRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if listingId, ok := s.D.GetOkExists("listing_id"); ok {
		tmp := listingId.(string)
		request.ListingId = &tmp
	}

	if packageVersion, ok := s.D.GetOkExists("package_version"); ok {
		tmp := packageVersion.(string)
		request.PackageVersion = &tmp
	}

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(false, "marketplace")

	response, err := s.Client.GetPackage(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ListingPackage
	return nil
}

func (s *MarketplaceListingPackageDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("MarketplaceListingPackageDataSource-", MarketplaceListingPackageDataSource(), s.D))
	switch v := (*s.Res).(type) {
	case oci_marketplace.ImageListingPackage:

		if v.AppCatalogListingId != nil {
			s.D.Set("app_catalog_listing_id", v.AppCatalogListingId)
		}

		if v.AppCatalogListingResourceVersion != nil {
			s.D.Set("app_catalog_listing_resource_version", v.AppCatalogListingResourceVersion)
		}

		if v.ImageId != nil {
			s.D.Set("image_id", v.ImageId)
		}

		if v.Description != nil {
			s.D.Set("description", v.Description)
		}

		if v.ListingId != nil {
			s.D.Set("Listing_id", v.ListingId)
		}

		s.D.Set("package_type", oci_marketplace.PackageTypeEnumImage)

		if v.Pricing != nil {
			s.D.Set("pricing", []interface{}{MarketplaceListingPackagePricingModelToMap(v.Pricing)})
		}

		if v.Regions != nil {
			regions := []interface{}{}
			for _, item := range v.Regions {
				regions = append(regions, MarketplaceListingPackageRegionToMap(item))
			}
			s.D.Set("regions", regions)
		}

		if v.ResourceId != nil {
			s.D.Set("resource_id", v.ResourceId)
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.OperatingSystem != nil {
			s.D.Set("operating_system", []interface{}{MarketplaceListingPackageOperatingSystemToMap(v.OperatingSystem)})
		} else {
			s.D.Set("operating_system", nil)
		}

		if v.Version != nil {
			s.D.Set("version", v.Version)
		}
	case oci_marketplace.OrchestrationListingPackage:
		if v.Description != nil {
			s.D.Set("description", v.Description)
		}

		if v.ListingId != nil {
			s.D.Set("Listing_id", v.ListingId)
		}

		s.D.Set("package_type", oci_marketplace.PackageTypeEnumOrchestration)

		if v.Pricing != nil {
			s.D.Set("pricing", []interface{}{MarketplaceListingPackagePricingModelToMap(v.Pricing)})
		}

		if v.ResourceId != nil {
			s.D.Set("resource_id", v.ResourceId)
		}

		if v.ResourceLink != nil {
			s.D.Set("resource_link", v.ResourceLink)
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.OperatingSystem != nil {
			s.D.Set("operating_system", []interface{}{MarketplaceListingPackageOperatingSystemToMap(v.OperatingSystem)})
		} else {
			s.D.Set("operating_system", nil)
		}

		if v.Variables != nil {
			variables := []interface{}{}
			for _, item := range v.Variables {
				variables = append(variables, MarketplaceListingPackageOrchestrationVariableToMap(item))
			}
			s.D.Set("variables", variables)
		}

		if v.Version != nil {
			s.D.Set("version", v.Version)
		}
	default:
		log.Printf("[WARN] Received 'ListingPackage' of unknown type %v", *s.Res)
		return nil
	}
	return nil
}

func MarketplaceListingPackageRegionToMap(obj oci_marketplace.Region) interface{} {
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

func MarketplaceListingPackageOperatingSystemToMap(obj *oci_marketplace.OperatingSystem) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}

func MarketplaceListingPackageOrchestrationVariableToMap(obj oci_marketplace.OrchestrationVariable) map[string]interface{} {
	result := map[string]interface{}{}

	result["data_type"] = string(obj.DataType)

	if obj.DefaultValue != nil {
		result["default_value"] = string(*obj.DefaultValue)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.HintMessage != nil {
		result["hint_message"] = string(*obj.HintMessage)
	}

	if obj.IsMandatory != nil {
		result["is_mandatory"] = bool(*obj.IsMandatory)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}

func MarketplaceListingPackagePricingModelToMap(obj *oci_marketplace.PricingModel) map[string]interface{} {
	result := map[string]interface{}{}

	result["currency"] = string(obj.Currency)

	result["pay_go_strategy"] = string(obj.PayGoStrategy)

	if obj.Rate != nil {
		result["rate"] = float32(*obj.Rate)
	}

	result["type"] = string(obj.Type)

	return result
}
