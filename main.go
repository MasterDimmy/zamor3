/*
	GNU GPL FREE
	Обязательное указание автора при использовании
	Эдель Дмитрий, 2016
*/

package main

import (
	"fmt"
)

const (
	kenguru_left = iota + 1
	kenguru_right
	suslik_left
	suslik_right
	cerepaha_left
	cerepaha_right
	koala_left
	koala_right
)

type karta struct {
	l     int
	r     int
	u     int
	d     int
	level int
}

var zero = karta{0, 0, 0, 0, 0}

var all = []karta{
	{suslik_right, suslik_left, kenguru_left, koala_left, 0},
	{cerepaha_right, cerepaha_right, koala_right, kenguru_left, 0},
	{suslik_right, kenguru_right, koala_left, suslik_left, 0},
	{kenguru_right, koala_left, suslik_left, kenguru_right, 0},
	{kenguru_left, koala_left, cerepaha_right, koala_right, 0},
	{cerepaha_right, koala_right, cerepaha_right, suslik_right, 0},
	{koala_left, cerepaha_right, suslik_left, kenguru_left, 0},
	{cerepaha_right, cerepaha_right, kenguru_right, koala_left, 0},
	{kenguru_left, suslik_left, cerepaha_left, suslik_left, 0},
}

//по часовой
//pos = 0 начальная, если pos==4 то сделали полный круг
func rotate(a karta) karta {
	return karta{a.d, a.u, a.l, a.r, a.level}
}

type Tfield [4][4]karta

func karta_to_str(a karta) string {
	return string(enum_to_str(a.l)[0]) + string(enum_to_str(a.r)[0]) + string(enum_to_str(a.u)[0]) + string(enum_to_str(a.d)[0])
}

func enum_to_str(a int) string {
	switch a {
	case kenguru_left:
		return "ul" //keng
	case kenguru_right:
		return "ur" //keng
	case suslik_left:
		return "sl"
	case suslik_right:
		return "sr"
	case cerepaha_left:
		return "cl"
	case cerepaha_right:
		return "cr"
	case koala_left:
		return "kl"
	case koala_right:
		return "kr"
	}
	return "????"
}

var map_limit int = 3

/*
--- --- ---
--- --- ---
--- --- ---

--- --- ---
--- --- ---
--- --- ---

--- --- ---
--- --- ---
--- --- ---
*/
func print_em2(i int) {
	if i == 0 {
		fmt.Printf("┌")
	} else {
		if i == map_limit*map_limit-1 {
			fmt.Printf("└")
		} else {
			fmt.Printf("|")
		}
	}
	for j := 0; j < map_limit*map_limit*2+2; j++ { //столбец
		fmt.Printf("-")
	}
	if i == 0 {
		fmt.Printf("┐")
	} else {
		if i == map_limit*map_limit-1 {
			fmt.Printf("┘")
		} else {
			fmt.Printf("|")
		}
	}
	fmt.Printf("\n")
}

func print_em() {
	var sa [][]string
	for i := 0; i < map_limit*map_limit; i++ { //ряд
		var st []string
		for j := 0; j < map_limit*map_limit; j++ { //ряд
			st = append(st, "  ")
		}
		sa = append(sa, st)
	}

	for i := 0; i < map_limit; i++ { //ряд
		for j := 0; j < map_limit; j++ { //столбец
			//sk := karta_to_str(field[i][j])
			sa[i*3][j*3+1] = enum_to_str(field[i][j].u)   //u
			sa[i*3+2][j*3+1] = enum_to_str(field[i][j].d) //d
			sa[i*3+1][j*3] = enum_to_str(field[i][j].l)   //l
			sa[i*3+1][j*3+2] = enum_to_str(field[i][j].r) //r
			//sa[i*3+1][j*3+1] = fmt.Sprintf("%d", field[i][j].level)[0]
		}
	}
	for i := 0; i < map_limit*map_limit; i++ { //ряд
		if i%3 == 0 {
			print_em2(i)
		}

		for j := 0; j < map_limit*map_limit; j++ { //столбец
			if j%3 == 0 {
				fmt.Printf("|")
			}
			fmt.Printf("%s", string(sa[i][j]))
			if j == map_limit*map_limit-1 {
				fmt.Printf("|")
			}
		}
		fmt.Printf("\n")
		if i == map_limit*map_limit-1 {
			print_em2(i)
			fmt.Printf("\n")
		}
	}
	fmt.Printf("\n")
}

//проверить соответствие для карты
func check(i int, j int) bool {
	//слева
	if j > 0 && ((field[i][j].l != field[i][j-1].r) || (field[i][j-1].r == 0)) {
		return false
	}
	//сверху
	if i > 0 && ((field[i][j].u != field[i-1][j].d) || (field[i-1][j].d == 0)) {
		return false
	}
	return true
}

var field Tfield

type Tused map[int]bool

func print_used(u Tused) {
	for i := 0; i < 9; i++ {
		r := u[i]
		s := ' '
		if r {
			s = 'T'
		}
		fmt.Printf("%d[%c] ", i, s)
	}
	fmt.Printf("\n")
}

var sol = 1

//поиск i,j начальная текущая
func search(i int, j int, level int, used Tused) bool {
	//fmt.Printf("search: %d %d ", i, j)
	//print_used(used)
	for ka := 0; ka < 9; ka++ {
		a, ok := used[ka]
		if ok && a {
			continue
		}
		field[i][j] = all[ka]
		field[i][j].level = level
		//print_em()
		nowused := make(Tused, 9)
		for r, t := range used {
			nowused[r] = t
		}
		nowused[ka] = true
		for pos := 0; pos < 4; pos++ {
			ok := func(i int, j int, used Tused) bool {
				if !check(i, j) {
					return false
				}
				for k := i; k < 3; k++ {
					for l := 0; l < 3; l++ {
						if field[k][l].level == 0 {
							ok := search(k, l, level+1, used)
							if ok {
								return true
							}
						}
					}
				}
				if level == 9 {
					if sol == 1 {
						fmt.Printf(`Обозначения: 
u - кенгуру
s - суслик
c - черепаха
k - коала

`)
					}
					fmt.Printf("Решение #%d\n", sol)
					sol++
					print_em()
					return true
				}
				return false
			}(i, j, nowused)
			if ok {
				if level != 1 {
					return true
				}
			}
			field[i][j] = rotate(field[i][j])
			//print_em()
		}
	}
	field[i][j] = zero
	return false
}

func main() {
	used := make(Tused, 9)
	for i := 0; i < 9; i++ {
		used[i] = false
	}
	search(0, 0, 1, used)
	if sol == 1 {
		fmt.Printf("\nрешений не найдено\n")
	}
}
