package appsmstemplate

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

func assertTemplate(t *testing.T, actual, expected *npool.AppSMSTemplate) {
	assert.Equal(t, actual.AppID, expected.AppID)
	assert.Equal(t, actual.LangID, expected.LangID)
	assert.Equal(t, actual.UsedFor, expected.UsedFor)
	assert.Equal(t, actual.Subject, expected.Subject)
	assert.Equal(t, actual.Message, expected.Message)
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	template := npool.AppSMSTemplate{
		AppID:   uuid.New().String(),
		LangID:  uuid.New().String(),
		UsedFor: "For test",
		Subject: "For test",
		Message: "Test test test",
	}

	resp, err := Create(context.Background(), &npool.CreateAppSMSTemplateRequest{
		Info: &template,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{}.String())
		assertTemplate(t, resp.Info, &template)
	}

	resp1, err := Get(context.Background(), &npool.GetAppSMSTemplateRequest{
		ID: resp.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assertTemplate(t, resp1.Info, &template)
	}

	template.ID = resp.Info.ID
	resp2, err := Update(context.Background(), &npool.UpdateAppSMSTemplateRequest{
		Info: &template,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp2.Info.ID, resp.Info.ID)
		assertTemplate(t, resp2.Info, &template)
	}

	template.ID = resp.Info.ID
	resp3, err := GetByAppLangUsedFor(context.Background(), &npool.GetAppSMSTemplateByAppLangUsedForRequest{
		AppID:   template.AppID,
		LangID:  template.LangID,
		UsedFor: template.UsedFor,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp3.Info.ID, resp.Info.ID)
		assertTemplate(t, resp3.Info, &template)
	}
}
