package redis

import (
	"strconv"
)

// AddViewNums 阅读加一
func AddViewNums(articleID string) error {
	nums, _ := GetViewNums(articleID)
	set := client.HSet("article_view", articleID, nums+1)
	return set.Err()
}

// GetViewNums 查询阅读量
func GetViewNums(articleID string) (int64, error) {
	get := client.HGet("article_view", articleID)
	value, _ := strconv.ParseInt(get.Val(), 10, 64)
	return value, get.Err()
}

// AddLikeNums 喜欢加一
func AddLikeNums(articleID string) error {
	nums, _ := GetLikeNums(articleID)
	set := client.HSet("article_like", articleID, nums+1)
	return set.Err()
}

// GetLikeNums 查询喜欢数量
func GetLikeNums(articleID string) (int64, error) {
	get := client.HGet("article_like", articleID)
	value, _ := strconv.ParseInt(get.Val(), 10, 64)
	return value, get.Err()
}
