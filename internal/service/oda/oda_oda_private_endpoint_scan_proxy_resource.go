// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oda

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_oda "github.com/oracle/oci-go-sdk/v65/oda"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OdaOdaPrivateEndpointScanProxyResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(45 * time.Minute),
			Update: schema.DefaultTimeout(45 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},
		Create: createOdaOdaPrivateEndpointScanProxy,
		Read:   readOdaOdaPrivateEndpointScanProxy,
		Delete: deleteOdaOdaPrivateEndpointScanProxy,
		Schema: map[string]*schema.Schema{
			// Required
			"oda_private_endpoint_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"protocol": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"scan_listener_infos": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"scan_listener_fqdn": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"scan_listener_ip": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"scan_listener_port": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"scan_listener_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional

			// Computed
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createOdaOdaPrivateEndpointScanProxy(d *schema.ResourceData, m interface{}) error {
	sync := &OdaOdaPrivateEndpointScanProxyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagementClient()
	sync.OdaClient = m.(*client.OracleClients).OdaClient()

	return tfresource.CreateResource(d, sync)
}

func readOdaOdaPrivateEndpointScanProxy(d *schema.ResourceData, m interface{}) error {
	sync := &OdaOdaPrivateEndpointScanProxyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagementClient()
	sync.OdaClient = m.(*client.OracleClients).OdaClient()

	return tfresource.ReadResource(sync)
}

func deleteOdaOdaPrivateEndpointScanProxy(d *schema.ResourceData, m interface{}) error {
	sync := &OdaOdaPrivateEndpointScanProxyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagementClient()
	sync.OdaClient = m.(*client.OracleClients).OdaClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type OdaOdaPrivateEndpointScanProxyResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_oda.ManagementClient
	OdaClient              *oci_oda.OdaClient
	Res                    *oci_oda.OdaPrivateEndpointScanProxy
	DisableNotFoundRetries bool
}

func (s *OdaOdaPrivateEndpointScanProxyResourceCrud) ID() string {
	return GetOdaPrivateEndpointScanProxyCompositeId(s.D.Get("oda_private_endpoint_id").(string), *s.Res.Id)
}

func (s *OdaOdaPrivateEndpointScanProxyResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_oda.OdaPrivateEndpointScanProxyLifecycleStateCreating),
	}
}

func (s *OdaOdaPrivateEndpointScanProxyResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_oda.OdaPrivateEndpointScanProxyLifecycleStateActive),
	}
}

func (s *OdaOdaPrivateEndpointScanProxyResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_oda.OdaPrivateEndpointScanProxyLifecycleStateDeleting),
	}
}

func (s *OdaOdaPrivateEndpointScanProxyResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_oda.OdaPrivateEndpointScanProxyLifecycleStateDeleted),
	}
}

func (s *OdaOdaPrivateEndpointScanProxyResourceCrud) Create() error {
	request := oci_oda.CreateOdaPrivateEndpointScanProxyRequest{}

	if odaPrivateEndpointId, ok := s.D.GetOkExists("oda_private_endpoint_id"); ok {
		tmp := odaPrivateEndpointId.(string)
		request.OdaPrivateEndpointId = &tmp
	}

	if protocol, ok := s.D.GetOkExists("protocol"); ok {
		request.Protocol = oci_oda.OdaPrivateEndpointScanProxyProtocolEnum(protocol.(string))
	}

	if scanListenerInfos, ok := s.D.GetOkExists("scan_listener_infos"); ok {
		interfaces := scanListenerInfos.([]interface{})
		tmp := make([]oci_oda.ScanListenerInfo, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "scan_listener_infos", stateDataIndex)
			converted, err := s.mapToScanListenerInfo(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("scan_listener_infos") {
			request.ScanListenerInfos = tmp
		}
	}

	if scanListenerType, ok := s.D.GetOkExists("scan_listener_type"); ok {
		request.ScanListenerType = oci_oda.OdaPrivateEndpointScanProxyScanListenerTypeEnum(scanListenerType.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "oda")

	response, err := s.Client.CreateOdaPrivateEndpointScanProxy(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	s.setIdFromWorkRequest(workId)
	return s.getOdaPrivateEndpointScanProxyFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "oda"), oci_oda.WorkRequestResourceResourceActionCreate, s.D.Timeout(schema.TimeoutCreate))
}

func (s *OdaOdaPrivateEndpointScanProxyResourceCrud) setIdFromWorkRequest(workId *string) {
	var identifier *string
	var err error

	workRequestResponse := oci_oda.GetWorkRequestResponse{}
	workRequestResponse, err = s.OdaClient.GetWorkRequest(context.Background(),
		oci_oda.GetWorkRequestRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "oda"),
			},
		})
	if err == nil {
		// The work request response contains an array of objects
		for _, res := range workRequestResponse.Resources {
			if res.ResourceType != nil && strings.Contains(strings.ToLower(*res.ResourceType), "oda") {
				identifier = res.ResourceId
				break
			}
		}
	}
	if identifier != nil {
		s.D.SetId(*identifier)
	}
}

func (s *OdaOdaPrivateEndpointScanProxyResourceCrud) getOdaPrivateEndpointScanProxyFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_oda.WorkRequestResourceResourceActionEnum, timeout time.Duration) error {

	// Wait until it finishes
	odaPrivateEndpointScanProxyId, err := odaPrivateEndpointScanProxyWaitForWorkRequest(workId, "oda",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.OdaClient)

	if err != nil {
		return err
	}
	s.D.SetId(*odaPrivateEndpointScanProxyId)

	return s.Get()
}

func odaPrivateEndpointScanProxyWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "oda", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_oda.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func odaPrivateEndpointScanProxyWaitForWorkRequest(wId *string, entityType string, action oci_oda.WorkRequestResourceResourceActionEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_oda.OdaClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "oda")
	retryPolicy.ShouldRetryOperation = odaPrivateEndpointScanProxyWorkRequestShouldRetryFunc(timeout)

	response := oci_oda.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_oda.WorkRequestStatusInProgress),
			string(oci_oda.WorkRequestStatusAccepted),
			string(oci_oda.WorkRequestStatusCanceling),
		},
		Target: []string{
			string(oci_oda.WorkRequestStatusSucceeded),
			string(oci_oda.WorkRequestStatusFailed),
			string(oci_oda.WorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_oda.GetWorkRequestRequest{
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
		if strings.Contains(strings.ToLower(*res.ResourceType), entityType) {
			if res.ResourceAction == action {
				identifier = res.ResourceId
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_oda.WorkRequestStatusFailed || response.Status == oci_oda.WorkRequestStatusCanceled {
		return nil, getErrorFromOdaOdaPrivateEndpointScanProxyWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromOdaOdaPrivateEndpointScanProxyWorkRequest(client *oci_oda.OdaClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_oda.WorkRequestResourceResourceActionEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_oda.ListWorkRequestErrorsRequest{
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

func (s *OdaOdaPrivateEndpointScanProxyResourceCrud) Get() error {
	request := oci_oda.GetOdaPrivateEndpointScanProxyRequest{}

	if odaPrivateEndpointId, ok := s.D.GetOkExists("oda_private_endpoint_id"); ok {
		tmp := odaPrivateEndpointId.(string)
		request.OdaPrivateEndpointId = &tmp
	}

	tmp := s.D.Id()
	request.OdaPrivateEndpointScanProxyId = &tmp

	odaPrivateEndpointId, odaPrivateEndpointScanProxyId, err := parseOdaPrivateEndpointScanProxyCompositeId(s.D.Id())
	if err == nil {
		request.OdaPrivateEndpointId = &odaPrivateEndpointId
		request.OdaPrivateEndpointScanProxyId = &odaPrivateEndpointScanProxyId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "oda")

	response, err := s.Client.GetOdaPrivateEndpointScanProxy(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.OdaPrivateEndpointScanProxy
	return nil
}

func (s *OdaOdaPrivateEndpointScanProxyResourceCrud) Delete() error {
	request := oci_oda.DeleteOdaPrivateEndpointScanProxyRequest{}

	if odaPrivateEndpointId, ok := s.D.GetOkExists("oda_private_endpoint_id"); ok {
		tmp := odaPrivateEndpointId.(string)
		request.OdaPrivateEndpointId = &tmp
	}

	tmp := s.D.Id()
	request.OdaPrivateEndpointScanProxyId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "oda")

	response, err := s.Client.DeleteOdaPrivateEndpointScanProxy(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := odaPrivateEndpointScanProxyWaitForWorkRequest(workId, "oda",
		oci_oda.WorkRequestResourceResourceActionDelete, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.OdaClient)
	return delWorkRequestErr
}

func (s *OdaOdaPrivateEndpointScanProxyResourceCrud) SetData() error {

	odaPrivateEndpointId, odaPrivateEndpointScanProxyId, err := parseOdaPrivateEndpointScanProxyCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("oda_private_endpoint_id", &odaPrivateEndpointId)
		s.D.SetId(odaPrivateEndpointScanProxyId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	s.D.Set("protocol", s.Res.Protocol)

	scanListenerInfos := []interface{}{}
	for _, item := range s.Res.ScanListenerInfos {
		scanListenerInfos = append(scanListenerInfos, ScanListenerInfoToMap(item))
	}
	s.D.Set("scan_listener_infos", scanListenerInfos)

	s.D.Set("scan_listener_type", s.Res.ScanListenerType)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func GetOdaPrivateEndpointScanProxyCompositeId(odaPrivateEndpointId string, odaPrivateEndpointScanProxyId string) string {
	odaPrivateEndpointId = url.PathEscape(odaPrivateEndpointId)
	odaPrivateEndpointScanProxyId = url.PathEscape(odaPrivateEndpointScanProxyId)
	compositeId := "odaPrivateEndpoints/" + odaPrivateEndpointId + "/odaPrivateEndpointScanProxies/" + odaPrivateEndpointScanProxyId
	return compositeId
}

func parseOdaPrivateEndpointScanProxyCompositeId(compositeId string) (odaPrivateEndpointId string, odaPrivateEndpointScanProxyId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("odaPrivateEndpoints/.*/odaPrivateEndpointScanProxies/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	odaPrivateEndpointId, _ = url.PathUnescape(parts[1])
	odaPrivateEndpointScanProxyId, _ = url.PathUnescape(parts[3])

	return
}

func OdaPrivateEndpointScanProxySummaryToMap(obj oci_oda.OdaPrivateEndpointScanProxySummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	result["protocol"] = string(obj.Protocol)

	scanListenerInfos := []interface{}{}
	for _, item := range obj.ScanListenerInfos {
		scanListenerInfos = append(scanListenerInfos, ScanListenerInfoToMap(item))
	}
	result["scan_listener_infos"] = scanListenerInfos

	result["scan_listener_type"] = string(obj.ScanListenerType)

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	return result
}

func (s *OdaOdaPrivateEndpointScanProxyResourceCrud) mapToScanListenerInfo(fieldKeyFormat string) (oci_oda.ScanListenerInfo, error) {
	result := oci_oda.ScanListenerInfo{}

	if scanListenerFqdn, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "scan_listener_fqdn")); ok {
		tmp := scanListenerFqdn.(string)
		result.ScanListenerFqdn = &tmp
	}

	if scanListenerIp, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "scan_listener_ip")); ok {
		tmp := scanListenerIp.(string)
		result.ScanListenerIp = &tmp
	}

	if scanListenerPort, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "scan_listener_port")); ok {
		tmp := scanListenerPort.(int)
		result.ScanListenerPort = &tmp
	}

	return result, nil
}

func ScanListenerInfoToMap(obj oci_oda.ScanListenerInfo) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ScanListenerFqdn != nil {
		result["scan_listener_fqdn"] = string(*obj.ScanListenerFqdn)
	}

	if obj.ScanListenerIp != nil {
		result["scan_listener_ip"] = string(*obj.ScanListenerIp)
	}

	if obj.ScanListenerPort != nil {
		result["scan_listener_port"] = int(*obj.ScanListenerPort)
	}

	return result
}
