// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package streaming

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_streaming "github.com/oracle/oci-go-sdk/v56/streaming"
)

func StreamingConnectHarnessResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createStreamingConnectHarness,
		Read:     readStreamingConnectHarness,
		Update:   updateStreamingConnectHarness,
		Delete:   deleteStreamingConnectHarness,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
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
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"lifecycle_state_details": {
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

func createStreamingConnectHarness(d *schema.ResourceData, m interface{}) error {
	sync := &StreamingConnectHarnessResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StreamAdminClient()

	return tfresource.CreateResource(d, sync)
}

func readStreamingConnectHarness(d *schema.ResourceData, m interface{}) error {
	sync := &StreamingConnectHarnessResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StreamAdminClient()

	return tfresource.ReadResource(sync)
}

func updateStreamingConnectHarness(d *schema.ResourceData, m interface{}) error {
	sync := &StreamingConnectHarnessResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StreamAdminClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteStreamingConnectHarness(d *schema.ResourceData, m interface{}) error {
	sync := &StreamingConnectHarnessResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StreamAdminClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type StreamingConnectHarnessResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_streaming.StreamAdminClient
	Res                    *oci_streaming.ConnectHarness
	DisableNotFoundRetries bool
}

func (s *StreamingConnectHarnessResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *StreamingConnectHarnessResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_streaming.ConnectHarnessLifecycleStateCreating),
	}
}

func (s *StreamingConnectHarnessResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_streaming.ConnectHarnessLifecycleStateActive),
	}
}

func (s *StreamingConnectHarnessResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_streaming.ConnectHarnessLifecycleStateDeleting),
	}
}

func (s *StreamingConnectHarnessResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_streaming.ConnectHarnessLifecycleStateDeleted),
	}
}

func (s *StreamingConnectHarnessResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_streaming.ConnectHarnessLifecycleStateUpdating),
	}
}

func (s *StreamingConnectHarnessResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_streaming.ConnectHarnessLifecycleStateActive),
	}
}

func (s *StreamingConnectHarnessResourceCrud) Create() error {
	request := oci_streaming.CreateConnectHarnessRequest{}

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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "streaming")

	response, err := s.Client.CreateConnectHarness(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ConnectHarness
	return nil
}

func (s *StreamingConnectHarnessResourceCrud) Get() error {
	request := oci_streaming.GetConnectHarnessRequest{}

	tmp := s.D.Id()
	request.ConnectHarnessId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "streaming")

	response, err := s.Client.GetConnectHarness(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ConnectHarness
	return nil
}

func (s *StreamingConnectHarnessResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_streaming.UpdateConnectHarnessRequest{}

	tmp := s.D.Id()
	request.ConnectHarnessId = &tmp

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "streaming")

	response, err := s.Client.UpdateConnectHarness(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ConnectHarness
	return nil
}

func (s *StreamingConnectHarnessResourceCrud) Delete() error {
	request := oci_streaming.DeleteConnectHarnessRequest{}

	tmp := s.D.Id()
	request.ConnectHarnessId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "streaming")

	_, err := s.Client.DeleteConnectHarness(context.Background(), request)
	return err
}

func (s *StreamingConnectHarnessResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleStateDetails != nil {
		s.D.Set("lifecycle_state_details", *s.Res.LifecycleStateDetails)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func (s *StreamingConnectHarnessResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_streaming.ChangeConnectHarnessCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.ConnectHarnessId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "streaming")

	_, err := s.Client.ChangeConnectHarnessCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
