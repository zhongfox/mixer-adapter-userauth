// nolint:lll
// Generates the auth adapter's resource yaml. It contains the adapter's configuration, name, supported template
// names (metric in this case), and whether it is session or no-session based.
//go:generate $GOPATH/src/istio.io/istio/bin/mixer_codegen.sh -a mixer/adapter/userauth/config/config.proto -x "-s=false -n userauth -t authorization"

package userauth

import (
	"context"
	"fmt"
	"net"
	"time"

	"google.golang.org/grpc"

	"istio.io/api/mixer/adapter/model/v1beta1"
	policy "istio.io/api/policy/v1beta1"
	"istio.io/istio/mixer/adapter/userauth/config"
	"istio.io/istio/mixer/pkg/status"
	"istio.io/istio/mixer/template/authorization"
)

type (
	// Server is basic server interface
	Server interface {
		Addr() string
		Close() error
		Run(shutdown chan error)
	}

	// AuthAdapter supports authorization template.
	AuthAdapter struct {
		listener net.Listener
		server   *grpc.Server
	}
)

var _ authorization.HandleAuthorizationServiceServer = &AuthAdapter{}

// HandleAuthorization token validate
func (s *AuthAdapter) HandleAuthorization(ctx context.Context, r *authorization.HandleAuthorizationRequest) (*v1beta1.CheckResult, error) {

	fmt.Printf("received request %v\n", *r)

	cfg := &config.Params{}

	if r.AdapterConfig != nil {
		if err := cfg.Unmarshal(r.AdapterConfig.Value); err != nil {
			fmt.Printf("error unmarshalling adapter config: %v", err)
			return nil, err
		}
	}

	decodeValue := func(in interface{}) interface{} {
		switch t := in.(type) {
		case *policy.Value_StringValue:
			return t.StringValue
		case *policy.Value_Int64Value:
			return t.Int64Value
		case *policy.Value_DoubleValue:
			return t.DoubleValue
		default:
			return fmt.Sprintf("%v", in)
		}
	}

	decodeValueMap := func(in map[string]*policy.Value) map[string]interface{} {
		out := make(map[string]interface{}, len(in))
		for k, v := range in {
			out[k] = decodeValue(v.GetValue())
		}
		return out
	}

	fmt.Println(cfg.Token)

	props := decodeValueMap(r.Instance.Subject.Properties)
	fmt.Printf("%v\n", props)

	for k, v := range props {
		fmt.Println("k:", k, "v:", v)
		if (k == "token") && v == cfg.Token {
			fmt.Println("success!!")
			return &v1beta1.CheckResult{
				Status:        status.OK,
				ValidDuration: time.Second * 3,
				ValidUseCount: 3,
			}, nil
		}
	}

	fmt.Println("failure; header not provided")
	return &v1beta1.CheckResult{
		Status: status.WithPermissionDenied("Unauthorized..."),
	}, nil
}

// Addr returns the listening address of the server
func (s *AuthAdapter) Addr() string {
	return s.listener.Addr().String()
}

// Run starts the server run
func (s *AuthAdapter) Run(shutdown chan error) {
	shutdown <- s.server.Serve(s.listener)
}

// Close gracefully shuts down the server; used for testing
func (s *AuthAdapter) Close() error {
	if s.server != nil {
		s.server.GracefulStop()
	}

	if s.listener != nil {
		_ = s.listener.Close()
	}

	return nil
}

// NewAuthAdapter creates a new IBP adapter that listens at provided port.
func NewAuthAdapter(addr string) (Server, error) {
	if addr == "" {
		addr = "0"
	}
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", addr))
	if err != nil {
		return nil, fmt.Errorf("unable to listen on socket: %v", err)
	}
	s := &AuthAdapter{
		listener: listener,
	}
	fmt.Printf("listening on \"%v\"\n", s.Addr())
	s.server = grpc.NewServer()
	authorization.RegisterHandleAuthorizationServiceServer(s.server, s)
	return s, nil
}
