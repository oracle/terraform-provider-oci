// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Load Balancing Service API
//
// API for the Load Balancing Service
//

package loadbalancer

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ConnectionConfiguration Configuration details for the connection between the client and backend servers.
type ConnectionConfiguration struct {

	// The maximum idle time, in seconds, allowed between two successive receive or two successive send operations
	// between the client and backend servers. A send operation does not reset the timer for receive operations. A
	// receive operation does not reset the timer for send operations.
	// The default values are:
	// *  300 seconds for TCP
	// *  60 seconds for HTTP and WebSocket protocols.
	// Note: The protocol is set at the listener.
	// Modify this parameter if the client or backend server stops transmitting data for more than the default time.
	// Some examples include:
	// *  The client sends a database query to the backend server and the database takes over 300 seconds to execute.
	//    Therefore, the backend server does not transmit any data within 300 seconds.
	// *  The client uploads data using the HTTP protocol. During the upload, the backend does not transmit any data
	//    to the client for more than 60 seconds.
	// *  The client downloads data using the HTTP protocol.  After the initial request, it stops transmitting data to
	//    the backend server for more than 60 seconds.
	// *  The client starts transmitting data after establishing a WebSocket connection, but the backend server does
	//    not transmit data for more than 60 seconds.
	// *  The backend server starts transmitting data after establishing a WebSocket connection, but the client does
	//    not transmit data for more than 60 seconds.
	// The maximum value is 7200 seconds. Contact My Oracle Support to file a service request if you want to increase
	// this limit for your tenancy. For more information, see Service Limits (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/servicelimits.htm).
	// Example: `1200`
	IdleTimeout *int `mandatory:"true" json:"idleTimeout"`
}

func (m ConnectionConfiguration) String() string {
	return common.PointerString(m)
}
