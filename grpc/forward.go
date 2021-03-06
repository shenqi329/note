package grpc

import (
	easynoteError "easynote/error"
	"github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	//"google.golang.org/grpc"
	"easynote/util/key"
	grpcPb "im/logicserver/grpc/pb"
	"log"
)

type HandleFunc func(ctx context.Context, request proto.Message) (proto.Message, error)

type HandleFuncInfo struct {
	handle       HandleFunc
	responseType grpcPb.MessageType
}

type Forward struct {
	handleFuncMap map[grpcPb.MessageType]*HandleFuncInfo
}

func (r *Forward) AddHandleFunc(messageType grpcPb.MessageType, responseType grpcPb.MessageType, handle HandleFunc) {
	if r.handleFuncMap == nil {
		r.handleFuncMap = make(map[grpcPb.MessageType]*HandleFuncInfo)
	}
	r.handleFuncMap[messageType] = &HandleFuncInfo{
		handle:       handle,
		responseType: responseType,
	}
}

func (r *Forward) ForwardTLP(ctx context.Context, request *grpcPb.ForwardTLPRequest) (*grpcPb.ForwardTLPResponse, error) {

	rpcResponse := &grpcPb.ForwardTLPResponse{
		Code:   easynoteError.CommonInternalServerError,
		Desc:   easynoteError.ErrorCodeToText(easynoteError.CommonInternalServerError),
		ConnId: request.RpcInfo.ConnId,
	}

	if request.RpcInfo == nil {
		log.Println("RpcInfo为空")
		return rpcResponse, nil
	}

	ctx = context.WithValue(ctx, key.RpcInfo, request.RpcInfo)

	handleFuncInfo := r.handleFuncMap[(grpcPb.MessageType)(request.MessageType)]

	if handleFuncInfo == nil {
		log.Print("没有对应的处理函数")
		return rpcResponse, nil
	}

	protoMessage := grpcPb.Factory((grpcPb.MessageType)(request.MessageType))
	if protoMessage == nil {
		log.Println("未知消息类型")
		return rpcResponse, nil
	}
	log.Print(protoMessage.String())
	err := proto.Unmarshal(request.ProtoBuf, protoMessage)
	if err != nil {
		log.Println(err.Error())
		return rpcResponse, nil
	}

	protoReply, err := handleFuncInfo.handle(ctx, protoMessage)
	if err != nil {
		log.Println(err.Error())
		return rpcResponse, nil
	}
	if handleFuncInfo.responseType == 0 {
		rpcResponse = &grpcPb.ForwardTLPResponse{
			Code:        easynoteError.CommonSuccess,
			Desc:        easynoteError.ErrorCodeToText(easynoteError.CommonSuccess),
			MessageType: (uint32)(handleFuncInfo.responseType),
		}
		return rpcResponse, nil
	}
	if protoReply == nil {
		return rpcResponse, nil
	}

	log.Println(protoReply.String())
	log.Println(request.RpcInfo.ConnId)
	protoBuf, err := proto.Marshal(protoReply)
	log.Println(protoBuf)

	if err != nil {
		log.Println(err.Error())
		return rpcResponse, nil
	}

	rpcResponse = &grpcPb.ForwardTLPResponse{
		Code:        easynoteError.CommonSuccess,
		Desc:        easynoteError.ErrorCodeToText(easynoteError.CommonSuccess),
		MessageType: (uint32)(handleFuncInfo.responseType),
		ProtoBuf:    protoBuf,
		ConnId:      request.RpcInfo.ConnId,
	}
	return rpcResponse, nil
}
