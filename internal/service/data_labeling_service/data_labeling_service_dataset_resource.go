// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_labeling_service

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	oci_common "github.com/oracle/oci-go-sdk/v58/common"
	oci_data_labeling_service "github.com/oracle/oci-go-sdk/v58/datalabelingservice"
)

func DataLabelingServiceDatasetResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDataLabelingServiceDataset,
		Read:     readDataLabelingServiceDataset,
		Update:   updateDataLabelingServiceDataset,
		Delete:   deleteDataLabelingServiceDataset,
		Schema: map[string]*schema.Schema{
			// Required
			"annotation_format": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"dataset_format_details": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"format_type": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"dataset_source_details": {
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
						"source_type": {
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
			"label_set": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"items": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"name": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional

									// Computed
								},
							},
						},

						// Optional

						// Computed
					},
				},
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
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
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"initial_record_generation_configuration": {
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

						// Computed
					},
				},
			},
			"labeling_instructions": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"lifecycle_details": {
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
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDataLabelingServiceDataset(d *schema.ResourceData, m interface{}) error {
	sync := &DataLabelingServiceDatasetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataLabelingManagementClient()

	return tfresource.CreateResource(d, sync)
}

func readDataLabelingServiceDataset(d *schema.ResourceData, m interface{}) error {
	sync := &DataLabelingServiceDatasetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataLabelingManagementClient()

	return tfresource.ReadResource(sync)
}

func updateDataLabelingServiceDataset(d *schema.ResourceData, m interface{}) error {
	sync := &DataLabelingServiceDatasetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataLabelingManagementClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDataLabelingServiceDataset(d *schema.ResourceData, m interface{}) error {
	sync := &DataLabelingServiceDatasetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataLabelingManagementClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DataLabelingServiceDatasetResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_data_labeling_service.DataLabelingManagementClient
	Res                    *oci_data_labeling_service.Dataset
	DisableNotFoundRetries bool
}

func (s *DataLabelingServiceDatasetResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DataLabelingServiceDatasetResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_data_labeling_service.DatasetLifecycleStateCreating),
	}
}

func (s *DataLabelingServiceDatasetResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_data_labeling_service.DatasetLifecycleStateActive),
		string(oci_data_labeling_service.DatasetLifecycleStateNeedsAttention),
	}
}

func (s *DataLabelingServiceDatasetResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_data_labeling_service.DatasetLifecycleStateDeleting),
	}
}

func (s *DataLabelingServiceDatasetResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_data_labeling_service.DatasetLifecycleStateDeleted),
	}
}

func (s *DataLabelingServiceDatasetResourceCrud) Create() error {
	request := oci_data_labeling_service.CreateDatasetRequest{}

	if annotationFormat, ok := s.D.GetOkExists("annotation_format"); ok {
		tmp := annotationFormat.(string)
		request.AnnotationFormat = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if datasetFormatDetails, ok := s.D.GetOkExists("dataset_format_details"); ok {
		if tmpList := datasetFormatDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "dataset_format_details", 0)
			tmp, err := s.mapToDatasetFormatDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.DatasetFormatDetails = tmp
		}
	}

	if datasetSourceDetails, ok := s.D.GetOkExists("dataset_source_details"); ok {
		if tmpList := datasetSourceDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "dataset_source_details", 0)
			tmp, err := s.mapToDatasetSourceDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.DatasetSourceDetails = tmp
		}
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if initialRecordGenerationConfiguration, ok := s.D.GetOkExists("initial_record_generation_configuration"); ok {
		if tmpList := initialRecordGenerationConfiguration.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "initial_record_generation_configuration", 0)
			tmp, err := s.mapToInitialRecordGenerationConfiguration(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.InitialRecordGenerationConfiguration = &tmp
		}
	}

	if labelSet, ok := s.D.GetOkExists("label_set"); ok {
		if tmpList := labelSet.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "label_set", 0)
			tmp, err := s.mapToLabelSet(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.LabelSet = &tmp
		}
	}

	if labelingInstructions, ok := s.D.GetOkExists("labeling_instructions"); ok {
		tmp := labelingInstructions.(string)
		request.LabelingInstructions = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_labeling_service")

	response, err := s.Client.CreateDataset(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getDatasetFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_labeling_service"), oci_data_labeling_service.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DataLabelingServiceDatasetResourceCrud) getDatasetFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_data_labeling_service.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	datasetId, err := datasetWaitForWorkRequest(workId, "dataset",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*datasetId)

	return s.Get()
}

func datasetWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "data_labeling_service", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_data_labeling_service.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func datasetWaitForWorkRequest(wId *string, entityType string, action oci_data_labeling_service.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_data_labeling_service.DataLabelingManagementClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "data_labeling_service")
	retryPolicy.ShouldRetryOperation = datasetWorkRequestShouldRetryFunc(timeout)

	response := oci_data_labeling_service.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_data_labeling_service.OperationStatusInProgress),
			string(oci_data_labeling_service.OperationStatusAccepted),
			string(oci_data_labeling_service.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_data_labeling_service.OperationStatusSucceeded),
			string(oci_data_labeling_service.OperationStatusFailed),
			string(oci_data_labeling_service.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_data_labeling_service.GetWorkRequestRequest{
					WorkRequestId: wId,
					RequestMetadata: oci_common.RequestMetadata{
						RetryPolicy: retryPolicy,
					},
				})
			wr := &response.WorkRequest
			return wr, string(wr.Status), err
		},
		Timeout: timeout,
	}
	if _, e := stateConf.WaitForState(); e != nil {
		return nil, e
	}

	var identifier *string
	// The work request response contains an array of objects that finished the operation
	for _, res := range response.Resources {
		if strings.Contains(strings.ToLower(*res.EntityType), entityType) {
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_data_labeling_service.OperationStatusFailed || response.Status == oci_data_labeling_service.OperationStatusCanceled {
		return nil, getErrorFromDataLabelingServiceDatasetWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDataLabelingServiceDatasetWorkRequest(client *oci_data_labeling_service.DataLabelingManagementClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_data_labeling_service.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_data_labeling_service.ListWorkRequestErrorsRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: retryPolicy,
			},
		})
	if err != nil {
		return err
	}

	allErrs := make([]string, 0)
	for _, wrkErr := range response.Items {
		allErrs = append(allErrs, *wrkErr.Message)
	}
	errorMessage := strings.Join(allErrs, "\n")

	workRequestErr := fmt.Errorf("work request did not succeed, workId: %s, entity: %s, action: %s. Message: %s", *workId, entityType, action, errorMessage)

	return workRequestErr
}

func (s *DataLabelingServiceDatasetResourceCrud) Get() error {
	request := oci_data_labeling_service.GetDatasetRequest{}

	tmp := s.D.Id()
	request.DatasetId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_labeling_service")

	response, err := s.Client.GetDataset(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Dataset
	return nil
}

func (s *DataLabelingServiceDatasetResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}

	if _, ok := s.D.GetOkExists("label_set"); ok && s.D.HasChange(fmt.Sprintf("%s.%d.%s", "label_set", 0, "items")) {
		oldRaw, newRaw := s.D.GetChange(fmt.Sprintf("%s.%d.%s", "label_set", 0, "items"))

		oldLabelSet, err := mapToLabelSetObject(oldRaw)
		if err != nil {
			return err
		}
		newLabelSet, err := mapToLabelSetObject(newRaw)
		if err != nil {
			return err
		}
		err = s.updateLabelSet(oldLabelSet, newLabelSet)
		if err != nil {
			return err
		}
	}

	request := oci_data_labeling_service.UpdateDatasetRequest{}

	tmp := s.D.Id()
	request.DatasetId = &tmp

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if labelingInstructions, ok := s.D.GetOkExists("labeling_instructions"); ok {
		tmp := labelingInstructions.(string)
		request.LabelingInstructions = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_labeling_service")

	response, err := s.Client.UpdateDataset(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Dataset
	return nil
}

func (s *DataLabelingServiceDatasetResourceCrud) Delete() error {
	request := oci_data_labeling_service.DeleteDatasetRequest{}

	tmp := s.D.Id()
	request.DatasetId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_labeling_service")

	response, err := s.Client.DeleteDataset(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := datasetWaitForWorkRequest(workId, "dataset",
		oci_data_labeling_service.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *DataLabelingServiceDatasetResourceCrud) SetData() error {
	if s.Res.AnnotationFormat != nil {
		s.D.Set("annotation_format", *s.Res.AnnotationFormat)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DatasetFormatDetails != nil {
		datasetFormatDetailsArray := []interface{}{}
		if datasetFormatDetailsMap := DatasetFormatDetailsToMap(&s.Res.DatasetFormatDetails); datasetFormatDetailsMap != nil {
			datasetFormatDetailsArray = append(datasetFormatDetailsArray, datasetFormatDetailsMap)
		}
		s.D.Set("dataset_format_details", datasetFormatDetailsArray)
	} else {
		s.D.Set("dataset_format_details", nil)
	}

	if s.Res.DatasetSourceDetails != nil {
		datasetSourceDetailsArray := []interface{}{}
		if datasetSourceDetailsMap := DatasetSourceDetailsToMap(&s.Res.DatasetSourceDetails); datasetSourceDetailsMap != nil {
			datasetSourceDetailsArray = append(datasetSourceDetailsArray, datasetSourceDetailsMap)
		}
		s.D.Set("dataset_source_details", datasetSourceDetailsArray)
	} else {
		s.D.Set("dataset_source_details", nil)
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

	if s.Res.InitialRecordGenerationConfiguration != nil {
		s.D.Set("initial_record_generation_configuration", []interface{}{InitialRecordGenerationConfigurationToMap(s.Res.InitialRecordGenerationConfiguration)})
	} else {
		s.D.Set("initial_record_generation_configuration", nil)
	}

	if s.Res.LabelSet != nil {
		s.D.Set("label_set", []interface{}{LabelSetToMap(s.Res.LabelSet)})
	} else {
		s.D.Set("label_set", nil)
	}

	if s.Res.LabelingInstructions != nil {
		s.D.Set("labeling_instructions", *s.Res.LabelingInstructions)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func (s *DataLabelingServiceDatasetResourceCrud) mapToDatasetFormatDetails(fieldKeyFormat string) (oci_data_labeling_service.DatasetFormatDetails, error) {
	var baseObject oci_data_labeling_service.DatasetFormatDetails
	//discriminator
	formatTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "format_type"))
	var formatType string
	if ok {
		formatType = formatTypeRaw.(string)
	} else {
		formatType = "" // default value
	}
	switch strings.ToLower(formatType) {
	case strings.ToLower("DOCUMENT"):
		baseObject = oci_data_labeling_service.DocumentDatasetFormatDetails{}
	case strings.ToLower("IMAGE"):
		baseObject = oci_data_labeling_service.ImageDatasetFormatDetails{}
	case strings.ToLower("TEXT"):
		baseObject = oci_data_labeling_service.TextDatasetFormatDetails{}
	default:
		return nil, fmt.Errorf("unknown format_type '%v' was specified", formatType)
	}
	return baseObject, nil
}

func DatasetFormatDetailsToMap(obj *oci_data_labeling_service.DatasetFormatDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_data_labeling_service.DocumentDatasetFormatDetails:
		result["format_type"] = "DOCUMENT"
	case oci_data_labeling_service.ImageDatasetFormatDetails:
		result["format_type"] = "IMAGE"
	case oci_data_labeling_service.TextDatasetFormatDetails:
		result["format_type"] = "TEXT"
	default:
		log.Printf("[WARN] Received 'format_type' of unknown type %v of type %v", *obj, v)
		return nil
	}

	return result
}

func (s *DataLabelingServiceDatasetResourceCrud) mapToDatasetSourceDetails(fieldKeyFormat string) (oci_data_labeling_service.DatasetSourceDetails, error) {
	var baseObject oci_data_labeling_service.DatasetSourceDetails
	//discriminator
	sourceTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_type"))
	var sourceType string
	if ok {
		sourceType = sourceTypeRaw.(string)
	} else {
		sourceType = "" // default value
	}
	switch strings.ToLower(sourceType) {
	case strings.ToLower("OBJECT_STORAGE"):
		details := oci_data_labeling_service.ObjectStorageSourceDetails{}
		if bucket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
			tmp := bucket.(string)
			details.Bucket = &tmp
		}
		if namespace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
			tmp := namespace.(string)
			details.Namespace = &tmp
		}
		if prefix, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "prefix")); ok {
			tmp := prefix.(string)
			details.Prefix = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown source_type '%v' was specified", sourceType)
	}
	return baseObject, nil
}

func DatasetSourceDetailsToMap(obj *oci_data_labeling_service.DatasetSourceDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_data_labeling_service.ObjectStorageSourceDetails:
		result["source_type"] = "OBJECT_STORAGE"

		if v.Bucket != nil {
			result["bucket"] = string(*v.Bucket)
		}

		if v.Namespace != nil {
			result["namespace"] = string(*v.Namespace)
		}

		if v.Prefix != nil {
			result["prefix"] = string(*v.Prefix)
		}
	default:
		log.Printf("[WARN] Received 'source_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func DatasetSummaryToMap(obj oci_data_labeling_service.DatasetSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AnnotationFormat != nil {
		result["annotation_format"] = string(*obj.AnnotationFormat)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DatasetFormatDetails != nil {
		datasetFormatDetailsArray := []interface{}{}
		if datasetFormatDetailsMap := DatasetFormatDetailsToMap(&obj.DatasetFormatDetails); datasetFormatDetailsMap != nil {
			datasetFormatDetailsArray = append(datasetFormatDetailsArray, datasetFormatDetailsMap)
		}
		result["dataset_format_details"] = datasetFormatDetailsArray
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func (s *DataLabelingServiceDatasetResourceCrud) mapToInitialRecordGenerationConfiguration(fieldKeyFormat string) (oci_data_labeling_service.InitialRecordGenerationConfiguration, error) {
	result := oci_data_labeling_service.InitialRecordGenerationConfiguration{}

	if limit, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "limit")); ok {
		tmp := limit.(float32)
		result.Limit = &tmp
	}

	return result, nil
}

func InitialRecordGenerationConfigurationToMap(obj *oci_data_labeling_service.InitialRecordGenerationConfiguration) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Limit != nil {
		result["limit"] = float32(*obj.Limit)
	}

	return result
}

func mapToLabelSetObject(labeSetRaw interface{}) (oci_data_labeling_service.LabelSet, error) {
	result := oci_data_labeling_service.LabelSet{}

	interfaces := labeSetRaw.([]interface{})
	tmp := make([]oci_data_labeling_service.Label, len(interfaces))
	for i, labelRaw := range interfaces {
		label := oci_data_labeling_service.Label{}
		labelObj := labelRaw.(map[string]interface{})
		labelStr := labelObj["name"].(string)
		label.Name = &labelStr

		tmp[i] = label
	}
	if len(tmp) != 0 {
		result.Items = tmp
	}
	return result, nil
}

func (s *DataLabelingServiceDatasetResourceCrud) mapToLabel(fieldKeyFormat string) (oci_data_labeling_service.Label, error) {
	result := oci_data_labeling_service.Label{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	return result, nil
}

func LabelToMap(obj oci_data_labeling_service.Label) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}

func (s *DataLabelingServiceDatasetResourceCrud) mapToLabelSet(fieldKeyFormat string) (oci_data_labeling_service.LabelSet, error) {
	result := oci_data_labeling_service.LabelSet{}

	if items, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "items")); ok {
		interfaces := items.([]interface{})
		tmp := make([]oci_data_labeling_service.Label, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "items"), stateDataIndex)
			converted, err := s.mapToLabel(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "items")) {
			result.Items = tmp
		}
	}

	return result, nil
}

func LabelSetToMap(obj *oci_data_labeling_service.LabelSet) map[string]interface{} {
	result := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range obj.Items {
		items = append(items, LabelToMap(item))
	}
	result["items"] = items

	return result
}

func (s *DataLabelingServiceDatasetResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_data_labeling_service.ChangeDatasetCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.DatasetId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_labeling_service")

	response, err := s.Client.ChangeDatasetCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getDatasetFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_labeling_service"), oci_data_labeling_service.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DataLabelingServiceDatasetResourceCrud) updateLabelSet(oldLabelSet, newLabelSet oci_data_labeling_service.LabelSet) error {
	addDatasetLabelsRequest := oci_data_labeling_service.AddDatasetLabelsRequest{}
	labelSetDiffAdd := make([]oci_data_labeling_service.Label, 0)

	for _, newLabel := range newLabelSet.Items {
		foundLabel := false
		for _, oldLabel := range oldLabelSet.Items {
			if *oldLabel.Name == *newLabel.Name {
				foundLabel = true
				break
			}
		}
		if !foundLabel {
			labelSetDiffAdd = append(labelSetDiffAdd, newLabel)
		}
	}

	if len(labelSetDiffAdd) > 0 {
		idTmp := s.D.Id()
		addDatasetLabelsRequest.DatasetId = &idTmp

		addDatasetLabelsRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_labeling_service")
		addDatasetLabelsRequest.LabelSet = &oci_data_labeling_service.LabelSet{Items: labelSetDiffAdd}

		response, err := s.Client.AddDatasetLabels(context.Background(), addDatasetLabelsRequest)
		if err != nil {
			return err
		}

		workId := response.OpcWorkRequestId
		err = s.getDatasetFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_labeling_service"), oci_data_labeling_service.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return err
		}
	}
	removeDatasetLabelsRequest := oci_data_labeling_service.RemoveDatasetLabelsRequest{}
	labelSetDiffRemove := make([]oci_data_labeling_service.Label, 0)

	for _, oldLabel := range oldLabelSet.Items {
		foundLabel := false
		for _, newLabel := range newLabelSet.Items {
			if *oldLabel.Name == *newLabel.Name {
				foundLabel = true
				break
			}
		}
		if !foundLabel {
			labelSetDiffRemove = append(labelSetDiffRemove, oldLabel)
		}
	}

	if len(labelSetDiffRemove) > 0 {
		idTmp := s.D.Id()
		removeDatasetLabelsRequest.DatasetId = &idTmp

		removeDatasetLabelsRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_labeling_service")
		removeDatasetLabelsRequest.LabelSet = &oci_data_labeling_service.LabelSet{Items: labelSetDiffRemove}

		response, err := s.Client.RemoveDatasetLabels(context.Background(), removeDatasetLabelsRequest)
		if err != nil {
			return err
		}

		workId := response.OpcWorkRequestId
		err = s.getDatasetFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_labeling_service"), oci_data_labeling_service.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return err
		}
	}

	return nil
}
