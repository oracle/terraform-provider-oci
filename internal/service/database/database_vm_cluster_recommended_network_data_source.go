// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"fmt"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v58/database"
)

func DatabaseVmClusterRecommendedNetworkDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDatabaseVmClusterRecommendedNetwork,
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
			"exadata_infrastructure_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"networks": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"cidr": {
							Type:     schema.TypeString,
							Required: true,
						},
						"domain": {
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
						"prefix": {
							Type:     schema.TypeString,
							Required: true,
						},
						"vlan_id": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},

			// Optional
			"defined_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"dns": {
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
			"ntp": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"scan_listener_port_tcp": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"scan_listener_port_tcp_ssl": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			// Computed
			"scans": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"hostname": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ips": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"port": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"scan_listener_port_tcp": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"scan_listener_port_tcp_ssl": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"vm_networks": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"domain_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"gateway": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"netmask": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"network_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"nodes": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"hostname": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"ip": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"vip": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"vip_hostname": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"vlan_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readSingularDatabaseVmClusterRecommendedNetwork(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseVmClusterRecommendedNetworkDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseVmClusterRecommendedNetworkDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.GenerateRecommendedVmClusterNetworkResponse
}

func (s *DatabaseVmClusterRecommendedNetworkDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseVmClusterRecommendedNetworkDataSourceCrud) Get() error {
	request := oci_database.GenerateRecommendedVmClusterNetworkRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
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

	if dns, ok := s.D.GetOkExists("dns"); ok {
		interfaces := dns.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("dns") {
			request.Dns = tmp
		}
	}

	if exadataInfrastructureId, ok := s.D.GetOkExists("exadata_infrastructure_id"); ok {
		tmp := exadataInfrastructureId.(string)
		request.ExadataInfrastructureId = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if networks, ok := s.D.GetOkExists("networks"); ok {
		interfaces := networks.([]interface{})
		tmp := make([]oci_database.InfoForNetworkGenDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "networks", stateDataIndex)
			converted, err := s.mapToInfoForNetworkGenDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("networks") {
			request.Networks = tmp
		}
	}

	if ntp, ok := s.D.GetOkExists("ntp"); ok {
		interfaces := ntp.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("ntp") {
			request.Ntp = tmp
		}
	}

	if scanListenerPortTcp, ok := s.D.GetOkExists("scan_listener_port_tcp"); ok {
		tmp := scanListenerPortTcp.(int)
		request.ScanListenerPortTcp = &tmp
	}

	if scanListenerPortTcpSsl, ok := s.D.GetOkExists("scan_listener_port_tcp_ssl"); ok {
		tmp := scanListenerPortTcpSsl.(int)
		request.ScanListenerPortTcpSsl = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.GenerateRecommendedVmClusterNetwork(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseVmClusterRecommendedNetworkDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseVmClusterRecommendedNetworkDataSource-", DatabaseVmClusterRecommendedNetworkDataSource(), s.D))

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("dns", s.Res.Dns)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("ntp", s.Res.Ntp)

	scans := []interface{}{}
	for _, item := range s.Res.Scans {
		scans = append(scans, ScanDetailsToMap(item))
	}
	s.D.Set("scans", scans)

	vmNetworks := []interface{}{}
	for _, item := range s.Res.VmNetworks {
		vmNetworks = append(vmNetworks, VmNetworkDetailsToMap(item, true))
	}
	s.D.Set("vm_networks", vmNetworks)

	return nil
}

func (s *DatabaseVmClusterRecommendedNetworkDataSourceCrud) mapToInfoForNetworkGenDetails(fieldKeyFormat string) (oci_database.InfoForNetworkGenDetails, error) {
	result := oci_database.InfoForNetworkGenDetails{}

	if cidr, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cidr")); ok {
		tmp := cidr.(string)
		result.Cidr = &tmp
	}

	if domain, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "domain")); ok {
		tmp := domain.(string)
		result.Domain = &tmp
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
		result.NetworkType = oci_database.InfoForNetworkGenDetailsNetworkTypeEnum(networkType.(string))
	}

	if prefix, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "prefix")); ok {
		tmp := prefix.(string)
		result.Prefix = &tmp
	}

	if vlanId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vlan_id")); ok {
		tmp := vlanId.(string)
		result.VlanId = &tmp
	}

	return result, nil
}
