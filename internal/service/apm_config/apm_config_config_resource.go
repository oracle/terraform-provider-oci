// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package apm_config

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_apm_config "github.com/oracle/oci-go-sdk/v65/apmconfig"
)

func ApmConfigConfigResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createApmConfigConfig,
		Read:     readApmConfigConfig,
		Update:   updateApmConfigConfig,
		Delete:   deleteApmConfigConfig,
		Schema: map[string]*schema.Schema{
			// Required
			"apm_domain_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"config_type": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"AGENT",
					"APDEX",
					"MACS_APM_EXTENSION",
					"METRIC_GROUP",
					"OPTIONS",
					"SPAN_FILTER",
				}, true),
			},

			// Optional
			"agent_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attach_install_dir": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"config_map": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"file_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Optional
									"body": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"content_type": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},

						// Computed
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
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dimensions": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"value_source": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"filter_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"filter_text": {
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
			"group": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"management_agent_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"match_agents_with_attribute_value": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"metrics": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"description": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"unit": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"value_source": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"namespace": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"opc_dry_run": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"options": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.JsonStringDiffSuppressFunction,
			},
			"overrides": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"override_list": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"agent_filter": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"override_map": {
										Type:     schema.TypeMap,
										Optional: true,
										Computed: true,
										Elem:     schema.TypeString,
									},

									// Computed
								},
							},
						},

						// Computed
					},
				},
			},
			"process_filter": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"rules": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"display_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"filter_text": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"is_apply_to_error_spans": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"is_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"priority": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"satisfied_response_time": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"tolerating_response_time": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"run_as_user": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"service_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"created_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"in_use_by": {
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"config_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"options_group": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"match_agents_with_attribute_key": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
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
			"updated_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createApmConfigConfig(d *schema.ResourceData, m interface{}) error {
	sync := &ApmConfigConfigResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ConfigClient()

	return tfresource.CreateResource(d, sync)
}

func readApmConfigConfig(d *schema.ResourceData, m interface{}) error {
	sync := &ApmConfigConfigResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ConfigClient()

	return tfresource.ReadResource(sync)
}

func updateApmConfigConfig(d *schema.ResourceData, m interface{}) error {
	sync := &ApmConfigConfigResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ConfigClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteApmConfigConfig(d *schema.ResourceData, m interface{}) error {
	sync := &ApmConfigConfigResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ConfigClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type ApmConfigConfigResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_apm_config.ConfigClient
	Res                    *oci_apm_config.Config
	DisableNotFoundRetries bool
}

func (s *ApmConfigConfigResourceCrud) ID() string {
	config := *s.Res
	return GetConfigCompositeId(*config.GetId(), s.D.Get("apm_domain_id").(string))
}

func (s *ApmConfigConfigResourceCrud) Create() error {
	request := oci_apm_config.CreateConfigRequest{}
	err := s.populateTopLevelPolymorphicCreateConfigRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apm_config")

	response, err := s.Client.CreateConfig(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Config
	return nil
}

func (s *ApmConfigConfigResourceCrud) Get() error {
	request := oci_apm_config.GetConfigRequest{}

	configId, apmDomainId, err := parseConfigCompositeId(s.D.Id())
	if err == nil {
		request.ConfigId = &configId
		request.ApmDomainId = &apmDomainId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apm_config")

	response, err := s.Client.GetConfig(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Config
	return nil
}

func (s *ApmConfigConfigResourceCrud) Update() error {
	request := oci_apm_config.UpdateConfigRequest{}
	err := s.populateTopLevelPolymorphicUpdateConfigRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apm_config")

	response, err := s.Client.UpdateConfig(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Config
	return nil
}

func (s *ApmConfigConfigResourceCrud) Delete() error {
	request := oci_apm_config.DeleteConfigRequest{}

	if tmp := s.D.Id(); tmp != "" {
		configId, apmDomainId, err := parseConfigCompositeId(s.D.Id())
		if err == nil {
			request.ConfigId = &configId
			request.ApmDomainId = &apmDomainId
		} else {
			log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apm_config")

	_, err := s.Client.DeleteConfig(context.Background(), request)
	return err
}

func (s *ApmConfigConfigResourceCrud) SetData() error {

	s.D.Set("process_filter", nil)
	s.D.Set("match_agents_with_attribute_key", nil)

	switch v := (*s.Res).(type) {
	case oci_apm_config.AgentConfig:
		s.D.Set("config_type", "AGENT")

		if v.Config != nil {
			s.D.Set("config", []interface{}{AgentConfigMapToMap(v.Config)})
		} else {
			s.D.Set("config", nil)
		}

		s.D.Set("match_agents_with_attribute_key", v.MatchAgentsWithAttributeKey)

		if v.MatchAgentsWithAttributeValue != nil {
			s.D.Set("match_agents_with_attribute_value", *v.MatchAgentsWithAttributeValue)
		}

		if v.Overrides != nil {
			s.D.Set("overrides", []interface{}{AgentConfigOverridesToMap(v.Overrides)})
		} else {
			s.D.Set("overrides", nil)
		}

		if v.CreatedBy != nil {
			s.D.Set("created_by", *v.CreatedBy)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Etag != nil {
			s.D.Set("etag", *v.Etag)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.UpdatedBy != nil {
			s.D.Set("updated_by", *v.UpdatedBy)
		}
	case oci_apm_config.ApdexRules:
		s.D.Set("config_type", "APDEX")

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		rules := []interface{}{}
		for _, item := range v.Rules {
			rules = append(rules, ApdexToMap(item))
		}
		s.D.Set("rules", rules)

		if v.CreatedBy != nil {
			s.D.Set("created_by", *v.CreatedBy)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.Etag != nil {
			s.D.Set("etag", *v.Etag)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.UpdatedBy != nil {
			s.D.Set("updated_by", *v.UpdatedBy)
		}
	case oci_apm_config.MacsApmExtension:
		s.D.Set("config_type", "MACS_APM_EXTENSION")

		if v.AgentVersion != nil {
			s.D.Set("agent_version", *v.AgentVersion)
		}

		if v.AttachInstallDir != nil {
			s.D.Set("attach_install_dir", *v.AttachInstallDir)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.ManagementAgentId != nil {
			s.D.Set("management_agent_id", *v.ManagementAgentId)
		}

		s.D.Set("process_filter", v.ProcessFilter)

		if v.RunAsUser != nil {
			s.D.Set("run_as_user", *v.RunAsUser)
		}

		if v.ServiceName != nil {
			s.D.Set("service_name", *v.ServiceName)
		}

		if v.CreatedBy != nil {
			s.D.Set("created_by", *v.CreatedBy)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Etag != nil {
			s.D.Set("etag", *v.Etag)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.UpdatedBy != nil {
			s.D.Set("updated_by", *v.UpdatedBy)
		}
	case oci_apm_config.MetricGroup:
		s.D.Set("config_type", "METRIC_GROUP")

		dimensions := []interface{}{}
		for _, item := range v.Dimensions {
			dimensions = append(dimensions, DimensionToMap(item))
		}
		s.D.Set("dimensions", dimensions)

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.FilterId != nil {
			s.D.Set("filter_id", *v.FilterId)
		}

		metrics := []interface{}{}
		for _, item := range v.Metrics {
			metrics = append(metrics, MetricToMap(item))
		}
		s.D.Set("metrics", metrics)

		if v.Namespace != nil {
			s.D.Set("namespace", *v.Namespace)
		}

		if v.CreatedBy != nil {
			s.D.Set("created_by", *v.CreatedBy)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.Etag != nil {
			s.D.Set("etag", *v.Etag)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.UpdatedBy != nil {
			s.D.Set("updated_by", *v.UpdatedBy)
		}
	case oci_apm_config.Options:
		s.D.Set("config_type", "OPTIONS")

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.Group != nil {
			s.D.Set("group", *v.Group)
		}

		if v.Options != nil {
			s.D.Set("options", optionsToMap(v.Options))
		} else {
			s.D.Set("options", nil)
		}

		if v.CreatedBy != nil {
			s.D.Set("created_by", *v.CreatedBy)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.Etag != nil {
			s.D.Set("etag", *v.Etag)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.UpdatedBy != nil {
			s.D.Set("updated_by", *v.UpdatedBy)
		}
	case oci_apm_config.SpanFilter:
		s.D.Set("config_type", "SPAN_FILTER")

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.FilterText != nil {
			s.D.Set("filter_text", *v.FilterText)
		}

		inUseBy := []interface{}{}
		for _, item := range v.InUseBy {
			inUseBy = append(inUseBy, SpanFilterReferenceToMap(item))
		}
		s.D.Set("in_use_by", inUseBy)

		if v.CreatedBy != nil {
			s.D.Set("created_by", *v.CreatedBy)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.Etag != nil {
			s.D.Set("etag", *v.Etag)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.UpdatedBy != nil {
			s.D.Set("updated_by", *v.UpdatedBy)
		}
	default:
		log.Printf("[WARN] Received 'config_type' of unknown type %v", *s.Res)
		return nil
	}
	return nil
}

func GetConfigCompositeId(configId string, apmDomainId string) string {
	apmDomainId = url.PathEscape(apmDomainId)
	configId = url.PathEscape(configId)
	compositeId := "configs/" + configId + "/apmDomainId/" + apmDomainId
	return compositeId
}

func parseConfigCompositeId(compositeId string) (configId string, apmDomainId string, err error) {
	parts := strings.Split(compositeId, "/")

	match, _ := regexp.MatchString("configs/.*/apmDomainId/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	configId, _ = url.PathUnescape(parts[1])
	apmDomainId, _ = url.PathUnescape(parts[3])

	return
}

func (s *ApmConfigConfigResourceCrud) mapToAgentConfigMap(fieldKeyFormat string) (oci_apm_config.AgentConfigMap, error) {
	result := oci_apm_config.AgentConfigMap{}
	result.ConfigMap = make(map[string]oci_apm_config.AgentConfigFile)

	if configMap, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "config_map")); ok {
		for _, row := range configMap.([]interface{}) {
			rowMap := tfresource.ObjectMapToStringMap(row.(map[string]interface{}))
			rowFileName := rowMap["file_name"]
			rowContentType := rowMap["content_type"]
			rowBody := rowMap["body"]
			result.ConfigMap[rowFileName] = oci_apm_config.AgentConfigFile{ContentType: &rowContentType, Body: &rowBody}
		}
	}

	return result, nil
}

func AgentConfigMapToMap(obj *oci_apm_config.AgentConfigMap) map[string]interface{} {
	result := map[string]interface{}{}
	var resultConfigMap []interface{}

	configMap := obj.ConfigMap

	for fileName := range configMap {
		body := configMap[fileName].Body
		contentType := configMap[fileName].ContentType
		temp := map[string]string{
			"file_name":    fileName,
			"body":         *body,
			"content_type": *contentType,
		}

		resultConfigMap = append(resultConfigMap, temp)
	}

	result["config_map"] = resultConfigMap
	return result
}

func (s *ApmConfigConfigResourceCrud) mapToAgentConfigOverride(fieldKeyFormat string) (oci_apm_config.AgentConfigOverride, error) {
	result := oci_apm_config.AgentConfigOverride{}

	if agentFilter, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "agent_filter")); ok {
		tmp := agentFilter.(string)
		result.AgentFilter = &tmp
	}

	if overrideMap, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "override_map")); ok {
		result.OverrideMap = tfresource.ObjectMapToStringMap(overrideMap.(map[string]interface{}))
	}

	return result, nil
}

func AgentConfigOverrideToMap(obj oci_apm_config.AgentConfigOverride) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AgentFilter != nil {
		result["agent_filter"] = string(*obj.AgentFilter)
	}

	result["override_map"] = obj.OverrideMap

	return result
}

func (s *ApmConfigConfigResourceCrud) mapToAgentConfigOverrides(fieldKeyFormat string) (oci_apm_config.AgentConfigOverrides, error) {
	result := oci_apm_config.AgentConfigOverrides{}

	if overrideList, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "override_list")); ok {
		interfaces := overrideList.([]interface{})
		tmp := make([]oci_apm_config.AgentConfigOverride, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "override_list"), stateDataIndex)
			converted, err := s.mapToAgentConfigOverride(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "override_list")) {
			result.OverrideList = tmp
		}
	}

	return result, nil
}

func AgentConfigOverridesToMap(obj *oci_apm_config.AgentConfigOverrides) map[string]interface{} {
	result := map[string]interface{}{}

	overrideList := []interface{}{}
	for _, item := range obj.OverrideList {
		overrideList = append(overrideList, AgentConfigOverrideToMap(item))
	}
	result["override_list"] = overrideList

	return result
}

func (s *ApmConfigConfigResourceCrud) mapToApdex(fieldKeyFormat string) (oci_apm_config.Apdex, error) {
	result := oci_apm_config.Apdex{}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	if filterText, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "filter_text")); ok {
		tmp := filterText.(string)
		result.FilterText = &tmp
	}

	if isApplyToErrorSpans, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_apply_to_error_spans")); ok {
		tmp := isApplyToErrorSpans.(bool)
		result.IsApplyToErrorSpans = &tmp
	}

	if isEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_enabled")); ok {
		tmp := isEnabled.(bool)
		result.IsEnabled = &tmp
	}

	if priority, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "priority")); ok {
		tmp := priority.(int)
		result.Priority = &tmp
	}

	if satisfiedResponseTime, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "satisfied_response_time")); ok {
		tmp := satisfiedResponseTime.(int)
		result.SatisfiedResponseTime = &tmp
	}

	if toleratingResponseTime, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tolerating_response_time")); ok {
		tmp := toleratingResponseTime.(int)
		result.ToleratingResponseTime = &tmp
	}

	return result, nil
}

func ApdexToMap(obj oci_apm_config.Apdex) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.FilterText != nil {
		result["filter_text"] = string(*obj.FilterText)
	}

	if obj.IsApplyToErrorSpans != nil {
		result["is_apply_to_error_spans"] = bool(*obj.IsApplyToErrorSpans)
	}

	if obj.IsEnabled != nil {
		result["is_enabled"] = bool(*obj.IsEnabled)
	}

	if obj.Priority != nil {
		result["priority"] = int(*obj.Priority)
	}

	if obj.SatisfiedResponseTime != nil {
		result["satisfied_response_time"] = int(*obj.SatisfiedResponseTime)
	}

	if obj.ToleratingResponseTime != nil {
		result["tolerating_response_time"] = int(*obj.ToleratingResponseTime)
	}

	return result
}

func ConfigSummaryToMap(obj oci_apm_config.ConfigSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.GetId() != nil {
		result["id"] = *obj.GetId()
	}

	if obj.GetDefinedTags() != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.GetDefinedTags())
	}

	result["freeform_tags"] = obj.GetFreeformTags()

	if obj.GetTimeCreated() != nil {
		result["time_created"] = obj.GetTimeCreated().String()
	}

	if obj.GetTimeUpdated() != nil {
		result["time_updated"] = obj.GetTimeUpdated().String()
	}

	switch v := (obj).(type) {
	case oci_apm_config.AgentConfigSummary:
		result["config_type"] = "AGENT"

		if v.Config != nil {
			result["config"] = []interface{}{AgentConfigMapToMap(v.Config)}
		}

		result["match_agents_with_attribute_key"] = v.MatchAgentsWithAttributeKey

		if v.MatchAgentsWithAttributeValue != nil {
			result["match_agents_with_attribute_value"] = string(*v.MatchAgentsWithAttributeValue)
		}

		if v.Overrides != nil {
			result["overrides"] = []interface{}{AgentConfigOverridesToMap(v.Overrides)}
		}
	case oci_apm_config.ApdexRulesSummary:
		result["config_type"] = "APDEX"

		if v.DisplayName != nil {
			result["display_name"] = string(*v.DisplayName)
		}

		rules := []interface{}{}
		for _, item := range v.Rules {
			rules = append(rules, ApdexToMap(item))
		}
		result["rules"] = rules
	case oci_apm_config.MacsApmExtensionSummary:
		result["config_type"] = "MACS_APM_EXTENSION"

		if v.AgentVersion != nil {
			result["agent_version"] = string(*v.AgentVersion)
		}

		if v.AttachInstallDir != nil {
			result["attach_install_dir"] = string(*v.AttachInstallDir)
		}

		if v.DisplayName != nil {
			result["display_name"] = string(*v.DisplayName)
		}

		if v.ManagementAgentId != nil {
			result["management_agent_id"] = string(*v.ManagementAgentId)
		}

		result["process_filter"] = v.ProcessFilter

		if v.RunAsUser != nil {
			result["run_as_user"] = string(*v.RunAsUser)
		}

		if v.ServiceName != nil {
			result["service_name"] = string(*v.ServiceName)
		}
	case oci_apm_config.MetricGroupSummary:
		result["config_type"] = "METRIC_GROUP"

		dimensions := []interface{}{}
		for _, item := range v.Dimensions {
			dimensions = append(dimensions, DimensionToMap(item))
		}
		result["dimensions"] = dimensions

		if v.DisplayName != nil {
			result["display_name"] = string(*v.DisplayName)
		}

		if v.FilterId != nil {
			result["filter_id"] = string(*v.FilterId)
		}

		metrics := []interface{}{}
		for _, item := range v.Metrics {
			metrics = append(metrics, MetricToMap(item))
		}
		result["metrics"] = metrics

		if v.Namespace != nil {
			result["namespace"] = string(*v.Namespace)
		}
	case oci_apm_config.OptionsSummary:
		result["config_type"] = "OPTIONS"

		if v.Description != nil {
			result["description"] = string(*v.Description)
		}

		if v.DisplayName != nil {
			result["display_name"] = string(*v.DisplayName)
		}

		if v.Group != nil {
			result["group"] = string(*v.Group)
		}

		if v.Options != nil {
			result["options"] = optionsToMap(v.Options)
		}
	case oci_apm_config.SpanFilterSummary:
		result["config_type"] = "SPAN_FILTER"

		if v.Description != nil {
			result["description"] = string(*v.Description)
		}

		if v.DisplayName != nil {
			result["display_name"] = string(*v.DisplayName)
		}

		if v.FilterText != nil {
			result["filter_text"] = string(*v.FilterText)
		}

		inUseBy := []interface{}{}
		for _, item := range v.InUseBy {
			inUseBy = append(inUseBy, SpanFilterReferenceToMap(item))
		}
		result["in_use_by"] = inUseBy
	default:
		log.Printf("[WARN] Received 'config_type' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *ApmConfigConfigResourceCrud) mapToDimension(fieldKeyFormat string) (oci_apm_config.Dimension, error) {
	result := oci_apm_config.Dimension{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if valueSource, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value_source")); ok {
		tmp := valueSource.(string)
		result.ValueSource = &tmp
	}

	return result, nil
}

func DimensionToMap(obj oci_apm_config.Dimension) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.ValueSource != nil {
		result["value_source"] = string(*obj.ValueSource)
	}

	return result
}

func (s *ApmConfigConfigResourceCrud) mapToMetric(fieldKeyFormat string) (oci_apm_config.Metric, error) {
	result := oci_apm_config.Metric{}

	if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
		tmp := description.(string)
		result.Description = &tmp
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if unit, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "unit")); ok {
		tmp := unit.(string)
		result.Unit = &tmp
	}

	if valueSource, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value_source")); ok {
		tmp := valueSource.(string)
		result.ValueSource = &tmp
	}

	return result, nil
}

func MetricToMap(obj oci_apm_config.Metric) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Unit != nil {
		result["unit"] = string(*obj.Unit)
	}

	if obj.ValueSource != nil {
		result["value_source"] = string(*obj.ValueSource)
	}

	return result
}

func SpanFilterReferenceToMap(obj oci_apm_config.SpanFilterReference) map[string]interface{} {
	result := map[string]interface{}{}

	result["config_type"] = string(obj.ConfigType)

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.OptionsGroup != nil {
		result["options_group"] = string(*obj.OptionsGroup)
	}

	return result
}

func mapToOptions(options string) (*interface{}, error) {
	var result interface{}
	var err error

	var obj interface{}
	err = json.Unmarshal([]byte(options), &obj)
	result = &obj

	return &result, err
}

func optionsToMap(obj *interface{}) string {
	var result string

	if obj != nil {
		var bytes, _ = json.Marshal(obj)
		result = string(bytes)
	}

	return result
}

func (s *ApmConfigConfigResourceCrud) populateTopLevelPolymorphicCreateConfigRequest(request *oci_apm_config.CreateConfigRequest) error {
	//discriminator
	configTypeRaw, ok := s.D.GetOkExists("config_type")
	var configType string
	if ok {
		configType = configTypeRaw.(string)
	} else {
		configType = "" // default value
	}
	switch strings.ToLower(configType) {
	case strings.ToLower("AGENT"):
		details := oci_apm_config.CreateAgentConfigDetails{}
		if config, ok := s.D.GetOkExists("config"); ok {
			if tmpList := config.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "config", 0)
				tmp, err := s.mapToAgentConfigMap(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.Config = &tmp
			}
		}
		if matchAgentsWithAttributeValue, ok := s.D.GetOkExists("match_agents_with_attribute_value"); ok {
			tmp := matchAgentsWithAttributeValue.(string)
			details.MatchAgentsWithAttributeValue = &tmp
		}
		if overrides, ok := s.D.GetOkExists("overrides"); ok {
			if tmpList := overrides.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "overrides", 0)
				tmp, err := s.mapToAgentConfigOverrides(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.Overrides = &tmp
			}
		}
		if apmDomainId, ok := s.D.GetOkExists("apm_domain_id"); ok {
			tmp := apmDomainId.(string)
			request.ApmDomainId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if opcDryRun, ok := s.D.GetOkExists("opc_dry_run"); ok {
			tmp := opcDryRun.(string)
			request.OpcDryRun = &tmp
		}
		request.CreateConfigDetails = details
	case strings.ToLower("APDEX"):
		details := oci_apm_config.CreateApdexRulesDetails{}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if rules, ok := s.D.GetOkExists("rules"); ok {
			interfaces := rules.([]interface{})
			tmp := make([]oci_apm_config.Apdex, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "rules", stateDataIndex)
				converted, err := s.mapToApdex(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("rules") {
				details.Rules = tmp
			}
		}
		if apmDomainId, ok := s.D.GetOkExists("apm_domain_id"); ok {
			tmp := apmDomainId.(string)
			request.ApmDomainId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if opcDryRun, ok := s.D.GetOkExists("opc_dry_run"); ok {
			tmp := opcDryRun.(string)
			request.OpcDryRun = &tmp
		}
		request.CreateConfigDetails = details
	case strings.ToLower("MACS_APM_EXTENSION"):
		details := oci_apm_config.CreateMacsApmExtensionDetails{}
		if agentVersion, ok := s.D.GetOkExists("agent_version"); ok {
			tmp := agentVersion.(string)
			details.AgentVersion = &tmp
		}
		if attachInstallDir, ok := s.D.GetOkExists("attach_install_dir"); ok {
			tmp := attachInstallDir.(string)
			details.AttachInstallDir = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if managementAgentId, ok := s.D.GetOkExists("management_agent_id"); ok {
			tmp := managementAgentId.(string)
			details.ManagementAgentId = &tmp
		}
		if processFilter, ok := s.D.GetOkExists("process_filter"); ok {
			interfaces := processFilter.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("process_filter") {
				details.ProcessFilter = tmp
			}
		}
		if runAsUser, ok := s.D.GetOkExists("run_as_user"); ok {
			tmp := runAsUser.(string)
			details.RunAsUser = &tmp
		}
		if serviceName, ok := s.D.GetOkExists("service_name"); ok {
			tmp := serviceName.(string)
			details.ServiceName = &tmp
		}
		if apmDomainId, ok := s.D.GetOkExists("apm_domain_id"); ok {
			tmp := apmDomainId.(string)
			request.ApmDomainId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if opcDryRun, ok := s.D.GetOkExists("opc_dry_run"); ok {
			tmp := opcDryRun.(string)
			request.OpcDryRun = &tmp
		}
		request.CreateConfigDetails = details
	case strings.ToLower("METRIC_GROUP"):
		details := oci_apm_config.CreateMetricGroupDetails{}
		if dimensions, ok := s.D.GetOkExists("dimensions"); ok {
			interfaces := dimensions.([]interface{})
			tmp := make([]oci_apm_config.Dimension, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "dimensions", stateDataIndex)
				converted, err := s.mapToDimension(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("dimensions") {
				details.Dimensions = tmp
			}
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if filterId, ok := s.D.GetOkExists("filter_id"); ok {
			tmp := filterId.(string)
			details.FilterId = &tmp
		}
		if metrics, ok := s.D.GetOkExists("metrics"); ok {
			interfaces := metrics.([]interface{})
			tmp := make([]oci_apm_config.Metric, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "metrics", stateDataIndex)
				converted, err := s.mapToMetric(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("metrics") {
				details.Metrics = tmp
			}
		}
		if namespace, ok := s.D.GetOkExists("namespace"); ok {
			tmp := namespace.(string)
			details.Namespace = &tmp
		}
		if apmDomainId, ok := s.D.GetOkExists("apm_domain_id"); ok {
			tmp := apmDomainId.(string)
			request.ApmDomainId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if opcDryRun, ok := s.D.GetOkExists("opc_dry_run"); ok {
			tmp := opcDryRun.(string)
			request.OpcDryRun = &tmp
		}
		request.CreateConfigDetails = details
	case strings.ToLower("OPTIONS"):
		details := oci_apm_config.CreateOptionsDetails{}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if group, ok := s.D.GetOkExists("group"); ok {
			tmp := group.(string)
			details.Group = &tmp
		}
		if options, ok := s.D.GetOkExists("options"); ok {
			tmp, err := mapToOptions(options.(string))
			if err != nil {
				return err
			}
			details.Options = tmp
		}
		/*if options, ok := s.D.GetOkExists("options"); ok {
			if tmpList := options.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "options", 0)
				tmp, err := s.mapToOptions(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.Options = tmp
			}
		}*/
		if apmDomainId, ok := s.D.GetOkExists("apm_domain_id"); ok {
			tmp := apmDomainId.(string)
			request.ApmDomainId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if opcDryRun, ok := s.D.GetOkExists("opc_dry_run"); ok {
			tmp := opcDryRun.(string)
			request.OpcDryRun = &tmp
		}
		request.CreateConfigDetails = details
	case strings.ToLower("SPAN_FILTER"):
		details := oci_apm_config.CreateSpanFilterDetails{}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if filterText, ok := s.D.GetOkExists("filter_text"); ok {
			tmp := filterText.(string)
			details.FilterText = &tmp
		}
		if apmDomainId, ok := s.D.GetOkExists("apm_domain_id"); ok {
			tmp := apmDomainId.(string)
			request.ApmDomainId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if opcDryRun, ok := s.D.GetOkExists("opc_dry_run"); ok {
			tmp := opcDryRun.(string)
			request.OpcDryRun = &tmp
		}
		request.CreateConfigDetails = details
	default:
		return fmt.Errorf("unknown config_type '%v' was specified", configType)
	}
	return nil
}

func (s *ApmConfigConfigResourceCrud) populateTopLevelPolymorphicUpdateConfigRequest(request *oci_apm_config.UpdateConfigRequest) error {
	//discriminator
	configTypeRaw, ok := s.D.GetOkExists("config_type")
	var configType string
	if ok {
		configType = configTypeRaw.(string)
	} else {
		configType = "" // default value
	}
	configId, apmDomainId, err := parseConfigCompositeId(s.D.Id())
	if err == nil {
		request.ConfigId = &configId
		request.ApmDomainId = &apmDomainId
	} else {
		log.Printf("[WARN] populateTopLevelPolymorphicUpdateConfigRequest() unable to parse current ID: %s", s.D.Id())
	}
	switch strings.ToLower(configType) {
	case strings.ToLower("AGENT"):
		details := oci_apm_config.UpdateAgentConfigDetails{}
		if config, ok := s.D.GetOkExists("config"); ok {
			if tmpList := config.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "config", 0)
				tmp, err := s.mapToAgentConfigMap(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.Config = &tmp
			}
		}
		if overrides, ok := s.D.GetOkExists("overrides"); ok {
			if tmpList := overrides.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "overrides", 0)
				tmp, err := s.mapToAgentConfigOverrides(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.Overrides = &tmp
			}
		}
		if apmDomainId, ok := s.D.GetOkExists("apm_domain_id"); ok {
			tmp := apmDomainId.(string)
			request.ApmDomainId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if opcDryRun, ok := s.D.GetOkExists("opc_dry_run"); ok {
			tmp := opcDryRun.(string)
			request.OpcDryRun = &tmp
		}
		request.UpdateConfigDetails = details
	case strings.ToLower("APDEX"):
		details := oci_apm_config.UpdateApdexRulesDetails{}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if rules, ok := s.D.GetOkExists("rules"); ok {
			interfaces := rules.([]interface{})
			tmp := make([]oci_apm_config.Apdex, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "rules", stateDataIndex)
				converted, err := s.mapToApdex(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("rules") {
				details.Rules = tmp
			}
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if opcDryRun, ok := s.D.GetOkExists("opc_dry_run"); ok {
			tmp := opcDryRun.(string)
			request.OpcDryRun = &tmp
		}
		request.UpdateConfigDetails = details
	case strings.ToLower("MACS_APM_EXTENSION"):
		details := oci_apm_config.UpdateMacsApmExtensionDetails{}
		if agentVersion, ok := s.D.GetOkExists("agent_version"); ok {
			tmp := agentVersion.(string)
			details.AgentVersion = &tmp
		}
		if attachInstallDir, ok := s.D.GetOkExists("attach_install_dir"); ok {
			tmp := attachInstallDir.(string)
			details.AttachInstallDir = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if processFilter, ok := s.D.GetOkExists("process_filter"); ok {
			interfaces := processFilter.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("process_filter") {
				details.ProcessFilter = tmp
			}
		}
		if runAsUser, ok := s.D.GetOkExists("run_as_user"); ok {
			tmp := runAsUser.(string)
			details.RunAsUser = &tmp
		}
		if serviceName, ok := s.D.GetOkExists("service_name"); ok {
			tmp := serviceName.(string)
			details.ServiceName = &tmp
		}
		if apmDomainId, ok := s.D.GetOkExists("apm_domain_id"); ok {
			tmp := apmDomainId.(string)
			request.ApmDomainId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if opcDryRun, ok := s.D.GetOkExists("opc_dry_run"); ok {
			tmp := opcDryRun.(string)
			request.OpcDryRun = &tmp
		}
		request.UpdateConfigDetails = details
	case strings.ToLower("METRIC_GROUP"):
		details := oci_apm_config.UpdateMetricGroupDetails{}
		if dimensions, ok := s.D.GetOkExists("dimensions"); ok {
			interfaces := dimensions.([]interface{})
			tmp := make([]oci_apm_config.Dimension, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "dimensions", stateDataIndex)
				converted, err := s.mapToDimension(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("dimensions") {
				details.Dimensions = tmp
			}
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if filterId, ok := s.D.GetOkExists("filter_id"); ok {
			tmp := filterId.(string)
			details.FilterId = &tmp
		}
		if metrics, ok := s.D.GetOkExists("metrics"); ok {
			interfaces := metrics.([]interface{})
			tmp := make([]oci_apm_config.Metric, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "metrics", stateDataIndex)
				converted, err := s.mapToMetric(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("metrics") {
				details.Metrics = tmp
			}
		}
		if namespace, ok := s.D.GetOkExists("namespace"); ok {
			tmp := namespace.(string)
			details.Namespace = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if opcDryRun, ok := s.D.GetOkExists("opc_dry_run"); ok {
			tmp := opcDryRun.(string)
			request.OpcDryRun = &tmp
		}
		request.UpdateConfigDetails = details
	case strings.ToLower("OPTIONS"):
		details := oci_apm_config.UpdateOptionsDetails{}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if group, ok := s.D.GetOkExists("group"); ok {
			tmp := group.(string)
			details.Group = &tmp
		}
		if options, ok := s.D.GetOkExists("options"); ok {
			tmp, err := mapToOptions(options.(string))
			if err != nil {
				return err
			}
			details.Options = tmp
		}
		/*
			if options, ok := s.D.GetOkExists("options"); ok {
				if tmpList := options.([]interface{}); len(tmpList) > 0 {
					fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "options", 0)
					tmp, err := s.mapToOptions(fieldKeyFormat)
					if err != nil {
						return err
					}
					details.Options = tmp
				}
			}
		*/
		if apmDomainId, ok := s.D.GetOkExists("apm_domain_id"); ok {
			tmp := apmDomainId.(string)
			request.ApmDomainId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if opcDryRun, ok := s.D.GetOkExists("opc_dry_run"); ok {
			tmp := opcDryRun.(string)
			request.OpcDryRun = &tmp
		}
		request.UpdateConfigDetails = details
	case strings.ToLower("SPAN_FILTER"):
		details := oci_apm_config.UpdateSpanFilterDetails{}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if filterText, ok := s.D.GetOkExists("filter_text"); ok {
			tmp := filterText.(string)
			details.FilterText = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if opcDryRun, ok := s.D.GetOkExists("opc_dry_run"); ok {
			tmp := opcDryRun.(string)
			request.OpcDryRun = &tmp
		}
		request.UpdateConfigDetails = details
	default:
		return fmt.Errorf("unknown config_type '%v' was specified", configType)
	}
	return nil
}
