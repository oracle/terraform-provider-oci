// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package capacity_management

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_capacity_management "github.com/oracle/oci-go-sdk/v65/capacitymanagement"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CapacityManagementInternalOccmDemandSignalDeliveryResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCapacityManagementInternalOccmDemandSignalDelivery,
		Read:     readCapacityManagementInternalOccmDemandSignalDelivery,
		Update:   updateCapacityManagementInternalOccmDemandSignalDelivery,
		Delete:   deleteCapacityManagementInternalOccmDemandSignalDelivery,
		Schema: map[string]*schema.Schema{
			// Required
			"accepted_quantity": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateFunc:     tfresource.ValidateInt64TypeString,
				DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"demand_signal_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"demand_signal_item_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"occ_customer_group_id": {
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
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"justification": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"notes": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"lifecycle_details": {
				Type:     schema.TypeString,
				Optional: true,
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
			"time_delivered": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createCapacityManagementInternalOccmDemandSignalDelivery(d *schema.ResourceData, m interface{}) error {
	sync := &CapacityManagementInternalOccmDemandSignalDeliveryResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).InternalDemandSignalClient()

	return tfresource.CreateResource(d, sync)
}

func readCapacityManagementInternalOccmDemandSignalDelivery(d *schema.ResourceData, m interface{}) error {
	sync := &CapacityManagementInternalOccmDemandSignalDeliveryResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).InternalDemandSignalClient()

	return tfresource.ReadResource(sync)
}

func updateCapacityManagementInternalOccmDemandSignalDelivery(d *schema.ResourceData, m interface{}) error {
	sync := &CapacityManagementInternalOccmDemandSignalDeliveryResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).InternalDemandSignalClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteCapacityManagementInternalOccmDemandSignalDelivery(d *schema.ResourceData, m interface{}) error {
	sync := &CapacityManagementInternalOccmDemandSignalDeliveryResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).InternalDemandSignalClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type CapacityManagementInternalOccmDemandSignalDeliveryResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_capacity_management.InternalDemandSignalClient
	Res                    *oci_capacity_management.InternalOccmDemandSignalDelivery
	DisableNotFoundRetries bool
}

func (s *CapacityManagementInternalOccmDemandSignalDeliveryResourceCrud) ID() string {
	return GetInternalOccmDemandSignalDeliveryCompositeId(*s.Res.Id)
}

func (s *CapacityManagementInternalOccmDemandSignalDeliveryResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_capacity_management.InternalOccmDemandSignalDeliveryLifecycleStateCreating),
	}
}

func (s *CapacityManagementInternalOccmDemandSignalDeliveryResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_capacity_management.InternalOccmDemandSignalDeliveryLifecycleStateActive),
	}
}

func (s *CapacityManagementInternalOccmDemandSignalDeliveryResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_capacity_management.InternalOccmDemandSignalDeliveryLifecycleStateDeleting),
	}
}

func (s *CapacityManagementInternalOccmDemandSignalDeliveryResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_capacity_management.InternalOccmDemandSignalDeliveryLifecycleStateDeleted),
	}
}

func (s *CapacityManagementInternalOccmDemandSignalDeliveryResourceCrud) Create() error {
	request := oci_capacity_management.CreateInternalOccmDemandSignalDeliveryRequest{}

	if acceptedQuantity, ok := s.D.GetOkExists("accepted_quantity"); ok {
		tmp := acceptedQuantity.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert acceptedQuantity string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.AcceptedQuantity = &tmpInt64
	}

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

	if demandSignalId, ok := s.D.GetOkExists("demand_signal_id"); ok {
		tmp := demandSignalId.(string)
		request.DemandSignalId = &tmp
	}

	if demandSignalItemId, ok := s.D.GetOkExists("demand_signal_item_id"); ok {
		tmp := demandSignalItemId.(string)
		request.DemandSignalItemId = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if justification, ok := s.D.GetOkExists("justification"); ok {
		tmp := justification.(string)
		request.Justification = &tmp
	}

	if notes, ok := s.D.GetOkExists("notes"); ok {
		tmp := notes.(string)
		request.Notes = &tmp
	}

	if occCustomerGroupId, ok := s.D.GetOkExists("occ_customer_group_id"); ok {
		tmp := occCustomerGroupId.(string)
		request.OccCustomerGroupId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "capacity_management")

	response, err := s.Client.CreateInternalOccmDemandSignalDelivery(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.InternalOccmDemandSignalDelivery
	return nil
}

func (s *CapacityManagementInternalOccmDemandSignalDeliveryResourceCrud) Get() error {
	request := oci_capacity_management.GetInternalOccmDemandSignalDeliveryRequest{}

	occmDemandSignalDeliveryId, err := parseInternalOccmDemandSignalDeliveryCompositeId(s.D.Id())
	if err == nil {
		request.OccmDemandSignalDeliveryId = &occmDemandSignalDeliveryId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "capacity_management")

	response, err := s.Client.GetInternalOccmDemandSignalDelivery(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.InternalOccmDemandSignalDelivery
	return nil
}

func (s *CapacityManagementInternalOccmDemandSignalDeliveryResourceCrud) Update() error {
	request := oci_capacity_management.UpdateInternalOccmDemandSignalDeliveryRequest{}

	if acceptedQuantity, ok := s.D.GetOkExists("accepted_quantity"); ok {
		tmp := acceptedQuantity.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert acceptedQuantity string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.AcceptedQuantity = &tmpInt64
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if justification, ok := s.D.GetOkExists("justification"); ok {
		tmp := justification.(string)
		request.Justification = &tmp
	}

	if lifecycleDetails, ok := s.D.GetOkExists("lifecycle_details"); ok {
		if s.D.HasChange("lifecycle_details") {
			request.LifecycleDetails = oci_capacity_management.InternalOccmDemandSignalDeliveryLifecycleDetailsEnum(lifecycleDetails.(string))
		}
	}

	if notes, ok := s.D.GetOkExists("notes"); ok {
		tmp := notes.(string)
		request.Notes = &tmp
	}

	occmDemandSignalDeliveryId, err := parseInternalOccmDemandSignalDeliveryCompositeId(s.D.Id())
	if err != nil {
		return fmt.Errorf("error parsing resource ID for update: %w", err)
	}
	request.OccmDemandSignalDeliveryId = &occmDemandSignalDeliveryId

	if timeDelivered, ok := s.D.GetOkExists("time_delivered"); ok {
		tmp, err := time.Parse(time.RFC3339, timeDelivered.(string))
		if err != nil {
			return err
		}
		request.TimeDelivered = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "capacity_management")

	response, err := s.Client.UpdateInternalOccmDemandSignalDelivery(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.InternalOccmDemandSignalDelivery
	return nil
}

func (s *CapacityManagementInternalOccmDemandSignalDeliveryResourceCrud) Delete() error {
	request := oci_capacity_management.DeleteInternalOccmDemandSignalDeliveryRequest{}

	occmDemandSignalDeliveryId, err := parseInternalOccmDemandSignalDeliveryCompositeId(s.D.Id())
	if err == nil {
		request.OccmDemandSignalDeliveryId = &occmDemandSignalDeliveryId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "capacity_management")

	_, err = s.Client.DeleteInternalOccmDemandSignalDelivery(context.Background(), request)
	return err
}

func (s *CapacityManagementInternalOccmDemandSignalDeliveryResourceCrud) SetData() error {

	_, err := parseInternalOccmDemandSignalDeliveryCompositeId(s.D.Id())
	if err == nil {
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.AcceptedQuantity != nil {
		s.D.Set("accepted_quantity", strconv.FormatInt(*s.Res.AcceptedQuantity, 10))
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DemandSignalId != nil {
		s.D.Set("demand_signal_id", *s.Res.DemandSignalId)
	}

	if s.Res.DemandSignalItemId != nil {
		s.D.Set("demand_signal_item_id", *s.Res.DemandSignalItemId)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Justification != nil {
		s.D.Set("justification", *s.Res.Justification)
	}

	s.D.Set("lifecycle_details", s.Res.LifecycleDetails)

	if s.Res.Notes != nil {
		s.D.Set("notes", *s.Res.Notes)
	}

	if s.Res.OccCustomerGroupId != nil {
		s.D.Set("occ_customer_group_id", *s.Res.OccCustomerGroupId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeDelivered != nil {
		s.D.Set("time_delivered", s.Res.TimeDelivered.String())
	}

	return nil
}

func GetInternalOccmDemandSignalDeliveryCompositeId(occmDemandSignalDeliveryId string) string {
	occmDemandSignalDeliveryId = url.PathEscape(occmDemandSignalDeliveryId)
	compositeId := "internal/occmDemandSignalDeliveries/" + occmDemandSignalDeliveryId + ""
	return compositeId
}

func parseInternalOccmDemandSignalDeliveryCompositeId(compositeId string) (occmDemandSignalDeliveryId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("internal/occmDemandSignalDeliveries/.*", compositeId)
	if !match || len(parts) != 3 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	occmDemandSignalDeliveryId, _ = url.PathUnescape(parts[2])

	return
}

func InternalOccmDemandSignalDeliverySummaryToMap(obj oci_capacity_management.InternalOccmDemandSignalDeliverySummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AcceptedQuantity != nil {
		result["accepted_quantity"] = strconv.FormatInt(*obj.AcceptedQuantity, 10)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DemandSignalId != nil {
		result["demand_signal_id"] = string(*obj.DemandSignalId)
	}

	if obj.DemandSignalItemId != nil {
		result["demand_signal_item_id"] = string(*obj.DemandSignalItemId)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.Justification != nil {
		result["justification"] = string(*obj.Justification)
	}

	result["lifecycle_details"] = string(obj.LifecycleDetails)

	if obj.Notes != nil {
		result["notes"] = string(*obj.Notes)
	}

	if obj.OccCustomerGroupId != nil {
		result["occ_customer_group_id"] = string(*obj.OccCustomerGroupId)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeDelivered != nil {
		result["time_delivered"] = obj.TimeDelivered.String()
	}

	return result
}
