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

	oci_apm_synthetics "github.com/oracle/oci-go-sdk/v65/apmsynthetics"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ApmSyntheticsOnPremiseVantagePointResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createApmSyntheticsOnPremiseVantagePoint,
		Read:     readApmSyntheticsOnPremiseVantagePoint,
		Update:   updateApmSyntheticsOnPremiseVantagePoint,
		Delete:   deleteApmSyntheticsOnPremiseVantagePoint,
		Schema: map[string]*schema.Schema{
			// Required
			"apm_domain_id": {
				Type:     schema.TypeString,
				Required: true,
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
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"display_name": {
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
			"workers_summary": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"available": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"available_capabilities": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"capability": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"on_premise_vantage_point_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"disabled": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"min_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"total": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"used": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func createApmSyntheticsOnPremiseVantagePoint(d *schema.ResourceData, m interface{}) error {
	sync := &ApmSyntheticsOnPremiseVantagePointResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApmSyntheticClient()

	return tfresource.CreateResource(d, sync)
}

func readApmSyntheticsOnPremiseVantagePoint(d *schema.ResourceData, m interface{}) error {
	sync := &ApmSyntheticsOnPremiseVantagePointResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApmSyntheticClient()

	return tfresource.ReadResource(sync)
}

func updateApmSyntheticsOnPremiseVantagePoint(d *schema.ResourceData, m interface{}) error {
	sync := &ApmSyntheticsOnPremiseVantagePointResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApmSyntheticClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteApmSyntheticsOnPremiseVantagePoint(d *schema.ResourceData, m interface{}) error {
	sync := &ApmSyntheticsOnPremiseVantagePointResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApmSyntheticClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type ApmSyntheticsOnPremiseVantagePointResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_apm_synthetics.ApmSyntheticClient
	Res                    *oci_apm_synthetics.OnPremiseVantagePoint
	DisableNotFoundRetries bool
}

func (s *ApmSyntheticsOnPremiseVantagePointResourceCrud) ID() string {
	return GetOnPremiseVantagePointCompositeId(*s.Res.Id, s.D.Get("apm_domain_id").(string))
}

func (s *ApmSyntheticsOnPremiseVantagePointResourceCrud) Create() error {
	request := oci_apm_synthetics.CreateOnPremiseVantagePointRequest{}

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

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if type_, ok := s.D.GetOkExists("type"); ok {
		request.Type = oci_apm_synthetics.CreateOnPremiseVantagePointDetailsTypeEnum(type_.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apm_synthetics")

	response, err := s.Client.CreateOnPremiseVantagePoint(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.OnPremiseVantagePoint
	return nil
}

func (s *ApmSyntheticsOnPremiseVantagePointResourceCrud) Get() error {
	request := oci_apm_synthetics.GetOnPremiseVantagePointRequest{}

	if apmDomainId, ok := s.D.GetOkExists("apm_domain_id"); ok {
		tmp := apmDomainId.(string)
		request.ApmDomainId = &tmp
	}

	onPremiseVantagePointId, apmDomainId, err := parseOnPremiseVantagePointCompositeId(s.D.Id())
	if err == nil {
		request.OnPremiseVantagePointId = &onPremiseVantagePointId
		request.ApmDomainId = &apmDomainId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apm_synthetics")

	response, err := s.Client.GetOnPremiseVantagePoint(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.OnPremiseVantagePoint
	return nil
}

func (s *ApmSyntheticsOnPremiseVantagePointResourceCrud) Update() error {
	request := oci_apm_synthetics.UpdateOnPremiseVantagePointRequest{}

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

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	onPremiseVantagePointId, apmDomainId, err := parseOnPremiseVantagePointCompositeId(s.D.Id())
	if err == nil {
		request.OnPremiseVantagePointId = &onPremiseVantagePointId
		request.ApmDomainId = &apmDomainId
	} else {
		log.Printf("[WARN] Update() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apm_synthetics")

	response, err := s.Client.UpdateOnPremiseVantagePoint(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.OnPremiseVantagePoint
	return nil
}

func (s *ApmSyntheticsOnPremiseVantagePointResourceCrud) Delete() error {
	request := oci_apm_synthetics.DeleteOnPremiseVantagePointRequest{}

	if apmDomainId, ok := s.D.GetOkExists("apm_domain_id"); ok {
		tmp := apmDomainId.(string)
		request.ApmDomainId = &tmp
	}

	onPremiseVantagePointId, apmDomainId, err1 := parseOnPremiseVantagePointCompositeId(s.D.Id())
	if err1 == nil {
		request.OnPremiseVantagePointId = &onPremiseVantagePointId
		request.ApmDomainId = &apmDomainId
	} else {
		log.Printf("[WARN] Delete() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apm_synthetics")

	_, err := s.Client.DeleteOnPremiseVantagePoint(context.Background(), request)
	return err
}

func (s *ApmSyntheticsOnPremiseVantagePointResourceCrud) SetData() error {

	_, apmDomainId, err := parseOnPremiseVantagePointCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("apm_domain_id", apmDomainId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
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

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	s.D.Set("type", s.Res.Type)

	if s.Res.WorkersSummary != nil {
		s.D.Set("workers_summary", []interface{}{WorkersSummaryToMap(s.Res.WorkersSummary)})
	} else {
		s.D.Set("workers_summary", nil)
	}

	return nil
}

func GetOnPremiseVantagePointCompositeId(onPremiseVantagePointId string, apmDomainId string) string {
	apmDomainId = url.PathEscape(apmDomainId)
	onPremiseVantagePointId = url.PathEscape(onPremiseVantagePointId)
	compositeId := "onPremiseVantagePoints/" + onPremiseVantagePointId + "/apmDomainId/" + apmDomainId
	return compositeId
}

func parseOnPremiseVantagePointCompositeId(compositeId string) (onPremiseVantagePointId string, apmDomainId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("onPremiseVantagePoints/.*/apmDomainId/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	onPremiseVantagePointId, _ = url.PathUnescape(parts[1])
	apmDomainId, _ = url.PathUnescape(parts[3])

	return
}

func AvailableCapabilityToMap(obj oci_apm_synthetics.AvailableCapability) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Capability != nil {
		result["capability"] = string(*obj.Capability)
	}

	if obj.Count != nil {
		result["on_premise_vantage_point_count"] = int(*obj.Count)
	}

	return result
}

func OnPremiseVantagePointSummaryToMap(obj oci_apm_synthetics.OnPremiseVantagePointSummary) map[string]interface{} {
	result := map[string]interface{}{}

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

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	result["type"] = string(obj.Type)

	if obj.WorkersSummary != nil {
		result["workers_summary"] = []interface{}{WorkersSummaryToMap(obj.WorkersSummary)}
	}

	return result
}

func WorkersSummaryToMap(obj *oci_apm_synthetics.WorkersSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Available != nil {
		result["available"] = int(*obj.Available)
	}

	availableCapabilities := []interface{}{}
	for _, item := range obj.AvailableCapabilities {
		availableCapabilities = append(availableCapabilities, AvailableCapabilityToMap(item))
	}
	result["available_capabilities"] = availableCapabilities

	if obj.Disabled != nil {
		result["disabled"] = int(*obj.Disabled)
	}

	if obj.MinVersion != nil {
		result["min_version"] = string(*obj.MinVersion)
	}

	if obj.Total != nil {
		result["total"] = int(*obj.Total)
	}

	if obj.Used != nil {
		result["used"] = int(*obj.Used)
	}

	return result
}
