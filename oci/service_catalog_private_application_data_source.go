// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_service_catalog "github.com/oracle/oci-go-sdk/v41/servicecatalog"
)

func init() {
	RegisterDatasource("oci_service_catalog_private_application", ServiceCatalogPrivateApplicationDataSource())
}

func ServiceCatalogPrivateApplicationDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["private_application_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(ServiceCatalogPrivateApplicationResource(), fieldMap, readSingularServiceCatalogPrivateApplication)
}

func readSingularServiceCatalogPrivateApplication(d *schema.ResourceData, m interface{}) error {
	sync := &ServiceCatalogPrivateApplicationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).serviceCatalogClient()

	return ReadResource(sync)
}

type ServiceCatalogPrivateApplicationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_service_catalog.ServiceCatalogClient
	Res    *oci_service_catalog.GetPrivateApplicationResponse
}

func (s *ServiceCatalogPrivateApplicationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ServiceCatalogPrivateApplicationDataSourceCrud) Get() error {
	request := oci_service_catalog.GetPrivateApplicationRequest{}

	if privateApplicationId, ok := s.D.GetOkExists("private_application_id"); ok {
		tmp := privateApplicationId.(string)
		request.PrivateApplicationId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "service_catalog")

	response, err := s.Client.GetPrivateApplication(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ServiceCatalogPrivateApplicationDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Logo != nil {
		s.D.Set("logo", []interface{}{SCUploadDataToMap(s.Res.Logo)})
	} else {
		s.D.Set("logo", nil)
	}

	if s.Res.LongDescription != nil {
		s.D.Set("long_description", *s.Res.LongDescription)
	}

	s.D.Set("package_type", s.Res.PackageType)

	if s.Res.ShortDescription != nil {
		s.D.Set("short_description", *s.Res.ShortDescription)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
