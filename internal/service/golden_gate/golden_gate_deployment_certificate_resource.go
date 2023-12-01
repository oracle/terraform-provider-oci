// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package golden_gate

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_golden_gate "github.com/oracle/oci-go-sdk/v65/goldengate"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func GoldenGateDeploymentCertificateResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createGoldenGateDeploymentCertificate,
		Read:     readGoldenGateDeploymentCertificate,
		Delete:   deleteGoldenGateDeploymentCertificate,
		Schema: map[string]*schema.Schema{
			// Required
			"certificate_content": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"deployment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"key": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional

			// Computed
			"authority_key_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_ca": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_self_signed": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"issuer": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"md5hash": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"public_key": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"public_key_algorithm": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"public_key_size": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"serial": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sha1hash": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"subject": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"subject_key_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_valid_from": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_valid_to": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"version": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createGoldenGateDeploymentCertificate(d *schema.ResourceData, m interface{}) error {
	sync := &GoldenGateDeploymentCertificateResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GoldenGateClient()

	return tfresource.CreateResource(d, sync)
}

func readGoldenGateDeploymentCertificate(d *schema.ResourceData, m interface{}) error {
	sync := &GoldenGateDeploymentCertificateResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GoldenGateClient()

	return tfresource.ReadResource(sync)
}

func deleteGoldenGateDeploymentCertificate(d *schema.ResourceData, m interface{}) error {
	sync := &GoldenGateDeploymentCertificateResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GoldenGateClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type GoldenGateDeploymentCertificateResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_golden_gate.GoldenGateClient
	Res                    *oci_golden_gate.Certificate
	DisableNotFoundRetries bool
}

func (s *GoldenGateDeploymentCertificateResourceCrud) ID() string {
	return GetDeploymentCertificateCompositeId(s.D.Get("key").(string), s.D.Get("deployment_id").(string))
}

func GetDeploymentCertificateCompositeId(certificateKey string, deploymentId string) string {
	certificateKey = url.PathEscape(certificateKey)
	deploymentId = url.PathEscape(deploymentId)
	compositeId := "deployments/" + deploymentId + "/certificates/" + certificateKey
	return compositeId
}

func parseDeploymentCertificateCompositeId(compositeId string) (certificateKey string, deploymentId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("deployments/.*/certificates/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	deploymentId, _ = url.PathUnescape(parts[1])
	certificateKey, _ = url.PathUnescape(parts[3])

	return
}

func (s *GoldenGateDeploymentCertificateResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_golden_gate.CertificateLifecycleStateCreating),
	}
}

func (s *GoldenGateDeploymentCertificateResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_golden_gate.CertificateLifecycleStateActive),
		string(oci_golden_gate.CertificateLifecycleStateFailed),
	}
}

func (s *GoldenGateDeploymentCertificateResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_golden_gate.CertificateLifecycleStateDeleting),
	}
}

func (s *GoldenGateDeploymentCertificateResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_golden_gate.CertificateLifecycleStateDeleted),
	}
}

func (s *GoldenGateDeploymentCertificateResourceCrud) Create() error {
	request := oci_golden_gate.CreateCertificateRequest{}

	if certificateContent, ok := s.D.GetOkExists("certificate_content"); ok {
		tmp := certificateContent.(string)
		request.CertificateContent = &tmp
	}

	if deploymentId, ok := s.D.GetOkExists("deployment_id"); ok {
		tmp := deploymentId.(string)
		request.DeploymentId = &tmp
	}

	if key, ok := s.D.GetOkExists("key"); ok {
		tmp := key.(string)
		request.Key = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "golden_gate")

	response, err := s.Client.CreateCertificate(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getDeploymentCertificateFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "golden_gate"), oci_golden_gate.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *GoldenGateDeploymentCertificateResourceCrud) getDeploymentCertificateFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_golden_gate.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	certificateKey, deploymentId, err := deploymentCertificateWaitForWorkRequest(workId,
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(GetDeploymentCertificateCompositeId(*certificateKey, *deploymentId))

	return s.Get()
}

func deploymentCertificateWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "golden_gate", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_golden_gate.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func deploymentCertificateWaitForWorkRequest(wId *string, action oci_golden_gate.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_golden_gate.GoldenGateClient) (*string, *string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "golden_gate")
	retryPolicy.ShouldRetryOperation = deploymentCertificateWorkRequestShouldRetryFunc(timeout)

	response := oci_golden_gate.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_golden_gate.OperationStatusInProgress),
			string(oci_golden_gate.OperationStatusAccepted),
		},
		Target: []string{
			string(oci_golden_gate.OperationStatusSucceeded),
			string(oci_golden_gate.OperationStatusFailed),
			string(oci_golden_gate.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_golden_gate.GetWorkRequestRequest{
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
		return nil, nil, e
	}

	var certificateKey *string
	var deploymentId *string
	// The work request response contains an array of objects that finished the operation
	for _, res := range response.Resources {
		if strings.Contains(strings.ToLower(*res.EntityType), "certificate") {
			if res.ActionType == action {
				certificateKey = res.Identifier
			}
		}
		if strings.Contains(strings.ToLower(*res.EntityType), "deployment") {
			if res.ActionType == action {
				deploymentId = res.Identifier
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if certificateKey == nil || response.Status == oci_golden_gate.OperationStatusFailed || response.Status == oci_golden_gate.OperationStatusCanceled {
		return nil, nil, getErrorFromGoldenGateDeploymentCertificateWorkRequest(client, wId, retryPolicy, "certificate", action)
	}

	return certificateKey, deploymentId, nil
}

func getErrorFromGoldenGateDeploymentCertificateWorkRequest(client *oci_golden_gate.GoldenGateClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_golden_gate.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_golden_gate.ListWorkRequestErrorsRequest{
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

func (s *GoldenGateDeploymentCertificateResourceCrud) Get() error {
	request := oci_golden_gate.GetCertificateRequest{}

	certificateKey, deploymentId, err := parseDeploymentCertificateCompositeId(s.D.Id())
	if err == nil {
		request.CertificateKey = &certificateKey
		request.DeploymentId = &deploymentId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "golden_gate")

	response, err := s.Client.GetCertificate(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Certificate
	return nil
}

func (s *GoldenGateDeploymentCertificateResourceCrud) Delete() error {
	request := oci_golden_gate.DeleteCertificateRequest{}

	if key, ok := s.D.GetOkExists("key"); ok {
		tmp := key.(string)
		request.CertificateKey = &tmp
	}

	if deploymentId, ok := s.D.GetOkExists("deployment_id"); ok {
		tmp := deploymentId.(string)
		request.DeploymentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "golden_gate")

	response, err := s.Client.DeleteCertificate(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, _, delWorkRequestErr := deploymentCertificateWaitForWorkRequest(workId,
		oci_golden_gate.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *GoldenGateDeploymentCertificateResourceCrud) SetData() error {

	if s.Res.AuthorityKeyId != nil {
		s.D.Set("authority_key_id", *s.Res.AuthorityKeyId)
	}

	if s.Res.CertificateContent != nil {
		s.D.Set("certificate_content", *s.Res.CertificateContent)
	}

	if s.Res.DeploymentId != nil {
		s.D.Set("deployment_id", *s.Res.DeploymentId)
	}

	if s.Res.IsCa != nil {
		s.D.Set("is_ca", *s.Res.IsCa)
	}

	if s.Res.IsSelfSigned != nil {
		s.D.Set("is_self_signed", *s.Res.IsSelfSigned)
	}

	if s.Res.Issuer != nil {
		s.D.Set("issuer", *s.Res.Issuer)
	}

	if s.Res.Key != nil {
		s.D.Set("key", *s.Res.Key)
	}

	if s.Res.Md5Hash != nil {
		s.D.Set("md5hash", *s.Res.Md5Hash)
	}

	if s.Res.PublicKey != nil {
		s.D.Set("public_key", *s.Res.PublicKey)
	}

	if s.Res.PublicKeyAlgorithm != nil {
		s.D.Set("public_key_algorithm", *s.Res.PublicKeyAlgorithm)
	}

	if s.Res.PublicKeySize != nil {
		s.D.Set("public_key_size", strconv.FormatInt(*s.Res.PublicKeySize, 10))
	}

	if s.Res.Serial != nil {
		s.D.Set("serial", *s.Res.Serial)
	}

	if s.Res.Sha1Hash != nil {
		s.D.Set("sha1hash", *s.Res.Sha1Hash)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.Subject != nil {
		s.D.Set("subject", *s.Res.Subject)
	}

	if s.Res.SubjectKeyId != nil {
		s.D.Set("subject_key_id", *s.Res.SubjectKeyId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeValidFrom != nil {
		s.D.Set("time_valid_from", s.Res.TimeValidFrom.String())
	}

	if s.Res.TimeValidTo != nil {
		s.D.Set("time_valid_to", s.Res.TimeValidTo.String())
	}

	if s.Res.Version != nil {
		s.D.Set("version", *s.Res.Version)
	}

	return nil
}

func CertificateSummaryToMap(obj oci_golden_gate.CertificateSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsSelfSigned != nil {
		result["is_self_signed"] = bool(*obj.IsSelfSigned)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.Subject != nil {
		result["subject"] = string(*obj.Subject)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeValidTo != nil {
		result["time_valid_to"] = obj.TimeValidTo.String()
	}

	return result
}
