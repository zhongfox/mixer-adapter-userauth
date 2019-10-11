// nolint:lll
// Generates the auth adapter's resource yaml. It contains the adapter's configuration, name, supported template
// names (metric in this case), and whether it is session or no-session based.
//go:generate $GOPATH/src/istio.io/istio/bin/mixer_codegen.sh -a mixer/adapter/userauth/config/config.proto -x "-s=false -n userauth -t authorization"

package userauth

import (
	"context"
	"fmt"
	"google.golang.org/grpc/metadata"
	"net"
	"userauth/ftl_mixadp"

	"google.golang.org/grpc"
	"errors"

	"istio.io/api/mixer/adapter/model/v1beta1"
	policy "istio.io/api/policy/v1beta1"
	"istio.io/istio/mixer/pkg/status"
	"istio.io/istio/mixer/template/authorization"
	"userauth/config"
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

 func decodeValue (in interface{}) interface{} {
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

func decodeValueMap (in map[string]*policy.Value) map[string]interface{} {
	out := make(map[string]interface{}, len(in))
	for k, v := range in {
		out[k] = decodeValue(v.GetValue())
	}
	return out
}

// HandleAuthorization token validate
func (s *AuthAdapter) HandleAuthorization(ctx context.Context, r *authorization.HandleAuthorizationRequest) (*v1beta1.CheckResult, error) {

	fmt.Printf("received request %v\n", *r)

	cfg := &config.Params{}

	if r.AdapterConfig != nil {
		if err := cfg.Unmarshal(r.AdapterConfig.Value); err != nil {
			fmt.Printf("error unmarshalling adapter config: %v", err)
			return nil, err // TODO test
		}
	}

	addr := cfg.Token
	if addr == "" {
		err := errors.New("no auth address provided")
		fmt.Println(err)
		return nil, err
	}

	props := decodeValueMap(r.Instance.Subject.Properties)
	fmt.Printf("checking with attrs: %v\n", props)

	if t, ok := props["token"]; ok {
		cookie := t.(string)
		md := metadata.Pairs("cookie", cookie)
		ctx := metadata.NewOutgoingContext(context.Background(), md)

		conn, err := grpc.Dial(addr, grpc.WithInsecure())
		defer conn.Close()
		if err != nil {
			fmt.Printf("can not connect: %v\n", err)

			return &v1beta1.CheckResult{ Status: status.WithUnavailable("connect error") }, nil
		}

		client := ftl_mixadp.NewSigClient(conn)
		// var header, trailer metadata.MD
		response, err := client.Verify(ctx, &ftl_mixadp.VerifyReq{})
		if err != nil {
			return &v1beta1.CheckResult{ Status: status.WithUnavailable("verify error") }, nil
		}

		if response.ErrCode == "ok" {
			// return &v1beta1.CheckResult{ Status: status.OK, ValidDuration: time.Second * 3, ValidUseCount: 3}, nil
			return &v1beta1.CheckResult{ Status: status.OK}, nil
		} else if response.ErrCode =="sig-err" {
			message := fmt.Sprintf("Unauthorized: %s", response.ErrMsg)
			return &v1beta1.CheckResult{ Status: status.WithPermissionDenied(message)}, nil
		} else {
			return &v1beta1.CheckResult{ Status: status.WithUnavailable(response.ErrMsg) }, nil
		}
	}

	fmt.Println("failure; header not provided")
	return &v1beta1.CheckResult{ Status: status.WithPermissionDenied("Unauthorized...") }, nil
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
