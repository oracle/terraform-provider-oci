// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package service_catalog

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_service_catalog "github.com/oracle/oci-go-sdk/v58/servicecatalog"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

func ServiceCatalogServiceCatalogResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createServiceCatalogServiceCatalog,
		Read:     readServiceCatalogServiceCatalog,
		Update:   updateServiceCatalogServiceCatalog,
		Delete:   deleteServiceCatalogServiceCatalog,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createServiceCatalogServiceCatalog(d *schema.ResourceData, m interface{}) error {
	sync := &ServiceCatalogServiceCatalogResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ServiceCatalogClient()

	return tfresource.CreateResource(d, sync)
}

func readServiceCatalogServiceCatalog(d *schema.ResourceData, m interface{}) error {
	sync := &ServiceCatalogServiceCatalogResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ServiceCatalogClient()

	return tfresource.ReadResource(sync)
}

func updateServiceCatalogServiceCatalog(d *schema.ResourceData, m interface{}) error {
	sync := &ServiceCatalogServiceCatalogResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ServiceCatalogClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteServiceCatalogServiceCatalog(d *schema.ResourceData, m interface{}) error {
	sync := &ServiceCatalogServiceCatalogResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ServiceCatalogClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type ServiceCatalogServiceCatalogResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_service_catalog.ServiceCatalogClient
	Res                    *oci_service_catalog.ServiceCatalog
	DisableNotFoundRetries bool
}

func (s *ServiceCatalogServiceCatalogResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *ServiceCatalogServiceCatalogResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *ServiceCatalogServiceCatalogResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_service_catalog.ServiceCatalogLifecycleStateActive),
	}
}

func (s *ServiceCatalogServiceCatalogResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *ServiceCatalogServiceCatalogResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_service_catalog.ServiceCatalogLifecycleStateDeleted),
	}
}

func (s *ServiceCatalogServiceCatalogResourceCrud) Create() error {
	request := oci_service_catalog.CreateServiceCatalogRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "service_catalog")

	response, err := s.Client.CreateServiceCatalog(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ServiceCatalog
	return nil
}

func (s *ServiceCatalogServiceCatalogResourceCrud) Get() error {
	request := oci_service_catalog.GetServiceCatalogRequest{}

	tmp := s.D.Id()
	request.ServiceCatalogId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "service_catalog")

	response, err := s.Client.GetServiceCatalog(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ServiceCatalog
	return nil
}

func (s *ServiceCatalogServiceCatalogResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_service_catalog.UpdateServiceCatalogRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.ServiceCatalogId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "service_catalog")

	response, err := s.Client.UpdateServiceCatalog(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ServiceCatalog
	return nil
}

func (s *ServiceCatalogServiceCatalogResourceCrud) Delete() error {
	request := oci_service_catalog.DeleteServiceCatalogRequest{}

	tmp := s.D.Id()
	request.ServiceCatalogId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "service_catalog")

	_, err := s.Client.DeleteServiceCatalog(context.Background(), request)
	return err
}

func (s *ServiceCatalogServiceCatalogResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func ServiceCatalogSummaryToMap(obj oci_service_catalog.ServiceCatalogSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	return result
}

func (s *ServiceCatalogServiceCatalogResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_service_catalog.ChangeServiceCatalogCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.ServiceCatalogId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "service_catalog")

	_, err := s.Client.ChangeServiceCatalogCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
