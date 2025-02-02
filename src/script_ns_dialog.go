package opennox

import (
	"github.com/noxworld-dev/noxscript/ns/v4"
	"github.com/noxworld-dev/noxscript/ns/v4/audio"
	"github.com/noxworld-dev/opennox-lib/strman"

	"github.com/noxworld-dev/opennox/v1/common/sound"
	"github.com/noxworld-dev/opennox/v1/server"
)

func (s noxScriptNS) SetShopkeeperText(obj ns.Obj, text ns.StringID) {
	//TODO implement me
	panic("implement me")
}

func (s noxScriptNS) SetShopkeeperTextStr(obj ns.Obj, text string) {
	//TODO implement me
	panic("implement me")
}

func (s noxScriptNS) SetDialog(cobj ns.Obj, typ ns.DialogType, start, end ns.Func) {
	if cobj == nil {
		return
	}
	flags := server.ParseDialogFlags(string(typ))
	obj := cobj.(nsObj)
	starti := s.s.noxScript.AsFuncIndex(start)
	endi := s.s.noxScript.AsFuncIndex(end)
	nox_xxx_scriptSetDialog_548C80(obj.SObj(), flags, starti, endi)
}

func (s noxScriptNS) CancelDialog(cobj ns.Obj) {
	if cobj == nil {
		return
	}
	obj := cobj.(nsObj)
	scriptCancelDialog(obj.SObj())
}

func (s noxScriptNS) StoryPic(cobj ns.Obj, name string) {
	if cobj == nil {
		return
	}
	obj := cobj.(nsObj)
	obj.SObj().SetDialogPortrait(name)
}

func (s noxScriptNS) TellStory(aud audio.Name, story ns.StringID) {
	caller := s.s.noxScript.Caller()
	trigger := s.s.noxScript.Trigger()
	snd := sound.ByName(string(aud))
	nox_xxx_startShopDialog_548DE0(caller, trigger, snd, strman.ID(story))
}

func (s noxScriptNS) TellStoryStr(audio audio.Name, story string) {
	//TODO implement me
	panic("implement me")
}

func (s noxScriptNS) StartDialog(cobj ns.Obj, cother ns.Obj) {
	if cobj == nil || cother == nil {
		return
	}
	obj := cobj.(nsObj)
	other := cother.(nsObj)
	nox_xxx_script_forcedialog_548CD0(other.SObj(), obj.SObj())
}

func (s noxScriptNS) GetAnswer(cobj ns.Obj) ns.DialogAnswer {
	if cobj == nil {
		return 0
	}
	obj := cobj.(nsObj)
	return ns.DialogAnswer(sub_548F40(obj.SObj()))
}
