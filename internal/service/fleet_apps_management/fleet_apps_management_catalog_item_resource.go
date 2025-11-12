// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package fleet_apps_management

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_fleet_apps_management "github.com/oracle/oci-go-sdk/v65/fleetappsmanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func FleetAppsManagementCatalogItemResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createFleetAppsManagementCatalogItem,
		Read:     readFleetAppsManagementCatalogItem,
		Update:   updateFleetAppsManagementCatalogItem,
		Delete:   deleteFleetAppsManagementCatalogItem,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"config_source_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"package_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"catalog_source_payload": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"config_source_type": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"GIT_CATALOG_SOURCE",
								"MARKETPLACE_CATALOG_SOURCE",
								"PAR_CATALOG_SOURCE",
								"STACK_TEMPLATE_CATALOG_SOURCE",
							}, true),
						},

						// Optional
						"access_uri": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ForceNew:         true,
							DiffSuppressFunc: MaskedUriSuppressDiff,
						},
						"branch_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"bucket": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"configuration_source_provider_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"listing_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"long_description": {
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
						"object": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"repository_url": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"template_display_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"time_expires": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
						},
						"version": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"working_directory": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"zip_file_base64encoded": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"defined_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				// DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					// k looks like "defined_tags.%", "defined_tags.<key>"
					if strings.HasPrefix(k, "defined_tags.Oracle-Tags.CreatedBy") ||
						strings.HasPrefix(k, "defined_tags.Oracle-Tags.CreatedOn") {
						return true
					}
					return false
				},
				Elem: schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"listing_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"listing_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"short_description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"time_released": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
			},
			"version_description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"clone_catalog_item_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			// Computed
			"catalog_result_payload": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"branch_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"config_result_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"configuration_source_provider_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"package_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"repository_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"template_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_expires": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"working_directory": {
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
			"should_list_public_items": {
				Type:     schema.TypeBool,
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
			"time_backfill_last_checked": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_last_checked": {
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

func MaskedUriSuppressDiff(k, old, new string, d *schema.ResourceData) bool {
	// Check if the value is a masked PAR URL
	if strings.Contains(old, "/p/***/n/") {
		// Suppress the difference
		return true
	}
	// Otherwise, allow the difference to be displayed
	return false
}
func createFleetAppsManagementCatalogItem(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementCatalogItemResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementCatalogClient()
	sync.WorkRequestClient = m.(*client.OracleClients).FleetAppsManagementFleetAppsManagementWorkRequestClient()

	if e := tfresource.CreateResource(d, sync); e != nil {
		return e
	}

	if _, ok := sync.D.GetOkExists("clone_catalog_item_trigger"); ok {
		err := sync.CloneCatalogItem()
		if err != nil {
			return err
		}
	}
	return nil

}

func readFleetAppsManagementCatalogItem(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementCatalogItemResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementCatalogClient()

	return tfresource.ReadResource(sync)
}

func updateFleetAppsManagementCatalogItem(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementCatalogItemResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementCatalogClient()
	sync.WorkRequestClient = m.(*client.OracleClients).FleetAppsManagementFleetAppsManagementWorkRequestClient()

	if _, ok := sync.D.GetOkExists("clone_catalog_item_trigger"); ok && sync.D.HasChange("clone_catalog_item_trigger") {
		oldRaw, newRaw := sync.D.GetChange("clone_catalog_item_trigger")
		oldValue := oldRaw.(int)
		newValue := newRaw.(int)
		if oldValue < newValue {
			err := sync.CloneCatalogItem()

			if err != nil {
				return err
			}
		} else {
			sync.D.Set("clone_catalog_item_trigger", oldRaw)
			return fmt.Errorf("new value of trigger should be greater than the old value")
		}
	}

	if err := tfresource.UpdateResource(d, sync); err != nil {
		return err
	}

	return nil
}

func deleteFleetAppsManagementCatalogItem(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementCatalogItemResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementCatalogClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).FleetAppsManagementFleetAppsManagementWorkRequestClient()

	return tfresource.DeleteResource(d, sync)
}

type FleetAppsManagementCatalogItemResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_fleet_apps_management.FleetAppsManagementCatalogClient
	Res                    *oci_fleet_apps_management.CatalogItem
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_fleet_apps_management.FleetAppsManagementWorkRequestClient
}

func (s *FleetAppsManagementCatalogItemResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *FleetAppsManagementCatalogItemResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_fleet_apps_management.CatalogItemLifecycleStateCreating),
	}
}

func (s *FleetAppsManagementCatalogItemResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_fleet_apps_management.CatalogItemLifecycleStateActive),
	}
}

func (s *FleetAppsManagementCatalogItemResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_fleet_apps_management.CatalogItemLifecycleStateDeleting),
	}
}

func (s *FleetAppsManagementCatalogItemResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_fleet_apps_management.CatalogItemLifecycleStateDeleted),
	}
}

func (s *FleetAppsManagementCatalogItemResourceCrud) Create() error {
	request := oci_fleet_apps_management.CreateCatalogItemRequest{}

	if catalogSourcePayload, ok := s.D.GetOkExists("catalog_source_payload"); ok {
		if tmpList := catalogSourcePayload.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "catalog_source_payload", 0)
			tmp, err := s.mapToCatalogSourcePayload(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.CatalogSourcePayload = tmp
		}
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if configSourceType, ok := s.D.GetOkExists("config_source_type"); ok {
		request.ConfigSourceType = oci_fleet_apps_management.CatalogItemConfigSourceTypeEnum(configSourceType.(string))
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

	if listingId, ok := s.D.GetOkExists("listing_id"); ok {
		tmp := listingId.(string)
		request.ListingId = &tmp
	}

	if listingVersion, ok := s.D.GetOkExists("listing_version"); ok {
		tmp := listingVersion.(string)
		request.ListingVersion = &tmp
	}

	if packageType, ok := s.D.GetOkExists("package_type"); ok {
		request.PackageType = oci_fleet_apps_management.CatalogItemPackageTypeEnum(packageType.(string))
	}

	if shortDescription, ok := s.D.GetOkExists("short_description"); ok {
		tmp := shortDescription.(string)
		request.ShortDescription = &tmp
	}

	if timeReleased, ok := s.D.GetOkExists("time_released"); ok {
		tmp, err := time.Parse(time.RFC3339, timeReleased.(string))
		if err != nil {
			return err
		}
		request.TimeReleased = &oci_common.SDKTime{Time: tmp}
	}

	if versionDescription, ok := s.D.GetOkExists("version_description"); ok {
		tmp := versionDescription.(string)
		request.VersionDescription = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	response, err := s.Client.CreateCatalogItem(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	workRequestResponse := oci_fleet_apps_management.GetWorkRequestResponse{}
	workRequestResponse, err = s.WorkRequestClient.GetWorkRequest(context.Background(),
		oci_fleet_apps_management.GetWorkRequestRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management"),
			},
		})
	if err == nil {
		// The work request response contains an array of objects
		for _, res := range workRequestResponse.Resources {
			if res.EntityType != nil && strings.Contains(strings.ToLower(*res.EntityType), "famscatalogitem") && res.Identifier != nil {
				s.D.SetId(*res.Identifier)
				break
			}
		}
	}
	return s.getCatalogItemFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management"), oci_fleet_apps_management.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *FleetAppsManagementCatalogItemResourceCrud) getCatalogItemFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_fleet_apps_management.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	catalogItemId, err := catalogItemWaitForWorkRequest(workId, "famscatalogitem",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.WorkRequestClient)

	if err != nil {
		return err
	}
	s.D.SetId(*catalogItemId)

	return s.Get()
}

func catalogItemWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "fleet_apps_management", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_fleet_apps_management.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func catalogItemWaitForWorkRequest(wId *string, entityType string, action oci_fleet_apps_management.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_fleet_apps_management.FleetAppsManagementWorkRequestClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "fleet_apps_management")
	retryPolicy.ShouldRetryOperation = catalogItemWorkRequestShouldRetryFunc(timeout)

	response := oci_fleet_apps_management.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_fleet_apps_management.OperationStatusInProgress),
			string(oci_fleet_apps_management.OperationStatusAccepted),
			string(oci_fleet_apps_management.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_fleet_apps_management.OperationStatusSucceeded),
			string(oci_fleet_apps_management.OperationStatusFailed),
			string(oci_fleet_apps_management.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_fleet_apps_management.GetWorkRequestRequest{
					WorkRequestId: wId,
					RequestMetadata: oci_common.RequestMetadata{
						RetryPolicy: retryPolicy,
					},
				})
			wr := &response.WorkRequest
			return wr, string(wr.Status), err
		},
		Timeout: timeout,
	}
	if _, e := stateConf.WaitForState(); e != nil {
		return nil, e
	}

	var identifier *string
	// The work request response contains an array of objects that finished the operation
	for _, res := range response.Resources {
		if strings.Contains(strings.ToLower(*res.EntityType), entityType) {
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_fleet_apps_management.OperationStatusFailed || response.Status == oci_fleet_apps_management.OperationStatusCanceled {
		return nil, getErrorFromFleetAppsManagementCatalogItemWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromFleetAppsManagementCatalogItemWorkRequest(client *oci_fleet_apps_management.FleetAppsManagementWorkRequestClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_fleet_apps_management.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_fleet_apps_management.ListWorkRequestErrorsRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: retryPolicy,
			},
		})
	if err != nil {
		return err
	}

	allErrs := make([]string, 0)
	for _, wrkErr := range response.Items {
		allErrs = append(allErrs, *wrkErr.Message)
	}
	errorMessage := strings.Join(allErrs, "\n")

	workRequestErr := fmt.Errorf("work request did not succeed, workId: %s, entity: %s, action: %s. Message: %s", *workId, entityType, action, errorMessage)

	return workRequestErr
}

func (s *FleetAppsManagementCatalogItemResourceCrud) Get() error {
	request := oci_fleet_apps_management.GetCatalogItemRequest{}

	tmp := s.D.Id()
	request.CatalogItemId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	response, err := s.Client.GetCatalogItem(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.CatalogItem
	return nil
}

func (s *FleetAppsManagementCatalogItemResourceCrud) Update() error {

	if _, ok := s.D.GetOkExists("compartmentId"); ok && s.D.HasChange("compartmentId") {
		err := s.ChangeCatalogItemCompartment()
		if err != nil {
			return err
		}
	}
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_fleet_apps_management.UpdateCatalogItemRequest{}

	tmp := s.D.Id()
	request.CatalogItemId = &tmp

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

	if shortDescription, ok := s.D.GetOkExists("short_description"); ok {
		tmp := shortDescription.(string)
		request.ShortDescription = &tmp
	}

	if versionDescription, ok := s.D.GetOkExists("version_description"); ok {
		tmp := versionDescription.(string)
		request.VersionDescription = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	response, err := s.Client.UpdateCatalogItem(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getCatalogItemFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management"), oci_fleet_apps_management.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *FleetAppsManagementCatalogItemResourceCrud) Delete() error {
	request := oci_fleet_apps_management.DeleteCatalogItemRequest{}

	tmp := s.D.Id()
	request.CatalogItemId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	response, err := s.Client.DeleteCatalogItem(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := catalogItemWaitForWorkRequest(workId, "famscatalogitem",
		oci_fleet_apps_management.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.WorkRequestClient)
	return delWorkRequestErr
}

func (s *FleetAppsManagementCatalogItemResourceCrud) SetData() error {
	if s.Res.CatalogResultPayload != nil {
		catalogResultPayloadArray := []interface{}{}
		if catalogResultPayloadMap := CatalogResultPayloadToMap(&s.Res.CatalogResultPayload); catalogResultPayloadMap != nil {
			catalogResultPayloadArray = append(catalogResultPayloadArray, catalogResultPayloadMap)
		}
		s.D.Set("catalog_result_payload", catalogResultPayloadArray)
	} else {
		s.D.Set("catalog_result_payload", nil)
	}

	if s.Res.CatalogSourcePayload != nil {
		catalogSourcePayloadArray := []interface{}{}
		if catalogSourcePayloadMap := CatalogSourcePayloadToMap(&s.Res.CatalogSourcePayload); catalogSourcePayloadMap != nil {
			catalogSourcePayloadArray = append(catalogSourcePayloadArray, catalogSourcePayloadMap)
		}
		s.D.Set("catalog_source_payload", catalogSourcePayloadArray)
	} else {
		s.D.Set("catalog_source_payload", nil)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("config_source_type", s.Res.ConfigSourceType)

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

	if s.Res.ListingId != nil {
		s.D.Set("listing_id", *s.Res.ListingId)
	}

	if s.Res.ListingVersion != nil {
		s.D.Set("listing_version", *s.Res.ListingVersion)
	}

	s.D.Set("package_type", s.Res.PackageType)

	if s.Res.ShortDescription != nil {
		s.D.Set("short_description", *s.Res.ShortDescription)
	}

	if s.Res.ShouldListPublicItems != nil {
		s.D.Set("should_list_public_items", *s.Res.ShouldListPublicItems)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeBackfillLastChecked != nil {
		s.D.Set("time_backfill_last_checked", s.Res.TimeBackfillLastChecked.String())
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeLastChecked != nil {
		s.D.Set("time_last_checked", s.Res.TimeLastChecked.String())
	}

	if s.Res.TimeReleased != nil {
		s.D.Set("time_released", s.Res.TimeReleased.Format(time.RFC3339Nano))
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.VersionDescription != nil {
		s.D.Set("version_description", *s.Res.VersionDescription)
	}

	return nil
}

func (s *FleetAppsManagementCatalogItemResourceCrud) CloneCatalogItem() error {
	request := oci_fleet_apps_management.CloneCatalogItemRequest{}

	idTmp := s.D.Id()
	request.CatalogItemId = &idTmp

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if versionDescription, ok := s.D.GetOkExists("version_description"); ok {
		tmp := versionDescription.(string)
		request.VersionDescription = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	response, err := s.Client.CloneCatalogItem(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	val := s.D.Get("clone_catalog_item_trigger")
	s.D.Set("clone_catalog_item_trigger", val)

	s.Res = &response.CatalogItem
	return nil
}

func (s *FleetAppsManagementCatalogItemResourceCrud) ChangeCatalogItemCompartment() error {
	request := oci_fleet_apps_management.ChangeCatalogItemCompartmentRequest{}

	idTmp := s.D.Id()
	request.CatalogItemId = &idTmp

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	_, err := s.Client.ChangeCatalogItemCompartment(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

func CatalogItemSummaryToMap(obj oci_fleet_apps_management.CatalogItemSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CatalogResultPayload != nil {
		catalogResultPayloadArray := []interface{}{}
		if catalogResultPayloadMap := CatalogResultPayloadToMap(&obj.CatalogResultPayload); catalogResultPayloadMap != nil {
			catalogResultPayloadArray = append(catalogResultPayloadArray, catalogResultPayloadMap)
		}
		result["catalog_result_payload"] = catalogResultPayloadArray
	}

	if obj.CatalogSourcePayload != nil {
		catalogSourcePayloadArray := []interface{}{}
		if catalogSourcePayloadMap := CatalogSourcePayloadToMap(&obj.CatalogSourcePayload); catalogSourcePayloadMap != nil {
			catalogSourcePayloadArray = append(catalogSourcePayloadArray, catalogSourcePayloadMap)
		}
		result["catalog_source_payload"] = catalogSourcePayloadArray
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	result["config_source_type"] = string(obj.ConfigSourceType)

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

	if obj.ListingId != nil {
		result["listing_id"] = string(*obj.ListingId)
	}

	if obj.ListingVersion != nil {
		result["listing_version"] = string(*obj.ListingVersion)
	}

	result["package_type"] = string(obj.PackageType)

	if obj.ShortDescription != nil {
		result["short_description"] = string(*obj.ShortDescription)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeReleased != nil {
		result["time_released"] = obj.TimeReleased.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	if obj.VersionDescription != nil {
		result["version_description"] = string(*obj.VersionDescription)
	}

	return result
}

func CatalogResultPayloadToMap(obj *oci_fleet_apps_management.CatalogResultPayload) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_fleet_apps_management.CatalogGitResultConfig:
		result["config_result_type"] = "GIT_RESULT_CONFIG"

		if v.BranchName != nil {
			result["branch_name"] = string(*v.BranchName)
		}

		if v.ConfigurationSourceProviderId != nil {
			result["configuration_source_provider_id"] = string(*v.ConfigurationSourceProviderId)
		}

		if v.RepositoryUrl != nil {
			result["repository_url"] = string(*v.RepositoryUrl)
		}

		if v.WorkingDirectory != nil {
			result["working_directory"] = string(*v.WorkingDirectory)
		}
	case oci_fleet_apps_management.CatalogParResultConfig:
		result["config_result_type"] = "PAR_RESULT_CONFIG"

		if v.PackageUrl != nil {
			result["package_url"] = string(*v.PackageUrl)
		}

		if v.TimeExpires != nil {
			result["time_expires"] = v.TimeExpires.Format(time.RFC3339Nano)
		}

		if v.WorkingDirectory != nil {
			result["working_directory"] = string(*v.WorkingDirectory)
		}
	case oci_fleet_apps_management.CatalogTemplateResultConfig:
		result["config_result_type"] = "TEMPLATE_RESULT_CONFIG"

		if v.TemplateId != nil {
			result["template_id"] = string(*v.TemplateId)
		}

		if v.WorkingDirectory != nil {
			result["working_directory"] = string(*v.WorkingDirectory)
		}
	default:
		log.Printf("[WARN] Received 'config_result_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *FleetAppsManagementCatalogItemResourceCrud) mapToCatalogSourcePayload(fieldKeyFormat string) (oci_fleet_apps_management.CatalogSourcePayload, error) {
	var baseObject oci_fleet_apps_management.CatalogSourcePayload
	//discriminator
	configSourceTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "config_source_type"))
	var configSourceType string
	if ok {
		configSourceType = configSourceTypeRaw.(string)
	} else {
		configSourceType = "" // default value
	}
	switch strings.ToLower(configSourceType) {
	case strings.ToLower("GIT_CATALOG_SOURCE"):
		details := oci_fleet_apps_management.CatalogGitSourceConfig{}
		if branchName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "branch_name")); ok {
			tmp := branchName.(string)
			details.BranchName = &tmp
		}
		if configurationSourceProviderId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "configuration_source_provider_id")); ok {
			tmp := configurationSourceProviderId.(string)
			details.ConfigurationSourceProviderId = &tmp
		}
		if repositoryUrl, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "repository_url")); ok {
			tmp := repositoryUrl.(string)
			details.RepositoryUrl = &tmp
		}
		if workingDirectory, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "working_directory")); ok {
			tmp := workingDirectory.(string)
			details.WorkingDirectory = &tmp
		}
		baseObject = details
	case strings.ToLower("MARKETPLACE_CATALOG_SOURCE"):
		details := oci_fleet_apps_management.CatalogMarketplaceSourceConfig{}
		if listingId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "listing_id")); ok {
			tmp := listingId.(string)
			details.ListingId = &tmp
		}
		if version, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "version")); ok {
			tmp := version.(string)
			details.Version = &tmp
		}
		if workingDirectory, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "working_directory")); ok {
			tmp := workingDirectory.(string)
			details.WorkingDirectory = &tmp
		}
		baseObject = details
	case strings.ToLower("PAR_CATALOG_SOURCE"):
		details := oci_fleet_apps_management.CatalogParSourceConfig{}
		if accessUri, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "access_uri")); ok {
			tmp := accessUri.(string)
			details.AccessUri = &tmp
		}
		if bucket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
			tmp := bucket.(string)
			details.BucketName = &tmp
		}
		if namespace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
			tmp := namespace.(string)
			details.NamespaceName = &tmp
		}
		if object, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object")); ok {
			tmp := object.(string)
			details.ObjectName = &tmp
		}
		if timeExpires, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_expires")); ok {
			tmp, err := time.Parse(time.RFC3339, timeExpires.(string))
			if err != nil {
				return details, err
			}
			details.TimeExpires = &oci_common.SDKTime{Time: tmp}
		}
		if workingDirectory, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "working_directory")); ok {
			tmp := workingDirectory.(string)
			details.WorkingDirectory = &tmp
		}
		baseObject = details
	case strings.ToLower("STACK_TEMPLATE_CATALOG_SOURCE"):
		details := oci_fleet_apps_management.CatalogSourceTemplateConfig{}
		if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if longDescription, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "long_description")); ok {
			tmp := longDescription.(string)
			details.LongDescription = &tmp
		}
		if templateDisplayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "template_display_name")); ok {
			tmp := templateDisplayName.(string)
			details.TemplateDisplayName = &tmp
		}
		if zipFileBase64Encoded, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "zip_file_base64encoded")); ok {
			tmp := zipFileBase64Encoded.(string)
			details.ZipFileBase64Encoded = &tmp
		}
		if workingDirectory, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "working_directory")); ok {
			tmp := workingDirectory.(string)
			details.WorkingDirectory = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown config_source_type '%v' was specified", configSourceType)
	}
	return baseObject, nil
}

func CatalogSourcePayloadToMap(obj *oci_fleet_apps_management.CatalogSourcePayload) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_fleet_apps_management.CatalogGitSourceConfig:
		result["config_source_type"] = "GIT_CATALOG_SOURCE"

		if v.BranchName != nil {
			result["branch_name"] = string(*v.BranchName)
		}

		if v.ConfigurationSourceProviderId != nil {
			result["configuration_source_provider_id"] = string(*v.ConfigurationSourceProviderId)
		}

		if v.RepositoryUrl != nil {
			result["repository_url"] = string(*v.RepositoryUrl)
		}

		if v.WorkingDirectory != nil {
			result["working_directory"] = string(*v.WorkingDirectory)
		}
	case oci_fleet_apps_management.CatalogMarketplaceSourceConfig:
		result["config_source_type"] = "MARKETPLACE_CATALOG_SOURCE"

		if v.ListingId != nil {
			result["listing_id"] = string(*v.ListingId)
		}

		if v.Version != nil {
			result["version"] = string(*v.Version)
		}

		if v.WorkingDirectory != nil {
			result["working_directory"] = string(*v.WorkingDirectory)
		}
	case oci_fleet_apps_management.CatalogParSourceConfig:
		result["config_source_type"] = "PAR_CATALOG_SOURCE"

		if v.AccessUri != nil {
			result["access_uri"] = string(*v.AccessUri)
		}

		if v.BucketName != nil {
			result["bucket"] = string(*v.BucketName)
		}

		if v.NamespaceName != nil {
			result["namespace"] = string(*v.NamespaceName)
		}

		if v.ObjectName != nil {
			result["object"] = string(*v.ObjectName)
		}

		if v.TimeExpires != nil {
			result["time_expires"] = v.TimeExpires.Format(time.RFC3339Nano)
		}

		if v.WorkingDirectory != nil {
			result["working_directory"] = string(*v.WorkingDirectory)
		}
	case oci_fleet_apps_management.CatalogSourceTemplateConfig:
		result["config_source_type"] = "STACK_TEMPLATE_CATALOG_SOURCE"

		if v.Description != nil {
			result["description"] = string(*v.Description)
		}

		if v.LongDescription != nil {
			result["long_description"] = string(*v.LongDescription)
		}

		if v.TemplateDisplayName != nil {
			result["template_display_name"] = string(*v.TemplateDisplayName)
		}

		if v.ZipFileBase64Encoded != nil {
			result["zip_file_base64encoded"] = string(*v.ZipFileBase64Encoded)
		}

		if v.WorkingDirectory != nil {
			result["working_directory"] = string(*v.WorkingDirectory)
		}
	default:
		log.Printf("[WARN] Received 'config_source_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *FleetAppsManagementCatalogItemResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_fleet_apps_management.ChangeCatalogItemCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.CatalogItemId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	response, err := s.Client.ChangeCatalogItemCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getCatalogItemFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management"), oci_fleet_apps_management.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
