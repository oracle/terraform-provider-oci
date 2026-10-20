// Copyright (c) 2017, 2026, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package generative_ai

import "testing"

type hostedDeploymentServiceError struct {
	statusCode int
	code       string
	message    string
}

func (e hostedDeploymentServiceError) Error() string          { return e.message }
func (e hostedDeploymentServiceError) GetHTTPStatusCode() int { return e.statusCode }
func (e hostedDeploymentServiceError) GetCode() string        { return e.code }
func (e hostedDeploymentServiceError) GetMessage() string     { return e.message }
func (e hostedDeploymentServiceError) GetOpcRequestID() string {
	return "test-request-id"
}

func TestIsActiveDeploymentDeleteError(t *testing.T) {
	crud := &GenerativeAiHostedDeploymentResourceCrud{}

	tests := []struct {
		name string
		err  error
		want bool
	}{
		{
			name: "active deployment",
			err: hostedDeploymentServiceError{
				statusCode: 403,
				code:       "NotAllowed",
				message:    "Deployment deletion NOT allowed: deployment is activeDeployment",
			},
			want: true,
		},
		{
			name: "different service error",
			err: hostedDeploymentServiceError{
				statusCode: 403,
				code:       "NotAllowed",
				message:    "deployment cannot be deleted",
			},
			want: false,
		},
		{
			name: "different status",
			err: hostedDeploymentServiceError{
				statusCode: 409,
				code:       "Conflict",
				message:    "deployment is activeDeployment",
			},
			want: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := crud.isActiveDeploymentDeleteError(test.err); got != test.want {
				t.Fatalf("isActiveDeploymentDeleteError() = %t, want %t", got, test.want)
			}
		})
	}
}
