package raft

import (
	"CrazyCollin/personalProjects/raft-db/pkg/log"
	"CrazyCollin/personalProjects/raft-db/pkg/protocol"
	"google.golang.org/grpc"
)

type RaftClientEnd struct {
	id             uint64
	addr           string
	conns          []*grpc.ClientConn
	raftServiceCli *protocol.RaftServiceClient
}

func NewRaftClientEnd(id uint64, addr string) *RaftClientEnd {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.MainLogger.Error().Msgf("failed to connect:%v", err)
	}
	conns := []*grpc.ClientConn{}
	conns = append(conns, conn)
	rpcClient := protocol.NewRaftServiceClient(conn)
	return &RaftClientEnd{
		id:             id,
		addr:           addr,
		conns:          conns,
		raftServiceCli: &rpcClient,
	}
}

func (rce *RaftClientEnd) Id() uint64 {
	return rce.id
}

func (rce *RaftClientEnd) GetRaftServiceCli() *protocol.RaftServiceClient {
	return rce.raftServiceCli
}

func (rce *RaftClientEnd) CloseConn() {
	for _, conn := range rce.conns {
		conn.Close()
	}
}
