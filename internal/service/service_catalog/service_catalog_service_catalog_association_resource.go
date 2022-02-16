// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package service_catalog

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_service_catalog "github.com/oracle/oci-go-sdk/v58/servicecatalog"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func ServiceCatalogServiceCatalogAssociationResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createServiceCatalogServiceCatalogAssociation,
		Read:     readServiceCatalogServiceCatalogAssociation,
		Delete:   deleteServiceCatalogServiceCatalogAssociation,
		Schema: map[string]*schema.Schema{
			// Required
			"entity_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"service_catalog_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"entity_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createServiceCatalogServiceCatalogAssociation(d *schema.ResourceData, m interface{}) error {
	sync := &ServiceCatalogServiceCatalogAssociationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ServiceCatalogClient()

	return tfresource.CreateResource(d, sync)
}

func readServiceCatalogServiceCatalogAssociation(d *schema.ResourceData, m interface{}) error {
	sync := &ServiceCatalogServiceCatalogAssociationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ServiceCatalogClient()

	return tfresource.ReadResource(sync)
}

func deleteServiceCatalogServiceCatalogAssociation(d *schema.ResourceData, m interface{}) error {
	sync := &ServiceCatalogServiceCatalogAssociationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ServiceCatalogClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type ServiceCatalogServiceCatalogAssociationResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_service_catalog.ServiceCatalogClient
	Res                    *oci_service_catalog.ServiceCatalogAssociation
	DisableNotFoundRetries bool
}

func (s *ServiceCatalogServiceCatalogAssociationResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *ServiceCatalogServiceCatalogAssociationResourceCrud) Create() error {
	request := oci_service_catalog.CreateServiceCatalogAssociationRequest{}

	if entityId, ok := s.D.GetOkExists("entity_id"); ok {
		tmp := entityId.(string)
		request.EntityId = &tmp
	}

	if entityType, ok := s.D.GetOkExists("entity_type"); ok {
		tmp := entityType.(string)
		request.EntityType = &tmp
	}

	if serviceCatalogId, ok := s.D.GetOkExists("service_catalog_id"); ok {
		tmp := serviceCatalogId.(string)
		request.ServiceCatalogId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "service_catalog")

	response, err := s.Client.CreateServiceCatalogAssociation(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ServiceCatalogAssociation
	return nil
}

func (s *ServiceCatalogServiceCatalogAssociationResourceCrud) Get() error {
	request := oci_service_catalog.GetServiceCatalogAssociationRequest{}

	tmp := s.D.Id()
	request.ServiceCatalogAssociationId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "service_catalog")

	response, err := s.Client.GetServiceCatalogAssociation(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ServiceCatalogAssociation
	return nil
}

func (s *ServiceCatalogServiceCatalogAssociationResourceCrud) Delete() error {
	request := oci_service_catalog.DeleteServiceCatalogAssociationRequest{}

	tmp := s.D.Id()
	request.ServiceCatalogAssociationId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "service_catalog")

	_, err := s.Client.DeleteServiceCatalogAssociation(context.Background(), request)
	return err
}

func (s *ServiceCatalogServiceCatalogAssociationResourceCrud) SetData() error {
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

func ServiceCatalogAssociationSummaryToMap(obj oci_service_catalog.ServiceCatalogAssociationSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.EntityId != nil {
		result["entity_id"] = string(*obj.EntityId)
	}

	if obj.EntityType != nil {
		result["entity_type"] = string(*obj.EntityType)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.ServiceCatalogId != nil {
		result["service_catalog_id"] = string(*obj.ServiceCatalogId)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	return result
}
