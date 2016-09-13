package core

import (
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/crud"
	"github.com/hashicorp/terraform/helper/schema"
)

// // See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/DHCPDNSOption/
// type DHCPDNSOption struct {
// 	Type             string   `json:"type"`
// 	CustomDNSServers []string `json:"customDnsServers"`
// 	ServerType       string   `json:"serverType"`
// }

// // DHCPOptions contains a set of dhcp options
// //
// // See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/DHCPOptions/
// type DHCPOptions struct {
// 	RequestableResource
// 	ETaggedResource
// 	CompartmentID string          `json:"compartmentId"`
// 	DisplayName   string          `json:"displayName"`
// 	ID            string          `json:"id"`
// 	Options       []DHCPDNSOption `json:"options"`
// 	State         string          `json:"lifecycleState"`
// 	TimeCreated   Time            `json:"timeCreated"`
// }

func DHCPOptionsResource() *schema.Resource {
	return &schema.Resource{
		Create: createDHCPOptions,
		Read:   readDHCPOptions,
		Update: updateDHCPOptions,
		Delete: deleteDHCPOptions,
		Schema: map[string]*schema.Schema{
			"compartment_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"display_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"options": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"custom_dns_servers": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"server_type": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vcn_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createDHCPOptions(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	crd := &DHCPOptionsResourceCrud{D: d, Client: client}
	return crud.CreateResource(d, crd)
}

func readDHCPOptions(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	crd := &DHCPOptionsResourceCrud{D: d, Client: client}
	return crud.ReadResource(crd)
}

func updateDHCPOptions(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	crd := &DHCPOptionsResourceCrud{D: d, Client: client}
	return crud.UpdateResource(d, crd)
}

func deleteDHCPOptions(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	crd := &DHCPOptionsResourceCrud{D: d, Client: client}
	return crud.DeleteResource(crd)
}
