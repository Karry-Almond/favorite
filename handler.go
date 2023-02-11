package main

import (
	"context"
	// "errors"
	api "favorite/kitex_gen/api"

	"github.com/karry-almond/packages/favoriteDB"
	// "log"
	// "golang.org/x/tools/go/analysis/passes/nilfunc"
)

// HelloImpl implements the last service interface defined in the IDL.
type FavoriteImpl struct{}

// Action implements the FavoriteImpl interface.
func (s *FavoriteImpl) Action(ctx context.Context, req *api.DouyinFavoriteActionRequest) (resp *api.DouyinFavoriteActionResponse, err error) {
	// 类型是点赞请求
	if req.ActionType == 1 {
		if _, err = favoriteDB.NewFavorite(req.UserId, req.VideoId); err != nil {
			return &api.DouyinFavoriteActionResponse{
				// status_code = 1 表示点赞失败
				StatusCode: 1,
				StatusMsg:  err.Error()}, nil
		}
	} else if req.ActionType == 2 {
		if _, err = favoriteDB.CancelFavorite(req.UserId, req.VideoId); err != nil {
			return &api.DouyinFavoriteActionResponse{
				// status_code = 2 表示取消点赞失败
				StatusCode: 2,
				StatusMsg:  err.Error()}, nil
		}
	} else {
		return &api.DouyinFavoriteActionResponse{
			// status_code = 3 表示类型信息错误
			StatusCode: 3,
			StatusMsg:  "action_type error"}, nil
	}

	return &api.DouyinFavoriteActionResponse{
		// status_code = 0 表示操作成功
		StatusCode: 0,
		StatusMsg:  "action success"}, nil
}

// List implements the FavoriteImpl interface.
func (s *FavoriteImpl) List(ctx context.Context, req *api.DouyinFavoriteListResponse) (resp *api.DouyinFavoriteListRequest, err error) {
	// TODO: Your code here...
	return
}
