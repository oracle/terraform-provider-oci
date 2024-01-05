// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package apm_synthetics

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_apm_synthetics "github.com/oracle/oci-go-sdk/v65/apmsynthetics"
)

func ApmSyntheticsScriptResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createApmSyntheticsScript,
		Read:     readApmSyntheticsScript,
		Update:   updateApmSyntheticsScript,
		Delete:   deleteApmSyntheticsScript,
		Schema: map[string]*schema.Schema{
			// Required
			"apm_domain_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"content": {
				Type:     schema.TypeString,
				Required: true,
			},
			"content_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"content_file_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"parameters": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"param_name": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"is_secret": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"param_value": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
						"is_overwritten": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"script_parameter": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"is_secret": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"param_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"param_value": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},

			// Computed
			"content_size_in_bytes": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"monitor_status_count_map": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"disabled": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"enabled": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"invalid": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"total": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_uploaded": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createApmSyntheticsScript(d *schema.ResourceData, m interface{}) error {
	sync := &ApmSyntheticsScriptResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApmSyntheticClient()

	return tfresource.CreateResource(d, sync)
}

func readApmSyntheticsScript(d *schema.ResourceData, m interface{}) error {
	sync := &ApmSyntheticsScriptResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApmSyntheticClient()

	return tfresource.ReadResource(sync)
}

func updateApmSyntheticsScript(d *schema.ResourceData, m interface{}) error {
	sync := &ApmSyntheticsScriptResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApmSyntheticClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteApmSyntheticsScript(d *schema.ResourceData, m interface{}) error {
	sync := &ApmSyntheticsScriptResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApmSyntheticClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type ApmSyntheticsScriptResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_apm_synthetics.ApmSyntheticClient
	Res                    *oci_apm_synthetics.Script
	DisableNotFoundRetries bool
}

func (s *ApmSyntheticsScriptResourceCrud) ID() string {
	return GetScriptCompositeId(*s.Res.Id, s.D.Get("apm_domain_id").(string))
}

func (s *ApmSyntheticsScriptResourceCrud) Create() error {
	request := oci_apm_synthetics.CreateScriptRequest{}

	if apmDomainId, ok := s.D.GetOkExists("apm_domain_id"); ok {
		tmp := apmDomainId.(string)
		request.ApmDomainId = &tmp
	}

	if content, ok := s.D.GetOkExists("content"); ok {
		tmp := content.(string)
		request.Content = &tmp
	}

	if contentFileName, ok := s.D.GetOkExists("content_file_name"); ok {
		tmp := contentFileName.(string)
		request.ContentFileName = &tmp
	}

	if contentType, ok := s.D.GetOkExists("content_type"); ok {
		request.ContentType = oci_apm_synthetics.ContentTypesEnum(contentType.(string))
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
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
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if parameters, ok := s.D.GetOkExists("parameters"); ok {
		interfaces := parameters.([]interface{})
		tmp := make([]oci_apm_synthetics.ScriptParameter, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "parameters", stateDataIndex)
			converted, err := s.mapToScriptParameter(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("parameters") {
			request.Parameters = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apm_synthetics")

	response, err := s.Client.CreateScript(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Script
	return nil
}

func (s *ApmSyntheticsScriptResourceCrud) Get() error {
	request := oci_apm_synthetics.GetScriptRequest{}

	scriptId, apmDomainId, err := parseScriptCompositeId(s.D.Id())
	if err == nil {
		request.ScriptId = &scriptId
		request.ApmDomainId = &apmDomainId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apm_synthetics")

	response, err := s.Client.GetScript(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Script
	return nil
}

func (s *ApmSyntheticsScriptResourceCrud) Update() error {
	request := oci_apm_synthetics.UpdateScriptRequest{}

	if apmDomainId, ok := s.D.GetOkExists("apm_domain_id"); ok {
		tmp := apmDomainId.(string)
		request.ApmDomainId = &tmp
	}

	if content, ok := s.D.GetOkExists("content"); ok {
		tmp := content.(string)
		request.Content = &tmp
	}

	if contentFileName, ok := s.D.GetOkExists("content_file_name"); ok {
		tmp := contentFileName.(string)
		request.ContentFileName = &tmp
	}

	if contentType, ok := s.D.GetOkExists("content_type"); ok {
		request.ContentType = oci_apm_synthetics.ContentTypesEnum(contentType.(string))
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
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
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if parameters, ok := s.D.GetOkExists("parameters"); ok {
		interfaces := parameters.([]interface{})
		tmp := make([]oci_apm_synthetics.ScriptParameter, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "parameters", stateDataIndex)
			converted, err := s.mapToScriptParameter(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("parameters") {
			request.Parameters = tmp
		}
	}

	tmp := s.D.Id()
	request.ScriptId = &tmp
	scriptId, apmDomainId, err := parseScriptCompositeId(s.D.Id())
	if err == nil {
		request.ScriptId = &scriptId
		request.ApmDomainId = &apmDomainId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apm_synthetics")

	response, err := s.Client.UpdateScript(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Script
	return nil
}

func (s *ApmSyntheticsScriptResourceCrud) Delete() error {
	request := oci_apm_synthetics.DeleteScriptRequest{}

	if apmDomainId, ok := s.D.GetOkExists("apm_domain_id"); ok {
		tmp := apmDomainId.(string)
		request.ApmDomainId = &tmp
	}

	tmp := s.D.Id()
	request.ScriptId = &tmp

	if tmp != "" {
		scriptId, apmDomainId, err := parseScriptCompositeId(s.D.Id())
		if err == nil {
			request.ScriptId = &scriptId
			request.ApmDomainId = &apmDomainId
		} else {
			log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
		}
	}
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apm_synthetics")

	_, err := s.Client.DeleteScript(context.Background(), request)
	return err
}

func (s *ApmSyntheticsScriptResourceCrud) SetData() error {

	if s.Res.Content != nil {
		s.D.Set("content", *s.Res.Content)
	}

	if s.Res.ContentFileName != nil {
		s.D.Set("content_file_name", *s.Res.ContentFileName)
	}

	if s.Res.ContentSizeInBytes != nil {
		s.D.Set("content_size_in_bytes", *s.Res.ContentSizeInBytes)
	}

	s.D.Set("content_type", s.Res.ContentType)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.MonitorStatusCountMap != nil {
		s.D.Set("monitor_status_count_map", []interface{}{MonitorStatusCountMapToMap(s.Res.MonitorStatusCountMap)})
	} else {
		s.D.Set("monitor_status_count_map", nil)
	}

	parameters := []interface{}{}
	for _, item := range s.Res.Parameters {
		parameters = append(parameters, ScriptParameterInfoToMap(item))
	}
	s.D.Set("parameters", parameters)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.TimeUploaded != nil {
		s.D.Set("time_uploaded", s.Res.TimeUploaded.String())
	}

	return nil
}

func GetScriptCompositeId(scriptId string, apmDomainId string) string {
	apmDomainId = url.PathEscape(apmDomainId)
	scriptId = url.PathEscape(scriptId)
	compositeId := "scripts/" + scriptId + "/apmDomainId/" + apmDomainId
	return compositeId
}

func parseScriptCompositeId(compositeId string) (scriptId string, apmDomainId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("scripts/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	scriptId, _ = url.PathUnescape(parts[1])
	apmDomainId, _ = url.PathUnescape(parts[3])

	return
}

func MonitorStatusCountMapToMap(obj *oci_apm_synthetics.MonitorStatusCountMap) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Disabled != nil {
		result["disabled"] = int(*obj.Disabled)
	}

	if obj.Enabled != nil {
		result["enabled"] = int(*obj.Enabled)
	}

	if obj.Invalid != nil {
		result["invalid"] = int(*obj.Invalid)
	}

	if obj.Total != nil {
		result["total"] = int(*obj.Total)
	}

	return result
}

func ScriptParameterToMap(obj *oci_apm_synthetics.ScriptParameter) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsSecret != nil {
		result["is_secret"] = bool(*obj.IsSecret)
	}

	if obj.ParamName != nil {
		result["param_name"] = string(*obj.ParamName)
	}

	if obj.ParamValue != nil {
		result["param_value"] = string(*obj.ParamValue)
	}

	return result
}

func ScriptSummaryToMap(obj oci_apm_synthetics.ScriptSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["content_type"] = string(obj.ContentType)

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

	if obj.MonitorStatusCountMap != nil {
		result["monitor_status_count_map"] = []interface{}{MonitorStatusCountMapToMap(obj.MonitorStatusCountMap)}
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func (s *ApmSyntheticsScriptResourceCrud) mapToScriptParameter(fieldKeyFormat string) (oci_apm_synthetics.ScriptParameter, error) {
	result := oci_apm_synthetics.ScriptParameter{}

	if paramName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "param_name")); ok {
		tmp := paramName.(string)
		result.ParamName = &tmp
	}

	if paramValue, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "param_value")); ok {
		tmp := paramValue.(string)
		result.ParamValue = &tmp
	}

	if isSecret, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_secret")); ok {
		tmp := isSecret.(bool)
		result.IsSecret = &tmp
	}

	return result, nil
}

func ScriptParameterInfoToMap(obj oci_apm_synthetics.ScriptParameterInfo) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ScriptParameter != nil {
		ScriptParameters := []interface{}{}

		ScriptParameters = append(ScriptParameters, ScriptParameterToMap(obj.ScriptParameter))

		result["script_parameter"] = ScriptParameters
		result["is_secret"] = bool(*obj.ScriptParameter.IsSecret)
		result["param_name"] = string(*obj.ScriptParameter.ParamName)
		if obj.ScriptParameter.ParamValue != nil {
			result["param_value"] = string(*obj.ScriptParameter.ParamValue)
		}

	}

	if obj.IsOverwritten != nil {
		result["is_overwritten"] = obj.IsOverwritten
	}

	return result
}
