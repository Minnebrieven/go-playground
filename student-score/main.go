package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Nama  string
	Score int
}

func panicHandler() {
	isPanic := recover()

	if isPanic != nil {
		fmt.Println("Recover", isPanic)
	}
}

func minMaxScore(data []Student) (min, max Student) {
	var currentMin, currentMax Student = data[0], data[0]
	for _, value := range data {
		if value.Score > currentMax.Score {
			currentMax.Score = value.Score
			currentMax.Nama = value.Nama
		}

		if value.Score < currentMin.Score {
			currentMin.Score = value.Score
			currentMin.Nama = value.Nama
		}
	}
	return currentMin, currentMax
}

func averageScore(data []Student) int {
	var average, count int
	for _, value := range data {
		average += value.Score
		count += 1
	}
	return average / count
}

// func (s Student) NewStudent(namaStudent string, scoreStudent int) Student{
// 	student = Student{
// 		Nama: namaStudent,
// 		Score: scoreStudent,
// 	}
// 	return student
// }

func main() {
	var studentCount, scoreSiswa int
	var namaSiswa string
	var studentSlices = []Student{}

	defer panicHandler()
	fmt.Print("Masukan Jumlah siswa : ")
	fmt.Scan(&studentCount)
	switch {
	case reflect.TypeOf(studentCount).Kind() != reflect.Int:
		panic("Input haruslah angka")
	case studentCount <= 0:
		panic("Input tidak bisa angka 0")
	}

	i := 1
	for i <= studentCount {

		fmt.Printf("Masukan	Nama  siswa %d: ", i)
		fmt.Scan(&namaSiswa)
		fmt.Printf("Masukan	Score siswa %d: ", i)
		fmt.Scan(&scoreSiswa)
		switch {
		case reflect.TypeOf(scoreSiswa).Kind() != reflect.Int:
			panic("Input haruslah angka")
		case scoreSiswa <= 0:
			panic("Input tidak bisa angka 0")
		}
		student := Student{
			Nama:  namaSiswa,
			Score: scoreSiswa,
		}
		studentSlices = append(studentSlices, student)
		i++
	}
	fmt.Printf("Average Score = %d\n", averageScore(studentSlices))
	min, max := minMaxScore(studentSlices)
	fmt.Printf("Minimum Score = %s (%d)\nMaximum Score = %s (%d)", min.Nama, min.Score, max.Nama, max.Score)
}
