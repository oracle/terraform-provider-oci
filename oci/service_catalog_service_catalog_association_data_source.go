// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_service_catalog "github.com/oracle/oci-go-sdk/v42/servicecatalog"
)

func init() {
	RegisterDatasource("oci_service_catalog_service_catalog_association", ServiceCatalogServiceCatalogAssociationDataSource())
}

func ServiceCatalogServiceCatalogAssociationDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["service_catalog_association_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(ServiceCatalogServiceCatalogAssociationResource(), fieldMap, readSingularServiceCatalogServiceCatalogAssociation)
}

func readSingularServiceCatalogServiceCatalogAssociation(d *schema.ResourceData, m interface{}) error {
	sync := &ServiceCatalogServiceCatalogAssociationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).serviceCatalogClient()

	return ReadResource(sync)
}

type ServiceCatalogServiceCatalogAssociationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_service_catalog.ServiceCatalogClient
	Res    *oci_service_catalog.GetServiceCatalogAssociationResponse
}

func (s *ServiceCatalogServiceCatalogAssociationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ServiceCatalogServiceCatalogAssociationDataSourceCrud) Get() error {
	request := oci_service_catalog.GetServiceCatalogAssociationRequest{}

	if serviceCatalogAssociationId, ok := s.D.GetOkExists("service_catalog_association_id"); ok {
		tmp := serviceCatalogAssociationId.(string)
		request.ServiceCatalogAssociationId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "service_catalog")

	response, err := s.Client.GetServiceCatalogAssociation(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ServiceCatalogServiceCatalogAssociationDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.EntityId != nil {
		s.D.Set("entity_id", *s.Res.EntityId)
	}

	if s.Res.EntityType != nil {
		s.D.Set("entity_type", *s.Res.EntityType)
	}

	if s.Res.ServiceCatalogId != nil {
		s.D.Set("service_catalog_id", *s.Res.ServiceCatalogId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
