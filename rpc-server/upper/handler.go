package main

import (
	"context"
	"strings"
	upper "upper/kitex_gen/upper"
)

// UpperServiceImpl implements the last service interface defined in the IDL.
type UpperServiceImpl struct{}

// Toupper implements the UpperServiceImpl interface.
func (s *UpperServiceImpl) Toupper(ctx context.Context, req *upper.NormalRequest) (resp *upper.UpperResponse, err error) {
	// TODO: Your code here...
	resp = &upper.UpperResponse{
		Result_: strings.ToUpper(req.Message),
	}
	return
}
