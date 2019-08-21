// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/terraform/helper/validation"

	"github.com/hashicorp/terraform/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/common"
	oci_streaming "github.com/oracle/oci-go-sdk/streaming"
)

const (
	streamingArchiverRunningState = "running"
	streamingArchiverStoppedState = "stopped"
)

func StreamingStreamArchiverResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createStreamingStreamArchiver,
		Read:     readStreamingStreamArchiver,
		Update:   updateStreamingStreamArchiver,
		Delete:   deleteStreamingStreamArchiver,
		Schema: map[string]*schema.Schema{
			// Required
			"batch_rollover_size_in_mbs": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"batch_rollover_time_in_seconds": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"bucket": {
				Type:     schema.TypeString,
				Required: true,
			},
			"start_position": {
				Type:     schema.TypeString,
				Required: true,
			},
			"stream_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"use_existing_bucket": {
				Type:     schema.TypeBool,
				Required: true,
			},

			// Optional
			"state": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					streamingArchiverRunningState,
					streamingArchiverStoppedState,
				}, true),
			},

			// Computed
			"error": {
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"code": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"message": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createStreamingStreamArchiver(d *schema.ResourceData, m interface{}) error {
	sync := &StreamingStreamArchiverResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).streamAdminClient

	return CreateResource(d, sync)
}

func readStreamingStreamArchiver(d *schema.ResourceData, m interface{}) error {
	sync := &StreamingStreamArchiverResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).streamAdminClient

	return ReadResource(sync)
}

func updateStreamingStreamArchiver(d *schema.ResourceData, m interface{}) error {
	sync := &StreamingStreamArchiverResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).streamAdminClient

	return UpdateResource(d, sync)
}

func deleteStreamingStreamArchiver(d *schema.ResourceData, m interface{}) error {
	sync := &StreamingStreamArchiverResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).streamAdminClient
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type StreamingStreamArchiverResourceCrud struct {
	BaseCrud
	Client                 *oci_streaming.StreamAdminClient
	Res                    *oci_streaming.Archiver
	DisableNotFoundRetries bool
}

func (s *StreamingStreamArchiverResourceCrud) ID() string {
	return getStreamArchiverCompositeId(s.D.Get("stream_id").(string))
}

func (s *StreamingStreamArchiverResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_streaming.ArchiverLifecycleStateCreating),
		string(oci_streaming.ArchiverLifecycleStateStarting),
		string(oci_streaming.ArchiverLifecycleStateUpdating),
	}
}

func (s *StreamingStreamArchiverResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_streaming.ArchiverLifecycleStateRunning),
		string(oci_streaming.ArchiverLifecycleStateStopped),
	}
}

func (s *StreamingStreamArchiverResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_streaming.ArchiverLifecycleStateUpdating),
		string(oci_streaming.ArchiverLifecycleStateStarting),
		string(oci_streaming.ArchiverLifecycleStateStopping),
	}
}

func (s *StreamingStreamArchiverResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_streaming.ArchiverLifecycleStateRunning),
		string(oci_streaming.ArchiverLifecycleStateStopped),
	}
}

func (s *StreamingStreamArchiverResourceCrud) Create() error {
	//As API does not support delete provider checks if there is an existing archiver attached to the stream_id provided.
	//If yes then creation is skipped and update is done assuming change exists
	err := s.Get()
	desiredStateStr := streamingArchiverStoppedState
	request := oci_streaming.CreateArchiverRequest{}

	if batchRolloverSizeInMBs, ok := s.D.GetOkExists("batch_rollover_size_in_mbs"); ok {
		tmp := batchRolloverSizeInMBs.(int)
		request.BatchRolloverSizeInMBs = &tmp
	}

	if batchRolloverTimeInSeconds, ok := s.D.GetOkExists("batch_rollover_time_in_seconds"); ok {
		tmp := batchRolloverTimeInSeconds.(int)
		request.BatchRolloverTimeInSeconds = &tmp
	}

	if bucket, ok := s.D.GetOkExists("bucket"); ok {
		tmp := bucket.(string)
		request.BucketName = &tmp
	}

	if startPosition, ok := s.D.GetOkExists("start_position"); ok {
		request.StartPosition = oci_streaming.CreateArchiverDetailsStartPositionEnum(startPosition.(string))
	}

	if streamId, ok := s.D.GetOkExists("stream_id"); ok {
		tmp := streamId.(string)
		request.StreamId = &tmp
	}

	if useExistingBucket, ok := s.D.GetOkExists("use_existing_bucket"); ok {
		tmp := useExistingBucket.(bool)
		request.UseExistingBucket = &tmp
	}

	if desiredState, ok := s.D.GetOkExists("state"); ok {
		desiredStateStr = desiredState.(string)
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "streaming")

	if err != nil {
		response, err := s.Client.CreateArchiver(context.Background(), request)
		if err != nil {
			return err
		}
		s.Res = &response.Archiver
		err = s.SetData()
		if err != nil {
			return err
		}
		//Default creation of archiver is in STOPPED state, hence explicitly starting is required to be created in RUNNING state
		if strings.ToLower(desiredStateStr) == streamingArchiverRunningState {
			archiver, err := s.setStreamingArchiverDesiredState(request.StreamId, desiredStateStr)
			if err != nil {
				return err
			}
			s.Res = archiver
		}
		return nil
	}
	//If resource was previously destroyed and recreated in RUNNING state, archiver is stopped to perform any update and then started to RUNNING state
	remoteState := string(s.Res.LifecycleState)
	desiredStateStr = strings.ToLower(desiredStateStr)
	if desiredStateStr == streamingArchiverRunningState && strings.ToLower(remoteState) == desiredStateStr {
		_, err := s.setStreamingArchiverDesiredState(request.StreamId, streamingArchiverStoppedState)
		if err != nil {
			return err
		}
	}
	err = s.Update()
	if err != nil {
		return err
	}
	return nil
}

func (s *StreamingStreamArchiverResourceCrud) Get() error {
	request := oci_streaming.GetArchiverRequest{}

	if streamId, ok := s.D.GetOkExists("stream_id"); ok {
		tmp := streamId.(string)
		request.StreamId = &tmp
	}

	streamId, err := parseStreamArchiverCompositeId(s.D.Id())
	if err == nil {
		request.StreamId = &streamId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "streaming")

	response, err := s.Client.GetArchiver(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Archiver
	return nil
}

func (s *StreamingStreamArchiverResourceCrud) Delete() error {
	streamId, err := parseStreamArchiverCompositeId(s.D.Id())
	if err != nil {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	err = s.Get()
	if err != nil {
		return err
	}
	remoteState := string(s.Res.LifecycleState)
	if strings.ToLower(remoteState) != streamingArchiverStoppedState {
		_, err = s.setStreamingArchiverDesiredState(&streamId, streamingArchiverStoppedState)
	}
	return err
}

// Update has 3 API calls. STOP, UPDATE and START
// Stop is before update, because service does not allow updating the resource in RUNNING state.
// Lets consider the following example, resource is in RUNNING state on the cloud and if the new config value {state=stopped},
// it will never happen as there is update before changing the state to running which will fail for all the subsequent plan/apply.
// Hence we need to check if we need to stop before update, start after update
func (s *StreamingStreamArchiverResourceCrud) Update() error {
	request := oci_streaming.UpdateArchiverRequest{}

	if streamId, ok := s.D.GetOkExists("stream_id"); ok {
		tmp := streamId.(string)
		request.StreamId = &tmp
	}

	changeExists := s.D.HasChange("state")
	desiredStateStr := streamingArchiverStoppedState
	if desiredState, ok := s.D.GetOkExists("state"); ok {
		desiredStateStr = desiredState.(string)
	}
	desiredStateStr = strings.ToLower(desiredStateStr)
	err := s.Get()
	if err != nil {
		return err
	}
	remoteState := string(s.Res.LifecycleState)
	if changeExists && desiredStateStr == streamingArchiverStoppedState && strings.ToLower(remoteState) != desiredStateStr {
		_, err := s.setStreamingArchiverDesiredState(request.StreamId, desiredStateStr)
		if err != nil {
			return err
		}
	}

	if batchRolloverSizeInMBs, ok := s.D.GetOkExists("batch_rollover_size_in_mbs"); ok {
		tmp := batchRolloverSizeInMBs.(int)
		request.BatchRolloverSizeInMBs = &tmp
	}

	if batchRolloverTimeInSeconds, ok := s.D.GetOkExists("batch_rollover_time_in_seconds"); ok {
		tmp := batchRolloverTimeInSeconds.(int)
		request.BatchRolloverTimeInSeconds = &tmp
	}

	if bucket, ok := s.D.GetOkExists("bucket"); ok {
		tmp := bucket.(string)
		request.BucketName = &tmp
	}

	if startPosition, ok := s.D.GetOkExists("start_position"); ok {
		request.StartPosition = oci_streaming.UpdateArchiverDetailsStartPositionEnum(startPosition.(string))
	}

	if useExistingBucket, ok := s.D.GetOkExists("use_existing_bucket"); ok {
		tmp := useExistingBucket.(bool)
		request.UseExistingBucket = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "streaming")

	response, err := s.Client.UpdateArchiver(context.Background(), request)
	if err != nil {
		return err
	}

	if changeExists && desiredStateStr == streamingArchiverRunningState && strings.ToLower(remoteState) != desiredStateStr {
		archiver, err := s.setStreamingArchiverDesiredState(request.StreamId, desiredStateStr)
		if err != nil {
			return err
		}
		s.Res = archiver
		return nil
	}

	s.Res = &response.Archiver
	return nil
}

func (s *StreamingStreamArchiverResourceCrud) SetData() error {
	streamId, err := parseStreamArchiverCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("stream_id", &streamId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.BatchRolloverSizeInMBs != nil {
		s.D.Set("batch_rollover_size_in_mbs", *s.Res.BatchRolloverSizeInMBs)
	}

	if s.Res.BatchRolloverTimeInSeconds != nil {
		s.D.Set("batch_rollover_time_in_seconds", *s.Res.BatchRolloverTimeInSeconds)
	}

	if s.Res.BucketName != nil {
		s.D.Set("bucket", *s.Res.BucketName)
	}

	if s.Res.Error != nil {
		s.D.Set("error", []interface{}{ArchiverErrorToMap(s.Res.Error)})
	} else {
		s.D.Set("error", nil)
	}

	s.D.Set("start_position", s.Res.StartPosition)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.UseExistingBucket != nil {
		s.D.Set("use_existing_bucket", *s.Res.UseExistingBucket)
	}

	return nil
}

func getStreamArchiverCompositeId(streamId string) string {
	streamId = url.PathEscape(streamId)
	compositeId := "archiver/" + streamId
	return compositeId
}

func parseStreamArchiverCompositeId(compositeId string) (streamId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("archiver/.*", compositeId)
	if !match || len(parts) != 2 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	streamId, _ = url.PathUnescape(parts[1])

	return
}

func ArchiverErrorToMap(obj *oci_streaming.ArchiverError) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Code != nil {
		result["code"] = string(*obj.Code)
	}

	if obj.Message != nil {
		result["message"] = string(*obj.Message)
	}

	return result
}

func (s *StreamingStreamArchiverResourceCrud) setStreamingArchiverDesiredState(streamId *string, desiredState string) (*oci_streaming.Archiver, error) {
	switch strings.ToLower(desiredState) {
	case streamingArchiverRunningState:
		startRequest := oci_streaming.StartArchiverRequest{}
		startRequest.StreamId = streamId
		startRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "streaming")

		startResponse, err := s.Client.StartArchiver(context.Background(), startRequest)
		if err != nil {
			return &startResponse.Archiver, err
		}

		// Wait for archiver to not be in RUNNING state after the start.
		getArchiverResponse, err := waitForStreamArchiverUntilItIsInDesiredState(streamId, s.Client, s.D.Timeout(schema.TimeoutUpdate), desiredState)

		return &getArchiverResponse.Archiver, err
	case streamingArchiverStoppedState:
		stopRequest := oci_streaming.StopArchiverRequest{}
		stopRequest.StreamId = streamId
		stopRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "streaming")

		stopResponse, err := s.Client.StopArchiver(context.Background(), stopRequest)
		if err != nil {
			return &stopResponse.Archiver, err
		}

		// Wait for archiver to not be in STOPPED state after the stop.
		getArchiverResponse, err := waitForStreamArchiverUntilItIsInDesiredState(streamId, s.Client, s.D.Timeout(schema.TimeoutUpdate), desiredState)

		return &getArchiverResponse.Archiver, err
	default:
		return nil, fmt.Errorf("received unknown 'state' %s", desiredState)
	}

}

func waitForStreamArchiverUntilItIsInDesiredState(streamId *string, client *oci_streaming.StreamAdminClient, timeout time.Duration, desiredState string) (*oci_streaming.GetArchiverResponse, error) {
	getArchiverRequest := oci_streaming.GetArchiverRequest{}

	getArchiverRequest.StreamId = streamId

	archiverUpdating := func(response oci_common.OCIOperationResponse) bool {
		if getArchiverResponse, ok := response.Response.(oci_streaming.GetArchiverResponse); ok {
			switch strings.ToLower(desiredState) {
			case streamingArchiverRunningState:
				if getArchiverResponse.LifecycleState != oci_streaming.ArchiverLifecycleStateRunning {
					return true
				}
			case streamingArchiverStoppedState:
				if getArchiverResponse.LifecycleState != oci_streaming.ArchiverLifecycleStateStopped {
					return true
				}
			default:
				return false
			}
		}
		return false
	}

	getArchiverRequest.RequestMetadata.RetryPolicy = getRetryPolicyWithAdditionalRetryCondition(timeout, archiverUpdating, "streaming")
	getArchiverResponse, err := client.GetArchiver(context.Background(), getArchiverRequest)
	return &getArchiverResponse, err
}
