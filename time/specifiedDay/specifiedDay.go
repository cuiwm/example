package main

import (
	"fmt"
	"time"
)

//func main() {
//	t := time.Now()
//	fmt.Println(t)
//	fmt.Errorf("time:%s", t.UTC().Format())
//
//	fmt.Println(t.UTC().Format(time.UnixDate))
//
//	fmt.Println(t.Unix())
//
//	timestamp := strconv.FormatInt(t.UTC().UnixNano(), 10)
//	fmt.Println(timestamp)
//	timestamp = timestamp[:10]
//	fmt.Println(timestamp)
//}

func main() {
	//2018/9/23 0:0:0 Beijing
	var tutc int64 = 1537632000
	tdays := tutc / (24 * 3600)
	tutcsec := tutc % (24 * 3600)
	tbj := time.Unix(1537632000, 0)
	fmt.Println("2018/9/23 0:0:0 Beijing indeed:", tbj.Format("2006-01-02 15:04:05"),
		"day:", tdays, "sec", tutcsec)
	//time.Time代表一个纳秒精度的时间点
	var t time.Time
	//返回当前时间
	t = time.Now()
	fmt.Printf("%v\n", t)
	//反回所在时区
	fmt.Printf("%v\n", t.Location())
	//返回UTC时间和UTC时区
	fmt.Printf("%v %v\n", t.UTC(), t.UTC().Location())
	//同上，In()返回指定时区的时间
	fmt.Printf("%v %v\n", t.In(time.UTC), t.In(time.UTC).Location())
	//返回当地时区的时间
	fmt.Printf("%v %v\n", t.Local(), t.Local().Location())

	//根据时间戳返回本地时间
	//参数分别表示秒数和纳秒数
	t2 := time.Unix(1487780010, 0)
	fmt.Println(t2)

	//根据指定时间返回time.Time
	//分别指定年，月，日，时，分，秒，纳秒，时区
	t3 := time.Date(2017, time.Month(5), 26, 15, 30, 20, 0, t.Location())
	fmt.Println(t3)

	//格式化输出时间
	t4 := time.Now()
	fmt.Println(t4.Format("2006-01-02 15:04:05"))

	//获取时间信息
	t5 := time.Now()
	//返回日期
	fmt.Println(t5.Date())
	//返回年
	fmt.Println(t5.Year())
	//返回月
	fmt.Println(t5.Month())
	//返回日
	fmt.Println(t5.Day())
	//返回星期
	fmt.Println(t5.Weekday())
	//返回ISO 9601标准下的年份和星期编号
	fmt.Println(t5.ISOWeek())
	//返回时分秒
	fmt.Println(t5.Clock())
	//返回小时
	fmt.Println(t5.Hour())
	//返回分钟
	fmt.Println(t5.Minute())
	//返回秒
	fmt.Println(t5.Second())
	//返回纳秒
	fmt.Println(t5.Nanosecond())
	//返回一年中对应的天
	fmt.Println(t5.YearDay())
	//返回时区
	fmt.Println(t5.Location())
	//返回时区的规范名,时区相对于UTC的时间偏移量(秒)
	fmt.Println(t5.Zone())
	//返回时间戳
	fmt.Println(t5.Unix())
	//返回纳秒时间戳
	fmt.Println(t5.UnixNano())

	//时间的比较与计算
	t6 := time.Now()
	//是否零时时间
	fmt.Println(t6.IsZero())
	//t6时间在t5时间之后，返回真
	fmt.Println(t6.After(t5))
	//t5时间在t6时间之前，返回真
	fmt.Println(t5.Before(t6))
	//时间是否相同
	fmt.Println(t6.Equal(t6))
	//返回t6加上纳秒的时间
	fmt.Println(t6.Add(10000))
	//返回两个时间之差的纳秒数
	fmt.Println(t6.Sub(t5))
	//返回t6加1年，1月，1天的时间
	fmt.Println(t6.AddDate(1, 1, 1))

	//时间的序列化
	t7 := time.Now()
	//序列化二进制
	bin, _ := t7.MarshalBinary()
	//反序列化二进制
	t7.UnmarshalBinary(bin)
	fmt.Println(t7)
	//序列化json
	json, _ := t7.MarshalJSON()
	fmt.Println(string(json))
	//反序列化json
	t7.UnmarshalJSON(json)
	fmt.Println(t7)
	//序列化文本
	txt, _ := t7.MarshalText()
	fmt.Println(string(txt))
	//反序列化文本
	t7.UnmarshalText(txt)
	fmt.Println(t7)
	//gob编码
	gob, _ := t7.GobEncode()
	t7.GobDecode(gob)
	fmt.Println(t7)

	//时间段time.Duration
	dur := time.Duration(6666666600000)
	//返回字符串表示
	fmt.Println(dur.String())
	//返回小时表示
	fmt.Println(dur.Hours())
	//返回分钟表示
	fmt.Println(dur.Minutes())
	//返回秒表示
	fmt.Println(dur.Seconds())
	//返回纳秒表示
	fmt.Println(dur.Nanoseconds())

	//时区time.Location
	//返回时区名
	fmt.Println(time.Local.String())

	//通过地点名和时间偏移量返回时区
	fmt.Println(time.FixedZone("Shanghai", 800))

	//通过给定时区名称，返回时区
	loc, _ := time.LoadLocation("Asia/Shanghai")
	fmt.Println(loc)

	//阻塞当前进程3秒
	time.Sleep(time.Second * 3)

	//定时器time.Timer
	//创建一个1秒后触发定时器
	timer1 := time.NewTimer(time.Second * 1)
	<-timer1.C
	fmt.Println("timer1 end")

	//1秒后运行函数
	time.AfterFunc(time.Second*1, func() {
		fmt.Println("wait 1 second")
	})
	time.Sleep(time.Second * 3)

	//打点器time.Ticker
	//创建一个打点器，在固定1秒内重复执行
	ticker := time.NewTicker(time.Second)
	num := 1
	for {
		if num > 5 {
			//大于5次关闭打点器
			ticker.Stop()
			break
		}
		//否则从打点器中获取chan
		select {
		case <-ticker.C:
			num++
			fmt.Println("1 second...")
		}
	}
}
