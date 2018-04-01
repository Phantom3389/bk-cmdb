/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.
 * Copyright (C) 2017-2018 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except 
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and 
 * limitations under the License.
 */
 
package userapi

import (
	"configcenter/src/common"
	"configcenter/src/source_controller/common/commondata"
	"fmt"
)

//GetUserAPI  获取user api 列表
func (cli *Client) GetUserAPI(input commondata.ObjQueryInput) (int, *common.APIRsp, error) {
	url := fmt.Sprintf("%s/host/v1/userapi/search", cli.GetAddress())
	return cli.GetRequestInfoEx(common.HTTPSelectPost, input, url)
}

//Detail 获取一个user api 的详情
func (cli *Client) Detail(appID, id string) (int, *common.APIRsp, error) {
	url := fmt.Sprintf("%s/host/v1/userapi/detail/%s/%s", cli.GetAddress(), appID, id)
	return cli.GetRequestInfoEx(common.HTTPSelectGet, nil, url)
}