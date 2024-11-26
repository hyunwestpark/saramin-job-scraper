# Saramin Job Scraper 📝🤖

![프로젝트 이미지](./home.png)

이 프로젝트는 사용자가 특정 직무 관련 용어를 웹사이트에 입력했을 때, 사람인(Saramin) 사이트에서 관련 채용 공고를 스크래핑하여 CSV 파일로 저장할 수 있게 만든 프로젝트입니다. 이 프로젝트는 **노마드 코더의 [쉽고 빠른 Go 시작하기](https://nomadcoders.co/go-for-beginners)** 수업을 수강하며 학습 목적으로 생성되었습니다. 📚💻

## 주요 기능 ✨

- 사용자가 특정 직무 용어(예: "프론트엔드 개발자")를 입력하면, 사람인에서 해당 용어와 관련된 채용 공고를 검색합니다.
- 검색 결과로 나온 채용 공고의 회사명, 직무명, 위치, 고유 ID를 스크래핑합니다.
- 모든 스크랩한 데이터를 CSV 파일로 저장하여, 사용자가 쉽게 열람하고 분석할 수 있도록 합니다.

## 기술 스택 🛠️

- **언어**: ![Go Badge](https://img.shields.io/badge/Go-1.19-blue?logo=go&logoColor=white)
- **패키지**:
  - `net/http`: HTTP 요청을 보내고 응답을 받기 위해 사용.
  - `encoding/csv`: CSV 파일 작성에 사용.
  - `github.com/PuerkitoBio/goquery`: 웹 스크래핑을 위해 사용된 Go용 HTML 파서.
  - `github.com/labstack/echo/v4`: 웹 서버를 쉽게 구축하기 위해 사용된 Go용 프레임워크.

## 설치 및 실행 방법 🚀

1. **Go 설치**: 이 프로젝트는 Go 언어로 작성되었기 때문에, Go가 설치되어 있어야 합니다. [Go 설치 가이드](https://go.dev/doc/install)를 참고하세요.

2. **클론하기**:

   ```sh
   git clone https://github.com/hyunwestpark/saramin-job-scraper.git
   cd saramin-job-scraper
   ```

3. **프로젝트 실행**:
   서버를 실행하여 웹 페이지를 통해 스크래핑 기능을 이용할 수 있습니다.

   ```sh
   go run main.go
   ```

4. **웹 인터페이스 사용**:

   - 브라우저에서 `http://localhost:1323`에 접속합니다.
   - 직무 용어를 입력하고 'Search' 버튼을 누르면 해당 데이터를 스크래핑하여 CSV 파일로 다운로드할 수 있습니다.

5. **결과 확인**:
   프로그램이 실행되면 `<입력한 용어>_jobs.csv` 파일이 다운로드됩니다. 이 파일에는 스크래핑된 채용 공고의 상세 정보가 저장되어 있습니다.

## 프로젝트 구조 📁

```
├── main.go                // 주요 실행 파일 (Echo 서버 및 라우팅 구현)
├── scrapper/              // 웹 스크래핑 관련 기능 구현 폴더
│   └── scrapper.go        // 스크래핑 로직 구현 파일
├── go.mod                 // Go 모듈 설정 파일
├── home.html              // 사용자 입력을 받는 웹 페이지 파일
├── LICENSE.txt            // 라이센스 파일
└── README.md              // 프로젝트 설명 파일
```

## 참고 자료 📖

- **노마드 코더: [쉽고 빠른 Go 시작하기](https://nomadcoders.co/go-for-beginners)**: 이 프로젝트는 해당 강의를 통해 배우며 개발되었습니다.

## 주의사항 ⚠️

- 이 프로젝트는 **학습 목적으로만 사용**되었으며, Saramin 웹사이트의 데이터 이용 약관을 준수했습니다.
- Saramin 웹사이트의 `robots.txt` 파일을 확인하여 허용된 경로만 스크래핑해주세요.

## 라이센스 📄

이 프로젝트는 MIT 라이센스를 따르며, 관련 라이브러리의 라이센스 역시 포함하고 있습니다. 자세한 내용은 `LICENSE` 파일을 참조하십시오.
