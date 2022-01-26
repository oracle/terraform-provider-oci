// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package streaming

import (
	"context"
	"fmt"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_streaming "github.com/oracle/oci-go-sdk/v56/streaming"
)

func StreamingStreamPoolResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
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
			"custom_encryption_key": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"kms_key_id": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
						"key_state": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
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
			"private_endpoint_settings": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"nsg_ids": {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Set:      utils.LiteralTypeHashCodeForSets,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"private_endpoint_ip": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"subnet_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},

			// Computed
			"endpoint_fqdn": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_private": {
				Type:     schema.TypeBool,
				Computed: true,
			},
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
	sync.Client = m.(*client.OracleClients).StreamAdminClient()

	return tfresource.CreateResource(d, sync)
}

func readStreamingStreamPool(d *schema.ResourceData, m interface{}) error {
	sync := &StreamingStreamPoolResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StreamAdminClient()

	return tfresource.ReadResource(sync)
}

func updateStreamingStreamPool(d *schema.ResourceData, m interface{}) error {
	sync := &StreamingStreamPoolResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StreamAdminClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteStreamingStreamPool(d *schema.ResourceData, m interface{}) error {
	sync := &StreamingStreamPoolResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StreamAdminClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type StreamingStreamPoolResourceCrud struct {
	tfresource.BaseCrud
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

	if customEncryptionKey, ok := s.D.GetOkExists("custom_encryption_key"); ok {
		if tmpList := customEncryptionKey.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "custom_encryption_key", 0)
			tmp, err := s.mapToCustomEncryptionKeyDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.CustomEncryptionKeyDetails = &tmp
		}
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

	if privateEndpointSettings, ok := s.D.GetOkExists("private_endpoint_settings"); ok {
		if tmpList := privateEndpointSettings.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "private_endpoint_settings", 0)
			tmp, err := s.mapToPrivateEndpointDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.PrivateEndpointDetails = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "streaming")

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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "streaming")

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

	if customEncryptionKey, ok := s.D.GetOkExists("custom_encryption_key"); ok && s.D.HasChange("custom_encryption_key") {
		if tmpList := customEncryptionKey.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "custom_encryption_key", 0)
			tmp, err := s.mapToCustomEncryptionKeyDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.CustomEncryptionKeyDetails = &tmp
		}
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "streaming")

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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "streaming")

	_, err := s.Client.DeleteStreamPool(context.Background(), request)
	return err
}

func (s *StreamingStreamPoolResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CustomEncryptionKey != nil {
		s.D.Set("custom_encryption_key", []interface{}{CustomEncryptionKeyToMap(s.Res.CustomEncryptionKey)})
	} else {
		s.D.Set("custom_encryption_key", nil)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.EndpointFqdn != nil {
		s.D.Set("endpoint_fqdn", *s.Res.EndpointFqdn)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsPrivate != nil {
		s.D.Set("is_private", *s.Res.IsPrivate)
	}

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

	if s.Res.PrivateEndpointSettings != nil {
		s.D.Set("private_endpoint_settings", []interface{}{PrivateEndpointSettingsToMap(s.Res.PrivateEndpointSettings, false)})
	} else {
		s.D.Set("private_endpoint_settings", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func (s *StreamingStreamPoolResourceCrud) mapToCustomEncryptionKeyDetails(fieldKeyFormat string) (oci_streaming.CustomEncryptionKeyDetails, error) {
	result := oci_streaming.CustomEncryptionKeyDetails{}

	if kmsKeyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kms_key_id")); ok {
		tmp := kmsKeyId.(string)
		result.KmsKeyId = &tmp
	}

	return result, nil
}

func CustomEncryptionKeyToMap(obj *oci_streaming.CustomEncryptionKey) map[string]interface{} {
	result := map[string]interface{}{}

	result["key_state"] = string(obj.KeyState)

	if obj.KmsKeyId != nil {
		result["kms_key_id"] = string(*obj.KmsKeyId)
	}

	return result
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

func (s *StreamingStreamPoolResourceCrud) mapToPrivateEndpointDetails(fieldKeyFormat string) (oci_streaming.PrivateEndpointDetails, error) {
	result := oci_streaming.PrivateEndpointDetails{}

	if nsgIds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "nsg_ids")); ok {
		set := nsgIds.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "nsg_ids")) {
			result.NsgIds = tmp
		}
	}

	if privateEndpointIp, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "private_endpoint_ip")); ok {
		tmp := privateEndpointIp.(string)
		result.PrivateEndpointIp = &tmp
	}

	if subnetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "subnet_id")); ok {
		tmp := subnetId.(string)
		result.SubnetId = &tmp
	}

	return result, nil
}

func PrivateEndpointSettingsToMap(obj *oci_streaming.PrivateEndpointSettings, datasource bool) map[string]interface{} {
	result := map[string]interface{}{}

	nsgIds := []interface{}{}
	for _, item := range obj.NsgIds {
		nsgIds = append(nsgIds, item)
	}
	if datasource {
		result["nsg_ids"] = nsgIds
	} else {
		result["nsg_ids"] = schema.NewSet(utils.LiteralTypeHashCodeForSets, nsgIds)
	}

	if obj.PrivateEndpointIp != nil {
		result["private_endpoint_ip"] = string(*obj.PrivateEndpointIp)
	}

	if obj.SubnetId != nil {
		result["subnet_id"] = string(*obj.SubnetId)
	}

	return result
}

func (s *StreamingStreamPoolResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_streaming.ChangeStreamPoolCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.StreamPoolId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "streaming")

	_, err := s.Client.ChangeStreamPoolCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
