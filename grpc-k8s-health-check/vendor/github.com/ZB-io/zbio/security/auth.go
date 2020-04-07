package security

import (
	"context"
	"github.com/ZB-io/zbio/log"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func BuildAuthFunction(expectedScheme string, expectedToken string) func(ctx context.Context) (context.Context, error) {
	return func(ctx context.Context) (context.Context, error) {
		token, err := grpc_auth.AuthFromMD(ctx, expectedScheme)
		if err != nil {
			return nil, err
		}
		if token != expectedToken {
			return nil, status.Errorf(codes.PermissionDenied, "BuildAuthFunction bad token")
		} else {
			log.Debugf("Successful authentication: %v", expectedToken)
		}
		return context.WithValue(ctx, "some_context_marker", "marker_exists"), nil
	}
}
