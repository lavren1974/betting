package main


import (
	"fmt"
	crand "crypto/rand"
	"math/rand"
)

func random(max int, size int) {
//Мой банк
	MyBanc_A := 100000.0
//анк после	
	MyBanc_B := MyBanc_A
	//MyStavka := 10.0
//коэффициент	
	MyKef := 2.0
	
//шаг	
	MyStop := 0
	//MyStavka_B := MyStavka * 1.5
	MyStavka := []float64{10.0, 20.0, 40.0, 80.0, 160.0, 320.0, 640.0, 1280.0, 2550.0}
	//p := []int{2, 3, 5, 7, 11, 13}
	
	
	nominal := make([]int, size) 
	a := make([]int, max)
	
	b := make([]byte, 1)
    crand.Read(b)
    rand.Seed(int64(b[0]))
   
    for i := 0; i < max; i++ {
    
    	
        r := rand.Intn(size) 
        fmt.Println(i, r)
        a[i] = r
        nominal[r]+=1
        
        //for i := 0; i < MyStop; i++ {
        
        //if MyStop == 0{
        	MyBanc_B = MyBanc_B - MyStavka[MyStop]
	        if r == 1 {
				//fmt.Println("+")
				MyBanc_B = MyBanc_B +( MyStavka[MyStop] * MyKef)
				MyStop=0
				fmt.Println(MyBanc_B)
			}else if r == 0{
				MyStop = MyStop + 1	
				 if MyStop == len(MyStavka) {
				 	MyStop=0
				 }
				//fmt.Println(MyBanc_B)		
			}
        //}
        
        //}
        
        
        //if r == 1 {
			//fmt.Println("+")
			//MyBanc_B = MyBanc_B +( MyStavka * MyKef)
		//} 
    }
    fmt.Println(a)
    
    for i := 0; i < len(nominal); i++ {
        fmt.Println(nominal[i])
    }
    
    fmt.Println("MyBanc_A - ", MyBanc_A)
    fmt.Println("MyStavka - ", MyStavka)
    fmt.Println("MyBanc_B - ", MyBanc_B)

}


func main() {

	
	random (10, 2)
}