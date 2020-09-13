# Golang을 이용한 Scrapper 만들기

---

> Golang의 개념을 잡기 위해 노마드코더 강의를 보고 따라 만들어봤다.
> 주프로젝트는 Scrapper이며 연습하기 위해 서브 프로젝트 Bank, Dictionary, URL Checker, Go Routines도 같이 진행했다.

- 시작하기
  - Golang을 운영체제에 맞게 설치한다.
  - 설치 완료 후 go env GOROOT 명령어를 실행해서 올바른 경로에 잘 설치가 되었는지 확인한다. 맥에 경우 /usr/local/go 가 나오면 정상적으로 설치가 완료된 것이다.
  - go env GOPATH 명령어를 실행해서 경로를 확인한다. 맥에 경우 /Users/사용자명/go 가 나오는데 여기가 작업 디렉토리의 루트가 된다. 폴더에 들어가면 bin, pkg, src 폴더가 있다. 만약에 없다면 만들어준다. src 폴더 아래에서 프로젝트를 시작하면 된다. 프로젝트들은 도메인 주소로 관리가 돤다. 만약에 프로젝트의 깃주소가 github.com/sungjun1313/goScrapper 라면 src/github.com/sungjun1313/goScrapper 이 경로에서 프로젝트를 관리하면 된다.
  - vsCode 툴은 이용하는데 처음에 go 파일을 열면 필요한 모듈을 설치하라고 경고문이 뜬다. 모두 설치한 후 껐다가 다시 실행하면 된다.
  - go run main.go 명령어를 실행하면 main.go 파일이 실행된다.

---

### Bank

- Golang의 struct를 생성하고 get, set method를 만들어보았다.
- method 호출할 때 생길 수 있는 에러를 어떻게 처리하는 지 알아보았다.
- Golang에서 리스트나 오브젝트 등을 프린트하면 자동으로 Golang의 String method가 실행되어 보이는데 이를 덮어씌워서 커스텀화하는 방법을 알아보았다.

---

### Dictionary

- Dictionary를 이용해 map을 다루는 방법을 알아봤다.
- search, add, update, delete method를 만들어 기능을 추가하였다.

---

### URL Checker & Go Routines

- 각각 URL을 체크하여 결과값을 map에 저장하였다.
- 기존 for구문이라면 URL을 순서대로 하나씩 체크해야하지만 goRoutine을 사용하면 URL 여러개를 동시에 체크하여 빠르게 작업할 수 있게 된다.

---

### Scrapper

- goquery 라이브러리를 이용하여 특정 사이트의 필요한 정보를 가져와 csv파일로 저장한다. 나중에는 echo 라이브러리를 이용하여 웹서버를 열어 웹페이지에서 검색해서 필요한 정보를 수집한 뒤 csv파일로 다운받을 수 있게 하였다.
- 처음에는 일반적인 방식으로 코딩하였으며, 나중에는 goRoutine 방식으로 코딩하였다. 처리해야 하는 데이터가 많으면 많을수록 확실히 goRoutine 방식으로 다루는 것이 빠르다는 것을 체감할 수 있었다.

---
