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

package y3_9_202002181444

import (
	"context"
	"fmt"
	"time"

	"configcenter/src/common"
	"configcenter/src/common/metadata"
	mCommon "configcenter/src/scene_server/admin_server/common"
	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/storage/dal"
)

var (
	groupBaseInfo = mCommon.BaseInfo
)

func initPlatAttr(ctx context.Context, db dal.RDB, conf *upgrader.Config) error {
	objID := common.BKInnerObjIDPlat
	dataRows := []*Attribute{
		{ObjectID: objID, PropertyID: "bk_status", PropertyName: "状态", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeEnum, Option: statusEnum},
		{ObjectID: objID, PropertyID: "bk_cloud_vendor", PropertyName: "云厂商", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeEnum, Option: cloudVendorEnum},
		{ObjectID: objID, PropertyID: "bk_state_name", PropertyName: "所在国家", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeEnum, Option: stateEnum},
		{ObjectID: objID, PropertyID: "bk_province_name", PropertyName: "所在省份", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeEnum, Option: provincesEnum},
		{ObjectID: objID, PropertyID: "bk_vpc_id", PropertyName: "VPC唯一标识", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeInt, Option: ""},
		{ObjectID: objID, PropertyID: "bk_vpc_name", PropertyName: "VPC名称", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		{ObjectID: objID, PropertyID: "bk_account_id", PropertyName: "云账户ID", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeInt, Option: ""},
		{ObjectID: objID, PropertyID: "bk_creator", PropertyName: "创建者", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		{ObjectID: objID, PropertyID: "bk_last_editor", PropertyName: "最后修改人", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
	}

	t := new(time.Time)
	for _, r := range dataRows {
		r.OwnerID = conf.OwnerID
		r.IsPre = true
		if false != r.IsEditable {
			r.IsEditable = true
		}
		r.IsReadOnly = false
		r.CreateTime = t
		r.Creator = common.CCSystemOperatorUserName
		r.LastTime = r.CreateTime
		r.LastEditor = common.CCSystemOperatorUserName
		r.Description = ""

		id, err := db.NextSequence(ctx, common.BKTableNameObjAttDes)
		if err != nil {
			return fmt.Errorf("upgrade y3.9.202002181444, insert plat attrName: %s, but get NextSequence failed, err: %v", r.PropertyName, err)
		}
		r.ID = int64(id)

		if err := db.Table(common.BKTableNameObjAttDes).Insert(ctx, r); err != nil {
			return fmt.Errorf("upgrade y3.9.202002181444, but insert plat attrName: %s, failed, err: %v", r.PropertyName, err)
		}
	}

	return nil
}

var statusEnum = []metadata.EnumVal{
	{ID: "1", Name: "正常", Type: "text"},
	{ID: "2", Name: "异常", Type: "text"},
	{ID: "3", Name: "同步中", Type: "text"},
}

var cloudVendorEnum = []metadata.EnumVal{
	{ID: "1", Name: "aws", Type: "text"},
	{ID: "2", Name: "tencent_cloud", Type: "text"},
}

var stateEnum = []metadata.EnumVal{
	{ID: "AR", Name: "阿根廷", Type: "text"},
	{ID: "AD", Name: "安道尔", Type: "text"},
	{ID: "AE", Name: "阿联酋", Type: "text"},
	{ID: "AF", Name: "阿富汗", Type: "text"},
	{ID: "AG", Name: "安提瓜和巴布达", Type: "text"},
	{ID: "AI", Name: "安圭拉", Type: "text"},
	{ID: "AL", Name: "阿尔巴尼亚", Type: "text"},
	{ID: "AM", Name: "亚美尼亚", Type: "text"},
	{ID: "AO", Name: "安哥拉", Type: "text"},
	{ID: "AQ", Name: "南极洲", Type: "text"},
	{ID: "AS", Name: "美属萨摩亚", Type: "text"},
	{ID: "AT", Name: "奥地利", Type: "text"},
	{ID: "AU", Name: "澳大利亚", Type: "text"},
	{ID: "AW", Name: "阿鲁巴", Type: "text"},
	{ID: "AX", Name: "奥兰群岛", Type: "text"},
	{ID: "AZ", Name: "阿塞拜疆", Type: "text"},
	{ID: "BA", Name: "波黑", Type: "text"},
	{ID: "BB", Name: "巴巴多斯", Type: "text"},
	{ID: "BD", Name: "孟加拉", Type: "text"},
	{ID: "BE", Name: "比利时", Type: "text"},
	{ID: "BF", Name: "布基纳法索", Type: "text"},
	{ID: "BG", Name: "保加利亚", Type: "text"},
	{ID: "BH", Name: "巴林", Type: "text"},
	{ID: "BI", Name: "布隆迪", Type: "text"},
	{ID: "BJ", Name: "贝宁", Type: "text"},
	{ID: "BL", Name: "圣巴泰勒米岛", Type: "text"},
	{ID: "BM", Name: "百慕大", Type: "text"},
	{ID: "BN", Name: "文莱", Type: "text"},
	{ID: "BO", Name: "玻利维亚", Type: "text"},
	{ID: "BQ", Name: "荷兰加勒比区", Type: "text"},
	{ID: "BR", Name: "巴西", Type: "text"},
	{ID: "BS", Name: "巴哈马", Type: "text"},
	{ID: "BT", Name: "不丹", Type: "text"},
	{ID: "BV", Name: "布韦岛", Type: "text"},
	{ID: "BW", Name: "博茨瓦纳", Type: "text"},
	{ID: "BY", Name: "白俄罗斯", Type: "text"},
	{ID: "BZ", Name: "伯利兹", Type: "text"},
	{ID: "CA", Name: "加拿大", Type: "text"},
	{ID: "CC", Name: "科科斯群岛", Type: "text"},
	{ID: "CD", Name: "刚果（金）", Type: "text"},
	{ID: "CF", Name: "中非", Type: "text"},
	{ID: "CG", Name: "刚果（布）", Type: "text"},
	{ID: "CH", Name: "瑞士", Type: "text"},
	{ID: "CI", Name: "科特迪瓦", Type: "text"},
	{ID: "CK", Name: "库克群岛", Type: "text"},
	{ID: "CL", Name: "智利", Type: "text"},
	{ID: "CM", Name: "喀麦隆", Type: "text"},
	{ID: "CN", Name: "中国", Type: "text"},
	{ID: "CO", Name: "哥伦比亚", Type: "text"},
	{ID: "CR", Name: "哥斯达黎加", Type: "text"},
	{ID: "CU", Name: "古巴", Type: "text"},
	{ID: "CV", Name: "佛得角", Type: "text"},
	{ID: "CW", Name: "库拉索", Type: "text"},
	{ID: "CX", Name: "圣诞岛", Type: "text"},
	{ID: "CY", Name: "塞浦路斯", Type: "text"},
	{ID: "CZ", Name: "捷克", Type: "text"},
	{ID: "DE", Name: "德国", Type: "text"},
	{ID: "DJ", Name: "吉布提", Type: "text"},
	{ID: "DK", Name: "丹麦", Type: "text"},
	{ID: "DM", Name: "多米尼克", Type: "text"},
	{ID: "DO", Name: "多米尼加", Type: "text"},
	{ID: "DZ", Name: "阿尔及利亚", Type: "text"},
	{ID: "EC", Name: "厄瓜多尔", Type: "text"},
	{ID: "EE", Name: "爱沙尼亚", Type: "text"},
	{ID: "EG", Name: "埃及", Type: "text"},
	{ID: "EH", Name: "西撒哈拉", Type: "text"},
	{ID: "ER", Name: "厄立特里亚", Type: "text"},
	{ID: "ES", Name: "西班牙", Type: "text"},
	{ID: "ET", Name: "埃塞俄比亚", Type: "text"},
	{ID: "FI", Name: "芬兰", Type: "text"},
	{ID: "FJ", Name: "斐济群岛", Type: "text"},
	{ID: "FK", Name: "马尔维纳斯群岛（福克兰）", Type: "text"},
	{ID: "FM", Name: "密克罗尼西亚联邦", Type: "text"},
	{ID: "FO", Name: "法罗群岛", Type: "text"},
	{ID: "FR", Name: "法国", Type: "text"},
	{ID: "GA", Name: "加蓬", Type: "text"},
	{ID: "GB", Name: "英国", Type: "text"},
	{ID: "GD", Name: "格林纳达", Type: "text"},
	{ID: "GE", Name: "格鲁吉亚", Type: "text"},
	{ID: "GF", Name: "法属圭亚那", Type: "text"},
	{ID: "GG", Name: "根西岛", Type: "text"},
	{ID: "GH", Name: "加纳", Type: "text"},
	{ID: "GI", Name: "直布罗陀", Type: "text"},
	{ID: "GL", Name: "格陵兰", Type: "text"},
	{ID: "GM", Name: "冈比亚", Type: "text"},
	{ID: "GN", Name: "几内亚", Type: "text"},
	{ID: "GP", Name: "瓜德罗普", Type: "text"},
	{ID: "GQ", Name: "赤道几内亚", Type: "text"},
	{ID: "GR", Name: "希腊", Type: "text"},
	{ID: "GS", Name: "南乔治亚岛和南桑威奇群岛", Type: "text"},
	{ID: "GT", Name: "危地马拉", Type: "text"},
	{ID: "GU", Name: "关岛", Type: "text"},
	{ID: "GW", Name: "几内亚比绍", Type: "text"},
	{ID: "GY", Name: "圭亚那", Type: "text"},
	{ID: "HM", Name: "赫德岛和麦克唐纳群岛", Type: "text"},
	{ID: "HN", Name: "洪都拉斯", Type: "text"},
	{ID: "HR", Name: "克罗地亚", Type: "text"},
	{ID: "HT", Name: "海地", Type: "text"},
	{ID: "HU", Name: "匈牙利", Type: "text"},
	{ID: "ID", Name: "印尼", Type: "text"},
	{ID: "IE", Name: "爱尔兰", Type: "text"},
	{ID: "IL", Name: "以色列", Type: "text"},
	{ID: "IM", Name: "马恩岛", Type: "text"},
	{ID: "IN", Name: "印度", Type: "text"},
	{ID: "IO", Name: "英属印度洋领地", Type: "text"},
	{ID: "IQ", Name: "伊拉克", Type: "text"},
	{ID: "IR", Name: "伊朗", Type: "text"},
	{ID: "IS", Name: "冰岛", Type: "text"},
	{ID: "IT", Name: "意大利", Type: "text"},
	{ID: "JE", Name: "泽西岛", Type: "text"},
	{ID: "JM", Name: "牙买加", Type: "text"},
	{ID: "JO", Name: "约旦", Type: "text"},
	{ID: "JP", Name: "日本", Type: "text"},
	{ID: "KE", Name: "肯尼亚", Type: "text"},
	{ID: "KG", Name: "吉尔吉斯斯坦", Type: "text"},
	{ID: "KH", Name: "柬埔寨", Type: "text"},
	{ID: "KI", Name: "基里巴斯", Type: "text"},
	{ID: "KM", Name: "科摩罗", Type: "text"},
	{ID: "KN", Name: "圣基茨和尼维斯", Type: "text"},
	{ID: "KP", Name: "朝鲜", Type: "text"},
	{ID: "KR", Name: "韩国", Type: "text"},
	{ID: "KW", Name: "科威特", Type: "text"},
	{ID: "KY", Name: "开曼群岛", Type: "text"},
	{ID: "KZ", Name: "哈萨克斯坦", Type: "text"},
	{ID: "LA", Name: "老挝", Type: "text"},
	{ID: "LB", Name: "黎巴嫩", Type: "text"},
	{ID: "LC", Name: "圣卢西亚", Type: "text"},
	{ID: "LI", Name: "列支敦士登", Type: "text"},
	{ID: "LK", Name: "斯里兰卡", Type: "text"},
	{ID: "LR", Name: "利比里亚", Type: "text"},
	{ID: "LS", Name: "莱索托", Type: "text"},
	{ID: "LT", Name: "立陶宛", Type: "text"},
	{ID: "LU", Name: "卢森堡", Type: "text"},
	{ID: "LV", Name: "拉脱维亚", Type: "text"},
	{ID: "LY", Name: "利比亚", Type: "text"},
	{ID: "MA", Name: "摩洛哥", Type: "text"},
	{ID: "MC", Name: "摩纳哥", Type: "text"},
	{ID: "MD", Name: "摩尔多瓦", Type: "text"},
	{ID: "ME", Name: "黑山", Type: "text"},
	{ID: "MF", Name: "法属圣马丁", Type: "text"},
	{ID: "MG", Name: "马达加斯加", Type: "text"},
	{ID: "MH", Name: "马绍尔群岛", Type: "text"},
	{ID: "MK", Name: "马其顿", Type: "text"},
	{ID: "ML", Name: "马里", Type: "text"},
	{ID: "MM", Name: "缅甸", Type: "text"},
	{ID: "MN", Name: "蒙古国", Type: "text"},
	{ID: "MP", Name: "北马里亚纳群岛", Type: "text"},
	{ID: "MQ", Name: "马提尼克", Type: "text"},
	{ID: "MR", Name: "毛里塔尼亚", Type: "text"},
	{ID: "MS", Name: "蒙塞拉特岛", Type: "text"},
	{ID: "MT", Name: "马耳他", Type: "text"},
	{ID: "MU", Name: "毛里求斯", Type: "text"},
	{ID: "MV", Name: "马尔代夫", Type: "text"},
	{ID: "MW", Name: "马拉维", Type: "text"},
	{ID: "MX", Name: "墨西哥", Type: "text"},
	{ID: "MY", Name: "马来西亚", Type: "text"},
	{ID: "MZ", Name: "莫桑比克", Type: "text"},
	{ID: "NA", Name: "纳米比亚", Type: "text"},
	{ID: "NC", Name: "新喀里多尼亚", Type: "text"},
	{ID: "NE", Name: "尼日尔", Type: "text"},
	{ID: "NF", Name: "诺福克岛", Type: "text"},
	{ID: "NG", Name: "尼日利亚", Type: "text"},
	{ID: "NI", Name: "尼加拉瓜", Type: "text"},
	{ID: "NL", Name: "荷兰", Type: "text"},
	{ID: "NO", Name: "挪威", Type: "text"},
	{ID: "NP", Name: "尼泊尔", Type: "text"},
	{ID: "NR", Name: "瑙鲁", Type: "text"},
	{ID: "NU", Name: "纽埃", Type: "text"},
	{ID: "NZ", Name: "新西兰", Type: "text"},
	{ID: "OM", Name: "阿曼", Type: "text"},
	{ID: "PA", Name: "巴拿马", Type: "text"},
	{ID: "PE", Name: "秘鲁", Type: "text"},
	{ID: "PF", Name: "法属波利尼西亚", Type: "text"},
	{ID: "PG", Name: "巴布亚新几内亚", Type: "text"},
	{ID: "PH", Name: "菲律宾", Type: "text"},
	{ID: "PK", Name: "巴基斯坦", Type: "text"},
	{ID: "PL", Name: "波兰", Type: "text"},
	{ID: "PM", Name: "圣皮埃尔和密克隆", Type: "text"},
	{ID: "PN", Name: "皮特凯恩群岛", Type: "text"},
	{ID: "PR", Name: "波多黎各", Type: "text"},
	{ID: "PS", Name: "巴勒斯坦", Type: "text"},
	{ID: "PT", Name: "葡萄牙", Type: "text"},
	{ID: "PW", Name: "帕劳", Type: "text"},
	{ID: "PY", Name: "巴拉圭", Type: "text"},
	{ID: "QA", Name: "卡塔尔", Type: "text"},
	{ID: "RE", Name: "留尼汪", Type: "text"},
	{ID: "RO", Name: "罗马尼亚", Type: "text"},
	{ID: "RS", Name: "塞尔维亚", Type: "text"},
	{ID: "RU", Name: "俄罗斯", Type: "text"},
	{ID: "RW", Name: "卢旺达", Type: "text"},
	{ID: "SA", Name: "沙特阿拉伯", Type: "text"},
	{ID: "SB", Name: "所罗门群岛", Type: "text"},
	{ID: "SC", Name: "塞舌尔", Type: "text"},
	{ID: "SD", Name: "苏丹", Type: "text"},
	{ID: "SE", Name: "瑞典", Type: "text"},
	{ID: "SG", Name: "新加坡", Type: "text"},
	{ID: "SH", Name: "圣赫勒拿", Type: "text"},
	{ID: "SI", Name: "斯洛文尼亚", Type: "text"},
	{ID: "SJ", Name: "斯瓦尔巴群岛和扬马延岛", Type: "text"},
	{ID: "SK", Name: "斯洛伐克", Type: "text"},
	{ID: "SL", Name: "塞拉利昂", Type: "text"},
	{ID: "SM", Name: "圣马力诺", Type: "text"},
	{ID: "SN", Name: "塞内加尔", Type: "text"},
	{ID: "SO", Name: "索马里", Type: "text"},
	{ID: "SR", Name: "苏里南", Type: "text"},
	{ID: "SS", Name: "南苏丹", Type: "text"},
	{ID: "ST", Name: "圣多美和普林西比", Type: "text"},
	{ID: "SV", Name: "萨尔瓦多", Type: "text"},
	{ID: "SX", Name: "荷属圣马丁", Type: "text"},
	{ID: "SY", Name: "叙利亚", Type: "text"},
	{ID: "SZ", Name: "斯威士兰", Type: "text"},
	{ID: "TC", Name: "特克斯和凯科斯群岛", Type: "text"},
	{ID: "TD", Name: "乍得", Type: "text"},
	{ID: "TF", Name: "法属南部领地", Type: "text"},
	{ID: "TG", Name: "多哥", Type: "text"},
	{ID: "TH", Name: "泰国", Type: "text"},
	{ID: "TJ", Name: "塔吉克斯坦", Type: "text"},
	{ID: "TK", Name: "托克劳", Type: "text"},
	{ID: "TL", Name: "东帝汶", Type: "text"},
	{ID: "TM", Name: "土库曼斯坦", Type: "text"},
	{ID: "TN", Name: "突尼斯", Type: "text"},
	{ID: "TO", Name: "汤加", Type: "text"},
	{ID: "TR", Name: "土耳其", Type: "text"},
	{ID: "TT", Name: "特立尼达和多巴哥", Type: "text"},
	{ID: "TV", Name: "图瓦卢", Type: "text"},
	{ID: "TZ", Name: "坦桑尼亚", Type: "text"},
	{ID: "UA", Name: "乌克兰", Type: "text"},
	{ID: "UG", Name: "乌干达", Type: "text"},
	{ID: "UM", Name: "美国本土外小岛屿", Type: "text"},
	{ID: "UY", Name: "乌拉圭", Type: "text"},
	{ID: "UZ", Name: "乌兹别克斯坦", Type: "text"},
	{ID: "VA", Name: "梵蒂冈", Type: "text"},
	{ID: "VC", Name: "圣文森特和格林纳丁斯", Type: "text"},
	{ID: "VE", Name: "委内瑞拉", Type: "text"},
	{ID: "VG", Name: "英属维尔京群岛", Type: "text"},
	{ID: "VI", Name: "美属维尔京群岛", Type: "text"},
	{ID: "VN", Name: "越南", Type: "text"},
	{ID: "US", Name: "美国", Type: "text"},
	{ID: "VU", Name: "瓦努阿图", Type: "text"},
	{ID: "WF", Name: "瓦利斯和富图纳", Type: "text"},
	{ID: "WS", Name: "萨摩亚", Type: "text"},
	{ID: "YE", Name: "也门", Type: "text"},
	{ID: "YT", Name: "马约特", Type: "text"},
	{ID: "ZA", Name: "南非", Type: "text"},
	{ID: "ZM", Name: "赞比亚", Type: "text"},
	{ID: "ZW", Name: "津巴布韦", Type: "text"},
}

var provincesEnum = []metadata.EnumVal{
	{ID: "110000", Name: "北京市", Type: "text"},
	{ID: "120000", Name: "天津市", Type: "text"},
	{ID: "130000", Name: "河北省", Type: "text"},
	{ID: "140000", Name: "山西省", Type: "text"},
	{ID: "150000", Name: "内蒙古自治区", Type: "text"},
	{ID: "210000", Name: "辽宁省", Type: "text"},
	{ID: "220000", Name: "吉林省", Type: "text"},
	{ID: "230000", Name: "黑龙江省", Type: "text"},
	{ID: "310000", Name: "上海市", Type: "text"},
	{ID: "320000", Name: "江苏省", Type: "text"},
	{ID: "330000", Name: "浙江省", Type: "text"},
	{ID: "340000", Name: "安徽省", Type: "text"},
	{ID: "350000", Name: "福建省", Type: "text"},
	{ID: "360000", Name: "江西省", Type: "text"},
	{ID: "370000", Name: "山东省", Type: "text"},
	{ID: "410000", Name: "河南省", Type: "text"},
	{ID: "420000", Name: "湖北省", Type: "text"},
	{ID: "430000", Name: "湖南省", Type: "text"},
	{ID: "440000", Name: "广东省", Type: "text"},
	{ID: "450000", Name: "广西壮族自治区", Type: "text"},
	{ID: "460000", Name: "海南省", Type: "text"},
	{ID: "500000", Name: "重庆市", Type: "text"},
	{ID: "510000", Name: "四川省", Type: "text"},
	{ID: "520000", Name: "贵州省", Type: "text"},
	{ID: "530000", Name: "云南省", Type: "text"},
	{ID: "540000", Name: "西藏自治区", Type: "text"},
	{ID: "610000", Name: "陕西省", Type: "text"},
	{ID: "620000", Name: "甘肃省", Type: "text"},
	{ID: "630000", Name: "青海省", Type: "text"},
	{ID: "640000", Name: "宁夏回族自治区", Type: "text"},
	{ID: "650000", Name: "新疆维吾尔自治区", Type: "text"},
	{ID: "710000", Name: "台湾省", Type: "text"},
	{ID: "810000", Name: "香港特别行政区", Type: "text"},
	{ID: "820000", Name: "澳门特别行政区", Type: "text"},
}

type Attribute struct {
	ID                int64       `field:"id" json:"id" bson:"id"`
	OwnerID           string      `field:"bk_supplier_account" json:"bk_supplier_account" bson:"bk_supplier_account"`
	ObjectID          string      `field:"bk_obj_id" json:"bk_obj_id" bson:"bk_obj_id"`
	PropertyID        string      `field:"bk_property_id" json:"bk_property_id" bson:"bk_property_id"`
	PropertyName      string      `field:"bk_property_name" json:"bk_property_name" bson:"bk_property_name"`
	PropertyGroup     string      `field:"bk_property_group" json:"bk_property_group" bson:"bk_property_group"`
	PropertyGroupName string      `field:"bk_property_group_name,ignoretomap" json:"bk_property_group_name" bson:"-"`
	PropertyIndex     int64       `field:"bk_property_index" json:"bk_property_index" bson:"bk_property_index"`
	Unit              string      `field:"unit" json:"unit" bson:"unit"`
	Placeholder       string      `field:"placeholder" json:"placeholder" bson:"placeholder"`
	IsEditable        bool        `field:"editable" json:"editable" bson:"editable"`
	IsPre             bool        `field:"ispre" json:"ispre" bson:"ispre"`
	IsRequired        bool        `field:"isrequired" json:"isrequired" bson:"isrequired"`
	IsReadOnly        bool        `field:"isreadonly" json:"isreadonly" bson:"isreadonly"`
	IsOnly            bool        `field:"isonly" json:"isonly" bson:"isonly"`
	IsSystem          bool        `field:"bk_issystem" json:"bk_issystem" bson:"bk_issystem"`
	IsAPI             bool        `field:"bk_isapi" json:"bk_isapi" bson:"bk_isapi"`
	PropertyType      string      `field:"bk_property_type" json:"bk_property_type" bson:"bk_property_type"`
	Option            interface{} `field:"option" json:"option" bson:"option"`
	Description       string      `field:"description" json:"description" bson:"description"`
	Creator           string      `field:"creator" json:"creator" bson:"creator"`
	CreateTime        *time.Time  `json:"create_time" bson:"create_time"`
	LastEditor        string      `json:"bk_last_editor" bson:"bk_last_editor"`
	LastTime          *time.Time  `json:"last_time" bson:"last_time"`
}
