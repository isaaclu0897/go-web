3-實戰GoLang 建立Web 應用
===

### 其他練習

練習：
* 實現一個可登入頁面，成功登入後，跳轉後臺
* 實現一個畫面

> 對於網站設計實在是沒什麼概念，所以找了一個範本網站來做參考。所謂程式開發...從複製、貼上開始。所以來參考一下吧～[網站設計範本](https://zh.wix.com/website/templates/html/portfolio-cv)

#### SA

```mermaid
graph LR;

%% defined node
page_Home(Home頁面);
page_login(登入頁面);
page_admin(後臺頁面);

act_login([登入]);
act_browser([訪客模式<br>瀏覽其他頁面]);

data_usr&pwd[/賬號&密碼/];

%% defined flow
page_Home --> act_browser;

page_Home -.- page_login --> data_usr&pwd --> act_login --> page_admin;

%% defined subgraph
subgraph 圖例說明;
    a(頁面);
    b([動作]);
    c[/資料/];
    
    a -.- b;
    b -.- c;
end;
```

###### tags: `Go` `GoLang` `Web` 
