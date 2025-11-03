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

func FleetAppsManagementFleetResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createFleetAppsManagementFleet,
		Read:     readFleetAppsManagementFleet,
		Update:   updateFleetAppsManagementFleet,
		Delete:   deleteFleetAppsManagementFleet,
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
			"resource_selection": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"resource_selection_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"DYNAMIC",
								"MANUAL",
							}, true),
						},

						// Optional
						"rule_selection_criteria": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"match_condition": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"rules": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"basis": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"compartment_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"compartment_id_in_subtree": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
												},
												"conditions": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional
															"attr_group": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"attr_key": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"attr_value": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},

															// Computed
														},
													},
												},
												"match_condition": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"resource_compartment_id": {
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

						// Computed
					},
				},
			},

			// Optional
			"credentials": {
				Type:             schema.TypeList,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool { return true },
				Elem: &schema.Resource{
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
							ForceNew: true,
						},
						"entity_specifics": {
							Type:     schema.TypeList,
							Required: true,
							ForceNew: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"credential_level": {
										Type:             schema.TypeString,
										Required:         true,
										ForceNew:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"FLEET",
											"RESOURCE",
											"TARGET",
										}, true),
									},

									// Optional
									"resource_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"target": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"variables": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										ForceNew: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"name": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"value": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},

												// Computed
											},
										},
									},

									// Computed
								},
							},
						},
						"password": {
							Type:     schema.TypeList,
							Required: true,
							ForceNew: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"credential_type": {
										Type:             schema.TypeString,
										Required:         true,
										ForceNew:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"KEY_ENCRYPTION",
											"PLAIN_TEXT",
											"VAULT_SECRET",
										}, true),
									},

									// Optional
									"key_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"key_version": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"secret_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"secret_version": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"value": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"vault_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},

									// Computed
								},
							},
						},
						"user": {
							Type:     schema.TypeList,
							Required: true,
							ForceNew: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"credential_type": {
										Type:             schema.TypeString,
										Required:         true,
										ForceNew:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"KEY_ENCRYPTION",
											"PLAIN_TEXT",
											"VAULT_SECRET",
										}, true),
									},

									// Optional
									"key_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"key_version": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"secret_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"secret_version": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"value": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"vault_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},

									// Computed
								},
							},
						},

						// Optional

						// Computed
					},
				},
			},
			"defined_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				// DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool { return true },
				Elem:             schema.TypeString,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"details": {
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
						"fleet_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"environment_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				//ForceNew: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"is_target_auto_confirm": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"notification_preferences": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"compartment_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"topic_id": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"preferences": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"on_job_canceled": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"on_job_failure": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"on_job_schedule_change": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"on_job_start": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"on_job_success": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"on_resource_non_compliance": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"on_runbook_newer_version": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"on_task_failure": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"on_task_pause": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"on_task_success": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"on_topology_modification": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"upcoming_schedule": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"notify_before": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"on_upcoming_schedule": {
													Type:     schema.TypeBool,
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

						// Computed
					},
				},
			},
			"parent_fleet_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"products": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				//ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"properties": {
				Type:             schema.TypeList,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool { return true },
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"compartment_id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"fleet_property_type": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"display_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"is_required": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"value": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"resources": {
				Type:             schema.TypeList,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool { return true },
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"compartment_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"resource_id": {
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
						"fleet_resource_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
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
		},
	}
}

func createFleetAppsManagementFleet(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementFleetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementClient()
	sync.WorkRequestClient = m.(*client.OracleClients).FleetAppsManagementFleetAppsManagementWorkRequestClient()

	return tfresource.CreateResource(d, sync)
}

func readFleetAppsManagementFleet(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementFleetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementClient()

	return tfresource.ReadResource(sync)
}

func updateFleetAppsManagementFleet(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementFleetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementClient()
	sync.WorkRequestClient = m.(*client.OracleClients).FleetAppsManagementFleetAppsManagementWorkRequestClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteFleetAppsManagementFleet(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementFleetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).FleetAppsManagementFleetAppsManagementWorkRequestClient()

	return tfresource.DeleteResource(d, sync)
}

type FleetAppsManagementFleetResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_fleet_apps_management.FleetAppsManagementClient
	Res                    *oci_fleet_apps_management.Fleet
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_fleet_apps_management.FleetAppsManagementWorkRequestClient
}

func (s *FleetAppsManagementFleetResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *FleetAppsManagementFleetResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_fleet_apps_management.FleetLifecycleStateCreating),
	}
}

func (s *FleetAppsManagementFleetResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_fleet_apps_management.FleetLifecycleStateActive),
		string(oci_fleet_apps_management.FleetLifecycleStateNeedsAttention),
	}
}

func (s *FleetAppsManagementFleetResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_fleet_apps_management.FleetLifecycleStateDeleting),
	}
}

func (s *FleetAppsManagementFleetResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_fleet_apps_management.FleetLifecycleStateDeleted),
	}
}

func (s *FleetAppsManagementFleetResourceCrud) Create() error {
	request := oci_fleet_apps_management.CreateFleetRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if credentials, ok := s.D.GetOkExists("credentials"); ok {
		interfaces := credentials.([]interface{})
		tmp := make([]oci_fleet_apps_management.AssociatedFleetCredentialDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "credentials", stateDataIndex)
			converted, err := s.mapToAssociatedFleetCredentialDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("credentials") {
			request.Credentials = tmp
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

	if details, ok := s.D.GetOkExists("details"); ok {
		if tmpList := details.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "details", 0)
			tmp, err := s.mapToFleetDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Details = tmp
		}
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if environmentType, ok := s.D.GetOkExists("environment_type"); ok {
		tmp := environmentType.(string)
		request.EnvironmentType = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isTargetAutoConfirm, ok := s.D.GetOkExists("is_target_auto_confirm"); ok {
		tmp := isTargetAutoConfirm.(bool)
		request.IsTargetAutoConfirm = &tmp
	}

	if notificationPreferences, ok := s.D.GetOkExists("notification_preferences"); ok {
		interfaces := notificationPreferences.([]interface{})
		tmp := make([]oci_fleet_apps_management.NotificationPreference, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "notification_preferences", stateDataIndex)
			converted, err := s.mapToNotificationPreference(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("notification_preferences") {
			request.NotificationPreferences = tmp
		}
	}

	if parentFleetId, ok := s.D.GetOkExists("parent_fleet_id"); ok {
		tmp := parentFleetId.(string)
		request.ParentFleetId = &tmp
	}

	if products, ok := s.D.GetOkExists("products"); ok {
		interfaces := products.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("products") {
			request.Products = tmp
		}
	}

	if properties, ok := s.D.GetOkExists("properties"); ok {
		interfaces := properties.([]interface{})
		tmp := make([]oci_fleet_apps_management.AssociatedFleetPropertyDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "properties", stateDataIndex)
			converted, err := s.mapToAssociatedFleetPropertyDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("properties") {
			request.Properties = tmp
		}
	}

	if resourceSelection, ok := s.D.GetOkExists("resource_selection"); ok {
		if tmpList := resourceSelection.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "resource_selection", 0)
			tmp, err := s.mapToResourceSelection(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ResourceSelection = tmp
		}
	}

	if resources, ok := s.D.GetOkExists("resources"); ok {
		interfaces := resources.([]interface{})
		tmp := make([]oci_fleet_apps_management.AssociatedFleetResourceDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "resources", stateDataIndex)
			converted, err := s.mapToAssociatedFleetResourceDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("resources") {
			request.Resources = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	response, err := s.Client.CreateFleet(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getFleetFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management"), oci_fleet_apps_management.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *FleetAppsManagementFleetResourceCrud) getFleetFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_fleet_apps_management.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	fleetId, err := fleetWaitForWorkRequest(workId, "fleet",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.WorkRequestClient)

	if err != nil {
		return err
	}
	s.D.SetId(*fleetId)

	return s.Get()
}

func fleetWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func fleetWaitForWorkRequest(wId *string, entityType string, action oci_fleet_apps_management.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_fleet_apps_management.FleetAppsManagementWorkRequestClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "fleet_apps_management")
	retryPolicy.ShouldRetryOperation = fleetWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromFleetAppsManagementFleetWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromFleetAppsManagementFleetWorkRequest(client *oci_fleet_apps_management.FleetAppsManagementWorkRequestClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_fleet_apps_management.ActionTypeEnum) error {
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

func (s *FleetAppsManagementFleetResourceCrud) Get() error {
	request := oci_fleet_apps_management.GetFleetRequest{}

	tmp := s.D.Id()
	request.FleetId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	response, err := s.Client.GetFleet(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Fleet
	return nil
}

func (s *FleetAppsManagementFleetResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_fleet_apps_management.UpdateFleetRequest{}

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

	if environmentType, ok := s.D.GetOkExists("environment_type"); ok {
		tmp := environmentType.(string)
		request.EnvironmentType = &tmp
	}

	tmp := s.D.Id()
	request.FleetId = &tmp

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isTargetAutoConfirm, ok := s.D.GetOkExists("is_target_auto_confirm"); ok {
		tmp := isTargetAutoConfirm.(bool)
		request.IsTargetAutoConfirm = &tmp
	}

	if notificationPreferences, ok := s.D.GetOkExists("notification_preferences"); ok {
		interfaces := notificationPreferences.([]interface{})
		tmp := make([]oci_fleet_apps_management.NotificationPreference, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "notification_preferences", stateDataIndex)
			converted, err := s.mapToNotificationPreference(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("notification_preferences") {
			request.NotificationPreferences = tmp
		}
	}

	if products, ok := s.D.GetOkExists("products"); ok {
		interfaces := products.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("products") {
			request.Products = tmp
		}
	}

	if resourceSelection, ok := s.D.GetOkExists("resource_selection"); ok {
		if tmpList := resourceSelection.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "resource_selection", 0)
			tmp, err := s.mapToResourceSelection(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ResourceSelection = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	response, err := s.Client.UpdateFleet(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Fleet
	return nil
}

func (s *FleetAppsManagementFleetResourceCrud) Delete() error {
	request := oci_fleet_apps_management.DeleteFleetRequest{}

	tmp := s.D.Id()
	request.FleetId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	response, err := s.Client.DeleteFleet(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := fleetWaitForWorkRequest(workId, "fleet",
		oci_fleet_apps_management.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.WorkRequestClient)
	return delWorkRequestErr
}

func (s *FleetAppsManagementFleetResourceCrud) SetData() error {

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	} else {
		s.D.Set("compartment_id", nil)
	}

	credentials := []interface{}{}
	for _, item := range s.Res.Credentials {
		credentials = append(credentials, AssociatedFleetCredentialDetailsToMap(item))
	}
	s.D.Set("credentials", credentials)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.Details != nil {
		detailsArray := []interface{}{}
		if detailsMap := FleetDetailsToMap(&s.Res.Details); detailsMap != nil {
			detailsArray = append(detailsArray, detailsMap)
		}
		s.D.Set("details", detailsArray)
	} else {
		s.D.Set("details", nil)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.EnvironmentType != nil {
		s.D.Set("environment_type", *s.Res.EnvironmentType)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsTargetAutoConfirm != nil {
		s.D.Set("is_target_auto_confirm", *s.Res.IsTargetAutoConfirm)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	notificationPreferences := []interface{}{}
	for _, item := range s.Res.NotificationPreferences {
		notificationPreferences = append(notificationPreferences, NotificationPreferenceToMap(item))
	}
	s.D.Set("notification_preferences", notificationPreferences)

	if s.Res.ParentFleetId != nil {
		s.D.Set("parent_fleet_id", *s.Res.ParentFleetId)
	}

	s.D.Set("products", s.Res.Products)

	properties := []interface{}{}
	for _, item := range s.Res.Properties {
		properties = append(properties, AssociatedFleetPropertyDetailsToMap(item))
	}
	s.D.Set("properties", properties)

	if s.Res.ResourceRegion != nil {
		s.D.Set("resource_region", *s.Res.ResourceRegion)
	}

	if s.Res.ResourceSelection != nil {
		resourceSelectionArray := []interface{}{}
		if resourceSelectionMap := ResourceSelectionToMap(&s.Res.ResourceSelection); resourceSelectionMap != nil {
			resourceSelectionArray = append(resourceSelectionArray, resourceSelectionMap)
		}
		s.D.Set("resource_selection", resourceSelectionArray)
	} else {
		s.D.Set("resource_selection", nil)
	}

	resources := []interface{}{}
	for _, item := range s.Res.Resources {
		resources = append(resources, AssociatedFleetResourceDetailsToMap(item))
	}
	s.D.Set("resources", resources)

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

func (s *FleetAppsManagementFleetResourceCrud) mapToAssociatedFleetCredentialDetails(fieldKeyFormat string) (oci_fleet_apps_management.AssociatedFleetCredentialDetails, error) {
	result := oci_fleet_apps_management.AssociatedFleetCredentialDetails{}

	if compartmentId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compartment_id")); ok {
		tmp := compartmentId.(string)
		result.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	if entitySpecifics, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "entity_specifics")); ok {
		if tmpList := entitySpecifics.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "entity_specifics"), 0)
			tmp, err := s.mapToCredentialEntitySpecificDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert entity_specifics, encountered error: %v", err)
			}
			result.EntitySpecifics = tmp
		}
	}

	if password, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "password")); ok {
		if tmpList := password.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "password"), 0)
			tmp, err := s.mapToCredentialDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert password, encountered error: %v", err)
			}
			result.Password = tmp
		}
	}

	if user, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "user")); ok {
		if tmpList := user.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "user"), 0)
			tmp, err := s.mapToCredentialDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert user, encountered error: %v", err)
			}
			result.User = tmp
		}
	}

	return result, nil
}

func AssociatedFleetCredentialDetailsToMap(obj oci_fleet_apps_management.AssociatedFleetCredentialDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.EntitySpecifics != nil {
		entitySpecificsArray := []interface{}{}
		if entitySpecificsMap := CredentialEntitySpecificDetailsToMap(&obj.EntitySpecifics); entitySpecificsMap != nil {
			entitySpecificsArray = append(entitySpecificsArray, entitySpecificsMap)
		}
		result["entity_specifics"] = entitySpecificsArray
	}

	if obj.Password != nil {
		passwordArray := []interface{}{}
		if passwordMap := CredentialDetailsToMap(&obj.Password); passwordMap != nil {
			passwordArray = append(passwordArray, passwordMap)
		}
		result["password"] = passwordArray
	}

	if obj.User != nil {
		userArray := []interface{}{}
		if userMap := CredentialDetailsToMap(&obj.User); userMap != nil {
			userArray = append(userArray, userMap)
		}
		result["user"] = userArray
	}

	return result
}

func (s *FleetAppsManagementFleetResourceCrud) mapToAssociatedFleetPropertyDetails(fieldKeyFormat string) (oci_fleet_apps_management.AssociatedFleetPropertyDetails, error) {
	result := oci_fleet_apps_management.AssociatedFleetPropertyDetails{}

	if compartmentId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compartment_id")); ok {
		tmp := compartmentId.(string)
		result.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	if fleetPropertyType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "fleet_property_type")); ok {
		result.FleetPropertyType = oci_fleet_apps_management.AssociatedFleetPropertyDetailsFleetPropertyTypeEnum(fleetPropertyType.(string))
	}

	if isRequired, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_required")); ok {
		tmp := isRequired.(bool)
		result.IsRequired = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func AssociatedFleetPropertyDetailsToMap(obj oci_fleet_apps_management.AssociatedFleetPropertyDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["fleet_property_type"] = string(obj.FleetPropertyType)

	if obj.IsRequired != nil {
		result["is_required"] = bool(*obj.IsRequired)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *FleetAppsManagementFleetResourceCrud) mapToAssociatedFleetResourceDetails(fieldKeyFormat string) (oci_fleet_apps_management.AssociatedFleetResourceDetails, error) {
	result := oci_fleet_apps_management.AssociatedFleetResourceDetails{}

	if compartmentId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compartment_id")); ok {
		tmp := compartmentId.(string)
		result.CompartmentId = &tmp
	}

	if fleetResourceType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "fleet_resource_type")); ok {
		tmp := fleetResourceType.(string)
		result.FleetResourceType = &tmp
	}

	if resourceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "resource_id")); ok {
		tmp := resourceId.(string)
		result.ResourceId = &tmp
	}

	if tenancyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tenancy_id")); ok {
		tmp := tenancyId.(string)
		result.TenancyId = &tmp
	}

	return result, nil
}

func AssociatedFleetResourceDetailsToMap(obj oci_fleet_apps_management.AssociatedFleetResourceDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.FleetResourceType != nil {
		result["fleet_resource_type"] = string(*obj.FleetResourceType)
	}

	if obj.ResourceId != nil {
		result["resource_id"] = string(*obj.ResourceId)
	}

	if obj.TenancyId != nil {
		result["tenancy_id"] = string(*obj.TenancyId)
	}

	return result
}

func (s *FleetAppsManagementFleetResourceCrud) mapToCondition(fieldKeyFormat string) (oci_fleet_apps_management.Condition, error) {
	result := oci_fleet_apps_management.Condition{}

	if attrGroup, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "attr_group")); ok {
		tmp := attrGroup.(string)
		result.AttrGroup = &tmp
	}

	if attrKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "attr_key")); ok {
		tmp := attrKey.(string)
		result.AttrKey = &tmp
	}

	if attrValue, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "attr_value")); ok {
		tmp := attrValue.(string)
		result.AttrValue = &tmp
	}

	return result, nil
}

func ConditionToMap(obj oci_fleet_apps_management.Condition) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AttrGroup != nil {
		result["attr_group"] = string(*obj.AttrGroup)
	}

	if obj.AttrKey != nil {
		result["attr_key"] = string(*obj.AttrKey)
	}

	if obj.AttrValue != nil {
		result["attr_value"] = string(*obj.AttrValue)
	}

	return result
}

func (s *FleetAppsManagementFleetResourceCrud) mapToCredentialDetails(fieldKeyFormat string) (oci_fleet_apps_management.CredentialDetails, error) {
	var baseObject oci_fleet_apps_management.CredentialDetails
	//discriminator
	credentialTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "credential_type"))
	var credentialType string
	if ok {
		credentialType = credentialTypeRaw.(string)
	} else {
		credentialType = "" // default value
	}
	switch strings.ToLower(credentialType) {
	case strings.ToLower("KEY_ENCRYPTION"):
		details := oci_fleet_apps_management.KeyEncryptionCredentialDetails{}
		if keyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_id")); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if keyVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_version")); ok {
			tmp := keyVersion.(string)
			details.KeyVersion = &tmp
		}
		if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
			tmp := value.(string)
			details.Value = &tmp
		}
		if vaultId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vault_id")); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		baseObject = details
	case strings.ToLower("PLAIN_TEXT"):
		details := oci_fleet_apps_management.PlainTextCredentialDetails{}
		if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
			tmp := value.(string)
			details.Value = &tmp
		}
		baseObject = details
	case strings.ToLower("VAULT_SECRET"):
		details := oci_fleet_apps_management.VaultSecretCredentialDetails{}
		if secretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "secret_id")); ok {
			tmp := secretId.(string)
			details.SecretId = &tmp
		}
		if secretVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "secret_version")); ok {
			tmp := secretVersion.(string)
			details.SecretVersion = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown credential_type '%v' was specified", credentialType)
	}
	return baseObject, nil
}

func (s *FleetAppsManagementFleetResourceCrud) mapToCredentialEntitySpecificDetails(fieldKeyFormat string) (oci_fleet_apps_management.CredentialEntitySpecificDetails, error) {
	var baseObject oci_fleet_apps_management.CredentialEntitySpecificDetails
	//discriminator
	credentialLevelRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "credential_level"))
	var credentialLevel string
	if ok {
		credentialLevel = credentialLevelRaw.(string)
	} else {
		credentialLevel = "" // default value
	}
	switch strings.ToLower(credentialLevel) {
	case strings.ToLower("FLEET"):
		details := oci_fleet_apps_management.FleetCredentialEntitySpecificDetails{}
		if variables, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "variables")); ok {
			interfaces := variables.([]interface{})
			tmp := make([]oci_fleet_apps_management.Variable, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "variables"), stateDataIndex)
				converted, err := s.mapToVariable(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "variables")) {
				details.Variables = tmp
			}
		}
		baseObject = details
	case strings.ToLower("RESOURCE"):
		details := oci_fleet_apps_management.ResourceCredentialEntitySpecificDetails{}
		if resourceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "resource_id")); ok {
			tmp := resourceId.(string)
			details.ResourceId = &tmp
		}
		baseObject = details
	case strings.ToLower("TARGET"):
		details := oci_fleet_apps_management.TargetCredentialEntitySpecificDetails{}
		if resourceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "resource_id")); ok {
			tmp := resourceId.(string)
			details.ResourceId = &tmp
		}
		if target, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "target")); ok {
			tmp := target.(string)
			details.Target = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown credential_level '%v' was specified", credentialLevel)
	}
	return baseObject, nil
}

func (s *FleetAppsManagementFleetResourceCrud) mapToFleetDetails(fieldKeyFormat string) (oci_fleet_apps_management.FleetDetails, error) {
	var baseObject oci_fleet_apps_management.FleetDetails
	//discriminator
	fleetTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "fleet_type"))
	var fleetType string
	if ok {
		fleetType = fleetTypeRaw.(string)
	} else {
		fleetType = "" // default value
	}
	switch strings.ToLower(fleetType) {
	case strings.ToLower("ENVIRONMENT"):
		details := oci_fleet_apps_management.EnvironmentFleetDetails{}
		baseObject = details
	case strings.ToLower("GENERIC"):
		details := oci_fleet_apps_management.GenericFleetDetails{}
		baseObject = details
	case strings.ToLower("GROUP"):
		details := oci_fleet_apps_management.GroupFleetDetails{}
		baseObject = details
	case strings.ToLower("PRODUCT"):
		details := oci_fleet_apps_management.ProductFleetDetails{}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown fleet_type '%v' was specified", fleetType)
	}
	return baseObject, nil
}

func FleetDetailsToMap(obj *oci_fleet_apps_management.FleetDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch (*obj).(type) {
	case oci_fleet_apps_management.EnvironmentFleetDetails:
		result["fleet_type"] = "ENVIRONMENT"
	case oci_fleet_apps_management.GenericFleetDetails:
		result["fleet_type"] = "GENERIC"
	case oci_fleet_apps_management.GroupFleetDetails:
		result["fleet_type"] = "GROUP"
	case oci_fleet_apps_management.ProductFleetDetails:
		result["fleet_type"] = "PRODUCT"
	default:
		log.Printf("[WARN] Received 'fleet_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func FleetSummaryToMap(obj oci_fleet_apps_management.FleetSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.Details != nil {
		detailsArray := []interface{}{}
		if detailsMap := FleetDetailsToMap(&obj.Details); detailsMap != nil {
			detailsArray = append(detailsArray, detailsMap)
		}
		result["details"] = detailsArray
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.EnvironmentType != nil {
		result["environment_type"] = string(*obj.EnvironmentType)
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

	return result
}

func (s *FleetAppsManagementFleetResourceCrud) mapToNotificationPreference(fieldKeyFormat string) (oci_fleet_apps_management.NotificationPreference, error) {
	result := oci_fleet_apps_management.NotificationPreference{}

	if compartmentId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compartment_id")); ok {
		tmp := compartmentId.(string)
		result.CompartmentId = &tmp
	}

	if preferences, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "preferences")); ok {
		if tmpList := preferences.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "preferences"), 0)
			tmp, err := s.mapToPreferences(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert preferences, encountered error: %v", err)
			}
			result.Preferences = &tmp
		}
	}

	if topicId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "topic_id")); ok {
		tmp := topicId.(string)
		result.TopicId = &tmp
	}

	return result, nil
}

func NotificationPreferenceToMap(obj oci_fleet_apps_management.NotificationPreference) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.Preferences != nil {
		result["preferences"] = []interface{}{PreferencesToMap(obj.Preferences)}
	}

	if obj.TopicId != nil {
		result["topic_id"] = string(*obj.TopicId)
	}

	return result
}

func (s *FleetAppsManagementFleetResourceCrud) mapToPreferences(fieldKeyFormat string) (oci_fleet_apps_management.Preferences, error) {
	result := oci_fleet_apps_management.Preferences{}

	if onJobCanceled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "on_job_canceled")); ok {
		tmp := onJobCanceled.(bool)
		result.OnJobCanceled = &tmp
	}

	if onJobFailure, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "on_job_failure")); ok {
		tmp := onJobFailure.(bool)
		result.OnJobFailure = &tmp
	}

	if onJobScheduleChange, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "on_job_schedule_change")); ok {
		tmp := onJobScheduleChange.(bool)
		result.OnJobScheduleChange = &tmp
	}

	if onJobStart, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "on_job_start")); ok {
		tmp := onJobStart.(bool)
		result.OnJobStart = &tmp
	}

	if onJobSuccess, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "on_job_success")); ok {
		tmp := onJobSuccess.(bool)
		result.OnJobSuccess = &tmp
	}

	if onResourceNonCompliance, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "on_resource_non_compliance")); ok {
		tmp := onResourceNonCompliance.(bool)
		result.OnResourceNonCompliance = &tmp
	}

	if onRunbookNewerVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "on_runbook_newer_version")); ok {
		tmp := onRunbookNewerVersion.(bool)
		result.OnRunbookNewerVersion = &tmp
	}

	if onTaskFailure, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "on_task_failure")); ok {
		tmp := onTaskFailure.(bool)
		result.OnTaskFailure = &tmp
	}

	if onTaskPause, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "on_task_pause")); ok {
		tmp := onTaskPause.(bool)
		result.OnTaskPause = &tmp
	}

	if onTaskSuccess, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "on_task_success")); ok {
		tmp := onTaskSuccess.(bool)
		result.OnTaskSuccess = &tmp
	}

	if onTopologyModification, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "on_topology_modification")); ok {
		tmp := onTopologyModification.(bool)
		result.OnTopologyModification = &tmp
	}

	if upcomingSchedule, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "upcoming_schedule")); ok {
		if tmpList := upcomingSchedule.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "upcoming_schedule"), 0)
			tmp, err := s.mapToUpcomingSchedule(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert upcoming_schedule, encountered error: %v", err)
			}
			result.UpcomingSchedule = &tmp
		}
	}

	return result, nil
}

func PreferencesToMap(obj *oci_fleet_apps_management.Preferences) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.OnJobCanceled != nil {
		result["on_job_canceled"] = bool(*obj.OnJobCanceled)
	}

	if obj.OnJobFailure != nil {
		result["on_job_failure"] = bool(*obj.OnJobFailure)
	}

	if obj.OnJobScheduleChange != nil {
		result["on_job_schedule_change"] = bool(*obj.OnJobScheduleChange)
	}

	if obj.OnJobStart != nil {
		result["on_job_start"] = bool(*obj.OnJobStart)
	}

	if obj.OnJobSuccess != nil {
		result["on_job_success"] = bool(*obj.OnJobSuccess)
	}

	if obj.OnResourceNonCompliance != nil {
		result["on_resource_non_compliance"] = bool(*obj.OnResourceNonCompliance)
	}

	if obj.OnRunbookNewerVersion != nil {
		result["on_runbook_newer_version"] = bool(*obj.OnRunbookNewerVersion)
	}

	if obj.OnTaskFailure != nil {
		result["on_task_failure"] = bool(*obj.OnTaskFailure)
	}

	if obj.OnTaskPause != nil {
		result["on_task_pause"] = bool(*obj.OnTaskPause)
	}

	if obj.OnTaskSuccess != nil {
		result["on_task_success"] = bool(*obj.OnTaskSuccess)
	}

	if obj.OnTopologyModification != nil {
		result["on_topology_modification"] = bool(*obj.OnTopologyModification)
	}

	if obj.UpcomingSchedule != nil {
		result["upcoming_schedule"] = []interface{}{UpcomingScheduleToMap(obj.UpcomingSchedule)}
	}

	return result
}

func (s *FleetAppsManagementFleetResourceCrud) mapToResourceSelection(fieldKeyFormat string) (oci_fleet_apps_management.ResourceSelection, error) {
	var baseObject oci_fleet_apps_management.ResourceSelection
	//discriminator
	resourceSelectionTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "resource_selection_type"))
	var resourceSelectionType string
	if ok {
		resourceSelectionType = resourceSelectionTypeRaw.(string)
	} else {
		resourceSelectionType = "" // default value
	}
	switch strings.ToLower(resourceSelectionType) {
	case strings.ToLower("DYNAMIC"):
		details := oci_fleet_apps_management.DynamicResourceSelection{}
		if ruleSelectionCriteria, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "rule_selection_criteria")); ok {
			if tmpList := ruleSelectionCriteria.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "rule_selection_criteria"), 0)
				tmp, err := s.mapToSelectionCriteria(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert rule_selection_criteria, encountered error: %v", err)
				}
				details.RuleSelectionCriteria = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("MANUAL"):
		details := oci_fleet_apps_management.ManualResourceSelection{}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown resource_selection_type '%v' was specified", resourceSelectionType)
	}
	return baseObject, nil
}

func ResourceSelectionToMap(obj *oci_fleet_apps_management.ResourceSelection) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_fleet_apps_management.DynamicResourceSelection:
		result["resource_selection_type"] = "DYNAMIC"

		if v.RuleSelectionCriteria != nil {
			result["rule_selection_criteria"] = []interface{}{SelectionCriteriaToMap(v.RuleSelectionCriteria)}
		}
	case oci_fleet_apps_management.ManualResourceSelection:
		result["resource_selection_type"] = "MANUAL"
	default:
		log.Printf("[WARN] Received 'resource_selection_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *FleetAppsManagementFleetResourceCrud) mapToRule(fieldKeyFormat string) (oci_fleet_apps_management.Rule, error) {
	result := oci_fleet_apps_management.Rule{}

	if basis, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "basis")); ok {
		tmp := basis.(string)
		result.Basis = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compartment_id")); ok {
		tmp := compartmentId.(string)
		result.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compartment_id_in_subtree")); ok {
		tmp := compartmentIdInSubtree.(bool)
		result.CompartmentIdInSubtree = &tmp
	}

	if conditions, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "conditions")); ok {
		interfaces := conditions.([]interface{})
		tmp := make([]oci_fleet_apps_management.Condition, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "conditions"), stateDataIndex)
			converted, err := s.mapToCondition(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "conditions")) {
			result.Conditions = tmp
		}
	}

	if matchCondition, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "match_condition")); ok {
		result.MatchCondition = oci_fleet_apps_management.RuleMatchConditionEnum(matchCondition.(string))
	}

	if resourceCompartmentId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "resource_compartment_id")); ok {
		tmp := resourceCompartmentId.(string)
		result.ResourceCompartmentId = &tmp
	}

	return result, nil
}

func RuleToMap(obj oci_fleet_apps_management.Rule) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Basis != nil {
		result["basis"] = string(*obj.Basis)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.CompartmentIdInSubtree != nil {
		result["compartment_id_in_subtree"] = bool(*obj.CompartmentIdInSubtree)
	}

	conditions := []interface{}{}
	for _, item := range obj.Conditions {
		conditions = append(conditions, ConditionToMap(item))
	}
	result["conditions"] = conditions

	result["match_condition"] = string(obj.MatchCondition)

	if obj.ResourceCompartmentId != nil {
		result["resource_compartment_id"] = string(*obj.ResourceCompartmentId)
	}

	return result
}

func (s *FleetAppsManagementFleetResourceCrud) mapToSelectionCriteria(fieldKeyFormat string) (oci_fleet_apps_management.SelectionCriteria, error) {
	result := oci_fleet_apps_management.SelectionCriteria{}

	if matchCondition, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "match_condition")); ok {
		result.MatchCondition = oci_fleet_apps_management.SelectionCriteriaMatchConditionEnum(matchCondition.(string))
	}

	if rules, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "rules")); ok {
		interfaces := rules.([]interface{})
		tmp := make([]oci_fleet_apps_management.Rule, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "rules"), stateDataIndex)
			converted, err := s.mapToRule(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "rules")) {
			result.Rules = tmp
		}
	}

	return result, nil
}

func SelectionCriteriaToMap(obj *oci_fleet_apps_management.SelectionCriteria) map[string]interface{} {
	result := map[string]interface{}{}

	result["match_condition"] = string(obj.MatchCondition)

	rules := []interface{}{}
	for _, item := range obj.Rules {
		rules = append(rules, RuleToMap(item))
	}
	result["rules"] = rules

	return result
}

func (s *FleetAppsManagementFleetResourceCrud) mapToUpcomingSchedule(fieldKeyFormat string) (oci_fleet_apps_management.UpcomingSchedule, error) {
	result := oci_fleet_apps_management.UpcomingSchedule{}

	if notifyBefore, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "notify_before")); ok {
		tmp := notifyBefore.(string)
		result.NotifyBefore = &tmp
	}

	if onUpcomingSchedule, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "on_upcoming_schedule")); ok {
		tmp := onUpcomingSchedule.(bool)
		result.OnUpcomingSchedule = &tmp
	}

	return result, nil
}

func UpcomingScheduleToMap(obj *oci_fleet_apps_management.UpcomingSchedule) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.NotifyBefore != nil {
		result["notify_before"] = string(*obj.NotifyBefore)
	}

	if obj.OnUpcomingSchedule != nil {
		result["on_upcoming_schedule"] = bool(*obj.OnUpcomingSchedule)
	}

	return result
}

func (s *FleetAppsManagementFleetResourceCrud) mapToVariable(fieldKeyFormat string) (oci_fleet_apps_management.Variable, error) {
	result := oci_fleet_apps_management.Variable{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func (s *FleetAppsManagementFleetResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_fleet_apps_management.ChangeFleetCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.FleetId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	response, err := s.Client.ChangeFleetCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getFleetFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management"), oci_fleet_apps_management.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

// func suppressIfMissingInState(k, old, new string, d *schema.ResourceData) bool {
// 	return new == "" && old != ""
// }

func suppressIfMissingInState(k, old, new string, d *schema.ResourceData) bool {
	// Normalize whitespace
	old = strings.TrimSpace(old)
	new = strings.TrimSpace(new)

	// Suppress when new is effectively empty or null but old had a value
	if new == "" || new == "[]" || new == "<nil>" || new == "null" || new == "{}" {
		if old != "" && old != "[]" && old != "{}" && old != "null" && old != "<nil>" {
			return true
		}
	}
	return false
}
