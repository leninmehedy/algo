package main

// Author: Lenin Mehedy (leninmehedy@gmail.com)

// LCS finds the longest common subsequence between two strings
// it returns the LCS if found
func LCS(s1, s2 string) string {
	// convert into runes so that we can also support UTF8 characters
	r1 := []rune(s1)
	r2 := []rune(s2)

	// create the length matrix for dynamic programming
	m := make([][]int, len(r1)+1)

	// initialize the length matrix
	// O(n)
	for i := 0; i <= len(r1); i++ {
		m[i] = make([]int, len(r2)+1)
	}

	// calculate the LCS length using dynamic programming
	// O(mn)
	for i := 1; i <= len(r1); i++ {
		for j := 1; j <= len(r2); j++ {
			if r1[i-1] == r2[j-1] {
				m[i][j] = m[i-1][j-1] + 1
			} else {
				if m[i-1][j] >= m[i][j-1] {
					m[i][j] = m[i-1][j]
				} else {
					m[i][j] = m[i][j-1]
				}
			}
		}
	}

	// traverse back for the LCS
	// O(max(m,n))
	r3 := make([]rune, 0)
	for i, j := len(r1), len(r2); i > 0 && j > 0; {
		if m[i-1][j-1] < m[i][j] {
			r3 = append(r3, r1[i-1])
			i--
			j--
		} else if m[i-1][j] == m[i][j] {
			i--
		} else {
			j--
		}
	}

	// reverse to get the correct sequence for LCS
	// O(n/2)
	for i, j := 0, len(r3)-1; i < j; i, j = i+1, j-1 {
		r3[i], r3[j] = r3[j], r3[i]
	}

	// return LCS
	return string(r3)
}
