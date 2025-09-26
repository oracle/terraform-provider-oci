// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package iot

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	oci_iot "github.com/oracle/oci-go-sdk/v65/iot"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func IotDigitalTwinAdapterResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createIotDigitalTwinAdapterWithContext,
		ReadContext:   readIotDigitalTwinAdapterWithContext,
		UpdateContext: updateIotDigitalTwinAdapterWithContext,
		DeleteContext: deleteIotDigitalTwinAdapterWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"iot_domain_id": {
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
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"digital_twin_model_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"digital_twin_model_spec_uri": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
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
			"inbound_envelope": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"reference_endpoint": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"envelope_mapping": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"time_observed": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"reference_payload": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"data": {
										Type:     schema.TypeMap,
										Required: true,
										Elem:     schema.TypeString,
									},
									"data_format": {
										Type:             schema.TypeString,
										Required:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"JSON",
										}, true),
									},

									// Optional

									// Computed
								},
							},
						},

						// Computed
					},
				},
			},
			"inbound_routes": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"condition": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"description": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"payload_mapping": {
							Type:     schema.TypeMap,
							Optional: true,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"reference_payload": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"data": {
										Type:     schema.TypeMap,
										Required: true,
										Elem:     schema.TypeString,
									},
									"data_format": {
										Type:             schema.TypeString,
										Required:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"JSON",
										}, true),
									},

									// Optional

									// Computed
								},
							},
						},

						// Computed
					},
				},
			},

			// Computed
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
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

func createIotDigitalTwinAdapterWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &IotDigitalTwinAdapterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IotClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readIotDigitalTwinAdapterWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &IotDigitalTwinAdapterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IotClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateIotDigitalTwinAdapterWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &IotDigitalTwinAdapterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IotClient()

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func deleteIotDigitalTwinAdapterWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &IotDigitalTwinAdapterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IotClient()
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type IotDigitalTwinAdapterResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_iot.IotClient
	Res                    *oci_iot.DigitalTwinAdapter
	DisableNotFoundRetries bool
}

func (s *IotDigitalTwinAdapterResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *IotDigitalTwinAdapterResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *IotDigitalTwinAdapterResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_iot.LifecycleStateActive),
	}
}

func (s *IotDigitalTwinAdapterResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *IotDigitalTwinAdapterResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_iot.LifecycleStateDeleted),
	}
}

func (s *IotDigitalTwinAdapterResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_iot.CreateDigitalTwinAdapterRequest{}

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

	if digitalTwinModelId, ok := s.D.GetOkExists("digital_twin_model_id"); ok {
		tmp := digitalTwinModelId.(string)
		request.DigitalTwinModelId = &tmp
	}

	if digitalTwinModelSpecUri, ok := s.D.GetOkExists("digital_twin_model_spec_uri"); ok {
		tmp := digitalTwinModelSpecUri.(string)
		request.DigitalTwinModelSpecUri = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if inboundEnvelope, ok := s.D.GetOkExists("inbound_envelope"); ok {
		if tmpList := inboundEnvelope.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "inbound_envelope", 0)
			tmp, err := s.mapToDigitalTwinAdapterInboundEnvelope(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.InboundEnvelope = &tmp
		}
	}

	if inboundRoutes, ok := s.D.GetOkExists("inbound_routes"); ok {
		interfaces := inboundRoutes.([]interface{})
		tmp := make([]oci_iot.DigitalTwinAdapterInboundRoute, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "inbound_routes", stateDataIndex)
			converted, err := s.mapToDigitalTwinAdapterInboundRoute(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("inbound_routes") {
			request.InboundRoutes = tmp
		}
	}

	if iotDomainId, ok := s.D.GetOkExists("iot_domain_id"); ok {
		tmp := iotDomainId.(string)
		request.IotDomainId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "iot")

	response, err := s.Client.CreateDigitalTwinAdapter(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.DigitalTwinAdapter
	return nil
}

func (s *IotDigitalTwinAdapterResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_iot.GetDigitalTwinAdapterRequest{}

	tmp := s.D.Id()
	request.DigitalTwinAdapterId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "iot")

	response, err := s.Client.GetDigitalTwinAdapter(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.DigitalTwinAdapter
	return nil
}

func (s *IotDigitalTwinAdapterResourceCrud) UpdateWithContext(ctx context.Context) error {
	request := oci_iot.UpdateDigitalTwinAdapterRequest{}

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

	tmp := s.D.Id()
	request.DigitalTwinAdapterId = &tmp

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if inboundEnvelope, ok := s.D.GetOkExists("inbound_envelope"); ok {
		if tmpList := inboundEnvelope.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "inbound_envelope", 0)
			tmp, err := s.mapToDigitalTwinAdapterInboundEnvelope(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.InboundEnvelope = &tmp
		}
	}

	if inboundRoutes, ok := s.D.GetOkExists("inbound_routes"); ok {
		interfaces := inboundRoutes.([]interface{})
		tmp := make([]oci_iot.DigitalTwinAdapterInboundRoute, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "inbound_routes", stateDataIndex)
			converted, err := s.mapToDigitalTwinAdapterInboundRoute(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("inbound_routes") {
			request.InboundRoutes = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "iot")

	response, err := s.Client.UpdateDigitalTwinAdapter(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.DigitalTwinAdapter
	return nil
}

func (s *IotDigitalTwinAdapterResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_iot.DeleteDigitalTwinAdapterRequest{}

	tmp := s.D.Id()
	request.DigitalTwinAdapterId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "iot")

	_, err := s.Client.DeleteDigitalTwinAdapter(ctx, request)
	return err
}

func (s *IotDigitalTwinAdapterResourceCrud) SetData() error {
	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DigitalTwinModelId != nil {
		s.D.Set("digital_twin_model_id", *s.Res.DigitalTwinModelId)
	}

	if s.Res.DigitalTwinModelSpecUri != nil {
		s.D.Set("digital_twin_model_spec_uri", *s.Res.DigitalTwinModelSpecUri)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.InboundEnvelope != nil {
		s.D.Set("inbound_envelope", []interface{}{DigitalTwinAdapterInboundEnvelopeToMap(s.Res.InboundEnvelope)})
	} else {
		s.D.Set("inbound_envelope", nil)
	}

	inboundRoutes := []interface{}{}
	for _, item := range s.Res.InboundRoutes {
		inboundRoutes = append(inboundRoutes, DigitalTwinAdapterInboundRouteToMap(item))
	}
	s.D.Set("inbound_routes", inboundRoutes)

	if s.Res.IotDomainId != nil {
		s.D.Set("iot_domain_id", *s.Res.IotDomainId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func (s *IotDigitalTwinAdapterResourceCrud) mapToDigitalTwinAdapterEnvelopeMapping(fieldKeyFormat string) (oci_iot.DigitalTwinAdapterEnvelopeMapping, error) {
	result := oci_iot.DigitalTwinAdapterEnvelopeMapping{}

	if timeObserved, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_observed")); ok {
		tmp := timeObserved.(string)
		result.TimeObserved = &tmp
	}

	return result, nil
}

func DigitalTwinAdapterEnvelopeMappingToMap(obj *oci_iot.DigitalTwinAdapterEnvelopeMapping) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.TimeObserved != nil {
		result["time_observed"] = string(*obj.TimeObserved)
	}

	return result
}

func (s *IotDigitalTwinAdapterResourceCrud) mapToDigitalTwinAdapterInboundEnvelope(fieldKeyFormat string) (oci_iot.DigitalTwinAdapterInboundEnvelope, error) {
	result := oci_iot.DigitalTwinAdapterInboundEnvelope{}

	if envelopeMapping, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "envelope_mapping")); ok {
		if tmpList := envelopeMapping.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "envelope_mapping"), 0)
			tmp, err := s.mapToDigitalTwinAdapterEnvelopeMapping(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert envelope_mapping, encountered error: %v", err)
			}
			result.EnvelopeMapping = &tmp
		}
	}

	if referenceEndpoint, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "reference_endpoint")); ok {
		tmp := referenceEndpoint.(string)
		result.ReferenceEndpoint = &tmp
	}

	if referencePayload, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "reference_payload")); ok {
		if tmpList := referencePayload.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "reference_payload"), 0)
			tmp, err := s.mapToDigitalTwinAdapterPayload(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert reference_payload, encountered error: %v", err)
			}
			result.ReferencePayload = tmp
		}
	}

	return result, nil
}

func DigitalTwinAdapterInboundEnvelopeToMap(obj *oci_iot.DigitalTwinAdapterInboundEnvelope) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.EnvelopeMapping != nil {
		result["envelope_mapping"] = []interface{}{DigitalTwinAdapterEnvelopeMappingToMap(obj.EnvelopeMapping)}
	}

	if obj.ReferenceEndpoint != nil {
		result["reference_endpoint"] = string(*obj.ReferenceEndpoint)
	}

	if obj.ReferencePayload != nil {
		referencePayloadArray := []interface{}{}
		if referencePayloadMap := DigitalTwinAdapterPayloadToMap(&obj.ReferencePayload); referencePayloadMap != nil {
			referencePayloadArray = append(referencePayloadArray, referencePayloadMap)
		}
		result["reference_payload"] = referencePayloadArray
	}

	return result
}

func (s *IotDigitalTwinAdapterResourceCrud) mapToDigitalTwinAdapterInboundRoute(fieldKeyFormat string) (oci_iot.DigitalTwinAdapterInboundRoute, error) {
	result := oci_iot.DigitalTwinAdapterInboundRoute{}

	if condition, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "condition")); ok {
		tmp := condition.(string)
		result.Condition = &tmp
	}

	if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
		tmp := description.(string)
		result.Description = &tmp
	}

	if payloadMapping, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "payload_mapping")); ok {
		result.PayloadMapping = tfresource.ObjectMapToStringMap(payloadMapping.(map[string]interface{}))
	}

	if referencePayload, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "reference_payload")); ok {
		if tmpList := referencePayload.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "reference_payload"), 0)
			tmp, err := s.mapToDigitalTwinAdapterPayload(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert reference_payload, encountered error: %v", err)
			}
			result.ReferencePayload = tmp
		}
	}

	return result, nil
}

func DigitalTwinAdapterInboundRouteToMap(obj oci_iot.DigitalTwinAdapterInboundRoute) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Condition != nil {
		result["condition"] = string(*obj.Condition)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	result["payload_mapping"] = obj.PayloadMapping

	if obj.ReferencePayload != nil {
		referencePayloadArray := []interface{}{}
		if referencePayloadMap := DigitalTwinAdapterPayloadToMap(&obj.ReferencePayload); referencePayloadMap != nil {
			referencePayloadArray = append(referencePayloadArray, referencePayloadMap)
		}
		result["reference_payload"] = referencePayloadArray
	}

	return result
}

func (s *IotDigitalTwinAdapterResourceCrud) mapToDigitalTwinAdapterPayload(fieldKeyFormat string) (oci_iot.DigitalTwinAdapterPayload, error) {
	var baseObject oci_iot.DigitalTwinAdapterPayload
	//discriminator
	dataFormatRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "data_format"))
	var dataFormat string
	if ok {
		dataFormat = dataFormatRaw.(string)
	} else {
		dataFormat = "JSON" // default value
	}
	switch strings.ToLower(dataFormat) {
	case strings.ToLower("JSON"):
		details := oci_iot.DigitalTwinAdapterJsonPayload{}
		if data, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "data")); ok {
			details.Data = data.(map[string]interface{})
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown data_format '%v' was specified", dataFormat)
	}
	return baseObject, nil
}

func DigitalTwinAdapterPayloadToMap(obj *oci_iot.DigitalTwinAdapterPayload) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_iot.DigitalTwinAdapterJsonPayload:
		result["data_format"] = "JSON"

		result["data"] = tfresource.ObjectMapToStringMap(v.Data)
	default:
		log.Printf("[WARN] Received 'data_format' of unknown type %v", *obj)
		return nil
	}

	return result
}

func DigitalTwinAdapterSummaryToMap(obj oci_iot.DigitalTwinAdapterSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DigitalTwinModelId != nil {
		result["digital_twin_model_id"] = string(*obj.DigitalTwinModelId)
	}

	if obj.DigitalTwinModelSpecUri != nil {
		result["digital_twin_model_spec_uri"] = string(*obj.DigitalTwinModelSpecUri)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IotDomainId != nil {
		result["iot_domain_id"] = string(*obj.IotDomainId)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}
