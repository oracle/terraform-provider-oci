// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"strconv"
	"testing"
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/stretchr/testify/suite"

	"github.com/oracle/terraform-provider-baremetal/client/mocks"
)

type ResourceLoadBalancerBackendsetTestSuite struct {
	suite.Suite
	Client      *mocks.BareMetalClient
	Providers   map[string]terraform.ResourceProvider
	TimeCreated baremetal.Time
}

func (s *ResourceLoadBalancerBackendsetTestSuite) SetupTest() {
	s.Client = &mocks.BareMetalClient{}

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": Provider(
			func(d *schema.ResourceData) (interface{}, error) {
				return s.Client, nil
			},
		),
	}

	s.TimeCreated = baremetal.Time{Time: time.Now()}
}

func (s *ResourceLoadBalancerBackendsetTestSuite) TestCreateResourceLoadBalancerBackendsetrMaximal() {
	resourceName := "baremetal_load_balancer_backendset.t"
	config := `
resource "baremetal_load_balancer_backendset" "t" {
  load_balancer_id = "ocid1.loadbalancer.stub_id"
  name             = "stub_backendset_name"
  policy           = "stub_policy"

  health_checker {
    interval_ms         = 30001
    port                = 1234
    protocol            = "stub_protocol"
    response_body_regex = "stub_regex"
  }

  ssl_configuration {
    certificate_name        = "stub_certificate_name"
    verify_depth            = 6
    verify_peer_certificate = false
  }
}
`
	config += testProviderConfig

	loadBalancerID := "ocid1.loadbalancer.stub_id"
	res := &baremetal.BackendSet{
		Name:   "stub_backendset_name",
		Policy: "stub_policy",
		HealthChecker: &baremetal.HealthChecker{
			IntervalInMS:      30001,
			Port:              1234,
			Protocol:          "stub_protocol",
			ResponseBodyRegex: "stub_regex",
		},
		SSLConfig: &baremetal.SSLConfiguration{
			CertificateName:       "stub_certificate_name",
			VerifyDepth:           6,
			VerifyPeerCertificate: false,
		},
		// empty?
		Backends: []baremetal.Backend{},
	}
	res.RequestID = "stub_opc_request_id"
	// opts := baremetal.LoadBalancerOptions{}

	deletedRes := &baremetal.BackendSet{}
	*deletedRes = *res

	workReqID := "stub_work_req_id"
	s.Client.On(
		"CreateBackendSet",
		loadBalancerID,
		res.Name,
		res.Policy,
		res.Backends,
		res.HealthChecker,
		res.SSLConfig,
		// &opts,
		(*baremetal.LoadBalancerOptions)(nil),
	).Return(workReqID, nil)

	workReqCreated := &baremetal.WorkRequest{
		ID:             workReqID,
		LoadBalancerID: loadBalancerID,
		State:          baremetal.WorkRequestSucceeded,
	}

	s.Client.On(
		"GetWorkRequest",
		workReqID,
		(*baremetal.ClientRequestOptions)(nil),
	).Return(workReqCreated, nil)

	s.Client.On(
		"GetBackendSet",
		loadBalancerID,
		res.Name,
		(*baremetal.ClientRequestOptions)(nil),
	).Return(res, nil)

	s.Client.On(
		"DeleteBackendSet",
		loadBalancerID,
		res.Name,
		(*baremetal.ClientRequestOptions)(nil),
	).Return(workReqID, nil)

	s.Client.On(
		"GetWorkRequest",
		workReqID,
		(*baremetal.ClientRequestOptions)(nil),
	).Return(workReqCreated, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "load_balancer_id", loadBalancerID),
					resource.TestCheckResourceAttr(resourceName, "name", res.Name),

					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.certificate_name", res.SSLConfig.CertificateName),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.verify_depth", strconv.Itoa(res.SSLConfig.VerifyDepth)),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.verify_peer_certificate", strconv.FormatBool(res.SSLConfig.VerifyPeerCertificate)),
				),
			},
		},
	})
}

func TestResourceLoadBalancerBackendsetTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceLoadBalancerBackendsetTestSuite))
}
