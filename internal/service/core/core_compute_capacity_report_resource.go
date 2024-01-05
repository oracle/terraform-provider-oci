// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_core "github.com/oracle/oci-go-sdk/v65/core"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CoreComputeCapacityReportResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCoreComputeCapacityReport,
		Read:     readCoreComputeCapacityReport,
		Delete:   deleteCoreComputeCapacityReport,
		Schema: map[string]*schema.Schema{
			// Required
			"availability_domain": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
			},
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
						"instance_shape": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"fault_domain": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"instance_shape_config": {
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
										Type:     schema.TypeFloat,
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
										Type:     schema.TypeFloat,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},

									// Computed
								},
							},
						},

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

			// Optional

			// Computed
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createCoreComputeCapacityReport(d *schema.ResourceData, m interface{}) error {
	sync := &CoreComputeCapacityReportResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.CreateResource(d, sync)
}

func readCoreComputeCapacityReport(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteCoreComputeCapacityReport(d *schema.ResourceData, m interface{}) error {
	return nil
}

type CoreComputeCapacityReportResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_core.ComputeClient
	Res                    *oci_core.ComputeCapacityReport
	DisableNotFoundRetries bool
}

func (s *CoreComputeCapacityReportResourceCrud) ID() string {
	return *s.Res.CompartmentId
}

func (s *CoreComputeCapacityReportResourceCrud) Create() error {
	request := oci_core.CreateComputeCapacityReportRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if shapeAvailabilities, ok := s.D.GetOkExists("shape_availabilities"); ok {
		interfaces := shapeAvailabilities.([]interface{})
		tmp := make([]oci_core.CreateCapacityReportShapeAvailabilityDetails, len(interfaces))
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateComputeCapacityReport(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ComputeCapacityReport
	return nil
}

func (s *CoreComputeCapacityReportResourceCrud) SetData() error {
	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

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

func (s *CoreComputeCapacityReportResourceCrud) mapToCapacityReportInstanceShapeConfig(fieldKeyFormat string) (oci_core.CapacityReportInstanceShapeConfig, error) {
	result := oci_core.CapacityReportInstanceShapeConfig{}

	if memoryInGBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "memory_in_gbs")); ok {
		tmp := float32(memoryInGBs.(float64))
		result.MemoryInGBs = &tmp
	}

	if nvmes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "nvmes")); ok {
		tmp := nvmes.(int)
		result.Nvmes = &tmp
	}

	if ocpus, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ocpus")); ok {
		tmp := float32(ocpus.(float64))
		result.Ocpus = &tmp
	}

	return result, nil
}

func CapacityReportInstanceShapeConfigToMap(obj *oci_core.CapacityReportInstanceShapeConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.MemoryInGBs != nil {
		result["memory_in_gbs"] = float32(*obj.MemoryInGBs)
	}

	if obj.Nvmes != nil {
		result["nvmes"] = int(*obj.Nvmes)
	}

	if obj.Ocpus != nil {
		result["ocpus"] = float32(*obj.Ocpus)
	}

	return result
}

func (s *CoreComputeCapacityReportResourceCrud) mapToCreateCapacityReportShapeAvailabilityDetails(fieldKeyFormat string) (oci_core.CreateCapacityReportShapeAvailabilityDetails, error) {
	result := oci_core.CreateCapacityReportShapeAvailabilityDetails{}

	if faultDomain, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "fault_domain")); ok {
		tmp := faultDomain.(string)
		result.FaultDomain = &tmp
	}

	if instanceShape, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "instance_shape")); ok {
		tmp := instanceShape.(string)
		result.InstanceShape = &tmp
	}

	if instanceShapeConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "instance_shape_config")); ok {
		if tmpList := instanceShapeConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "instance_shape_config"), 0)
			tmp, err := s.mapToCapacityReportInstanceShapeConfig(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert instance_shape_config, encountered error: %v", err)
			}
			result.InstanceShapeConfig = &tmp
		}
	}

	return result, nil
}

func CapacityReportShapeAvailabilityToMap(obj oci_core.CapacityReportShapeAvailability) map[string]interface{} {
	result := map[string]interface{}{}

	result["availability_status"] = string(obj.AvailabilityStatus)

	if obj.AvailableCount != nil {
		result["available_count"] = strconv.FormatInt(*obj.AvailableCount, 10)
	}

	if obj.FaultDomain != nil {
		result["fault_domain"] = string(*obj.FaultDomain)
	}

	if obj.InstanceShape != nil {
		result["instance_shape"] = string(*obj.InstanceShape)
	}

	if obj.InstanceShapeConfig != nil {
		result["instance_shape_config"] = []interface{}{CapacityReportInstanceShapeConfigToMap(obj.InstanceShapeConfig)}
	}

	return result
}
