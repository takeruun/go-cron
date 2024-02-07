# go-cron
golang で cron 実行

## 使用ライブラリ

https://github.com/robfig/cron

4年前だけど、動くからいいかな

## 使い方
見ての通り

```go
// 文的なもので指定
c.AddFunc("@every 1s", func() { log.Println("every 1s") })
c.AddFunc("@every 5s", SampleFunction)

// cron 形式で指定
c.AddFunc("* * * * *", func() { log.Println("every minute") })

c.Start()
```