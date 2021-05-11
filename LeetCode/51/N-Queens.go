package main

import "fmt"
var res [][]string
func solveNQueens(n int) [][]string {
    res=[][]string{}    //用来保存最终结果
    queens:=make([]int,n)   //用来保存第n行的皇后放在第int的位置
    for i:=0;i<n;i++{
        queens[i]=-1    //每行皇后都放在-1上，表示没有放置
    }
    col,dia1,dia2:=make(map[int]bool),make(map[int]bool),make(map[int]bool)
    backtrace(n,0,queens,col,dia1,dia2)
    return res
}
//回溯五要素之一，递归函数，其传入参数中一定会有一个代表“层“的参数，对于此题为row,也就是棋盘的行
func backtrace(n,row int,queens []int,col,dia1,dia2 map[int]bool){
    //回溯五要素之一，结束条件，即满足题目要求的情况
    if row==n{
        ss:=formet(n,queens)
        res=append(res,ss)
        return
    }
    //回溯五要素之一，每层的循环，对于此题每层对列进行循环，所以最终是对全部的格子进行了遍历
    for i:=0;i<n;i++{
        //对于此题，怎么放,放的条件必须得达到，不属于回溯范畴
        if col[i]{  //列有了，不能放
            continue
        }
        if dia1[row-i]{//正对角线有了，不能放
            continue
        }
        if dia2[row+i]{//反対角线有了，不能放
            continue
        }
        //此位置可以放置
        //回溯五要素之一，“吞“，达成了放置的条件后，就要实实在在的放上去
        col[i]=true //此列有人了
        dia1[row-i]=true    //此对角线有人了
        dia2[row+i]=true    //此反对角线有人了
        queens[row]=i   //第row行的皇后放在了i位置
        backtrace(n, row+1, queens, col, dia1, dia2)    //递归。注意要进一层了
        //回溯五要素之一，“吐“，一次一次的进层，没有满足条件的话就要吐出来
        queens[row]=-1  //第row行不放皇后了（不在第i处放，继续下一个i++处再测试去了）
        delete(dia2,row+i)  //对角线，反对角线，列的放置都要清空
        delete(dia1,row-i)
        delete(col,i)
    }
}
func formet(n int,queens []int)[]string{    //仅做格式化作用
    ret:=make([]string,0)
    for i:=0;i<n;i++{
        a:=make([]byte,n)
        for j:=0;j<n;j++{
            a[j]='.'
        }
        a[queens[i]]='Q'
        ret=append(ret,string(a))
    }
    return ret
}
func main() {
	a:=solveNQueens(4)
        fmt.Println(a)
}
