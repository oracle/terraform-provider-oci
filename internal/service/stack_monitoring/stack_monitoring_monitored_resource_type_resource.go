// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package stack_monitoring

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_stack_monitoring "github.com/oracle/oci-go-sdk/v65/stackmonitoring"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func StackMonitoringMonitoredResourceTypeResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createStackMonitoringMonitoredResourceType,
		Read:     readStackMonitoringMonitoredResourceType,
		Update:   updateStackMonitoringMonitoredResourceType,
		Delete:   deleteStackMonitoringMonitoredResourceType,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
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
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"metadata": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"format": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"SYSTEM_FORMAT",
							}, true),
						},

						// Optional
						"agent_properties": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"required_properties": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"unique_property_sets": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"properties": {
										Type:     schema.TypeList,
										Required: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									// Optional

									// Computed
								},
							},
						},
						"valid_properties_for_create": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"valid_properties_for_update": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"valid_property_values": {
							Type:     schema.TypeMap,
							Optional: true,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"valid_sub_resource_types": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						// Computed
					},
				},
			},
			"metric_namespace": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resource_category": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"additional_namespace_map": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"availability_metrics_config": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"collection_interval_in_seconds": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"metrics": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"handler_config": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"collectd_resource_name_config": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"exclude_properties": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"include_properties": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"suffix": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"collector_types": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"handler_properties": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"value": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"metric_mappings": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"collector_metric_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_skip_upload": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"metric_upload_interval_in_seconds": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"telemetry_metric_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"metric_name_config": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"exclude_pattern_on_prefix": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_prefix_with_collector_type": {
										Type:     schema.TypeBool,
										Computed: true,
									},
								},
							},
						},
						"metric_upload_interval_in_seconds": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"telegraf_resource_name_config": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"exclude_tags": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"include_tags": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"is_use_tags_only": {
										Type:     schema.TypeBool,
										Computed: true,
									},
								},
							},
						},
						"telemetry_resource_group": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"is_system_defined": {
				Type:     schema.TypeBool,
				Computed: true,
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
			"tenancy_id": {
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

func createStackMonitoringMonitoredResourceType(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringMonitoredResourceTypeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()

	return tfresource.CreateResource(d, sync)
}

func readStackMonitoringMonitoredResourceType(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringMonitoredResourceTypeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()

	return tfresource.ReadResource(sync)
}

func updateStackMonitoringMonitoredResourceType(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringMonitoredResourceTypeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteStackMonitoringMonitoredResourceType(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringMonitoredResourceTypeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type StackMonitoringMonitoredResourceTypeResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_stack_monitoring.StackMonitoringClient
	Res                    *oci_stack_monitoring.MonitoredResourceType
	DisableNotFoundRetries bool
}

func (s *StackMonitoringMonitoredResourceTypeResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *StackMonitoringMonitoredResourceTypeResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_stack_monitoring.ResourceTypeLifecycleStateCreating),
	}
}

func (s *StackMonitoringMonitoredResourceTypeResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_stack_monitoring.ResourceTypeLifecycleStateActive),
	}
}

func (s *StackMonitoringMonitoredResourceTypeResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_stack_monitoring.ResourceTypeLifecycleStateDeleting),
	}
}

func (s *StackMonitoringMonitoredResourceTypeResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_stack_monitoring.ResourceTypeLifecycleStateDeleted),
	}
}

func (s *StackMonitoringMonitoredResourceTypeResourceCrud) Create() error {
	request := oci_stack_monitoring.CreateMonitoredResourceTypeRequest{}

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

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if metadata, ok := s.D.GetOkExists("metadata"); ok {
		if tmpList := metadata.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "metadata", 0)
			tmp, err := s.mapToResourceTypeMetadataDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Metadata = tmp
		}
	}

	if metricNamespace, ok := s.D.GetOkExists("metric_namespace"); ok {
		tmp := metricNamespace.(string)
		request.MetricNamespace = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if resourceCategory, ok := s.D.GetOkExists("resource_category"); ok {
		request.ResourceCategory = oci_stack_monitoring.ResourceCategoryEnum(resourceCategory.(string))
	}

	if sourceType, ok := s.D.GetOkExists("source_type"); ok {
		request.SourceType = oci_stack_monitoring.SourceTypeEnum(sourceType.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring")

	response, err := s.Client.CreateMonitoredResourceType(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.MonitoredResourceType
	return nil
}

func (s *StackMonitoringMonitoredResourceTypeResourceCrud) Get() error {
	request := oci_stack_monitoring.GetMonitoredResourceTypeRequest{}

	tmp := s.D.Id()
	request.MonitoredResourceTypeId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring")

	response, err := s.Client.GetMonitoredResourceType(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.MonitoredResourceType
	return nil
}

func (s *StackMonitoringMonitoredResourceTypeResourceCrud) Update() error {
	request := oci_stack_monitoring.UpdateMonitoredResourceTypeRequest{}

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
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if metadata, ok := s.D.GetOkExists("metadata"); ok {
		if tmpList := metadata.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "metadata", 0)
			tmp, err := s.mapToResourceTypeMetadataDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Metadata = tmp
		}
	}

	if metricNamespace, ok := s.D.GetOkExists("metric_namespace"); ok {
		tmp := metricNamespace.(string)
		request.MetricNamespace = &tmp
	}

	tmp := s.D.Id()
	request.MonitoredResourceTypeId = &tmp

	if resourceCategory, ok := s.D.GetOkExists("resource_category"); ok {
		request.ResourceCategory = oci_stack_monitoring.ResourceCategoryEnum(resourceCategory.(string))
	}

	if sourceType, ok := s.D.GetOkExists("source_type"); ok {
		request.SourceType = oci_stack_monitoring.SourceTypeEnum(sourceType.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring")

	response, err := s.Client.UpdateMonitoredResourceType(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.MonitoredResourceType
	return nil
}

func (s *StackMonitoringMonitoredResourceTypeResourceCrud) Delete() error {
	request := oci_stack_monitoring.DeleteMonitoredResourceTypeRequest{}

	tmp := s.D.Id()
	request.MonitoredResourceTypeId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring")

	_, err := s.Client.DeleteMonitoredResourceType(context.Background(), request)
	return err
}

func (s *StackMonitoringMonitoredResourceTypeResourceCrud) SetData() error {
	s.D.Set("additional_namespace_map", s.Res.AdditionalNamespaceMap)

	if s.Res.AvailabilityMetricsConfig != nil {
		s.D.Set("availability_metrics_config", []interface{}{AvailabilityMetricsDetailsToMap(s.Res.AvailabilityMetricsConfig)})
	} else {
		s.D.Set("availability_metrics_config", nil)
	}

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

	if s.Res.HandlerConfig != nil {
		s.D.Set("handler_config", []interface{}{AgentExtensionHandlerConfigurationToMap(s.Res.HandlerConfig)})
	} else {
		s.D.Set("handler_config", nil)
	}

	if s.Res.IsSystemDefined != nil {
		s.D.Set("is_system_defined", *s.Res.IsSystemDefined)
	}

	if s.Res.Metadata != nil {
		metadataArray := []interface{}{}
		if metadataMap := ResourceTypeMetadataDetailsToMap(&s.Res.Metadata); metadataMap != nil {
			metadataArray = append(metadataArray, metadataMap)
		}
		s.D.Set("metadata", metadataArray)
	} else {
		s.D.Set("metadata", nil)
	}

	if s.Res.MetricNamespace != nil {
		s.D.Set("metric_namespace", *s.Res.MetricNamespace)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("resource_category", s.Res.ResourceCategory)

	s.D.Set("source_type", s.Res.SourceType)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TenancyId != nil {
		s.D.Set("tenancy_id", *s.Res.TenancyId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func AgentExtensionHandlerConfigurationToMap(obj *oci_stack_monitoring.AgentExtensionHandlerConfiguration) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CollectdResourceNameConfig != nil {
		result["collectd_resource_name_config"] = []interface{}{CollectdResourceNameConfigurationDetailsToMap(obj.CollectdResourceNameConfig)}
	}

	result["collector_types"] = obj.CollectorTypes

	handlerProperties := []interface{}{}
	for _, item := range obj.HandlerProperties {
		handlerProperties = append(handlerProperties, AgentExtensionHandlerConfigurationPropertyToMap(item))
	}
	result["handler_properties"] = handlerProperties

	metricMappings := []interface{}{}
	for _, item := range obj.MetricMappings {
		metricMappings = append(metricMappings, AgentExtensionHandlerMetricMappingDetailsToMap(item))
	}
	result["metric_mappings"] = metricMappings

	if obj.MetricNameConfig != nil {
		result["metric_name_config"] = []interface{}{MetricNameConfigurationDetailsToMap(obj.MetricNameConfig)}
	}

	if obj.MetricUploadIntervalInSeconds != nil {
		result["metric_upload_interval_in_seconds"] = int(*obj.MetricUploadIntervalInSeconds)
	}

	if obj.TelegrafResourceNameConfig != nil {
		result["telegraf_resource_name_config"] = []interface{}{TelegrafResourceNameConfigurationDetailsToMap(obj.TelegrafResourceNameConfig)}
	}

	if obj.TelemetryResourceGroup != nil {
		result["telemetry_resource_group"] = string(*obj.TelemetryResourceGroup)
	}

	return result
}

func AgentExtensionHandlerConfigurationPropertyToMap(obj oci_stack_monitoring.AgentExtensionHandlerConfigurationProperty) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func AgentExtensionHandlerMetricMappingDetailsToMap(obj oci_stack_monitoring.AgentExtensionHandlerMetricMappingDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CollectorMetricName != nil {
		result["collector_metric_name"] = string(*obj.CollectorMetricName)
	}

	if obj.IsSkipUpload != nil {
		result["is_skip_upload"] = bool(*obj.IsSkipUpload)
	}

	if obj.MetricUploadIntervalInSeconds != nil {
		result["metric_upload_interval_in_seconds"] = int(*obj.MetricUploadIntervalInSeconds)
	}

	if obj.TelemetryMetricName != nil {
		result["telemetry_metric_name"] = string(*obj.TelemetryMetricName)
	}

	return result
}

func AvailabilityMetricsDetailsToMap(obj *oci_stack_monitoring.AvailabilityMetricsDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CollectionIntervalInSeconds != nil {
		result["collection_interval_in_seconds"] = int(*obj.CollectionIntervalInSeconds)
	}

	result["metrics"] = obj.Metrics

	return result
}

func CollectdResourceNameConfigurationDetailsToMap(obj *oci_stack_monitoring.CollectdResourceNameConfigurationDetails) map[string]interface{} {
	result := map[string]interface{}{}

	result["exclude_properties"] = obj.ExcludeProperties

	result["include_properties"] = obj.IncludeProperties

	if obj.Suffix != nil {
		result["suffix"] = string(*obj.Suffix)
	}

	return result
}

func MetricNameConfigurationDetailsToMap(obj *oci_stack_monitoring.MetricNameConfigurationDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ExcludePatternOnPrefix != nil {
		result["exclude_pattern_on_prefix"] = string(*obj.ExcludePatternOnPrefix)
	}

	if obj.IsPrefixWithCollectorType != nil {
		result["is_prefix_with_collector_type"] = bool(*obj.IsPrefixWithCollectorType)
	}

	return result
}

func MonitoredResourceTypeSummaryToMap(obj oci_stack_monitoring.MonitoredResourceTypeSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["additional_namespace_map"] = obj.AdditionalNamespaceMap

	if obj.AvailabilityMetricsConfig != nil {
		result["availability_metrics_config"] = []interface{}{AvailabilityMetricsDetailsToMap(obj.AvailabilityMetricsConfig)}
	}

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

	if obj.HandlerConfig != nil {
		result["handler_config"] = []interface{}{AgentExtensionHandlerConfigurationToMap(obj.HandlerConfig)}
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsSystemDefined != nil {
		result["is_system_defined"] = bool(*obj.IsSystemDefined)
	}

	if obj.Metadata != nil {
		metadataArray := []interface{}{}
		if metadataMap := ResourceTypeMetadataDetailsToMap(&obj.Metadata); metadataMap != nil {
			metadataArray = append(metadataArray, metadataMap)
		}
		result["metadata"] = metadataArray
	}

	if obj.MetricNamespace != nil {
		result["metric_namespace"] = string(*obj.MetricNamespace)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	result["resource_category"] = string(obj.ResourceCategory)

	result["source_type"] = string(obj.SourceType)

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

func (s *StackMonitoringMonitoredResourceTypeResourceCrud) mapToResourceTypeMetadataDetails(fieldKeyFormat string) (oci_stack_monitoring.ResourceTypeMetadataDetails, error) {
	var baseObject oci_stack_monitoring.ResourceTypeMetadataDetails
	//discriminator
	formatRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "format"))
	var format string
	if ok {
		format = formatRaw.(string)
	} else {
		format = "" // default value
	}
	switch strings.ToLower(format) {
	case strings.ToLower("SYSTEM_FORMAT"):
		details := oci_stack_monitoring.SystemFormatResourceTypeMetadataDetails{}
		if agentProperties, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "agent_properties")); ok {
			interfaces := agentProperties.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "agent_properties")) {
				details.AgentProperties = tmp
			}
		}
		if requiredProperties, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "required_properties")); ok {
			interfaces := requiredProperties.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "required_properties")) {
				details.RequiredProperties = tmp
			}
		}
		if uniquePropertySets, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "unique_property_sets")); ok {
			interfaces := uniquePropertySets.([]interface{})
			tmp := make([]oci_stack_monitoring.UniquePropertySet, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "unique_property_sets"), stateDataIndex)
				converted, err := s.mapToUniquePropertySet(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "unique_property_sets")) {
				details.UniquePropertySets = tmp
			}
		}
		if validPropertiesForCreate, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "valid_properties_for_create")); ok {
			interfaces := validPropertiesForCreate.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "valid_properties_for_create")) {
				details.ValidPropertiesForCreate = tmp
			}
		}
		if validPropertiesForUpdate, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "valid_properties_for_update")); ok {
			interfaces := validPropertiesForUpdate.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "valid_properties_for_update")) {
				details.ValidPropertiesForUpdate = tmp
			}
		}
		if validPropertyValues, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "valid_property_values")); ok {
			details.ValidPropertyValues = ObjectMapToStringListMap(validPropertyValues.(map[string]interface{}))
		}
		if validSubResourceTypes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "valid_sub_resource_types")); ok {
			interfaces := validSubResourceTypes.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "valid_sub_resource_types")) {
				details.ValidSubResourceTypes = tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown format '%v' was specified", format)
	}
	return baseObject, nil
}

func ResourceTypeMetadataDetailsToMap(obj *oci_stack_monitoring.ResourceTypeMetadataDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_stack_monitoring.SystemFormatResourceTypeMetadataDetails:
		result["format"] = "SYSTEM_FORMAT"

		result["agent_properties"] = v.AgentProperties

		result["required_properties"] = v.RequiredProperties

		uniquePropertySets := []interface{}{}
		for _, item := range v.UniquePropertySets {
			uniquePropertySets = append(uniquePropertySets, UniquePropertySetToMap(item))
		}
		result["unique_property_sets"] = uniquePropertySets

		result["valid_properties_for_create"] = v.ValidPropertiesForCreate

		result["valid_properties_for_update"] = v.ValidPropertiesForUpdate

		result["valid_property_values"] = StringListMapToObjectMap(v.ValidPropertyValues)

		result["valid_sub_resource_types"] = v.ValidSubResourceTypes

	default:
		log.Printf("[WARN] Received 'format' of unknown type %v", *obj)
		return nil
	}

	return result
}

func TelegrafResourceNameConfigurationDetailsToMap(obj *oci_stack_monitoring.TelegrafResourceNameConfigurationDetails) map[string]interface{} {
	result := map[string]interface{}{}

	result["exclude_tags"] = obj.ExcludeTags

	result["include_tags"] = obj.IncludeTags

	if obj.IsUseTagsOnly != nil {
		result["is_use_tags_only"] = bool(*obj.IsUseTagsOnly)
	}

	return result
}

/*
Doing the conversion.
input 	-> map of string to comma-separated-string
output 	-> map of string to string array
*/
func ObjectMapToStringListMap(rm map[string]interface{}) map[string][]string {
	result := make(map[string][]string)
	for k, v := range rm {
		result[k] = strings.Split(v.(string), ",")
	}
	return result
}

/*
Doing the conversion
input	-> map of string to string array
output 	-> map of string to comma-separated-string
*/
func StringListMapToObjectMap(mapOfStrToList map[string][]string) map[string]interface{} {
	result := make(map[string]interface{})
	if mapOfStrToList == nil {
		return result
	}
	for k, v := range mapOfStrToList {
		result[k] = strings.Join(v, ",")
	}
	return result
}

func (s *StackMonitoringMonitoredResourceTypeResourceCrud) mapToUniquePropertySet(fieldKeyFormat string) (oci_stack_monitoring.UniquePropertySet, error) {
	result := oci_stack_monitoring.UniquePropertySet{}

	if properties, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "properties")); ok {
		interfaces := properties.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "properties")) {
			result.Properties = tmp
		}
	}

	return result, nil
}

func UniquePropertySetToMap(obj oci_stack_monitoring.UniquePropertySet) map[string]interface{} {
	result := map[string]interface{}{}

	result["properties"] = obj.Properties

	return result
}
