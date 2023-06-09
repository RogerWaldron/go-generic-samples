package main

import (
	"context"
	"fmt"
)

const (
	RetrieveUserByUsernamePath         = "/users/%s"
	RetrieveFollowagePath              = "/followage/%s/%s"
	RetrieveChannelChattersPath        = "/chatters/%s"
	RetrieveChannelSubscriberCountPath = "/subs/%s/count?token=%s"
	RetrieveUserStreamPath             = "/streams/%s"
	RetrieveChannelEmotesPath          = "/emotes/%s"
)

type Service struct {
	httpclient.HttpClient
}

func (service *Service) RetrieveUserByUsername(ctx context.Context, username string) (*types.UserResponse, error) {
	res, err := service.Get(ctx, &httpclient.RequestOptions{Path: fmt.Sprintf(RetrieveUserByUsernamePath, username)})
	return utils.HandleResponse(res, err, &types.UserResponse{})
}

func (service *Service) RetrieveChannelChatters(ctx context.Context, channelName string) (*types.ChannelChattersResponse, error) {
	res, err := service.Get(ctx, &httpclient.RequestOptions{Path: fmt.Sprintf(RetrieveChannelChattersPath, channelName)})
	return utils.HandleResponse(res, err, &types.ChannelChattersResponse{})
}

func (service *Service) RetrieveChannelSubscriberCount(ctx context.Context, channelName, token string) (*types.SubscriberCountResponse, error) {
	res, err := service.Get(ctx, &httpclient.RequestOptions{Path: fmt.Sprintf(RetrieveChannelSubscriberCountPath, channelName, token)})
	return utils.HandleResponse(res, err, &types.SubscriberCountResponse{})
}

func (service *Service) RetireveChannelStream(ctx context.Context, username string) (*types.Stream, error) {
	res, err := service.Get(ctx, &httpclient.RequestOptions{Path: fmt.Sprintf(RetrieveUserStreamPath, username)})
	return utils.HandleResponse(res, err, &types.Stream{})
}

func (service *Service) RetrieveChannelEmotes(ctx context.Context, channelName string) (*map[string]map[string]common.Emote, error) {
	res, err := service.Get(ctx, &httpclient.RequestOptions{Path: fmt.Sprintf(RetrieveChannelEmotesPath, channelName)})
	return utils.HandleResponse(res, err, &map[string]map[string]common.Emote{})
}