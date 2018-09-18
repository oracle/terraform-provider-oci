// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform/helper/schema"

	oci_core "github.com/oracle/oci-go-sdk/core"
)

func CrossConnectResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createCrossConnect,
		Read:     readCrossConnect,
		Update:   updateCrossConnect,
		Delete:   deleteCrossConnect,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
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

func createCrossConnect(d *schema.ResourceData, m interface{}) error {
	sync := &CrossConnectResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

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

func readCrossConnect(d *schema.ResourceData, m interface{}) error {
	sync := &CrossConnectResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return ReadResource(sync)
}

func updateCrossConnect(d *schema.ResourceData, m interface{}) error {
	sync := &CrossConnectResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return UpdateResource(d, sync)
}

func deleteCrossConnect(d *schema.ResourceData, m interface{}) error {
	sync := &CrossConnectResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type CrossConnectResourceCrud struct {
	BaseCrud
	Client                 *oci_core.VirtualNetworkClient
	Res                    *oci_core.CrossConnect
	DisableNotFoundRetries bool
}

func (s *CrossConnectResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CrossConnectResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.CrossConnectLifecycleStateProvisioning),
	}
}

func (s *CrossConnectResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.CrossConnectLifecycleStatePendingCustomer),
		string(oci_core.CrossConnectLifecycleStateProvisioned),
	}
}

func (s *CrossConnectResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_core.CrossConnectLifecycleStateProvisioning),
	}
}

func (s *CrossConnectResourceCrud) UpdatedTarget() []string {
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

func (s *CrossConnectResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.CrossConnectLifecycleStateTerminating),
	}
}

func (s *CrossConnectResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.CrossConnectLifecycleStateTerminated),
	}
}

func (s *CrossConnectResourceCrud) Create() error {
	request := oci_core.CreateCrossConnectRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if crossConnectGroupId, ok := s.D.GetOkExists("cross_connect_group_id"); ok {
		tmp := crossConnectGroupId.(string)
		request.CrossConnectGroupId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if farCrossConnectOrCrossConnectGroupId, ok := s.D.GetOkExists("far_cross_connect_or_cross_connect_group_id"); ok {
		tmp := farCrossConnectOrCrossConnectGroupId.(string)
		request.FarCrossConnectOrCrossConnectGroupId = &tmp
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

func (s *CrossConnectResourceCrud) Get() error {
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

func (s *CrossConnectResourceCrud) Update() error {
	request := oci_core.UpdateCrossConnectRequest{}

	tmp := s.D.Id()
	request.CrossConnectId = &tmp

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
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

func (s *CrossConnectResourceCrud) Delete() error {
	request := oci_core.DeleteCrossConnectRequest{}

	tmp := s.D.Id()
	request.CrossConnectId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteCrossConnect(context.Background(), request)
	return err
}

func (s *CrossConnectResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CrossConnectGroupId != nil {
		s.D.Set("cross_connect_group_id", *s.Res.CrossConnectGroupId)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

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
