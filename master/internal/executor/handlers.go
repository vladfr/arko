package executor

import (
	"bytes"
	"fmt"

	"github.com/fullstorydev/grpcurl"
	"github.com/golang/protobuf/proto"
	"github.com/jhump/protoreflect/desc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// BufferedEventHandler implements grpcurl.InvocationEventHandler
type BufferedEventHandler struct {
	buf          *bytes.Buffer
	descSource   grpcurl.DescriptorSource
	formatter    grpcurl.Formatter
	NumResponses int
	Status       *status.Status
}

// NewBufferedEventHandler returns an InvocationEventHandler that logs the response messages to a string buffer
func NewBufferedEventHandler(buf *bytes.Buffer, descSource grpcurl.DescriptorSource, formatter grpcurl.Formatter) *BufferedEventHandler {
	return &BufferedEventHandler{
		buf:        buf,
		descSource: descSource,
		formatter:  formatter,
	}
}

// OnResolveMethod is called with a descriptor of the method that is being invoked.
func (h *BufferedEventHandler) OnResolveMethod(dsc *desc.MethodDescriptor) {

}

// OnSendHeaders is called with the request metadata that is being sent.
func (h *BufferedEventHandler) OnSendHeaders(m metadata.MD) {

}

// OnReceiveHeaders is called when response headers have been received.
func (h *BufferedEventHandler) OnReceiveHeaders(m metadata.MD) {

}

// OnReceiveResponse is called for each response message received.
func (h *BufferedEventHandler) OnReceiveResponse(resp proto.Message) {
	h.NumResponses++
	if respStr, err := h.formatter(resp); err != nil {
		fmt.Printf("Failed to format response message %d: %v\n", h.NumResponses, err)
	} else {
		fmt.Fprintln(h.buf, respStr)
	}
}

// OnReceiveTrailers is called when response trailers and final RPC status have been received.
func (h *BufferedEventHandler) OnReceiveTrailers(stat *status.Status, m metadata.MD) {
	h.Status = stat
}
