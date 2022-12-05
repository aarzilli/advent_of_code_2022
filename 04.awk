function contains(a0, a1, b0, b1) {
	return (a0 <= b0) && (b1 <= a1)
}

function overlaps(a0, a1, b0, b1) {
	return contains1(a0, a1, b0) || contains1(a0, a1, b1) || contains1(b0, b1, a0) || contains1(b0, b1, a1)
}

function contains1(a0, a1, x) {
	return (x >= a0) && (x <= a1)
}

BEGIN { FS = "[,-]" }

{
	if (contains($1, $2, $3, $4) || contains($3, $4, $1, $2)) {
		part1++
	}
	if (overlaps($1, $2, $3, $4)) {
		part2++
	}
}

END {
	print "PART 1: " part1
	print "PART 2: " part2
}
