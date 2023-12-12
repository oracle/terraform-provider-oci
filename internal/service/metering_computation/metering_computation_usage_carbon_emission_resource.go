// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package metering_computation

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_metering_computation "github.com/oracle/oci-go-sdk/v65/usageapi"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func MeteringComputationUsageCarbonEmissionResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createMeteringComputationUsageCarbonEmission,
		Read:     readMeteringComputationUsageCarbonEmission,
		Delete:   deleteMeteringComputationUsageCarbonEmission,
		Schema: map[string]*schema.Schema{
			// Required
			"tenant_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"time_usage_ended": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
			},
			"time_usage_started": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
			},

			// Optional
			"compartment_depth": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
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
			"usage_carbon_emission_filter": {
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
						"computed_carbon_emission": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"emission_calculation_method": {
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
					},
				},
			},
		},
	}
}

func createMeteringComputationUsageCarbonEmission(d *schema.ResourceData, m interface{}) error {
	sync := &MeteringComputationUsageCarbonEmissionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).UsageapiClient()

	return tfresource.CreateResource(d, sync)
}

func readMeteringComputationUsageCarbonEmission(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteMeteringComputationUsageCarbonEmission(d *schema.ResourceData, m interface{}) error {
	return nil
}

type MeteringComputationUsageCarbonEmissionResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_metering_computation.UsageapiClient
	Res                    *oci_metering_computation.UsageCarbonEmissionAggregation
	DisableNotFoundRetries bool
}

func (s *MeteringComputationUsageCarbonEmissionResourceCrud) ID() string {
	return tfresource.Timestamp()
}

func (s *MeteringComputationUsageCarbonEmissionResourceCrud) Create() error {
	request := oci_metering_computation.RequestUsageCarbonEmissionsRequest{}

	if compartmentDepth, ok := s.D.GetOkExists("compartment_depth"); ok {
		tmp := compartmentDepth.(int)
		request.CompartmentDepth = &tmp
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

	if usageCarbonEmissionFilter, ok := s.D.GetOkExists("usage_carbon_emission_filter"); ok {
		tmp := usageCarbonEmissionFilter.(string)
		var usageCarbonEmission_filterObj oci_metering_computation.Filter
		err := json.Unmarshal([]byte(tmp), &usageCarbonEmission_filterObj)
		if err != nil {
			return err
		}
		request.Filter = &usageCarbonEmission_filterObj
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "metering_computation")

	response, err := s.Client.RequestUsageCarbonEmissions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.UsageCarbonEmissionAggregation
	return nil
}

func (s *MeteringComputationUsageCarbonEmissionResourceCrud) SetData() error {
	s.D.Set("group_by", s.Res.GroupBy)

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, UsageCarbonEmissionSummaryToMap(item))
	}
	s.D.Set("items", items)

	return nil
}

func (s *MeteringComputationUsageCarbonEmissionResourceCrud) mapToTag(fieldKeyFormat string) (oci_metering_computation.Tag, error) {
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

func UsageCarbonEmissionSummaryToMap(obj oci_metering_computation.UsageCarbonEmissionSummary) map[string]interface{} {
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

	if obj.ComputedCarbonEmission != nil {
		result["computed_carbon_emission"] = float64(*obj.ComputedCarbonEmission)
	}

	if obj.EmissionCalculationMethod != nil {
		result["emission_calculation_method"] = string(*obj.EmissionCalculationMethod)
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
		tags = append(tags, TagToMap(item))
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

	return result
}
