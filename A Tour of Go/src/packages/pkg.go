package packages // chương trình này lỗi

//  + go chạy từ package main, nếu package chứa hàm main nhưng tên package không phải main thì khi chạy sẽ lỗi
//  + mỗi một package chỉ có một hàm main nếu package có 2 file go mà cả hai file đều có hàm main thì sẽ lỗi

// trong vd này cần sửa 2 chỗ để chương trình chạy
// 1 sửa tên package từ 'packages' thành 'main'  
// 2 sửa package file pkg2 thành main hoặc xóa pkg2 đi  ( ??? )

import (
"fmt"
//"packages"
)

func main(){
	fmt.Println("************ loi do file go co phan package khong phai main *********")
}