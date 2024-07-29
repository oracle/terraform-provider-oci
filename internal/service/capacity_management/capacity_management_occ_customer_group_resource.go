// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package capacity_management

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_capacity_management "github.com/oracle/oci-go-sdk/v65/capacitymanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CapacityManagementOccCustomerGroupResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCapacityManagementOccCustomerGroup,
		Read:     readCapacityManagementOccCustomerGroup,
		Update:   updateCapacityManagementOccCustomerGroup,
		Delete:   deleteCapacityManagementOccCustomerGroup,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"customers_list": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"display_name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"tenancy_id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"description": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"status": {
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
					},
				},
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
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
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

func createCapacityManagementOccCustomerGroup(d *schema.ResourceData, m interface{}) error {
	sync := &CapacityManagementOccCustomerGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CapacityManagementClient()

	return tfresource.CreateResource(d, sync)
}

func readCapacityManagementOccCustomerGroup(d *schema.ResourceData, m interface{}) error {
	sync := &CapacityManagementOccCustomerGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CapacityManagementClient()

	return tfresource.ReadResource(sync)
}

func updateCapacityManagementOccCustomerGroup(d *schema.ResourceData, m interface{}) error {
	sync := &CapacityManagementOccCustomerGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CapacityManagementClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteCapacityManagementOccCustomerGroup(d *schema.ResourceData, m interface{}) error {
	sync := &CapacityManagementOccCustomerGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CapacityManagementClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type CapacityManagementOccCustomerGroupResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_capacity_management.CapacityManagementClient
	Res                    *oci_capacity_management.OccCustomerGroup
	DisableNotFoundRetries bool
}

func (s *CapacityManagementOccCustomerGroupResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CapacityManagementOccCustomerGroupResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_capacity_management.OccCustomerGroupLifecycleStateCreating),
	}
}

func (s *CapacityManagementOccCustomerGroupResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_capacity_management.OccCustomerGroupLifecycleStateActive),
	}
}

func (s *CapacityManagementOccCustomerGroupResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_capacity_management.OccCustomerGroupLifecycleStateDeleting),
	}
}

func (s *CapacityManagementOccCustomerGroupResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_capacity_management.OccCustomerGroupLifecycleStateDeleted),
	}
}

func (s *CapacityManagementOccCustomerGroupResourceCrud) Create() error {
	request := oci_capacity_management.CreateOccCustomerGroupRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if customersList, ok := s.D.GetOkExists("customers_list"); ok {
		interfaces := customersList.([]interface{})
		tmp := make([]oci_capacity_management.CreateOccCustomerDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "customers_list", stateDataIndex)
			converted, err := s.mapToCreateOccCustomerDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("customers_list") {
			request.CustomersList = tmp
		}
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

	if status, ok := s.D.GetOkExists("status"); ok {
		request.Status = oci_capacity_management.CreateOccCustomerGroupDetailsStatusEnum(status.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "capacity_management")

	response, err := s.Client.CreateOccCustomerGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.OccCustomerGroup
	return nil
}

func (s *CapacityManagementOccCustomerGroupResourceCrud) Get() error {
	request := oci_capacity_management.GetOccCustomerGroupRequest{}

	tmp := s.D.Id()
	request.OccCustomerGroupId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "capacity_management")

	response, err := s.Client.GetOccCustomerGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.OccCustomerGroup
	return nil
}

func (s *CapacityManagementOccCustomerGroupResourceCrud) Update() error {
	request := oci_capacity_management.UpdateOccCustomerGroupRequest{}

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

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.OccCustomerGroupId = &tmp

	if status, ok := s.D.GetOkExists("status"); ok {
		request.Status = oci_capacity_management.UpdateOccCustomerGroupDetailsStatusEnum(status.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "capacity_management")

	response, err := s.Client.UpdateOccCustomerGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.OccCustomerGroup
	return nil
}

func (s *CapacityManagementOccCustomerGroupResourceCrud) Delete() error {
	request := oci_capacity_management.DeleteOccCustomerGroupRequest{}

	tmp := s.D.Id()
	request.OccCustomerGroupId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "capacity_management")

	_, err := s.Client.DeleteOccCustomerGroup(context.Background(), request)
	return err
}

func (s *CapacityManagementOccCustomerGroupResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	customersList := []interface{}{}
	for _, item := range s.Res.CustomersList {
		customersList = append(customersList, OccCustomerToMap(item))
	}
	s.D.Set("customers_list", customersList)

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

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("status", s.Res.Status)

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

func (s *CapacityManagementOccCustomerGroupResourceCrud) mapToCreateOccCustomerDetails(fieldKeyFormat string) (oci_capacity_management.CreateOccCustomerDetails, error) {
	result := oci_capacity_management.CreateOccCustomerDetails{}

	if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
		tmp := description.(string)
		result.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	if status, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "status")); ok {
		result.Status = oci_capacity_management.CreateOccCustomerDetailsStatusEnum(status.(string))
	}

	if tenancyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tenancy_id")); ok {
		tmp := tenancyId.(string)
		result.TenancyId = &tmp
	}

	return result, nil
}

func OccCustomerToMap(obj oci_capacity_management.OccCustomer) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.OccCustomerGroupId != nil {
		result["occ_customer_group_id"] = string(*obj.OccCustomerGroupId)
	}

	result["status"] = string(obj.Status)

	if obj.TenancyId != nil {
		result["tenancy_id"] = string(*obj.TenancyId)
	}

	return result
}

func OccCustomerGroupSummaryToMap(obj oci_capacity_management.OccCustomerGroupSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
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

	result["state"] = string(obj.LifecycleState)

	result["status"] = string(obj.Status)

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
