package main

//设计一个简化版的推特(Twitter)，可以让用户实现发送推文，关注/取消关注其他用户，能够看见关注人（包括自己）的最近十条推文。你的设计需要支持以下的几个功能：
//
//postTweet(userId, tweetId): 创建一条新的推文
//getNewsFeed(userId): 检索最近的十条推文。每个推文都必须是由此用户关注的人或者是用户自己发出的。推文必须按照时间顺序由最近的开始排序。
//follow(followerId, followeeId): 关注一个用户
//unfollow(followerId, followeeId): 取消关注一个用户

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */

import (
	"sort"
	"time"
)

type Twitter struct {
	Users map[int]*User
}

type User struct {
	Id         int
	Tweets     []*Tweet
	Followlist map[int]*User
}

type Tweet struct {
	Id   int
	Time int64
}

/** Initialize your data structure here. */
// func Constructor() Twitter {
// 	return Twitter{Users: make(map[int]*User)}
// }

func (this *Twitter) CreateUser(Id int) {
	if this.Users[Id] == nil {
		this.Users[Id] = &User{
			Id:         Id,
			Followlist: make(map[int]*User),
		}
	}
}

/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.CreateUser(userId)
	t := &Tweet{Id: tweetId, Time: time.Now().UnixNano()}
	this.Users[userId].Tweets = append(this.Users[userId].Tweets, t)
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) (ids []int) {
	this.CreateUser(userId)

	tws := make([]*Tweet, 0)
	tws = append(tws, this.Users[userId].Tweets...)

	for _, u := range this.Users[userId].Followlist {
		if u.Id == userId {
			continue
		}
		tws = append(tws, u.Tweets...)
	}
	sort.Slice(tws, func(i, j int) bool {
		return tws[i].Time > tws[j].Time
	})

	for i := 0; i < len(tws); i++ {
		ids = append(ids, tws[i].Id)
		if len(ids) == 10 {
			break
		}
	}
	return
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int) {
	this.CreateUser(followerId)
	this.CreateUser(followeeId)
	this.Users[followerId].Followlist[followeeId] = this.Users[followeeId]
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int) {
	this.CreateUser(followerId)
	this.CreateUser(followeeId)
	if this.Users[followerId].Followlist[followeeId] != nil {
		delete(this.Users[followerId].Followlist, followeeId)
	}
}
