// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package marketplace

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_marketplace "github.com/oracle/oci-go-sdk/v56/marketplace"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func MarketplaceListingDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularMarketplaceListing,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"listing_id": {
				Type:     schema.TypeString,
				Required: true,
			},
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
			"categories": {
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
			"is_featured": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"keywords": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"languages": {
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
			"license_model_description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"links": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"href": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"rel": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"listing_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"long_description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"package_type": {
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
						"contact_email": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"contact_phone": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"hq_address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"links": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"href": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"rel": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"logo": {
							Type:     schema.TypeList,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
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
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"website_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"year_founded": {
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
			"release_notes": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"screenshots": {
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
						"description": {
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
			"short_description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"support_contacts": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"email": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"phone": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"subject": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"support_links": {
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
						"url": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
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
			"system_requirements": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tagline": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_released": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"usage_information": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"videos": {
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
						"url": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readSingularMarketplaceListing(d *schema.ResourceData, m interface{}) error {
	sync := &MarketplaceListingDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MarketplaceClient()

	return tfresource.ReadResource(sync)
}

type MarketplaceListingDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_marketplace.MarketplaceClient
	Res    *oci_marketplace.GetListingResponse
}

func (s *MarketplaceListingDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MarketplaceListingDataSourceCrud) Get() error {
	request := oci_marketplace.GetListingRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if listingId, ok := s.D.GetOkExists("listing_id"); ok {
		tmp := listingId.(string)
		request.ListingId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "marketplace")

	response, err := s.Client.GetListing(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *MarketplaceListingDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.Banner != nil {
		s.D.Set("banner", []interface{}{MarketplaceListingUploadDataToMap(s.Res.Banner)})
	} else {
		s.D.Set("banner", nil)
	}

	s.D.Set("categories", s.Res.Categories)

	if s.Res.DefaultPackageVersion != nil {
		s.D.Set("default_package_version", *s.Res.DefaultPackageVersion)
	}

	documentationLinks := []interface{}{}
	for _, item := range s.Res.DocumentationLinks {
		documentationLinks = append(documentationLinks, DocumentationLinkToMap(item))
	}
	s.D.Set("documentation_links", documentationLinks)

	if s.Res.Icon != nil {
		s.D.Set("icon", []interface{}{MarketplaceListingUploadDataToMap(s.Res.Icon)})
	} else {
		s.D.Set("icon", nil)
	}

	if s.Res.IsFeatured != nil {
		s.D.Set("is_featured", *s.Res.IsFeatured)
	}

	if s.Res.Keywords != nil {
		s.D.Set("keywords", *s.Res.Keywords)
	}

	languages := []interface{}{}
	for _, item := range s.Res.Languages {
		languages = append(languages, MarketplaceListingItemToMap(item))
	}
	s.D.Set("languages", languages)

	if s.Res.LicenseModelDescription != nil {
		s.D.Set("license_model_description", *s.Res.LicenseModelDescription)
	}

	links := []interface{}{}
	for _, item := range s.Res.Links {
		links = append(links, MarketplaceListingLinkToMap(item))
	}
	s.D.Set("links", links)

	s.D.Set("listing_type", s.Res.ListingType)

	if s.Res.LongDescription != nil {
		s.D.Set("long_description", *s.Res.LongDescription)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("package_type", s.Res.PackageType)

	if s.Res.Publisher != nil {
		s.D.Set("publisher", []interface{}{MarketplaceListingPublisherToMap(s.Res.Publisher)})
	} else {
		s.D.Set("publisher", nil)
	}

	regions := []interface{}{}
	for _, item := range s.Res.Regions {
		regions = append(regions, MarketplaceListingRegionToMap(item))
	}
	s.D.Set("regions", regions)

	if s.Res.ReleaseNotes != nil {
		s.D.Set("release_notes", *s.Res.ReleaseNotes)
	}

	screenshots := []interface{}{}
	for _, item := range s.Res.Screenshots {
		screenshots = append(screenshots, ScreenshotToMap(item))
	}
	s.D.Set("screenshots", screenshots)

	if s.Res.ShortDescription != nil {
		s.D.Set("short_description", *s.Res.ShortDescription)
	}

	supportContacts := []interface{}{}
	for _, item := range s.Res.SupportContacts {
		supportContacts = append(supportContacts, MarketplaceListingSupportContactToMap(item))
	}
	s.D.Set("support_contacts", supportContacts)

	supportLinks := []interface{}{}
	for _, item := range s.Res.SupportLinks {
		supportLinks = append(supportLinks, MarketplaceListingNamedLinkToMap(item))
	}
	s.D.Set("support_links", supportLinks)

	supportedOperatingSystems := []interface{}{}
	for _, item := range s.Res.SupportedOperatingSystems {
		supportedOperatingSystems = append(supportedOperatingSystems, MarketplaceListingOperatingSystemToMap(item))
	}
	s.D.Set("supported_operating_systems", supportedOperatingSystems)

	if s.Res.SystemRequirements != nil {
		s.D.Set("system_requirements", *s.Res.SystemRequirements)
	}

	if s.Res.Tagline != nil {
		s.D.Set("tagline", *s.Res.Tagline)
	}

	if s.Res.TimeReleased != nil {
		s.D.Set("time_released", s.Res.TimeReleased.String())
	}

	if s.Res.UsageInformation != nil {
		s.D.Set("usage_information", *s.Res.UsageInformation)
	}

	if s.Res.Version != nil {
		s.D.Set("version", *s.Res.Version)
	}

	videos := []interface{}{}
	for _, item := range s.Res.Videos {
		videos = append(videos, MarketplaceListingNamedLinkToMap(item))
	}
	s.D.Set("videos", videos)

	return nil
}

func DocumentationLinkToMap(obj oci_marketplace.DocumentationLink) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DocumentCategory != nil {
		result["document_category"] = string(*obj.DocumentCategory)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Url != nil {
		result["url"] = string(*obj.Url)
	}

	return result
}

func MarketplaceListingItemToMap(obj oci_marketplace.Item) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Code != nil {
		result["code"] = string(*obj.Code)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}

func MarketplaceListingLinkToMap(obj oci_marketplace.Link) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Href != nil {
		result["href"] = string(*obj.Href)
	}

	result["rel"] = string(obj.Rel)

	return result
}

func MarketplaceListingNamedLinkToMap(obj oci_marketplace.NamedLink) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Url != nil {
		result["url"] = string(*obj.Url)
	}

	return result
}

func MarketplaceListingOperatingSystemToMap(obj oci_marketplace.OperatingSystem) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}

func MarketplaceListingPublisherToMap(obj *oci_marketplace.Publisher) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ContactEmail != nil {
		result["contact_email"] = string(*obj.ContactEmail)
	}

	if obj.ContactPhone != nil {
		result["contact_phone"] = string(*obj.ContactPhone)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.HqAddress != nil {
		result["hq_address"] = string(*obj.HqAddress)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	links := []interface{}{}
	for _, item := range obj.Links {
		links = append(links, MarketplaceListingLinkToMap(item))
	}
	result["links"] = links

	if obj.Logo != nil {
		result["logo"] = []interface{}{MarketplaceListingUploadDataToMap(obj.Logo)}
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.WebsiteUrl != nil {
		result["website_url"] = string(*obj.WebsiteUrl)
	}

	if obj.YearFounded != nil {
		result["year_founded"] = strconv.FormatInt(*obj.YearFounded, 10)
	}

	return result
}

func MarketplaceListingRegionToMap(obj oci_marketplace.Region) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Code != nil {
		result["code"] = string(*obj.Code)
	}

	countries := []interface{}{}
	for _, item := range obj.Countries {
		countries = append(countries, MarketplaceListingItemToMap(item))
	}
	result["countries"] = countries

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}

func ScreenshotToMap(obj oci_marketplace.Screenshot) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ContentUrl != nil {
		result["content_url"] = string(*obj.ContentUrl)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.FileExtension != nil {
		result["file_extension"] = string(*obj.FileExtension)
	}

	if obj.MimeType != nil {
		result["mime_type"] = string(*obj.MimeType)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}

func MarketplaceListingSupportContactToMap(obj oci_marketplace.SupportContact) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Email != nil {
		result["email"] = string(*obj.Email)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Phone != nil {
		result["phone"] = string(*obj.Phone)
	}

	if obj.Subject != nil {
		result["subject"] = string(*obj.Subject)
	}

	return result
}

func MarketplaceListingUploadDataToMap(obj *oci_marketplace.UploadData) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ContentUrl != nil {
		result["content_url"] = string(*obj.ContentUrl)
	}

	if obj.FileExtension != nil {
		result["file_extension"] = string(*obj.FileExtension)
	}

	if obj.MimeType != nil {
		result["mime_type"] = string(*obj.MimeType)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}
