/*




Классическая задача на указатели. 

Вам надо написать функцию changeStrings, которая должна принимать указатели на две строки и затем менять местами их значения.

Выводить или возвращать ничего не нужно.

Sample Input:

Stepik Hello

Sample Output:

Hello Stepik


*/

// WRONG
// func changeStrings(pointer1 *int, pointer2 *int){
// 	var temp1 int;
// 	var temp2 int;

// 	temp1 = &pointer1
// 	temp2 = &pointer2
// 	pointer1 = *temp2
// 	pointer2 = *temp1
// }

// // нужно написать только функцию

// GOOD
func changeStrings(pointer1 *string, pointer2 *string){
	
    var temp string;
    var temp2 string;
    
    temp1 = *pointer2
    temp2 = *pointer1
    
    *pointer1 = temp1
    *pointer2 = temp2
    
}


