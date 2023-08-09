// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: eth/service/node_service.proto

/*
Package service is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package service

import (
	"context"
	"io"
	"net/http"

	"github.com/golang/protobuf/ptypes/empty"
	emptypb "github.com/golang/protobuf/ptypes/empty"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	github_com_prysmaticlabs_prysm_v4_consensus_types_primitives "github.com/prysmaticlabs/prysm/v4/consensus-types/primitives"
	v1 "github.com/prysmaticlabs/prysm/v4/eth/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join
var _ = github_com_prysmaticlabs_prysm_v4_consensus_types_primitives.Epoch(0)
var _ = emptypb.Empty{}
var _ = empty.Empty{}

func request_BeaconNode_GetIdentity_0(ctx context.Context, marshaler runtime.Marshaler, client BeaconNodeClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq emptypb.Empty
	var metadata runtime.ServerMetadata

	msg, err := client.GetIdentity(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_BeaconNode_GetIdentity_0(ctx context.Context, marshaler runtime.Marshaler, server BeaconNodeServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq emptypb.Empty
	var metadata runtime.ServerMetadata

	msg, err := server.GetIdentity(ctx, &protoReq)
	return msg, metadata, err

}

var (
	filter_BeaconNode_ListPeers_0 = &utilities.DoubleArray{Encoding: map[string]int{}, Base: []int(nil), Check: []int(nil)}
)

func request_BeaconNode_ListPeers_0(ctx context.Context, marshaler runtime.Marshaler, client BeaconNodeClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq v1.PeersRequest
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_BeaconNode_ListPeers_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.ListPeers(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_BeaconNode_ListPeers_0(ctx context.Context, marshaler runtime.Marshaler, server BeaconNodeServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq v1.PeersRequest
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_BeaconNode_ListPeers_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.ListPeers(ctx, &protoReq)
	return msg, metadata, err

}

func request_BeaconNode_GetPeer_0(ctx context.Context, marshaler runtime.Marshaler, client BeaconNodeClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq v1.PeerRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["peer_id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "peer_id")
	}

	peer_id, err := runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "peer_id", err)
	}
	protoReq.PeerId = (peer_id)

	msg, err := client.GetPeer(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_BeaconNode_GetPeer_0(ctx context.Context, marshaler runtime.Marshaler, server BeaconNodeServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq v1.PeerRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["peer_id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "peer_id")
	}

	peer_id, err := runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "peer_id", err)
	}
	protoReq.PeerId = (peer_id)

	msg, err := server.GetPeer(ctx, &protoReq)
	return msg, metadata, err

}

func request_BeaconNode_PeerCount_0(ctx context.Context, marshaler runtime.Marshaler, client BeaconNodeClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq emptypb.Empty
	var metadata runtime.ServerMetadata

	msg, err := client.PeerCount(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_BeaconNode_PeerCount_0(ctx context.Context, marshaler runtime.Marshaler, server BeaconNodeServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq emptypb.Empty
	var metadata runtime.ServerMetadata

	msg, err := server.PeerCount(ctx, &protoReq)
	return msg, metadata, err

}

func request_BeaconNode_GetVersion_0(ctx context.Context, marshaler runtime.Marshaler, client BeaconNodeClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq emptypb.Empty
	var metadata runtime.ServerMetadata

	msg, err := client.GetVersion(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_BeaconNode_GetVersion_0(ctx context.Context, marshaler runtime.Marshaler, server BeaconNodeServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq emptypb.Empty
	var metadata runtime.ServerMetadata

	msg, err := server.GetVersion(ctx, &protoReq)
	return msg, metadata, err

}

func request_BeaconNode_GetHealth_0(ctx context.Context, marshaler runtime.Marshaler, client BeaconNodeClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq emptypb.Empty
	var metadata runtime.ServerMetadata

	msg, err := client.GetHealth(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_BeaconNode_GetHealth_0(ctx context.Context, marshaler runtime.Marshaler, server BeaconNodeServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq emptypb.Empty
	var metadata runtime.ServerMetadata

	msg, err := server.GetHealth(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterBeaconNodeHandlerServer registers the http handlers for service BeaconNode to "mux".
// UnaryRPC     :call BeaconNodeServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterBeaconNodeHandlerFromEndpoint instead.
func RegisterBeaconNodeHandlerServer(ctx context.Context, mux *runtime.ServeMux, server BeaconNodeServer) error {

	mux.Handle("GET", pattern_BeaconNode_GetIdentity_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/ethereum.eth.service.BeaconNode/GetIdentity")
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_BeaconNode_GetIdentity_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_BeaconNode_GetIdentity_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_BeaconNode_ListPeers_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/ethereum.eth.service.BeaconNode/ListPeers")
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_BeaconNode_ListPeers_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_BeaconNode_ListPeers_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_BeaconNode_GetPeer_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/ethereum.eth.service.BeaconNode/GetPeer")
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_BeaconNode_GetPeer_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_BeaconNode_GetPeer_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_BeaconNode_PeerCount_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/ethereum.eth.service.BeaconNode/PeerCount")
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_BeaconNode_PeerCount_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_BeaconNode_PeerCount_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_BeaconNode_GetVersion_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/ethereum.eth.service.BeaconNode/GetVersion")
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_BeaconNode_GetVersion_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_BeaconNode_GetVersion_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_BeaconNode_GetHealth_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/ethereum.eth.service.BeaconNode/GetHealth")
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_BeaconNode_GetHealth_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_BeaconNode_GetHealth_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterBeaconNodeHandlerFromEndpoint is same as RegisterBeaconNodeHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterBeaconNodeHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterBeaconNodeHandler(ctx, mux, conn)
}

// RegisterBeaconNodeHandler registers the http handlers for service BeaconNode to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterBeaconNodeHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterBeaconNodeHandlerClient(ctx, mux, NewBeaconNodeClient(conn))
}

// RegisterBeaconNodeHandlerClient registers the http handlers for service BeaconNode
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "BeaconNodeClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "BeaconNodeClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "BeaconNodeClient" to call the correct interceptors.
func RegisterBeaconNodeHandlerClient(ctx context.Context, mux *runtime.ServeMux, client BeaconNodeClient) error {

	mux.Handle("GET", pattern_BeaconNode_GetIdentity_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/ethereum.eth.service.BeaconNode/GetIdentity")
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_BeaconNode_GetIdentity_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_BeaconNode_GetIdentity_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_BeaconNode_ListPeers_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/ethereum.eth.service.BeaconNode/ListPeers")
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_BeaconNode_ListPeers_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_BeaconNode_ListPeers_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_BeaconNode_GetPeer_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/ethereum.eth.service.BeaconNode/GetPeer")
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_BeaconNode_GetPeer_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_BeaconNode_GetPeer_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_BeaconNode_PeerCount_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/ethereum.eth.service.BeaconNode/PeerCount")
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_BeaconNode_PeerCount_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_BeaconNode_PeerCount_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_BeaconNode_GetVersion_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/ethereum.eth.service.BeaconNode/GetVersion")
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_BeaconNode_GetVersion_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_BeaconNode_GetVersion_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_BeaconNode_GetHealth_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/ethereum.eth.service.BeaconNode/GetHealth")
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_BeaconNode_GetHealth_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_BeaconNode_GetHealth_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_BeaconNode_GetIdentity_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3, 2, 4}, []string{"internal", "eth", "v1", "node", "identity"}, ""))

	pattern_BeaconNode_ListPeers_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3, 2, 4}, []string{"internal", "eth", "v1", "node", "peers"}, ""))

	pattern_BeaconNode_GetPeer_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3, 2, 4, 1, 0, 4, 1, 5, 5}, []string{"internal", "eth", "v1", "node", "peers", "peer_id"}, ""))

	pattern_BeaconNode_PeerCount_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3, 2, 4}, []string{"internal", "eth", "v1", "node", "peer_count"}, ""))

	pattern_BeaconNode_GetVersion_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3, 2, 4}, []string{"internal", "eth", "v1", "node", "version"}, ""))

	pattern_BeaconNode_GetHealth_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3, 2, 4}, []string{"internal", "eth", "v1", "node", "health"}, ""))
)

var (
	forward_BeaconNode_GetIdentity_0 = runtime.ForwardResponseMessage

	forward_BeaconNode_ListPeers_0 = runtime.ForwardResponseMessage

	forward_BeaconNode_GetPeer_0 = runtime.ForwardResponseMessage

	forward_BeaconNode_PeerCount_0 = runtime.ForwardResponseMessage

	forward_BeaconNode_GetVersion_0 = runtime.ForwardResponseMessage

	forward_BeaconNode_GetHealth_0 = runtime.ForwardResponseMessage
)
