package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	count := 150
	for i := 0; i < count; i++ { //设置顶点的颜色，count 为顶点个数
		colors := []string{
			"0,0,1,",
			"0,1,0,",
			"1,0,0,",
			"0,0,0.5,",
			"0,0.5,0,",
			"0.5,0,0,",
			"0.5,0.5,1,",
			"0.5,1,0.5,",
			"1,0.5,0.5,",
			"1,1,0.5,",
			"1,0.5,1,",
			"0.5,1,1,",
		}
		color := colors[rand.Intn(12)]
		fmt.Println(color)
	}
}

/*
c++ 设置索引的，穷举每一个三角形的所有面
#include <iostream>
#include <algorithm>
#include <vector>
using namespace std;
const int N = 1e6+10;

void solve()
{
	for(int i = 0;i < 300;i +=3)
	{
		vector<int> cfd(3);
		cfd[0] = i;
		cfd[1] = i+1;
		cfd[2] = i+2;
		do{
			for(auto x:cfd)
			{
				cout << x << ",";
			}
			cout << endl;
		}while(next_permutation(cfd.begin(),cfd.end()));
	}
}

signed main()
{

    ios::sync_with_stdio(false);
    cin.tie(0);
    cout.tie(0);

    // caseT
    solve();

    return 0;
}
*/
