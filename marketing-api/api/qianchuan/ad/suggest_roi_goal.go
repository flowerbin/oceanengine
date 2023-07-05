package ad

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/ad"
)

// SuggestRoiGoal 获取支付ROI目标建议
func SuggestRoiGoal(clt *core.SDKClient, accessToken string, req *ad.SuggestRoiGoalRequest) (*ad.SuggestRoiResult, error) {
	var resp ad.SuggestRoiGoalResponse
	if err := clt.Get("v1.0/qianchuan/suggest/roi/goal/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
