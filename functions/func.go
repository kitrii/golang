package main

import "strconv"

func main() {

	fn := func(input uint) uint {
		if input == 0 {
			return 100
		}
		s := strconv.FormatUint(uint64(input), 10)
		return s
	}

}
