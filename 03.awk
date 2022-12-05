function tomap(a, r, i) {
	for(i = 0; i <= length(a); i++) {
		r[substr(a, i, 1)] = 1
	}
}

BEGIN { split("", part2) }

{
	part2[length(part2)] = $0
	
	if (length(part2) == 3) {
		split("", a)
		split("", b)
		split("", c)
		tomap(part2[0], a)
		tomap(part2[1], b)
		tomap(part2[2], c)
		
		for(key in a) {
			if ((b[key] == 1) && (c[key] == 1)) {
				part2sol = part2sol key
			}
		}
		
		
		split("", part2)
	}
	
	split("", a)
	split("", b)
	tomap(substr($0, 0, length($0)/2), a)
	tomap(substr($0, length($0)/2+1), b)
	for(key in a) {
		if (b[key] == 1) {
			part1 = part1 key
		}
	}
}

END {
	print "PART 1: " part1
	print "PART 2: " part2sol
}
