package backend

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"strings"

	"github.com/fullstorydev/grpcurl"
	"github.com/golang/protobuf/jsonpb"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/grpcreflect"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
)

type GrpcClient struct {
	descSource grpcurl.DescriptorSource
	conn       *grpc.ClientConn
	ctx        context.Context
}

func NewGrpcClient(ctx context.Context, endpoint string) (*GrpcClient, error) {
	network := "tcp"
	var opts []grpc.DialOption
	var creds credentials.TransportCredentials
	conn, err := grpcurl.BlockingDial(ctx, network, endpoint, creds, opts...)
	if err != nil {
		return nil, err
	}
	reflectClient := grpcreflect.NewClientAuto(ctx, conn)
	return &GrpcClient{
		descSource: grpcurl.DescriptorSourceFromServer(ctx, reflectClient),
		conn:       conn,
		ctx:        ctx,
	}, nil
}

func (client *GrpcClient) Close() {
	err := client.conn.Close()
	if err != nil {
		slog.Error(err.Error())
	}
}

// ListGrpcServices list grpc services
func (client *GrpcClient) ListGrpcServices() ([]string, error) {
	svcs, err := grpcurl.ListServices(client.descSource)
	if err != nil {
		return nil, err
	}
	return svcs, nil
}

// ListGrpcMethods list grpc methods
func (client *GrpcClient) ListGrpcMethods(service string) ([]string, error) {
	methods, err := grpcurl.ListMethods(client.descSource, service)
	if err != nil {
		return nil, err
	}
	return methods, nil
}

// DescribeGrpcSymbol grpc symbol can be service, method, message...
func (client *GrpcClient) DescribeGrpcSymbol(symbol string) (string, error) {
	if symbol[0] == '.' {
		symbol = symbol[1:]
	}
	descriptor, err := client.descSource.FindSymbol(symbol)
	if err != nil {
		return "", err
	}
	// fullyQualifiedName := descriptor.GetFullyQualifiedName()
	text, err := grpcurl.GetDescriptorText(descriptor, client.descSource)
	if err != nil {
		errMsg := fmt.Sprintf("for symbol %s, fail to get descriptor text", symbol)
		return "", errors.New(errMsg)
	}
	return text, nil
}

func (client *GrpcClient) GetGrpcMessageTemplate(symbol string) (string, error) {
	if symbol[0] == '.' {
		symbol = symbol[1:]
	}
	descriptor, err := client.descSource.FindSymbol(symbol)
	if err != nil {
		return "", err
	}

	if messageDescriptor, ok := descriptor.(*desc.MessageDescriptor); ok {
		// for messages, also show a template in JSON
		message := grpcurl.MakeTemplate(messageDescriptor)
		jsm := jsonpb.Marshaler{
			EnumsAsInts:  false,
			EmitDefaults: true,
		}
		template, err := jsm.MarshalToString(message)
		if err != nil {
			return "", err
		}
		return template, nil
	}

	errMsg := fmt.Sprintf("symbol %s is not message", symbol)
	return "", errors.New(errMsg)
}

func (client *GrpcClient) InvokeGrpc(grpcMethod string, requestBody string, rpcHeaders []string) (*CommonResp, error) {
	in := strings.NewReader(requestBody)
	options := grpcurl.FormatOptions{
		EmitJSONDefaultFields: true,
		AllowUnknownFields:    false,
		IncludeTextSeparator:  false,
	}
	rf, formatter, _ := grpcurl.RequestParserAndFormatter(grpcurl.FormatJSON, client.descSource, in, options)

	eventHandler := &GrpcEventHandler{
		Formatter: formatter,
	}
	invokeErr := grpcurl.InvokeRPC(client.ctx, client.descSource, client.conn, grpcMethod,
		rpcHeaders, eventHandler, rf.Next)
	if invokeErr != nil {
		return nil, invokeErr
	}

	code := eventHandler.GrpcStatus.Code()
	if code != codes.OK {
		m := make(map[string]string)
		m["code"] = code.String()
		m["message"] = eventHandler.GrpcStatus.Message()
		jsonBytes, _ := json.Marshal(m)
		return &CommonResp{
			StatusCode: code.String(),
			RespBody:   string(jsonBytes),
			RespHeader: eventHandler.GrpcMetadata,
		}, nil
	}

	return &CommonResp{
		StatusCode: code.String(),
		RespBody:   eventHandler.GrpcBody,
		RespHeader: eventHandler.GrpcMetadata,
	}, nil
}
