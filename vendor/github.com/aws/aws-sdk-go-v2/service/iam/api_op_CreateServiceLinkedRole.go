// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/CreateServiceLinkedRoleRequest
type CreateServiceLinkedRoleInput struct {
	_ struct{} `type:"structure"`

	// The service principal for the AWS service to which this role is attached.
	// You use a string similar to a URL but without the http:// in front. For example:
	// elasticbeanstalk.amazonaws.com.
	//
	// Service principals are unique and case-sensitive. To find the exact service
	// principal for your service-linked role, see AWS Services That Work with IAM
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-services-that-work-with-iam.html)
	// in the IAM User Guide. Look for the services that have Yes in the Service-Linked
	// Role column. Choose the Yes link to view the service-linked role documentation
	// for that service.
	//
	// AWSServiceName is a required field
	AWSServiceName *string `min:"1" type:"string" required:"true"`

	// A string that you provide, which is combined with the service-provided prefix
	// to form the complete role name. If you make multiple requests for the same
	// service, then you must supply a different CustomSuffix for each request.
	// Otherwise the request fails with a duplicate role name error. For example,
	// you could add -1 or -debug to the suffix.
	//
	// Some services do not support the CustomSuffix parameter. If you provide an
	// optional suffix and the operation fails, try the operation again without
	// the suffix.
	CustomSuffix *string `min:"1" type:"string"`

	// The description of the role.
	Description *string `type:"string"`
}

// String returns the string representation
func (s CreateServiceLinkedRoleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateServiceLinkedRoleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateServiceLinkedRoleInput"}

	if s.AWSServiceName == nil {
		invalidParams.Add(aws.NewErrParamRequired("AWSServiceName"))
	}
	if s.AWSServiceName != nil && len(*s.AWSServiceName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AWSServiceName", 1))
	}
	if s.CustomSuffix != nil && len(*s.CustomSuffix) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CustomSuffix", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/CreateServiceLinkedRoleResponse
type CreateServiceLinkedRoleOutput struct {
	_ struct{} `type:"structure"`

	// A Role object that contains details about the newly created role.
	Role *Role `type:"structure"`
}

// String returns the string representation
func (s CreateServiceLinkedRoleOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateServiceLinkedRole = "CreateServiceLinkedRole"

// CreateServiceLinkedRoleRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Creates an IAM role that is linked to a specific AWS service. The service
// controls the attached policies and when the role can be deleted. This helps
// ensure that the service is not broken by an unexpectedly changed or deleted
// role, which could put your AWS resources into an unknown state. Allowing
// the service to control the role helps improve service stability and proper
// cleanup when a service and its role are no longer needed. For more information,
// see Using Service-Linked Roles (https://docs.aws.amazon.com/IAM/latest/UserGuide/using-service-linked-roles.html)
// in the IAM User Guide.
//
// To attach a policy to this service-linked role, you must make the request
// using the AWS service that depends on this role.
//
//    // Example sending a request using CreateServiceLinkedRoleRequest.
//    req := client.CreateServiceLinkedRoleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/CreateServiceLinkedRole
func (c *Client) CreateServiceLinkedRoleRequest(input *CreateServiceLinkedRoleInput) CreateServiceLinkedRoleRequest {
	op := &aws.Operation{
		Name:       opCreateServiceLinkedRole,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateServiceLinkedRoleInput{}
	}

	req := c.newRequest(op, input, &CreateServiceLinkedRoleOutput{})
	return CreateServiceLinkedRoleRequest{Request: req, Input: input, Copy: c.CreateServiceLinkedRoleRequest}
}

// CreateServiceLinkedRoleRequest is the request type for the
// CreateServiceLinkedRole API operation.
type CreateServiceLinkedRoleRequest struct {
	*aws.Request
	Input *CreateServiceLinkedRoleInput
	Copy  func(*CreateServiceLinkedRoleInput) CreateServiceLinkedRoleRequest
}

// Send marshals and sends the CreateServiceLinkedRole API request.
func (r CreateServiceLinkedRoleRequest) Send(ctx context.Context) (*CreateServiceLinkedRoleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateServiceLinkedRoleResponse{
		CreateServiceLinkedRoleOutput: r.Request.Data.(*CreateServiceLinkedRoleOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateServiceLinkedRoleResponse is the response type for the
// CreateServiceLinkedRole API operation.
type CreateServiceLinkedRoleResponse struct {
	*CreateServiceLinkedRoleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateServiceLinkedRole request.
func (r *CreateServiceLinkedRoleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
