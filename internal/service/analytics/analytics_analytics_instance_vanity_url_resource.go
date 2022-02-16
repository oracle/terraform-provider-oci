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

func AnalyticsAnalyticsInstanceVanityUrlResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createAnalyticsAnalyticsInstanceVanityUrl,
		Read:     readAnalyticsAnalyticsInstanceVanityUrl,
		Update:   updateAnalyticsAnalyticsInstanceVanityUrl,
		Delete:   deleteAnalyticsAnalyticsInstanceVanityUrl,
		Schema: map[string]*schema.Schema{
			// Required
			"analytics_instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ca_certificate": {
				Type:     schema.TypeString,
				Required: true,
			},
			"hosts": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"private_key": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"public_certificate": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				//Computed: true,
				ForceNew: true,
			},
			"passphrase": {
				Type:     schema.TypeString,
				Optional: true,
				//Computed:  true,
				Sensitive: true,
			},
			// Computed
		},
	}
}

func createAnalyticsAnalyticsInstanceVanityUrl(d *schema.ResourceData, m interface{}) error {
	sync := &AnalyticsAnalyticsInstanceVanityUrlResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AnalyticsClient()

	return tfresource.CreateResource(d, sync)
}

func readAnalyticsAnalyticsInstanceVanityUrl(d *schema.ResourceData, m interface{}) error {
	sync := &AnalyticsAnalyticsInstanceVanityUrlResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AnalyticsClient()

	return tfresource.ReadResource(sync)
}

func updateAnalyticsAnalyticsInstanceVanityUrl(d *schema.ResourceData, m interface{}) error {
	sync := &AnalyticsAnalyticsInstanceVanityUrlResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AnalyticsClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteAnalyticsAnalyticsInstanceVanityUrl(d *schema.ResourceData, m interface{}) error {
	sync := &AnalyticsAnalyticsInstanceVanityUrlResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AnalyticsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type AnalyticsAnalyticsInstanceVanityUrlResourceCrud struct {
	tfresource.BaseCrud
	Client      *oci_analytics.AnalyticsClient
	Res         *oci_analytics.VanityUrlDetails
	WorkRequest *oci_analytics.WorkRequest
	//	VanityUrlKey		   string
	DisableNotFoundRetries bool
}

func (s *AnalyticsAnalyticsInstanceVanityUrlResourceCrud) ID() string {
	if s.WorkRequest != nil {
		if s.WorkRequest.Status == oci_analytics.WorkRequestStatusSucceeded {
			return s.D.Id()
		} else {
			return *s.WorkRequest.Id
		}
	}
	return ""
}

func (s *AnalyticsAnalyticsInstanceVanityUrlResourceCrud) Create() error {
	request := oci_analytics.CreateVanityUrlRequest{}

	if analyticsInstanceId, ok := s.D.GetOkExists("analytics_instance_id"); ok {
		tmp := analyticsInstanceId.(string)
		request.AnalyticsInstanceId = &tmp
	}

	if caCertificate, ok := s.D.GetOkExists("ca_certificate"); ok {
		tmp := caCertificate.(string)
		request.CaCertificate = &tmp
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if hosts, ok := s.D.GetOkExists("hosts"); ok {
		interfaces := hosts.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("hosts") {
			request.Hosts = tmp
		}
	}

	if passphrase, ok := s.D.GetOkExists("passphrase"); ok {
		tmp := passphrase.(string)
		request.Passphrase = &tmp
	}

	if privateKey, ok := s.D.GetOkExists("private_key"); ok {
		tmp := privateKey.(string)
		request.PrivateKey = &tmp
	}

	if publicCertificate, ok := s.D.GetOkExists("public_certificate"); ok {
		tmp := publicCertificate.(string)
		request.PublicCertificate = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "analytics")
	response, err := s.Client.CreateVanityUrl(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId

	returnError := s.getAnalyticsInstanceVanityUrlFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "analytics"), oci_analytics.WorkRequestActionResultVanityUrlCreated, s.D.Timeout(schema.TimeoutCreate))
	getWorkRequestRequest := oci_analytics.GetWorkRequestRequest{}
	getWorkRequestRequest.WorkRequestId = workId
	getWorkRequestRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "analytics")
	workRequestResponse, err := s.Client.GetWorkRequest(context.Background(), getWorkRequestRequest)
	s.WorkRequest = &workRequestResponse.WorkRequest
	return returnError
}

// TODO:  Make sure this isn't being used anywhere and delete it
func (s *AnalyticsAnalyticsInstanceVanityUrlResourceCrud) getAnalyticsInstanceVanityUrlFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_analytics.WorkRequestActionResultEnum, timeout time.Duration) error {

	// Wait until it finishes
	analyticsInstanceId, err := analyticsInstanceVanityUrlWaitForWorkRequest(workId, "analytics",
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
	// Figure out what the vanity url is and set the composite id

	request := oci_analytics.GetAnalyticsInstanceRequest{}

	// Analytics instance id will always be in here, so it's ok to just get it from there
	if analyticsInstanceId, ok := s.D.GetOkExists("analytics_instance_id"); ok {
		tmp := analyticsInstanceId.(string)
		request.AnalyticsInstanceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "analytics")

	response, err := s.Client.GetAnalyticsInstance(context.Background(), request)
	if err != nil {
		return err
	}

	// If we have the key in s, we can just use it.  Otherwise, have to figure out the key by searching the hosts in the instance for one of the ones
	// the user specified
	vanityUrlKey := ""

	if len(response.AnalyticsInstance.VanityUrlDetails) == 0 {
		return fmt.Errorf("[ERROR] unable to find vanity url resource from analytics instance resource")
	} else {
		createHosts, _ := s.D.GetOk("hosts")
		interfaces := createHosts.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		hostToSearchFor := tmp[0]
	finish:
		for key, value := range response.AnalyticsInstance.VanityUrlDetails {
			for _, host := range value.Hosts {
				if host == hostToSearchFor {
					vanityUrlKey = key
					break finish
				}
			}
		}
	}

	if vanityUrlKey == "" {
		return fmt.Errorf("[ERROR] Could not determine vanity url key value")
	}

	compositeId := getAnalyticsInstanceVanityUrlCompositeId(*analyticsInstanceId, vanityUrlKey)
	s.D.SetId(compositeId)
	return s.Get()
}

func analyticsInstanceVanityUrlWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func analyticsInstanceVanityUrlWaitForWorkRequest(wId *string, entityType string, action oci_analytics.WorkRequestActionResultEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_analytics.AnalyticsClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "analytics")
	retryPolicy.ShouldRetryOperation = analyticsInstanceVanityUrlWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromAnalyticsAnalyticsInstanceVanityUrlWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromAnalyticsAnalyticsInstanceVanityUrlWorkRequest(client *oci_analytics.AnalyticsClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_analytics.WorkRequestActionResultEnum) error {
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

func (s *AnalyticsAnalyticsInstanceVanityUrlResourceCrud) Get() error {
	analyticsInstanceId, vanityUrlKey, err := parseAnalyticsInstanceVanityUrlCompositeId(s.D.Id())
	if err != nil {
		return fmt.Errorf("[ERROR] unable to find parse vanity url key from id %v", s.D.Id())
	}

	request := oci_analytics.GetAnalyticsInstanceRequest{}
	request.AnalyticsInstanceId = &analyticsInstanceId

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "analytics")

	response, err := s.Client.GetAnalyticsInstance(context.Background(), request)
	if err != nil {
		return err
	}

	if vanityUrlKey == "" {
		return fmt.Errorf("[ERROR] Could not determine vanity url key value")
	}
	vanityUrlValue := response.AnalyticsInstance.VanityUrlDetails[vanityUrlKey]
	s.Res = &vanityUrlValue

	return nil
}

func (s *AnalyticsAnalyticsInstanceVanityUrlResourceCrud) Update() error {
	request := oci_analytics.UpdateVanityUrlRequest{}

	if analyticsInstanceId, ok := s.D.GetOkExists("analytics_instance_id"); ok {
		tmp := analyticsInstanceId.(string)
		request.AnalyticsInstanceId = &tmp
	}

	if caCertificate, ok := s.D.GetOkExists("ca_certificate"); ok {
		tmp := caCertificate.(string)
		request.CaCertificate = &tmp
	}

	if passphrase, ok := s.D.GetOkExists("passphrase"); ok {
		tmp := passphrase.(string)
		request.Passphrase = &tmp
	}

	if privateKey, ok := s.D.GetOkExists("private_key"); ok {
		tmp := privateKey.(string)
		request.PrivateKey = &tmp
	}

	if publicCertificate, ok := s.D.GetOkExists("public_certificate"); ok {
		tmp := publicCertificate.(string)
		request.PublicCertificate = &tmp
	}

	_, vanityUrlKey, err := parseAnalyticsInstanceVanityUrlCompositeId(s.D.Id())
	if err == nil {
		request.VanityUrlKey = &vanityUrlKey
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "analytics")

	response, err := s.Client.UpdateVanityUrl(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getAnalyticsInstanceVanityUrlFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "analytics"), oci_analytics.WorkRequestActionResultVanityUrlUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *AnalyticsAnalyticsInstanceVanityUrlResourceCrud) Delete() error {
	request := oci_analytics.DeleteVanityUrlRequest{}

	if analyticsInstanceId, ok := s.D.GetOkExists("analytics_instance_id"); ok {
		tmp := analyticsInstanceId.(string)
		request.AnalyticsInstanceId = &tmp
	}

	_, vanityUrlKey, err := parseAnalyticsInstanceVanityUrlCompositeId(s.D.Id())
	if err == nil {
		request.VanityUrlKey = &vanityUrlKey
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "analytics")

	response, err := s.Client.DeleteVanityUrl(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := analyticsInstanceVanityUrlWaitForWorkRequest(workId, "analytics",
		oci_analytics.WorkRequestActionResultVanityUrlDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *AnalyticsAnalyticsInstanceVanityUrlResourceCrud) SetData() error {
	analyticsInstanceId, _, err := parseAnalyticsInstanceVanityUrlCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("analytics_instance_id", &analyticsInstanceId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}
	s.D.Set("hosts", s.Res.Hosts)

	s.D.Set("urls", s.Res.Urls)

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.PublicCertificate != nil {
		s.D.Set("public_certificate", *s.Res.PublicCertificate)
	}

	return nil
}

func getAnalyticsInstanceVanityUrlCompositeId(analyticsInstanceId string, vanityUrlKey string) string {
	analyticsInstanceId = url.PathEscape(analyticsInstanceId)
	vanityUrlKey = url.PathEscape(vanityUrlKey)
	compositeId := "analyticsInstances/" + analyticsInstanceId + "/vanityUrls/" + vanityUrlKey
	return compositeId
}

func parseAnalyticsInstanceVanityUrlCompositeId(compositeId string) (analyticsInstanceId string, vanityUrlKey string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("analyticsInstances/.*/vanityUrls/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	analyticsInstanceId, _ = url.PathUnescape(parts[1])
	vanityUrlKey, _ = url.PathUnescape(parts[3])

	return
}
