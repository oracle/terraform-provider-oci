// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform/helper/validation"

	"github.com/hashicorp/terraform/helper/schema"

	oci_oda "github.com/oracle/oci-go-sdk/v25/oda"
)

func init() {
	RegisterResource("oci_oda_oda_instance", OdaOdaInstanceResource())
}

func OdaOdaInstanceResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createOdaOdaInstance,
		Read:     readOdaOdaInstance,
		Update:   updateOdaOdaInstance,
		Delete:   deleteOdaOdaInstance,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"shape_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: definedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"connector_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_sub_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					string(oci_oda.OdaInstanceLifecycleStateActive),
					string(oci_oda.OdaInstanceLifecycleStateInactive),
				}, true),
			},
			"state_message": {
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
			"web_app_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createOdaOdaInstance(d *schema.ResourceData, m interface{}) error {
	sync := &OdaOdaInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).odaClient()

	var isInactiveRequest = false
	if configState, ok := sync.D.GetOkExists("state"); ok {
		wantedState := oci_oda.OdaInstanceLifecycleStateEnum(strings.ToUpper(configState.(string)))
		if wantedState == oci_oda.OdaInstanceLifecycleStateInactive {
			isInactiveRequest = true
		}
	}

	if error := CreateResource(d, sync); error != nil {
		return error
	}

	if isInactiveRequest {
		return inactiveOdaIfNeeded(d, sync)
	}

	return nil
}

func inactiveOdaIfNeeded(d *schema.ResourceData, sync *OdaOdaInstanceResourceCrud) error {
	if err := sync.StopOdaInstance(); err != nil {
		return err
	}
	return ReadResource(sync)
}

func readOdaOdaInstance(d *schema.ResourceData, m interface{}) error {
	sync := &OdaOdaInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).odaClient()

	return ReadResource(sync)
}

func updateOdaOdaInstance(d *schema.ResourceData, m interface{}) error {
	sync := &OdaOdaInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).odaClient()

	// Start/Stop ODA instance
	stateActive, stateInactive := false, false

	if sync.D.HasChange("state") {
		wantedState := strings.ToUpper(sync.D.Get("state").(string))
		if oci_oda.OdaInstanceLifecycleStateActive == oci_oda.OdaInstanceLifecycleStateEnum(wantedState) {
			stateActive = true
			stateInactive = false
		} else if oci_oda.OdaInstanceLifecycleStateInactive == oci_oda.OdaInstanceLifecycleStateEnum(wantedState) {
			stateInactive = true
			stateActive = false
		} else {
			return fmt.Errorf("[ERROR] Invalid state input for update %v", wantedState)
		}
	}

	if stateActive {
		if err := sync.StartOdaInstance(); err != nil {
			return err
		}
		if err := sync.D.Set("state", oci_oda.OdaInstanceLifecycleStateActive); err != nil {
			return err
		}
	}

	// when state is inactive, it is invalid to update resource
	if err := UpdateResource(d, sync); err != nil {
		return err
	}

	if stateInactive {
		if err := sync.StopOdaInstance(); err != nil {
			return err
		}
		if err := sync.D.Set("state", oci_oda.OdaInstanceLifecycleStateInactive); err != nil {
			return err
		}
	}

	return nil
}

func deleteOdaOdaInstance(d *schema.ResourceData, m interface{}) error {
	sync := &OdaOdaInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).odaClient()
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type OdaOdaInstanceResourceCrud struct {
	BaseCrud
	Client                 *oci_oda.OdaClient
	Res                    *oci_oda.OdaInstance
	DisableNotFoundRetries bool
}

func (s *OdaOdaInstanceResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *OdaOdaInstanceResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_oda.OdaInstanceLifecycleStateCreating),
	}
}

func (s *OdaOdaInstanceResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_oda.OdaInstanceLifecycleStateActive),
	}
}

func (s *OdaOdaInstanceResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_oda.OdaInstanceLifecycleStateDeleting),
	}
}

func (s *OdaOdaInstanceResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_oda.OdaInstanceLifecycleStateDeleted),
	}
}

func (s *OdaOdaInstanceResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_oda.OdaInstanceLifecycleStateUpdating),
	}
}

func (s *OdaOdaInstanceResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_oda.OdaInstanceLifecycleStateActive),
	}
}

func (s *OdaOdaInstanceResourceCrud) Create() error {
	request := oci_oda.CreateOdaInstanceRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if shapeName, ok := s.D.GetOkExists("shape_name"); ok {
		request.ShapeName = oci_oda.CreateOdaInstanceDetailsShapeNameEnum(shapeName.(string))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "oda")

	response, err := s.Client.CreateOdaInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.OdaInstance
	return nil
}

func (s *OdaOdaInstanceResourceCrud) Get() error {
	request := oci_oda.GetOdaInstanceRequest{}

	tmp := s.D.Id()
	request.OdaInstanceId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "oda")

	response, err := s.Client.GetOdaInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.OdaInstance
	return nil
}

func (s *OdaOdaInstanceResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_oda.UpdateOdaInstanceRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.OdaInstanceId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "oda")

	response, err := s.Client.UpdateOdaInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.OdaInstance
	return nil
}

func (s *OdaOdaInstanceResourceCrud) Delete() error {
	request := oci_oda.DeleteOdaInstanceRequest{}

	tmp := s.D.Id()
	request.OdaInstanceId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "oda")

	_, err := s.Client.DeleteOdaInstance(context.Background(), request)
	return err
}

func (s *OdaOdaInstanceResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ConnectorUrl != nil {
		s.D.Set("connector_url", *s.Res.ConnectorUrl)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("lifecycle_sub_state", s.Res.LifecycleSubState)

	s.D.Set("shape_name", s.Res.ShapeName)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.StateMessage != nil {
		s.D.Set("state_message", *s.Res.StateMessage)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.WebAppUrl != nil {
		s.D.Set("web_app_url", *s.Res.WebAppUrl)
	}

	return nil
}

func (s *OdaOdaInstanceResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_oda.ChangeOdaInstanceCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.OdaInstanceId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "oda")

	_, err := s.Client.ChangeOdaInstanceCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}
	return nil
}

func (s *OdaOdaInstanceResourceCrud) StartOdaInstance() error {
	state := oci_oda.OdaInstanceLifecycleStateActive
	if err := s.Get(); err != nil {
		return err
	}
	if s.Res.LifecycleState == state {
		fmt.Printf("[WARN] The ODA instance already in the wanted state: %v", state)
		return nil
	}
	request := oci_oda.StartOdaInstanceRequest{}

	tmp := s.D.Id()
	request.OdaInstanceId = &tmp

	if _, err := s.Client.StartOdaInstance(context.Background(), request); err != nil {
		return err
	}
	retentionPolicyFunc := func() bool { return s.Res.LifecycleState == state }

	return WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *OdaOdaInstanceResourceCrud) StopOdaInstance() error {
	state := oci_oda.OdaInstanceLifecycleStateInactive
	if err := s.Get(); err != nil {
		return err
	}
	if s.Res.LifecycleState == state {
		fmt.Printf("[WARN] The ODA instance already in the wanted state: %v", state)
		return nil
	}
	request := oci_oda.StopOdaInstanceRequest{}

	tmp := s.D.Id()
	request.OdaInstanceId = &tmp

	if _, err := s.Client.StopOdaInstance(context.Background(), request); err != nil {
		return err
	}
	retentionPolicyFunc := func() bool { return s.Res.LifecycleState == state }

	return WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}
