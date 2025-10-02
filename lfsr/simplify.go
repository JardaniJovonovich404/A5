package lfsr

import (
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func splitAlphaNum(s string) []string {
	re := regexp.MustCompile(`\d+|\D+`)
	return re.FindAllString(s, -1)
}

func lessLex(a, b string) bool {
	as := splitAlphaNum(a)
	bs := splitAlphaNum(b)

	for i := 0; i < len(as) && i < len(bs); i++ {
		if ai, err1 := strconv.Atoi(as[i]); err1 == nil {
			if bi, err2 := strconv.Atoi(bs[i]); err2 == nil {
				if ai != bi {
					return ai < bi
				}
				continue
			}
		}
		if as[i] != bs[i] {
			return as[i] < bs[i]
		}
	}
	return len(as) < len(bs)
}

func SimplifyXOR(expr string) string {
	parts := strings.Split(expr, "⊕")
	counts := make(map[string]int)
	for _, p := range parts {
		term := strings.TrimSpace(p)
		if term == "" || term == "0" {
			continue
		}
		counts[term]++
	}

	var result []string
	for term, c := range counts {
		if c%2 == 1 {
			result = append(result, term)
		}
	}

	if len(result) == 0 {
		return "0"
	}

	sort.Slice(result, func(i, j int) bool {
		return lessLex(result[i], result[j])
	})

	return strings.Join(result, " ⊕ ")
}
