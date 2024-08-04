package block

import (
	"strconv"
)

type CrimsonFence struct {
	Waterlogged bool
	West bool
	East bool
	North bool
	South bool
}

func (b CrimsonFence) Encode() (string, BlockProperties) {
	return "minecraft:crimson_fence", BlockProperties{
		"south": strconv.FormatBool(b.South),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"west": strconv.FormatBool(b.West),
		"east": strconv.FormatBool(b.East),
		"north": strconv.FormatBool(b.North),
	}
}

func (b CrimsonFence) New(props BlockProperties) Block {
	return CrimsonFence{
		Waterlogged: props["waterlogged"] != "false",
		West: props["west"] != "false",
		East: props["east"] != "false",
		North: props["north"] != "false",
		South: props["south"] != "false",
	}
}