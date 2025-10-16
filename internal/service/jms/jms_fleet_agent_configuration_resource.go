// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package jms

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_jms "github.com/oracle/oci-go-sdk/v65/jms"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func JmsFleetAgentConfigurationResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createJmsFleetAgentConfiguration,
		Read:     readJmsFleetAgentConfiguration,
		Update:   updateJmsFleetAgentConfiguration,
		Delete:   deleteJmsFleetAgentConfiguration,
		Schema: map[string]*schema.Schema{
			// Required
			"fleet_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"agent_polling_interval_in_minutes": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"is_capturing_ip_address_and_fqdn_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_collecting_managed_instance_metrics_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_collecting_usernames_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_libraries_scan_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"java_usage_tracker_processing_frequency_in_minutes": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"jre_scan_frequency_in_minutes": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"linux_configuration": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"exclude_paths": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"include_paths": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						// Optional

						// Computed
					},
				},
			},
			"mac_os_configuration": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"exclude_paths": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"include_paths": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						// Optional

						// Computed
					},
				},
			},
			"windows_configuration": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"exclude_paths": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"include_paths": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						// Optional

						// Computed
					},
				},
			},
			"work_request_validity_period_in_days": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			// Computed
			"time_last_modified": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createJmsFleetAgentConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &JmsFleetAgentConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaManagementServiceClient()

	return tfresource.CreateResource(d, sync)
}

func readJmsFleetAgentConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &JmsFleetAgentConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaManagementServiceClient()

	return tfresource.ReadResource(sync)
}

func updateJmsFleetAgentConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &JmsFleetAgentConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaManagementServiceClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteJmsFleetAgentConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &JmsFleetAgentConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaManagementServiceClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type JmsFleetAgentConfigurationResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_jms.JavaManagementServiceClient
	Res                    *oci_jms.FleetAgentConfiguration
	DisableNotFoundRetries bool
}

func (s *JmsFleetAgentConfigurationResourceCrud) ID() string {
	return GetFleetAgentConfigurationCompositeId(s.D.Get("fleet_id").(string))
}

func (s *JmsFleetAgentConfigurationResourceCrud) Create() error {
	request := oci_jms.UpdateFleetAgentConfigurationRequest{}

	if agentPollingIntervalInMinutes, ok := s.D.GetOkExists("agent_polling_interval_in_minutes"); ok {
		tmp := agentPollingIntervalInMinutes.(int)
		request.AgentPollingIntervalInMinutes = &tmp
	}

	if fleetId, ok := s.D.GetOkExists("fleet_id"); ok {
		tmp := fleetId.(string)
		request.FleetId = &tmp
	}

	if isCapturingIpAddressAndFqdnEnabled, ok := s.D.GetOkExists("is_capturing_ip_address_and_fqdn_enabled"); ok {
		tmp := isCapturingIpAddressAndFqdnEnabled.(bool)
		request.IsCapturingIpAddressAndFqdnEnabled = &tmp
	}

	if isCollectingManagedInstanceMetricsEnabled, ok := s.D.GetOkExists("is_collecting_managed_instance_metrics_enabled"); ok {
		tmp := isCollectingManagedInstanceMetricsEnabled.(bool)
		request.IsCollectingManagedInstanceMetricsEnabled = &tmp
	}

	if isCollectingUsernamesEnabled, ok := s.D.GetOkExists("is_collecting_usernames_enabled"); ok {
		tmp := isCollectingUsernamesEnabled.(bool)
		request.IsCollectingUsernamesEnabled = &tmp
	}

	if isLibrariesScanEnabled, ok := s.D.GetOkExists("is_libraries_scan_enabled"); ok {
		tmp := isLibrariesScanEnabled.(bool)
		request.IsLibrariesScanEnabled = &tmp
	}

	if javaUsageTrackerProcessingFrequencyInMinutes, ok := s.D.GetOkExists("java_usage_tracker_processing_frequency_in_minutes"); ok {
		tmp := javaUsageTrackerProcessingFrequencyInMinutes.(int)
		request.JavaUsageTrackerProcessingFrequencyInMinutes = &tmp
	}

	if jreScanFrequencyInMinutes, ok := s.D.GetOkExists("jre_scan_frequency_in_minutes"); ok {
		tmp := jreScanFrequencyInMinutes.(int)
		request.JreScanFrequencyInMinutes = &tmp
	}

	if linuxConfiguration, ok := s.D.GetOkExists("linux_configuration"); ok {
		if tmpList := linuxConfiguration.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "linux_configuration", 0)
			tmp, err := s.mapToFleetAgentOsConfiguration(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.LinuxConfiguration = &tmp
		}
	}

	if macOsConfiguration, ok := s.D.GetOkExists("mac_os_configuration"); ok {
		if tmpList := macOsConfiguration.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "mac_os_configuration", 0)
			tmp, err := s.mapToFleetAgentOsConfiguration(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.MacOsConfiguration = &tmp
		}
	}

	if windowsConfiguration, ok := s.D.GetOkExists("windows_configuration"); ok {
		if tmpList := windowsConfiguration.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "windows_configuration", 0)
			tmp, err := s.mapToFleetAgentOsConfiguration(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.WindowsConfiguration = &tmp
		}
	}

	if workRequestValidityPeriodInDays, ok := s.D.GetOkExists("work_request_validity_period_in_days"); ok {
		tmp := workRequestValidityPeriodInDays.(int)
		request.WorkRequestValidityPeriodInDays = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "jms")

	response, err := s.Client.UpdateFleetAgentConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getFleetAgentConfigurationFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "jms"), oci_jms.ActionTypeRelated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *JmsFleetAgentConfigurationResourceCrud) getFleetAgentConfigurationFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_jms.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	fleetAgentConfigurationId, err := fleetAgentConfigurationWaitForWorkRequest(workId, "fleet",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, fleetAgentConfigurationId)
		_, cancelErr := s.Client.CancelWorkRequest(context.Background(),
			oci_jms.CancelWorkRequestRequest{
				WorkRequestId: workId,
				RequestMetadata: oci_common.RequestMetadata{
					RetryPolicy: retryPolicy,
				},
			})
		if cancelErr != nil {
			log.Printf("[DEBUG] cleanup cancelWorkRequest failed with the error: %v\n", cancelErr)
		}
		return err
	}
	s.D.SetId(*fleetAgentConfigurationId)

	return s.Get()
}

func fleetAgentConfigurationWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "jms", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_jms.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func fleetAgentConfigurationWaitForWorkRequest(wId *string, entityType string, action oci_jms.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_jms.JavaManagementServiceClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "jms")
	retryPolicy.ShouldRetryOperation = fleetAgentConfigurationWorkRequestShouldRetryFunc(timeout)

	response := oci_jms.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_jms.OperationStatusInProgress),
			string(oci_jms.OperationStatusAccepted),
			string(oci_jms.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_jms.OperationStatusSucceeded),
			string(oci_jms.OperationStatusFailed),
			string(oci_jms.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_jms.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_jms.OperationStatusFailed || response.Status == oci_jms.OperationStatusCanceled {
		return nil, getErrorFromJmsFleetAgentConfigurationWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	compositeId := GetFleetAgentConfigurationCompositeId(*identifier)

	return &compositeId, nil
}

func getErrorFromJmsFleetAgentConfigurationWorkRequest(client *oci_jms.JavaManagementServiceClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_jms.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_jms.ListWorkRequestErrorsRequest{
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

func (s *JmsFleetAgentConfigurationResourceCrud) Get() error {
	request := oci_jms.GetFleetAgentConfigurationRequest{}

	if fleetId, ok := s.D.GetOkExists("fleet_id"); ok {
		tmp := fleetId.(string)
		request.FleetId = &tmp
	}

	fleetId, err := parseFleetAgentConfigurationCompositeId(s.D.Id())
	if err == nil {
		request.FleetId = &fleetId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "jms")

	response, err := s.Client.GetFleetAgentConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.FleetAgentConfiguration
	return nil
}

func (s *JmsFleetAgentConfigurationResourceCrud) Update() error {
	request := oci_jms.UpdateFleetAgentConfigurationRequest{}

	if agentPollingIntervalInMinutes, ok := s.D.GetOkExists("agent_polling_interval_in_minutes"); ok {
		tmp := agentPollingIntervalInMinutes.(int)
		request.AgentPollingIntervalInMinutes = &tmp
	}

	if fleetId, ok := s.D.GetOkExists("fleet_id"); ok {
		tmp := fleetId.(string)
		request.FleetId = &tmp
	}

	if isCapturingIpAddressAndFqdnEnabled, ok := s.D.GetOkExists("is_capturing_ip_address_and_fqdn_enabled"); ok {
		tmp := isCapturingIpAddressAndFqdnEnabled.(bool)
		request.IsCapturingIpAddressAndFqdnEnabled = &tmp
	}

	if isCollectingManagedInstanceMetricsEnabled, ok := s.D.GetOkExists("is_collecting_managed_instance_metrics_enabled"); ok {
		tmp := isCollectingManagedInstanceMetricsEnabled.(bool)
		request.IsCollectingManagedInstanceMetricsEnabled = &tmp
	}

	if isCollectingUsernamesEnabled, ok := s.D.GetOkExists("is_collecting_usernames_enabled"); ok {
		tmp := isCollectingUsernamesEnabled.(bool)
		request.IsCollectingUsernamesEnabled = &tmp
	}

	if isLibrariesScanEnabled, ok := s.D.GetOkExists("is_libraries_scan_enabled"); ok {
		tmp := isLibrariesScanEnabled.(bool)
		request.IsLibrariesScanEnabled = &tmp
	}

	if javaUsageTrackerProcessingFrequencyInMinutes, ok := s.D.GetOkExists("java_usage_tracker_processing_frequency_in_minutes"); ok {
		tmp := javaUsageTrackerProcessingFrequencyInMinutes.(int)
		request.JavaUsageTrackerProcessingFrequencyInMinutes = &tmp
	}

	if jreScanFrequencyInMinutes, ok := s.D.GetOkExists("jre_scan_frequency_in_minutes"); ok {
		tmp := jreScanFrequencyInMinutes.(int)
		request.JreScanFrequencyInMinutes = &tmp
	}

	if linuxConfiguration, ok := s.D.GetOkExists("linux_configuration"); ok {
		if tmpList := linuxConfiguration.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "linux_configuration", 0)
			tmp, err := s.mapToFleetAgentOsConfiguration(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.LinuxConfiguration = &tmp
		}
	}

	if macOsConfiguration, ok := s.D.GetOkExists("mac_os_configuration"); ok {
		if tmpList := macOsConfiguration.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "mac_os_configuration", 0)
			tmp, err := s.mapToFleetAgentOsConfiguration(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.MacOsConfiguration = &tmp
		}
	}

	if windowsConfiguration, ok := s.D.GetOkExists("windows_configuration"); ok {
		if tmpList := windowsConfiguration.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "windows_configuration", 0)
			tmp, err := s.mapToFleetAgentOsConfiguration(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.WindowsConfiguration = &tmp
		}
	}

	if workRequestValidityPeriodInDays, ok := s.D.GetOkExists("work_request_validity_period_in_days"); ok {
		tmp := workRequestValidityPeriodInDays.(int)
		request.WorkRequestValidityPeriodInDays = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "jms")

	response, err := s.Client.UpdateFleetAgentConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getFleetAgentConfigurationFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "jms"), oci_jms.ActionTypeRelated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *JmsFleetAgentConfigurationResourceCrud) Delete() error {
	return nil
}

func (s *JmsFleetAgentConfigurationResourceCrud) SetData() error {

	fleetId, err := parseFleetAgentConfigurationCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("fleet_id", &fleetId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.AgentPollingIntervalInMinutes != nil {
		s.D.Set("agent_polling_interval_in_minutes", *s.Res.AgentPollingIntervalInMinutes)
	}

	if s.Res.IsCapturingIpAddressAndFqdnEnabled != nil {
		s.D.Set("is_capturing_ip_address_and_fqdn_enabled", *s.Res.IsCapturingIpAddressAndFqdnEnabled)
	}

	if s.Res.IsCollectingManagedInstanceMetricsEnabled != nil {
		s.D.Set("is_collecting_managed_instance_metrics_enabled", *s.Res.IsCollectingManagedInstanceMetricsEnabled)
	}

	if s.Res.IsCollectingUsernamesEnabled != nil {
		s.D.Set("is_collecting_usernames_enabled", *s.Res.IsCollectingUsernamesEnabled)
	}

	if s.Res.IsLibrariesScanEnabled != nil {
		s.D.Set("is_libraries_scan_enabled", *s.Res.IsLibrariesScanEnabled)
	}

	if s.Res.JavaUsageTrackerProcessingFrequencyInMinutes != nil {
		s.D.Set("java_usage_tracker_processing_frequency_in_minutes", *s.Res.JavaUsageTrackerProcessingFrequencyInMinutes)
	}

	if s.Res.JreScanFrequencyInMinutes != nil {
		s.D.Set("jre_scan_frequency_in_minutes", *s.Res.JreScanFrequencyInMinutes)
	}

	if s.Res.LinuxConfiguration != nil {
		s.D.Set("linux_configuration", []interface{}{FleetAgentOsConfigurationToMap(s.Res.LinuxConfiguration)})
	} else {
		s.D.Set("linux_configuration", nil)
	}

	if s.Res.MacOsConfiguration != nil {
		s.D.Set("mac_os_configuration", []interface{}{FleetAgentOsConfigurationToMap(s.Res.MacOsConfiguration)})
	} else {
		s.D.Set("mac_os_configuration", nil)
	}

	if s.Res.TimeLastModified != nil {
		s.D.Set("time_last_modified", s.Res.TimeLastModified.String())
	}

	if s.Res.WindowsConfiguration != nil {
		s.D.Set("windows_configuration", []interface{}{FleetAgentOsConfigurationToMap(s.Res.WindowsConfiguration)})
	} else {
		s.D.Set("windows_configuration", nil)
	}

	if s.Res.WorkRequestValidityPeriodInDays != nil {
		s.D.Set("work_request_validity_period_in_days", *s.Res.WorkRequestValidityPeriodInDays)
	}

	return nil
}

func GetFleetAgentConfigurationCompositeId(fleetId string) string {
	fleetId = url.PathEscape(fleetId)
	compositeId := "fleets/" + fleetId + "/agentConfiguration"
	return compositeId
}

func parseFleetAgentConfigurationCompositeId(compositeId string) (fleetId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("fleets/.*/agentConfiguration", compositeId)
	if !match || len(parts) != 3 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	fleetId, _ = url.PathUnescape(parts[1])

	return
}

func (s *JmsFleetAgentConfigurationResourceCrud) mapToFleetAgentOsConfiguration(fieldKeyFormat string) (oci_jms.FleetAgentOsConfiguration, error) {
	result := oci_jms.FleetAgentOsConfiguration{}

	if excludePaths, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "exclude_paths")); ok {
		interfaces := excludePaths.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "exclude_paths")) {
			result.ExcludePaths = tmp
		}
	}

	if includePaths, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "include_paths")); ok {
		interfaces := includePaths.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "include_paths")) {
			result.IncludePaths = tmp
		}
	}

	return result, nil
}

func FleetAgentOsConfigurationToMap(obj *oci_jms.FleetAgentOsConfiguration) map[string]interface{} {
	result := map[string]interface{}{}

	result["exclude_paths"] = obj.ExcludePaths

	result["include_paths"] = obj.IncludePaths

	return result
}
