package bike

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var (
	battleReg    = regexp.MustCompile(`\w{8} :参戦ID\n参加者募集！\n`)
)

type BattleInfo struct {
	Id       int64  `json:"id"`
	Level    string `json:"level"`
	RoomId   string `json:"roomId"`
	MobName  string `json:"mobName"`
	Creator  string `json:"creator"`
	CreateAt string `json:"createAt"`
}

func IsGBFBattle(msg string) bool {
	return battleReg.MatchString(msg)
}

func ConvertGBFBattleInfo(msg string) (*BattleInfo, error) {
	if msg == "" {
		return nil, errors.New("Msg is empty")
	}
	if !IsGBFBattle(msg) {
		return nil, fmt.Errorf("%s is not GBF Battle", msg)
	}
	msg = strings.Replace(msg, ":参戦ID", "", -1)
	msg = strings.Replace(msg, "参加者募集！", "", -1)
	msg = strings.Replace(msg, "\n", " ", -1)
	msg = strings.Replace(msg, "Lv", " ", -1)
	strs := strings.Split(msg, " ")
	strs = delete_empty(strs)
	if len(strs) > 3 {
		return nil, fmt.Errorf("%s result not match. %#v", msg, strs)
	}
	roomId := strs[0]
	level := ""
	mobName := ""
	if len(strs) == 3 {
		level = strs[1]
		mobName = strs[2]
	} else {
		mobName = strs[1]
	}
	ret := &BattleInfo{
		RoomId:  roomId,
		Level:   level,
		MobName: mobName,
	}
	return ret, nil
}

func delete_empty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}
