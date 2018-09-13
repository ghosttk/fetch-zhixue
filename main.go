package main
import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() { 
    client := &http.Client{}
    url := "http://www.zhixue.com/paperfresh/api/question/show/course/getTopics?_=1536654169486&pageIndex=2&chapterId=19130501-001_04&pressCode=01&pressName=%E9%AB%98%E4%B8%AD%E7%94%9F%E7%89%A9%E4%BA%BA%E6%95%99%E7%89%88&paperId=ea11e891-d0c8-49ae-88cf-4d2b39cd01c5&level=0&sectionCode=&difficultyCode=&paperTypeCode=&areas=&sortField=default&sortDirection=true&keyWord=&chapterType=1&year="
    reqest, err := http.NewRequest("GET", url, nil)
    if err != nil {
        fmt.Println(err)
        panic(err)
    }

    reqest.Header.Add("Cookie", "JSESSIONID=95327ABA948ADD6DAC458B175D9933FB; __DAYU_PP=6RQrny7jaZjIan27Vaur37031128ffe8; tlsysSessionId=e1514a5f-86d1-4d52-8244-c2d1c8f48b61; 6e12c6a9-fda1-41b3-82ec-cc496762788d=webim-visitor-JE7Y88KPFXGHRYH634GF; loginUserName=zxt450698; HEADER_ROLE_SID=0.9528420846167329; OPENPAPERID=ea11e891-d0c8-49ae-88cf-4d2b39cd01c5; CartTotalCount=0")
    reqest.Header.Add("Accept", "application/json, text/plain, */*")
    reqest.Header.Add("Accept-Encoding", "gzip, deflate")
    reqest.Header.Add("Accept-Language", "zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7")
    reqest.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Ubuntu Chromium/68.0.3440.106 Chrome/68.0.3440.106 Safari/537.36" )
    reqest.Header.Add("Referer", "http://www.zhixue.com/paperfresh/dist/")
    reqest.Header.Add("Connection","keep-alive")

    //处理返回结果
    response, errDo:= client.Do(reqest)
    if errDo != nil {
        panic(errDo)
    }
    defer response.Body.Close()
    var by []byte
    by,_ = ioutil.ReadAll(response.Body)
    fmt.Println(string(by))
}
