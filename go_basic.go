package main
import ("fmt")


func main(){
	name := "Sourav Singha"
	fmt.Println(name)
	var arr = [3]int {1,2,3}
	var arrIn = [...]int{1,2,3,4,5}
	fmt.Println(arr[0])
	fmt.Println(arrIn)
	length := len(arrIn)
	fmt.Printf("Length of array is : %d\n",length)
}