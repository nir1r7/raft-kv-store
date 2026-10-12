
package transport

import (
	"context"
	"github.com/nir1r7/raft-kv-store/internal/kv"
	"github.com/nir1r7/raft-kv-store/proto/kvpb"
)

type KVServer struct {
	kvpb.UnimplementedKVServer
	store *kv.Store
}

func NewKVServer(store *kv.Store) *KVServer {
	return &KVServer{store: store}
}

func (s *KVServer) Get(ctx context.Context, req *kvpb.GetRequest) (*kvpb.GetResponse, error){
	v, found := s.store.Get(req.Key)
	return &kvpb.GetResponse{Value: v, Found: found}, nil
}

func (s *KVServer) Put(ctx context.Context, req *kvpb.PutRequest) (*kvpb.PutResponse, error){
	s.store.Put(req.Key, req.Value)
	return &kvpb.PutResponse{}, nil
}

func (s *KVServer) Delete(ctx context.Context, req *kvpb.DeleteRequest) (*kvpb.DeleteResponse, error){
	s.store.Delete(req.Key)
	return &kvpb.DeleteResponse{}, nil
}