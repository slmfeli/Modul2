package main 
import "fmt" 

func main() {     
	var fe int 

    fmt.Print("bilangan: ")     
	fmt.Scanln(&fe) 
    fmt.Print("Faktor: ")     
	for i := 1; i <= fe; i++ {         
		if fe%i == 0 {             
			fmt.Print(i, " ") 
        }     
	}     
	fmt.Println() 
	
    var jumlahFaktor int = 0    
	for i := 1; i <= fe; i++ {         
		if fe%i == 0 {             
			jumlahFaktor++ 
        } 
    }      
	if jumlahFaktor == 2 { 
        fmt.Println("Prima: true") 
    } else {         
		fmt.Println("Prima: false") 
    } 
} 
 