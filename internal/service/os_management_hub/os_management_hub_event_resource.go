// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package os_management_hub

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_os_management_hub "github.com/oracle/oci-go-sdk/v65/osmanagementhub"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OsManagementHubEventResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createOsManagementHubEvent,
		Read:     readOsManagementHubEvent,
		Update:   updateOsManagementHubEvent,
		Delete:   deleteOsManagementHubEvent,
		Schema: map[string]*schema.Schema{
			// Required
			"event_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
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

			// Computed
			"data": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"additional_details": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"exploit_cves": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"initiator_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"vmcore": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"backtrace": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"component": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"work_request_ids": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"content": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"content_availability": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"content_location": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"exploit_detection_log_content": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"exploit_object_store_location": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"size": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"type": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"event_fingerprint": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"event_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"operation_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"reason": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_first_occurred": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"event_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"event_summary": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_managed_by_autonomous_linux": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"resource_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_details": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"architecture": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ksplice_effective_kernel_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"os_family": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"os_kernel_release": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"os_kernel_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"os_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"os_system_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
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
			"time_occurred": {
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

func createOsManagementHubEvent(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubEventResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OsmhEventClient()
	sync.WorkRequestClient = m.(*client.OracleClients).OsManagementHubWorkRequestClient()
	return tfresource.CreateResource(d, sync)
}

func readOsManagementHubEvent(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubEventResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OsmhEventClient()

	return tfresource.ReadResource(sync)
}

func updateOsManagementHubEvent(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubEventResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OsmhEventClient()
	sync.WorkRequestClient = m.(*client.OracleClients).OsManagementHubWorkRequestClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteOsManagementHubEvent(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubEventResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OsmhEventClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).OsManagementHubWorkRequestClient()

	return tfresource.DeleteResource(d, sync)
}

type OsManagementHubEventResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_os_management_hub.EventClient
	Res                    *oci_os_management_hub.Event
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_os_management_hub.WorkRequestClient
}

func (s *OsManagementHubEventResourceCrud) ID() string {
	event := *s.Res
	return *event.GetId()
}

func (s *OsManagementHubEventResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_os_management_hub.EventLifecycleStateCreating),
	}
}

func (s *OsManagementHubEventResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_os_management_hub.EventLifecycleStateActive),
	}
}

func (s *OsManagementHubEventResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_os_management_hub.EventLifecycleStateDeleting),
	}
}

func (s *OsManagementHubEventResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_os_management_hub.EventLifecycleStateDeleted),
	}
}

func (s *OsManagementHubEventResourceCrud) Create() error {
	request := oci_os_management_hub.UpdateEventRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if eventId, ok := s.D.GetOkExists("event_id"); ok {
		tmp := eventId.(string)
		request.EventId = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub")

	response, err := s.Client.UpdateEvent(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Event
	return nil
}

func (s *OsManagementHubEventResourceCrud) getEventFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_os_management_hub.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	eventId, err := eventWaitForWorkRequest(workId, "event",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.WorkRequestClient)

	if err != nil {
		return err
	}
	s.D.SetId(*eventId)

	return s.Get()
}

func eventWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "os_management_hub", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_os_management_hub.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func eventWaitForWorkRequest(wId *string, entityType string, action oci_os_management_hub.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_os_management_hub.WorkRequestClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "os_management_hub")
	retryPolicy.ShouldRetryOperation = eventWorkRequestShouldRetryFunc(timeout)

	response := oci_os_management_hub.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_os_management_hub.OperationStatusInProgress),
			string(oci_os_management_hub.OperationStatusAccepted),
			string(oci_os_management_hub.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_os_management_hub.OperationStatusSucceeded),
			string(oci_os_management_hub.OperationStatusFailed),
			string(oci_os_management_hub.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_os_management_hub.GetWorkRequestRequest{
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
		if strings.Contains(strings.ToLower(string(res.EntityType)), entityType) {
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_os_management_hub.OperationStatusFailed || response.Status == oci_os_management_hub.OperationStatusCanceled {
		return nil, getErrorFromOsManagementHubEventWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromOsManagementHubEventWorkRequest(client *oci_os_management_hub.WorkRequestClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_os_management_hub.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_os_management_hub.ListWorkRequestErrorsRequest{
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

func (s *OsManagementHubEventResourceCrud) Get() error {
	request := oci_os_management_hub.GetEventRequest{}

	tmp := s.D.Id()
	request.EventId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub")

	response, err := s.Client.GetEvent(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Event
	return nil
}

func (s *OsManagementHubEventResourceCrud) Update() error {

	if _, ok := s.D.GetOkExists("compartmentId"); ok && s.D.HasChange("compartmentId") {
		err := s.ChangeEventCompartment()
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
	request := oci_os_management_hub.UpdateEventRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	//tmp := s.D.Id()
	//request.EventId = &tmp

	if eventId, ok := s.D.GetOkExists("event_id"); ok {
		tmp := eventId.(string)
		request.EventId = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub")

	response, err := s.Client.UpdateEvent(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Event
	return nil
}

func (s *OsManagementHubEventResourceCrud) Delete() error {
	request := oci_os_management_hub.DeleteEventRequest{}

	tmp := s.D.Id()
	request.EventId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub")

	response, err := s.Client.DeleteEvent(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	if workId != nil {
		_, delWorkRequestErr := eventWaitForWorkRequest(workId, "event",
			oci_os_management_hub.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.WorkRequestClient)
		return delWorkRequestErr
	}
	return nil
}

func (s *OsManagementHubEventResourceCrud) SetData() error {
	switch v := (*s.Res).(type) {
	case oci_os_management_hub.AgentEvent:
		s.D.Set("type", "AGENT")

		if v.Data != nil {
			s.D.Set("data", []interface{}{AgentEventDataToMap(v.Data)})
		} else {
			s.D.Set("data", nil)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.EventDetails != nil {
			s.D.Set("event_details", *v.EventDetails)
		}

		if v.EventSummary != nil {
			s.D.Set("event_summary", *v.EventSummary)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.Set("id", *v.Id)
		}

		if v.IsManagedByAutonomousLinux != nil {
			s.D.Set("is_managed_by_autonomous_linux", *v.IsManagedByAutonomousLinux)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ResourceId != nil {
			s.D.Set("resource_id", *v.ResourceId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemDetails != nil {
			s.D.Set("system_details", []interface{}{SystemDetailsToMap(v.SystemDetails)})
		} else {
			s.D.Set("system_details", nil)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeOccurred != nil {
			s.D.Set("time_occurred", v.TimeOccurred.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_os_management_hub.ExploitAttemptEvent:
		s.D.Set("type", "EXPLOIT_ATTEMPT")

		if v.Data != nil {
			s.D.Set("data", []interface{}{ExploitAttemptEventDataToMap(v.Data)})
		} else {
			s.D.Set("data", nil)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.EventDetails != nil {
			s.D.Set("event_details", *v.EventDetails)
		}

		if v.EventSummary != nil {
			s.D.Set("event_summary", *v.EventSummary)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.Set("event_id", *v.Id)
		}

		if v.IsManagedByAutonomousLinux != nil {
			s.D.Set("is_managed_by_autonomous_linux", *v.IsManagedByAutonomousLinux)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ResourceId != nil {
			s.D.Set("resource_id", *v.ResourceId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemDetails != nil {
			s.D.Set("system_details", []interface{}{SystemDetailsToMap(v.SystemDetails)})
		} else {
			s.D.Set("system_details", nil)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeOccurred != nil {
			s.D.Set("time_occurred", v.TimeOccurred.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_os_management_hub.KernelCrashEvent:
		s.D.Set("type", "KERNEL_CRASH")

		if v.Data != nil {
			s.D.Set("data", []interface{}{KernelEventDataToMap(v.Data)})
		} else {
			s.D.Set("data", nil)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.EventDetails != nil {
			s.D.Set("event_details", *v.EventDetails)
		}

		if v.EventSummary != nil {
			s.D.Set("event_summary", *v.EventSummary)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.Set("id", *v.Id)
		}

		if v.IsManagedByAutonomousLinux != nil {
			s.D.Set("is_managed_by_autonomous_linux", *v.IsManagedByAutonomousLinux)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ResourceId != nil {
			s.D.Set("resource_id", *v.ResourceId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemDetails != nil {
			s.D.Set("system_details", []interface{}{SystemDetailsToMap(v.SystemDetails)})
		} else {
			s.D.Set("system_details", nil)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeOccurred != nil {
			s.D.Set("time_occurred", v.TimeOccurred.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_os_management_hub.KernelOopsEvent:
		s.D.Set("type", "KERNEL_OOPS")

		if v.Data != nil {
			s.D.Set("data", []interface{}{KernelEventDataToMap(v.Data)})
		} else {
			s.D.Set("data", nil)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.EventDetails != nil {
			s.D.Set("event_details", *v.EventDetails)
		}

		if v.EventSummary != nil {
			s.D.Set("event_summary", *v.EventSummary)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.Set("id", *v.Id)
		}

		if v.IsManagedByAutonomousLinux != nil {
			s.D.Set("is_managed_by_autonomous_linux", *v.IsManagedByAutonomousLinux)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ResourceId != nil {
			s.D.Set("resource_id", *v.ResourceId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemDetails != nil {
			s.D.Set("system_details", []interface{}{SystemDetailsToMap(v.SystemDetails)})
		} else {
			s.D.Set("system_details", nil)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeOccurred != nil {
			s.D.Set("time_occurred", v.TimeOccurred.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_os_management_hub.KspliceUpdateEvent:
		s.D.Set("type", "KSPLICE_UPDATE")

		if v.Data != nil {
			s.D.Set("data", []interface{}{KspliceUpdateEventDataToMap(v.Data)})
		} else {
			s.D.Set("data", nil)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.EventDetails != nil {
			s.D.Set("event_details", *v.EventDetails)
		}

		if v.EventSummary != nil {
			s.D.Set("event_summary", *v.EventSummary)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.Set("id", *v.Id)
		}

		if v.IsManagedByAutonomousLinux != nil {
			s.D.Set("is_managed_by_autonomous_linux", *v.IsManagedByAutonomousLinux)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ResourceId != nil {
			s.D.Set("resource_id", *v.ResourceId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemDetails != nil {
			s.D.Set("system_details", []interface{}{SystemDetailsToMap(v.SystemDetails)})
		} else {
			s.D.Set("system_details", nil)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeOccurred != nil {
			s.D.Set("time_occurred", v.TimeOccurred.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_os_management_hub.ManagementStationEvent:
		s.D.Set("type", "MANAGEMENT_STATION")

		if v.Data != nil {
			s.D.Set("data", []interface{}{ManagementStationEventDataToMap(v.Data)})
		} else {
			s.D.Set("data", nil)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.EventDetails != nil {
			s.D.Set("event_details", *v.EventDetails)
		}

		if v.EventSummary != nil {
			s.D.Set("event_summary", *v.EventSummary)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.Set("id", *v.Id)
		}

		if v.IsManagedByAutonomousLinux != nil {
			s.D.Set("is_managed_by_autonomous_linux", *v.IsManagedByAutonomousLinux)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ResourceId != nil {
			s.D.Set("resource_id", *v.ResourceId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemDetails != nil {
			s.D.Set("system_details", []interface{}{SystemDetailsToMap(v.SystemDetails)})
		} else {
			s.D.Set("system_details", nil)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeOccurred != nil {
			s.D.Set("time_occurred", v.TimeOccurred.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_os_management_hub.SoftwareSourceEvent:
		s.D.Set("type", "SOFTWARE_SOURCE")

		if v.Data != nil {
			s.D.Set("data", []interface{}{SoftwareSourceEventDataToMap(v.Data)})
		} else {
			s.D.Set("data", nil)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.EventDetails != nil {
			s.D.Set("event_details", *v.EventDetails)
		}

		if v.EventSummary != nil {
			s.D.Set("event_summary", *v.EventSummary)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.Set("id", *v.Id)
		}

		if v.IsManagedByAutonomousLinux != nil {
			s.D.Set("is_managed_by_autonomous_linux", *v.IsManagedByAutonomousLinux)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ResourceId != nil {
			s.D.Set("resource_id", *v.ResourceId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemDetails != nil {
			s.D.Set("system_details", []interface{}{SystemDetailsToMap(v.SystemDetails)})
		} else {
			s.D.Set("system_details", nil)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeOccurred != nil {
			s.D.Set("time_occurred", v.TimeOccurred.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_os_management_hub.SoftwareUpdateEvent:
		s.D.Set("type", "SOFTWARE_UPDATE")

		if v.Data != nil {
			s.D.Set("data", []interface{}{SoftwareUpdateEventDataToMap(v.Data)})
		} else {
			s.D.Set("data", nil)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.EventDetails != nil {
			s.D.Set("event_details", *v.EventDetails)
		}

		if v.EventSummary != nil {
			s.D.Set("event_summary", *v.EventSummary)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.Set("id", *v.Id)
		}

		if v.IsManagedByAutonomousLinux != nil {
			s.D.Set("is_managed_by_autonomous_linux", *v.IsManagedByAutonomousLinux)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ResourceId != nil {
			s.D.Set("resource_id", *v.ResourceId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemDetails != nil {
			s.D.Set("system_details", []interface{}{SystemDetailsToMap(v.SystemDetails)})
		} else {
			s.D.Set("system_details", nil)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeOccurred != nil {
			s.D.Set("time_occurred", v.TimeOccurred.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *s.Res)
		return nil
	}
	return nil
}

func (s *OsManagementHubEventResourceCrud) ChangeEventCompartment() error {
	request := oci_os_management_hub.ChangeEventCompartmentRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	idTmp := s.D.Id()
	request.EventId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub")

	_, err := s.Client.ChangeEventCompartment(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

func (s *OsManagementHubEventResourceCrud) mapToAgentEventData(fieldKeyFormat string) (oci_os_management_hub.AgentEventData, error) {
	result := oci_os_management_hub.AgentEventData{}

	if additionalDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "additional_details")); ok {
		if tmpList := additionalDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "additional_details"), 0)
			tmp, err := s.mapToWorkRequestEventDataAdditionalDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert additional_details, encountered error: %v", err)
			}
			result.AdditionalDetails = &tmp
		}
	}

	if operationType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "operation_type")); ok {
		result.OperationType = oci_os_management_hub.AgentEventDataOperationTypeEnum(operationType.(string))
	}

	if status, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "status")); ok {
		result.Status = oci_os_management_hub.EventStatusEnum(status.(string))
	}

	return result, nil
}

func AgentEventDataToMap(obj *oci_os_management_hub.AgentEventData) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AdditionalDetails != nil {
		result["additional_details"] = []interface{}{WorkRequestEventDataAdditionalDetailsToMap(obj.AdditionalDetails)}
	}

	result["operation_type"] = string(obj.OperationType)

	result["status"] = string(obj.Status)

	return result
}

func EventSummaryToMap(obj oci_os_management_hub.EventSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.EventSummary != nil {
		result["event_summary"] = string(*obj.EventSummary)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsManagedByAutonomousLinux != nil {
		result["is_managed_by_autonomous_linux"] = bool(*obj.IsManagedByAutonomousLinux)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.ResourceId != nil {
		result["resource_id"] = string(*obj.ResourceId)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeOccurred != nil {
		result["time_occurred"] = obj.TimeOccurred.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	result["type"] = string(obj.Type)

	return result
}

func (s *OsManagementHubEventResourceCrud) mapToExploitAttemptAdditionalDetails(fieldKeyFormat string) (oci_os_management_hub.ExploitAttemptAdditionalDetails, error) {
	result := oci_os_management_hub.ExploitAttemptAdditionalDetails{}

	if exploitCves, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "exploit_cves")); ok {
		interfaces := exploitCves.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "exploit_cves")) {
			result.ExploitCves = tmp
		}
	}

	return result, nil
}

func ExploitAttemptAdditionalDetailsToMap(obj *oci_os_management_hub.ExploitAttemptAdditionalDetails) map[string]interface{} {
	result := map[string]interface{}{}

	result["exploit_cves"] = obj.ExploitCves

	return result
}

func (s *OsManagementHubEventResourceCrud) mapToExploitAttemptEventContent(fieldKeyFormat string) (oci_os_management_hub.ExploitAttemptEventContent, error) {
	result := oci_os_management_hub.ExploitAttemptEventContent{}

	if exploitDetectionLogContent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "exploit_detection_log_content")); ok {
		tmp := exploitDetectionLogContent.(string)
		result.ExploitDetectionLogContent = &tmp
	}

	if exploitObjectStoreLocation, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "exploit_object_store_location")); ok {
		tmp := exploitObjectStoreLocation.(string)
		result.ExploitObjectStoreLocation = &tmp
	}

	// TODO: figure out how to support inherited field
	//if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
	//	result.Type = oci_os_management_hub.EventContentTypeEnum(type_.(string))
	//}

	return result, nil
}

func ExploitAttemptEventContentToMap(obj *oci_os_management_hub.ExploitAttemptEventContent) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ExploitDetectionLogContent != nil {
		result["exploit_detection_log_content"] = string(*obj.ExploitDetectionLogContent)
	}

	if obj.ExploitObjectStoreLocation != nil {
		result["exploit_object_store_location"] = string(*obj.ExploitObjectStoreLocation)
	}

	result["type"] = oci_os_management_hub.EventContentTypeExploitAttempt

	return result
}

func (s *OsManagementHubEventResourceCrud) mapToExploitAttemptEventData(fieldKeyFormat string) (oci_os_management_hub.ExploitAttemptEventData, error) {
	result := oci_os_management_hub.ExploitAttemptEventData{}

	if additionalDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "additional_details")); ok {
		if tmpList := additionalDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "additional_details"), 0)
			tmp, err := s.mapToExploitAttemptAdditionalDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert additional_details, encountered error: %v", err)
			}
			result.AdditionalDetails = &tmp
		}
	}

	if content, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "content")); ok {
		if tmpList := content.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "content"), 0)
			tmp, err := s.mapToExploitAttemptEventContent(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert content, encountered error: %v", err)
			}
			result.Content = &tmp
		}
	}

	if eventCount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "event_count")); ok {
		tmp := eventCount.(int)
		result.Count = &tmp
	}

	return result, nil
}

func ExploitAttemptEventDataToMap(obj *oci_os_management_hub.ExploitAttemptEventData) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AdditionalDetails != nil {
		result["additional_details"] = []interface{}{ExploitAttemptAdditionalDetailsToMap(obj.AdditionalDetails)}
	}

	if obj.Content != nil {
		result["content"] = []interface{}{ExploitAttemptEventContentToMap(obj.Content)}
	}

	if obj.Count != nil {
		result["event_count"] = int(*obj.Count)
	}

	return result
}

func (s *OsManagementHubEventResourceCrud) mapToKernelEventAdditionalDetails(fieldKeyFormat string) (oci_os_management_hub.KernelEventAdditionalDetails, error) {
	result := oci_os_management_hub.KernelEventAdditionalDetails{}

	if vmcore, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vmcore")); ok {
		if tmpList := vmcore.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "vmcore"), 0)
			tmp, err := s.mapToVmcoreDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert vmcore, encountered error: %v", err)
			}
			result.Vmcore = &tmp
		}
	}

	return result, nil
}

func KernelEventAdditionalDetailsToMap(obj *oci_os_management_hub.KernelEventAdditionalDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Vmcore != nil {
		result["vmcore"] = []interface{}{VmcoreDetailsToMap(obj.Vmcore)}
	}

	return result
}

func (s *OsManagementHubEventResourceCrud) mapToKernelEventContent(fieldKeyFormat string) (oci_os_management_hub.KernelEventContent, error) {
	result := oci_os_management_hub.KernelEventContent{}

	if contentAvailability, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "content_availability")); ok {
		result.ContentAvailability = oci_os_management_hub.KernelEventContentContentAvailabilityEnum(contentAvailability.(string))
	}

	if contentLocation, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "content_location")); ok {
		tmp := contentLocation.(string)
		result.ContentLocation = &tmp
	}

	if size, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "size")); ok {
		tmp := size.(int)
		result.Size = &tmp
	}

	// TODO: figure out how to support inherited field
	//if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
	//	result.Type = oci_os_management_hub.EventContentTypeEnum(type_.(string))
	//}

	return result, nil
}

func KernelEventContentToMap(obj *oci_os_management_hub.KernelEventContent) map[string]interface{} {
	result := map[string]interface{}{}

	result["content_availability"] = string(obj.ContentAvailability)

	if obj.ContentLocation != nil {
		result["content_location"] = string(*obj.ContentLocation)
	}

	if obj.Size != nil {
		result["size"] = int(*obj.Size)
	}

	result["type"] = oci_os_management_hub.EventContentTypeKernel

	return result
}

func (s *OsManagementHubEventResourceCrud) mapToKernelEventData(fieldKeyFormat string) (oci_os_management_hub.KernelEventData, error) {
	result := oci_os_management_hub.KernelEventData{}

	if additionalDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "additional_details")); ok {
		if tmpList := additionalDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "additional_details"), 0)
			tmp, err := s.mapToKernelEventAdditionalDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert additional_details, encountered error: %v", err)
			}
			result.AdditionalDetails = &tmp
		}
	}

	if content, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "content")); ok {
		if tmpList := content.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "content"), 0)
			tmp, err := s.mapToKernelEventContent(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert content, encountered error: %v", err)
			}
			result.Content = &tmp
		}
	}

	if eventFingerprint, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "event_fingerprint")); ok {
		tmp := eventFingerprint.(string)
		result.EventFingerprint = &tmp
	}

	if eventCount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "event_count")); ok {
		tmp := eventCount.(int)
		result.Count = &tmp
	}

	if reason, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "reason")); ok {
		tmp := reason.(string)
		result.Reason = &tmp
	}

	if timeFirstOccurred, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_first_occurred")); ok {
		tmp, err := time.Parse(time.RFC3339, timeFirstOccurred.(string))
		if err != nil {
			return result, err
		}
		result.TimeFirstOccurred = &oci_common.SDKTime{Time: tmp}
	}

	return result, nil
}

func KernelEventDataToMap(obj *oci_os_management_hub.KernelEventData) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AdditionalDetails != nil {
		result["additional_details"] = []interface{}{KernelEventAdditionalDetailsToMap(obj.AdditionalDetails)}
	}

	if obj.Content != nil {
		result["content"] = []interface{}{KernelEventContentToMap(obj.Content)}
	}

	if obj.EventFingerprint != nil {
		result["event_fingerprint"] = string(*obj.EventFingerprint)
	}

	if obj.Count != nil {
		result["event_count"] = int(*obj.Count)
	}

	if obj.Reason != nil {
		result["reason"] = string(*obj.Reason)
	}

	if obj.TimeFirstOccurred != nil {
		result["time_first_occurred"] = obj.TimeFirstOccurred.Format(time.RFC3339Nano)
	}

	return result
}

func (s *OsManagementHubEventResourceCrud) mapToKspliceUpdateEventData(fieldKeyFormat string) (oci_os_management_hub.KspliceUpdateEventData, error) {
	result := oci_os_management_hub.KspliceUpdateEventData{}

	if additionalDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "additional_details")); ok {
		if tmpList := additionalDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "additional_details"), 0)
			tmp, err := s.mapToWorkRequestEventDataAdditionalDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert additional_details, encountered error: %v", err)
			}
			result.AdditionalDetails = &tmp
		}
	}

	if operationType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "operation_type")); ok {
		result.OperationType = oci_os_management_hub.KspliceUpdateEventDataOperationTypeEnum(operationType.(string))
	}

	if status, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "status")); ok {
		result.Status = oci_os_management_hub.EventStatusEnum(status.(string))
	}

	return result, nil
}

func KspliceUpdateEventDataToMap(obj *oci_os_management_hub.KspliceUpdateEventData) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AdditionalDetails != nil {
		result["additional_details"] = []interface{}{WorkRequestEventDataAdditionalDetailsToMap(obj.AdditionalDetails)}
	}

	result["operation_type"] = string(obj.OperationType)

	result["status"] = string(obj.Status)

	return result
}

func (s *OsManagementHubEventResourceCrud) mapToManagementStationEventData(fieldKeyFormat string) (oci_os_management_hub.ManagementStationEventData, error) {
	result := oci_os_management_hub.ManagementStationEventData{}

	if additionalDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "additional_details")); ok {
		if tmpList := additionalDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "additional_details"), 0)
			tmp, err := s.mapToWorkRequestEventDataAdditionalDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert additional_details, encountered error: %v", err)
			}
			result.AdditionalDetails = &tmp
		}
	}

	if operationType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "operation_type")); ok {
		result.OperationType = oci_os_management_hub.ManagementStationEventDataOperationTypeEnum(operationType.(string))
	}

	if status, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "status")); ok {
		result.Status = oci_os_management_hub.EventStatusEnum(status.(string))
	}

	return result, nil
}

func ManagementStationEventDataToMap(obj *oci_os_management_hub.ManagementStationEventData) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AdditionalDetails != nil {
		result["additional_details"] = []interface{}{WorkRequestEventDataAdditionalDetailsToMap(obj.AdditionalDetails)}
	}

	result["operation_type"] = string(obj.OperationType)

	result["status"] = string(obj.Status)

	return result
}

func (s *OsManagementHubEventResourceCrud) mapToSoftwareSourceEventData(fieldKeyFormat string) (oci_os_management_hub.SoftwareSourceEventData, error) {
	result := oci_os_management_hub.SoftwareSourceEventData{}

	if additionalDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "additional_details")); ok {
		if tmpList := additionalDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "additional_details"), 0)
			tmp, err := s.mapToWorkRequestEventDataAdditionalDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert additional_details, encountered error: %v", err)
			}
			result.AdditionalDetails = &tmp
		}
	}

	if operationType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "operation_type")); ok {
		result.OperationType = oci_os_management_hub.SoftwareSourceEventDataOperationTypeEnum(operationType.(string))
	}

	if status, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "status")); ok {
		result.Status = oci_os_management_hub.EventStatusEnum(status.(string))
	}

	return result, nil
}

func SoftwareSourceEventDataToMap(obj *oci_os_management_hub.SoftwareSourceEventData) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AdditionalDetails != nil {
		result["additional_details"] = []interface{}{WorkRequestEventDataAdditionalDetailsToMap(obj.AdditionalDetails)}
	}

	result["operation_type"] = string(obj.OperationType)

	result["status"] = string(obj.Status)

	return result
}

func (s *OsManagementHubEventResourceCrud) mapToSoftwareUpdateEventData(fieldKeyFormat string) (oci_os_management_hub.SoftwareUpdateEventData, error) {
	result := oci_os_management_hub.SoftwareUpdateEventData{}

	if additionalDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "additional_details")); ok {
		if tmpList := additionalDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "additional_details"), 0)
			tmp, err := s.mapToWorkRequestEventDataAdditionalDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert additional_details, encountered error: %v", err)
			}
			result.AdditionalDetails = &tmp
		}
	}

	if operationType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "operation_type")); ok {
		result.OperationType = oci_os_management_hub.SoftwareUpdateEventDataOperationTypeEnum(operationType.(string))
	}

	if status, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "status")); ok {
		result.Status = oci_os_management_hub.EventStatusEnum(status.(string))
	}

	return result, nil
}

func SoftwareUpdateEventDataToMap(obj *oci_os_management_hub.SoftwareUpdateEventData) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AdditionalDetails != nil {
		result["additional_details"] = []interface{}{WorkRequestEventDataAdditionalDetailsToMap(obj.AdditionalDetails)}
	}

	result["operation_type"] = string(obj.OperationType)

	result["status"] = string(obj.Status)

	return result
}

func SystemDetailsToMap(obj *oci_os_management_hub.SystemDetails) map[string]interface{} {
	result := map[string]interface{}{}

	result["architecture"] = string(obj.Architecture)

	if obj.KspliceEffectiveKernelVersion != nil {
		result["ksplice_effective_kernel_version"] = string(*obj.KspliceEffectiveKernelVersion)
	}

	result["os_family"] = string(obj.OsFamily)

	if obj.OsKernelRelease != nil {
		result["os_kernel_release"] = string(*obj.OsKernelRelease)
	}

	if obj.OsKernelVersion != nil {
		result["os_kernel_version"] = string(*obj.OsKernelVersion)
	}

	if obj.OsName != nil {
		result["os_name"] = string(*obj.OsName)
	}

	if obj.OsSystemVersion != nil {
		result["os_system_version"] = string(*obj.OsSystemVersion)
	}

	return result
}

func (s *OsManagementHubEventResourceCrud) mapToVmcoreDetails(fieldKeyFormat string) (oci_os_management_hub.VmcoreDetails, error) {
	result := oci_os_management_hub.VmcoreDetails{}

	if backtrace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "backtrace")); ok {
		tmp := backtrace.(string)
		result.Backtrace = &tmp
	}

	if component, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "component")); ok {
		tmp := component.(string)
		result.Component = &tmp
	}

	return result, nil
}

func VmcoreDetailsToMap(obj *oci_os_management_hub.VmcoreDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Backtrace != nil {
		result["backtrace"] = string(*obj.Backtrace)
	}

	if obj.Component != nil {
		result["component"] = string(*obj.Component)
	}

	return result
}

func (s *OsManagementHubEventResourceCrud) mapToWorkRequestEventDataAdditionalDetails(fieldKeyFormat string) (oci_os_management_hub.WorkRequestEventDataAdditionalDetails, error) {
	result := oci_os_management_hub.WorkRequestEventDataAdditionalDetails{}

	if initiatorId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "initiator_id")); ok {
		tmp := initiatorId.(string)
		result.InitiatorId = &tmp
	}

	if workRequestIds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "work_request_ids")); ok {
		interfaces := workRequestIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "work_request_ids")) {
			result.WorkRequestIds = tmp
		}
	}

	return result, nil
}

func WorkRequestEventDataAdditionalDetailsToMap(obj *oci_os_management_hub.WorkRequestEventDataAdditionalDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.InitiatorId != nil {
		result["initiator_id"] = string(*obj.InitiatorId)
	}

	result["work_request_ids"] = obj.WorkRequestIds

	return result
}

func (s *OsManagementHubEventResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_os_management_hub.ChangeEventCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.EventId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub")

	_, err := s.Client.ChangeEventCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
