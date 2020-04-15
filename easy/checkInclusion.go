package easy

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) == 0 || len(s2) == 0 || len(s1) > len(s2) {
		return false
	}

	if s1 == s2 {
		return true
	}

	subMap := string2Map(s1)
	window := string2Map(s2[0:len(s1)])

	var l, r = 0, len(s1)
	for {
		//if len(subMap) > len(window) {
		//	return false
		//}

		if equalsMap(subMap, window) {
			return true
		}

		delete(window, s2[l])

		l++
		r++

		if r < len(s2)+1 {
			var b byte
			if r == len(s2) {
				break
			} else {
				b = s2[r]
			}
			if _, ok := window[b]; !ok {
				window[b] = 1
			} else {
				window[b] = window[b] + 1
			}
		} else {
			break
		}
		fmt.Println(s2, s1, s2[l:r])
	}

	return false
}

func string2Map(s string) map[byte]int {
	var m = map[byte]int{}
	for _, k := range s {
		r := byte(k)
		if _, ok := m[r]; !ok {
			m[r] = 1
		} else {
			m[r] = m[r] + 1
		}
	}
	return m
}

func equalsMap(m1, m2 map[byte]int) bool {
	if len(m1) != len(m2) {
		return false
	}

	for k, v := range m1 {
		if v2, ok := m2[k]; !ok || v2 != v {
			return false
		}
	}

	return true
}

//func checkInclusion(s1 string, s2 string) bool {
//	if len(s1) == 0 || len(s2) == 0 || len(s1) > len(s2) {
//		return false
//	}
//
//	if s1 == s2 {
//		return true
//	}
//
//	getMap := func() map[byte]int {
//		var m = map[byte]int{}
//		for i := 0; i < len(s1); i++ {
//			_, ok := m[s1[i]]
//			if ok {
//				m[s1[i]] = m[s1[i]] + 1
//			} else {
//				m[s1[i]] = 1
//			}
//		}
//		return m
//	}
//	m := getMap()
//
//	for i := 0; i < len(s2); i++ {
//		if !exists(s2, i, m) {
//			continue
//		} else { // matched current index
//			if len(m) == 1 && m[s2[i]] == 1 {
//				return true
//			}
//
//			//if exists(s2, i-1, m) {
//			// match left
//			matched := match(s2, i, m, true, len(s1))
//			if matched {
//				return true
//			}
//			//}
//
//			//if exists(s2, i+1, m) {
//			// match right
//			matched = match(s2, i, m, false, len(s1))
//			if matched {
//				return true
//			}
//			//}
//
//		}
//	}
//	return false
//}
//
//func exists(s string, i int, m map[byte]int) bool {
//	if i >= 0 && i < len(s) {
//		_, ok := m[s[i]]
//		return ok
//	}
//	return false
//}
//
//func match(s string, i int, m map[byte]int, goLeft bool, total int) bool {
//	var matched = map[byte]int{}
//
//	for i >= 0 && i < len(s) && total > 0 {
//		n++
//		total--
//		r := s[i]
//		// not exists in map
//		if _, ok := m[r]; !ok {
//			break
//		}
//
//		// count matched
//		if _, ok := matched[r]; ok {
//			matched[r] = matched[r] + 1
//		} else {
//			matched[r] = 1
//		}
//
//		if matched[r] > m[r] {
//			return false
//		}
//
//		if goLeft {
//			i--
//		} else {
//			i++
//		}
//	}
//
//	if len(matched) != len(m) {
//		return false
//	}
//
//	for k, v := range matched {
//		c, ok := m[k]
//		if !ok {
//			return false
//		}
//
//		if c != v {
//			return false
//		}
//	}
//
//	return true
//}
