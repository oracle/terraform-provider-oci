// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/oracle/oci-go-sdk/core"
	"github.com/stretchr/testify/suite"
)

type DatasourceCoreVolumeTestSuite struct {
	suite.Suite
	Config       string
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *DatasourceCoreVolumeTestSuite) SetupTest() {
	s.Providers = testAccProviders
	s.Config = legacyTestProviderConfig() + `
	data "oci_identity_availability_domains" "ADs" {
		compartment_id = "${var.compartment_id}"
	}
	resource "oci_core_volume" "t" {
		availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
		compartment_id = "${var.compartment_id}"
		display_name = "-tf-volume"
		size_in_gbs = 50
	}
	resource "oci_core_volume" "u" {
		availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
		compartment_id = "${var.compartment_id}"
		display_name = "-tf-volume-clone"
		size_in_gbs = 50
		source_details {
			type = "volume"
			id = "${oci_core_volume.t.id}"
		}
	}`
	s.ResourceName = "data.oci_core_volumes.t"
}

func (s *DatasourceCoreVolumeTestSuite) TestAccDatasourceCoreVolume_basic() {
	compartmentID := getCompartmentIDForLegacyTests()
	resource.Test(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + `
				data "oci_core_volumes" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${oci_core_volume.t.compartment_id}"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					TestCheckResourceAttributesEqual(s.ResourceName, "availability_domain", "data.oci_identity_availability_domains.ADs", "availability_domains.0.name"),
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", compartmentID),
					resource.TestCheckResourceAttrSet(s.ResourceName, "volumes.#"),
				),
			},
			// Server-side filtering tests.
			{
				// This test exercises filtering by state. Adding client-side filtering by display_name
				// to restrict the results to the volumes created in this test.
				Config: s.Config + `
				data "oci_core_volumes" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${oci_core_volume.t.compartment_id}"
					state = "` + string(core.VolumeLifecycleStateAvailable) + `"
					filter {
						name = "display_name"
						values = ["-tf.*"]
						regex = true
					}
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					TestCheckResourceAttributesEqual(s.ResourceName, "availability_domain", "data.oci_identity_availability_domains.ADs", "availability_domains.0.name"),
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", compartmentID),
					resource.TestCheckResourceAttr(s.ResourceName, "volumes.#", "2"),
					// TODO: Add id checks once order-by is supported.
				),
			},
			{
				// This test exercises filtering by display_name. Adding filtering by state filter as
				// well to limit the scope to available volumes because the service can return terminated
				// if they have the same display name.
				Config: s.Config + `
				data "oci_core_volumes" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${oci_core_volume.t.compartment_id}"
					state = "` + string(core.VolumeLifecycleStateAvailable) + `"
					display_name = "${oci_core_volume.t.display_name}"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					TestCheckResourceAttributesEqual(s.ResourceName, "availability_domain", "data.oci_identity_availability_domains.ADs", "availability_domains.0.name"),
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", compartmentID),
					resource.TestCheckResourceAttr(s.ResourceName, "volumes.#", "1"),
					TestCheckResourceAttributesEqual(s.ResourceName, "volumes.0.id", "oci_core_volume.t", "id"),
				),
			},
			// Client-side filtering tests.
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + `
				data "oci_core_volumes" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${oci_core_volume.t.compartment_id}"
					filter {
						name = "id"
						values = ["${oci_core_volume.t.id}"]
					}
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					TestCheckResourceAttributesEqual(s.ResourceName, "availability_domain", "data.oci_identity_availability_domains.ADs", "availability_domains.0.name"),
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", compartmentID),
					resource.TestCheckResourceAttr(s.ResourceName, "volumes.#", "1"),
					TestCheckResourceAttributesEqual(s.ResourceName, "volumes.0.availability_domain", "oci_core_volume.t", "availability_domain"),
					TestCheckResourceAttributesEqual(s.ResourceName, "volumes.0.compartment_id", "oci_core_volume.t", "compartment_id"),
					TestCheckResourceAttributesEqual(s.ResourceName, "volumes.0.id", "oci_core_volume.t", "id"),
					TestCheckResourceAttributesEqual(s.ResourceName, "volumes.0.time_created", "oci_core_volume.t", "time_created"),
					resource.TestCheckResourceAttr(s.ResourceName, "volumes.0.state", string(core.VolumeLifecycleStateAvailable)),
					resource.TestCheckResourceAttr(s.ResourceName, "volumes.0.display_name", "-tf-volume"),
					resource.TestCheckResourceAttr(s.ResourceName, "volumes.0.size_in_gbs", "50"),
					resource.TestCheckResourceAttr(s.ResourceName, "volumes.0.size_in_mbs", "51200"),
					resource.TestCheckResourceAttr(s.ResourceName, "volumes.0.source_details.#", "0"),
					resource.TestCheckResourceAttr(s.ResourceName, "volumes.0.is_hydrated", "true"),
				),
			},
			{
				Config: s.Config + `
				data "oci_core_volumes" "u" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${oci_core_volume.u.compartment_id}"
					filter {
						name = "id"
						values = ["${oci_core_volume.u.id}"]
					}
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					TestCheckResourceAttributesEqual("data.oci_core_volumes.u", "availability_domain", "data.oci_identity_availability_domains.ADs", "availability_domains.0.name"),
					resource.TestCheckResourceAttr("data.oci_core_volumes.u", "compartment_id", compartmentID),
					resource.TestCheckResourceAttr("data.oci_core_volumes.u", "volumes.#", "1"),
					resource.TestCheckResourceAttr("data.oci_core_volumes.u", "volumes.0.source_details.#", "1"),
					TestCheckResourceAttributesEqual("data.oci_core_volumes.u", "volumes.0.availability_domain", "oci_core_volume.u", "availability_domain"),
					TestCheckResourceAttributesEqual("data.oci_core_volumes.u", "volumes.0.compartment_id", "oci_core_volume.u", "compartment_id"),
					TestCheckResourceAttributesEqual("data.oci_core_volumes.u", "volumes.0.id", "oci_core_volume.u", "id"),
					resource.TestCheckResourceAttr("data.oci_core_volumes.u", "volumes.0.source_details.0.type", "volume"),
					TestCheckResourceAttributesEqual("data.oci_core_volumes.u", "volumes.0.source_details.0.id", "oci_core_volume.u", "source_details.0.id"),
					resource.TestCheckResourceAttr("data.oci_core_volumes.u", "volumes.0.state", string(core.VolumeLifecycleStateAvailable)),
					resource.TestCheckResourceAttr("data.oci_core_volumes.u", "volumes.0.display_name", "-tf-volume-clone"),
					resource.TestCheckResourceAttr("data.oci_core_volumes.u", "volumes.0.size_in_gbs", "50"),
					resource.TestCheckResourceAttr("data.oci_core_volumes.u", "volumes.0.size_in_mbs", "51200"),
					TestCheckResourceAttributesEqual("data.oci_core_volumes.u", "volumes.0.is_hydrated", "oci_core_volume.u", "is_hydrated"),
					TestCheckResourceAttributesEqual("data.oci_core_volumes.u", "volumes.0.time_created", "oci_core_volume.u", "time_created"),
				),
			},
		},
	},
	)
}

func TestDatasourceCoreVolumeTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceCoreVolumeTestSuite))
}
