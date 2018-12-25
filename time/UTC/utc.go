package main

import (
	"fmt"
	"time"
)

func DayStartTime(start time.Time) time.Time {
	y, m, d := start.Date()
	loc := time.FixedZone("SH", 8*60*60)
	return time.Date(y, m, d, 0, 0, 0, 0, loc)
}

func Weekth(start, check time.Time) int64 {
	sec := check.Unix() - start.Unix() + 6*86400 - 1
	return 1 + sec/(7*86400)
}

func RoundDownTime(t time.Time) time.Time {
	delta := (t.UTC().Unix() + 8*3600) / 86400
	return time.Unix(delta*86400-8*3600, 0)
}

func RoundUpTime(start time.Time) time.Time {
	y, m, d := start.Date()
	loc := time.FixedZone("SH", 8*60*60)
	return time.Date(y, m, d, 23, 59, 59, 0, loc)
}

func CalcEndTime(start time.Time, weeks int64) time.Time {
	wkd := start.Weekday()
	st := time.Unix(start.Unix()-int64(wkd)*86400+86400, 0)
	endTimeSec := st.Unix() + weeks*7*24*60*60 - 1
	return time.Unix(endTimeSec, 0)

}

func inTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}

func GeneTime(start, end time.Time, sels []int, weeks int64) []time.Time {
	fmt.Println(start, end, sels, weeks)
	result := make([]time.Time, 0)

	mt := start
	var j int
	for mt.Before(end) {

		fmt.Println("[", j, "]", "mt:", mt, Weekth(start, mt))

		j++

		w := int(mt.Weekday())
		for i, sel := range sels {
			fmt.Println("\t\t[", i, "]", w, "?=", sel, "time:", mt)

			if w == sel {
				result = append(result, mt)
				break
			}

		}

		mt = time.Unix(mt.Unix()+86400, 0)
	}

	fmt.Println(start, end, sels, weeks)
	return result

}

func main() {

	//wkd := start.Weekday()
	//mt := time.Unix(start.Unix()-int64(wkd)*86400, 0)

	d1 := time.Now()
	weeks := int64(4)
	endTime := CalcEndTime(d1, weeks)
	startTime := DayStartTime(d1)
	var days []int = []int{1, 6, 0}

	fmt.Println(GeneTime(startTime, CalcEndTime(startTime, weeks), days, weeks))

	fmt.Println("[ ", startTime, endTime, "]", startTime.Unix(), float64(startTime.Unix())/86400)
	fmt.Println("[", RoundUpTime(d1), "]", RoundUpTime(d1).Unix(), float64(RoundUpTime(d1).Unix())/86400)

	// xt1 := (d1.UTC().Unix() + 8*3600) / 86400
	// fmt.Println(time.Unix(xt1*86400-8*3600, 0))

	return

	fmt.Println(d1.Format("2006-01-02 15:04:05 Z07:00:00 MST"))

	fmt.Println(time.Now())
	date := "2009-11-10 23:00:00 +0000 UTC"
	t, err := time.Parse("2006-01-02 15:04:05 -0700 MST", date)
	if err != nil {
		fmt.Println("parse error", err.Error())
	}
	fmt.Println(t.Format(time.ANSIC))

	a1 := time.Now()
	fmt.Println("a1:{time.Now}\t\t\t\t", a1)
	a2 := a1.Format("2006-01-02 15:04:05")
	fmt.Println("a2:{a1.Format(\"\")} \t\t\t", a2)
	a3, _ := time.Parse("2006-01-02 15:04:05", a2)
	//time.ParseInLocation()
	fmt.Println("a3: {time.Parse(\"\", a2)}\t", a3)
	a4 := time.Now()
	fmt.Println("a4:{ time.Now()}\t\t\t", a4)
	a5 := a4.UTC()
	fmt.Println("a5: {a4.UTC()}\t\t\t\t", a5)
	fmt.Println("a5.Sub(a3):", a5.Sub(a3))
	fmt.Println("a4.Sub(a1):", a4.Sub(a1))
	fmt.Println("-----------------------------------")
	b1 := time.Now().UTC()
	fmt.Println("Now()\t\t\t\t\t", time.Now())
	fmt.Println("b1 {time.Now().UTC()}:\t", b1)
	b2 := b1.Format("2006-01-02 15:04:05")
	fmt.Println("b2:{b1.Format(\"\")}\t\t", b2)
	b3, _ := time.Parse("2006-01-02 15:04:05", b2)
	fmt.Println("b3:{time.Parse(\"\", b2)}\t", b3)
	b4 := time.Now().UTC()
	fmt.Println("b4{time.Now().UTC()}:\t", b4)
	fmt.Println("b4.Sub(b1):", b4.Sub(b1))
	fmt.Println("b4.Sub(b3):", b4.Sub(b1))
}

/*
美国时间一般被认为是美国本土的时间。

美国本土横跨西五区至西八区，共四个时区，每个时区对应一个标准时间。

从东向西分别为

东部时间(EST)（西五区时间）、

中部时间(CST)（西六区时间）、

山地时间(MST)（西七区时间）、

太平洋时间（西部时间）(PST)（西八区时间）。

[阿拉斯加时间(AKST)（西九区时间）和夏威夷时间(HST)（西十区时间）]按照“东早西晚”的规律，各递减一小时。美国从每年3月的第二个星期日至11月的第一个星期日采用夏令时，夏令时比正常时间快一小时。

美国的时区界限并不完全按照经线划分，基本上照顾了各州的自然边界。不同的时区覆盖的州市大小、多少不同。东部时间(EST)包括大西洋沿岸及近大陆的19个州和华盛顿特区，代表城市华盛顿、纽约。

中部时间(CST)代表城市芝加哥、新奥尔良。

山地时间(MST)代表城市盐湖城、丹佛，

太平洋时间(PST)包括太平洋沿岸的4个州，代表城市旧金山、洛杉矶、西雅图，

阿拉斯加时间(AKST)只限于阿拉斯加。

夏威夷时间(HST)只限于夏威夷。
美国时间与中国时间时差(秋冬季)

太平洋时区、山地时区、中部时区、东部时区
太平洋时区、山地时区、中部时区、东部时区 [1]
太平洋时区：代表城市洛杉矶，与北京相差16小时；
山地时区：代表城市盐湖城，与北京相差15小时；
中部时区：代表城市芝加哥，与北京相差14小时；
东部时区：代表城市纽约、华盛顿，与北京相差13小时；
夏威夷时区：代表城市：火奴鲁鲁，与北京相差18小时；
阿拉斯加时区：代表城市：费尔班克斯，与北京相差17小时。

例如：

1.北京时间2013年1月1日早上8点减去16小时就是美国太平洋时间2012年12月31日下午16点。

2.北京时间20：00，美国太平洋时间就是04：00。

3.在夏令时，时差少一小时。如在美国夏令时。北京时间20：00，美国太平洋时间就是05：00。



美国时间与中国时间时差(夏令时)
太平洋时区：代表城市洛杉矶，与北京相差15小时；
山地时区：代表城市盐湖城，与北京相差14小时；
中部时区：代表城市芝加哥，与北京相差13小时；
东部时区：代表城市纽约、华盛顿，与北京相差12小时；
夏威夷时区：代表城市：火奴鲁鲁，与北京相差18小时（夏威夷没有夏时制）；
阿拉斯加时区：代表城市：费尔班克斯，与北京相差16小时


*/
