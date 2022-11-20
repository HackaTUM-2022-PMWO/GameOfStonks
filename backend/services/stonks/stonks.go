package stonks

type StonkName string

const (
	StonkPaperClip StonkName = "paperClip"
	StonkScissors  StonkName = "scissors"
	StonkPencil    StonkName = "pencil"
	StonkHouse     StonkName = "house"
	StonkMate      StonkName = "mate"
)

var AllStonkNames = []StonkName{
	StonkPaperClip,
	StonkScissors,
	StonkPencil,
	StonkHouse,
	StonkMate,
}

func (s StonkName) IsValid() bool {
	for _, st := range AllStonkNames {
		if s == st {
			return true
		}
	}
	return false
}
