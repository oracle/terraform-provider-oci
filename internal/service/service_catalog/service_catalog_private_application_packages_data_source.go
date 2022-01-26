// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package service_catalog

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_service_catalog "github.com/oracle/oci-go-sdk/v56/servicecatalog"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func ServiceCatalogPrivateApplicationPackagesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readServiceCatalogPrivateApplicationPackages,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"package_type": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"private_application_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"private_application_package_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"private_application_package_collection": {
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
									"content_url": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"mime_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"package_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"private_application_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_created": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"version": {
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

func readServiceCatalogPrivateApplicationPackages(d *schema.ResourceData, m interface{}) error {
	sync := &ServiceCatalogPrivateApplicationPackagesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ServiceCatalogClient()

	return tfresource.ReadResource(sync)
}

type ServiceCatalogPrivateApplicationPackagesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_service_catalog.ServiceCatalogClient
	Res    *oci_service_catalog.ListPrivateApplicationPackagesResponse
}

func (s *ServiceCatalogPrivateApplicationPackagesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ServiceCatalogPrivateApplicationPackagesDataSourceCrud) Get() error {
	request := oci_service_catalog.ListPrivateApplicationPackagesRequest{}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if packageType, ok := s.D.GetOkExists("package_type"); ok {
		interfaces := packageType.([]interface{})
		tmp := make([]oci_service_catalog.PackageTypeEnumEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_service_catalog.PackageTypeEnumEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("package_type") {
			request.PackageType = tmp
		}
	}

	if privateApplicationId, ok := s.D.GetOkExists("private_application_id"); ok {
		tmp := privateApplicationId.(string)
		request.PrivateApplicationId = &tmp
	}

	if privateApplicationPackageId, ok := s.D.GetOkExists("id"); ok {
		tmp := privateApplicationPackageId.(string)
		request.PrivateApplicationPackageId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "service_catalog")

	response, err := s.Client.ListPrivateApplicationPackages(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListPrivateApplicationPackages(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ServiceCatalogPrivateApplicationPackagesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ServiceCatalogPrivateApplicationPackagesDataSource-", ServiceCatalogPrivateApplicationPackagesDataSource(), s.D))
	resources := []map[string]interface{}{}
	privateApplicationPackage := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, PrivateApplicationPackagesSummaryToMap(item))
	}
	privateApplicationPackage["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, ServiceCatalogPrivateApplicationPackagesDataSource().Schema["private_application_package_collection"].Elem.(*schema.Resource).Schema)
		privateApplicationPackage["items"] = items
	}

	resources = append(resources, privateApplicationPackage)
	if err := s.D.Set("private_application_package_collection", resources); err != nil {
		return err
	}

	return nil
}

func PrivateApplicationPackagesSummaryToMap(obj oci_service_catalog.PrivateApplicationPackageSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	result["package_type"] = string(obj.PackageType)

	if obj.PrivateApplicationId != nil {
		result["private_application_id"] = string(*obj.PrivateApplicationId)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.Version != nil {
		result["version"] = string(*obj.Version)
	}

	return result
}
