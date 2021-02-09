# go-memo

## 커멘드memo
go build
gofmt -s -w .
GOPATH=~/workspace

## 로컬에서 package import하기
- 디렉토리 단위가 된다.
- 디렉토리명과 package 선언명이 일치해야한다.
- package 선언한 부분이 import될때 쓰인다.
- 디렉토리 내부의 파일이름은 아무거나 괜찮다.

## go mod에 관해서
- go.mod에서 go패키지 의존성 관리해준다.
- 시스템 전체에 영향을 끼치는 GOPATH사용하는 의존성 관리가 아니다.
- go mod init은 패키지명도 같이 적어줘야 한다.
  - `go mod init jiho_local` 같이
  - go mod를 사용중이면 로컬에서 만들 패키지 import도 go.mod의 모듈명을 추가해야 한다.
    - `import "jiho_local/basic"` 같이

## go test에 색갈 입히기
- 인스톨 : `go get -u github.com/rakyll/gotest`
- 실행 : `gotest -v ./testing/`

## unit test에 관해서
- `t.Error`, `t.Fatal` 가 실행되지 않은 테스트 케이스는 성공으로 간주한다.
- `t.Error` 는 테스트 케이스의 실패를 의미한다. 그래도 테스트 케이스를 끝까지 수행한다.
- `t.Fatal` 는 테스트 케이스의 실패를 의미한다. 테스트 케이스를 바로 중지한다.