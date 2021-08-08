package main
import (
    "fmt"
    // "io/ioutil"
    "net/http"
    // "os"
    "golang.org/x/net/html"
)


// Copied from gopl.io/ch5/outline2.
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
    if pre != nil {
        pre(n)
    }
    for c := n.FirstChild; c != nil; c = c.NextSibling {
        forEachNode(c, pre, post)
    }
    if post != nil {
        post(n)
    }
}

func main()  {
    var url string = "https://www.op.gg/summoner/matches/ajax/detail/gold/gameId=5351809074&summonerId=3458580&moreLoad=1&"
    
    // idValid []int

    resp, err := http.Get(url)
    if err != nil {
        return 
    }
    fmt.Println("111")
    if resp.StatusCode != http.StatusOK {
        resp.Body.Close()
        return 
    }

    doc, err := html.Parse(resp.Body)
    resp.Body.Close()
    if err != nil {
        return 
    }
    fmt.Println("222")
    // var links []string
    visitNode := func(n *html.Node) {
        if n.Type == html.ElementNode && n.Data == "script" {
            // fmt.Println("ok")
            fmt.Println(n.FirstChild.Data)
            // links = append(links, link.String())
        }
    }
    forEachNode(doc, visitNode, nil)

}

// 网址 https://game.gtimg.cn/images/lol/act/img/js/hero/2.js
// 分类：
// 
// 2、英雄形象

// 1、战队信息

// 排名
// 英雄
// 位置
// 胜率
// 登场率
// 被禁率
// 使用次数
// 击杀
// 死亡
// 助攻
// 造成英雄伤害
// 承受英雄伤害
// 总治疗量
// 补刀数
// 击杀敌方野怪
// 击杀己方野怪
// 金钱
