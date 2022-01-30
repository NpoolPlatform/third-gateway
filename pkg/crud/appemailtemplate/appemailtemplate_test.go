package appemailtemplate

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

func assertTemplate(t *testing.T, actual, expected *npool.AppEmailTemplate) {
	assert.Equal(t, actual.AppID, expected.AppID)
	assert.Equal(t, actual.LangID, expected.LangID)
	assert.Equal(t, actual.UsedFor, expected.UsedFor)
	assert.Equal(t, actual.Sender, expected.Sender)
	assert.Equal(t, actual.ReplyTos, expected.ReplyTos)
	assert.Equal(t, actual.CCTos, expected.CCTos)
	assert.Equal(t, actual.Subject, expected.Subject)
	assert.Equal(t, actual.Body, expected.Body)
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	template := npool.AppEmailTemplate{
		AppID:    uuid.New().String(),
		LangID:   uuid.New().String(),
		UsedFor:  "For test",
		Sender:   "kikakkz@hotmail.com",
		ReplyTos: []string{"kkkk@npool.cc", "kkkkkkkkk@npool.cc"},
		CCTos:    []string{"aaaa@npool.cc", "aaaaaaaaa@npool.cc"},
		Subject:  "For test",
		Body:     "<html><body>For test</body></html>",
	}

	resp, err := Create(context.Background(), &npool.CreateAppEmailTemplateRequest{
		Info: &template,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{}.String())
		assertTemplate(t, resp.Info, &template)
	}

	resp1, err := Get(context.Background(), &npool.GetAppEmailTemplateRequest{
		ID: resp.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assertTemplate(t, resp1.Info, &template)
	}

	template.ID = resp.Info.ID
	resp2, err := Update(context.Background(), &npool.UpdateAppEmailTemplateRequest{
		Info: &template,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp2.Info.ID, resp.Info.ID)
		assertTemplate(t, resp2.Info, &template)
	}

	template.ID = resp.Info.ID
	resp3, err := GetByAppLangUsedFor(context.Background(), &npool.GetAppEmailTemplateByAppLangUsedForRequest{
		AppID:   template.AppID,
		LangID:  template.LangID,
		UsedFor: template.UsedFor,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp3.Info.ID, resp.Info.ID)
		assertTemplate(t, resp3.Info, &template)
	}
}
