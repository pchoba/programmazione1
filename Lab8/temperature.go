package main

import (
	"fmt"
	"sort"
)

func main() {
	var temp, media, mediana int
	var db_temp, temp_basse []int
	for {
		fmt.Scanf("%d", &temp)
		if temp == 999 {
			break
		} else {
			db_temp = append(db_temp, temp)
		}
	}
	//media
	for i := 0; i < len(db_temp); i++ {
		media += db_temp[i]
	}
	media = media / len(db_temp)
	fmt.Println("La media delle temperature è", media)
	//mediana
	sort.Ints(db_temp)
	if len(db_temp)%2 == 0 {
		num1, num2 := db_temp[len(db_temp)/2], db_temp[len(db_temp)/2+1]
		mediana = (num1 + num2) / 2
		fmt.Println("La mediana è", mediana)
	} else if len(db_temp)%2 == 1 {
		mediana = (len(db_temp) - 1) / 2
		fmt.Println("La mediana è", mediana)
	}
	//temperature sotto la media
	for i := 0; i < len(db_temp); i++ {
		if db_temp[i] < media {
			temp_basse = append(temp_basse, db_temp[i])
		}
	}
	fmt.Println("Le temperature sotto la media sono: ", temp_basse)
	//3 temperature più alte
	temp_basse = db_temp[:3]
	fmt.Println("Le tre temperature più basse sono", temp_basse)
	//3 temperature più basse
	temp_basse = nil
	temp_basse = db_temp[len(db_temp)-3:]
	fmt.Println("Le tre temperature più alte sono", temp_basse)
}
