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

func FleetAppsManagementPlatformConfigurationResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createFleetAppsManagementPlatformConfiguration,
		Read:     readFleetAppsManagementPlatformConfiguration,
		Update:   updateFleetAppsManagementPlatformConfiguration,
		Delete:   deleteFleetAppsManagementPlatformConfiguration,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"config_category_details": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"config_category": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"CREDENTIAL",
								"ENVIRONMENT",
								"PATCH_TYPE",
								"PRODUCT",
								"PRODUCT_STACK",
							}, true),
						},

						// Optional
						"compatible_products": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"display_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"components": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"credentials": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"display_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"patch_types": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"display_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"products": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"display_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"sub_category_details": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"sub_category": {
										Type:             schema.TypeString,
										Required:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"PRODUCT_STACK_AS_PRODUCT",
											"PRODUCT_STACK_GENERIC",
										}, true),
									},

									// Optional
									"components": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"credentials": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"display_name": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},

												// Computed
											},
										},
									},
									"patch_types": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"display_name": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},

												// Computed
											},
										},
									},
									"versions": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									// Computed
								},
							},
						},
						"versions": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						// Computed
					},
				},
			},

			// Optional
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"defined_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
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

func createFleetAppsManagementPlatformConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementPlatformConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementAdminClient()
	sync.FleetClient = m.(*client.OracleClients).FleetAppsManagementClient()

	return tfresource.CreateResource(d, sync)
}

func readFleetAppsManagementPlatformConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementPlatformConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementAdminClient()
	sync.FleetClient = m.(*client.OracleClients).FleetAppsManagementClient()

	return tfresource.ReadResource(sync)
}

func updateFleetAppsManagementPlatformConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementPlatformConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementAdminClient()
	sync.FleetClient = m.(*client.OracleClients).FleetAppsManagementClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteFleetAppsManagementPlatformConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementPlatformConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementAdminClient()
	sync.FleetClient = m.(*client.OracleClients).FleetAppsManagementClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type FleetAppsManagementPlatformConfigurationResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_fleet_apps_management.FleetAppsManagementAdminClient
	FleetClient            *oci_fleet_apps_management.FleetAppsManagementClient
	Res                    *oci_fleet_apps_management.PlatformConfiguration
	DisableNotFoundRetries bool
}

func (s *FleetAppsManagementPlatformConfigurationResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *FleetAppsManagementPlatformConfigurationResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *FleetAppsManagementPlatformConfigurationResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_fleet_apps_management.PlatformConfigurationLifecycleStateActive),
	}
}

func (s *FleetAppsManagementPlatformConfigurationResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_fleet_apps_management.PlatformConfigurationLifecycleStateDeleting),
	}
}

func (s *FleetAppsManagementPlatformConfigurationResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_fleet_apps_management.PlatformConfigurationLifecycleStateDeleted),
	}
}

func (s *FleetAppsManagementPlatformConfigurationResourceCrud) Create() error {
	request := oci_fleet_apps_management.CreatePlatformConfigurationRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if configCategoryDetails, ok := s.D.GetOkExists("config_category_details"); ok {
		if tmpList := configCategoryDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "config_category_details", 0)
			tmp, err := s.mapToConfigCategoryDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ConfigCategoryDetails = tmp
		}
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	response, err := s.Client.CreatePlatformConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.PlatformConfiguration
	return nil
}

func (s *FleetAppsManagementPlatformConfigurationResourceCrud) getPlatformConfigurationFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_fleet_apps_management.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	platformConfigurationId, err := platformConfigurationWaitForWorkRequest(workId, "platformconfiguration",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.FleetClient)

	if err != nil {
		return err
	}
	s.D.SetId(*platformConfigurationId)

	return s.Get()
}

func platformConfigurationWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func platformConfigurationWaitForWorkRequest(wId *string, entityType string, action oci_fleet_apps_management.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_fleet_apps_management.FleetAppsManagementClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "fleet_apps_management")
	retryPolicy.ShouldRetryOperation = platformConfigurationWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromFleetAppsManagementPlatformConfigurationWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromFleetAppsManagementPlatformConfigurationWorkRequest(client *oci_fleet_apps_management.FleetAppsManagementClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_fleet_apps_management.ActionTypeEnum) error {
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

func (s *FleetAppsManagementPlatformConfigurationResourceCrud) Get() error {
	request := oci_fleet_apps_management.GetPlatformConfigurationRequest{}

	tmp := s.D.Id()
	request.PlatformConfigurationId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	response, err := s.Client.GetPlatformConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.PlatformConfiguration
	return nil
}

func (s *FleetAppsManagementPlatformConfigurationResourceCrud) Update() error {
	request := oci_fleet_apps_management.UpdatePlatformConfigurationRequest{}

	if configCategoryDetails, ok := s.D.GetOkExists("config_category_details"); ok {
		if tmpList := configCategoryDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "config_category_details", 0)
			tmp, err := s.mapToConfigCategoryDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ConfigCategoryDetails = tmp
		}
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	tmp := s.D.Id()
	request.PlatformConfigurationId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	response, err := s.Client.UpdatePlatformConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getPlatformConfigurationFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management"), oci_fleet_apps_management.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *FleetAppsManagementPlatformConfigurationResourceCrud) Delete() error {
	request := oci_fleet_apps_management.DeletePlatformConfigurationRequest{}

	tmp := s.D.Id()
	request.PlatformConfigurationId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	response, err := s.Client.DeletePlatformConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := platformConfigurationWaitForWorkRequest(workId, "platformconfiguration",
		oci_fleet_apps_management.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.FleetClient)
	return delWorkRequestErr
}

func (s *FleetAppsManagementPlatformConfigurationResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ConfigCategoryDetails != nil {
		configCategoryDetailsArray := []interface{}{}
		if configCategoryDetailsMap := ConfigCategoryDetailsToMap(&s.Res.ConfigCategoryDetails); configCategoryDetailsMap != nil {
			configCategoryDetailsArray = append(configCategoryDetailsArray, configCategoryDetailsMap)
		}
		s.D.Set("config_category_details", configCategoryDetailsArray)
	} else {
		s.D.Set("config_category_details", nil)
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

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.ResourceRegion != nil {
		s.D.Set("resource_region", *s.Res.ResourceRegion)
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

	s.D.Set("type", s.Res.Type)

	return nil
}

func (s *FleetAppsManagementPlatformConfigurationResourceCrud) mapToConfigAssociationDetails(fieldKeyFormat string) (oci_fleet_apps_management.ConfigAssociationDetails, error) {
	result := oci_fleet_apps_management.ConfigAssociationDetails{}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	if id, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "id")); ok {
		tmp := id.(string)
		result.Id = &tmp
	}

	return result, nil
}

func ConfigAssociationDetailsToMap(obj oci_fleet_apps_management.ConfigAssociationDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	return result
}

func (s *FleetAppsManagementPlatformConfigurationResourceCrud) mapToConfigCategoryDetails(fieldKeyFormat string) (oci_fleet_apps_management.ConfigCategoryDetails, error) {
	var baseObject oci_fleet_apps_management.ConfigCategoryDetails
	//discriminator
	configCategoryRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "config_category"))
	var configCategory string
	if ok {
		configCategory = configCategoryRaw.(string)
	} else {
		configCategory = "" // default value
	}
	switch strings.ToLower(configCategory) {
	case strings.ToLower("CREDENTIAL"):
		details := oci_fleet_apps_management.CredentialConfigCategoryDetails{}
		baseObject = details
	case strings.ToLower("ENVIRONMENT"):
		details := oci_fleet_apps_management.EnvironmentConfigCategoryDetails{}
		baseObject = details
	case strings.ToLower("PATCH_TYPE"):
		details := oci_fleet_apps_management.PatchTypeConfigCategoryDetails{}
		baseObject = details
	case strings.ToLower("PRODUCT"):
		details := oci_fleet_apps_management.ProductConfigCategoryDetails{}
		if compatibleProducts, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compatible_products")); ok {
			interfaces := compatibleProducts.([]interface{})
			tmp := make([]oci_fleet_apps_management.ConfigAssociationDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "compatible_products"), stateDataIndex)
				converted, err := s.mapToConfigAssociationDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "compatible_products")) {
				details.CompatibleProducts = tmp
			}
		}
		if components, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "components")); ok {
			interfaces := components.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "components")) {
				details.Components = tmp
			}
		}
		if credentials, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "credentials")); ok {
			interfaces := credentials.([]interface{})
			tmp := make([]oci_fleet_apps_management.ConfigAssociationDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "credentials"), stateDataIndex)
				converted, err := s.mapToConfigAssociationDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "credentials")) {
				details.Credentials = tmp
			}
		}
		if patchTypes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "patch_types")); ok {
			interfaces := patchTypes.([]interface{})
			tmp := make([]oci_fleet_apps_management.ConfigAssociationDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "patch_types"), stateDataIndex)
				converted, err := s.mapToConfigAssociationDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "patch_types")) {
				details.PatchTypes = tmp
			}
		}
		if versions, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "versions")); ok {
			interfaces := versions.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "versions")) {
				details.Versions = tmp
			}
		}
		baseObject = details
	case strings.ToLower("PRODUCT_STACK"):
		details := oci_fleet_apps_management.ProductStackConfigCategoryDetails{}
		if products, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "products")); ok {
			interfaces := products.([]interface{})
			tmp := make([]oci_fleet_apps_management.ConfigAssociationDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "products"), stateDataIndex)
				converted, err := s.mapToConfigAssociationDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "products")) {
				details.Products = tmp
			}
		}
		if subCategoryDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "sub_category_details")); ok {
			if tmpList := subCategoryDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "sub_category_details"), 0)
				tmp, err := s.mapToProductStackSubCategoryDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert sub_category_details, encountered error: %v", err)
				}
				details.SubCategoryDetails = tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown config_category '%v' was specified", configCategory)
	}
	return baseObject, nil
}

func ConfigCategoryDetailsToMap(obj *oci_fleet_apps_management.ConfigCategoryDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_fleet_apps_management.CredentialConfigCategoryDetails:
		result["config_category"] = "CREDENTIAL"
	case oci_fleet_apps_management.EnvironmentConfigCategoryDetails:
		result["config_category"] = "ENVIRONMENT"
	case oci_fleet_apps_management.PatchTypeConfigCategoryDetails:
		result["config_category"] = "PATCH_TYPE"
	case oci_fleet_apps_management.ProductConfigCategoryDetails:
		result["config_category"] = "PRODUCT"

		compatibleProducts := []interface{}{}
		for _, item := range v.CompatibleProducts {
			compatibleProducts = append(compatibleProducts, ConfigAssociationDetailsToMap(item))
		}
		result["compatible_products"] = compatibleProducts

		result["components"] = v.Components

		credentials := []interface{}{}
		for _, item := range v.Credentials {
			credentials = append(credentials, ConfigAssociationDetailsToMap(item))
		}
		result["credentials"] = credentials

		patchTypes := []interface{}{}
		for _, item := range v.PatchTypes {
			patchTypes = append(patchTypes, ConfigAssociationDetailsToMap(item))
		}
		result["patch_types"] = patchTypes

		result["versions"] = v.Versions
	case oci_fleet_apps_management.ProductStackConfigCategoryDetails:
		result["config_category"] = "PRODUCT_STACK"

		products := []interface{}{}
		for _, item := range v.Products {
			products = append(products, ConfigAssociationDetailsToMap(item))
		}
		result["products"] = products

		if v.SubCategoryDetails != nil {
			subCategoryDetailsArray := []interface{}{}
			if subCategoryDetailsMap := ProductStackSubCategoryDetailsToMap(&v.SubCategoryDetails); subCategoryDetailsMap != nil {
				subCategoryDetailsArray = append(subCategoryDetailsArray, subCategoryDetailsMap)
			}
			result["sub_category_details"] = subCategoryDetailsArray
		}
	default:
		log.Printf("[WARN] Received 'config_category' of unknown type %v", *obj)
		return nil
	}

	return result
}

func PlatformConfigurationSummaryToMap(obj oci_fleet_apps_management.PlatformConfigurationSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.ConfigCategoryDetails != nil {
		configCategoryDetailsArray := []interface{}{}
		if configCategoryDetailsMap := ConfigCategoryDetailsToMap(&obj.ConfigCategoryDetails); configCategoryDetailsMap != nil {
			configCategoryDetailsArray = append(configCategoryDetailsArray, configCategoryDetailsMap)
		}
		result["config_category_details"] = configCategoryDetailsArray
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

	if obj.ResourceRegion != nil {
		result["resource_region"] = string(*obj.ResourceRegion)
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

	result["type"] = string(obj.Type)

	return result
}

func (s *FleetAppsManagementPlatformConfigurationResourceCrud) mapToProductStackSubCategoryDetails(fieldKeyFormat string) (oci_fleet_apps_management.ProductStackSubCategoryDetails, error) {
	var baseObject oci_fleet_apps_management.ProductStackSubCategoryDetails
	//discriminator
	subCategoryRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "sub_category"))
	var subCategory string
	if ok {
		subCategory = subCategoryRaw.(string)
	} else {
		subCategory = "" // default value
	}
	switch strings.ToLower(subCategory) {
	case strings.ToLower("PRODUCT_STACK_AS_PRODUCT"):
		details := oci_fleet_apps_management.ProductStackAsProductSubCategoryDetails{}
		if components, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "components")); ok {
			interfaces := components.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "components")) {
				details.Components = tmp
			}
		}
		if credentials, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "credentials")); ok {
			interfaces := credentials.([]interface{})
			tmp := make([]oci_fleet_apps_management.ConfigAssociationDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "credentials"), stateDataIndex)
				converted, err := s.mapToConfigAssociationDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "credentials")) {
				details.Credentials = tmp
			}
		}
		if patchTypes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "patch_types")); ok {
			interfaces := patchTypes.([]interface{})
			tmp := make([]oci_fleet_apps_management.ConfigAssociationDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "patch_types"), stateDataIndex)
				converted, err := s.mapToConfigAssociationDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "patch_types")) {
				details.PatchTypes = tmp
			}
		}
		if versions, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "versions")); ok {
			interfaces := versions.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "versions")) {
				details.Versions = tmp
			}
		}
		baseObject = details
	case strings.ToLower("PRODUCT_STACK_GENERIC"):
		details := oci_fleet_apps_management.ProductStackGenericSubCategoryDetails{}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown sub_category '%v' was specified", subCategory)
	}
	return baseObject, nil
}

func ProductStackSubCategoryDetailsToMap(obj *oci_fleet_apps_management.ProductStackSubCategoryDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_fleet_apps_management.ProductStackAsProductSubCategoryDetails:
		result["sub_category"] = "PRODUCT_STACK_AS_PRODUCT"

		result["components"] = v.Components

		credentials := []interface{}{}
		for _, item := range v.Credentials {
			credentials = append(credentials, ConfigAssociationDetailsToMap(item))
		}
		result["credentials"] = credentials

		patchTypes := []interface{}{}
		for _, item := range v.PatchTypes {
			patchTypes = append(patchTypes, ConfigAssociationDetailsToMap(item))
		}
		result["patch_types"] = patchTypes

		result["versions"] = v.Versions
	case oci_fleet_apps_management.ProductStackGenericSubCategoryDetails:
		result["sub_category"] = "PRODUCT_STACK_GENERIC"
	default:
		log.Printf("[WARN] Received 'sub_category' of unknown type %v", *obj)
		return nil
	}

	return result
}
