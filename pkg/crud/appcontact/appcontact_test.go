package appcontact

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	npool "github.com/NpoolPlatform/message/npool/thirdgateway"
	"github.com/NpoolPlatform/third-gateway/pkg/test-init" //nolint

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

func assertTemplate(t *testing.T, actual, expected *npool.AppContact) {
	assert.Equal(t, actual.AppID, expected.AppID)
	assert.Equal(t, actual.UsedFor, expected.UsedFor)
	assert.Equal(t, actual.Account, expected.Account)
	assert.Equal(t, actual.AccountType, expected.AccountType)
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	template := npool.AppContact{
		AppID:       uuid.New().String(),
		UsedFor:     "For test",
		Account:     "kikakkz@hotmail.com",
		AccountType: "email",
	}

	resp, err := Create(context.Background(), &npool.CreateAppContactRequest{
		Info: &template,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{}.String())
		assertTemplate(t, resp.Info, &template)
	}

	resp1, err := Get(context.Background(), &npool.GetAppContactRequest{
		ID: resp.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assertTemplate(t, resp1.Info, &template)
	}

	template.ID = resp.Info.ID
	resp2, err := Update(context.Background(), &npool.UpdateAppContactRequest{
		Info: &template,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp2.Info.ID, resp.Info.ID)
		assertTemplate(t, resp2.Info, &template)
	}

	template.ID = resp.Info.ID
	resp3, err := GetByAppUsedForAccountType(context.Background(), &npool.GetAppContactByAppUsedForAccountTypeRequest{
		AppID:       template.AppID,
		UsedFor:     template.UsedFor,
		AccountType: template.AccountType,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp3.Info.ID, resp.Info.ID)
		assertTemplate(t, resp3.Info, &template)
	}
}
