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

func DataccVmInstanceResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(40 * time.Minute),
			Update: schema.DefaultTimeout(40 * time.Minute),
			Delete: tfresource.DefaultTimeout.Delete,
		},
		CreateContext: createDataccVmInstanceWithContext,
		ReadContext:   readDataccVmInstanceWithContext,
		UpdateContext: updateDataccVmInstanceWithContext,
		DeleteContext: deleteDataccVmInstanceWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cpus_enabled": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"infrastructure_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ssh_public_keys": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Optional
			"boot_storage_size_in_gbs": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"data_storage_size_in_gb": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_servers": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"domain_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"gateway": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"hostname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"image_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"ip_address": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"memory_size_in_gbs": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"metadata": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem:     schema.TypeString,
			},
			"netmask": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"ntp_servers": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"server_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem:     schema.TypeString,
			},
			"time_zone": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"userdata": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"vlan_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"vm_network_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
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

func createDataccVmInstanceWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataccVmInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BaseinfraClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readDataccVmInstanceWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataccVmInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BaseinfraClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateDataccVmInstanceWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataccVmInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BaseinfraClient()

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func deleteDataccVmInstanceWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataccVmInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BaseinfraClient()
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type DataccVmInstanceResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_datacc.BaseinfraClient
	Res                    *oci_datacc.VmInstance
	DisableNotFoundRetries bool
}

func (s *DataccVmInstanceResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DataccVmInstanceResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_datacc.VmInstanceLifecycleStateCreating),
	}
}

func (s *DataccVmInstanceResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_datacc.VmInstanceLifecycleStateActive),
	}
}

func (s *DataccVmInstanceResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_datacc.VmInstanceLifecycleStateDeleting),
	}
}

func (s *DataccVmInstanceResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_datacc.VmInstanceLifecycleStateDeleted),
	}
}

func (s *DataccVmInstanceResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_datacc.CreateVmInstanceRequest{}

	if bootStorageSizeInGBs, ok := s.D.GetOkExists("boot_storage_size_in_gbs"); ok {
		tmp := bootStorageSizeInGBs.(float64)
		request.BootStorageSizeInGBs = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if cpusEnabled, ok := s.D.GetOkExists("cpus_enabled"); ok {
		tmp := cpusEnabled.(int)
		request.CpusEnabled = &tmp
	}

	if dataStorageSizeInGB, ok := s.D.GetOkExists("data_storage_size_in_gb"); ok {
		tmp := dataStorageSizeInGB.(float64)
		request.DataStorageSizeInGBs = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
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

	if domainName, ok := s.D.GetOkExists("domain_name"); ok {
		tmp := domainName.(string)
		request.DomainName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if gateway, ok := s.D.GetOkExists("gateway"); ok {
		tmp := gateway.(string)
		request.Gateway = &tmp
	}

	if hostname, ok := s.D.GetOkExists("hostname"); ok {
		tmp := hostname.(string)
		request.Hostname = &tmp
	}

	if imageId, ok := s.D.GetOkExists("image_id"); ok {
		tmp := imageId.(string)
		request.ImageId = &tmp
	}

	if infrastructureId, ok := s.D.GetOkExists("infrastructure_id"); ok {
		tmp := infrastructureId.(string)
		request.InfrastructureId = &tmp
	}

	if ipAddress, ok := s.D.GetOkExists("ip_address"); ok {
		tmp := ipAddress.(string)
		request.IpAddress = &tmp
	}

	if memorySizeInGBs, ok := s.D.GetOkExists("memory_size_in_gbs"); ok {
		tmp := memorySizeInGBs.(float64)
		request.MemorySizeInGBs = &tmp
	}

	if metadata, ok := s.D.GetOkExists("metadata"); ok {
		request.Metadata = tfresource.ObjectMapToStringMap(metadata.(map[string]interface{}))
	}

	if netmask, ok := s.D.GetOkExists("netmask"); ok {
		tmp := netmask.(string)
		request.Netmask = &tmp
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

	if serverId, ok := s.D.GetOkExists("server_id"); ok {
		tmp := serverId.(string)
		request.ServerId = &tmp
	}

	if sshPublicKeys, ok := s.D.GetOkExists("ssh_public_keys"); ok {
		interfaces := sshPublicKeys.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("ssh_public_keys") {
			request.SshPublicKeys = tmp
		}
	}

	if systemTags, ok := s.D.GetOkExists("system_tags"); ok {
		convertedSystemTags, err := tfresource.MapToSystemTags(systemTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.SystemTags = convertedSystemTags
	}

	if timeZone, ok := s.D.GetOkExists("time_zone"); ok {
		tmp := timeZone.(string)
		request.TimeZone = &tmp
	}

	if userdata, ok := s.D.GetOkExists("userdata"); ok {
		tmp := userdata.(string)
		request.Userdata = &tmp
	}

	if vlanId, ok := s.D.GetOkExists("vlan_id"); ok {
		tmp := vlanId.(string)
		request.VlanId = &tmp
	}

	if vmNetworkId, ok := s.D.GetOkExists("vm_network_id"); ok {
		tmp := vmNetworkId.(string)
		request.VmNetworkId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datacc")

	response, err := s.Client.CreateVmInstance(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getVmInstanceFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datacc"), oci_datacc.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DataccVmInstanceResourceCrud) getVmInstanceFromWorkRequest(ctx context.Context, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_datacc.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	vmInstanceId, err := vmInstanceWaitForWorkRequest(ctx, workId, "instance",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*vmInstanceId)

	return s.GetWithContext(ctx)
}

func vmInstanceWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func vmInstanceWaitForWorkRequest(ctx context.Context, wId *string, entityType string, action oci_datacc.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_datacc.BaseinfraClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "datacc")
	retryPolicy.ShouldRetryOperation = vmInstanceWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromDataccVmInstanceWorkRequest(ctx, client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDataccVmInstanceWorkRequest(ctx context.Context, client *oci_datacc.BaseinfraClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_datacc.ActionTypeEnum) error {
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

func (s *DataccVmInstanceResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_datacc.GetVmInstanceRequest{}

	tmp := s.D.Id()
	request.VmInstanceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datacc")

	response, err := s.Client.GetVmInstance(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.VmInstance
	return nil
}

func (s *DataccVmInstanceResourceCrud) UpdateWithContext(ctx context.Context) error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(ctx, compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_datacc.UpdateVmInstanceRequest{}

	if dataStorageSizeInGB, ok := s.D.GetOkExists("data_storage_size_in_gb"); ok {
		tmp := dataStorageSizeInGB.(float64)
		request.DataStorageSizeInGBs = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.VmInstanceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datacc")

	response, err := s.Client.UpdateVmInstance(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getVmInstanceFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datacc"), oci_datacc.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DataccVmInstanceResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_datacc.DeleteVmInstanceRequest{}

	tmp := s.D.Id()
	request.VmInstanceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datacc")

	response, err := s.Client.DeleteVmInstance(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := vmInstanceWaitForWorkRequest(ctx, workId, "instance",
		oci_datacc.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *DataccVmInstanceResourceCrud) SetData() error {
	if s.Res.BootStorageSizeInGBs != nil {
		s.D.Set("boot_storage_size_in_gbs", *s.Res.BootStorageSizeInGBs)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CpusEnabled != nil {
		s.D.Set("cpus_enabled", *s.Res.CpusEnabled)
	}

	if s.Res.DataStorageSizeInGBs != nil {
		s.D.Set("data_storage_size_in_gb", *s.Res.DataStorageSizeInGBs)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("dns_servers", s.Res.DnsServers)

	if s.Res.DomainName != nil {
		s.D.Set("domain_name", *s.Res.DomainName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Gateway != nil {
		s.D.Set("gateway", *s.Res.Gateway)
	}

	if s.Res.Hostname != nil {
		s.D.Set("hostname", *s.Res.Hostname)
	}

	if s.Res.ImageId != nil {
		s.D.Set("image_id", *s.Res.ImageId)
	}

	if s.Res.InfrastructureId != nil {
		s.D.Set("infrastructure_id", *s.Res.InfrastructureId)
	}

	if s.Res.IpAddress != nil {
		s.D.Set("ip_address", *s.Res.IpAddress)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.MemorySizeInGBs != nil {
		s.D.Set("memory_size_in_gbs", *s.Res.MemorySizeInGBs)
	}

	s.D.Set("metadata", s.Res.Metadata)

	if s.Res.Netmask != nil {
		s.D.Set("netmask", *s.Res.Netmask)
	}

	s.D.Set("ntp_servers", s.Res.NtpServers)

	if s.Res.ServerId != nil {
		s.D.Set("server_id", *s.Res.ServerId)
	}

	s.D.Set("ssh_public_keys", s.Res.SshPublicKeys)

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

	if s.Res.TimeZone != nil {
		s.D.Set("time_zone", *s.Res.TimeZone)
	}

	if s.Res.Userdata != nil {
		s.D.Set("userdata", *s.Res.Userdata)
	}

	if s.Res.VlanId != nil {
		s.D.Set("vlan_id", *s.Res.VlanId)
	}

	if s.Res.VmNetworkId != nil {
		s.D.Set("vm_network_id", *s.Res.VmNetworkId)
	}

	return nil
}

func VmInstanceSummaryToMap(obj oci_datacc.VmInstanceSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.CpusEnabled != nil {
		result["cpus_enabled"] = int(*obj.CpusEnabled)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.DomainName != nil {
		result["domain_name"] = string(*obj.DomainName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.ImageId != nil {
		result["image_id"] = string(*obj.ImageId)
	}

	if obj.InfrastructureId != nil {
		result["infrastructure_id"] = string(*obj.InfrastructureId)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.MemorySizeInGBs != nil {
		result["memory_size_in_gbs"] = float64(*obj.MemorySizeInGBs)
	}

	if obj.ServerId != nil {
		result["server_id"] = string(*obj.ServerId)
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

	if obj.TimeZone != nil {
		result["time_zone"] = string(*obj.TimeZone)
	}

	if obj.VmNetworkId != nil {
		result["vm_network_id"] = string(*obj.VmNetworkId)
	}

	return result
}

func (s *DataccVmInstanceResourceCrud) updateCompartment(ctx context.Context, compartment interface{}) error {
	changeCompartmentRequest := oci_datacc.ChangeVmInstanceCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.VmInstanceId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datacc")

	response, err := s.Client.ChangeVmInstanceCompartment(ctx, changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getVmInstanceFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datacc"), oci_datacc.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
