package bike

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBattleMsg(t *testing.T) {
	ast := assert.New(t)
	msg := `91A14694 :参戦ID
参加者募集！
Lv10 ユグドラシル・マグナ`
	ast.True(IsGBFBattle(msg))
	msg = `2F1E6FF1 :参戦ID
参加者募集！
黄龍・黒麒麟HL`
	ast.True(IsGBFBattle(msg))
}

func TestBattleMsgFail(t *testing.T) {
	ast := assert.New(t)
	msg := `参加者募集！参戦I：37F7B348
ALv60 リヴァイアサン・マグナ
https://t.co/RbqZBFIUBz`
	ast.False(IsGBFBattle(msg))
}

func TestConvertBattleMsg(t *testing.T) {
	ast := assert.New(t)
	msg := `91A14694 :参戦ID
参加者募集！
Lv100 ユグドラシル・マグナ`
	result, err := ConvertGBFBattleInfo(msg)
	if err != nil {
		ast.Fail(err.Error())
		return
	}
	ast.Equal("91A14694", result.RoomId)
	ast.Equal("100", result.Level)
	ast.Equal("ユグドラシル・マグナ", result.MobName)
	msg = `2F1E6FF1 :参戦ID
参加者募集！
黄龍・黒麒麟HL`
	result, err = ConvertGBFBattleInfo(msg)
	if err != nil {
		ast.Fail(err.Error())
		return
	}
	ast.Equal("2F1E6FF1", result.RoomId)
	ast.Equal("", result.Level)
	ast.Equal("黄龍・黒麒麟HL", result.MobName)
}
