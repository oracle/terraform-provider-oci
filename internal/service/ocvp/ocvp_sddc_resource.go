// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ocvp

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_ocvp "github.com/oracle/oci-go-sdk/v65/ocvp"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OcvpSddcResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("6h"),
			Update: tfresource.GetTimeoutDuration("2h"),
			Delete: tfresource.GetTimeoutDuration("2h"),
		},
		Create: createOcvpSddc,
		Read:   readOcvpSddc,
		Update: updateOcvpSddc,
		Delete: deleteOcvpSddc,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compute_availability_domain": {
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"esxi_software_version", "initial_configuration"},
				Deprecated:    tfresource.FieldDeprecatedForAnother("compute_availability_domain", "initial_configuration"),
			},
			"esxi_hosts_count": {
				Type:          schema.TypeInt,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"esxi_software_version", "initial_configuration"},
				Deprecated:    tfresource.FieldDeprecatedForAnother("esxi_hosts_count", "initial_configuration"),
			},
			"nsx_edge_uplink1vlan_id": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"esxi_software_version", "initial_configuration"},
				Deprecated:    tfresource.FieldDeprecatedForAnother("nsx_edge_uplink1vlan_id", "initial_configuration"),
			},
			"nsx_edge_uplink2vlan_id": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"esxi_software_version", "initial_configuration"},
				Deprecated:    tfresource.FieldDeprecatedForAnother("nsx_edge_uplink2vlan_id", "initial_configuration"),
			},
			"nsx_edge_vtep_vlan_id": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"esxi_software_version", "initial_configuration"},
				Deprecated:    tfresource.FieldDeprecatedForAnother("nsx_edge_vtep_vlan_id", "initial_configuration"),
			},
			"nsx_vtep_vlan_id": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"esxi_software_version", "initial_configuration"},
				Deprecated:    tfresource.FieldDeprecatedForAnother("nsx_vtep_vlan_id", "initial_configuration"),
			},
			"provisioning_subnet_id": {
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"esxi_software_version", "initial_configuration"},
				Deprecated:    tfresource.FieldDeprecatedForAnother("provisioning_subnet_id", "initial_configuration"),
			},
			"ssh_authorized_keys": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vmotion_vlan_id": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"esxi_software_version", "initial_configuration"},
				Deprecated:    tfresource.FieldDeprecatedForAnother("vmotion_vlan_id", "initial_configuration"),
			},
			"vmware_software_version": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vsan_vlan_id": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"esxi_software_version", "initial_configuration"},
				Deprecated:    tfresource.FieldDeprecatedForAnother("vsan_vlan_id", "initial_configuration"),
			},
			"vsphere_vlan_id": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"esxi_software_version", "initial_configuration"},
				Deprecated:    tfresource.FieldDeprecatedForAnother("vsphere_vlan_id", "initial_configuration"),
			},

			// Optional
			"datastores": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"block_volume_ids": {
							Type:     schema.TypeList,
							Required: true,
							ForceNew: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"datastore_type": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional

						// Computed
						"capacity": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
					},
				},
				ConflictsWith: []string{"esxi_software_version", "initial_configuration"},
				Deprecated:    tfresource.FieldDeprecatedForAnother("datastores", "initial_configuration"),
			},
			"capacity_reservation_id": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ForceNew:      true,
				ConflictsWith: []string{"esxi_software_version", "initial_configuration"},
				Deprecated:    tfresource.FieldDeprecatedForAnother("capacity_reservation_id", "initial_configuration"),
			},
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"display_name": {
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
			"hcx_vlan_id": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"esxi_software_version", "initial_configuration"},
				Deprecated:    tfresource.FieldDeprecatedForAnother("hcx_vlan_id", "initial_configuration"),
			},
			"hcx_action": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					UpgradeHcxAction,
					DowngradeHcxAction,
					CancelDowngradeHcxAction,
				}, true),
			},
			"initial_host_ocpu_count": {
				Type:          schema.TypeFloat,
				Optional:      true,
				Computed:      true,
				ForceNew:      true,
				ConflictsWith: []string{"esxi_software_version", "initial_configuration"},
				Deprecated:    tfresource.FieldDeprecatedForAnother("initial_host_ocpu_count", "initial_configuration"),
			},
			"initial_host_shape_name": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ForceNew:      true,
				ConflictsWith: []string{"esxi_software_version", "initial_configuration"},
				Deprecated:    tfresource.FieldDeprecatedForAnother("initial_host_shape_name", "initial_configuration"),
			},
			"initial_sku": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ForceNew:      true,
				ConflictsWith: []string{"esxi_software_version", "initial_configuration"},
				Deprecated:    tfresource.FieldDeprecatedForAnother("initial_sku", "initial_configuration"),
			},
			"instance_display_name_prefix": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ForceNew:      true,
				ConflictsWith: []string{"esxi_software_version", "initial_configuration"},
				Deprecated:    tfresource.FieldDeprecatedForAnother("instance_display_name_prefix", "initial_configuration"),
			},
			"is_hcx_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"is_shielded_instance_enabled": {
				Type:          schema.TypeBool,
				Optional:      true,
				Computed:      true,
				ForceNew:      true,
				ConflictsWith: []string{"esxi_software_version", "initial_configuration"},
				Deprecated:    tfresource.FieldDeprecatedForAnother("is_shielded_instance_enabled", "initial_configuration"),
			},
			"is_single_host_sddc": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"provisioning_vlan_id": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"esxi_software_version", "initial_configuration"},
				Deprecated:    tfresource.FieldDeprecatedForAnother("provisioning_vlan_id", "initial_configuration"),
			},
			"refresh_hcx_license_status": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"replication_vlan_id": {
				Type:       schema.TypeString,
				Optional:   true,
				Computed:   true,
				Deprecated: tfresource.FieldDeprecatedForAnother("replication_vlan_id", "initial_configuration"),
			},
			"reserving_hcx_on_premise_license_keys": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"workload_network_cidr": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ForceNew:      true,
				ConflictsWith: []string{"esxi_software_version", "initial_configuration"},
				Deprecated:    tfresource.FieldDeprecatedForAnother("workload_network_cidr", "initial_configuration"),
			},

			// Computed
			"actual_esxi_hosts_count": {
				Type:       schema.TypeInt,
				Computed:   true,
				Deprecated: tfresource.FieldDeprecatedForAnother("actual_esxi_hosts_count", "initial_configuration"),
			},
			"hcx_fqdn": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hcx_initial_password": {
				Type:       schema.TypeString,
				Computed:   true,
				Deprecated: tfresource.FieldDeprecatedButSupportedThroughAnotherDataSource("hcx_initial_password", "oci_ocvp_retrieve_password"),
			},
			"hcx_on_prem_key": {
				Type:       schema.TypeString,
				Computed:   true,
				Deprecated: tfresource.FieldDeprecatedAndAvoidReferences("hcx_on_prem_key"),
			},
			"hcx_on_prem_licenses": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"activation_key": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"system_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"hcx_private_ip_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_hcx_enterprise_enabled": {
				Type:       schema.TypeBool,
				Computed:   true,
				Deprecated: tfresource.FieldDeprecatedAndAvoidReferences("is_hcx_enterprise_enabled"),
			},
			"is_hcx_pending_downgrade": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"nsx_edge_uplink_ip_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"nsx_manager_fqdn": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"nsx_manager_initial_password": {
				Type:       schema.TypeString,
				Computed:   true,
				Deprecated: tfresource.FieldDeprecatedButSupportedThroughAnotherDataSource("nsx_manager_initial_password", "oci_ocvp_retrieve_password"),
			},
			"nsx_manager_private_ip_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"nsx_manager_username": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"nsx_overlay_segment_name": {
				Type:       schema.TypeString,
				Computed:   true,
				Deprecated: tfresource.FieldDeprecatedAndAvoidReferences("nsx_overlay_segment_name"),
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_hcx_billing_cycle_end": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_hcx_license_status_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"upgrade_licenses": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"license_key": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"license_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
				Deprecated: tfresource.FieldDeprecatedButSupportedThroughAnotherResource("upgrade_licenses", "ocvp_cluster_resource"),
			},
			"vcenter_fqdn": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vcenter_initial_password": {
				Type:       schema.TypeString,
				Computed:   true,
				Deprecated: tfresource.FieldDeprecatedButSupportedThroughAnotherDataSource("vcenter_initial_password", "oci_ocvp_retrieve_password"),
			},
			"vcenter_private_ip_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vcenter_username": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vsphere_upgrade_guide": {
				Type:       schema.TypeString,
				Computed:   true,
				Deprecated: tfresource.FieldDeprecatedAndAvoidReferences("vsphere_upgrade_guide"),
			},
			"vsphere_upgrade_objects": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"download_link": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"link_description": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
				Deprecated: tfresource.FieldDeprecatedButSupportedThroughAnotherResource("vsphere_upgrade_objects", "ocvp_cluster_resource"),
			},
			// New API attributes
			//optional
			"esxi_software_version": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"compute_availability_domain", "esxi_hosts_count", "nsx_edge_uplink1vlan_id", "nsx_edge_uplink2vlan_id", "nsx_edge_vtep_vlan_id", "nsx_vtep_vlan_id", "provisioning_subnet_id", "vmotion_vlan_id", "vsan_vlan_id", "vsphere_vlan_id", "datastores", "capacity_reservation_id", "hcx_vlan_id", "initial_host_ocpu_count", "initial_host_shape_name", "initial_sku", "instance_display_name_prefix", "is_shielded_instance_enabled", "provisioning_vlan_id", "replication_vlan_id", "workload_network_cidr"},
			},
			"initial_configuration": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"initial_cluster_configurations": {
							Type:     schema.TypeList,
							Required: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"compute_availability_domain": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},
									"esxi_hosts_count": {
										Type:     schema.TypeInt,
										Required: true,
										ForceNew: true,
									},
									"network_configuration": {
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"nsx_edge_vtep_vlan_id": {
													Type:     schema.TypeString,
													Required: true,
													ForceNew: true,
												},
												"nsx_vtep_vlan_id": {
													Type:     schema.TypeString,
													Required: true,
													ForceNew: true,
												},
												"provisioning_subnet_id": {
													Type:     schema.TypeString,
													Required: true,
													ForceNew: true,
												},
												"vmotion_vlan_id": {
													Type:     schema.TypeString,
													Required: true,
													ForceNew: true,
												},
												"vsan_vlan_id": {
													Type:     schema.TypeString,
													Required: true,
													ForceNew: true,
												},

												// Optional
												"hcx_vlan_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"nsx_edge_uplink1vlan_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"nsx_edge_uplink2vlan_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"provisioning_vlan_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"replication_vlan_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"vsphere_vlan_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},

												// Computed
											},
										},
									},
									"vsphere_type": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},

									// Optional
									"capacity_reservation_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"datastores": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										ForceNew: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"block_volume_ids": {
													Type:     schema.TypeList,
													Required: true,
													ForceNew: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"datastore_type": {
													Type:     schema.TypeString,
													Required: true,
													ForceNew: true,
												},

												// Optional

												// Computed
											},
										},
									},
									"display_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"initial_commitment": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"initial_host_ocpu_count": {
										Type:     schema.TypeFloat,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"initial_host_shape_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"instance_display_name_prefix": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"is_shielded_instance_enabled": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"workload_network_cidr": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},

									// Computed
									"actual_esxi_hosts_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
					},
				},
				ConflictsWith: []string{"compute_availability_domain", "esxi_hosts_count", "nsx_edge_uplink1vlan_id", "nsx_edge_uplink2vlan_id", "nsx_edge_vtep_vlan_id", "nsx_vtep_vlan_id", "provisioning_subnet_id", "vmotion_vlan_id", "vsan_vlan_id", "vsphere_vlan_id", "datastores", "capacity_reservation_id", "hcx_vlan_id", "initial_host_ocpu_count", "initial_host_shape_name", "initial_sku", "instance_display_name_prefix", "is_shielded_instance_enabled", "provisioning_vlan_id", "replication_vlan_id", "workload_network_cidr"},
			},
			//computed
			"hcx_mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"clusters_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

const (
	UpgradeHcxAction         string = "UPGRADE"
	DowngradeHcxAction       string = "DOWNGRADE"
	CancelDowngradeHcxAction string = "CANCEL_DOWNGRADE"
)

func createOcvpSddc(d *schema.ResourceData, m interface{}) error {
	sync := &OcvpSddcResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SddcClient()
	sync.ClusterClient = m.(*client.OracleClients).ClusterClient()
	sync.WorkRequestClient = m.(*client.OracleClients).OcvpWorkRequestClient()

	return tfresource.CreateResource(d, sync)
}

func readOcvpSddc(d *schema.ResourceData, m interface{}) error {
	sync := &OcvpSddcResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SddcClient()
	sync.ClusterClient = m.(*client.OracleClients).ClusterClient()
	return tfresource.ReadResource(sync)
}

func updateOcvpSddc(d *schema.ResourceData, m interface{}) error {
	sync := &OcvpSddcResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SddcClient()
	sync.WorkRequestClient = m.(*client.OracleClients).OcvpWorkRequestClient()
	sync.ClusterClient = m.(*client.OracleClients).ClusterClient()
	return tfresource.UpdateResource(d, sync)
}

func deleteOcvpSddc(d *schema.ResourceData, m interface{}) error {
	sync := &OcvpSddcResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SddcClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).OcvpWorkRequestClient()
	sync.ClusterClient = m.(*client.OracleClients).ClusterClient()
	return tfresource.DeleteResource(d, sync)
}

type OcvpSddcResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_ocvp.SddcClient
	ClusterClient          *oci_ocvp.ClusterClient
	Res                    *oci_ocvp.Sddc
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_ocvp.WorkRequestClient
}

func (s *OcvpSddcResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *OcvpSddcResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_ocvp.LifecycleStatesCreating),
	}
}

func (s *OcvpSddcResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_ocvp.LifecycleStatesActive),
	}
}

func (s *OcvpSddcResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_ocvp.LifecycleStatesDeleting),
	}
}

func (s *OcvpSddcResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_ocvp.LifecycleStatesDeleted),
	}
}

func (s *OcvpSddcResourceCrud) getOkExistsClusterConfigurationProperty(property string) (interface{}, bool) {
	if initialConfiguration, ok := s.D.GetOkExists("initial_configuration"); ok {
		if tmpList := initialConfiguration.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%s.%d.%%s", "initial_configuration", 0, "initial_cluster_configurations", 0)
			if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, property)); ok {
				return value, ok
			}
		}
	}
	return nil, false
}

func (s *OcvpSddcResourceCrud) getOkExistsNetworkConfigurationProperty(property string) (interface{}, bool) {
	if _, ok := s.getOkExistsClusterConfigurationProperty("network_configuration"); ok {
		fieldKeyFormat := fmt.Sprintf("%s.%d.%s.%d.%s.%d.%%s", "initial_configuration", 0, "initial_cluster_configurations", 0, "network_configuration", 0)
		if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, property)); ok {
			return value, ok
		}
	}
	return nil, false
}

func (s *OcvpSddcResourceCrud) Create() error {
	request := oci_ocvp.CreateSddcRequest{}
	networkConfiguration := oci_ocvp.NetworkConfiguration{}
	initialClusterConfiguration := oci_ocvp.InitialClusterConfiguration{NetworkConfiguration: &networkConfiguration}

	capacityReservationIdDeprecated, okDeprecated := s.D.GetOkExists("capacity_reservation_id")
	capacityReservationId, ok := s.getOkExistsClusterConfigurationProperty("capacity_reservation_id")

	if okDeprecated {
		tmp := capacityReservationIdDeprecated.(string)
		initialClusterConfiguration.CapacityReservationId = &tmp
	} else if ok {
		tmp := capacityReservationId.(string)
		initialClusterConfiguration.CapacityReservationId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	computeAvailabilityDomainDeprecated, okDeprecated := s.D.GetOkExists("compute_availability_domain")
	computeAvailabilityDomain, ok := s.getOkExistsClusterConfigurationProperty("compute_availability_domain")
	if !okDeprecated && !ok {
		return errors.New("one of compute_availability_domain or initial_cluster_configurations.compute_availability_domain must be configured")
	}
	if okDeprecated {
		tmp := computeAvailabilityDomainDeprecated.(string)
		initialClusterConfiguration.ComputeAvailabilityDomain = &tmp
	} else {
		tmp := computeAvailabilityDomain.(string)
		initialClusterConfiguration.ComputeAvailabilityDomain = &tmp
	}

	datastoresDeprecated, okDeprecated := s.D.GetOkExists("datastores")
	datastores, ok := s.getOkExistsClusterConfigurationProperty("datastores")

	if okDeprecated {
		interfaces := datastoresDeprecated.([]interface{})
		tmp := make([]oci_ocvp.DatastoreInfo, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "datastores", stateDataIndex)
			converted, err := s.mapToDatastoreInfo(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("datastores") {
			initialClusterConfiguration.Datastores = tmp
		}
	} else if ok {
		interfaces := datastores.([]interface{})
		tmp := make([]oci_ocvp.DatastoreInfo, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%s.%d.%s.%d.%%s", "initial_configuration", 0, "initial_cluster_configurations", 0, "datastores", stateDataIndex)
			log.Printf("[DEBUG] changing data store keyformat %s", fieldKeyFormat)
			converted, err := s.mapToDatastoreInfo(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf("%s.%d.%s.%d.%s", "initial_configuration", 0, "initial_cluster_configurations", 0, "datastores")) {
			initialClusterConfiguration.Datastores = tmp
			log.Printf("[DEBUG] changing data store %s", datastores)
		} else {
			log.Printf("[DEBUG] failed to detect change data store %s", datastores)
		}
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

	if esxiSoftwareVersion, ok := s.D.GetOkExists("esxi_software_version"); ok {
		tmp := esxiSoftwareVersion.(string)
		request.EsxiSoftwareVersion = &tmp
	}

	if displayName, ok := s.getOkExistsClusterConfigurationProperty("display_name"); ok {
		tmp := displayName.(string)
		initialClusterConfiguration.DisplayName = &tmp
	}

	esxiHostsCountDeprecated, okDeprecated := s.D.GetOkExists("esxi_hosts_count")
	esxiHostsCount, ok := s.getOkExistsClusterConfigurationProperty("esxi_hosts_count")

	if okDeprecated {
		tmp := esxiHostsCountDeprecated.(int)
		initialClusterConfiguration.EsxiHostsCount = &tmp
	} else if ok {
		tmp := esxiHostsCount.(int)
		initialClusterConfiguration.EsxiHostsCount = &tmp
	}

	if vsphereType, ok := s.getOkExistsClusterConfigurationProperty("vsphere_type"); ok {
		initialClusterConfiguration.VsphereType = oci_ocvp.VsphereTypesEnum(vsphereType.(string))
	} else {
		initialClusterConfiguration.VsphereType = oci_ocvp.VsphereTypesManagement
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	hcxVlanIdDeprecated, okDeprecated := s.D.GetOkExists("hcx_vlan_id")
	hcxVlanId, ok := s.getOkExistsNetworkConfigurationProperty("hcx_vlan_id")

	if okDeprecated {
		tmp := hcxVlanIdDeprecated.(string)
		initialClusterConfiguration.NetworkConfiguration.HcxVlanId = &tmp
	} else if ok {
		tmp := hcxVlanId.(string)
		initialClusterConfiguration.NetworkConfiguration.HcxVlanId = &tmp
	}

	initialHostOcpuCountDeprecated, okDeprecated := s.D.GetOkExists("initial_host_ocpu_count")
	initialHostOcpuCount, ok := s.getOkExistsClusterConfigurationProperty("initial_host_ocpu_count")

	if okDeprecated {
		tmp := float32(initialHostOcpuCountDeprecated.(float64))
		initialClusterConfiguration.InitialHostOcpuCount = &tmp
	} else if ok {
		tmp := float32(initialHostOcpuCount.(float64))
		initialClusterConfiguration.InitialHostOcpuCount = &tmp
	}

	initialHostShapeNameDeprecated, okDeprecated := s.D.GetOkExists("initial_host_shape_name")
	initialHostShapeName, ok := s.getOkExistsClusterConfigurationProperty("initial_host_shape_name")

	if okDeprecated {
		tmp := initialHostShapeNameDeprecated.(string)
		initialClusterConfiguration.InitialHostShapeName = &tmp
	} else if ok {
		tmp := initialHostShapeName.(string)
		initialClusterConfiguration.InitialHostShapeName = &tmp
	}

	initialSkuDeprecated, okDeprecated := s.D.GetOkExists("initial_sku")
	initialSku, ok := s.getOkExistsClusterConfigurationProperty("initial_commitment")

	if okDeprecated {
		tmp := initialSkuDeprecated.(string)
		if enum, ok := oci_ocvp.GetMappingCommitmentEnum(tmp); ok {
			initialClusterConfiguration.InitialCommitment = enum
		} else {
			return errors.New(fmt.Sprintf("cannot map initial commitment %s", tmp))
		}

	} else if ok {
		tmp := initialSku.(string)
		if enum, ok := oci_ocvp.GetMappingCommitmentEnum(tmp); ok {
			initialClusterConfiguration.InitialCommitment = enum
		} else {
			return errors.New(fmt.Sprintf("cannot map initial commitment %s", tmp))
		}
	}

	instanceDisplayNamePrefixDeprecated, okDeprecated := s.D.GetOkExists("instance_display_name_prefix")
	instanceDisplayNamePrefix, ok := s.getOkExistsClusterConfigurationProperty("instance_display_name_prefix")

	if okDeprecated {
		tmp := instanceDisplayNamePrefixDeprecated.(string)
		initialClusterConfiguration.InstanceDisplayNamePrefix = &tmp
	} else if ok {
		tmp := instanceDisplayNamePrefix.(string)
		initialClusterConfiguration.InstanceDisplayNamePrefix = &tmp
	}

	isHcxEnabled, isHcxEnabledExist := s.D.GetOkExists("is_hcx_enabled")
	if isHcxEnabledExist && isHcxEnabled.(bool) {
		request.HcxMode = oci_ocvp.HcxModesAdvanced
	} else {
		request.HcxMode = oci_ocvp.HcxModesDisabled
	}

	if hcxAction, ok := s.D.GetOk("hcx_action"); ok {
		hcxAction = strings.ToUpper(hcxAction.(string))
		if hcxAction == UpgradeHcxAction {
			if isHcxEnabledExist && isHcxEnabled.(bool) {
				request.HcxMode = oci_ocvp.HcxModesEnterprise
			} else {
				return fmt.Errorf("is_hcx_enabled must be set to true when hcx_action is 'UPGRADE'")
			}

		} else {
			return fmt.Errorf("hcx_action '%s' is not supported during SDDC creation. ", hcxAction)
		}
	}

	isShieldedInstanceEnabledDeprecated, okDeprecated := s.D.GetOkExists("is_shielded_instance_enabled")
	isShieldedInstanceEnabled, ok := s.getOkExistsClusterConfigurationProperty("is_shielded_instance_enabled")

	if okDeprecated {
		tmp := isShieldedInstanceEnabledDeprecated.(bool)
		initialClusterConfiguration.IsShieldedInstanceEnabled = &tmp
	} else if ok {
		tmp := isShieldedInstanceEnabled.(bool)
		initialClusterConfiguration.IsShieldedInstanceEnabled = &tmp
	}

	if isSingleHostSddc, ok := s.D.GetOkExists("is_single_host_sddc"); ok {
		tmp := isSingleHostSddc.(bool)
		request.IsSingleHostSddc = &tmp
	}

	nsxEdgeUplink1VlanIdDeprecated, okDeprecated := s.D.GetOkExists("nsx_edge_uplink1vlan_id")
	nsxEdgeUplink1VlanId, ok := s.getOkExistsNetworkConfigurationProperty("nsx_edge_uplink1vlan_id")

	if okDeprecated {
		tmp := nsxEdgeUplink1VlanIdDeprecated.(string)
		initialClusterConfiguration.NetworkConfiguration.NsxEdgeUplink1VlanId = &tmp
	} else if ok {
		tmp := nsxEdgeUplink1VlanId.(string)
		initialClusterConfiguration.NetworkConfiguration.NsxEdgeUplink1VlanId = &tmp
	}

	nsxEdgeUplink2VlanIdDeprecated, okDeprecated := s.D.GetOkExists("nsx_edge_uplink2vlan_id")
	nsxEdgeUplink2VlanId, ok := s.getOkExistsNetworkConfigurationProperty("nsx_edge_uplink2vlan_id")

	if okDeprecated {
		tmp := nsxEdgeUplink2VlanIdDeprecated.(string)
		initialClusterConfiguration.NetworkConfiguration.NsxEdgeUplink2VlanId = &tmp
	} else if ok {
		tmp := nsxEdgeUplink2VlanId.(string)
		initialClusterConfiguration.NetworkConfiguration.NsxEdgeUplink2VlanId = &tmp
	}

	nsxEdgeVTepVlanIdDep, okDep := s.D.GetOkExists("nsx_edge_vtep_vlan_id")
	nsxEdgeVTepVlanId, ok := s.getOkExistsNetworkConfigurationProperty("nsx_edge_vtep_vlan_id")
	if !okDep && !ok {
		return errors.New("one of nsx_edge_vtep_vlan_id or initial_cluster_configurations.networkConfiguration.nsx_edge_vtep_vlan_id must be configured")
	}
	if okDep {
		tmp := nsxEdgeVTepVlanIdDep.(string)
		initialClusterConfiguration.NetworkConfiguration.NsxEdgeVTepVlanId = &tmp
	} else {
		tmp := nsxEdgeVTepVlanId.(string)
		initialClusterConfiguration.NetworkConfiguration.NsxEdgeVTepVlanId = &tmp
	}

	nsxVTepVlanIdDeprecated, okDeprecated := s.D.GetOkExists("nsx_vtep_vlan_id")
	nsxVTepVlanId, ok := s.getOkExistsNetworkConfigurationProperty("nsx_vtep_vlan_id")
	if !okDeprecated && !ok {
		return errors.New("one of nsx_vtep_vlan_id or initial_cluster_configurations.networkConfiguration.nsx_vtep_vlan_id must be configured")
	}
	if okDeprecated {
		tmp := nsxVTepVlanIdDeprecated.(string)
		initialClusterConfiguration.NetworkConfiguration.NsxVTepVlanId = &tmp
	} else {
		tmp := nsxVTepVlanId.(string)
		initialClusterConfiguration.NetworkConfiguration.NsxVTepVlanId = &tmp
	}

	provisioningSubnetIdDeprecated, okDeprecated := s.D.GetOkExists("provisioning_subnet_id")
	provisioningSubnetId, ok := s.getOkExistsNetworkConfigurationProperty("provisioning_subnet_id")

	if !okDeprecated && !ok {
		return errors.New("one of provisioning_subnet_id or initial_cluster_configurations.networkConfiguration.provisioning_subnet_id must be configured")
	}
	if okDeprecated {
		tmp := provisioningSubnetIdDeprecated.(string)
		initialClusterConfiguration.NetworkConfiguration.ProvisioningSubnetId = &tmp
	} else {
		tmp := provisioningSubnetId.(string)
		initialClusterConfiguration.NetworkConfiguration.ProvisioningSubnetId = &tmp
	}

	provisioningVlanIdDeprecated, okDeprecated := s.D.GetOkExists("provisioning_vlan_id")
	provisioningVlanId, ok := s.getOkExistsNetworkConfigurationProperty("provisioning_vlan_id")

	if okDeprecated {
		tmp := provisioningVlanIdDeprecated.(string)
		initialClusterConfiguration.NetworkConfiguration.ProvisioningVlanId = &tmp
	} else if ok {
		tmp := provisioningVlanId.(string)
		initialClusterConfiguration.NetworkConfiguration.ProvisioningVlanId = &tmp
	}

	replicationVlanIdDeprecated, okDeprecated := s.D.GetOkExists("replication_vlan_id")
	replicationVlanId, ok := s.getOkExistsNetworkConfigurationProperty("replication_vlan_id")

	if okDeprecated {
		tmp := replicationVlanIdDeprecated.(string)
		initialClusterConfiguration.NetworkConfiguration.ReplicationVlanId = &tmp
	} else if ok {
		tmp := replicationVlanId.(string)
		initialClusterConfiguration.NetworkConfiguration.ReplicationVlanId = &tmp
	}

	if sshAuthorizedKeys, ok := s.D.GetOkExists("ssh_authorized_keys"); ok {
		tmp := sshAuthorizedKeys.(string)
		request.SshAuthorizedKeys = &tmp
	}

	vmotionVlanIdDeprecated, okDeprecated := s.D.GetOkExists("vmotion_vlan_id")
	vmotionVlanId, ok := s.getOkExistsNetworkConfigurationProperty("vmotion_vlan_id")
	if !okDeprecated && !ok {
		return errors.New("one of vmotion_vlan_id or initial_cluster_configurations.networkConfiguration.vmotion_vlan_id must be configured")
	}
	if okDeprecated {
		tmp := vmotionVlanIdDeprecated.(string)
		initialClusterConfiguration.NetworkConfiguration.VmotionVlanId = &tmp
	} else {
		tmp := vmotionVlanId.(string)
		initialClusterConfiguration.NetworkConfiguration.VmotionVlanId = &tmp
	}

	if vmwareSoftwareVersion, ok := s.D.GetOkExists("vmware_software_version"); ok {
		tmp := vmwareSoftwareVersion.(string)
		request.VmwareSoftwareVersion = &tmp
	}

	vsanVlanIdDeprecated, okDeprecated := s.D.GetOkExists("vsan_vlan_id")
	vsanVlanId, ok := s.getOkExistsNetworkConfigurationProperty("vsan_vlan_id")
	if !okDeprecated && !ok {
		return errors.New("one of vsan_vlan_id or initial_cluster_configurations.networkConfiguration.vsan_vlan_id must be configured")
	}
	if okDeprecated {
		tmp := vsanVlanIdDeprecated.(string)
		initialClusterConfiguration.NetworkConfiguration.VsanVlanId = &tmp
	} else {
		tmp := vsanVlanId.(string)
		initialClusterConfiguration.NetworkConfiguration.VsanVlanId = &tmp
	}

	vsphereVlanIdDeprecated, okDeprecated := s.D.GetOkExists("vsphere_vlan_id")
	vsphereVlanId, ok := s.getOkExistsNetworkConfigurationProperty("vsphere_vlan_id")

	if okDeprecated {
		tmp := vsphereVlanIdDeprecated.(string)
		initialClusterConfiguration.NetworkConfiguration.VsphereVlanId = &tmp
	} else if ok {
		tmp := vsphereVlanId.(string)
		initialClusterConfiguration.NetworkConfiguration.VsphereVlanId = &tmp
	}

	workloadNetworkCidrDeprecated, okDeprecated := s.D.GetOkExists("workload_network_cidr")
	workloadNetworkCidr, ok := s.getOkExistsClusterConfigurationProperty("workload_network_cidr")

	if okDeprecated {
		tmp := workloadNetworkCidrDeprecated.(string)
		initialClusterConfiguration.WorkloadNetworkCidr = &tmp
	} else if ok {
		tmp := workloadNetworkCidr.(string)
		initialClusterConfiguration.WorkloadNetworkCidr = &tmp
	}

	if _, ok := s.D.GetOk("reserving_hcx_on_premise_license_keys"); ok {
		return fmt.Errorf("reserving_hcx_on_premise_license_keys should not be provided during SDDC creation.")
	}

	initialClusterConfigurations := []oci_ocvp.InitialClusterConfiguration{initialClusterConfiguration}
	request.InitialConfiguration = &oci_ocvp.InitialConfiguration{InitialClusterConfigurations: initialClusterConfigurations}
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp")

	response, err := s.Client.CreateSddc(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	workRequestResponse := oci_ocvp.GetWorkRequestResponse{}
	workRequestResponse, err = s.WorkRequestClient.GetWorkRequest(context.Background(),
		oci_ocvp.GetWorkRequestRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp"),
			},
		})
	if err == nil {
		// The work request response contains an array of objects
		for _, res := range workRequestResponse.Resources {
			if res.EntityType != nil && strings.Contains(strings.ToLower(*res.EntityType), "sddc") && res.Identifier != nil {
				s.D.SetId(*res.Identifier)
				break
			}
		}
	}
	creationError := s.getSddcFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp"), oci_ocvp.ActionTypesCreated, s.D.Timeout(schema.TimeoutCreate))

	if creationError != nil {
		return creationError
	} else if hcxAction, ok := s.D.GetOk("hcx_action"); ok {
		s.D.Set("hcx_action", hcxAction)
	}

	if refresh, ok := s.D.GetOk("refresh_hcx_license_status"); ok {
		tmp := s.D.Id()
		return s.refreshHcxLicenseStatus(&tmp, refresh)
	}

	return nil
}

func (s *OcvpSddcResourceCrud) getSddcFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_ocvp.ActionTypesEnum, timeout time.Duration) error {

	// Wait until it finishes
	sddcId, err := sddcWaitForWorkRequest(workId, "sddc",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.WorkRequestClient)

	if err != nil {
		return err
	}
	s.D.SetId(*sddcId)

	return s.Get()
}

func sddcWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "ocvp", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_ocvp.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func sddcWaitForWorkRequest(wId *string, entityType string, action oci_ocvp.ActionTypesEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_ocvp.WorkRequestClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "ocvp")
	retryPolicy.ShouldRetryOperation = sddcWorkRequestShouldRetryFunc(timeout)

	response := oci_ocvp.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_ocvp.OperationStatusInProgress),
			string(oci_ocvp.OperationStatusAccepted),
			string(oci_ocvp.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_ocvp.OperationStatusSucceeded),
			string(oci_ocvp.OperationStatusFailed),
			string(oci_ocvp.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_ocvp.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_ocvp.OperationStatusFailed || response.Status == oci_ocvp.OperationStatusCanceled {
		return nil, getErrorFromOcvpSddcWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromOcvpSddcWorkRequest(client *oci_ocvp.WorkRequestClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_ocvp.ActionTypesEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_ocvp.ListWorkRequestErrorsRequest{
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

func (s *OcvpSddcResourceCrud) Get() error {
	request := oci_ocvp.GetSddcRequest{}

	tmp := s.D.Id()
	request.SddcId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp")

	response, err := s.Client.GetSddc(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Sddc
	return nil
}

func (s *OcvpSddcResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_ocvp.UpdateSddcRequest{}

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

	if esxiSoftwareVersion, ok := s.D.GetOkExists("esxi_software_version"); ok {
		tmp := esxiSoftwareVersion.(string)
		request.EsxiSoftwareVersion = &tmp
	}

	if sshAuthorizedKeys, ok := s.D.GetOkExists("ssh_authorized_keys"); ok {
		tmp := sshAuthorizedKeys.(string)
		request.SshAuthorizedKeys = &tmp
	}

	if vmwareSoftwareVersion, ok := s.D.GetOkExists("vmware_software_version"); ok {
		tmp := vmwareSoftwareVersion.(string)
		request.VmwareSoftwareVersion = &tmp
	}

	sddcId := s.D.Id()
	request.SddcId = &sddcId

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp")

	response, err := s.Client.UpdateSddc(context.Background(), request)
	if err != nil {
		return err
	}
	s.Res = &response.Sddc

	if _, exists := s.D.GetOk("reserving_hcx_on_premise_license_keys"); exists {
		if hcxAction, ok := s.D.GetOk("hcx_action"); !ok || strings.ToUpper(hcxAction.(string)) != DowngradeHcxAction {
			return fmt.Errorf("reserving_hcx_on_premise_license_keys can only be set when hcx_action is DOWNGRADE")
		}
	}

	if action, exists := s.D.GetOk("hcx_action"); exists && strings.ToUpper(action.(string)) == DowngradeHcxAction {
		if _, ok := s.D.GetOk("reserving_hcx_on_premise_license_keys"); !ok {
			return fmt.Errorf("reserving_hcx_on_premise_license_keys must exist when hcx_action is DOWNGRADE")
		} else if s.D.HasChange("reserving_hcx_on_premise_license_keys") && !s.D.HasChange("hcx_action") {
			return fmt.Errorf("reserving_hcx_on_premise_license_keys cannot be changed when hcx_action is already DOWNGRADE")
		}
	}

	var updateHcxError error

	if hcxAction, ok := s.D.GetOk("hcx_action"); ok && s.D.HasChange("hcx_action") {
		action := strings.ToUpper(hcxAction.(string))
		sddcId := s.D.Id()

		if action == UpgradeHcxAction {
			hcxRequest := oci_ocvp.UpgradeHcxRequest{}
			hcxRequest.SddcId = &sddcId
			hcxRes, hcxErr := s.Client.UpgradeHcx(context.Background(), hcxRequest)
			if hcxErr != nil {
				return hcxErr
			}
			workId := hcxRes.OpcWorkRequestId
			updateHcxError = s.getSddcFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp"), oci_ocvp.ActionTypesUpdated, s.D.Timeout(schema.TimeoutUpdate))
			if updateHcxError == nil {
				s.D.Set("hcx_action", hcxAction)
			}
		} else if action == DowngradeHcxAction {
			hcxRequest := oci_ocvp.DowngradeHcxRequest{}
			hcxRequest.SddcId = &sddcId
			if reservingKeys, ok := s.D.GetOk("reserving_hcx_on_premise_license_keys"); ok {
				var keys []string
				for _, key := range reservingKeys.([]interface{}) {
					keys = append(keys, key.(string))
				}
				hcxRequest.ReservingHcxOnPremiseLicenseKeys = keys
			}
			hcxRes, hcxErr := s.Client.DowngradeHcx(context.Background(), hcxRequest)
			if hcxErr != nil {
				return hcxErr
			}
			workId := hcxRes.OpcWorkRequestId
			updateHcxError = s.getSddcFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp"), oci_ocvp.ActionTypesUpdated, s.D.Timeout(schema.TimeoutUpdate))
			if updateHcxError == nil {
				s.D.Set("hcx_action", hcxAction)
			}
		} else if action == CancelDowngradeHcxAction {
			hcxRequest := oci_ocvp.CancelDowngradeHcxRequest{}
			hcxRequest.SddcId = &sddcId
			hcxRes, hcxErr := s.Client.CancelDowngradeHcx(context.Background(), hcxRequest)
			if hcxErr != nil {
				return hcxErr
			}
			workId := hcxRes.OpcWorkRequestId
			updateHcxError = s.getSddcFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp"), oci_ocvp.ActionTypesUpdated, s.D.Timeout(schema.TimeoutUpdate))
			if updateHcxError == nil {
				s.D.Set("hcx_action", hcxAction)
			}
		} else {
			return fmt.Errorf("hcx_action '%s' is not supported. ", hcxAction)
		}
	}

	if updateHcxError != nil {
		return updateHcxError
	}

	if refresh, ok := s.D.GetOk("refresh_hcx_license_status"); ok && s.D.HasChange("refresh_hcx_license_status") {
		return s.refreshHcxLicenseStatus(&sddcId, refresh)
	}

	updateClusterRequest := oci_ocvp.UpdateClusterRequest{}
	networkConfiguration := oci_ocvp.NetworkConfiguration{}
	if nsxEdgeUplink1VlanIdDeprecated, ok := s.D.GetOkExists("nsx_edge_uplink1vlan_id"); ok {
		tmp := nsxEdgeUplink1VlanIdDeprecated.(string)
		networkConfiguration.NsxEdgeUplink1VlanId = &tmp
	}

	if nsxEdgeUplink2VlanIdDeprecated, ok := s.D.GetOkExists("nsx_edge_uplink2vlan_id"); ok {
		tmp := nsxEdgeUplink2VlanIdDeprecated.(string)
		networkConfiguration.NsxEdgeUplink2VlanId = &tmp
	}

	if nsxEdgeVTepVlanIdDep, ok := s.D.GetOkExists("nsx_edge_vtep_vlan_id"); ok {
		tmp := nsxEdgeVTepVlanIdDep.(string)
		networkConfiguration.NsxEdgeVTepVlanId = &tmp
	}

	if nsxVTepVlanIdDeprecated, ok := s.D.GetOkExists("nsx_vtep_vlan_id"); ok {
		tmp := nsxVTepVlanIdDeprecated.(string)
		networkConfiguration.NsxVTepVlanId = &tmp
	}

	if provisioningSubnetIdDeprecated, ok := s.D.GetOkExists("provisioning_subnet_id"); ok {
		tmp := provisioningSubnetIdDeprecated.(string)
		networkConfiguration.ProvisioningSubnetId = &tmp
		// To be consistent with old API behaviour, if deprecated Terraform fields are used, we update the cluster vmware version as well
		if vmwareSoftwareVersion, ok := s.D.GetOkExists("vmware_software_version"); ok {
			tmp := vmwareSoftwareVersion.(string)
			updateClusterRequest.VmwareSoftwareVersion = &tmp
		}
	}

	if provisioningVlanIdDeprecated, ok := s.D.GetOkExists("provisioning_vlan_id"); ok {
		tmp := provisioningVlanIdDeprecated.(string)
		networkConfiguration.ProvisioningVlanId = &tmp
	}

	if replicationVlanIdDeprecated, ok := s.D.GetOkExists("replication_vlan_id"); ok {
		tmp := replicationVlanIdDeprecated.(string)
		networkConfiguration.ReplicationVlanId = &tmp
	}

	if vmotionVlanIdDeprecated, ok := s.D.GetOkExists("vmotion_vlan_id"); ok {
		tmp := vmotionVlanIdDeprecated.(string)
		networkConfiguration.VmotionVlanId = &tmp
	}

	if vsanVlanIdDeprecated, ok := s.D.GetOkExists("vsan_vlan_id"); ok {
		tmp := vsanVlanIdDeprecated.(string)
		networkConfiguration.VsanVlanId = &tmp
	}

	if vSphereVlanIdDeprecated, ok := s.D.GetOkExists("vsphere_vlan_id"); ok {
		tmp := vSphereVlanIdDeprecated.(string)
		networkConfiguration.VsphereVlanId = &tmp
	}

	if hcxVlanIdDeprecated, ok := s.D.GetOkExists("hcx_vlan_id"); ok {
		tmp := hcxVlanIdDeprecated.(string)
		networkConfiguration.HcxVlanId = &tmp
	}

	clusterSummary, err := GetManagementClusterSummary(&sddcId, s.Res.CompartmentId, s.ClusterClient)
	if err != nil {
		return err
	}
	updateClusterRequest.ClusterId = clusterSummary.Id
	updateClusterRequest.NetworkConfiguration = &networkConfiguration
	updateClusterRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp")
	_, err = s.ClusterClient.UpdateCluster(context.Background(), updateClusterRequest)
	if err != nil {
		return err
	}

	getSddcRequest := oci_ocvp.GetSddcRequest{}
	sddcId = s.D.Id()
	getSddcRequest.SddcId = &sddcId
	getSddcRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp")
	getSddcResponse, err := s.Client.GetSddc(context.Background(), getSddcRequest)
	if err != nil {
		return err
	}
	s.Res = &getSddcResponse.Sddc

	return nil
}

func (s *OcvpSddcResourceCrud) refreshHcxLicenseStatus(sddcId *string, refresh interface{}) error {
	hcxRequest := oci_ocvp.RefreshHcxLicenseStatusRequest{}
	hcxRequest.SddcId = sddcId
	hcxRes, hcxErr := s.Client.RefreshHcxLicenseStatus(context.Background(), hcxRequest)
	if hcxErr != nil {
		return hcxErr
	}
	workId := hcxRes.OpcWorkRequestId
	err := s.getSddcFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp"), oci_ocvp.ActionTypesUpdated, s.D.Timeout(schema.TimeoutUpdate))
	if err == nil {
		s.D.Set("refresh_hcx_license_status", refresh)
	}
	return err
}

func (s *OcvpSddcResourceCrud) Delete() error {
	request := oci_ocvp.DeleteSddcRequest{}

	tmp := s.D.Id()
	request.SddcId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp")

	response, err := s.Client.DeleteSddc(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := sddcWaitForWorkRequest(workId, "sddc",
		oci_ocvp.ActionTypesDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.WorkRequestClient)
	return delWorkRequestErr
}

func (s *OcvpSddcResourceCrud) SetData() error {

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if subnetId, ok := s.D.GetOkExists("provisioning_subnet_id"); ok && subnetId != "" {
		log.Printf("[DEBUG] provisioning_subnet_id %s is configured. Using old API fields.", subnetId)
		s.D.Set("initial_configuration", nil)
		actualEsxiHostCount, err := CalculateActualEsxiHostCount(s.Res.Id, s.Res.CompartmentId, s.ClusterClient)
		if err != nil {
			return nil
		}
		// Setting actual count for deprecated
		s.D.Set("actual_esxi_hosts_count", actualEsxiHostCount)
		// We Update value of esxi_hosts_count in state file only if the esxi_hosts_count of the
		// SDDC is modified in the TF config by the user.
		// As there could a scenario where the SDDC esxi_hosts_count on the cloud could be different as esxi host can be attached to the SDDC
		// Then we do not Update the size but instead Update the actual_esxi_hosts_count in the state file.
		_, ok := s.D.GetOk("esxi_hosts_count") // This checks if size is in the state or not. If not and size in response is not nil it could be that user is importing and hence we need to updated the size
		if !ok {
			log.Printf("[DEBUG] esxi_hosts_count does not exists in state, hence assuming user is importing resource")
		}
		if s.D.HasChange("esxi_hosts_count") || !ok {
			oldValue, newValue := s.D.GetChange("esxi_hosts_count")
			log.Printf("[DEBUG] esxi_hosts_count has been updated in config from %v to %v", oldValue, newValue)
			s.D.Set("esxi_hosts_count", actualEsxiHostCount)
		}

		s.D.Set("nsx_overlay_segment_name", "WORKLOAD")
		if len(s.Res.HcxOnPremLicenses) > 0 {
			s.D.Set("hcx_on_prem_key", s.Res.HcxOnPremLicenses[0].ActivationKey)
		}

		if s.Res.HcxMode != oci_ocvp.HcxModesDisabled {
			hcxPassword, err := GetSddcPassword(s.Client, s.D.Id(), oci_ocvp.RetrievePasswordTypeHcx)
			if err != nil {
				return err
			}
			if hcxPassword != nil {
				s.D.Set("hcx_initial_password", *hcxPassword)
			}
		}

		nsxPassword, err := GetSddcPassword(s.Client, s.D.Id(), oci_ocvp.RetrievePasswordTypeNsx)
		if err != nil {
			return err
		}
		if nsxPassword != nil {
			s.D.Set("nsx_manager_initial_password", *nsxPassword)
		}

		vCenterPassword, err := GetSddcPassword(s.Client, s.D.Id(), oci_ocvp.RetrievePasswordTypeVcenter)
		if err != nil {
			return err
		}
		if vCenterPassword != nil {
			s.D.Set("vcenter_initial_password", *vCenterPassword)
		}

		err = s.SetDataClusterValues(s.Res.Id, s.Res.CompartmentId, s.ClusterClient)

		if err != nil {
			return err
		}
	} else {
		if s.Res.InitialConfiguration != nil {
			s.D.Set("initial_configuration", []interface{}{InitialConfigurationToMap(s.Res.InitialConfiguration,
				s.D.GetOk, s.D.HasChange, s.D.GetChange, false)})
		} else {
			s.D.Set("initial_configuration", nil)
		}

		if s.Res.EsxiSoftwareVersion != nil {
			s.D.Set("esxi_software_version", *s.Res.EsxiSoftwareVersion)
		}

		if s.Res.ClustersCount != nil {
			s.D.Set("clusters_count", *s.Res.ClustersCount)
		}
		// set deprecated fields to nil to avoid plan diff when user migrates from deprecated fields to new fields
		s.D.Set("hcx_mode", s.Res.HcxMode)
		s.D.Set("upgrade_licenses", nil)
		s.D.Set("vsphere_upgrade_objects", nil)
		s.D.Set("compute_availability_domain", nil)
		s.D.Set("esxi_hosts_count", nil)
		s.D.Set("actual_esxi_hosts_count", nil)
		s.D.Set("nsx_edge_uplink1vlan_id", nil)
		s.D.Set("nsx_edge_uplink2vlan_id", nil)
		s.D.Set("nsx_edge_vtep_vlan_id", nil)
		s.D.Set("nsx_vtep_vlan_id", nil)
		s.D.Set("provisioning_subnet_id", nil)
		s.D.Set("vmotion_vlan_id", nil)
		s.D.Set("vsan_vlan_id", nil)
		s.D.Set("vsphere_vlan_id", nil)
		s.D.Set("datastores", nil)
		s.D.Set("capacity_reservation_id", nil)
		s.D.Set("hcx_vlan_id", nil)
		s.D.Set("initial_host_ocpu_count", nil)
		s.D.Set("initial_host_shape_name", nil)
		s.D.Set("initial_sku", nil)
		s.D.Set("instance_display_name_prefix", nil)
		s.D.Set("is_shielded_instance_enabled", nil)
		s.D.Set("provisioning_vlan_id", nil)
		s.D.Set("replication_vlan_id", nil)
		s.D.Set("workload_network_cidr", nil)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.HcxFqdn != nil {
		s.D.Set("hcx_fqdn", *s.Res.HcxFqdn)
	}

	hcxOnPremLicenses := []interface{}{}
	for _, item := range s.Res.HcxOnPremLicenses {
		hcxOnPremLicenses = append(hcxOnPremLicenses, HcxLicenseSummaryToMap(item))
	}
	s.D.Set("hcx_on_prem_licenses", hcxOnPremLicenses)

	if s.Res.HcxPrivateIpId != nil {
		s.D.Set("hcx_private_ip_id", *s.Res.HcxPrivateIpId)
	}

	switch s.Res.HcxMode {
	case oci_ocvp.HcxModesDisabled:
		s.D.Set("is_hcx_enabled", false)
		s.D.Set("is_hcx_enterprise_enabled", false)

	case oci_ocvp.HcxModesAdvanced:
		s.D.Set("is_hcx_enabled", true)
		s.D.Set("is_hcx_enterprise_enabled", false)

	case oci_ocvp.HcxModesEnterprise:
		s.D.Set("is_hcx_enabled", true)
		s.D.Set("is_hcx_enterprise_enabled", true)
	}

	if s.Res.IsHcxPendingDowngrade != nil {
		s.D.Set("is_hcx_pending_downgrade", *s.Res.IsHcxPendingDowngrade)
	}

	if s.Res.IsSingleHostSddc != nil {
		s.D.Set("is_single_host_sddc", *s.Res.IsSingleHostSddc)
	}

	if s.Res.NsxEdgeUplinkIpId != nil {
		s.D.Set("nsx_edge_uplink_ip_id", *s.Res.NsxEdgeUplinkIpId)
	}

	if s.Res.NsxManagerFqdn != nil {
		s.D.Set("nsx_manager_fqdn", *s.Res.NsxManagerFqdn)
	}

	if s.Res.NsxManagerPrivateIpId != nil {
		s.D.Set("nsx_manager_private_ip_id", *s.Res.NsxManagerPrivateIpId)
	}

	if s.Res.NsxManagerUsername != nil {
		s.D.Set("nsx_manager_username", *s.Res.NsxManagerUsername)
	}

	if s.Res.SshAuthorizedKeys != nil {
		s.D.Set("ssh_authorized_keys", *s.Res.SshAuthorizedKeys)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeHcxBillingCycleEnd != nil {
		s.D.Set("time_hcx_billing_cycle_end", s.Res.TimeHcxBillingCycleEnd.String())
	}

	if s.Res.TimeHcxLicenseStatusUpdated != nil {
		s.D.Set("time_hcx_license_status_updated", s.Res.TimeHcxLicenseStatusUpdated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.VcenterFqdn != nil {
		s.D.Set("vcenter_fqdn", *s.Res.VcenterFqdn)
	}

	if s.Res.VcenterPrivateIpId != nil {
		s.D.Set("vcenter_private_ip_id", *s.Res.VcenterPrivateIpId)
	}

	if s.Res.VcenterUsername != nil {
		s.D.Set("vcenter_username", *s.Res.VcenterUsername)
	}

	if s.Res.VmwareSoftwareVersion != nil {
		s.D.Set("vmware_software_version", *s.Res.VmwareSoftwareVersion)
	}

	return nil
}

func GetSddcPassword(sddcClient *oci_ocvp.SddcClient, sddcId string, passwordType oci_ocvp.RetrievePasswordTypeEnum) (*string, error) {
	request := oci_ocvp.RetrievePasswordRequest{}
	request.SddcId = &sddcId
	request.Type = passwordType
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "ocvp")
	response, err := sddcClient.RetrievePassword(context.Background(), request)
	return response.SddcPassword.Value, err
}

func GetManagementClusterSummary(sddcId *string, compartmentId *string, clusterClient *oci_ocvp.ClusterClient) (*oci_ocvp.ClusterSummary, error) {
	request := oci_ocvp.ListClustersRequest{}
	request.SddcId = sddcId
	request.CompartmentId = compartmentId
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "ocvp")
	response, err := clusterClient.ListClusters(context.Background(), request)

	if err != nil {
		return nil, fmt.Errorf("failed to list clusters sddcId : '%s'", *sddcId)
	}
	for _, clusterSummary := range response.Items {
		// only update management cluster
		if clusterSummary.VsphereType == oci_ocvp.VsphereTypesManagement {
			return &clusterSummary, nil
		}
	}
	return nil, fmt.Errorf("cannot find management Cluster for SDDC %s", *sddcId)
}

func CalculateActualEsxiHostCount(sddcId *string, compartmentId *string, clusterClient *oci_ocvp.ClusterClient) (int, error) {
	log.Printf("[DEBUG] getting esxi host count from cluster")
	request := oci_ocvp.ListClustersRequest{}
	request.SddcId = sddcId
	request.CompartmentId = compartmentId
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "ocvp")
	response, err := clusterClient.ListClusters(context.Background(), request)

	if err != nil {
		return 0, fmt.Errorf("failed to list clusters sddcId : '%s'", *sddcId)
	}
	count := 0
	for _, item := range response.Items {
		count += *item.EsxiHostsCount
	}
	return count, nil

}

func (s *OcvpSddcResourceCrud) SetDataClusterValues(sddcId *string, compartmentId *string, clusterClient *oci_ocvp.ClusterClient) error {
	clusterSummary, err := GetManagementClusterSummary(sddcId, compartmentId, clusterClient)
	if err != nil {
		return err
	}
	clusterId := clusterSummary.Id
	log.Printf("[DEBUG] setting values from cluster %s", *clusterId)

	req := oci_ocvp.GetClusterRequest{}
	req.ClusterId = clusterId
	clusterResponse, err := clusterClient.GetCluster(context.Background(), req)

	if err != nil {
		log.Printf("[ERROR] failed to get cluster id : '%s'", *clusterId)
		return err
	}
	log.Printf("[DEBUG] setting vshere upgrade objects")
	vsphereUpgradeObjects := []interface{}{}
	for _, item := range clusterResponse.VsphereUpgradeObjects {
		vsphereUpgradeObjects = append(vsphereUpgradeObjects, VsphereUpgradeObjectToMap(item))
	}
	err = s.D.Set("vsphere_upgrade_objects", vsphereUpgradeObjects)
	if err != nil {
		return err
	}

	if len(vsphereUpgradeObjects) > 0 {
		s.D.Set("vsphere_upgrade_guide", "vsphereUpgradeGuide_place_holder")
	}

	log.Printf("[DEBUG] setting upgrade licenses")
	upgradeLicenses := []interface{}{}
	for _, item := range clusterResponse.UpgradeLicenses {
		upgradeLicenses = append(upgradeLicenses, VsphereLicenseToMap(item))
	}
	err = s.D.Set("upgrade_licenses", upgradeLicenses)
	if err != nil {
		return err
	}

	if clusterResponse.CapacityReservationId != nil {
		s.D.Set("capacity_reservation_id", clusterResponse.CapacityReservationId)
	}

	if clusterResponse.ComputeAvailabilityDomain != nil {
		s.D.Set("compute_availability_domain", clusterResponse.ComputeAvailabilityDomain)
	}

	datastores := []interface{}{}
	for _, item := range clusterResponse.Datastores {
		datastores = append(datastores, DatastoreDetailsToMap(item))
	}
	s.D.Set("datastores", datastores)

	if clusterResponse.InitialHostOcpuCount != nil {
		s.D.Set("initial_host_ocpu_count", clusterResponse.InitialHostOcpuCount)
	}

	if clusterResponse.InitialHostShapeName != nil {
		s.D.Set("initial_host_shape_name", clusterResponse.InitialHostShapeName)
	}

	s.D.Set("initial_sku", clusterResponse.InitialCommitment)

	if clusterResponse.InstanceDisplayNamePrefix != nil {
		s.D.Set("instance_display_name_prefix", clusterResponse.InstanceDisplayNamePrefix)
	}

	if clusterResponse.IsShieldedInstanceEnabled != nil {
		s.D.Set("is_shielded_instance_enabled", clusterResponse.IsShieldedInstanceEnabled)
	}

	if clusterResponse.WorkloadNetworkCidr != nil {
		s.D.Set("workload_network_cidr", clusterResponse.WorkloadNetworkCidr)
	}

	networkConfiguration := clusterResponse.NetworkConfiguration
	if networkConfiguration.NsxEdgeUplink1VlanId != nil {
		s.D.Set("nsx_edge_uplink1vlan_id", networkConfiguration.NsxEdgeUplink1VlanId)
	}
	if networkConfiguration.NsxEdgeUplink2VlanId != nil {
		s.D.Set("nsx_edge_uplink2vlan_id", networkConfiguration.NsxEdgeUplink2VlanId)
	}
	if networkConfiguration.NsxEdgeVTepVlanId != nil {
		s.D.Set("nsx_edge_vtep_vlan_id", networkConfiguration.NsxEdgeVTepVlanId)
	}
	if networkConfiguration.NsxVTepVlanId != nil {
		s.D.Set("nsx_vtep_vlan_id", networkConfiguration.NsxVTepVlanId)
	}
	if networkConfiguration.ProvisioningSubnetId != nil {
		s.D.Set("provisioning_subnet_id", networkConfiguration.ProvisioningSubnetId)
	}
	if networkConfiguration.ProvisioningVlanId != nil {
		s.D.Set("provisioning_vlan_id", networkConfiguration.ProvisioningVlanId)
	}
	if networkConfiguration.ReplicationVlanId != nil {
		s.D.Set("replication_vlan_id", networkConfiguration.ReplicationVlanId)
	}
	if networkConfiguration.VmotionVlanId != nil {
		s.D.Set("vmotion_vlan_id", networkConfiguration.VmotionVlanId)
	}
	if networkConfiguration.VsanVlanId != nil {
		s.D.Set("vsan_vlan_id", networkConfiguration.VsanVlanId)
	}
	if networkConfiguration.VsphereVlanId != nil {
		s.D.Set("vsphere_vlan_id", networkConfiguration.VsphereVlanId)
	}

	if networkConfiguration.HcxVlanId != nil {
		s.D.Set("hcx_vlan_id", networkConfiguration.HcxVlanId)
	}
	return nil
}

func InitialConfigurationToMap(obj *oci_ocvp.InitialConfiguration, getOk func(val string) (interface{}, bool), hasChange func(val string) bool,
	getChange func(key string) (interface{}, interface{}), isDataSource bool) map[string]interface{} {
	result := map[string]interface{}{}

	initialClusterConfigurations := []interface{}{}
	for _, item := range obj.InitialClusterConfigurations {
		initialClusterConfigurations = append(initialClusterConfigurations, InitialClusterConfigurationToMap(item, getOk, hasChange, getChange, isDataSource))
	}
	result["initial_cluster_configurations"] = initialClusterConfigurations

	return result
}

func InitialClusterConfigurationToMap(obj oci_ocvp.InitialClusterConfiguration,
	getOk func(val string) (interface{}, bool), hasChange func(val string) bool,
	getChange func(key string) (interface{}, interface{}), isDataSource bool) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CapacityReservationId != nil {
		result["capacity_reservation_id"] = string(*obj.CapacityReservationId)
	}

	if obj.ComputeAvailabilityDomain != nil {
		result["compute_availability_domain"] = string(*obj.ComputeAvailabilityDomain)
	}

	datastores := []interface{}{}
	for _, item := range obj.Datastores {
		datastores = append(datastores, DatastoreSummaryToMap(item))
	}
	result["datastores"] = datastores

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	// We Update value of esxi_hosts_count in state file only if the esxi_hosts_count of the
	// SDDC is modified in the TF config by the user.
	// As there could a scenario where the SDDC esxi_hosts_count on the cloud could be different as esxi host can be attached to the SDDC
	// Then we do not Update the size but instead Update the actual_esxi_hosts_count in the state file.
	fieldKey := fmt.Sprintf("%s.%d.%s.%d.%s", "initial_configuration", 0, "initial_cluster_configurations", 0, "esxi_hosts_count")
	if obj.EsxiHostsCount != nil {
		log.Printf("[DEBUG] esxi_hosts_count is not nill %s", obj)
		configuredHostsCount, ok := getOk(fieldKey) // This checks if size is in the state or not. If not and size in response is not nil it could be that user is importing and hence we need to updated the size
		if !ok {
			log.Printf("[DEBUG] esxi_hosts_count does not exists in state, hence assuming user is importing resource")
		}
		actualHostsCount := *obj.EsxiHostsCount
		if hasChange(fieldKey) || !ok || isDataSource {
			oldValue, newValue := getChange(fieldKey)
			log.Printf("[DEBUG] esxi_hosts_count has been updated in config from %v to %v", oldValue, newValue)
			log.Printf("[DEBUG] setting esxi_hosts_count value to %d", actualHostsCount)
			result["esxi_hosts_count"] = actualHostsCount
		} else {
			result["esxi_hosts_count"] = configuredHostsCount
		}
		result["actual_esxi_hosts_count"] = actualHostsCount
	}

	result["initial_commitment"] = string(obj.InitialCommitment)

	if obj.InstanceDisplayNamePrefix != nil {
		result["instance_display_name_prefix"] = string(*obj.InstanceDisplayNamePrefix)
	}

	if obj.IsShieldedInstanceEnabled != nil {
		result["is_shielded_instance_enabled"] = bool(*obj.IsShieldedInstanceEnabled)
	}

	if obj.InitialHostOcpuCount != nil {
		result["initial_host_ocpu_count"] = float32(*obj.InitialHostOcpuCount)
	}

	if obj.InitialHostShapeName != nil {
		result["initial_host_shape_name"] = string(*obj.InitialHostShapeName)
	}

	if obj.NetworkConfiguration != nil {
		result["network_configuration"] = []interface{}{NetworkConfigurationToMap(obj.NetworkConfiguration)}
	}

	result["vsphere_type"] = string(obj.VsphereType)

	if obj.WorkloadNetworkCidr != nil {
		result["workload_network_cidr"] = string(*obj.WorkloadNetworkCidr)
	}

	return result
}
func (s *OcvpSddcResourceCrud) mapToDatastoreInfo(fieldKeyFormat string) (oci_ocvp.DatastoreInfo, error) {
	result := oci_ocvp.DatastoreInfo{}

	if blockVolumeIds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "block_volume_ids")); ok {
		interfaces := blockVolumeIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "block_volume_ids")) {
			result.BlockVolumeIds = tmp
		}
	}

	if datastoreType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "datastore_type")); ok {
		result.DatastoreType = oci_ocvp.DatastoreTypesEnum(datastoreType.(string))
	}

	return result, nil
}

func DatastoreSummaryToMap(obj oci_ocvp.DatastoreInfo) map[string]interface{} {
	result := map[string]interface{}{}

	result["block_volume_ids"] = obj.BlockVolumeIds

	result["datastore_type"] = string(obj.DatastoreType)

	return result
}

func HcxLicenseSummaryToMap(obj oci_ocvp.HcxLicenseSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ActivationKey != nil {
		result["activation_key"] = string(*obj.ActivationKey)
	}

	result["status"] = string(obj.Status)

	if obj.SystemName != nil {
		result["system_name"] = string(*obj.SystemName)
	}

	return result
}

func SddcSummaryToMap(obj oci_ocvp.SddcSummary, sddcClient *oci_ocvp.SddcClient, clusterClient *oci_ocvp.ClusterClient) (map[string]interface{}, error) {
	result := map[string]interface{}{}

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

	if obj.HcxFqdn != nil {
		result["hcx_fqdn"] = string(*obj.HcxFqdn)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	result["hcx_mode"] = obj.HcxMode

	switch obj.HcxMode {
	case oci_ocvp.HcxModesDisabled:
		result["is_hcx_enabled"] = false
		result["is_hcx_enterprise_enabled"] = false

	case oci_ocvp.HcxModesAdvanced:
		result["is_hcx_enabled"] = true
		result["is_hcx_enterprise_enabled"] = false

	case oci_ocvp.HcxModesEnterprise:
		result["is_hcx_enabled"] = true
		result["is_hcx_enterprise_enabled"] = true
	}

	if obj.IsSingleHostSddc != nil {
		result["is_single_host_sddc"] = bool(*obj.IsSingleHostSddc)
	}

	if obj.ClustersCount != nil {
		result["clusters_count"] = int(*obj.ClustersCount)
	}

	if obj.NsxManagerFqdn != nil {
		result["nsx_manager_fqdn"] = string(*obj.NsxManagerFqdn)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	if obj.VcenterFqdn != nil {
		result["vcenter_fqdn"] = string(*obj.VcenterFqdn)
	}

	if obj.VmwareSoftwareVersion != nil {
		result["vmware_software_version"] = string(*obj.VmwareSoftwareVersion)
	}

	count, err := CalculateActualEsxiHostCount(obj.Id, obj.CompartmentId, clusterClient)
	if err != nil {
		return nil, err
	}
	result["esxi_hosts_count"] = &count
	result["actual_esxi_hosts_count"] = &count

	clusterSummary, err := GetManagementClusterSummary(obj.Id, obj.CompartmentId, clusterClient)
	if err != nil {
		return nil, err
	}
	if clusterSummary.ComputeAvailabilityDomain != nil {
		result["compute_availability_domain"] = clusterSummary.ComputeAvailabilityDomain
	}

	if clusterSummary.InitialHostOcpuCount != nil {
		result["initial_host_ocpu_count"] = clusterSummary.InitialHostOcpuCount
	}

	if clusterSummary.InitialHostShapeName != nil {
		result["initial_host_shape_name"] = clusterSummary.InitialHostShapeName
	}

	if clusterSummary.IsShieldedInstanceEnabled != nil {
		result["is_shielded_instance_enabled"] = clusterSummary.IsShieldedInstanceEnabled
	}

	return result, nil
}

func (s *OcvpSddcResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_ocvp.ChangeSddcCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.SddcId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp")

	_, err := s.Client.ChangeSddcCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
