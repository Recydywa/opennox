package opennox

import (
	"fmt"
	"math"
	"unsafe"

	"github.com/noxworld-dev/opennox-lib/object"
	"github.com/noxworld-dev/opennox-lib/spell"
	"github.com/noxworld-dev/opennox-lib/types"

	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/common/sound"
	"github.com/noxworld-dev/opennox/v1/common/unit/ai"
	"github.com/noxworld-dev/opennox/v1/legacy"
	"github.com/noxworld-dev/opennox/v1/legacy/common/alloc"
	"github.com/noxworld-dev/opennox/v1/server"
)

func init() {
	server.RegisterAIAction(AIActionIdle{})
	server.RegisterAIAction(AIActionWait{})
	server.RegisterAIAction(AIActionWaitRel{})
	server.RegisterAIAction(AIActionCastOnObj{})
	server.RegisterAIAction(AIActionCastOnPos{})
	server.RegisterAIAction(AIActionCastDuration{})
	server.RegisterAIAction(AIActionDropObj{})
	server.RegisterAIAction(AIActionFindObj{})
	server.RegisterAIAction(AIActionMorphIntoChest{})
	server.RegisterAIAction(AIActionMorphBackToSelf{})
	server.RegisterAIAction(AIActionReport{})
}

type aiData struct {
	s                  *Server
	allocListen        alloc.ClassT[MonsterListen]
	listenHead         *MonsterListen
	soundMuteThreshold int
	lastHeard          types.Pointf
	stackChanged       bool
}

func (a *aiData) Init(s *Server) {
	a.s = s
}

func nox_ai_debug_print(str string) {
	ai.Log.Printf("%s", str)
}

func aiStackSetArgs(s *server.AIStackItem, args ...any) {
	if s == nil {
		return
	}
	off := 0
	for _, v := range args {
		switch v := v.(type) {
		case int:
			s.Args[off] = uintptr(uint32(int32(v)))
		case uint:
			s.Args[off] = uintptr(uint32(v))
		case int32:
			s.Args[off] = uintptr(uint32(v))
		case uint32:
			s.Args[off] = uintptr(v)
		case unsafe.Pointer:
			s.Args[off] = uintptr(v)
		case float32:
			s.Args[off] = uintptr(math.Float32bits(v))
		case server.Obj:
			s.Args[off] = uintptr(unsafe.Pointer(toObjectS(v)))
		case types.Pointf:
			s.Args[off+0] = uintptr(math.Float32bits(v.X))
			off++
			s.Args[off+1] = uintptr(math.Float32bits(v.Y))
		default:
			panic(fmt.Errorf("unsupported arg: %T", v))
		}
		off++
	}
}

func aiStackArgObj(s *server.AIStackItem, i int) *Object {
	return asObject(unsafe.Pointer(s.Args[i]))
}

func aiStackArgUnit(s *server.AIStackItem, i int) *Object {
	return asObject(unsafe.Pointer(s.Args[i]))
}

func (obj *Object) maybePrintAIStack(event string) {
	if noxflags.HasEngine(noxflags.EngineShowAI) {
		obj.UpdateDataMonster().PrintAIStack(obj.getServer().Frame(), event)
	}
}

func (a *aiData) nox_xxx_mobActionDependency(u *Object) {
	ud := u.UpdateDataMonster()
	stack := ud.GetAIStack()
	if len(stack) == 0 {
		return
	}
	isOR := false
	okPrev := false
	for i := len(stack) - 1; i >= 0; i-- {
		ok := true
		st := &stack[i]
		typ := st.Type()
		if !typ.IsCondition() {
			isOR = false
			continue
		}
		switch typ {
		case ai.DEPENDENCY_OR:
			isOR = true
			okPrev = false
			continue
		case ai.DEPENDENCY_TIME:
			ok = st.ArgU32(0) > a.s.Frame()
		case ai.DEPENDENCY_ALIVE:
			obj := aiStackArgObj(st, 0)
			h := obj.HealthData
			if obj == nil || !obj.Class().HasAny(object.MaskUnits) || (h.Cur == 0) && h.Max != 0 {
				ok = false
				ud.Field97 = 0
				ud.Field101 = a.s.Frame() + a.s.TickRate()
			}
		case ai.DEPENDENCY_UNDER_ATTACK:
			if sub_5347A0(u) {
				if u.Obj130 != nil {
					v26 := getOwnerUnit(u.Obj130)
					if v26 != nil && v26.Class().HasAny(object.MaskUnits) {
						st.Args[0] = uintptr(a.s.Frame())
					}
				}
			}
			ok = a.s.Frame()-st.ArgU32(0) <= 10*a.s.TickRate()
		case ai.DEPENDENCY_NOT_UNDER_ATTACK:
			if !sub_5347A0(u) {
				break
			}
			if u.Obj130 == nil {
				break
			}
			attacker := getOwnerUnit(u.Obj130)
			if attacker == nil {
				break
			}
			if attacker.Class().HasAny(object.MaskUnits) {
				ok = false
			}
		case ai.DEPENDENCY_CAN_SEE:
			obj := aiStackArgObj(st, 0)
			if obj == nil || !nox_xxx_unitCanInteractWith_5370E0(u, obj, 0) {
				ok = false
			}
		case ai.DEPENDENCY_CANNOT_SEE:
			obj := aiStackArgObj(st, 0)
			if obj == nil || legacy.Sub_533360(u.SObj(), obj.SObj()) != 0 {
				ok = false
			}
		case ai.DEPENDENCY_BLOCKED_LINE_OF_FIRE:
			obj := aiStackArgObj(st, 0)
			if obj == nil {
				ok = false
				break
			}
			pos, pos2 := u.Pos(), obj.Pos()
			ok = !a.s.MapTraceObstacles(u.SObj(), pos, pos2)
		case ai.DEPENDENCY_OBJECT_AT_VISIBLE_LOCATION:
			v28 := aiStackArgObj(st, 2)
			v29 := false
			if v28 != nil && nox_xxx_unitCanInteractWith_5370E0(u, v28, 0) {
				v29 = true
				pos := v28.Pos()
				st.Args[0] = uintptr(math.Float32bits(pos.X))
				st.Args[1] = uintptr(math.Float32bits(pos.Y))
			}
			if a.s.MapTraceRay(u.Pos(), st.ArgPos(0), server.MapTraceFlag1); !v29 {
				ok = false
			}
		case ai.DEPENDENCY_OBJECT_FARTHER_THAN:
			obj, r := aiStackArgObj(st, 2), st.ArgF32(0)
			ok = obj != nil && nox_xxx_calcDistance_4E6C00(u, obj) > r
		case ai.DEPENDENCY_OBJECT_CLOSER_THAN:
			obj, r := aiStackArgObj(st, 2), st.ArgF32(0)
			ok = obj != nil && nox_xxx_calcDistance_4E6C00(u, obj) <= r
		case ai.DEPENDENCY_LOCATION_FARTHER_THAN:
			pos := u.Pos()
			pos2 := st.ArgPos(2)
			dx := pos2.X - pos.X
			dy := pos2.Y - pos.Y
			ok = float32(math.Sqrt(float64(dy*dy+dx*dx))) >= st.ArgF32(0)
		case ai.DEPENDENCY_LOCATION_CLOSER_THAN:
			pos := u.Pos()
			pos2 := st.ArgPos(2)
			dx := pos2.X - pos.X
			dy := pos2.Y - pos.Y
			ok = float32(math.Sqrt(float64(dy*dy+dx*dx))) <= st.ArgF32(0)
		case ai.DEPENDENCY_VISIBLE_ENEMY:
			ok = ud.CurrentEnemy != nil
		case ai.DEPENDENCY_NO_VISIBLE_ENEMY:
			ok = ud.CurrentEnemy == nil
		case ai.DEPENDENCY_NO_VISIBLE_FOOD:
			r := 250.0
			if noxflags.HasGame(noxflags.GameModeQuest) {
				r = 640.0
			}
			ok = legacy.Nox_xxx_mobSearchEdible_544A00(u.SObj(), float32(r)) == 0
		case ai.DEPENDENCY_NO_INTERESTING_SOUND:
			if ud.Field97 != 0 && a.s.Frame()-ud.Field101 < 3*a.s.TickRate() {
				ok = false
			}
		case ai.DEPENDENCY_NO_NEW_ENEMY:
			old := aiStackArgUnit(st, 0)
			if old == nil {
				ok = false
				break
			}
			enemy := asObjectS(ud.CurrentEnemy)
			if enemy == nil {
				break
			}
			pos := u.Pos()

			pos2 := old.Pos()
			dx1 := pos2.X - pos.X
			dy1 := pos2.Y - pos.Y
			rOld := dx1*dx1 + dy1*dy1

			pos3 := enemy.Pos()
			dx2 := pos3.X - pos.X
			dy2 := pos3.Y - pos.Y
			rNew := dx2*dx2 + dy2*dy2

			ok = rNew >= rOld
		case ai.DEPENDENCY_UNINTERRUPTABLE:
			return
		case ai.DEPENDENCY_IS_ENCHANTED:
			ok = u.HasEnchant(server.EnchantID(st.Args[0]))
		case ai.DEPENDENCY_ENEMY_CLOSER_THAN:
			enemy := asObjectS(ud.CurrentEnemy)
			ok = enemy != nil && nox_xxx_calcDistance_4E6C00(u, enemy) <= st.ArgF32(0)
		case ai.DEPENDENCY_NOT_HEALTHY:
			h := u.HealthData
			perc := float32(1.0)
			if h.Max != 0 {
				perc = float32(h.Cur) / float32(h.Max)
			}
			if perc >= ud.ResumeLevel {
				ok = false
			}
		case ai.DEPENDENCY_WAIT_FOR_STAMINA:
			if int(ud.Field282_0) >= legacy.Nox_xxx_weaponGetStaminaByType_4F7E80(int(ud.Field514)) {
				ok = false
			}
		case ai.DEPENDENCY_ENEMY_FARTHER_THAN:
			enemy := asObjectS(ud.CurrentEnemy)
			if enemy != nil && nox_xxx_calcDistance_4E6C00(u, enemy) < st.ArgF32(0) {
				ok = false
			}
		case ai.DEPENDENCY_UNDER_CURSOR:
			hu := legacy.HostPlayerUnit()
			if hu == nil || legacy.Nox_xxx_findObjectAtCursor_54AF40(hu.SObj()) != u.SObj() {
				ok = false
			}
		case ai.DEPENDENCY_NOT_CORNERED:
			ok = legacy.Nox_xxx_mobGetMoveAttemptTime_534810(u.SObj()) == 0
		case ai.DEPENDENCY_LOCATION_IS_SAFE:
			legacy.Set_dword_5d4594_2489460(1)
			a.s.Map.EachObjInCircle(st.ArgPos(0), 50.0, func(it *server.Object) bool {
				legacy.Nox_xxx_unitIsDangerous_547120(asObjectS(it).SObj(), u.SObj())
				return true
			})
			if legacy.Get_dword_5d4594_2489460() == 0 {
				ok = false
			}
		case ai.DEPENDENCY_NOT_FRUSTRATED:
			if ud.Flags360&0x200000 != 0 {
				ok = false
				ud.Flags360 &= 0xFFDFFFFF
			}
		case ai.DEPENDENCY_NOT_MOVED:
			ok = u.Pos() == u.PrevPos
		}
		if !isOR {
			if ok {
				continue
			}
			// failed, pop actions
		} else {
			if ok {
				okPrev = true
			}
			if i > 0 {
				if stack[i-1].Type().IsCondition() {
					continue
				}
			}
			if okPrev {
				continue
			}
		}
		if noxflags.HasEngine(noxflags.EngineShowAI) {
			ai.Log.Printf("%d: %s DEPENDENCY '%s'@%d failed, popping:\n", a.s.Frame(), u.String(), typ.String(), i)
		}
		for {
			u.monsterPopAction()
			if !(int(ud.AIStackInd) >= i && legacy.Sub_5341F0(u.SObj()) == 0) {
				break
			}
		}
		stack = ud.GetAIStack()
		i = len(stack) - 1
		u.monsterActionReset()
		u.maybePrintAIStack("procDep")
	}
}

func nox_xxx_checkMobAction_50A0D0(u *Object, act ai.ActionType) bool {
	return legacy.Nox_xxx_checkMobAction_50A0D0(u.SObj(), act) != 0
}

func sub_545E60(a1c *server.Object) int {
	u := asObjectS(a1c)
	s := u.getServer()

	ud := u.UpdateDataMonster()
	ts := u.Field134
	if ud.Field129 >= ts || s.Frame()-ts >= 10*s.TickRate() {
		return 0
	}
	ud.Field129 = ts
	if u.Obj130 != nil {
		if obj4 := getOwnerUnit(u.Obj130); obj4 != nil {
			if !u.isEnemyTo(obj4) {
				return 0
			}
			canInteract := nox_xxx_unitCanInteractWith_5370E0(u, obj4, 0)
			if u.isPlant() {
				if !canInteract {
					return 0
				}
				u.monsterPushAction(ai.DEPENDENCY_ENEMY_CLOSER_THAN, float32(ud.Field328)*1.05)
			} else {
				u.monsterPushAction(ai.DEPENDENCY_UNDER_ATTACK, s.Frame())
			}
			u.monsterPushAction(ai.ACTION_FIGHT, obj4.Pos(), s.Frame())
			if !canInteract {
				u.monsterPushAction(ai.DEPENDENCY_NO_VISIBLE_ENEMY)
				if nox_xxx_monsterCanAttackAtWill_534390(u) {
					u.monsterPushAction(ai.DEPENDENCY_NO_INTERESTING_SOUND)
				}
				u.monsterPushAction(ai.ACTION_MOVE_TO, obj4.Pos(), 0)
			}
			return 1
		}
	}
	if !nox_xxx_checkMobAction_50A0D0(u, ai.ACTION_ROAM) {
		u.monsterPushAction(ai.DEPENDENCY_TIME, 5*s.TickRate())
		u.monsterPushAction(ai.DEPENDENCY_NO_VISIBLE_ENEMY)
		if nox_xxx_monsterCanAttackAtWill_534390(u) {
			u.monsterPushAction(ai.DEPENDENCY_NO_INTERESTING_SOUND)
		}
		u.monsterPushAction(ai.ACTION_ROAM, 0, 0, -128)
	}
	return 0
}

var _ = [1]struct{}{}[24-unsafe.Sizeof(MonsterListen{})]

type MonsterListen struct {
	snd   sound.ID       // 0, 0
	obj   *server.Object // 1, 4
	pos   types.Pointf   // 2, 8
	frame uint32         // 4, 16
	next  *MonsterListen // 5, 20
}

func sub_50D1C0() {
	noxServer.ai.Reset()
}

func (a *aiData) Reset() {
	a.allocListen.FreeAllObjects()
	//allocMonsterListen.Class = nil
	a.listenHead = nil
}

func (a *aiData) Free() {
	if a.allocListen.Class != nil {
		a.allocListen.Free()
	}
	a.listenHead = nil
}

func (a *aiData) NewSound(snd sound.ID, obj *server.Object, pos types.Pointf) {
	if a.s.Audio.Flags(snd) == 0 {
		return
	}
	if a.allocListen.Class == nil {
		a.allocListen = alloc.NewClassT("MonsterListen", MonsterListen{}, 128)
	}
	if getOwnerUnit(obj) == nil {
		return
	}

	p := a.allocListen.NewObject()
	if p == nil {
		return
	}
	*p = MonsterListen{
		snd:   snd,
		obj:   obj,
		pos:   pos,
		frame: a.s.Frame(),
	}
	p.next = a.listenHead
	a.listenHead = p
}

func getOwnerUnit(obj *server.Object) *server.Object {
	for it := obj; it != nil; it = it.ObjOwner {
		if it.Class().HasAny(object.MaskUnits) {
			return it
		}
	}
	return nil
}

func (a *aiData) aiListenToSounds(u *Object) {
	// Not sure about this check. If unit is invulnerable, don't listen for anything?
	// Is present in vanilla though, so this might be some kind of weird fix for strange behaviour
	if legacy.Nox_xxx_checkIsKillable_528190(u.SObj()) == 0 {
		return
	}

	// Do not listen to anything if you are a fish, frog or rat
	if u.isFish() || u.isFrog() || u.isRat() {
		return
	}

	var (
		prev     *MonsterListen
		next     *MonsterListen
		maxHeard *MonsterListen
		maxDist  int
	)
	ud := u.UpdateDataMonster()
	for it := a.listenHead; it != nil; it = next {
		next = it.next
		if a.s.Frame() < it.frame || a.s.Frame()-it.frame > 2 {
			if prev != nil {
				prev.next = next
			} else {
				a.listenHead = it.next
			}
			a.allocListen.FreeObjectFirst(it)
		} else {
			if it.obj != nil && it.obj.Flags().Has(object.FlagDestroyed) {
				it.obj = nil
			}
			if ud.Field101 <= it.frame && a.shouldUnitListen(u, it) {
				dist := a.traceSound(u, it)
				// This finds the farthest?
				if dist > 0 && dist > maxDist {
					maxDist = dist
					maxHeard = it
				}
			}
			prev = it
		}
	}
	if maxHeard != nil && (maxHeard.frame > ud.Field101 || maxDist > int(ud.Field102)) {
		a.nox_xxx_unitEmitHearEvent_50D110(u, maxHeard, maxDist)
	}
}

func (a *aiData) traceSound(u *Object, p *MonsterListen) int {
	flags := a.s.Audio.Flags(p.snd)
	perc := a.soundFadePerc(p.snd, p.pos, u.Pos())
	if !a.checkSoundThreshold(flags, perc) {
		return -1
	}
	if !a.s.MapTraceRayAt(u.Pos(), p.pos, nil, nil, 5) {
		perc = int(float64(perc) * 0.5)
	}
	if !a.checkSoundThreshold(flags, perc) {
		return -1
	}
	return perc
}

func nox_xxx_gameSetAudioFadeoutMb_501AC0(v int) {
	noxServer.ai.nox_xxx_gameSetAudioFadeoutMb(v)
}

func (a *aiData) nox_xxx_gameSetAudioFadeoutMb(v int) {
	v = clamp(v, 0, 100)
	a.soundMuteThreshold = v
}

func (a *aiData) soundFadePerc(snd sound.ID, p1, p2 types.Pointf) int {
	max := a.s.Audio.MaxDist(snd)
	if max <= 0 {
		return 0
	}
	dx := float64(p1.X - p2.X)
	dy := float64(p1.Y - p2.Y)
	if abs(dx) >= float64(max) || abs(dy) >= float64(max) {
		return 0
	}
	dist := int(math.Sqrt(dy*dy + dx*dx + 0.1))
	if dist >= max {
		return 0
	}
	v := 100 * (max - dist) / max
	v = clamp(v, 0, 100)
	if v <= a.soundMuteThreshold {
		return 0
	}
	return v
}

func (a *aiData) checkSoundThreshold(flags, perc int) bool {
	threshold := 50
	if flags&0x20 != 0 { // walk?
		threshold = 89
	} else if flags&0x40 != 0 { // run?
		threshold = 20
	}
	return perc >= threshold
}

func (a *aiData) shouldUnitListen(u *Object, lis *MonsterListen) bool {
	ud := u.UpdateDataMonster()
	punit := lis.obj.FindOwnerChainPlayer()
	flags := a.s.Audio.Flags(lis.snd)
	if ud.Field101 > a.s.Frame() {
		return false
	}
	if punit == nil {
		if flags == 0 {
			return false
		}
	} else {
		if u.isEnemyTo(punit) {
			if flags == 0 {
				return false
			}
		} else if flags&0x10 == 0 {
			return false
		}
	}
	if v12 := legacy.Sub_501C00(lis.pos, lis.obj); v12 != 0 {
		pos := u.Pos()
		cpi, free := alloc.Make([]int32{}, 2)
		defer free()
		cpi[0] = int32(pos.X)
		cpi[1] = int32(pos.Y)
		resp := legacy.Nox_xxx_polygonIsPlayerInPolygon_4217B0(unsafe.Pointer(&cpi[0]), 0)
		if resp != nil && byte(v12) != *(*byte)(unsafe.Add(unsafe.Pointer(&resp.Field_0[32]), 2)) {
			return false
		}
	}
	return true
}

func (a *aiData) nox_xxx_unitEmitHearEvent_50D110(u *Object, lis *MonsterListen, dist int) {
	ud := u.UpdateDataMonster()
	ud.Field97 = uint32(lis.snd)
	ud.Field101 = a.s.Frame()
	ud.Field102 = uint32(dist)
	if lis.obj != nil {
		ud.Field98 = lis.obj.NetCode
	} else {
		ud.Field98 = 0
	}
	ud.Field99X = lis.pos.X
	ud.Field99Y = lis.pos.Y
	a.lastHeard = lis.pos
	obj5 := lis.obj.FindOwnerChainPlayer()
	a.s.noxScript.ScriptCallback(&ud.ScriptHearEnemy, obj5, u.SObj(), server.NoxEventMonsterHearEnemy)
}

func (a *aiData) lastHeardEvent() types.Pointf {
	return a.lastHeard
}

func nox_xxx_monsterPopAction_50A160(a1 *server.Object) int {
	return asObjectS(a1).monsterPopAction()
}

func (obj *Object) monsterPopAction() int {
	s := obj.getServer()
	ud := obj.UpdateDataMonster()
	if noxflags.HasEngine(noxflags.EngineShowAI) {
		typ := ud.AIStackHead().Type()
		ai.Log.Printf("%d: PopActionStack( %s(#%d) ) = %s@%d:\n", s.Frame(), obj, obj.NetCode, typ, ud.AIStackInd)
	}
	if cur := ud.AIStackHead(); cur != nil {
		if act := cur.Type(); !act.IsCondition() && cur.Field5 != 0 {
			if a := server.GetAIAction(act); a != nil {
				a.End(obj.SObj())
			}
		}
	}
	s.ai.stackChanged = true
	// pop action
	ud.AIStackInd--
	// pop related conditions (if any)
	for ; ud.AIStackInd >= 0; ud.AIStackInd-- {
		cur := &ud.AIStack[ud.AIStackInd]
		if !cur.Type().IsCondition() {
			break
		}
	}
	obj.monsterActionReset()
	si := ud.AIStackInd
	if si < 0 {
		ud.AIStackInd = 0
		ud.AIStack[0].Action = 0
	}
	return int(si)
}

func nox_xxx_monsterPushAction_50A260_impl(u *server.Object, act int, file string, line int) unsafe.Pointer {
	return unsafe.Pointer(asObjectS(u).monsterPushActionImpl(ai.ActionType(act), file, line))
}

func (obj *Object) monsterPushActionImpl(act ai.ActionType, file string, line int) *server.AIStackItem {
	if !obj.Class().Has(object.ClassMonster) {
		return nil
	}
	s := obj.getServer()
	ud := obj.UpdateDataMonster()
	if int(ud.AIStackInd) >= len(ud.AIStack)-1 {
		return nil
	}
	if cur := ud.AIStackHead(); cur == nil {
		ud.AIStackInd = -1
	} else {
		if ud.AIStackHead().Type() == ai.ACTION_DEAD && act != ai.ACTION_DYING {
			return nil
		}
		if curAct := cur.Type(); curAct == ai.ACTION_IDLE && ud.AIStackInd == 0 {
			ud.AIStackInd = -1
		} else if !curAct.IsCondition() && cur.Field5 != 0 {
			if a := server.GetAIAction(curAct); a != nil {
				a.Cancel(obj.SObj())
			}
		}
	}
	ud.AIStackInd++
	ud.AIStack[ud.AIStackInd] = server.AIStackItem{
		Action: uint32(act), Field5: 0,
	}
	obj.monsterActionReset()
	if noxflags.HasEngine(noxflags.EngineShowAI) {
		ai.Log.Printf("%d: PushActionStack( %s(#%d), %s ), result: (%s:%d)\n", s.Frame(), obj, obj.NetCode, act, file, line)
	}
	s.ai.stackChanged = true
	return ud.AIStackHead()
}

func (obj *Object) monsterActionReset() {
	ud := obj.UpdateDataMonster()
	ud.Field2 = 0
	ud.Field67 = 0
	ud.Field74 = 0
	ud.Field91 = 0
	ud.Field120_1 = 0
	ud.Field120_2 = 0
	ud.Field120_3 = 0
	ud.Field124 = noxServer.Frame()
	ud.Field137 = noxServer.Frame()
}

func nox_xxx_unitUpdateMonster_50A5C0(a1 *server.Object) {
	u := asObjectS(a1)
	s := u.getServer()
	ud := u.UpdateDataMonster()

	if ud.Field523_2 != 0 {
		ud.Field523_2--
	}

	if obj4 := ud.Field548; obj4 != nil && obj4.Flags().HasAny(object.FlagDead|object.FlagDestroyed) {
		ud.Field548 = nil
	}
	legacy.Nox_xxx_mobAction_50A910(u.SObj())
	s.ai.aiListenToSounds(u)
	if !u.Flags().Has(object.FlagEnabled) {
		return
	}
	s.ai.stackChanged = false
	if ud.Field121 == nil {
		return
	}
	if !u.Flags().HasAny(object.FlagDead | object.FlagDestroyed) {
		if ud.Flags360&0x200 != 0 {
			if v7 := legacy.Nox_xxx_monsterGetSoundSet_424300(u.SObj()); v7 != nil {
				s.AudioEventObj(sound.ID(*(*uint32)(unsafe.Add(v7, 64))), u, 0, 0)
			}
			s.noxScript.ScriptCallback(&ud.ScriptIsHit, u.Obj130, u.SObj(), server.NoxEventMonsterIsHit)
			if noxflags.HasEngine(noxflags.EngineShowAI) {
				cur, max := u.Health()
				ai.Log.Printf("%d: HP = %d/%d\n", s.Frame(), cur, max)
			}
		}
		if v8 := ud.Field130; v8 != 0 && int(s.Frame()-v8) >= int(s.TickRate()) {
			legacy.Nox_xxx_monsterPlayHurtSound_532800(u.SObj())
			ud.Field130 = 0
		}
		legacy.Nox_xxx_mobAction_5469B0(u.SObj())
	}

	if h := u.HealthData; h != nil {
		if !u.Flags().Has(object.FlagDead) && int(s.Frame()-u.Field134) > int(s.TickRate()) {
			if h.Cur < h.Max && h.Max != 0 && s.Frame()%(180*s.TickRate()/uint32(h.Max)) == 0 {
				legacy.Nox_xxx_unitAdjustHP_4EE460(u.SObj(), 1)
			}
		}
	}
	legacy.Nox_xxx_unitUpdateSightMB_5281F0(u.SObj())
	legacy.Nox_xxx_monsterMainAIFn_547210(u.SObj())
	s.ai.nox_xxx_mobActionDependency(u)
	legacy.Nox_xxx_updateNPCAnimData_50A850(u.SObj())
	changedPrev := s.ai.stackChanged
	curAct := ai.ACTION_IDLE
	for {
		cur := ud.AIStackHead()
		act := cur.Type()
		curAct = act
		if cur.Field5 != 0 {
			if !s.ai.stackChanged {
				s.ai.stackChanged = changedPrev
			}
			break
		}
		cur.Field5 = 1
		s.ai.stackChanged = false
		if a := server.GetAIAction(act); a != nil {
			a.Start(u.SObj())
		}
		if !s.ai.stackChanged {
			s.ai.stackChanged = changedPrev
			break
		}
	}
	if a := server.GetAIAction(curAct); a != nil {
		a.Update(u.SObj())
	}
	if s.ai.stackChanged {
		u.maybePrintAIStack("stack changed")
	}
	ud.Flags360 &= 0xFFFFFDFF
	legacy.Nox_xxx_monsterPolygonEnter_421FF0(u.SObj())

	if v := ud.Field282_0; v < 100 {
		ud.Field282_0 = v + uint8(100/s.TickRate())
	}
	if u.isMimic() {
		legacy.Nox_xxx_monsterMimicCheckMorph_534950(u.SObj())
	}
	if s.Frame()-u.Field134 > 3*s.TickRate() {
		ud.Flags360 &= 0xFFF7FFFF
	}
}

type AIActionDropObj struct{}

func (AIActionDropObj) Type() ai.ActionType {
	return ai.ACTION_DROP_OBJECT
}
func (AIActionDropObj) Start(_ *server.Object)  {}
func (AIActionDropObj) Update(_ *server.Object) {}
func (AIActionDropObj) End(_ *server.Object)    {}
func (AIActionDropObj) Cancel(_ *server.Object) {}

type AIActionFindObj struct{}

func (AIActionFindObj) Type() ai.ActionType {
	return ai.ACTION_FIND_OBJECT
}
func (AIActionFindObj) Start(_ *server.Object)  {}
func (AIActionFindObj) Update(_ *server.Object) {}
func (AIActionFindObj) End(_ *server.Object)    {}
func (AIActionFindObj) Cancel(_ *server.Object) {}

type AIActionIdle struct{}

func (AIActionIdle) Type() ai.ActionType {
	return ai.ACTION_IDLE
}
func (AIActionIdle) Start(obj *server.Object) {
	u := asObjectS(obj)
	s := u.getServer()
	ud := u.UpdateDataMonster()
	cur := ud.AIStackHead()
	cur.Args[0] = uintptr(s.Frame())
}
func (AIActionIdle) End(_ *server.Object)    {}
func (AIActionIdle) Cancel(_ *server.Object) {}

func sub_5343C0(u *Object) bool {
	ud := u.UpdateDataMonster()
	return ud.Aggression < 0.66000003 && ud.Aggression > 0.33000001
}

func nox_xxx_monsterCanAttackAtWill_534390(u *Object) bool {
	return u.UpdateDataMonster().Aggression > 0.66000003
}

func sub_534440(u *Object) bool {
	return u.UpdateDataMonster().Aggression < 0.079999998
}

func sub_5347A0(u *Object) bool {
	return (u.UpdateDataMonster().Flags360>>9)&1 != 0
}

func nox_xxx_monsterLookAtDamager_5466B0(u *Object) bool {
	if !sub_5347A0(u) {
		return false
	}
	u.monsterPushAction(ai.ACTION_FACE_LOCATION, u.Pos132)
	return true
}

func (AIActionIdle) Update(obj *server.Object) {
	u := asObjectS(obj)
	s := u.getServer()
	ud := u.UpdateDataMonster()
	if uint16(s.Frame()-ud.AIStackHead().ArgU32(0)) == uint16(ud.Field305) {
		s.noxScript.ScriptCallback(&ud.ScriptLookingForEnemy, nil, u.SObj(), server.NoxEventMonsterIdle)
	}
	if u.Flags().Has(object.FlagEnabled) && (sub_5343C0(u) || nox_xxx_monsterCanAttackAtWill_534390(u)) {
		if enemy := asObjectS(ud.CurrentEnemy); enemy != nil {
			u.monsterPushAction(ai.ACTION_FIGHT, enemy.Pos(), s.Frame())
			return
		}
		if legacy.Sub_5466F0(u.SObj()) != 0 {
			return
		}
	}
	if !u.Flags().Has(object.FlagEnabled) || sub_534440(u) || sub_545E60(u.SObj()) == 0 {
		if u.isMimic() {
			if !u.HasEnchant(server.ENCHANT_ANTI_MAGIC) {
				legacy.Nox_xxx_mobHealSomeone_5411A0(u.SObj())
			}
		} else {
			if nox_xxx_monsterLookAtDamager_5466B0(u) {
				return
			}
			if s.Frame()-ud.Field137 <= s.TickRate()/2 || u.PosVec == u.PrevPos {
				if !u.HasEnchant(server.ENCHANT_ANTI_MAGIC) {
					legacy.Nox_xxx_mobHealSomeone_5411A0(u.SObj())
				}
			} else {
				u.monsterPushAction(ai.ACTION_FACE_LOCATION, u.PrevPos)
			}
		}
	}
}

type AIActionWait struct{}

func (AIActionWait) Type() ai.ActionType {
	return ai.ACTION_WAIT
}
func (AIActionWait) Start(_ *server.Object) {}
func (AIActionWait) End(_ *server.Object)   {}

func (AIActionWait) Update(obj *server.Object) {
	u := asObjectS(obj)
	s := u.getServer()
	ud := u.UpdateDataMonster()
	if s.Frame() >= ud.AIStackHead().ArgU32(0) {
		u.monsterPopAction()
	}
}

func (AIActionWait) Cancel(obj *server.Object) {
	u := asObjectS(obj)
	u.monsterPopAction()
}

type AIActionWaitRel struct{}

func (AIActionWaitRel) Type() ai.ActionType {
	return ai.ACTION_WAIT_RELATIVE
}
func (AIActionWaitRel) Start(_ *server.Object) {}
func (AIActionWaitRel) End(_ *server.Object)   {}

func (AIActionWaitRel) Update(obj *server.Object) {
	u := asObjectS(obj)
	s := u.getServer()
	ud := u.UpdateDataMonster()
	if s.Frame() > ud.Field137+ud.AIStackHead().ArgU32(0) {
		u.monsterPopAction()
	}
}

func (AIActionWaitRel) Cancel(obj *server.Object) {
	u := asObjectS(obj)
	u.monsterPopAction()
}

type AIActionCastOnObj struct{}

func (AIActionCastOnObj) Type() ai.ActionType {
	return ai.ACTION_CAST_SPELL_ON_OBJECT
}
func (AIActionCastOnObj) Start(_ *server.Object)  {}
func (AIActionCastOnObj) End(_ *server.Object)    {}
func (AIActionCastOnObj) Cancel(_ *server.Object) {}

func (AIActionCastOnObj) Update(obj *server.Object) {
	u := asObjectS(obj)
	ud := u.UpdateDataMonster()
	if ud.AIStackHead().ArgU32(2) == 0 {
		u.monsterPopAction()
		return
	}
	legacy.Nox_xxx_mobActionCast_5413B0(u.SObj(), 0)
	if ud.Field120_3 != 0 {
		u.monsterPopAction()
	}
}

type AIActionCastOnPos struct{}

func (AIActionCastOnPos) Type() ai.ActionType {
	return ai.ACTION_CAST_SPELL_ON_LOCATION
}
func (AIActionCastOnPos) Start(_ *server.Object)  {}
func (AIActionCastOnPos) End(_ *server.Object)    {}
func (AIActionCastOnPos) Cancel(_ *server.Object) {}

func (AIActionCastOnPos) Update(obj *server.Object) {
	u := asObjectS(obj)
	ud := u.UpdateDataMonster()
	legacy.Nox_xxx_mobActionCast_5413B0(u.SObj(), 1)
	if ud.Field120_3 != 0 {
		u.monsterPopAction()
	}
}

type AIActionCastDuration struct{}

func (AIActionCastDuration) Type() ai.ActionType {
	return ai.ACTION_CAST_DURATION_SPELL
}

func (AIActionCastDuration) Start(_ *server.Object) {}

func (AIActionCastDuration) Update(obj *server.Object) {
	u := asObjectS(obj)
	ud := u.UpdateDataMonster()
	if ud.AIStackHead().ArgU32(2) != 0 {
		legacy.Nox_xxx_mobActionCast_5413B0(u.SObj(), 0)
	} else {
		u.monsterPopAction()
	}
}

func (AIActionCastDuration) End(obj *server.Object) {
	u := asObjectS(obj)
	s := u.getServer()
	ud := u.UpdateDataMonster()
	s.spells.duration.CancelFor(spell.ID(ud.AIStackHead().ArgU32(0)), u)
}

func (AIActionCastDuration) Cancel(obj *server.Object) {
	u := asObjectS(obj)
	s := u.getServer()
	ud := u.UpdateDataMonster()
	s.spells.duration.CancelFor(spell.ID(ud.AIStackHead().ArgU32(0)), u)
	u.monsterPopAction()
}

type AIActionMorphIntoChest struct{}

func (AIActionMorphIntoChest) Type() ai.ActionType {
	return ai.ACTION_MORPH_INTO_CHEST
}

func (AIActionMorphIntoChest) Start(_ *server.Object)  {}
func (AIActionMorphIntoChest) End(_ *server.Object)    {}
func (AIActionMorphIntoChest) Cancel(_ *server.Object) {}

func (AIActionMorphIntoChest) Update(obj *server.Object) {
	u := asObjectS(obj)
	ud := u.UpdateDataMonster()
	if ud.Field120_3 != 0 {
		u.monsterPopAction()
		ud.Flags360 |= 0x40000
		legacy.Nox_xxx_monsterMarkUpdate_4E8020(u.SObj())
	}
}

type AIActionMorphBackToSelf struct{}

func (AIActionMorphBackToSelf) Type() ai.ActionType {
	return ai.ACTION_MORPH_BACK_TO_SELF
}

func (AIActionMorphBackToSelf) Start(_ *server.Object)  {}
func (AIActionMorphBackToSelf) End(_ *server.Object)    {}
func (AIActionMorphBackToSelf) Cancel(_ *server.Object) {}

func (AIActionMorphBackToSelf) Update(obj *server.Object) {
	u := asObjectS(obj)
	ud := u.UpdateDataMonster()
	if ud.Field120_3 != 0 {
		u.monsterPopAction()
		ud.Flags360 &^= 0x40000
		legacy.Nox_xxx_monsterMarkUpdate_4E8020(u.SObj())
	}
}

type AIActionReport struct{}

func (AIActionReport) Type() ai.ActionType {
	return ai.ACTION_REPORT
}

func (AIActionReport) Start(_ *server.Object)  {}
func (AIActionReport) End(_ *server.Object)    {}
func (AIActionReport) Cancel(_ *server.Object) {}

func (AIActionReport) Update(obj *server.Object) {
	u := asObjectS(obj)
	s := u.getServer()
	ud := u.UpdateDataMonster()
	u.monsterPopAction()
	s.noxScript.ScriptCallback(&ud.ScriptEndOfWaypoint, nil, u.SObj(), server.NoxEventMonsterDone)
}
