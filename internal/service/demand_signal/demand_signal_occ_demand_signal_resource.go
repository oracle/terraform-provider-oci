// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package demand_signal

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_demand_signal "github.com/oracle/oci-go-sdk/v65/demandsignal"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DemandSignalOccDemandSignalResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDemandSignalOccDemandSignal,
		Read:     readDemandSignalOccDemandSignal,
		Update:   updateDemandSignalOccDemandSignal,
		Delete:   deleteDemandSignalOccDemandSignal,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"is_active": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"occ_demand_signal_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"occ_demand_signals": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"resource_type": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"units": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"values": {
							Type:     schema.TypeList,
							Required: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"time_expected": {
										Type:             schema.TypeString,
										Required:         true,
										ForceNew:         true,
										DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
									},
									"value": {
										Type:     schema.TypeFloat,
										Required: true,
										ForceNew: true,
									},

									// Optional
									"comments": {
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
			"patch_operations": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"from": {
							Type:     schema.TypeString,
							Required: true,
						},
						"operation": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"INSERT",
								"INSERT_MULTIPLE",
								"MERGE",
								"MOVE",
								"PROHIBIT",
								"REMOVE",
								"REPLACE",
								"REQUIRE",
							}, true),
						},
						"selection": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:             schema.TypeMap,
							Required:         true,
							DiffSuppressFunc: tfresource.JsonStringDiffSuppressFunction,
						},
						//"values": {
						//	Type:     schema.TypeList,
						//	Required: true,
						//},

						// Optional
						"position": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"selected_item": {
							Type:     schema.TypeString,
							Optional: true,
						},

						// Computed
					},
				},
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

func createDemandSignalOccDemandSignal(d *schema.ResourceData, m interface{}) error {
	sync := &DemandSignalOccDemandSignalResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OccDemandSignalClient()

	return tfresource.CreateResource(d, sync)
}

func readDemandSignalOccDemandSignal(d *schema.ResourceData, m interface{}) error {
	sync := &DemandSignalOccDemandSignalResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OccDemandSignalClient()

	return tfresource.ReadResource(sync)
}

func updateDemandSignalOccDemandSignal(d *schema.ResourceData, m interface{}) error {
	sync := &DemandSignalOccDemandSignalResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OccDemandSignalClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDemandSignalOccDemandSignal(d *schema.ResourceData, m interface{}) error {
	sync := &DemandSignalOccDemandSignalResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OccDemandSignalClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DemandSignalOccDemandSignalResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_demand_signal.OccDemandSignalClient
	Res                    *oci_demand_signal.OccDemandSignal
	PatchResponse          *oci_demand_signal.OccDemandSignal
	DisableNotFoundRetries bool
}

func (s *DemandSignalOccDemandSignalResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DemandSignalOccDemandSignalResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_demand_signal.OccDemandSignalLifecycleStateCreating),
	}
}

func (s *DemandSignalOccDemandSignalResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_demand_signal.OccDemandSignalLifecycleStateActive),
	}
}

func (s *DemandSignalOccDemandSignalResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_demand_signal.OccDemandSignalLifecycleStateDeleting),
	}
}

func (s *DemandSignalOccDemandSignalResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_demand_signal.OccDemandSignalLifecycleStateDeleted),
	}
}

func (s *DemandSignalOccDemandSignalResourceCrud) Create() error {
	request := oci_demand_signal.CreateOccDemandSignalRequest{}

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

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isActive, ok := s.D.GetOkExists("is_active"); ok {
		tmp := isActive.(bool)
		request.IsActive = &tmp
	}

	if occDemandSignals, ok := s.D.GetOkExists("occ_demand_signals"); ok {
		interfaces := occDemandSignals.([]interface{})
		tmp := make([]oci_demand_signal.OccDemandSignalData, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "occ_demand_signals", stateDataIndex)
			converted, err := s.mapToOccDemandSignalData(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("occ_demand_signals") {
			request.OccDemandSignals = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "demand_signal")

	response, err := s.Client.CreateOccDemandSignal(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.OccDemandSignal
	return nil
}

func (s *DemandSignalOccDemandSignalResourceCrud) Get() error {
	request := oci_demand_signal.GetOccDemandSignalRequest{}

	tmp := s.D.Id()
	request.OccDemandSignalId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "demand_signal")

	response, err := s.Client.GetOccDemandSignal(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.OccDemandSignal
	return nil
}

func (s *DemandSignalOccDemandSignalResourceCrud) Update() error {
	request := oci_demand_signal.UpdateOccDemandSignalRequest{}

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

	if isActive, ok := s.D.GetOkExists("is_active"); ok {
		tmp := isActive.(bool)
		request.IsActive = &tmp
	}

	tmp := s.D.Id()
	request.OccDemandSignalId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "demand_signal")

	response, err := s.Client.UpdateOccDemandSignal(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.OccDemandSignal
	return nil
}

func (s *DemandSignalOccDemandSignalResourceCrud) Delete() error {
	request := oci_demand_signal.DeleteOccDemandSignalRequest{}

	tmp := s.D.Id()
	request.OccDemandSignalId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "demand_signal")

	_, err := s.Client.DeleteOccDemandSignal(context.Background(), request)
	return err
}

func (s *DemandSignalOccDemandSignalResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsActive != nil {
		s.D.Set("is_active", *s.Res.IsActive)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	occDemandSignals := []interface{}{}
	for _, item := range s.Res.OccDemandSignals {
		occDemandSignals = append(occDemandSignals, OccDemandSignalDataToMap(item))
	}
	s.D.Set("occ_demand_signals", occDemandSignals)

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

func (s *DemandSignalOccDemandSignalResourceCrud) mapToOccDemandSignalData(fieldKeyFormat string) (oci_demand_signal.OccDemandSignalData, error) {
	result := oci_demand_signal.OccDemandSignalData{}

	if resourceType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "resource_type")); ok {
		tmp := resourceType.(string)
		result.ResourceType = &tmp
	}

	if units, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "units")); ok {
		tmp := units.(string)
		result.Units = &tmp
	}

	if values, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "values")); ok {
		interfaces := values.([]interface{})
		tmp := make([]oci_demand_signal.OccDemandSignalValue, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "values"), stateDataIndex)
			converted, err := s.mapToOccDemandSignalValue(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "values")) {
			result.Values = tmp
		}
	}

	return result, nil
}

func OccDemandSignalDataToMap(obj oci_demand_signal.OccDemandSignalData) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ResourceType != nil {
		result["resource_type"] = string(*obj.ResourceType)
	}

	if obj.Units != nil {
		result["units"] = string(*obj.Units)
	}

	values := []interface{}{}
	for _, item := range obj.Values {
		values = append(values, OccDemandSignalValueToMap(item))
	}
	result["values"] = values

	return result
}

func OccDemandSignalSummaryToMap(obj oci_demand_signal.OccDemandSignalSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
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

	if obj.IsActive != nil {
		result["is_active"] = bool(*obj.IsActive)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
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

func (s *DemandSignalOccDemandSignalResourceCrud) mapToOccDemandSignalValue(fieldKeyFormat string) (oci_demand_signal.OccDemandSignalValue, error) {
	result := oci_demand_signal.OccDemandSignalValue{}

	if comments, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "comments")); ok {
		tmp := comments.(string)
		result.Comments = &tmp
	}

	if timeExpected, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_expected")); ok {
		tmp, err := time.Parse(time.RFC3339, timeExpected.(string))
		if err != nil {
			return result, err
		}
		result.TimeExpected = &oci_common.SDKTime{Time: tmp}
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp, ok := value.(float32)
		if !ok {
			tmp = float32(value.(float64))
		}
		result.Value = &tmp
	}

	return result, nil
}

func OccDemandSignalValueToMap(obj oci_demand_signal.OccDemandSignalValue) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Comments != nil {
		result["comments"] = string(*obj.Comments)
	}

	if obj.TimeExpected != nil {
		result["time_expected"] = obj.TimeExpected.Format(time.RFC3339Nano)
	}

	if obj.Value != nil {
		result["value"] = float32(*obj.Value)
	}

	return result
}
