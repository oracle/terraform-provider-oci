// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dataintegration

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_dataintegration "github.com/oracle/oci-go-sdk/v65/dataintegration"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataintegrationWorkspaceApplicationScheduleResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDataintegrationWorkspaceApplicationSchedule,
		Read:     readDataintegrationWorkspaceApplicationSchedule,
		Update:   updateDataintegrationWorkspaceApplicationSchedule,
		Delete:   deleteDataintegrationWorkspaceApplicationSchedule,
		Schema: map[string]*schema.Schema{
			// Required
			"application_key": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"identifier": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"workspace_id": {
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
			"frequency_details": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"model_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"CUSTOM",
								"DAILY",
								"HOURLY",
								"MONTHLY",
								"MONTHLY_RULE",
								"WEEKLY",
							}, true),
						},

						// Optional
						"custom_expression": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"day_of_week": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"days": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
						},
						"frequency": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"interval": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"time": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"hour": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"minute": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"second": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"week_of_month": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"is_daylight_adjustment_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"key": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"model_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"object_status": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"object_version": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"registry_metadata": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"aggregator_key": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"is_favorite": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"key": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"labels": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"registry_version": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"timezone": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"metadata": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"aggregator": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"identifier": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"key": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"aggregator_key": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"count_statistics": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"object_type_count_list": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"object_count": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"object_type": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
						"created_by": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"created_by_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"identifier_path": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"info_fields": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"is_favorite": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"labels": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"registry_version": {
							Type:     schema.TypeInt,
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
						"updated_by": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"updated_by_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"model_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_ref": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"parent": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"root_doc_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func createDataintegrationWorkspaceApplicationSchedule(d *schema.ResourceData, m interface{}) error {
	sync := &DataintegrationWorkspaceApplicationScheduleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataIntegrationClient()

	return tfresource.CreateResource(d, sync)
}

func readDataintegrationWorkspaceApplicationSchedule(d *schema.ResourceData, m interface{}) error {
	sync := &DataintegrationWorkspaceApplicationScheduleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataIntegrationClient()

	return tfresource.ReadResource(sync)
}

func updateDataintegrationWorkspaceApplicationSchedule(d *schema.ResourceData, m interface{}) error {
	sync := &DataintegrationWorkspaceApplicationScheduleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataIntegrationClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDataintegrationWorkspaceApplicationSchedule(d *schema.ResourceData, m interface{}) error {
	sync := &DataintegrationWorkspaceApplicationScheduleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataIntegrationClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DataintegrationWorkspaceApplicationScheduleResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_dataintegration.DataIntegrationClient
	Res                    *oci_dataintegration.Schedule
	DisableNotFoundRetries bool
}

func (s *DataintegrationWorkspaceApplicationScheduleResourceCrud) ID() string {
	return GetWorkspaceApplicationScheduleCompositeId(s.D.Get("application_key").(string), *s.Res.Key, s.D.Get("workspace_id").(string))
}

func (s *DataintegrationWorkspaceApplicationScheduleResourceCrud) Create() error {
	request := oci_dataintegration.CreateScheduleRequest{}

	if applicationKey, ok := s.D.GetOkExists("application_key"); ok {
		tmp := applicationKey.(string)
		request.ApplicationKey = &tmp
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if frequencyDetails, ok := s.D.GetOkExists("frequency_details"); ok {
		if tmpList := frequencyDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "frequency_details", 0)
			tmp, err := s.mapToAbstractFrequencyDetailsForSchedule(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.FrequencyDetails = tmp
		}
	}

	if identifier, ok := s.D.GetOkExists("identifier"); ok {
		tmp := identifier.(string)
		request.Identifier = &tmp
	}

	if isDaylightAdjustmentEnabled, ok := s.D.GetOkExists("is_daylight_adjustment_enabled"); ok {
		tmp := isDaylightAdjustmentEnabled.(bool)
		request.IsDaylightAdjustmentEnabled = &tmp
	}

	if key, ok := s.D.GetOkExists("key"); ok {
		tmp := key.(string)
		request.Key = &tmp
	}

	if modelVersion, ok := s.D.GetOkExists("model_version"); ok {
		tmp := modelVersion.(string)
		request.ModelVersion = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if objectStatus, ok := s.D.GetOkExists("object_status"); ok {
		tmp := objectStatus.(int)
		request.ObjectStatus = &tmp
	}

	if objectVersion, ok := s.D.GetOkExists("object_version"); ok {
		tmp := objectVersion.(int)
		request.ObjectVersion = &tmp
	}

	if registryMetadata, ok := s.D.GetOkExists("registry_metadata"); ok {
		if tmpList := registryMetadata.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "registry_metadata", 0)
			tmp, err := s.mapToRegistryMetadata(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.RegistryMetadata = &tmp
		}
	}

	if timezone, ok := s.D.GetOkExists("timezone"); ok {
		tmp := timezone.(string)
		request.Timezone = &tmp
	}

	if workspaceId, ok := s.D.GetOkExists("workspace_id"); ok {
		tmp := workspaceId.(string)
		request.WorkspaceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dataintegration")

	response, err := s.Client.CreateSchedule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Schedule
	return nil
}

func (s *DataintegrationWorkspaceApplicationScheduleResourceCrud) Get() error {
	request := oci_dataintegration.GetScheduleRequest{}

	if applicationKey, ok := s.D.GetOkExists("application_key"); ok {
		tmp := applicationKey.(string)
		request.ApplicationKey = &tmp
	}

	if scheduleKey, ok := s.D.GetOkExists("key"); ok {
		tmp := scheduleKey.(string)
		request.ScheduleKey = &tmp
	}

	if workspaceId, ok := s.D.GetOkExists("workspace_id"); ok {
		tmp := workspaceId.(string)
		request.WorkspaceId = &tmp
	}

	applicationKey, scheduleKey, workspaceId, err := parseWorkspaceApplicationScheduleCompositeId(s.D.Id())
	if err == nil {
		request.ApplicationKey = &applicationKey
		request.ScheduleKey = &scheduleKey
		request.WorkspaceId = &workspaceId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dataintegration")

	response, err := s.Client.GetSchedule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Schedule
	return nil
}

func (s *DataintegrationWorkspaceApplicationScheduleResourceCrud) Update() error {
	request := oci_dataintegration.UpdateScheduleRequest{}

	if applicationKey, ok := s.D.GetOkExists("application_key"); ok {
		tmp := applicationKey.(string)
		request.ApplicationKey = &tmp
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if frequencyDetails, ok := s.D.GetOkExists("frequency_details"); ok {
		if tmpList := frequencyDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "frequency_details", 0)
			tmp, err := s.mapToAbstractFrequencyDetailsForSchedule(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.FrequencyDetails = tmp
		}
	}

	if identifier, ok := s.D.GetOkExists("identifier"); ok {
		tmp := identifier.(string)
		request.Identifier = &tmp
	}

	if isDaylightAdjustmentEnabled, ok := s.D.GetOkExists("is_daylight_adjustment_enabled"); ok {
		tmp := isDaylightAdjustmentEnabled.(bool)
		request.IsDaylightAdjustmentEnabled = &tmp
	}

	if key, ok := s.D.GetOkExists("key"); ok {
		tmp := key.(string)
		request.Key = &tmp
	}

	if modelType, ok := s.D.GetOkExists("model_type"); ok {
		tmp := modelType.(string)
		request.ModelType = &tmp
	}

	if modelVersion, ok := s.D.GetOkExists("model_version"); ok {
		tmp := modelVersion.(string)
		request.ModelVersion = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if objectStatus, ok := s.D.GetOkExists("object_status"); ok {
		tmp := objectStatus.(int)
		request.ObjectStatus = &tmp
	}

	if objectVersion, ok := s.D.GetOkExists("object_version"); ok {
		tmp := objectVersion.(int)
		request.ObjectVersion = &tmp
	}

	if parentRef, ok := s.D.GetOkExists("parent_ref"); ok {
		if tmpList := parentRef.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "parent_ref", 0)
			tmp, err := s.mapToParentReference(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ParentRef = &tmp
		}
	}

	if registryMetadata, ok := s.D.GetOkExists("registry_metadata"); ok {
		if tmpList := registryMetadata.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "registry_metadata", 0)
			tmp, err := s.mapToRegistryMetadata(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.RegistryMetadata = &tmp
		}
	}

	if scheduleKey, ok := s.D.GetOkExists("key"); ok {
		tmp := scheduleKey.(string)
		request.ScheduleKey = &tmp
	}

	if timezone, ok := s.D.GetOkExists("timezone"); ok {
		tmp := timezone.(string)
		request.Timezone = &tmp
	}

	if workspaceId, ok := s.D.GetOkExists("workspace_id"); ok {
		tmp := workspaceId.(string)
		request.WorkspaceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dataintegration")

	response, err := s.Client.UpdateSchedule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Schedule
	return nil
}

func (s *DataintegrationWorkspaceApplicationScheduleResourceCrud) Delete() error {
	request := oci_dataintegration.DeleteScheduleRequest{}

	if applicationKey, ok := s.D.GetOkExists("application_key"); ok {
		tmp := applicationKey.(string)
		request.ApplicationKey = &tmp
	}

	if scheduleKey, ok := s.D.GetOkExists("key"); ok {
		tmp := scheduleKey.(string)
		request.ScheduleKey = &tmp
	}

	if workspaceId, ok := s.D.GetOkExists("workspace_id"); ok {
		tmp := workspaceId.(string)
		request.WorkspaceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dataintegration")

	_, err := s.Client.DeleteSchedule(context.Background(), request)
	return err
}

func (s *DataintegrationWorkspaceApplicationScheduleResourceCrud) SetData() error {

	applicationKey, scheduleKey, workspaceId, err := parseWorkspaceApplicationScheduleCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("application_key", &applicationKey)
		s.D.Set("key", &scheduleKey)
		s.D.Set("workspace_id", &workspaceId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.FrequencyDetails != nil {
		frequencyDetailsArray := []interface{}{}
		if frequencyDetailsMap := AbstractFrequencyDetailsToMapForSchedule(&s.Res.FrequencyDetails); frequencyDetailsMap != nil {
			frequencyDetailsArray = append(frequencyDetailsArray, frequencyDetailsMap)
		}
		s.D.Set("frequency_details", frequencyDetailsArray)
	} else {
		s.D.Set("frequency_details", nil)
	}

	if s.Res.Identifier != nil {
		s.D.Set("identifier", *s.Res.Identifier)
	}

	if s.Res.IsDaylightAdjustmentEnabled != nil {
		s.D.Set("is_daylight_adjustment_enabled", *s.Res.IsDaylightAdjustmentEnabled)
	}

	if s.Res.Key != nil {
		s.D.Set("key", *s.Res.Key)
	}

	if s.Res.Metadata != nil {
		s.D.Set("metadata", []interface{}{ObjectMetadataToMapForSchedule(s.Res.Metadata)})
	} else {
		s.D.Set("metadata", nil)
	}

	if s.Res.ModelType != nil {
		s.D.Set("model_type", *s.Res.ModelType)
	}

	if s.Res.ModelVersion != nil {
		s.D.Set("model_version", *s.Res.ModelVersion)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.ObjectStatus != nil {
		s.D.Set("object_status", *s.Res.ObjectStatus)
	}

	if s.Res.ObjectVersion != nil {
		s.D.Set("object_version", *s.Res.ObjectVersion)
	}

	if s.Res.ParentRef != nil {
		s.D.Set("parent_ref", []interface{}{ParentReferenceToMapForSchedule(s.Res.ParentRef)})
	} else {
		s.D.Set("parent_ref", nil)
	}

	if s.Res.Timezone != nil {
		s.D.Set("timezone", *s.Res.Timezone)
	}

	return nil
}

func GetWorkspaceApplicationScheduleCompositeId(applicationKey string, scheduleKey string, workspaceId string) string {
	applicationKey = url.PathEscape(applicationKey)
	scheduleKey = url.PathEscape(scheduleKey)
	workspaceId = url.PathEscape(workspaceId)
	compositeId := "workspaces/" + workspaceId + "/applications/" + applicationKey + "/schedules/" + scheduleKey
	return compositeId
}

func parseWorkspaceApplicationScheduleCompositeId(compositeId string) (applicationKey string, scheduleKey string, workspaceId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("workspaces/.*/applications/.*/schedules/.*", compositeId)
	if !match || len(parts) != 6 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	workspaceId, _ = url.PathUnescape(parts[1])
	applicationKey, _ = url.PathUnescape(parts[3])
	scheduleKey, _ = url.PathUnescape(parts[5])

	return
}

func AbstractFrequencyDetailsToMapForSchedule(obj *oci_dataintegration.AbstractFrequencyDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_dataintegration.CustomFrequencyDetails:
		result["model_type"] = "CUSTOM"

		if v.CustomExpression != nil {
			result["custom_expression"] = string(*v.CustomExpression)
		}
	case oci_dataintegration.DailyFrequencyDetails:
		result["model_type"] = "DAILY"

		if v.Interval != nil {
			result["interval"] = int(*v.Interval)
		}

		if v.Time != nil {
			result["time"] = []interface{}{TimeToMapForSchedule(v.Time)}
		}
	case oci_dataintegration.HourlyFrequencyDetails:
		result["model_type"] = "HOURLY"

		if v.Interval != nil {
			result["interval"] = int(*v.Interval)
		}

		if v.Time != nil {
			result["time"] = []interface{}{TimeToMapForSchedule(v.Time)}
		}
	case oci_dataintegration.MonthlyFrequencyDetails:
		result["model_type"] = "MONTHLY"

		result["days"] = v.Days
		result["days"] = v.Days

		if v.Interval != nil {
			result["interval"] = int(*v.Interval)
		}

		if v.Time != nil {
			result["time"] = []interface{}{TimeToMapForSchedule(v.Time)}
		}
	case oci_dataintegration.MonthlyRuleFrequencyDetails:
		result["model_type"] = "MONTHLY_RULE"

		result["day_of_week"] = string(v.DayOfWeek)

		if v.Interval != nil {
			result["interval"] = int(*v.Interval)
		}

		if v.Time != nil {
			result["time"] = []interface{}{TimeToMapForSchedule(v.Time)}
		}

		result["week_of_month"] = string(v.WeekOfMonth)
	case oci_dataintegration.WeeklyFrequencyDetails:
		result["model_type"] = "WEEKLY"

		result["days"] = v.Days
		result["days"] = v.Days

		if v.Time != nil {
			result["time"] = []interface{}{TimeToMapForSchedule(v.Time)}
		}
	default:
		log.Printf("[WARN] Received 'model_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func AggregatorSummaryToMapForSchedule(obj *oci_dataintegration.AggregatorSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.Identifier != nil {
		result["identifier"] = string(*obj.Identifier)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Type != nil {
		result["type"] = string(*obj.Type)
	}

	return result
}

func CountStatisticToMapForSchedule(obj *oci_dataintegration.CountStatistic) map[string]interface{} {
	result := map[string]interface{}{}

	objectTypeCountList := []interface{}{}
	for _, item := range obj.ObjectTypeCountList {
		objectTypeCountList = append(objectTypeCountList, CountStatisticSummaryToMapForSchedule(item))
	}
	result["object_type_count_list"] = objectTypeCountList

	return result
}

func CountStatisticSummaryToMapForSchedule(obj oci_dataintegration.CountStatisticSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ObjectCount != nil {
		result["object_count"] = strconv.FormatInt(*obj.ObjectCount, 10)
	}

	result["object_type"] = string(obj.ObjectType)

	return result
}

func ObjectMetadataToMapForSchedule(obj *oci_dataintegration.ObjectMetadata) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Aggregator != nil {
		result["aggregator"] = []interface{}{AggregatorSummaryToMapForSchedule(obj.Aggregator)}
	}

	if obj.AggregatorKey != nil {
		result["aggregator_key"] = string(*obj.AggregatorKey)
	}

	if obj.CountStatistics != nil {
		result["count_statistics"] = []interface{}{CountStatisticToMapForSchedule(obj.CountStatistics)}
	}

	if obj.CreatedBy != nil {
		result["created_by"] = string(*obj.CreatedBy)
	}

	if obj.CreatedByName != nil {
		result["created_by_name"] = string(*obj.CreatedByName)
	}

	if obj.IdentifierPath != nil {
		result["identifier_path"] = string(*obj.IdentifierPath)
	}

	result["info_fields"] = obj.InfoFields
	result["info_fields"] = obj.InfoFields

	if obj.IsFavorite != nil {
		result["is_favorite"] = bool(*obj.IsFavorite)
	}

	result["labels"] = obj.Labels
	result["labels"] = obj.Labels

	if obj.RegistryVersion != nil {
		result["registry_version"] = int(*obj.RegistryVersion)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	if obj.UpdatedBy != nil {
		result["updated_by"] = string(*obj.UpdatedBy)
	}

	if obj.UpdatedByName != nil {
		result["updated_by_name"] = string(*obj.UpdatedByName)
	}

	return result
}

func (s *DataintegrationWorkspaceApplicationScheduleResourceCrud) mapToParentReference(fieldKeyFormat string) (oci_dataintegration.ParentReference, error) {
	result := oci_dataintegration.ParentReference{}

	if parent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "parent")); ok {
		tmp := parent.(string)
		result.Parent = &tmp
	}

	if rootDocId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "root_doc_id")); ok {
		tmp := rootDocId.(string)
		result.RootDocId = &tmp
	}

	return result, nil
}

func ParentReferenceToMapForSchedule(obj *oci_dataintegration.ParentReference) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Parent != nil {
		result["parent"] = string(*obj.Parent)
	}

	if obj.RootDocId != nil {
		result["root_doc_id"] = string(*obj.RootDocId)
	}

	return result
}

func (s *DataintegrationWorkspaceApplicationScheduleResourceCrud) mapToRegistryMetadata(fieldKeyFormat string) (oci_dataintegration.RegistryMetadata, error) {
	result := oci_dataintegration.RegistryMetadata{}

	if aggregatorKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "aggregator_key")); ok {
		tmp := aggregatorKey.(string)
		result.AggregatorKey = &tmp
	}

	if isFavorite, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_favorite")); ok {
		tmp := isFavorite.(bool)
		result.IsFavorite = &tmp
	}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
	}

	if labels, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "labels")); ok {
		interfaces := labels.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "labels")) {
			result.Labels = tmp
		}
	}

	if registryVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "registry_version")); ok {
		tmp := registryVersion.(int)
		result.RegistryVersion = &tmp
	}

	return result, nil
}

func RegistryMetadataToMapForSchedule(obj *oci_dataintegration.RegistryMetadata) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AggregatorKey != nil {
		result["aggregator_key"] = string(*obj.AggregatorKey)
	}

	if obj.IsFavorite != nil {
		result["is_favorite"] = bool(*obj.IsFavorite)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	result["labels"] = obj.Labels
	result["labels"] = obj.Labels

	if obj.RegistryVersion != nil {
		result["registry_version"] = int(*obj.RegistryVersion)
	}

	return result
}

func ScheduleSummaryToMap(obj oci_dataintegration.ScheduleSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.FrequencyDetails != nil {
		frequencyDetailsArray := []interface{}{}
		if frequencyDetailsMap := AbstractFrequencyDetailsToMapForSchedule(&obj.FrequencyDetails); frequencyDetailsMap != nil {
			frequencyDetailsArray = append(frequencyDetailsArray, frequencyDetailsMap)
		}
		result["frequency_details"] = frequencyDetailsArray
	}

	if obj.Identifier != nil {
		result["identifier"] = string(*obj.Identifier)
	}

	if obj.IsDaylightAdjustmentEnabled != nil {
		result["is_daylight_adjustment_enabled"] = bool(*obj.IsDaylightAdjustmentEnabled)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.Metadata != nil {
		result["metadata"] = []interface{}{ObjectMetadataToMapForSchedule(obj.Metadata)}
	}

	if obj.ModelType != nil {
		result["model_type"] = string(*obj.ModelType)
	}

	if obj.ModelVersion != nil {
		result["model_version"] = string(*obj.ModelVersion)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.ObjectStatus != nil {
		result["object_status"] = int(*obj.ObjectStatus)
	}

	if obj.ObjectVersion != nil {
		result["object_version"] = int(*obj.ObjectVersion)
	}

	if obj.ParentRef != nil {
		result["parent_ref"] = []interface{}{ParentReferenceToMapForSchedule(obj.ParentRef)}
	}

	if obj.Timezone != nil {
		result["timezone"] = string(*obj.Timezone)
	}

	return result
}

func (s *DataintegrationWorkspaceApplicationScheduleResourceCrud) mapToTime(fieldKeyFormat string) (oci_dataintegration.Time, error) {
	result := oci_dataintegration.Time{}

	if hour, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "hour")); ok {
		tmp := hour.(int)
		result.Hour = &tmp
	}

	if minute, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "minute")); ok {
		tmp := minute.(int)
		result.Minute = &tmp
	}

	if second, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "second")); ok {
		tmp := second.(int)
		result.Second = &tmp
	}

	return result, nil
}

func TimeToMapForSchedule(obj *oci_dataintegration.Time) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Hour != nil {
		result["hour"] = int(*obj.Hour)
	}

	if obj.Minute != nil {
		result["minute"] = int(*obj.Minute)
	}

	if obj.Second != nil {
		result["second"] = int(*obj.Second)
	}

	return result
}

func (s *DataintegrationWorkspaceApplicationScheduleResourceCrud) mapToAbstractFrequencyDetailsForSchedule(fieldKeyFormat string) (oci_dataintegration.AbstractFrequencyDetails, error) {
	var baseObject oci_dataintegration.AbstractFrequencyDetails
	//discriminator
	modelTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "model_type"))
	var modelType string
	if ok {
		modelType = modelTypeRaw.(string)
	} else {
		modelType = "" // default value
	}
	switch strings.ToLower(modelType) {
	case strings.ToLower("CUSTOM"):
		details := oci_dataintegration.CustomFrequencyDetails{}
		if customExpression, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "custom_expression")); ok {
			tmp := customExpression.(string)
			details.CustomExpression = &tmp
		}
		if frequency, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "frequency")); ok {
			details.Frequency = oci_dataintegration.AbstractFrequencyDetailsFrequencyEnum(frequency.(string))
		}
		baseObject = details
	case strings.ToLower("DAILY"):
		details := oci_dataintegration.DailyFrequencyDetails{}
		if interval, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "interval")); ok {
			tmp := interval.(int)
			details.Interval = &tmp
		}
		if time, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time")); ok {
			if tmpList := time.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "time"), 0)
				tmp, err := s.mapToTime(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert time, encountered error: %v", err)
				}
				details.Time = &tmp
			}
		}
		if frequency, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "frequency")); ok {
			details.Frequency = oci_dataintegration.AbstractFrequencyDetailsFrequencyEnum(frequency.(string))
		}
		baseObject = details
	case strings.ToLower("HOURLY"):
		details := oci_dataintegration.HourlyFrequencyDetails{}
		if interval, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "interval")); ok {
			tmp := interval.(int)
			details.Interval = &tmp
		}
		if time, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time")); ok {
			if tmpList := time.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "time"), 0)
				tmp, err := s.mapToTime(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert time, encountered error: %v", err)
				}
				details.Time = &tmp
			}
		}
		if frequency, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "frequency")); ok {
			details.Frequency = oci_dataintegration.AbstractFrequencyDetailsFrequencyEnum(frequency.(string))
		}
		baseObject = details
	case strings.ToLower("MONTHLY"):
		details := oci_dataintegration.MonthlyFrequencyDetails{}
		if days, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "days")); ok {
			interfaces := days.([]interface{})
			tmp := make([]int, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(int)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "days")) {
				details.Days = tmp
			}
		}
		if interval, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "interval")); ok {
			tmp := interval.(int)
			details.Interval = &tmp
		}
		if time, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time")); ok {
			if tmpList := time.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "time"), 0)
				tmp, err := s.mapToTime(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert time, encountered error: %v", err)
				}
				details.Time = &tmp
			}
		}
		if frequency, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "frequency")); ok {
			details.Frequency = oci_dataintegration.AbstractFrequencyDetailsFrequencyEnum(frequency.(string))
		}
		baseObject = details
	case strings.ToLower("MONTHLY_RULE"):
		details := oci_dataintegration.MonthlyRuleFrequencyDetails{}
		if dayOfWeek, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "day_of_week")); ok {
			details.DayOfWeek = oci_dataintegration.MonthlyRuleFrequencyDetailsDayOfWeekEnum(dayOfWeek.(string))
		}
		if interval, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "interval")); ok {
			tmp := interval.(int)
			details.Interval = &tmp
		}
		if time, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time")); ok {
			if tmpList := time.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "time"), 0)
				tmp, err := s.mapToTime(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert time, encountered error: %v", err)
				}
				details.Time = &tmp
			}
		}
		if weekOfMonth, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "week_of_month")); ok {
			details.WeekOfMonth = oci_dataintegration.MonthlyRuleFrequencyDetailsWeekOfMonthEnum(weekOfMonth.(string))
		}
		if frequency, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "frequency")); ok {
			details.Frequency = oci_dataintegration.AbstractFrequencyDetailsFrequencyEnum(frequency.(string))
		}
		baseObject = details
	case strings.ToLower("WEEKLY"):
		details := oci_dataintegration.WeeklyFrequencyDetails{}
		if days, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "days")); ok {
			interfaces := days.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "days")) {
				details.Days = make([]oci_dataintegration.WeeklyFrequencyDetailsDaysEnum, len(tmp))
				for i := range tmp {
					details.Days[i] = oci_dataintegration.WeeklyFrequencyDetailsDaysEnum(tmp[i])
				}
			}
		}
		if time, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time")); ok {
			if tmpList := time.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "time"), 0)
				tmp, err := s.mapToTime(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert time, encountered error: %v", err)
				}
				details.Time = &tmp
			}
		}
		if frequency, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "frequency")); ok {
			details.Frequency = oci_dataintegration.AbstractFrequencyDetailsFrequencyEnum(frequency.(string))
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown model_type '%v' was specified", modelType)
	}
	return baseObject, nil
}
