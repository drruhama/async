package main

import (
	"fmt"
	"log"
)

type Examination struct {
	Name  string
	Score int
}

func main() {

	examResult := []Examination{
		{
			Name:  "Budi",
			Score: 90,
		},
		{
			Name:  "Andi",
			Score: 85,
		},
		{
			Name:  "Nayla",
			Score: 87,
		},
		{
			Name:  "Danu",
			Score: 80,
		},
		{
			Name:  "Rahmat",
			Score: 75,
		},
		{
			Name:  "Erika",
			Score: 83,
		},
		{
			Name:  "Siska",
			Score: 87,
		},
		{
			Name:  "Mita",
			Score: 94,
		},
		{
			Name:  "Shinta",
			Score: 82,
		},
		{
			Name:  "Jojo",
			Score: 83,
		},
	}

	var unsorted = []int{}

	for _, tampung := range examResult {
		unsorted = append(unsorted, tampung.Score)
	}
	fmt.Print("Bilangan yang dicari median : ", unsorted, "\n")
	sorted := BubbleSort(unsorted...)
	log.Println(sorted)
	//cari median
	var median int
	dataCopy := []int{75, 80, 82, 83, 83, 85, 87, 87, 90, 94}
	l := len(dataCopy)
	if l == 0 {
		fmt.Println("data tidak boleh kosong")
	} else if l%2 == 0 {
		median = (dataCopy[l/2-1] + dataCopy[l/2]) / 2
	} else {
		median = dataCopy[l/2]
	}

	fmt.Println("Median :", median)
	fmt.Println("Min :", dataCopy[0])
	fmt.Println("Max :", dataCopy[l-1])
	rataRata := ratarata(dataCopy, l)
	fmt.Println("Rata-rata :", rataRata)
	//binary search
	find := 81
	fmt.Println("Data Array Terurut : ", dataCopy)
	fmt.Println("Cari :", find)
	var hasil = Binary(find, dataCopy)
	fmt.Println("Hasil diatas 80 :", hasil)

}

func BubbleSort(Nums ...int) (sorted []int) {
	for i := 0; i < len(Nums)-1; i++ {
		for j := i + 1; j < len(Nums); j++ {
			log.Println(Nums[i], ">", Nums[j])
			if Nums[i] > Nums[j] {
				log.Println("change", Nums[i], "in position", i, "to", Nums[j], "in position", j)
				temp := Nums[i]
				Nums[i] = Nums[j]
				Nums[j] = temp
				log.Println(Nums)
			}
		}
		log.Println("hasil sort :", Nums)
	}
	return
}

func ratarata(x []int, n int) int {
	var i int = 0
	var y int = 0
	for i = 0; i < n; i++ {
		y += x[i]
	}
	var hasilRata = y / n
	return hasilRata
}

func Binary(find int, sortedNums []int) (found bool) {
	if len(sortedNums) <= 1 {
		if sortedNums[0] >= find {
			return true
		}
		return false
	}

	for {

		middleIndex := len(sortedNums) / 2

		log.Println("middle index", middleIndex, "sorted nums", sortedNums, "target data", find)

		if sortedNums[middleIndex] == find {
			log.Println("done with middle index", middleIndex, "sorted nums", sortedNums, "target data", find)
			return true
		} else if find < sortedNums[middleIndex] {
			sortedNums = sortedNums[:middleIndex]
		} else if find > sortedNums[middleIndex] {
			sortedNums = sortedNums[middleIndex:]
		}

		if middleIndex == 0 {
			break
		}
	}

	return false
}
