// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"

	oci_streaming "github.com/oracle/oci-go-sdk/streaming"
)

func init() {
	RegisterResource("oci_streaming_stream_pool", StreamingStreamPoolResource())
}

func StreamingStreamPoolResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createStreamingStreamPool,
		Read:     readStreamingStreamPool,
		Update:   updateStreamingStreamPool,
		Delete:   deleteStreamingStreamPool,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
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
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"kafka_settings": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"auto_create_topics_enable": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"bootstrap_servers": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"log_retention_hours": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"num_partitions": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
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

func createStreamingStreamPool(d *schema.ResourceData, m interface{}) error {
	sync := &StreamingStreamPoolResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).streamAdminClient

	return CreateResource(d, sync)
}

func readStreamingStreamPool(d *schema.ResourceData, m interface{}) error {
	sync := &StreamingStreamPoolResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).streamAdminClient

	return ReadResource(sync)
}

func updateStreamingStreamPool(d *schema.ResourceData, m interface{}) error {
	sync := &StreamingStreamPoolResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).streamAdminClient

	return UpdateResource(d, sync)
}

func deleteStreamingStreamPool(d *schema.ResourceData, m interface{}) error {
	sync := &StreamingStreamPoolResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).streamAdminClient
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type StreamingStreamPoolResourceCrud struct {
	BaseCrud
	Client                 *oci_streaming.StreamAdminClient
	Res                    *oci_streaming.StreamPool
	DisableNotFoundRetries bool
}

func (s *StreamingStreamPoolResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *StreamingStreamPoolResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_streaming.StreamPoolLifecycleStateCreating),
	}
}

func (s *StreamingStreamPoolResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_streaming.StreamPoolLifecycleStateActive),
	}
}

func (s *StreamingStreamPoolResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_streaming.StreamPoolLifecycleStateDeleting),
	}
}

func (s *StreamingStreamPoolResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_streaming.StreamPoolLifecycleStateDeleted),
	}
}

func (s *StreamingStreamPoolResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_streaming.StreamPoolLifecycleStateUpdating),
	}
}

func (s *StreamingStreamPoolResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_streaming.StreamPoolLifecycleStateActive),
	}
}

func (s *StreamingStreamPoolResourceCrud) Create() error {
	request := oci_streaming.CreateStreamPoolRequest{}

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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if kafkaSettings, ok := s.D.GetOkExists("kafka_settings"); ok {
		if tmpList := kafkaSettings.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "kafka_settings", 0)
			tmp, err := s.mapToKafkaSettings(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.KafkaSettings = &tmp
		}
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "streaming")

	response, err := s.Client.CreateStreamPool(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.StreamPool
	return nil
}

func (s *StreamingStreamPoolResourceCrud) Get() error {
	request := oci_streaming.GetStreamPoolRequest{}

	tmp := s.D.Id()
	request.StreamPoolId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "streaming")

	response, err := s.Client.GetStreamPool(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.StreamPool
	return nil
}

func (s *StreamingStreamPoolResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_streaming.UpdateStreamPoolRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if kafkaSettings, ok := s.D.GetOkExists("kafka_settings"); ok {
		if tmpList := kafkaSettings.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "kafka_settings", 0)
			tmp, err := s.mapToKafkaSettings(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.KafkaSettings = &tmp
		}
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	tmp := s.D.Id()
	request.StreamPoolId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "streaming")

	response, err := s.Client.UpdateStreamPool(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.StreamPool
	return nil
}

func (s *StreamingStreamPoolResourceCrud) Delete() error {
	request := oci_streaming.DeleteStreamPoolRequest{}

	tmp := s.D.Id()
	request.StreamPoolId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "streaming")

	_, err := s.Client.DeleteStreamPool(context.Background(), request)
	return err
}

func (s *StreamingStreamPoolResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.KafkaSettings != nil {
		s.D.Set("kafka_settings", []interface{}{KafkaSettingsToMap(s.Res.KafkaSettings)})
	} else {
		s.D.Set("kafka_settings", nil)
	}

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

func (s *StreamingStreamPoolResourceCrud) mapToKafkaSettings(fieldKeyFormat string) (oci_streaming.KafkaSettings, error) {
	result := oci_streaming.KafkaSettings{}

	if autoCreateTopicsEnable, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "auto_create_topics_enable")); ok {
		tmp := autoCreateTopicsEnable.(bool)
		result.AutoCreateTopicsEnable = &tmp
	}

	if logRetentionHours, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "log_retention_hours")); ok {
		tmp := logRetentionHours.(int)
		result.LogRetentionHours = &tmp
	}

	if numPartitions, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "num_partitions")); ok {
		tmp := numPartitions.(int)
		result.NumPartitions = &tmp
	}

	return result, nil
}

func KafkaSettingsToMap(obj *oci_streaming.KafkaSettings) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AutoCreateTopicsEnable != nil {
		result["auto_create_topics_enable"] = bool(*obj.AutoCreateTopicsEnable)
	}

	if obj.BootstrapServers != nil {
		result["bootstrap_servers"] = string(*obj.BootstrapServers)
	}

	if obj.LogRetentionHours != nil {
		result["log_retention_hours"] = int(*obj.LogRetentionHours)
	}

	if obj.NumPartitions != nil {
		result["num_partitions"] = int(*obj.NumPartitions)
	}

	return result
}

func (s *StreamingStreamPoolResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_streaming.ChangeStreamPoolCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.StreamPoolId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "streaming")

	_, err := s.Client.ChangeStreamPoolCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}
	retentionPolicyFunc := func() bool { return s.Res.LifecycleState == oci_streaming.StreamPoolLifecycleStateActive }
	return WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}
