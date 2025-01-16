// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package opensearch

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_opensearch "github.com/oracle/oci-go-sdk/v65/opensearch"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OpensearchOpensearchClusterResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("45m"),
			Update: tfresource.GetTimeoutDuration("45m"),
			Delete: tfresource.GetTimeoutDuration("45m"),
		},
		Create: createOpensearchOpensearchCluster,
		Read:   readOpensearchOpensearchCluster,
		Update: updateOpensearchOpensearchCluster,
		Delete: deleteOpensearchOpensearchCluster,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"data_node_count": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"data_node_host_memory_gb": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"data_node_host_ocpu_count": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"data_node_host_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"data_node_storage_gb": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"master_node_count": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"master_node_host_memory_gb": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"master_node_host_ocpu_count": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"master_node_host_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"opendashboard_node_count": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"opendashboard_node_host_memory_gb": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"opendashboard_node_host_ocpu_count": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"software_version": {
				Type:     schema.TypeString,
				Required: true,
			},
			"subnet_compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"vcn_compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"vcn_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"data_node_host_bare_metal_shape": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
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
			"inbound_cluster_ids": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"maintenance_details": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"notification_email_ids": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						// Computed
						"end_time": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"start_time": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"master_node_host_bare_metal_shape": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"outbound_cluster_config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"is_enabled": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"outbound_clusters": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"display_name": {
										Type:     schema.TypeString,
										Required: true,
									},
									"seed_cluster_id": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"is_skip_unavailable": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"mode": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"ping_schedule": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
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
			"reverse_connection_endpoint_customer_ips": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"security_master_user_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"security_master_user_password_hash": {
				Type:      schema.TypeString,
				Optional:  true,
				Computed:  true,
				Sensitive: true,
			},
			"security_mode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem:     schema.TypeString,
			},
			"configure_outbound_cluster_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			// Computed
			"availability_domains": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"fqdn": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"opendashboard_fqdn": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"opendashboard_private_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"opensearch_fqdn": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"opensearch_private_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"reverse_connection_endpoints": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"customer_ip": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"nat_ip": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_deleted": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"total_storage_gb": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func createOpensearchOpensearchCluster(d *schema.ResourceData, m interface{}) error {
	sync := &OpensearchOpensearchClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OpensearchClusterClient()

	if e := tfresource.CreateResource(d, sync); e != nil {
		return e
	}

	if _, ok := sync.D.GetOkExists("configure_outbound_cluster_trigger"); ok {
		err := sync.ConfigureOutboundCluster()
		if err != nil {
			return err
		}
	}
	return nil

}

func readOpensearchOpensearchCluster(d *schema.ResourceData, m interface{}) error {
	sync := &OpensearchOpensearchClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OpensearchClusterClient()

	return tfresource.ReadResource(sync)
}

func updateOpensearchOpensearchCluster(d *schema.ResourceData, m interface{}) error {
	sync := &OpensearchOpensearchClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OpensearchClusterClient()

	if _, ok := sync.D.GetOkExists("configure_outbound_cluster_trigger"); ok && sync.D.HasChange("configure_outbound_cluster_trigger") {
		oldRaw, newRaw := sync.D.GetChange("configure_outbound_cluster_trigger")
		oldValue := oldRaw.(int)
		newValue := newRaw.(int)
		if oldValue < newValue {
			err := sync.ConfigureOutboundCluster()

			if err != nil {
				return err
			}
		} else {
			sync.D.Set("configure_outbound_cluster_trigger", oldRaw)
			return fmt.Errorf("new value of trigger should be greater than the old value")
		}
	}

	if err := tfresource.UpdateResource(d, sync); err != nil {
		return err
	}

	return nil
}

func deleteOpensearchOpensearchCluster(d *schema.ResourceData, m interface{}) error {
	sync := &OpensearchOpensearchClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OpensearchClusterClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type OpensearchOpensearchClusterResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_opensearch.OpensearchClusterClient
	Res                    *oci_opensearch.OpensearchCluster
	DisableNotFoundRetries bool
}

func (s *OpensearchOpensearchClusterResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *OpensearchOpensearchClusterResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_opensearch.OpensearchClusterLifecycleStateCreating),
	}
}

func (s *OpensearchOpensearchClusterResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_opensearch.OpensearchClusterLifecycleStateActive),
	}
}

func (s *OpensearchOpensearchClusterResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_opensearch.OpensearchClusterLifecycleStateDeleting),
	}
}

func (s *OpensearchOpensearchClusterResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_opensearch.OpensearchClusterLifecycleStateDeleted),
	}
}

func (s *OpensearchOpensearchClusterResourceCrud) Create() error {
	request := oci_opensearch.CreateOpensearchClusterRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if dataNodeCount, ok := s.D.GetOkExists("data_node_count"); ok {
		tmp := dataNodeCount.(int)
		request.DataNodeCount = &tmp
	}

	if dataNodeHostBareMetalShape, ok := s.D.GetOkExists("data_node_host_bare_metal_shape"); ok {
		tmp := dataNodeHostBareMetalShape.(string)
		request.DataNodeHostBareMetalShape = &tmp
	}

	if dataNodeHostMemoryGB, ok := s.D.GetOkExists("data_node_host_memory_gb"); ok {
		tmp := dataNodeHostMemoryGB.(int)
		request.DataNodeHostMemoryGB = &tmp
	}

	if dataNodeHostOcpuCount, ok := s.D.GetOkExists("data_node_host_ocpu_count"); ok {
		tmp := dataNodeHostOcpuCount.(int)
		request.DataNodeHostOcpuCount = &tmp
	}

	if dataNodeHostType, ok := s.D.GetOkExists("data_node_host_type"); ok {
		request.DataNodeHostType = oci_opensearch.DataNodeHostTypeEnum(dataNodeHostType.(string))
	}

	if dataNodeStorageGB, ok := s.D.GetOkExists("data_node_storage_gb"); ok {
		tmp := dataNodeStorageGB.(int)
		request.DataNodeStorageGB = &tmp
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

	if inboundClusterIds, ok := s.D.GetOkExists("inbound_cluster_ids"); ok {
		interfaces := inboundClusterIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("inbound_cluster_ids") {
			request.InboundClusterIds = tmp
		}
	}

	if maintenanceDetails, ok := s.D.GetOkExists("maintenance_details"); ok {
		if tmpList := maintenanceDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "maintenance_details", 0)
			tmp, err := s.mapToCreateMaintenanceDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.MaintenanceDetails = &tmp
		}
	}

	if masterNodeCount, ok := s.D.GetOkExists("master_node_count"); ok {
		tmp := masterNodeCount.(int)
		request.MasterNodeCount = &tmp
	}

	if masterNodeHostBareMetalShape, ok := s.D.GetOkExists("master_node_host_bare_metal_shape"); ok {
		tmp := masterNodeHostBareMetalShape.(string)
		request.MasterNodeHostBareMetalShape = &tmp
	}

	if masterNodeHostMemoryGB, ok := s.D.GetOkExists("master_node_host_memory_gb"); ok {
		tmp := masterNodeHostMemoryGB.(int)
		request.MasterNodeHostMemoryGB = &tmp
	}

	if masterNodeHostOcpuCount, ok := s.D.GetOkExists("master_node_host_ocpu_count"); ok {
		tmp := masterNodeHostOcpuCount.(int)
		request.MasterNodeHostOcpuCount = &tmp
	}

	if masterNodeHostType, ok := s.D.GetOkExists("master_node_host_type"); ok {
		request.MasterNodeHostType = oci_opensearch.MasterNodeHostTypeEnum(masterNodeHostType.(string))
	}

	if opendashboardNodeCount, ok := s.D.GetOkExists("opendashboard_node_count"); ok {
		tmp := opendashboardNodeCount.(int)
		request.OpendashboardNodeCount = &tmp
	}

	if opendashboardNodeHostMemoryGB, ok := s.D.GetOkExists("opendashboard_node_host_memory_gb"); ok {
		tmp := opendashboardNodeHostMemoryGB.(int)
		request.OpendashboardNodeHostMemoryGB = &tmp
	}

	if opendashboardNodeHostOcpuCount, ok := s.D.GetOkExists("opendashboard_node_host_ocpu_count"); ok {
		tmp := opendashboardNodeHostOcpuCount.(int)
		request.OpendashboardNodeHostOcpuCount = &tmp
	}

	if outboundClusterConfig, ok := s.D.GetOkExists("outbound_cluster_config"); ok {
		if tmpList := outboundClusterConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "outbound_cluster_config", 0)
			tmp, err := s.mapToOutboundClusterConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.OutboundClusterConfig = &tmp
		}
	}

	if reverseConnectionEndpointCustomerIps, ok := s.D.GetOkExists("reverse_connection_endpoint_customer_ips"); ok {
		interfaces := reverseConnectionEndpointCustomerIps.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("reverse_connection_endpoint_customer_ips") {
			request.ReverseConnectionEndpointCustomerIps = tmp
		}
	}

	if securityMasterUserName, ok := s.D.GetOkExists("security_master_user_name"); ok {
		tmp := securityMasterUserName.(string)
		request.SecurityMasterUserName = &tmp
	}

	if securityMasterUserPasswordHash, ok := s.D.GetOkExists("security_master_user_password_hash"); ok {
		tmp := securityMasterUserPasswordHash.(string)
		request.SecurityMasterUserPasswordHash = &tmp
	}

	if securityMode, ok := s.D.GetOkExists("security_mode"); ok {
		request.SecurityMode = oci_opensearch.SecurityModeEnum(securityMode.(string))
	}

	if softwareVersion, ok := s.D.GetOkExists("software_version"); ok {
		tmp := softwareVersion.(string)
		request.SoftwareVersion = &tmp
	}

	if subnetCompartmentId, ok := s.D.GetOkExists("subnet_compartment_id"); ok {
		tmp := subnetCompartmentId.(string)
		request.SubnetCompartmentId = &tmp
	}

	if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
		tmp := subnetId.(string)
		request.SubnetId = &tmp
	}

	if systemTags, ok := s.D.GetOkExists("system_tags"); ok {
		convertedSystemTags, err := tfresource.MapToSystemTags(systemTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.SystemTags = convertedSystemTags
	}

	if vcnCompartmentId, ok := s.D.GetOkExists("vcn_compartment_id"); ok {
		tmp := vcnCompartmentId.(string)
		request.VcnCompartmentId = &tmp
	}

	if vcnId, ok := s.D.GetOkExists("vcn_id"); ok {
		tmp := vcnId.(string)
		request.VcnId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opensearch")

	response, err := s.Client.CreateOpensearchCluster(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	workRequestResponse := oci_opensearch.GetWorkRequestResponse{}
	workRequestResponse, err = s.Client.GetWorkRequest(context.Background(),
		oci_opensearch.GetWorkRequestRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opensearch"),
			},
		})
	if err == nil {
		// The work request response contains an array of objects
		for _, res := range workRequestResponse.Resources {
			if res.EntityType != nil && strings.Contains(strings.ToLower(*res.EntityType), "opensearch") && res.Identifier != nil {
				s.D.SetId(*res.Identifier)
				break
			}
		}
	}
	return s.getOpensearchClusterFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opensearch"), oci_opensearch.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *OpensearchOpensearchClusterResourceCrud) getOpensearchClusterFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_opensearch.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	opensearchClusterId, err := opensearchClusterWaitForWorkRequest(workId, "opensearch",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*opensearchClusterId)

	return s.Get()
}

func opensearchClusterWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "opensearch", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_opensearch.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func opensearchClusterWaitForWorkRequest(wId *string, entityType string, action oci_opensearch.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_opensearch.OpensearchClusterClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "opensearch")
	retryPolicy.ShouldRetryOperation = opensearchClusterWorkRequestShouldRetryFunc(timeout)

	response := oci_opensearch.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_opensearch.OperationStatusInProgress),
			string(oci_opensearch.OperationStatusAccepted),
			string(oci_opensearch.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_opensearch.OperationStatusSucceeded),
			string(oci_opensearch.OperationStatusFailed),
			string(oci_opensearch.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_opensearch.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_opensearch.OperationStatusFailed || response.Status == oci_opensearch.OperationStatusCanceled {
		return nil, getErrorFromOpensearchOpensearchClusterWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromOpensearchOpensearchClusterWorkRequest(client *oci_opensearch.OpensearchClusterClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_opensearch.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_opensearch.ListWorkRequestErrorsRequest{
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

func (s *OpensearchOpensearchClusterResourceCrud) Get() error {
	request := oci_opensearch.GetOpensearchClusterRequest{}

	tmp := s.D.Id()
	request.OpensearchClusterId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opensearch")

	response, err := s.Client.GetOpensearchCluster(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.OpensearchCluster
	return nil
}

func (s *OpensearchOpensearchClusterResourceCrud) HorizontalConditionMet() (result bool) {
	if _, ok := s.D.GetOkExists("data_node_count"); ok && s.D.HasChange("data_node_count") {
		return true
	} else if _, ok := s.D.GetOkExists("master_node_count"); ok && s.D.HasChange("master_node_count") {
		return true
	} else if _, ok := s.D.GetOkExists("opendashboard_node_count"); ok && s.D.HasChange("opendashboard_node_count") {
		return true
	}
	return false
}

func (s *OpensearchOpensearchClusterResourceCrud) VerticalConditionMet() (result bool) {
	if _, ok := s.D.GetOkExists("master_node_host_ocpu_count"); ok && s.D.HasChange("master_node_host_ocpu_count") {
		return true
	} else if _, ok := s.D.GetOkExists("master_node_host_memory_gb"); ok && s.D.HasChange("master_node_host_memory_gb") {
		return true
	} else if _, ok := s.D.GetOkExists("data_node_host_ocpu_count"); ok && s.D.HasChange("data_node_host_ocpu_count") {
		return true
	} else if _, ok := s.D.GetOkExists("data_node_host_memory_gb"); ok && s.D.HasChange("data_node_host_memory_gb") {
		return true
	} else if _, ok := s.D.GetOkExists("data_node_storage_gb"); ok && s.D.HasChange("data_node_storage_gb") {
		return true
	} else if _, ok := s.D.GetOkExists("opendashboard_node_host_ocpu_count"); ok && s.D.HasChange("opendashboard_node_host_ocpu_count") {
		return true
	} else if _, ok := s.D.GetOkExists("opendashboard_node_host_memory_gb"); ok && s.D.HasChange("opendashboard_node_host_memory_gb") {
		return true
	}
	return false
}

func (s *OpensearchOpensearchClusterResourceCrud) Update() error {
	if s.HorizontalConditionMet() {
		log.Println("Horizontal Resize Begin...")
		err := s.ResizeOpensearchClusterHorizontal()
		if err != nil {
			return err
		}
	}

	if s.VerticalConditionMet() {
		log.Println("Vertical Resize Begin...")
		err := s.ResizeOpensearchClusterVertical()
		if err != nil {
			return err
		}
	}

	request := oci_opensearch.UpdateOpensearchClusterRequest{}

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

	if maintenanceDetails, ok := s.D.GetOkExists("maintenance_details"); ok {
		if tmpList := maintenanceDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "maintenance_details", 0)
			tmp, err := s.mapToUpdateMaintenanceDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.MaintenanceDetails = &tmp
		}
	}

	tmp := s.D.Id()
	request.OpensearchClusterId = &tmp

	if outboundClusterConfig, ok := s.D.GetOkExists("outbound_cluster_config"); ok {
		if tmpList := outboundClusterConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "outbound_cluster_config", 0)
			tmp, err := s.mapToOutboundClusterConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.OutboundClusterConfig = &tmp
		}
	}

	if reverseConnectionEndpointCustomerIps, ok := s.D.GetOkExists("reverse_connection_endpoint_customer_ips"); ok {
		interfaces := reverseConnectionEndpointCustomerIps.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("reverse_connection_endpoint_customer_ips") {
			request.ReverseConnectionEndpointCustomerIps = tmp
		}
	}

	if securityMasterUserName, ok := s.D.GetOkExists("security_master_user_name"); ok {
		tmp := securityMasterUserName.(string)
		request.SecurityMasterUserName = &tmp
	}

	if securityMasterUserPasswordHash, ok := s.D.GetOkExists("security_master_user_password_hash"); ok {
		tmp := securityMasterUserPasswordHash.(string)
		request.SecurityMasterUserPasswordHash = &tmp
	}

	if securityMode, ok := s.D.GetOkExists("security_mode"); ok {
		request.SecurityMode = oci_opensearch.SecurityModeEnum(securityMode.(string))
	}

	if softwareVersion, ok := s.D.GetOkExists("software_version"); ok {
		tmp := softwareVersion.(string)
		request.SoftwareVersion = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opensearch")

	response, err := s.Client.UpdateOpensearchCluster(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getOpensearchClusterFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opensearch"), oci_opensearch.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *OpensearchOpensearchClusterResourceCrud) Delete() error {
	request := oci_opensearch.DeleteOpensearchClusterRequest{}

	tmp := s.D.Id()
	request.OpensearchClusterId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opensearch")

	response, err := s.Client.DeleteOpensearchCluster(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := opensearchClusterWaitForWorkRequest(workId, "opensearch",
		oci_opensearch.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *OpensearchOpensearchClusterResourceCrud) SetData() error {
	s.D.Set("availability_domains", s.Res.AvailabilityDomains)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DataNodeCount != nil {
		s.D.Set("data_node_count", *s.Res.DataNodeCount)
	}

	if s.Res.DataNodeHostBareMetalShape != nil {
		s.D.Set("data_node_host_bare_metal_shape", *s.Res.DataNodeHostBareMetalShape)
	}

	if s.Res.DataNodeHostMemoryGB != nil {
		s.D.Set("data_node_host_memory_gb", *s.Res.DataNodeHostMemoryGB)
	}

	if s.Res.DataNodeHostOcpuCount != nil {
		s.D.Set("data_node_host_ocpu_count", *s.Res.DataNodeHostOcpuCount)
	}

	s.D.Set("data_node_host_type", s.Res.DataNodeHostType)

	if s.Res.DataNodeStorageGB != nil {
		s.D.Set("data_node_storage_gb", *s.Res.DataNodeStorageGB)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.Fqdn != nil {
		s.D.Set("fqdn", *s.Res.Fqdn)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("inbound_cluster_ids", s.Res.InboundClusterIds)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.MaintenanceDetails != nil {
		s.D.Set("maintenance_details", []interface{}{MaintenanceDetailsToMap(s.Res.MaintenanceDetails)})
	} else {
		s.D.Set("maintenance_details", nil)
	}

	if s.Res.MasterNodeCount != nil {
		s.D.Set("master_node_count", *s.Res.MasterNodeCount)
	}

	if s.Res.MasterNodeHostBareMetalShape != nil {
		s.D.Set("master_node_host_bare_metal_shape", *s.Res.MasterNodeHostBareMetalShape)
	}

	if s.Res.MasterNodeHostMemoryGB != nil {
		s.D.Set("master_node_host_memory_gb", *s.Res.MasterNodeHostMemoryGB)
	}

	if s.Res.MasterNodeHostOcpuCount != nil {
		s.D.Set("master_node_host_ocpu_count", *s.Res.MasterNodeHostOcpuCount)
	}

	s.D.Set("master_node_host_type", s.Res.MasterNodeHostType)

	if s.Res.OpendashboardFqdn != nil {
		s.D.Set("opendashboard_fqdn", *s.Res.OpendashboardFqdn)
	}

	if s.Res.OpendashboardNodeCount != nil {
		s.D.Set("opendashboard_node_count", *s.Res.OpendashboardNodeCount)
	}

	if s.Res.OpendashboardNodeHostMemoryGB != nil {
		s.D.Set("opendashboard_node_host_memory_gb", *s.Res.OpendashboardNodeHostMemoryGB)
	}

	if s.Res.OpendashboardNodeHostOcpuCount != nil {
		s.D.Set("opendashboard_node_host_ocpu_count", *s.Res.OpendashboardNodeHostOcpuCount)
	}

	if s.Res.OpendashboardPrivateIp != nil {
		s.D.Set("opendashboard_private_ip", *s.Res.OpendashboardPrivateIp)
	}

	if s.Res.OpensearchFqdn != nil {
		s.D.Set("opensearch_fqdn", *s.Res.OpensearchFqdn)
	}

	if s.Res.OpensearchPrivateIp != nil {
		s.D.Set("opensearch_private_ip", *s.Res.OpensearchPrivateIp)
	}

	if s.Res.OutboundClusterConfig != nil {
		s.D.Set("outbound_cluster_config", []interface{}{OutboundClusterConfigToMap(s.Res.OutboundClusterConfig)})
	} else {
		s.D.Set("outbound_cluster_config", nil)
	}

	s.D.Set("reverse_connection_endpoint_customer_ips", s.Res.ReverseConnectionEndpointCustomerIps)

	reverseConnectionEndpoints := []interface{}{}
	for _, item := range s.Res.ReverseConnectionEndpoints {
		reverseConnectionEndpoints = append(reverseConnectionEndpoints, ReverseConnectionEndpointToMap(item))
	}
	s.D.Set("reverse_connection_endpoints", reverseConnectionEndpoints)

	if s.Res.SecurityMasterUserName != nil {
		s.D.Set("security_master_user_name", *s.Res.SecurityMasterUserName)
	}

	if s.Res.SecurityMasterUserPasswordHash != nil {
		s.D.Set("security_master_user_password_hash", *s.Res.SecurityMasterUserPasswordHash)
	}

	s.D.Set("security_mode", s.Res.SecurityMode)

	if s.Res.SoftwareVersion != nil {
		s.D.Set("software_version", *s.Res.SoftwareVersion)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SubnetCompartmentId != nil {
		s.D.Set("subnet_compartment_id", *s.Res.SubnetCompartmentId)
	}

	if s.Res.SubnetId != nil {
		s.D.Set("subnet_id", *s.Res.SubnetId)
	}

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeDeleted != nil {
		s.D.Set("time_deleted", s.Res.TimeDeleted.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.TotalStorageGB != nil {
		s.D.Set("total_storage_gb", *s.Res.TotalStorageGB)
	}

	if s.Res.VcnCompartmentId != nil {
		s.D.Set("vcn_compartment_id", *s.Res.VcnCompartmentId)
	}

	if s.Res.VcnId != nil {
		s.D.Set("vcn_id", *s.Res.VcnId)
	}

	return nil
}

func (s *OpensearchOpensearchClusterResourceCrud) ConfigureOutboundCluster() error {
	request := oci_opensearch.ConfigureOutboundClusterRequest{}

	if inboundClusterIds, ok := s.D.GetOkExists("inbound_cluster_ids"); ok {
		interfaces := inboundClusterIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("inbound_cluster_ids") {
			request.InboundClusterIds = tmp
		}
	}

	idTmp := s.D.Id()
	request.OpensearchClusterId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opensearch")

	_, err := s.Client.ConfigureOutboundCluster(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	val := s.D.Get("configure_outbound_cluster_trigger")
	s.D.Set("configure_outbound_cluster_trigger", val)

	return nil
}

func (s *OpensearchOpensearchClusterResourceCrud) ResizeOpensearchClusterHorizontal() error {
	tfresource.ShortRetryTime = tfresource.LongRetryTime * 5
	request := oci_opensearch.ResizeOpensearchClusterHorizontalRequest{}

	if dataNodeCount, ok := s.D.GetOkExists("data_node_count"); ok {
		tmp := dataNodeCount.(int)
		request.DataNodeCount = &tmp
	}

	if masterNodeCount, ok := s.D.GetOkExists("master_node_count"); ok {
		tmp := masterNodeCount.(int)
		request.MasterNodeCount = &tmp
	}

	if opendashboardNodeCount, ok := s.D.GetOkExists("opendashboard_node_count"); ok {
		tmp := opendashboardNodeCount.(int)
		request.OpendashboardNodeCount = &tmp
	}

	idTmp := s.D.Id()
	request.OpensearchClusterId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opensearch")

	_, err := s.Client.ResizeOpensearchClusterHorizontal(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

func (s *OpensearchOpensearchClusterResourceCrud) ResizeOpensearchClusterVertical() error {
	tfresource.ShortRetryTime = tfresource.LongRetryTime * 5
	request := oci_opensearch.ResizeOpensearchClusterVerticalRequest{}

	if dataNodeHostMemoryGB, ok := s.D.GetOkExists("data_node_host_memory_gb"); ok {
		tmp := dataNodeHostMemoryGB.(int)
		request.DataNodeHostMemoryGB = &tmp
	}

	if dataNodeHostOcpuCount, ok := s.D.GetOkExists("data_node_host_ocpu_count"); ok {
		tmp := dataNodeHostOcpuCount.(int)
		request.DataNodeHostOcpuCount = &tmp
	}

	if dataNodeStorageGB, ok := s.D.GetOkExists("data_node_storage_gb"); ok {
		tmp := dataNodeStorageGB.(int)
		request.DataNodeStorageGB = &tmp
	}

	if masterNodeHostMemoryGB, ok := s.D.GetOkExists("master_node_host_memory_gb"); ok {
		tmp := masterNodeHostMemoryGB.(int)
		request.MasterNodeHostMemoryGB = &tmp
	}

	if masterNodeHostOcpuCount, ok := s.D.GetOkExists("master_node_host_ocpu_count"); ok {
		tmp := masterNodeHostOcpuCount.(int)
		request.MasterNodeHostOcpuCount = &tmp
	}

	if opendashboardNodeHostMemoryGB, ok := s.D.GetOkExists("opendashboard_node_host_memory_gb"); ok {
		tmp := opendashboardNodeHostMemoryGB.(int)
		request.OpendashboardNodeHostMemoryGB = &tmp
	}

	if opendashboardNodeHostOcpuCount, ok := s.D.GetOkExists("opendashboard_node_host_ocpu_count"); ok {
		tmp := opendashboardNodeHostOcpuCount.(int)
		request.OpendashboardNodeHostOcpuCount = &tmp
	}

	idTmp := s.D.Id()
	request.OpensearchClusterId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opensearch")

	_, err := s.Client.ResizeOpensearchClusterVertical(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

func (s *OpensearchOpensearchClusterResourceCrud) mapToCreateMaintenanceDetails(fieldKeyFormat string) (oci_opensearch.CreateMaintenanceDetails, error) {
	result := oci_opensearch.CreateMaintenanceDetails{}

	if notificationEmailIds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "notification_email_ids")); ok {
		interfaces := notificationEmailIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "notification_email_ids")) {
			result.NotificationEmailIds = tmp
		}
	}

	return result, nil
}

func (s *OpensearchOpensearchClusterResourceCrud) mapToUpdateMaintenanceDetails(fieldKeyFormat string) (oci_opensearch.UpdateMaintenanceDetails, error) {
	result := oci_opensearch.UpdateMaintenanceDetails{}

	if notificationEmailIds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "notification_email_ids")); ok {
		interfaces := notificationEmailIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "notification_email_ids")) {
			result.NotificationEmailIds = tmp
		}
	}

	return result, nil
}

func MaintenanceDetailsToMap(obj *oci_opensearch.MaintenanceDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.EndTime != nil {
		result["end_time"] = obj.EndTime.String()
	}

	result["notification_email_ids"] = obj.NotificationEmailIds

	if obj.StartTime != nil {
		result["start_time"] = obj.StartTime.String()
	}

	result["state"] = string(obj.State)

	return result
}

func OpensearchClusterSummaryToMap(obj oci_opensearch.OpensearchClusterSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["availability_domains"] = obj.AvailabilityDomains

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.OutboundClusterConfig != nil {
		result["outbound_cluster_config"] = []interface{}{OutboundClusterConfigToMap(obj.OutboundClusterConfig)}
	}

	result["security_mode"] = string(obj.SecurityMode)

	if obj.SoftwareVersion != nil {
		result["software_version"] = string(*obj.SoftwareVersion)
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

	if obj.TotalStorageGB != nil {
		result["total_storage_gb"] = int(*obj.TotalStorageGB)
	}

	return result
}

func (s *OpensearchOpensearchClusterResourceCrud) mapToOutboundClusterConfig(fieldKeyFormat string) (oci_opensearch.OutboundClusterConfig, error) {
	result := oci_opensearch.OutboundClusterConfig{}

	if isEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_enabled")); ok {
		tmp := isEnabled.(bool)
		result.IsEnabled = &tmp
	}

	if outboundClusters, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "outbound_clusters")); ok {
		interfaces := outboundClusters.([]interface{})
		tmp := make([]oci_opensearch.OutboundClusterSummary, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "outbound_clusters"), stateDataIndex)
			converted, err := s.mapToOutboundClusterSummary(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "outbound_clusters")) {
			result.OutboundClusters = tmp
		}
	}

	return result, nil
}

func OutboundClusterConfigToMap(obj *oci_opensearch.OutboundClusterConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsEnabled != nil {
		result["is_enabled"] = bool(*obj.IsEnabled)
	}

	outboundClusters := []interface{}{}
	for _, item := range obj.OutboundClusters {
		outboundClusters = append(outboundClusters, OutboundClusterSummaryToMap(item))
	}
	result["outbound_clusters"] = outboundClusters

	return result
}

func (s *OpensearchOpensearchClusterResourceCrud) mapToOutboundClusterSummary(fieldKeyFormat string) (oci_opensearch.OutboundClusterSummary, error) {
	result := oci_opensearch.OutboundClusterSummary{}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	if isSkipUnavailable, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_skip_unavailable")); ok {
		tmp := isSkipUnavailable.(bool)
		result.IsSkipUnavailable = &tmp
	}

	if mode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mode")); ok {
		result.Mode = oci_opensearch.CccModeEnum(mode.(string))
	}

	if pingSchedule, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ping_schedule")); ok {
		tmp := pingSchedule.(string)
		result.PingSchedule = &tmp
	}

	if seedClusterId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "seed_cluster_id")); ok {
		tmp := seedClusterId.(string)
		result.SeedClusterId = &tmp
	}

	return result, nil
}

func OutboundClusterSummaryToMap(obj oci_opensearch.OutboundClusterSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.IsSkipUnavailable != nil {
		result["is_skip_unavailable"] = bool(*obj.IsSkipUnavailable)
	}

	result["mode"] = string(obj.Mode)

	if obj.PingSchedule != nil {
		result["ping_schedule"] = string(*obj.PingSchedule)
	}

	if obj.SeedClusterId != nil {
		result["seed_cluster_id"] = string(*obj.SeedClusterId)
	}

	return result
}

func ReverseConnectionEndpointToMap(obj oci_opensearch.ReverseConnectionEndpoint) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CustomerIp != nil {
		result["customer_ip"] = string(*obj.CustomerIp)
	}

	if obj.NatIp != nil {
		result["nat_ip"] = string(*obj.NatIp)
	}

	return result
}

func (s *OpensearchOpensearchClusterResourceCrud) mapToSecuritySamlConfig(fieldKeyFormat string) (oci_opensearch.SecuritySamlConfig, error) {
	result := oci_opensearch.SecuritySamlConfig{}

	if adminBackendRole, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "admin_backend_role")); ok {
		tmp := adminBackendRole.(string)
		result.AdminBackendRole = &tmp
	}

	if idpEntityId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "idp_entity_id")); ok {
		tmp := idpEntityId.(string)
		result.IdpEntityId = &tmp
	}

	if idpMetadataContent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "idp_metadata_content")); ok {
		tmp := idpMetadataContent.(string)
		result.IdpMetadataContent = &tmp
	}

	if isEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_enabled")); ok {
		tmp := isEnabled.(bool)
		result.IsEnabled = &tmp
	}

	if opendashboardUrl, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "opendashboard_url")); ok {
		tmp := opendashboardUrl.(string)
		result.OpendashboardUrl = &tmp
	}

	if rolesKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "roles_key")); ok {
		tmp := rolesKey.(string)
		result.RolesKey = &tmp
	}

	if subjectKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "subject_key")); ok {
		tmp := subjectKey.(string)
		result.SubjectKey = &tmp
	}

	return result, nil
}

func SecuritySamlConfigToMap(obj *oci_opensearch.SecuritySamlConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AdminBackendRole != nil {
		result["admin_backend_role"] = string(*obj.AdminBackendRole)
	}

	if obj.IdpEntityId != nil {
		result["idp_entity_id"] = string(*obj.IdpEntityId)
	}

	if obj.IdpMetadataContent != nil {
		result["idp_metadata_content"] = string(*obj.IdpMetadataContent)
	}

	if obj.IsEnabled != nil {
		result["is_enabled"] = bool(*obj.IsEnabled)
	}

	if obj.OpendashboardUrl != nil {
		result["opendashboard_url"] = string(*obj.OpendashboardUrl)
	}

	if obj.RolesKey != nil {
		result["roles_key"] = string(*obj.RolesKey)
	}

	if obj.SubjectKey != nil {
		result["subject_key"] = string(*obj.SubjectKey)
	}

	return result
}
