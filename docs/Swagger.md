# Swagger
API 문서화 자동화 Tool

## 개요  
---
API 문서를 코드 개발과 분리하여 처리하는 것이 아닌 
코드 개발 단계에서 미리 작성하고, 자동화된 툴을 이용해
외부에 공개하는 방식 

## 패키지 다운로드  
---
```bash
$ go get github.com/swaggo/swag/cmd/swag 
$ go get github.com/swaggo/echo-swagger
```

## 서버에 스웨커 문서 페이지 Path 추가  
---
Echo Framework를 사용한다면 다음과 같이 GET + 사용자 정의 PATH를 정의하고, 
echoSwagger.WarpHandler를 추가하여 문서를 외부에서 확인할 수 있도록 제공 가능하다.  
```go
// swagger setting
e.GET("/swagger/*", echoSwagger.WrapHandler)
```

## 작성 방법  
---
- 대문 작성  
    Swagger 문서의 시작 부분을 main 함수에 추가한다. 
- Option list
    Link: [main.go][generallink]
    [generallink]: https://github.com/swaggo/swag/blob/master/example/celler/main.go
- Sample 
    main.go
    ```go
    // @title Platform-sample Swagger API
    // @version 1.0.0
    // @host localhost:8080
    // @contact.name API Support
    // @contact.url http://www.swagger.io/support
    // @contact.email support@swagger.io
    func main(){
    ...
    }
    ```
- Path 별 작성  
    handler 함수 위치에 작성
    ```go
    // @Summary Get user
    // @Description Get user's info
    // @Accept json
    // @Produce json
    // @Param id path string true "id of the user"
    // @Success 200 {object} model.User
    // @Router /users/{id} [get]
    func (userController *UserController) GetUser(c echo.Context) error {
        id, err := strconv.Atoi(c.Param("id"))
        if err != nil {
            return c.JSON(http.StatusBadRequest, err)
        }

        user, err := userController.UserService.GetUser(id)
        if err != nil {
            return c.JSON(http.StatusInternalServerError, err)
        }
        return c.JSON(http.StatusOK, user)
    }
    ```

## 문서 생성  
```bash
$ swag init
```
그러나 데이터 모델에 패키지 외부 참조가 있다면 다음과 같은 에러가 발생할 수 있다.  
```bash
 ParseComment error :cannot find type definition: ...
```
이런 경우 다음과 같이 명령을 입력한다.  
```bash
$ swag init --parseDependency --parseInternal
```

## 문서 확인  
프로그램을 실행하고 해당 
```
http://{서버 접속 정보}/swagger/index.html
```
