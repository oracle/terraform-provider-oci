// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package metering_computation

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/hashicorp/terraform-plugin-sdk/helper/structure"
	oci_common "github.com/oracle/oci-go-sdk/v56/common"
	oci_metering_computation "github.com/oracle/oci-go-sdk/v56/usageapi"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

func MeteringComputationUsageResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		Create:   createMeteringComputationUsage,
		Read:     readMeteringComputationUsage,
		Delete:   deleteMeteringComputationUsage,
		Schema: map[string]*schema.Schema{
			// Required
			"granularity": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"tenant_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"time_usage_ended": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: utils.TimeDiffSuppressFunction,
			},
			"time_usage_started": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: utils.TimeDiffSuppressFunction,
			},

			// Optional
			"compartment_depth": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"filter": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validateFilterJson,
			},
			"forecast": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"time_forecast_ended": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: utils.TimeDiffSuppressFunction,
						},

						// Optional
						"forecast_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"time_forecast_started": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ForceNew:         true,
							DiffSuppressFunc: utils.TimeDiffSuppressFunction,
						},

						// Computed
					},
				},
			},
			"group_by": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"group_by_tag": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"key": {
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
						"value": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"is_aggregate_by_time": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"query_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"items": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"ad": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"compartment_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"compartment_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"compartment_path": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"computed_amount": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"computed_quantity": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"currency": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"discount": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"is_forecast": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"list_rate": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"overage": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"overages_flag": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"platform": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"region": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"resource_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"resource_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"service": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"shape": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"sku_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"sku_part_number": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"subscription_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"tags": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"key": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"namespace": {
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
						"tenant_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"tenant_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_usage_ended": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_usage_started": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"unit": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"unit_price": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"weight": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func createMeteringComputationUsage(d *schema.ResourceData, m interface{}) error {
	sync := &MeteringComputationUsageResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).UsageapiClient()

	return tfresource.CreateResource(d, sync)
}

func readMeteringComputationUsage(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteMeteringComputationUsage(d *schema.ResourceData, m interface{}) error {
	return nil
}

type MeteringComputationUsageResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_metering_computation.UsageapiClient
	Res                    *oci_metering_computation.UsageAggregation
	DisableNotFoundRetries bool
}

func (s *MeteringComputationUsageResourceCrud) ID() string {
	return utils.Timestamp()
}

func (s *MeteringComputationUsageResourceCrud) Create() error {
	request := oci_metering_computation.RequestSummarizedUsagesRequest{}

	if compartmentDepth, ok := s.D.GetOkExists("compartment_depth"); ok {
		tmp := float32(compartmentDepth.(float64))
		request.CompartmentDepth = &tmp
	}

	if filter, ok := s.D.GetOkExists("filter"); ok {
		tmp := filter.(string)
		var filterObj oci_metering_computation.Filter
		err := json.Unmarshal([]byte(tmp), &filterObj)
		if err != nil {
			return err
		}
		request.Filter = &filterObj
	}

	if forecast, ok := s.D.GetOkExists("forecast"); ok {
		if tmpList := forecast.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "forecast", 0)
			tmp, err := s.mapToForecast(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Forecast = &tmp
		}
	}

	if granularity, ok := s.D.GetOkExists("granularity"); ok {
		request.Granularity = oci_metering_computation.RequestSummarizedUsagesDetailsGranularityEnum(granularity.(string))
	}

	if groupBy, ok := s.D.GetOkExists("group_by"); ok {
		interfaces := groupBy.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("group_by") {
			request.GroupBy = tmp
		}
	}

	if groupByTag, ok := s.D.GetOkExists("group_by_tag"); ok {
		interfaces := groupByTag.([]interface{})
		tmp := make([]oci_metering_computation.Tag, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "group_by_tag", stateDataIndex)
			converted, err := s.mapToTag(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("group_by_tag") {
			request.GroupByTag = tmp
		}
	}

	if isAggregateByTime, ok := s.D.GetOkExists("is_aggregate_by_time"); ok {
		tmp := isAggregateByTime.(bool)
		request.IsAggregateByTime = &tmp
	}

	if queryType, ok := s.D.GetOkExists("query_type"); ok {
		request.QueryType = oci_metering_computation.RequestSummarizedUsagesDetailsQueryTypeEnum(queryType.(string))
	}

	if tenantId, ok := s.D.GetOkExists("tenant_id"); ok {
		tmp := tenantId.(string)
		request.TenantId = &tmp
	}

	if timeUsageEnded, ok := s.D.GetOkExists("time_usage_ended"); ok {
		tmp, err := time.Parse(time.RFC3339, timeUsageEnded.(string))
		if err != nil {
			return err
		}
		request.TimeUsageEnded = &oci_common.SDKTime{Time: tmp}
	}

	if timeUsageStarted, ok := s.D.GetOkExists("time_usage_started"); ok {
		tmp, err := time.Parse(time.RFC3339, timeUsageStarted.(string))
		if err != nil {
			return err
		}
		request.TimeUsageStarted = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "metering_computation")

	response, err := s.Client.RequestSummarizedUsages(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.UsageAggregation
	return nil
}

func (s *MeteringComputationUsageResourceCrud) SetData() error {
	s.D.Set("group_by", s.Res.GroupBy)

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, UsageSummaryToMap(item))
	}
	s.D.Set("items", items)

	return nil
}

func (s *MeteringComputationUsageResourceCrud) mapToForecast(fieldKeyFormat string) (oci_metering_computation.Forecast, error) {
	result := oci_metering_computation.Forecast{}

	if forecastType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "forecast_type")); ok {
		result.ForecastType = oci_metering_computation.ForecastForecastTypeEnum(forecastType.(string))
	}

	if timeForecastEnded, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_forecast_ended")); ok {
		tmp, err := time.Parse(time.RFC3339, timeForecastEnded.(string))
		if err != nil {
			return result, err
		}
		result.TimeForecastEnded = &oci_common.SDKTime{Time: tmp}
	}

	if timeForecastStarted, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_forecast_started")); ok {
		tmp, err := time.Parse(time.RFC3339, timeForecastStarted.(string))
		if err != nil {
			return result, err
		}
		result.TimeForecastStarted = &oci_common.SDKTime{Time: tmp}
	}

	return result, nil
}
func (s *MeteringComputationUsageResourceCrud) mapToTag(fieldKeyFormat string) (oci_metering_computation.Tag, error) {
	result := oci_metering_computation.Tag{}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
	}

	if namespace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
		tmp := namespace.(string)
		result.Namespace = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}
	return result, nil
}

func TagToMapInUsage(obj oci_metering_computation.Tag) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.Namespace != nil {
		result["namespace"] = string(*obj.Namespace)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func UsageSummaryToMap(obj oci_metering_computation.UsageSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Ad != nil {
		result["ad"] = string(*obj.Ad)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.CompartmentName != nil {
		result["compartment_name"] = string(*obj.CompartmentName)
	}

	if obj.CompartmentPath != nil {
		result["compartment_path"] = string(*obj.CompartmentPath)
	}

	if obj.ComputedAmount != nil {
		result["computed_amount"] = float32(*obj.ComputedAmount)
	}

	if obj.ComputedQuantity != nil {
		result["computed_quantity"] = float32(*obj.ComputedQuantity)
	}

	if obj.Currency != nil {
		result["currency"] = string(*obj.Currency)
	}

	if obj.Discount != nil {
		result["discount"] = float32(*obj.Discount)
	}

	if obj.IsForecast != nil {
		result["is_forecast"] = bool(*obj.IsForecast)
	}

	if obj.ListRate != nil {
		result["list_rate"] = float32(*obj.ListRate)
	}

	if obj.Overage != nil {
		result["overage"] = string(*obj.Overage)
	}

	if obj.OveragesFlag != nil {
		result["overages_flag"] = string(*obj.OveragesFlag)
	}

	if obj.Platform != nil {
		result["platform"] = string(*obj.Platform)
	}

	if obj.Region != nil {
		result["region"] = string(*obj.Region)
	}

	if obj.ResourceId != nil {
		result["resource_id"] = string(*obj.ResourceId)
	}

	if obj.ResourceName != nil {
		result["resource_name"] = string(*obj.ResourceName)
	}

	if obj.Service != nil {
		result["service"] = string(*obj.Service)
	}

	if obj.Shape != nil {
		result["shape"] = string(*obj.Shape)
	}

	if obj.SkuName != nil {
		result["sku_name"] = string(*obj.SkuName)
	}

	if obj.SkuPartNumber != nil {
		result["sku_part_number"] = string(*obj.SkuPartNumber)
	}

	if obj.SubscriptionId != nil {
		result["subscription_id"] = string(*obj.SubscriptionId)
	}

	tags := []interface{}{}
	for _, item := range obj.Tags {
		tags = append(tags, TagToMapInUsage(item))
	}
	result["tags"] = tags

	if obj.TenantId != nil {
		result["tenant_id"] = string(*obj.TenantId)
	}

	if obj.TenantName != nil {
		result["tenant_name"] = string(*obj.TenantName)
	}

	if obj.TimeUsageEnded != nil {
		result["time_usage_ended"] = obj.TimeUsageEnded.String()
	}

	if obj.TimeUsageStarted != nil {
		result["time_usage_started"] = obj.TimeUsageStarted.String()
	}

	if obj.Unit != nil {
		result["unit"] = string(*obj.Unit)
	}

	if obj.UnitPrice != nil {
		result["unit_price"] = float32(*obj.UnitPrice)
	}

	if obj.Weight != nil {
		result["weight"] = float32(*obj.Weight)
	}

	return result
}

func validateFilterJson(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)
	if len(value) < 1 {
		errors = append(errors, fmt.Errorf("%q contains an invalid JSON policy", k))
		return
	}
	if value[:1] != "{" {
		errors = append(errors, fmt.Errorf("%q contains an invalid JSON policy", k))
		return
	}
	if _, err := structure.NormalizeJsonString(v); err != nil {
		errors = append(errors, fmt.Errorf("%q contains an invalid JSON: %s", k, err))
	}
	return
}
