# Gothumb
bithumb용 go언어 API 클라이언트.

**유사제품: goinone**

**주의!** bithumb은 응답이 유연하게 설계되어 있습니다. api 문서를 참조해주세요. 

## API
* Info
  * balance
  
    ```golang
    // token string, secret string 
    // resp *APIResponse
      client := NewClient(token, secret)
      resp, err := client.Balance()
    ```
