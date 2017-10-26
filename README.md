# Gothumb
bithumb용 go언어 API 클라이언트, goinone 과 유사합니다.
** 주의 ** bithumb은 응답이 유연하게 설계되어 있습니다. api 문서를 참조해주세요. 

## API
* Info
  * balance
  
    ```golang
    // token string, secret string 
    // balance BalanceResponse
      client := NewClient(token, secret)
      balance, err := client.Balance()
    ```
