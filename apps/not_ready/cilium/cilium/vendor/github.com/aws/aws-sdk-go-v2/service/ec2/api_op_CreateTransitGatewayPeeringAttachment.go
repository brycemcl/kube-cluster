// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateTransitGatewayPeeringAttachmentInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The AWS account ID of the owner of the peer transit gateway.
	//
	// PeerAccountId is a required field
	PeerAccountId *string `type:"string" required:"true"`

	// The Region where the peer transit gateway is located.
	//
	// PeerRegion is a required field
	PeerRegion *string `type:"string" required:"true"`

	// The ID of the peer transit gateway with which to create the peering attachment.
	//
	// PeerTransitGatewayId is a required field
	PeerTransitGatewayId *string `type:"string" required:"true"`

	// The tags to apply to the transit gateway peering attachment.
	TagSpecifications []TagSpecification `locationName:"TagSpecification" locationNameList:"item" type:"list"`

	// The ID of the transit gateway.
	//
	// TransitGatewayId is a required field
	TransitGatewayId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateTransitGatewayPeeringAttachmentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateTransitGatewayPeeringAttachmentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateTransitGatewayPeeringAttachmentInput"}

	if s.PeerAccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PeerAccountId"))
	}

	if s.PeerRegion == nil {
		invalidParams.Add(aws.NewErrParamRequired("PeerRegion"))
	}

	if s.PeerTransitGatewayId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PeerTransitGatewayId"))
	}

	if s.TransitGatewayId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TransitGatewayId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateTransitGatewayPeeringAttachmentOutput struct {
	_ struct{} `type:"structure"`

	// The transit gateway peering attachment.
	TransitGatewayPeeringAttachment *TransitGatewayPeeringAttachment `locationName:"transitGatewayPeeringAttachment" type:"structure"`
}

// String returns the string representation
func (s CreateTransitGatewayPeeringAttachmentOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateTransitGatewayPeeringAttachment = "CreateTransitGatewayPeeringAttachment"

// CreateTransitGatewayPeeringAttachmentRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Requests a transit gateway peering attachment between the specified transit
// gateway (requester) and a peer transit gateway (accepter). The transit gateways
// must be in different Regions. The peer transit gateway can be in your account
// or a different AWS account.
//
// After you create the peering attachment, the owner of the accepter transit
// gateway must accept the attachment request.
//
//    // Example sending a request using CreateTransitGatewayPeeringAttachmentRequest.
//    req := client.CreateTransitGatewayPeeringAttachmentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateTransitGatewayPeeringAttachment
func (c *Client) CreateTransitGatewayPeeringAttachmentRequest(input *CreateTransitGatewayPeeringAttachmentInput) CreateTransitGatewayPeeringAttachmentRequest {
	op := &aws.Operation{
		Name:       opCreateTransitGatewayPeeringAttachment,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateTransitGatewayPeeringAttachmentInput{}
	}

	req := c.newRequest(op, input, &CreateTransitGatewayPeeringAttachmentOutput{})
	return CreateTransitGatewayPeeringAttachmentRequest{Request: req, Input: input, Copy: c.CreateTransitGatewayPeeringAttachmentRequest}
}

// CreateTransitGatewayPeeringAttachmentRequest is the request type for the
// CreateTransitGatewayPeeringAttachment API operation.
type CreateTransitGatewayPeeringAttachmentRequest struct {
	*aws.Request
	Input *CreateTransitGatewayPeeringAttachmentInput
	Copy  func(*CreateTransitGatewayPeeringAttachmentInput) CreateTransitGatewayPeeringAttachmentRequest
}

// Send marshals and sends the CreateTransitGatewayPeeringAttachment API request.
func (r CreateTransitGatewayPeeringAttachmentRequest) Send(ctx context.Context) (*CreateTransitGatewayPeeringAttachmentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateTransitGatewayPeeringAttachmentResponse{
		CreateTransitGatewayPeeringAttachmentOutput: r.Request.Data.(*CreateTransitGatewayPeeringAttachmentOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateTransitGatewayPeeringAttachmentResponse is the response type for the
// CreateTransitGatewayPeeringAttachment API operation.
type CreateTransitGatewayPeeringAttachmentResponse struct {
	*CreateTransitGatewayPeeringAttachmentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateTransitGatewayPeeringAttachment request.
func (r *CreateTransitGatewayPeeringAttachmentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
