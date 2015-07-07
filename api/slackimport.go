// Copyright (c) 2015 Spinpunch, Inc. All Rights Reserved.
// See License.txt for license information.

package api

import (
	l4g "code.google.com/p/log4go"
	"encoding/json"
	"github.com/mattermost/platform/model"
	"io"
)

type SlackChannel struct {
	Id      string            `json:"id"`
	Name    string            `json:"name"`
	Members []string          `json:"members"`
	Topic   map[string]string `json:"topic"`
}

type SlackUser struct {
	Id       string            `json:"id"`
	UserName string            `json:"name"`
	RealName string            `json:"real_name"`
	Profile  map[string]string `json:"profile"`
}

type SlackPost struct {
	User      string `json:"user"`
	Text      string `json:"text"`
	TimeStamp string `json:"ts"`
}

func SlackParseChannels(data io.Reader) []SlackChannel {
	decoder := json.NewDecoder(data)

	var channels []SlackChannel
	if err := decoder.Decode(&channels); err != nil {
		return make([]SlackChannel, 0)
	}
	return channels
}

func SlackParseUsers(data io.Reader) []SlackUser {
	decoder := json.NewDecoder(data)

	var users []SlackUser
	if err := decoder.Decode(&users); err != nil {
		return make([]SlackUser, 0)
	}
	return users
}

func SlackParsePosts(data io.Reader) []SlackPost {
	decoder := json.NewDecoder(data)

	var posts []SlackPost
	if err := decoder.Decode(&posts); err != nil {
		return make([]SlackPost, 0)
	}
	return posts
}

func SlackAddUsers(c *Context, teamId string, slackusers []SlackUser) map[string]*model.User {
	// Get team
	var team *model.Team
	if result := <-Srv.Store.Team().Get(teamId); result.Err != nil {
		c.Err = result.Err
		return make(map[string]*model.User)
	} else {
		team = result.Data.(*model.Team)
	}

	// Add users
	addedUsers := make(map[string]*model.User)
	for _, sUser := range slackusers {
		newUser := model.User{
			TeamId:   teamId,
			Username: sUser.UserName,
			FullName: sUser.RealName,
			Email:    sUser.Profile["email"],
		}
		mUser := CreateUser(c, team, &newUser)
		addedUsers[sUser.Id] = mUser
	}

	return addedUsers
}

func SlackAddPosts(c *Context, channel *model.Channel, posts []SlackPost, users map[string]*model.User) {
	for _, sPost := range posts {
		newPost := model.Post{
			UserId:    users[sPost.User].Id,
			ChannelId: channel.Id,
			Message:   sPost.Text,
		}
		CreatePost(c, &newPost, false)
	}
}

func SlackAddChannels(c *Context, teamId string, slackchannels []SlackChannel, posts map[string][]SlackPost, users map[string]*model.User) map[string]*model.Channel {
	// Add channels
	addedChannels := make(map[string]*model.Channel)
	for _, sChannel := range slackchannels {
		newChannel := model.Channel{
			TeamId:      teamId,
			Type:        model.CHANNEL_OPEN,
			DisplayName: sChannel.Name,
			Name:        sChannel.Name,
			Description: sChannel.Topic["value"],
		}
		mChannel, err := CreateChannel(c, &newChannel, "", false)
		if err != nil {
			l4g.Debug("Failed to import: %s", newChannel.DisplayName)
		}
		addedChannels[sChannel.Id] = mChannel
		SlackAddPosts(c, mChannel, posts[sChannel.Name], users)
	}

	return addedChannels
}
