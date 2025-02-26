// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ai_document

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_ai_document "github.com/oracle/oci-go-sdk/v65/aidocument"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func AiDocumentProcessorJobResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createAiDocumentProcessorJob,
		Read:     readAiDocumentProcessorJob,
		Delete:   deleteAiDocumentProcessorJob,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"input_location": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"source_type": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"INLINE_DOCUMENT_CONTENT",
								"OBJECT_STORAGE_LOCATIONS",
							}, true),
						},

						// Optional
						"data": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
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
									"page_range": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										ForceNew: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									// Computed
								},
							},
						},
						"page_range": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						// Computed
					},
				},
			},
			"output_location": {
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
						"prefix": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"processor_config": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"processor_type": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"GENERAL",
								"INVOICE",
							}, true),
						},

						// Optional
						"document_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"features": {
							Type:     schema.TypeList,
							Required: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"feature_type": {
										Type:             schema.TypeString,
										Required:         true,
										ForceNew:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"DOCUMENT_CLASSIFICATION",
											"DOCUMENT_ELEMENTS_EXTRACTION",
											"KEY_VALUE_EXTRACTION",
											"LANGUAGE_CLASSIFICATION",
											"TABLE_EXTRACTION",
											"TEXT_EXTRACTION",
										}, true),
									},

									// Optional
									"generate_searchable_pdf": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"max_results": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"model_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"selection_mark_detection": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"tenancy_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},

									// Computed
								},
							},
						},
						"is_zip_output_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"language": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"model_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"normalization_fields": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"map": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										ForceNew: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"normalization_type": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
											},

											// Computed
										},
									},
								},
							},
						},

						// Computed
					},
				},
			},

			// Optional
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"percent_complete": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
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

func createAiDocumentProcessorJob(d *schema.ResourceData, m interface{}) error {
	sync := &AiDocumentProcessorJobResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AiServiceDocumentClient()

	return tfresource.CreateResource(d, sync)
}

func readAiDocumentProcessorJob(d *schema.ResourceData, m interface{}) error {
	sync := &AiDocumentProcessorJobResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AiServiceDocumentClient()

	return tfresource.ReadResource(sync)
}

func deleteAiDocumentProcessorJob(d *schema.ResourceData, m interface{}) error {
	return nil
}

type AiDocumentProcessorJobResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_ai_document.AIServiceDocumentClient
	Res                    *oci_ai_document.ProcessorJob
	DisableNotFoundRetries bool
}

func (s *AiDocumentProcessorJobResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *AiDocumentProcessorJobResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_ai_document.ProcessorJobLifecycleStateInProgress),
	}
}

func (s *AiDocumentProcessorJobResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_ai_document.ProcessorJobLifecycleStateAccepted),
		string(oci_ai_document.ProcessorJobLifecycleStateSucceeded),
	}
}

func (s *AiDocumentProcessorJobResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *AiDocumentProcessorJobResourceCrud) DeletedTarget() []string {
	return []string{}
}

func (s *AiDocumentProcessorJobResourceCrud) Create() error {
	request := oci_ai_document.CreateProcessorJobRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if inputLocation, ok := s.D.GetOkExists("input_location"); ok {
		if tmpList := inputLocation.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "input_location", 0)
			tmp, err := s.mapToInputLocation(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.InputLocation = tmp
		}
	}

	if outputLocation, ok := s.D.GetOkExists("output_location"); ok {
		if tmpList := outputLocation.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "output_location", 0)
			tmp, err := s.mapToOutputLocation(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.OutputLocation = &tmp
		}
	}

	if processorConfig, ok := s.D.GetOkExists("processor_config"); ok {
		if tmpList := processorConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "processor_config", 0)
			tmp, err := s.mapToProcessorConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ProcessorConfig = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_document")

	response, err := s.Client.CreateProcessorJob(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ProcessorJob
	return nil
}

func (s *AiDocumentProcessorJobResourceCrud) Get() error {
	request := oci_ai_document.GetProcessorJobRequest{}

	tmp := s.D.Id()
	request.ProcessorJobId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_document")

	response, err := s.Client.GetProcessorJob(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ProcessorJob
	return nil
}

func (s *AiDocumentProcessorJobResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", s.Res.DisplayName)
	}

	if s.Res.InputLocation != nil {
		inputLocationArray := []interface{}{}
		if inputLocationMap := InputLocationToMap(&s.Res.InputLocation); inputLocationMap != nil {
			inputLocationArray = append(inputLocationArray, inputLocationMap)
		}
		s.D.Set("input_location", inputLocationArray)
	} else {
		s.D.Set("input_location", nil)
	}

	s.D.Set("lifecycle_details", s.Res.LifecycleDetails)

	if s.Res.OutputLocation != nil {
		s.D.Set("output_location", []interface{}{OutputLocationToMap(s.Res.OutputLocation)})
	} else {
		s.D.Set("output_location", nil)
	}

	if s.Res.PercentComplete != nil {
		s.D.Set("percent_complete", *s.Res.PercentComplete)
	}

	if s.Res.ProcessorConfig != nil {
		processorConfigArray := []interface{}{}
		if processorConfigMap := ProcessorConfigToMap(&s.Res.ProcessorConfig); processorConfigMap != nil {
			processorConfigArray = append(processorConfigArray, processorConfigMap)
		}
		s.D.Set("processor_config", processorConfigArray)
	} else {
		s.D.Set("processor_config", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

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

func (s *AiDocumentProcessorJobResourceCrud) mapToDocumentFeature(fieldKeyFormat string) (oci_ai_document.DocumentFeature, error) {
	var baseObject oci_ai_document.DocumentFeature
	//discriminator
	featureTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "feature_type"))
	var featureType string
	if ok {
		featureType = featureTypeRaw.(string)
	} else {
		featureType = "" // default value
	}
	switch strings.ToLower(featureType) {
	case strings.ToLower("DOCUMENT_CLASSIFICATION"):
		details := oci_ai_document.DocumentClassificationFeature{}
		if maxResults, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max_results")); ok {
			tmp := maxResults.(int)
			details.MaxResults = &tmp
		}
		if modelId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "model_id")); ok {
			tmp := modelId.(string)
			details.ModelId = &tmp
		}
		if tenancyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tenancy_id")); ok {
			tmp := tenancyId.(string)
			details.TenancyId = &tmp
		}
		baseObject = details
	case strings.ToLower("DOCUMENT_ELEMENTS_EXTRACTION"):
		details := oci_ai_document.DocumentElementsExtractionFeature{}
		if modelId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "model_id")); ok {
			tmp := modelId.(string)
			details.ModelId = &tmp
		}
		baseObject = details
	case strings.ToLower("KEY_VALUE_EXTRACTION"):
		details := oci_ai_document.DocumentKeyValueExtractionFeature{}
		if modelId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "model_id")); ok {
			tmp := modelId.(string)
			details.ModelId = &tmp
		}
		if tenancyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tenancy_id")); ok {
			tmp := tenancyId.(string)
			details.TenancyId = &tmp
		}
		baseObject = details
	case strings.ToLower("LANGUAGE_CLASSIFICATION"):
		details := oci_ai_document.DocumentLanguageClassificationFeature{}
		if maxResults, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max_results")); ok {
			tmp := maxResults.(int)
			details.MaxResults = &tmp
		}
		baseObject = details
	case strings.ToLower("TABLE_EXTRACTION"):
		details := oci_ai_document.DocumentTableExtractionFeature{}
		if modelId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "model_id")); ok {
			tmp := modelId.(string)
			details.ModelId = &tmp
		}
		baseObject = details
	case strings.ToLower("TEXT_EXTRACTION"):
		details := oci_ai_document.DocumentTextExtractionFeature{}
		if generateSearchablePdf, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "generate_searchable_pdf")); ok {
			tmp := generateSearchablePdf.(bool)
			details.GenerateSearchablePdf = &tmp
		}
		if modelId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "model_id")); ok {
			tmp := modelId.(string)
			details.ModelId = &tmp
		}
		if selectionMarkDetection, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "selection_mark_detection")); ok {
			tmp := selectionMarkDetection.(bool)
			details.SelectionMarkDetection = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown feature_type '%v' was specified", featureType)
	}
	return baseObject, nil
}

func DocumentFeatureToMap(obj oci_ai_document.DocumentFeature) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_ai_document.DocumentClassificationFeature:
		result["feature_type"] = "DOCUMENT_CLASSIFICATION"

		if v.MaxResults != nil {
			result["max_results"] = int(*v.MaxResults)
		}

		if v.ModelId != nil {
			result["model_id"] = string(*v.ModelId)
		}

		if v.TenancyId != nil {
			result["tenancy_id"] = string(*v.TenancyId)
		}
	case oci_ai_document.DocumentElementsExtractionFeature:
		result["feature_type"] = "DOCUMENT_ELEMENTS_EXTRACTION"

		if v.ModelId != nil {
			result["model_id"] = string(*v.ModelId)
		}
	case oci_ai_document.DocumentKeyValueExtractionFeature:
		result["feature_type"] = "KEY_VALUE_EXTRACTION"

		if v.ModelId != nil {
			result["model_id"] = string(*v.ModelId)
		}

		if v.TenancyId != nil {
			result["tenancy_id"] = string(*v.TenancyId)
		}
	case oci_ai_document.DocumentLanguageClassificationFeature:
		result["feature_type"] = "LANGUAGE_CLASSIFICATION"

		if v.MaxResults != nil {
			result["max_results"] = int(*v.MaxResults)
		}
	case oci_ai_document.DocumentTableExtractionFeature:
		result["feature_type"] = "TABLE_EXTRACTION"

		if v.ModelId != nil {
			result["model_id"] = string(*v.ModelId)
		}
	case oci_ai_document.DocumentTextExtractionFeature:
		result["feature_type"] = "TEXT_EXTRACTION"

		if v.GenerateSearchablePdf != nil {
			result["generate_searchable_pdf"] = bool(*v.GenerateSearchablePdf)
		}

		if v.ModelId != nil {
			result["model_id"] = string(*v.ModelId)
		}

		if v.SelectionMarkDetection != nil {
			result["selection_mark_detection"] = bool(*v.SelectionMarkDetection)
		}
	default:
		log.Printf("[WARN] Received 'feature_type' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *AiDocumentProcessorJobResourceCrud) mapToInputLocation(fieldKeyFormat string) (oci_ai_document.InputLocation, error) {
	var baseObject oci_ai_document.InputLocation
	//discriminator
	sourceTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_type"))
	var sourceType string
	if ok {
		sourceType = sourceTypeRaw.(string)
	} else {
		sourceType = "" // default value
	}
	switch strings.ToLower(sourceType) {
	case strings.ToLower("INLINE_DOCUMENT_CONTENT"):
		details := oci_ai_document.InlineDocumentContent{}
		if data, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "data")); ok {
			tmp := data.(string)
			details.Data = []byte(tmp)
		}
		if pageRange, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "page_range")); ok {
			interfaces := pageRange.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "page_range")) {
				details.PageRange = tmp
			}
		}
		baseObject = details
	case strings.ToLower("OBJECT_STORAGE_LOCATIONS"):
		details := oci_ai_document.ObjectStorageLocations{}
		if objectLocations, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_locations")); ok {
			interfaces := objectLocations.([]interface{})
			tmp := make([]oci_ai_document.ObjectLocation, len(interfaces))
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
		return nil, fmt.Errorf("unknown source_type '%v' was specified", sourceType)
	}
	return baseObject, nil
}

func InputLocationToMap(obj *oci_ai_document.InputLocation) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_ai_document.InlineDocumentContent:
		result["source_type"] = "INLINE_DOCUMENT_CONTENT"

		if v.Data != nil {
			result["data"] = string(v.Data)
		}

		result["page_range"] = v.PageRange
	case oci_ai_document.ObjectStorageLocations:
		result["source_type"] = "OBJECT_STORAGE_LOCATIONS"

		result["data"] = ""
		result["page_range"] = []string{}

		objectLocations := []interface{}{}
		for _, item := range v.ObjectLocations {
			objectLocations = append(objectLocations, ObjectLocationToMap(item))
		}
		result["object_locations"] = objectLocations
	default:
		log.Printf("[WARN] Received 'source_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *AiDocumentProcessorJobResourceCrud) mapToNormalizationFields(fieldKeyFormat string) (oci_ai_document.NormalizationFields, error) {
	result := oci_ai_document.NormalizationFields{}

	//The expected pattern for NormalizationFields here is that there will be only 1 value against normalization_fields i.e. map,
	//only 1 value against map i.e. normalization_type, and only 1 value against normalization_type i.e. some string value.
	if map_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "map")); ok {
		newMap := make(map[string]oci_ai_document.NormalizationFieldsMapValue)
		for k, v := range map_.([]interface{}) {
			fmt.Println("map key: ", k)
			value := v.(map[string]interface{})
			for k, v := range value {
				fmt.Println("map key: ", k)
				value := v.(interface{})
				strVal := fmt.Sprint(value)
				res := oci_ai_document.NormalizationFieldsMapValue{}
				res.NormalizationType = &strVal
				newMap[k] = res
			}
		}
		result.PropertiesMap = newMap
	}
	return result, nil
}

func NormalizationFieldsToMap(obj *oci_ai_document.NormalizationFields) map[string]interface{} {
	result := map[string]interface{}{}

	if obj != nil {
		list := []interface{}{}
		result["map"] = append(list, obj.PropertiesMap)
	}

	return result
}

func (s *AiDocumentProcessorJobResourceCrud) mapToObjectLocation(fieldKeyFormat string) (oci_ai_document.ObjectLocation, error) {
	result := oci_ai_document.ObjectLocation{}

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

	if pageRange, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "page_range")); ok {
		interfaces := pageRange.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "page_range")) {
			result.PageRange = tmp
		}
	}

	return result, nil
}

func ObjectLocationToMap(obj oci_ai_document.ObjectLocation) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BucketName != nil {
		result["bucket"] = obj.BucketName
	}

	if obj.NamespaceName != nil {
		result["namespace"] = obj.NamespaceName
	}

	if obj.ObjectName != nil {
		result["object"] = obj.ObjectName
	}

	result["page_range"] = obj.PageRange

	return result
}

func (s *AiDocumentProcessorJobResourceCrud) mapToOutputLocation(fieldKeyFormat string) (oci_ai_document.OutputLocation, error) {
	result := oci_ai_document.OutputLocation{}

	if bucket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
		tmp := bucket.(string)
		result.BucketName = &tmp
	}

	if namespace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
		tmp := namespace.(string)
		result.NamespaceName = &tmp
	}

	if prefix, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "prefix")); ok {
		tmp := prefix.(string)
		result.Prefix = &tmp
	}

	return result, nil
}

func OutputLocationToMap(obj *oci_ai_document.OutputLocation) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BucketName != nil {
		result["bucket"] = obj.BucketName
	}

	if obj.NamespaceName != nil {
		result["namespace"] = obj.NamespaceName
	}

	if obj.Prefix != nil {
		result["prefix"] = obj.Prefix
	}

	return result
}

func (s *AiDocumentProcessorJobResourceCrud) mapToProcessorConfig(fieldKeyFormat string) (oci_ai_document.ProcessorConfig, error) {
	var baseObject oci_ai_document.ProcessorConfig
	//discriminator
	processorTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "processor_type"))
	var processorType string
	if ok {
		processorType = processorTypeRaw.(string)
	} else {
		processorType = "" // default value
	}
	switch strings.ToLower(processorType) {
	case strings.ToLower("GENERAL"):
		details := oci_ai_document.GeneralProcessorConfig{}
		if documentType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "document_type")); ok {
			details.DocumentType = oci_ai_document.DocumentTypeEnum(documentType.(string))
		}
		if features, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "features")); ok {
			interfaces := features.([]interface{})
			tmp := make([]oci_ai_document.DocumentFeature, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "features"), stateDataIndex)
				converted, err := s.mapToDocumentFeature(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "features")) {
				details.Features = tmp
			}
		}
		if isZipOutputEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_zip_output_enabled")); ok {
			tmp := isZipOutputEnabled.(bool)
			details.IsZipOutputEnabled = &tmp
		}
		if language, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "language")); ok {
			tmp := language.(string)
			details.Language = &tmp
		}

		if normalizationFields, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "normalization_fields")); ok {
			if tmpList := normalizationFields.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "normalization_fields"), 0)
				tmp, err := s.mapToNormalizationFields(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert normalization_fields, encountered error: %v", err)
				}
				fmt.Println("tmp: ", tmp)
				//For this type - GeneralProcessorConfig, there is no field of NormalizationFields in the configuration struct
				//it is part of next switch case InvoiceProcessorConfig only. Hence this is commented out
				//details.NormalizationFields = &tmp
			}
		}

		baseObject = details
	case strings.ToLower("INVOICE"):
		details := oci_ai_document.InvoiceProcessorConfig{}
		if modelId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "model_id")); ok {
			tmp := modelId.(string)
			details.ModelId = &tmp
		}
		if normalizationFields, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "normalization_fields")); ok {
			if tmpList := normalizationFields.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "normalization_fields"), 0)
				tmp, err := s.mapToNormalizationFields(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert normalization_fields, encountered error: %v", err)
				}
				details.NormalizationFields = &tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown processor_type '%v' was specified", processorType)
	}
	return baseObject, nil
}

func ProcessorConfigToMap(obj *oci_ai_document.ProcessorConfig) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_ai_document.GeneralProcessorConfig:
		result["processor_type"] = "GENERAL"

		result["document_type"] = string(v.DocumentType)

		features := []interface{}{}
		for _, item := range v.Features {
			features = append(features, DocumentFeatureToMap(item))
			result["model_id"] = DocumentFeatureToMap(item)["model_id"] //Added model_id in response
		}
		result["features"] = features

		if v.IsZipOutputEnabled != nil {
			result["is_zip_output_enabled"] = bool(*v.IsZipOutputEnabled)
		}

		if v.Language != nil {
			result["language"] = string(*v.Language)
		}

		//normalization_fields computation
		normTypeSampleValMap := make(map[string]string)
		normTypeSampleValMap["normalization_type"] = "normalization_type_sample_val"

		appendedVal := []interface{}{}
		appendedVal = append(appendedVal, normTypeSampleValMap)

		normValMap := make(map[string]interface{})
		normValMap["map"] = appendedVal

		normFields := []interface{}{}
		normFields = append(normFields, normValMap)

		result["normalization_fields"] = normFields

	case oci_ai_document.InvoiceProcessorConfig:
		result["processor_type"] = "INVOICE"

		if v.ModelId != nil {
			result["model_id"] = string(*v.ModelId)
		}

		if v.NormalizationFields != nil {

			res := oci_ai_document.NormalizationFieldsMapValue{}
			propMap := v.NormalizationFields.PropertiesMap
			res = propMap["normalization_type"]

			//normalization_fields computation
			normTypeSampleValMap := make(map[string]string)
			normTypeSampleValMap["normalization_type"] = fmt.Sprint(res.NormalizationType)

			appendedVal := []interface{}{}
			appendedVal = append(appendedVal, normTypeSampleValMap)

			normValMap := make(map[string]interface{})
			normValMap["map"] = appendedVal

			normFields := []interface{}{}
			normFields = append(normFields, normValMap)

			result["normalization_fields"] = normFields
		}
	default:
		log.Printf("[WARN] Received 'processor_type' of unknown type %v", *obj)
		return nil
	}

	return result
}
