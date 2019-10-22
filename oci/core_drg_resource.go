// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"

	oci_core "github.com/oracle/oci-go-sdk/core"
	oci_work_requests "github.com/oracle/oci-go-sdk/workrequests"
)

func CoreDrgResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createCoreDrg,
		Read:     readCoreDrg,
		Update:   updateCoreDrg,
		Delete:   deleteCoreDrg,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
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
		},
	}
}

func createCoreDrg(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDrgResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return CreateResource(d, sync)
}

func readCoreDrg(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDrgResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return ReadResource(sync)
}

func updateCoreDrg(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDrgResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient
	sync.workRequestClient = m.(*OracleClients).workRequestClient
	return UpdateResource(d, sync)
}

func deleteCoreDrg(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDrgResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type CoreDrgResourceCrud struct {
	BaseCrud
	Client                 *oci_core.VirtualNetworkClient
	workRequestClient      *oci_work_requests.WorkRequestClient
	Res                    *oci_core.Drg
	DisableNotFoundRetries bool
}

func (s *CoreDrgResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CoreDrgResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.DrgLifecycleStateProvisioning),
	}
}

func (s *CoreDrgResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.DrgLifecycleStateAvailable),
	}
}

func (s *CoreDrgResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.DrgLifecycleStateTerminating),
	}
}

func (s *CoreDrgResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.DrgLifecycleStateTerminated),
	}
}

func (s *CoreDrgResourceCrud) Create() error {
	request := oci_core.CreateDrgRequest{}

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

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateDrg(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Drg
	return nil
}

func (s *CoreDrgResourceCrud) Get() error {
	request := oci_core.GetDrgRequest{}

	tmp := s.D.Id()
	request.DrgId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetDrg(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Drg
	return nil
}

func (s *CoreDrgResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_core.UpdateDrgRequest{}

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

	tmp := s.D.Id()
	request.DrgId = &tmp

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateDrg(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Drg
	return nil
}

func (s *CoreDrgResourceCrud) Delete() error {
	request := oci_core.DeleteDrgRequest{}

	tmp := s.D.Id()
	request.DrgId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteDrg(context.Background(), request)
	return err
}

func (s *CoreDrgResourceCrud) SetData() error {
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

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func (s *CoreDrgResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_core.ChangeDrgCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.DrgId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.ChangeDrgCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}
	workId := response.OpcWorkRequestId
	// work request doesn't return identifier once succeeded
	_, err = WaitForWorkRequest(s.workRequestClient, workId, "core", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, false)
	if err != nil {
		return err
	}
	return nil
}
