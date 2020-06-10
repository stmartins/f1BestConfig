package main

import "fmt"

type bestAverage struct {
	moyenne int
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

func main() {

	freins = map[string] []int {
		"freinsGeko" : {7, 4, 12, 4},
		"freinsTheKeeper" : {7, 3, 10, 12},
		"freinsVaccum" : {8, 8, 15, 5},
		"freinsMinimax" : {7, 9, 14, 4},
		"freinsFeather" : {7, 4, 18, 6},
	}
	boites = map[string] []int {
		"boiteEngager" : {8, 5, 5, 6},
		"boiteVortex" : {7, 6, 3, 3},
		"boiteSliders" : {8, 5, 5, 13},
		"boiteGateway" : {7, 4, 6, 4},
	}
	aileronArrieres = map[string] []int {
		"alarStealth" : {7, 2, 3, 10},
		"alarPhazer" : {8, 11, 7, 5},
		"alarContrail" : {10, 10, 4, 4},
		"alarPeregrine" : {13, 2, 3, 3},
		"alarBase" : {10, 8, 4, 4},
	}
	aileronAvants = map[string] []int {
		"alavCarver" : {7, 3, 3, 10},
		"alavLockOn" : {8, 5, 5, 5},
		"alavBigBite" : {13, 4, 4, 5},
		"alavBlazer" : {8, 8, 5, 5},
		"alavBullet" : {12, 4, 5, 4},
	}
	suspensions = map[string] []int {
		"suspBungee" : {7, 4, 4, 4},
		"suspCompressor" : {8, 11, 5, 12},
		"suspInfluencer" : {7, 8, 4, 8},
		"suspPinPoint" : {7, 4, 4, 5},
		"suspMantis" : {7, 3, 3, 5},
	}
	moteurs = map[string] []int {
		"moteurStickler" : {17, 5, 5, 5},
		"moteurGorilla" : {19, 4, 7, 4},
		"moteurBlinker" : {11, 4, 4, 4},
		"moteurBrute" : {16, 5, 11, 5},
		"moteurBigBore" : {19, 4, 4, 4},
	}

	initAllTools(freins, boites, aileronArrieres, aileronAvants, suspensions, moteurs)
	bestConfig := calculBestAverage()
	printBestConfig(bestConfig)

}

func printBestConfig(bestConfig bestAverage)  {
	println("With ",bestConfig.moyenne, " average, the best equipement is:" )
	println(bestConfig.names[0], printStats(freins[bestConfig.names[0]]))
	println(bestConfig.names[1], printStats(boites[bestConfig.names[1]]))
	println(bestConfig.names[2], printStats(aileronArrieres[bestConfig.names[2]]))
	println(bestConfig.names[3], printStats(aileronAvants[bestConfig.names[3]]))
	println(bestConfig.names[4], printStats(suspensions[bestConfig.names[4]]))
	println(bestConfig.names[5], printStats(moteurs[bestConfig.names[5]]))
}

func printStats(st []int) string {
	return fmt.Sprint("puissance: ", st[0], ", aerodynamique: ", st[1], ", adherence: ",st[2], ", fiabilite: ",st[3])
}

func calculBestAverage() bestAverage {
	tmpAvg := 0
	bestAvg := bestAverage{}

	a := len(allFreins)
	b := len(allBoites)
	c := len(allAileAr)
	d := len(allAileAv)
	e := len(allSuspen)
	f := len(allMoteur)

	for i:=0; i < a; i++ {
		for j:=0; j < b; j++ {
			for k:=0; k < c; k++ {
				for l:=0; l < d; l++ {
					for m:=0; m < e; m++ {
						for n:=0; n < f; n++ {
							puissanceAvg := (allFreins[i].puissance + allBoites[j].puissance + allAileAr[k].puissance +
								allAileAv[l].puissance + allSuspen[m].puissance + allMoteur[n].puissance) / 6
							aerodynamiqueAvg := (allFreins[i].aerodynamique + allBoites[j].aerodynamique + allAileAr[k].aerodynamique +
								allAileAv[l].aerodynamique + allSuspen[m].aerodynamique + allMoteur[n].aerodynamique) / 6
							adherenceAvg := (allFreins[i].adherence + allBoites[j].adherence + allAileAr[k].adherence +
								allAileAv[l].adherence + allSuspen[m].adherence + allMoteur[n].adherence) / 6
							fiabiliteAvg := (allFreins[i].fiabilite + allBoites[j].fiabilite + allAileAr[k].fiabilite +
								allAileAv[l].fiabilite + allSuspen[m].fiabilite + allMoteur[n].fiabilite) / 6

							tmpAvg = (puissanceAvg + aerodynamiqueAvg + adherenceAvg + fiabiliteAvg) / 4
							if tmpAvg > bestAvg.moyenne {
								bestAvg.moyenne = tmpAvg
								bestAvg.names[0] = allFreins[i].name
								bestAvg.names[1] = allBoites[j].name
								bestAvg.names[2] = allAileAr[k].name
								bestAvg.names[3] = allAileAv[l].name
								bestAvg.names[4] = allSuspen[m].name
								bestAvg.names[5] = allMoteur[n].name
							}
						}
					}
				}
			}
		}
	}
	return bestAvg
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