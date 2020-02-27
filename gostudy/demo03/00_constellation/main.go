package main

import "fmt"

var constellationMatchings = [12][12]int{
	{80, 70, 90, 50, 100, 40, 60, 40, 100, 50, 90, 70},
	{70, 80, 70, 90, 50, 100, 40, 60, 40, 100, 50, 90},
	{90, 70, 80, 70, 90, 50, 100, 40, 60, 40, 100, 50},
	{50, 90, 70, 80, 70, 90, 50, 100, 40, 60, 40, 100},
	{100, 50, 90, 70, 80, 70, 90, 50, 100, 40, 60, 40},
	{40, 100, 50, 90, 70, 80, 70, 90, 50, 100, 40, 60},
	{60, 40, 100, 50, 90, 70, 80, 70, 90, 50, 100, 40},
	{40, 60, 40, 100, 50, 90, 70, 80, 70, 90, 50, 100},
	{100, 40, 60, 40, 100, 50, 90, 70, 80, 70, 90, 50},
	{50, 100, 40, 60, 40, 100, 50, 90, 70, 80, 70, 90},
	{90, 50, 100, 40, 60, 40, 100, 50, 90, 70, 80, 70},
	{70, 90, 50, 100, 40, 60, 40, 100, 50, 90, 70, 80},
}

var constellationArray = []*Constellation{
	&Constellation{0, "Aries", "白羊座"},
	&Constellation{1, "Taurus", "金牛座"},
	&Constellation{2, "Gemini", "双子座"},
	&Constellation{3, "Cancer", "巨蟹座"},
	&Constellation{4, "Leo", "狮子座"},
	&Constellation{5, "Virgo", "处女座"},
	&Constellation{6, "Libra", "白羊座"},
	&Constellation{7, "Scorpio", "天蝎座"},
	&Constellation{8, "Sagittarius", "射手座"},
	&Constellation{9, "Capricorn", "摩羯座"},
	&Constellation{10, "Aquarius", "水瓶座"},
	&Constellation{11, "Pisces", "双鱼座"},
}

// Constellation 星座
type Constellation struct {
	Index  int
	EnName string
	CnName string
}

type zodiac struct {
	Constellations []*Constellation
}

var zodiacList zodiac

func (z *zodiac) compareByEnName(male, female string) int {
	var mIndex, fIndex int
	for _, constellation := range z.Constellations {
		if constellation.EnName == male {
			mIndex = constellation.Index
		} else if constellation.EnName == female {
			fIndex = constellation.Index
		}
	}
	return constellationMatchings[mIndex][fIndex]
}

func (z *zodiac) compareByCnName(male, female string) int {
	var mIndex, fIndex int
	for _, constellation := range z.Constellations {
		if constellation.CnName == male {
			mIndex = constellation.Index
		} else if constellation.CnName == female {
			fIndex = constellation.Index
		}
	}
	return constellationMatchings[mIndex][fIndex]
}

func (z *zodiac) compareByIndex(male, female int) int {
	return constellationMatchings[male][female]
}

func init() {

	zodiacList = zodiac{constellationArray}
	name := zodiacList.Constellations
	fmt.Print("       ")
	for i := 0; i < 12; i++ {
		fmt.Printf("%4s", name[i].CnName)
	}
	fmt.Println()
	for i := 0; i < 12; i++ {
		fmt.Printf("%4s", name[i].CnName)
		for j := 0; j < 12; j++ {
			fmt.Printf("%7d", constellationMatchings[i][j])
		}
		fmt.Println()
	}
}

func main() {
	fmt.Println(zodiacList.compareByIndex(1, 3))
	fmt.Println(zodiacList.compareByEnName("Sagittarius", "Aries"))
	fmt.Println(zodiacList.compareByCnName("巨蟹座", "天蝎座"))
}
