package handler

import (
	"context"
	"strings"
	"sync"

	searchPB "github.com/ygpark2/njro/service/search/proto"
)

// SearchGRPCServer implements the standard gRPC SearchServiceServer interface
type SearchGRPCServer struct {
	searchPB.UnimplementedSearchServiceServer
	mu   sync.RWMutex
	docs map[string]string // id -> text (In-memory search index)
}

// NewSearchGRPCServer returns an instance of SearchGRPCServer
func NewSearchGRPCServer() *SearchGRPCServer {
	return &SearchGRPCServer{
		docs: make(map[string]string),
	}
}

// Index adds or updates a document in the search index
func (s *SearchGRPCServer) Index(ctx context.Context, req *searchPB.IndexRequest) (*searchPB.IndexResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	doc := req.GetDocument()
	if doc != nil && doc.GetId() != "" {
		s.docs[doc.GetId()] = doc.GetText()
	}

	return &searchPB.IndexResponse{}, nil
}

// Search queries the search index by keyword
func (s *SearchGRPCServer) Search(ctx context.Context, req *searchPB.SearchRequest) (*searchPB.SearchResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	keyword := strings.ToLower(req.GetKeyword())
	var matches []*searchPB.Document

	for id, text := range s.docs {
		if strings.Contains(strings.ToLower(text), keyword) {
			matches = append(matches, &searchPB.Document{
				Id:   id,
				Text: text,
			})
		}
	}

	return &searchPB.SearchResponse{
		Documents: matches,
	}, nil
}
