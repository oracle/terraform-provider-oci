// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package marketplace

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_marketplace "github.com/oracle/oci-go-sdk/v56/marketplace"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func MarketplacePublicationPackageDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularMarketplacePublicationPackage,
		Schema: map[string]*schema.Schema{
			"package_version": {
				Type:     schema.TypeString,
				Required: true,
			},
			"publication_id": {
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
			"listing_id": {
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

func readSingularMarketplacePublicationPackage(d *schema.ResourceData, m interface{}) error {
	sync := &MarketplacePublicationPackageDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MarketplaceClient()

	return tfresource.ReadResource(sync)
}

type MarketplacePublicationPackageDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_marketplace.MarketplaceClient
	Res    *oci_marketplace.GetPublicationPackageResponse
}

func (s *MarketplacePublicationPackageDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MarketplacePublicationPackageDataSourceCrud) Get() error {
	request := oci_marketplace.GetPublicationPackageRequest{}

	if packageVersion, ok := s.D.GetOkExists("package_version"); ok {
		tmp := packageVersion.(string)
		request.PackageVersion = &tmp
	}

	if publicationId, ok := s.D.GetOkExists("publication_id"); ok {
		tmp := publicationId.(string)
		request.PublicationId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "marketplace")

	response, err := s.Client.GetPublicationPackage(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *MarketplacePublicationPackageDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("MarketplacePublicationPackageDataSource-", MarketplacePublicationPackageDataSource(), s.D))
	switch v := (s.Res.PublicationPackage).(type) {
	case oci_marketplace.ImagePublicationPackage:
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
			s.D.Set("listing_id", v.ListingId)
		}

		s.D.Set("package_type", oci_marketplace.PackageTypeEnumImage)

		if v.ResourceId != nil {
			s.D.Set("resource_id", v.ResourceId)
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.OperatingSystem != nil {
			s.D.Set("operating_system", []interface{}{MarketplacePublicationPackageOperatingSystemToMap(v.OperatingSystem)})
		} else {
			s.D.Set("operating_system", nil)
		}

		if v.Version != nil {
			s.D.Set("version", v.Version)
		}

	case oci_marketplace.OrchestrationPublicationPackage:
		if v.Description != nil {
			s.D.Set("description", v.Description)
		}

		if v.ListingId != nil {
			s.D.Set("Listing_id", v.ListingId)
		}

		s.D.Set("package_type", oci_marketplace.PackageTypeEnumOrchestration)

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
			s.D.Set("operating_system", []interface{}{MarketplacePublicationPackageOperatingSystemToMap(v.OperatingSystem)})
		} else {
			s.D.Set("operating_system", nil)
		}

		if v.Variables != nil {
			variables := []interface{}{}
			for _, item := range v.Variables {
				variables = append(variables, MarketplacePublicationPackageOrchestrationVariableToMap(item))
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

func MarketplacePublicationPackageOperatingSystemToMap(obj *oci_marketplace.OperatingSystem) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}

func MarketplacePublicationPackageOrchestrationVariableToMap(obj oci_marketplace.OrchestrationVariable) map[string]interface{} {
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
