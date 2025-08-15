// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dbmulticloud

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_dbmulticloud "github.com/oracle/oci-go-sdk/v65/dbmulticloud"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DbmulticloudOracleDbAzureConnectorResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDbmulticloudOracleDbAzureConnector,
		Read:     readDbmulticloudOracleDbAzureConnector,
		Update:   updateDbmulticloudOracleDbAzureConnector,
		Delete:   deleteDbmulticloudOracleDbAzureConnector,
		Schema: map[string]*schema.Schema{
			// Required
			"azure_identity_mechanism": {
				Type:     schema.TypeString,
				Required: true,
			},
			"azure_resource_group": {
				Type:     schema.TypeString,
				Required: true,
			},
			"azure_subscription_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"azure_tenant_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"db_cluster_resource_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"access_token": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Computed
			"arc_agent_nodes": {
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"current_arc_agent_version": {
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
						},
						"host_id": {
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
						},
						"host_name": {
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
						},
						"time_last_checked": {
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
						},
					},
				},
			},
			"azure_identity_connectivity_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_modification": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"lifecycle_state_details": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"state": {
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

func createDbmulticloudOracleDbAzureConnector(d *schema.ResourceData, m interface{}) error {
	sync := &DbmulticloudOracleDbAzureConnectorResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OracleDBAzureConnectorClient()
	sync.WorkRequestClient = m.(*client.OracleClients).DbmulticloudWorkRequestClient()

	return tfresource.CreateResource(d, sync)
}

func readDbmulticloudOracleDbAzureConnector(d *schema.ResourceData, m interface{}) error {
	sync := &DbmulticloudOracleDbAzureConnectorResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OracleDBAzureConnectorClient()

	return tfresource.ReadResource(sync)
}

func updateDbmulticloudOracleDbAzureConnector(d *schema.ResourceData, m interface{}) error {
	sync := &DbmulticloudOracleDbAzureConnectorResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OracleDBAzureConnectorClient()
	sync.WorkRequestClient = m.(*client.OracleClients).DbmulticloudWorkRequestClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDbmulticloudOracleDbAzureConnector(d *schema.ResourceData, m interface{}) error {
	sync := &DbmulticloudOracleDbAzureConnectorResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OracleDBAzureConnectorClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).DbmulticloudWorkRequestClient()

	return tfresource.DeleteResource(d, sync)
}

type DbmulticloudOracleDbAzureConnectorResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_dbmulticloud.OracleDBAzureConnectorClient
	Res                    *oci_dbmulticloud.OracleDbAzureConnector
	PatchResponse          *oci_dbmulticloud.OracleDbAzureConnector
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_dbmulticloud.WorkRequestClient
}

func (s *DbmulticloudOracleDbAzureConnectorResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DbmulticloudOracleDbAzureConnectorResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_dbmulticloud.OracleDbAzureConnectorLifecycleStateCreating),
	}
}

func (s *DbmulticloudOracleDbAzureConnectorResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_dbmulticloud.OracleDbAzureConnectorLifecycleStateActive),
	}
}

func (s *DbmulticloudOracleDbAzureConnectorResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_dbmulticloud.OracleDbAzureConnectorLifecycleStateDeleting),
	}
}

func (s *DbmulticloudOracleDbAzureConnectorResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_dbmulticloud.OracleDbAzureConnectorLifecycleStateDeleted),
	}
}

func (s *DbmulticloudOracleDbAzureConnectorResourceCrud) Create() error {
	request := oci_dbmulticloud.CreateOracleDbAzureConnectorRequest{}

	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> INTO DbmulticloudOracleDbAzureConnectorResourceCrud) Create()  >>>>>>>>>>")

	if accessToken, ok := s.D.GetOkExists("access_token"); ok {
		tmp := accessToken.(string)
		request.AccessToken = &tmp
	}

	if azureIdentityMechanism, ok := s.D.GetOkExists("azure_identity_mechanism"); ok {
		request.AzureIdentityMechanism = oci_dbmulticloud.OracleDbAzureConnectorAzureIdentityMechanismEnum(azureIdentityMechanism.(string))
	}

	if azureResourceGroup, ok := s.D.GetOkExists("azure_resource_group"); ok {
		tmp := azureResourceGroup.(string)
		request.AzureResourceGroup = &tmp
	}

	if azureSubscriptionId, ok := s.D.GetOkExists("azure_subscription_id"); ok {
		tmp := azureSubscriptionId.(string)
		request.AzureSubscriptionId = &tmp
	}

	if azureTenantId, ok := s.D.GetOkExists("azure_tenant_id"); ok {
		tmp := azureTenantId.(string)
		request.AzureTenantId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if dbClusterResourceId, ok := s.D.GetOkExists("db_cluster_resource_id"); ok {
		tmp := dbClusterResourceId.(string)
		request.DbClusterResourceId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if privateEndpointDnsAlias, ok := s.D.GetOkExists("private_endpoint_dns_alias"); ok {
		tmp := privateEndpointDnsAlias.(string)
		request.PrivateEndpointDnsAlias = &tmp
	}

	if privateEndpointIpAddress, ok := s.D.GetOkExists("private_endpoint_ip_address"); ok {
		tmp := privateEndpointIpAddress.(string)
		request.PrivateEndpointIpAddress = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dbmulticloud")

	response, err := s.Client.CreateOracleDbAzureConnector(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}

	err = s.getOracleDbAzureConnectorFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dbmulticloud"), oci_dbmulticloud.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
	if err != nil {
		return err
	}
	return nil
}

func (s *DbmulticloudOracleDbAzureConnectorResourceCrud) getOracleDbAzureConnectorFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_dbmulticloud.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	oracleDbAzureConnectorId, err := oracleDbAzureConnectorWaitForWorkRequest(workId, "oracledbazureconnector",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.WorkRequestClient)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, oracleDbAzureConnectorId)
		_, cancelErr := s.WorkRequestClient.CancelWorkRequest(context.Background(),
			oci_dbmulticloud.CancelWorkRequestRequest{
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
	s.D.SetId(*oracleDbAzureConnectorId)

	return s.Get()
}

func oracleDbAzureConnectorWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "dbmulticloud", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_dbmulticloud.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func oracleDbAzureConnectorWaitForWorkRequest(wId *string, entityType string, action oci_dbmulticloud.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_dbmulticloud.WorkRequestClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "dbmulticloud")
	retryPolicy.ShouldRetryOperation = oracleDbAzureConnectorWorkRequestShouldRetryFunc(timeout)

	response := oci_dbmulticloud.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_dbmulticloud.OperationStatusInProgress),
			string(oci_dbmulticloud.OperationStatusAccepted),
			string(oci_dbmulticloud.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_dbmulticloud.OperationStatusSucceeded),
			string(oci_dbmulticloud.OperationStatusFailed),
			string(oci_dbmulticloud.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_dbmulticloud.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_dbmulticloud.OperationStatusFailed || response.Status == oci_dbmulticloud.OperationStatusCanceled {
		return nil, getErrorFromDbmulticloudOracleDbAzureConnectorWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDbmulticloudOracleDbAzureConnectorWorkRequest(client *oci_dbmulticloud.WorkRequestClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_dbmulticloud.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_dbmulticloud.ListWorkRequestErrorsRequest{
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

func (s *DbmulticloudOracleDbAzureConnectorResourceCrud) Get() error {
	request := oci_dbmulticloud.GetOracleDbAzureConnectorRequest{}

	tmp := s.D.Id()
	request.OracleDbAzureConnectorId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dbmulticloud")

	response, err := s.Client.GetOracleDbAzureConnector(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.OracleDbAzureConnector
	return nil
}

func (s *DbmulticloudOracleDbAzureConnectorResourceCrud) Update() error {

	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_dbmulticloud.UpdateOracleDbAzureConnectorRequest{}

	if accessToken, ok := s.D.GetOkExists("access_token"); ok {
		tmp := accessToken.(string)
		request.AccessToken = &tmp
	}

	if azureIdentityMechanism, ok := s.D.GetOkExists("azure_identity_mechanism"); ok {
		request.AzureIdentityMechanism = oci_dbmulticloud.OracleDbAzureConnectorAzureIdentityMechanismEnum(azureIdentityMechanism.(string))
	}

	if azureResourceGroup, ok := s.D.GetOkExists("azure_resource_group"); ok {
		tmp := azureResourceGroup.(string)
		request.AzureResourceGroup = &tmp
	}

	if azureSubscriptionId, ok := s.D.GetOkExists("azure_subscription_id"); ok {
		tmp := azureSubscriptionId.(string)
		request.AzureSubscriptionId = &tmp
	}

	if azureTenantId, ok := s.D.GetOkExists("azure_tenant_id"); ok {
		tmp := azureTenantId.(string)
		request.AzureTenantId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if dbClusterResourceId, ok := s.D.GetOkExists("db_cluster_resource_id"); ok {
		tmp := dbClusterResourceId.(string)
		request.DbClusterResourceId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	tmp := s.D.Id()
	request.OracleDbAzureConnectorId = &tmp

	if privateEndpointDnsAlias, ok := s.D.GetOkExists("private_endpoint_dns_alias"); ok {
		tmp := privateEndpointDnsAlias.(string)
		request.PrivateEndpointDnsAlias = &tmp
	}

	if privateEndpointIpAddress, ok := s.D.GetOkExists("private_endpoint_ip_address"); ok {
		tmp := privateEndpointIpAddress.(string)
		request.PrivateEndpointIpAddress = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dbmulticloud")

	response, err := s.Client.UpdateOracleDbAzureConnector(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	err = s.getOracleDbAzureConnectorFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dbmulticloud"), oci_dbmulticloud.ActionTypeCreated, s.D.Timeout(schema.TimeoutUpdate))
	if err != nil {
		return err
	}
	return nil
}

func (s *DbmulticloudOracleDbAzureConnectorResourceCrud) Delete() error {
	request := oci_dbmulticloud.DeleteOracleDbAzureConnectorRequest{}

	tmp := s.D.Id()
	request.OracleDbAzureConnectorId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dbmulticloud")

	response, err := s.Client.DeleteOracleDbAzureConnector(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := oracleDbAzureConnectorWaitForWorkRequest(workId, "oracledbazureconnector",
		oci_dbmulticloud.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.WorkRequestClient)
	return delWorkRequestErr
}

func (s *DbmulticloudOracleDbAzureConnectorResourceCrud) SetData() error {

	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> INTO *DbmulticloudOracleDbAzureConnectorResourceCrud) SetData()  >>>>>>>>>>")
	if s.Res.AccessToken != nil {
		s.D.Set("access_token", *s.Res.AccessToken)
	}

	arcAgentNodes := []interface{}{}
	for _, item := range s.Res.ArcAgentNodes {
		arcAgentNodes = append(arcAgentNodes, ArcAgentNodesToMap(item))
	}
	s.D.Set("arc_agent_nodes", arcAgentNodes)

	s.D.Set("azure_identity_connectivity_status", s.Res.AzureIdentityConnectivityStatus)
	s.D.Set("azure_identity_mechanism", s.Res.AzureIdentityMechanism)

	if s.Res.AzureResourceGroup != nil {
		s.D.Set("azure_resource_group", *s.Res.AzureResourceGroup)
	}

	if s.Res.AzureSubscriptionId != nil {
		s.D.Set("azure_subscription_id", *s.Res.AzureSubscriptionId)
	}

	if s.Res.AzureTenantId != nil {
		s.D.Set("azure_tenant_id", *s.Res.AzureTenantId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DbClusterResourceId != nil {
		s.D.Set("db_cluster_resource_id", *s.Res.DbClusterResourceId)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.LastModification != nil {
		s.D.Set("last_modification", *s.Res.LastModification)
	}

	if s.Res.LifecycleStateDetails != nil {
		s.D.Set("lifecycle_state_details", *s.Res.LifecycleStateDetails)
	}

	if s.Res.PrivateEndpointDnsAlias != nil {
		s.D.Set("private_endpoint_dns_alias", *s.Res.PrivateEndpointDnsAlias)
	}

	if s.Res.PrivateEndpointIpAddress != nil {
		s.D.Set("private_endpoint_ip_address", *s.Res.PrivateEndpointIpAddress)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}
	return nil
}

func ArcAgentNodesToMap(obj oci_dbmulticloud.ArcAgentNodes) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CurrentArcAgentVersion != nil {
		result["current_arc_agent_version"] = string(*obj.CurrentArcAgentVersion)
	}

	if obj.HostId != nil {
		result["host_id"] = string(*obj.HostId)
	}

	if obj.HostName != nil {
		result["host_name"] = string(*obj.HostName)
	}

	result["status"] = string(obj.Status)

	if obj.TimeLastChecked != nil {
		result["time_last_checked"] = obj.TimeLastChecked.String()
	}

	return result
}

func OracleDbAzureConnectorSummaryToMap(obj oci_dbmulticloud.OracleDbAzureConnectorSummary) map[string]interface{} {
	result := map[string]interface{}{}

	arcAgentNodes := []interface{}{}

	result["arc_agent_nodes"] = arcAgentNodes

	result["azure_identity_connectivity_status"] = string(obj.AzureIdentityConnectivityStatus)

	result["azure_identity_mechanism"] = string(obj.AzureIdentityMechanism)

	if obj.AzureResourceGroup != nil {
		result["azure_resource_group"] = string(*obj.AzureResourceGroup)
	}

	if obj.AzureSubscriptionId != nil {
		result["azure_subscription_id"] = string(*obj.AzureSubscriptionId)
	}

	if obj.AzureTenantId != nil {
		result["azure_tenant_id"] = string(*obj.AzureTenantId)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DbClusterResourceId != nil {
		result["db_cluster_resource_id"] = string(*obj.DbClusterResourceId)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LastModification != nil {
		result["last_modification"] = string(*obj.LastModification)
	}

	if obj.LifecycleStateDetails != nil {
		result["lifecycle_state_details"] = string(*obj.LifecycleStateDetails)
	}

	if obj.PrivateEndpointDnsAlias != nil {
		result["private_endpoint_dns_alias"] = string(*obj.PrivateEndpointDnsAlias)
	}

	if obj.PrivateEndpointIpAddress != nil {
		result["private_endpoint_ip_address"] = string(*obj.PrivateEndpointIpAddress)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func (s *DbmulticloudOracleDbAzureConnectorResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_dbmulticloud.ChangeOracleDbAzureConnectorCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.OracleDbAzureConnectorId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dbmulticloud")

	response, err := s.Client.ChangeOracleDbAzureConnectorCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getOracleDbAzureConnectorFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dbmulticloud"), oci_dbmulticloud.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
