package main
import ("fmt"
		"math"
			)
const NMAX int = 100
func main(){
	var A[NMAX]int
	var n, i, biner, j int
	var x, jum, final float64
	biner = 0
	fmt.Scan(&n)
	for i =0; i < n; i++{
		fmt.Scan(&A[i])
	}
	for i =0; i < n; i++{
		biner = (10*biner)+A[i]
	}
	i = 0
	for biner > 0{
		j = biner%10
		x = math.Pow(float64(2), float64(i))
		final = x*float64(j)
		i++
		jum += final
		biner = biner/10
	}
	fmt.Print(jum)
}