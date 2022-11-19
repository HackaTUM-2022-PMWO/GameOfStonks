package stonks

type StonkName string

const (
	stonkEmpty     StonkName = "" // INVALID!
	StonkPaperClip StonkName = "paperClip"
	StonkScissors  StonkName = "scissors"
	StonkPencil    StonkName = "pencil"
	StonkHouse     StonkName = "house"
	StonkMate      StonkName = "mate"
)

var allStonkNames = []StonkName{
	StonkPaperClip,
	StonkScissors,
	StonkPencil,
	StonkHouse,
	StonkMate,
}

func (s StonkName) IsValid() bool {
	for _, st := range allStonkNames {
		if s == st {
			return true
		}
	}
	return false
}
