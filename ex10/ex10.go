package main

import (
	"fmt"
	"sort"
)

func main() {
	seq := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 0, 3.1, 3.9, -50.4} // Входные данные
	result := make(map[int][]float32)                                                       // Результирующая мапа. В качестве ключа - минимальный порог. В качестве значения слайс данных попадающих в порог

	// Сортируем последовательность
	sort.Slice(seq, func(i, j int) bool {
		return seq[i] < seq[j]
	})

	// Получаем порог
	min, max := getMinMax(seq[0])
	for _, v := range seq {
		// Если элемент попадает в порог - добавляем в группу
		if int(v) >= min && int(v) <= max {
			result[min] = append(result[min], v)
		} else { // Иначе обновляем порог и добавляем в соответствующую группу
			min, max = getMinMax(v)
			result[min] = append(result[min], v)
		}
	}
	fmt.Println(result)
}

// Принимает элемент
// Возвращает порог
func getMinMax(a float32) (min, max int) {
	min = int(a) - int(a)%10
	if min < 0 {
		min -= 10
	}
	max = min + 10
	return int(min), int(max)
}
