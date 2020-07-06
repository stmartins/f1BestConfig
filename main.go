package main

import "fmt"

type bestAverage struct {
	moyenne float32
	names [6]string
}

type stats struct {
	name string
	puissance int
	aerodynamique int
	adherence int
	fiabilite int
}

var (
	allFreins []stats
	allBoites []stats
	allAileAr []stats
	allAileAv []stats
	allSuspen []stats
	allMoteur []stats
	freins map[string][]int
	boites map[string][]int
	aileronArrieres map[string][]int
	aileronAvants map[string][]int
	suspensions map[string][]int
	moteurs map[string][]int
)

func initElements() {
	freins = map[string][]int{
		"freinsGeko":      {8, 5, 15, 5},
		"freinsTheKeeper": {7, 4, 12, 13},
		"freinsVaccum":    {8, 9, 19, 5},
		"freinsMinimax":   {7, 0, 18, 4},
		"freinsFeather":   {8, 5, 24, 8},
	}
	boites = map[string][]int{
		"boiteEngager": {8, 6, 6, 6},
		"boiteVortex":  {7, 8, 4, 4},
		"boiteMsm":  {7, 3, 3, 7},
		"boiteSliders": {8, 5, 5, 15},
		"boiteGateway": {7, 4, 7, 4},
	}
	aileronArrieres = map[string][]int{
		"alarStealth":   {7, 2, 3, 10},
		"alarPhazer":    {8, 12, 8, 5},
		"alarContrail":  {10, 11, 4, 4},
		"alarPeregrine": {13, 2, 3, 3},
		"alarLock&Load": {7, 12, 4, 4},
		"alarBase":      {10, 14, 5, 5},
	}
	aileronAvants = map[string][]int{
		"alavCarver":  {7, 3, 3, 10},
		"alavLockOn":  {8, 5, 5, 5},
		"alavBigBite": {13, 4, 4, 5},
		"alavBlazer":  {8, 11, 5, 6},
		"alavBullet":  {15, 4, 7, 4},
	}
	suspensions = map[string][]int{
		"suspBungee":     {7, 4, 4, 4},
		"suspCompressor": {8, 12, 5, 12},
		"suspInfluencer": {7, 9, 4, 9},
		"suspPinPoint":   {8, 5, 5, 8},
		"suspMantis":     {7, 4, 4, 10},
	}
	moteurs = map[string][]int{
		"moteurStickler": {19, 5, 5, 5},
		"moteurGorilla":  {21, 5, 7, 5},
		"moteurBlinker":  {11, 4, 4, 4},
		"moteurBrute":    {17, 5, 11, 5},
		"moteurBigBore":  {24, 4, 4, 4},
		"moteurSmoothOp":  {17, 4, 4, 7},
	}
}

func printBestConfig(bestConfig bestAverage)  {
	fmt.Println("With ",bestConfig.moyenne, " average, the best equipement is:" )
	fmt.Println(bestConfig.names[0], printStats(freins[bestConfig.names[0]]))
	fmt.Println(bestConfig.names[1], printStats(boites[bestConfig.names[1]]))
	fmt.Println(bestConfig.names[2], printStats(aileronArrieres[bestConfig.names[2]]))
	fmt.Println(bestConfig.names[3], printStats(aileronAvants[bestConfig.names[3]]))
	fmt.Println(bestConfig.names[4], printStats(suspensions[bestConfig.names[4]]))
	fmt.Println(bestConfig.names[5], printStats(moteurs[bestConfig.names[5]]))
}

func printStats(st []int) string {
	return fmt.Sprint("puissance: ", st[0], ", aerodynamique: ", st[1], ", adherence: ",st[2], ", fiabilite: ",st[3])
}

func calculBestAverage() bestAverage {
	tmpAvg := float32(0)
	bestAvg := bestAverage{}

	a, b, c, d, e, f := setLimits()
tt:= 0
	for i := 0; i < a; i++ {
		for j := 0; j < b; j++ {
			for k := 0; k < c; k++ {
				for l := 0; l < d; l++ {
					for m := 0; m < e; m++ {
						for n := 0; n < f; n++ {
							tt++
							println(tt, "/", a*b*c*d*e*f)
							puissanceAvg, aerodynamiqueAvg, adherenceAvg, fiabiliteAvg := elementsAverage(i, j, k, l, m, n)
							tmpAvg = (puissanceAvg + aerodynamiqueAvg + adherenceAvg + fiabiliteAvg) / 4
							if tmpAvg > bestAvg.moyenne {
								fmt.Println("in best average puissance:", puissanceAvg*6, "aero:", aerodynamiqueAvg*6, "adher:", adherenceAvg*6, "fiab:", fiabiliteAvg*6, "total:", tmpAvg)
								funcName(&bestAvg, tmpAvg, i, j, k, l, m, n)
							}
						}
					}
				}
			}
		}
	}
	return bestAvg
}

func funcName(bestAvg *bestAverage, tmpAvg float32, i int, j int, k int, l int, m int, n int) {
	bestAvg.moyenne = tmpAvg
	bestAvg.names[0] = allFreins[i].name
	bestAvg.names[1] = allBoites[j].name
	bestAvg.names[2] = allAileAr[k].name
	bestAvg.names[3] = allAileAv[l].name
	bestAvg.names[4] = allSuspen[m].name
	bestAvg.names[5] = allMoteur[n].name
}

func elementsAverage(i int, j int, k int, l int, m int, n int) (float32, float32, float32, float32) {
	puissanceAvg := float32((allFreins[i].puissance + allBoites[j].puissance + allAileAr[k].puissance +
		allAileAv[l].puissance + allSuspen[m].puissance + allMoteur[n].puissance) / 6)
	aerodynamiqueAvg := float32((allFreins[i].aerodynamique + allBoites[j].aerodynamique + allAileAr[k].aerodynamique +
		allAileAv[l].aerodynamique + allSuspen[m].aerodynamique + allMoteur[n].aerodynamique) / 6)
	adherenceAvg := float32((allFreins[i].adherence + allBoites[j].adherence + allAileAr[k].adherence +
		allAileAv[l].adherence + allSuspen[m].adherence + allMoteur[n].adherence) / 6)
	fiabiliteAvg := float32((allFreins[i].fiabilite + allBoites[j].fiabilite + allAileAr[k].fiabilite +
		allAileAv[l].fiabilite + allSuspen[m].fiabilite + allMoteur[n].fiabilite) / 6)
	return puissanceAvg, aerodynamiqueAvg, adherenceAvg, fiabiliteAvg
}

func setLimits() (int, int, int, int, int, int) {
	a := len(allFreins)
	b := len(allBoites)
	c := len(allAileAr)
	d := len(allAileAv)
	e := len(allSuspen)
	f := len(allMoteur)
	return a, b, c, d, e, f
}

func initAllTools(freins, boites, aileronArrieres, aileronAvants, suspensions, moteurs map[string][]int)  {
	allFreins = initFields(freins)
	allBoites = initFields(boites)
	allAileAr = initFields(aileronArrieres)
	allAileAv = initFields(aileronAvants)
	allSuspen = initFields(suspensions)
	allMoteur = initFields(moteurs)
}

func initFields(tool map[string][]int) []stats {
	i := 0
	allTools := make([]stats, len(tool))
	for name, stat := range tool {
		allTools[i] = initStats(name, stat[0], stat[1], stat[2], stat[3])
		i++
	}
	return allTools
}

func initStats(name string, puissance, aerodynamique, adherance, fiabilite int) stats {
	return stats{
		name: 		   name,
		puissance:     puissance,
		aerodynamique: aerodynamique,
		adherence:     adherance,
		fiabilite:     fiabilite,
	}
}

func main() {

	initElements()
	initAllTools(freins, boites, aileronArrieres, aileronAvants, suspensions, moteurs)
	bestConfig := calculBestAverage()
	printBestConfig(bestConfig)
}
