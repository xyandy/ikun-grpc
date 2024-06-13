package backend

import (
	"github.com/fullstorydev/grpcurl"
	"github.com/golang/protobuf/proto"
	"github.com/jhump/protoreflect/desc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type CommonResp struct {
	StatusCode string
	RespBody   string
	RespHeader map[string][]string
}

// GrpcEventHandler implements InvocationEventHandler
// https://github.com/fullstorydev/grpcurl/blob/master/invoke.go#L27
type GrpcEventHandler struct {
	Formatter    grpcurl.Formatter
	GrpcStatus   *status.Status
	GrpcBody     string
	GrpcMetadata metadata.MD
}

func (h *GrpcEventHandler) OnResolveMethod(md *desc.MethodDescriptor) {
}

func (h *GrpcEventHandler) OnSendHeaders(md metadata.MD) {
}

// OnReceiveHeaders if grpc code == ok, md will be filled
func (h *GrpcEventHandler) OnReceiveHeaders(md metadata.MD) {
	if len(md) > 0 {
		h.GrpcMetadata = md
	}
}

// OnReceiveResponse if grpc code == ok, this func is called
func (h *GrpcEventHandler) OnReceiveResponse(resp proto.Message) {
	respStr, _ := h.Formatter(resp)
	h.GrpcBody = respStr
}

// OnReceiveTrailers if grpc code != ok, stat and md will be filled
func (h *GrpcEventHandler) OnReceiveTrailers(stat *status.Status, md metadata.MD) {
	h.GrpcStatus = stat
	if len(md) > 0 {
		h.GrpcMetadata = md
	}
}
