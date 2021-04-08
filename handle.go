package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/mcoo/OPQBot"

	"github.com/icepie/xiaoice-beauty/model"
)

func buildResult(rte model.AnalyzeImgRte) string {
	score := int64(rte.Content.Metadata.Score)
	if score == 0 {

		Result := fmt.Sprintf(`
## 鉴定结果
	- 得分: %.1f
	- 简述: %s
		
		`, rte.Content.Metadata.Score, rte.Content.Text)

		return Result
	}

	Result := fmt.Sprintf(`
## 鉴定结果
	- 得分: %.1f
	- 简述: %s

## 人物信息
	- 性别: %s
	- 是否为名人: %s
	- 是否为表情包: %s
	- 面部像素区域: %s

## 评分细则
	- %s: %.1f
	- %s: %.1f
	- %s: %.1f
`, rte.Content.Metadata.Score, rte.Content.Text, rte.Content.Metadata.Gender, rte.Content.Metadata.Isceleb, rte.Content.Metadata.Isemoji, rte.Content.Metadata.FacePoints, rte.Content.Metadata.FbrKey0, rte.Content.Metadata.FbrScore0, rte.Content.Metadata.FbrKey1, rte.Content.Metadata.FbrScore1, rte.Content.Metadata.FbrKey2, rte.Content.Metadata.FbrScore2)
	return Result
}

// 好友消息处理
func friendMsgHandle(botQQ int64, packet OPQBot.FriendMsgPack) {
	log.Println(botQQ, packet)

	if packet.MsgType == "PicMsg" {
		// 解析消息
		fpc := FriendPicContent{}
		err := json.Unmarshal([]byte(packet.Content), &fpc)
		if err != nil {
			log.Println(err)
			return
		}

		if fpc.Content == nil {
			return
		}

		if strings.Contains(fpc.Content.(string), "颜") {

			for i := 0; i < len(fpc.Friendpic); i++ {
				Bot.Send(OPQBot.SendMsgPack{
					SendToType: OPQBot.SendToTypeFriend,
					ToUserUid:  packet.FromUin,
					Content:    OPQBot.SendTypeTextMsgContent{Content: "正在生成报告...请稍等"},
				})

				// 从图片链接分析
				rte, err := IB.AnalyzeImgByUrl(fpc.Friendpic[i].Url)
				if err != nil {
					log.Println("error: ", err)
					Bot.Send(OPQBot.SendMsgPack{
						SendToType: OPQBot.SendToTypeFriend,
						ToUserUid:  packet.FromUin,
						Content:    OPQBot.SendTypeTextMsgContent{Content: "分析失败啦～换张图试试吧"},
					})
					return
				}

				Bot.Send(OPQBot.SendMsgPack{
					SendToType: OPQBot.SendToTypeFriend,
					ToUserUid:  packet.FromUin,
					Content:    OPQBot.SendTypeTextMsgContent{Content: buildResult(rte)},
				})

				u, err := url.Parse(rte.Content.Metadata.Reportimgurl)
				if err != nil {
					log.Println("error: ", err)
					return
				}

				// 这玩意ssl证书有问题
				u.Scheme = "http"

				Bot.Send(OPQBot.SendMsgPack{
					SendToType: OPQBot.SendToTypeFriend,
					ToUserUid:  packet.FromUin,
					Content:    OPQBot.SendTypePicMsgByUrlContent{Content: "图片报告来了", PicUrl: u.String()},
				})
			}

		}

	}
}

// 群消息处理
func groupMsgHandle(botQQ int64, packet OPQBot.GroupMsgPack) {
	log.Println(botQQ, packet)

	if packet.MsgType == "PicMsg" {
		// 解析消息
		gpc := GroupPicContent{}
		err := json.Unmarshal([]byte(packet.Content), &gpc)
		if err != nil {
			return
		}

		if gpc.Content == nil {
			return
		}

		if strings.Contains(gpc.Content.(string), "颜") {

			for i := 0; i < len(gpc.GroupPic); i++ {
				Bot.Send(OPQBot.SendMsgPack{
					SendToType: OPQBot.SendToTypeGroup,
					ToUserUid:  packet.FromGroupID,
					Content:    OPQBot.SendTypeTextMsgContent{Content: "正在生成报告...请稍等"},
				})

				// 从图片链接分析
				rte, err := IB.AnalyzeImgByUrl(gpc.GroupPic[i].Url)
				if err != nil {
					log.Println("error: ", err)
					Bot.Send(OPQBot.SendMsgPack{
						SendToType: OPQBot.SendToTypeGroup,
						ToUserUid:  packet.FromGroupID,
						Content:    OPQBot.SendTypeTextMsgContent{Content: "分析失败啦～换张图试试吧"},
					})
					return
				}

				Bot.Send(OPQBot.SendMsgPack{
					SendToType: OPQBot.SendToTypeGroup,
					ToUserUid:  packet.FromGroupID,
					Content:    OPQBot.SendTypeTextMsgContent{Content: buildResult(rte)},
				})

				u, err := url.Parse(rte.Content.Metadata.Reportimgurl)
				if err != nil {
					log.Println("error: ", err)
					return
				}

				// 这玩意ssl证书有问题
				u.Scheme = "http"

				Bot.Send(OPQBot.SendMsgPack{
					SendToType: OPQBot.SendToTypeGroup,
					ToUserUid:  packet.FromGroupID,
					Content:    OPQBot.SendTypePicMsgByUrlContent{Content: "图片报告来了", PicUrl: u.String()},
				})

			}

		}

	}

}
