// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package bds

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_bds "github.com/oracle/oci-go-sdk/v65/bds"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func BdsBdsInstanceOperationCertificateManagementsManagementResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		Create:   createBdsBdsInstanceOperationCertificateManagementsManagement,
		Read:     readBdsBdsInstanceOperationCertificateManagementsManagement,
		Update:   updateBdsBdsInstanceOperationCertificateManagementsManagement,
		Delete:   deleteBdsBdsInstanceOperationCertificateManagementsManagement,
		Schema: map[string]*schema.Schema{
			// Required
			"bds_instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cluster_admin_password": {
				Type:      schema.TypeString,
				Required:  true,
				ForceNew:  true,
				Sensitive: true,
			},
			"services": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"enable_operation_certificate_management": {
				Type:     schema.TypeBool,
				Required: true,
			},
			// Required
			"renew_operation_certificate_management": {
				Type:     schema.TypeBool,
				Required: true,
			},
			// Optional
			"host_cert_details": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Optional
						"certificate": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"host_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"private_key": {
							Type:      schema.TypeString,
							Optional:  true,
							Computed:  true,
							ForceNew:  true,
							Sensitive: true,
						},

						// Computed
					},
				},
			},
			"root_certificate": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"server_key_password": {
				Type:      schema.TypeString,
				Optional:  true,
				Computed:  true,
				ForceNew:  true,
				Sensitive: true,
			},

			// Computed
		},
	}
}

func createBdsBdsInstanceOperationCertificateManagementsManagement(d *schema.ResourceData, m interface{}) error {
	sync := &BdsBdsInstanceOperationCertificateManagementsManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()
	sync.Res = &BdsBdsInstanceOperationCertificateManagementsManagementResponse{}

	return tfresource.CreateResource(d, sync)
}

func readBdsBdsInstanceOperationCertificateManagementsManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

func updateBdsBdsInstanceOperationCertificateManagementsManagement(d *schema.ResourceData, m interface{}) error {
	sync := &BdsBdsInstanceOperationCertificateManagementsManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()
	sync.Res = &BdsBdsInstanceOperationCertificateManagementsManagementResponse{}

	return tfresource.UpdateResource(d, sync)
}

func deleteBdsBdsInstanceOperationCertificateManagementsManagement(d *schema.ResourceData, m interface{}) error {
	sync := &BdsBdsInstanceOperationCertificateManagementsManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()
	sync.Res = &BdsBdsInstanceOperationCertificateManagementsManagementResponse{}
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type BdsBdsInstanceOperationCertificateManagementsManagementResponse struct {
	enableResponse  *oci_bds.EnableCertificateResponse
	renewResponse   *oci_bds.RenewCertificateResponse
	disableResponse *oci_bds.DisableCertificateResponse
}

type BdsBdsInstanceOperationCertificateManagementsManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_bds.BdsClient
	Res                    *BdsBdsInstanceOperationCertificateManagementsManagementResponse
	DisableNotFoundRetries bool
}

func (s *BdsBdsInstanceOperationCertificateManagementsManagementResourceCrud) ID() string {
	return tfresource.GenerateDataSourceHashID("BdsBdsInstanceOperationCertificateManagementsManagementResource-", BdsBdsInstanceOperationCertificateManagementsManagementResource(), s.D)
}

func (s *BdsBdsInstanceOperationCertificateManagementsManagementResourceCrud) Create() error {
	var operationEnable bool
	var operationRenew bool
	if enableOperation, ok := s.D.GetOkExists("enable_operation_certificate_management"); ok {
		operationEnable = enableOperation.(bool)
	}
	if renewOperation, ok := s.D.GetOkExists("renew_operation_certificate_management"); ok {
		operationRenew = renewOperation.(bool)
	}

	if operationEnable && operationRenew {
		return fmt.Errorf("Both enable_operation_certificate_management and renew_operation_certificate_management can't be true")
	}

	if operationEnable {
		request := oci_bds.EnableCertificateRequest{}

		if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
			tmp := bdsInstanceId.(string)
			request.BdsInstanceId = &tmp
		}

		if clusterAdminPassword, ok := s.D.GetOkExists("cluster_admin_password"); ok {
			tmp := clusterAdminPassword.(string)
			request.ClusterAdminPassword = &tmp
		}

		if hostCertDetails, ok := s.D.GetOkExists("host_cert_details"); ok {
			interfaces := hostCertDetails.([]interface{})
			tmp := make([]oci_bds.HostCertDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "host_cert_details", stateDataIndex)
				converted, err := s.mapToHostCertDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("host_cert_details") {
				request.HostCertDetails = tmp
			}
		} else {
			request.HostCertDetails = []oci_bds.HostCertDetails{}
		}

		if rootCertificate, ok := s.D.GetOkExists("root_certificate"); ok {
			tmp := rootCertificate.(string)
			request.RootCertificate = &tmp
		} else {
			tmp := ""
			request.RootCertificate = &tmp
		}

		if serverKeyPassword, ok := s.D.GetOkExists("server_key_password"); ok {
			tmp := serverKeyPassword.(string)
			request.ServerKeyPassword = &tmp
		} else {
			tmp := ""
			request.ServerKeyPassword = &tmp
		}

		if services, ok := s.D.GetOkExists("services"); ok {
			interfaces := services.([]interface{})
			tmp := make([]oci_bds.ServiceEnum, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = oci_bds.ServiceEnum(interfaces[i].(string))
				}
			}
			if len(tmp) != 0 || s.D.HasChange("services") {
				request.Services = tmp
			}
		}

		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

		response, err := s.Client.EnableCertificate(context.Background(), request)
		if err != nil {
			return err
		}

		workId := response.OpcWorkRequestId
		err = s.getBdsInstanceOperationCertificateManagementsManagementFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds"), oci_bds.ActionTypesUpdated, s.D.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return err
		}
		s.Res.enableResponse = &response
		return nil
	}

	if operationRenew {
		request := oci_bds.RenewCertificateRequest{}

		if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
			tmp := bdsInstanceId.(string)
			request.BdsInstanceId = &tmp
		}

		if clusterAdminPassword, ok := s.D.GetOkExists("cluster_admin_password"); ok {
			tmp := clusterAdminPassword.(string)
			request.ClusterAdminPassword = &tmp
		}

		if hostCertDetails, ok := s.D.GetOkExists("host_cert_details"); ok {
			interfaces := hostCertDetails.([]interface{})
			tmp := make([]oci_bds.HostCertDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "host_cert_details", stateDataIndex)
				converted, err := s.mapToHostCertDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("host_cert_details") {
				request.HostCertDetails = tmp
			}
		} else {
			request.HostCertDetails = []oci_bds.HostCertDetails{}
		}

		if rootCertificate, ok := s.D.GetOkExists("root_certificate"); ok {
			tmp := rootCertificate.(string)
			request.RootCertificate = &tmp
		} else {
			tmp := ""
			request.RootCertificate = &tmp
		}

		if serverKeyPassword, ok := s.D.GetOkExists("server_key_password"); ok {
			tmp := serverKeyPassword.(string)
			request.ServerKeyPassword = &tmp
		} else {
			tmp := ""
			request.ServerKeyPassword = &tmp
		}

		if services, ok := s.D.GetOkExists("services"); ok {
			interfaces := services.([]interface{})
			tmp := make([]oci_bds.ServiceEnum, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = oci_bds.ServiceEnum(interfaces[i].(string))
				}
			}
			if len(tmp) != 0 || s.D.HasChange("services") {
				request.Services = tmp
			}
		}

		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

		response, err := s.Client.RenewCertificate(context.Background(), request)
		if err != nil {
			return err
		}

		workId := response.OpcWorkRequestId
		err = s.getBdsInstanceOperationCertificateManagementsManagementFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds"), oci_bds.ActionTypesUpdated, s.D.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return err
		}
		s.Res.renewResponse = &response
		return nil
	}

	request := oci_bds.DisableCertificateRequest{}

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	if clusterAdminPassword, ok := s.D.GetOkExists("cluster_admin_password"); ok {
		tmp := clusterAdminPassword.(string)
		request.ClusterAdminPassword = &tmp
	}

	if services, ok := s.D.GetOkExists("services"); ok {
		interfaces := services.([]interface{})
		tmp := make([]oci_bds.ServiceEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_bds.ServiceEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("services") {
			request.Services = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.DisableCertificate(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	err = s.getBdsInstanceOperationCertificateManagementsManagementFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds"), oci_bds.ActionTypesUpdated, s.D.Timeout(schema.TimeoutUpdate))
	if err != nil {
		return err
	}
	s.Res.disableResponse = &response
	return nil
}

func (s *BdsBdsInstanceOperationCertificateManagementsManagementResourceCrud) getBdsInstanceOperationCertificateManagementsManagementFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_bds.ActionTypesEnum, timeout time.Duration) error {

	// Wait until it finishes
	_, err := bdsInstanceOperationCertificateManagementsManagementWaitForWorkRequest(workId, "bds",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}

	return nil
}

func bdsInstanceOperationCertificateManagementsManagementWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "bds", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_bds.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func bdsInstanceOperationCertificateManagementsManagementWaitForWorkRequest(wId *string, entityType string, action oci_bds.ActionTypesEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_bds.BdsClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "bds")
	retryPolicy.ShouldRetryOperation = bdsInstanceOperationCertificateManagementsManagementWorkRequestShouldRetryFunc(timeout)

	response := oci_bds.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_bds.OperationStatusInProgress),
			string(oci_bds.OperationStatusAccepted),
			string(oci_bds.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_bds.OperationStatusSucceeded),
			string(oci_bds.OperationStatusFailed),
			string(oci_bds.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_bds.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_bds.OperationStatusFailed || response.Status == oci_bds.OperationStatusCanceled {
		return nil, getErrorFromBdsBdsInstanceOperationCertificateManagementsManagementWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromBdsBdsInstanceOperationCertificateManagementsManagementWorkRequest(client *oci_bds.BdsClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_bds.ActionTypesEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_bds.ListWorkRequestErrorsRequest{
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

func (s *BdsBdsInstanceOperationCertificateManagementsManagementResourceCrud) Update() error {
	var operationEnable bool
	var operationRenew bool
	if enableOperation, ok := s.D.GetOkExists("enable_operation_certificate_management"); ok {
		operationEnable = enableOperation.(bool)
	}
	if renewOperation, ok := s.D.GetOkExists("renew_operation_certificate_management"); ok {
		operationRenew = renewOperation.(bool)
	}

	if operationEnable && operationRenew {
		return fmt.Errorf("Both enable_operation_certificate_management and renew_operation_certificate_management can't be true")
	}

	if operationEnable {
		request := oci_bds.EnableCertificateRequest{}

		if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
			tmp := bdsInstanceId.(string)
			request.BdsInstanceId = &tmp
		}

		if clusterAdminPassword, ok := s.D.GetOkExists("cluster_admin_password"); ok {
			tmp := clusterAdminPassword.(string)
			request.ClusterAdminPassword = &tmp
		}

		if hostCertDetails, ok := s.D.GetOkExists("host_cert_details"); ok {
			interfaces := hostCertDetails.([]interface{})
			tmp := make([]oci_bds.HostCertDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "host_cert_details", stateDataIndex)
				converted, err := s.mapToHostCertDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("host_cert_details") {
				request.HostCertDetails = tmp
			}
		} else {
			request.HostCertDetails = []oci_bds.HostCertDetails{}
		}

		if rootCertificate, ok := s.D.GetOkExists("root_certificate"); ok {
			tmp := rootCertificate.(string)
			request.RootCertificate = &tmp
		} else {
			tmp := ""
			request.RootCertificate = &tmp
		}

		if serverKeyPassword, ok := s.D.GetOkExists("server_key_password"); ok {
			tmp := serverKeyPassword.(string)
			request.ServerKeyPassword = &tmp
		} else {
			tmp := ""
			request.ServerKeyPassword = &tmp
		}

		if services, ok := s.D.GetOkExists("services"); ok {
			interfaces := services.([]interface{})
			tmp := make([]oci_bds.ServiceEnum, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = oci_bds.ServiceEnum(interfaces[i].(string))
				}
			}
			if len(tmp) != 0 || s.D.HasChange("services") {
				request.Services = tmp
			}
		}

		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

		response, err := s.Client.EnableCertificate(context.Background(), request)
		if err != nil {
			return err
		}

		workId := response.OpcWorkRequestId
		err = s.getBdsInstanceOperationCertificateManagementsManagementFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds"), oci_bds.ActionTypesUpdated, s.D.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return err
		}
		s.Res.enableResponse = &response
		return nil
	}

	if operationRenew {
		request := oci_bds.RenewCertificateRequest{}

		if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
			tmp := bdsInstanceId.(string)
			request.BdsInstanceId = &tmp
		}

		if clusterAdminPassword, ok := s.D.GetOkExists("cluster_admin_password"); ok {
			tmp := clusterAdminPassword.(string)
			request.ClusterAdminPassword = &tmp
		}

		if hostCertDetails, ok := s.D.GetOkExists("host_cert_details"); ok {
			interfaces := hostCertDetails.([]interface{})
			tmp := make([]oci_bds.HostCertDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "host_cert_details", stateDataIndex)
				converted, err := s.mapToHostCertDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("host_cert_details") {
				request.HostCertDetails = tmp
			}
		} else {
			request.HostCertDetails = []oci_bds.HostCertDetails{}
		}

		if rootCertificate, ok := s.D.GetOkExists("root_certificate"); ok {
			tmp := rootCertificate.(string)
			request.RootCertificate = &tmp
		} else {
			tmp := ""
			request.RootCertificate = &tmp
		}

		if serverKeyPassword, ok := s.D.GetOkExists("server_key_password"); ok {
			tmp := serverKeyPassword.(string)
			request.ServerKeyPassword = &tmp
		} else {
			tmp := ""
			request.ServerKeyPassword = &tmp
		}

		if services, ok := s.D.GetOkExists("services"); ok {
			interfaces := services.([]interface{})
			tmp := make([]oci_bds.ServiceEnum, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = oci_bds.ServiceEnum(interfaces[i].(string))
				}
			}
			if len(tmp) != 0 || s.D.HasChange("services") {
				request.Services = tmp
			}
		}

		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

		response, err := s.Client.RenewCertificate(context.Background(), request)
		if err != nil {
			return err
		}

		workId := response.OpcWorkRequestId
		err = s.getBdsInstanceOperationCertificateManagementsManagementFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds"), oci_bds.ActionTypesUpdated, s.D.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return err
		}
		s.Res.renewResponse = &response
		return nil
	}

	request := oci_bds.DisableCertificateRequest{}

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	if clusterAdminPassword, ok := s.D.GetOkExists("cluster_admin_password"); ok {
		tmp := clusterAdminPassword.(string)
		request.ClusterAdminPassword = &tmp
	}

	if services, ok := s.D.GetOkExists("services"); ok {
		interfaces := services.([]interface{})
		tmp := make([]oci_bds.ServiceEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_bds.ServiceEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("services") {
			request.Services = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.DisableCertificate(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	err = s.getBdsInstanceOperationCertificateManagementsManagementFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds"), oci_bds.ActionTypesUpdated, s.D.Timeout(schema.TimeoutUpdate))
	if err != nil {
		return err
	}
	s.Res.disableResponse = &response
	return nil
}

func (s *BdsBdsInstanceOperationCertificateManagementsManagementResourceCrud) Delete() error {
	var operation bool
	if enableOperation, ok := s.D.GetOkExists("enable_operation_certificate_management"); ok {
		operation = enableOperation.(bool)
	}

	if !operation {
		return nil
	}

	request := oci_bds.DisableCertificateRequest{}

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	if clusterAdminPassword, ok := s.D.GetOkExists("cluster_admin_password"); ok {
		tmp := clusterAdminPassword.(string)
		request.ClusterAdminPassword = &tmp
	}

	if services, ok := s.D.GetOkExists("services"); ok {
		interfaces := services.([]interface{})
		tmp := make([]oci_bds.ServiceEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_bds.ServiceEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("services") {
			request.Services = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.DisableCertificate(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	err = s.getBdsInstanceOperationCertificateManagementsManagementFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds"), oci_bds.ActionTypesUpdated, s.D.Timeout(schema.TimeoutUpdate))
	if err != nil {
		return err
	}

	s.Res.disableResponse = &response
	return nil
}

func (s *BdsBdsInstanceOperationCertificateManagementsManagementResourceCrud) SetData() error {
	return nil
}

func (s *BdsBdsInstanceOperationCertificateManagementsManagementResourceCrud) mapToHostCertDetails(fieldKeyFormat string) (oci_bds.HostCertDetails, error) {
	result := oci_bds.HostCertDetails{}

	if certificate, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "certificate")); ok {
		tmp := certificate.(string)
		result.Certificate = &tmp
	}

	if hostName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "host_name")); ok {
		tmp := hostName.(string)
		result.HostName = &tmp
	}

	if privateKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "private_key")); ok {
		tmp := privateKey.(string)
		result.PrivateKey = &tmp
	}

	return result, nil
}

func HostCertDetailsToMap(obj oci_bds.HostCertDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Certificate != nil {
		result["certificate"] = string(*obj.Certificate)
	}

	if obj.HostName != nil {
		result["host_name"] = string(*obj.HostName)
	}

	if obj.PrivateKey != nil {
		result["private_key"] = string(*obj.PrivateKey)
	}

	return result
}
