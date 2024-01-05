// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package stack_monitoring

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_stack_monitoring "github.com/oracle/oci-go-sdk/v65/stackmonitoring"
)

func StackMonitoringMonitoredResourceResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createStackMonitoringMonitoredResource,
		Read:     readStackMonitoringMonitoredResource,
		Update:   updateStackMonitoringMonitoredResource,
		Delete:   deleteStackMonitoringMonitoredResource,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"additional_aliases": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"credential": {
							Type:     schema.TypeList,
							Required: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"name": {
										Type:     schema.TypeString,
										Required: true,
									},
									"service": {
										Type:     schema.TypeString,
										Required: true,
									},
									"source": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional

									// Computed
								},
							},
						},
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"source": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"additional_credentials": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"credential_type": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"ENCRYPTED",
								"EXISTING",
								"PLAINTEXT",
							}, true),
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"key_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"properties": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"value": {
										Type:     schema.TypeString,
										Optional: true,
									},

									// Computed
								},
							},
						},
						"source": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"type": {
							Type:     schema.TypeString,
							Optional: true,
						},

						// Computed
					},
				},
			},
			"aliases": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"credential": {
							Type:     schema.TypeList,
							Required: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"name": {
										Type:     schema.TypeString,
										Required: true,
									},
									"service": {
										Type:     schema.TypeString,
										Required: true,
									},
									"source": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional

									// Computed
								},
							},
						},
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"source": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"credentials": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"credential_type": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"ENCRYPTED",
								"EXISTING",
								"PLAINTEXT",
							}, true),
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"key_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"properties": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"value": {
										Type:     schema.TypeString,
										Optional: true,
									},

									// Computed
								},
							},
						},
						"source": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"type": {
							Type:     schema.TypeString,
							Optional: true,
						},

						// Computed
					},
				},
			},
			"database_connection_details": {
				Type:     schema.TypeList,
				Optional: true,
				//	Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"port": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"protocol": {
							Type:     schema.TypeString,
							Required: true,
						},
						"service_name": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"connector_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"db_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"db_unique_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"ssl_secret_id": {
							Type:     schema.TypeString,
							Optional: true,
						},

						// Computed
					},
				},
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_resource_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"host_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"license": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"management_agent_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"properties": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"value": {
							Type:     schema.TypeString,
							Optional: true,
						},

						// Computed
					},
				},
			},
			"resource_time_zone": {
				Type:     schema.TypeString,
				Optional: true,
			},

			// Computed
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
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"tenant_id": {
				Type:     schema.TypeString,
				Computed: true,
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

func createStackMonitoringMonitoredResource(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringMonitoredResourceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()

	return tfresource.CreateResource(d, sync)
}

func readStackMonitoringMonitoredResource(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringMonitoredResourceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()

	return tfresource.ReadResource(sync)
}

func updateStackMonitoringMonitoredResource(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringMonitoredResourceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteStackMonitoringMonitoredResource(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringMonitoredResourceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type StackMonitoringMonitoredResourceResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_stack_monitoring.StackMonitoringClient
	Res                    *oci_stack_monitoring.MonitoredResource
	DisableNotFoundRetries bool
}

func (s *StackMonitoringMonitoredResourceResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *StackMonitoringMonitoredResourceResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_stack_monitoring.ResourceLifecycleStateCreating),
	}
}

func (s *StackMonitoringMonitoredResourceResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_stack_monitoring.ResourceLifecycleStateActive),
	}
}

func (s *StackMonitoringMonitoredResourceResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_stack_monitoring.ResourceLifecycleStateDeleting),
	}
}

func (s *StackMonitoringMonitoredResourceResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_stack_monitoring.ResourceLifecycleStateDeleted),
	}
}

func (s *StackMonitoringMonitoredResourceResourceCrud) Create() error {
	request := oci_stack_monitoring.CreateMonitoredResourceRequest{}

	if additionalAliases, ok := s.D.GetOkExists("additional_aliases"); ok {
		interfaces := additionalAliases.([]interface{})
		tmp := make([]oci_stack_monitoring.MonitoredResourceAliasCredential, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "additional_aliases", stateDataIndex)
			converted, err := s.mapToMonitoredResourceAliasCredential(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("additional_aliases") {
			request.AdditionalAliases = tmp
		}
	}

	if additionalCredentials, ok := s.D.GetOkExists("additional_credentials"); ok {
		interfaces := additionalCredentials.([]interface{})
		tmp := make([]oci_stack_monitoring.MonitoredResourceCredential, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "additional_credentials", stateDataIndex)
			converted, err := s.mapToMonitoredResourceCredential(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("additional_credentials") {
			request.AdditionalCredentials = tmp
		}
	}

	if aliases, ok := s.D.GetOkExists("aliases"); ok {
		if tmpList := aliases.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "aliases", 0)
			tmp, err := s.mapToMonitoredResourceAliasCredential(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Aliases = &tmp
		}
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if credentials, ok := s.D.GetOkExists("credentials"); ok {
		if tmpList := credentials.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "credentials", 0)
			tmp, err := s.mapToMonitoredResourceCredential(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Credentials = tmp
		}
	}

	if databaseConnectionDetails, ok := s.D.GetOkExists("database_connection_details"); ok {
		if tmpList := databaseConnectionDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "database_connection_details", 0)
			tmp, err := s.mapToConnectionDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.DatabaseConnectionDetails = &tmp
		}
	}

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

	if externalResourceId, ok := s.D.GetOkExists("external_resource_id"); ok {
		tmp := externalResourceId.(string)
		request.ExternalResourceId = &tmp
	}

	if externalId, ok := s.D.GetOkExists("external_id"); ok {
		tmp := externalId.(string)
		request.ExternalId = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if hostName, ok := s.D.GetOkExists("host_name"); ok {
		tmp := hostName.(string)
		request.HostName = &tmp
	}

	if license, ok := s.D.GetOkExists("license"); ok {
		request.License = oci_stack_monitoring.LicenseTypeEnum(license.(string))
	}

	if managementAgentId, ok := s.D.GetOkExists("management_agent_id"); ok {
		tmp := managementAgentId.(string)
		request.ManagementAgentId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if properties, ok := s.D.GetOkExists("properties"); ok {
		interfaces := properties.([]interface{})
		tmp := make([]oci_stack_monitoring.MonitoredResourceProperty, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "properties", stateDataIndex)
			converted, err := s.mapToMonitoredResourceProperty(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("properties") {
			request.Properties = tmp
		}
	}

	if resourceTimeZone, ok := s.D.GetOkExists("resource_time_zone"); ok {
		tmp := resourceTimeZone.(string)
		request.ResourceTimeZone = &tmp
	}

	if type_, ok := s.D.GetOkExists("type"); ok {
		tmp := type_.(string)
		request.Type = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring")

	response, err := s.Client.CreateMonitoredResource(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getMonitoredResourceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring"), oci_stack_monitoring.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *StackMonitoringMonitoredResourceResourceCrud) getMonitoredResourceFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_stack_monitoring.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	monitoredResourceId, err := monitoredResourceWaitForWorkRequest(workId, "stackmonitoringresource",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*monitoredResourceId)

	return s.Get()
}

func monitoredResourceWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "stack_monitoring", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_stack_monitoring.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func monitoredResourceWaitForWorkRequest(wId *string, entityType string, action oci_stack_monitoring.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_stack_monitoring.StackMonitoringClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "stack_monitoring")
	retryPolicy.ShouldRetryOperation = monitoredResourceWorkRequestShouldRetryFunc(timeout)

	response := oci_stack_monitoring.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_stack_monitoring.OperationStatusInProgress),
			string(oci_stack_monitoring.OperationStatusAccepted),
			string(oci_stack_monitoring.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_stack_monitoring.OperationStatusSucceeded),
			string(oci_stack_monitoring.OperationStatusFailed),
			string(oci_stack_monitoring.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_stack_monitoring.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_stack_monitoring.OperationStatusFailed || response.Status == oci_stack_monitoring.OperationStatusCanceled {
		return nil, getErrorFromStackMonitoringMonitoredResourceWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromStackMonitoringMonitoredResourceWorkRequest(client *oci_stack_monitoring.StackMonitoringClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_stack_monitoring.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_stack_monitoring.ListWorkRequestErrorsRequest{
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

func (s *StackMonitoringMonitoredResourceResourceCrud) Get() error {
	request := oci_stack_monitoring.GetMonitoredResourceRequest{}

	tmp := s.D.Id()
	request.MonitoredResourceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring")

	response, err := s.Client.GetMonitoredResource(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.MonitoredResource
	return nil
}

func (s *StackMonitoringMonitoredResourceResourceCrud) Update() error {

	if _, ok := s.D.GetOkExists("license"); ok && s.D.HasChange("license") {
		err := s.ManageLicense()
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
	request := oci_stack_monitoring.UpdateMonitoredResourceRequest{}

	if additionalAliases, ok := s.D.GetOkExists("additional_aliases"); ok {
		interfaces := additionalAliases.([]interface{})
		tmp := make([]oci_stack_monitoring.MonitoredResourceAliasCredential, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "additional_aliases", stateDataIndex)
			converted, err := s.mapToMonitoredResourceAliasCredential(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("additional_aliases") {
			request.AdditionalAliases = tmp
		}
	}

	if additionalCredentials, ok := s.D.GetOkExists("additional_credentials"); ok {
		interfaces := additionalCredentials.([]interface{})
		tmp := make([]oci_stack_monitoring.MonitoredResourceCredential, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "additional_credentials", stateDataIndex)
			converted, err := s.mapToMonitoredResourceCredential(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("additional_credentials") {
			request.AdditionalCredentials = tmp
		}
	}

	if aliases, ok := s.D.GetOkExists("aliases"); ok {
		if tmpList := aliases.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "aliases", 0)
			tmp, err := s.mapToMonitoredResourceAliasCredential(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Aliases = &tmp
		}
	}

	if credentials, ok := s.D.GetOkExists("credentials"); ok {
		if tmpList := credentials.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "credentials", 0)
			tmp, err := s.mapToMonitoredResourceCredential(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Credentials = tmp
		}
	}

	if databaseConnectionDetails, ok := s.D.GetOkExists("database_connection_details"); ok {
		if tmpList := databaseConnectionDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "database_connection_details", 0)
			tmp, err := s.mapToConnectionDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.DatabaseConnectionDetails = &tmp
		}
	}

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

	if hostName, ok := s.D.GetOkExists("host_name"); ok {
		tmp := hostName.(string)
		request.HostName = &tmp
	}

	tmp := s.D.Id()
	request.MonitoredResourceId = &tmp

	if properties, ok := s.D.GetOkExists("properties"); ok {
		interfaces := properties.([]interface{})
		tmp := make([]oci_stack_monitoring.MonitoredResourceProperty, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "properties", stateDataIndex)
			converted, err := s.mapToMonitoredResourceProperty(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("properties") {
			request.Properties = tmp
		}
	}

	if resourceTimeZone, ok := s.D.GetOkExists("resource_time_zone"); ok {
		tmp := resourceTimeZone.(string)
		request.ResourceTimeZone = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring")

	response, err := s.Client.UpdateMonitoredResource(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getMonitoredResourceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring"), oci_stack_monitoring.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *StackMonitoringMonitoredResourceResourceCrud) Delete() error {
	request := oci_stack_monitoring.DeleteMonitoredResourceRequest{}

	if isDeleteMembers, ok := s.D.GetOkExists("is_delete_members"); ok {
		tmp := isDeleteMembers.(bool)
		request.IsDeleteMembers = &tmp
	}

	tmp := s.D.Id()
	request.MonitoredResourceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring")

	response, err := s.Client.DeleteMonitoredResource(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := monitoredResourceWaitForWorkRequest(workId, "stackmonitoringresource",
		oci_stack_monitoring.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *StackMonitoringMonitoredResourceResourceCrud) SetData() error {
	if s.Res.Aliases != nil {
		s.D.Set("aliases", []interface{}{MonitoredResourceAliasCredentialToMap(s.Res.Aliases)})
	} else {
		s.D.Set("aliases", nil)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.Credentials != nil {
		credentialsArray := []interface{}{}
		if credentialsMap := MonitoredResourceCredentialToMap(&s.Res.Credentials); credentialsMap != nil {
			credentialsArray = append(credentialsArray, credentialsMap)
		}
		s.D.Set("credentials", credentialsArray)
	} else {
		s.D.Set("credentials", nil)
	}

	if s.Res.DatabaseConnectionDetails != nil {
		s.D.Set("database_connection_details", []interface{}{ConnectionDetailsToMap(s.Res.DatabaseConnectionDetails)})
	} else {
		s.D.Set("database_connection_details", nil)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.ExternalId != nil {
		s.D.Set("external_id", *s.Res.ExternalId)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.HostName != nil {
		s.D.Set("host_name", *s.Res.HostName)
	}

	s.D.Set("license", s.Res.License)

	if s.Res.ManagementAgentId != nil {
		s.D.Set("management_agent_id", *s.Res.ManagementAgentId)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	properties := []interface{}{}
	for _, item := range s.Res.Properties {
		properties = append(properties, MonitoredResourcePropertyToMap(item))
	}
	s.D.Set("properties", properties)

	if s.Res.ResourceTimeZone != nil {
		s.D.Set("resource_time_zone", *s.Res.ResourceTimeZone)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TenantId != nil {
		s.D.Set("tenant_id", *s.Res.TenantId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.Type != nil {
		s.D.Set("type", *s.Res.Type)
	}

	return nil
}

func (s *StackMonitoringMonitoredResourceResourceCrud) ManageLicense() error {
	request := oci_stack_monitoring.ManageLicenseRequest{}

	if license, ok := s.D.GetOkExists("license"); ok {
		request.License = oci_stack_monitoring.LicenseTypeEnum(license.(string))
	}

	idTmp := s.D.Id()
	request.MonitoredResourceId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring")

	_, err := s.Client.ManageLicense(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

func (s *StackMonitoringMonitoredResourceResourceCrud) mapToConnectionDetails(fieldKeyFormat string) (oci_stack_monitoring.ConnectionDetails, error) {
	result := oci_stack_monitoring.ConnectionDetails{}

	if connectorId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "connector_id")); ok {
		tmp := connectorId.(string)
		result.ConnectorId = &tmp
	}

	if dbId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "db_id")); ok {
		tmp := dbId.(string)
		result.DbId = &tmp
	}

	if dbUniqueName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "db_unique_name")); ok {
		tmp := dbUniqueName.(string)
		result.DbUniqueName = &tmp
	}

	if port, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "port")); ok {
		tmp := port.(int)
		result.Port = &tmp
	}

	if protocol, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "protocol")); ok {
		result.Protocol = oci_stack_monitoring.ConnectionDetailsProtocolEnum(protocol.(string))
	}

	if serviceName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "service_name")); ok {
		tmp := serviceName.(string)
		result.ServiceName = &tmp
	}

	if sslSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ssl_secret_id")); ok {
		tmp := sslSecretId.(string)
		result.SslSecretId = &tmp
	}

	return result, nil
}

func ConnectionDetailsToMap(obj *oci_stack_monitoring.ConnectionDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ConnectorId != nil {
		result["connector_id"] = string(*obj.ConnectorId)
	}

	if obj.DbId != nil {
		result["db_id"] = string(*obj.DbId)
	}

	if obj.DbUniqueName != nil {
		result["db_unique_name"] = string(*obj.DbUniqueName)
	}

	if obj.Port != nil {
		result["port"] = int(*obj.Port)
	}

	result["protocol"] = string(obj.Protocol)

	if obj.ServiceName != nil {
		result["service_name"] = string(*obj.ServiceName)
	}

	if obj.SslSecretId != nil {
		result["ssl_secret_id"] = string(*obj.SslSecretId)
	}

	return result
}

func (s *StackMonitoringMonitoredResourceResourceCrud) mapToCredentialProperty(fieldKeyFormat string) (oci_stack_monitoring.CredentialProperty, error) {
	result := oci_stack_monitoring.CredentialProperty{}

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

func CredentialPropertyToMap(obj oci_stack_monitoring.CredentialProperty) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *StackMonitoringMonitoredResourceResourceCrud) mapToMonitoredResourceAliasCredential(fieldKeyFormat string) (oci_stack_monitoring.MonitoredResourceAliasCredential, error) {
	result := oci_stack_monitoring.MonitoredResourceAliasCredential{}

	if credential, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "credential")); ok {
		if tmpList := credential.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "credential"), 0)
			tmp, err := s.mapToMonitoredResourceAliasSourceCredential(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert credential, encountered error: %v", err)
			}
			result.Credential = &tmp
		}
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if source, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source")); ok {
		tmp := source.(string)
		result.Source = &tmp
	}

	return result, nil
}

func MonitoredResourceAliasCredentialToMap(obj *oci_stack_monitoring.MonitoredResourceAliasCredential) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Credential != nil {
		result["credential"] = []interface{}{MonitoredResourceAliasSourceCredentialToMap(obj.Credential)}
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Source != nil {
		result["source"] = string(*obj.Source)
	}

	return result
}

func (s *StackMonitoringMonitoredResourceResourceCrud) mapToMonitoredResourceAliasSourceCredential(fieldKeyFormat string) (oci_stack_monitoring.MonitoredResourceAliasSourceCredential, error) {
	result := oci_stack_monitoring.MonitoredResourceAliasSourceCredential{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if service, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "service")); ok {
		tmp := service.(string)
		result.Service = &tmp
	}

	if source, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source")); ok {
		tmp := source.(string)
		result.Source = &tmp
	}

	return result, nil
}

func MonitoredResourceAliasSourceCredentialToMap(obj *oci_stack_monitoring.MonitoredResourceAliasSourceCredential) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Service != nil {
		result["service"] = string(*obj.Service)
	}

	if obj.Source != nil {
		result["source"] = string(*obj.Source)
	}

	return result
}

func (s *StackMonitoringMonitoredResourceResourceCrud) mapToMonitoredResourceCredential(fieldKeyFormat string) (oci_stack_monitoring.MonitoredResourceCredential, error) {
	var baseObject oci_stack_monitoring.MonitoredResourceCredential
	//discriminator
	credentialTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "credential_type"))
	var credentialType string
	if ok {
		credentialType = credentialTypeRaw.(string)
	} else {
		credentialType = "" // default value
	}
	switch strings.ToLower(credentialType) {
	case strings.ToLower("ENCRYPTED"):
		details := oci_stack_monitoring.EncryptedCredentials{}
		if keyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_id")); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if properties, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "properties")); ok {
			interfaces := properties.([]interface{})
			tmp := make([]oci_stack_monitoring.CredentialProperty, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "properties"), stateDataIndex)
				converted, err := s.mapToCredentialProperty(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "properties")) {
				details.Properties = tmp
			}
		}
		if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		if source, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source")); ok {
			tmp := source.(string)
			details.Source = &tmp
		}
		if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
			tmp := type_.(string)
			details.Type = &tmp
		}
		baseObject = details
	case strings.ToLower("EXISTING"):
		details := oci_stack_monitoring.PreExistingCredentials{}
		if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		if source, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source")); ok {
			tmp := source.(string)
			details.Source = &tmp
		}
		if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
			tmp := type_.(string)
			details.Type = &tmp
		}
		baseObject = details
	case strings.ToLower("PLAINTEXT"):
		details := oci_stack_monitoring.PlainTextCredentials{}
		if properties, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "properties")); ok {
			interfaces := properties.([]interface{})
			tmp := make([]oci_stack_monitoring.CredentialProperty, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "properties"), stateDataIndex)
				converted, err := s.mapToCredentialProperty(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "properties")) {
				details.Properties = tmp
			}
		}
		if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		if source, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source")); ok {
			tmp := source.(string)
			details.Source = &tmp
		}
		if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
			tmp := type_.(string)
			details.Type = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown credential_type '%v' was specified", credentialType)
	}
	return baseObject, nil
}

func MonitoredResourceCredentialToMap(obj *oci_stack_monitoring.MonitoredResourceCredential) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_stack_monitoring.EncryptedCredentials:
		result["credential_type"] = "ENCRYPTED"

		if v.KeyId != nil {
			result["key_id"] = string(*v.KeyId)
		}

		properties := []interface{}{}
		for _, item := range v.Properties {
			properties = append(properties, CredentialPropertyToMap(item))
		}
		result["properties"] = properties

		if v.Description != nil {
			result["description"] = string(*v.Description)
		}

		if v.Name != nil {
			result["name"] = string(*v.Name)
		}

		if v.Source != nil {
			result["source"] = string(*v.Source)
		}

		if v.Type != nil {
			result["type"] = string(*v.Type)
		}
	case oci_stack_monitoring.PreExistingCredentials:
		result["credential_type"] = "EXISTING"

		if v.Description != nil {
			result["description"] = string(*v.Description)
		}

		if v.Name != nil {
			result["name"] = string(*v.Name)
		}

		if v.Source != nil {
			result["source"] = string(*v.Source)
		}

		if v.Type != nil {
			result["type"] = string(*v.Type)
		}
	case oci_stack_monitoring.PlainTextCredentials:
		result["credential_type"] = "PLAINTEXT"

		properties := []interface{}{}
		for _, item := range v.Properties {
			properties = append(properties, CredentialPropertyToMap(item))
		}
		result["properties"] = properties

		if v.Description != nil {
			result["description"] = string(*v.Description)
		}

		if v.Name != nil {
			result["name"] = string(*v.Name)
		}

		if v.Source != nil {
			result["source"] = string(*v.Source)
		}

		if v.Type != nil {
			result["type"] = string(*v.Type)
		}
	default:
		log.Printf("[WARN] Received 'credential_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *StackMonitoringMonitoredResourceResourceCrud) mapToMonitoredResourceProperty(fieldKeyFormat string) (oci_stack_monitoring.MonitoredResourceProperty, error) {
	result := oci_stack_monitoring.MonitoredResourceProperty{}

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

func MonitoredResourcePropertyToMap(obj oci_stack_monitoring.MonitoredResourceProperty) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *StackMonitoringMonitoredResourceResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_stack_monitoring.ChangeMonitoredResourceCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.MonitoredResourceId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring")

	response, err := s.Client.ChangeMonitoredResourceCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getMonitoredResourceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring"), oci_stack_monitoring.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
