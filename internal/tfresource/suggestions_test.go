package tfresource

import "testing"

func TestUnit_getSuggestionFor400(t *testing.T) {
	type args struct {
		tfError customError
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test returned value is as expected for InvalidParameter",
			args: args{
				tfError: customError{
					ErrorCodeName: "InvalidParameter",
					Service:       "core",
					Message:       "testMessage",
				},
			},
			want: "Please Update the parameter(s) in the Terraform config as per error message testMessage",
		},
		{
			name: "Test returned value is as expected for LimitExceeded",
			args: args{
				tfError: customError{
					ErrorCodeName: "LimitExceeded",
					Service:       "core",
					Message:       "testMessage",
				},
			},
			want: "Request a service limit increase for this resource core",
		},
		{
			name: "Test returned value is as expected for QuotaExceeded",
			args: args{
				tfError: customError{
					ErrorCodeName: "QuotaExceeded",
					Service:       "core",
					Message:       "testMessage",
				},
			},
			want: "Contact your administrator to increase limit for your account or compartment for this service: core",
		},
		{
			name: "Test returned value is as expected for default message",
			args: args{
				tfError: customError{
					ErrorCodeName: "Default",
					Service:       "core",
					Message:       "testMessage",
				},
			},
			want: "Please retry or contact support for help with service: core",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSuggestionFor400(tt.args.tfError); got != tt.want {
				t.Errorf("getSuggestionFor400() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnit_getSuggestionForServiceError(t *testing.T) {
	type args struct {
		tfError customError
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test returned value is as expected for 400 code",
			args: args{
				tfError: customError{
					TypeOfError:   ServiceError,
					ErrorCode:     400,
					ErrorCodeName: "InvalidParameter",
					Service:       "core",
					Message:       "testMessage",
				},
			},
			want: "Please Update the parameter(s) in the Terraform config as per error message testMessage",
		},
		{
			name: "Test returned value is as expected for 404 code",
			args: args{
				tfError: customError{
					ErrorCode: 404,
					Service:   "core",
				},
			},
			want: "Either the resource has been deleted or service core need policy to access this resource. Policy reference: https://docs.oracle.com/en-us/iaas/Content/Identity/Reference/policyreference.htm",
		},
		{
			name: "Test returned value is as expected for 409 code",
			args: args{
				tfError: customError{
					ErrorCode: 409,
					Service:   "core",
				},
			},
			want: "The resource is in a conflicted state. Please retry again or contact support for help with service: core",
		},
		{
			name: "Test returned value is as expected for 429 code",
			args: args{
				tfError: customError{
					ErrorCode: 429,
				},
			},
			want: "Please re-apply your Terraform config and/or increase the retry timeout using this document: https://docs.oracle.com/en-us/iaas/Content/API/SDKDocs/terraformtroubleshooting.htm#common_issues__automaticretries",
		},
		{
			name: "Test returned value is as expected for 500 code",
			args: args{
				tfError: customError{
					ErrorCode: 500,
					Service:   "core",
				},
			},
			want: "The service for this resource encountered an error. Please contact support for help with service: core",
		},

		{
			name: "Test returned value is as expected for default message",
			args: args{
				tfError: customError{
					ErrorCode:     501,
					ErrorCodeName: "Default",
					Service:       "core",
				},
			},
			want: "Please retry or contact support for help with service: core",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSuggestionForServiceError(tt.args.tfError); got != tt.want {
				t.Errorf("getSuggestionForServiceError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnit_getSuggestionFromError(t *testing.T) {
	type args struct {
		tfError customError
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test returned value is as expected for ServiceError",
			args: args{
				tfError: customError{
					TypeOfError:   ServiceError,
					ErrorCode:     400,
					ErrorCodeName: "InvalidParameter",
					Service:       "core",
					Message:       "testMessage",
				},
			},
			want: "Please Update the parameter(s) in the Terraform config as per error message testMessage",
		},
		{
			name: "Test returned value is as expected for TimeoutError",
			args: args{
				tfError: customError{
					TypeOfError: TimeoutError,
				},
			},
			want: "Try increasing the timeout by referring to this document: https://docs.oracle.com/en-us/iaas/Content/API/SDKDocs/terraformtroubleshooting.htm#common_issues__timeoutwhilewaiting",
		},
		{
			name: "Test returned value is as expected for UnexpectedStateError",
			args: args{
				tfError: customError{
					TypeOfError: UnexpectedStateError,
					Service:     "core",
				},
			},
			want: "Please retry or contact support for help with service: core",
		},
		{
			name: "Test returned value is as expected for Default",
			args: args{
				tfError: customError{
					TypeOfError: WorkRequestError,
					Service:     "core",
				},
			},
			want: "Please retry or contact support for help with service: core",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSuggestionFromError(tt.args.tfError); got != tt.want {
				t.Errorf("getSuggestionFromError() = %v, want %v", got, tt.want)
			}
		})
	}
}
