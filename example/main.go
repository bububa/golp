package main

import (
	"fmt"
	"github.com/bububa/golp"
)

var (
	maxCost  float64 = 50000
	minSales float64 = 100000
	cols             = []string{"淘宝首页焦点图3", "淘宝首页焦点图4", "淘宝首页1屏banner", "淘宝首页2屏右侧大图", "淘宝首页3屏通栏", "淘宝首页3屏小图", "淘宝首页焦点图2", "listing页面右侧banner1", "listing页面右侧banner2", "旺旺聊天窗口右上角banner", "旺旺聊天窗口底部文字链", "淘宝首页3屏banner", "旺旺每日弹窗焦点图1", "旺旺每日弹窗焦点图1", "旺旺每日弹窗小图2", "旺旺每日弹窗焦点图2", "收藏夹底部通栏轮播2", "收藏夹底部通栏轮播1", "旺旺每日弹窗焦点图3", "淘金币新焦点图1", "淘金币新焦点图2", "旺旺每日弹窗焦点图3"}
	traffics         = []float64{17107.485076, 16736.42942, 12055.959664, 6511.818442, 3055.455805, 1895.28462, 13998.978238, 1309.735094, 1305.598224, 915.052257, 551.022855, 666.2381, 1062.1194, 251.428164, 313.38054, 1136.2219, 574.412652, 502.872093, 618.701904, 1827.706333, 1285.98213, 89.618364}
	cost             = []float64{16631.0740992, 15267.9427096, 9429.4827372, 5881.9868222, 2191.499336, 1042.406541, 12258.172654, 1040.8947326, 1023.8638704, 649.2513633, 400.4099413, 439.717146, 981.3983256, 188.571123, 231.9015996, 1022.59971, 499.5649428, 488.9997594, 528.474543, 2242.4915008, 1559.831562, 56.2151556}
	sales            = []float64{43573.414139904, 37559.139065616, 25836.782699928, 9646.458388408, 3835.123838, 1855.48364298, 33219.64789234, 3892.946299924, 3911.159984928, 5362.816260858, 180.184473585, 809.07954864, 4062.989067984, 1614.16881288, 1486.489253436, 4029.0428574, 934.186443036, 1378.979321508, 2029.34224512, 8992.390918208, 5147.4441546, 354.15548028}
)

func main() {
	numCols := len(traffics)
	lp := golp.NewLP(0, numCols)
	lp.AddConstraint(cost, golp.LE, maxCost)
	lp.AddConstraint(sales, golp.GE, minSales)
	lp.SetObjFn(traffics)
	for col := 0; col < lp.NumCols(); col++ {
		lp.SetBinary(col, true)
		lp.SetColName(col, cols[col])
		fmt.Printf("Col:%d, name:%s\n", col, cols[col])
	}
	lp.SetMaximize()
	lp.Solve()
	fmt.Printf("Objective value: %v\n", lp.Objective())
	fmt.Printf("Number Rows: %d\n", lp.NumRows())
	fmt.Printf("Number Cols: %d\n", lp.NumCols())
	vars := lp.Variables()
	fmt.Printf("Variable values:\n")
	var totalCost, totalSales float64
	for i := 0; i < lp.NumCols(); i++ {
		fmt.Printf("%s = %v\n", lp.ColName(i), vars[i])
		if vars[i] == 1 {
			totalCost += cost[i]
			totalSales += sales[i]
		}
	}
	fmt.Printf("Total Cost: %v\n", totalCost)
	fmt.Printf("Total Sales: %v\n", totalSales)
}
