# weather_cli
天気を取得してくれるCLIツール

## 設定
### 天気APIのアクセストークンの取得
[ここから](https://openweathermap.org/)取得
環境変数の追加
```
export WEATHER_API="hogehoge"
```
### go getする
```
go get https://github.com/mytheta/weather
```

## 使い方
```
weather 都市名
```
で，都市の天気，気温，湿度が取得できる．

## example
```
☁  ~ weather Koganei
Koganeiの天気はClouds, 気温:19.2200度, 湿度:80%
```