// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package service_catalog

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_service_catalog "github.com/oracle/oci-go-sdk/v65/servicecatalog"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

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
	sync.Client = m.(*client.OracleClients).ServiceCatalogClient()

	return tfresource.ReadResource(sync)
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "service_catalog")

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

	s.D.SetId(tfresource.GenerateDataSourceHashID("ServiceCatalogPrivateApplicationPackageDataSource-", ServiceCatalogPrivateApplicationPackageDataSource(), s.D))
	switch v := (s.Res.PrivateApplicationPackage).(type) {
	case oci_service_catalog.PrivateApplicationStackPackage:
		s.D.Set("package_type", "STACK")

		if v.ContentUrl != nil {
			s.D.Set("content_url", *v.ContentUrl)
		}

		if v.MimeType != nil {
			s.D.Set("mime_type", *v.MimeType)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.PrivateApplicationId != nil {
			s.D.Set("private_application_id", *v.PrivateApplicationId)
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.Version != nil {
			s.D.Set("version", *v.Version)
		}
	default:
		log.Printf("[WARN] Received 'package_type' of unknown type %v", s.Res.PrivateApplicationPackage)
		return nil
	}

	return nil
}
