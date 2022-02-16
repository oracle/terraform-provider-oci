// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package analytics

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_analytics "github.com/oracle/oci-go-sdk/v58/analytics"
	oci_common "github.com/oracle/oci-go-sdk/v58/common"
)

func AnalyticsAnalyticsInstancePrivateAccessChannelResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("2h0m"),
			Update: tfresource.GetTimeoutDuration("2h0m"),
			Delete: tfresource.GetTimeoutDuration("2h0m"),
		},
		Create: createAnalyticsAnalyticsInstancePrivateAccessChannel,
		Read:   readAnalyticsAnalyticsInstancePrivateAccessChannel,
		Update: updateAnalyticsAnalyticsInstancePrivateAccessChannel,
		Delete: deleteAnalyticsAnalyticsInstancePrivateAccessChannel,
		Schema: map[string]*schema.Schema{
			// Required
			"analytics_instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"private_source_dns_zones": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"dns_zone": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"description": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vcn_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional

			// Computed
			"egress_source_ip_addresses": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"ip_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"key": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createAnalyticsAnalyticsInstancePrivateAccessChannel(d *schema.ResourceData, m interface{}) error {
	sync := &AnalyticsAnalyticsInstancePrivateAccessChannelResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AnalyticsClient()

	return tfresource.CreateResource(d, sync)
}

func readAnalyticsAnalyticsInstancePrivateAccessChannel(d *schema.ResourceData, m interface{}) error {
	sync := &AnalyticsAnalyticsInstancePrivateAccessChannelResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AnalyticsClient()

	return tfresource.ReadResource(sync)
}

func updateAnalyticsAnalyticsInstancePrivateAccessChannel(d *schema.ResourceData, m interface{}) error {
	sync := &AnalyticsAnalyticsInstancePrivateAccessChannelResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AnalyticsClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteAnalyticsAnalyticsInstancePrivateAccessChannel(d *schema.ResourceData, m interface{}) error {
	sync := &AnalyticsAnalyticsInstancePrivateAccessChannelResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AnalyticsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type AnalyticsAnalyticsInstancePrivateAccessChannelResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_analytics.AnalyticsClient
	Res                    *oci_analytics.PrivateAccessChannel
	WorkRequest            *oci_analytics.WorkRequest
	DisableNotFoundRetries bool
}

func (s *AnalyticsAnalyticsInstancePrivateAccessChannelResourceCrud) ID() string {
	if s.WorkRequest != nil {
		if s.WorkRequest.Status == oci_analytics.WorkRequestStatusSucceeded {
			return s.D.Id()
		} else {
			return *s.WorkRequest.Id
		}
	}
	return ""
}

func (s *AnalyticsAnalyticsInstancePrivateAccessChannelResourceCrud) Create() error {
	request := oci_analytics.CreatePrivateAccessChannelRequest{}

	if analyticsInstanceId, ok := s.D.GetOkExists("analytics_instance_id"); ok {
		tmp := analyticsInstanceId.(string)
		request.AnalyticsInstanceId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if privateSourceDnsZones, ok := s.D.GetOkExists("private_source_dns_zones"); ok {
		interfaces := privateSourceDnsZones.([]interface{})
		tmp := make([]oci_analytics.PrivateSourceDnsZone, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "private_source_dns_zones", stateDataIndex)
			converted, err := s.mapToPrivateSourceDnsZone(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("private_source_dns_zones") {
			request.PrivateSourceDnsZones = tmp
		}
	}

	if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
		tmp := subnetId.(string)
		request.SubnetId = &tmp
	}

	if vcnId, ok := s.D.GetOkExists("vcn_id"); ok {
		tmp := vcnId.(string)
		request.VcnId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "analytics")

	response, err := s.Client.CreatePrivateAccessChannel(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	returnError := s.getAnalyticsInstancePrivateAccessChannelFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "analytics"), oci_analytics.WorkRequestActionResultPrivateAccessChannelCreated, s.D.Timeout(schema.TimeoutCreate))
	getWorkRequestRequest := oci_analytics.GetWorkRequestRequest{}
	getWorkRequestRequest.WorkRequestId = workId
	getWorkRequestRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "analytics")
	workRequestResponse, _ := s.Client.GetWorkRequest(context.Background(), getWorkRequestRequest)
	s.WorkRequest = &workRequestResponse.WorkRequest
	return returnError
}

func (s *AnalyticsAnalyticsInstancePrivateAccessChannelResourceCrud) getAnalyticsInstancePrivateAccessChannelFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_analytics.WorkRequestActionResultEnum, timeout time.Duration) error {

	// Wait until it finishes
	analyticsInstanceId, err := analyticsInstancePrivateAccessChannelWaitForWorkRequest(workId, "analytics",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, analyticsInstanceId)
		_, cancelErr := s.Client.DeleteWorkRequest(context.Background(),
			oci_analytics.DeleteWorkRequestRequest{
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

	request := oci_analytics.GetAnalyticsInstanceRequest{}
	request.AnalyticsInstanceId = analyticsInstanceId

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "analytics")

	response, err := s.Client.GetAnalyticsInstance(context.Background(), request)
	if err != nil {
		return err
	}

	pacKey := ""

	if len(response.AnalyticsInstance.PrivateAccessChannels) == 0 {
		return fmt.Errorf("[ERROR] unable to find private access channel resource from analytics instance resource")
	} else {
		for key, _ := range response.AnalyticsInstance.PrivateAccessChannels {
			pacKey = key
		}
	}

	compositeId := getAnalyticsInstancePrivateAccessChannelCompositeId(*analyticsInstanceId, pacKey)
	s.D.SetId(compositeId)

	return s.Get()
}

func analyticsInstancePrivateAccessChannelWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "analytics", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_analytics.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func analyticsInstancePrivateAccessChannelWaitForWorkRequest(wId *string, entityType string, action oci_analytics.WorkRequestActionResultEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_analytics.AnalyticsClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "analytics")
	retryPolicy.ShouldRetryOperation = analyticsInstancePrivateAccessChannelWorkRequestShouldRetryFunc(timeout)

	response := oci_analytics.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_analytics.WorkRequestStatusInProgress),
			string(oci_analytics.WorkRequestStatusAccepted),
			string(oci_analytics.WorkRequestStatusCanceling),
		},
		Target: []string{
			string(oci_analytics.WorkRequestStatusSucceeded),
			string(oci_analytics.WorkRequestStatusFailed),
			string(oci_analytics.WorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_analytics.GetWorkRequestRequest{
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
	for _, res := range response.WorkRequest.Resources {
		if res.ResourceType == "ANALYTICS_INSTANCE" {
			if res.ActionResult == action {
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_analytics.WorkRequestStatusFailed || response.Status == oci_analytics.WorkRequestStatusCanceled {
		return nil, getErrorFromAnalyticsAnalyticsInstancePrivateAccessChannelWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromAnalyticsAnalyticsInstancePrivateAccessChannelWorkRequest(client *oci_analytics.AnalyticsClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_analytics.WorkRequestActionResultEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_analytics.ListWorkRequestErrorsRequest{
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

func (s *AnalyticsAnalyticsInstancePrivateAccessChannelResourceCrud) Get() error {
	request := oci_analytics.GetPrivateAccessChannelRequest{}
	analyticsInstanceId, privateAccessChannelKey, err := parseAnalyticsInstancePrivateAccessChannelCompositeId(s.D.Id())
	if err == nil {
		request.AnalyticsInstanceId = &analyticsInstanceId
		request.PrivateAccessChannelKey = &privateAccessChannelKey
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "analytics")

	response, err := s.Client.GetPrivateAccessChannel(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.PrivateAccessChannel
	return nil
}

func (s *AnalyticsAnalyticsInstancePrivateAccessChannelResourceCrud) Update() error {
	request := oci_analytics.UpdatePrivateAccessChannelRequest{}
	// The PAC api will give an error if certain values are specified in Update that have not changed.  Therefore, we must get the current value of the PAC and compare
	// the values specified in the terraform payload with the current values, and only include those that are different.

	if analyticsInstanceId, ok := s.D.GetOkExists("analytics_instance_id"); ok {
		tmp := analyticsInstanceId.(string)
		request.AnalyticsInstanceId = &tmp
	}

	if privateAccessChannelKey, ok := s.D.GetOkExists("key"); ok {
		tmp := privateAccessChannelKey.(string)
		request.PrivateAccessChannelKey = &tmp
	}

	getRequest := oci_analytics.GetPrivateAccessChannelRequest{}
	getRequest.AnalyticsInstanceId = request.AnalyticsInstanceId
	getRequest.PrivateAccessChannelKey = request.PrivateAccessChannelKey
	getRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "analytics")

	getResponse, err := s.Client.GetPrivateAccessChannel(context.Background(), getRequest)
	if err != nil {
		return err
	}

	currentPAC := getResponse.PrivateAccessChannel

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		if tmp != *currentPAC.DisplayName {
			request.DisplayName = &tmp
		}
	}

	if privateSourceDnsZones, ok := s.D.GetOkExists("private_source_dns_zones"); ok {
		interfaces := privateSourceDnsZones.([]interface{})
		tmp := make([]oci_analytics.PrivateSourceDnsZone, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "private_source_dns_zones", stateDataIndex)
			converted, err := s.mapToPrivateSourceDnsZone(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("private_source_dns_zones") {
			request.PrivateSourceDnsZones = tmp
		}
	}

	if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
		tmp := subnetId.(string)
		if tmp != *currentPAC.SubnetId {
			request.SubnetId = &tmp
		}
	}

	if vcnId, ok := s.D.GetOkExists("vcn_id"); ok {
		tmp := vcnId.(string)
		if tmp != *currentPAC.VcnId {
			request.VcnId = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "analytics")

	response, err := s.Client.UpdatePrivateAccessChannel(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getAnalyticsInstancePrivateAccessChannelFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "analytics"), oci_analytics.WorkRequestActionResultPrivateAccessChannelUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *AnalyticsAnalyticsInstancePrivateAccessChannelResourceCrud) Delete() error {
	request := oci_analytics.DeletePrivateAccessChannelRequest{}

	if analyticsInstanceId, ok := s.D.GetOkExists("analytics_instance_id"); ok {
		tmp := analyticsInstanceId.(string)
		request.AnalyticsInstanceId = &tmp
	}

	if privateAccessChannelKey, ok := s.D.GetOkExists("key"); ok {
		tmp := privateAccessChannelKey.(string)
		request.PrivateAccessChannelKey = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "analytics")

	response, err := s.Client.DeletePrivateAccessChannel(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := analyticsInstancePrivateAccessChannelWaitForWorkRequest(workId, "analytics",
		oci_analytics.WorkRequestActionResultPrivateAccessChannelDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *AnalyticsAnalyticsInstancePrivateAccessChannelResourceCrud) SetData() error {

	analyticsInstanceId, privateAccessChannelKey, err := parseAnalyticsInstancePrivateAccessChannelCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("analytics_instance_id", &analyticsInstanceId)
		s.D.Set("private_access_channel_key", &privateAccessChannelKey)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("egress_source_ip_addresses", s.Res.EgressSourceIpAddresses)

	if s.Res.IpAddress != nil {
		s.D.Set("ip_address", *s.Res.IpAddress)
	}

	if s.Res.Key != nil {
		s.D.Set("key", *s.Res.Key)
	}

	privateSourceDnsZones := []interface{}{}
	for _, item := range s.Res.PrivateSourceDnsZones {
		privateSourceDnsZones = append(privateSourceDnsZones, PrivateSourceDnsZoneToMap(item))
	}
	s.D.Set("private_source_dns_zones", privateSourceDnsZones)

	if s.Res.SubnetId != nil {
		s.D.Set("subnet_id", *s.Res.SubnetId)
	}

	if s.Res.VcnId != nil {
		s.D.Set("vcn_id", *s.Res.VcnId)
	}

	return nil
}

func getAnalyticsInstancePrivateAccessChannelCompositeId(analyticsInstanceId string, privateAccessChannelKey string) string {
	analyticsInstanceId = url.PathEscape(analyticsInstanceId)
	privateAccessChannelKey = url.PathEscape(privateAccessChannelKey)
	compositeId := "analyticsInstances/" + analyticsInstanceId + "/privateAccessChannels/" + privateAccessChannelKey
	return compositeId
}

func parseAnalyticsInstancePrivateAccessChannelCompositeId(compositeId string) (analyticsInstanceId string, privateAccessChannelKey string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("analyticsInstances/.*/privateAccessChannels/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	analyticsInstanceId, _ = url.PathUnescape(parts[1])
	privateAccessChannelKey, _ = url.PathUnescape(parts[3])

	return
}

func (s *AnalyticsAnalyticsInstancePrivateAccessChannelResourceCrud) mapToPrivateSourceDnsZone(fieldKeyFormat string) (oci_analytics.PrivateSourceDnsZone, error) {
	result := oci_analytics.PrivateSourceDnsZone{}

	if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
		tmp := description.(string)
		result.Description = &tmp
	}

	if dnsZone, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "dns_zone")); ok {
		tmp := dnsZone.(string)
		result.DnsZone = &tmp
	}

	return result, nil
}
