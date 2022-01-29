package code

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"

	"github.com/go-redis/redis/v8"

	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

const (
	redisTimeout = 5 * time.Second
)

type UserCode struct {
	AppID       uuid.UUID
	Account     string
	AccountType string
	UsedFor     string
	Code        string
	ExpireAt    time.Time
}

func (c *UserCode) Key() string {
	return fmt.Sprintf("verify-code:%v:%v:%v:%v", c.AppID, c.AccountType, c.Account, c.UsedFor)
}

func CreateCodeCache(ctx context.Context, code *UserCode) error {
	cli, err := redis2.GetClient()
	if err != nil {
		return xerrors.Errorf("fail get redis client: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, redisTimeout)
	defer cancel()

	body, err := json.Marshal(code)
	if err != nil {
		return xerrors.Errorf("fail marshal code: %v", err)
	}

	err = cli.Set(ctx, code.Key(), body, time.Until(code.ExpireAt)).Err()
	if err != nil {
		return xerrors.Errorf("fail create code cache: %v", err)
	}

	return nil
}

func VerifyCodeCache(ctx context.Context, code *UserCode) error {
	cli, err := redis2.GetClient()
	if err != nil {
		return xerrors.Errorf("fail get redis client: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, redisTimeout)
	defer cancel()

	val, err := cli.Get(ctx, code.Key()).Result()
	if err == redis.Nil {
		return xerrors.Errorf("code not found in redis")
	} else if err != nil {
		return xerrors.Errorf("fail get code: %v", err)
	}

	user := UserCode{}
	err = json.Unmarshal([]byte(val), &user)
	if err != nil {
		return xerrors.Errorf("fail unmarshal val: %v", err)
	}

	if user.Code != code.Code {
		return xerrors.Errorf("invalid code")
	}

	if time.Now().After(user.ExpireAt) {
		return xerrors.Errorf("code expired")
	}

	return nil
}

func DeleteCodeCache(ctx context.Context, code *UserCode) error {
	cli, err := redis2.GetClient()
	if err != nil {
		return xerrors.Errorf("fail get redis client: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, redisTimeout)
	defer cancel()

	err = cli.Del(ctx, code.Key()).Err()
	if err != nil {
		return xerrors.Errorf("fail delete code: %v", err)
	}

	return nil
}