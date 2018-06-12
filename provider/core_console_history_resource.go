// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"

	oci_core "github.com/oracle/oci-go-sdk/core"
)

func ConsoleHistoryResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createConsoleHistory,
		Read:     readConsoleHistory,
		Update:   updateConsoleHistory,
		Delete:   deleteConsoleHistory,
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
			"availability_domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"id": {
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

func createConsoleHistory(d *schema.ResourceData, m interface{}) error {
	sync := &ConsoleHistoryResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient

	return crud.CreateResource(d, sync)
}

func readConsoleHistory(d *schema.ResourceData, m interface{}) error {
	sync := &ConsoleHistoryResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient

	return crud.ReadResource(sync)
}

func updateConsoleHistory(d *schema.ResourceData, m interface{}) error {
	sync := &ConsoleHistoryResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient

	return crud.UpdateResource(d, sync)
}

func deleteConsoleHistory(d *schema.ResourceData, m interface{}) error {
	sync := &ConsoleHistoryResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient
	sync.DisableNotFoundRetries = true

	return crud.DeleteResource(d, sync)
}

type ConsoleHistoryResourceCrud struct {
	crud.BaseCrud
	Client                 *oci_core.ComputeClient
	Res                    *oci_core.ConsoleHistory
	DisableNotFoundRetries bool
}

func (s *ConsoleHistoryResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *ConsoleHistoryResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.ConsoleHistoryLifecycleStateRequested),
		string(oci_core.ConsoleHistoryLifecycleStateGettingHistory),
	}
}

func (s *ConsoleHistoryResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.ConsoleHistoryLifecycleStateSucceeded),
	}
}

func (s *ConsoleHistoryResourceCrud) Create() error {
	request := oci_core.CaptureConsoleHistoryRequest{}

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

	if instanceId, ok := s.D.GetOkExists("instance_id"); ok {
		tmp := instanceId.(string)
		request.InstanceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CaptureConsoleHistory(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ConsoleHistory
	return nil
}

func (s *ConsoleHistoryResourceCrud) Get() error {
	request := oci_core.GetConsoleHistoryRequest{}

	tmp := s.D.Id()
	request.InstanceConsoleHistoryId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetConsoleHistory(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ConsoleHistory
	return nil
}

func (s *ConsoleHistoryResourceCrud) Update() error {
	request := oci_core.UpdateConsoleHistoryRequest{}

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

	tmp := s.D.Id()
	request.InstanceConsoleHistoryId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateConsoleHistory(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ConsoleHistory
	return nil
}

func (s *ConsoleHistoryResourceCrud) Delete() error {
	// Do not delete console history.
	return nil
}

func (s *ConsoleHistoryResourceCrud) SetData() {
	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

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

	if s.Res.Id != nil {
		s.D.Set("id", *s.Res.Id)
	}

	if s.Res.InstanceId != nil {
		s.D.Set("instance_id", *s.Res.InstanceId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

}
