# logrus-hook-caller
[![Go Report Card](https://goreportcard.com/badge/github.com/zhl11b/logrus-hook-caller)](https://goreportcard.com/report/github.com/zhl11b/logrus-hook-caller)

make logrus display linenum and filename no matter use logrus.Debugf("xx") or logrus.Withfield("xx","xx").Debugf("xx")

cation!! because Debugf or Debugln's call deepth are diferent with Debug.so you have to use Debugf instead Debug, then you can't get right caller.
注意！！因为Debuf,Debugln和Debug的调用深度不一样。所以需要使用Debugf取代Debug，那么就可以正确的显示行了。