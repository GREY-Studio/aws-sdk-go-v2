// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
)

// WaitUntilResourceRecordSetsChanged uses the Route 53 API operation
// GetChange to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Route53) WaitUntilResourceRecordSetsChanged(ctx context.Context, input *GetChangeInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilResourceRecordSetsChanged",
		MaxAttempts: 60,
		Delay:       aws.ConstantWaiterDelay(30 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "ChangeInfo.Status",
				Expected: "INSYNC",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *GetChangeInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.GetChangeRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}
