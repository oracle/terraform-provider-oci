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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_apm_synthetics "github.com/oracle/oci-go-sdk/v65/apmsynthetics"
)

func ApmSyntheticsDedicatedVantagePointResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createApmSyntheticsDedicatedVantagePoint,
		Read:     readApmSyntheticsDedicatedVantagePoint,
		Update:   updateApmSyntheticsDedicatedVantagePoint,
		Delete:   deleteApmSyntheticsDedicatedVantagePoint,
		Schema: map[string]*schema.Schema{
			// Required
			"apm_domain_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"dvp_stack_details": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"dvp_stack_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"dvp_stack_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"ORACLE_RM_STACK",
							}, true),
						},
						"dvp_stream_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"dvp_version": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"region": {
				Type:     schema.TypeString,
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
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
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
			"name": {
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

func createApmSyntheticsDedicatedVantagePoint(d *schema.ResourceData, m interface{}) error {
	sync := &ApmSyntheticsDedicatedVantagePointResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApmSyntheticClient()

	return tfresource.CreateResource(d, sync)
}

func readApmSyntheticsDedicatedVantagePoint(d *schema.ResourceData, m interface{}) error {
	sync := &ApmSyntheticsDedicatedVantagePointResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApmSyntheticClient()

	return tfresource.ReadResource(sync)
}

func updateApmSyntheticsDedicatedVantagePoint(d *schema.ResourceData, m interface{}) error {
	sync := &ApmSyntheticsDedicatedVantagePointResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApmSyntheticClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteApmSyntheticsDedicatedVantagePoint(d *schema.ResourceData, m interface{}) error {
	sync := &ApmSyntheticsDedicatedVantagePointResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApmSyntheticClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type ApmSyntheticsDedicatedVantagePointResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_apm_synthetics.ApmSyntheticClient
	Res                    *oci_apm_synthetics.DedicatedVantagePoint
	DisableNotFoundRetries bool
}

func (s *ApmSyntheticsDedicatedVantagePointResourceCrud) ID() string {
	return GetDedicatedVantagePointCompositeId(*s.Res.Id, s.D.Get("apm_domain_id").(string))
}

func (s *ApmSyntheticsDedicatedVantagePointResourceCrud) Create() error {
	request := oci_apm_synthetics.CreateDedicatedVantagePointRequest{}

	if apmDomainId, ok := s.D.GetOkExists("apm_domain_id"); ok {
		tmp := apmDomainId.(string)
		request.ApmDomainId = &tmp
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

	if dvpStackDetails, ok := s.D.GetOkExists("dvp_stack_details"); ok {
		if tmpList := dvpStackDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "dvp_stack_details", 0)
			tmp, err := s.mapToDvpStackDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.DvpStackDetails = tmp
		}
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if region, ok := s.D.GetOkExists("region"); ok {
		tmp := region.(string)
		request.Region = &tmp
	}

	if status, ok := s.D.GetOkExists("status"); ok {
		request.Status = oci_apm_synthetics.DedicatedVantagePointStatusEnum(status.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apm_synthetics")

	response, err := s.Client.CreateDedicatedVantagePoint(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DedicatedVantagePoint
	return nil
}

func (s *ApmSyntheticsDedicatedVantagePointResourceCrud) Get() error {
	request := oci_apm_synthetics.GetDedicatedVantagePointRequest{}

	if apmDomainId, ok := s.D.GetOkExists("apm_domain_id"); ok {
		tmp := apmDomainId.(string)
		request.ApmDomainId = &tmp
	}

	tmp := s.D.Id()
	request.DedicatedVantagePointId = &tmp

	dedicatedVantagePointId, apmDomainId, err := parseDedicatedVantagePointCompositeId(s.D.Id())
	if err == nil {
		request.DedicatedVantagePointId = &dedicatedVantagePointId
		request.ApmDomainId = &apmDomainId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s apmDomainId: %s", s.D.Id(), apmDomainId)
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apm_synthetics")

	response, err := s.Client.GetDedicatedVantagePoint(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DedicatedVantagePoint
	return nil
}

func (s *ApmSyntheticsDedicatedVantagePointResourceCrud) Update() error {
	request := oci_apm_synthetics.UpdateDedicatedVantagePointRequest{}

	if apmDomainId, ok := s.D.GetOkExists("apm_domain_id"); ok {
		tmp := apmDomainId.(string)
		request.ApmDomainId = &tmp
	}

	tmp := s.D.Id()
	request.DedicatedVantagePointId = &tmp

	dedicatedVantagePointId, apmDomainId, err := parseDedicatedVantagePointCompositeId(s.D.Id())
	if err == nil {
		request.DedicatedVantagePointId = &dedicatedVantagePointId
		request.ApmDomainId = &apmDomainId
	} else {
		log.Printf("[WARN] Update() unable to parse current ID: %s apmDomainId: %s", s.D.Id(), apmDomainId)
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if dvpStackDetails, ok := s.D.GetOkExists("dvp_stack_details"); ok {
		if tmpList := dvpStackDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "dvp_stack_details", 0)
			tmp, err := s.mapToDvpStackDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.DvpStackDetails = tmp
		}
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if region, ok := s.D.GetOkExists("region"); ok {
		tmp := region.(string)
		request.Region = &tmp
	}

	if status, ok := s.D.GetOkExists("status"); ok {
		request.Status = oci_apm_synthetics.DedicatedVantagePointStatusEnum(status.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apm_synthetics")

	response, err := s.Client.UpdateDedicatedVantagePoint(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DedicatedVantagePoint
	return nil
}

func (s *ApmSyntheticsDedicatedVantagePointResourceCrud) Delete() error {
	request := oci_apm_synthetics.DeleteDedicatedVantagePointRequest{}

	if apmDomainId, ok := s.D.GetOkExists("apm_domain_id"); ok {
		tmp := apmDomainId.(string)
		request.ApmDomainId = &tmp
	}

	tmp := s.D.Id()
	request.DedicatedVantagePointId = &tmp

	dedicatedVantagePointId, apmDomainId, err1 := parseDedicatedVantagePointCompositeId(s.D.Id())
	if err1 == nil {
		request.DedicatedVantagePointId = &dedicatedVantagePointId
	} else {
		log.Printf("[WARN] Delete() unable to parse current ID: %s apmDomainId: %s", s.D.Id(), apmDomainId)
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apm_synthetics")

	_, err := s.Client.DeleteDedicatedVantagePoint(context.Background(), request)
	return err
}

func (s *ApmSyntheticsDedicatedVantagePointResourceCrud) SetData() error {

	dedicatedVantagePointId, apmDomainId, err := parseDedicatedVantagePointCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("dedicated_vantage_point_id", dedicatedVantagePointId)
		if apmDomainId != "" {

		}
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.DvpStackDetails != nil {
		dvpStackDetailsArray := []interface{}{}
		if dvpStackDetailsMap := DvpStackDetailsToMap(&s.Res.DvpStackDetails); dvpStackDetailsMap != nil {
			dvpStackDetailsArray = append(dvpStackDetailsArray, dvpStackDetailsMap)
		}
		s.D.Set("dvp_stack_details", dvpStackDetailsArray)
	} else {
		s.D.Set("dvp_stack_details", nil)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.MonitorStatusCountMap != nil {
		s.D.Set("monitor_status_count_map", []interface{}{MonitorStatusCountMapToMap(s.Res.MonitorStatusCountMap)})
	} else {
		s.D.Set("monitor_status_count_map", nil)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.Region != nil {
		s.D.Set("region", *s.Res.Region)
	}

	s.D.Set("status", s.Res.Status)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func GetDedicatedVantagePointCompositeId(dedicatedVantagePointId string, apmDomainId string) string {
	apmDomainId = url.PathEscape(apmDomainId)
	dedicatedVantagePointId = url.PathEscape(dedicatedVantagePointId)
	compositeId := "dedicatedVantagePoints/" + dedicatedVantagePointId + "/apmDomainId/" + apmDomainId
	return compositeId
}

func parseDedicatedVantagePointCompositeId(compositeId string) (dedicatedVantagePointId string, apmDomainId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("dedicatedVantagePoints/.*/apmDomainId/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	dedicatedVantagePointId, _ = url.PathUnescape(parts[1])
	apmDomainId, _ = url.PathUnescape(parts[3])

	return
}

func DedicatedVantagePointSummaryToMap(obj oci_apm_synthetics.DedicatedVantagePointSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.DvpStackDetails != nil {
		dvpStackDetailsArray := []interface{}{}
		if dvpStackDetailsMap := DvpStackDetailsToMap(&obj.DvpStackDetails); dvpStackDetailsMap != nil {
			dvpStackDetailsArray = append(dvpStackDetailsArray, dvpStackDetailsMap)
		}
		result["dvp_stack_details"] = dvpStackDetailsArray
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.MonitorStatusCountMap != nil {
		result["monitor_status_count_map"] = []interface{}{MonitorStatusCountMapToMap(obj.MonitorStatusCountMap)}
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Region != nil {
		result["region"] = string(*obj.Region)
	}

	result["status"] = string(obj.Status)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func DvpStackDetailsToMap(obj *oci_apm_synthetics.DvpStackDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_apm_synthetics.OracleRmStack:
		result["dvp_stack_type"] = "ORACLE_RM_STACK"

		if v.DvpStackId != nil {
			result["dvp_stack_id"] = string(*v.DvpStackId)
		}

		if v.DvpStreamId != nil {
			result["dvp_stream_id"] = string(*v.DvpStreamId)
		}

		if v.DvpVersion != nil {
			result["dvp_version"] = string(*v.DvpVersion)
		}
	default:
		log.Printf("[WARN] Received 'dvp_stack_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *ApmSyntheticsDedicatedVantagePointResourceCrud) mapToDvpStackDetails(fieldKeyFormat string) (oci_apm_synthetics.OracleRmStack, error) {
	result := oci_apm_synthetics.OracleRmStack{}

	if dvpVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "dvp_version")); ok {
		tmp := dvpVersion.(string)
		result.DvpVersion = &tmp
	}

	if dvpStackId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "dvp_stack_id")); ok {
		tmp := dvpStackId.(string)
		result.DvpStackId = &tmp
	}

	if dvpStreamId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "dvp_stream_id")); ok {
		tmp := dvpStreamId.(string)
		result.DvpStreamId = &tmp
	}

	return result, nil
}
