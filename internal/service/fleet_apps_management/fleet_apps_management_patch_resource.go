// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package fleet_apps_management

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_fleet_apps_management "github.com/oracle/oci-go-sdk/v65/fleetappsmanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func FleetAppsManagementPatchResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createFleetAppsManagementPatch,
		Read:     readFleetAppsManagementPatch,
		Update:   updateFleetAppsManagementPatch,
		Delete:   deleteFleetAppsManagementPatch,
		Schema: map[string]*schema.Schema{
			// Required
			"artifact_details": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"category": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"GENERIC",
								"PLATFORM_SPECIFIC",
							}, true),
						},

						// Optional
						"artifact": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"content": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"bucket": {
													Type:     schema.TypeString,
													Required: true,
												},
												"checksum": {
													Type:     schema.TypeString,
													Required: true,
												},
												"namespace": {
													Type:     schema.TypeString,
													Required: true,
												},
												"object": {
													Type:     schema.TypeString,
													Required: true,
												},
												"source_type": {
													Type:             schema.TypeString,
													Required:         true,
													DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
													ValidateFunc: validation.StringInSlice([]string{
														"OBJECT_STORAGE_BUCKET",
													}, true),
												},

												// Optional

												// Computed
											},
										},
									},

									// Computed
								},
							},
						},
						"artifacts": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"architecture": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"content": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"bucket": {
													Type:     schema.TypeString,
													Required: true,
												},
												"checksum": {
													Type:     schema.TypeString,
													Required: true,
												},
												"namespace": {
													Type:     schema.TypeString,
													Required: true,
												},
												"object": {
													Type:     schema.TypeString,
													Required: true,
												},
												"source_type": {
													Type:             schema.TypeString,
													Required:         true,
													DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
													ValidateFunc: validation.StringInSlice([]string{
														"OBJECT_STORAGE_BUCKET",
													}, true),
												},

												// Optional

												// Computed
											},
										},
									},
									"os_type": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},

						// Computed
					},
				},
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"patch_type": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"platform_configuration_id": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"product": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"platform_configuration_id": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"version": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"severity": {
				Type:     schema.TypeString,
				Required: true,
			},
			"time_released": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"dependent_patches": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"id": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
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

			// Computed
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"resource_region": {
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
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createFleetAppsManagementPatch(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementPatchResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementOperationsClient()
	sync.FleetClient = m.(*client.OracleClients).FleetAppsManagementClient()

	return tfresource.CreateResource(d, sync)
}

func readFleetAppsManagementPatch(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementPatchResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementOperationsClient()
	sync.FleetClient = m.(*client.OracleClients).FleetAppsManagementClient()

	return tfresource.ReadResource(sync)
}

func updateFleetAppsManagementPatch(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementPatchResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementOperationsClient()
	sync.FleetClient = m.(*client.OracleClients).FleetAppsManagementClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteFleetAppsManagementPatch(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementPatchResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementOperationsClient()
	sync.FleetClient = m.(*client.OracleClients).FleetAppsManagementClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type FleetAppsManagementPatchResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_fleet_apps_management.FleetAppsManagementOperationsClient
	FleetClient            *oci_fleet_apps_management.FleetAppsManagementClient
	Res                    *oci_fleet_apps_management.Patch
	DisableNotFoundRetries bool
}

func (s *FleetAppsManagementPatchResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *FleetAppsManagementPatchResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *FleetAppsManagementPatchResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_fleet_apps_management.PatchLifecycleStateActive),
	}
}

func (s *FleetAppsManagementPatchResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_fleet_apps_management.PatchLifecycleStateDeleting),
	}
}

func (s *FleetAppsManagementPatchResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_fleet_apps_management.PatchLifecycleStateDeleted),
	}
}

func (s *FleetAppsManagementPatchResourceCrud) Create() error {
	request := oci_fleet_apps_management.CreatePatchRequest{}

	if artifactDetails, ok := s.D.GetOkExists("artifact_details"); ok {
		if tmpList := artifactDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "artifact_details", 0)
			tmp, err := s.mapToArtifactDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ArtifactDetails = tmp
		}
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

	if dependentPatches, ok := s.D.GetOkExists("dependent_patches"); ok {
		interfaces := dependentPatches.([]interface{})
		tmp := make([]oci_fleet_apps_management.DependentPatchDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "dependent_patches", stateDataIndex)
			converted, err := s.mapToDependentPatchDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("dependent_patches") {
			request.DependentPatches = tmp
		}
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

	if patchType, ok := s.D.GetOkExists("patch_type"); ok {
		if tmpList := patchType.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "patch_type", 0)
			tmp, err := s.mapToPatchType(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.PatchType = &tmp
		}
	}

	if product, ok := s.D.GetOkExists("product"); ok {
		if tmpList := product.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "product", 0)
			tmp, err := s.mapToPatchProduct(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Product = &tmp
		}
	}

	if severity, ok := s.D.GetOkExists("severity"); ok {
		request.Severity = oci_fleet_apps_management.PatchSeverityEnum(severity.(string))
	}

	if timeReleased, ok := s.D.GetOkExists("time_released"); ok {
		tmp, err := time.Parse(time.RFC3339, timeReleased.(string))
		if err != nil {
			return err
		}
		request.TimeReleased = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	response, err := s.Client.CreatePatch(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Patch
	return nil
}

func (s *FleetAppsManagementPatchResourceCrud) getPatchFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_fleet_apps_management.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	patchId, err := patchWaitForWorkRequest(workId, "patch",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.FleetClient)

	if err != nil {
		return err
	}
	s.D.SetId(*patchId)

	return s.Get()
}

func patchWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func patchWaitForWorkRequest(wId *string, entityType string, action oci_fleet_apps_management.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_fleet_apps_management.FleetAppsManagementClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "fleet_apps_management")
	retryPolicy.ShouldRetryOperation = patchWorkRequestShouldRetryFunc(timeout)

	response := oci_fleet_apps_management.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
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
		return nil, getErrorFromFleetAppsManagementPatchWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromFleetAppsManagementPatchWorkRequest(client *oci_fleet_apps_management.FleetAppsManagementClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_fleet_apps_management.ActionTypeEnum) error {
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

func (s *FleetAppsManagementPatchResourceCrud) Get() error {
	request := oci_fleet_apps_management.GetPatchRequest{}

	tmp := s.D.Id()
	request.PatchId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	response, err := s.Client.GetPatch(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Patch
	return nil
}

func (s *FleetAppsManagementPatchResourceCrud) Update() error {
	request := oci_fleet_apps_management.UpdatePatchRequest{}

	if artifactDetails, ok := s.D.GetOkExists("artifact_details"); ok {
		if tmpList := artifactDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "artifact_details", 0)
			tmp, err := s.mapToArtifactDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ArtifactDetails = tmp
		}
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

	if dependentPatches, ok := s.D.GetOkExists("dependent_patches"); ok {
		interfaces := dependentPatches.([]interface{})
		tmp := make([]oci_fleet_apps_management.DependentPatchDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "dependent_patches", stateDataIndex)
			converted, err := s.mapToDependentPatchDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("dependent_patches") {
			request.DependentPatches = tmp
		}
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.PatchId = &tmp

	if patchType, ok := s.D.GetOkExists("patch_type"); ok {
		if tmpList := patchType.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "patch_type", 0)
			tmp, err := s.mapToPatchType(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.PatchType = &tmp
		}
	}

	if product, ok := s.D.GetOkExists("product"); ok {
		if tmpList := product.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "product", 0)
			tmp, err := s.mapToPatchProduct(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Product = &tmp
		}
	}

	if severity, ok := s.D.GetOkExists("severity"); ok {
		request.Severity = oci_fleet_apps_management.PatchSeverityEnum(severity.(string))
	}

	if timeReleased, ok := s.D.GetOkExists("time_released"); ok {
		tmp, err := time.Parse(time.RFC3339, timeReleased.(string))
		if err != nil {
			return err
		}
		request.TimeReleased = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	response, err := s.Client.UpdatePatch(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getPatchFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management"), oci_fleet_apps_management.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *FleetAppsManagementPatchResourceCrud) Delete() error {
	request := oci_fleet_apps_management.DeletePatchRequest{}

	tmp := s.D.Id()
	request.PatchId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	response, err := s.Client.DeletePatch(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := patchWaitForWorkRequest(workId, "patch",
		oci_fleet_apps_management.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.FleetClient)
	return delWorkRequestErr
}

func (s *FleetAppsManagementPatchResourceCrud) SetData() error {
	if s.Res.ArtifactDetails != nil {
		artifactDetailsArray := []interface{}{}
		if artifactDetailsMap := ArtifactDetailsToMap(&s.Res.ArtifactDetails); artifactDetailsMap != nil {
			artifactDetailsArray = append(artifactDetailsArray, artifactDetailsMap)
		}
		s.D.Set("artifact_details", artifactDetailsArray)
	} else {
		s.D.Set("artifact_details", nil)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	dependentPatches := []interface{}{}
	for _, item := range s.Res.DependentPatches {
		dependentPatches = append(dependentPatches, DependentPatchDetailsToMap(item))
	}
	s.D.Set("dependent_patches", dependentPatches)

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.PatchType != nil {
		s.D.Set("patch_type", []interface{}{PatchTypeToMap(s.Res.PatchType)})
	} else {
		s.D.Set("patch_type", nil)
	}

	if s.Res.Product != nil {
		s.D.Set("product", []interface{}{PatchProductToMap(s.Res.Product)})
	} else {
		s.D.Set("product", nil)
	}

	if s.Res.ResourceRegion != nil {
		s.D.Set("resource_region", *s.Res.ResourceRegion)
	}

	s.D.Set("severity", s.Res.Severity)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeReleased != nil {
		s.D.Set("time_released", s.Res.TimeReleased.Format(time.RFC3339Nano))
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	s.D.Set("type", s.Res.Type)

	return nil
}

func (s *FleetAppsManagementPatchResourceCrud) mapToArtifactDetails(fieldKeyFormat string) (oci_fleet_apps_management.ArtifactDetails, error) {
	var baseObject oci_fleet_apps_management.ArtifactDetails
	//discriminator
	categoryRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "category"))
	var category string
	if ok {
		category = categoryRaw.(string)
	} else {
		category = "" // default value
	}
	switch strings.ToLower(category) {
	case strings.ToLower("GENERIC"):
		details := oci_fleet_apps_management.GenericArtifactDetails{}
		if artifact, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "artifact")); ok {
			if tmpList := artifact.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "artifact"), 0)
				tmp, err := s.mapToGenericArtifact(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert artifact, encountered error: %v", err)
				}
				details.Artifact = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("PLATFORM_SPECIFIC"):
		details := oci_fleet_apps_management.PlatformSpecificArtifactDetails{}
		if artifacts, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "artifacts")); ok {
			interfaces := artifacts.([]interface{})
			tmp := make([]oci_fleet_apps_management.PlatformSpecificArtifact, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "artifacts"), stateDataIndex)
				converted, err := s.mapToPlatformSpecificArtifact(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "artifacts")) {
				details.Artifacts = tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown category '%v' was specified", category)
	}
	return baseObject, nil
}

func ArtifactDetailsToMap(obj *oci_fleet_apps_management.ArtifactDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_fleet_apps_management.GenericArtifactDetails:
		result["category"] = "GENERIC"

		if v.Artifact != nil {
			result["artifact"] = []interface{}{GenericArtifactToMap(v.Artifact)}
		}
	case oci_fleet_apps_management.PlatformSpecificArtifactDetails:
		result["category"] = "PLATFORM_SPECIFIC"

		artifacts := []interface{}{}
		for _, item := range v.Artifacts {
			artifacts = append(artifacts, PlatformSpecificArtifactToMap(item))
		}
		result["artifacts"] = artifacts
	default:
		log.Printf("[WARN] Received 'category' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *FleetAppsManagementPatchResourceCrud) mapToContentDetails(fieldKeyFormat string) (oci_fleet_apps_management.ContentDetails, error) {
	var baseObject oci_fleet_apps_management.ContentDetails
	//discriminator
	sourceTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_type"))
	var sourceType string
	if ok {
		sourceType = sourceTypeRaw.(string)
	} else {
		sourceType = "" // default value
	}
	switch strings.ToLower(sourceType) {
	case strings.ToLower("OBJECT_STORAGE_BUCKET"):
		details := oci_fleet_apps_management.ObjectStorageBucketContentDetails{}
		if bucket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
			tmp := bucket.(string)
			details.BucketName = &tmp
		}
		if checksum, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "checksum")); ok {
			tmp := checksum.(string)
			details.Checksum = &tmp
		}
		if namespace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
			tmp := namespace.(string)
			details.NamespaceName = &tmp
		}
		if object, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object")); ok {
			tmp := object.(string)
			details.ObjectName = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown source_type '%v' was specified", sourceType)
	}
	return baseObject, nil
}

func (s *FleetAppsManagementPatchResourceCrud) mapToDependentPatchDetails(fieldKeyFormat string) (oci_fleet_apps_management.DependentPatchDetails, error) {
	result := oci_fleet_apps_management.DependentPatchDetails{}

	if id, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "id")); ok {
		tmp := id.(string)
		result.Id = &tmp
	}

	return result, nil
}

func DependentPatchDetailsToMap(obj oci_fleet_apps_management.DependentPatchDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	return result
}

func (s *FleetAppsManagementPatchResourceCrud) mapToGenericArtifact(fieldKeyFormat string) (oci_fleet_apps_management.GenericArtifact, error) {
	result := oci_fleet_apps_management.GenericArtifact{}

	if content, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "content")); ok {
		if tmpList := content.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "content"), 0)
			tmp, err := s.mapToContentDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert content, encountered error: %v", err)
			}
			result.Content = tmp
		}
	}

	return result, nil
}

func GenericArtifactToMap(obj *oci_fleet_apps_management.GenericArtifact) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Content != nil {
		contentArray := []interface{}{}
		if contentMap := ContentDetailsToMap(&obj.Content); contentMap != nil {
			contentArray = append(contentArray, contentMap)
		}
		result["content"] = contentArray
	}

	return result
}

func (s *FleetAppsManagementPatchResourceCrud) mapToPatchProduct(fieldKeyFormat string) (oci_fleet_apps_management.PatchProduct, error) {
	result := oci_fleet_apps_management.PatchProduct{}

	if platformConfigurationId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "platform_configuration_id")); ok {
		tmp := platformConfigurationId.(string)
		result.PlatformConfigurationId = &tmp
	}

	if version, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "version")); ok {
		tmp := version.(string)
		result.Version = &tmp
	}

	return result, nil
}

func PatchProductToMap(obj *oci_fleet_apps_management.PatchProduct) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.PlatformConfigurationId != nil {
		result["platform_configuration_id"] = string(*obj.PlatformConfigurationId)
	}

	if obj.Version != nil {
		result["version"] = string(*obj.Version)
	}

	return result
}

func PatchSummaryToMap(obj oci_fleet_apps_management.PatchSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ArtifactDetails != nil {
		artifactDetailsArray := []interface{}{}
		if artifactDetailsMap := ArtifactDetailsToMap(&obj.ArtifactDetails); artifactDetailsMap != nil {
			artifactDetailsArray = append(artifactDetailsArray, artifactDetailsMap)
		}
		result["artifact_details"] = artifactDetailsArray
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.PatchType != nil {
		result["patch_type"] = []interface{}{PatchTypeToMap(obj.PatchType)}
	}

	if obj.Product != nil {
		result["product"] = []interface{}{PatchProductToMap(obj.Product)}
	}

	if obj.ResourceRegion != nil {
		result["resource_region"] = string(*obj.ResourceRegion)
	}

	result["severity"] = string(obj.Severity)

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

	result["type"] = string(obj.Type)

	return result
}

func (s *FleetAppsManagementPatchResourceCrud) mapToPatchType(fieldKeyFormat string) (oci_fleet_apps_management.PatchType, error) {
	result := oci_fleet_apps_management.PatchType{}

	if platformConfigurationId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "platform_configuration_id")); ok {
		tmp := platformConfigurationId.(string)
		result.PlatformConfigurationId = &tmp
	}

	return result, nil
}

func PatchTypeToMap(obj *oci_fleet_apps_management.PatchType) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.PlatformConfigurationId != nil {
		result["platform_configuration_id"] = string(*obj.PlatformConfigurationId)
	}

	return result
}

func (s *FleetAppsManagementPatchResourceCrud) mapToPlatformSpecificArtifact(fieldKeyFormat string) (oci_fleet_apps_management.PlatformSpecificArtifact, error) {
	result := oci_fleet_apps_management.PlatformSpecificArtifact{}

	if architecture, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "architecture")); ok {
		result.Architecture = oci_fleet_apps_management.PlatformSpecificArtifactArchitectureEnum(architecture.(string))
	}

	if content, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "content")); ok {
		if tmpList := content.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "content"), 0)
			tmp, err := s.mapToContentDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert content, encountered error: %v", err)
			}
			result.Content = tmp
		}
	}

	if osType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "os_type")); ok {
		result.OsType = oci_fleet_apps_management.PlatformSpecificArtifactOsTypeEnum(osType.(string))
	}

	return result, nil
}

func PlatformSpecificArtifactToMap(obj oci_fleet_apps_management.PlatformSpecificArtifact) map[string]interface{} {
	result := map[string]interface{}{}

	result["architecture"] = string(obj.Architecture)

	if obj.Content != nil {
		contentArray := []interface{}{}
		if contentMap := ContentDetailsToMap(&obj.Content); contentMap != nil {
			contentArray = append(contentArray, contentMap)
		}
		result["content"] = contentArray
	}

	result["os_type"] = string(obj.OsType)

	return result
}
