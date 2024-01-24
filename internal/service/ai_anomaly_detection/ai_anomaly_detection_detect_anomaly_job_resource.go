// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ai_anomaly_detection

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_ai_anomaly_detection "github.com/oracle/oci-go-sdk/v65/aianomalydetection"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func AiAnomalyDetectionDetectAnomalyJobResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createAiAnomalyDetectionDetectAnomalyJob,
		Read:     readAiAnomalyDetectionDetectAnomalyJob,
		Update:   updateAiAnomalyDetectionDetectAnomalyJob,
		Delete:   deleteAiAnomalyDetectionDetectAnomalyJob,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"input_details": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"input_type": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"BASE64_ENCODED",
								"INLINE",
								"OBJECT_LIST",
							}, true),
						},

						// Optional
						"content": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"content_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"data": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"timestamp": {
										Type:             schema.TypeString,
										Optional:         true,
										Computed:         true,
										ForceNew:         true,
										DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
									},
									"values": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										ForceNew: true,
										Elem: &schema.Schema{
											Type: schema.TypeFloat,
										},
									},

									// Computed
								},
							},
						},
						"object_locations": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"bucket": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"namespace": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"object": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},

									// Computed
								},
							},
						},
						"signal_names": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						// Computed
						"message": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"model_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"output_details": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"bucket": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"namespace": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"output_type": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"OBJECT_STORAGE",
							}, true),
						},

						// Optional
						"prefix": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},

			// Optional
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sensitivity": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"defined_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"lifecycle_state_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"project_id": {
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"time_accepted": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_finished": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_started": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createAiAnomalyDetectionDetectAnomalyJob(d *schema.ResourceData, m interface{}) error {
	sync := &AiAnomalyDetectionDetectAnomalyJobResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AnomalyDetectionClient()

	return tfresource.CreateResource(d, sync)
}

func readAiAnomalyDetectionDetectAnomalyJob(d *schema.ResourceData, m interface{}) error {
	sync := &AiAnomalyDetectionDetectAnomalyJobResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AnomalyDetectionClient()

	return tfresource.ReadResource(sync)
}

func updateAiAnomalyDetectionDetectAnomalyJob(d *schema.ResourceData, m interface{}) error {
	sync := &AiAnomalyDetectionDetectAnomalyJobResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AnomalyDetectionClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteAiAnomalyDetectionDetectAnomalyJob(d *schema.ResourceData, m interface{}) error {
	sync := &AiAnomalyDetectionDetectAnomalyJobResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AnomalyDetectionClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type AiAnomalyDetectionDetectAnomalyJobResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_ai_anomaly_detection.AnomalyDetectionClient
	Res                    *oci_ai_anomaly_detection.DetectAnomalyJob
	DisableNotFoundRetries bool
}

func (s *AiAnomalyDetectionDetectAnomalyJobResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *AiAnomalyDetectionDetectAnomalyJobResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_ai_anomaly_detection.DetectAnomalyJobLifecycleStateInProgress),
	}
}

func (s *AiAnomalyDetectionDetectAnomalyJobResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_ai_anomaly_detection.DetectAnomalyJobLifecycleStateAccepted),
	}
}

func (s *AiAnomalyDetectionDetectAnomalyJobResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_ai_anomaly_detection.DetectAnomalyJobLifecycleStateInProgress),
	}
}

func (s *AiAnomalyDetectionDetectAnomalyJobResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_ai_anomaly_detection.DetectAnomalyJobLifecycleStateCanceled),
	}
}

func (s *AiAnomalyDetectionDetectAnomalyJobResourceCrud) Create() error {
	request := oci_ai_anomaly_detection.CreateDetectAnomalyJobRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if inputDetails, ok := s.D.GetOkExists("input_details"); ok {
		if tmpList := inputDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "input_details", 0)
			tmp, err := s.mapToInputDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.InputDetails = tmp
		}
	}

	if modelId, ok := s.D.GetOkExists("model_id"); ok {
		tmp := modelId.(string)
		request.ModelId = &tmp
	}

	if outputDetails, ok := s.D.GetOkExists("output_details"); ok {
		if tmpList := outputDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "output_details", 0)
			tmp, err := s.mapToOutputDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.OutputDetails = tmp
		}
	}

	if sensitivity, ok := s.D.GetOkExists("sensitivity"); ok {
		tmp := sensitivity.(float64)
		f32 := float32(tmp)
		request.Sensitivity = &f32
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_anomaly_detection")

	response, err := s.Client.CreateDetectAnomalyJob(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DetectAnomalyJob
	return nil
}

func (s *AiAnomalyDetectionDetectAnomalyJobResourceCrud) Get() error {
	request := oci_ai_anomaly_detection.GetDetectAnomalyJobRequest{}

	tmp := s.D.Id()
	request.DetectAnomalyJobId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_anomaly_detection")

	response, err := s.Client.GetDetectAnomalyJob(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DetectAnomalyJob
	return nil
}

func (s *AiAnomalyDetectionDetectAnomalyJobResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_ai_anomaly_detection.UpdateDetectAnomalyJobRequest{}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	tmp := s.D.Id()
	request.DetectAnomalyJobId = &tmp

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_anomaly_detection")

	response, err := s.Client.UpdateDetectAnomalyJob(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DetectAnomalyJob
	return nil
}

func (s *AiAnomalyDetectionDetectAnomalyJobResourceCrud) Delete() error {
	request := oci_ai_anomaly_detection.DeleteDetectAnomalyJobRequest{}

	tmp := s.D.Id()
	request.DetectAnomalyJobId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_anomaly_detection")

	_, err := s.Client.DeleteDetectAnomalyJob(context.Background(), request)
	return err
}

func (s *AiAnomalyDetectionDetectAnomalyJobResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.InputDetails != nil {
		inputDetailsArray := []interface{}{}
		if inputDetailsMap := InputJobDetailsToMap(&s.Res.InputDetails); inputDetailsMap != nil {
			inputDetailsArray = append(inputDetailsArray, inputDetailsMap)
		}
		s.D.Set("input_details", inputDetailsArray)
	} else {
		s.D.Set("input_details", nil)
	}

	if s.Res.LifecycleStateDetails != nil {
		s.D.Set("lifecycle_state_details", *s.Res.LifecycleStateDetails)
	}

	if s.Res.ModelId != nil {
		s.D.Set("model_id", *s.Res.ModelId)
	}

	if s.Res.OutputDetails != nil {
		outputDetailsArray := []interface{}{}
		if outputDetailsMap := OutputJobDetailsToMap(&s.Res.OutputDetails); outputDetailsMap != nil {
			outputDetailsArray = append(outputDetailsArray, outputDetailsMap)
		}
		s.D.Set("output_details", outputDetailsArray)
	} else {
		s.D.Set("output_details", nil)
	}

	if s.Res.ProjectId != nil {
		s.D.Set("project_id", *s.Res.ProjectId)
	}

	if s.Res.Sensitivity != nil {
		s.D.Set("sensitivity", *s.Res.Sensitivity)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		//s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeAccepted != nil {
		s.D.Set("time_accepted", s.Res.TimeAccepted.String())
	}

	if s.Res.TimeFinished != nil {
		s.D.Set("time_finished", s.Res.TimeFinished.String())
	}

	if s.Res.TimeStarted != nil {
		s.D.Set("time_started", s.Res.TimeStarted.String())
	}

	return nil
}

func (s *AiAnomalyDetectionDetectAnomalyJobResourceCrud) mapToDataItem(fieldKeyFormat string) (oci_ai_anomaly_detection.DataItem, error) {
	result := oci_ai_anomaly_detection.DataItem{}

	if timestamp, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "timestamp")); ok {
		tmp, err := time.Parse(time.RFC3339, timestamp.(string))
		if err != nil {
			return result, err
		}
		result.Timestamp = &oci_common.SDKTime{Time: tmp}
	}

	if values, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "values")); ok {
		interfaces := values.([]interface{})
		tmp := make([]float64, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(float64)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "values")) {
			result.Values = tmp
		}
	}

	return result, nil
}

func DataItemToMap(obj oci_ai_anomaly_detection.DataItem) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Timestamp != nil {
		result["timestamp"] = obj.Timestamp.Format(time.RFC3339Nano)
	}

	result["values"] = obj.Values

	return result
}

func DetectAnomalyJobSummaryToMap(obj oci_ai_anomaly_detection.DetectAnomalyJobSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleStateDetails != nil {
		result["lifecycle_state_details"] = string(*obj.LifecycleStateDetails)
	}

	if obj.ModelId != nil {
		result["model_id"] = string(*obj.ModelId)
	}

	if obj.ProjectId != nil {
		result["project_id"] = string(*obj.ProjectId)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		//result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeAccepted != nil {
		result["time_accepted"] = obj.TimeAccepted.String()
	}

	if obj.TimeFinished != nil {
		result["time_finished"] = obj.TimeFinished.String()
	}

	if obj.TimeStarted != nil {
		result["time_started"] = obj.TimeStarted.String()
	}

	return result
}

func (s *AiAnomalyDetectionDetectAnomalyJobResourceCrud) mapToInputDetails(fieldKeyFormat string) (oci_ai_anomaly_detection.InputDetails, error) {
	var baseObject oci_ai_anomaly_detection.InputDetails
	//discriminator
	inputTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "input_type"))
	var inputType string
	if ok {
		inputType = inputTypeRaw.(string)
	} else {
		inputType = "" // default value
	}
	switch strings.ToLower(inputType) {
	case strings.ToLower("BASE64_ENCODED"):
		details := oci_ai_anomaly_detection.EmbeddedInputDetails{}
		if content, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "content")); ok {
			tmp := content.(string)
			details.Content = []byte(tmp)
		}
		if contentType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "content_type")); ok {
			details.ContentType = oci_ai_anomaly_detection.EmbeddedInputDetailsContentTypeEnum(contentType.(string))
		}
		baseObject = details
	case strings.ToLower("INLINE"):
		details := oci_ai_anomaly_detection.InlineInputDetails{}
		if data, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "data")); ok {
			interfaces := data.([]interface{})
			tmp := make([]oci_ai_anomaly_detection.DataItem, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "data"), stateDataIndex)
				converted, err := s.mapToDataItem(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "data")) {
				details.Data = tmp
			}
		}
		if signalNames, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "signal_names")); ok {
			interfaces := signalNames.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "signal_names")) {
				details.SignalNames = tmp
			}
		}
		baseObject = details
	case strings.ToLower("OBJECT_LIST"):
		details := oci_ai_anomaly_detection.ObjectListInputJobDetails{}
		if objectLocations, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_locations")); ok {
			interfaces := objectLocations.([]interface{})
			tmp := make([]oci_ai_anomaly_detection.ObjectLocation, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "object_locations"), stateDataIndex)
				converted, err := s.mapToObjectLocation(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "object_locations")) {
				details.ObjectLocations = tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown input_type '%v' was specified", inputType)
	}
	return baseObject, nil
}

func InputJobDetailsToMap(obj *oci_ai_anomaly_detection.InputJobDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_ai_anomaly_detection.EmbeddedInputDetails:
		result["input_type"] = "BASE64_ENCODED"

		if v.Content != nil {
			tmp := v.Content
			result["content"] = string(tmp[:])
		}

		result["content_type"] = string(v.ContentType)
	case oci_ai_anomaly_detection.InlineInputDetails:
		result["input_type"] = "INLINE"

		data := []interface{}{}
		for _, item := range v.Data {
			data = append(data, DataItemToMap(item))
		}
		result["data"] = data
		result["signal_names"] = v.SignalNames
	case oci_ai_anomaly_detection.ObjectListInputJobDetails:
		result["input_type"] = "OBJECT_LIST"

		objectLocations := []interface{}{}
		for _, item := range v.ObjectLocations {
			objectLocations = append(objectLocations, ObjectLocationToMap(item))
		}
		result["object_locations"] = objectLocations
	default:
		log.Printf("[WARN] Received 'input_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *AiAnomalyDetectionDetectAnomalyJobResourceCrud) mapToObjectLocation(fieldKeyFormat string) (oci_ai_anomaly_detection.ObjectLocation, error) {
	result := oci_ai_anomaly_detection.ObjectLocation{}

	if bucket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
		tmp := bucket.(string)
		result.BucketName = &tmp
	}

	if namespace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
		tmp := namespace.(string)
		result.NamespaceName = &tmp
	}

	if object, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object")); ok {
		tmp := object.(string)
		result.ObjectName = &tmp
	}

	return result, nil
}

func ObjectLocationToMap(obj oci_ai_anomaly_detection.ObjectLocation) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BucketName != nil {
		result["bucket"] = string(*obj.BucketName)
	}

	if obj.NamespaceName != nil {
		result["namespace"] = string(*obj.NamespaceName)
	}

	if obj.ObjectName != nil {
		result["object"] = string(*obj.ObjectName)
	}
	return result
}

func (s *AiAnomalyDetectionDetectAnomalyJobResourceCrud) mapToOutputDetails(fieldKeyFormat string) (oci_ai_anomaly_detection.OutputDetails, error) {
	var baseObject oci_ai_anomaly_detection.OutputDetails
	//discriminator
	outputTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "output_type"))
	var outputType string
	if ok {
		outputType = outputTypeRaw.(string)
	} else {
		outputType = "" // default value
	}
	switch strings.ToLower(outputType) {
	case strings.ToLower("OBJECT_STORAGE"):
		details := oci_ai_anomaly_detection.ObjectStoreOutputDetails{}
		if bucket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
			tmp := bucket.(string)
			details.BucketName = &tmp
		}
		if namespace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
			tmp := namespace.(string)
			details.NamespaceName = &tmp
		}
		if prefix, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "prefix")); ok {
			tmp := prefix.(string)
			details.Prefix = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown output_type '%v' was specified", outputType)
	}
	return baseObject, nil
}

func OutputJobDetailsToMap(obj *oci_ai_anomaly_detection.OutputJobDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_ai_anomaly_detection.ObjectStorageLocation:
		result["output_type"] = "OBJECT_STORAGE"

		if v.BucketName != nil {
			result["bucket"] = string(*v.BucketName)
		}

		if v.NamespaceName != nil {
			result["namespace"] = string(*v.NamespaceName)
		}

		if v.Prefix != nil {
			result["prefix"] = string(*v.Prefix)
		}
	default:
		log.Printf("[WARN] Received 'output_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *AiAnomalyDetectionDetectAnomalyJobResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_ai_anomaly_detection.ChangeDetectAnomalyJobCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.DetectAnomalyJobId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_anomaly_detection")

	_, err := s.Client.ChangeDetectAnomalyJobCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
