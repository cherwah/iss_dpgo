package asset

type Asset struct {
	power  int
	health int
}

type Human_Image struct {
	Image [128]byte
}

type Monster_Image struct {
	Image [512]byte
}

type Human struct {
	Asset
	Image *Human_Image
}

type Monster struct {
	Asset
	Image *Monster_Image
}

type IPlayable interface {
	Get_Type() string
}

var human_img *Human_Image
var monster_img *Monster_Image

func New_Human() Human {
	if human_img == nil {
		human_img = &Human_Image{}
	}

	return Human{
		Asset: Asset{
			power:  10,
			health: 100,
		},
		Image: human_img,
	}
}

func New_Monster() Monster {
	if monster_img == nil {
		monster_img = &Monster_Image{}
	}

	return Monster{
		Asset: Asset{
			power:  50,
			health: 100,
		},
		Image: monster_img,
	}
}

func (h *Human) Get_Type() string {
	return "H"
}

func (m *Monster) Get_Type() string {
	return "M"
}
