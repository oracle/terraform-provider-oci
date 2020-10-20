// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_core "github.com/oracle/oci-go-sdk/v27/core"
)

func init() {
	RegisterResource("oci_core_cross_connect", CoreCrossConnectResource())
}

func CoreCrossConnectResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createCoreCrossConnect,
		Read:     readCoreCrossConnect,
		Update:   updateCoreCrossConnect,
		Delete:   deleteCoreCrossConnect,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"location_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"port_speed_shape_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"cross_connect_group_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"customer_reference_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: definedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"far_cross_connect_or_cross_connect_group_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"near_cross_connect_or_cross_connect_group_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"is_active": {
				Type:         schema.TypeBool,
				Optional:     true,
				ValidateFunc: validateBoolInSlice([]bool{true}),
			},

			// Computed
			"port_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createCoreCrossConnect(d *schema.ResourceData, m interface{}) error {
	sync := &CoreCrossConnectResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient()

	err := CreateResource(d, sync)
	if err != nil {
		return err
	}

	// Issue an Update if 'is_active' is set to true
	if _, ok := sync.D.GetOkExists("is_active"); ok {
		log.Printf("[DEBUG] CrossConnect resource is set to be active, calling 'Update' for the resource")
		return UpdateResource(d, sync)
	}

	return nil
}

func readCoreCrossConnect(d *schema.ResourceData, m interface{}) error {
	sync := &CoreCrossConnectResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient()

	return ReadResource(sync)
}

func updateCoreCrossConnect(d *schema.ResourceData, m interface{}) error {
	sync := &CoreCrossConnectResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient()

	return UpdateResource(d, sync)
}

func deleteCoreCrossConnect(d *schema.ResourceData, m interface{}) error {
	sync := &CoreCrossConnectResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient()
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type CoreCrossConnectResourceCrud struct {
	BaseCrud
	Client                 *oci_core.VirtualNetworkClient
	Res                    *oci_core.CrossConnect
	DisableNotFoundRetries bool
}

func (s *CoreCrossConnectResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CoreCrossConnectResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.CrossConnectLifecycleStateProvisioning),
	}
}

func (s *CoreCrossConnectResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.CrossConnectLifecycleStatePendingCustomer),
		string(oci_core.CrossConnectLifecycleStateProvisioned),
	}
}

func (s *CoreCrossConnectResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.CrossConnectLifecycleStateTerminating),
	}
}

func (s *CoreCrossConnectResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.CrossConnectLifecycleStateTerminated),
	}
}

func (s *CoreCrossConnectResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_core.CrossConnectLifecycleStateProvisioning),
	}
}

func (s *CoreCrossConnectResourceCrud) UpdatedTarget() []string {
	if _, ok := s.D.GetOkExists("is_active"); ok {
		log.Printf("[DEBUG] CrossConnect resource is set to be active, wait until the state is '%s'", string(oci_core.CrossConnectLifecycleStateProvisioned))
		return []string{
			string(oci_core.CrossConnectLifecycleStateProvisioned),
		}
	}

	return []string{
		string(oci_core.CrossConnectLifecycleStatePendingCustomer),
		string(oci_core.CrossConnectLifecycleStateProvisioned),
	}
}

func (s *CoreCrossConnectResourceCrud) Create() error {
	request := oci_core.CreateCrossConnectRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if crossConnectGroupId, ok := s.D.GetOkExists("cross_connect_group_id"); ok {
		tmp := crossConnectGroupId.(string)
		request.CrossConnectGroupId = &tmp
	}

	if customerReferenceName, ok := s.D.GetOkExists("customer_reference_name"); ok {
		tmp := customerReferenceName.(string)
		request.CustomerReferenceName = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if farCrossConnectOrCrossConnectGroupId, ok := s.D.GetOkExists("far_cross_connect_or_cross_connect_group_id"); ok {
		tmp := farCrossConnectOrCrossConnectGroupId.(string)
		request.FarCrossConnectOrCrossConnectGroupId = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if locationName, ok := s.D.GetOkExists("location_name"); ok {
		tmp := locationName.(string)
		request.LocationName = &tmp
	}

	if nearCrossConnectOrCrossConnectGroupId, ok := s.D.GetOkExists("near_cross_connect_or_cross_connect_group_id"); ok {
		tmp := nearCrossConnectOrCrossConnectGroupId.(string)
		request.NearCrossConnectOrCrossConnectGroupId = &tmp
	}

	if portSpeedShapeName, ok := s.D.GetOkExists("port_speed_shape_name"); ok {
		tmp := portSpeedShapeName.(string)
		request.PortSpeedShapeName = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateCrossConnect(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.CrossConnect
	return nil
}

func (s *CoreCrossConnectResourceCrud) Get() error {
	request := oci_core.GetCrossConnectRequest{}

	tmp := s.D.Id()
	request.CrossConnectId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetCrossConnect(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.CrossConnect
	return nil
}

func (s *CoreCrossConnectResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_core.UpdateCrossConnectRequest{}

	tmp := s.D.Id()
	request.CrossConnectId = &tmp

	if customerReferenceName, ok := s.D.GetOkExists("customer_reference_name"); ok {
		tmp := customerReferenceName.(string)
		request.CustomerReferenceName = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
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
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	// Cross Connect Resource can be set to 'Active' only once when the resource is 'PENDING_CUSTOMER' and not 'PROVISIONED'
	if isActive, ok := s.D.GetOkExists("is_active"); ok {
		if state, ok := s.D.GetOkExists("state"); ok && state.(string) == string(oci_core.CrossConnectLifecycleStatePendingCustomer) {
			log.Printf("[DEBUG] Cross Connect is in a valid state: '%s' to be set to active", state.(string))
			tmp := isActive.(bool)
			request.IsActive = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateCrossConnect(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.CrossConnect
	return nil
}

func (s *CoreCrossConnectResourceCrud) Delete() error {
	request := oci_core.DeleteCrossConnectRequest{}

	tmp := s.D.Id()
	request.CrossConnectId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteCrossConnect(context.Background(), request)
	return err
}

func (s *CoreCrossConnectResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CrossConnectGroupId != nil {
		s.D.Set("cross_connect_group_id", *s.Res.CrossConnectGroupId)
	}

	if s.Res.CustomerReferenceName != nil {
		s.D.Set("customer_reference_name", *s.Res.CustomerReferenceName)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LocationName != nil {
		s.D.Set("location_name", *s.Res.LocationName)
	}

	if s.Res.PortName != nil {
		s.D.Set("port_name", *s.Res.PortName)
	}

	if s.Res.PortSpeedShapeName != nil {
		s.D.Set("port_speed_shape_name", *s.Res.PortSpeedShapeName)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func (s *CoreCrossConnectResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_core.ChangeCrossConnectCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.CrossConnectId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.ChangeCrossConnectCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}
	return nil
}
