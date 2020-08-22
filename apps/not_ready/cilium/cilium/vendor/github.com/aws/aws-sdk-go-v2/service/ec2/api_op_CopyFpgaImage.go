// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CopyFpgaImageInput struct {
	_ struct{} `type:"structure"`

	// Unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request. For more information, see Ensuring Idempotency (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Run_Instance_Idempotency.html).
	ClientToken *string `type:"string"`

	// The description for the new AFI.
	Description *string `type:"string"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The name for the new AFI. The default is the name of the source AFI.
	Name *string `type:"string"`

	// The ID of the source AFI.
	//
	// SourceFpgaImageId is a required field
	SourceFpgaImageId *string `type:"string" required:"true"`

	// The Region that contains the source AFI.
	//
	// SourceRegion is a required field
	SourceRegion *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CopyFpgaImageInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CopyFpgaImageInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CopyFpgaImageInput"}

	if s.SourceFpgaImageId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SourceFpgaImageId"))
	}

	if s.SourceRegion == nil {
		invalidParams.Add(aws.NewErrParamRequired("SourceRegion"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CopyFpgaImageOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the new AFI.
	FpgaImageId *string `locationName:"fpgaImageId" type:"string"`
}

// String returns the string representation
func (s CopyFpgaImageOutput) String() string {
	return awsutil.Prettify(s)
}

const opCopyFpgaImage = "CopyFpgaImage"

// CopyFpgaImageRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Copies the specified Amazon FPGA Image (AFI) to the current Region.
//
//    // Example sending a request using CopyFpgaImageRequest.
//    req := client.CopyFpgaImageRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CopyFpgaImage
func (c *Client) CopyFpgaImageRequest(input *CopyFpgaImageInput) CopyFpgaImageRequest {
	op := &aws.Operation{
		Name:       opCopyFpgaImage,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CopyFpgaImageInput{}
	}

	req := c.newRequest(op, input, &CopyFpgaImageOutput{})
	return CopyFpgaImageRequest{Request: req, Input: input, Copy: c.CopyFpgaImageRequest}
}

// CopyFpgaImageRequest is the request type for the
// CopyFpgaImage API operation.
type CopyFpgaImageRequest struct {
	*aws.Request
	Input *CopyFpgaImageInput
	Copy  func(*CopyFpgaImageInput) CopyFpgaImageRequest
}

// Send marshals and sends the CopyFpgaImage API request.
func (r CopyFpgaImageRequest) Send(ctx context.Context) (*CopyFpgaImageResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CopyFpgaImageResponse{
		CopyFpgaImageOutput: r.Request.Data.(*CopyFpgaImageOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CopyFpgaImageResponse is the response type for the
// CopyFpgaImage API operation.
type CopyFpgaImageResponse struct {
	*CopyFpgaImageOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CopyFpgaImage request.
func (r *CopyFpgaImageResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
