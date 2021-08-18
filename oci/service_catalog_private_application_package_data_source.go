// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_service_catalog "github.com/oracle/oci-go-sdk/v46/servicecatalog"
)

func init() {
	RegisterDatasource("oci_service_catalog_private_application_package", ServiceCatalogPrivateApplicationPackageDataSource())
}

func ServiceCatalogPrivateApplicationPackageDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularServiceCatalogPrivateApplicationPackage,
		Schema: map[string]*schema.Schema{
			"private_application_package_id": {
				Type:     schema.TypeString,
				Required: true,
			},
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
	}
}

func readSingularServiceCatalogPrivateApplicationPackage(d *schema.ResourceData, m interface{}) error {
	sync := &ServiceCatalogPrivateApplicationPackageDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).serviceCatalogClient()

	return ReadResource(sync)
}

type ServiceCatalogPrivateApplicationPackageDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_service_catalog.ServiceCatalogClient
	Res    *oci_service_catalog.GetPrivateApplicationPackageResponse
}

func (s *ServiceCatalogPrivateApplicationPackageDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ServiceCatalogPrivateApplicationPackageDataSourceCrud) Get() error {
	request := oci_service_catalog.GetPrivateApplicationPackageRequest{}

	if privateApplicationPackageId, ok := s.D.GetOkExists("private_application_package_id"); ok {
		tmp := privateApplicationPackageId.(string)
		request.PrivateApplicationPackageId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "service_catalog")

	response, err := s.Client.GetPrivateApplicationPackage(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ServiceCatalogPrivateApplicationPackageDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("ServiceCatalogPrivateApplicationPackageDataSource-", ServiceCatalogPrivateApplicationPackageDataSource(), s.D))
	switch v := (s.Res.PrivateApplicationPackage).(type) {
	case oci_service_catalog.PrivateApplicationStackPackage:
		if v.ContentUrl != nil {
			s.D.Set("content_url", v.ContentUrl)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", v.DisplayName)
		}

		if v.MimeType != nil {
			s.D.Set("mime_type", v.MimeType)
		}

		s.D.Set("package_type", oci_service_catalog.PackageTypeEnumStack)

		if v.PrivateApplicationId != nil {
			s.D.Set("private_application_id", v.PrivateApplicationId)
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.Version != nil {
			s.D.Set("version", v.Version)
		}

	default:
		log.Printf("[WARN] Received 'PrivateApplicationPackage' of unknown type %v", *s.Res)
		return nil
	}

	return nil
}

func PrivateApplicationPackageSummaryToMap(obj oci_service_catalog.PrivateApplicationPackageSummary) map[string]interface{} {
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
