package main

import "fmt"

/*
给定一个 n × n 的二维矩阵表示一个图像。

将图像顺时针旋转 90 度。

说明：

你必须在原地旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要使用另一个矩阵来旋转图像。

示例 1:

给定 matrix =
[
  [1,2,3],
  [4,5,6],
  [7,8,9]
],

原地旋转输入矩阵，使其变为:
[
  [7,4,1],
  [8,5,2],
  [9,6,3]
]
示例 2:

给定 matrix =
[
  [ 5, 1, 9,11],
  [ 2, 4, 8,10],
  [13, 3, 6, 7],
  [15,14,12,16]
],

原地旋转输入矩阵，使其变为:
[
  [15,13, 2, 5],
  [14, 3, 4, 1],
  [12, 6, 8, 9],
  [16, 7,10,11]
]

*/

/*
 那么顺时针方向脚标变化的规律是：

[x][y] --> [y][n - 1 - x]

可以得到逆时针方向脚标变化的规律是：

[x][y] --> [n - 1 - y][x]

那么可以得到变化的代码：
            for (int j = 0; j < 3; j++) {
                x2 = n - 1 - y1;
                y2 = x1;

                matrix[x1][y1] = matrix[x1][y1] ^ matrix[x2][y2];
                matrix[x2][y2] = matrix[x1][y1] ^ matrix[x2][y2];
                matrix[x1][y1] = matrix[x1][y1] ^ matrix[x2][y2];

                x1 = x2;
                y1 = y2;
            }
还有一种解法，首先以从对角线为轴翻转，然后再以x轴中线上下翻转即可得到结果，如下图所示(其中蓝色数字表示翻转轴)：
*/
func rotate(matrix [][]int) {
	length := len(matrix)
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i; j++ {
			matrix[i][j], matrix[length-1-j][length-1-i] = matrix[length-1-j][length-1-i], matrix[i][j]
		}
	}
	for i := 0; i < length/2; i++ {
		for j := 0; j < length; j++ {
			matrix[i][j], matrix[length-1-i][j] = matrix[length-1-i][j], matrix[i][j]
		}
	}

}

func main() {
	data := [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	rotate(data)
	fmt.Println(data)
}
