# GO로 만들어보고 싶은 것
1. 신용카드 BIN번호(1~6번째 자리까지 숫자)로 카드사 인식하는 프로그램
2. LUHN 알고리즘으로 신용카드 번호 Validation 확인하는 프로그램
3. 슬랙봇으로 웹이랑 연결해서 값 주고받는거(Jira나 Confluence에서 댓글달리면 슬랙에 알림오고 거기에 댓댓글 달 수 있는 이런 동작가능하게?)
4. 계속 업데이트 중..

# GOPATH와 GOROOT

## GOPATH
과거 GO는 워크스페이스를 지정하고 해당 워크스페이스 하위에 코드가 존재해야 했는데, 이 워크스페이스가 GOPATH이고 기본 경로는 C:\Users\<username>\go\로 되어있다.
코드는 GOPATH\src\에 위치해야 했는데, 이런 제약사항을 자유롭게 해주기 위해서 GO MODULE이 생겼다. 보통 GOPATH에는 GO 소스코드 및 외부 패키지가 저장된다고 한다.

## GOROOT
GOROOT는 GO의 컴파일러와 기본 모듈들이 있는 곳으로 GO가 설치되어있는 경로이다.

# 컴퓨터 원리
컴퓨터는 0과 1밖에 모름  
트랜지스터는 연산을 수행하는 가장 기본이 되는 소자  
트랜지스터는 성질이 다른 2가지 실리콘 N형과 P형을 겹쳐서 만든다  
트랜지스터는 크게 증폭과 스위치, 2가지 기능을 제공한다  
스위치는 전류가 들어오는 1, 들어오지 않는 0으로 구분하며 이 때 이 숫자값을 나타내는 공간을 비트라고 한다  
트랜지스터 1개로 0,1을 표현할 수 있고 2진수로 한자리를 차지한다  
트랜지스터로 계산 기능을 만들기 위해서는 논리 소자가 필요하다  
논리조사는 AND, OR, XOR, NOT (NAND,NOR)이 있다  
#### 1비트 가산기로 더하기 표현하기
| 입력A  | 입력B  |  결과  |
|:----:|:----:|:----:|
|  0   |  0   |  00  |
|  0   |  1   |  01  |
|  1   |  0   |  01  |
|  1   |  1   |  10  |  
  
이러한 1비트 가산기를 응용해 2비트 가산기를 만들고 XX비트까지 계산한다  
위와 동일하게 1비트 감산기를 만들거나 1비트 가산기를 여러번 사용해 곱셈을, 1비트 감산기를 여러번 사용해 나눗셈을 계산한다

#### 프로그램을 실행하면 컴퓨터는 4단계 일을 수행한다
1. 프로그램 로드
2. 데이터 로드 및 캐싱
3. 연산 및 저장
4. 프로그램 종료까지 2와 3을 반복  
  
#### 프로그램 로드
프로그램 실행 > 운영체제가 프로그램 실행 파일을 메모리에 복사 : 로드(load)
#### 데이터 로드 및 캐싱
CPU가 연산을 처리하려면 연산에 필요한 데이터를 가져와야 하는데 이를 위해 CPU내부에 캐시(cache) 메모리 공간을 가지고 있다  
메모리에서 연산에 필요한 데이터와 근처 데이터(참조지역성 locality of reference)를 캐시로 복사한다
#### 연산 및 저장
필요한 데이터를 캐시에 저장 완료 > CPU가 연산에 사용할 데이터를 register로 복사  
레지스터는 실제 연산이 수행되는 특수한 데이터 공간  

# 프로그래밍 언어

## 초창기 프로그래밍 언어
컴퓨터는 0,1밖에 몰?루기 때문에 기계어로 스위치 조절하는 역할을 했음  
이렇게 수행할 명령어를 나타내는 부호를 OP(Operation code)라고 함  
### ADD 3 4를 OP코드로 표현할 경우
|0|0|1|1|
|---|---|---|---|
ADD  
  
|0|0|1|1|
|---|---|---|---|
3  
  
|0|1|0|0|
|---|---|---|---|
4  

(0011은 OP코드로 ADD를 의미함)  

## 어셈블리어
기계어와 1:1로 대응되는 언어로 인간이 쉽게 읽고 쓸 수 있음  
기계어보다는 쉽지만 칩셋마다 명령어를 새로 익혀야함  

## 고수준 언어
어셈블리어는 코드와 기계어가 1:1로 매칭되기 떄문에 변환 과정이 쉬우나 고수준 언어는 기계어로 변환하기 위해 컴파일이 필요함  

## 프로그래밍 언어 구분

### 정적 컴파일 언어
미리 기계어로 변환해두었다가(실행파일) 사용하는 방식의 언어  
실행 시 변환과정이 필요 없어서 빠르고, 타입 에러를 컴파일 시점에서 발견할 수 있어 타입 안정성이 뛰어남  
칩셋과 운영체제마다 0과 1로 된 바이너리 코드를 표현하는 형식이 달라 칩셋에 맞게 변환해줘야 함  
실행 환경(CPU나 OS)의 다양성을 지원하려면 그만큼 많은 빌드 과정이 필요  

### 동적 컴파일 언어
실행 시점(runtime)에 기계어로 변환하는 방식의 언어  
하나의 코드로 모든 플랫폼에서 실행이 가능(범용성)  
  
GO는 정적 컴파일 언어지만 내부 환경 변수만 바꿔서 다양한 플랫폼에 맞도록 컴파일 할 수 있어 비교적 쉽게 범용성을 가질 수 있음  

### 약 타입 언어
타입 검사를 약하게 하는 언어, 동적 타입 언어라고도 함  
서로 다른 타입 간 연산에 관대함  
편하지만 예기치 못한 버그가 발생할 수 있음  

### 강 타입 언어
타입 검사를 강하게 하는 언어, 정적 타입 언어라고도 함  
서로 다른 타입 간 연산에 엄격함  
까다롭지만 타입으로 생길 수 있는 문제를 미연에 방지 가능  
  
GO는 강 타입 언어에서 지원하는 자동 타입 변환까지도 지원하지 않는 최강 타입 언어이다(아잇 싯팔)  

### 가비지 컬렉터
메모리에서 불필요한 영역을 치워주는 역할  
없는 언어는 개발자가 메모리 할당과 해제를 직접 해야하고, 할당한 메모리를 해제하지 않으면 메모리 누수가 나거나 이미 해제된 영역을 다시 해제해 버그가 발생하기도 함  
있는 언어는 메모리를 자동으로 해제하는 편리함은 있지만 CPU성능을 끌어다 쓰기 때문에 성능상의 이슈가 있음  
  
GO는 가비지 컬렉터를 사용하지만 최신 GC라서 성능이 가비지 컬렉터 언어중에 ㅅㅌㅊ이다(GC쓰는게 얼마나 있다고.. 사기같은데..)  

# GO에 대해

## 역사
대충 2009년에 나왔고 업데이트 계속하고 있다.  
온라인 GO 컴파일러인 play.golang.org가 있다.  

## 특징
|개념|있다/없다|설명|
|---|---|---|
|클래스|없다|클래스는 없지만, 메서드를 가지는 구조체를 지원|
|상속|없다|상속을 지원하지 않음|
|메서드|있다|구조체가 메서드를 가질 수 있음|
|인터페이스|있다|상속이 없지만 인터페이스는 있음|
|익명 함수|있다|함수 리터럴이라는 이름으로 제공|
|가비지 컬렉터|있다|고성능? 가비지 컬렉터를 제공|
|포인터|있다|메모리 주소를 가리키는 포인터가 존재|
|제네릭 프로그래밍|없다|제네릭 프로그래밍을 지원하지 않음|
|네임스페이스|없다|네임스페이스를 제공하지 않음. 모든 코드는 패키지 단위로 분리|
  
## 코드 실행 단계
1. 폴더 생성
2. .go 파일 생성 및 작성
3. Go 모듈 생성
4. 빌드
5. 실행
  
### 폴더 생성
Go 언어에서 모든 코드는 패키지 단위로 작성되며 같은 폴더에 위치한 .go 파일은 모두 같은 패키지에 포함되고, 패키지명으로 폴더명을 사용

### .go 파일 생성 및 작성
확장자는 반드시 .go로 끝나야 한다

### Go 모듈 생성
Go 1.16 버전 이후로 Go 모듈이 기본으로 적용되기에 모든 Go 코드는 빌드하기 전 모듈을 생성해야함  
모듈 생성은 `go mod init`명령으로 실행하며 그 결과로 `go.mod` 파일이 생성된다  

### 빌드
`go build` 명령은 Go 코드를 기계어로 변환하여 실행 파일을 생성  
ex) AMD64 계열 칩셋을 사용하는 리눅스 실행파일 생성 시 `GOOS=linux GOARCH=amd64 go build`  

### 실행
빌드로 생성된 실행 파일을 명령어로 실행