// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package demand_signal

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	oci_demand_signal "github.com/oracle/oci-go-sdk/v65/demandsignal"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DemandSignalOccMetricAlarmResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createDemandSignalOccMetricAlarmWithContext,
		ReadContext:   readDemandSignalOccMetricAlarmWithContext,
		UpdateContext: updateDemandSignalOccMetricAlarmWithContext,
		DeleteContext: deleteDemandSignalOccMetricAlarmWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"frequency": {
				Type:     schema.TypeString,
				Required: true,
			},
			"is_active": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"resource_configuration": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"resource": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"COMPUTE",
								"EXADATA",
								"NETWORK",
								"STORAGE",
							}, true),
						},
						"usage_type": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"compute_hw_generation": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"hw_generation": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"link_role": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"node_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"occ_metric_alarm_provider": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"shape": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"storage_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"threshold": {
				Type:     schema.TypeInt,
				Required: true,
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
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"subscribers": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"threshold_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
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

func createDemandSignalOccMetricAlarmWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DemandSignalOccMetricAlarmResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OccMetricAlarmClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readDemandSignalOccMetricAlarmWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DemandSignalOccMetricAlarmResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OccMetricAlarmClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateDemandSignalOccMetricAlarmWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DemandSignalOccMetricAlarmResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OccMetricAlarmClient()

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func deleteDemandSignalOccMetricAlarmWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DemandSignalOccMetricAlarmResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OccMetricAlarmClient()
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type DemandSignalOccMetricAlarmResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_demand_signal.OccMetricAlarmClient
	Res                    *oci_demand_signal.OccMetricAlarm
	DisableNotFoundRetries bool
}

func (s *DemandSignalOccMetricAlarmResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DemandSignalOccMetricAlarmResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_demand_signal.OccMetricAlarmLifecycleStateCreating),
	}
}

func (s *DemandSignalOccMetricAlarmResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_demand_signal.OccMetricAlarmLifecycleStateActive),
	}
}

func (s *DemandSignalOccMetricAlarmResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_demand_signal.OccMetricAlarmLifecycleStateDeleting),
	}
}

func (s *DemandSignalOccMetricAlarmResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_demand_signal.OccMetricAlarmLifecycleStateDeleted),
	}
}

func (s *DemandSignalOccMetricAlarmResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_demand_signal.CreateOccMetricAlarmRequest{}

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

	if frequency, ok := s.D.GetOkExists("frequency"); ok {
		request.Frequency = oci_demand_signal.OccMetricAlarmFrequencyEnum(frequency.(string))
	}

	if isActive, ok := s.D.GetOkExists("is_active"); ok {
		tmp := isActive.(bool)
		request.IsActive = &tmp
	}

	if resourceConfiguration, ok := s.D.GetOkExists("resource_configuration"); ok {
		if tmpList := resourceConfiguration.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "resource_configuration", 0)
			tmp, err := s.mapToBaseResourceConfiguration(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ResourceConfiguration = tmp
		}
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_demand_signal.CreateOccMetricAlarmDetailsLifecycleStateEnum(state.(string))
	}

	if subscribers, ok := s.D.GetOkExists("subscribers"); ok {
		interfaces := subscribers.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("subscribers") {
			request.Subscribers = tmp
		}
	}

	if threshold, ok := s.D.GetOkExists("threshold"); ok {
		tmp := threshold.(int)
		request.Threshold = &tmp
	}

	if thresholdType, ok := s.D.GetOkExists("threshold_type"); ok {
		request.ThresholdType = oci_demand_signal.CreateOccMetricAlarmDetailsThresholdTypeEnum(thresholdType.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "demand_signal")

	response, err := s.Client.CreateOccMetricAlarm(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.OccMetricAlarm
	return nil
}

func (s *DemandSignalOccMetricAlarmResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_demand_signal.GetOccMetricAlarmRequest{}

	tmp := s.D.Id()
	request.OccMetricAlarmId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "demand_signal")

	response, err := s.Client.GetOccMetricAlarm(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.OccMetricAlarm
	return nil
}

func (s *DemandSignalOccMetricAlarmResourceCrud) UpdateWithContext(ctx context.Context) error {
	request := oci_demand_signal.UpdateOccMetricAlarmRequest{}

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

	if frequency, ok := s.D.GetOkExists("frequency"); ok {
		request.Frequency = oci_demand_signal.OccMetricAlarmFrequencyEnum(frequency.(string))
	}

	if isActive, ok := s.D.GetOkExists("is_active"); ok {
		tmp := isActive.(bool)
		request.IsActive = &tmp
	}

	tmp := s.D.Id()
	request.OccMetricAlarmId = &tmp

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_demand_signal.UpdateOccMetricAlarmDetailsLifecycleStateEnum(state.(string))
	}

	if subscribers, ok := s.D.GetOkExists("subscribers"); ok {
		interfaces := subscribers.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("subscribers") {
			request.Subscribers = tmp
		}
	}

	if threshold, ok := s.D.GetOkExists("threshold"); ok {
		tmp := threshold.(int)
		request.Threshold = &tmp
	}

	if thresholdType, ok := s.D.GetOkExists("threshold_type"); ok {
		request.ThresholdType = oci_demand_signal.UpdateOccMetricAlarmDetailsThresholdTypeEnum(thresholdType.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "demand_signal")

	response, err := s.Client.UpdateOccMetricAlarm(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.OccMetricAlarm
	return nil
}

func (s *DemandSignalOccMetricAlarmResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_demand_signal.DeleteOccMetricAlarmRequest{}

	tmp := s.D.Id()
	request.OccMetricAlarmId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "demand_signal")

	_, err := s.Client.DeleteOccMetricAlarm(ctx, request)
	return err
}

func (s *DemandSignalOccMetricAlarmResourceCrud) SetData() error {
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

	s.D.Set("frequency", s.Res.Frequency)

	if s.Res.IsActive != nil {
		s.D.Set("is_active", *s.Res.IsActive)
	}

	if s.Res.ResourceConfiguration != nil {
		resourceConfigurationArray := []interface{}{}
		if resourceConfigurationMap := BaseResourceConfigurationToMap(&s.Res.ResourceConfiguration); resourceConfigurationMap != nil {
			resourceConfigurationArray = append(resourceConfigurationArray, resourceConfigurationMap)
		}
		s.D.Set("resource_configuration", resourceConfigurationArray)
	} else {
		s.D.Set("resource_configuration", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("subscribers", s.Res.Subscribers)

	s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))

	if s.Res.Threshold != nil {
		s.D.Set("threshold", *s.Res.Threshold)
	}

	s.D.Set("threshold_type", s.Res.ThresholdType)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func (s *DemandSignalOccMetricAlarmResourceCrud) mapToBaseResourceConfiguration(fieldKeyFormat string) (oci_demand_signal.BaseResourceConfiguration, error) {
	var baseObject oci_demand_signal.BaseResourceConfiguration
	//discriminator
	resourceRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "resource"))
	var resource_ string
	if ok {
		resource_ = resourceRaw.(string)
	} else {
		resource_ = "" // default value
	}
	switch strings.ToLower(resource_) {
	case strings.ToLower("COMPUTE"):
		details := oci_demand_signal.ComputeResourceConfiguration{}
		if computeHwGeneration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compute_hw_generation")); ok {
			tmp := computeHwGeneration.(string)
			details.ComputeHwGeneration = &tmp
		}
		if shape, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "shape")); ok {
			tmp := shape.(string)
			details.Shape = &tmp
		}
		if usageType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "usage_type")); ok {
			tmp := usageType.(string)
			details.UsageType = &tmp
		}
		baseObject = details
	case strings.ToLower("EXADATA"):
		details := oci_demand_signal.ExadataResourceConfiguration{}
		if hwGeneration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "hw_generation")); ok {
			tmp := hwGeneration.(string)
			details.HwGeneration = &tmp
		}
		if nodeType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "node_type")); ok {
			tmp := nodeType.(string)
			details.NodeType = &tmp
		}
		if usageType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "usage_type")); ok {
			tmp := usageType.(string)
			details.UsageType = &tmp
		}
		baseObject = details
	case strings.ToLower("NETWORK"):
		details := oci_demand_signal.NetworkResourceConfiguration{}
		if linkRole, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "link_role")); ok {
			tmp := linkRole.(string)
			details.LinkRole = &tmp
		}
		if occMetricAlarmProvider, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "occ_metric_alarm_provider")); ok {
			tmp := occMetricAlarmProvider.(string)
			details.Provider = &tmp
		}
		if usageType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "usage_type")); ok {
			tmp := usageType.(string)
			details.UsageType = &tmp
		}
		baseObject = details
	case strings.ToLower("STORAGE"):
		details := oci_demand_signal.StorageResourceConfiguration{}
		if storageType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "storage_type")); ok {
			tmp := storageType.(string)
			details.StorageType = &tmp
		}
		if usageType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "usage_type")); ok {
			tmp := usageType.(string)
			details.UsageType = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown resource '%v' was specified", resource_)
	}
	return baseObject, nil
}

func BaseResourceConfigurationToMap(obj *oci_demand_signal.BaseResourceConfiguration) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_demand_signal.ComputeResourceConfiguration:
		result["resource"] = "COMPUTE"

		if v.ComputeHwGeneration != nil {
			result["compute_hw_generation"] = string(*v.ComputeHwGeneration)
		}

		if v.Shape != nil {
			result["shape"] = string(*v.Shape)
		}

		if v.UsageType != nil {
			result["usage_type"] = string(*v.UsageType)
		}
	case oci_demand_signal.ExadataResourceConfiguration:
		result["resource"] = "EXADATA"

		if v.HwGeneration != nil {
			result["hw_generation"] = string(*v.HwGeneration)
		}

		if v.NodeType != nil {
			result["node_type"] = string(*v.NodeType)
		}

		if v.UsageType != nil {
			result["usage_type"] = string(*v.UsageType)
		}
	case oci_demand_signal.NetworkResourceConfiguration:
		result["resource"] = "NETWORK"

		if v.LinkRole != nil {
			result["link_role"] = string(*v.LinkRole)
		}

		if v.Provider != nil {
			result["occ_metric_alarm_provider"] = string(*v.Provider)
		}

		if v.UsageType != nil {
			result["usage_type"] = string(*v.UsageType)
		}
	case oci_demand_signal.StorageResourceConfiguration:
		result["resource"] = "STORAGE"

		if v.StorageType != nil {
			result["storage_type"] = string(*v.StorageType)
		}

		if v.UsageType != nil {
			result["usage_type"] = string(*v.UsageType)
		}
	default:
		log.Printf("[WARN] Received 'resource' of unknown type %v", *obj)
		return nil
	}

	return result
}

func OccMetricAlarmSummaryToMap(obj oci_demand_signal.OccMetricAlarmSummary) map[string]interface{} {
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

	result["frequency"] = string(obj.Frequency)

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsActive != nil {
		result["is_active"] = bool(*obj.IsActive)
	}

	if obj.ResourceConfiguration != nil {
		resourceConfigurationArray := []interface{}{}
		if resourceConfigurationMap := BaseResourceConfigurationToMap(&obj.ResourceConfiguration); resourceConfigurationMap != nil {
			resourceConfigurationArray = append(resourceConfigurationArray, resourceConfigurationMap)
		}
		result["resource_configuration"] = resourceConfigurationArray
	}

	result["state"] = string(obj.LifecycleState)

	result["subscribers"] = obj.Subscribers

	result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)

	if obj.Threshold != nil {
		result["threshold"] = int(*obj.Threshold)
	}

	result["threshold_type"] = string(obj.ThresholdType)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}
