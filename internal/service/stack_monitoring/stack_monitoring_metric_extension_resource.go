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

func StackMonitoringMetricExtensionResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createStackMonitoringMetricExtension,
		Read:     readStackMonitoringMetricExtension,
		Update:   updateStackMonitoringMetricExtension,
		Delete:   deleteStackMonitoringMetricExtension,
		Schema: map[string]*schema.Schema{
			// Required
			"collection_recurrences": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"metric_list": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"data_type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"compute_expression": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  " ",
						},
						"display_name": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  " ",
						},
						"is_dimension": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"is_hidden": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"metric_category": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"unit": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  " ",
						},

						// Computed
					},
				},
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"query_properties": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"collection_method": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"JMX",
								"OS_COMMAND",
								"SQL",
							}, true),
						},

						// Optional
						"arguments": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"auto_row_prefix": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"command": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"delimiter": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"identity_metric": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"in_param_details": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"in_param_position": {
										Type:     schema.TypeInt,
										Required: true,
									},
									"in_param_value": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Computed
								},
							},
						},
						"is_metric_service_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"jmx_attributes": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"managed_bean_query": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"out_param_details": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"out_param_position": {
										Type:     schema.TypeInt,
										Required: true,
									},
									"out_param_type": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Computed
								},
							},
						},
						"script_details": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"content": {
										Type:     schema.TypeString,
										Required: true,
									},
									"name": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Computed
								},
							},
						},
						"sql_details": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"content": {
										Type:     schema.TypeString,
										Required: true,
									},
									"script_file_name": {
										Type:     schema.TypeString,
										Optional: true,
										Default:  " ",
									},

									// Computed
								},
							},
						},
						"sql_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"starts_with": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"resource_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"publish_trigger": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			// Computed
			"collection_method": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enabled_on_resources": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"resource_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"enabled_on_resources_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"resource_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tenant_id": {
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

func createStackMonitoringMetricExtension(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringMetricExtensionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()

	publishTrigger := false

	if _, ok := sync.D.GetOkExists("publish_trigger"); ok {
		publishTrigger = d.Get("publish_trigger").(bool)
	}

	if e := tfresource.CreateResource(d, sync); e != nil {
		return e
	}

	if publishTrigger {
		err := sync.PublishMetricExtension()
		if err != nil {
			return tfresource.HandleError(sync, err)
		}

		err = sync.SetData()
		if err != nil {
			return err
		}
	}

	return nil

}

func readStackMonitoringMetricExtension(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringMetricExtensionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()

	return tfresource.ReadResource(sync)
}

func updateStackMonitoringMetricExtension(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringMetricExtensionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()

	// Either it can be an update metric extension request or a publish metric extension request
	// both can not be clubbed together in a single change. Added additional logic such that if publish
	// trigger is changing from 0 to 1 then only do publish action and don't do any update. If publish value is
	// either not passed or is same as previous value i.e. it has not changed and its 0 , then only continue with
	// normal update of metric extension

	if _, ok := sync.D.GetOkExists("publish_trigger"); ok && sync.D.HasChange("publish_trigger") {

		if sync.D.HasChangeExcept("publish_trigger") {
			return fmt.Errorf("publishing of metric extension change cannot be combined with other edits to the metric extension")
		}

		oldRaw, newRaw := sync.D.GetChange("publish_trigger")
		newValue := newRaw.(bool)
		if newValue {
			err := sync.PublishMetricExtension()
			if err != nil {
				return tfresource.HandleError(sync, err)
			}

			err = sync.SetData()
			if err != nil {
				return err
			}

			return nil
		} else {
			sync.D.Set("publish_trigger", oldRaw)
			return fmt.Errorf("value of publish trigger can only be true or false. Once set to true it can not be set back to false i.e. published status of metric extension cannot be changed back to draft")
		}
	}

	if _, ok := sync.D.GetOkExists("publish_trigger"); ok {
		publishTrigger := sync.D.Get("publish_trigger").(bool)

		if publishTrigger {
			return fmt.Errorf("published metric extension cannot be edited")
		}

	}

	if err := tfresource.UpdateResource(d, sync); err != nil {
		return err
	}

	return nil
}

func deleteStackMonitoringMetricExtension(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringMetricExtensionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type StackMonitoringMetricExtensionResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_stack_monitoring.StackMonitoringClient
	Res                    *oci_stack_monitoring.MetricExtension
	DisableNotFoundRetries bool
}

func (s *StackMonitoringMetricExtensionResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *StackMonitoringMetricExtensionResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *StackMonitoringMetricExtensionResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_stack_monitoring.MetricExtensionLifeCycleStatesActive),
	}
}

func (s *StackMonitoringMetricExtensionResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *StackMonitoringMetricExtensionResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_stack_monitoring.MetricExtensionLifeCycleStatesDeleted),
	}
}

func (s *StackMonitoringMetricExtensionResourceCrud) Create() error {
	request := oci_stack_monitoring.CreateMetricExtensionRequest{}

	if collectionRecurrences, ok := s.D.GetOkExists("collection_recurrences"); ok {
		tmp := collectionRecurrences.(string)
		request.CollectionRecurrences = &tmp
	}

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

	if metricList, ok := s.D.GetOkExists("metric_list"); ok {
		interfaces := metricList.([]interface{})
		tmp := make([]oci_stack_monitoring.Metric, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "metric_list", stateDataIndex)
			converted, err := s.mapToMetric(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("metric_list") {
			request.MetricList = tmp
		}
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if queryProperties, ok := s.D.GetOkExists("query_properties"); ok {
		if tmpList := queryProperties.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "query_properties", 0)
			tmp, err := s.mapToMetricExtensionQueryProperties(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.QueryProperties = tmp
		}
	}

	if resourceType, ok := s.D.GetOkExists("resource_type"); ok {
		tmp := resourceType.(string)
		request.ResourceType = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring")

	response, err := s.Client.CreateMetricExtension(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.MetricExtension
	return nil
}

func (s *StackMonitoringMetricExtensionResourceCrud) Get() error {
	request := oci_stack_monitoring.GetMetricExtensionRequest{}

	tmp := s.D.Id()
	request.MetricExtensionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring")

	response, err := s.Client.GetMetricExtension(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.MetricExtension

	return nil
}

func (s *StackMonitoringMetricExtensionResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_stack_monitoring.UpdateMetricExtensionRequest{}

	if collectionRecurrences, ok := s.D.GetOkExists("collection_recurrences"); ok {
		tmp := collectionRecurrences.(string)
		request.CollectionRecurrences = &tmp
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	tmp := s.D.Id()
	request.MetricExtensionId = &tmp

	if metricList, ok := s.D.GetOkExists("metric_list"); ok {
		interfaces := metricList.([]interface{})
		tmp := make([]oci_stack_monitoring.Metric, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "metric_list", stateDataIndex)
			converted, err := s.mapToMetric(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("metric_list") {
			request.MetricList = tmp
		}
	}

	if queryProperties, ok := s.D.GetOkExists("query_properties"); ok {
		if tmpList := queryProperties.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "query_properties", 0)
			tmp, err := s.mapToMetricExtensionUpdateQueryProperties(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.QueryProperties = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring")

	response, err := s.Client.UpdateMetricExtension(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.MetricExtension
	return nil
}

func (s *StackMonitoringMetricExtensionResourceCrud) Delete() error {
	request := oci_stack_monitoring.DeleteMetricExtensionRequest{}

	tmp := s.D.Id()
	request.MetricExtensionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring")

	_, err := s.Client.DeleteMetricExtension(context.Background(), request)
	return err
}

func (s *StackMonitoringMetricExtensionResourceCrud) SetData() error {
	if s.Res.CollectionMethod != nil {
		s.D.Set("collection_method", *s.Res.CollectionMethod)
	}

	if s.Res.CollectionRecurrences != nil {
		s.D.Set("collection_recurrences", *s.Res.CollectionRecurrences)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CreatedBy != nil {
		s.D.Set("created_by", *s.Res.CreatedBy)
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	enabledOnResources := []interface{}{}
	for _, item := range s.Res.EnabledOnResources {
		enabledOnResources = append(enabledOnResources, EnabledResourceDetailsToMap(item))
	}
	s.D.Set("enabled_on_resources", enabledOnResources)

	if s.Res.EnabledOnResourcesCount != nil {
		s.D.Set("enabled_on_resources_count", *s.Res.EnabledOnResourcesCount)
	}

	if s.Res.LastUpdatedBy != nil {
		s.D.Set("last_updated_by", *s.Res.LastUpdatedBy)
	}

	metricList := []interface{}{}
	for _, item := range s.Res.MetricList {
		metricList = append(metricList, MetricToMap(item))
	}
	s.D.Set("metric_list", metricList)

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.QueryProperties != nil {
		queryPropertiesArray := []interface{}{}

		if queryPropertiesMap := MetricExtensionQueryPropertiesToMap(&s.Res.QueryProperties); queryPropertiesMap != nil {
			queryPropertiesArray = append(queryPropertiesArray, queryPropertiesMap)
		}
		s.D.Set("query_properties", queryPropertiesArray)
	} else {
		s.D.Set("query_properties", nil)
	}

	if s.Res.ResourceType != nil {
		s.D.Set("resource_type", *s.Res.ResourceType)
	}

	if s.Res.ResourceUri != nil {
		s.D.Set("resource_uri", *s.Res.ResourceUri)
	}

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("status", s.Res.Status)

	if s.Res.Status == oci_stack_monitoring.MetricExtensionLifeCycleDetailsDraft {
		s.D.Set("publish_trigger", false)
	}

	if s.Res.Status == oci_stack_monitoring.MetricExtensionLifeCycleDetailsPublished {
		s.D.Set("publish_trigger", true)
	}

	if s.Res.TenantId != nil {
		s.D.Set("tenant_id", *s.Res.TenantId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func (s *StackMonitoringMetricExtensionResourceCrud) PublishMetricExtension() error {
	request := oci_stack_monitoring.PublishMetricExtensionRequest{}

	idTmp := s.D.Id()
	request.MetricExtensionId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring")

	response, err := s.Client.PublishMetricExtension(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	val := s.D.Get("publish_trigger")
	s.D.Set("publish_trigger", val)

	s.Res = &response.MetricExtension
	return nil
}

func EnabledResourceDetailsToMap(obj oci_stack_monitoring.EnabledResourceDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ResourceId != nil {
		result["resource_id"] = string(*obj.ResourceId)
	}

	return result
}

func (s *StackMonitoringMetricExtensionResourceCrud) mapToMetric(fieldKeyFormat string) (oci_stack_monitoring.Metric, error) {
	result := oci_stack_monitoring.Metric{}

	if computeExpression, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compute_expression")); ok {
		tmp := computeExpression.(string)
		result.ComputeExpression = &tmp
	}

	if dataType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "data_type")); ok {
		result.DataType = oci_stack_monitoring.MetricDataTypeEnum(dataType.(string))
	}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	if isDimension, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_dimension")); ok {
		tmp := isDimension.(bool)
		result.IsDimension = &tmp
	}

	if isHidden, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_hidden")); ok {
		tmp := isHidden.(bool)
		result.IsHidden = &tmp
	}

	if metricCategory, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "metric_category")); ok {
		result.MetricCategory = oci_stack_monitoring.MetricMetricCategoryEnum(metricCategory.(string))
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if unit, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "unit")); ok {
		tmp := unit.(string)
		result.Unit = &tmp
	}

	return result, nil
}

func MetricToMap(obj oci_stack_monitoring.Metric) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ComputeExpression != nil {
		result["compute_expression"] = string(*obj.ComputeExpression)
	}

	result["data_type"] = string(obj.DataType)

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.IsDimension != nil {
		result["is_dimension"] = bool(*obj.IsDimension)
	}

	if obj.IsHidden != nil {
		result["is_hidden"] = bool(*obj.IsHidden)
	}

	result["metric_category"] = string(obj.MetricCategory)

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Unit != nil {
		result["unit"] = string(*obj.Unit)
	}

	return result
}

func (s *StackMonitoringMetricExtensionResourceCrud) mapToMetricExtensionQueryProperties(fieldKeyFormat string) (oci_stack_monitoring.MetricExtensionQueryProperties, error) {
	var baseObject oci_stack_monitoring.MetricExtensionQueryProperties
	//discriminator
	collectionMethodRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "collection_method"))
	var collectionMethod string
	if ok {
		collectionMethod = collectionMethodRaw.(string)
	} else {
		collectionMethod = "" // default value
	}
	switch strings.ToLower(collectionMethod) {
	case strings.ToLower("JMX"):
		details := oci_stack_monitoring.JmxUpdateQueryProperties{}
		if autoRowPrefix, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "auto_row_prefix")); ok {
			tmp := autoRowPrefix.(string)
			details.AutoRowPrefix = &tmp
		}
		if identityMetric, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "identity_metric")); ok {
			tmp := identityMetric.(string)
			details.IdentityMetric = &tmp
		}
		if isMetricServiceEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_metric_service_enabled")); ok {
			tmp := isMetricServiceEnabled.(bool)
			details.IsMetricServiceEnabled = &tmp
		}
		if jmxAttributes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "jmx_attributes")); ok {
			tmp := jmxAttributes.(string)
			details.JmxAttributes = &tmp
		}
		if managedBeanQuery, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "managed_bean_query")); ok {
			tmp := managedBeanQuery.(string)
			details.ManagedBeanQuery = &tmp
		}
		baseObject = details
	case strings.ToLower("OS_COMMAND"):
		details := oci_stack_monitoring.OsCommandUpdateQueryProperties{}
		if arguments, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "arguments")); ok {
			tmp := arguments.(string)
			details.Arguments = &tmp
		}
		if command, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "command")); ok {
			tmp := command.(string)
			details.Command = &tmp
		}
		if delimiter, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "delimiter")); ok {
			tmp := delimiter.(string)
			details.Delimiter = &tmp
		}
		if scriptDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "script_details")); ok {
			if tmpList := scriptDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "script_details"), 0)
				tmp, err := s.mapToScriptFileDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert script_details, encountered error: %v", err)
				}
				details.ScriptDetails = &tmp
			}
		}
		if startsWith, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "starts_with")); ok {
			tmp := startsWith.(string)
			details.StartsWith = &tmp
		}
		baseObject = details
	case strings.ToLower("SQL"):
		details := oci_stack_monitoring.SqlUpdateQueryProperties{}
		if inParamDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "in_param_details")); ok {
			interfaces := inParamDetails.([]interface{})
			tmp := make([]oci_stack_monitoring.SqlInParamDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "in_param_details"), stateDataIndex)
				converted, err := s.mapToSqlInParamDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "in_param_details")) {
				details.InParamDetails = tmp
			}
		}
		if outParamDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "out_param_details")); ok {
			if tmpList := outParamDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "out_param_details"), 0)
				tmp, err := s.mapToSqlOutParamDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert out_param_details, encountered error: %v", err)
				}
				details.OutParamDetails = &tmp
			}
		}
		if sqlDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "sql_details")); ok {
			if tmpList := sqlDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "sql_details"), 0)
				tmp, err := s.mapToSqlDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert sql_details, encountered error: %v", err)
				}
				details.SqlDetails = &tmp
			}
		}
		if sqlType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "sql_type")); ok {
			details.SqlType = oci_stack_monitoring.SqlQueryTypesEnum(sqlType.(string))
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown collection_method '%v' was specified", collectionMethod)
	}
	return baseObject, nil
}

func (s *StackMonitoringMetricExtensionResourceCrud) mapToMetricExtensionUpdateQueryProperties(fieldKeyFormat string) (oci_stack_monitoring.MetricExtensionUpdateQueryProperties, error) {
	var baseObject oci_stack_monitoring.MetricExtensionUpdateQueryProperties
	//discriminator
	collectionMethodRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "collection_method"))
	var collectionMethod string
	if ok {
		collectionMethod = collectionMethodRaw.(string)
	} else {
		collectionMethod = "" // default value
	}
	switch strings.ToLower(collectionMethod) {
	case strings.ToLower("JMX"):
		details := oci_stack_monitoring.JmxUpdateQueryProperties{}
		if autoRowPrefix, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "auto_row_prefix")); ok {
			tmp := autoRowPrefix.(string)
			details.AutoRowPrefix = &tmp
		}
		if identityMetric, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "identity_metric")); ok {
			tmp := identityMetric.(string)
			details.IdentityMetric = &tmp
		}
		if isMetricServiceEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_metric_service_enabled")); ok {
			tmp := isMetricServiceEnabled.(bool)
			details.IsMetricServiceEnabled = &tmp
		}
		if jmxAttributes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "jmx_attributes")); ok {
			tmp := jmxAttributes.(string)
			details.JmxAttributes = &tmp
		}
		if managedBeanQuery, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "managed_bean_query")); ok {
			tmp := managedBeanQuery.(string)
			details.ManagedBeanQuery = &tmp
		}
		baseObject = details
	case strings.ToLower("OS_COMMAND"):
		details := oci_stack_monitoring.OsCommandUpdateQueryProperties{}
		if arguments, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "arguments")); ok {
			tmp := arguments.(string)
			details.Arguments = &tmp
		}
		if command, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "command")); ok {
			tmp := command.(string)
			details.Command = &tmp
		}
		if delimiter, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "delimiter")); ok {
			tmp := delimiter.(string)
			details.Delimiter = &tmp
		}
		if scriptDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "script_details")); ok {
			if tmpList := scriptDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "script_details"), 0)
				tmp, err := s.mapToScriptFileDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert script_details, encountered error: %v", err)
				}
				details.ScriptDetails = &tmp
			}
		}
		if startsWith, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "starts_with")); ok {
			tmp := startsWith.(string)
			details.StartsWith = &tmp
		}
		baseObject = details
	case strings.ToLower("SQL"):
		details := oci_stack_monitoring.SqlUpdateQueryProperties{}
		if inParamDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "in_param_details")); ok {
			interfaces := inParamDetails.([]interface{})
			tmp := make([]oci_stack_monitoring.SqlInParamDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "in_param_details"), stateDataIndex)
				converted, err := s.mapToSqlInParamDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "in_param_details")) {
				details.InParamDetails = tmp
			}
		}
		if outParamDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "out_param_details")); ok {
			if tmpList := outParamDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "out_param_details"), 0)
				tmp, err := s.mapToSqlOutParamDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert out_param_details, encountered error: %v", err)
				}
				details.OutParamDetails = &tmp
			}
		}
		if sqlDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "sql_details")); ok {
			if tmpList := sqlDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "sql_details"), 0)
				tmp, err := s.mapToSqlDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert sql_details, encountered error: %v", err)
				}
				details.SqlDetails = &tmp
			}
		}
		if sqlType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "sql_type")); ok {
			details.SqlType = oci_stack_monitoring.SqlQueryTypesEnum(sqlType.(string))
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown collection_method '%v' was specified", collectionMethod)
	}
	return baseObject, nil
}

func MetricExtensionQueryPropertiesToMap(obj *oci_stack_monitoring.MetricExtensionQueryProperties) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_stack_monitoring.JmxQueryProperties:
		result["collection_method"] = "JMX"

		if v.AutoRowPrefix != nil {
			result["auto_row_prefix"] = string(*v.AutoRowPrefix)
		}

		if v.IdentityMetric != nil {
			result["identity_metric"] = string(*v.IdentityMetric)
		}

		if v.IsMetricServiceEnabled != nil {
			result["is_metric_service_enabled"] = bool(*v.IsMetricServiceEnabled)
		}

		if v.JmxAttributes != nil {
			result["jmx_attributes"] = string(*v.JmxAttributes)
		}

		if v.ManagedBeanQuery != nil {
			result["managed_bean_query"] = string(*v.ManagedBeanQuery)
		}
	case oci_stack_monitoring.OsCommandQueryProperties:
		result["collection_method"] = "OS_COMMAND"

		if v.Arguments != nil {
			result["arguments"] = string(*v.Arguments)
		}

		if v.Command != nil {
			result["command"] = string(*v.Command)
		}

		if v.Delimiter != nil {
			result["delimiter"] = string(*v.Delimiter)
		}

		if v.ScriptDetails != nil {
			result["script_details"] = []interface{}{ScriptFileDetailsToMap(v.ScriptDetails)}
		}

		if v.StartsWith != nil {
			result["starts_with"] = string(*v.StartsWith)
		}
	case oci_stack_monitoring.SqlQueryProperties:
		result["collection_method"] = "SQL"

		inParamDetails := []interface{}{}
		for _, item := range v.InParamDetails {
			inParamDetails = append(inParamDetails, SqlInParamDetailsToMap(item))
		}
		result["in_param_details"] = inParamDetails

		if v.OutParamDetails != nil {
			result["out_param_details"] = []interface{}{SqlOutParamDetailsToMap(v.OutParamDetails)}
		}

		if v.SqlDetails != nil {
			result["sql_details"] = []interface{}{SqlDetailsToMap(v.SqlDetails)}
		}

		result["sql_type"] = string(v.SqlType)
	default:
		log.Printf("[WARN] Received 'collection_method' of unknown type %v", *obj)
		return nil
	}

	return result
}

func MetricExtensionSummaryToMap(obj oci_stack_monitoring.MetricExtensionSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["collection_method"] = string(obj.CollectionMethod)

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.EnabledOnResourcesCount != nil {
		result["enabled_on_resources_count"] = int(*obj.EnabledOnResourcesCount)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.ResourceType != nil {
		result["resource_type"] = string(*obj.ResourceType)
	}

	if obj.ResourceUri != nil {
		result["resource_uri"] = string(*obj.ResourceUri)
	}

	result["state"] = string(obj.LifecycleState)

	result["status"] = string(obj.Status)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func (s *StackMonitoringMetricExtensionResourceCrud) mapToScriptFileDetails(fieldKeyFormat string) (oci_stack_monitoring.ScriptFileDetails, error) {
	result := oci_stack_monitoring.ScriptFileDetails{}

	if content, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "content")); ok {
		tmp := content.(string)
		result.Content = &tmp
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	return result, nil
}

func ScriptFileDetailsToMap(obj *oci_stack_monitoring.ScriptFileDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Content != nil {
		result["content"] = string(*obj.Content)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}

func (s *StackMonitoringMetricExtensionResourceCrud) mapToSqlDetails(fieldKeyFormat string) (oci_stack_monitoring.SqlDetails, error) {
	result := oci_stack_monitoring.SqlDetails{}

	if content, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "content")); ok {
		tmp := content.(string)
		result.Content = &tmp
	}

	if scriptFileName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "script_file_name")); ok {
		tmp := scriptFileName.(string)
		result.ScriptFileName = &tmp
	}

	return result, nil
}

func SqlDetailsToMap(obj *oci_stack_monitoring.SqlDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Content != nil {
		result["content"] = string(*obj.Content)
	}

	if obj.ScriptFileName != nil {
		result["script_file_name"] = string(*obj.ScriptFileName)
	}

	return result
}

func (s *StackMonitoringMetricExtensionResourceCrud) mapToSqlInParamDetails(fieldKeyFormat string) (oci_stack_monitoring.SqlInParamDetails, error) {
	result := oci_stack_monitoring.SqlInParamDetails{}

	if inParamPosition, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "in_param_position")); ok {
		tmp := inParamPosition.(int)
		result.InParamPosition = &tmp
	}

	if inParamValue, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "in_param_value")); ok {
		tmp := inParamValue.(string)
		result.InParamValue = &tmp
	}

	return result, nil
}

func SqlInParamDetailsToMap(obj oci_stack_monitoring.SqlInParamDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.InParamPosition != nil {
		result["in_param_position"] = int(*obj.InParamPosition)
	}

	if obj.InParamValue != nil {
		result["in_param_value"] = string(*obj.InParamValue)
	}

	return result
}

func (s *StackMonitoringMetricExtensionResourceCrud) mapToSqlOutParamDetails(fieldKeyFormat string) (oci_stack_monitoring.SqlOutParamDetails, error) {
	result := oci_stack_monitoring.SqlOutParamDetails{}

	if outParamPosition, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "out_param_position")); ok {
		tmp := outParamPosition.(int)
		result.OutParamPosition = &tmp
	}

	if outParamType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "out_param_type")); ok {
		result.OutParamType = oci_stack_monitoring.SqlOutParamTypesEnum(outParamType.(string))
	}

	return result, nil
}

func SqlOutParamDetailsToMap(obj *oci_stack_monitoring.SqlOutParamDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.OutParamPosition != nil {
		result["out_param_position"] = int(*obj.OutParamPosition)
	}

	result["out_param_type"] = string(obj.OutParamType)

	return result
}

func (s *StackMonitoringMetricExtensionResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_stack_monitoring.ChangeMetricExtensionCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.MetricExtensionId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring")

	_, err := s.Client.ChangeMetricExtensionCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
