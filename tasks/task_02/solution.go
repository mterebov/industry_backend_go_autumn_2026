package main

func rotateRunes(s string, shift int) string {
	// Приводим строку к слайсу рун и вычисляем ее длину
	runes := []rune(s)
	n := len(runes)

	// Пустую строку сдвигать некуда (и делить на ноль нельзя)
	if n == 0 {
		return ""
	}

	// Вычисляем позицию для разделения строки на 2 части.
	// Сдвиг вправо на k — то же самое, что сдвиг влево на n-k,
	// поэтому приводим сдвиг к диапазону [0, n)
	cutPos := shift % n
	if cutPos < 0 {
		cutPos += n
	}

	// Выполняем сдвиг влево: сначала хвост, потом голова
	result := make([]rune, 0, n)
	result = append(result, runes[cutPos:]...)
	result = append(result, runes[:cutPos]...)

	return string(result)
}
