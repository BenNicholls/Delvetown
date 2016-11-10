package data

//enumerated stats for modifiers.
const (
	HP int = iota
	MP
	ATT
	DEF
	ATTSPEED
	MVSPEED
	SIGHTRANGE
	LIGHTSTRENGTH
	MIND
	BODY
	SPIRIT
	MAX_STATS
)

type Stats struct {
	HP int
	MP int
	Attack int
	Defense int
	AttackSpeed int
	MoveSpeed int
    SightRange int
    LightStrength int
	Mind, Body, Spirit int
}

type Modifier struct {
	name string
	desc string
	stat int //one of the enumerated values above
	value int
	absolute bool //if true, modifier is flat. otherwise, a percentage muliplier (ex 110 for 10% increase)
	permanent bool
	duration int
}

