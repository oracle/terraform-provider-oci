// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_core "github.com/oracle/oci-go-sdk/v65/core"
)

func CoreConsoleHistoryResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCoreConsoleHistory,
		Read:     readCoreConsoleHistory,
		Update:   updateCoreConsoleHistory,
		Delete:   deleteCoreConsoleHistory,
		Schema: map[string]*schema.Schema{
			// Required
			"instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
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
			"availability_domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": {
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

func createCoreConsoleHistory(d *schema.ResourceData, m interface{}) error {
	sync := &CoreConsoleHistoryResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.CreateResource(d, sync)
}

func readCoreConsoleHistory(d *schema.ResourceData, m interface{}) error {
	sync := &CoreConsoleHistoryResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.ReadResource(sync)
}

func updateCoreConsoleHistory(d *schema.ResourceData, m interface{}) error {
	sync := &CoreConsoleHistoryResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteCoreConsoleHistory(d *schema.ResourceData, m interface{}) error {
	sync := &CoreConsoleHistoryResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type CoreConsoleHistoryResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_core.ComputeClient
	Res                    *oci_core.ConsoleHistory
	DisableNotFoundRetries bool
}

func (s *CoreConsoleHistoryResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CoreConsoleHistoryResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.ConsoleHistoryLifecycleStateRequested),
		string(oci_core.ConsoleHistoryLifecycleStateGettingHistory),
	}
}

func (s *CoreConsoleHistoryResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.ConsoleHistoryLifecycleStateSucceeded),
	}
}

func (s *CoreConsoleHistoryResourceCrud) Create() error {
	request := oci_core.CaptureConsoleHistoryRequest{}

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
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if instanceId, ok := s.D.GetOkExists("instance_id"); ok {
		tmp := instanceId.(string)
		request.InstanceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CaptureConsoleHistory(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ConsoleHistory
	return nil
}

func (s *CoreConsoleHistoryResourceCrud) Get() error {
	request := oci_core.GetConsoleHistoryRequest{}

	tmp := s.D.Id()
	request.InstanceConsoleHistoryId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetConsoleHistory(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ConsoleHistory
	return nil
}

func (s *CoreConsoleHistoryResourceCrud) Update() error {
	request := oci_core.UpdateConsoleHistoryRequest{}

	tmp := s.D.Id()
	request.InstanceConsoleHistoryId = &tmp

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
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateConsoleHistory(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ConsoleHistory
	return nil
}

func (s *CoreConsoleHistoryResourceCrud) Delete() error {
	request := oci_core.DeleteConsoleHistoryRequest{}

	tmp := s.D.Id()
	request.InstanceConsoleHistoryId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteConsoleHistory(context.Background(), request)
	return err
}

func (s *CoreConsoleHistoryResourceCrud) SetData() error {
	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

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

	if s.Res.InstanceId != nil {
		s.D.Set("instance_id", *s.Res.InstanceId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
