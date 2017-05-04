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

	"github.com/oracle/terraform-provider-baremetal/client"
	"github.com/oracle/terraform-provider-baremetal/client/mocks"
)

type ResourceLoadBalancerBackendTestSuite struct {
	suite.Suite
	Client      client.BareMetalClient
	Providers   map[string]terraform.ResourceProvider
	TimeCreated baremetal.Time
}

func (s *ResourceLoadBalancerBackendTestSuite) SetupTest() {
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

func (s *ResourceLoadBalancerBackendTestSuite) TestCreateResourceLoadBalancerBackendMaximal() {
	resourceName := "baremetal_load_balancer_backend.t"
	config := `
resource "baremetal_load_balancer_backend" "t" {
  load_balancer_id = "ocid1.loadbalancer.stub_id"
  backendset_name  = "stub_backendset_name"
  name             = "stub_backend_name"
  ip_address       = "1.2.3.4"
  port             = 1234
  backup           = true
  drain            = true
  offline          = true
  weight           = 1
}
`
	config += testProviderConfig

	loadBalancerID := "ocid1.loadbalancer.stub_id"
	backendsetName := "stub_backendset_name"
	res := &baremetal.Backend{
		Name:      "stub_backend_name",
		IPAddress: "1.2.3.4",
		Port:      1234,
		Backup:    true,
		Drain:     true,
		Offline:   true,
		Weight:    1,
	}
	res.RequestID = "stub_opc_request_id"
	opts := &baremetal.CreateLoadBalancerBackendOptions{
		Backup:  res.Backup,
		Drain:   res.Drain,
		Offline: res.Offline,
		Weight:  res.Weight,
	}
	// opts := (*baremetal.CreateLoadBalancerBackendOptions)(nil)

	deletedRes := &baremetal.Backend{}
	*deletedRes = *res

	workReqID := "stub_work_req_id"
	s.Client.On(
		"CreateBackend",
		loadBalancerID,
		backendsetName,
		res.IPAddress,
		res.Port,
		opts,
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
		"GetBackend",
		loadBalancerID,
		backendsetName,
		res.Name,
		(*baremetal.ClientRequestOptions)(nil),
	).Return(res, nil)

	s.Client.On(
		"DeleteBackend",
		loadBalancerID,
		backendsetName,
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
					resource.TestCheckResourceAttr(resourceName, "backendset_name", backendsetName),
					resource.TestCheckResourceAttr(resourceName, "name", res.Name),

					resource.TestCheckResourceAttr(resourceName, "ip_address", res.IPAddress),

					resource.TestCheckResourceAttr(resourceName, "backup", strconv.FormatBool(opts.Backup)),
					resource.TestCheckResourceAttr(resourceName, "drain", strconv.FormatBool(opts.Drain)),
					resource.TestCheckResourceAttr(resourceName, "offline", strconv.FormatBool(opts.Offline)),
					resource.TestCheckResourceAttr(resourceName, "weight", strconv.Itoa(opts.Weight)),
				),
			},
		},
	})
}

func TestResourceLoadBalancerBackendTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceLoadBalancerBackendTestSuite))
}
