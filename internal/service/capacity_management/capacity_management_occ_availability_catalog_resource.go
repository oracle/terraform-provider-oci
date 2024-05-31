// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package capacity_management

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_capacity_management "github.com/oracle/oci-go-sdk/v65/capacitymanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CapacityManagementOccAvailabilityCatalogResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCapacityManagementOccAvailabilityCatalog,
		Read:     readCapacityManagementOccAvailabilityCatalog,
		Update:   updateCapacityManagementOccAvailabilityCatalog,
		Delete:   deleteCapacityManagementOccAvailabilityCatalog,
		Schema: map[string]*schema.Schema{
			// Required
			"base64encoded_catalog_details": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
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
			"metadata_details": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"format_version": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional

						// Computed
					},
				},
			},

			// Computed
			"catalog_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"details": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"available_quantity": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"catalog_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"date_expected_capacity_handover": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"date_final_customer_order": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"namespace": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"resource_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"resource_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"unit": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"workload_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"lifecycle_details": {
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

func createCapacityManagementOccAvailabilityCatalog(d *schema.ResourceData, m interface{}) error {
	sync := &CapacityManagementOccAvailabilityCatalogResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CapacityManagementClient()

	return tfresource.CreateResource(d, sync)
}

func readCapacityManagementOccAvailabilityCatalog(d *schema.ResourceData, m interface{}) error {
	sync := &CapacityManagementOccAvailabilityCatalogResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CapacityManagementClient()

	return tfresource.ReadResource(sync)
}

func updateCapacityManagementOccAvailabilityCatalog(d *schema.ResourceData, m interface{}) error {
	sync := &CapacityManagementOccAvailabilityCatalogResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CapacityManagementClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteCapacityManagementOccAvailabilityCatalog(d *schema.ResourceData, m interface{}) error {
	sync := &CapacityManagementOccAvailabilityCatalogResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CapacityManagementClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type CapacityManagementOccAvailabilityCatalogResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_capacity_management.CapacityManagementClient
	Res                    *oci_capacity_management.OccAvailabilityCatalog
	DisableNotFoundRetries bool
}

func (s *CapacityManagementOccAvailabilityCatalogResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CapacityManagementOccAvailabilityCatalogResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_capacity_management.OccAvailabilityCatalogLifecycleStateCreating),
	}
}

func (s *CapacityManagementOccAvailabilityCatalogResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_capacity_management.OccAvailabilityCatalogLifecycleStateActive),
	}
}

func (s *CapacityManagementOccAvailabilityCatalogResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_capacity_management.OccAvailabilityCatalogLifecycleStateDeleting),
	}
}

func (s *CapacityManagementOccAvailabilityCatalogResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_capacity_management.OccAvailabilityCatalogLifecycleStateDeleted),
	}
}

func (s *CapacityManagementOccAvailabilityCatalogResourceCrud) Create() error {
	request := oci_capacity_management.CreateOccAvailabilityCatalogRequest{}

	if base64EncodedCatalogDetails, ok := s.D.GetOkExists("base64encoded_catalog_details"); ok {
		tmp := base64EncodedCatalogDetails.(string)
		request.Base64EncodedCatalogDetails = &tmp
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

	if metadataDetails, ok := s.D.GetOkExists("metadata_details"); ok {
		if tmpList := metadataDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "metadata_details", 0)
			tmp, err := s.mapToMetadataDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.MetadataDetails = &tmp
		}
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		request.Namespace = oci_capacity_management.NamespaceEnum(namespace.(string))
	}

	if occCustomerGroupId, ok := s.D.GetOkExists("occ_customer_group_id"); ok {
		tmp := occCustomerGroupId.(string)
		request.OccCustomerGroupId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "capacity_management")

	response, err := s.Client.CreateOccAvailabilityCatalog(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.OccAvailabilityCatalog
	return nil
}

func (s *CapacityManagementOccAvailabilityCatalogResourceCrud) Get() error {
	request := oci_capacity_management.GetOccAvailabilityCatalogRequest{}

	tmp := s.D.Id()
	request.OccAvailabilityCatalogId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "capacity_management")

	response, err := s.Client.GetOccAvailabilityCatalog(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.OccAvailabilityCatalog
	return nil
}

func (s *CapacityManagementOccAvailabilityCatalogResourceCrud) Update() error {
	request := oci_capacity_management.UpdateOccAvailabilityCatalogRequest{}

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
	request.OccAvailabilityCatalogId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "capacity_management")

	response, err := s.Client.UpdateOccAvailabilityCatalog(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.OccAvailabilityCatalog
	return nil
}

func (s *CapacityManagementOccAvailabilityCatalogResourceCrud) Delete() error {
	request := oci_capacity_management.DeleteOccAvailabilityCatalogRequest{}

	tmp := s.D.Id()
	request.OccAvailabilityCatalogId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "capacity_management")

	_, err := s.Client.DeleteOccAvailabilityCatalog(context.Background(), request)
	return err
}

func (s *CapacityManagementOccAvailabilityCatalogResourceCrud) SetData() error {
	s.D.Set("catalog_state", s.Res.CatalogState)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	details := []interface{}{}
	for _, item := range s.Res.Details {
		details = append(details, OccAvailabilitySummaryToMap(item))
	}
	s.D.Set("details", details)

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.MetadataDetails != nil {
		s.D.Set("metadata_details", []interface{}{MetadataDetailsToMap(s.Res.MetadataDetails)})
	} else {
		s.D.Set("metadata_details", nil)
	}

	s.D.Set("namespace", s.Res.Namespace)

	if s.Res.OccCustomerGroupId != nil {
		s.D.Set("occ_customer_group_id", *s.Res.OccCustomerGroupId)
	}

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

func (s *CapacityManagementOccAvailabilityCatalogResourceCrud) mapToMetadataDetails(fieldKeyFormat string) (oci_capacity_management.MetadataDetails, error) {
	result := oci_capacity_management.MetadataDetails{}

	if formatVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "format_version")); ok {
		result.FormatVersion = oci_capacity_management.MetadataDetailsFormatVersionEnum(formatVersion.(string))
	}

	return result, nil
}

func MetadataDetailsToMap(obj *oci_capacity_management.MetadataDetails) map[string]interface{} {
	result := map[string]interface{}{}

	result["format_version"] = string(obj.FormatVersion)

	return result
}

func OccAvailabilityCatalogSummaryToMap(obj oci_capacity_management.OccAvailabilityCatalogSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["catalog_state"] = string(obj.CatalogState)

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

	if obj.MetadataDetails != nil {
		result["metadata_details"] = []interface{}{MetadataDetailsToMap(obj.MetadataDetails)}
	}

	result["namespace"] = string(obj.Namespace)

	if obj.OccCustomerGroupId != nil {
		result["occ_customer_group_id"] = string(*obj.OccCustomerGroupId)
	}

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

func OccAvailabilitySummaryToMap(obj oci_capacity_management.OccAvailabilitySummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AvailableQuantity != nil {
		result["available_quantity"] = strconv.FormatInt(*obj.AvailableQuantity, 10)
	}

	if obj.CatalogId != nil {
		result["catalog_id"] = string(*obj.CatalogId)
	}

	if obj.DateExpectedCapacityHandover != nil {
		result["date_expected_capacity_handover"] = obj.DateExpectedCapacityHandover.String()
	}

	if obj.DateFinalCustomerOrder != nil {
		result["date_final_customer_order"] = obj.DateFinalCustomerOrder.String()
	}

	result["namespace"] = string(obj.Namespace)

	if obj.ResourceName != nil {
		result["resource_name"] = string(*obj.ResourceName)
	}

	if obj.ResourceType != nil {
		result["resource_type"] = string(*obj.ResourceType)
	}

	if obj.Unit != nil {
		result["unit"] = string(*obj.Unit)
	}

	if obj.WorkloadType != nil {
		result["workload_type"] = string(*obj.WorkloadType)
	}

	return result
}
