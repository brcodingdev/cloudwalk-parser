package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/brcodingdev/cloudwalk-parser/internal/pkg/model"
	"github.com/redis/go-redis/v9"
)

// Match repository to store and get information from the database
type Match interface {
	// Add insert new match record
	Add(match *model.Match) error
	// FindAll find all matches stored
	FindAll() ([]model.Match, error)
}

// RedisMatch structure of redis database
type RedisMatch struct {
	ctx context.Context
	rds *redis.Client
}

// NewRedisMatch creates new instance of implementation of redis database
func NewRedisMatch(
	context context.Context,
	client *redis.Client,
) *RedisMatch {
	return &RedisMatch{
		ctx: context,
		rds: client,
	}
}

// Add ...
func (r RedisMatch) Add(match *model.Match) error {
	serial, err := json.Marshal(match)
	if err != nil {
		return err
	}

	err = r.rds.HSet(
		r.ctx,
		fmt.Sprintf("game"),
		fmt.Sprintf("game%d", match.ID),
		serial,
	).Err()

	if err != nil {
		return err
	}

	return nil
}

// Clean cleans all records on database
func (r RedisMatch) Clean() {
	r.rds.FlushAll(r.ctx)
}

// FindAll ...
func (r RedisMatch) FindAll() ([]model.Match, error) {
	hash := r.rds.HGetAll(r.ctx, "game").Val()
	var matches []model.Match
	for _, v := range hash {
		var match model.Match
		err := json.Unmarshal([]byte(v), &match)
		if err != nil {
			return nil, err
		}
		matches = append(matches, match)
	}
	return matches, nil
}
