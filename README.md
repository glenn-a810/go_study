# GOPATH와 GOROOT

## GOPATH
과거 GO는 워크스페이스를 지정하고 해당 워크스페이스 하위에 코드가 존재해야 했는데, 이 워크스페이스가 GOPATH이고 기본 경로는 C:\Users\<username>\go\로 되어있다.
코드는 GOPATH\src\에 위치해야 했는데, 이런 제약사항을 자유롭게 해주기 위해서 GO MODULE이 생겼다. 보통 GOPATH에는 GO 소스코드 및 외부 패키지가 저장된다고 한다.

## GOROOT
GOROOT는 GO의 컴파일러와 기본 모듈들이 있는 곳으로 GO가 설치되어있는 경로이다.