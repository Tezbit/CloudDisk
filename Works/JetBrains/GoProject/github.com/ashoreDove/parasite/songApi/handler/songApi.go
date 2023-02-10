package handler

import (
	"context"
	"encoding/json"
	"errors"
	serviceSong "github.com/ashoreDove/parasite-song/proto/song"
	songApi "github.com/ashoreDove/parasite-songApi/proto/songApi"
	log "github.com/micro/go-micro/v2/logger"
)

type SongApi struct {
	SongModuleService serviceSong.SongService
}

func (s *SongApi) GetSongUrl(ctx context.Context, req *songApi.Request, resp *songApi.Response) error {
	log.Info("收到对 /songApi/getSongUrl 访问请求")
	var post map[string]interface{}
	err := json.Unmarshal([]byte(req.Body), &post)
	if err != nil {
		log.Error(err)
		return err
	}
	params := post["params"]
	param := params.(map[string]interface{})
	if _, ok := param["song_id"]; !ok {
		resp.StatusCode = 500
		return errors.New("参数song_id异常")
	}
	log.Info(req)
	info, err := s.SongModuleService.GetSongInfo(context.TODO(), &serviceSong.SongIdRequest{
		SongId: int64(param["song_id"].(float64)),
	})
	if err != nil {
		resp.StatusCode = 400
		return err
	}
	//转类型
	b, err := json.Marshal(info)
	if err != nil {
		return err
	}
	resp.StatusCode = 200
	resp.Body = string(b)
	return nil
}

// SongApi.Call 通过API向外暴露为/songApi/call，接收http请求
// 即：/songApi/call请求会调用go.micro.api.songApi 服务的SongApi.Call方法
func (s *SongApi) Search(ctx context.Context, req *songApi.Request, resp *songApi.Response) error {
	log.Info("收到对 /songApi/search 访问请求")
	var post map[string]interface{}
	err := json.Unmarshal([]byte(req.Body), &post)
	if err != nil {
		log.Error(err)
		return err
	}
	params := post["params"]
	param := params.(map[string]interface{})
	if _, ok := param["keyword"]; !ok {
		resp.StatusCode = 500
		return errors.New("参数keyword异常")
	}
	log.Info(req)
	searchResp, err := s.SongModuleService.Search(context.TODO(), &serviceSong.SearchRequest{Keyword: param["keyword"].(string)})
	if err != nil {
		resp.StatusCode = 400
		return err
	}
	//转类型
	b, err := json.Marshal(searchResp)
	if err != nil {
		return err
	}
	resp.StatusCode = 200
	resp.Body = string(b)
	return nil
}
