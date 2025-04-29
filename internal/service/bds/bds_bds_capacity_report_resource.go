// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package bds

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_bds "github.com/oracle/oci-go-sdk/v65/bds"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func BdsBdsCapacityReportResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createBdsBdsCapacityReport,
		Read:     readBdsBdsCapacityReport,
		Delete:   deleteBdsBdsCapacityReport,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"shape_availabilities": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"shape": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"shape_config": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"memory_in_gbs": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"nvmes": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"ocpus": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},

									// Computed
								},
							},
						},

						// Computed
						"domain_level_capacity_reports": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"availability_domain": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"capacity_availability": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"availability_status": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"available_count": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"domain_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"fault_domain": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},

			// Optional

			// Computed
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createBdsBdsCapacityReport(d *schema.ResourceData, m interface{}) error {
	sync := &BdsBdsCapacityReportResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()

	return tfresource.CreateResource(d, sync)
}

func readBdsBdsCapacityReport(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteBdsBdsCapacityReport(d *schema.ResourceData, m interface{}) error {
	return nil
}

type BdsBdsCapacityReportResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_bds.BdsClient
	Res                    *oci_bds.BdsCapacityReport
	DisableNotFoundRetries bool
}

func (s *BdsBdsCapacityReportResourceCrud) ID() string {
	return tfresource.GenerateDataSourceHashID("BdsBdsCapacityReportResource-", BdsBdsCapacityReportResource(), s.D)
}

func (s *BdsBdsCapacityReportResourceCrud) Create() error {
	request := oci_bds.CreateBdsCapacityReportRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if shapeAvailabilities, ok := s.D.GetOkExists("shape_availabilities"); ok {
		interfaces := shapeAvailabilities.([]interface{})
		tmp := make([]oci_bds.CreateCapacityReportShapeAvailabilityDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "shape_availabilities", stateDataIndex)
			converted, err := s.mapToCreateCapacityReportShapeAvailabilityDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("shape_availabilities") {
			request.ShapeAvailabilities = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.CreateBdsCapacityReport(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.BdsCapacityReport
	return nil
}

func (s *BdsBdsCapacityReportResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	shapeAvailabilities := []interface{}{}
	for _, item := range s.Res.ShapeAvailabilities {
		shapeAvailabilities = append(shapeAvailabilities, CapacityReportShapeAvailabilityToMap(item))
	}
	s.D.Set("shape_availabilities", shapeAvailabilities)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func (s *BdsBdsCapacityReportResourceCrud) mapToCapacityAvailability(fieldKeyFormat string) (oci_bds.CapacityAvailability, error) {
	result := oci_bds.CapacityAvailability{}

	if availabilityStatus, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "availability_status")); ok {
		result.AvailabilityStatus = oci_bds.AvailabilityStatusEnum(availabilityStatus.(string))
	}

	if availableCount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "available_count")); ok {
		tmp := availableCount.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return result, fmt.Errorf("unable to convert availableCount string: %s to an int64 and encountered error: %v", tmp, err)
		}
		result.AvailableCount = &tmpInt64
	}

	return result, nil
}

func CapacityAvailabilityToMap(obj *oci_bds.CapacityAvailability) map[string]interface{} {
	result := map[string]interface{}{}

	result["availability_status"] = string(obj.AvailabilityStatus)

	if obj.AvailableCount != nil {
		result["available_count"] = strconv.FormatInt(*obj.AvailableCount, 10)
	}

	return result
}

func (s *BdsBdsCapacityReportResourceCrud) mapToCreateCapacityReportShapeAvailabilityDetails(fieldKeyFormat string) (oci_bds.CreateCapacityReportShapeAvailabilityDetails, error) {
	result := oci_bds.CreateCapacityReportShapeAvailabilityDetails{}

	if shape, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "shape")); ok {
		tmp := shape.(string)
		result.Shape = &tmp
	}

	if shapeConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "shape_config")); ok {
		if tmpList := shapeConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "shape_config"), 0)
			tmp, err := s.mapToShapeConfigDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert shape_config, encountered error: %v", err)
			}
			result.ShapeConfig = &tmp
		}
	}

	return result, nil
}

func CapacityReportShapeAvailabilityToMap(obj oci_bds.CapacityReportShapeAvailability) map[string]interface{} {
	result := map[string]interface{}{}

	domainLevelCapacityReports := []interface{}{}
	for _, item := range obj.DomainLevelCapacityReports {
		domainLevelCapacityReports = append(domainLevelCapacityReports, DomainTypeCapacityReportToMap(item))
	}
	result["domain_level_capacity_reports"] = domainLevelCapacityReports

	if obj.Shape != nil {
		result["shape"] = string(*obj.Shape)
	}

	if obj.ShapeConfig != nil {
		result["shape_config"] = []interface{}{ShapeConfigDetailsToMap(obj.ShapeConfig)}
	}

	return result
}

func DomainTypeCapacityReportToMap(obj oci_bds.DomainTypeCapacityReport) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_bds.MultiAdCapacityReport:
		result["domain_type"] = "MULTI_AD"

		if v.AvailabilityDomain != nil {
			result["availability_domain"] = string(*v.AvailabilityDomain)
		}

		if v.CapacityAvailability != nil {
			result["capacity_availability"] = []interface{}{CapacityAvailabilityToMap(v.CapacityAvailability)}
		}
	case oci_bds.SingleAdCapacityReport:
		result["domain_type"] = "SINGLE_AD"

		if v.CapacityAvailability != nil {
			result["capacity_availability"] = []interface{}{CapacityAvailabilityToMap(v.CapacityAvailability)}
		}

		if v.FaultDomain != nil {
			result["fault_domain"] = string(*v.FaultDomain)
		}
	default:
		log.Printf("[WARN] Received 'domain_type' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *BdsBdsCapacityReportResourceCrud) mapToShapeConfigDetails(fieldKeyFormat string) (oci_bds.ShapeConfigDetails, error) {
	result := oci_bds.ShapeConfigDetails{}

	if memoryInGBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "memory_in_gbs")); ok {
		tmp := memoryInGBs.(int)
		result.MemoryInGBs = &tmp
	}

	if nvmes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "nvmes")); ok {
		tmp := nvmes.(int)
		result.Nvmes = &tmp
	}

	if ocpus, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ocpus")); ok {
		tmp := ocpus.(int)
		result.Ocpus = &tmp
	}

	return result, nil
}

func ShapeConfigDetailsToMap(obj *oci_bds.ShapeConfigDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.MemoryInGBs != nil {
		result["memory_in_gbs"] = int(*obj.MemoryInGBs)
	}

	if obj.Nvmes != nil {
		result["nvmes"] = int(*obj.Nvmes)
	}

	if obj.Ocpus != nil {
		result["ocpus"] = int(*obj.Ocpus)
	}

	return result
}
