// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
)

type DeleteNetworkAclEntryInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// Indicates whether the rule is an egress rule.
	//
	// Egress is a required field
	Egress *bool `locationName:"egress" type:"boolean" required:"true"`

	// The ID of the network ACL.
	//
	// NetworkAclId is a required field
	NetworkAclId *string `locationName:"networkAclId" type:"string" required:"true"`

	// The rule number of the entry to delete.
	//
	// RuleNumber is a required field
	RuleNumber *int64 `locationName:"ruleNumber" type:"integer" required:"true"`
}

// String returns the string representation
func (s DeleteNetworkAclEntryInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteNetworkAclEntryInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteNetworkAclEntryInput"}

	if s.Egress == nil {
		invalidParams.Add(aws.NewErrParamRequired("Egress"))
	}

	if s.NetworkAclId == nil {
		invalidParams.Add(aws.NewErrParamRequired("NetworkAclId"))
	}

	if s.RuleNumber == nil {
		invalidParams.Add(aws.NewErrParamRequired("RuleNumber"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteNetworkAclEntryOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteNetworkAclEntryOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteNetworkAclEntry = "DeleteNetworkAclEntry"

// DeleteNetworkAclEntryRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Deletes the specified ingress or egress entry (rule) from the specified network
// ACL.
//
//    // Example sending a request using DeleteNetworkAclEntryRequest.
//    req := client.DeleteNetworkAclEntryRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DeleteNetworkAclEntry
func (c *Client) DeleteNetworkAclEntryRequest(input *DeleteNetworkAclEntryInput) DeleteNetworkAclEntryRequest {
	op := &aws.Operation{
		Name:       opDeleteNetworkAclEntry,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteNetworkAclEntryInput{}
	}

	req := c.newRequest(op, input, &DeleteNetworkAclEntryOutput{})
	req.Handlers.Unmarshal.Remove(ec2query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteNetworkAclEntryRequest{Request: req, Input: input, Copy: c.DeleteNetworkAclEntryRequest}
}

// DeleteNetworkAclEntryRequest is the request type for the
// DeleteNetworkAclEntry API operation.
type DeleteNetworkAclEntryRequest struct {
	*aws.Request
	Input *DeleteNetworkAclEntryInput
	Copy  func(*DeleteNetworkAclEntryInput) DeleteNetworkAclEntryRequest
}

// Send marshals and sends the DeleteNetworkAclEntry API request.
func (r DeleteNetworkAclEntryRequest) Send(ctx context.Context) (*DeleteNetworkAclEntryResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteNetworkAclEntryResponse{
		DeleteNetworkAclEntryOutput: r.Request.Data.(*DeleteNetworkAclEntryOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteNetworkAclEntryResponse is the response type for the
// DeleteNetworkAclEntry API operation.
type DeleteNetworkAclEntryResponse struct {
	*DeleteNetworkAclEntryOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteNetworkAclEntry request.
func (r *DeleteNetworkAclEntryResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
