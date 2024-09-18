// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package capacity_management

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_capacity_management "github.com/oracle/oci-go-sdk/v65/capacitymanagement"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CapacityManagementOccCapacityRequestResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCapacityManagementOccCapacityRequest,
		Read:     readCapacityManagementOccCapacityRequest,
		Update:   updateCapacityManagementOccCapacityRequest,
		Delete:   deleteCapacityManagementOccCapacityRequest,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"date_expected_capacity_handover": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
			},
			"details": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"demand_quantity": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							ValidateFunc:     tfresource.ValidateInt64TypeString,
							DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
						},
						"resource_name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"resource_type": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"workload_type": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"actual_handover_quantity": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ForceNew:         true,
							ValidateFunc:     tfresource.ValidateInt64TypeString,
							DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
						},
						"associated_occ_handover_resource_block_list": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"handover_quantity": {
										Type:             schema.TypeString,
										Optional:         true,
										Computed:         true,
										ForceNew:         true,
										ValidateFunc:     tfresource.ValidateInt64TypeString,
										DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
									},
									"occ_handover_resource_block_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},

									// Computed
								},
							},
						},
						"availability_domain": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
						},
						"date_actual_handover": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
						},
						"date_expected_handover": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
						},
						"expected_handover_quantity": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ForceNew:         true,
							ValidateFunc:     tfresource.ValidateInt64TypeString,
							DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
						},
						"source_workload_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"occ_availability_catalog_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			//"occ_capacity_request_id": {
			//	Type:     schema.TypeString,
			//	Required: true,
			//	ForceNew: true,
			//},
			"region": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"availability_domain": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
			},
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
				ForceNew: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
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
						//	Elem:     schema.TypeString,
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
			"request_state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"request_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"occ_customer_group_id": {
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

func createCapacityManagementOccCapacityRequest(d *schema.ResourceData, m interface{}) error {
	sync := &CapacityManagementOccCapacityRequestResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CapacityManagementClient()

	return tfresource.CreateResource(d, sync)
}

func readCapacityManagementOccCapacityRequest(d *schema.ResourceData, m interface{}) error {
	sync := &CapacityManagementOccCapacityRequestResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CapacityManagementClient()

	return tfresource.ReadResource(sync)
}

func updateCapacityManagementOccCapacityRequest(d *schema.ResourceData, m interface{}) error {
	sync := &CapacityManagementOccCapacityRequestResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CapacityManagementClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteCapacityManagementOccCapacityRequest(d *schema.ResourceData, m interface{}) error {
	sync := &CapacityManagementOccCapacityRequestResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CapacityManagementClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type CapacityManagementOccCapacityRequestResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_capacity_management.CapacityManagementClient
	Res                    *oci_capacity_management.OccCapacityRequest
	PatchResponse          *oci_capacity_management.OccCapacityRequest
	DisableNotFoundRetries bool
}

func (s *CapacityManagementOccCapacityRequestResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CapacityManagementOccCapacityRequestResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_capacity_management.OccCapacityRequestLifecycleStateCreating),
	}
}

func (s *CapacityManagementOccCapacityRequestResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_capacity_management.OccCapacityRequestLifecycleStateActive),
	}
}

func (s *CapacityManagementOccCapacityRequestResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_capacity_management.OccCapacityRequestLifecycleStateDeleting),
	}
}

func (s *CapacityManagementOccCapacityRequestResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_capacity_management.OccCapacityRequestLifecycleStateDeleted),
	}
}

func (s *CapacityManagementOccCapacityRequestResourceCrud) Create() error {
	request := oci_capacity_management.CreateOccCapacityRequestRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if dateExpectedCapacityHandover, ok := s.D.GetOkExists("date_expected_capacity_handover"); ok {
		tmp, err := time.Parse(time.RFC3339, dateExpectedCapacityHandover.(string))
		if err != nil {
			return err
		}
		request.DateExpectedCapacityHandover = &oci_common.SDKTime{Time: tmp}
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

	if details, ok := s.D.GetOkExists("details"); ok {
		interfaces := details.([]interface{})
		tmp := make([]oci_capacity_management.OccCapacityRequestBaseDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "details", stateDataIndex)
			converted, err := s.mapToOccCapacityRequestBaseDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("details") {
			request.Details = tmp
		}
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if lifecycleDetails, ok := s.D.GetOkExists("lifecycle_details"); ok {
		tmp := lifecycleDetails.(string)
		request.LifecycleDetails = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		request.Namespace = oci_capacity_management.NamespaceEnum(namespace.(string))
	}

	if occAvailabilityCatalogId, ok := s.D.GetOkExists("occ_availability_catalog_id"); ok {
		tmp := occAvailabilityCatalogId.(string)
		request.OccAvailabilityCatalogId = &tmp
	}

	if region, ok := s.D.GetOkExists("region"); ok {
		tmp := region.(string)
		request.Region = &tmp
	}

	if requestState, ok := s.D.GetOkExists("request_state"); ok {
		request.RequestState = oci_capacity_management.CreateOccCapacityRequestDetailsRequestStateEnum(requestState.(string))
	}

	if requestType, ok := s.D.GetOkExists("request_type"); ok {
		request.RequestType = oci_capacity_management.OccCapacityRequestRequestTypeEnum(requestType.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "capacity_management")

	response, err := s.Client.CreateOccCapacityRequest(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.OccCapacityRequest
	//err = s.Patch()
	//if err != nil {
	//	log.Printf("[ERROR] Failed to execute Patch operation: %v", err)
	//	return err
	//}
	return nil
}
func (s *CapacityManagementOccCapacityRequestResourceCrud) Patch() error {
	request := oci_capacity_management.PatchOccCapacityRequestRequest{}

	if occCapacityRequestId, ok := s.D.GetOkExists("id"); ok {
		tmp := occCapacityRequestId.(string)
		request.OccCapacityRequestId = &tmp
	}

	if patchOperations, ok := s.D.GetOkExists("patch_operations"); ok {
		interfaces := patchOperations.([]interface{})
		tmp := make([]oci_capacity_management.PatchInstruction, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "patch_operations", stateDataIndex)
			converted, err := s.mapToPatchInstruction(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("patch_operations") {
			request.Items = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "capacity_management")
	response, err := s.Client.PatchOccCapacityRequest(context.Background(), request)
	if err != nil {
		return err
	}

	s.PatchResponse = &response.OccCapacityRequest
	return nil
}

func (s *CapacityManagementOccCapacityRequestResourceCrud) Get() error {
	request := oci_capacity_management.GetOccCapacityRequestRequest{}

	tmp := s.D.Id()
	request.OccCapacityRequestId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "capacity_management")

	response, err := s.Client.GetOccCapacityRequest(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.OccCapacityRequest
	return nil
}

func (s *CapacityManagementOccCapacityRequestResourceCrud) Update() error {
	request := oci_capacity_management.UpdateOccCapacityRequestRequest{}

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

	tmp := s.D.Id()
	request.OccCapacityRequestId = &tmp

	if requestState, ok := s.D.GetOkExists("request_state"); ok {
		request.RequestState = oci_capacity_management.UpdateOccCapacityRequestDetailsRequestStateEnum(requestState.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "capacity_management")

	response, err := s.Client.UpdateOccCapacityRequest(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.OccCapacityRequest
	return nil
}

func (s *CapacityManagementOccCapacityRequestResourceCrud) Delete() error {
	request := oci_capacity_management.DeleteOccCapacityRequestRequest{}

	tmp := s.D.Id()
	request.OccCapacityRequestId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "capacity_management")

	_, err := s.Client.DeleteOccCapacityRequest(context.Background(), request)
	return err
}

func (s *CapacityManagementOccCapacityRequestResourceCrud) SetData() error {
	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DateExpectedCapacityHandover != nil {
		s.D.Set("date_expected_capacity_handover", s.Res.DateExpectedCapacityHandover.Format(time.RFC3339Nano))
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	details := []interface{}{}
	for _, item := range s.Res.Details {
		details = append(details, OccCapacityRequestBaseDetailsToMap(item))
	}
	s.D.Set("details", details)

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("namespace", s.Res.Namespace)

	if s.Res.OccAvailabilityCatalogId != nil {
		s.D.Set("occ_availability_catalog_id", *s.Res.OccAvailabilityCatalogId)
	}

	if s.Res.OccCustomerGroupId != nil {
		s.D.Set("occ_customer_group_id", *s.Res.OccCustomerGroupId)
	}

	if s.Res.Region != nil {
		s.D.Set("region", *s.Res.Region)
	}

	s.D.Set("request_state", s.Res.RequestState)

	s.D.Set("request_type", s.Res.RequestType)

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

func (s *CapacityManagementOccCapacityRequestResourceCrud) mapToAssociatedOccHandoverResourceBlock(fieldKeyFormat string) (oci_capacity_management.AssociatedOccHandoverResourceBlock, error) {
	result := oci_capacity_management.AssociatedOccHandoverResourceBlock{}

	if handoverQuantity, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "handover_quantity")); ok {
		tmp := handoverQuantity.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return result, fmt.Errorf("unable to convert handoverQuantity string: %s to an int64 and encountered error: %v", tmp, err)
		}
		result.HandoverQuantity = &tmpInt64
	}

	if occHandoverResourceBlockId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "occ_handover_resource_block_id")); ok {
		tmp := occHandoverResourceBlockId.(string)
		result.OccHandoverResourceBlockId = &tmp
	}

	return result, nil
}

func AssociatedOccHandoverResourceBlockToMap(obj oci_capacity_management.AssociatedOccHandoverResourceBlock) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.HandoverQuantity != nil {
		result["handover_quantity"] = strconv.FormatInt(*obj.HandoverQuantity, 10)
	}

	if obj.OccHandoverResourceBlockId != nil {
		result["occ_handover_resource_block_id"] = string(*obj.OccHandoverResourceBlockId)
	}

	return result
}

func (s *CapacityManagementOccCapacityRequestResourceCrud) mapToOccCapacityRequestBaseDetails(fieldKeyFormat string) (oci_capacity_management.OccCapacityRequestBaseDetails, error) {
	result := oci_capacity_management.OccCapacityRequestBaseDetails{}

	if actualHandoverQuantity, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "actual_handover_quantity")); ok {
		tmp := actualHandoverQuantity.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return result, fmt.Errorf("unable to convert actualHandoverQuantity string: %s to an int64 and encountered error: %v", tmp, err)
		}
		result.ActualHandoverQuantity = &tmpInt64
	}

	if associatedOccHandoverResourceBlockList, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "associated_occ_handover_resource_block_list")); ok {
		interfaces := associatedOccHandoverResourceBlockList.([]interface{})
		tmp := make([]oci_capacity_management.AssociatedOccHandoverResourceBlock, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "associated_occ_handover_resource_block_list"), stateDataIndex)
			converted, err := s.mapToAssociatedOccHandoverResourceBlock(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "associated_occ_handover_resource_block_list")) {
			result.AssociatedOccHandoverResourceBlockList = tmp
		}
	}

	if availabilityDomain, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "availability_domain")); ok {
		tmp := availabilityDomain.(string)
		result.AvailabilityDomain = &tmp
	}

	if dateActualHandover, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "date_actual_handover")); ok {
		tmp, err := time.Parse(time.RFC3339, dateActualHandover.(string))
		if err != nil {
			return result, err
		}
		result.DateActualHandover = &oci_common.SDKTime{Time: tmp}
	}

	if dateExpectedHandover, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "date_expected_handover")); ok {
		tmp, err := time.Parse(time.RFC3339, dateExpectedHandover.(string))
		if err != nil {
			return result, err
		}
		result.DateExpectedHandover = &oci_common.SDKTime{Time: tmp}
	}

	if demandQuantity, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "demand_quantity")); ok {
		tmp := demandQuantity.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return result, fmt.Errorf("unable to convert demandQuantity string: %s to an int64 and encountered error: %v", tmp, err)
		}
		result.DemandQuantity = &tmpInt64
	}

	if expectedHandoverQuantity, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "expected_handover_quantity")); ok {
		tmp := expectedHandoverQuantity.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return result, fmt.Errorf("unable to convert expectedHandoverQuantity string: %s to an int64 and encountered error: %v", tmp, err)
		}
		result.ExpectedHandoverQuantity = &tmpInt64
	}

	if resourceName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "resource_name")); ok {
		tmp := resourceName.(string)
		result.ResourceName = &tmp
	}

	if resourceType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "resource_type")); ok {
		tmp := resourceType.(string)
		result.ResourceType = &tmp
	}

	if sourceWorkloadType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_workload_type")); ok {
		tmp := sourceWorkloadType.(string)
		result.SourceWorkloadType = &tmp
	}

	if workloadType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "workload_type")); ok {
		tmp := workloadType.(string)
		result.WorkloadType = &tmp
	}

	return result, nil
}

func OccCapacityRequestBaseDetailsToMap(obj oci_capacity_management.OccCapacityRequestBaseDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ActualHandoverQuantity != nil {
		result["actual_handover_quantity"] = strconv.FormatInt(*obj.ActualHandoverQuantity, 10)
	}

	associatedOccHandoverResourceBlockList := []interface{}{}
	for _, item := range obj.AssociatedOccHandoverResourceBlockList {
		associatedOccHandoverResourceBlockList = append(associatedOccHandoverResourceBlockList, AssociatedOccHandoverResourceBlockToMap(item))
	}
	result["associated_occ_handover_resource_block_list"] = associatedOccHandoverResourceBlockList

	if obj.AvailabilityDomain != nil {
		result["availability_domain"] = string(*obj.AvailabilityDomain)
	}

	if obj.DateActualHandover != nil {
		result["date_actual_handover"] = obj.DateActualHandover.Format(time.RFC3339Nano)
	}

	if obj.DateExpectedHandover != nil {
		result["date_expected_handover"] = obj.DateExpectedHandover.Format(time.RFC3339Nano)
	}

	if obj.DemandQuantity != nil {
		result["demand_quantity"] = strconv.FormatInt(*obj.DemandQuantity, 10)
	}

	if obj.ExpectedHandoverQuantity != nil {
		result["expected_handover_quantity"] = strconv.FormatInt(*obj.ExpectedHandoverQuantity, 10)
	}

	if obj.ResourceName != nil {
		result["resource_name"] = string(*obj.ResourceName)
	}

	if obj.ResourceType != nil {
		result["resource_type"] = string(*obj.ResourceType)
	}

	if obj.SourceWorkloadType != nil {
		result["source_workload_type"] = string(*obj.SourceWorkloadType)
	}

	if obj.WorkloadType != nil {
		result["workload_type"] = string(*obj.WorkloadType)
	}

	return result
}

func OccCapacityRequestSummaryToMap(obj oci_capacity_management.OccCapacityRequestSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AvailabilityDomain != nil {
		result["availability_domain"] = string(*obj.AvailabilityDomain)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DateExpectedCapacityHandover != nil {
		result["date_expected_capacity_handover"] = obj.DateExpectedCapacityHandover.String()
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

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	result["namespace"] = string(obj.Namespace)

	if obj.OccAvailabilityCatalogId != nil {
		result["occ_availability_catalog_id"] = string(*obj.OccAvailabilityCatalogId)
	}

	if obj.OccCustomerGroupId != nil {
		result["occ_customer_group_id"] = string(*obj.OccCustomerGroupId)
	}

	if obj.Region != nil {
		result["region"] = string(*obj.Region)
	}

	result["request_state"] = string(obj.RequestState)

	result["request_type"] = string(obj.RequestType)

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

func (s *CapacityManagementOccCapacityRequestResourceCrud) mapToPatchInstruction(fieldKeyFormat string) (oci_capacity_management.PatchInstruction, error) {
	var baseObject oci_capacity_management.PatchInstruction
	//discriminator
	operationRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "operation"))
	var operation string
	if ok {
		operation = operationRaw.(string)
	} else {
		operation = "" // default value
	}
	switch strings.ToLower(operation) {
	case strings.ToLower("INSERT"):
		details := oci_capacity_management.PatchInsertInstruction{}
		if position, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "position")); ok {
			details.Position = oci_capacity_management.PatchInsertInstructionPositionEnum(position.(string))
		}
		if selectedItem, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "selected_item")); ok {
			tmp := selectedItem.(string)
			details.SelectedItem = &tmp
		}
		if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
			details.Value = &value
		}
		if selection, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "selection")); ok {
			tmp := selection.(string)
			details.Selection = &tmp
		}
		baseObject = details
	case strings.ToLower("INSERT_MULTIPLE"):
		details := oci_capacity_management.PatchInsertMultipleInstruction{}
		if position, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "position")); ok {
			details.Position = oci_capacity_management.PatchInsertMultipleInstructionPositionEnum(position.(string))
		}
		if selectedItem, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "selected_item")); ok {
			tmp := selectedItem.(string)
			details.SelectedItem = &tmp
		}
		if values, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "values")); ok {
			interfaces := values.([]interface{})
			if s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "values")) {
				details.Values = interfaces
			}
		}
		if selection, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "selection")); ok {
			tmp := selection.(string)
			details.Selection = &tmp
		}
		baseObject = details
	case strings.ToLower("MERGE"):
		details := oci_capacity_management.PatchMergeInstruction{}
		if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
			details.Value = &value
		}
		if selection, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "selection")); ok {
			tmp := selection.(string)
			details.Selection = &tmp
		}
		baseObject = details
	case strings.ToLower("MOVE"):
		details := oci_capacity_management.PatchMoveInstruction{}
		if from, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "from")); ok {
			tmp := from.(string)
			details.From = &tmp
		}
		if position, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "position")); ok {
			details.Position = oci_capacity_management.PatchMoveInstructionPositionEnum(position.(string))
		}
		if selection, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "selection")); ok {
			tmp := selection.(string)
			details.Selection = &tmp
		}
		baseObject = details
	case strings.ToLower("PROHIBIT"):
		details := oci_capacity_management.PatchProhibitInstruction{}
		if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
			details.Value = &value
		}
		if selection, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "selection")); ok {
			tmp := selection.(string)
			details.Selection = &tmp
		}
		baseObject = details
	case strings.ToLower("REMOVE"):
		details := oci_capacity_management.PatchRemoveInstruction{}
		if selection, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "selection")); ok {
			tmp := selection.(string)
			details.Selection = &tmp
		}
		baseObject = details
	case strings.ToLower("REPLACE"):
		details := oci_capacity_management.PatchReplaceInstruction{}
		if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
			details.Value = &value
		}
		if selection, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "selection")); ok {
			tmp := selection.(string)
			details.Selection = &tmp
		}
		baseObject = details
	case strings.ToLower("REQUIRE"):
		details := oci_capacity_management.PatchRequireInstruction{}
		if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
			details.Value = &value
		}
		if selection, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "selection")); ok {
			tmp := selection.(string)
			details.Selection = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown operation '%v' was specified", operation)
	}
	return baseObject, nil
}

func (s *CapacityManagementOccCapacityRequestResourceCrud) mapToobject(fieldKeyFormat string) (oci_capacity_management.OccCapacityRequest, error) {
	result := oci_capacity_management.OccCapacityRequest{}

	return result, nil
}
