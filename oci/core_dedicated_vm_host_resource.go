// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_core "github.com/oracle/oci-go-sdk/v31/core"
	oci_work_requests "github.com/oracle/oci-go-sdk/v31/workrequests"
)

func init() {
	RegisterResource("oci_core_dedicated_vm_host", CoreDedicatedVmHostResource())
}

func CoreDedicatedVmHostResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createCoreDedicatedVmHost,
		Read:     readCoreDedicatedVmHost,
		Update:   updateCoreDedicatedVmHost,
		Delete:   deleteCoreDedicatedVmHost,
		Schema: map[string]*schema.Schema{
			// Required
			"availability_domain": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"dedicated_vm_host_shape": {
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
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fault_domain": {
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

			// Computed
			"remaining_ocpus": {
				Type:     schema.TypeFloat,
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
			"total_ocpus": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
		},
	}
}

func createCoreDedicatedVmHost(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDedicatedVmHostResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient()
	sync.workRequestClient = m.(*OracleClients).workRequestClient

	return CreateResource(d, sync)
}

func readCoreDedicatedVmHost(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDedicatedVmHostResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient()

	return ReadResource(sync)
}

func updateCoreDedicatedVmHost(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDedicatedVmHostResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient()
	sync.workRequestClient = m.(*OracleClients).workRequestClient

	return UpdateResource(d, sync)
}

func deleteCoreDedicatedVmHost(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDedicatedVmHostResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient()
	sync.workRequestClient = m.(*OracleClients).workRequestClient
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type CoreDedicatedVmHostResourceCrud struct {
	BaseCrud
	Client                 *oci_core.ComputeClient
	workRequestClient      *oci_work_requests.WorkRequestClient
	Res                    *oci_core.DedicatedVmHost
	DisableNotFoundRetries bool
}

func (s *CoreDedicatedVmHostResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CoreDedicatedVmHostResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.DedicatedVmHostLifecycleStateCreating),
	}
}

func (s *CoreDedicatedVmHostResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.DedicatedVmHostLifecycleStateActive),
	}
}

func (s *CoreDedicatedVmHostResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.DedicatedVmHostLifecycleStateDeleting),
	}
}

func (s *CoreDedicatedVmHostResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.DedicatedVmHostLifecycleStateDeleted),
	}
}

func (s *CoreDedicatedVmHostResourceCrud) Create() error {
	request := oci_core.CreateDedicatedVmHostRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if dedicatedVmHostShape, ok := s.D.GetOkExists("dedicated_vm_host_shape"); ok {
		tmp := dedicatedVmHostShape.(string)
		request.DedicatedVmHostShape = &tmp
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

	if faultDomain, ok := s.D.GetOkExists("fault_domain"); ok {
		tmp := faultDomain.(string)
		request.FaultDomain = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateDedicatedVmHost(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		identifier, err := WaitForWorkRequestWithErrorHandling(s.workRequestClient, workId, "dedicatedvmhost", oci_work_requests.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if identifier != nil {
			s.D.SetId(*identifier)
		}
		if err != nil {
			return err
		}
	}

	return s.Get()
}

func (s *CoreDedicatedVmHostResourceCrud) Get() error {
	request := oci_core.GetDedicatedVmHostRequest{}

	tmp := s.D.Id()
	request.DedicatedVmHostId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetDedicatedVmHost(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DedicatedVmHost
	return nil
}

func (s *CoreDedicatedVmHostResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_core.UpdateDedicatedVmHostRequest{}

	tmp := s.D.Id()
	request.DedicatedVmHostId = &tmp

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

	response, err := s.Client.UpdateDedicatedVmHost(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DedicatedVmHost
	return nil
}

func (s *CoreDedicatedVmHostResourceCrud) Delete() error {
	request := oci_core.DeleteDedicatedVmHostRequest{}

	tmp := s.D.Id()
	request.DedicatedVmHostId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.DeleteDedicatedVmHost(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = WaitForWorkRequestWithErrorHandling(s.workRequestClient, workId, "dedicatedvmhost", oci_work_requests.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *CoreDedicatedVmHostResourceCrud) SetData() error {
	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DedicatedVmHostShape != nil {
		s.D.Set("dedicated_vm_host_shape", *s.Res.DedicatedVmHostShape)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.FaultDomain != nil {
		s.D.Set("fault_domain", *s.Res.FaultDomain)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.RemainingOcpus != nil {
		s.D.Set("remaining_ocpus", *s.Res.RemainingOcpus)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TotalOcpus != nil {
		s.D.Set("total_ocpus", *s.Res.TotalOcpus)
	}

	return nil
}

func (s *CoreDedicatedVmHostResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_core.ChangeDedicatedVmHostCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.DedicatedVmHostId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.ChangeDedicatedVmHostCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}
	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = WaitForWorkRequestWithErrorHandling(s.workRequestClient, workId, "dedicatedvmhost", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}
	return nil
}
