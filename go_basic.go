package main
import ("fmt")


func compare(x int, y int){
	if (x>y){
		fmt.Println("Fuck ",x ,"is greater than",y)
	}else{
		fmt.Println("Fuck ",x ,"is less than",y)
	}
}

func main(){
	name := "Sourav Singha"
	fmt.Println(name)
	var arr = [3]int {1,2,3}
	var arrIn = [...]int{1,2,3,4,5}
	fmt.Println(arr[0])
	fmt.Println(arrIn)
	length := len(arrIn)
	fmt.Printf("Length of array is : %d\n",length)
	new_arr := arr[1:3]
	fmt.Println(new_arr)
	fmt.Println("Length:", len(arrIn))
	fmt.Println("Capacity:", cap(arrIn))
	compare(5,4)

	for i:=0; i < len(arrIn); i ++ {
		fmt.Println(arrIn[i])
	}
}
