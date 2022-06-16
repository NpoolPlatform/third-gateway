package client

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	npool "github.com/NpoolPlatform/message/npool/thirdgateway"
	"github.com/NpoolPlatform/third-gateway/pkg/test-init" //nolint
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

func TestClient(t *testing.T) {
	_ = NotifyEmail(context.Background(), &npool.NotifyEmailRequest{}) //nolint
	// Here won't pass test due to we always test with localhost
}
