package opennox

import (
	"unsafe"

	"github.com/noxworld-dev/opennox-lib/object"
	"github.com/noxworld-dev/opennox-lib/spell"
	"github.com/noxworld-dev/opennox-lib/things"

	"github.com/noxworld-dev/opennox/v1/common/unit/ai"
	"github.com/noxworld-dev/opennox/v1/legacy"
	"github.com/noxworld-dev/opennox/v1/server"
)

func randomColor3() server.Color3 {
	return server.Color3{
		R: byte(randomIntMinMax(0, 255)),
		G: byte(randomIntMinMax(0, 255)),
		B: byte(randomIntMinMax(0, 255)),
	}
}

func objectMonsterInit(sobj *server.Object) {
	obj := asObjectS(sobj)
	s := obj.getServer()
	ud := obj.UpdateDataMonster()
	if !obj.Flags().HasAny(object.FlagDead | object.FlagDestroyed) {
		switch int(obj.TypeInd) {
		case s.CarnivorousPlantID():
			obj.clearActionStack()
			ud.Field328 = float32(float64(*(*float32)(unsafe.Add(ud.Field121, 112))+obj.Shape.Circle.R) + 10.0)
			ud.AIAction340 = uint32(ai.ACTION_GUARD)
		case s.RatID():
			obj.clearActionStack()
			obj.monsterPushAction(ai.ACTION_RANDOM_WALK)
			ud.Aggression = 0.16
			ud.AIAction340 = uint32(ai.ACTION_INVALID)
		case s.FishSmallID(), s.FishBigID():
			obj.clearActionStack()
			obj.monsterPushAction(ai.ACTION_ROAM, 0, 0, 0xff)
			ud.Aggression = 0.16
			ud.AIAction340 = uint32(ai.ACTION_INVALID)
		case s.GreenFrogID():
			obj.clearActionStack()
			obj.monsterPushAction(ai.ACTION_IDLE)
			ud.Aggression = 0.16
			ud.AIAction340 = uint32(ai.ACTION_INVALID)
			ud.Flags360 |= 0x100
		case s.NPCID():
			for i := 0; i < 6; i++ {
				cl := randomColor3()
				nox_xxx_setNPCColor_4E4A90(obj.SObj(), byte(i), &cl)
			}
		}
	}
	switch ai.ActionType(ud.AIAction340) {
	case ai.ACTION_ESCORT:
		obj.monsterPushAction(ai.ACTION_ESCORT, obj.Pos())
	case ai.ACTION_GUARD:
		obj.monsterPushAction(ai.ACTION_GUARD, obj.Pos(), int(obj.Direction1))
	case ai.ACTION_ROAM:
		if nox_xxx_monsterCanAttackAtWill_534390(obj) {
			obj.monsterPushAction(ai.ACTION_HUNT)
		} else {
			obj.monsterPushAction(ai.ACTION_ROAM, 0, 0, uint32(uint8(ud.Field333)))
		}
	case ai.ACTION_FIGHT:
		obj.monsterPushAction(ai.ACTION_FIGHT, obj.Pos(), uint32(s.Frame()))
	case ai.ACTION_INVALID:
		// nop
	default:
		if ud.AIStackInd < 0 {
			obj.monsterPushAction(ai.ACTION_IDLE)
		}
	}
	ud.AIAction340 = uint32(ai.ACTION_INVALID)
	ud.Direction94 = uint32(obj.Direction1)
	ud.Pos95 = obj.Pos()
	h := obj.HealthData
	if h.Cur == h.Max {
		legacy.Nox_xxx_unitSetHP_4E4560(obj.SObj(), uint16(float64(h.Max)*float64(ud.Field338)))
	}
	h.Field2 = h.Cur
	for i := range ud.HealthGraph103 {
		ud.HealthGraph103[i] = h.Cur
	}
	if obj.SubClass().AsMonster().HasAny(object.MonsterNPC | object.MonsterFemaleNPC) {
		obj.SpeedBase = float32(1.7 + float64(ud.Field332)*0.5)
	} else {
		obj.SpeedBase = float32(float64(obj.SpeedBase) * noxRndCounter1.FloatClamp(0.94999999, 1.05))
	}
	if legacy.Nox_xxx_monsterCanCast_534300(obj.SObj()) {
		ud.Field339 = 100
	}
	if ud.Flags360&0x40 != 0 {
		ud.Field339 = 0
	}
	if ud.Flags360&0x8000 != 0 {
		ud.Flags360 |= 0x4000
	}
}

func (obj *Object) monsterCast(spellInd spell.ID, target *server.Object) {
	s := obj.getServer()
	ud := obj.UpdateDataMonster()
	obj.monsterPushAction(ai.DEPENDENCY_UNINTERRUPTABLE)
	sp := s.SpellDefByInd(spellInd)
	if sp.Def.Flags.Has(things.SpellDuration) {
		ts := s.Frame() + uint32(noxRndCounter1.IntClamp(int(s.TickRate()/2), int(2*s.TickRate())))
		obj.monsterPushAction(ai.DEPENDENCY_TIME, ts)
		obj.monsterPushAction(ai.ACTION_CAST_DURATION_SPELL, uint32(spellInd), 0, target)
	} else {
		obj.monsterPushAction(ai.ACTION_CAST_SPELL_ON_OBJECT, uint32(spellInd), 0, target)
	}
	if target.SObj() != obj.SObj() && !obj.monsterActionIsScheduled(ai.ACTION_FLEE) {
		if !sp.Def.Flags.Has(things.SpellTargeted) { // TODO: looks like the flag name is incorrect
			obj.monsterPushAction(ai.ACTION_FACE_OBJECT, target)
			obj.monsterPushAction(ai.DEPENDENCY_BLOCKED_LINE_OF_FIRE, target)
		}

		v9 := *(*uint32)(unsafe.Add(ud.Field121, 212))
		obj.monsterPushAction(ai.DEPENDENCY_OBJECT_FARTHER_THAN, v9, 0, target)
		obj.monsterPushAction(ai.DEPENDENCY_OR)
		pos2 := target.Pos()
		obj.monsterPushAction(ai.ACTION_MOVE_TO, pos2.X, pos2.Y, target)
	}
}

func nox_xxx_monsterCast_540A30(cu *server.Object, spellInd int, a3p *server.Object) {
	asObjectS(cu).monsterCast(spell.ID(spellInd), a3p)
}

func nox_xxx_setNPCColor_4E4A90(obj *server.Object, ind byte, cl *server.Color3) {
	ud := obj.UpdateDataMonster()
	obj.NeedSync()
	ud.Color[ind] = *cl
	if obj.Class().HasAny(object.ClassClientPersist | object.ClassImmobile | object.ClassPlayer) {
		for i, v := range obj.Field140 {
			obj.Field140[i] = v&0xFFFFF000 | 0x4000000
		}
	} else {
		val := sub_4E4C90(obj, 0x400)
		legacy.Sub_4E4500(obj, 0x4000000, 1024, val)
	}
}
