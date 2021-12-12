## MSA-ContainerApp-GO-Itemapi
Golang으로 작성된 간단한 REST API 프로그램입니다.

### 연관 프로젝트
[msa-containerapp-java-employeeapi](https://github.com/daintree-henry/msa-containerapp-java-employeeapi)  
[msa-containerapp-nodejs-webapp](https://github.com/daintree-henry/msa-containerapp-nodejs-webapp)

### API

### 소스 코드 다운로드
```shell
git clone https://github.com/daintree-henry/msa-containerapp-go-itemapi
```

### Version 1 - 컨테이너 엔진 한 대에서 배포
- image registry url (docker hub):
- 주요 file 구조:
```shell
  $ tree
    .
    |-- Dockerfile
    |-- LICENSE
    |-- README.md
    |-- controllers
    |   `-- controller.go
    |-- domain
    |   `-- item.go
    |-- go.mod
    |-- go.sum
    |-- k8s
    |   |-- deployment.yaml
    |   `-- service.yaml
    `-- main.go
```

### 실습
- 요구사항: 도커 컨테이너 엔진이 설치된 Linux OS에서 작업

#### 실습 1
도커 허브에 배포되어 있는 이미지를 사용하여 컨테이너를 실행합니다.
```shell
#1. 이미지 다운로드


#1. 컨테이너 배포


#2. API 테스트

```

#### 실습 2
직접 소스코드의 도커파일을 통해 이미지를 생성하고 자신의 레파지토리에 배포합니다.
배포한 이미지를 사용하여 컨테이너를 실행합니다.
```shell
#1. 실습 소스코드 다운로드
git clone 

#2. 도커 이미지 빌드
cd 
docker build -t .

#3. 도커 이미지 푸시
docker push 

#4. 로컬 레지스트리의 이미지 삭제
docker image rm 

#5. 컨테이너 배포
docker run 

#6. API 테스트

```

### Version 2 - 컨테이너 오케스트레이션 도구를 활용한 배포

