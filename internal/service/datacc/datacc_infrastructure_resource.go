// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package datacc

import (
	"context"
	"fmt"
	"strconv"
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

func DataccInfrastructureResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createDataccInfrastructureWithContext,
		ReadContext:   readDataccInfrastructureWithContext,
		UpdateContext: updateDataccInfrastructureWithContext,
		DeleteContext: deleteDataccInfrastructureWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"cloud_control_plane_server1": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cloud_control_plane_server2": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"dns_servers": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"gateway": {
				Type:     schema.TypeString,
				Required: true,
			},
			"netmask": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ntp_servers": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"shape": {
				Type:     schema.TypeString,
				Required: true,
			},
			"system_model": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"acfs_file_system_storage_in_gbs": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"admin_networkcidr": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"backup_network_bonding_interface": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"backup_network_bonding_mode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"client_network_bonding_interface": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"client_network_bonding_mode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"contacts": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"email": {
							Type:     schema.TypeString,
							Required: true,
						},
						"is_primary": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"is_contact_mos_validated": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"phone_number": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"corporate_proxy": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cps_network_bonding_interface": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cps_network_bonding_mode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"data_disk_percentage": {
				Type:     schema.TypeInt,
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
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"maintenance_window": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"custom_action_timeout_in_mins": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"days_of_week": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"hours_of_day": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 20,
							MinItems: 0,
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
						},
						"is_custom_action_timeout_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"is_monthly_patching_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"lead_time_in_weeks": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"months": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"patching_mode": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"preference": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"weeks_of_month": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 4,
							MinItems: 1,
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
						},

						// Computed
					},
				},
			},
			"vlan_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scale_storage_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			// Computed
			"acfs_file_system_used_storage_in_gbs": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"compute_capacity": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"allocated_cores": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"available_cores": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"available_memory_in_gbs": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"reserved_cores": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"reserved_memory_in_gbs": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"total_cores": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"total_memory_in_gbs": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"used_memory_in_gbs": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"lifecycle_state_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"network_adapter_configuration": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"rack_serial_number": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"reco_disk_percentage": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"servers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"base_vm_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"compute_capacity": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"allocated_cores": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"available_cores": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"available_memory_in_gbs": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"reserved_cores": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"reserved_memory_in_gbs": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"total_cores": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"total_memory_in_gbs": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"used_memory_in_gbs": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ilom_ip_address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ilom_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"instance_vm_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"server_ip_address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"server_name": {
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
			"ssd_configuration_requested": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"storage_capacity": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"disk_group": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"logical_free_space_in_gbs": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"phy_free_space_in_gbs": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"phy_reserved_space_in_gbs": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"phy_total_space_in_gbs": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"storage_disk_redundancy": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"subscription_plan_number": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_storage_capacity": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"acfs": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"free_space_in_gbs": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"total_space_in_gbs": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
								},
							},
						},
						"disk_groups": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"free_space_in_gbs": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"reserved_space_in_gbs": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"storage_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"total_space_in_gbs": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"time_activated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_last_state_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_validated": {
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

func createDataccInfrastructureWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataccInfrastructureResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BaseinfraClient()

	if e := tfresource.CreateResourceWithContext(ctx, d, sync); e != nil {
		return tfresource.HandleDiagError(m, e)
	}

	if _, ok := sync.D.GetOkExists("scale_storage_trigger"); ok {
		err := sync.ScaleInfrastructureStorage(ctx)
		if err != nil {
			return tfresource.HandleDiagError(m, err)
		}
	}
	return nil

}

func readDataccInfrastructureWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataccInfrastructureResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BaseinfraClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateDataccInfrastructureWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataccInfrastructureResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BaseinfraClient()

	if _, ok := sync.D.GetOkExists("scale_storage_trigger"); ok && sync.D.HasChange("scale_storage_trigger") {
		oldRaw, newRaw := sync.D.GetChange("scale_storage_trigger")
		oldValue := oldRaw.(int)
		newValue := newRaw.(int)
		if oldValue < newValue {
			err := sync.ScaleInfrastructureStorage(ctx)

			if err != nil {
				return tfresource.HandleDiagError(m, err)
			}
		} else {
			sync.D.Set("scale_storage_trigger", oldRaw)
			err := fmt.Errorf("new value of trigger should be greater than the old value")
			return tfresource.HandleDiagError(m, err)
		}
	}

	if err := tfresource.UpdateResourceWithContext(ctx, d, sync); err != nil {
		return tfresource.HandleDiagError(m, err)
	}

	return nil
}

func deleteDataccInfrastructureWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataccInfrastructureResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BaseinfraClient()
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type DataccInfrastructureResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_datacc.BaseinfraClient
	Res                    *oci_datacc.Infrastructure
	DisableNotFoundRetries bool
}

func (s *DataccInfrastructureResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DataccInfrastructureResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_datacc.InfrastructureLifecycleStateCreating),
		string(oci_datacc.InfrastructureLifecycleStateValidating),
		string(oci_datacc.InfrastructureLifecycleStateActivating),
	}
}

func (s *DataccInfrastructureResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_datacc.InfrastructureLifecycleStateRequiresValidation),
		string(oci_datacc.InfrastructureLifecycleStateRequiresActivation),
		string(oci_datacc.InfrastructureLifecycleStateActive),
	}
}

func (s *DataccInfrastructureResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_datacc.InfrastructureLifecycleStateDeleting),
	}
}

func (s *DataccInfrastructureResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_datacc.InfrastructureLifecycleStateDeleted),
	}
}

func (s *DataccInfrastructureResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_datacc.CreateInfrastructureRequest{}

	if acfsFileSystemStorageInGbs, ok := s.D.GetOkExists("acfs_file_system_storage_in_gbs"); ok {
		tmp := acfsFileSystemStorageInGbs.(float64)
		request.AcfsFileSystemStorageInGbs = &tmp
	}

	if adminNetworkcidr, ok := s.D.GetOkExists("admin_networkcidr"); ok {
		tmp := adminNetworkcidr.(string)
		request.AdminNetworkcidr = &tmp
	}

	if backupNetworkBondingInterface, ok := s.D.GetOkExists("backup_network_bonding_interface"); ok {
		request.BackupNetworkBondingInterface = oci_datacc.NetworkBondingInterfaceEnumEnum(backupNetworkBondingInterface.(string))
	}

	if backupNetworkBondingMode, ok := s.D.GetOkExists("backup_network_bonding_mode"); ok {
		request.BackupNetworkBondingMode = oci_datacc.NetworkBondingModeEnum(backupNetworkBondingMode.(string))
	}

	if clientNetworkBondingInterface, ok := s.D.GetOkExists("client_network_bonding_interface"); ok {
		request.ClientNetworkBondingInterface = oci_datacc.NetworkBondingInterfaceEnumEnum(clientNetworkBondingInterface.(string))
	}

	if clientNetworkBondingMode, ok := s.D.GetOkExists("client_network_bonding_mode"); ok {
		request.ClientNetworkBondingMode = oci_datacc.NetworkBondingModeEnum(clientNetworkBondingMode.(string))
	}

	if cloudControlPlaneServer1, ok := s.D.GetOkExists("cloud_control_plane_server1"); ok {
		tmp := cloudControlPlaneServer1.(string)
		request.CloudControlPlaneServer1 = &tmp
	}

	if cloudControlPlaneServer2, ok := s.D.GetOkExists("cloud_control_plane_server2"); ok {
		tmp := cloudControlPlaneServer2.(string)
		request.CloudControlPlaneServer2 = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if contacts, ok := s.D.GetOkExists("contacts"); ok {
		interfaces := contacts.([]interface{})
		tmp := make([]oci_datacc.InfrastructureContact, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "contacts", stateDataIndex)
			converted, err := s.mapToInfrastructureContact(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("contacts") {
			request.Contacts = tmp
		}
	}

	if corporateProxy, ok := s.D.GetOkExists("corporate_proxy"); ok {
		tmp := corporateProxy.(string)
		request.CorporateProxy = &tmp
	}

	if cpsNetworkBondingInterface, ok := s.D.GetOkExists("cps_network_bonding_interface"); ok {
		request.CpsNetworkBondingInterface = oci_datacc.NetworkBondingInterfaceEnumEnum(cpsNetworkBondingInterface.(string))
	}

	if cpsNetworkBondingMode, ok := s.D.GetOkExists("cps_network_bonding_mode"); ok {
		request.CpsNetworkBondingMode = oci_datacc.NetworkBondingModeEnum(cpsNetworkBondingMode.(string))
	}

	if dataDiskPercentage, ok := s.D.GetOkExists("data_disk_percentage"); ok {
		tmp := dataDiskPercentage.(int)
		request.DataDiskPercentage = &tmp
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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if gateway, ok := s.D.GetOkExists("gateway"); ok {
		tmp := gateway.(string)
		request.Gateway = &tmp
	}

	if maintenanceWindow, ok := s.D.GetOkExists("maintenance_window"); ok {
		if tmpList := maintenanceWindow.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "maintenance_window", 0)
			tmp, err := s.mapToMaintenanceWindow(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.MaintenanceWindow = &tmp
		}
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

	if shape, ok := s.D.GetOkExists("shape"); ok {
		request.Shape = oci_datacc.ShapeEnumEnum(shape.(string))
	}

	if systemModel, ok := s.D.GetOkExists("system_model"); ok {
		request.SystemModel = oci_datacc.SystemModelEnumEnum(systemModel.(string))
	}

	if vlanId, ok := s.D.GetOkExists("vlan_id"); ok {
		tmp := vlanId.(string)
		request.VlanId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datacc")

	response, err := s.Client.CreateInfrastructure(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getInfrastructureFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datacc"), oci_datacc.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DataccInfrastructureResourceCrud) getInfrastructureFromWorkRequest(ctx context.Context, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_datacc.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	infrastructureId, err := infrastructureWaitForWorkRequest(ctx, workId, "infrastructure",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*infrastructureId)

	return s.GetWithContext(ctx)
}

func infrastructureWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func infrastructureWaitForWorkRequest(ctx context.Context, wId *string, entityType string, action oci_datacc.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_datacc.BaseinfraClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "datacc")
	retryPolicy.ShouldRetryOperation = infrastructureWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromDataccInfrastructureWorkRequest(ctx, client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDataccInfrastructureWorkRequest(ctx context.Context, client *oci_datacc.BaseinfraClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_datacc.ActionTypeEnum) error {
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

func (s *DataccInfrastructureResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_datacc.GetInfrastructureRequest{}

	tmp := s.D.Id()
	request.InfrastructureId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datacc")

	response, err := s.Client.GetInfrastructure(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.Infrastructure
	return nil
}

func (s *DataccInfrastructureResourceCrud) UpdateWithContext(ctx context.Context) error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(ctx, compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_datacc.UpdateInfrastructureRequest{}

	if acfsFileSystemStorageInGbs, ok := s.D.GetOkExists("acfs_file_system_storage_in_gbs"); ok {
		tmp := acfsFileSystemStorageInGbs.(float64)
		request.AcfsFileSystemStorageInGbs = &tmp
	}

	if adminNetworkcidr, ok := s.D.GetOkExists("admin_networkcidr"); ok {
		tmp := adminNetworkcidr.(string)
		request.AdminNetworkcidr = &tmp
	}

	if backupNetworkBondingInterface, ok := s.D.GetOkExists("backup_network_bonding_interface"); ok {
		request.BackupNetworkBondingInterface = oci_datacc.NetworkBondingInterfaceEnumEnum(backupNetworkBondingInterface.(string))
	}

	if backupNetworkBondingMode, ok := s.D.GetOkExists("backup_network_bonding_mode"); ok {
		request.BackupNetworkBondingMode = oci_datacc.NetworkBondingModeEnum(backupNetworkBondingMode.(string))
	}

	if clientNetworkBondingInterface, ok := s.D.GetOkExists("client_network_bonding_interface"); ok {
		request.ClientNetworkBondingInterface = oci_datacc.NetworkBondingInterfaceEnumEnum(clientNetworkBondingInterface.(string))
	}

	if clientNetworkBondingMode, ok := s.D.GetOkExists("client_network_bonding_mode"); ok {
		request.ClientNetworkBondingMode = oci_datacc.NetworkBondingModeEnum(clientNetworkBondingMode.(string))
	}

	if cloudControlPlaneServer1, ok := s.D.GetOkExists("cloud_control_plane_server1"); ok {
		tmp := cloudControlPlaneServer1.(string)
		request.CloudControlPlaneServer1 = &tmp
	}

	if cloudControlPlaneServer2, ok := s.D.GetOkExists("cloud_control_plane_server2"); ok {
		tmp := cloudControlPlaneServer2.(string)
		request.CloudControlPlaneServer2 = &tmp
	}

	if contacts, ok := s.D.GetOkExists("contacts"); ok {
		interfaces := contacts.([]interface{})
		tmp := make([]oci_datacc.InfrastructureContact, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "contacts", stateDataIndex)
			converted, err := s.mapToInfrastructureContact(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("contacts") {
			request.Contacts = tmp
		}
	}

	if corporateProxy, ok := s.D.GetOkExists("corporate_proxy"); ok {
		tmp := corporateProxy.(string)
		request.CorporateProxy = &tmp
	}

	if cpsNetworkBondingInterface, ok := s.D.GetOkExists("cps_network_bonding_interface"); ok {
		request.CpsNetworkBondingInterface = oci_datacc.NetworkBondingInterfaceEnumEnum(cpsNetworkBondingInterface.(string))
	}

	if cpsNetworkBondingMode, ok := s.D.GetOkExists("cps_network_bonding_mode"); ok {
		request.CpsNetworkBondingMode = oci_datacc.NetworkBondingModeEnum(cpsNetworkBondingMode.(string))
	}

	if dataDiskPercentage, ok := s.D.GetOkExists("data_disk_percentage"); ok {
		tmp := dataDiskPercentage.(int)
		request.DataDiskPercentage = &tmp
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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if gateway, ok := s.D.GetOkExists("gateway"); ok {
		tmp := gateway.(string)
		request.Gateway = &tmp
	}

	tmp := s.D.Id()
	request.InfrastructureId = &tmp

	if maintenanceWindow, ok := s.D.GetOkExists("maintenance_window"); ok {
		if tmpList := maintenanceWindow.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "maintenance_window", 0)
			tmp, err := s.mapToMaintenanceWindow(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.MaintenanceWindow = &tmp
		}
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

	if shape, ok := s.D.GetOkExists("shape"); ok {
		request.Shape = oci_datacc.ShapeEnumEnum(shape.(string))
	}

	if systemModel, ok := s.D.GetOkExists("system_model"); ok {
		request.SystemModel = oci_datacc.SystemModelEnumEnum(systemModel.(string))
	}

	if vlanId, ok := s.D.GetOkExists("vlan_id"); ok {
		tmp := vlanId.(string)
		request.VlanId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datacc")

	response, err := s.Client.UpdateInfrastructure(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getInfrastructureFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datacc"), oci_datacc.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DataccInfrastructureResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_datacc.DeleteInfrastructureRequest{}

	tmp := s.D.Id()
	request.InfrastructureId = &tmp

	if isForceDeleteInfrastructure, ok := s.D.GetOkExists("is_force_delete_infrastructure"); ok {
		tmp := isForceDeleteInfrastructure.(bool)
		request.IsForceDeleteInfrastructure = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datacc")

	response, err := s.Client.DeleteInfrastructure(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := infrastructureWaitForWorkRequest(ctx, workId, "infrastructure",
		oci_datacc.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *DataccInfrastructureResourceCrud) SetData() error {
	if s.Res.AcfsFileSystemStorageInGbs != nil {
		s.D.Set("acfs_file_system_storage_in_gbs", *s.Res.AcfsFileSystemStorageInGbs)
	}

	if s.Res.AcfsFileSystemUsedStorageInGbs != nil {
		s.D.Set("acfs_file_system_used_storage_in_gbs", *s.Res.AcfsFileSystemUsedStorageInGbs)
	}

	if s.Res.AdminNetworkcidr != nil {
		s.D.Set("admin_networkcidr", *s.Res.AdminNetworkcidr)
	}

	s.D.Set("backup_network_bonding_interface", s.Res.BackupNetworkBondingInterface)

	s.D.Set("backup_network_bonding_mode", s.Res.BackupNetworkBondingMode)

	s.D.Set("client_network_bonding_interface", s.Res.ClientNetworkBondingInterface)

	s.D.Set("client_network_bonding_mode", s.Res.ClientNetworkBondingMode)

	if s.Res.CloudControlPlaneServer1 != nil {
		s.D.Set("cloud_control_plane_server1", *s.Res.CloudControlPlaneServer1)
	}

	if s.Res.CloudControlPlaneServer2 != nil {
		s.D.Set("cloud_control_plane_server2", *s.Res.CloudControlPlaneServer2)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ComputeCapacity != nil {
		s.D.Set("compute_capacity", []interface{}{ComputeCapacityDetailsToMap(s.Res.ComputeCapacity)})
	} else {
		s.D.Set("compute_capacity", nil)
	}

	contacts := []interface{}{}
	for _, item := range s.Res.Contacts {
		contacts = append(contacts, InfrastructureContactToMap(item))
	}
	s.D.Set("contacts", contacts)

	if s.Res.CorporateProxy != nil {
		s.D.Set("corporate_proxy", *s.Res.CorporateProxy)
	}

	s.D.Set("cps_network_bonding_interface", s.Res.CpsNetworkBondingInterface)

	s.D.Set("cps_network_bonding_mode", s.Res.CpsNetworkBondingMode)

	if s.Res.DataDiskPercentage != nil {
		s.D.Set("data_disk_percentage", *s.Res.DataDiskPercentage)
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

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Gateway != nil {
		s.D.Set("gateway", *s.Res.Gateway)
	}

	if s.Res.LifecycleStateDetails != nil {
		s.D.Set("lifecycle_state_details", *s.Res.LifecycleStateDetails)
	}

	if s.Res.MaintenanceWindow != nil {
		s.D.Set("maintenance_window", []interface{}{MaintenanceWindowToMap(s.Res.MaintenanceWindow)})
	} else {
		s.D.Set("maintenance_window", nil)
	}

	if s.Res.Netmask != nil {
		s.D.Set("netmask", *s.Res.Netmask)
	}

	if s.Res.NetworkAdapterConfiguration != nil {
		s.D.Set("network_adapter_configuration", *s.Res.NetworkAdapterConfiguration)
	}

	s.D.Set("ntp_servers", s.Res.NtpServers)

	if s.Res.RackSerialNumber != nil {
		s.D.Set("rack_serial_number", *s.Res.RackSerialNumber)
	}

	if s.Res.RecoDiskPercentage != nil {
		s.D.Set("reco_disk_percentage", *s.Res.RecoDiskPercentage)
	}

	servers := []interface{}{}
	for _, item := range s.Res.Servers {
		servers = append(servers, InfrastructureServerToMap(item))
	}
	s.D.Set("servers", servers)

	s.D.Set("shape", s.Res.Shape)

	s.D.Set("ssd_configuration_requested", s.Res.SsdConfigurationRequested)

	s.D.Set("state", s.Res.LifecycleState)

	storageCapacity := []interface{}{}
	for _, item := range s.Res.StorageCapacity {
		storageCapacity = append(storageCapacity, StorageCapacityDetailsToMap(item))
	}
	s.D.Set("storage_capacity", storageCapacity)

	if s.Res.SubscriptionPlanNumber != nil {
		s.D.Set("subscription_plan_number", *s.Res.SubscriptionPlanNumber)
	}

	s.D.Set("system_model", s.Res.SystemModel)

	if s.Res.SystemStorageCapacity != nil {
		s.D.Set("system_storage_capacity", []interface{}{SystemStorageCapacityDetailsToMap(s.Res.SystemStorageCapacity)})
	} else {
		s.D.Set("system_storage_capacity", nil)
	}

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeActivated != nil {
		s.D.Set("time_activated", s.Res.TimeActivated.String())
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeLastStateUpdated != nil {
		s.D.Set("time_last_state_updated", s.Res.TimeLastStateUpdated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.TimeValidated != nil {
		s.D.Set("time_validated", s.Res.TimeValidated.String())
	}

	if s.Res.Version != nil {
		s.D.Set("version", *s.Res.Version)
	}

	if s.Res.VlanId != nil {
		s.D.Set("vlan_id", *s.Res.VlanId)
	}

	return nil
}

func (s *DataccInfrastructureResourceCrud) ScaleInfrastructureStorage(ctx context.Context) error {
	request := oci_datacc.ScaleInfrastructureStorageRequest{}

	idTmp := s.D.Id()
	request.InfrastructureId = &idTmp

	if ssdConfigurationRequested, ok := s.D.GetOkExists("ssd_configuration_requested"); ok {
		request.SsdConfigurationRequested = oci_datacc.ShapeEnumEnum(ssdConfigurationRequested.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datacc")

	response, err := s.Client.ScaleInfrastructureStorage(ctx, request)
	if err != nil {
		return err
	}
	var workRequestId *string
	if response.RawResponse != nil {
		if wr := response.RawResponse.Header.Get("opc-work-request-id"); wr != "" {
			workRequestId = &wr
		}
	}
	if workRequestId != nil {
		_, err = infrastructureWaitForWorkRequest(ctx, workRequestId, "infrastructure", oci_datacc.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.Client)
		if err != nil {
			return err
		}
	}

	if waitErr := tfresource.WaitForUpdatedStateWithContext(ctx, s.D, s); waitErr != nil {
		return waitErr
	}

	val := s.D.Get("scale_storage_trigger")
	s.D.Set("scale_storage_trigger", val)

	return nil
}

func AcfsCapacityDetailsToMap(obj *oci_datacc.AcfsCapacityDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.FreeSpaceInGbs != nil {
		result["free_space_in_gbs"] = float64(*obj.FreeSpaceInGbs)
	}

	if obj.TotalSpaceInGbs != nil {
		result["total_space_in_gbs"] = float64(*obj.TotalSpaceInGbs)
	}

	return result
}

func ComputeCapacityDetailsToMap(obj *oci_datacc.ComputeCapacityDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AllocatedCores != nil {
		result["allocated_cores"] = int(*obj.AllocatedCores)
	}

	if obj.AvailableCores != nil {
		result["available_cores"] = int(*obj.AvailableCores)
	}

	if obj.AvailableMemoryInGBs != nil {
		result["available_memory_in_gbs"] = strconv.FormatInt(*obj.AvailableMemoryInGBs, 10)
	}

	if obj.ReservedCores != nil {
		result["reserved_cores"] = int(*obj.ReservedCores)
	}

	if obj.ReservedMemoryInGBs != nil {
		result["reserved_memory_in_gbs"] = strconv.FormatInt(*obj.ReservedMemoryInGBs, 10)
	}

	if obj.TotalCores != nil {
		result["total_cores"] = int(*obj.TotalCores)
	}

	if obj.TotalMemoryInGBs != nil {
		result["total_memory_in_gbs"] = strconv.FormatInt(*obj.TotalMemoryInGBs, 10)
	}

	if obj.UsedMemoryInGBs != nil {
		result["used_memory_in_gbs"] = strconv.FormatInt(*obj.UsedMemoryInGBs, 10)
	}

	return result
}

func DiskGroupCapacityDetailsToMap(obj oci_datacc.DiskGroupCapacityDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.FreeSpaceInGbs != nil {
		result["free_space_in_gbs"] = float64(*obj.FreeSpaceInGbs)
	}

	if obj.ReservedSpaceInGbs != nil {
		result["reserved_space_in_gbs"] = float64(*obj.ReservedSpaceInGbs)
	}

	if obj.StorageType != nil {
		result["storage_type"] = string(*obj.StorageType)
	}

	if obj.TotalSpaceInGbs != nil {
		result["total_space_in_gbs"] = float64(*obj.TotalSpaceInGbs)
	}

	return result
}

func (s *DataccInfrastructureResourceCrud) mapToInfrastructureContact(fieldKeyFormat string) (oci_datacc.InfrastructureContact, error) {
	result := oci_datacc.InfrastructureContact{}

	if email, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "email")); ok {
		tmp := email.(string)
		result.Email = &tmp
	}

	if isContactMosValidated, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_contact_mos_validated")); ok {
		tmp := isContactMosValidated.(bool)
		result.IsContactMosValidated = &tmp
	}

	if isPrimary, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_primary")); ok {
		tmp := isPrimary.(bool)
		result.IsPrimary = &tmp
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if phoneNumber, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "phone_number")); ok {
		tmp := phoneNumber.(string)
		result.PhoneNumber = &tmp
	}

	return result, nil
}

func InfrastructureContactToMap(obj oci_datacc.InfrastructureContact) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Email != nil {
		result["email"] = string(*obj.Email)
	}

	if obj.IsContactMosValidated != nil {
		result["is_contact_mos_validated"] = bool(*obj.IsContactMosValidated)
	}

	if obj.IsPrimary != nil {
		result["is_primary"] = bool(*obj.IsPrimary)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.PhoneNumber != nil {
		result["phone_number"] = string(*obj.PhoneNumber)
	}

	return result
}

func InfrastructureServerToMap(obj oci_datacc.InfrastructureServer) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BaseVmCount != nil {
		result["base_vm_count"] = int(*obj.BaseVmCount)
	}

	if obj.ComputeCapacity != nil {
		result["compute_capacity"] = []interface{}{ComputeCapacityDetailsToMap(obj.ComputeCapacity)}
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IlomIpAddress != nil {
		result["ilom_ip_address"] = string(*obj.IlomIpAddress)
	}

	if obj.IlomName != nil {
		result["ilom_name"] = string(*obj.IlomName)
	}

	if obj.InstanceVmCount != nil {
		result["instance_vm_count"] = int(*obj.InstanceVmCount)
	}

	if obj.ServerIpAddress != nil {
		result["server_ip_address"] = string(*obj.ServerIpAddress)
	}

	if obj.ServerName != nil {
		result["server_name"] = string(*obj.ServerName)
	}

	result["state"] = string(obj.LifecycleState)

	return result
}

func InfrastructureSummaryToMap(obj oci_datacc.InfrastructureSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AcfsFileSystemStorageInGbs != nil {
		result["acfs_file_system_storage_in_gbs"] = float64(*obj.AcfsFileSystemStorageInGbs)
	}

	if obj.AcfsFileSystemUsedStorageInGbs != nil {
		result["acfs_file_system_used_storage_in_gbs"] = float64(*obj.AcfsFileSystemUsedStorageInGbs)
	}

	if obj.AdminNetworkcidr != nil {
		result["admin_networkcidr"] = string(*obj.AdminNetworkcidr)
	}

	result["backup_network_bonding_interface"] = string(obj.BackupNetworkBondingInterface)

	result["backup_network_bonding_mode"] = string(obj.BackupNetworkBondingMode)

	result["client_network_bonding_interface"] = string(obj.ClientNetworkBondingInterface)

	result["client_network_bonding_mode"] = string(obj.ClientNetworkBondingMode)

	if obj.CloudControlPlaneServer1 != nil {
		result["cloud_control_plane_server1"] = string(*obj.CloudControlPlaneServer1)
	}

	if obj.CloudControlPlaneServer2 != nil {
		result["cloud_control_plane_server2"] = string(*obj.CloudControlPlaneServer2)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.ComputeCapacity != nil {
		result["compute_capacity"] = []interface{}{ComputeCapacityDetailsToMap(obj.ComputeCapacity)}
	}

	contacts := []interface{}{}
	for _, item := range obj.Contacts {
		contacts = append(contacts, InfrastructureContactToMap(item))
	}
	result["contacts"] = contacts

	if obj.CorporateProxy != nil {
		result["corporate_proxy"] = string(*obj.CorporateProxy)
	}

	result["cps_network_bonding_interface"] = string(obj.CpsNetworkBondingInterface)

	result["cps_network_bonding_mode"] = string(obj.CpsNetworkBondingMode)

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["dns_servers"] = obj.DnsServers

	result["freeform_tags"] = obj.FreeformTags

	if obj.Gateway != nil {
		result["gateway"] = string(*obj.Gateway)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleStateDetails != nil {
		result["lifecycle_state_details"] = string(*obj.LifecycleStateDetails)
	}

	if obj.MaintenanceWindow != nil {
		result["maintenance_window"] = []interface{}{MaintenanceWindowToMap(obj.MaintenanceWindow)}
	}

	if obj.Netmask != nil {
		result["netmask"] = string(*obj.Netmask)
	}

	result["ntp_servers"] = obj.NtpServers

	result["shape"] = string(obj.Shape)

	result["ssd_configuration_requested"] = string(obj.SsdConfigurationRequested)

	result["state"] = string(obj.LifecycleState)

	storageCapacity := []interface{}{}
	for _, item := range obj.StorageCapacity {
		storageCapacity = append(storageCapacity, StorageCapacityDetailsToMap(item))
	}
	result["storage_capacity"] = storageCapacity

	if obj.SubscriptionPlanNumber != nil {
		result["subscription_plan_number"] = string(*obj.SubscriptionPlanNumber)
	}

	result["system_model"] = string(obj.SystemModel)

	if obj.SystemStorageCapacity != nil {
		result["system_storage_capacity"] = []interface{}{SystemStorageCapacityDetailsToMap(obj.SystemStorageCapacity)}
	}

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeActivated != nil {
		result["time_activated"] = obj.TimeActivated.String()
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeLastStateUpdated != nil {
		result["time_last_state_updated"] = obj.TimeLastStateUpdated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	if obj.TimeValidated != nil {
		result["time_validated"] = obj.TimeValidated.String()
	}

	if obj.Version != nil {
		result["version"] = string(*obj.Version)
	}

	return result
}

func (s *DataccInfrastructureResourceCrud) mapToMaintenanceWindow(fieldKeyFormat string) (oci_datacc.MaintenanceWindow, error) {
	result := oci_datacc.MaintenanceWindow{}

	if customActionTimeoutInMins, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "custom_action_timeout_in_mins")); ok {
		tmp := customActionTimeoutInMins.(int)
		result.CustomActionTimeoutInMins = &tmp
	}

	if daysOfWeek, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "days_of_week")); ok {
		interfaces := daysOfWeek.([]interface{})
		tmp := make([]oci_datacc.MaintenanceWindowDaysOfWeekEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_datacc.MaintenanceWindowDaysOfWeekEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "days_of_week")) {
			result.DaysOfWeek = tmp
		}
	}

	if hoursOfDay, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "hours_of_day")); ok {
		interfaces := hoursOfDay.([]interface{})
		tmp := make([]int, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(int)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "hours_of_day")) {
			result.HoursOfDay = tmp
		}
	}

	if isCustomActionTimeoutEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_custom_action_timeout_enabled")); ok {
		tmp := isCustomActionTimeoutEnabled.(bool)
		result.IsCustomActionTimeoutEnabled = &tmp
	}

	if isMonthlyPatchingEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_monthly_patching_enabled")); ok {
		tmp := isMonthlyPatchingEnabled.(bool)
		result.IsMonthlyPatchingEnabled = &tmp
	}

	if leadTimeInWeeks, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "lead_time_in_weeks")); ok {
		tmp := leadTimeInWeeks.(int)
		result.LeadTimeInWeeks = &tmp
	}

	if months, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "months")); ok {
		interfaces := months.([]interface{})
		tmp := make([]oci_datacc.MaintenanceWindowMonthsEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_datacc.MaintenanceWindowMonthsEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "months")) {
			result.Months = tmp
		}
	}

	if patchingMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "patching_mode")); ok {
		result.PatchingMode = oci_datacc.BasePatchingModeEnum(patchingMode.(string))
	}

	if preference, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "preference")); ok {
		result.Preference = oci_datacc.MaintenancePreferenceEnum(preference.(string))
	}

	if weeksOfMonth, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "weeks_of_month")); ok {
		interfaces := weeksOfMonth.([]interface{})
		tmp := make([]int, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(int)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "weeks_of_month")) {
			result.WeeksOfMonth = tmp
		}
	}

	return result, nil
}

func MaintenanceWindowToMap(obj *oci_datacc.MaintenanceWindow) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CustomActionTimeoutInMins != nil {
		result["custom_action_timeout_in_mins"] = int(*obj.CustomActionTimeoutInMins)
	}

	result["days_of_week"] = obj.DaysOfWeek

	result["hours_of_day"] = obj.HoursOfDay

	if obj.IsCustomActionTimeoutEnabled != nil {
		result["is_custom_action_timeout_enabled"] = bool(*obj.IsCustomActionTimeoutEnabled)
	}

	if obj.IsMonthlyPatchingEnabled != nil {
		result["is_monthly_patching_enabled"] = bool(*obj.IsMonthlyPatchingEnabled)
	}

	if obj.LeadTimeInWeeks != nil {
		result["lead_time_in_weeks"] = int(*obj.LeadTimeInWeeks)
	}

	result["months"] = obj.Months

	result["patching_mode"] = string(obj.PatchingMode)

	result["preference"] = string(obj.Preference)

	result["weeks_of_month"] = obj.WeeksOfMonth

	return result
}

func StorageCapacityDetailsToMap(obj oci_datacc.StorageCapacityDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DiskGroup != nil {
		result["disk_group"] = string(*obj.DiskGroup)
	}

	if obj.LogicalFreeSpaceInGBs != nil {
		result["logical_free_space_in_gbs"] = strconv.FormatInt(*obj.LogicalFreeSpaceInGBs, 10)
	}

	if obj.PhyFreeSpaceInGBs != nil {
		result["phy_free_space_in_gbs"] = strconv.FormatInt(*obj.PhyFreeSpaceInGBs, 10)
	}

	if obj.PhyReservedSpaceInGBs != nil {
		result["phy_reserved_space_in_gbs"] = strconv.FormatInt(*obj.PhyReservedSpaceInGBs, 10)
	}

	if obj.PhyTotalSpaceInGBs != nil {
		result["phy_total_space_in_gbs"] = strconv.FormatInt(*obj.PhyTotalSpaceInGBs, 10)
	}

	result["storage_disk_redundancy"] = string(obj.StorageDiskRedundancy)

	return result
}

func SystemStorageCapacityDetailsToMap(obj *oci_datacc.SystemStorageCapacityDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Acfs != nil {
		result["acfs"] = []interface{}{AcfsCapacityDetailsToMap(obj.Acfs)}
	}

	diskGroups := []interface{}{}
	for _, item := range obj.DiskGroups {
		diskGroups = append(diskGroups, DiskGroupCapacityDetailsToMap(item))
	}
	result["disk_groups"] = diskGroups

	return result
}

func (s *DataccInfrastructureResourceCrud) updateCompartment(ctx context.Context, compartment interface{}) error {
	changeCompartmentRequest := oci_datacc.ChangeInfrastructureCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.InfrastructureId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datacc")

	response, err := s.Client.ChangeInfrastructureCompartment(ctx, changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getInfrastructureFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datacc"), oci_datacc.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
