package main 
import (     
	"fmt" 
    "math" 
) 

func f(k int) float64 {
	numerator := math.Pow(float64(4*k+2), 2)     
	denominator := float64((4*k + 1) * (4*k + 3))     
	return numerator / denominator 
}  

func sqrt2Approximation(K int) float64 {     
	result := 1.0     
	for k := 0; k < K; k++ {         
		result *= f(k) 
    }     
	return result 
	}  

func main() {     
	var K int 
    fmt.Print("Nilai K = ")     
	fmt.Scan(&K) 
      
    approx := sqrt2Approximation(K)       
    fmt.Printf("Nilai akar 2 = %.10f\n", approx) 
} 