package wows

import "github.com/imroc/req"

func (c *Client) GetPlayers(search string) []Players {
	res := c.httpGet("account/list/", req.Param{
		"search": search,
	})
	var ret PlayersResponse
	c.mustUnmarshal(res.Bytes(), &ret)
	return ret.Data
}

func (c *Client) GetPlayerPersonalData(accountId int) PlayerPersonalData {
	res := c.httpGet("account/info/", req.Param{
		"account_id": accountId,
	})
	var ret PlayerPersonalDataResponse
	c.mustUnmarshal(res.Bytes(), &ret)
	return ret.Data[accountId]
}

func (c *Client) GetPlayerAchievements(accountId int) PlayerAchievements {
	res := c.httpGet("account/achievements/", req.Param{
		"account_id": accountId,
	})
	var ret PlayerAchievementsResponse
	c.mustUnmarshal(res.Bytes(), &ret)
	return ret.Data[accountId]
}
