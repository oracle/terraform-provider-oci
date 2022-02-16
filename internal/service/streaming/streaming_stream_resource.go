// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package streaming

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_streaming "github.com/oracle/oci-go-sdk/v58/streaming"
)

func StreamingStreamResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createStreamingStream,
		Read:     readStreamingStream,
		Update:   updateStreamingStream,
		Delete:   deleteStreamingStream,
		Schema: map[string]*schema.Schema{
			// Required
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"partitions": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"compartment_id": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"stream_pool_id"},
			},
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
			"retention_in_hours": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"stream_pool_id": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"compartment_id"},
			},

			// Computed
			"lifecycle_state_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"messages_endpoint": {
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

func createStreamingStream(d *schema.ResourceData, m interface{}) error {
	sync := &StreamingStreamResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StreamAdminClient()

	return tfresource.CreateResource(d, sync)
}

func readStreamingStream(d *schema.ResourceData, m interface{}) error {
	sync := &StreamingStreamResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StreamAdminClient()

	return tfresource.ReadResource(sync)
}

func updateStreamingStream(d *schema.ResourceData, m interface{}) error {
	sync := &StreamingStreamResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StreamAdminClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteStreamingStream(d *schema.ResourceData, m interface{}) error {
	sync := &StreamingStreamResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StreamAdminClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type StreamingStreamResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_streaming.StreamAdminClient
	Res                    *oci_streaming.Stream
	DisableNotFoundRetries bool
}

func (s *StreamingStreamResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *StreamingStreamResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_streaming.StreamLifecycleStateCreating),
	}
}

func (s *StreamingStreamResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_streaming.StreamLifecycleStateActive),
	}
}

func (s *StreamingStreamResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_streaming.StreamLifecycleStateDeleting),
	}
}

func (s *StreamingStreamResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_streaming.StreamLifecycleStateDeleted),
	}
}

func (s *StreamingStreamResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_streaming.StreamLifecycleStateUpdating),
	}
}

func (s *StreamingStreamResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_streaming.StreamLifecycleStateActive),
	}
}

func (s *StreamingStreamResourceCrud) Create() error {
	request := oci_streaming.CreateStreamRequest{}

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

	if partitions, ok := s.D.GetOkExists("partitions"); ok {
		tmp := partitions.(int)
		request.Partitions = &tmp
	}

	if retentionInHours, ok := s.D.GetOkExists("retention_in_hours"); ok {
		tmp := retentionInHours.(int)
		request.RetentionInHours = &tmp
	}

	if streamPoolId, ok := s.D.GetOkExists("stream_pool_id"); ok {
		tmp := streamPoolId.(string)
		request.StreamPoolId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "streaming")

	response, err := s.Client.CreateStream(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Stream
	return nil
}

func (s *StreamingStreamResourceCrud) Get() error {
	request := oci_streaming.GetStreamRequest{}

	tmp := s.D.Id()
	request.StreamId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "streaming")

	response, err := s.Client.GetStream(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Stream
	return nil
}

func (s *StreamingStreamResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_streaming.UpdateStreamRequest{}

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

	tmp := s.D.Id()
	request.StreamId = &tmp

	if streamPoolId, ok := s.D.GetOkExists("stream_pool_id"); ok && s.D.HasChange("stream_pool_id") {
		tmp := streamPoolId.(string)
		request.StreamPoolId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "streaming")

	response, err := s.Client.UpdateStream(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Stream
	return nil
}

func (s *StreamingStreamResourceCrud) Delete() error {
	request := oci_streaming.DeleteStreamRequest{}

	tmp := s.D.Id()
	request.StreamId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "streaming")

	_, err := s.Client.DeleteStream(context.Background(), request)
	return err
}

func (s *StreamingStreamResourceCrud) SetData() error {
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

	if s.Res.MessagesEndpoint != nil {
		s.D.Set("messages_endpoint", *s.Res.MessagesEndpoint)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.Partitions != nil {
		s.D.Set("partitions", *s.Res.Partitions)
	}

	if s.Res.RetentionInHours != nil {
		s.D.Set("retention_in_hours", *s.Res.RetentionInHours)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.StreamPoolId != nil {
		s.D.Set("stream_pool_id", *s.Res.StreamPoolId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func (s *StreamingStreamResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_streaming.ChangeStreamCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.StreamId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "streaming")

	_, err := s.Client.ChangeStreamCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
