// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package datacc

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_datacc "github.com/oracle/oci-go-sdk/v65/datacc"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataccVmClusterNetworkResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createDataccVmClusterNetworkWithContext,
		ReadContext:   readDataccVmClusterNetworkWithContext,
		UpdateContext: updateDataccVmClusterNetworkWithContext,
		DeleteContext: deleteDataccVmClusterNetworkWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"infrastructure_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"vm_networks": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"domain_name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"gateway": {
							Type:     schema.TypeString,
							Required: true,
						},
						"netmask": {
							Type:     schema.TypeString,
							Required: true,
						},
						"network_type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"nodes": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"hostname": {
										Type:     schema.TypeString,
										Required: true,
									},
									"ip": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"vip": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"vip_hostname": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},

						// Optional
						"prefix": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vlan_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},

			// Optional
			"consumer_type": {
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
			"dns_servers": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"listener_port": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"listener_port_ssl": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"node_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"ntp_servers": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"scans": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"hostname": {
							Type:     schema.TypeString,
							Required: true,
						},
						"ips": {
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

			// Computed
			"associated_resource_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"base_vm_cluster_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_scan_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
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

func createDataccVmClusterNetworkWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataccVmClusterNetworkResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BaseinfraClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readDataccVmClusterNetworkWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataccVmClusterNetworkResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BaseinfraClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateDataccVmClusterNetworkWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataccVmClusterNetworkResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BaseinfraClient()

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func deleteDataccVmClusterNetworkWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataccVmClusterNetworkResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BaseinfraClient()
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type DataccVmClusterNetworkResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_datacc.BaseinfraClient
	Res                    *oci_datacc.VmClusterNetwork
	DisableNotFoundRetries bool
}

func (s *DataccVmClusterNetworkResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DataccVmClusterNetworkResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_datacc.VmClusterNetworkLifecycleStateCreating),
		string(oci_datacc.VmClusterNetworkLifecycleStateValidating),
	}
}

func (s *DataccVmClusterNetworkResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_datacc.VmClusterNetworkLifecycleStateRequiresValidation),
		string(oci_datacc.VmClusterNetworkLifecycleStateValidated),
	}
}

func (s *DataccVmClusterNetworkResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_datacc.VmClusterNetworkLifecycleStateDeleting),
	}
}

func (s *DataccVmClusterNetworkResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_datacc.VmClusterNetworkLifecycleStateDeleted),
	}
}

func (s *DataccVmClusterNetworkResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_datacc.CreateVmClusterNetworkRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if consumerType, ok := s.D.GetOkExists("consumer_type"); ok {
		request.ConsumerType = oci_datacc.VmNetworkConsumerTypeEnum(consumerType.(string))
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

	if dnsServers, ok := s.D.GetOkExists("dns_servers"); ok {
		interfaces := dnsServers.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("dns_servers") {
			request.DnsServers = tmp
		}
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if infrastructureId, ok := s.D.GetOkExists("infrastructure_id"); ok {
		tmp := infrastructureId.(string)
		request.InfrastructureId = &tmp
	}

	if listenerPort, ok := s.D.GetOkExists("listener_port"); ok {
		tmp := listenerPort.(int)
		request.ListenerPort = &tmp
	}

	if listenerPortSsl, ok := s.D.GetOkExists("listener_port_ssl"); ok {
		tmp := listenerPortSsl.(int)
		request.ListenerPortSsl = &tmp
	}

	if nodeCount, ok := s.D.GetOkExists("node_count"); ok {
		tmp := nodeCount.(int)
		request.NodeCount = &tmp
	}

	if ntpServers, ok := s.D.GetOkExists("ntp_servers"); ok {
		interfaces := ntpServers.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("ntp_servers") {
			request.NtpServers = tmp
		}
	}

	if scans, ok := s.D.GetOkExists("scans"); ok {
		interfaces := scans.([]interface{})
		tmp := make([]oci_datacc.ScanDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "scans", stateDataIndex)
			converted, err := s.mapToScanDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("scans") {
			request.Scans = tmp
		}
	}

	if vmNetworks, ok := s.D.GetOkExists("vm_networks"); ok {
		interfaces := vmNetworks.([]interface{})
		tmp := make([]oci_datacc.VmNetworkDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "vm_networks", stateDataIndex)
			converted, err := s.mapToVmNetworkDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("vm_networks") {
			request.VmNetworks = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datacc")

	response, err := s.Client.CreateVmClusterNetwork(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getVmClusterNetworkFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datacc"), oci_datacc.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DataccVmClusterNetworkResourceCrud) getVmClusterNetworkFromWorkRequest(ctx context.Context, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_datacc.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	vmClusterNetworkId, err := vmClusterNetworkWaitForWorkRequest(ctx, workId, "vmclusternetwork",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*vmClusterNetworkId)

	return s.GetWithContext(ctx)
}

func vmClusterNetworkWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "datacc", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_datacc.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func vmClusterNetworkWaitForWorkRequest(ctx context.Context, wId *string, entityType string, action oci_datacc.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_datacc.BaseinfraClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "datacc")
	retryPolicy.ShouldRetryOperation = vmClusterNetworkWorkRequestShouldRetryFunc(timeout)

	response := oci_datacc.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_datacc.OperationStatusInProgress),
			string(oci_datacc.OperationStatusAccepted),
			string(oci_datacc.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_datacc.OperationStatusSucceeded),
			string(oci_datacc.OperationStatusFailed),
			string(oci_datacc.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(ctx,
				oci_datacc.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_datacc.OperationStatusFailed || response.Status == oci_datacc.OperationStatusCanceled {
		return nil, getErrorFromDataccVmClusterNetworkWorkRequest(ctx, client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDataccVmClusterNetworkWorkRequest(ctx context.Context, client *oci_datacc.BaseinfraClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_datacc.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(ctx,
		oci_datacc.ListWorkRequestErrorsRequest{
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

func (s *DataccVmClusterNetworkResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_datacc.GetVmClusterNetworkRequest{}

	tmp := s.D.Id()
	request.VmClusterNetworkId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datacc")

	response, err := s.Client.GetVmClusterNetwork(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.VmClusterNetwork
	return nil
}

func (s *DataccVmClusterNetworkResourceCrud) UpdateWithContext(ctx context.Context) error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(ctx, compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_datacc.UpdateVmClusterNetworkRequest{}

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

	if dnsServers, ok := s.D.GetOkExists("dns_servers"); ok {
		interfaces := dnsServers.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("dns_servers") {
			request.DnsServers = tmp
		}
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if listenerPort, ok := s.D.GetOkExists("listener_port"); ok {
		tmp := listenerPort.(int)
		request.ListenerPort = &tmp
	}

	if listenerPortSsl, ok := s.D.GetOkExists("listener_port_ssl"); ok {
		tmp := listenerPortSsl.(int)
		request.ListenerPortSsl = &tmp
	}

	if ntpServers, ok := s.D.GetOkExists("ntp_servers"); ok {
		interfaces := ntpServers.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("ntp_servers") {
			request.NtpServers = tmp
		}
	}

	if scans, ok := s.D.GetOkExists("scans"); ok {
		interfaces := scans.([]interface{})
		tmp := make([]oci_datacc.ScanDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "scans", stateDataIndex)
			converted, err := s.mapToScanDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("scans") {
			request.Scans = tmp
		}
	}

	tmp := s.D.Id()
	request.VmClusterNetworkId = &tmp

	if vmNetworks, ok := s.D.GetOkExists("vm_networks"); ok {
		interfaces := vmNetworks.([]interface{})
		tmp := make([]oci_datacc.VmNetworkDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "vm_networks", stateDataIndex)
			converted, err := s.mapToVmNetworkDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("vm_networks") {
			request.VmNetworks = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datacc")

	response, err := s.Client.UpdateVmClusterNetwork(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getVmClusterNetworkFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datacc"), oci_datacc.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DataccVmClusterNetworkResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_datacc.DeleteVmClusterNetworkRequest{}

	tmp := s.D.Id()
	request.VmClusterNetworkId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datacc")

	response, err := s.Client.DeleteVmClusterNetwork(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := vmClusterNetworkWaitForWorkRequest(ctx, workId, "vmclusternetwork",
		oci_datacc.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *DataccVmClusterNetworkResourceCrud) SetData() error {
	if s.Res.AssociatedResourceId != nil {
		s.D.Set("associated_resource_id", *s.Res.AssociatedResourceId)
	}

	if s.Res.BaseVmClusterId != nil {
		s.D.Set("base_vm_cluster_id", *s.Res.BaseVmClusterId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("consumer_type", s.Res.ConsumerType)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("dns_servers", s.Res.DnsServers)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.InfrastructureId != nil {
		s.D.Set("infrastructure_id", *s.Res.InfrastructureId)
	}

	if s.Res.IsScanEnabled != nil {
		s.D.Set("is_scan_enabled", *s.Res.IsScanEnabled)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.ListenerPort != nil {
		s.D.Set("listener_port", *s.Res.ListenerPort)
	}

	if s.Res.ListenerPortSsl != nil {
		s.D.Set("listener_port_ssl", *s.Res.ListenerPortSsl)
	}

	if s.Res.NodeCount != nil {
		s.D.Set("node_count", *s.Res.NodeCount)
	}

	s.D.Set("ntp_servers", s.Res.NtpServers)

	scans := []interface{}{}
	for _, item := range s.Res.Scans {
		scans = append(scans, ScanDetailsToMap(item))
	}
	s.D.Set("scans", scans)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	vmNetworks := []interface{}{}
	for _, item := range s.Res.VmNetworks {
		vmNetworks = append(vmNetworks, VmNetworkDetailsToMap(item))
	}
	s.D.Set("vm_networks", vmNetworks)

	return nil
}

func (s *DataccVmClusterNetworkResourceCrud) mapToNodeDetails(fieldKeyFormat string) (oci_datacc.NodeDetails, error) {
	result := oci_datacc.NodeDetails{}

	if hostname, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "hostname")); ok {
		tmp := hostname.(string)
		result.Hostname = &tmp
	}

	if ip, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ip")); ok {
		tmp := ip.(string)
		result.Ip = &tmp
	}

	if vip, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vip")); ok && vip != "" {
		tmp := vip.(string)
		result.Vip = &tmp
	}

	if vipHostname, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vip_hostname")); ok && vipHostname != "" {
		tmp := vipHostname.(string)
		result.VipHostname = &tmp
	}

	return result, nil
}

func NodeDetailsToMap(obj oci_datacc.NodeDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Hostname != nil {
		result["hostname"] = string(*obj.Hostname)
	}

	if obj.Ip != nil {
		result["ip"] = string(*obj.Ip)
	}

	if obj.Vip != nil {
		result["vip"] = string(*obj.Vip)
	}

	if obj.VipHostname != nil {
		result["vip_hostname"] = string(*obj.VipHostname)
	}

	return result
}

func (s *DataccVmClusterNetworkResourceCrud) mapToScanDetails(fieldKeyFormat string) (oci_datacc.ScanDetails, error) {
	result := oci_datacc.ScanDetails{}

	if hostname, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "hostname")); ok {
		tmp := hostname.(string)
		result.Hostname = &tmp
	}

	if ips, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ips")); ok {
		interfaces := ips.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "ips")) {
			result.Ips = tmp
		}
	}

	return result, nil
}

func ScanDetailsToMap(obj oci_datacc.ScanDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Hostname != nil {
		result["hostname"] = string(*obj.Hostname)
	}

	result["ips"] = obj.Ips

	return result
}

func VmClusterNetworkSummaryToMap(obj oci_datacc.VmClusterNetworkSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AssociatedResourceId != nil {
		result["associated_resource_id"] = string(*obj.AssociatedResourceId)
	}

	if obj.BaseVmClusterId != nil {
		result["base_vm_cluster_id"] = string(*obj.BaseVmClusterId)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	result["consumer_type"] = string(obj.ConsumerType)

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["dns_servers"] = obj.DnsServers

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.InfrastructureId != nil {
		result["infrastructure_id"] = string(*obj.InfrastructureId)
	}

	if obj.IsScanEnabled != nil {
		result["is_scan_enabled"] = bool(*obj.IsScanEnabled)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.ListenerPort != nil {
		result["listener_port"] = int(*obj.ListenerPort)
	}

	if obj.ListenerPortSsl != nil {
		result["listener_port_ssl"] = int(*obj.ListenerPortSsl)
	}

	if obj.NodeCount != nil {
		result["node_count"] = int(*obj.NodeCount)
	}

	result["ntp_servers"] = obj.NtpServers

	scans := []interface{}{}
	for _, item := range obj.Scans {
		scans = append(scans, ScanDetailsToMap(item))
	}
	result["scans"] = scans

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

	vmNetworks := []interface{}{}
	for _, item := range obj.VmNetworks {
		vmNetworks = append(vmNetworks, VmNetworkDetailsToMap(item))
	}
	result["vm_networks"] = vmNetworks

	return result
}

func (s *DataccVmClusterNetworkResourceCrud) mapToVmNetworkDetails(fieldKeyFormat string) (oci_datacc.VmNetworkDetails, error) {
	result := oci_datacc.VmNetworkDetails{}

	if domainName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "domain_name")); ok {
		tmp := domainName.(string)
		result.DomainName = &tmp
	}

	if gateway, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "gateway")); ok {
		tmp := gateway.(string)
		result.Gateway = &tmp
	}

	if netmask, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "netmask")); ok {
		tmp := netmask.(string)
		result.Netmask = &tmp
	}

	if networkType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "network_type")); ok {
		result.NetworkType = oci_datacc.VmClusterNetworkTypeEnum(networkType.(string))
	}

	if nodes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "nodes")); ok {
		interfaces := nodes.([]interface{})
		tmp := make([]oci_datacc.NodeDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "nodes"), stateDataIndex)
			converted, err := s.mapToNodeDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "nodes")) {
			result.Nodes = tmp
		}
	}

	if prefix, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "prefix")); ok && prefix != "" {
		tmp := prefix.(string)
		result.Prefix = &tmp
	}

	if vlanId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vlan_id")); ok && vlanId != "" {
		tmp := vlanId.(string)
		result.VlanId = &tmp
	}

	return result, nil
}

func VmNetworkDetailsToMap(obj oci_datacc.VmNetworkDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DomainName != nil {
		result["domain_name"] = string(*obj.DomainName)
	}

	if obj.Gateway != nil {
		result["gateway"] = string(*obj.Gateway)
	}

	if obj.Netmask != nil {
		result["netmask"] = string(*obj.Netmask)
	}

	result["network_type"] = string(obj.NetworkType)

	nodes := []interface{}{}
	for _, item := range obj.Nodes {
		nodes = append(nodes, NodeDetailsToMap(item))
	}
	result["nodes"] = nodes

	if obj.Prefix != nil {
		result["prefix"] = string(*obj.Prefix)
	}

	if obj.VlanId != nil {
		result["vlan_id"] = string(*obj.VlanId)
	}

	return result
}

func (s *DataccVmClusterNetworkResourceCrud) updateCompartment(ctx context.Context, compartment interface{}) error {
	changeCompartmentRequest := oci_datacc.ChangeVmClusterNetworkCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.VmClusterNetworkId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datacc")

	response, err := s.Client.ChangeVmClusterNetworkCompartment(ctx, changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getVmClusterNetworkFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datacc"), oci_datacc.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
