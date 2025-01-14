package main

import "strconv"

/* Postgres can return duration from AGE() method. This return has []uint8 type.
* This method transforms that to the duration in seconds */

func Fromuint8ToInt(entry []uint8) int {

	entryLen := len(entry)

	total, multiplier := 0, 1

	if entryLen == 0 {
		return total
	}

	for i := entryLen - 1; i >= 0; i-- {

		holder, _ := strconv.Atoi(string(entry[i]))

		switch i {
		case entryLen - 1:
			total += holder * multiplier // 1
			multiplier *= 10
		case entryLen - 2:
			total += holder * multiplier // 10
			multiplier *= 10
		case entryLen - 3:
			multiplier *= 6
			continue
		case entryLen - 4:
			total += holder * multiplier // 60
		case entryLen - 5:
			total += holder * multiplier // 60
		case entryLen - 6:
			multiplier *= 6
			continue
		default:
			total += holder * multiplier
			multiplier *= 10
		}
	}

	return total
}
